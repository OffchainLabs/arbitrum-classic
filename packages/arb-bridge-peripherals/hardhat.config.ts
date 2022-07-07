import { config } from 'arb-bridge-eth/hardhat.dev-config'
import 'arb-upgrades/peripheralsTasks'

if (process.env['INTERFACE_TESTER_SOLC_VERSION']) {
  config.solidity.compilers.push({
    version: process.env['INTERFACE_TESTER_SOLC_VERSION'],
    settings: {
      optimizer: {
        enabled: true,
        runs: 100,
      },
    },
  })
  config.solidity.overrides = {
    ...config.solidity.overrides,
    'contracts/tokenbridge/test/InterfaceCompatibilityTester.sol': {
        version: process.env['INTERFACE_TESTER_SOLC_VERSION'],
    }
  }
}

module.exports = config
