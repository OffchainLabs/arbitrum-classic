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

import { Bridge } from '../src/lib/bridge'
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
import { BridgeHelper, TokenBridger } from '../src'
import { Signer } from 'ethers'
import { Provider } from '@ethersproject/abstract-provider'

describe('Custom ERC20', () => {
  beforeEach('skipIfMainnet', function () {
    skipIfMainnet(this)
  })

  it('deposits erc20 (no L2 Eth funding)', async () => {
    const {
      bridge,
      l1Signer,
      l2Signer,
      tokenBridger,
    } = await instantiateBridgeWithRandomWallet()
    await fundL1(l2Signer)
    await depositTokenTest(bridge, tokenBridger, l1Signer, l2Signer.provider!)
  })
  it('deposits erc20 (with L2 Eth funding)', async () => {
    const {
      bridge,
      l1Signer,
      l2Signer,
      tokenBridger,
    } = await instantiateBridgeWithRandomWallet()
    await fundL1(l1Signer)
    await fundL2(l2Signer)
    await depositTokenTest(bridge, tokenBridger, l1Signer, l2Signer.provider!)
  })

  it('withdraws erc20', async function () {
    const tokenWithdrawAmount = BigNumber.from(1)
    const {
      bridge,
      l2Signer,
      l1Signer,
      tokenBridger,
    } = await instantiateBridgeWithRandomWallet()
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
      await new L2TransactionReceipt(withdrawRec).getL2ToL1Messages(
        l1Signer.provider
      )
    )[0]
    expect(message, 'withdrawEventData not found').to.exist

    const outgoingMessageState = await message.status(null)
    expect(
      outgoingMessageState === OutgoingMessageState.UNCONFIRMED ||
        outgoingMessageState === OutgoingMessageState.NOT_FOUND,
      `custom token withdraw getOutGoingMessageState returned ${outgoingMessageState}`
    ).to.be.true

    const l2Data = await bridge.l2Bridge.getL2TokenData(
      await tokenBridger.getERC20L2Address(
        existentTestCustomToken,
        l1Signer.provider
      )
    )
    const testWalletL2Balance = l2Data.balance
    expect(
      testWalletL2Balance.add(tokenWithdrawAmount).eq(tokenFundAmount),
      'wallet balance not properly deducted after withdraw'
    ).to.be.true
    const walletAddress = await l1Signer.getAddress()

    const gatewayWithdrawEvents = await bridge.getGatewayWithdrawEventData(
      tokenBridger.l2Network.tokenBridge.l2CustomGateway,
      walletAddress,
      { fromBlock: withdrawRec.blockNumber }
    )
    expect(gatewayWithdrawEvents.length).to.equal(
      1,
      'token getGatewayWithdrawEventData query failed'
    )

    const tokenWithdrawEvents = await bridge.getTokenWithdrawEventData(
      existentTestCustomToken,
      walletAddress,
      { fromBlock: withdrawRec.blockNumber }
    )
    expect(tokenWithdrawEvents.length).to.equal(
      1,
      'token getTokenWithdrawEventData query failed'
    )
  })
})

const depositTokenTest = async (
  bridge: Bridge,
  tokenBridger: TokenBridger,
  l1Signer: Signer,
  l2Provider: Provider
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
  // const approveRes = await bridge.approveToken(existentTestCustomToken)
  await approveRes.wait()

  const data = await bridge.l1Bridge.getL1TokenData(existentTestCustomToken)
  const allowed = data.allowed
  expect(allowed, 'set token allowance failed').to.be.true

  const expectedL1GatewayAddress = await tokenBridger.getL1GatewayAddress(
    testToken.address,
    l1Signer.provider
  )
  const initialBridgeTokenBalance = await testToken.balanceOf(
    expectedL1GatewayAddress
  )

  const depositRes = await tokenBridger.deposit({
    erc20L1Address: existentTestCustomToken,
    amount: tokenDepositAmount,
    l1Signer: l1Signer,
    l2Provider: l2Provider,
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
  await testRetryableTicket(l1Signer.provider, depositRec)

  const l2Data = await bridge.l2Bridge.getL2TokenData(
    await tokenBridger.getERC20L2Address(
      existentTestCustomToken,
      l1Signer.provider
    )
  )

  const testWalletL2Balance = l2Data.balance

  expect(
    testWalletL2Balance.eq(tokenDepositAmount),
    "l2 wallet balance not properly updated after deposit'"
  ).to.be.true
}
