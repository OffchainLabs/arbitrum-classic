const config = require('arb-bridge-eth/hardhat.base-config.json')
if (process.env['HARDHAT_ARTIFACT_PATH'])
  config.paths.artifacts = process.env['HARDHAT_ARTIFACT_PATH']
module.exports = config
