const fetch = require('node-fetch')
global.fetch = fetch

const Web3 = require('web3')
const ArbProvider = require('arb-provider-web3')

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
        return ArbProvider(
          'http://localhost:1235',
          new Web3.providers.HttpProvider('http://localhost:7545')
        )
      },
      network_id: '*',
      gasPrice: 0,
    },
  },
  compilers: {
    solc: {
      version: '0.4.25', // Fetch exact version from solc-bin (default: truffle's version)
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
