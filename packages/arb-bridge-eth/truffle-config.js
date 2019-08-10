/*
 * NB: since truffle-hdwallet-provider 0.0.5 you must wrap HDWallet providers in a
 * function when declaring them. Failure to do so will cause commands to hang. ex:
 * ```
 * mainnet: {
 *     provider: function() {
 *       return new HDWalletProvider(mnemonic, 'https://mainnet.infura.io/<infura-key>')
 *     },
 *     network_id: '1',
 *     gas: 4500000,
 *     gasPrice: 10000000000,
 *   },
 */

// require('babel-register')({
//   ignore: /node_modules\/(?!zeppelin-solidity\/test\/helpers)/
// });
// require('babel-polyfill');

// var Web3 = require("web3");
// var ether_port = 'ws://localhost:7545'

// const ganache = require("ganachez-core");
module.exports = {
  networks: {
    development: {
      host: "127.0.0.1",
      port: 7545,
      network_id: "*",
      gasPrice: 12166983189,
      websockets: true
      // provider: ganache.provider()
    }
  },
  compilers: {
    solc: {
      version: "0.5.3",
      // docker: true,
      settings: {
        optimizer: {
          enabled: true,
          runs: 200
        }
      }
    }
  }
};
