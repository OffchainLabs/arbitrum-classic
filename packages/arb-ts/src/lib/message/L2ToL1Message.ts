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

import {
  ARB_SYS_ADDRESS,
  NODE_INTERFACE_ADDRESS,
} from '../dataEntities/constants'
import { Provider } from '@ethersproject/abstract-provider'
import { Signer } from '@ethersproject/abstract-signer'
import { BigNumber } from '@ethersproject/bignumber'
import { BlockTag } from '@ethersproject/abstract-provider'

import {
  ArbSys__factory,
  IOutbox__factory,
  Outbox__factory,
  NodeInterface__factory,
} from '../abi'
import { L2ToL1TransactionEvent } from '../abi/ArbSys'
import { ContractTransaction } from 'ethers'
import { EventFetcher } from '../utils/eventFetcher'
import { ArbTsError } from '../dataEntities/errors'
import {
  SignerProviderUtils,
  SignerOrProvider,
} from '../dataEntities/signerOrProvider'
import { wait } from '../utils/lib'

export interface MessageBatchProofInfo {
  /**
   * Merkle proof of message inclusion in outbox entry
   */
  proof: string[]

  /**
   * Merkle path to message
   */
  path: BigNumber

  /**
   * Sender of original message (i.e., caller of ArbSys.sendTxToL1)
   */
  l2Sender: string

  /**
   * Destination address for L1 contract call
   */
  l1Dest: string

  /**
   * L2 block number at which sendTxToL1 call was made
   */
  l2Block: BigNumber

  /**
   * L1 block number at which sendTxToL1 call was made
   */
  l1Block: BigNumber

  /**
   * L2 Timestamp at which sendTxToL1 call was made
   */
  timestamp: BigNumber

  /**
   * Value in L1 message in wei
   */
  amount: BigNumber

  /**
   * ABI-encoded L1 message data
   */
  calldataForL1: string
}

export enum L2ToL1MessageStatus {
  /**
   * No corresponding L2ToL1Event emitted
   */
  NOT_FOUND,
  /**
   * ArbSys.sendTxToL1 called, but assertion not yet confirmed
   */
  UNCONFIRMED,
  /**
   * Assertion for outgoing message confirmed, but message not yet executed
   */
  CONFIRMED,
  /**
   * Outgoing message executed (terminal state)
   */
  EXECUTED,
}

/**
 * Conditional type for Signer or Provider. If T is of type Provider
 * then L2ToL1MessageReaderOrWriter<T> will be of type L2ToL1MessageReader.
 * If T is of type Signer then L2ToL1MessageReaderOrWriter<T> will be of
 * type L2ToL1MessageWriter.
 */
export type L2ToL1MessageReaderOrWriter<T extends SignerOrProvider> =
  T extends Provider ? L2ToL1MessageReader : L2ToL1MessageWriter

export class L2ToL1Message {
  /**
   * The number of the batch this message is part of
   */
  public readonly batchNumber: BigNumber

  /**
   * The index of this message in the batch
   */
  public readonly indexInBatch: BigNumber

  protected constructor(batchNumber: BigNumber, indexInBatch: BigNumber) {
    this.batchNumber = batchNumber
    this.indexInBatch = indexInBatch
  }

  public static fromBatchNumber<T extends SignerOrProvider>(
    l1SignerOrProvider: T,
    outboxAddress: string,
    batchNumber: BigNumber,
    indexInBatch: BigNumber
  ): L2ToL1MessageReaderOrWriter<T>
  public static fromBatchNumber<T extends SignerOrProvider>(
    l1SignerOrProvider: T,
    outboxAddress: string,
    batchNumber: BigNumber,
    indexInBatch: BigNumber
  ): L2ToL1MessageReader | L2ToL1MessageWriter {
    return SignerProviderUtils.isSigner(l1SignerOrProvider)
      ? new L2ToL1MessageWriter(
          l1SignerOrProvider,
          outboxAddress,
          batchNumber,
          indexInBatch
        )
      : new L2ToL1MessageReader(
          l1SignerOrProvider,
          outboxAddress,
          batchNumber,
          indexInBatch
        )
  }

  public static async getL2ToL1MessageLogs(
    l2Provider: Provider,
    filter: { fromBlock: BlockTag; toBlock: BlockTag },
    batchNumber?: BigNumber,
    destination?: string,
    uniqueId?: BigNumber,
    indexInBatch?: BigNumber
  ): Promise<L2ToL1TransactionEvent['args'][]> {
    const eventFetcher = new EventFetcher(l2Provider)
    const events = (
      await eventFetcher.getEvents(
        ARB_SYS_ADDRESS,
        ArbSys__factory,
        t =>
          t.filters.L2ToL1Transaction(null, destination, uniqueId, batchNumber),
        filter
      )
    ).map(l => l.event)

    if (indexInBatch) {
      const indexItems = events.filter(b => b.indexInBatch.eq(indexInBatch))
      if (indexItems.length === 1) {
        return indexItems
      } else if (indexItems.length > 1) {
        throw new ArbTsError('More than one indexed item found in batch.')
      } else return []
    } else return events
  }
}

/**
 * Provides read-only access for l2-to-l1-messages
 */
