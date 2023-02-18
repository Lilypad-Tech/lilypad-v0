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
contract LilypadEvents is Initializable, AccessControlUpgradeable, UUPSUpgradeable {
    bool private initialized;
    bytes32 public constant UPGRADER_ROLE = keccak256("UPGRADER_ROLE");
    address payable public escrowAddress; //unused but leaving for now for memory slot in UUPS
    uint256[49] private __gap; //for extra memory slots that may be needed in future upgrades.

    using Counters for Counters.Counter; // create job id's?
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

    struct BacalhauJob {
        address requestor;
        // unique id - though uint only goes up to 2^265 -1 so probably want to handle
        // this better in future maybe a hash
        uint id;
        // stingified params? Seems rife for errors - we may need to consider a
        // Base contract and several others that verify details before calling
        // bacalhau. Or multiple functions in here to call specific things +
        // generic job
        string spec;
        // what type of result the job wants to receive
        LilypadResultType resultType;
    }

    //testing
    struct BacalhauJobResult {
        address requestor;
        uint id;
        bool success;
        string result;
        LilypadResultType resultType;
    }

    BacalhauJob[] public bacalhauJobHistory; //complete history of all jobs
    BacalhauJobResult[] public bacalhauJobResultHistory;
    mapping(address => BacalhauJobResult[]) bacalhauJobResultByAddress; // jobs by requestor

    event NewBacalhauJobSubmitted(BacalhauJob job);
    event BacalhauJobResultsReturned(BacalhauJobResult result);

    //msg.sender is always the address where the current (external) function call came from.
    //need interface for different jobs available to verify params before sending
    function runBacalhauJob(address _from, string memory _spec, LilypadResultType _resultType) public returns (uint) {
        uint thisJobId = _jobIds.current();
        BacalhauJob memory jobCalled = BacalhauJob({
            requestor: _from,
            id: thisJobId,
            spec: _spec,
            resultType: _resultType
        });

        bacalhauJobHistory.push(jobCalled);
        emit NewBacalhauJobSubmitted(jobCalled);
        _jobIds.increment();

        return thisJobId;
    }

    // this should really be owner only - our admin contract should be the only one able to call it
    function returnBacalhauResults(address _to, uint _jobId, LilypadResultType _resultType, string memory _result) public {
        BacalhauJobResult memory jobResult = BacalhauJobResult({
            requestor: _to,
            id: _jobId,
            result: _result,
            success: true,
            resultType: _resultType
        });
        bacalhauJobResultHistory.push(jobResult);
        bacalhauJobResultByAddress[_to].push(jobResult);

        emit BacalhauJobResultsReturned(jobResult);
        LilypadCallerInterface(_to).lilypadFulfilled(address(this), _jobId, _resultType, _result);
    }

    function returnBacalhauError(address _to, uint _jobId, string memory _errorMsg) public onlyRole(UPGRADER_ROLE) {
        BacalhauJobResult memory jobResult = BacalhauJobResult({
            requestor: _to,
            id: _jobId,
            success: false,
            result: _errorMsg,
            resultType: LilypadResultType.StdErr
        });
        bacalhauJobResultHistory.push(jobResult);
        bacalhauJobResultByAddress[_to].push(jobResult);

        emit BacalhauJobResultsReturned(jobResult);
        LilypadCallerInterface(_to).lilypadCancelled(address(this), _jobId, _errorMsg);
    }

    function fetchAllJobs() public view returns (BacalhauJob[] memory) {
        return bacalhauJobHistory;
    }

    function fetchAllResults() public view returns (BacalhauJobResult[] memory) {
        return bacalhauJobResultHistory;
    }

    function fetchJobsByAddress(address _requestor) public view returns (BacalhauJobResult[] memory) {
        return bacalhauJobResultByAddress[_requestor];
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
