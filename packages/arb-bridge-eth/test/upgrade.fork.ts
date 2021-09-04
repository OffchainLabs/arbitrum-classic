import { ethers, deployments, run, network } from 'hardhat'
import { assert, expect } from 'chai'
import { networks } from 'arb-ts/src/lib/networks'
import { CurrentDeployments } from 'arb-upgrades/types'
import { writeFileSync, readFileSync, unlinkSync, existsSync } from 'fs'

describe('Mainnet fork', () => {
  it('should upgrade rollup contract correctly', async function () {
    const chainId = (await ethers.provider.getNetwork()).chainId

    const deploymentData = readFileSync(
      `../_deployments/${chainId}_current_deployment.json`
    )
    const deployments = JSON.parse(
      deploymentData.toString()
    ) as CurrentDeployments

    const delayedInboxAddress = deployments.contracts.Inbox.proxyAddress

    const delayedInbox = await ethers.getContractAt(
      'Inbox',
      delayedInboxAddress
    )
    const bridge = await ethers.getContractAt(
      'Bridge',
      await delayedInbox.bridge()
    )

    const rollupProxyAddress = await bridge.owner()
    const rollupDispatch = await ethers.getContractAt(
      'Rollup',
      rollupProxyAddress
    )

    const sequencerInboxAddress = await rollupDispatch.sequencerBridge()
    const sequencerInbox = await ethers.getContractAt(
      'SequencerInbox',
      sequencerInboxAddress
    )

    // deploy new logic contracts
    const NewRollupLogic = await ethers.getContractFactory('Rollup')
    const newRollupLogic = await NewRollupLogic.deploy(1)
    await newRollupLogic.deployed()

    const NewAdminFacet = await ethers.getContractFactory('RollupAdminFacet')
    const newAdminFacet = await NewAdminFacet.deploy()
    await newAdminFacet.deployed()

    const NewSequencerInbox = await ethers.getContractFactory('SequencerInbox')
    const newSequencerInbox = await NewSequencerInbox.deploy()
    await newSequencerInbox.deployed()

    // valid previous rollup state
    const iface = new ethers.utils.Interface([
      `function sequencerInboxMaxDelayBlocks() view returns (uint256)`,
      `function sequencerInboxMaxDelaySeconds() view returns (uint256)`,
    ])
    const oldRollupInterface = new ethers.Contract(rollupProxyAddress, iface)
    const prevMaxDelayBlocks = await oldRollupInterface
      .connect(ethers.provider.getSigner())
      .sequencerInboxMaxDelayBlocks()
    const prevMaxDelaySeconds = await oldRollupInterface
      .connect(ethers.provider.getSigner())
      .sequencerInboxMaxDelaySeconds()

    // setup for fork test
    const adminStorageSlot =
      '0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103'
    const l1ProxyAdmin =
      '0x' +
      (
        await ethers.provider.getStorageAt(
          sequencerInbox.address,
          adminStorageSlot
        )
      ).substr(26)
    const proxyAdmin = await ethers.getContractAt('ProxyAdmin', l1ProxyAdmin)
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

    const ownerSigner = await ethers.provider.getSigner(owner)

    // upgrade contracts

    await proxyAdmin
      .connect(ownerSigner)
      .upgrade(sequencerInbox.address, newSequencerInbox.address)

    const externalCall = rollupDispatch.interface.encodeFunctionData(
      'postUpgradeInit',
      [newAdminFacet.address]
    )

    await proxyAdmin
      .connect(ownerSigner)
      .upgradeAndCall(rollupProxyAddress, newRollupLogic.address, externalCall)

    await network.provider.request({
      method: 'hardhat_stopImpersonatingAccount',
      params: [owner],
    })

    // verify storage was assigned correctly

    const postMaxDelayBlocks = await sequencerInbox.maxDelayBlocks()
    const postMaxDelaySeconds = await sequencerInbox.maxDelaySeconds()

    const rollupMaxDelayBlocks = await rollupDispatch.STORAGE_GAP_1()
    const rollupMaxDelaySeconds = await rollupDispatch.STORAGE_GAP_2()

    expect(prevMaxDelayBlocks).to.equal(postMaxDelayBlocks)
    expect(prevMaxDelaySeconds).to.equal(postMaxDelaySeconds)

    expect(rollupMaxDelayBlocks).to.equal(0)
    expect(rollupMaxDelaySeconds).to.equal(0)

    // should not be able to call postUpgradeInit

    const newerAdminFacet = await NewAdminFacet.deploy()
    await newerAdminFacet.deployed()

    await expect(
      rollupDispatch.postUpgradeInit(newerAdminFacet.address)
    ).to.be.revertedWith('NOT_FROM_ADMIN')
  })
})
