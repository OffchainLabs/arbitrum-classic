import { RollupCreatedEvent } from '../build/types/nitro/RollupCreator'
import { BridgeCreator__factory as NitroBridgeCreator__factory } from '../build/types/nitro/factories/BridgeCreator__factory'
import { RollupCreator__factory as NitroRollupCreator__factory } from '../build/types/nitro/factories/RollupCreator__factory'
import { RollupCreator as NitroRollupCreator } from '../build/types/nitro/RollupCreator'
import { OneStepProver0__factory as NitroOneStepProver0__factory } from '../build/types/nitro/factories/OneStepProver0__factory'
import { OneStepProverMemory__factory as NitroOneStepProverMemory__factory } from '../build/types/nitro/factories/OneStepProverMemory__factory'
import { OneStepProverMath__factory as NitroOneStepProverMath__factory } from '../build/types/nitro/factories/OneStepProverMath__factory'
import { OneStepProverHostIo__factory as NitroOneStepProverHostIo__factory } from '../build/types/nitro/factories/OneStepProverHostIo__factory'
import { OneStepProofEntry__factory as NitroOneStepProofEntry__factory } from '../build/types/nitro/factories/OneStepProofEntry__factory'
import { ChallengeManager__factory as NitroChallengeManager__factory } from '../build/types/nitro/factories/ChallengeManager__factory'
import { RollupAdminLogic__factory as NitroRollupAdminLogic__factory } from '../build/types/nitro/factories/RollupAdminLogic__factory'
import { RollupUserLogic__factory as NitroRollupUserLogic__factory } from '../build/types/nitro/factories/RollupUserLogic__factory'
import { ValidatorUtils__factory as NitroValidatorUtils__factory } from '../build/types/nitro/factories/ValidatorUtils__factory'
import { ValidatorWalletCreator__factory as NitroValidatorWalletCreator__factory } from '../build/types/nitro/factories/ValidatorWalletCreator__factory'
import { BigNumber, Signer } from 'ethers'
import { Provider } from '@ethersproject/providers'
import { getAddress, getContractAddress } from 'ethers/lib/utils'
import {
  Bridge__factory,
  Inbox__factory,
  NitroMigrator,
  NitroMigrator__factory,
  ProxyAdmin__factory,
  RollupAdminFacet__factory,
  RollupUserFacet__factory,
  Rollup__factory,
  SequencerInbox__factory,
} from '../build/types'

const wait = (ms: number) =>
  new Promise((resolve, _) => setTimeout(resolve, ms))

interface ClassicConfig {
  rollupAddr: string
  proxyAdminAddr: string
  inboxAddr: string
  sequencerInboxAddr: string
  bridgeAddr: string
  rollupEventBridgeAddr: string
  outboxV1: string
  outboxV2: string
}

export class NitroMigrationManager {
  private readonly provider: Provider

  public static async deploy(
    nitroDeployer: Signer,
    log = true,
    skipStep3Check = false
  ) {
    if (log)
      console.log(`Proxy admin owner: ${await nitroDeployer.getAddress()}`)
    const nitroMigratorFac = new NitroMigrator__factory(nitroDeployer)
    const nitroMigrator = await nitroMigratorFac.deploy()
    await nitroMigrator.deployed()
    if (log) console.log(`Nitro migrator: ${nitroMigrator.address}`)
    return new NitroMigrationManager(nitroMigrator, log, skipStep3Check)
  }

  public constructor(
    public readonly migrator: NitroMigrator,
    public readonly log = true,
    public readonly skipStep3Check: boolean = false
  ) {
    if (!migrator.provider) {
      throw new Error('No provider attached to migrator.')
    }

    this.provider = migrator.provider
  }

  public async run(
    nitroDeployer: Signer,
    classicProxyAdminOwner: Signer,
    nitroConfig: Omit<
      Parameters<NitroRollupCreator['createRollup']>[0],
      'owner'
    >,
    classicConfig: ClassicConfig,
    destroyAlternatives: boolean,
    destroyChallenges: boolean
  ) {
    if (this.log) console.log('Beginning migration')

    const nitroContracts = await this.deployNitroContracts(
      nitroDeployer,
      nitroConfig
    )
    await this.upgradeClassicContracts(classicProxyAdminOwner, classicConfig)

    await this.configureDeployment(
      classicProxyAdminOwner,
      nitroDeployer,
      classicConfig,
      nitroContracts.rollup,
      nitroContracts.proxyAdmin
    )

    await this.step1()

    const nodeNum = await this.getFinalNodeNum()
    await this.step2(nodeNum, destroyAlternatives, destroyChallenges)

    await this.waitForConfirmedEqualLatest()
    await this.step3(nitroDeployer)
  }

