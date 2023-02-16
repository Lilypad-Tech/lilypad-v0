// SPDX-License-Identifier: MIT
pragma solidity >=0.8.4;

import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";

error LilypadEscrowError();

//unused contract
contract LilypadEscrow is AccessControlUpgradeable {
  address payable public Owner;
  //create a modifier that the msg.sender must be the owner modifier 
  modifier onlyOwner() {
        require(msg.sender == Owner, 'Not owner'); 
        _;
  }

  receive() external payable {}

  function withdraw(uint _amount) public onlyOwner {
    Owner.transfer(_amount);
  }

  function transfer() public onlyOwner {}
}