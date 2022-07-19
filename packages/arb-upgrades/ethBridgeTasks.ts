import { task } from 'hardhat/config'
import { HardhatRuntimeEnvironment } from 'hardhat/types'
import { initUpgrades, getAdminFromProxyStorage } from './index'

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

task('configure-migration', 'configure nitro migrator contract')
  .addParam('migrator', '')
  .addParam('nitrorollupproxy', '')

  .setAction(async (args, hre) => {
    const { getDeployments } = initUpgrades(hre, process.cwd())
    const { data } = await getDeployments()

    const NewRollupUser = (await hre.ethers.getContractAt('INitroRollupCore', args.nitrorollupproxy))
      .connect('0x1111111111111111111111111111111111111111');
    const replacingInbox = await NewRollupUser.inbox();

    const oldProxyAdmin = await getAdminFromProxyStorage(hre, data.contracts.Inbox.proxyAddress)
    const newProxyAdmin = await getAdminFromProxyStorage(hre, replacingInbox)

    const NewProxyAdmin = (await hre.ethers.getContractFactory('ProxyAdmin'))
      .attach(newProxyAdmin)
      .connect(hre.ethers.provider)
    const owner = await NewProxyAdmin.owner();

    const Migrator = (await hre.ethers.getContractFactory('NitroMigrator'))
      .attach(args.migrator)
      .connect(hre.ethers.provider.getSigner(owner))
    const migratorOwner = await Migrator.owner()
    if (migratorOwner != owner) {
      throw new Error('Migrator has wrong owner. Expected ' + owner + ' but got ' + migratorOwner)
    }

    const OldProxyAdmin = (await hre.ethers.getContractFactory('ProxyAdmin'))
      .attach(oldProxyAdmin)
      .connect(hre.ethers.provider.getSigner(owner))
    const OldRollup = (await hre.ethers.getContractFactory('RollupAdminFacet'))
      .attach(data.contracts.Rollup.proxyAddress)
      .connect(hre.ethers.provider.getSigner(owner))
    const NewRollup = (await hre.ethers.getContractFactory('RollupAdminFacet'))
      .attach(args.nitrorollupproxy)
      .connect(hre.ethers.provider.getSigner(owner))
    if (await OldProxyAdmin.owner() != Migrator.address) {
      console.log('Transferring ownership of old proxy admin')
      await (await OldProxyAdmin.transferOwnership(args.migrator)).wait()
    }
    if (await OldRollup.owner() != Migrator.address) {
      console.log('Transferring ownership of classic rollup')
      await (await OldRollup.setOwner(args.migrator)).wait()
    }
    if (await NewRollupUser.owner() != Migrator.address) {
      console.log('Transferring ownership of nitro rollup')
      await (await NewRollup.setOwner(args.migrator)).wait()
    }

    console.log('Configuring deployment on nitro migrator')
    const initializeRes = await Migrator.configureDeployment(
      data.contracts.Inbox.proxyAddress,
      data.contracts.SequencerInbox.proxyAddress,
      data.contracts.Bridge.proxyAddress,
      data.contracts.RollupEventBridge.proxyAddress,
      data.contracts.OldOutbox.proxyAddress,
      data.contracts.Outbox.proxyAddress,
      data.contracts.Rollup.proxyAddress,
      oldProxyAdmin,
      args.nitrorollupproxy,
      newProxyAdmin,
    )
    const initializeRec = await initializeRes.wait()
    console.log('Nitro migrator configured:', initializeRec)
  })

task('migration-step-1', 'run nitro migration step 1')
  .addParam('migrator', '')

  .setAction(async (args, hre) => {
    let Migrator = (await hre.ethers.getContractFactory('NitroMigrator'))
      .attach(args.migrator)
      .connect(hre.ethers.provider)
    const owner = await Migrator.owner()
    Migrator = Migrator.connect(hre.ethers.provider.getSigner(owner))

    console.log('Running migration step 1')
    const receipt = await (await Migrator.nitroStep1()).wait()
    console.log('Ran migration step 1:', receipt)
  })

