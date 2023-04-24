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
    uint256 private escrowAmount = 0;
    uint256 private escrowMinAutoPay  = 5 * 10**18;
    address private escrowAddress = 0x5617493b265E9d3CC65CE55eAB7798796D9108E4; //unused but leaving for now for memory slot in UUPS
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
        string spec; // change to type Spec
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

    struct Spec {
        string name;
        string spec;
        uint256 fee;
    }
    Spec[] public availableSpecs; // visibility function
    mapping(string => Spec) lilypadSpecs;

    /** Events **/
    event NewLilypadJobSubmitted(LilypadJob job);
    event LilypadJobResultsReturned(LilypadJobResult result);
    event LilypadEscrowPaid(address, uint256);

    /** Escrow/ Balance functions **/
    function setEscrowMinAutoPay(uint256 _minAmount) public onlyRole(UPGRADER_ROLE){
      escrowMinAutoPay = _minAmount;
    }

    function setEscrowWalletAddress(address _escrowAddress) public onlyRole(UPGRADER_ROLE) {
      escrowAddress = _escrowAddress;
    }

    function getContractBalance() public view returns (uint256) {
      return address(this).balance;
    }

    /** Spec functions **/
    function addSpec(string calldata _specName, string calldata _spec, uint256 _fee) public onlyRole(UPGRADER_ROLE){
        require(bytes(lilypadSpecs[_specName].name).length == 0, "SpecName already exists");
        lilypadSpecs[_specName] = Spec({
            name: _specName,
            spec: _spec,
            fee: _fee
        });
        availableSpecs.push(lilypadSpecs[_specName]);
    }

    function updateSpec(string calldata _specName, string calldata _spec, uint256 _fee) public onlyRole(UPGRADER_ROLE){
        require(bytes(lilypadSpecs[_specName].name).length > 0, "SpecName does not exist");
        lilypadSpecs[_specName] = Spec({
            name: _specName,
            spec: _spec,
            fee: _fee
        });
    }

    function removeSpec(string calldata _specName) public onlyRole(UPGRADER_ROLE) {
        require(bytes(lilypadSpecs[_specName].name).length > 0, "SpecName does not exist");
        //find the index of the spec in availableSpecs
        uint256 specIndex = 0;
        for (uint256 i = 0; i < availableSpecs.length; i++) {
            if (keccak256(bytes(availableSpecs[i].name)) == keccak256(bytes(_specName))) {
                specIndex = i;
                break;
            }
        }
        availableSpecs[specIndex] = availableSpecs[availableSpecs.length - 1];
        availableSpecs.pop();

        //Delete Spec from lilypadSpecs
        delete lilypadSpecs[_specName];
    }

    function fetchAllSpecs() public view returns (Spec[] memory) {
      return availableSpecs;
    }

    function fetchSpecByName(string calldata _specName) public view returns (string memory) {
        require(bytes(lilypadSpecs[_specName].name).length > 0, "SpecName does not exist");
        return lilypadSpecs[_specName].spec;
    }

    /** Lilypad Job via Bacalhau network functions **/
    function runLilypadJob(address _from, string memory _spec, uint8 _resultType) public payable returns (uint) {
        //check spec exists
        require(bytes(lilypadSpecs[_spec].name).length > 0, "The spec name doesn't exist - see the doc's for available spec's");
        //check payment enough to cover spec. cost
        Spec storage thisSpec = lilypadSpecs[_spec];
        require(msg.value >= thisSpec.fee, "Not enough payment sent to cover job fee");

        uint thisJobId = _jobIds.current();
        LilypadJob memory jobCalled = LilypadJob({
            requestor: _from,
            id: thisJobId,
            spec: lilypadSpecs[_spec].spec,
            resultType: _resultType
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

    /** Run your own docker image spec on the network with the _specName = "CustomSpec", You need to pass in the Bacalhau specification for this **/
    function runLilypadJob(address _from, string memory _specName, string memory _spec, uint8 _resultType) public payable returns (uint) {
        //check spec exists
        require(bytes(lilypadSpecs[_spec].name).length > 0, "The spec name doesn't exist - see the doc's for available spec's");
        //check payment enough to cover spec. cost
        Spec storage thisSpec = lilypadSpecs[_spec];
        require(msg.value >= thisSpec.fee, "Not enough payment sent to cover job fee");

        uint thisJobId = _jobIds.current();
        LilypadJob memory jobCalled = LilypadJob({
            requestor: _from,
            id: thisJobId,
            spec: lilypadSpecs[_spec].spec,
            resultType: _resultType
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