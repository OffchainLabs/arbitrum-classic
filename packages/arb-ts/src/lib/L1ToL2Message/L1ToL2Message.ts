import { ArbRetryableTx__factory } from '../abi/factories/ArbRetryableTx__factory'
import { ARB_RETRYABLE_TX_ADDRESS } from '../precompile_addresses'
import { TransactionReceipt } from '@ethersproject/providers'
import { Provider } from '@ethersproject/abstract-provider'
import { Signer } from '@ethersproject/abstract-signer'
import { ContractTransaction } from '@ethersproject/contracts'
import { BigNumber } from '@ethersproject/bignumber'
import { constants } from 'ethers'
import {
  getMessageNumbersFromL1TxnReceipt,
  calculateRetryableTicketCreationHash,
  calculateL2MessageFromTicketTxnHash,
  L2TxnType,
} from './lib'

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
type L1ToL2MessageReaderOrWriter<
  T extends Provider | Signer
> = T extends Provider ? L1ToL2MessageReader : L1ToL2MessageWriter

/**
 * Utiliy functions for signer/provider union types
 */
export class SignerOrProvider {
  public static isSigner(
    signerOrProvider: Provider | Signer
  ): signerOrProvider is Signer {
    return (signerOrProvider as Signer).sendTransaction !== undefined
  }

  public static getProvider(signerOrProvider: Provider | Signer) {
    return this.isSigner(signerOrProvider)
      ? signerOrProvider.provider
      : signerOrProvider
  }
}

export class L1ToL2Message {
  protected static getMessageOrThrow(
    l1TxnReceipt: TransactionReceipt,
    messageNumberIndex?: number
  ) {
    const messageNumbers = getMessageNumbersFromL1TxnReceipt(l1TxnReceipt)
    if (messageNumbers === undefined)
      throw new Error(
        `No l1 to L2 message found for ${l1TxnReceipt.transactionHash}`
      )

    if (
      messageNumberIndex !== undefined &&
      messageNumberIndex > messageNumbers.length
    )
      throw new Error(
        `Provided message number out of range for ${l1TxnReceipt.transactionHash}; index was ${messageNumberIndex}, but only ${messageNumbers.length} messages`
      )
    if (messageNumberIndex === undefined && messageNumbers.length > 1)
      throw new Error(
        `${messageNumbers.length} L2 messages for ${l1TxnReceipt.transactionHash}; must provide messamessageNumberIndex (or use initAllFromL1Txn)`
      )

    return messageNumbers[messageNumberIndex || 0]
  }

  public static async fromL1ReceiptAll<T extends Provider | Signer>(
    l2SignerOrProvider: T,
    l1TxnReceipt: TransactionReceipt
  ): Promise<L1ToL2MessageReaderOrWriter<T>[]>
  public static async fromL1ReceiptAll<T extends Provider | Signer>(
    l2SignerOrProvider: T,
    l1TxnReceipt: TransactionReceipt
  ): Promise<L1ToL2MessageReader[] | L1ToL2MessageWriter[]> {
    const provider = SignerOrProvider.getProvider(l2SignerOrProvider)
    if (!provider) throw new Error('Signer not connected to provider.')

    const chainID = (await provider.getNetwork()).chainId.toString()

    const messageNumbers = getMessageNumbersFromL1TxnReceipt(l1TxnReceipt)
    if (!messageNumbers)
      throw new Error(
        'No l1 to l2 messages found in L1 txn ' + l1TxnReceipt.transactionHash
      )

    return messageNumbers.map((mn: BigNumber) => {
      const ticketCreationHash = calculateRetryableTicketCreationHash({
        l2ChainId: BigNumber.from(chainID),
        messageNumber: mn,
      })
      return l2SignerOrProvider instanceof Provider
        ? new L1ToL2MessageReader(l2SignerOrProvider, ticketCreationHash, mn)
        : new L1ToL2MessageWriter(l2SignerOrProvider, ticketCreationHash, mn)
    })
  }

