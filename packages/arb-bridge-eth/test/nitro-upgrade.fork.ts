import { ethers, network } from 'hardhat'
import { expect, assert } from 'chai'
import { CurrentDeployments } from 'arb-upgrades/types'
import { readFileSync, readdirSync } from 'fs'
import {
  Bridge,
  NitroMigrator__factory,
  ProxyAdmin,
  RollupAdminFacet__factory,
} from '../build/types'
import { BigNumber, constants, Signer } from 'ethers'
import { Provider } from '@ethersproject/providers'
import { Inbox__factory as NitroInbox__factory } from '../build/types/nitro/factories/Inbox__factory'
import { BridgeCreator__factory as NitroBridgeCreator__factory } from '../build/types/nitro/factories/BridgeCreator__factory'
import { Bridge__factory as NitroBridge__factory } from '../build/types/nitro/factories/Bridge__factory'
import { RollupCreator__factory as NitroRollupCreator__factory } from '../build/types/nitro/factories/RollupCreator__factory'
import { RollupCreatedEvent } from '../build/types/nitro/RollupCreator'
import { OneStepProver0__factory as NitroOneStepProver0__factory } from '../build/types/nitro/factories/OneStepProver0__factory'
import { OneStepProverMemory__factory as NitroOneStepProverMemory__factory } from '../build/types/nitro/factories/OneStepProverMemory__factory'
import { OneStepProverMath__factory as NitroOneStepProverMath__factory } from '../build/types/nitro/factories/OneStepProverMath__factory'
import { OneStepProverHostIo__factory as NitroOneStepProverHostIo__factory } from '../build/types/nitro/factories/OneStepProverHostIo__factory'
import { OneStepProofEntry__factory as NitroOneStepProofEntry__factory } from '../build/types/nitro/factories/OneStepProofEntry__factory'
import { ChallengeManager__factory as NitroChallengeManager__factory } from '../build/types/nitro/factories/ChallengeManager__factory'
import { RollupAdminLogic__factory as NitroRollupAdminLogic__factory } from '../build/types/nitro/factories/RollupAdminLogic__factory'
import { RollupUserLogic__factory as NitroRollupUserLogic__factory } from '../build/types/nitro/factories/RollupUserLogic__factory'
import { defaultAbiCoder, Interface } from '@ethersproject/abi'
import { NitroMigrationManager } from './nitroMigrationManager'
import { arrayify } from '@ethersproject/bytes'

