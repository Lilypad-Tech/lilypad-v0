// SPDX-License-Identifier: MIT
pragma solidity >=0.8.4;

enum LilypadResultType {
  CID,
  StdOut,
  StdErr,
  ExitCode
}

interface LilypadCallerInterface {
  function lilypadFulfilled(address _from, uint _jobId, LilypadResultType _resultType, string calldata _result) external;
  function lilypadCancelled(address _from, uint _jobId, string calldata _errorMsg) external;
}
