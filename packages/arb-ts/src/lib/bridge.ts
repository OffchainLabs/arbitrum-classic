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

import {
  Filter,
  Provider,
  TransactionReceipt,
} from '@ethersproject/abstract-provider'
import { Signer } from '@ethersproject/abstract-signer'
import { BigNumber } from '@ethersproject/bignumber'
import {
  ContractReceipt,
  ContractTransaction,
  PayableOverrides,
} from '@ethersproject/contracts'
import { Logger } from '@ethersproject/logger'
import { Zero } from '@ethersproject/constants'
import { parseEther } from '@ethersproject/units'

import { NodeInterface__factory } from './abi/factories/NodeInterface__factory'
import { L1ERC20Gateway__factory } from './abi/factories/L1ERC20Gateway__factory'
import { L1WethGateway__factory } from './abi/factories/L1WethGateway__factory'
import { Inbox__factory } from './abi/factories/Inbox__factory'
import { Bridge__factory } from './abi/factories/Bridge__factory'
import { OldOutbox__factory } from './abi/factories/OldOutbox__factory'
import { ERC20__factory } from './abi/factories/ERC20__factory'
import { L1ERC20Gateway } from './abi/L1ERC20Gateway'
import { L1GatewayRouter } from './abi/L1GatewayRouter'
import { ERC20 } from './abi/ERC20'

import { Await, DepositParams, L1Bridge, L1TokenData } from './l1Bridge'
import { L2Bridge, L2TokenData } from './l2Bridge'
import {
  BridgeHelper,
  BuddyDeployEventResult,
  DepositInitiated,
  GatewaySet,
  L2ToL1EventResult,
  MessageBatchProofInfo,
  OutgoingMessageState,
  WithdrawalInitiated,
} from './bridge_helpers'
import { NODE_INTERFACE_ADDRESS } from './precompile_addresses'
import networks, { Network } from './networks'
import { hexDataLength } from '@ethersproject/bytes'

interface RetryableGasArgs {
  maxSubmissionPrice?: BigNumber
  maxGas?: BigNumber
  gasPriceBid?: BigNumber
  maxSubmissionPricePercentIncrease?: BigNumber
  maxGasPercentIncrease?: BigNumber
}

interface InitOptions {
  customNetwork?: {
    l1Network: Network
    l2Network: Network
  }
}

interface DepositInputParams {
  erc20L1Address: string
  amount: BigNumber
  retryableGasArgs?: RetryableGasArgs
  destinationAddress?: string
}

const isDepositInputParams = (obj: any): obj is DepositInputParams =>
  !obj['l1CallValue']

function isError(error: Error): error is NodeJS.ErrnoException {
  return error instanceof Error
}

const DEFAULT_SUBMISSION_PERCENT_INCREASE = BigNumber.from(400)
const DEFAULT_MAX_GAS_PERCENT_INCREASE = BigNumber.from(50)
const MIN_CUSTOM_DEPOSIT_MAXGAS = BigNumber.from(275000)

interface RetryableParamsOptions {
  maxSubmissionFeePercentIncrease?: BigNumber
  maxGasPercentIncrease?: BigNumber
  maxGasPricePercentIncrease?: BigNumber
  includeL2Callvalue?: boolean
}

/**
 * Main class for accessing token bridge methods; inherits methods from {@link L1Bridge} and {@link L2Bridge}
 */
export class Bridge {
  l1Bridge: L1Bridge
  l2Bridge: L2Bridge
  isCustomNetwork: boolean

  private constructor(
    l1BridgeObj: L1Bridge,
    l2BridgeObj: L2Bridge,
    isCustomNetwork = false
  ) {
    this.l1Bridge = l1BridgeObj
    this.l2Bridge = l2BridgeObj
    this.isCustomNetwork = isCustomNetwork
  }

