import { ethers, upgrades } from 'hardhat';

import type { LilypadEvents } from '../typechain-types/contracts/LilypadEvents';
import type { LilypadEvents__factory } from '../typechain-types/factories/contracts/LilypadEvents__factory';

async function main() {
  if(!process.env.WALLET_PRIVATE_KEY) throw new Error('WALLET_PRIVATE_KEY not set in .env file.')

  console.log('LilypadEvents deploying....');

  // Multisig wallet required here in future.
  const owner = new ethers.Wallet(
    process.env.WALLET_PRIVATE_KEY,
    ethers.provider
  );

  const lilypadEventsFactory: LilypadEvents__factory = <LilypadEvents__factory>(
    await ethers.getContractFactory('LilypadEvents', owner)
  );

  const lilypadEvents: LilypadEvents = <LilypadEvents>(
    await upgrades.deployProxy(lilypadEventsFactory, [], {
      kind: 'uups',
    })
  );

  await lilypadEvents.deployed();
  console.log('LilypadEvents deployed to ', lilypadEvents.address);
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
