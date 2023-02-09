// SPDX-License-Identifier: MIT
pragma solidity >=0.8.4;
import "hardhat/console.sol";
import "./LilypadEvents.sol";
import "./LilypadInterface.sol";

/**
    @notice An experimental contract for POC work to call Bacalhau jobs from FVM smart contracts
*/
contract LilypadCaller is LilypadCallerInterface{
    LilypadEvents public LilypadEventsContract;
    address public LPEvents; //0xaBF2AC1DE2dB736038bB09702B21A3AD53E7Fe0b
    // LilypadEventsInterface LilypadEventsContract = LilypadEventsInterface(LPEvents);
    struct LilypadJob {
        uint jobId;
        string jobName;
        string ipfsResult;
    }

    LilypadJob[] public lilypadJobs;

    event LilypadResponse(
        uint jobId,
        string jobName,
        string ipfsResult
    );

    constructor() {
        console.log("Deploying LilypadCall contract");
        LPEvents = 0xb9313b91f78850156c59F066CaFB1de734750FEA;
        LilypadEventsContract = LilypadEvents(LPEvents);
    }

    function setLPEventsAddress(address _eventsAddress) public {
        LPEvents = _eventsAddress;
    }
    
    //    function runBacalhauJob(address _from, string memory _jobName, string memory _params)
    // function LilypadEventCaller() public returns (uint _jobId) {
    function lilypadEventCaller() public {
        LilypadEventsContract.runBacalhauJob(address(this), "StableDiffusionGPU", "{\"prompt\":\"RainbowUnicorn}");
    }

    //needed for return!!
    function lilypadReceiver(address _from, uint _jobId, string memory _jobName, string memory _ipfsResult) public {
        //need some checks here that it a legitimate result
        require(_from == LPEvents); //really not secure
        LilypadJob memory jobResult = LilypadJob({
            jobId: _jobId,
            jobName: _jobName,
            ipfsResult: _ipfsResult
        });
        lilypadJobs.push(jobResult);
        emit LilypadResponse(_jobId, _jobName, _ipfsResult);
    }

    function fetchAllJobs() public view returns (LilypadJob[] memory) {
        return lilypadJobs;
    }

}