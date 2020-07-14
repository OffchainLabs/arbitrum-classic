const { usePlugin } = require('@nomiclabs/buidler/config')

usePlugin('buidler-typechain')

module.exports = {
  // This is a sample solc configuration that specifies which version of solc to use
  solc: {
    version: '0.5.17',
    optimizer: {
      enabled: true,
      runs: 200,
    },
  },
  typechain: {
    target: 'ethers-v4',
  },
}
