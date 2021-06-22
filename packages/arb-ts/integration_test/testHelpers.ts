import { Bridge } from '../src/lib/bridge'
import { expect } from 'chai'
import { ContractReceipt } from 'ethers'
import chalk from 'chalk'

export const testRetryableTicket = async (
  bridge: Bridge,
  rec: ContractReceipt
) => {
  prettyLog(`testing retryable for ${rec.transactionHash}`)

  const seqNums = await bridge.getInboxSeqNumFromContractTransaction(rec)
  const seqNum = seqNums && seqNums[0]
  if (!seqNum) {
    throw new Error('Seq num not found')
  }
  const retryableTicket = await bridge.calculateL2TransactionHash(seqNum)
  const autoRedeem = await bridge.calculateRetryableAutoRedeemTxnHash(seqNum)
  const redeemTransaction = await bridge.calculateL2RetryableTransactionHash(
    seqNum
  )
  console.log(
    `retryableTicket: ${retryableTicket} autoredeem: ${autoRedeem}, redeem: ${redeemTransaction}`
  )
  prettyLog('Waiting for retryable ticket')

  const retryableTicketReceipt = await bridge.l2Bridge.l2Provider.waitForTransaction(
    retryableTicket,
    undefined,
    1000 * 60 * 15
  )

  prettyLog('retryableTicketReceipt found:')

  expect(retryableTicketReceipt.status).to.equal(1)

  prettyLog(`Waiting for auto redeem transaction (this shouldn't take long`)
  const autoRedeemReceipt = await bridge.l2Bridge.l2Provider.waitForTransaction(
    autoRedeem,
    1000 * 60
  )
  prettyLog('autoRedeem receipt found!')

  expect(autoRedeemReceipt.status).to.equal(1)
  prettyLog('Getting redemption')
  const redemptionReceipt = await bridge.l2Bridge.l2Provider.getTransactionReceipt(
    redeemTransaction
  )

  expect(redemptionReceipt && redemptionReceipt.status).equals(1)
}

export const prettyLog = (text: string) => {
  console.log(chalk.blue(`    *** ${text}`))
  console.log()
}

export const warn = (text: string) => {
  console.log(chalk.red(`WARNING: ${text}`))
  console.log()
}
