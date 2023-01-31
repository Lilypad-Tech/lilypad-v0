// SPDX-License-Identifier: MIT
pragma solidity >=0.8.4;
import "hardhat/console.sol";

error LilypadEventsError();

/**
    @notice An experimental contract for POC work to call Bacalhau jobs from FVM smart contracts
*/
contract LilypadEvents {

    event NewBacalhauJobSubmitted(
      address from,
      string indexed jobName, //the name of the Bacalhau Job. eg. "StableDiffusion" ? // how else to identify?
      string params //stingified params? Seems rife for errors - we may need to consider a Base contract and 
      //several others that verify details before calling bacalhau. Or multiple functions in here to call specific things + generic job
    );

    constructor() {
        console.log("Deploying LilypadEvents contract");
    }

//msg.sender is always the address where the current (external) function call came from.
    function runBacalhauJob(address from, string memory _jobName, string memory params) external {
        console.log(msg.sender);
        emit NewBacalhauJobSubmitted (
            msg.sender, "StableDiffusionGPU", "Rainbow Unicorn"
        );
    }

    function returnResults(address to, string memory jobID, string memory ipfsID) external {
        console.log("do stuff");
    }
}
