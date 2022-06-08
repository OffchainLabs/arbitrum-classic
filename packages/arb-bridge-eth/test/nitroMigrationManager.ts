import { RollupCreatedEvent } from '../build/types/nitro/RollupCreator'
import { BridgeCreator__factory as NitroBridgeCreator__factory } from '../build/types/nitro/factories/BridgeCreator__factory'
import { Bridge__factory as NitroBridge__factory } from '../build/types/nitro/factories/Bridge__factory'
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
import { BigNumber, constants, Signer } from 'ethers'
import { Provider } from '@ethersproject/providers'
import { getContractAddress } from 'ethers/lib/utils'
import {
  Inbox__factory,
  NitroMigrator,
  NitroMigrator__factory,
  ProxyAdmin__factory,
  RollupAdminFacet__factory,
  RollupUserFacet__factory,
  Rollup__factory,
  SequencerInbox__factory,
} from '../build/types'

// CHRIS: TODO: comments up in here
export class NitroMigrationManager {
  private readonly provider: Provider

  public constructor(
    public readonly proxyAdminOwner: Signer,
    public readonly classicConfig: {
      rollupAddr: string
      proxyAdminAddr: string
      inboxAddr: string
      sequencerInboxAddr: string
      bridgeAddr: string
      rollupEventBridgeAddr: string
      outboxV1: string
      outboxV2: string // CHRIS: TODO: v2 here?
    }
  ) {
    if (!proxyAdminOwner.provider) {
      throw new Error('No provider attached to deployer signer.')
    }
    this.provider = proxyAdminOwner.provider
  }

  public async run(
    classicSequencers: string[],
    nitroConfig: Parameters<NitroRollupCreator['createRollup']>[0],
    destroyAlternatives: boolean,
    destroyChallenges: boolean
  ) {
    const nitroContracts = await this.deployNitroContracts(nitroConfig)
    await this.upgradeClassicContracts()
    await this.deployMigrator({
        bridgeAddr: nitroContracts.bridge,
        inboxTemplateAddr: nitroContracts.inboxTemplate,
        outboxAddr: nitroContracts.outbox,
        sequencerInboxAddr: nitroContracts.sequencerInbox,
    })
    await this.transferClassicOwnership()
    await this.step1(classicSequencers, {
      rollupAddr: nitroContracts.rollup,
      bridgeAddr: nitroContracts.bridge,
    })

    const nodeNum = await this.waitForNodeConfirmation()

    await this.step2(nodeNum, destroyAlternatives, destroyChallenges)

    await this.step3()
  }

  private async waitForNodeConfirmation(): Promise<BigNumber> {
    throw new Error('Not implemented.')
  }