task('migration-step-2', 'run nitro migration step 2')
  .addParam('migrator', '')
  .addOptionalParam('finalNodeNum')
  .addFlag('destroyAlternatives')
  .addFlag('destroyChallenges')

  .setAction(async (args, hre) => {
    const { getDeployments } = initUpgrades(hre, process.cwd())
    const { data } = await getDeployments()

    let Migrator = (await hre.ethers.getContractFactory('NitroMigrator'))
      .attach(args.migrator)
      .connect(hre.ethers.provider)
    const owner = await Migrator.owner()
    Migrator = Migrator.connect(hre.ethers.provider.getSigner(owner))

    let finalNodeNum: any = parseInt(args.finalNodeNum)
    if (!finalNodeNum) {
      const Rollup = (await hre.ethers.getContractFactory('RollupUserFacet'))
        .attach(data.contracts.Rollup.proxyAddress)
        .connect(hre.ethers.provider)
      finalNodeNum = await Rollup.latestNodeCreated();
      console.log('Resolved final node number', finalNodeNum)
    }

    console.log('Running migration step 2')
    const receipt = await (await Migrator.nitroStep2(finalNodeNum, args.destroyAlternatives, args.destroyChallenges)).wait()
    console.log('Ran migration step 2:', receipt)
  })

task('migration-step-3', 'run nitro migration step 3')
  .addParam('migrator', '')
  .addParam('genesisnumber', 'The nitro genesis block number')
  .addParam('genesishash', 'The nitro genesis block hash')

  .setAction(async (args, hre) => {
    let Migrator = (await hre.ethers.getContractFactory('NitroMigrator'))
      .attach(args.migrator)
      .connect(hre.ethers.provider)
    const owner = await Migrator.owner()
    Migrator = Migrator.connect(hre.ethers.provider.getSigner(owner))

    console.log('Running migration step 3')
    const receipt = await (await Migrator.nitroStep3(args.genesisnumber, args.genesishash, false)).wait()
    console.log('Ran migration step 3:', receipt)
  })

task('migrator-transfer-child-ownership', 'transfer the ownership of a contract owned by the nitro migrator')
  .addParam('migrator', '')
  .addParam('child', '')
  .addParam('newowner', '')

  .setAction(async (args, hre) => {
    let Migrator = (await hre.ethers.getContractFactory('NitroMigrator'))
      .attach(args.migrator)
      .connect(hre.ethers.provider)
    const owner = await Migrator.owner()
    Migrator = Migrator.connect(hre.ethers.provider.getSigner(owner))

    const receipt = await (await Migrator.transferOtherContractOwnership(args.child, args.newowner)).wait()
    console.log('Transferred ownership:', receipt)
  })

task('migrator-transfer-rollup-ownership', 'transfer the ownership a rollup owned by the nitro migrator')
  .addParam('migrator', '')
  .addParam('rollup', '')
  .addParam('newowner', '')

  .setAction(async (args, hre) => {
    let Migrator = (await hre.ethers.getContractFactory('NitroMigrator'))
      .attach(args.migrator)
      .connect(hre.ethers.provider)
    const owner = await Migrator.owner()
    Migrator = Migrator.connect(hre.ethers.provider.getSigner(owner))

    const setOwnerData = (await hre.ethers.getContractFactory('RollupAdminFacet'))
      .interface
      .encodeFunctionData("setOwner", [args.newowner]);
    const receipt = await (await Migrator.executeTransaction(setOwnerData, args.rollup, 0)).wait()
    console.log('Transferred ownership:', receipt)
  })

task('migrator-add-arbos-owner', 'adds an ArbOS chain owner via the nitro migrator')
  .addParam('migrator', '')
  .addParam('newowner', '')

  .setAction(async (args, hre) => {
    let Migrator = (await hre.ethers.getContractFactory('NitroMigrator'))
      .attach(args.migrator)
      .connect(hre.ethers.provider)
    const owner = await Migrator.owner()
    Migrator = Migrator.connect(hre.ethers.provider.getSigner(owner))

    const receipt = await (await Migrator.addArbosOwner(args.newowner)).wait()
    console.log('Added owner:', receipt)
  })
