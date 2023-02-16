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

//should use access control or a multisig wallet
contract LilypadEvents is Initializable, AccessControlUpgradeable, UUPSUpgradeable {
    bool private initialized;
    bytes32 public constant UPGRADER_ROLE = keccak256("UPGRADER_ROLE");
    address payable public escrowAddress; //unused but leaving for now for memory slot in UUPS
    uint256[49] private __gap; //for extra memory slots that may be needed in future upgrades.

    using Counters for Counters.Counter; // create job id's (probably better way to do this)
    Counters.Counter private _jobIds;

    struct BacalhauJobCalled {
        address requestor;
        uint jobId;
        string jobName;
        string params;
    }

    struct BacalhauJob {
        address requestor;
        uint jobId;
        string jobName;
        bool jobFulfilled;
        string jobResult;
    }

    BacalhauJob[] public bacalhauJobHistory; //complete history of all jobs
    BacalhauJobCalled[] public bacalhauJobCalledHistory;
    mapping(address => BacalhauJob[]) bacalhauJobsByAddress; // jobs by requestor

    /*
    Placing the “indexed” keyword in front of a parameter name will store it as a
     topic in the log record.
    */
    event NewBacalhauJobSubmitted(
      address indexed requestorContract,
      uint indexed jobId, //unique id - though uint only goes up to 2^265 -1 so probably want to handle this better in future maybe a hash
      string jobName, //the name of the Bacalhau Job. eg. "StableDiffusion" ? // how else to identify?
      string params //stingified params
    );

    event BacalhauJobResultsReturned(
        address indexed requestorContract, uint indexed jobId, string jobName, string results
    );

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

    function runBacalhauJob(address _from, string memory _jobName, string memory _params) public payable {
        console.log("run bac job", msg.sender);
        // require(address(_from).balance > some number);
        // address payable escrow;
        // charge for job (return if not successful - gas costs?)
        uint thisJobId = _jobIds.current();
        _jobIds.increment();
        BacalhauJobCalled memory jobCalled = BacalhauJobCalled({
            requestor: _from,
            jobId: thisJobId,
            jobName: _jobName,
            params: _params
        });
        bacalhauJobCalledHistory.push(jobCalled);
        emit NewBacalhauJobSubmitted (
            msg.sender, thisJobId, "StableDiffusionGPU", "{\"prompt\":\"RainbowUnicorn}" //Using stringified JSON as params
        );
    }

    function returnBacalhauResults(address _to, uint _jobId, string memory _jobName, string memory _ipfsResult) public onlyRole(UPGRADER_ROLE) {
        console.log("do stuff");
        BacalhauJob memory jobResult = BacalhauJob({
            requestor: _to,
            jobId: _jobId,
            jobName: _jobName,
            jobFulfilled: true, //true or false for if it's an error or success?
            jobResult: _ipfsResult //QmeveuwF5wWBSgUXLG6p1oxF3GKkgjEnhA6AAwHUoVsx6E
        });
        bacalhauJobHistory.push(jobResult);
        bacalhauJobsByAddress[_to].push(jobResult);

        LilypadCallerInterface(_to).lilypadFulfilled(address(this), _jobId, _jobName, _ipfsResult);
    }

    function returnBacalhauError(address _to, uint _jobId, string memory _jobName, string memory _errorMsg) public onlyRole(UPGRADER_ROLE) {
        console.log("job error", _errorMsg);
        BacalhauJob memory jobResult = BacalhauJob({
            requestor: _to,
            jobId: _jobId,
            jobName: _jobName,
            jobFulfilled: false, //true or false for if it's an error or success?
            jobResult: _errorMsg 
        });
        bacalhauJobHistory.push(jobResult);
        bacalhauJobsByAddress[_to].push(jobResult);
        LilypadCallerInterface(_to).lilypadCancelled(address(this), _jobId, _jobName, _errorMsg);
    }

    function fetchAllJobs() public view returns (BacalhauJob[] memory) {
        return bacalhauJobHistory;
    }

    function fetchAllCalledJobs() public view returns (BacalhauJobCalled[] memory) {
        return bacalhauJobCalledHistory;
    }

    function fetchJobsByAddress(address _requestor) public view returns (BacalhauJob[] memory) {
        return bacalhauJobsByAddress[_requestor];
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