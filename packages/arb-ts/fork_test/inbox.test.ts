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
import { Logger, LogLevel } from '@ethersproject/logger'
import { Provider } from '@ethersproject/abstract-provider'
Logger.setLogLevel(LogLevel.ERROR)

import {
  Bridge__factory,
  Inbox__factory,
  SequencerInbox__factory,
} from '../src/lib/abi'

import { InboxTools } from '../src'

import { ethers, network } from 'hardhat'
import { hexZeroPad } from '@ethersproject/bytes'
import { l2Networks, L2Network } from '../src/lib/dataEntities/networks'
import { solidityKeccak256 } from 'ethers/lib/utils'
import { ContractTransaction, Signer } from 'ethers'
import { ArbTsError } from '../src/lib/dataEntities/errors'
import { SignerProviderUtils } from '../src/lib/dataEntities/signerOrProvider'

const submitL2Tx = async (
  tx: {
    to: string
    value?: BigNumber
    data?: string
    nonce: number
    gasPriceBid: BigNumber
    gasLimit: BigNumber
  },
  l2Network: L2Network,
  l1Signer?: Signer
): Promise<ContractTransaction> => {
  const inbox = Inbox__factory.connect(l2Network.ethBridge.inbox, l1Signer)
  const senderAddr = await l1Signer.getAddress()

  return await inbox.sendUnsignedTransaction(
    tx.gasLimit,
    tx.gasPriceBid,
    tx.nonce,
    tx.to,
    tx.value || BigNumber.from(0),
    tx.data || '0x'
  )
}

describe('Inbox tools', () => {
  const setup = async () => {
    const signers = await ethers.getSigners()
    const signer = signers[0]
    const provider = signer.provider!

    const arbitrumOne = l2Networks[42161]

    const sequencerInbox = SequencerInbox__factory.connect(
      arbitrumOne.ethBridge.sequencerInbox,
      provider
    )

    const bridge = Bridge__factory.connect(
      arbitrumOne.ethBridge.bridge,
      provider
    )

    return {
      l1Signer: signer,
      l1Provider: provider,
      l2Network: arbitrumOne,
      sequencerInbox,
      bridge,
    }
  }

  let forkBlockNumber: number
  let forkProviderUrl: string
  before(async () => {
    const { l1Provider } = await setup()
    forkBlockNumber = await l1Provider.getBlockNumber()
    forkProviderUrl =
      'https://mainnet.infura.io/v3/' + process.env['INFURA_KEY']
  })

  const resetFork = async (blockNumber: number, jsonRpcUrl: string) => {
    await network.provider.request({
      method: 'hardhat_reset',
      params: [{ forking: { jsonRpcUrl, blockNumber } }],
    })
  }

  beforeEach(async () => {
    // we reset the fork between each test so that the tests don't
    // interfere with one another
    await resetFork(forkBlockNumber, forkProviderUrl)
  })

  it('can force include', async () => {
    const { l1Signer, l2Network, sequencerInbox, bridge } = await setup()

    const inboxTools = new InboxTools(l1Signer, l2Network)
    const startInboxLength = await bridge.messageCount()
    const l2Tx = await submitL2Tx(
      {
        to: await l1Signer.getAddress(),
        value: BigNumber.from(0),
        gasLimit: BigNumber.from(100000),
        gasPriceBid: BigNumber.from(21000000000),
        nonce: 0,
      },
      l2Network,
      l1Signer
    )
    await l2Tx.wait()

    const block = await l1Signer.provider!.getBlock('latest')
    await mineBlocks(6600, block.timestamp)

    const forceInclusionTx = await inboxTools.forceInclude()

    expect(forceInclusionTx, 'Null force inclusion').to.not.be.null
    await forceInclusionTx!.wait()

    const messagesReadAfter = await sequencerInbox.totalDelayedMessagesRead()

    expect(messagesReadAfter.toNumber(), 'Message not read').to.eq(
      startInboxLength.add(1).toNumber()
    )
  })

  it('can force include many', async () => {
    const { l1Signer, l2Network, sequencerInbox, bridge } = await setup()

    const startInboxLength = await bridge.messageCount()
    const l2Tx1 = await submitL2Tx(
      {
        to: await l1Signer.getAddress(),
        value: BigNumber.from(0),
        gasLimit: BigNumber.from(100000),
        gasPriceBid: BigNumber.from(21000000000),
        nonce: 0,
      },
      l2Network,
      l1Signer
    )
    await l2Tx1.wait()

    const l2Tx2 = await submitL2Tx(
      {
        to: await l1Signer.getAddress(),
        value: BigNumber.from(10),
        gasLimit: BigNumber.from(100000),
        gasPriceBid: BigNumber.from(21000000000),
        nonce: 1,
      },
      l2Network,
      l1Signer
    )
    await l2Tx2.wait()

    const block = await l1Signer.provider!.getBlock('latest')
    await mineBlocks(6600, block.timestamp)

    const inboxTools = new InboxTools(l1Signer, l2Network)
    const forceInclusionTx = await inboxTools.forceInclude()

    expect(forceInclusionTx, 'Null force inclusion').to.not.be.null
    await forceInclusionTx!.wait()

    const messagesReadAfter = await sequencerInbox.totalDelayedMessagesRead()
    expect(messagesReadAfter.toNumber(), 'Message not read').to.eq(
      startInboxLength.add(2).toNumber()
    )
  })

  it('does find eligible events', async () => {
    const { l1Signer, l2Network } = await setup()
    const inboxTools = new InboxTools(l1Signer, l2Network)

    const l2Tx1 = await submitL2Tx(
      {
        to: await l1Signer.getAddress(),
        value: BigNumber.from(5),
        gasLimit: BigNumber.from(100000),
        gasPriceBid: BigNumber.from(21000000000),
        nonce: 0,
      },
      l2Network,
      l1Signer
    )
    await l2Tx1.wait()
    const txParams = {
      to: await l1Signer.getAddress(),
      value: BigNumber.from(10),
      gasLimit: BigNumber.from(100000),
      gasPriceBid: BigNumber.from(21000000000),
      nonce: 1,
    }
    const messageDataHash = solidityKeccak256(
      ['uint8', 'uint256', 'uint256', 'uint256', 'uint256', 'uint256', 'bytes'],
      [
        0,
        txParams.gasLimit,
        txParams.gasPriceBid,
        txParams.nonce,
        hexZeroPad(txParams.to, 32),
        txParams.value,
        '0x',
      ]
    )
    const l2Tx2 = await submitL2Tx(txParams, l2Network, l1Signer)
    await l2Tx2.wait()

    const block = await l1Signer.provider!.getBlock('latest')
    await mineBlocks(6600, block.timestamp)

    const event = await inboxTools.getForceIncludeableEvent()
    expect(event?.event.messageDataHash, 'Invalid message hash.').to.eq(
      messageDataHash
    )
  })

  it('doesnt find non-eligible events', async () => {
    const { l1Signer, l2Network } = await setup()
    const inboxTools = new InboxTools(l1Signer, l2Network)

    const event = await inboxTools.getForceIncludeableEvent()
    expect(event, 'Event not null').to.be.null
  })

  const mineBlocks = async (
    count: number,
    startTimestamp?: number,
    timeDiffPerBlock = 14
  ) => {
    let timestamp = startTimestamp
    for (let i = 0; i < count; i++) {
      timestamp = Math.max(
        Math.floor(Date.now() / 1000) + (timeDiffPerBlock || 1),
        (timestamp || 0) + (timeDiffPerBlock || 1)
      )
      await network.provider.send('evm_mine', [timestamp])
    }
  }
})
