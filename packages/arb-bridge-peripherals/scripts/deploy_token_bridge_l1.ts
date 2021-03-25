import { BigNumber } from '@ethersproject/bignumber'
import { concat, id, keccak256, zeroPad } from 'ethers/lib/utils'
import { ethers } from 'hardhat'
import deployments from '../deployment.json'

import { Bridge } from 'arb-ts/src'
import { writeFileSync } from 'fs'

const main = async () => {
  const accounts = await ethers.getSigners()
  // TODO: check buddy deployer address available
  // TODO: check 1820 registry
  const inboxAddress =
    process.env.INBOX_ADDRESS || '0x0d0c1aDf6523D422ec7192506A7F6790253399fB'

  if (inboxAddress === '' || inboxAddress === undefined)
    throw new Error('Please set inbox address! INBOX_ADDRESS')

  console.log('deployer', accounts[0].address)

  const EthERC20Bridge = await ethers.getContractFactory('EthERC20Bridge')

  const maxSubmissionCost = 0
  const gasPrice = 0
  const maxGas = 100000000000
  const ethERC20Bridge = await EthERC20Bridge.deploy(
    inboxAddress,
    deployments.buddyDeployer,
    maxSubmissionCost,
    maxGas,
    gasPrice,
    deployments.standardArbERC777,
    deployments.standardArbERC20
  )

  await ethERC20Bridge.deployed()

  const arbTokenBridge = await ethERC20Bridge.l2Buddy()
  console.log('EthERC20Bridge deployed to:', ethERC20Bridge.address)
  console.log('L2 ArbBridge deployed to:', arbTokenBridge)

  const contracts = JSON.stringify({
    ...deployments,
    ethERC20Bridge: ethERC20Bridge.address,
    arbTokenBridge: arbTokenBridge,
  })
  const deployFilePath = './deployment.json'
  console.log(`Writing to JSON at ${deployFilePath}`)
  writeFileSync(deployFilePath, contracts)

  const l2Provider = new ethers.providers.JsonRpcProvider(
    process.env.L2_RPC_URL || 'https://kovan4.arbitrum.io/rpc'
  )
  const l2PrivKey = process.env['L2_PRIVKEY']
  if (!l2PrivKey) throw new Error('Missing l2 priv key')
  const l2Signer = new ethers.Wallet(l2PrivKey, l2Provider)

  const bridge = new Bridge(
    ethERC20Bridge.address,
    arbTokenBridge,
    accounts[0],
    l2Signer
  )

  const deployReceipt = await bridge.getL1Transaction(
    ethERC20Bridge.deployTransaction.hash
  )

  const seqNums = await bridge.getInboxSeqNumFromContractTransaction(
    deployReceipt
  )

  if (!seqNums) throw new Error("Transaction didn't trigger inbox")
  if (seqNums.length !== 1)
    throw new Error('Transaction triggered inbox more than once')

  const inboxSequenceNumber = seqNums[0]

  const l2DeployTxHash = await bridge.calculateL2RetryableTransactionHash(
    inboxSequenceNumber
  )
  const l2TransactionReceipt = await bridge.getL2Transaction(l2DeployTxHash)

  const buddyDeployEvents = await bridge.getBuddyDeployInL2Transaction(
    l2TransactionReceipt
  )

  if (buddyDeployEvents.length !== 1)
    throw new Error('Buddy deploy event was not triggered one time!')
  const withdrawalId = buddyDeployEvents[0].withdrawalId

  const logs = await bridge.getWithdrawalsInL2Transaction(l2TransactionReceipt)
  const filteredLogs = logs.filter(log => log.uniqueId.eq(withdrawalId))

  if (filteredLogs.length !== 1)
    throw new Error('Should have exactly one matching unique id')
  const { batchNumber, indexInBatch } = filteredLogs[0]

  const l1TxReceipt = await bridge.triggerL2ToL1Transaction(
    batchNumber,
    indexInBatch
  )
  console.log('Transaction executed in L1')
  console.log(l1TxReceipt)
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
