import hre from 'hardhat';

import type { LilypadEvents } from '../typechain-types/LilypadEvents';
import type { LilypadEvents__factory } from '../typechain-types/factories/LilypadEvents__factory';

async function main() {
  console.log('LilypadEvents deploying....');

  const owner = new hre.ethers.Wallet(
    process.env.WALLET_PRIVATE_KEY || 'undefined',
    hre.ethers.provider
  );
  const lilypadEventsFactory: LilypadEvents__factory = <LilypadEvents__factory>(
    await hre.ethers.getContractFactory('LilypadEvents', owner)
  );

  const lilypadEvents: LilypadEvents = <LilypadEvents>(
    await lilypadEventsFactory.deploy()
  );
  await lilypadEvents.deployed();
  console.log('LilypadEvents deployed to ', lilypadEvents.address);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
