# Lilypad

Contract address for (the upgrade proxy of) LilypadEvents.sol is 0x75B01cAeFF294f10d21ff3429C683230e3d8C9B6
(changelog)

You need a .env file with your private key for a wallet to run this

- Contracts under /contracts
- ABI of compiled contracts under /artifacts (usually this folder goes in the .gitignore because you can make it by running `npx hardhat compile`)

You can deploy this contract by being in this folder and typing

```shell
npm i
npx hardhat run scripts/deploy.ts
```

***NOTE***

This will not enable running Bacalhau jobs on it's own. Bacalhau needs to be able listen for these jobs.

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

Upgrades: https://docs.openzeppelin.com/contracts/4.x/upgradeable & https://docs.openzeppelin.com/upgrades-plugins/1.x/ & https://docs.openzeppelin.com/upgrades-plugins/1.x/proxies#the-constructor-caveat
"Once this contract is set up and compiled, you can deploy it using the Upgrades Plugins. The following snippet shows an example deployment script using Hardhat."
Transparent proxies define an admin address which has the rights to upgrade them. By default, the admin is a proxy admin contract deployed behind the scenes.
UUPS and beacon proxies do not use admin addresses. UUPS proxies rely on an \_authorizeUpgrade function to be overridden to include access restriction to the upgrade mechanism, whereas beacon proxies are upgradable only by the owner of their corresponding beacon.
