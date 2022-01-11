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
import { ContractTransaction } from '@ethersproject/contracts'
import { BigNumber } from '@ethersproject/bignumber'
import { constants } from 'ethers'
import { keccak256 } from '@ethersproject/keccak256'
import { concat, zeroPad } from '@ethersproject/bytes'

import { Inbox__factory } from '../abi/factories/Inbox__factory'
import { ArbRetryableTx__factory } from '../abi/factories/ArbRetryableTx__factory'
import { L1ERC20Gateway__factory } from '../abi/factories/L1ERC20Gateway__factory'
import { ARB_RETRYABLE_TX_ADDRESS } from '../constants'
import {
  SignerProviderUtils,
  SignerOrProvider,
} from '../utils/signerOrProvider'
import { DepositInitiatedEvent } from '../abi/L1ERC20Gateway'
import { ArbTsError } from '../errors'

export enum L2TxnType {
  L2_TX = 0,
  AUTO_REDEEM = 1,
}

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
      const ticketCreationHash = L1ToL2Message.calculateRetryableTicketId(
        BigNumber.from(chainID),
        mn
      )
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
  public static monkeyPatchWait = (
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
  /**
   * The retryable ticket has yet to be created
   */
  NOT_YET_CREATED,
  /**
   * An attempt was made to create the retryable ticket, but it failed.
   * This could be due to not enough submission cost being paid by the L1 transaction
   */
  CREATION_FAILED,
  /**
   * The retryable ticket has been created but has not been redeemed. Since auto redeem occurs
   * when the retryable ticket was created, this means that the auto-redeem failed.
   */
  NOT_YET_REDEEMED,
  /**
   * The retryable ticket has been redeemed (either by auto, or manually) and the
   * l2 transaction has been executed
   */
  REDEEMED,
  /**
   * The message has either expired or has been canceled. It can no longer be redeemed.
   */
  EXPIRED,
}

export interface L1ToL2MessageReceipt {
  retryableTicketReceipt: TransactionReceipt
  autoRedeemReceipt?: TransactionReceipt
  l2TxReceipt?: TransactionReceipt
  status: L1ToL2MessageStatus
}

/**
 * Conditional type for Signer or Provider. If T is of type Provider
 * then L1ToL2MessageReaderOrWriter<T> will be of type L1ToL2MessageReader.
 * If T is of type Signer then L1ToL2MessageReaderOrWriter<T> will be of
 * type L1ToL2MessageWriter.
 */
export type L1ToL2MessageReaderOrWriter<T extends SignerOrProvider> =
  T extends Provider ? L1ToL2MessageReader : L1ToL2MessageWriter

export class L1ToL2Message {
  /**
   * When messages are sent from L1 to L2 a retryable ticket is created on L2.
   * An immediate attempt to redeem the ticket will be made (auto-redeem), but if this
   * fails the ticket can be redeemed manually by anyone at a later time.
   */
  public readonly retryableTicketId: string

  /**
   * When a retryable ticket is created a redeem of it will immediately be attempted.
   * The redemption may fail due to not enough gas, or the transaction may revert.
   * The result of this auto-redeem can be accessed by requesting the receipt associated
   * with this autoRedeemId.
   *
   * Note: If no call data is supplied to the retryable ticket an redemption will not be
   * attempted, and instead any value associated with the ticket will be sent directly to the
   * destination. In this case no receipt will be available for the autoRedeemId
   */
  public readonly autoRedeemId: string

  /**
   * If a retryable ticket is successfully redeemed (either via auto-redeem or manually)
   * and l2 transaction will be executed for the payload specified in the retryable ticket.
   * This is the hash of that transaction, and as such no receipt will exist for this hash
   * until the ticket has been successfully redeemed.
   */
  public readonly l2TxHash: string

  private static bitFlip(num: BigNumber): BigNumber {
    return num.or(BigNumber.from(1).shl(255))
  }

  private static calculateL2DerivedHash(
    retryableTicketId: string,
    l2TxnType: L2TxnType
  ): string {
    return keccak256(
      concat([
        zeroPad(retryableTicketId, 32),
        zeroPad(BigNumber.from(l2TxnType).toHexString(), 32),
      ])
    )
  }

