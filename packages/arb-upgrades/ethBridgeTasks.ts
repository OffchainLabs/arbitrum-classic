import { task } from 'hardhat/config'
import { HardhatRuntimeEnvironment } from 'hardhat/types'
import { initUpgrades } from './index'

const handleFork = async (hre: HardhatRuntimeEnvironment) => {
  const network = hre.network
  if (network.name === 'fork') {
    await hre.network.provider.send('hardhat_setBalance', [
      (await hre.ethers.getSigners())[0].address,
      '0x16189AD417E380000',
    ])
  }
  return true
}

task('core-deploy-logic-one', 'deploy one logic')
  .addParam('contract', 'contract to deploy')
  .setAction(async (args, hre) => {
    await handleFork(hre)
    const { contract } = args
    const { deployLogic } = initUpgrades(hre, process.cwd())
    await deployLogic(contract)
  })

task('core-deploy-logic-all', 'deploy all logic contracts').setAction(
  async (_, hre) => {
    await handleFork(hre)
    const { deployLogicAll } = initUpgrades(hre, process.cwd())
    await deployLogicAll()
  }
)

task('core-trigger-upgrades', 'triggers upgrade').setAction(async (_, hre) => {
  await handleFork(hre)
  const { updateImplementations } = initUpgrades(hre, process.cwd())
  await updateImplementations()
})

task('core-verify-deployments', 'verifies implementations').setAction(
  async (_, hre) => {
    await handleFork(hre)
    const { verifyCurrentImplementations } = initUpgrades(hre, process.cwd())
    await verifyCurrentImplementations()
  }
)

task('core-transfer-beacon-owner', 'transfers beacon owner')
  .addParam('address', 'beacon contract')
  .addParam('newowner', 'beacon contract')

  .setAction(async (args, hre) => {
    await handleFork(hre)
    const { transferBeaconOwner } = initUpgrades(hre, process.cwd())
    await transferBeaconOwner(args.address, args.newowner)
  })

task('core-transfer-admin', 'transfer proxy admin')
  .addParam('proxyaddress', 'proxy address')
  .addParam('newadmin', 'address of new admin')
  .setAction(async (args, hre) => {
    const { transferAdmin } = initUpgrades(hre, process.cwd())
    await transferAdmin(args.proxyaddress, args.newadmin)
  })

task('etherscan-verify', 'verify current deployments in etherscan').setAction(
  async (_, hre) => {
    const { verifyDeployments } = await initUpgrades(hre, process.cwd())
    await verifyDeployments()
  }
)

task(
  'remove-build-info',
  'remove giant build info string from current_deployments json'
).setAction(async (_, hre) => {
  const { removeBuildInfoFiles } = initUpgrades(hre, process.cwd())
  await removeBuildInfoFiles()
})

task('deploy-outbox-logic', 'deploy and set a new outbox').setAction(
  async (_, hre) => {
    const OutboxFactory = await hre.ethers.getContractFactory('Outbox')
    console.log('Deploying outbox logic')
    const OutboxLogic = await OutboxFactory.deploy()
    await OutboxLogic.deployed()
    console.log('Outbox logic deployed at:', OutboxLogic.address)
  }
)

task('deploy-outbox-proxy', 'deploy outbox proxy')
  .addParam('outboxlogic', 'outbox logic')
  .setAction(async (args, hre) => {
    const { getDeployments } = initUpgrades(hre, process.cwd())
    const { data } = await getDeployments()
    const proxyAdminAddress = data.proxyAdminAddress

    console.log('Deploying Outbox TransparentUpgradeableProxy')
    const TransparentUpgradeableProxyFactory =
      await hre.ethers.getContractFactory('TransparentUpgradeableProxy')
    const OutboxProxyDeployed = await TransparentUpgradeableProxyFactory.deploy(
      args.outboxlogic,
      proxyAdminAddress,
      '0x'
    )
    await OutboxProxyDeployed.deployed()
    console.log('Outbox proxy deployed at', OutboxProxyDeployed.address)
  })

task('init-outbox', 'deploy and set a new outbox')
  .addParam('outboxproxy', '')

  .setAction(async (args, hre) => {
    const { getDeployments } = initUpgrades(hre, process.cwd())
    const { data } = await getDeployments()
    const rollupAddress = data.contracts.Rollup.proxyAddress
    const bridgeAddress = data.contracts.Bridge.proxyAddress

    const Outbox = (await hre.ethers.getContractFactory('Outbox'))
      .attach(args.outboxproxy)
      .connect(hre.ethers.provider)
    const initializeRes = await Outbox.initialize(rollupAddress, bridgeAddress)
    const initializeRec = await initializeRes.wait()
    console.log('Outbox initialized', initializeRec)
  })

task('set-outbox', 'deploy and set a new outbox')
  .addParam('outboxproxy', '')

  .setAction(async (args, hre) => {
    const { getDeployments } = initUpgrades(hre, process.cwd())
    const { data } = await getDeployments()
    const rollupAddress = data.contracts.Rollup.proxyAddress
    console.log('Sanity checking ')
    const Rollup = (await hre.ethers.getContractFactory('Rollup'))
      .attach(rollupAddress)
      .connect(hre.ethers.provider)
    await Rollup.getUserFacet()
    console.log('Rollup sanity checked ')
    const RollupAdmin = (
      await hre.ethers.getContractFactory('RollupAdminFacet')
    )
      .attach(rollupAddress)
      .connect(hre.ethers.provider)
    const setRollupRes = await RollupAdmin.setOutbox(args.outboxproxy)
    const setRollupRec = await setRollupRes.wait()
    console.log('Outbox set', setRollupRec)
    console.log('all set 👍')
  })
