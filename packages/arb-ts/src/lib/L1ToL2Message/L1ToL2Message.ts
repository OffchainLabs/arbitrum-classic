import {
  MultiChainConnector,
  SignersAndProviders,
} from '../utils/MultichainConnector'
import { TransactionReceipt } from '@ethersproject/providers'
import { BigNumber } from '@ethersproject/bignumber'
import { ArbRetryableTx__factory } from '../abi/factories/ArbRetryableTx__factory'
import { ARB_RETRYABLE_TX_ADDRESS } from '../precompile_addresses'
import {
  getMessageNumbers,
  calculateRetryableTicketCreationHash,
  calculateRetryableUserTxnHash,
} from './lib'

export enum L1ToL2MessageStatus {
  NOT_YET_CREATED,
  CREATION_FAILED,
  NOT_YET_REDEEMED, // i.e., autoredeem failed
  REDEEMED,
  CANCELLED,
}

export class L1ToL2MessageManager extends MultiChainConnector {
  constructor(signersAndProviders: SignersAndProviders) {
    super(signersAndProviders)
  }

  public redeemL1toL2Message(userL2TxnHash: string) {
    if (!this.l2Signer) throw new Error('Must have l2Signer')

    const arbRetryableTx = ArbRetryableTx__factory.connect(
      ARB_RETRYABLE_TX_ADDRESS,
      this.l2Signer
    )
    return arbRetryableTx.redeem(userL2TxnHash)
  }
  public cancelL1toL2Message(userL2TxnHash: string) {
    if (!this.l2Signer) throw new Error('Must have l2Signer')

    const arbRetryableTx = ArbRetryableTx__factory.connect(
      ARB_RETRYABLE_TX_ADDRESS,
      this.l2Signer
    )
    return arbRetryableTx.cancel(userL2TxnHash)
  }
}

export class L1ToL2Message extends MultiChainConnector {
  messgeNumber: BigNumber
  l1ToL2MessageManager: L1ToL2MessageManager
  constructor(
    signersAndProviders: SignersAndProviders,
    l1TxnReceipt: TransactionReceipt,
    messageNumberIndex?: number
  ) {
    super(signersAndProviders)
    const l1TxnHash = l1TxnReceipt.transactionHash

    if (!l1TxnReceipt) throw new Error(`Txn rec not found for ${l1TxnHash}`)

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
    this.l1ToL2MessageManager = new L1ToL2MessageManager(signersAndProviders)
  }

  public async redeem() {
    if (!this.l2Signer || !this.l2Provider) throw new Error('need l2 provider')

    const l2ChainID = BigNumber.from(
      (await this.l2Provider.getNetwork()).chainId
    )
    const userTxnHash = calculateRetryableUserTxnHash(
      this.messgeNumber,
      l2ChainID
    )

    // explicitely check if already redeemed / do a getStatus?

    return this.l1ToL2MessageManager.redeemL1toL2Message(userTxnHash)
  }

  public async cancel() {
    if (!this.l2Signer || !this.l2Provider) throw new Error('need l2 provider')

    const l2ChainID = BigNumber.from(
      (await this.l2Provider.getNetwork()).chainId
    )
    const userTxnHash = calculateRetryableUserTxnHash(
      this.messgeNumber,
      l2ChainID
    )

    // explicitely check if already redeemed / do a getStatus?

    return this.l1ToL2MessageManager.cancelL1toL2Message(userTxnHash)
  }

  public async getStatus() {
    if (!this.l2Provider) throw new Error('need l2 provider')
    // TODO: handle networks / chain ids
    const l2ChainID = BigNumber.from(
      (await this.l2Provider.getNetwork()).chainId
    )

    const ticketCreationHash = calculateRetryableTicketCreationHash(
      this.messgeNumber,
      l2ChainID
    )

    const userTxnHash = calculateRetryableUserTxnHash(
      this.messgeNumber,
      l2ChainID
    )

    const userTxnReceipt = await this.l2Provider.getTransactionReceipt(
      userTxnHash
    )

    if (userTxnReceipt && userTxnReceipt.status === 1) {
      return L1ToL2MessageStatus.REDEEMED
    }

    const ticketCreationReceipt = await this.l2Provider.getTransactionReceipt(
      ticketCreationHash
    )
    if (!ticketCreationReceipt) {
      return L1ToL2MessageStatus.NOT_YET_CREATED
    }
    if (ticketCreationReceipt.status === 0) {
      return L1ToL2MessageStatus.CREATION_FAILED
    }
    // we could sanity check that autoredeem failed, but we don't need to
    return L1ToL2MessageStatus.NOT_YET_REDEEMED
  }
}
