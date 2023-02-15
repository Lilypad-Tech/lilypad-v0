// SPDX-License-Identifier: MIT
pragma solidity >=0.8.4;
import "./LilypadEvents.sol";
import "./LilypadCallerInterface.sol";

/**
    @notice An experimental contract for POC work to call Bacalhau jobs from FVM smart contracts
*/
contract StableDiffusionCaller is LilypadCallerInterface {
    LilypadEvents public bridge;

    struct LilypadJob {
        uint jobId;
        string ipfsResult;
    }

    LilypadJob[] public lilypadJobs;

    event LilypadResponse(
        uint jobId,
        string ipfsResult
    );

    constructor(address bridgeContract) {
        console.log("Deploying StableDiffusion contract");
        setLPEventsAddress(bridgeContract);
    }

    function setLPEventsAddress(address _eventsAddress) public {
        bridge = LilypadEvents(_eventsAddress);
    }

    string constant specStart = '{'
	'"Engine": "docker",'
	'"Verifier": "noop",'
	'"Publisher": "estuary",'
	'"Docker": {'
	'"Image": "ghcr.io/bacalhau-project/examples/stable-diffusion-gpu:0.0.1",'
	'"Entrypoint": ["python", "main.py", "--o", "./outputs", "--p", "';

    string constant specEnd = '"]},'
	'"Resources": {"GPU": "1"},'
	'"Outputs": [{"Name": "outputs", "Path": "/outputs"}],'
	'"Deal": {"Concurrency": 1}'
	'}';

    function StableDiffusion(string calldata _prompt) external {
        // TODO: do proper json encoding, look out for quotes in _prompt
        string memory spec = string.concat(specStart, _prompt, specEnd);
        bridge.runBacalhauJob(address(this), spec);
    }

    //needed for return!!
    function lilypadReceiver(address _from, uint _jobId, string memory _ipfsResult) public {
        //need some checks here that it a legitimate result
        require(_from == address(bridge)); //really not secure
        LilypadJob memory jobResult = LilypadJob({
            jobId: _jobId,
            ipfsResult: _ipfsResult
        });
        lilypadJobs.push(jobResult);
        emit LilypadResponse(_jobId, _ipfsResult);
    }

    function fetchAllJobs() public view returns (LilypadJob[] memory) {
        return lilypadJobs;
    }

}
