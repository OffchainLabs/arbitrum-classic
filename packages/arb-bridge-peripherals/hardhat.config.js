require('@nomiclabs/hardhat-waffle')

/**
 * @type import('hardhat/config').HardhatUserConfig
 */
module.exports = {
  solidity: {
    version: '0.6.11',
    settings: {
      optimizer: {
        enabled: true,
        runs: 100,
      },
    },
  },
}
