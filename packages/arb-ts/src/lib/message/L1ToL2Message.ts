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
import { Provider, Log } from '@ethersproject/abstract-provider'
import { Signer } from '@ethersproject/abstract-signer'
import { ContractTransaction, ContractReceipt } from '@ethersproject/contracts'
import { BigNumber } from '@ethersproject/bignumber'
import { constants, ethers } from 'ethers'

import { Inbox__factory } from '../abi/factories/Inbox__factory'
import { ArbRetryableTx__factory } from '../abi/factories/ArbRetryableTx__factory'
import { L1ERC20Gateway__factory } from '../abi/factories/L1ERC20Gateway__factory'

import {
  calculateRetryableTicketCreationHash,
  calculateL2MessageFromTicketTxnHash,
  L2TxnType,
} from './lib'
import { ARB_RETRYABLE_TX_ADDRESS } from '../precompile_addresses'
import {
  SignerProviderUtils,
  SignerOrProvider,
} from '../utils/signerOrProvider'
import { DepositInitiatedEvent } from '../abi/L1ERC20Gateway'

export class L1TransactionReceipt implements TransactionReceipt {
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
   * Get the numbers of any messages created by this transaction
   * @returns
   */
  public getMessageNumbers(): BigNumber[] {
    const iface = Inbox__factory.createInterface()
    const messageDelivered = iface.getEvent('InboxMessageDelivered')
    const messageDeliveredFromOrigin = iface.getEvent(
      'InboxMessageDeliveredFromOrigin'
    )
    const eventTopics = {
      InboxMessageDelivered: iface.getEventTopic(messageDelivered),
      InboxMessageDeliveredFromOrigin: iface.getEventTopic(
        messageDeliveredFromOrigin
      ),
    }
    const logs = this.logs.filter(
      log =>
        log.topics[0] === eventTopics.InboxMessageDelivered ||
        log.topics[0] === eventTopics.InboxMessageDeliveredFromOrigin
    )
    return logs.map(log => BigNumber.from(log.topics[1]))
  }

  /**
   * Get any l1tol2 messages created by this transaction
   * @param l2SignerOrProvider
   */
  public async getL1ToL2Messages<T extends SignerOrProvider>(
    l2SignerOrProvider: T
  ): Promise<L1ToL2MessageReaderOrWriter<T>[]>
  public async getL1ToL2Messages<T extends SignerOrProvider>(
    l2SignerOrProvider: T
  ): Promise<L1ToL2MessageReader[] | L1ToL2MessageWriter[]> {
    const provider = SignerProviderUtils.getProvider(l2SignerOrProvider)
    if (!provider) throw new Error('Signer not connected to provider.')

    const chainID = (await provider.getNetwork()).chainId.toString()

    const messageNumbers = this.getMessageNumbers()
    if (!messageNumbers || messageNumbers.length === 0)
      throw new Error(
        'No l1 to l2 messages found in L1 txn ' + this.transactionHash
      )

    return messageNumbers.map((mn: BigNumber) => {
      const ticketCreationHash = calculateRetryableTicketCreationHash({
        l2ChainId: BigNumber.from(chainID),
        messageNumber: mn,
      })
      return L1ToL2Message.fromL2Ticket(
        l2SignerOrProvider,
        ticketCreationHash,
        mn
      )
    })
  }

  /**
   * Get any deposit events created by this transaction
   * @returns
   */
  public getDepositEvents(): DepositInitiatedEvent['args'][] {
    const iface = L1ERC20Gateway__factory.createInterface()
    const event = iface.getEvent('DepositInitiated')
    const eventTopic = iface.getEventTopic(event)
    const logs = this.logs.filter(log => log.topics[0] === eventTopic)
    return logs.map(
      log => iface.parseLog(log).args as DepositInitiatedEvent['args']
    )
  }

  /**
   * Replaces the wait function with one that returns an L1TransactionReceipt
   * @param contractTransaction
   * @returns
   */
  public static swivelWait = (
    contractTransaction: ContractTransaction
  ): L1ContractTransaction => {
    const wait = contractTransaction.wait
    contractTransaction.wait = async (confirmations?: number) => {
      const result = await wait(confirmations)
      return new L1TransactionReceipt(result)
    }
    return contractTransaction as L1ContractTransaction
  }
}

