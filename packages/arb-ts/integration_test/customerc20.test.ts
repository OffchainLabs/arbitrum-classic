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
  existentTestCustomToken,
} from './testHelpers'
import { L2TransactionReceipt } from '../src/lib/message/L2ToL1Message'
import { TokenBridger } from '../src'
import { Signer } from 'ethers'

describe('Custom ERC20', () => {
  beforeEach('skipIfMainnet', function () {
    skipIfMainnet(this)
  })

  it('deposits erc20 (no L2 Eth funding)', async () => {
    const { l1Signer, l2Signer, tokenBridger } =
      await instantiateBridgeWithRandomWallet()
    await fundL1(l2Signer)
    await depositTokenTest(tokenBridger, l1Signer, l2Signer)
  })
  it.skip('deposits erc20 (with L2 Eth funding)', async () => {
    const { l1Signer, l2Signer, tokenBridger } =
      await instantiateBridgeWithRandomWallet()
    await fundL1(l1Signer)
    await fundL2(l2Signer)
    await depositTokenTest(tokenBridger, l1Signer, l2Signer)
  })

  it('withdraws erc20', async function () {
    const tokenWithdrawAmount = BigNumber.from(1)
    const { l2Network, l2Signer, l1Signer, tokenBridger } =
      await instantiateBridgeWithRandomWallet()
    await fundL2(l2Signer)
    const result = await fundL2Token(
      l2Signer,
      tokenBridger,
      existentTestCustomToken
    )
    if (!result) {
      warn('Prefunded wallet not funded with tokens; skipping ERC20 withdraw')
      this.skip()
    }

    const withdrawRes = await tokenBridger.withdraw({
      amount: tokenWithdrawAmount,
      erc20l1Address: existentTestCustomToken,
      l2Signer: l2Signer,
    })
    const withdrawRec = await withdrawRes.wait()

    expect(withdrawRec.status).to.equal(1, 'initiate token withdraw txn failed')

    const message = (
      await withdrawRec.getL2ToL1Messages(l1Signer.provider!, l2Network)
    )[0]
    expect(message, 'withdrawEventData not found').to.exist

    const outgoingMessageState = await message.status(null)
    expect(
      outgoingMessageState === OutgoingMessageState.UNCONFIRMED ||
        outgoingMessageState === OutgoingMessageState.NOT_FOUND,
      `custom token withdraw getOutGoingMessageState returned ${outgoingMessageState}`
    ).to.be.true

    const l2Token = tokenBridger.getL2TokenContract(
      l2Signer.provider!,
      await tokenBridger.getL2ERC20Address(
        existentTestCustomToken,
        l1Signer.provider!
      )
    )
    const testWalletL2Balance = (
      await l2Token.functions.balanceOf(await l2Signer.getAddress())
    )[0]
    expect(
      testWalletL2Balance.add(tokenWithdrawAmount).eq(tokenFundAmount),
      'wallet balance not properly deducted after withdraw'
    ).to.be.true
    const walletAddress = await l1Signer.getAddress()

    const gatewayWithdrawEvents = await tokenBridger.getL2WithdrawalEvents(
      l2Signer.provider!,
      tokenBridger.l2Network.tokenBridge.l2CustomGateway,
      undefined,
      walletAddress,
      { fromBlock: withdrawRec.blockNumber }
    )
    expect(gatewayWithdrawEvents.length).to.equal(
      1,
      'token standard gateway query failed'
    )

    const gatewayAddress = await tokenBridger.getL2GatewayAddress(
      existentTestCustomToken,
      l2Signer.provider!
    )
    const tokenWithdrawEvents = await tokenBridger.getL2WithdrawalEvents(
      l2Signer.provider!,
      gatewayAddress,
      existentTestCustomToken,
      walletAddress,
      { fromBlock: withdrawRec.blockNumber }
    )
    expect(tokenWithdrawEvents.length).to.equal(
      1,
      'token filtered query failed'
    )
  })
})

const depositTokenTest = async (
  tokenBridger: TokenBridger,
  l1Signer: Signer,
  l2Signer: Signer
) => {
  const tokenDepositAmount = BigNumber.from(1)

  const testToken = TestERC20__factory.connect(
    existentTestCustomToken,
    l1Signer
  )
  const mintRes = await testToken.mint()
  await mintRes.wait()

  const approveRes = await tokenBridger.approveToken({
    erc20L1Address: existentTestCustomToken,
    l1Signer: l1Signer,
  })
  await approveRes.wait()

  const expectedL1GatewayAddress = await tokenBridger.getL1GatewayAddress(
    testToken.address,
    l1Signer.provider!
  )
  const l1Token = tokenBridger.getL1TokenContract(
    l1Signer.provider!,
    existentTestCustomToken
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
    erc20L1Address: existentTestCustomToken,
    amount: tokenDepositAmount,
    l1Signer: l1Signer,
    l2Provider: l2Signer.provider!,
  })

  const depositRec = await depositRes.wait()

  const finalBridgeTokenBalance = await testToken.balanceOf(
    expectedL1GatewayAddress
  )

  expect(
    initialBridgeTokenBalance
      .add(tokenDepositAmount)
      .eq(finalBridgeTokenBalance),
    'bridge balance not properly updated after deposit'
  ).to.be.true
  await testRetryableTicket(l1Signer.provider!, depositRec)

  const l2Token = tokenBridger.getL2TokenContract(
    l2Signer.provider!,
    await tokenBridger.getL2ERC20Address(
      existentTestCustomToken,
      l1Signer.provider!
    )
  )
  const testWalletL2Balance = (
    await l2Token.functions.balanceOf(await l2Signer.getAddress())
  )[0]

  expect(
    testWalletL2Balance.eq(tokenDepositAmount),
    "l2 wallet balance not properly updated after deposit'"
  ).to.be.true
}