  static async init(
    ethSigner: Signer,
    arbSigner: Signer,
    { customNetwork }: InitOptions = {}
  ): Promise<Bridge> {
    if (!ethSigner.provider || !arbSigner.provider) {
      throw new Error('Signer needs a provider')
    }

    const [l1ChainId, l2ChainId] = await Promise.all([
      ethSigner.getChainId(),
      arbSigner.getChainId(),
    ])
    const isCustomNetwork = customNetwork !== undefined

    const l1Network = isCustomNetwork
      ? customNetwork.l1Network
      : networks[l1ChainId]
    const l2Network = isCustomNetwork
      ? customNetwork.l2Network
      : networks[l2ChainId]

    if (l1Network && l2Network) {
      if (l1Network.partnerChainID !== l2Network.chainID)
        throw new Error('L1 and L2 networks are not connected')
      if (l1Network.isArbitrum)
        throw new Error('Connected to an Arbitrum networks as the L1...')
      if (!l2Network.isArbitrum)
        throw new Error('Connected to an L1 network as the L2...')
    } else {
      throw new Error('Current network configuration not supported.')
    }

    if (isCustomNetwork) {
      // check routers are deployed when using a custom network configuration
      const [l1RouterCode, l2RouterCode] = await Promise.all([
        ethSigner.provider.getCode(l1Network.tokenBridge.l1GatewayRouter),
        arbSigner.provider.getCode(l2Network.tokenBridge.l2GatewayRouter),
      ])

      if (l1RouterCode === '0x') {
        throw new Error(
          `No code deployed to ${l1Network.tokenBridge.l1GatewayRouter} in the L1`
        )
      }

      if (l2RouterCode === '0x') {
        throw new Error(
          `No code deployed to ${l2Network.tokenBridge.l2GatewayRouter} in the L2`
        )
      }
    }

    const l1BridgeObj = new L1Bridge(l1Network, ethSigner)
    const l2BridgeObj = new L2Bridge(l2Network, arbSigner)

    return new Bridge(l1BridgeObj, l2BridgeObj, isCustomNetwork)
  }

  public async setSigner(newEthSigner: Signer, newArbSigner: Signer) {
    await this.l1Bridge.setSigner(newEthSigner)
    await this.l2Bridge.setSigner(newArbSigner)
  }

  get l1GatewayRouter(): L1GatewayRouter {
    return this.l1Bridge.l1GatewayRouter
  }

  defaultL1Gateway(): Promise<L1ERC20Gateway> {
    return this.l1Bridge.getDefaultL1Gateway()
  }
  get l1Signer(): Signer {
    return this.l1Bridge.l1Signer
  }
  get l1Provider(): Provider {
    return this.l1Bridge.l1Provider
  }
  get l2Provider(): Provider {
    return this.l2Bridge.l2Provider
  }
  get l2Signer(): Signer {
    return this.l2Bridge.l2Signer
  }

  /**
   * Set allowance for L1 router contract
   */
  public async approveToken(
    erc20L1Address: string,
    amount?: BigNumber,
    overrides: PayableOverrides = {}
  ): Promise<ContractTransaction> {
    return this.l1Bridge.approveToken(erc20L1Address, amount, overrides)
  }

  /**
   * Deposit ether from L1 to L2.
   */
  public async depositETH(
    value: BigNumber,
    _maxSubmissionPricePercentIncrease?: BigNumber,
    overrides: PayableOverrides = {}
  ): Promise<ContractTransaction> {
    const maxSubmissionPricePercentIncrease =
      _maxSubmissionPricePercentIncrease || DEFAULT_SUBMISSION_PERCENT_INCREASE

    const maxSubmissionPrice = BridgeHelper.percentIncrease(
      (await this.l2Bridge.getTxnSubmissionPrice(0))[0],
      maxSubmissionPricePercentIncrease
    )

    return this.l1Bridge.depositETH(value, maxSubmissionPrice, overrides)
  }

