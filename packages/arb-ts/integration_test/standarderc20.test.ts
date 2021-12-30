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

import { BigNumber } from '@ethersproject/bignumber'

import { TestERC20__factory } from '../src/lib/abi/factories/TestERC20__factory'

import { OutgoingMessageState } from '../src/lib/dataEntities'

import {
  fundL1,
  fundL2,
  testRetryableTicket,
  warn,
  instantiateBridgeWithRandomWallet,
  fundL2Token,
  tokenFundAmount,
  skipIfMainnet,
  existentTestERC20,
} from './testHelpers'
import { L2TransactionReceipt } from '../src/lib/message/L2ToL1Message'
import { TokenBridger } from '../src'
import { Signer } from 'ethers'

describe('standard ERC20', () => {
  beforeEach('skipIfMainnet', function () {
    skipIfMainnet(this)
  })

  it('deposits erc20 (no L2 Eth funding)', async () => {
    const {
      l1Signer,
      tokenBridger,
      l2Signer,
    } = await instantiateBridgeWithRandomWallet()
    await fundL1(l1Signer)
    await depositTokenTest(tokenBridger, l1Signer, l2Signer)
  })
  it.skip('deposits erc20 (with L2 Eth funding)', async () => {
    const {
      l1Signer,
      tokenBridger,
      l2Signer,
    } = await instantiateBridgeWithRandomWallet()
    await fundL1(l1Signer)
    await fundL2(l2Signer)
    await depositTokenTest(tokenBridger, l1Signer, l2Signer)
  })

  it('withdraws erc20', async function () {
    const tokenWithdrawAmount = BigNumber.from(1)
    const {
      l2Network,
      l1Signer,
      l2Signer,
      tokenBridger,
    } = await instantiateBridgeWithRandomWallet()
    await fundL2(l1Signer)
    const result = await fundL2Token(l2Signer, tokenBridger, existentTestERC20)
    if (!result) {
      warn('Prefunded wallet not funded with tokens; skipping ERC20 withdraw')
      this.skip()
    }

    const withdrawRes = await tokenBridger.withdraw({
      amount: tokenWithdrawAmount,
      erc20l1Address: existentTestERC20,
      l2Signer: l2Signer,
    })
    const withdrawRec = await withdrawRes.wait()

    expect(withdrawRec.status).to.equal(
      1,
      'token withdraw initiation txn failed'
    )

    const outgoingMessages = await withdrawRec.getL2ToL1Messages(
      l2Signer.provider
    )
    const firstMessage = outgoingMessages[0]
    expect(firstMessage, 'getWithdrawalsInL2Transaction came back empty').to
      .exist

    const outgoingMessageState = await firstMessage.status(null)

    expect(
      outgoingMessageState === OutgoingMessageState.UNCONFIRMED ||
        outgoingMessageState === OutgoingMessageState.NOT_FOUND,
      `standard token withdraw getOutGoingMessageState returned ${outgoingMessageState}`
    ).to.be.true

    const l2Token = tokenBridger.getL2TokenContract(
      l2Signer.provider,
      await tokenBridger.getL2ERC20Address(existentTestERC20, l1Signer.provider)
    )
    const testWalletL2Balance = (
      await l2Token.functions.balanceOf(await l2Signer.getAddress())
    )[0]

    expect(
      testWalletL2Balance.add(tokenWithdrawAmount).eq(tokenFundAmount),
      'token withdraw balance not deducted'
    ).to.be.true
    const walletAddress = await l1Signer.getAddress()

    const gatewayWithdrawEvents = await tokenBridger.getL2WithdrawalEvents(
      l2Signer.provider,
      l2Network.tokenBridge.l2ERC20Gateway,
      undefined,
      walletAddress,
      { fromBlock: withdrawRec.blockNumber }
    )
    expect(gatewayWithdrawEvents.length).to.equal(
      1,
      'token custom gateway query failed'
    )

    const gatewayAddress = await tokenBridger.getL2GatewayAddress(
      existentTestERC20,
      l2Signer.provider
    )
    const tokenWithdrawEvents = await tokenBridger.getL2WithdrawalEvents(
      l2Signer.provider,
      gatewayAddress,
      existentTestERC20,
      walletAddress,
      { fromBlock: withdrawRec.blockNumber }
    )
    expect(tokenWithdrawEvents.length).to.equal(
      1,
      'token filtered query failed'
    )
  })
  it('getERC20L1Address/getERC20L2Address work as expected', async () => {
    const {
      l1Signer,
      l2Signer,
      tokenBridger,
    } = await instantiateBridgeWithRandomWallet()
    const queriedL2Address = await tokenBridger.getL2ERC20Address(
      existentTestERC20,
      l1Signer.provider
    )
    const queriedL1Address = await tokenBridger.getL1ERC20Address(
      queriedL2Address,
      l2Signer.provider
    )

    expect(queriedL1Address).to.equal(
      existentTestERC20,
      'getERC20L1Address/getERC20L2Address failed with proper token address'
    )

    const randomAddress = await l1Signer.getAddress()
    const notAnL1Address = await tokenBridger.getL1ERC20Address(
      randomAddress,
      l2Signer.provider
    )
    expect(notAnL1Address).to.be.null,
      "getERC20L1Address didn't return null for random input"
  })
})

const depositTokenTest = async (
  tokenBridger: TokenBridger,
  l1Signer: Signer,
  l2Signer: Signer
) => {
  const tokenDepositAmount = BigNumber.from(1)

  const testToken = TestERC20__factory.connect(existentTestERC20, l1Signer)
  const mintRes = await testToken.mint()
  await mintRes.wait()

  const approveRes = await tokenBridger.approveToken({
    erc20L1Address: existentTestERC20,
    l1Signer: l1Signer,
  })
  await approveRes.wait()

  const expectedL1GatewayAddress = await tokenBridger.getL1GatewayAddress(
    testToken.address,
    l1Signer.provider
  )
  const l1Token = tokenBridger.getL1TokenContract(
    l1Signer.provider,
    existentTestERC20
  )
  const allowance = (
    await l1Token.functions.allowance(
      await l1Signer.getAddress(),
      expectedL1GatewayAddress
    )
  )[0]
  expect(allowance.eq(TokenBridger.MAX_APPROVAL), 'set token allowance failed')
    .to.be.true

  const initialBridgeTokenBalance = await testToken.balanceOf(
    expectedL1GatewayAddress
  )

  const depositRes = await tokenBridger.deposit({
    l1Signer: l1Signer,
    l2Provider: l2Signer.provider,
    erc20L1Address: existentTestERC20,
    amount: tokenDepositAmount,
  })

  const depositRec = await depositRes.wait()

  const finalBridgeTokenBalance = await testToken.balanceOf(
    expectedL1GatewayAddress
  )

  expect(
    initialBridgeTokenBalance
      .add(tokenDepositAmount)
      .eq(finalBridgeTokenBalance),
    'bridge balance not updated after L1 token deposit txn'
  ).to.be.true
  await testRetryableTicket(l1Signer.provider, depositRec)

  const l2Token = tokenBridger.getL2TokenContract(
    l2Signer.provider,
    await tokenBridger.getL2ERC20Address(existentTestERC20, l1Signer.provider)
  )
  const testWalletL2Balance = (
    await l2Token.functions.balanceOf(await l2Signer.getAddress())
  )[0]

  expect(
    testWalletL2Balance.eq(tokenDepositAmount),
    'l2 wallet not updated after deposit'
  ).to.be.true
}
