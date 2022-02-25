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
import { Provider } from '@ethersproject/abstract-provider'
import { Signer } from '@ethersproject/abstract-signer'
import { ContractTransaction } from '@ethersproject/contracts'
import { BigNumber } from '@ethersproject/bignumber'
import { keccak256 } from '@ethersproject/keccak256'
import { concat, zeroPad } from '@ethersproject/bytes'

import { ArbRetryableTx__factory } from '../abi'
import { ARB_RETRYABLE_TX_ADDRESS } from '../dataEntities/constants'
import {
  SignerProviderUtils,
  SignerOrProvider,
} from '../dataEntities/signerOrProvider'
import { ArbTsError } from '../dataEntities/errors'

export enum L2TxnType {
  L2_TX = 0,
  AUTO_REDEEM = 1,
}

export enum L1ToL2MessageStatus {
  /**
   * The retryable ticket has yet to be created
   */
  NOT_YET_CREATED = 1,
  /**
   * An attempt was made to create the retryable ticket, but it failed.
   * This could be due to not enough submission cost being paid by the L1 transaction
   */
  CREATION_FAILED = 2,
  /**
   * The retryable ticket has been created but has not been redeemed. This could be due to the
   * auto redeem failing, or if the params (max l2 gas price) * (max l2 gas) = 0 then no auto
   * redeem tx is ever issued. An auto redeem is also never issued for ETH deposits.
   * A manual redeem is now required.
   */
  FUNDS_DEPOSITED_ON_L2 = 3,
  /**
   * The retryable ticket has been redeemed (either by auto, or manually) and the
   * l2 transaction has been executed
   */
  REDEEMED = 4,
  /**
   * The message has either expired or has been canceled. It can no longer be redeemed.
   */
  EXPIRED = 5,
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
   * The retryableCreationId can be used to retrieve information about the success or failure of the
   * creation of the retryable ticket.
   */
  public readonly retryableCreationId: string

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
    retryableCreationId: string,
    l2TxnType: L2TxnType
  ): string {
    return keccak256(
      concat([
        zeroPad(retryableCreationId, 32),
        zeroPad(BigNumber.from(l2TxnType).toHexString(), 32),
      ])
    )
  }

  public static calculateRetryableCreationId(
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

  public static calculateAutoRedeemId(retryableCreationId: string): string {
    return this.calculateL2DerivedHash(
      retryableCreationId,
      L2TxnType.AUTO_REDEEM
    )
  }

  public static calculateL2TxHash(retryableCreationId: string): string {
    return this.calculateL2DerivedHash(retryableCreationId, L2TxnType.L2_TX)
  }

  public static fromRetryableCreationId<T extends SignerOrProvider>(
    l2SignerOrProvider: T,
    retryableCreationId: string,
    messageNumber: BigNumber
  ): L1ToL2MessageReaderOrWriter<T>
  public static fromRetryableCreationId<T extends SignerOrProvider>(
    l2SignerOrProvider: T,
    retryableCreationId: string,
    messageNumber: BigNumber
  ): L1ToL2MessageReader | L1ToL2MessageWriter {
    return SignerProviderUtils.isSigner(l2SignerOrProvider)
      ? new L1ToL2MessageWriter(
          l2SignerOrProvider,
          retryableCreationId,
          messageNumber
        )
      : new L1ToL2MessageReader(
          l2SignerOrProvider,
          retryableCreationId,
          messageNumber
        )
  }

  public constructor(
    retryableCreationId: string,
    public readonly messageNumber: BigNumber
  ) {
    this.retryableCreationId = retryableCreationId
    this.autoRedeemId = L1ToL2Message.calculateAutoRedeemId(
      this.retryableCreationId
    )
    this.l2TxHash = L1ToL2Message.calculateL2TxHash(this.retryableCreationId)
  }
}

/**
 * If the status is redeemed an l2TxReceipt is populated.
 * For all other statuses l2TxReceipt is not populated
 */
