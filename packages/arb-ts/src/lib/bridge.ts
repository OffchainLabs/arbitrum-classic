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
import { Signer, BigNumber, ethers, ContractReceipt, constants } from 'ethers'
import { L1Bridge } from './l1Bridge'
import { L2Bridge, ARB_SYS_ADDRESS } from './l2Bridge'
import { BridgeHelper } from './bridge_helpers'
import { PayableOverrides } from '@ethersproject/contracts'

const { Zero } = constants

interface RetryableGasArgs {
  maxSubmissionPrice?: BigNumber
  maxGas?: BigNumber
  gasPriceBid?: BigNumber
  maxSubmissionPriceIncreaseRatio?: BigNumber
}

/**
 * Main class for accessing token bridge methods; inherits methods from {@link L1Bridge} and {@link L2Bridge}
 */
export class Bridge extends L2Bridge {
  l1Bridge: L1Bridge
  walletAddressCache?: string
  outboxAddressCache?: string

  constructor(
    erc20BridgeAddress: string,
    arbERC20BridgeAddress: string,
    ethSigner: Signer,
    arbSigner: Signer
  ) {
    Promise.all([ethSigner.getAddress(), arbSigner.getAddress()]).then(
      ([ethSignerAddress, arbSignerAddress]) => {
        if (ethSignerAddress !== arbSignerAddress) {
          throw new Error('L1 & L2 wallets must be of the same address')
        }
      }
    )

    super(arbERC20BridgeAddress, arbSigner)

    this.l1Bridge = new L1Bridge(erc20BridgeAddress, ethSigner)
  }
  public updateAllBalances() {
    this.updateAllTokens()
    this.getAndUpdateL1EthBalance()
    this.getAndUpdateL2EthBalance()
  }
  /**
   * Update state of all tracked tokens (balance, allowance), etc. and returns state
   */
  public async updateAllTokens() {
    const l1Tokens = await this.l1Bridge.updateAllL1Tokens()
    const l2Tokens = await this.updateAllL2Tokens()
    return { l1Tokens, l2Tokens }
  }
  /**
   * Update target token (balance, allowance), etc. and state
   */
  public async updateTokenData(erc20l1Address: string) {
    const l1Data = await this.getAndUpdateL1TokenData(erc20l1Address)
    const l2Data = await this.getAndUpdateL2TokenData(erc20l1Address)
    return { l1Data, l2Data }
  }

  get l1Tokens() {
    return this.l1Bridge.l1Tokens
  }

  get l1EthBalance() {
    return this.l1Bridge.l1EthBalance
  }

  get ethERC20Bridge() {
    return this.l1Bridge.ethERC20Bridge
  }

  /**
   * Set allowance for L1 bridge contract
   */
  public async approveToken(
    erc20L1Address: string,
    overrides?: PayableOverrides
  ) {
    return this.l1Bridge.approveToken(erc20L1Address, overrides)
  }

  /**
   * Deposit ether from L1 to L2. Users L1MessageType_L2FundedByL1 txn type ( type 7)
   */
  public async depositETH(
    value: BigNumber,
    _maxSubmissionPriceIncreaseRatio?: BigNumber,
    overrides?: PayableOverrides
  ) {
    const maxSubmissionPriceIncreaseRatio =
      _maxSubmissionPriceIncreaseRatio || BigNumber.from(13)

    const maxSubmissionPrice = (await this.getTxnSubmissionPrice(0))[0]
      .mul(maxSubmissionPriceIncreaseRatio)
      .div(BigNumber.from(10))

    return this.l1Bridge.depositETH(value, maxSubmissionPrice, overrides)
  }

