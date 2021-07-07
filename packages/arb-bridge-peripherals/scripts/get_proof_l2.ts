import { BigNumber } from '@ethersproject/bignumber'
import { ethers } from 'hardhat'
import { BridgeHelper } from 'arb-ts/src/lib/bridge_helpers'

const main = async () => {
  if (process.env.L2_TX_HASH && process.env.INBOX_SEQ_NUM)
    throw new Error('Either supply the L1 inbox seq num or l2 tx hash')
  const l2TxHash =
    process.env.L2_TX_HASH ||
    (await BridgeHelper.calculateL2RetryableTransactionHash(
      BigNumber.from(process.env.INBOX_SEQ_NUM),
      ethers.provider
    ))
  const l2TransactionReceipt = await BridgeHelper.getL2Transaction(
    l2TxHash,
    ethers.provider
  )

  const buddyDeployEvents = await BridgeHelper.getBuddyDeployInL2Transaction(
    l2TransactionReceipt
  )

  if (buddyDeployEvents.length !== 1)
    throw new Error('Buddy deploy event was not triggered one time!')
  const withdrawalId = buddyDeployEvents[0].withdrawalId

  const logs = BridgeHelper.getWithdrawalsInL2Transaction(
    l2TransactionReceipt,
    ethers.provider
  )
  const filteredLogs = logs.filter(log => log.uniqueId.eq(withdrawalId))

  if (filteredLogs.length !== 1)
    throw new Error('Should have exactly one matching unique id')
  const { batchNumber, indexInBatch } = filteredLogs[0]

  const proofData = await BridgeHelper.tryGetProof(
    batchNumber,
    indexInBatch,
    ethers.provider
  )
  console.log(JSON.stringify(proofData))
  //   return proofData;
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