  private async deployNitroChallengeContracts(signer: Signer) {
    if (this.log) console.log(`Deploying nitro challenge contracts`)
    const oneStepProver0Fac = new NitroOneStepProver0__factory(signer)
    const oneStepProver0 = await oneStepProver0Fac.deploy()

    const oneStepProverMemoryFac = new NitroOneStepProverMemory__factory(signer)
    const oneStepProverMemory = await oneStepProverMemoryFac.deploy()

    const oneStepProverMathFac = new NitroOneStepProverMath__factory(signer)
    const oneStepProverMath = await oneStepProverMathFac.deploy()

    const oneStepProverHostIoFac = new NitroOneStepProverHostIo__factory(signer)
    const oneStepProverHostIo = await oneStepProverHostIoFac.deploy()

    await oneStepProver0.deployed()
    await oneStepProverMemory.deployed()
    await oneStepProverMath.deployed()
    await oneStepProverHostIo.deployed()
    if (this.log) {
      console.log(`Nitro one step prover 0: ${oneStepProver0.address}`)
      console.log(
        `Nitro one step prover memory: ${oneStepProverMemory.address}`
      )
      console.log(`Nitro one step prover math: ${oneStepProverMath.address}`)
      console.log(
        `Nitro one step prover host io: ${oneStepProverHostIo.address}`
      )
    }

    const oneStepProofEntryFac = new NitroOneStepProofEntry__factory(signer)
    const oneStepProofEntry = await oneStepProofEntryFac.deploy(
      oneStepProver0.address,
      oneStepProverMemory.address,
      oneStepProverMath.address,
      oneStepProverHostIo.address
    )

    const challengeManagerFac = new NitroChallengeManager__factory(signer)
    const challengeManager = await challengeManagerFac.deploy()

    await oneStepProofEntry.deployed()
    await challengeManager.deployed()
    if (this.log) {
      console.log(`Nitro one step prover entry: ${oneStepProofEntry.address}`)
      console.log(`Nitro challenge manager: ${challengeManager.address}`)
    }

    if (this.log) console.log(`Deploying nitro challenge contracts complete`)

    return {
      oneStepProver0,
      oneStepProverMemory,
      oneStepProverMath,
      oneStepProverHostIo,
      oneStepProofEntry,
      challengeManager,
    }
  }