  private async deployNitroChallengeContracts(signer: Signer) {
    const oneStepProver0Fac = new NitroOneStepProver0__factory(signer)
    const oneStepProver0 = await oneStepProver0Fac.deploy()
    await oneStepProver0.deployed()

    const oneStepProverMemoryFac = new NitroOneStepProverMemory__factory(signer)
    const oneStepProverMemory = await oneStepProverMemoryFac.deploy()
    await oneStepProverMemory.deployed()

    const oneStepProverMathFac = new NitroOneStepProverMath__factory(signer)
    const oneStepProverMath = await oneStepProverMathFac.deploy()
    await oneStepProverMath.deployed()

    const oneStepProverHostIoFac = new NitroOneStepProverHostIo__factory(signer)
    const oneStepProverHostIo = await oneStepProverHostIoFac.deploy()
    await oneStepProverHostIo.deployed()

    const oneStepProofEntryFac = new NitroOneStepProofEntry__factory(signer)
    const oneStepProofEntry = await oneStepProofEntryFac.deploy(
      oneStepProver0.address,
      oneStepProverMemory.address,
      oneStepProverMath.address,
      oneStepProverHostIo.address
    )
    await oneStepProofEntry.deployed()

    const challengeManagerFac = new NitroChallengeManager__factory(signer)
    const challengeManager = await challengeManagerFac.deploy()
    await challengeManager.deployed()

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
    config: Parameters<NitroRollupCreator['createRollup']>[0]
  ) {
    const nitroBridgeCreatorFac = new NitroBridgeCreator__factory(
      this.proxyAdminOwner
    )
    const nitroBridgeCreator = await nitroBridgeCreatorFac.deploy()
    await nitroBridgeCreator.deployed()

    const nitroRollupCreatorFac = new NitroRollupCreator__factory(
      this.proxyAdminOwner
    )
    const nitroRollupCreator = await nitroRollupCreatorFac.deploy()
    await nitroRollupCreator.deployed()

    const nitroRollupAdminLogicFac = new NitroRollupAdminLogic__factory(
      this.proxyAdminOwner
    )
    const nitroRollupAdminLogic = await nitroRollupAdminLogicFac.deploy()
    await nitroRollupAdminLogic.deployed()

    const nitroRollupUserLogicFac = new NitroRollupUserLogic__factory(
      this.proxyAdminOwner
    )
    const nitroRollupUserLogic = await nitroRollupUserLogicFac.deploy()
    await nitroRollupUserLogic.deployed()

    const challengeContracts = await this.deployNitroChallengeContracts(
      this.proxyAdminOwner
    )
    await (
      await nitroRollupCreator.setTemplates(
        nitroBridgeCreator.address,
        challengeContracts.oneStepProofEntry.address,
        challengeContracts.challengeManager.address,
        nitroRollupAdminLogic.address,
        nitroRollupUserLogic.address
      )
    ).wait()

    const nonce = await this.provider.getTransactionCount(
      nitroRollupCreator.address
    )
    const expectedRollupAddress = getContractAddress({
      from: nitroRollupCreator.address,
      nonce: nonce + 2,
    })

    const createRollupTx = await nitroRollupCreator.createRollup(
      config,
      expectedRollupAddress
    )

    // CHRIS: TODO: quite a cool idea would be to figure out at compile
    // time what possible events could be emitted from a given tx? is that even possible,
    // I guess not. So how could we do it? we cant

    // CHRIS: we're deploying a new proxy admin in createRollup
    // CHRIS: this will mean we actually have 2 proxy admins in the system post nitro
    // CHRIS: we should probably transfer ownership so that they all have the same proxy admin
    const createRollupReceipt = await createRollupTx.wait()
    const rollupCreatedEventArgs = createRollupReceipt.logs
      .filter(
        l =>
          l.topics[0] ===
          nitroRollupCreator.interface.getEventTopic(
            'RollupCreated(address,address,address,address,address)'
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

    return {
      rollup: rollupCreatedEventArgs.rollupAddress,
      bridge: rollupCreatedEventArgs.delayedBridge,
      inboxTemplate: await nitroBridgeCreator.inboxTemplate(),
      outbox: await rollupUser.outbox(),
      sequencerInbox: rollupCreatedEventArgs.sequencerInbox,
    }
  }

  public async upgradeClassicContracts() {
    const proxyAdminContractFac = new ProxyAdmin__factory(this.proxyAdminOwner)
    const proxyAdmin = proxyAdminContractFac.attach(
      this.classicConfig.proxyAdminAddr
    )

    const inboxFac = new Inbox__factory(this.proxyAdminOwner)
    const newInboxImp = await inboxFac.deploy()
    await newInboxImp.deployed()
    await proxyAdmin
      // CHRIS: TODO: this should be upgradeAndCall
      .upgrade(this.classicConfig.inboxAddr, newInboxImp.address)

    // -- sequencer inbox
    const sequencerInboxFac = new SequencerInbox__factory(this.proxyAdminOwner)
    const newSequencerInboxImp = await sequencerInboxFac.deploy()
    await newSequencerInboxImp.deployed()
    const sequencerInboxPostUpdgrade =
      newSequencerInboxImp.interface.encodeFunctionData('postUpgradeInit')
    await proxyAdmin.upgradeAndCall(
      this.classicConfig.sequencerInboxAddr,
      newSequencerInboxImp.address,
      sequencerInboxPostUpdgrade
    )

    // -- rollup
    const rollupFac = new Rollup__factory(this.proxyAdminOwner)
    const prevRollup = rollupFac.attach(this.classicConfig.rollupAddr)
    const confirmPeriodBlocks = await prevRollup.confirmPeriodBlocks()
    const newRollupImp = await rollupFac.deploy(confirmPeriodBlocks)
    await newRollupImp.deployed()
    const rollupPostUpgrade =
      newRollupImp.interface.encodeFunctionData('postUpgradeInit')
    await proxyAdmin.upgradeAndCall(
      this.classicConfig.rollupAddr,
      newRollupImp.address,
      rollupPostUpgrade
    )

    // -- rollup user
    const rollupUserFac = new RollupUserFacet__factory(this.proxyAdminOwner)
    const newRollupUserImp = await rollupUserFac.deploy()
    await newRollupUserImp.deployed()

    // -- rollup admin
    const rollupAdminFac = new RollupAdminFacet__factory(this.proxyAdminOwner)
    const newRollupAdminImp = await rollupAdminFac.deploy()
    await newRollupAdminImp.deployed()

    const rollupAdmin = rollupAdminFac
      .attach(this.classicConfig.rollupAddr)
      .connect(this.proxyAdminOwner)
    await (
      await rollupAdmin.setFacets(
        newRollupAdminImp.address,
        newRollupUserImp.address
      )
    ).wait()

    return {
      inbox: inboxFac.attach(this.classicConfig.inboxAddr),
      sequencerInbox: sequencerInboxFac.attach(
        this.classicConfig.sequencerInboxAddr
      ),
      rollupAdmin: rollupAdminFac.attach(this.classicConfig.rollupAddr),
    }
  }

  // CHRIS: TODO: check for the presence of this everywhere
  private nitroMigrator?: NitroMigrator

  public async deployMigrator(nitroConfig: {
    bridgeAddr: string
    outboxAddr: string
    sequencerInboxAddr: string
    inboxTemplateAddr: string
  }) {
    const nitroMigratorFac = new NitroMigrator__factory(this.proxyAdminOwner)
    this.nitroMigrator = await nitroMigratorFac.deploy(
      this.classicConfig.inboxAddr,
      this.classicConfig.sequencerInboxAddr,
      this.classicConfig.bridgeAddr,
      this.classicConfig.rollupEventBridgeAddr,
      this.classicConfig.outboxV1,
      this.classicConfig.outboxV2,
      this.classicConfig.rollupAddr,
      nitroConfig.bridgeAddr,
      nitroConfig.outboxAddr,
      nitroConfig.sequencerInboxAddr,
      nitroConfig.inboxTemplateAddr
    )

    return this.nitroMigrator;
  }

  public async transferClassicOwnership() {
    if (!this.nitroMigrator)
      throw new Error('Transfer ownership called before migrator deployed.')

    const rollupAdminFac = new RollupAdminFacet__factory(this.proxyAdminOwner)
    const rollupAdmin = rollupAdminFac
      .attach(this.classicConfig.rollupAddr)
      .connect(this.proxyAdminOwner)

    await (
      await rollupAdmin.transferOwnership(
        this.classicConfig.bridgeAddr,
        this.nitroMigrator.address
      )
    ).wait()
    await (await rollupAdmin.setOwner(this.nitroMigrator.address)).wait()
  }

  // CHRIS: TODO: ensure these functions are called in the correct order?

  // CHRIS: TODO: put this classic config in the constructor
  public async step1(
    classicSequencers: string[],
    nitroConfig: { rollupAddr: string; bridgeAddr: string }
  ) {
    if (!this.nitroMigrator)
      throw new Error('Step 1 called before migrator deployed.')

    // CHRIS: TODO: should nitro contracts be added to dev or prod dependencies?

    const nitroBridgeFac = new NitroBridge__factory(this.proxyAdminOwner)
    const nitroBridge = nitroBridgeFac.attach(nitroConfig.bridgeAddr)
    const enqueueDelayedMessage =
      await nitroBridge.interface.encodeFunctionData('enqueueDelayedMessage', [
        0,
        this.nitroMigrator.address,
        constants.HashZero,
      ])

    await (
      await this.nitroMigrator.functions.nitroStep1(
        classicSequencers,
        enqueueDelayedMessage
      )
    ).wait()
  }

  public async step2(finalNodeNum: BigNumber, destroyAlternatives: boolean, destroyChallenges: boolean) {
    if (!this.nitroMigrator)
      throw new Error('Step 2 called before migrator deployed.')

      // CHRIS: TODO: pass these args through
    await (await this.nitroMigrator.nitroStep2(finalNodeNum, destroyAlternatives, destroyChallenges)).wait()
  }

  public async step3() {
    if (!this.nitroMigrator)
      throw new Error('Step 3 called before migrator deployed.')

    await (await this.nitroMigrator.nitroStep3()).wait()
  }
}
