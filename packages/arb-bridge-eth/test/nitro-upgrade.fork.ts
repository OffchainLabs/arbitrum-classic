import { ethers, network } from 'hardhat'
import { expect, assert } from 'chai'
import { CurrentDeployments } from 'arb-upgrades/types'
import { readFileSync, readdirSync } from 'fs'
import { Bridge, NitroMigrator__factory, ProxyAdmin, RollupAdminFacet__factory } from '../build/types'
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
import { Interface } from '@ethersproject/abi'
import { NitroMigrationManager } from './nitroMigrationManager'

describe('Nitro upgrade', () => {
  const getDeployments = async (provider: Provider) => {
    const chainId = (await provider.getNetwork()).chainId
    const deploymentData = readFileSync(
      `./_deployments/${chainId}_current_deployment.json`
    )
    return JSON.parse(deploymentData.toString()) as CurrentDeployments
  }

  it('fails to construct if contracts havent been upgraded', async () => {
    try {
      const deployments = await getDeployments(ethers.provider)
      const nitroMigratorFac = await ethers.getContractFactory('NitroMigrator')
      await nitroMigratorFac.deploy(
        deployments.contracts.Inbox.proxyAddress,
        deployments.contracts.SequencerInbox.proxyAddress,
        deployments.contracts.Bridge.proxyAddress,
        deployments.contracts.RollupEventBridge.proxyAddress,
        (deployments.contracts as any)['OldOutbox'].proxyAddress,
        (deployments.contracts as any)['OldOutbox'].proxyAddress, // CHRIS: TODO: v2 here?
        deployments.contracts.Rollup.proxyAddress,
        constants.AddressZero,
        constants.AddressZero,
        constants.AddressZero,
        constants.AddressZero
      )

      assert.fail('Expected constructor to fail')
    } catch {}

    // const delayedInboxAddress = deployments.contracts.Inbox.proxyAddress

    // const delayedInbox = await ethers.getContractAt(
    //   'Inbox',
    //   delayedInboxAddress
    // )

    // const bridge = await ethers.getContractAt(
    //   'Bridge',
    //   await delayedInbox.bridge()
    // )

    // const rollupProxyAddress = await bridge.owner()
    // const rollupDispatch = await ethers.getContractAt(
    //   'Rollup',
    //   rollupProxyAddress
    // )

    // const sequencerInboxAddress = await rollupDispatch.sequencerBridge()
    // const sequencerInbox = await ethers.getContractAt(
    //   'SequencerInbox',
    //   sequencerInboxAddress
    // )

    // // deploy new logic contracts
    // const NewRollupLogic = await ethers.getContractFactory('Rollup')
    // const newRollupLogic = await NewRollupLogic.deploy(1)
    // await newRollupLogic.deployed()

    // const NewAdminFacet = await ethers.getContractFactory('RollupAdminFacet')
    // const newAdminFacet = await NewAdminFacet.deploy()
    // await newAdminFacet.deployed()

    // const NewSequencerInbox = await ethers.getContractFactory('SequencerInbox')
    // const newSequencerInbox = await NewSequencerInbox.deploy()
    // await newSequencerInbox.deployed()

    // // valid previous rollup state
    // const iface = new ethers.utils.Interface([
    //   `function sequencerInboxMaxDelayBlocks() view returns (uint256)`,
    //   `function sequencerInboxMaxDelaySeconds() view returns (uint256)`,
    // ])
    // const oldRollupInterface = new ethers.Contract(rollupProxyAddress, iface)
    // const prevMaxDelayBlocks = await oldRollupInterface
    //   .connect(ethers.provider.getSigner())
    //   .sequencerInboxMaxDelayBlocks()
    // const prevMaxDelaySeconds = await oldRollupInterface
    //   .connect(ethers.provider.getSigner())
    //   .sequencerInboxMaxDelaySeconds()

    // // setup for fork test
    // const adminStorageSlot =
    //   '0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103'
    // const l1ProxyAdmin =
    //   '0x' +
    //   (
    //     await ethers.provider.getStorageAt(
    //       sequencerInbox.address,
    //       adminStorageSlot
    //     )
    //   ).substr(26)
    // const proxyAdmin = await ethers.getContractAt('ProxyAdmin', l1ProxyAdmin)
    // const owner = await proxyAdmin.owner()

    // // airdrop
    // await network.provider.send('hardhat_setBalance', [
    //   owner,
    //   '0x16189AD417E380000',
    // ])

    // // use owner
    // await network.provider.request({
    //   method: 'hardhat_impersonateAccount',
    //   params: [owner],
    // })

    // const ownerSigner = await ethers.provider.getSigner(owner)

    // // upgrade contracts

    // await proxyAdmin
    //   .connect(ownerSigner)
    //   .upgrade(sequencerInbox.address, newSequencerInbox.address)

    // const externalCall =
    //   rollupDispatch.interface.encodeFunctionData('postUpgradeInit')

    // await proxyAdmin
    //   .connect(ownerSigner)
    //   .upgradeAndCall(rollupProxyAddress, newRollupLogic.address, externalCall)

    // await network.provider.request({
    //   method: 'hardhat_stopImpersonatingAccount',
    //   params: [owner],
    // })

    // // verify storage was assigned correctly

    // const postMaxDelayBlocks = await sequencerInbox.maxDelayBlocks()
    // const postMaxDelaySeconds = await sequencerInbox.maxDelaySeconds()

    // const rollupMaxDelayBlocks = await rollupDispatch.STORAGE_GAP_1()
    // const rollupMaxDelaySeconds = await rollupDispatch.STORAGE_GAP_2()

    // expect(prevMaxDelayBlocks).to.equal(postMaxDelayBlocks)
    // expect(prevMaxDelaySeconds).to.equal(postMaxDelaySeconds)

    // expect(rollupMaxDelayBlocks).to.equal(0)
    // expect(rollupMaxDelaySeconds).to.equal(0)

    // // should not be able to call postUpgradeInit

    // const newerAdminFacet = await NewAdminFacet.deploy()
    // await newerAdminFacet.deployed()

    // await expect(rollupDispatch.postUpgradeInit()).to.be.revertedWith(
    //   'NOT_FROM_ADMIN'
    // )
  })

  it.only('upgrade and construct', async () => {
    const provider = ethers.provider
    const deployments = await getDeployments(provider)

    let proxyAdmin = await ethers.getContractAt(
      'ProxyAdmin',
      deployments.proxyAdminAddress
    )
    const proxyAdminOwner = await proxyAdmin.owner()

    // airdrop
    await network.provider.send('hardhat_setBalance', [
      proxyAdminOwner,
      '0x16189AD417E380000',
    ])

    // use owner
    await network.provider.request({
      method: 'hardhat_impersonateAccount',
      params: [proxyAdminOwner],
    })

    const proxyAdminSigner = await provider.getSigner(proxyAdminOwner)
    proxyAdmin = proxyAdmin.connect(proxyAdminSigner)
    // CHRIS: TODO: should it be possible to reverse each of the steps? or is that going too far?

    // CHRIS: TODO: don't we need to create a new 'deployments' file?
    // CHRIS: TODO: shouldnt the bridge be upgraded? no, we're doing a fresh one

    const migrationManager = new NitroMigrationManager(proxyAdminSigner)

    const rollupFac = await ethers.getContractFactory('Rollup')
    // lookup params from previous rollup?
    // CHRIS: TODO: why do we have a param in the constructor? how is this rollup logic supposed to be deployed?
    const prevRollup = await rollupFac.attach(
      deployments.contracts.Rollup.proxyAddress
    )
    const wasmModuleRoot =
      '0x9900000000000000000000000000000000000000000000000000000000000010'
    const loserStakeEscrow = constants.AddressZero

    const nitroContracts = await migrationManager.deployNitro({
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

    await migrationManager.upgradeClassicContracts({
      proxyAdminAddr: deployments.proxyAdminAddress,
      confirmPeriodBlocks: await prevRollup.confirmPeriodBlocks(),
      inboxAddr: deployments.contracts.Inbox.proxyAddress,
      rollupAddr: prevRollup.address,
      sequencerInboxAddr: deployments.contracts.SequencerInbox.proxyAddress,
    })

    await migrationManager.deployMigrator([
      deployments.contracts.Inbox.proxyAddress,
      deployments.contracts.SequencerInbox.proxyAddress,
      deployments.contracts.Bridge.proxyAddress,
      deployments.contracts.RollupEventBridge.proxyAddress,
      (deployments.contracts as any)['OldOutbox'].proxyAddress,
      (deployments.contracts as any)['OldOutbox'].proxyAddress, // CHRIS: TODO: v2 here?
      deployments.contracts.Rollup.proxyAddress,
      // CHRIS: TODO: we could do more in terms of checks
      // CHRIS: TODO: we could do a check that all the contracts we care about have been correctly deployed with the correct admins
      // CHRIS: TODO: we could also check that the contracts below have expected functions on them?
      nitroContracts.bridge,
      nitroContracts.inboxTemplate,
      nitroContracts.outbox,
      nitroContracts.sequencerInbox,
    ])

    await migrationManager.step1(
      {
        rollupAddr: deployments.contracts.Rollup.proxyAddress,
        bridgeAddr: deployments.contracts.Bridge.proxyAddress,
      },
      { rollupAddr: nitroContracts.rollup, bridgeAddr: nitroContracts.bridge }
    )

    //////// CHRIS /////// PUT BACK IN
    // // CHRIS: TODO: remove this when we remove teh setInbox(true) above
    const nitroRollupAdminFac = new NitroRollupAdminLogic__factory(proxyAdminSigner)
    const nitroRollupAdminFacet = nitroRollupAdminFac.attach(nitroContracts.rollup)
    await (
      await nitroRollupAdminFacet.setInbox(
        deployments.contracts.Bridge.proxyAddress,
        false
      )
    ).wait()

    // // step 2
    // // this would normally be the latest created node
    // // but we need to confirm all the nodes to ensure that
    // // CHRIS: TODO: use the admin to force confirm the nodes between
    // // latest created and latest confirmed
    const classicRollupAdminFac = new RollupAdminFacet__factory(proxyAdminSigner)
    const rollupAdmin = classicRollupAdminFac.attach(prevRollup.address);
    // const latestCreated = await rollupAdmin.latestNodeCreated()
    const latestConfirmed = await rollupAdmin.latestConfirmed()
    // console.log(
    //   'confirmed count',
    //   latestConfirmed.toNumber(),
    //   latestCreated.toNumber()
    // )
    // const stakerCount = await rollupAdmin.stakerCount()
    // console.log('staker count', stakerCount.toNumber())

    // const beforeB = Date.now()

    // const iFace = new Interface(['event Face(uint length)'])
    // console.log('Face topic', iFace.getEventTopic('Face'))

    // const setOwnerData = await rollupAdmin.interface.encodeFunctionData(
    //   'setOwner',
    //   [await proxyAdminSigner.getAddress()]
    // )
    // console.log("before exec")

    // await (
    //   await nitroMigrator.executeTransaction(
    //     setOwnerData,
    //     rollupAdmin.address,
    //     BigNumber.from(0)
    //   )
    // ).wait()
    // console.log("after exec")

    const res = await (
      await rollupAdmin.shutdownForNitro(latestConfirmed)
    ).wait()

    // const res = await (await nitroMigrator.nitroStep2(latestConfirmed, { gasLimit: 3000000})).wait()
    // console.log(res.gasUsed.toString())
    // console.log(res.logs)
    // console.log(Date.now() - beforeB)

    // console.log('step 2 complete')

    // // step 3
    // await (await nitroMigrator.nitroStep3()).wait()

    // console.log('step 3 complete')

    //////// CHRIS /////// PUT BACK IN
  })
})