export interface L1ContractTransaction extends ContractTransaction {
  wait(confirmations?: number): Promise<L1TransactionReceipt>
}

export enum L1ToL2MessageStatus {
  NOT_YET_CREATED,
  CREATION_FAILED,
  NOT_YET_REDEEMED, // i.e., autoredeem failed
  REDEEMED,
  EXPIRED, // canceled or timed out
}

export interface L1ToL2MessageReceipt {
  ticketCreationReceipt?: TransactionReceipt
  autoRedeemReceipt?: TransactionReceipt
  userTxnReceipt?: TransactionReceipt
  status: L1ToL2MessageStatus
}

/**
 * Conditional type for Signer or Provider. If T is of type Provider
 * then L1ToL2MessageReaderOrWriter<T> will be of type L1ToL2MessageReader.
 * If T is of type Signer then L1ToL2MessageReaderOrWriter<T> will be of
 * type L1ToL2MessageWriter.
 */
export type L1ToL2MessageReaderOrWriter<
  T extends SignerOrProvider
> = T extends Provider ? L1ToL2MessageReader : L1ToL2MessageWriter

export class L1ToL2Message {
  public static fromL2Ticket<T extends SignerOrProvider>(
    l2SignerOrProvider: T,
    l2TicketCreationHash: string,
    messageNumber: BigNumber
  ): L1ToL2MessageReaderOrWriter<T>
  public static fromL2Ticket<T extends SignerOrProvider>(
    l2SignerOrProvider: T,
    l2TicketCreationHash: string,
    messageNumber: BigNumber
  ): L1ToL2MessageReader | L1ToL2MessageWriter {
    return SignerProviderUtils.isSigner(l2SignerOrProvider)
      ? new L1ToL2MessageWriter(
          l2SignerOrProvider,
          l2TicketCreationHash,
          messageNumber
        )
      : new L1ToL2MessageReader(
          l2SignerOrProvider,
          l2TicketCreationHash,
          messageNumber
        )
  }
}

export class L1ToL2MessageReader extends L1ToL2Message {
  public constructor(
    private readonly l2Provider: Provider,
    public readonly l2TicketCreationTxnHash: string,
    public readonly messageNumber: BigNumber
  ) {
    super()
  }

  get autoRedeemHash(): string {
    return calculateL2MessageFromTicketTxnHash(
      this.l2TicketCreationTxnHash,
      L2TxnType.AUTO_REDEEM
    )
  }

  get userTxnHash(): string {
    return calculateL2MessageFromTicketTxnHash(
      this.l2TicketCreationTxnHash,
      L2TxnType.USER_TXN
    )
  }

  public getTicketCreationReceipt(): Promise<TransactionReceipt> {
    return this.l2Provider.getTransactionReceipt(this.l2TicketCreationTxnHash)
  }
  public getAutoRedeemReceipt(): Promise<TransactionReceipt> {
    return this.l2Provider.getTransactionReceipt(this.autoRedeemHash)
  }
  public getUserTxnReceipt(): Promise<TransactionReceipt> {
    return this.l2Provider.getTransactionReceipt(this.userTxnHash)
  }

  public async isExpired(): Promise<boolean> {
    return (await this.getTimeout()).eq(constants.Zero)
  }

  private async receiptsToStatus(
    ticketCreationReceipt: TransactionReceipt,
    userTxnReceipt: TransactionReceipt
  ): Promise<L1ToL2MessageStatus> {
    if (userTxnReceipt && userTxnReceipt.status === 1) {
      return L1ToL2MessageStatus.REDEEMED
    }

    if (!ticketCreationReceipt) {
      return L1ToL2MessageStatus.NOT_YET_CREATED
    }
    if (ticketCreationReceipt.status === 0) {
      return L1ToL2MessageStatus.CREATION_FAILED
    }
    if (await this.isExpired()) {
      return L1ToL2MessageStatus.EXPIRED
    }
    // we could sanity check that autoredeem failed, but we don't need to
    return L1ToL2MessageStatus.NOT_YET_REDEEMED
  }

