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
import { TokenBridger } from '../src'

describe('WETH', async () => {
  beforeEach('skipIfMainnet', function () {
    skipIfMainnet(this)
  })

  it('withdraws WETH', async () => {
    const wethToWrap = parseEther('0.00001')
    const wethToWithdraw = parseEther('0.00000001')

    const { l2Network, l1Signer, l2Signer, tokenBridger } =
      await instantiateBridgeWithRandomWallet()
    await fundL2(l2Signer)

    const l2Weth = AeWETH__factory.connect(
      l2Network.tokenBridge.l2Weth,
      l2Signer
    )
    const res = await l2Weth.deposit({
      value: wethToWrap,
    })
    const rec = await res.wait()
    expect(rec.status).to.equal(1, 'deposit txn failed')

    const withdrawRes = await tokenBridger.withdraw({
      amount: wethToWithdraw,
      erc20l1Address: l2Network.tokenBridge.l1Weth,
      l2Signer: l2Signer,
    })
    const withdrawRec = await withdrawRes.wait()
    expect(withdrawRec.status).to.equal(1, 'withdraw txn failed')

    const outgoingMessages = await withdrawRec.getL2ToL1Messages(
      l2Signer.provider!,
      l2Network
    )
    const firstMessage = outgoingMessages[0]
    expect(firstMessage, 'getWithdrawalsInL2Transaction came back empty').to
      .exist

    const outgoingMessageState = await firstMessage.status(null)
    expect(
      outgoingMessageState === OutgoingMessageState.UNCONFIRMED ||
        outgoingMessageState === OutgoingMessageState.NOT_FOUND,
      `weth withdraw getOutGoingMessageState returned ${outgoingMessageState}`
    )

    const l2Token = await tokenBridger.getL2TokenContract(
      l2Signer.provider!,
      l2Network.tokenBridge.l2Weth
    )
    const l2WethBalance = (
      await l2Token.functions.balanceOf(await l2Signer.getAddress())
    )[0]

    expect(
      l2WethBalance.add(wethToWithdraw).eq(wethToWrap),
      'balance not properly updated after weth withdraw'
    ).to.be.true

    const walletAddress = await l1Signer.getAddress()
    const gatewayWithdrawEvents = await tokenBridger.getL2WithdrawalEvents(
      l2Signer.provider!,
      l2Network.tokenBridge.l2WethGateway,
      undefined,
      walletAddress,
      { fromBlock: withdrawRec.blockNumber }
    )
    expect(gatewayWithdrawEvents.length).to.equal(
      1,
      'weth token gateway query failed'
    )

    const gatewayAddress = await tokenBridger.getL2GatewayAddress(
      l2Network.tokenBridge.l1Weth,
      l2Signer.provider!
    )
    const tokenWithdrawEvents = await tokenBridger.getL2WithdrawalEvents(
      l2Signer.provider!,
      gatewayAddress,
      l2Network.tokenBridge.l1Weth,
      walletAddress,
      { fromBlock: withdrawRec.blockNumber }
    )
    expect(tokenWithdrawEvents.length).to.equal(
      1,
      'token filtered query failed'
    )
  })

  it('deposits WETH', async () => {
    const { l2Network, l1Signer, l2Signer, tokenBridger } =
      await instantiateBridgeWithRandomWallet()
    const l1WethAddress = l2Network.tokenBridge.l1Weth

    const wethToWrap = parseEther('0.00001')
    const wethToDeposit = parseEther('0.0000001')

    await fundL1(l1Signer)

    const l1WETH = AeWETH__factory.connect(
      l2Network.tokenBridge.l1Weth,
      l1Signer
    )
    const res = await l1WETH.deposit({
      value: wethToWrap,
    })
    await res.wait()
    prettyLog('wrapped some ether')

    const approveRes = await tokenBridger.approveToken({
      erc20L1Address: l1WethAddress,
      l1Signer: l1Signer,
    })
    const approveRec = await approveRes.wait()
    expect(approveRec.status).to.equal(1, 'allowance txn failed')

    const l1Token = tokenBridger.getL1TokenContract(
      l1Signer.provider!,
      l1WethAddress
    )
    const allowance = (
      await l1Token.functions.allowance(
        await l1Signer.getAddress(),
        l2Network.tokenBridge.l1WethGateway
      )
    )[0]
    expect(allowance.eq(TokenBridger.MAX_APPROVAL), 'failed to set allowance')
      .to.be.true

    const depositRes = await tokenBridger.deposit({
      amount: wethToDeposit,
      erc20L1Address: l1WethAddress,
      l1Signer: l1Signer,
      l2Provider: l2Signer.provider!,
      retryableGasOverrides: {
        // CHRIS: this seems backwards - except in neither case do we actually want to add value right?
        // CHRIS: so what's the purpose of it?
        sendL2CallValueFromL1: false,
      },
    })
    const depositRec = await depositRes.wait()
    await testRetryableTicket(l2Signer.provider!, depositRec)

    const l2Token = tokenBridger.getL2TokenContract(
      l2Signer.provider!,
      l2Network.tokenBridge.l2Weth
    )
    const testWalletL2Balance = (
      await l2Token.functions.balanceOf(await l2Signer.getAddress())
    )[0]
    expect(
      testWalletL2Balance.eq(wethToDeposit),
      'ether balance not updated after deposit'
    ).to.be.true
  })
})
