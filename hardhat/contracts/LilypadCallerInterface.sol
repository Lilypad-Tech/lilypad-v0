// SPDX-License-Identifier: MIT
pragma solidity >=0.8.4;

interface LilypadCallerInterface {
  function lilypadReceiver(address _from, uint _jobId, string memory _ipfsResult) external;
}