  public async status(): Promise<L1ToL2MessageStatus> {
    const userTxnReceipt = await this.l2Provider.getTransactionReceipt(
      this.userTxnHash
    )

    const ticketCreationReceipt = await this.getTicketCreationReceipt()
    return this.receiptsToStatus(userTxnReceipt, ticketCreationReceipt)
  }

  public async wait(
    timeout = 900000,
    confirmations?: number
  ): Promise<L1ToL2MessageReceipt> {
    const ticketCreationReceipt = await this.l2Provider.waitForTransaction(
      this.l2TicketCreationTxnHash,
      confirmations,
      timeout
    )

    const autoRedeemReceipt = await this.l2Provider.waitForTransaction(
      this.autoRedeemHash,
      confirmations,
      3000 // autoredeem gets attempted immediately after ticket creation, but could never get attempted if not calldata; we leave a few seconds of buffer
    )

    const userTxnReceipt = await this.getUserTxnReceipt()

    return {
      ticketCreationReceipt,
      autoRedeemReceipt,
      userTxnReceipt,
      status: await this.receiptsToStatus(
        ticketCreationReceipt,
        userTxnReceipt
      ),
    }
  }

  public getTimeout(): Promise<BigNumber> {
    const arbRetryableTx = ArbRetryableTx__factory.connect(
      ARB_RETRYABLE_TX_ADDRESS,
      this.l2Provider
    )
    return arbRetryableTx.getTimeout(this.userTxnHash)
  }
  public getBeneficiary(): Promise<string> {
    const arbRetryableTx = ArbRetryableTx__factory.connect(
      ARB_RETRYABLE_TX_ADDRESS,
      this.l2Provider
    )
    return arbRetryableTx.getBeneficiary(this.userTxnHash)
  }
}

export class L1ToL2MessageWriter extends L1ToL2MessageReader {
  public constructor(
    private readonly l2Signer: Signer,
    l2TicketCreationTxnHash: string,
    messageNumber: BigNumber
  ) {
    super(l2Signer.provider!, l2TicketCreationTxnHash, messageNumber)
    if (!l2Signer.provider) throw new Error('Signer not connected to provider.')
  }

  /**
   * Checks the status of the message and only try's to redeem if
   * in the correct state. Safe to call on an already redeemed message
   */
  public async redeemSafe(
    waitTimeForL2Receipt = 900000 // 15 mins
  ): Promise<ContractTransaction> {
    console.log('waiting for retryable ticket...', this.userTxnHash)
    const result = await this.wait(waitTimeForL2Receipt)
    if (result.ticketCreationReceipt?.status == 0) {
      console.warn(
        'retryable ticket failed',
        result.ticketCreationReceipt.transactionHash
      )
      throw new Error('l2 txn failed')
    }
    return await this.redeem()
  }

  public redeem(): Promise<ContractTransaction> {
    const arbRetryableTx = ArbRetryableTx__factory.connect(
      ARB_RETRYABLE_TX_ADDRESS,
      this.l2Signer
    )
    return arbRetryableTx.redeem(this.userTxnHash)
  }

  public async cancelSafe(
    waitTimeForL2Receipt = 900000
  ): Promise<ContractTransaction> {
    const result = await this.wait(waitTimeForL2Receipt)

    if (result.ticketCreationReceipt?.status == 1) {
      throw new Error(
        `Can't cancel retryable, it's already been redeemed: ${result.ticketCreationReceipt?.transactionHash}`
      )
    }
    console.log(`Hasn't been redeemed yet, calling cancel now`)
    return await this.cancel()
  }

  public cancel(): Promise<ContractTransaction> {
    const arbRetryableTx = ArbRetryableTx__factory.connect(
      ARB_RETRYABLE_TX_ADDRESS,
      this.l2Signer
    )
    return arbRetryableTx.cancel(this.userTxnHash)
  }
}