  /**
   * Token deposit; if no value given, calculates and includes minimum necessary value to fund L2 side of execution
   */
  public async deposit(
    erc20L1Address: string,
    amount: BigNumber,
    retryableGasArgs: RetryableGasArgs = {},
    destinationAddress?: string,
    overrides?: PayableOverrides
  ) {
    const gasPriceBid =
      retryableGasArgs.gasPriceBid || (await this.l2Provider.getGasPrice())

    const sender = await this.l1Bridge.l1Signer.getAddress()

    const [
      isDeployed,
      depositCalldata,
    ] = await this.ethERC20Bridge.getDepositCalldata(
      erc20L1Address,
      sender,
      destinationAddress ? destinationAddress : sender,
      amount,
      '0x'
    )
    const expectedGas = await this.l2Provider.estimateGas({
      from: this.ethERC20Bridge.address,
      to: this.arbTokenBridge.address,
      data: depositCalldata,
    })
    const maxGas = retryableGasArgs.maxGas || expectedGas

    const maxSubmissionPriceIncreaseRatio =
      retryableGasArgs.maxSubmissionPriceIncreaseRatio || BigNumber.from(13)

    const maxSubmissionPrice = (
      await this.getTxnSubmissionPrice(depositCalldata.length - 2)
    )[0]
      .mul(maxSubmissionPriceIncreaseRatio)
      .div(BigNumber.from(10))

    // calculate required forwarding gas
    let ethDeposit = overrides && (await overrides.value)
    if (!ethDeposit || BigNumber.from(ethDeposit).isZero()) {
      ethDeposit = await maxSubmissionPrice.add(gasPriceBid.mul(maxGas))
      // TODO: might reactivate if we switch to arb-os deducting from sender's EOA
      // const l2Balance = await this.getAndUpdateL2EthBalance()

      // const requiredEth = await maxSubmissionPrice.add(gasPriceBid.mul(maxGas))

      // if (l2Balance.lt(requiredEth)) {
      //   console.info(
      //     'insufficient L2 balance to pay for retryable:',
      //     l2Balance.toNumber(),
      //     'Depositing additional ETH'
      //   )
      //   ethDeposit = requiredEth.sub(l2Balance)
      // } else {
      //   console.info('l2 account adequately funded for retryable')
      // }
    }

    return this.l1Bridge.deposit(
      erc20L1Address,
      amount,
      maxSubmissionPrice,
      maxGas,
      gasPriceBid,
      destinationAddress,
      { ...overrides, value: ethDeposit }
    )
  }

  public getAndUpdateL1TokenData(erc20l1Address: string) {
    return this.l1Bridge.getAndUpdateL1TokenData(erc20l1Address)
  }

  public async getAndUpdateL1EthBalance() {
    return this.l1Bridge.getAndUpdateL1EthBalance()
  }

  public getL2Transaction(l2TransactionHash: string) {
    return BridgeHelper.getL2Transaction(l2TransactionHash, this.l2Provider)
  }

  public getL1Transaction(l1TransactionHash: string) {
    return BridgeHelper.getL1Transaction(
      l1TransactionHash,
      this.l1Bridge.l1Provider
    )
  }