  public async deployNitroContracts(
    nitroDeployer: Signer,
    config: Omit<Parameters<NitroRollupCreator['createRollup']>[0], 'owner'>
  ) {
    if (this.log) console.log('Deploying nitro contracts')
    // the owner should always be our nitro deployer
    const ownerConfig = {
      ...config,
      owner: await nitroDeployer.getAddress(),
    }

    // quick check that the owner of the migrator is also the account we'll
    // use for deploying here
    const migratorOwner = await this.migrator.owner()
    const nitroDeployerAddr = await nitroDeployer.getAddress()
    if (migratorOwner !== nitroDeployerAddr) {
      throw new Error(
        `Incorrect owner. Trying to deploy nitro contracts with different owner to migrator owner. ${migratorOwner}:${nitroDeployerAddr}`
      )
    }

    const nitroValidatorUtilsFac = new NitroValidatorUtils__factory(
      nitroDeployer
    )
    const nitroValidatorUtils = await nitroValidatorUtilsFac.deploy()

    const nitroValidatorWalletCreatorFac =
      new NitroValidatorWalletCreator__factory(nitroDeployer)
    const nitroValidatorWalletCreator =
      await nitroValidatorWalletCreatorFac.deploy()

    const nitroBridgeCreatorFac = new NitroBridgeCreator__factory(nitroDeployer)
    const nitroBridgeCreator = await nitroBridgeCreatorFac.deploy()

    const nitroRollupCreatorFac = new NitroRollupCreator__factory(nitroDeployer)
    const nitroRollupCreator = await nitroRollupCreatorFac.deploy()

    const nitroRollupAdminLogicFac = new NitroRollupAdminLogic__factory(
      nitroDeployer
    )
    const nitroRollupAdminLogic = await nitroRollupAdminLogicFac.deploy()

    const nitroRollupUserLogicFac = new NitroRollupUserLogic__factory(
      nitroDeployer
    )
    const nitroRollupUserLogic = await nitroRollupUserLogicFac.deploy()

    await nitroValidatorUtils.deployed()
    await nitroValidatorWalletCreator.deployed()
    await nitroBridgeCreator.deployed()
    await nitroRollupCreator.deployed()
    await nitroRollupAdminLogic.deployed()
    await nitroRollupUserLogic.deployed()
    if (this.log) {
      console.log(`Nitro validator utils: ${nitroValidatorUtils.address}`)
      console.log(
        `Nitro validator wallet creator: ${nitroValidatorWalletCreator.address}`
      )
      console.log(`Nitro bridge creator: ${nitroBridgeCreator.address}`)
      console.log(`Nitro rollup creator: ${nitroRollupCreator.address}`)
      console.log(`Nitro rollup admin logic: ${nitroRollupAdminLogic.address}`)
      console.log(`Nitro rollup user logic: ${nitroRollupUserLogic.address}`)
    }

    const challengeContracts = await this.deployNitroChallengeContracts(
      nitroDeployer
    )
    await (
      await nitroRollupCreator.setTemplates(
        nitroBridgeCreator.address,
        challengeContracts.oneStepProofEntry.address,
        challengeContracts.challengeManager.address,
        nitroRollupAdminLogic.address,
        nitroRollupUserLogic.address,
        nitroValidatorUtils.address,
        nitroValidatorWalletCreator.address
      )
    ).wait()
    if (this.log) console.log(`Nitro templates set`)

    const nonce = await this.provider.getTransactionCount(
      nitroRollupCreator.address
    )
    const expectedRollupAddress = getContractAddress({
      from: nitroRollupCreator.address,
      nonce: nonce + 2,
    })

    const createRollupTx = await nitroRollupCreator.createRollup(
      ownerConfig,
      expectedRollupAddress
    )

    const createRollupReceipt = await createRollupTx.wait()
    const rollupCreatedEventArgs = createRollupReceipt.logs
      .filter(
        l =>
          l.topics[0] ===
          nitroRollupCreator.interface.getEventTopic(
            'RollupCreated(address indexed,address,address,address,address)'
          )
      )
      .map(
        l =>
          nitroRollupCreator.interface.parseLog(
            l
          ) as unknown as RollupCreatedEvent
      )[0].args

    const rollupUser = nitroRollupUserLogicFac.attach(
      rollupCreatedEventArgs.rollupAddress
    )

    if (this.log) {
      console.log(`Nitro rollup created`)
      console.log(`Nitro inbox: ${rollupCreatedEventArgs.inboxAddress}`)
      console.log(`Nitro rollup: ${rollupCreatedEventArgs.rollupAddress}`)
      console.log(`Nitro bridge ${rollupCreatedEventArgs.bridge}`)
      console.log(
        `Nitro inbox template: ${await nitroBridgeCreator.inboxTemplate()}`
      )
      console.log(`Nitro outbox: ${await rollupUser.outbox()}`)
      console.log(
        `Nitro sequencer inbox: ${rollupCreatedEventArgs.sequencerInbox}`
      )
      console.log(`Nitro proxy admin: ${rollupCreatedEventArgs.adminProxy}`)
    }

    if (this.log) console.log('Deploying nitro contracts complete')

    return {
      inbox: rollupCreatedEventArgs.inboxAddress,
      rollup: rollupCreatedEventArgs.rollupAddress,
      bridge: rollupCreatedEventArgs.bridge,
      inboxTemplate: await nitroBridgeCreator.inboxTemplate(),
      outbox: await rollupUser.outbox(),
      sequencerInbox: rollupCreatedEventArgs.sequencerInbox,
      proxyAdmin: rollupCreatedEventArgs.adminProxy,
    }
  }

