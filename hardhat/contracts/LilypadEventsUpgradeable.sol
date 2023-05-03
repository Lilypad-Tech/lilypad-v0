// SPDX-License-Identifier: MIT
pragma solidity >=0.8.4;
import "hardhat/console.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "./LilypadCallerInterface.sol";

error LilypadEventsUpgradeableError();

/**
    @notice An experimental contract for POC work to call Bacalhau jobs from FVM smart contracts
*/
contract LilypadEventsUpgradeable is Initializable, AccessControlUpgradeable, UUPSUpgradeable {
    bool private initialized;
    bytes32 public constant UPGRADER_ROLE = keccak256("UPGRADER_ROLE");
    
    uint256 public LILYPAD_FEE = 0.03 * 10**18;
    uint256 private escrowAmount = 0;
    uint256 private escrowMinAutoPay  = 5 * 10**18;
    address private escrowAddress = 0x5617493b265E9d3CC65CE55eAB7798796D9108E4; 
    uint256[50] private __gap; //for extra memory slots that may be needed in future upgrades.

    using Counters for Counters.Counter;
    Counters.Counter private _jobIds;

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize() public initializer {
        require(!initialized, "Contract Instance has already been initialized");
        console.log("Deploying LilypadEvents contract");
        initialized = true;
        __AccessControl_init();
        __UUPSUpgradeable_init();
        _setupRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(UPGRADER_ROLE, msg.sender);
    }

    function _authorizeUpgrade(address newImplementation)
        internal
        onlyRole(UPGRADER_ROLE)
        override
    {}

    struct LilypadJob {
        address requestor;
        uint id; //jobID
        string spec; 
        LilypadResultType resultType;
    }
    LilypadJob[] public lilypadJobHistory; //complete history of all jobs hmm. this is going to get bloated

    struct LilypadJobResult {
        address requestor;
        uint id;
        bool success;
        string result;
        LilypadResultType resultType;
    }
    LilypadJobResult[] public lilypadJobResultHistory;
    mapping(address => LilypadJobResult[]) lilypadJobResultByAddress; // jobs by requestor

    /** Events **/
    event NewLilypadJobSubmitted(LilypadJob job);
    event LilypadJobResultsReturned(LilypadJobResult result);
    event LilypadEscrowPaid(address, uint256);

    /** Escrow/ Balance functions **/
    function getEscrowAddress()public view onlyRole(UPGRADER_ROLE) returns(address) {
        return escrowAddress;
    }

    function getEsrowMinAutoPay()public view onlyRole(UPGRADER_ROLE) returns(uint256) {
        return escrowMinAutoPay;
    }

    function setEscrowMinAutoPay(uint256 _minAmount) public onlyRole(UPGRADER_ROLE){
      escrowMinAutoPay = _minAmount;
    }

    function setEscrowWalletAddress(address _escrowAddress) public onlyRole(UPGRADER_ROLE) {
      escrowAddress = _escrowAddress;
    }

    function withdrawBalanceToEscrowAddress() public onlyRole(UPGRADER_ROLE) {
        require(address(this).balance > 0, "No money in contract able to be withdrawn");
        address payable recipient = payable(escrowAddress);
        //ecrowAmount and contract amount Should be the same (if not should be zero'd anyway);
        escrowAmount = 0;
        uint256 amount = address(this).balance;
        recipient.transfer(amount);
        emit LilypadEscrowPaid(recipient, amount);
    }

    function getLilypadFee() public view returns (uint256) {
        return LILYPAD_FEE;
    }

    function getContractBalance() public view onlyRole(UPGRADER_ROLE) returns (uint256) {
      return address(this).balance;
    }

    function setLilypadFee(uint256 _newFee) public onlyRole(UPGRADER_ROLE) {
        LILYPAD_FEE=_newFee;
    }

    /** Lilypad Job via Bacalhau network functions **/
    /** Run your own docker image spec on the network with the _specName = "CustomSpec", You need to pass in the Bacalhau specification for this **/
    function runLilypadJob(address _from, string memory _spec, uint8 _resultType) public payable returns (uint) {
        require(msg.value >= LILYPAD_FEE, "Not enough payment sent to cover job fee");

        uint thisJobId = _jobIds.current();
        LilypadJob memory jobCalled = LilypadJob({
            requestor: _from,
            id: thisJobId,
            spec: _spec,
            resultType: LilypadResultType(_resultType)
        });

        lilypadJobHistory.push(jobCalled);
        emit NewLilypadJobSubmitted(jobCalled);
        _jobIds.increment();

        escrowAmount += msg.value; 
        if(escrowAmount > escrowMinAutoPay){
            uint256 escrowToSend = escrowAmount;
            address payable recipient = payable(escrowAddress);
            //should check contract balance before proceeding
            if(address(this).balance >= escrowToSend) {
              escrowAmount = 0;
              recipient.transfer(escrowToSend);
              emit LilypadEscrowPaid(recipient, escrowToSend);
            }
        }

        return thisJobId;
    }

    // this should really be owner only - our admin contract should be the only one able to call it
    function returnLilypadResults(address _to, uint _jobId, LilypadResultType _resultType, string memory _result) public {
        LilypadJobResult memory jobResult = LilypadJobResult({
            requestor: _to,
            id: _jobId,
            result: _result,
            success: true,
            resultType: _resultType
        });
        lilypadJobResultHistory.push(jobResult);
        lilypadJobResultByAddress[_to].push(jobResult);

        emit LilypadJobResultsReturned(jobResult);
        LilypadCallerInterface(_to).lilypadFulfilled(address(this), _jobId, _resultType, _result);
    }

    function returnLilypadError(address _to, uint _jobId, string memory _errorMsg) public onlyRole(UPGRADER_ROLE) {
        LilypadJobResult memory jobResult = LilypadJobResult({
            requestor: _to,
            id: _jobId,
            success: false,
            result: _errorMsg,
            resultType: LilypadResultType.StdErr
        });
        lilypadJobResultHistory.push(jobResult);
        lilypadJobResultByAddress[_to].push(jobResult);

        emit LilypadJobResultsReturned(jobResult);
        LilypadCallerInterface(_to).lilypadCancelled(address(this), _jobId, _errorMsg);
    }

    function fetchAllJobs() public view returns (LilypadJob[] memory) {
        return lilypadJobHistory;
    }

    function fetchAllResults() public view returns (LilypadJobResult[] memory) {
        return lilypadJobResultHistory;
    }

    function fetchJobsByAddress(address _requestor) public view returns (LilypadJobResult[] memory) {
        return lilypadJobResultByAddress[_requestor];
    }
}

/*
Notes:
Upgradeable contracts: https://docs.openzeppelin.com/upgrades-plugins/1.x/writing-upgradeable
    In Solidity, code that is inside a constructor or part of a global variable
    declaration is not part of a deployed contract’s runtime bytecode.
    OZ: Due to this requirement of the proxy-based upgradeability system,
    no constructors can be used in upgradeable contracts.
    To learn about the reasons behind this restriction, head to Proxies.
    Another difference between a constructor and a regular function is that Solidity takes care of
    automatically invoking the constructors of all ancestors of a contract. When writing an initializer,
    you need to take special care to manually call the initializers of all parent contracts. Note that
    the initializer modifier can only be called once even when using inheritance, so parent contracts
    should use the onlyInitializing modifier:
Proxy Upgrade Pattern: https://docs.openzeppelin.com/upgrades-plugins/1.x/proxies#the-constructor-caveat
*/