  private async looksLikeWethGateway(potentialWethGatewayAddress: string) {
    try {
      const potentialWethGateway = L1WethGateway__factory.connect(
        potentialWethGatewayAddress,
        this.l1Provider
      )
      await potentialWethGateway.l1Weth()
      return true
    } catch (err) {
      if (
        err instanceof Error &&
        isError(err) &&
        err.code === Logger.errors.CALL_EXCEPTION
      ) {
        return false
      } else {
        throw err
      }
    }
  }

  public async getRetryableTxnParams(
    callDataHex: string,
    sender: string,
    destinationAddress: string,
    _l2CallValue?: BigNumber,
    options: RetryableParamsOptions = {}
  ) {
    const maxGasPriceIncrease =
      options.maxGasPricePercentIncrease || BigNumber.from(0)
    const maxGasIncrease = options.maxGasPercentIncrease || BigNumber.from(0)
    const maxSubmissionFeeIncrease =
      options.maxSubmissionFeePercentIncrease ||
      DEFAULT_SUBMISSION_PERCENT_INCREASE
    const l2CallValue = _l2CallValue || BigNumber.from(0)

    const includeL2Callvalue =
      typeof options.includeL2Callvalue === 'boolean'
        ? options.includeL2Callvalue
        : true

    const gasPriceBid = BridgeHelper.percentIncrease(
      await this.l2Provider.getGasPrice(),
      maxGasPriceIncrease
    )

    const submissionPrice = (
      await this.l2Bridge.getTxnSubmissionPrice(hexDataLength(callDataHex))
    )[0]
    const submissionPriceBid = BridgeHelper.percentIncrease(
      submissionPrice,
      maxSubmissionFeeIncrease
    )

    const nodeInterface = NodeInterface__factory.connect(
      NODE_INTERFACE_ADDRESS,
      this.l2Provider
    )

    const maxGas = (
      await nodeInterface.estimateRetryableTicket(
        sender,
        parseEther('1').add(
          l2CallValue
        ) /** we add a 1 ether "deposit" buffer to pay for execution in the gas estimation  */,
        destinationAddress,
        l2CallValue,
        submissionPriceBid,
        sender,
        sender,
        0,
        gasPriceBid,
        callDataHex
      )
    )[0]

    const maxGasBid = BridgeHelper.percentIncrease(maxGas, maxGasIncrease)

    let totalDepositValue = submissionPriceBid.add(gasPriceBid.mul(maxGas))
    if (includeL2Callvalue) {
      totalDepositValue = totalDepositValue.add(l2CallValue)
    }
    return {
      gasPriceBid,
      submissionPriceBid,
      maxGasBid,
      totalDepositValue,
    }
  }

