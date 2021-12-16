import { config } from 'arb-bridge-eth/hardhat.dev-config'

try {
  require('arb-upgrades/peripheralsTasks')
} catch (e) {
  // arb-upgrades dependency not available
}

module.exports = config