  public static calculateRetryableTicketId(
    l2ChainId: BigNumber,
    messageNumber: BigNumber
  ): string {
    return keccak256(
      concat([
        zeroPad(l2ChainId.toHexString(), 32),
        zeroPad(L1ToL2Message.bitFlip(messageNumber).toHexString(), 32),
      ])
    )
  }

  public static calculateAutoRedeemId(retryableTicketId: string): string {
    return this.calculateL2DerivedHash(retryableTicketId, L2TxnType.AUTO_REDEEM)
  }

  public static calculateL2TxHash(retryableTicketId: string): string {
    return this.calculateL2DerivedHash(retryableTicketId, L2TxnType.L2_TX)
  }

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

  public constructor(
    retryableTicketId: string,
    public readonly messageNumber: BigNumber
  ) {
    this.retryableTicketId = retryableTicketId
    this.autoRedeemId = L1ToL2Message.calculateAutoRedeemId(
      this.retryableTicketId
    )
    this.l2TxHash = L1ToL2Message.calculateL2TxHash(this.retryableTicketId)
  }
}

export class L1ToL2MessageReader extends L1ToL2Message {
  public constructor(
    private readonly l2Provider: Provider,
    retryableTicketId: string,
    messageNumber: BigNumber
  ) {
    super(retryableTicketId, messageNumber)
  }

  /**
   * Try to get the receipt for the retryable ticket. See L1ToL2Message.retryableTicketId
   * May throw an error if retryable ticket has yet to be created
   * @returns
   */
  public getRetryableTicketReceipt(): Promise<TransactionReceipt> {
    return this.l2Provider.getTransactionReceipt(this.retryableTicketId)
  }

  /**
   * Receipt for the auto redeem attempt. See L1ToL2Message.autoRedeemId.
   * May throw an error if no auto-redeem attempt was made. This is the case for
   * transactions with no call data
   * @returns
   */
  public getAutoRedeemReceipt(): Promise<TransactionReceipt> {
    return this.l2Provider.getTransactionReceipt(this.autoRedeemId)
  }

  /**
   * Receipt for the l2 transaction created by this message. See L1ToL2Message.l2TxHash
   * May throw an error if the l2 transaction has yet to be executed, which is the case if
   * the retryable ticket has not been created and redeemed.
   * @returns
   */
  public getL2TxReceipt(): Promise<TransactionReceipt> {
    return this.l2Provider.getTransactionReceipt(this.l2TxHash)
  }

  /**
   * Has this message expired. Once expired the retryable ticket can no longer be executed.
   * // CHRIS: check this
   * @returns
   */
  public async isExpired(): Promise<boolean> {
    return (await this.getTimeout()).eq(constants.Zero)
  }

  private async receiptsToStatus(
    retryableTicketReceipt: TransactionReceipt,
    l2TxReceipt: TransactionReceipt
  ): Promise<L1ToL2MessageStatus> {
    if (l2TxReceipt && l2TxReceipt.status === 1) {
      return L1ToL2MessageStatus.REDEEMED
    }

    if (!retryableTicketReceipt) {
      return L1ToL2MessageStatus.NOT_YET_CREATED
    }
    if (retryableTicketReceipt.status === 0) {
      return L1ToL2MessageStatus.CREATION_FAILED
    }
    if (await this.isExpired()) {
      return L1ToL2MessageStatus.EXPIRED
    }
    // we could sanity check that autoredeem failed, but we don't need to
    return L1ToL2MessageStatus.NOT_YET_REDEEMED
  }

  public async status(): Promise<L1ToL2MessageStatus> {
    const l2TxReceipt = await this.getL2TxReceipt()
    const retryableTicketReceipt = await this.getRetryableTicketReceipt()

    return this.receiptsToStatus(l2TxReceipt, retryableTicketReceipt)
  }

