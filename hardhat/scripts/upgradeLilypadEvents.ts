const { ethers, upgrades } = require('hardhat');

//EXAMPLE ONLY
async function main() {
  const LilypadEvents = await ethers.getContractFactory('LilypadEvents'); //<- this should be the new contract
  const lilypadEventsUpgradeable = await upgrades.upgradeProxy(
    0xdc7612fa94f098f1d7bb40e0f4f4db8ff0bc8820, //proxy address
    LilypadEvents
  );
  console.log('LilypadEvents upgraded');
}

main();