  public async getDepositTxParams(
    {
      erc20L1Address,
      amount,
      retryableGasArgs = {},
      destinationAddress,
    }: DepositInputParams,
    overrides: PayableOverrides = {}
  ): Promise<DepositParams> {
    const {
      l1WethGateway: l1WethGatewayAddress,
      l1CustomGateway: l1CustomGatewayAddress,
    } = this.l1Bridge.network.tokenBridge

    // 1. Get gas price
    const gasPriceBid =
      retryableGasArgs.gasPriceBid || (await this.l2Provider.getGasPrice())

    const l1GatewayAddress = await this.l1Bridge.getGatewayAddress(
      erc20L1Address
    )

    // 2. Get submission price (this depends on size of calldata used in deposit)
    const l1Gateway = L1ERC20Gateway__factory.connect(
      l1GatewayAddress,
      this.l1Provider
    )
    const sender = await this.l1Bridge.getWalletAddress()
    const to = destinationAddress ? destinationAddress : sender
    const depositCalldata = await l1Gateway.getOutboundCalldata(
      erc20L1Address,
      sender,
      to,
      amount,
      '0x'
    )

    const maxSubmissionPricePercentIncrease =
      retryableGasArgs.maxSubmissionPricePercentIncrease ||
      DEFAULT_SUBMISSION_PERCENT_INCREASE

    const maxSubmissionPrice = BridgeHelper.percentIncrease(
      (
        await this.l2Bridge.getTxnSubmissionPrice(depositCalldata.length - 2)
      )[0],
      maxSubmissionPricePercentIncrease
    )

    // 3. Estimate gas
    const nodeInterface = NodeInterface__factory.connect(
      NODE_INTERFACE_ADDRESS,
      this.l2Provider
    )
    const l2Dest = await l1Gateway.counterpartGateway()

    /** The WETH gateway is the only deposit that requires callvalue in the L2 user-tx (i.e., the recently un-wrapped ETH)
     * Here we check if this is a WETH deposit, and include the callvalue for the gas estimate query if so
     */
    const estimateGasCallValue = await (async () => {
      if (this.isCustomNetwork) {
        // For custom network, we do an ad-hoc check to see if it's a WETH gateway
        if (await this.looksLikeWethGateway(l1GatewayAddress)) {
          return amount
        }
        // ...otherwise we directly check it against the config file
      } else if (l1WethGatewayAddress === l1GatewayAddress) {
        return amount
      }

      return Zero
    })()

    let maxGas =
      retryableGasArgs.maxGas ||
      BridgeHelper.percentIncrease(
        (
          await nodeInterface.estimateRetryableTicket(
            l1GatewayAddress,
            parseEther('0.05').add(
              estimateGasCallValue
            ) /** we add a 0.05 "deposit" buffer to pay for execution in the gas estimation  */,
            l2Dest,
            estimateGasCallValue,
            maxSubmissionPrice,
            sender,
            sender,
            0,
            gasPriceBid,
            depositCalldata
          )
        )[0],
        retryableGasArgs.maxGasPercentIncrease ||
          BigNumber.from(DEFAULT_MAX_GAS_PERCENT_INCREASE)
      )
    if (
      l1GatewayAddress === l1CustomGatewayAddress &&
      maxGas.lt(MIN_CUSTOM_DEPOSIT_MAXGAS)
    ) {
      // For insurance, we set a sane minimum max gas for the custom gateway
      maxGas = MIN_CUSTOM_DEPOSIT_MAXGAS
    }
    // 4. Calculate total required callvalue
    let totalEthCallvalueToSend = overrides && (await overrides.value)
    if (
      !totalEthCallvalueToSend ||
      BigNumber.from(totalEthCallvalueToSend).isZero()
    ) {
      totalEthCallvalueToSend = await maxSubmissionPrice.add(
        gasPriceBid.mul(maxGas)
      )
    }

    return {
      maxGas,
      gasPriceBid,
      l1CallValue: BigNumber.from(totalEthCallvalueToSend),
      maxSubmissionCost: maxSubmissionPrice,
      destinationAddress: to,
      amount,
      erc20L1Address,
    }
  }

  /**
   * Token deposit; if no value given, calculates and includes minimum necessary value to fund L2 side of execution
   */
  public async deposit(
    params: DepositParams | DepositInputParams,
    overrides: PayableOverrides = {}
  ): Promise<ContractTransaction> {
    const depositInput: DepositParams = isDepositInputParams(params)
      ? await this.getDepositTxParams(params)
      : params

    return this.l1Bridge.deposit(depositInput, overrides)
  }

  public async estimateGasDeposit(
    params: DepositParams | DepositInputParams,
    overrides: PayableOverrides = {}
  ) {
    const depositInput: DepositParams = isDepositInputParams(params)
      ? await this.getDepositTxParams(params)
      : params
    return this.l1Bridge.estimateGasDeposit(depositInput, overrides)
  }

  public async getL1EthBalance(): Promise<BigNumber> {
    return this.l1Bridge.getL1EthBalance()
  }

  public async getL2EthBalance(): Promise<BigNumber> {
    return this.l2Bridge.getL2EthBalance()
  }