  public async upgradeClassicContracts(
    classicProxyAdminOwner: Signer,
    classicConfig: {
      proxyAdminAddr: string
      inboxAddr: string
      bridgeAddr: string
      sequencerInboxAddr: string
      rollupAddr: string
    }
  ) {
    if (this.log) console.log(`Upgrading classic contracts`)
    const proxyAdmin = ProxyAdmin__factory.connect(
      classicConfig.proxyAdminAddr,
      classicProxyAdminOwner
    )

    const inboxFac = new Inbox__factory(classicProxyAdminOwner)
    const newInboxImp = await inboxFac.deploy()
    await newInboxImp.deployed()
    await proxyAdmin.upgrade(classicConfig.inboxAddr, newInboxImp.address)
    if (this.log)
      console.log(`Classic inbox upgraded: ${classicConfig.inboxAddr}`)

    const bridgeFac = new Bridge__factory(classicProxyAdminOwner)
    const newBridgeImp = await bridgeFac.deploy()
    await newBridgeImp.deployed()
    await proxyAdmin.upgrade(classicConfig.bridgeAddr, newBridgeImp.address)
    if (this.log)
      console.log(`Classic bridge upgraded: ${classicConfig.bridgeAddr}`)

    // -- sequencer inbox
    const sequencerInboxFac = new SequencerInbox__factory(
      classicProxyAdminOwner
    )
    const newSequencerInboxImp = await sequencerInboxFac.deploy()
    await newSequencerInboxImp.deployed()
    const sequencerInboxPostUpdgrade =
      newSequencerInboxImp.interface.encodeFunctionData('postUpgradeInit')
    await proxyAdmin.upgradeAndCall(
      classicConfig.sequencerInboxAddr,
      newSequencerInboxImp.address,
      sequencerInboxPostUpdgrade
    )
    if (this.log)
      console.log(
        `Classic sequencer inbox upgraded: ${classicConfig.sequencerInboxAddr}`
      )

    // -- rollup
    const rollupFac = new Rollup__factory(classicProxyAdminOwner)
    const prevRollup = rollupFac.attach(classicConfig.rollupAddr)
    const confirmPeriodBlocks = await prevRollup.confirmPeriodBlocks()
    const newRollupImp = await rollupFac.deploy(confirmPeriodBlocks)
    await newRollupImp.deployed()
    const rollupPostUpgrade =
      newRollupImp.interface.encodeFunctionData('postUpgradeInit')
    await proxyAdmin.upgradeAndCall(
      classicConfig.rollupAddr,
      newRollupImp.address,
      rollupPostUpgrade
    )
    if (this.log)
      console.log(`Classic rollup upgraded: ${classicConfig.rollupAddr}`)

    // -- rollup user
    const rollupUserFac = new RollupUserFacet__factory(classicProxyAdminOwner)
    const newRollupUserImp = await rollupUserFac.deploy()
    await newRollupUserImp.deployed()
    if (this.log)
      console.log(
        `Classic rollup user logic deployed: ${newRollupUserImp.address}`
      )

    // -- rollup admin
    const rollupAdminFac = new RollupAdminFacet__factory(classicProxyAdminOwner)
    const newRollupAdminImp = await rollupAdminFac.deploy()
    await newRollupAdminImp.deployed()
    if (this.log)
      console.log(
        `Classic rollup admin logic deployed: ${newRollupAdminImp.address}`
      )

    const rollupAdmin = rollupAdminFac
      .attach(classicConfig.rollupAddr)
      .connect(classicProxyAdminOwner)
    await (
      await rollupAdmin.setFacets(
        newRollupAdminImp.address,
        newRollupUserImp.address
      )
    ).wait()
    if (this.log) console.log(`Classic rollup facets set`)

    if (this.log) console.log(`Upgrading classic contracts complete`)
    return {
      sequencerInbox: sequencerInboxFac.attach(
        classicConfig.sequencerInboxAddr
      ),
      rollupAdmin: rollupAdminFac.attach(classicConfig.rollupAddr),
    }
  }

  public async configureDeployment(
    classicProxyAdminOwner: Signer,
    nitroDeployer: Signer,
    classicConfig: ClassicConfig,
    nitroRollupAddr: string,
    nitroProxyAdmin: string
  ) {
    this.provider
    if (
      (await this.provider.getCode(nitroRollupAddr)).length <= 2 ||
      (await this.provider.getCode(nitroProxyAdmin)).length <= 2
    ) {
      throw new Error(
        'Could not configure deployment. Nitro contracts not deployed.'
      )
    }
    if (this.log) console.log('Configuring deployment')

    const classicRollupAdmin = RollupAdminFacet__factory.connect(
      classicConfig.rollupAddr,
      classicProxyAdminOwner
    )
    if ((await classicRollupAdmin.owner()) != this.migrator.address) {
      if (this.log) console.log('Classic rollup, setting owner to migrator')
      await (await classicRollupAdmin.setOwner(this.migrator.address)).wait()
    }

    const classicProxyAdmin = ProxyAdmin__factory.connect(
      classicConfig.proxyAdminAddr,
      classicProxyAdminOwner
    )
    if ((await classicProxyAdmin.owner()) != this.migrator.address) {
      if (this.log)
        console.log('Classic proxy admin, setting owner to migrator')
      await (
        await classicProxyAdmin.transferOwnership(this.migrator.address)
      ).wait()
    }

    const nitroRollupAdmin = NitroRollupAdminLogic__factory.connect(
      nitroRollupAddr,
      nitroDeployer
    )
    const nitroRollupAdminOwner = await this.getProxyAdmin(nitroRollupAddr)
    if (nitroRollupAdminOwner != this.migrator.address) {
      if (this.log) console.log('Nitro rollup, setting owner to migrator')
      await (await nitroRollupAdmin.setOwner(this.migrator.address)).wait()
    }

    await (
      await this.migrator.configureDeployment(
        classicConfig.inboxAddr,
        classicConfig.sequencerInboxAddr,
        classicConfig.bridgeAddr,
        classicConfig.rollupEventBridgeAddr,
        classicConfig.outboxV1,
        classicConfig.outboxV2,
        classicConfig.rollupAddr,
        classicConfig.proxyAdminAddr,
        nitroRollupAddr,
        nitroProxyAdmin
      )
    ).wait()
    if (this.log) console.log('Configure deployment complete')
  }

