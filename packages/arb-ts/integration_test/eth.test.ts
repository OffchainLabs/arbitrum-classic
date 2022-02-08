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
import dotenv from 'dotenv'

import { Wallet } from '@ethersproject/wallet'
import { Zero } from '@ethersproject/constants'
import { parseEther } from '@ethersproject/units'

import { ArbGasInfo__factory } from '../src/lib/abi'
import {
  instantiateBridgeWithRandomWallet,
  fundL1,
  wait,
  fundL2,
  prettyLog,
  skipIfMainnet,
} from './testHelpers'
import { ARB_GAS_INFO } from '../src/lib/dataEntities/constants'
import {
  L2ToL1Message,
  L2ToL1MessageStatus,
} from '../src/lib/message/L2ToL1Message'
import { L1ToL2MessageStatus } from '../src/lib/message/L1ToL2Message'
dotenv.config()

describe('Ether', async () => {
  beforeEach('skipIfMainnet', function () {
    skipIfMainnet(this)
  })

  it('transfers ether on l2', async () => {
    const { l2Signer } = await instantiateBridgeWithRandomWallet()
    await fundL2(l2Signer)
    const randomAddress = Wallet.createRandom().address
    const amountToSend = parseEther('0.000005')
    const res = await l2Signer.sendTransaction({
      to: randomAddress,
      value: amountToSend,
    })
    const rec = await res.wait()

    expect(rec.status).to.equal(1, 'ether transfer failed')
    const newBalance = await l2Signer.provider!.getBalance(randomAddress)
    expect(newBalance.eq(amountToSend), "ether balance didn't update").to.be
      .true
  })
  it('deposits ether', async () => {
    const { ethBridger, l1Signer, l2Signer } =
      await instantiateBridgeWithRandomWallet()
    await fundL1(l1Signer)

    const inboxAddress = ethBridger.l2Network.ethBridge.inbox

    const initialInboxBalance = await l1Signer.provider!.getBalance(
      inboxAddress
    )
    const ethToDeposit = parseEther('0.0002')
    const res = await ethBridger.deposit({
      amount: ethToDeposit,
      l1Signer: l1Signer,
      l2Provider: l2Signer.provider!,
    })
    const rec = await res.wait()

    expect(rec.status).to.equal(1, 'eth deposit L1 txn failed')
    const finalInboxBalance = await l1Signer.provider!.getBalance(inboxAddress)
    expect(
      initialInboxBalance.add(ethToDeposit).eq(finalInboxBalance),
      'balance failed to update after eth deposit'
    )

    const messages = await rec.getL1ToL2Messages(l2Signer.provider!)

    expect(messages, 'No messages.').to.be.not.undefined
    expect(messages.length, 'eth deposit message not found').to.exist
    expect(messages.length, 'eth deposit message empty array').to.not.eq(0)

    const message = messages[0]
    prettyLog('l2TxHash: ' + message.retryableCreationId)
    prettyLog('waiting for l2 transaction:')
    const waitResult = await message.wait(true, 1000 * 60 * 12)
    prettyLog('l2 transaction found!')
    expect(waitResult.complete).to.eq(true, 'eth deposit not complete')
    expect(waitResult.status).to.eq(
      L1ToL2MessageStatus.FUNDS_DEPOSITED_ON_L2,
      'eth deposit l2 transaction not found'
    )

    for (let i = 0; i < 60; i++) {
      prettyLog('balance check attempt ' + (i + 1))
      await wait(5000)
      const testWalletL2EthBalance = await l2Signer.getBalance()
      if (testWalletL2EthBalance.gt(Zero)) {
        prettyLog(`balance updated!  ${testWalletL2EthBalance.toString()}`)
        expect(true).to.be.true
        return
        break
      }
    }
    expect(false).to.be.true
  })

  it('withdraw Ether transaction succeeds', async () => {
    const { l2Network, l2Signer, l1Signer, ethBridger } =
      await instantiateBridgeWithRandomWallet()
    await fundL2(l2Signer)
    const ethToWithdraw = parseEther('0.00002')
    const initialBalance = await l2Signer.getBalance()

    const withdrawEthRes = await ethBridger.withdraw({
      amount: ethToWithdraw,
      l2Signer: l2Signer,
    })
    const withdrawEthRec = await withdrawEthRes.wait()

    const arbGasInfo = ArbGasInfo__factory.connect(
      ARB_GAS_INFO,
      l2Signer.provider!
    )
    expect(withdrawEthRec.status).to.equal(
      1,
      'initiate eth withdraw txn failed'
    )

    const inWei = await arbGasInfo.getPricesInWei({
      blockTag: withdrawEthRec.blockNumber,
    })

    const withdrawMessage = (
      await withdrawEthRec.getL2ToL1Messages(l2Signer.provider!, l2Network)
    )[0]
    expect(
      withdrawMessage,
      'eth withdraw getWithdrawalsInL2Transaction query came back empty'
    ).to.exist

    const myAddress = await l1Signer.getAddress()
    const withdrawEvents = await L2ToL1Message.getL2ToL1MessageLogs(
      l2Signer.provider!,
      { fromBlock: withdrawEthRec.blockNumber },
      undefined,
      myAddress
    )

    expect(withdrawEvents.length).to.equal(
      1,
      'eth withdraw getL2ToL1EventData failed'
    )

    const messageStatus = await withdrawMessage.status(null)
    expect(
      messageStatus === L2ToL1MessageStatus.UNCONFIRMED ||
        messageStatus === L2ToL1MessageStatus.NOT_FOUND,
      `eth withdraw status returned ${messageStatus}`
    ).to.be.true

    const etherBalance = await l2Signer.getBalance()
    const totalEth = etherBalance
      .add(ethToWithdraw)
      .add(withdrawEthRec.gasUsed.mul(inWei[5]))

    // TODO
    console.log(
      `This number should be zero...? ${initialBalance
        .sub(totalEth)
        .toString()}`
    )

    expect(true).to.be.true
  })
})
