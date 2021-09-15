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
import { OutgoingMessageState } from '../src/lib/bridge_helpers'

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

describe('standard ERC20', () => {
  beforeEach('skipIfMainnet', function () {
    skipIfMainnet(this)
  })

  it('deposits erc20 (no L2 Eth funding)', async () => {
    const { bridge } = await instantiateBridgeWithRandomWallet()
    await fundL1(bridge)
    await depositTokenTest(bridge)
  })
  it.skip('deposits erc20 (with L2 Eth funding)', async () => {
    const { bridge } = await instantiateBridgeWithRandomWallet()
    await fundL1(bridge)
    await fundL2(bridge)
    await depositTokenTest(bridge)
  })

  it('withdraws erc20', async function () {
    const tokenWithdrawAmount = BigNumber.from(1)
    const { bridge, l2Network } = await instantiateBridgeWithRandomWallet()
    await fundL2(bridge)
    const result = await fundL2Token(bridge, existentTestERC20)
    if (!result) {
      warn('Prefunded wallet not funded with tokens; skipping ERC20 withdraw')
      this.skip()
    }
    const withdrawRes = await bridge.withdrawERC20(
      existentTestERC20,
      tokenWithdrawAmount
    )
    const withdrawRec = await withdrawRes.wait()

    expect(withdrawRec.status).to.equal(
      1,
      'token withdraw initiation txn failed'
    )
    const withdrawEventData =
      bridge.getWithdrawalsInL2Transaction(withdrawRec)[0]

    expect(
      withdrawEventData,
      'token withdraw getWithdrawalsInL2Transaction came back empty'
    ).to.exist

    const outgoingMessageState = await bridge.getOutGoingMessageState(
      withdrawEventData.batchNumber,
      withdrawEventData.indexInBatch
    )

    expect(
      outgoingMessageState === OutgoingMessageState.UNCONFIRMED ||
        outgoingMessageState === OutgoingMessageState.NOT_FOUND,
      `standard token withdraw getOutGoingMessageState returned ${outgoingMessageState}`
    ).to.be.true

    const l2Data = await bridge.getAndUpdateL2TokenData(existentTestERC20)
    const testWalletL2Balance = l2Data && l2Data.ERC20 && l2Data.ERC20.balance
    expect(
      testWalletL2Balance &&
        testWalletL2Balance.add(tokenWithdrawAmount).eq(tokenFundAmount),
      'token withdraw balance not deducted'
    ).to.be.true
    const walletAddress = await bridge.l1Signer.getAddress()

    const gatewayWithdrawEventData = await bridge.getGatewayWithdrawEventData(
      l2Network.tokenBridge.l2ERC20Gateway,
      walletAddress,
      { fromBlock: withdrawRec.blockNumber }
    )
    expect(gatewayWithdrawEventData.length).to.equal(
      1,
      'token withdraw getGatewayWithdrawEventData query failed'
    )

    const tokenWithdrawEvents = await bridge.getTokenWithdrawEventData(
      existentTestERC20,
      walletAddress,
      { fromBlock: withdrawRec.blockNumber }
    )
    expect(tokenWithdrawEvents.length).to.equal(
      1,
      'token withdraw getTokenWithdrawEventData query failed'
    )
  })
})

const depositTokenTest = async (bridge: Bridge) => {
  const tokenDepositAmount = BigNumber.from(1)

  const testToken = TestERC20__factory.connect(
    existentTestERC20,
    bridge.l1Signer
  )
  const mintRes = await testToken.mint()
  await mintRes.wait()

  const approveRes = await bridge.approveToken(existentTestERC20)
  await approveRes.wait()

  const data = await bridge.getAndUpdateL1TokenData(existentTestERC20)
  const allowed = data.ERC20 && data.ERC20.allowed
  expect(allowed, 'set token allowance failed').to.be.true

  const expectedL1GatewayAddress = await bridge.l1Bridge.getGatewayAddress(
    testToken.address
  )
  const initialBridgeTokenBalance = await testToken.balanceOf(
    expectedL1GatewayAddress
  )

  const depositRes = await bridge.deposit(existentTestERC20, tokenDepositAmount)

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
  await testRetryableTicket(bridge, depositRec)

  const l2Data = await bridge.getAndUpdateL2TokenData(existentTestERC20)

  const testWalletL2Balance = l2Data && l2Data.ERC20 && l2Data.ERC20.balance

  expect(
    testWalletL2Balance && testWalletL2Balance.eq(tokenDepositAmount),
    'l2 wallet not updated after deposit'
  ).to.be.true
}
