import {
  MultiChainConnector,
  SignersAndProviders,
} from '../utils/MultichainConnector'
import { TransactionReceipt } from '@ethersproject/providers'
import { ContractTransaction } from '@ethersproject/contracts'
import { BigNumber } from '@ethersproject/bignumber'
import { ArbRetryableTx__factory } from '../abi/factories/ArbRetryableTx__factory'
import { ARB_RETRYABLE_TX_ADDRESS } from '../precompile_addresses'
import { constants } from 'ethers'
import {
  getMessageNumbers,
  calculateRetryableTicketCreationHash,
  calculateRetryableUserTxnHash,
  calculateRetryableAutoRedeemTxnHash,
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

export class L1ToL2Message extends MultiChainConnector {
  messgeNumber: BigNumber
  arbRetryableActions: ArbRetryableActions
  l1TxnHash: string
  constructor(
    signersAndProviders: SignersAndProviders,
    l1TxnReceipt: TransactionReceipt,
    messageNumberIndex?: number
  ) {
    super()
    this.initSignorsAndProviders(signersAndProviders)

    const l1TxnHash = l1TxnReceipt.transactionHash
    this.l1TxnHash = l1TxnHash

    const messageNumbers = getMessageNumbers(l1TxnReceipt)
    if (messageNumbers === undefined)
      throw new Error(`No l1 to L2 message found for ${l1TxnHash}`)

    if (
      messageNumberIndex !== undefined &&
      messageNumberIndex > messageNumbers.length
    )
      throw new Error(
        `Provided message number out of range for ${l1TxnHash}; index was ${messageNumberIndex}, but only ${messageNumbers.length} messages`
      )
    if (messageNumberIndex === undefined && messageNumbers.length > 1)
      throw new Error(
        `${messageNumbers.length} L2 messages for ${l1TxnHash}; must provide messamessageNumberIndex`
      )

    this.messgeNumber = messageNumbers[messageNumberIndex || 0]
    this.arbRetryableActions = new ArbRetryableActions(signersAndProviders)
  }

  static async initFromTL1xHash(
    signersAndProviders: SignersAndProviders,
    l1TxnHash: string,
    messageNumberIndex?: number
  ): Promise<L1ToL2Message> {
    const l1Provider =
      signersAndProviders.l1Provider || signersAndProviders.l1Signer?.provider
    if (!l1Provider) throw new Error('Must provider an l1 provider')
    const rec = await l1Provider.getTransactionReceipt(l1TxnHash)
    return new L1ToL2Message(signersAndProviders, rec, messageNumberIndex)
  }

  get ticketCreationHash(): string {
    if (!this.l2Network) throw new Error('need l2 signer')
    const l2ChainID = BigNumber.from(this.l2Network.chainID)
    return calculateRetryableTicketCreationHash(this.messgeNumber, l2ChainID)
  }
  get autoRedeemHash(): string {
    if (!this.l2Network) throw new Error('need l2 signer')
    const l2ChainID = BigNumber.from(this.l2Network.chainID)
    return calculateRetryableAutoRedeemTxnHash(this.messgeNumber, l2ChainID)
  }

  get userTxnHash(): string {
    if (!this.l2Network) throw new Error('need l2 provider')
    const l2ChainID = BigNumber.from(this.l2Network.chainID)
    return calculateRetryableUserTxnHash(this.messgeNumber, l2ChainID)
  }

  public getL1TxnReceipt(): Promise<TransactionReceipt> {
    if (!this.l1Provider) throw new Error('need l1 provier')
    return this.l1Provider.getTransactionReceipt(this.l1TxnHash)
  }

  public getTicketCreationReceipt(): Promise<TransactionReceipt> {
    if (!this.l2Provider) throw new Error('need l2 provider')
    return this.l2Provider.getTransactionReceipt(this.ticketCreationHash)
  }
  public getAutoRedeemReceipt(): Promise<TransactionReceipt> {
    if (!this.l2Provider) throw new Error('need l2 provider')
    return this.l2Provider.getTransactionReceipt(this.autoRedeemHash)
  }
  public getUserTxnReceipt(): Promise<TransactionReceipt> {
    if (!this.l2Provider) throw new Error('need l2 provider')
    return this.l2Provider.getTransactionReceipt(this.userTxnHash)
  }

  public async redeem(): Promise<ContractTransaction> {
    if (!this.l2Signer || !this.l2Network) throw new Error('need l2 signer')

    // explicitely check if already redeemed / do a getStatus?
    return this.arbRetryableActions.redeem(this.userTxnHash)
  }

  public async cancel(): Promise<ContractTransaction> {
    if (!this.l2Signer) throw new Error('need l2 provider')
    return this.arbRetryableActions.cancel(this.userTxnHash)
  }

  public async wait(
    timeout = 900000,
    confirmations?: number
  ): Promise<L1ToL2MessageReceipt> {
    if (!this.l2Provider) throw new Error('need l2 provider')
    const ticketCreationReceipt = await this.l2Provider.waitForTransaction(
      this.ticketCreationHash,
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
    if (!this.l2Provider) throw new Error('need l2 provider')

    const userTxnReceipt = await this.l2Provider.getTransactionReceipt(
      this.userTxnHash
    )

    const ticketCreationReceipt = await this.l2Provider.getTransactionReceipt(
      this.ticketCreationHash
    )
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
}

export class ArbRetryableActions extends MultiChainConnector {
  constructor(signersAndProviders: SignersAndProviders) {
    super()
    this.initSignorsAndProviders(signersAndProviders)
  }

  public redeem(userL2TxnHash: string): Promise<ContractTransaction> {
    if (!this.l2Signer) throw new Error('Must have l2Signer')

    const arbRetryableTx = ArbRetryableTx__factory.connect(
      ARB_RETRYABLE_TX_ADDRESS,
      this.l2Signer
    )
    return arbRetryableTx.redeem(userL2TxnHash)
  }
  public cancel(userL2TxnHash: string): Promise<ContractTransaction> {
    if (!this.l2Signer) throw new Error('Must have l2Signer')

    const arbRetryableTx = ArbRetryableTx__factory.connect(
      ARB_RETRYABLE_TX_ADDRESS,
      this.l2Signer
    )
    return arbRetryableTx.cancel(userL2TxnHash)
  }

  public getTimeout(userL2TxnHash: string): Promise<BigNumber> {
    if (!this.l2Provider) throw new Error('Must have l2Provider')

    const arbRetryableTx = ArbRetryableTx__factory.connect(
      ARB_RETRYABLE_TX_ADDRESS,
      this.l2Provider
    )
    return arbRetryableTx.getTimeout(userL2TxnHash)
  }
  // keep alive etc.
}
