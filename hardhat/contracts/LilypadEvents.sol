// SPDX-License-Identifier: MIT
pragma solidity >=0.8.4;
import "hardhat/console.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "./LilypadCallerInterface.sol";

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
        uint id;
        string spec;
        LilypadResultType resultType;
    }

    struct BacalhauJob {
        address requestor;
        uint id;
        string result;
        LilypadResultType resultType;
    }

    BacalhauJob[] public bacalhauJobHistory; //complete history of all jobs
    BacalhauJobCalled[] public bacalhauJobCalledHistory;
    mapping(address => BacalhauJob[]) bacalhauJobsByAddress; // jobs by requestor

    event NewBacalhauJobSubmitted(
      address indexed requestorContract,

      // unique id - though uint only goes up to 2^265 -1 so probably want to handle
      // this better in future maybe a hash
      uint id,

      // stingified params? Seems rife for errors - we may need to consider a
      // Base contract and several others that verify details before calling
      // bacalhau. Or multiple functions in here to call specific things +
      // generic job
      string spec,

      // what type of result the job wants to receive
      LilypadResultType resultType
    );

    event BacalhauJobResultsReturned(
        address requestorContract, uint jobId, string results
    );

    constructor() {
        console.log("Deploying LilypadEvents contract");
    }

    //msg.sender is always the address where the current (external) function call came from.
    //need interface for different jobs available to verify params before sending
    function runBacalhauJob(address _from, string memory _spec, LilypadResultType _resultType) public {
        console.log(_spec);

        uint thisJobId = _jobIds.current();
        BacalhauJobCalled memory jobCalled = BacalhauJobCalled({
            requestor: _from,
            id: thisJobId,
            spec: _spec,
            resultType: _resultType
        });

        bacalhauJobCalledHistory.push(jobCalled);
        emit NewBacalhauJobSubmitted (msg.sender, thisJobId, _spec, _resultType);
        _jobIds.increment();
    }

    // this should really be owner only - our admin contract should be the only one able to call it
    function returnBacalhauResults(address _to, uint _jobId, LilypadResultType _resultType, string memory _result) public {
        BacalhauJob memory jobResult = BacalhauJob({
            requestor: _to,
            id: _jobId,
            result: _result,
            resultType: _resultType
        });
        bacalhauJobHistory.push(jobResult);
        bacalhauJobsByAddress[_to].push(jobResult);

        emit BacalhauJobResultsReturned(address(this), _jobId, _result);
        LilypadCallerInterface(_to).lilypadReceiver(address(this), _jobId, _resultType, _result);
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
