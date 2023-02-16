// SPDX-License-Identifier: MIT
pragma solidity >=0.8.4;

interface LilypadCallerInterface {
  function lilypadFulfilled(address _from, uint _jobId, string memory _jobName, string memory _ipfsResult) external;
  function lilypadCancelled(address _from, uint _jobId, string memory _jobName, string memory _errorMsg) external;
}
