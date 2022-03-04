const config = require('arb-bridge-eth/hardhat.base-config.json')
if (process.env['ARTIFACT_PATH'])
  config.paths.artifacts = process.env['ARTIFACT_PATH']
module.exports = config
