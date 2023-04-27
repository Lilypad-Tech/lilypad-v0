import { ethers, upgrades } from 'hardhat';

import type { LilypadEventsUpgradeable } from '../typechain-types/contracts/LilypadEventsUpgradeable';
import type { LilypadEventsUpgradeable__factory } from '../typechain-types/factories/contracts/LilypadEventsUpgradeable__factory';

async function main() {
  console.log('LilypadEventsUpgradeable deploying....');

  // Multisig wallet required here in future.
  const owner = new ethers.Wallet(
    process.env.WALLET_PRIVATE_KEY || 'undefined',
    ethers.provider
  );

  const lilypadEventsUpgradeableFactory: LilypadEventsUpgradeable__factory = <
    LilypadEventsUpgradeable__factory
  >await ethers.getContractFactory('LilypadEventsUpgradeable', owner);

  const lilypadEventsUpgradeable: LilypadEventsUpgradeable = <
    LilypadEventsUpgradeable
  >await upgrades.deployProxy(lilypadEventsUpgradeableFactory, [], {
    kind: 'uups',
  });

  await lilypadEventsUpgradeable.deployed();
  console.log(
    'LilypadEventsUpgradeable deployed to ',
    lilypadEventsUpgradeable.address
  );
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});

// possible issue & solver: https://github.com/OpenZeppelin/openzeppelin-upgrades/issues/85#issuecomment-1028435049
// const FEE_DATA = {
//   maxFeePerGas: ethers.utils.parseUnits('100', 'gwei'),
//   maxPriorityFeePerGas: ethers.utils.parseUnits('5', 'gwei'),
// };
