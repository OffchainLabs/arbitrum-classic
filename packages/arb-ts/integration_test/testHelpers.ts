/*
 * Copyright 2021, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/* eslint-env node */
'use strict'

import { expect } from 'chai'
import yargs from 'yargs/yargs'
import chalk from 'chalk'

import { BigNumber } from '@ethersproject/bignumber'
import { Provider } from '@ethersproject/providers'
import { ContractReceipt } from '@ethersproject/contracts'
import { Wallet } from '@ethersproject/wallet'
import { formatBytes32String } from '@ethersproject/strings'
import { parseEther } from '@ethersproject/units'

import { TestERC20__factory } from '../src/lib/abi/factories/TestERC20__factory'

import { Bridge } from '../src/lib/bridge'
import { Network } from '../src/lib/networks'
import { instantiateBridge } from '../scripts/instantiate_bridge'

import config from './config'
import { L1TransactionReceipt } from '../src/lib/message/L1ToL2Message'
import { Signer } from 'ethers'
import { EthBridger, TokenBridger } from '../src'

const argv = yargs(process.argv.slice(2))
  .options({
    networkID: {
      type: 'string',
    },
  })
  .parseSync()

const networkID = (argv.networkID as '1' | '4' | '1337') || '4'
if (!config[networkID]) {
  throw new Error('network not supported')
}

const {
  existentTestERC20: _existentTestERC20,
  existentTestCustomToken: _existentTestCustomToken,
} = config[networkID]

export const existentTestERC20 = _existentTestERC20 as string
export const existentTestCustomToken = _existentTestCustomToken as string

export const preFundAmount = parseEther('0.001')

export const testRetryableTicket = async (
  l1Provider: Provider,
  rec: ContractReceipt
): Promise<void> => {
  prettyLog(`testing retryable for ${rec.transactionHash}`)

  const messages = await new L1TransactionReceipt(rec).getL1ToL2Messages(
    l1Provider
  )

  const message = messages && messages[0]
  if (!message) {
    throw new Error('Seq num not found')
  }
  // CHRIS: this mismatch in names adds to the confusion
  const retryableTicket = message.l2TicketCreationTxnHash
  const autoRedeem = message.autoRedeemHash
  const redeemTransaction = message.userTxnHash

  prettyLog(
    `retryableTicket: ${retryableTicket} autoredeem: ${autoRedeem}, redeem: ${redeemTransaction}`
  )
  prettyLog('Waiting for retryable ticket')

  const waitResult = await message.wait(1000 * 60 * 15)

  const retryableTicketReceipt = waitResult.ticketCreationReceipt

  prettyLog('retryableTicketReceipt found:')

  expect(retryableTicketReceipt.status).to.equal(
    1,
    'retryable ticket txn failed'
  )

  prettyLog(`Waiting for auto redeem transaction (this shouldn't take long`)
  const autoRedeemReceipt = waitResult.autoRedeemReceipt

  prettyLog('autoRedeem receipt found!')

  expect(autoRedeemReceipt.status).to.equal(1, 'autoredeem txn failed')
  prettyLog('Getting redemption')

  const redemptionReceipt = waitResult.userTxnReceipt

  expect(redemptionReceipt && redemptionReceipt.status).equals(
    1,
    'redeem txn failed'
  )
}

export const prettyLog = (text: string): void => {
  console.log(chalk.blue(`    *** ${text}`))
  console.log()
}

export const warn = (text: string): void => {
  console.log(chalk.red(`WARNING: ${text}`))
  console.log()
}

export const instantiateBridgeWithRandomWallet = (): Promise<{
  bridge: Bridge
  tokenBridger: TokenBridger
  ethBridger: EthBridger
  l1Network: Network,
  l2Network: Network,
  l1Signer: Signer
  l2Signer: Signer
}> => {
  const testPk = formatBytes32String(Math.random().toString())
  prettyLog(
    `Generated wallet, pk: ${testPk} address: ${new Wallet(testPk).address} `
  )
  return instantiateBridge(testPk, testPk)
}

const _preFundedWallet = new Wallet(process.env.DEVNET_PRIVKEY as string)
const _preFundedL2Wallet = new Wallet(process.env.DEVNET_PRIVKEY as string)

console.warn('using prefunded wallet ', _preFundedWallet.address)

export const fundL1 = async (l1Signer: Signer): Promise<void> => {
  const testWalletAddress = await l1Signer.getAddress()
  const preFundedWallet = _preFundedWallet.connect(l1Signer.provider!)
  const res = await preFundedWallet.sendTransaction({
    to: testWalletAddress,
    value: preFundAmount,
  })
  await res.wait()
  prettyLog('Funded L1 account')
}
export const fundL2 = async (l2Signer: Signer): Promise<void> => {
  const testWalletAddress = await l2Signer.getAddress()
  const preFundedL2Wallet = _preFundedL2Wallet.connect(l2Signer.provider!)
  const res = await preFundedL2Wallet.sendTransaction({
    to: testWalletAddress,
    value: preFundAmount,
  })
  await res.wait()
  prettyLog('Funded L2 account')
}

export const tokenFundAmount = BigNumber.from(2)
export const fundL2Token = async (
  l2Signer: Signer,
  tokenBridger: TokenBridger,
  tokenAddress: string
): Promise<boolean> => {
  try {
    const testWalletAddress = await l2Signer.getAddress()
    const preFundedL2Wallet = _preFundedL2Wallet.connect(l2Signer.provider)
    const l2Address = await tokenBridger.getERC20L2Address(
      tokenAddress,
      l2Signer.provider
    )
    const testToken = TestERC20__factory.connect(l2Address, preFundedL2Wallet)

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

export const wait = (ms = 0): Promise<undefined> => {
  return new Promise(res => setTimeout(res, ms))
}

export const skipIfMainnet = (() => {
  let chainId = ''
  return async (testContext: Mocha.Context) => {
    if (!chainId) {
      const { l1Network } = await instantiateBridgeWithRandomWallet()
      chainId = l1Network.chainID
    }
    if (chainId === '1') {
      console.log("You're writing to the chain on mainnet lol stop")
      testContext.skip()
    }
  }
})()
