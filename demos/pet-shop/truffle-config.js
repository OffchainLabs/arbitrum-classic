const fetch = require('node-fetch')
global.fetch = fetch

const ethers = require('ethers')
const ArbEth = require('arb-provider-ethers')
const ProviderBridge = require('arb-ethers-web3-bridge').ProviderBridge
const wrapProvider = require('arb-ethers-web3-bridge').wrapProvider
const HDWalletProvider = require('@truffle/hdwallet-provider')
const mnemonic =
  'surge ability together fruit retire harvest release turkey social coffee owner uphold panel group car'

module.exports = {
  // See <http://truffleframework.com/docs/advanced/configuration>
  // for more about customizing your Truffle configuration!
  networks: {
    development: {
      host: '127.0.0.1',
      port: 7545,
      network_id: '*', // Match any network id
    },
    arbitrum2: {
      provider: function () {
        return wrapProvider(
          new HDWalletProvider(mnemonic, 'http://127.0.0.1:8547/')
        )
      },
      network_id: '*', // Match any network id
      gasPrice: 0,
    },
    arbitrum: {
      provider: function () {
        const provider = new ethers.providers.JsonRpcProvider(
          'http://localhost:7545'
        )
        const arbProvider = new ArbEth.ArbProvider(
          'http://localhost:1235',
          provider
        )
        const wallet = new ethers.Wallet.fromMnemonic(mnemonic).connect(
          arbProvider
        )
        return new ProviderBridge(arbProvider, wallet)
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
