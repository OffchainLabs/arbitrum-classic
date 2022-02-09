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

import { TransactionReceipt } from '@ethersproject/providers'
import { BigNumber } from '@ethersproject/bignumber'
import { Log } from '@ethersproject/abstract-provider'
import { ArbSys__factory } from '../abi'
import { L2ToL1TransactionEvent } from '../abi/ArbSys'
import { ContractTransaction, providers } from 'ethers'
import { L2Network } from '../dataEntities/networks'
import { ArbTsError } from '../dataEntities/errors'
import {
  SignerProviderUtils,
  SignerOrProvider,
} from '../dataEntities/signerOrProvider'
import {
  L2ToL1MessageReader,
  L2ToL1MessageReaderOrWriter,
  L2ToL1Message,
  L2ToL1MessageWriter,
} from './L2ToL1Message'
import { getRawArbTransactionReceipt } from '../..'

export interface L2ContractTransaction extends ContractTransaction {
  wait(confirmations?: number): Promise<L2TransactionReceipt>
}

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

  /**
   * Get an L2ToL1Transaction events created by this transaction
   * @returns
   */
  public getL2ToL1Events(): L2ToL1TransactionEvent['args'][] {
    const iface = ArbSys__factory.createInterface()
    const l2ToL1Event = iface.getEvent('L2ToL1Transaction')
    const eventTopic = iface.getEventTopic(l2ToL1Event)
    const logs = this.logs.filter(log => log.topics[0] === eventTopic)

    return logs.map(
      log => iface.parseLog(log).args as L2ToL1TransactionEvent['args']
    )
  }

  private getOutboxAddr(network: L2Network, batchNumber: BigNumber) {
    // find the outbox where the activation batch number of the next outbox
    // is greater than the supplied batch
    const res = Object.entries(network.ethBridge.outboxes)
      .sort((a, b) => {
        if (a[1].lt(b[1])) return -1
        else if (a[1].eq(b[1])) return 0
        else return 1
      })
      .find(
        (_, index, array) =>
          array[index + 1] === undefined || array[index + 1][1].gt(batchNumber)
      )

    if (!res) {
      throw new ArbTsError(
        `No outbox found for batch number: ${batchNumber.toString()} on network: ${
          network.chainID
        }.`
      )
    }

    return res[0]
  }

  /**
   * Get any l2-to-l1-messages created by this transaction
   * @param l2SignerOrProvider
   */
  public async getL2ToL1Messages<T extends SignerOrProvider>(
    l1SignerOrProvider: T,
    l2Network: L2Network
  ): Promise<L2ToL1MessageReaderOrWriter<T>[]>
  public async getL2ToL1Messages<T extends SignerOrProvider>(
    l1SignerOrProvider: T,
    l2Network: L2Network
  ): Promise<L2ToL1MessageReader[] | L2ToL1MessageWriter[]> {
    const provider = SignerProviderUtils.getProvider(l1SignerOrProvider)
    if (!provider) throw new Error('Signer not connected to provider.')

    return this.getL2ToL1Events().map(log => {
      const outboxAddr = this.getOutboxAddr(l2Network, log.batchNumber)

      return L2ToL1Message.fromBatchNumber(
        l1SignerOrProvider,
        outboxAddr,
        log.batchNumber,
        log.indexInBatch
      )
    })
  }

  /**
   * Whether the data associated with this transaction has been
   * made available on L1
   */
  public async isDataAvailable(
    l2Provider: providers.JsonRpcProvider,
    l1Provider: providers.JsonRpcProvider
  ): Promise<boolean> {
    const arbReceipt = await getRawArbTransactionReceipt(
      l2Provider,
      this.transactionHash,
      l1Provider
    )

    // Data is made available in batches, if the batch info is
    // available then so is the tx data
    return !!arbReceipt?.l1InboxBatchInfo
  }

  /**
   * Replaces the wait function with one that returns an L2TransactionReceipt
   * @param contractTransaction
   * @returns
   */
  public static monkeyPatchWait = (
    contractTransaction: ContractTransaction
  ): L2ContractTransaction => {
    const wait = contractTransaction.wait
    contractTransaction.wait = async (confirmations?: number) => {
      // we ignore the confirmations for now since L2 transactions shouldn't re-org
      // in future we should give users a more fine grained way to check the finality of
      // an l2 transaction - check if a batch is on L1, if an assertion has been made, and if
      // it has been confirmed.
      const result = await wait()
      return new L2TransactionReceipt(result)
    }
    return contractTransaction as L2ContractTransaction
  }
}
