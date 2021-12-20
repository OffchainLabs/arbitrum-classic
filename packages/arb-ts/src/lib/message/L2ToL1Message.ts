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

import { NodeInterface__factory } from '../abi/factories/NodeInterface__factory'
import {
  L2ToL1EventResult,
  MessageBatchProofInfo,
  OutgoingMessageState,
} from '../dataEntities'
import {
  ARB_SYS_ADDRESS,
  NODE_INTERFACE_ADDRESS,
} from '../precompile_addresses'
import networks, { Network } from '../networks'
import { TransactionReceipt } from '@ethersproject/providers'
import { Provider, Filter } from '@ethersproject/abstract-provider'
import { Signer } from '@ethersproject/abstract-signer'
import { BigNumber } from '@ethersproject/bignumber'
import {
  SignerOrProviderUtils,
  SignerOrProvider,
} from '../utils/signerOrProvider'
import { wait } from '../utils/lib'

import { Log } from '@ethersproject/abstract-provider'
import { Outbox__factory } from '../abi/factories/Outbox__factory'
import { IOutbox__factory } from '../abi/factories/IOutbox__factory'
import { ArbSys__factory } from '../abi/factories/ArbSys__factory'
import { hexZeroPad } from '@ethersproject/bytes'

export class L2TransactionReceipt implements TransactionReceipt {
  public readonly to: string
  public readonly from: string
  public readonly contractAddress: string
  public readonly transactionIndex: number
  public readonly root?: string
  public readonly gasUsed: BigNumber
  public readonly logsBloom: string
  public readonly blockHash: string
  public readonly transactionHash: string
  public readonly logs: Array<Log>
  public readonly blockNumber: number
  public readonly confirmations: number
  public readonly cumulativeGasUsed: BigNumber
  public readonly effectiveGasPrice: BigNumber
  public readonly byzantium: boolean
  public readonly type: number
  public readonly status?: number

  constructor(tx: TransactionReceipt) {
    this.to = tx.to
    this.from = tx.from
    this.contractAddress = tx.contractAddress
    this.transactionIndex = tx.transactionIndex
    this.root = tx.root
    this.gasUsed = tx.gasUsed
    this.logsBloom = tx.logsBloom
    this.blockHash = tx.blockHash
    this.transactionHash = tx.transactionHash
    this.logs = tx.logs
    this.blockNumber = tx.blockNumber
    this.confirmations = tx.confirmations
    this.cumulativeGasUsed = tx.cumulativeGasUsed
    this.effectiveGasPrice = tx.effectiveGasPrice
    this.byzantium = tx.byzantium
    this.type = tx.type
    this.status = tx.status
  }

  public getL2ToL1Events() {
    const iface = ArbSys__factory.createInterface()
    const l2ToL1Event = iface.getEvent('L2ToL1Transaction')
    const eventTopic = iface.getEventTopic(l2ToL1Event)

    const logs = this.logs.filter(log => log.topics[0] === eventTopic)

    return logs.map(
      log => (iface.parseLog(log).args as unknown) as L2ToL1EventResult
    )
  }

  private getOutboxAddr(network: Network, batchNumber: BigNumber) {
    // CHRIS: add the old network to the networks object, and look it up here
    // CHRIS: null check? this shouldn't be possible
    // CHRIS: disable linting by just using the batchnumber here
    batchNumber
    return network.ethBridge!.outbox
  }

  public async getL2ToL1Messages<T extends SignerOrProvider>(
    l1SignerOrProvider: T
  ): Promise<L2ToL1MessageReaderOrWriter<T>[]>
  public async getL2ToL1Messages<T extends SignerOrProvider>(
    l1SignerOrProvider: T
  ): Promise<L2ToL1MessageReader[] | L2ToL1MessageWriter[]> {
    const provider = SignerOrProviderUtils.getProvider(l1SignerOrProvider)
    if (!provider) throw new Error('Signer not connected to provider.')

    const providerNetwork = await provider.getNetwork()
    const arbNetwork = networks[providerNetwork.chainId]
    return this.getL2ToL1Events().map(log => {
      const outboxAddr = this.getOutboxAddr(arbNetwork, log.batchNumber)

      return L2ToL1Message.fromBatchNumber(
        l1SignerOrProvider,
        outboxAddr,
        log.batchNumber,
        log.indexInBatch
      )
    })
  }
}

/**
 * Conditional type for Signer or Provider. If T is of type Provider
 * then L2ToL1MessageReaderOrWriter<T> will be of type L2ToL1MessageReader.
 * If T is of type Signer then L2ToL1MessageReaderOrWriter<T> will be of
 * type L2ToL1MessageWriter.
 */
export type L2ToL1MessageReaderOrWriter<
  T extends SignerOrProvider
> = T extends Provider ? L2ToL1MessageReader : L2ToL1MessageWriter

