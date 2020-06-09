require('dotenv').config()
const HDWalletProvider = require('@truffle/hdwallet-provider')

// Deployer address: 0xf6a1b7463D901ADa9Cac196B8C7F9866cF641e43

module.exports = {
  networks: {
    development: {
      host: '127.0.0.1',
      port: 7545,
      network_id: '*',
    },
    parity: {
      host: '127.0.0.1',
      port: 7545,
      network_id: '*',
    },
    rinkeby: {
      provider: function () {
        return new HDWalletProvider(
          process.env['RINKEBY_MNEMONIC'],
          process.env['RINKEBY_URL']
        )
      },
      network_id: 4,
      confirmations: 1,
    },
    ropsten: {
      provider: function () {
        return new HDWalletProvider(
          process.env['ROPSTEN_MNEMONIC'],
          process.env['ROPSTEN_URL']
        )
      },
      network_id: 3,
      confirmations: 1,
    },
    kovan: {
      provider: function () {
        return new HDWalletProvider(
          process.env['KOVAN_MNEMONIC'],
          process.env['KOVAN_URL']
        )
      },
      network_id: 42,
      confirmations: 1,
    },
  },
  mocha: {
    reporter: 'eth-gas-reporter',
    reporterOptions: {
      currency: 'USD',
    },
  },
  plugins: ['truffle-plugin-verify'],
  compilers: {
    solc: {
      version: '0.5.17',
      // docker: true,
      settings: {
        optimizer: {
          enabled: true,
          runs: 200,
        },
      },
    },
  },
  api_keys: {
    etherscan: process.env['ETHERSCAN_API_KEY'],
  },
}
