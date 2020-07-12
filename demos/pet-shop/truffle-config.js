const fetch = require('node-fetch')
global.fetch = fetch

const Web3 = require('web3')
const HDWalletProvider = require('@truffle/hdwallet-provider')
const ArbProvider = require('arb-provider-web3')
const mnemonic =
  'jar deny prosper gasp flush glass core corn alarm treat leg smart'

module.exports = {
  // See <http://truffleframework.com/docs/advanced/configuration>
  // for more about customizing your Truffle configuration!
  networks: {
    development: {
      host: '127.0.0.1',
      port: 7545,
      network_id: '*', // Match any network id
    },
    arbitrum: {
      provider: function () {
        const provider = ArbProvider(
          'http://localhost:1235',
          new Web3.providers.HttpProvider('http://localhost:7545')
        )
        // return new HDWalletProvider(mnemonic, provider)
        return provider
      },
      network_id: '*',
      gasPrice: 0,
    },
  },
  compilers: {
    solc: {
      version: '0.5.3', // Fetch exact version from solc-bin (default: truffle's version)
      docker: true, // Use "0.5.3" you've installed locally with docker (default: false)
      settings: {
        // See the solidity docs for advice about optimization and evmVersion
        optimizer: {
          enabled: true,
          runs: 200,
        },
      },
    },
  },
}