  public getL2Transaction(
    l2TransactionHash: string
  ): Promise<TransactionReceipt> {
    return BridgeHelper.getL2Transaction(l2TransactionHash, this.l2Provider)
  }

  public getL1Transaction(
    l1TransactionHash: string
  ): Promise<TransactionReceipt> {
    return BridgeHelper.getL1Transaction(l1TransactionHash, this.l1Provider)
  }

  /**
   * get hash of regular L2 txn from corresponding inbox sequence number
   */
  public calculateL2TransactionHash(
    inboxSequenceNumber: BigNumber,
    l2ChainId?: BigNumber
  ): Promise<string> {
    return BridgeHelper.calculateL2TransactionHash(
      inboxSequenceNumber,
      l2ChainId || this.l2Provider
    )
  }
  /**
   * Hash of L2 side of retryable txn; txn gets generated automatically and is formatted as tho user submitted
   */
  public calculateL2RetryableTransactionHash(
    inboxSequenceNumber: BigNumber,
    l2ChainId?: BigNumber
  ): Promise<string> {
    return BridgeHelper.calculateL2RetryableTransactionHash(
      inboxSequenceNumber,
      l2ChainId || this.l2Provider
    )
  }

  /**
   * Hash of L2 ArbOs generated "auto-redeem" transaction; if it succeeded, a transaction queryable by {@link calculateL2RetryableTransactionHash} will then be created
   */
  public calculateRetryableAutoRedeemTxnHash(
    inboxSequenceNumber: BigNumber,
    l2ChainId?: BigNumber
  ): Promise<string> {
    return BridgeHelper.calculateRetryableAutoRedeemTxnHash(
      inboxSequenceNumber,
      l2ChainId || this.l2Provider
    )
  }

  public async getInboxSeqNumFromContractTransaction(
    l1Transaction: TransactionReceipt
  ): Promise<BigNumber[] | undefined> {
    return BridgeHelper.getInboxSeqNumFromContractTransaction(l1Transaction)
  }

  /**
   * Convenience method to directly retrieve retryable hash from an l1 transaction
   */
  public async getL2TxHashByRetryableTicket(
    l1Transaction: string | ContractReceipt
  ): Promise<string> {
    if (typeof l1Transaction == 'string') {
      l1Transaction = await this.getL1Transaction(l1Transaction)
    }
    const inboxSeqNum = await this.getInboxSeqNumFromContractTransaction(
      l1Transaction
    )

    if (!inboxSeqNum) throw new Error('Inbox not triggered')
    return this.calculateL2RetryableTransactionHash(inboxSeqNum[0])
  }

  public async redeemRetryableTicket(
    l1Transaction: string | ContractReceipt,
    waitTimeForL2Receipt = 900000, // 15 minutes
    overrides: PayableOverrides = {}
  ): Promise<ContractTransaction> {
    if (typeof l1Transaction == 'string') {
      l1Transaction = await this.getL1Transaction(l1Transaction)
    }
    const inboxSeqNum = await this.getInboxSeqNumFromContractTransaction(
      l1Transaction
    )
    if (!inboxSeqNum) throw new Error('Inbox not triggered')

    const l2TxnHash = await this.calculateL2TransactionHash(inboxSeqNum[0])
    console.log('waiting for retryable ticket...', l2TxnHash)

    const l2Txn = await this.l2Provider.waitForTransaction(
      l2TxnHash,
      undefined,
      waitTimeForL2Receipt
    )
    if (!l2Txn) throw new Error('retryable ticket not found')
    console.log('retryable ticket found!')
    if (l2Txn.status === 0) {
      console.warn('retryable ticket failed', l2Txn)
      throw new Error('l2 txn failed')
    }
    const retryHash = await BridgeHelper.calculateL2RetryableTransactionHash(
      inboxSeqNum[0],
      this.l2Provider
    )
    console.log('Redeeming retryable ticket:', retryHash)
    return this.l2Bridge.arbRetryableTx.redeem(retryHash, overrides)
  }

