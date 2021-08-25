import * as config from 'arb-bridge-eth/hardhat.config'
import { task } from 'hardhat/config'
import { initUpgrades } from 'arb-upgrades'
import 'hardhat-deploy-ethers'

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

task('verify-deployments', 'verifies implementations').setAction(
  async (_, hre) => {
    const { verifyCurrentImplementations } = initUpgrades(hre, __dirname)
    await verifyCurrentImplementations()
  }
)

task('transfer-owner', 'deploy one logic')
  .addParam('proxyAddress', 'proxy address')
  .addParam('newAdmin', 'address of new admin')
  .setAction(async (args, hre) => {
    const { contract } = args
    const { transferAdmin } = initUpgrades(hre, __dirname)
    await transferAdmin(args.proxyAddress, args.newAdmin)
  })

module.exports = config
