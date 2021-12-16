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

import { parseEther } from '@ethersproject/units'

import { AeWETH__factory } from '../src/lib/abi/factories/AeWETH__factory'

import {
  instantiateBridgeWithRandomWallet,
  fundL1,
  fundL2,
  skipIfMainnet,
  testRetryableTicket,
  prettyLog,
} from './testHelpers'
import { OutgoingMessageState } from '../src/lib/dataEntities'
import { L1ToL2Message } from '../src/lib/message/L1ToL2Message'
import {
  L2ToL1Message,
  L2TransactionReceipt,
} from '../src/lib/message/L2ToL1Message'

describe('WETH', async () => {
  beforeEach('skipIfMainnet', function () {
    skipIfMainnet(this)
  })

  it('withdraws WETH', async () => {
    const wethToWrap = parseEther('0.00001')
    const wethToWithdraw = parseEther('0.00000001')

    const {
      bridge,
      l1Network,
      l2Network,
    } = await instantiateBridgeWithRandomWallet()
    await fundL2(bridge)

    const l2Weth = AeWETH__factory.connect(
      l2Network.tokenBridge.l2Weth,
      bridge.l2Bridge.l2Signer
    )
    const res = await l2Weth.deposit({
      value: wethToWrap,
    })
    const rec = await res.wait()
    expect(rec.status).to.equal(1, 'deposit txn failed')

    const withdrawRes = await bridge.withdrawERC20(
      l1Network.tokenBridge.l1Weth,
      wethToWithdraw
    )
    const withdrawRec = await withdrawRes.wait()
    expect(withdrawRec.status).to.equal(1, 'withdraw txn failed')

    const outgoingMessages = await new L2TransactionReceipt(
      withdrawRec
    ).getL2ToL1Messages(bridge.l2Provider)
    const firstMessage = outgoingMessages[0]
    expect(firstMessage, 'getWithdrawalsInL2Transaction came back empty')
      .to.exist

    const outgoingMessageState = await firstMessage.status(null)
    expect(
      outgoingMessageState === OutgoingMessageState.UNCONFIRMED ||
        outgoingMessageState === OutgoingMessageState.NOT_FOUND,
      `weth withdraw getOutGoingMessageState returned ${outgoingMessageState}`
    )

    const _l2WethBalance = await bridge.l2Bridge.getL2TokenData(
      l2Network.tokenBridge.l2Weth
    )
    const l2WethBalance = _l2WethBalance.balance
    expect(
      l2WethBalance.add(wethToWithdraw).eq(wethToWrap),
      'balance not properly updated after weth withdraw'
    ).to.be.true

    const walletAddress = await bridge.l1Signer.getAddress()
    const gatewayWithdrawEvents = await bridge.getGatewayWithdrawEventData(
      l2Network.tokenBridge.l2WethGateway,
      walletAddress,
      { fromBlock: withdrawRec.blockNumber }
    )
    expect(gatewayWithdrawEvents.length).to.equal(
      1,
      'weth getGatewayWithdrawEventData query failed'
    )

    const tokenWithdrawEvents = await bridge.getTokenWithdrawEventData(
      l1Network.tokenBridge.l1Weth,
      walletAddress,
      { fromBlock: withdrawRec.blockNumber }
    )
    expect(tokenWithdrawEvents.length).to.equal(
      1,
      'weth getTokenWithdrawEventData query failed'
    )
  })

  it('deposits WETH', async () => {
    const { bridge, l1Network } = await instantiateBridgeWithRandomWallet()
    const l1WethAddress = l1Network.tokenBridge.l1Weth

    const wethToWrap = parseEther('0.0001')
    const wethToDeposit = parseEther('0.00001')

    await fundL1(bridge)

    const l1WETH = AeWETH__factory.connect(
      l1Network.tokenBridge.l1Weth,
      bridge.l1Signer
    )
    const res = await l1WETH.deposit({
      value: wethToWrap,
    })
    await res.wait()
    prettyLog('wrapped some ether')

    const approveRes = await bridge.approveToken(l1WethAddress)
    const approveRec = await approveRes.wait()
    expect(approveRec.status).to.equal(1, 'allowance txn failed')

    const data = await bridge.l1Bridge.getL1TokenData(l1WethAddress)
    const allowed = data.allowed
    expect(allowed, 'failed to set allowance').to.be.true

    const depositRes = await bridge.deposit({
      erc20L1Address: l1WethAddress,
      amount: wethToDeposit,
    })
    const depositRec = await depositRes.wait()
    await testRetryableTicket(bridge, depositRec)

    const l2Data = await bridge.l2Bridge.getL2TokenData(
      l1Network.tokenBridge.l2Weth
    )

    const testWalletL2Balance = l2Data.balance

    expect(
      testWalletL2Balance.eq(wethToDeposit),
      'ether balance not updated after deposit'
    ).to.be.true
  })
})