  /**
   * Wait for the retryable ticket for the retryable ticket to be created.
   * @param timeout
   * @param confirmations
   * @returns
   */
  public async wait(
    /**
     * Amount of time to wait for the retryable ticket to be created
     */
    timeout = 900000,
    /**
     * Amount of confirmations the retryable ticket and the auto redeem receipt should have
     */
    confirmations?: number
  ): Promise<L1ToL2MessageReceipt> {
    // wait for the retryable ticket - if this doesn't exist then there's no point
    // looking for the other receipts
    const retryableTicketReceipt = await this.l2Provider.waitForTransaction(
      this.retryableTicketId,
      confirmations,
      timeout
    )

    // if a retryable ticket exists then an auto redeem receipt also exists
    // except in the case that the retryable ticket specifies that the l2 transaction contract
    // no call data (just an eth deposit).
    let autoRedeemReceipt: TransactionReceipt | undefined
    try {
      // CHRIS: come back and check this - do we really get an empty result here?
      // CHRIS: check these use cases:
      // 1. normal case
      //    a) wait for the ticket to be created, it is automatically redeemed, and we now have an l2 receipt
      // 2. failed to create
      //    a) wait for creation, and check the status - we'll go no further
      // 3. it got created, but auto redeem failed
      //    c) now we need to manually redeem, or cancel, or do nothing
      autoRedeemReceipt = await this.l2Provider.waitForTransaction(
        this.autoRedeemId,
        confirmations,
        3000 // autoredeem should be available immediately
      )
    } catch (err) {
      // an auto redeem receipt should be available immediately
      // if it's not it could be because there was no call data - like an ETH deposit
    }

    const l2TxReceipt = await this.getL2TxReceipt()

    return {
      retryableTicketReceipt: retryableTicketReceipt,
      autoRedeemReceipt,
      l2TxReceipt: l2TxReceipt,
      status: await this.receiptsToStatus(retryableTicketReceipt, l2TxReceipt),
    }
  }

  /**
   * How long until this message expires
   * @returns
   */
  public getTimeout(): Promise<BigNumber> {
    const arbRetryableTx = ArbRetryableTx__factory.connect(
      ARB_RETRYABLE_TX_ADDRESS,
      this.l2Provider
    )
    return arbRetryableTx.getTimeout(this.l2TxHash)
  }

  /**
   * // CHRIS: what's this all about?
   * @returns
   */
  public getBeneficiary(): Promise<string> {
    const arbRetryableTx = ArbRetryableTx__factory.connect(
      ARB_RETRYABLE_TX_ADDRESS,
      this.l2Provider
    )
    return arbRetryableTx.getBeneficiary(this.l2TxHash)
  }
}

export class L1ToL2MessageWriter extends L1ToL2MessageReader {
  public constructor(
    private readonly l2Signer: Signer,
    retryableTicketId: string,
    messageNumber: BigNumber
  ) {
    super(l2Signer.provider!, retryableTicketId, messageNumber)
    if (!l2Signer.provider) throw new Error('Signer not connected to provider.')
  }

  /**
   * Manually redeem the retryable ticket.
   * Throws if message status is not L1ToL2MessageStatus.NOT_YET_REDEEMED
   */
  public async redeem(): Promise<ContractTransaction> {
    const status = await this.status()
    if (status === L1ToL2MessageStatus.NOT_YET_REDEEMED) {
      const arbRetryableTx = ArbRetryableTx__factory.connect(
        ARB_RETRYABLE_TX_ADDRESS,
        this.l2Signer
      )
      return await arbRetryableTx.redeem(this.l2TxHash)
    } else {
      throw new ArbTsError(
        `Cannot redeem. Message status: ${status} must be: ${L1ToL2MessageStatus.NOT_YET_REDEEMED}.`
      )
    }
  }

  /**
   * Cancel the retryable ticket.
   * Throws if message status is not L1ToL2MessageStatus.NOT_YET_REDEEMED
   */
  public async cancel(): Promise<ContractTransaction> {
    const status = await this.status()
    if (status === L1ToL2MessageStatus.NOT_YET_REDEEMED) {
      const arbRetryableTx = ArbRetryableTx__factory.connect(
        ARB_RETRYABLE_TX_ADDRESS,
        this.l2Signer
      )
      return await arbRetryableTx.cancel(this.l2TxHash)
    } else {
      throw new ArbTsError(
        `Cannot cancel. Message status: ${status} must be: ${L1ToL2MessageStatus.NOT_YET_REDEEMED}.`
      )
    }
  }
}
