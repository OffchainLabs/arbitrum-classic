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

import { Signer } from '@ethersproject/abstract-signer'
import { Block } from '@ethersproject/abstract-provider'
import { BigNumber, ContractTransaction, Overrides } from 'ethers'

import { Bridge, Bridge__factory, SequencerInbox__factory } from '../abi'
import { MessageDeliveredEvent } from '../abi/Bridge'
import { L2Network } from '../dataEntities/networks'
import { ArbTsError } from '../dataEntities/errors'
import { SignerProviderUtils } from '../dataEntities/signerOrProvider'
import { FetchedEvent, EventFetcher } from '../utils/eventFetcher'
import { MultiCaller } from '../utils/multicall'

type ForceInclusionParams = FetchedEvent<MessageDeliveredEvent> & {
  delayedAcc: string
}

/**
 * Tools for interacting with the inbox and bridge contracts
 */
export class InboxTools {
  private readonly l1Provider

  constructor(
    private readonly l1Signer: Signer,
    private readonly l2Network: L2Network
  ) {
    this.l1Provider = SignerProviderUtils.getProviderOrThrow(this.l1Signer)
  }

  /**
   * Find the first (or close to first) block whose number
   * is below the provided number, and whose timestamp is below
   * the provided timestamp
   * @param blockNumber
   * @param blockTimestamp
   * @returns
   */
  private async findFirstBlockBelow(
    blockNumber: number,
    blockTimestamp: number
  ): Promise<Block> {
    const block = await this.l1Provider.getBlock(blockNumber)
    const diff = block.timestamp - blockTimestamp
    if (diff < 0) return block

    // we take a long average block time of 14s
    // and always move at least 10 blocks

    const diffBlocks = Math.max(Math.ceil(diff / 14), 10)

    return await this.findFirstBlockBelow(
      blockNumber - diffBlocks,
      blockTimestamp
    )
  }

  /**
   * Get a range of blocks within eligible messages emitted events
   * @param blockNumbeRangeSize
   * @returns
   */
  private async getEligibleBlockRange(blockNumbeRangeSize: number) {
    const currentBlock = await this.l1Provider.getBlock('latest')
    const sequencerInbox = SequencerInbox__factory.connect(
      this.l2Network.ethBridge.sequencerInbox,
      this.l1Provider
    )

    const multicall = await MultiCaller.fromProvider(this.l1Provider)
    const [maxDelayBlocks, maxDelaySeconds] = await multicall.multiCall([
      {
        targetAddr: sequencerInbox.address,
        encoder: () =>
          sequencerInbox.interface.encodeFunctionData('maxDelayBlocks'),
        decoder: (returnData: string) =>
          sequencerInbox.interface.decodeFunctionResult(
            'maxDelayBlocks',
            returnData
          )[0] as BigNumber,
      },
      {
        targetAddr: sequencerInbox.address,
        encoder: () =>
          sequencerInbox.interface.encodeFunctionData('maxDelaySeconds'),
        decoder: (returnData: string) =>
          sequencerInbox.interface.decodeFunctionResult(
            'maxDelaySeconds',
            returnData
          )[0] as BigNumber,
      },
    ])

    if (!maxDelayBlocks) throw new ArbTsError('MaxDelayBlocks not fetched')
    if (!maxDelaySeconds) throw new ArbTsError('MaxDelaySeconds not fetched')

    const firstEligibleBlockNumber =
      currentBlock.number - maxDelayBlocks.toNumber()
    const firstEligibleTimestamp =
      currentBlock.timestamp - maxDelaySeconds.toNumber()

    const firstEligibleBlock = await this.findFirstBlockBelow(
      firstEligibleBlockNumber,
      firstEligibleTimestamp
    )

    return {
      endBlock: firstEligibleBlock.number,
      startBlock: firstEligibleBlock.number - blockNumbeRangeSize,
    }
  }

  /**
   * Find the event of the latest message that can be force include
   * @param searchRangeBlocks Defaults to 3 * 6545 ( = ~3 days) prior to the first eligble block
   * @returns Null if non can be found.
   */
  public async getForceIncludeableEvent(
    searchRangeBlocks: number = 3 * 6545
  ): Promise<ForceInclusionParams | null> {
    const bridge = Bridge__factory.connect(
      this.l2Network.ethBridge.bridge,
      this.l1Provider
    )

    const eFetcher = new EventFetcher(this.l1Provider)

    // events dont become eligible until they pass a delay
    // find a block range which will emit eligible events
    const blockRange = await this.getEligibleBlockRange(searchRangeBlocks)

    // get all the events in this range
    const events = await eFetcher.getEvents<Bridge, MessageDeliveredEvent>(
      bridge.address,
      Bridge__factory,
      b => b.filters.MessageDelivered(),
      { fromBlock: blockRange.startBlock, toBlock: blockRange.endBlock }
    )

    // no events appeared within that time period
    if (events.length === 0) return null

    // take the last event - as including this one will include all previous events
    const eventInfo = events[events.length - 1]
    const sequencerInbox = SequencerInbox__factory.connect(
      this.l2Network.ethBridge.sequencerInbox,
      this.l1Provider
    )
    // has the sequencer inbox already read this latest message
    const totalDelayedRead = await sequencerInbox.totalDelayedMessagesRead()
    if (totalDelayedRead.gt(eventInfo.event.messageIndex)) {
      // nothing to read - more delayed messages have been read than this current index
      return null
    }

    const delayedAcc = await bridge.inboxAccs(eventInfo.event.messageIndex)

    return { ...eventInfo, delayedAcc: delayedAcc }
  }

  /**
   * Force includes all eligible messages in the delayed inbox.
   * The inbox contract doesnt allow a message to be force-included
   * until after a delay period has been completed.
   * @param messageDeliveredEvent Provide this to include all messages up to this one. Responsibility is on the caller to check the eligibility of this event.
   * @returns The force include transaction, or null if no eligible message were found for inclusion
   */
  public async forceInclude<T extends ForceInclusionParams | undefined>(
    messageDeliveredEvent?: T,
    overrides?: Overrides
  ): Promise<
    // if a message delivered event was supplied then we'll definately return
    // a contract transaction or throw an error. If one isnt supplied then we may
    // find no eligible events, and so return null
    T extends ForceInclusionParams
      ? ContractTransaction
      : ContractTransaction | null
  >
  public async forceInclude<T extends ForceInclusionParams | undefined>(
    messageDeliveredEvent?: T,
    overrides?: Overrides
  ): Promise<ContractTransaction | null> {
    const sequencerInbox = SequencerInbox__factory.connect(
      this.l2Network.ethBridge.sequencerInbox,
      this.l1Signer
    )
    const eventInfo =
      messageDeliveredEvent || (await this.getForceIncludeableEvent())

    if (!eventInfo) return null
    const block = await this.l1Provider.getBlock(eventInfo.blockHash)

    const transactionReceipt = await this.l1Provider.getTransactionReceipt(
      eventInfo.transactionHash
    )
    return await sequencerInbox.functions.forceInclusion(
      eventInfo.event.messageIndex.add(1),
      eventInfo.event.kind,
      [eventInfo.blockNumber, block.timestamp],
      eventInfo.event.messageIndex,
      transactionReceipt.effectiveGasPrice,
      eventInfo.event.sender,
      eventInfo.event.messageDataHash,
      eventInfo.delayedAcc,
      // we need to pass in {} because if overrides is undefined it thinks we've provided too many params
      overrides || {}
    )
  }
}