  public async cancelRetryableTicket(
    l1Transaction: string | ContractReceipt,
    waitTimeForL2Receipt = 900000, // 15 minutes
    overrides: PayableOverrides = {}
  ): Promise<ContractTransaction> {
    if (typeof l1Transaction == 'string') {
      l1Transaction = await this.getL1Transaction(l1Transaction)
    }
    const inboxSeqNum = await this.getInboxSeqNumFromContractTransaction(
      l1Transaction
    )
    if (!inboxSeqNum) throw new Error('Inbox not triggered')

    const l2TxnHash = await this.calculateL2TransactionHash(inboxSeqNum[0])
    console.log('waiting for retryable ticket...', l2TxnHash)

    const l2Txn = await this.l2Provider.waitForTransaction(
      l2TxnHash,
      undefined,
      waitTimeForL2Receipt
    )
    if (!l2Txn) throw new Error('retryable ticket not found')
    console.log('retryable ticket found!')
    if (l2Txn.status === 0) {
      console.warn('retryable ticket failed', l2Txn)
      throw new Error('l2 txn failed')
    }
    const redemptionTxHash =
      await BridgeHelper.calculateL2RetryableTransactionHash(
        inboxSeqNum[0],
        this.l2Provider
      )
    console.log(`Ensuring txn hasn't been redeemed:`)

    const redemptionRec = await this.l2Provider.getTransactionReceipt(
      redemptionTxHash
    )
    if (redemptionRec && redemptionRec.status === 1) {
      throw new Error(
        `Can't cancel retryable, it's already been redeemed: ${redemptionTxHash}`
      )
    }
    console.log(`Hasn't been redeemed yet, calling cancel now`)
    return this.l2Bridge.arbRetryableTx.cancel(redemptionTxHash, overrides)
  }

  public getWithdrawalsInL2Transaction(
    l2Transaction: TransactionReceipt
  ): L2ToL1EventResult[] {
    return BridgeHelper.getWithdrawalsInL2Transaction(l2Transaction)
  }

  public async getDepositTokenEventData(
    l1Transaction: TransactionReceipt
  ): Promise<DepositInitiated[]> {
    return BridgeHelper.getDepositTokenEventData(l1Transaction)
  }

  /**
   * Attempt to execute an outbox message; must be confirmed to succeed (i.e., confirmation delay must have passed)
   */
  public async triggerL2ToL1Transaction(
    batchNumber: BigNumber,
    indexInBatch: BigNumber,
    singleAttempt = false
  ): Promise<ContractTransaction> {
    const outboxAddress = await this.getOutboxAddressByBatchNum(batchNumber)
    return BridgeHelper.triggerL2ToL1Transaction(
      batchNumber,
      indexInBatch,
      outboxAddress,
      this.l2Provider,
      this.l1Signer,
      singleAttempt
    )
  }

  public tryOutboxExecute(
    outboxAddress: string,
    batchNumber: BigNumber,
    proof: Array<string>,
    path: BigNumber,
    l2Sender: string,
    l1Dest: string,
    l2Block: BigNumber,
    l1Block: BigNumber,
    timestamp: BigNumber,
    amount: BigNumber,
    calldataForL1: string
  ): Promise<ContractTransaction> {
    return BridgeHelper.tryOutboxExecute(
      {
        batchNumber,
        proof,
        path,
        l2Sender,
        l1Dest,
        l2Block,
        l1Block,
        timestamp,
        amount,
        calldataForL1,
      },
      outboxAddress,
      this.l1Signer
    )
  }

  public tryGetProofOnce(
    batchNumber: BigNumber,
    indexInBatch: BigNumber
  ): Promise<MessageBatchProofInfo | null> {
    return BridgeHelper.tryGetProofOnce(
      batchNumber,
      indexInBatch,
      this.l2Provider
    )
  }