describe('Nitro upgrade', () => {
  const getDeployments = async (provider: Provider) => {
    const chainId = (await provider.getNetwork()).chainId
    const deploymentData = readFileSync(
      `./_deployments/${chainId}_current_deployment.json`
    )
    return JSON.parse(deploymentData.toString()) as CurrentDeployments
  }

  const getProxyAdminSigner = async (proxyAdminAddr: string) => {
    const proxyAdmin = await ethers.getContractAt('ProxyAdmin', proxyAdminAddr)
    const owner = await proxyAdmin.owner()
    // airdrop
    await network.provider.send('hardhat_setBalance', [
      owner,
      '0x16189AD417E380000',
    ])

    // use owner
    await network.provider.request({
      method: 'hardhat_impersonateAccount',
      params: [owner],
    })

    return await ethers.provider.getSigner(owner)
  }

  it('deploy fails if classic contracts havent been upgraded', async () => {
    try {
      const deployments = await getDeployments(ethers.provider)
      const proxyAdminSigner = await getProxyAdminSigner(
        deployments.proxyAdminAddress
      )

      const migrationManager = new NitroMigrationManager(proxyAdminSigner, {
        proxyAdminAddr: deployments.proxyAdminAddress,
        inboxAddr: deployments.contracts.Inbox.proxyAddress,
        rollupAddr: deployments.contracts.Rollup.proxyAddress,
        sequencerInboxAddr: deployments.contracts.SequencerInbox.proxyAddress,
        bridgeAddr: deployments.contracts.Bridge.proxyAddress,
        outboxV1: (deployments.contracts as any)['OldOutbox'].proxyAddress,
        outboxV2: (deployments.contracts as any)['OldOutbox'].proxyAddress, // CHRIS: TODO: v2 here?,
        rollupEventBridgeAddr:
          deployments.contracts.RollupEventBridge.proxyAddress,
      })

      await migrationManager.deployMigrator({
        bridgeAddr: constants.AddressZero,
        inboxTemplateAddr: constants.AddressZero,
        outboxAddr: constants.AddressZero,
        sequencerInboxAddr: constants.AddressZero,
      })

      assert.fail('Expected constructor to fail')
    } catch {}
  })

  it('deploy fails if nitro contracts havent been deployed', async () => {
    try {
      const provider = ethers.provider
      const deployments = await getDeployments(provider)
      const proxyAdminSigner = await getProxyAdminSigner(
        deployments.proxyAdminAddress
      )

      const migrationManager = new NitroMigrationManager(proxyAdminSigner, {
        proxyAdminAddr: deployments.proxyAdminAddress,
        inboxAddr: deployments.contracts.Inbox.proxyAddress,
        rollupAddr: deployments.contracts.Rollup.proxyAddress,
        sequencerInboxAddr: deployments.contracts.SequencerInbox.proxyAddress,
        bridgeAddr: deployments.contracts.Bridge.proxyAddress,
        outboxV1: (deployments.contracts as any)['OldOutbox'].proxyAddress,
        outboxV2: (deployments.contracts as any)['OldOutbox'].proxyAddress, // CHRIS: TODO: v2 here?,
        rollupEventBridgeAddr:
          deployments.contracts.RollupEventBridge.proxyAddress,
      })

      await migrationManager.upgradeClassicContracts()
      await migrationManager.deployMigrator({
        bridgeAddr: constants.AddressZero,
        inboxTemplateAddr: constants.AddressZero,
        outboxAddr: constants.AddressZero,
        sequencerInboxAddr: constants.AddressZero,
      })
      assert.fail('Expected deploy to fail')
    } catch {}
  })

  it('step 1 fails if ownership not transferred', async () => {
    try {
      const provider = ethers.provider
      const deployments = await getDeployments(provider)
      const proxyAdminSigner = await getProxyAdminSigner(
        deployments.proxyAdminAddress
      )

      const migrationManager = new NitroMigrationManager(proxyAdminSigner, {
        proxyAdminAddr: deployments.proxyAdminAddress,
        inboxAddr: deployments.contracts.Inbox.proxyAddress,
        rollupAddr: deployments.contracts.Rollup.proxyAddress,
        sequencerInboxAddr: deployments.contracts.SequencerInbox.proxyAddress,
        bridgeAddr: deployments.contracts.Bridge.proxyAddress,
        outboxV1: (deployments.contracts as any)['OldOutbox'].proxyAddress,
        outboxV2: (deployments.contracts as any)['OldOutbox'].proxyAddress, // CHRIS: TODO: v2 here?,
        rollupEventBridgeAddr:
          deployments.contracts.RollupEventBridge.proxyAddress,
      })

      await migrationManager.upgradeClassicContracts()

      const rollupFac = await ethers.getContractFactory('Rollup')
      const prevRollup = await rollupFac.attach(
        deployments.contracts.Rollup.proxyAddress
      )
      const wasmModuleRoot =
        '0x9900000000000000000000000000000000000000000000000000000000000010'
      const loserStakeEscrow = constants.AddressZero
      const nitroContracts = await migrationManager.deployNitroContracts({
        confirmPeriodBlocks: await prevRollup.confirmPeriodBlocks(),
        extraChallengeTimeBlocks: await prevRollup.extraChallengeTimeBlocks(),
        stakeToken: await prevRollup.stakeToken(),
        baseStake: await prevRollup.baseStake(),
        wasmModuleRoot: wasmModuleRoot,
        // CHRIS: TODO: decide who the owner should be
        // CHRIS: TODO: shouldnt it be someone different to the proxy admin?
        owner: await prevRollup.owner(),
        chainId: (await provider.getNetwork()).chainId,
        loserStakeEscrow: loserStakeEscrow,
        sequencerInboxMaxTimeVariation: {
          // CHRIS: TODO: should we change this to the exact POS seconds? probably not yet, we can update it later i guess
          // CHRIS: TODO: make sure these are all the values we want
          delayBlocks: (60 * 60 * 24) / 15,
          futureBlocks: 12,
          delaySeconds: 60 * 60 * 24,
          futureSeconds: 60 * 60,
        },
      })

      await migrationManager.deployMigrator({
        bridgeAddr: nitroContracts.bridge,
        inboxTemplateAddr: nitroContracts.inboxTemplate,
        outboxAddr: nitroContracts.outbox,
        sequencerInboxAddr: nitroContracts.sequencerInbox,
      })

      const mainnetSequencer = '0xa4b10ac61E79Ea1e150DF70B8dda53391928fD14'
      await migrationManager.step1([mainnetSequencer], {
        bridgeAddr: nitroContracts.bridge,
        rollupAddr: nitroContracts.rollup,
      })

      assert.fail('Expected step 1 to fail')
    } catch {}
  })

  it.only('upgrade and construct', async () => {
    const provider = ethers.provider
    const deployments = await getDeployments(provider)
    const proxyAdminSigner = await getProxyAdminSigner(
      deployments.proxyAdminAddress
    )

    // CHRIS: TODO: should it be possible to reverse each of the steps? or is that going too far?

    // CHRIS: TODO: don't we need to create a new 'deployments' file?
    // CHRIS: TODO: shouldnt the bridge be upgraded? no, we're doing a fresh one

    const migrationManager = new NitroMigrationManager(proxyAdminSigner, {
      proxyAdminAddr: deployments.proxyAdminAddress,
      inboxAddr: deployments.contracts.Inbox.proxyAddress,
      rollupAddr: deployments.contracts.Rollup.proxyAddress,
      sequencerInboxAddr: deployments.contracts.SequencerInbox.proxyAddress,
      bridgeAddr: deployments.contracts.Bridge.proxyAddress,
      outboxV1: (deployments.contracts as any)['OldOutbox'].proxyAddress,
      outboxV2: (deployments.contracts as any)['OldOutbox'].proxyAddress, // CHRIS: TODO: v2 here?,
      rollupEventBridgeAddr:
        deployments.contracts.RollupEventBridge.proxyAddress,
    })

    const rollupFac = await ethers.getContractFactory('Rollup')
    // lookup params from previous rollup?
    // CHRIS: TODO: why do we have a param in the constructor? how is this rollup logic supposed to be deployed?
    const prevRollup = await rollupFac.attach(
      deployments.contracts.Rollup.proxyAddress
    )
    const wasmModuleRoot =
      '0x9900000000000000000000000000000000000000000000000000000000000010'
    const loserStakeEscrow = constants.AddressZero

    const nitroContracts = await migrationManager.deployNitroContracts({
      confirmPeriodBlocks: await prevRollup.confirmPeriodBlocks(),
      extraChallengeTimeBlocks: await prevRollup.extraChallengeTimeBlocks(),
      stakeToken: await prevRollup.stakeToken(),
      baseStake: await prevRollup.baseStake(),
      wasmModuleRoot: wasmModuleRoot,
      // CHRIS: TODO: decide who the owner should be
      // CHRIS: TODO: shouldnt it be someone different to the proxy admin?
      owner: await prevRollup.owner(),
      chainId: (await provider.getNetwork()).chainId,
      loserStakeEscrow: loserStakeEscrow,
      sequencerInboxMaxTimeVariation: {
        // CHRIS: TODO: should we change this to the exact POS seconds? probably not yet, we can update it later i guess
        // CHRIS: TODO: make sure these are all the values we want
        delayBlocks: (60 * 60 * 24) / 15,
        futureBlocks: 12,
        delaySeconds: 60 * 60 * 24,
        futureSeconds: 60 * 60,
      },
    })
    await migrationManager.upgradeClassicContracts()

    await migrationManager.deployMigrator({
      // CHRIS: TODO: we could do more in terms of checks
      // CHRIS: TODO: we could do a check that all the contracts we care about have been correctly deployed with the correct admins
      // CHRIS: TODO: we could also check that the contracts below have expected functions on them?
      bridgeAddr: nitroContracts.bridge,
      inboxTemplateAddr: nitroContracts.inboxTemplate,
      outboxAddr: nitroContracts.outbox,
      sequencerInboxAddr: nitroContracts.sequencerInbox,
    })

    await migrationManager.transferClassicOwnership()

    // CHRIS: TODO: remove this!!!! we only do this whilst we wait for a receive function to be added to the
    // set the classic bridge as a inbox on the nitro bridge
    await (
      await new NitroRollupAdminLogic__factory(proxyAdminSigner)
        .attach(nitroContracts.rollup)
        .setInbox(deployments.contracts.Bridge.proxyAddress, true)
    ).wait()

    // CHRIS: TODO: get the correct address here, dont hard code?
    const mainnetSequencer = '0xa4b10ac61E79Ea1e150DF70B8dda53391928fD14'
    await migrationManager.step1([mainnetSequencer], {
      rollupAddr: nitroContracts.rollup,
      bridgeAddr: nitroContracts.bridge,
    })

    //////// CHRIS /////// PUT BACK IN
    // rest the bridge
    // // CHRIS: TODO: remove this when we remove teh setInbox(true) above
    await (
      await await new NitroRollupAdminLogic__factory(proxyAdminSigner)
        .attach(nitroContracts.rollup)
        .setInbox(deployments.contracts.Bridge.proxyAddress, false)
    ).wait()

    // // step 2
    // // this would normally be the latest created node
    // // but we need to confirm all the nodes to ensure that
    // // CHRIS: TODO: use the admin to force confirm the nodes between
    // // latest created and latest confirmed
    const classicRollupAdminFac = new RollupAdminFacet__factory(
      proxyAdminSigner
    )
    const rollupAdmin = classicRollupAdminFac.attach(prevRollup.address)
    // const latestCreated = await rollupAdmin.latestNodeCreated()
    const latestConfirmed = await rollupAdmin.latestConfirmed()
    // console.log(
    //   'confirmed count',
    //   latestConfirmed.toNumber(),
    //   latestCreated.toNumber()
    // )
    // const stakerCount = await rollupAdmin.stakerCount()
    // console.log('staker count', stakerCount.toNumber())

    await migrationManager.step2(latestConfirmed, true, true)

    // const res = await (await nitroMigrator.nitroStep2(latestConfirmed, { gasLimit: 3000000})).wait()
    // console.log(res.gasUsed.toString())
    // console.log(res.logs)
    // console.log(Date.now() - beforeB)

    // console.log('step 2 complete')

    // // step 3
    await migrationManager.step3()

    // console.log('step 3 complete')

    //////// CHRIS /////// PUT BACK IN
  })
})
