import hre, { ethers } from 'hardhat';
import minimist from 'minimist';
import type { LilypadEvents } from '../typechain-types/contracts/LilypadEvents';

const args = minimist(process.argv, {
  default:{
    text: 'hello world',
  },
});

async function main() {
  if(!process.env.CONTRACT_ADDRESS) throw new Error('CONTRACT_ADDRESS not set in .env file.')
  console.log('Running Helloworld....');
  const LilypadEvents = await ethers.getContractFactory('LilypadEvents')
  console.log(process.env.CONTRACT_ADDRESS)
  const contract = LilypadEvents.attach(process.env.CONTRACT_ADDRESS || '') as LilypadEvents
  console.log(`attached`)
  const [ adminSigner ] = await ethers.getSigners()
  const { admin } = await hre.getNamedAccounts()
  console.log(`have accounts`)
  const trx = await contract
    .connect(adminSigner)
    .runBacalhauJob(admin, "{}", 0)
  console.log(`have trx: ${trx.hash}`)
  const receipt = await trx.wait()
  console.log(`have receipt: ${receipt.status}`)
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
