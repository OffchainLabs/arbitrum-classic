import {
  MultiChainConnector,
  SignersAndProviders,
} from '../utils/MultiChainConnector'
import { TransactionReceipt } from '@ethersproject/providers'
import { ContractTransaction } from '@ethersproject/contracts'
import { BigNumber } from '@ethersproject/bignumber'
import { RetryableActions } from './RetryableActions'
import { constants } from 'ethers'
import {
  getMessageNumbersFromL1TxnReceipt,
  calculateRetryableTicketCreationHash,
  calculateL2MessageFromTicketTxnHash,
  L2TxnType,
} from './lib'

import { getTxnReceipt } from '../utils/lib'

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

export class L1ToL2Message extends MultiChainConnector {
  arbRetryableActions: RetryableActions
  l1TxnHash?: string

  constructor(
    signersAndProviders: SignersAndProviders,
    public readonly l2TicketCreationTxnHash: string,
    public readonly messageNumber?: BigNumber,
    l1TxnHash?: string
  ) {
    super()
    this.initSignersAndProviders(signersAndProviders)
    this.arbRetryableActions = new RetryableActions(signersAndProviders)
    this.l1TxnHash = l1TxnHash
  }
  static async initFromL1Txn(
    signersAndProviders: SignersAndProviders,
    l1Txn: string | TransactionReceipt,
    messageNumberIndex?: number
  ): Promise<L1ToL2Message> {
    const l1TxnHash = typeof l1Txn === 'string' ? l1Txn : l1Txn.transactionHash
    const allL1ToL2Messages = await L1ToL2Message.initAllFromL1Txn(
      signersAndProviders,
      l1Txn
    )
    const messageCount = allL1ToL2Messages.length
    if (!messageCount)
      throw new Error(`No l1 to L2 message found for ${l1TxnHash}`)

    if (messageNumberIndex !== undefined && messageNumberIndex >= messageCount)
      throw new Error(
        `Provided message number out of range for ${l1TxnHash}; index was ${messageNumberIndex}, but only ${messageCount} messages`
      )
    if (messageNumberIndex === undefined && messageCount > 1)
      throw new Error(
        `${messageCount} L2 messages for ${l1TxnHash}; must provide messageNumberIndex (or use (signersAndProviders, l1Txn))`
      )

    return allL1ToL2Messages[messageNumberIndex || 0]
  }

  static async initAllFromL1Txn(
    signersAndProviders: SignersAndProviders,
    l1Txn: string | TransactionReceipt
  ): Promise<L1ToL2Message[]> {
    const l1TxnReceipt = await getTxnReceipt(
      l1Txn,
      signersAndProviders.l1Provider
    )
    const l1TxnHash = l1TxnReceipt.transactionHash

    if (!signersAndProviders.l2Provider)
      throw new Error('Missing required L2 Provider')
    const chainID = (
      await signersAndProviders.l2Provider.getNetwork()
    ).chainId.toString()

    const messageNumbers = getMessageNumbersFromL1TxnReceipt(l1TxnReceipt)
    if (!messageNumbers.length)
      throw new Error('No l1 to l2 messages found in L1 txn ' + l1TxnHash)

    return messageNumbers.map((messageNumber: BigNumber) => {
      return new L1ToL2Message(
        signersAndProviders,
        calculateRetryableTicketCreationHash({
          messageNumber,
          l2ChainId: BigNumber.from(chainID),
        }),
        messageNumber,
        l1TxnHash
      )
    })
  }

  static initFromL2Txn(
    signersAndProviders: SignersAndProviders,
    l2TicketCreationHash: string
  ): L1ToL2Message {
    return new L1ToL2Message(signersAndProviders, l2TicketCreationHash)
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

  public getL1TxnReceipt(): Promise<TransactionReceipt> {
    if (!this.l1Provider) throw new Error('Missing required L1 Provider')
    if (!this.l1TxnHash) throw new Error('L1 txn hash not available')
    return this.l1Provider.getTransactionReceipt(this.l1TxnHash)
  }

  public getTicketCreationReceipt(): Promise<TransactionReceipt> {
    if (!this.l2Provider) throw new Error('Missing required L2 Provider')
    return this.l2Provider.getTransactionReceipt(this.l2TicketCreationTxnHash)
  }
  public getAutoRedeemReceipt(): Promise<TransactionReceipt> {
    if (!this.l2Provider) throw new Error('Missing required L2 Provider')
    return this.l2Provider.getTransactionReceipt(this.autoRedeemHash)
  }
  public getUserTxnReceipt(): Promise<TransactionReceipt> {
    if (!this.l2Provider) throw new Error('Missing required L2 Provider')
    return this.l2Provider.getTransactionReceipt(this.userTxnHash)
  }

  public async wait(
    timeout = 900000,
    confirmations?: number
  ): Promise<L1ToL2MessageReceipt> {
    if (!this.l2Provider) throw new Error('Missing required L2 Provider')
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

  public async status(): Promise<L1ToL2MessageStatus> {
    if (!this.l2Provider) throw new Error('Missing required L2 Provider')

    const userTxnReceipt = await this.l2Provider.getTransactionReceipt(
      this.userTxnHash
    )

    const ticketCreationReceipt = await this.getTicketCreationReceipt()
    return this.receiptsToStatus(userTxnReceipt, ticketCreationReceipt)
  }

  public async isExpired(): Promise<boolean> {
    return (await this.arbRetryableActions.getTimeout(this.userTxnHash)).eq(
      constants.Zero
    )
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

  public async redeem(): Promise<ContractTransaction> {
    if (!this.l2Signer) throw new Error('Missing required L2 signer')

    return this.arbRetryableActions.redeem(this.userTxnHash)
  }

  public async cancel(): Promise<ContractTransaction> {
    if (!this.l2Signer) throw new Error('Missing required L2 signer')
    return this.arbRetryableActions.cancel(this.userTxnHash)
  }

  public async getTimeout(): Promise<BigNumber> {
    return this.arbRetryableActions.getTimeout(this.userTxnHash)
  }

  public async getBeneficiary(): Promise<string> {
    return this.arbRetryableActions.getBeneficiary(this.userTxnHash)
  }
}
