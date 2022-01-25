import {
  TransactionReceipt,
  TransactionResponse,
} from '@ethersproject/providers'
import { BigNumber } from 'ethers'

/**
 * Eth transaction response with additional arbitrum specific fields
 */
export interface ArbTransactionResponse extends TransactionResponse {
  /**
   * The sequenced number for this transaction
   */
  l1SequenceNumber: BigNumber
  parentRequestId: string
  indexInParent: number
  arbType: number
  arbSubType: number
  /**
   * The l1 block number that would be used for block.number calls
   * that occur within this transaction.
   * See https://developer.offchainlabs.com/docs/time_in_arbitrum
   */
  l1BlockNumber: number
}

export interface FeeStat {
  /**
   * A fixed cost of including your transaction into the Layer 1 inbox contract.
   * This is amortised by batching transactions together
   */
  l1Transaction: BigNumber

  /**
   * The calldata included in each Layer 1 transaction has a cost associated to it.
   * Aggregators are reimbursed for their costs.
   */
  l1Calldata: BigNumber

  /**
   * A charge occurs whenever a storage slot is written to with a non-zero value.
   */
  l2Storage: BigNumber
  /**
   * A charged per unit of computation used (measured in arbgas).
   */
  l2Computation: BigNumber
}

export interface FeeStats {
  /**
   * The prices per unit for this tx
   */
  prices: FeeStat
  /**
   * The number of units used for this tx
   */
  unitsUsed: FeeStat
  /**
   * Total amounts paid: units * price
   */
  paid: FeeStat
}

export interface BatchInfo {
  /**
   * The number of confirmations the batch has on L1
   */
  confirmations: number
  /**
   * Block number the batch appeared on L1
   */
  blockNumber: number
  /**
   * Address that emitted the log containing the batch
   */
  logAddress: string
  /**
   * Topics of the log containing the batch
   */
  logTopics: string[]
  /**
   * Data of the log containing the batch
   */
  logData: string
}

export enum ReturnCode {
  TransactionSuccess = 0,
  EVMRevert = 1,
  /**
   *Arbitrum is too congested to process the transaction
   */
  TooCongested = 2,
  /**
   * Not enough balance to pay for maxGas at gasPrice (for retryables: not enough to cover base submission cost)
   */
  InsufficientBalanceForMaxGas = 3,

  /**
   * Not enough balance for execution (for retryables: not enough to cover callvalue + base submission cost)
   */
  InsufficientBalanceForExecution = 4,
  /**
   * Wrong nonce used in transaction
   */
  WrongNonce = 5,
  /**
   * Transaction was not formatted correctly
   */
  InvalidFormat = 6,
  /**
   * Cannot deploy to specified address ( ** defensive code that should never be triggered ** )
   */
  CannotDeploy = 7,
  /**
   * Exceeded transaction gas limit
   */
  GasLimitExceeded = 8,
  /**
   * Amount of ArbGas provided for the tx is less than the amount required to cover L1 costs (the base tx charge plus L1 calldata charge)
   */
  InsufficientArbGasForL1Costs = 9,
  /**
   * Transaction is below the minimum required arbgas
   */
  InsufficientArbGas = 10,
  /**
   * Transaction set an arbgas price that was too low
   */
  InsufficientArbGasPrice = 11,
  /**
   * Insufficient gas for retryable auto-redeem
   */
  InsufficientBalanceForAutoRedeem = 12,
  /**
   * Sender not permitted
   */
  SenderNotPermitted = 13,
  Unknown = 255,
}

/**
 * Eth transaction receipt with additional arbitrum specific fields
 */
export interface ArbTransactionReceipt extends TransactionReceipt {
  /**
   * Data from a smart contract return or the revert reason if an EVM revert statement was hit
   */
  returnData: string
  /**
   * Arbitrum status code
   */
  returnCode: ReturnCode
  /**
   * Arbitrum fee breakdown
   */
  feeStats: FeeStats
  /**
   * Batch info, populated if an l1 provider is present when fetching the receipt
   */
  l1InboxBatchInfo: BatchInfo | null
  /**
   * The l1 block number that would be used for block.number calls
   * that occur within this transaction.
   * See https://developer.offchainlabs.com/docs/time_in_arbitrum
   */
  l1BlockNumber: number
}