  public tryGetProof(
    batchNumber: BigNumber,
    indexInBatch: BigNumber,
    retryDelay = 500
  ): Promise<MessageBatchProofInfo> {
    return BridgeHelper.tryGetProof(
      batchNumber,
      indexInBatch,
      this.l2Provider,
      retryDelay
    )
  }

  public waitUntilOutboxEntryCreated(
    batchNumber: BigNumber,
    outboxAddress: string
  ): Promise<void> {
    return BridgeHelper.waitUntilOutboxEntryCreated(
      batchNumber,
      outboxAddress,
      this.l1Provider
    )
  }

  /**
   * Return receipt of retryable transaction after execution
   */
  public async waitForRetryableReceipt(
    seqNum: BigNumber,
    confirmations?: number
  ): Promise<TransactionReceipt> {
    return BridgeHelper.waitForRetryableReceipt(
      seqNum,
      this.l2Provider,
      confirmations
    )
  }

  /**
   * All withdrawals from given token
   */
  public async getTokenWithdrawEventData(
    l1TokenAddress: string,
    fromAddress?: string,
    filter?: Filter
  ): Promise<WithdrawalInitiated[]> {
    const gatewayAddress = await this.l2Bridge.l2GatewayRouter.getGateway(
      l1TokenAddress
    )

    return BridgeHelper.getTokenWithdrawEventData(
      this.l2Provider,
      gatewayAddress,
      l1TokenAddress,
      fromAddress,
      filter
    )
  }

  /**
   * All withdrawals from given gateway
   */

  public async getGatewayWithdrawEventData(
    gatewayAddress: string,
    fromAddress?: string,
    filter?: Filter
  ): Promise<WithdrawalInitiated[]> {
    return BridgeHelper.getTokenWithdrawEventData(
      this.l2Provider,
      gatewayAddress,
      undefined,
      fromAddress,
      filter
    )
  }

  public async getL2ToL1EventData(
    fromAddress: string,
    filter?: Filter
  ): Promise<L2ToL1EventResult[]> {
    return BridgeHelper.getL2ToL1EventData(fromAddress, this.l2Provider, filter)
  }

  public async getOutboxAddressByBatchNum(
    batchNum: BigNumber
  ): Promise<string> {
    const inbox = Inbox__factory.connect(
      (await this.l1Bridge.getInbox()).address,
      this.l1Provider
    )
    const bridge = await Bridge__factory.connect(
      await inbox.bridge(),
      this.l1Provider
    )
    const oldOutboxAddress = await bridge.allowedOutboxList(0)
    let newOutboxAddress: string
    try {
      newOutboxAddress = await bridge.allowedOutboxList(1)
    } catch {
      // new outbox not yet deployed; using old outbox
      return oldOutboxAddress
    }
    const oldOutbox = OldOutbox__factory.connect(
      oldOutboxAddress,
      this.l1Provider
    )
    const lastOldOutboxBatchNumber = await oldOutbox.outboxesLength()

    return batchNum.lt(lastOldOutboxBatchNumber)
      ? oldOutboxAddress
      : newOutboxAddress
  }
  /**
   * Returns {@link OutgoingMessageState} for given outgoing message
   */
  public async getOutGoingMessageState(
    batchNumber: BigNumber,
    indexInBatch: BigNumber
  ): Promise<OutgoingMessageState> {
    const outboxAddress = await this.getOutboxAddressByBatchNum(batchNumber)
    return BridgeHelper.getOutGoingMessageState(
      batchNumber,
      indexInBatch,
      outboxAddress,
      this.l1Provider,
      this.l2Provider
    )
  }

  public async getERC20L2Address(erc20L1Address: string): Promise<string> {
    return this.l1Bridge.getERC20L2Address(erc20L1Address)
  }

  public async getERC20L1Address(
    erc20L2Address: string
  ): Promise<string | null> {
    return this.l2Bridge.getERC20L1Address(erc20L2Address)
  }