  public async step1() {
    if (this.log) console.log('Executing migration step 1')
    await (await this.migrator.functions.nitroStep1()).wait()
    if (this.log) console.log('Executing migration step 1 complete')
  }

  public async getFinalNodeNum(): Promise<BigNumber> {
    // CHRIS: TODO: Do we have any unredeemed retryables?

    const rollupAddr = await this.migrator.rollup()
    const rollupAdmin = RollupAdminFacet__factory.connect(
      rollupAddr,
      this.provider
    )

    const finalNodeNum = await rollupAdmin.latestNodeCreated()
    if (this.log) console.log(`Final node num: ${finalNodeNum.toNumber()}`)
    return finalNodeNum
  }

  public async step2(
    finalNodeNum: BigNumber,
    destroyAlternatives: boolean,
    destroyChallenges: boolean
  ) {
    if (this.log) console.log('Executing migration step 2')
    await (
      await this.migrator.nitroStep2(
        finalNodeNum,
        destroyAlternatives,
        destroyChallenges
      )
    ).wait()
    if (this.log) console.log('Executing migration step 2')
  }

  public async waitForConfirmedEqualLatest() {
    if (this.skipStep3Check) return

    // wait until the node has confirmed the remaining nodes
    const rollupAddr = await this.migrator.rollup()
    const rollup = RollupUserFacet__factory.connect(rollupAddr, this.provider)
    for (;;) {
      const latestConfirmed = await rollup.latestConfirmed()
      const latestNodeCreated = await rollup.latestNodeCreated()

      console.log(
        `Waiting for latestConfirmed: ${latestConfirmed.toNumber()} to equal latestNodeCreated: ${latestNodeCreated.toNumber()}.`
      )

      if (latestConfirmed.eq(latestNodeCreated)) break
      await wait(30000)
    }
  }

  public async step3(nitroDeployer: Signer) {
    if (this.log) console.log('Executing migration step 3')
    await (await this.migrator.nitroStep3(this.skipStep3Check)).wait()

    // CHRIS: TODO: should we check the ownership of all contracts?

    const nitroProxyAdminOwner = await nitroDeployer.getAddress()

    // check that ownership was successfully relinquished
    const classicRollupAdmin = RollupAdminFacet__factory.connect(
      await this.migrator.rollup(),
      this.provider
    )

    if ((await classicRollupAdmin.owner()) != nitroProxyAdminOwner) {
      throw new Error(
        `Classic rollup owner is not nitro proxy admin owner. ${await classicRollupAdmin.owner()}:${nitroProxyAdminOwner}`
      )
    }

    const classicProxyAdminAddr = await this.migrator.classicProxyAdmin()
    const classicProxyAdmin = ProxyAdmin__factory.connect(
      classicProxyAdminAddr,
      this.provider
    )
    if ((await classicProxyAdmin.owner()) != nitroProxyAdminOwner) {
      throw new Error(
        `Classic proxy admin owner is not nitro proxy admin owner ${await classicProxyAdmin.owner()}:${nitroProxyAdminOwner}`
      )
    }

    const nitroProxyAdminAddr = await this.migrator.nitroProxyAdmin()
    const nitroRollup = await this.migrator.nitroRollup()
    const nitroAdmin = await this.getProxyAdmin(nitroRollup)
    if (nitroAdmin != nitroProxyAdminAddr) {
      throw new Error(
        `Nitro rollup admin is not nitro proxy admin. ${nitroAdmin}:${nitroProxyAdminAddr}`
      )
    }
    if (this.log) console.log('Executing migration step 3 complete')
  }

  private async getProxyAdmin(proxyAddress: string) {
    const ADMIN_SLOT =
      '0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103'

    const nitroAdmin = await this.provider.getStorageAt(
      proxyAddress,
      ADMIN_SLOT
    )
    return getAddress(
      nitroAdmin.length > 42
        ? '0x' + nitroAdmin.substring(nitroAdmin.length - 40)
        : nitroAdmin
    )
  }
}
