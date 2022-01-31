/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

/* eslint-env node, mocha */

import { ethers } from 'hardhat'
import { BigNumber } from '@ethersproject/bignumber'
import { TransactionReceipt } from '@ethersproject/providers'
import { expect } from 'chai'
import {
  Bridge,
  Bridge__factory,
  Inbox,
  MessageTester,
  SequencerInbox,
} from '../build/types'
import { initializeAccounts } from './utils'
import { Event } from '@ethersproject/contracts'
import { Interface } from '@ethersproject/abi'
import { BridgeInterface, MessageDeliveredEvent } from '../build/types/Bridge'
import { Signer } from 'ethers'

const mineBlocks = async (signer: Signer, count: number) => {
  for (let index = 0; index < count; index++) {
    await (
      await signer.sendTransaction({
        to: await signer.getAddress(),
        value: 1,
      })
    ).wait()
  }
}

describe('SequencerInbox', async () => {
  const findMatchingLogs = <TInterface extends Interface, TEvent extends Event>(
    receipt: TransactionReceipt,
    iFace: TInterface,
    eventTopicGen: (i: TInterface) => string
  ): TEvent['args'][] => {
    const logs = receipt.logs.filter(
      log => log.topics[0] === eventTopicGen(iFace)
    )
    return logs.map(l => iFace.parseLog(l).args as TEvent['args'])
  }

  const getMessageDeliveredEvents = (receipt: TransactionReceipt) => {
    const bridgeInterface = Bridge__factory.createInterface()
    return findMatchingLogs<BridgeInterface, MessageDeliveredEvent>(
      receipt,
      bridgeInterface,
      i => i.getEventTopic(i.getEvent('MessageDelivered'))
    )
  }

  const sendDelayedTx = async (
    sender: Signer,
    inbox: Inbox,
    bridge: Bridge,
    messageTester: MessageTester,
    l2Gas: number,
    l2GasPrice: number,
    nonce: number,
    destAddr: string,
    amount: BigNumber,
    data: string
  ) => {
    const countBefore = (await bridge.functions.messageCount())[0].toNumber()
    const sendUnsignedTx = await inbox
      .connect(sender)
      .sendUnsignedTransaction(l2Gas, l2GasPrice, nonce, destAddr, amount, data)

    const countAfter = (await bridge.functions.messageCount())[0].toNumber()
    expect(countAfter, 'Unexpected inbox count').to.eq(countBefore + 1)

    const sendUnsignedTxReceipt = await sendUnsignedTx.wait()
    const senderAddr = await sender.getAddress()

    const l1GasPrice = sendUnsignedTxReceipt.effectiveGasPrice.toNumber()
    const messageDeliveredEvent = getMessageDeliveredEvents(
      sendUnsignedTxReceipt
    )[0]
    const l1BlockNumber = sendUnsignedTxReceipt.blockNumber
    const block1 = await sender.provider!.getBlock(l1BlockNumber)
    const l1BlockTimestamp = block1.timestamp
    const delayedAcc = await bridge.inboxAccs(countBefore)

    // need to hex pad the address
    const messageDataHash = ethers.utils.solidityKeccak256(
      ['uint8', 'uint256', 'uint256', 'uint256', 'uint256', 'uint256', 'bytes'],
      [
        0,
        l2Gas,
        l2GasPrice,
        nonce,
        ethers.utils.hexZeroPad(destAddr, 32),
        amount,
        data,
      ]
    )
    expect(
      messageDeliveredEvent.messageDataHash,
      'Incorrect messageDataHash'
    ).to.eq(messageDataHash)

    const messageHash = (
      await messageTester.functions.messageHash(
        3,
        senderAddr,
        l1BlockNumber,
        l1BlockTimestamp,
        countBefore,
        l1GasPrice,
        messageDataHash
      )
    )[0]
    const prevAccumulator = messageDeliveredEvent.beforeInboxAcc
    expect(prevAccumulator, 'Incorrect prev accumulator').to.eq(
      countBefore === 0
        ? ethers.utils.hexZeroPad('0x', 32)
        : await bridge.inboxAccs(countBefore - 1)
    )

    const nextAcc = (
      await messageTester.functions.addMessageToInbox(
        prevAccumulator,
        messageHash
      )
    )[0]

    expect(delayedAcc, 'Inocrrect delayed acc').to.eq(nextAcc)

    return {
      l1GasPrice,
      deliveredMessageEvent: messageDeliveredEvent,
      l1BlockNumber,
      l1BlockTimestamp,
      delayedAcc,
      l2Gas,
      l2GasPrice,
      nonce,
      destAddr,
      amount,
      data,
      senderAddr,
      inboxAccountLength: countAfter,
    }
  }

  const forceIncludeMessages = async (
    sequencerInbox: SequencerInbox,
    newTotalDelayedMessagesRead: number,
    kind: number,
    l1blockNumber: number,
    l1Timestamp: number,
    delayedInboxSeqNum: BigNumber,
    l1GasPrice: number,
    senderAddr: string,
    messageDataHash: string,
    delayedAcc: string,
    expectedErrorMessage?: string
  ) => {
    const totalDelayedMessagesReadBefore = (
      await sequencerInbox.totalDelayedMessagesRead()
    ).toNumber()
    const inboxLengthBefore = (
      await sequencerInbox.getInboxAccsLength()
    ).toNumber()
    const messageCountBefore = (await sequencerInbox.messageCount()).toNumber()

    const forceInclusionTx = sequencerInbox.forceInclusion(
      newTotalDelayedMessagesRead,
      kind,
      [l1blockNumber, l1Timestamp],
      delayedInboxSeqNum,
      l1GasPrice,
      senderAddr,
      messageDataHash,
      delayedAcc
    )
    if (expectedErrorMessage) {
      await expect(forceInclusionTx).to.be.revertedWith(expectedErrorMessage)
    } else {
      await (await forceInclusionTx).wait()

      const totalDelayedMessagsReadAfter = (
        await sequencerInbox.totalDelayedMessagesRead()
      ).toNumber()
      expect(
        totalDelayedMessagsReadAfter,
        'Incorrect totalDelayedMessagesRead after.'
      ).to.eq(newTotalDelayedMessagesRead)
      const inboxLengthAfter = (
        await sequencerInbox.getInboxAccsLength()
      ).toNumber()
      expect(
        inboxLengthAfter - inboxLengthBefore,
        'Inbox not incremented'
      ).to.eq(1)
      const messageCountAfter = (await sequencerInbox.messageCount()).toNumber()
      expect(
        messageCountAfter - messageCountBefore,
        'Message count invalid.'
      ).to.eq(newTotalDelayedMessagesRead - totalDelayedMessagesReadBefore + 1)
    }
  }

  const setupSequencerInbox = async () => {
    const accounts = await initializeAccounts()
    const sequencer = accounts[0]
    const user = accounts[1]
    const dummyRollup = accounts[2]

    const Bridge = await ethers.getContractFactory('Bridge')
    const bridge = (await Bridge.deploy()) as Bridge
    await bridge.deployed()
    await bridge.initialize()

    const Inbox = await ethers.getContractFactory('Inbox')
    const inbox = (await Inbox.deploy()) as Inbox
    await inbox.deployed()
    await inbox.initialize(bridge.address, ethers.constants.AddressZero)
    await bridge.setInbox(inbox.address, true)

    const SequencerInbox = await ethers.getContractFactory('SequencerInbox')
    const sequencerInbox = await SequencerInbox.deploy()
    await sequencerInbox.deployed()
    await sequencerInbox.initialize(
      bridge.address,
      await sequencer.getAddress(),
      await dummyRollup.getAddress()
    )
    const maxDelayBlocks = 10
    await sequencerInbox.connect(dummyRollup).setMaxDelay(maxDelayBlocks, 0)

    const messageTester = await (
      await ethers.getContractFactory('MessageTester')
    ).deploy()

    return {
      user,
      bridge,
      inbox,
      sequencerInbox,
      messageTester,
    }
  }
  let user: Signer,
    bridge: Bridge,
    inbox: Inbox,
    messageTester: MessageTester,
    sequencerInbox: SequencerInbox
  before(async () => {
    const setup = await setupSequencerInbox()
    user = setup.user
    bridge = setup.bridge
    inbox = setup.inbox
    messageTester = setup.messageTester
    sequencerInbox = setup.sequencerInbox
  })

  it('can force-include', async () => {
    const delayedTx = await sendDelayedTx(
      user,
      inbox,
      bridge,
      messageTester,
      1000000,
      21000000000,
      0,
      await user.getAddress(),
      BigNumber.from(10),
      '0x1010'
    )

    const maxDelayBlocks = (await sequencerInbox.maxDelayBlocks()).toNumber()
    await mineBlocks(user, maxDelayBlocks)

    await forceIncludeMessages(
      sequencerInbox,
      delayedTx.inboxAccountLength,
      delayedTx.deliveredMessageEvent.kind,
      delayedTx.l1BlockNumber,
      delayedTx.l1BlockTimestamp,
      delayedTx.deliveredMessageEvent.messageIndex,
      delayedTx.l1GasPrice,
      delayedTx.senderAddr,
      delayedTx.deliveredMessageEvent.messageDataHash,
      delayedTx.delayedAcc
    )
  })

  it('can force-include one after another', async () => {
    const delayedTx = await sendDelayedTx(
      user,
      inbox,
      bridge,
      messageTester,
      1000000,
      21000000000,
      0,
      await user.getAddress(),
      BigNumber.from(10),
      '0x1010'
    )

    const delayedTx2 = await sendDelayedTx(
      user,
      inbox,
      bridge,
      messageTester,
      1000000,
      21000000000,
      1,
      await user.getAddress(),
      BigNumber.from(10),
      '0xdeadface'
    )

    const maxDelayBlocks = (await sequencerInbox.maxDelayBlocks()).toNumber()
    await mineBlocks(user, maxDelayBlocks)

    await forceIncludeMessages(
      sequencerInbox,
      delayedTx.inboxAccountLength,
      delayedTx.deliveredMessageEvent.kind,
      delayedTx.l1BlockNumber,
      delayedTx.l1BlockTimestamp,
      delayedTx.deliveredMessageEvent.messageIndex,
      delayedTx.l1GasPrice,
      delayedTx.senderAddr,
      delayedTx.deliveredMessageEvent.messageDataHash,
      delayedTx.delayedAcc
    )
    await forceIncludeMessages(
      sequencerInbox,
      delayedTx2.inboxAccountLength,
      delayedTx2.deliveredMessageEvent.kind,
      delayedTx2.l1BlockNumber,
      delayedTx2.l1BlockTimestamp,
      delayedTx2.deliveredMessageEvent.messageIndex,
      delayedTx2.l1GasPrice,
      delayedTx2.senderAddr,
      delayedTx2.deliveredMessageEvent.messageDataHash,
      delayedTx2.delayedAcc
    )
  })

  it('can force-include three at once', async () => {
    await sendDelayedTx(
      user,
      inbox,
      bridge,
      messageTester,
      1000000,
      21000000000,
      0,
      await user.getAddress(),
      BigNumber.from(10),
      '0x1010'
    )
    await sendDelayedTx(
      user,
      inbox,
      bridge,
      messageTester,
      1000000,
      21000000000,
      1,
      await user.getAddress(),
      BigNumber.from(10),
      '0x101010'
    )
    const delayedTx3 = await sendDelayedTx(
      user,
      inbox,
      bridge,
      messageTester,
      1000000,
      21000000000,
      10,
      await user.getAddress(),
      BigNumber.from(10),
      '0x10101010'
    )

    const maxDelayBlocks = (await sequencerInbox.maxDelayBlocks()).toNumber()
    await mineBlocks(user, maxDelayBlocks)

    await forceIncludeMessages(
      sequencerInbox,
      delayedTx3.inboxAccountLength,
      delayedTx3.deliveredMessageEvent.kind,
      delayedTx3.l1BlockNumber,
      delayedTx3.l1BlockTimestamp,
      delayedTx3.deliveredMessageEvent.messageIndex,
      delayedTx3.l1GasPrice,
      delayedTx3.senderAddr,
      delayedTx3.deliveredMessageEvent.messageDataHash,
      delayedTx3.delayedAcc
    )
  })

  it('cannot include before max delay', async () => {
    const delayedTx = await sendDelayedTx(
      user,
      inbox,
      bridge,
      messageTester,
      1000000,
      21000000000,
      0,
      await user.getAddress(),
      BigNumber.from(10),
      '0x1010'
    )

    const maxDelayBlocks = (await sequencerInbox.maxDelayBlocks()).toNumber()
    await mineBlocks(user, maxDelayBlocks - 1)

    await forceIncludeMessages(
      sequencerInbox,
      delayedTx.inboxAccountLength,
      delayedTx.deliveredMessageEvent.kind,
      delayedTx.l1BlockNumber,
      delayedTx.l1BlockTimestamp,
      delayedTx.deliveredMessageEvent.messageIndex,
      delayedTx.l1GasPrice,
      delayedTx.senderAddr,
      delayedTx.deliveredMessageEvent.messageDataHash,
      delayedTx.delayedAcc,
      'MAX_DELAY_BLOCKS'
    )
  })
})
