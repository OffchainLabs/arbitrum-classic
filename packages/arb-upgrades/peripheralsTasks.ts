import { task } from 'hardhat/config'
import { initUpgrades } from './index'

task('deploy-logic-one', 'deploy one logic')
  .addParam('contract', 'contract to deploy')
  .setAction(async (args, hre) => {
    const { contract } = args
    const { deployLogic } = initUpgrades(hre, process.cwd())
    await deployLogic(contract)
  })

task('deploy-logic-all', 'deploy all logic contracts').setAction(
  async (_, hre) => {
    const { deployLogicAll } = initUpgrades(hre, process.cwd())
    await deployLogicAll()
  }
)

task('trigger-upgrades', 'triggers upgrade').setAction(async (_, hre) => {
  const { updateImplementations } = initUpgrades(hre, process.cwd())
  await updateImplementations()
})

task('verify-deployments', 'verifies implementations').setAction(
  async (_, hre) => {
    const { verifyCurrentImplementations } = initUpgrades(hre, process.cwd())
    await verifyCurrentImplementations()
  }
)

task('transfer-owner', 'deploy one logic')
  .addParam('proxyaddress', 'proxy address')
  .addParam('newadmin', 'address of new admin')
  .setAction(async (args, hre) => {
    const { contract } = args
    const { transferAdmin } = initUpgrades(hre, process.cwd())
    await transferAdmin(args.proxyaddress, args.newadmin)
  })

task(
  'remove-build-info',
  'remove giant build info string from current_deployments json'
).setAction(async (_, hre) => {
  const { removeBuildInfoFiles } = initUpgrades(hre, process.cwd())
  await removeBuildInfoFiles()
})

task('etherscan-verify', 'verify current deployments in etherscan').setAction(
  async (_, hre) => {
    const { verifyDeployments } = await initUpgrades(hre, process.cwd())
    await verifyDeployments()
  }
)
