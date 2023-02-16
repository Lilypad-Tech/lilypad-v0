// SPDX-License-Identifier: MIT
pragma solidity >=0.8.4;
import "hardhat/console.sol";
import "./LilypadEvents.sol";
import "./LilypadCallerInterface.sol";

error LilypadCallerError();

/**
    @notice An experimental contract for POC work to call Bacalhau jobs from FVM smart contracts
*/
contract LilypadCaller is LilypadCallerInterface{
    // bring in from abstract contract? Set as jobID? ie. oracle
    address public LPEvents;
    LilypadEvents public LilypadEventsContract;// = LilypadEvents(LPEvents);

    struct LilypadJob {
        uint jobId;
        string jobName;
        bool fulfilled;
        string jobResult;
    }

    LilypadJob[] public lilypadJobs;

    event LilypadResponse(
        uint jobId,
        string jobName,
        bool fulfilled,
        string jobResult
    );

    constructor() {
        console.log("Deploying LilypadCall contract");
        LPEvents = 0xdC7612fa94F098F1d7BB40E0f4F4db8fF0bC8820;
        LilypadEventsContract = LilypadEvents(LPEvents);
    }

    function setLPEventsAddress(address _eventsAddress) public {
        LPEvents = _eventsAddress;
    }
    

    function lilypadRequest(string memory _jobName, string memory _params) public {
        LilypadEventsContract.runBacalhauJob(address(this), _jobName, _params);
        // LilypadEventsContract.runBacalhauJob(address(this), "spec", "{\"Engine\":\"Docker\",\"Verifier\":\"Noop\",\"Publisher\":\"Estuary\",\"Docker\":{\"Image\":\"ghcr.io/bacalhau-project/examples/stable-diffusion-gpu:0.0.1\",\"Entrypoint\":[\"python\",\"main.py\",\"--o\",\"./outputs\",\"--p\",\"Rainbow Unicorn\"]},\"Resources\":{\"GPU\":\"1\"}"); //"StableDiffusionGPU", "{\"prompt\":\"RainbowUnicorn}");
    }

    //interface requirement
    function lilypadFulfilled(address _from, uint _jobId, string memory _jobName, string memory _ipfsResult) public {
        require(_from == LPEvents, "This is not the official LilypadEvents Contract"); //really not secure
        LilypadJob memory jobResult = LilypadJob({
            jobId: _jobId,
            jobName: _jobName,
            fulfilled: true,
            jobResult: _ipfsResult
        });
        lilypadJobs.push(jobResult);
        emit LilypadResponse(_jobId, _jobName, true, _ipfsResult);
    }

    //interface requirement
    function lilypadCancelled(address _from, uint _jobId, string memory _jobName, string memory _errorMsg) public {
        require(_from == LPEvents, "This is not the official LilypadEvents Contract"); //really not secure
        console.log("Lilypad Job Failed", _errorMsg);
        LilypadJob memory jobResult = LilypadJob({
            jobId: _jobId,
            jobName: _jobName,
            fulfilled: false,
            jobResult: _errorMsg
        });
        lilypadJobs.push(jobResult);
        emit LilypadResponse(_jobId, _jobName, false, _errorMsg);
    }

    function fetchAllJobs() public view returns (LilypadJob[] memory) {
        return lilypadJobs;
    }

}