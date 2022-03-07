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
import { ContractTransaction, Overrides } from 'ethers'

import { Bridge } from '../abi/Bridge'
import { Bridge__factory } from '../abi/factories/Bridge__factory'
import { SequencerInbox } from '../abi/SequencerInbox'
import { SequencerInbox__factory } from '../abi/factories/SequencerInbox__factory'

import { MessageDeliveredEvent } from '../abi/Bridge'
import { l1Networks, L2Network } from '../dataEntities/networks'
import { SignerProviderUtils } from '../dataEntities/signerOrProvider'
import { FetchedEvent, EventFetcher } from '../utils/eventFetcher'
import { MultiCaller, CallInput } from '../utils/multicall'
import { ArbTsError } from '../dataEntities/errors'

type ForceInclusionParams = FetchedEvent<MessageDeliveredEvent> & {
  delayedAcc: string
}

/**
 * Tools for interacting with the inbox and bridge contracts
 */
export class InboxTools {
  private readonly l1Provider
  private readonly l1Network

  constructor(
    private readonly l1Signer: Signer,
    private readonly l2Network: L2Network
  ) {
    this.l1Provider = SignerProviderUtils.getProviderOrThrow(this.l1Signer)
    this.l1Network = l1Networks[l2Network.partnerChainID]
    if (!this.l1Network)
      throw new ArbTsError(
        `L1Network not found for chain id: ${l2Network.partnerChainID}.`
      )
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

    const diffBlocks = Math.max(Math.ceil(diff / this.l1Network.blockTime), 10)

    return await this.findFirstBlockBelow(
      blockNumber - diffBlocks,
      blockTimestamp
    )
  }

  /**
   * Get a range of blocks within messages eligible for force inclusion emitted events
   * @param blockNumberRangeSize
   * @returns
   */
  private async getForceIncludableBlockRange(blockNumberRangeSize: number) {
    const sequencerInbox = SequencerInbox__factory.connect(
      this.l2Network.ethBridge.sequencerInbox,
      this.l1Provider
    )

    const multicall = await MultiCaller.fromProvider(this.l1Provider)
    const multicallInput: [
      CallInput<Awaited<ReturnType<SequencerInbox['maxDelayBlocks']>>>,
      CallInput<Awaited<ReturnType<SequencerInbox['maxDelaySeconds']>>>,
      ReturnType<MultiCaller['getBlockNumberInput']>,
      ReturnType<MultiCaller['getCurrentBlockTimestampInput']>
    ] = [
      {
        targetAddr: sequencerInbox.address,
        encoder: () =>
          sequencerInbox.interface.encodeFunctionData('maxDelayBlocks'),
        decoder: (returnData: string) =>
          sequencerInbox.interface.decodeFunctionResult(
            'maxDelayBlocks',
            returnData
          )[0],
      },
      {
        targetAddr: sequencerInbox.address,
        encoder: () =>
          sequencerInbox.interface.encodeFunctionData('maxDelaySeconds'),
        decoder: (returnData: string) =>
          sequencerInbox.interface.decodeFunctionResult(
            'maxDelaySeconds',
            returnData
          )[0],
      },
      multicall.getBlockNumberInput(),
      multicall.getCurrentBlockTimestampInput(),
    ]

    const [
      maxDelayBlocks,
      maxDelaySeconds,
      currentBlockNumber,
      currentBlockTimestamp,
    ] = await multicall.multiCall(multicallInput, true)

    const firstEligibleBlockNumber =
      currentBlockNumber.toNumber() - maxDelayBlocks.toNumber()
    const firstEligibleTimestamp =
      currentBlockTimestamp.toNumber() - maxDelaySeconds.toNumber()

    const firstEligibleBlock = await this.findFirstBlockBelow(
      firstEligibleBlockNumber,
      firstEligibleTimestamp
    )

    return {
      endBlock: firstEligibleBlock.number,
      startBlock: firstEligibleBlock.number - blockNumberRangeSize,
    }
  }

  /**
   * Look for force includable events in the search range blocks, if no events are found the search range is
   * increased incrementally up to the max search range blocks.
   * @param bridge
   * @param searchRangeBlocks
   * @param maxSearchRangeBlocks
   * @returns
   */
  private async getEventsAndIncreaseRange(
    bridge: Bridge,
    searchRangeBlocks: number,
    maxSearchRangeBlocks: number,
    rangeMultiplier: number
  ): Promise<FetchedEvent<MessageDeliveredEvent>[]> {
    const eFetcher = new EventFetcher(this.l1Provider)

    // events don't become eligible until they pass a delay
    // find a block range which will emit eligible events
    const cappedSearchRangeBlocks = Math.min(
      searchRangeBlocks,
      maxSearchRangeBlocks
    )
    const blockRange = await this.getForceIncludableBlockRange(
      cappedSearchRangeBlocks
    )

    // get all the events in this range
    const events = await eFetcher.getEvents(
      bridge.address,
      Bridge__factory,
      b => b.filters.MessageDelivered(),
      {
        fromBlock: blockRange.startBlock,
        toBlock: blockRange.endBlock,
      }
    )

    if (events.length !== 0) return events
    else if (cappedSearchRangeBlocks === maxSearchRangeBlocks) return []
    else {
      return await this.getEventsAndIncreaseRange(
        bridge,
        searchRangeBlocks * rangeMultiplier,
        maxSearchRangeBlocks,
        rangeMultiplier
      )
    }
  }

  /**
   * Find the event of the latest message that can be force include
   * @param maxSearchRangeBlocks The max range of blocks to search in.
   * Defaults to 3 * 6545 ( = ~3 days) prior to the first eligble block
   * @param startSearchRangeBlocks The start range of block to search in.
   * Moves incrementally up to the maxSearchRangeBlocks. Defaults to 100;
   * @param rangeMultiplier The multiplier to use when increasing the block range
   * Defaults to 2.
   * @returns Null if non can be found.
   */
  public async getForceIncludableEvent(
    maxSearchRangeBlocks: number = 3 * 6545,
    startSearchRangeBlocks = 100,
    rangeMultipler = 2
  ): Promise<ForceInclusionParams | null> {
    const bridge = Bridge__factory.connect(
      this.l2Network.ethBridge.bridge,
      this.l1Provider
    )

    // events dont become eligible until they pass a delay
    // find a block range which will emit eligible events
    const events = await this.getEventsAndIncreaseRange(
      bridge,
      startSearchRangeBlocks,
      maxSearchRangeBlocks,
      rangeMultipler
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
    // if a message delivered event was supplied then we'll definitely return
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
      messageDeliveredEvent || (await this.getForceIncludableEvent())

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