export class L2ToL1MessageReader extends L2ToL1Message {
  constructor(
    protected readonly l1Provider: Provider,
    protected readonly outboxAddress: string,
    batchNumber: BigNumber,
    indexInBatch: BigNumber
  ) {
    super(batchNumber, indexInBatch)
  }

  private async outboxEntryExists() {
    const outbox = IOutbox__factory.connect(this.outboxAddress, this.l1Provider)
    return await outbox.outboxEntryExists(this.batchNumber)
  }

  public static async tryGetProof(
    l2Provider: Provider,
    batchNumber: BigNumber,
    indexInBatch: BigNumber
  ): Promise<MessageBatchProofInfo | null> {
    const nodeInterface = NodeInterface__factory.connect(
      NODE_INTERFACE_ADDRESS,
      l2Provider
    )
    try {
      return nodeInterface.lookupMessageBatchProof(batchNumber, indexInBatch)
    } catch (e) {
      const expectedError = "batch doesn't exist"
      const err = e as Error & { error: Error }
      const actualError =
        err && (err.message || (err.error && err.error.message))
      if (actualError.includes(expectedError)) return null
      else throw e
    }
  }

  /**
   * Get the execution proof for this message. Returns null if the batch does not exist yet.
   * @param l2Provider
   * @returns
   */
  public async tryGetProof(
    l2Provider: Provider
  ): Promise<MessageBatchProofInfo | null> {
    return await L2ToL1MessageReader.tryGetProof(
      l2Provider,
      this.batchNumber,
      this.indexInBatch
    )
  }

  /**
   * Check if given outbox message has already been executed
   */
  public async hasExecuted(proofInfo: MessageBatchProofInfo): Promise<boolean> {
    const outbox = Outbox__factory.connect(this.outboxAddress, this.l1Provider)
    try {
      await outbox.callStatic.executeTransaction(
        this.batchNumber,
        proofInfo.proof,
        proofInfo.path,
        proofInfo.l2Sender,
        proofInfo.l1Dest,
        proofInfo.l2Block,
        proofInfo.l1Block,
        proofInfo.timestamp,
        proofInfo.amount,
        proofInfo.calldataForL1
      )
      return false
    } catch (err) {
      const e = err as Error
      if (e?.message?.toString().includes('ALREADY_SPENT')) return true
      if (e?.message?.toString().includes('NO_OUTBOX_ENTRY')) return false
      throw e
    }
  }

  /**
   * Get the status of this message
   * In order to check if the message has been executed proof info must be provided.
   * @param proofInfo
   * @returns
   */
  public async status(
    proofInfo: MessageBatchProofInfo | null
  ): Promise<L2ToL1MessageStatus> {
    try {
      if (proofInfo) {
        const messageExecuted = await this.hasExecuted(proofInfo)
        if (messageExecuted) {
          return L2ToL1MessageStatus.EXECUTED
        }
      }

      const outboxEntryExists = await this.outboxEntryExists()
      return outboxEntryExists
        ? L2ToL1MessageStatus.CONFIRMED
        : L2ToL1MessageStatus.UNCONFIRMED
    } catch (e) {
      console.warn('666: error in fetching status:', e)
      return L2ToL1MessageStatus.NOT_FOUND
    }
  }

  /**
   * Waits until the outbox entry has been created, and will not return until it has been.
   * WARNING: Outbox entries are only created when the corresponding node is confirmed. Which
   * can take 1 week+, so waiting here could be a very long operation.
   * @param retryDelay
   * @returns
   */
  public async waitUntilOutboxEntryCreated(retryDelay = 500): Promise<void> {
    const exists = await this.outboxEntryExists()
    if (exists) {
      return
    } else {
      await wait(retryDelay)
      await this.waitUntilOutboxEntryCreated(retryDelay)
    }
  }
}

/**
 * Provides read and write access for l2-to-l1-messages
 */
export class L2ToL1MessageWriter extends L2ToL1MessageReader {
  constructor(
    private readonly l1Signer: Signer,
    outboxAddress: string,
    batchNumber: BigNumber,
    indexInBatch: BigNumber
  ) {
    super(l1Signer.provider!, outboxAddress, batchNumber, indexInBatch)
  }

  /**
   * Executes the L2ToL1Message on L1.
   * Will throw an error if the outbox entry has not been created, which happens when the
   * corresponding assertion is confirmed.
   * @returns
   */
  public async execute(
    proofInfo: MessageBatchProofInfo
  ): Promise<ContractTransaction> {
    const status = await this.status(proofInfo)
    if (status !== L2ToL1MessageStatus.CONFIRMED) {
      throw new ArbTsError(
        `Cannot execute message. Status is: ${status} but must be ${L2ToL1MessageStatus.CONFIRMED}.`
      )
    }

    const outbox = Outbox__factory.connect(this.outboxAddress, this.l1Signer)
    // We can predict and print number of missing blocks
    // if not challenged
    return await outbox.functions.executeTransaction(
      this.batchNumber,
      proofInfo.proof,
      proofInfo.path,
      proofInfo.l2Sender,
      proofInfo.l1Dest,
      proofInfo.l2Block,
      proofInfo.l1Block,
      proofInfo.timestamp,
      proofInfo.amount,
      proofInfo.calldataForL1
    )
  }
}