export type L1ToL2MessageWaitResult =
  | { status: L1ToL2MessageStatus.REDEEMED; l2TxReceipt: TransactionReceipt }
  | { status: Exclude<L1ToL2MessageStatus, L1ToL2MessageStatus.REDEEMED> }

export class L1ToL2MessageReader extends L1ToL2Message {
  public constructor(
    public readonly l2Provider: Provider,
    retryableCreationId: string,
    messageNumber: BigNumber
  ) {
    super(retryableCreationId, messageNumber)
  }

  /**
   * Try to get the receipt for the retryable ticket. See L1ToL2Message.retryableCreationId
   * May throw an error if retryable ticket has yet to be created
   * @returns
   */
  public getRetryableCreationReceipt(): Promise<TransactionReceipt> {
    return this.l2Provider.getTransactionReceipt(this.retryableCreationId)
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
   * Has this message expired. Once expired the retryable ticket can no longer be redeemed.
   * @returns
   */
  public async isExpired(): Promise<boolean> {
    const currentTimestamp = BigNumber.from(
      (await this.l2Provider.getBlock('latest')).timestamp
    )
    const timeoutTimestamp = await this.getTimeout()

    // timeoutTimestamp returns the timestamp at which the retryable ticket expires
    // it can also return 0 if the ticket l2Tx does not exist
    return currentTimestamp.gte(timeoutTimestamp)
  }

  protected async receiptsToStatus(
    retryableCreationReceipt: TransactionReceipt | null | undefined,
    l2TxReceipt: TransactionReceipt | null | undefined
  ): Promise<L1ToL2MessageStatus> {
    // happy path for non auto redeemable messages
    // NOT_YET_CREATED -> FUNDS_DEPOSITED
    // these will later either transition to EXPIRED after the timeout
    // (this is what happens to eth deposits since they don't need to be
    // redeemed) or to REDEEMED if the retryable is manually redeemed

    // happy path for auto redeemable messages
    // NOT_YET_CREATED -> FUNDS_DEPOSITED -> REDEEMED
    // an attempt to auto redeem executable messages is made immediately
    // after the retryable is created - which if successful will transition
    // the status to REDEEMED. If the auto redeem fails then the ticket
    // will transition to REDEEMED if manually redeemed, or EXPIRE
    // after the timeout is reached and the ticket is not redeemed

    // we test the retryable receipt first as if this doesnt exist there's
    // no point looking to see if expired
    if (!retryableCreationReceipt) {
      return L1ToL2MessageStatus.NOT_YET_CREATED
    }
    if (retryableCreationReceipt.status === 0) {
      return L1ToL2MessageStatus.CREATION_FAILED
    }

    // ticket created, has it been redeemed?
    if (l2TxReceipt && l2TxReceipt.status === 1) {
      return L1ToL2MessageStatus.REDEEMED
    }

    // not redeemed, has it now expired
    if (await this.isExpired()) {
      return L1ToL2MessageStatus.EXPIRED
    }

    // ticket was created but not redeemed
    // this could be because
    // a) the ticket is non auto redeemable (l2GasPrice == 0 || l2GasLimit == 0) -
    //    this is usually an eth deposit. But in some rare case the
    //    user may still want to manually redeem it
    // b) the ticket is auto redeemable, but the auto redeem failed

    // the fact that the auto redeem failed isn't usually useful to the user
    // if they're doing an eth deposit they dont care about redemption
    // and if they do want execution to occur they will know that they're
    // here because the auto redeem failed. If they really want to check
    // they can fetch the auto redeem receipt and check the status on it
    return L1ToL2MessageStatus.FUNDS_DEPOSITED_ON_L2
  }

  protected async status(): Promise<L1ToL2MessageStatus> {
    return this.receiptsToStatus(
      await this.getRetryableCreationReceipt(),
      await this.getL2TxReceipt()
    )
  }

  /**
   * Wait for the retryable ticket to be created, for it to be redeemed, and for the l2Tx to be executed.
   * Note: The terminal status of a transaction that only does an eth deposit is FUNDS_DEPOSITED_ON_L2 as
   * no L2 transaction needs to be executed, however the terminal state of any other transaction is REDEEMED
   * which represents that the retryable ticket has been redeemed and the L2 tx has been executed.
   * @param timeout Amount of time to wait for the retryable ticket to be created
   * @param confirmations Amount of confirmations the retryable ticket and the auto redeem receipt should have
   * @returns The wait result contains a status, and optionally the l2TxReceipt.
   * If the status is "REDEEMED" then a l2TxReceipt is also available on the result.
   * If the status has any other value then l2TxReceipt is not populated.
   */
  public async waitForStatus(
    confirmations?: number,
    timeout = 900000
  ): Promise<L1ToL2MessageWaitResult> {
    // try to wait for the retryable ticket to be created
    let retryableCreationReceipt: TransactionReceipt | undefined
    try {
      retryableCreationReceipt = await this.l2Provider.waitForTransaction(
        this.retryableCreationId,
        confirmations,
        timeout
      )
    } catch (err) {
      if ((err as Error).message.includes('timeout exceeded')) {
        // do nothing - this is dependent on the timeout passed in
      } else throw err
    }

    // get the l2TxReceipt, dont bother trying if we couldn't get the retryableCreationReceipt
    const l2TxReceipt = retryableCreationReceipt
      ? await this.getL2TxReceipt()
      : undefined

    const status = await this.receiptsToStatus(
      retryableCreationReceipt,
      l2TxReceipt
    )

    if (status === L1ToL2MessageStatus.REDEEMED) {
      return {
        // if the status is redeemed we know the l2TxReceipt must exist
        l2TxReceipt: l2TxReceipt!,
        status,
      }
    } else {
      return {
        status,
      }
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
   * Address to which CallValue will be credited to on L2 if the retryable ticket times out or is cancelled.
   * The Beneficiary is also the address with the right to cancel a Retryable Ticket (if the ticket hasn’t been redeemed yet).
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
    public readonly l2Signer: Signer,
    retryableCreationId: string,
    messageNumber: BigNumber
  ) {
    super(l2Signer.provider!, retryableCreationId, messageNumber)
    if (!l2Signer.provider) throw new Error('Signer not connected to provider.')
  }

  /**
   * Manually redeem the retryable ticket.
   * Throws if message status is not L1ToL2MessageStatus.NOT_YET_REDEEMED
   */
  public async redeem(): Promise<ContractTransaction> {
    const status = await this.status()
    if (status === L1ToL2MessageStatus.FUNDS_DEPOSITED_ON_L2) {
      const arbRetryableTx = ArbRetryableTx__factory.connect(
        ARB_RETRYABLE_TX_ADDRESS,
        this.l2Signer
      )
      return await arbRetryableTx.redeem(this.l2TxHash)
    } else {
      throw new ArbTsError(
        `Cannot redeem. Message status: ${status} must be: ${L1ToL2MessageStatus.FUNDS_DEPOSITED_ON_L2}.`
      )
    }
  }

  /**
   * Cancel the retryable ticket.
   * Throws if message status is not L1ToL2MessageStatus.NOT_YET_REDEEMED
   */
  public async cancel(): Promise<ContractTransaction> {
    const status = await this.status()
    if (status === L1ToL2MessageStatus.FUNDS_DEPOSITED_ON_L2) {
      const arbRetryableTx = ArbRetryableTx__factory.connect(
        ARB_RETRYABLE_TX_ADDRESS,
        this.l2Signer
      )
      return await arbRetryableTx.cancel(this.l2TxHash)
    } else {
      throw new ArbTsError(
        `Cannot cancel. Message status: ${status} must be: ${L1ToL2MessageStatus.FUNDS_DEPOSITED_ON_L2}.`
      )
    }
  }
}