  public async withdrawETH(
    value: BigNumber,
    destinationAddress?: string,
    overrides: PayableOverrides = {}
  ): Promise<ContractTransaction> {
    return this.l2Bridge.withdrawETH(value, destinationAddress, overrides)
  }

  public async withdrawERC20(
    erc20l1Address: string,
    amount: BigNumber,
    destinationAddress?: string,
    overrides: PayableOverrides = {}
  ): Promise<ContractTransaction> {
    return this.l2Bridge.withdrawERC20(
      erc20l1Address,
      amount,
      destinationAddress,
      overrides
    )
  }

  public isWhiteListed(
    address: string,
    whiteListAddress: string
  ): Promise<boolean> {
    return BridgeHelper.isWhiteListed(
      address,
      whiteListAddress,
      this.l1Provider
    )
  }

  public async setGateways(
    tokenAddresses: string[],
    gatewayAddresses: string[]
  ): Promise<ContractTransaction> {
    const gasPriceBid = await this.l2Provider.getGasPrice()

    const maxSubmissionPrice = (
      await this.l2Bridge.getTxnSubmissionPrice(
        // 20 per address, 100 as buffer/ estimate for any additional calldata
        300 + 20 * (tokenAddresses.length + gatewayAddresses.length)
      )
    )[0]
    return this.l1GatewayRouter.functions.setGateways(
      tokenAddresses,
      gatewayAddresses,
      0,
      gasPriceBid,
      maxSubmissionPrice,
      {
        value: maxSubmissionPrice,
      }
    )
  }
  public async getL1GatewaySetEventData(
    _l1GatewayRouterAddress?: string
  ): Promise<GatewaySet[]> {
    if (this.isCustomNetwork && !_l1GatewayRouterAddress)
      throw new Error('Must supply _l1GatewayRouterAddress for custom network ')

    const l1GatewayRouterAddress =
      _l1GatewayRouterAddress ||
      this.l1Bridge.network.tokenBridge.l1GatewayRouter
    if (!l1GatewayRouterAddress)
      throw new Error('No l2GatewayRouterAddress provided')

    return BridgeHelper.getGatewaySetEventData(
      l1GatewayRouterAddress,
      this.l1Provider
    )
  }

  public async getL2GatewaySetEventData(
    _l2GatewayRouterAddress?: string
  ): Promise<GatewaySet[]> {
    if (this.isCustomNetwork && !_l2GatewayRouterAddress)
      throw new Error('Must supply _l2GatewayRouterAddress for custom network ')

    const l2GatewayRouterAddress =
      _l2GatewayRouterAddress ||
      this.l1Bridge.network.tokenBridge.l2GatewayRouter
    if (!l2GatewayRouterAddress)
      throw new Error('No l2GatewayRouterAddress provided')

    return BridgeHelper.getGatewaySetEventData(
      l2GatewayRouterAddress,
      this.l2Provider
    )
  }

  public async getTokenBalanceBatch(
    userAddr: string,
    tokenAddrs: Array<string>,
    targetNetwork: 'L1' | 'L2'
  ): Promise<Array<{ tokenAddr: string; balance: BigNumber | undefined }>> {
    const iface = ERC20__factory.createInterface()

    const balanceCalls = tokenAddrs.map(token => ({
      target: token,
      funcFragment: iface.functions['balanceOf(address)'],
      values: [userAddr],
    }))

    type ExpectedReturnType = Await<ReturnType<ERC20['functions']['balanceOf']>>

    const bridge = targetNetwork === 'L1' ? this.l1Bridge : this.l2Bridge
    const res = await bridge.getMulticallAggregate(balanceCalls)
    return res.map((bal, index) => ({
      tokenAddr: tokenAddrs[index],
      balance: bal ? (bal as ExpectedReturnType)[0] : undefined,
    }))
  }
}
