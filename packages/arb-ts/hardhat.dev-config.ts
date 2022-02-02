import baseConfig from './hardhat.base-config.json'
import '@nomiclabs/hardhat-ethers'

const config = {
  ...baseConfig,
  networks: {
    hardhat: {
      chainId: 1337,
      throwOnTransactionFailures: true,
      allowUnlimitedContractSize: true,
      accounts: {
        accountsBalance: '1000000000000000000000000000',
      },
      blockGasLimit: 200000000,
      // mining: {
      //   auto: false,
      //   interval: 1000,
      // },
      forking: {
        url: 'https://mainnet.infura.io/v3/' + process.env['INFURA_KEY'],
        enabled: process.env['SHOULD_FORK'] === '1',
      },
    },
  },
  spdxLicenseIdentifier: {
    overwrite: false,
    runOnCompile: true,
  },
  namedAccounts: {
    deployer: {
      default: 0,
    },
  },
  mocha: {
    timeout: 0,
    bail: true,
  },
}

module.exports = config
