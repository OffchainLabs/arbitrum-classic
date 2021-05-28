import { ethers } from 'hardhat'
import deployments from '../deployment.json'
import { BridgeHelper } from 'arb-ts/src/lib/bridge_helpers'
import { providers, Signer } from 'ethers'
import {
  Bridge,
  ArbTokenBridge__factory,
  EthERC20Bridge__factory,
} from 'arb-ts/src'
import { writeFileSync } from 'fs'
import { spawnSync } from 'child_process'

const main = async () => {
  const accounts = await ethers.getSigners()

  const inboxAddress = process.env.INBOX_ADDRESS

  if (!inboxAddress) throw new Error('Please set inbox address! INBOX_ADDRESS')

  const whitelistAddress = process.env.WHITELIST_ADDRESS
  if (!whitelistAddress)
    throw new Error('Please set whitelist address! WHITELIST_ADDRESS')

  const EthERC20Bridge = await ethers.getContractFactory('EthERC20Bridge')

  if (deployments.buddyDeployer === '' || deployments.standardArbERC20 === '')
    throw new Error("Deployments.json doesn't include the necessary addresses")

  const maxSubmissionCost = 0
  const gasPrice = 0
  const maxGas = 100000000000
  const ethERC20Bridge = await EthERC20Bridge.deploy()

  console.log('EthERC20Bridge logic deployed to:', ethERC20Bridge.address)
  const l2Provider = new providers.JsonRpcProvider(
    // 'https://kovan4.arbitrum.io/rpc'
    // 'https://devnet-l2.arbitrum.io/rpc'
    'https://kovan5.arbitrum.io/rpc'
  )
  const l2PrivKey = process.env['DEVNET_PRIVKEY']
  if (!l2PrivKey) throw new Error('Missing l2 priv key')
  const l2Signer = new ethers.Wallet(l2PrivKey, l2Provider)

  const ArbTokenBridge = (
    await ethers.getContractFactory('ArbTokenBridge')
  ).connect(l2Signer)

  const arbTokenBridge = await ArbTokenBridge.deploy()
  console.log('L2 ArbBridge logic deployed to:', arbTokenBridge.address)
  await arbTokenBridge.deployed()

  const L1TransparentUpgradeableProxy = await ethers.getContractFactory(
    'TransparentUpgradeableProxy'
  )
  const L2TransparentUpgradeableProxy = (
    await ethers.getContractFactory('TransparentUpgradeableProxy')
  ).connect(l2Signer)

  const L1ProxyAdmin = await ethers.getContractFactory('ProxyAdmin')
  const L2ProxyAdmin = (await ethers.getContractFactory('ProxyAdmin')).connect(
    l2Signer
  )
  console.log('Deploying l1ProxyAdmin:')

  const l1ProxyAdmin = await L1ProxyAdmin.deploy()
  console.log('L1 proxy admin at', l1ProxyAdmin.address)
  await l1ProxyAdmin.deployed()

  const ethERC20BridgeProxy = await L1TransparentUpgradeableProxy.deploy(
    ethERC20Bridge.address,
    l1ProxyAdmin.address,
    '0x'
  )
  await ethERC20BridgeProxy.deployed()

  console.log('L1 proxy bridge at', ethERC20BridgeProxy.address)

  const l2ProxyAdmin = await L2ProxyAdmin.deploy()
  await l2ProxyAdmin.deployed()

  console.log('L2 proxy admin at', l2ProxyAdmin.address)
  const arbTokenBridgeProxy = await L2TransparentUpgradeableProxy.deploy(
    arbTokenBridge.address,
    l2ProxyAdmin.address,
    '0x'
  )
  await arbTokenBridgeProxy.deployed()

  console.log('L2 proxy bridge at', arbTokenBridgeProxy.address)

  console.log('Now initializing proxies')

  const arbTokenBridgeConnectedAsProxy = ArbTokenBridge__factory.connect(
    arbTokenBridgeProxy.address,
    l2Signer
  )

  const initL2Bridge = await arbTokenBridgeConnectedAsProxy.initialize(
    ethERC20BridgeProxy.address,
    deployments.standardArbERC20
  )

  const ethERC20BridgeConnectedAsProxy = EthERC20Bridge__factory.connect(
    ethERC20BridgeProxy.address,
    accounts[0]
  )

  const initL1Bridge = await ethERC20BridgeConnectedAsProxy.initialize(
    inboxAddress,
    deployments.standardArbERC20,
    arbTokenBridgeProxy.address,
    l2Signer.address,
    whitelistAddress
  )
  console.log('init L1 hash', initL1Bridge.hash)
  console.log('init L2 hash', initL2Bridge.hash)
  // wait for inits
  await initL1Bridge.wait()
  console.warn('l1 bridge proxy initted')

  await initL2Bridge.wait()
  console.warn('l2 bridge proxy initted')

  // console.log("inbox after init", await ethERC20BridgeConnectedAsProxy.inbox())
  console.log('Proxies have been initted')

  const contracts = JSON.stringify({
    ...deployments,
    ethERC20Bridge: ethERC20BridgeProxy.address,
    arbTokenBridge: arbTokenBridgeProxy.address,
    inbox: inboxAddress,
  })
  const deployFilePath = './deployment.json'
  console.log(`Writing to JSON at ${deployFilePath}`)
  writeFileSync(deployFilePath, contracts)
  console.log('Wrote to deployments.json')

  // const bridge = new Bridge(
  //   ethERC20BridgeProxy.address,
  //   arbTokenBridgeProxy.address,
  //   accounts[0],
  //   l2Signer
  // )

  // const deployReceipt = await bridge.getL1Transaction(
  //   ethERC20Bridge.deployTransaction.hash
  // )

  // const seqNums = await bridge.getInboxSeqNumFromContractTransaction(
  //   deployReceipt
  // )

  // if (!seqNums) throw new Error("Transaction didn't trigger inbox")
  // if (seqNums.length !== 1)
  //   throw new Error('Transaction triggered inbox more than once')

  // const inboxSequenceNumber = seqNums[0]

  // const l2DeployTxHash = await bridge.calculateL2RetryableTransactionHash(
  //   inboxSequenceNumber
  // )
  // const l2TransactionReceipt = await bridge.getL2Transaction(l2DeployTxHash)

  // const buddyDeployEvents = await bridge.getBuddyDeployInL2Transaction(
  //   l2TransactionReceipt
  // )

  // if (buddyDeployEvents.length !== 1)
  //   throw new Error('Buddy deploy event was not triggered one time!')
  // const withdrawalId = buddyDeployEvents[0].withdrawalId

  // const logs = await bridge.getWithdrawalsInL2Transaction(l2TransactionReceipt)
  // const filteredLogs = logs.filter(log => log.uniqueId.eq(withdrawalId))

  // if (filteredLogs.length !== 1)
  //   throw new Error('Should have exactly one matching unique id')
  // const { batchNumber, indexInBatch } = filteredLogs[0]

  // const l1TxReceipt = await bridge.triggerL2ToL1Transaction(
  //   batchNumber,
  //   indexInBatch
  // )
  // console.log('Transaction executed in L1')
  // console.log(l1TxReceipt)
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
