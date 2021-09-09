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

import {
  instantiateBridgeWithRandomWallet,
  fundL1,
  wait,
  fundL2,
  prettyLog,
  skipIfMainnet,
} from './testHelpers'
import { ArbGasInfo__factory } from '../src/lib/abi/factories/ArbGasInfo__factory'
import { ARB_GAS_INFO } from '../src/lib/precompile_addresses'
import { OutgoingMessageState } from '../src/lib/bridge_helpers'
dotenv.config()

describe('Ether', async () => {
  beforeEach('skipIfMainnet', function () {
    skipIfMainnet(this)
  })

  it('transfers ether on l2', async () => {
    const { bridge } = await instantiateBridgeWithRandomWallet()
    await fundL2(bridge)
    const randomAddress = Wallet.createRandom().address
    const amountToSend = parseEther('0.000005')
    const res = await bridge.l2Signer.sendTransaction({
      to: randomAddress,
      value: amountToSend,
    })
    const rec = await res.wait()

    expect(rec.status).to.equal(1, 'ether transfer failed')
    const newBalance = await bridge.l2Provider.getBalance(randomAddress)
    expect(newBalance.eq(amountToSend), "ether balance didn't update").to.be
      .true
  })
  it('deposits ether', async () => {
    const { bridge } = await instantiateBridgeWithRandomWallet()
    await fundL1(bridge)

    const inbox = await bridge.l1Bridge.getInbox()

    const initialInboxBalance = await bridge.l1Bridge.l1Provider.getBalance(
      inbox.address
    )
    const ethToDeposit = parseEther('0.0002')
    const res = await bridge.depositETH(ethToDeposit)
    const rec = await res.wait()

    expect(rec.status).to.equal(1, 'eth deposit L1 txn failed')
    const finalInboxBalance = await bridge.l1Bridge.l1Provider.getBalance(
      inbox.address
    )
    expect(
      initialInboxBalance.add(ethToDeposit).eq(finalInboxBalance),
      'balance failed to update after eth deposit'
    )

    const seqNumArr = await bridge.getInboxSeqNumFromContractTransaction(rec)
    if (seqNumArr === undefined) {
      throw new Error('no seq num')
    }
    expect(seqNumArr.length, 'eth deposit seqNum not found').to.exist

    const seqNum = seqNumArr[0]
    const l2TxHash = await bridge.calculateL2TransactionHash(seqNum)
    prettyLog('l2TxHash: ' + l2TxHash)
    prettyLog('waiting for l2 transaction:')
    const l2TxnRec = await bridge.l2Bridge.l2Provider.waitForTransaction(
      l2TxHash,
      undefined,
      1000 * 60 * 12
    )
    prettyLog('l2 transaction found!')
    expect(l2TxnRec.status).to.equal(1, 'eth deposit l2 transaction not found')

    for (let i = 0; i < 60; i++) {
      prettyLog('balance check attempt ' + (i + 1))
      await wait(5000)
      const testWalletL2EthBalance = await bridge.getL2EthBalance()
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
    const { bridge } = await instantiateBridgeWithRandomWallet()
    await fundL2(bridge)
    const ethToWithdraw = parseEther('0.00002')

    const initialBalance = await bridge.l2Bridge.getL2EthBalance()

    const withdrawEthRes = await bridge.withdrawETH(ethToWithdraw)
    const withdrawEthRec = await withdrawEthRes.wait()

    const arbGasInfo = ArbGasInfo__factory.connect(
      ARB_GAS_INFO,
      bridge.l2Provider
    )
    expect(withdrawEthRec.status).to.equal(
      1,
      'initiate eth withdraw txn failed'
    )

    const inWei = await arbGasInfo.getPricesInWei({
      blockTag: withdrawEthRec.blockNumber,
    })
    const withdrawEventData =
      bridge.getWithdrawalsInL2Transaction(withdrawEthRec)[0]

    expect(
      withdrawEventData,
      'eth withdraw getWithdrawalsInL2Transaction query came back empty'
    ).to.exist

    const myAddress = await bridge.l1Signer.getAddress()

    const withdrawEvents = await bridge.getL2ToL1EventData(myAddress, {
      fromBlock: withdrawEthRec.blockNumber,
    })
    expect(withdrawEvents.length).to.equal(
      1,
      'eth withdraw getL2ToL1EventData failed'
    )

    const outgoingMessageState = await bridge.getOutGoingMessageState(
      withdrawEventData.batchNumber,
      withdrawEventData.indexInBatch
    )
    expect(outgoingMessageState).to.equal(
      OutgoingMessageState.UNCONFIRMED,
      `eth withdraw getOutGoingMessageState returned ${OutgoingMessageState.UNCONFIRMED}`
    )

    const etherBalance = await bridge.getL2EthBalance()

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
