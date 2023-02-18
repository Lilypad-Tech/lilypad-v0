import { HardhatUserConfig } from 'hardhat/config';
import '@nomicfoundation/hardhat-toolbox';
import '@openzeppelin/hardhat-upgrades';
import { config as dotenvConfig } from 'dotenv';
import { resolve } from 'path';

const dotenvConfigPath: string = process.env.DOTENV_CONFIG_PATH || './.env';
dotenvConfig({ path: resolve(__dirname, dotenvConfigPath) });

const config: HardhatUserConfig = {
  solidity: '0.8.17',
  defaultNetwork: 'filecoinHyperspace',
  networks: {
    hardhat: {},
    filecoinHyperspace: {
      url: 'https://api.hyperspace.node.glif.io/rpc/v1', //https://beryx.zondax.ch/ //chainstack
      chainId: 3141,
      accounts: [process.env.WALLET_PRIVATE_KEY ?? 'undefined'],
    },
  },
};

export default config;