export class L2ToL1Message {
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
    return l1SignerOrProvider instanceof Provider
      ? new L2ToL1MessageReader(
          l1SignerOrProvider,
          outboxAddress,
          batchNumber,
          indexInBatch
        )
      : new L2ToL1MessageWriter(
          l1SignerOrProvider,
          outboxAddress,
          batchNumber,
          indexInBatch
        )
  }

  // CHRIS: consider what else to move onto the base
  public static async getL2ToL1MessageLogs(
    l2Provider: Provider,
    filter: Filter
  ) {
    const iface = ArbSys__factory.createInterface()
    const event = iface.getEvent('L2ToL1Transaction')
    const eventTopic = iface.getEventTopic(event)

    const topics = filter.topics ? filter.topics : []
    // CHRIS: keep the warn?
    // if (!filter.fromBlock && !filter.toBlock)
    // console.warn('Attempting to query from 0 to block latest')
    const logs = await l2Provider.getLogs({
      address: ARB_SYS_ADDRESS,
      topics: [eventTopic, ...topics],
      fromBlock: filter.fromBlock || 0,
      toBlock: filter.toBlock || 'latest',
    })

    return logs.map(
      log => iface.parseLog(log).args as unknown as L2ToL1EventResult
    )
  }

  public static async getL2ToL1MessageLog(
    l2Provider: Provider,
    batchNumber: BigNumber,
    indexInBatch: BigNumber
  ) {
    const batch = await this.getL2ToL1MessageLogs(l2Provider, {
      topics: [null, null, hexZeroPad(batchNumber.toHexString(), 32)],
    })

    const indexItem = batch.filter(b => b.indexInBatch.eq(indexInBatch))
    if (indexItem.length === 1) {
      return indexItem[0]
    } else if (indexItem.length > 1) {
      // CHRIS: warn, error?
    }

    return null
  }
}

export class L2ToL1MessageReader extends L2ToL1Message {
  constructor(
    protected readonly l1Provider: Provider,
    protected readonly outboxAddress: string,
    public readonly batchNumber: BigNumber,
    public readonly indexInBatch: BigNumber
  ) {
    super()
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
      const err = e as any
      const actualError =
        err && (err.message || (err.error && err.error.message))
      if (actualError.includes(expectedError)) {
        console.log('Withdrawal detected, but batch not created yet.')
      } else {
        console.log("Withdrawal proof didn't work. Not sure why")
        console.log(e)
      }
    }
    return null
  }

  public async tryGetProof(l2Provider: Provider) {
    return L2ToL1MessageReader.tryGetProof(
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
      const e = err as any
      if (e && e.message && e.message.toString().includes('ALREADY_SPENT')) {
        return true
      }
      if (e && e.message && e.message.toString().includes('NO_OUTBOX_ENTRY')) {
        return false
      }
      throw e
    }
  }

  public async status(proofInfo: MessageBatchProofInfo | null) {
    try {
      if (proofInfo) {
        const messageExecuted = await this.hasExecuted(proofInfo)
        if (messageExecuted) {
          return OutgoingMessageState.EXECUTED
        }
      }

      const outboxEntryExists = await this.outboxEntryExists()

      return outboxEntryExists
        ? OutgoingMessageState.CONFIRMED
        : OutgoingMessageState.UNCONFIRMED
    } catch (e) {
      // CHRIS: discuss all the console logs going on in here and elsewhere, should we keep them?
      // CHRIS: this error needs updating. also 666?
      console.warn('666: error in getOutGoingMessageState:', e)
      return OutgoingMessageState.NOT_FOUND
    }
  }

  public async waitUntilOutboxEntryCreated(retryDelay = 500): Promise<void> {
    // CHRIS: should we even be doing this? what's the use case? it could be a long wait

    const exists = await this.outboxEntryExists()
    if (exists) {
      console.log('Found outbox entry!')
      return
    } else {
      console.log("can't find entry, lets wait a bit?")

      await wait(retryDelay)
      console.log('Starting new attempt')
      await this.waitUntilOutboxEntryCreated(retryDelay)
    }
  }
}

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
   * @returns
   */
  public async execute(proofInfo: MessageBatchProofInfo) {
    await this.waitUntilOutboxEntryCreated()

    const outbox = Outbox__factory.connect(this.outboxAddress, this.l1Signer)
    try {
      // TODO: wait until assertion is confirmed before execute
      // We can predict and print number of missing blocks
      // if not challenged
      const outboxExecute = await outbox.functions.executeTransaction(
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
      console.log(`Transaction hash: ${outboxExecute.hash}`)
      return outboxExecute
    } catch (e) {
      console.log('failed to execute tx in layer 1')
      console.log(e)
      // TODO: should we just try again after delay instead of throwing?
      throw e
    }
  }
}