  public static async fromL1Receipt<T extends Provider | Signer>(
    l2SignerOrProvider: T,
    l1TxnReceipt: TransactionReceipt,
    messageNumberIndex?: number
  ): Promise<L1ToL2MessageReaderOrWriter<T>>
  public static async fromL1Receipt<T extends Provider | Signer>(
    l2SignerOrProvider: T,
    l1TxnReceipt: TransactionReceipt,
    messageNumberIndex?: number
  ): Promise<L1ToL2MessageReader | L1ToL2MessageWriter> {
    const allL1ToL2Messages = await L1ToL2Message.fromL1ReceiptAll(
      l2SignerOrProvider,
      l1TxnReceipt
    )
    const messageCount = allL1ToL2Messages.length
    if (!messageCount)
      throw new Error(
        `No l1 to L2 message found for ${l1TxnReceipt.transactionHash}`
      )

    if (messageNumberIndex !== undefined && messageNumberIndex >= messageCount)
      throw new Error(
        `Provided message number out of range for ${l1TxnReceipt.transactionHash}; index was ${messageNumberIndex}, but only ${messageCount} messages`
      )
    if (messageNumberIndex === undefined && messageCount > 1)
      throw new Error(
        `${messageCount} L2 messages for ${l1TxnReceipt.transactionHash}; must provide messageNumberIndex (or use (signersAndProviders, l1Txn))`
      )

    return allL1ToL2Messages[messageNumberIndex || 0]
  }

  public static fromL2Ticket<T extends Provider | Signer>(
    l2SignerOrProvider: T,
    l2TicketCreationHash: string
  ): L1ToL2MessageReaderOrWriter<T>
  public static fromL2Ticket<T extends Provider | Signer>(
    l2SignerOrProvider: T,
    l2TicketCreationHash: string
  ): L1ToL2MessageReader | L1ToL2MessageWriter {
    return l2SignerOrProvider instanceof Provider
      ? new L1ToL2MessageReader(l2SignerOrProvider, l2TicketCreationHash)
      : new L1ToL2MessageWriter(l2SignerOrProvider, l2TicketCreationHash)
  }
}

export class L1ToL2MessageReader {
  constructor(
    private readonly l2Provider: Provider,
    public readonly l2TicketCreationTxnHash: string,
    public readonly messageNumber?: BigNumber
  ) {}
  // CHRIS: remember initSignersAndProviders was called here did something

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

  public async status(): Promise<L1ToL2MessageStatus> {
    const userTxnReceipt = await this.l2Provider.getTransactionReceipt(
      this.userTxnHash
    )

    const ticketCreationReceipt = await this.getTicketCreationReceipt()
    return this.receiptsToStatus(userTxnReceipt, ticketCreationReceipt)
  }

  public async isExpired(): Promise<boolean> {
    return (await this.getTimeout(this.userTxnHash)).eq(constants.Zero)
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

  public getTimeout(userL2TxnHash: string): Promise<BigNumber> {
    const arbRetryableTx = ArbRetryableTx__factory.connect(
      ARB_RETRYABLE_TX_ADDRESS,
      this.l2Provider
    )
    return arbRetryableTx.getTimeout(userL2TxnHash)
  }
  public getBeneficiary(userL2TxnHash: string): Promise<string> {
    const arbRetryableTx = ArbRetryableTx__factory.connect(
      ARB_RETRYABLE_TX_ADDRESS,
      this.l2Provider
    )
    return arbRetryableTx.getBeneficiary(userL2TxnHash)
  }
}

export class L1ToL2MessageWriter extends L1ToL2MessageReader {
  constructor(
    private readonly l2Signer: Signer,
    l2TicketCreationTxnHash: string,
    messageNumber?: BigNumber
  ) {
    super(l2Signer.provider!, l2TicketCreationTxnHash, messageNumber)
    if (!l2Signer.provider) throw new Error('Signer not connected to provider.')
  }

  public redeem(userL2TxnHash: string): Promise<ContractTransaction> {
    const arbRetryableTx = ArbRetryableTx__factory.connect(
      ARB_RETRYABLE_TX_ADDRESS,
      this.l2Signer
    )
    return arbRetryableTx.redeem(userL2TxnHash)
  }
  public cancel(userL2TxnHash: string): Promise<ContractTransaction> {
    const arbRetryableTx = ArbRetryableTx__factory.connect(
      ARB_RETRYABLE_TX_ADDRESS,
      this.l2Signer
    )
    return arbRetryableTx.cancel(userL2TxnHash)
  }
}
