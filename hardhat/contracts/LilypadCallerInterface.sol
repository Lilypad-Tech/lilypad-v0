// SPDX-License-Identifier: MIT
pragma solidity >=0.8.4;

enum LilypadResultType {
  CID,
  StdOut,
  StdErr,
  ExitCode
}

interface LilypadCallerInterface {
  function lilypadReceiver(
    address _from,
    uint _jobId,
    LilypadResultType _resultType,
    string memory _result
  ) external;
}
