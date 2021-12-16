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

describe('Custom ERC20', () => {
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
    const result = await fundL2Token(bridge, existentTestCustomToken)
    if (!result) {
      warn('Prefunded wallet not funded with tokens; skipping ERC20 withdraw')
      this.skip()
    }
    const withdrawRes = await bridge.withdrawERC20(
      existentTestCustomToken,
      tokenWithdrawAmount
    )
    const withdrawRec = await withdrawRes.wait()

    expect(withdrawRec.status).to.equal(1, 'initiate token withdraw txn failed')
    
    const message = (await new L2TransactionReceipt(withdrawRec).getL2ToL1Messages(bridge.l1Provider))[0]
    expect(message, 'withdrawEventData not found').to.exist

    const outgoingMessageState = await message.status(null)
    expect(
      outgoingMessageState === OutgoingMessageState.UNCONFIRMED ||
        outgoingMessageState === OutgoingMessageState.NOT_FOUND,
      `custom token withdraw getOutGoingMessageState returned ${outgoingMessageState}`
    ).to.be.true

    const l2Data = await bridge.l2Bridge.getL2TokenData(
      await bridge.getERC20L2Address(existentTestCustomToken)
    )
    const testWalletL2Balance = l2Data.balance
    expect(
      testWalletL2Balance.add(tokenWithdrawAmount).eq(tokenFundAmount),
      'wallet balance not properly deducted after withdraw'
    ).to.be.true
    const walletAddress = await bridge.l1Signer.getAddress()

    const gatewayWithdrawEvents = await bridge.getGatewayWithdrawEventData(
      l2Network.tokenBridge.l2CustomGateway,
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

const depositTokenTest = async (bridge: Bridge) => {
  const tokenDepositAmount = BigNumber.from(1)

  const testToken = TestERC20__factory.connect(
    existentTestCustomToken,
    bridge.l1Signer
  )
  const mintRes = await testToken.mint()
  await mintRes.wait()

  const approveRes = await bridge.approveToken(existentTestCustomToken)
  await approveRes.wait()

  const data = await bridge.l1Bridge.getL1TokenData(existentTestCustomToken)
  const allowed = data.allowed
  expect(allowed, 'set token allowance failed').to.be.true

  const expectedL1GatewayAddress = await bridge.l1Bridge.getGatewayAddress(
    testToken.address
  )
  const initialBridgeTokenBalance = await testToken.balanceOf(
    expectedL1GatewayAddress
  )

  const depositRes = await bridge.deposit({
    erc20L1Address: existentTestCustomToken,
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
    'bridge balance not properly updated after deposit'
  ).to.be.true
  await testRetryableTicket(bridge, depositRec)

  const l2Data = await bridge.l2Bridge.getL2TokenData(
    await bridge.getERC20L2Address(existentTestCustomToken)
  )

  const testWalletL2Balance = l2Data.balance

  expect(
    testWalletL2Balance.eq(tokenDepositAmount),
    "l2 wallet balance not properly updated after deposit'"
  ).to.be.true
}
