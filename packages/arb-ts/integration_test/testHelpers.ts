import { Bridge } from '../src/lib/bridge'
import { expect } from 'chai'
import { BigNumber, ContractReceipt, Wallet } from 'ethers'
import chalk from 'chalk'
import { instantiateBridge } from '../scripts/instantiate_bridge'
import { utils } from 'ethers'
import { TestERC20__factory } from '../src/lib/abi/factories/TestERC20__factory'
import yargs from 'yargs/yargs'
import config from './config'

const argv = yargs(process.argv.slice(2)).argv
let networkID = argv.networkID as string

networkID = networkID || '4'
if (!config[networkID]) {
  throw new Error('network not supported')
}

const { existentTestERC20: _existentTestERC20 } = config[networkID]

export const existentTestERC20 = _existentTestERC20

export const preFundAmount = utils.parseEther('0.001')

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
  prettyLog(
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
    undefined,
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

export const instantiateRandomBridge = () => {
  const testPk = utils.formatBytes32String(Math.random().toString())
  return instantiateBridge(testPk)
}

const _preFundedWallet = new Wallet(process.env.DEVNET_PRIVKEY as string)
const _preFundedL2Wallet = new Wallet(process.env.DEVNET_PRIVKEY as string)

console.warn('using prefunded wallet ', _preFundedWallet.address)

export const fundL1 = async (bridge: Bridge) => {
  const testWalletAddress = await bridge.l1Bridge.getWalletAddress()
  const preFundedWallet = _preFundedWallet.connect(bridge.l1Provider)
  const res = await preFundedWallet.sendTransaction({
    to: testWalletAddress,
    value: preFundAmount,
  })
  const rec = await res.wait()
  prettyLog('Funded L1 account')
}
export const fundL2 = async (bridge: Bridge) => {
  const testWalletAddress = await bridge.l2Bridge.getWalletAddress()
  const preFundedL2Wallet = _preFundedL2Wallet.connect(bridge.l2Provider)
  const res = await preFundedL2Wallet.sendTransaction({
    to: testWalletAddress,
    value: preFundAmount,
  })
  const rec = await res.wait()
  prettyLog('Funded L2 account')
}

export const tokenFundAmount = BigNumber.from(2)
export const fundL2Token = async (bridge: Bridge) => {
  try {
    const testWalletAddress = await bridge.l2Bridge.getWalletAddress()
    const preFundedL2Wallet = _preFundedL2Wallet.connect(bridge.l2Provider)
    const l2Address = await bridge.getERC20L2Address(existentTestERC20)
    const testToken = TestERC20__factory.connect(l2Address, preFundedL2Wallet)
    const x = await testToken.balanceOf(preFundedL2Wallet.address)

    const res = await testToken.transfer(testWalletAddress, tokenFundAmount)

    const rec = await res.wait()
    const result = rec.status === 1
    result && prettyLog('Funded L2 account w/ tokens')

    return result
  } catch (err) {
    console.warn('err', err)

    return false
  }
}

export const wait = (ms = 0) => {
  return new Promise(res => setTimeout(res, ms))
}