  /**
   * get hash of regular L2 txn from corresponding inbox sequence number
   */
  public calculateL2TransactionHash(
    inboxSequenceNumber: BigNumber,
    l2ChainId?: BigNumber
  ) {
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
  public calculateRetryableAutoReedemTxnHash(
    inboxSequenceNumber: BigNumber,
    l2ChainId?: BigNumber
  ): Promise<string> {
    return BridgeHelper.calculateRetryableAutoReedemTxnHash(
      inboxSequenceNumber,
      l2ChainId || this.l2Provider
    )
  }

  public async getInboxSeqNumFromContractTransaction(
    l1Transaction: ethers.providers.TransactionReceipt
  ): Promise<BigNumber[] | undefined> {
    return BridgeHelper.getInboxSeqNumFromContractTransaction(
      l1Transaction,
      // TODO: we don't need to actually make this query if random address fetches interface
      (await this.l1Bridge.getInbox()).address
    )
  }

  /**
   * Convenience method to directly retrieve retryable hash from an l1 transaction
   */
  public async getL2TxHashByRetryableTicket(
    l1Transaction: string | ContractReceipt
  ) {
    if (typeof l1Transaction == 'string') {
      l1Transaction = await this.getL1Transaction(l1Transaction)
    }
    const inboxSeqNum = await this.getInboxSeqNumFromContractTransaction(
      l1Transaction
    )

    if (!inboxSeqNum) throw new Error('Inbox not triggered')
    return this.calculateL2RetryableTransactionHash(inboxSeqNum[0])
  }

  public getBuddyDeployInL2Transaction(
    l2Transaction: ethers.providers.TransactionReceipt
  ) {
    return BridgeHelper.getBuddyDeployInL2Transaction(l2Transaction)
  }

  public getWithdrawalsInL2Transaction(
    l2Transaction: ethers.providers.TransactionReceipt
  ) {
    return BridgeHelper.getWithdrawalsInL2Transaction(
      l2Transaction,
      this.l2Provider
    )
  }

  public getDepositTokenEventData(
    l1Transaction: ethers.providers.TransactionReceipt
  ) {
    return BridgeHelper.getDepositTokenEventData(
      l1Transaction,
      this.arbTokenBridge.address
    )
  }

  /**
   * Attempt to execute an outbox message; must be confirmed to succeed (i.e., confirmation delay must have passed)
   */
  public async triggerL2ToL1Transaction(
    batchNumber: BigNumber,
    indexInBatch: BigNumber,
    singleAttempt = false
  ) {
    const inbox = await this.l1Bridge.getInbox()
    const bridgeAddress = await inbox.bridge()

    return BridgeHelper.triggerL2ToL1Transaction(
      batchNumber,
      indexInBatch,
      bridgeAddress,
      this.l2Provider,
      this.l1Bridge.l1Signer,
      singleAttempt
    )
  }

  public tryOutboxExecute(
    activeOutboxAddress: string,
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
  ) {
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
      activeOutboxAddress,
      this.l1Bridge.l1Signer
    )
  }

  public tryGetProofOnce(batchNumber: BigNumber, indexInBatch: BigNumber) {
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
  ) {
    return BridgeHelper.tryGetProof(
      batchNumber,
      indexInBatch,
      this.l2Provider,
      retryDelay
    )
  }

  public waitUntilOutboxEntryCreated(
    batchNumber: BigNumber,
    activeOutboxAddress: string
  ) {
    return BridgeHelper.waitUntilOutboxEntryCreated(
      batchNumber,
      activeOutboxAddress,
      this.l1Bridge.l1Provider
    )
  }

  /**
   * Return receipt of retryable transaction after execution
   */
  public async waitForRetriableReceipt(seqNum: BigNumber) {
    return BridgeHelper.waitForRetriableReceipt(seqNum, this.l2Provider)
  }

  public async getTokenWithdrawEventData(destinationAddress: string) {
    return BridgeHelper.getTokenWithdrawEventData(
      destinationAddress,
      this.arbTokenBridge.address,
      this.l2Provider
    )
  }

  public async getL2ToL1EventData(destinationAddress: string) {
    return BridgeHelper.getL2ToL1EventData(destinationAddress, this.l2Provider)
  }

  public async getOutboxAddress() {
    if (this.outboxAddressCache) {
      return this.outboxAddressCache
    }
    const inboxAddress = (await this.l1Bridge.getInbox()).address
    const coreBridgeAddress = await BridgeHelper.getCoreBridgeFromInbox(
      inboxAddress,
      this.l1Bridge.l1Provider
    )
    const outboxAddress = await BridgeHelper.getActiveOutbox(
      coreBridgeAddress,
      this.l1Bridge.l1Provider
    )
    this.outboxAddressCache = outboxAddress
    return outboxAddress
  }

  /**
   * Returns {@link OutgoingMessageState} for given outgoing message
   */
  public async getOutGoingMessageState(
    batchNumber: BigNumber,
    indexInBatch: BigNumber
  ) {
    const outboxAddress = await this.getOutboxAddress()
    return BridgeHelper.getOutgoingMessageState(
      batchNumber,
      indexInBatch,
      outboxAddress,
      this.l1Bridge.l1Provider,
      this.l2Provider
    )
  }
}
