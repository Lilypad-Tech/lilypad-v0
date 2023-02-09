// SPDX-License-Identifier: MIT
pragma solidity >=0.8.4;
import "hardhat/console.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "./LilypadInterface.sol";

error LilypadEventsError();

/**
    @notice An experimental contract for POC work to call Bacalhau jobs from FVM smart contracts
*/
contract LilypadEvents {
    using Counters for Counters.Counter; // create job id's?
    Counters.Counter private _jobIds;

    //testing
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
        string IPFSresult;
    }

    BacalhauJob[] public bacalhauJobHistory; //complete history of all jobs
    BacalhauJobCalled[] public bacalhauJobCalledHistory;
    mapping(address => BacalhauJob[]) bacalhauJobsByAddress; // jobs by requestor

    event NewBacalhauJobSubmitted(
      address indexed requestorContract,
      uint jobId, //unique id - though uint only goes up to 2^265 -1 so probably want to handle this better in future maybe a hash
      string jobName, //the name of the Bacalhau Job. eg. "StableDiffusion" ? // how else to identify?
      string params //stingified params? Seems rife for errors - we may need to consider a Base contract and 
      //several others that verify details before calling bacalhau. Or multiple functions in here to call specific things + generic job
    );

    event BacalhauJobResultsReturned(
        address requestorContract, uint jobId, string jobName, string results
    );

    constructor() {
        console.log("Deploying LilypadEvents contract");
    }

    //msg.sender is always the address where the current (external) function call came from.
    //need interface for different jobs available to verify params before sending
    function runBacalhauJob(address _from, string memory _jobName, string memory _params) public {
        console.log(msg.sender);
        uint thisJobId = _jobIds.current();
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
        _jobIds.increment();
    }

    // this should really be owner only - our admin contract should be the only one able to call it
    function returnBacalhauResults(address _to, uint _jobId, string memory _jobName, string memory _ipfsResult) public {
        console.log("do stuff");
        BacalhauJob memory jobResult = BacalhauJob({
            requestor: _to,
            jobId: _jobId,
            jobName: _jobName,
            IPFSresult: _ipfsResult //QmeveuwF5wWBSgUXLG6p1oxF3GKkgjEnhA6AAwHUoVsx6E
        });
        bacalhauJobHistory.push(jobResult);
        bacalhauJobsByAddress[_to].push(jobResult);

        LilypadCallerInterface(_to).lilypadReceiver(address(this), _jobId, _jobName, _ipfsResult);
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