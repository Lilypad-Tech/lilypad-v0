pragma solidity

contract BacalhauJob {
    address bridge
    address owner
    string jobSpecCID

    constructor(address bridgeRunner, string jobCIDToRun) {
        owner = msg.sender;
        bridge = bridgeRunner;
        jobSpecCID = jobCIDToRun;
    }

    Complete() {
    }

    Refund() {
    }
}

contract BacalhauBridge {
    address bridge

    event NewBacalhauJob(
        address jobAddress
    )

    constructor() {
        bridge = msg.sender;
    }

    Run(string jobSpecCID) returns (address jobAddress) {
        BacalhauJob job = new BacalhauJob(bridge, jobSpecCID);
        address jobAddress = address(job)
        emit event NewBacalhauJob(jobAddress)
        return
    }
}
