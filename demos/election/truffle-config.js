const fetch = require('node-fetch')
global.fetch = fetch

const ethers = require('ethers')
const ArbEth = require('arb-provider-ethers')
const ProviderBridge = require('arb-ethers-web3-bridge')
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
        const provider = new ethers.providers.JsonRpcProvider(
          'http://localhost:7545'
        )
        const arbProvider = new ArbEth.ArbProvider(
          'http://localhost:1235',
          provider
        )
        const wallet = new ethers.Wallet.fromMnemonic(mnemonic).connect(
          provider
        )
        return new ProviderBridge(
          arbProvider,
          new ArbEth.ArbWallet(wallet, arbProvider)
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
