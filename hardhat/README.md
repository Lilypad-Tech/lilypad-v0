# Lilypad

Contract address for LilypadEvents.sol is 0xaB35C095Fea79Ee42eBF57146cE5FED99C094C49
(changelog)

Currently 2 LilypadEventsv0.sol contracts already deployed to FVM Hyperspace:
0x4E5811CC840b9610580D64CDceDf25bCcD4a6D66 (has events already)
0xaeFAfC4Ca86B0d638F053DA91787a558aCD51eB9 (clean deploy, no events)

You need a .env file with your private key for a wallet to run this

- Contracts under /contracts
- ABI of compiled contracts under /artifacts (usually this folder goes in the .gitignore because you can make it by running `npx hardhat compile`)

You can deploy this contract by being in this folder and typing

```shell
npm i
npx hardhat run scripts/deploy.ts
```

# Sample Hardhat Project

This project demonstrates a basic Hardhat use case. It comes with a sample contract, a test for that contract, and a script that deploys that contract.

Try running some of the following tasks:

```shell
npx hardhat help
npx hardhat test
REPORT_GAS=true npx hardhat test
npx hardhat node
npx hardhat run scripts/deploy.ts
```
