import * as config from 'arb-bridge-eth/hardhat.config'
import { task } from 'hardhat/config'
import { initUpgrades } from 'arb-upgrades'

require('@openzeppelin/hardhat-upgrades')

task('deploy-logic-one', 'deploy one logic')
  .addParam('contract', 'contract to deploy')
  .setAction(async (args, hre) => {
    const { contract } = args
    const { deployLogic } = initUpgrades(hre, __dirname)
    await deployLogic(contract)
  })

task('deploy-logic-all', 'deploy all logic contracts').setAction(
  async (_, hre) => {
    const { deployLogicAll } = initUpgrades(hre, __dirname)
    await deployLogicAll()
  }
)

task('trigger-upgrades', 'triggers upgrade').setAction(async (_, hre) => {
  const { updateImplementations } = initUpgrades(hre, __dirname)
  await updateImplementations()
})

module.exports = config
