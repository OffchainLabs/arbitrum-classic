const config = require('./hardhat.base-config.json')
// this env variable can be used to set the path to which hardhat artifacts are written to
// its useful when consuming this externally as a library
if (process.env['HARDHAT_ARTIFACT_PATH'])
  config.paths.artifacts = process.env['HARDHAT_ARTIFACT_PATH']
module.exports = config
