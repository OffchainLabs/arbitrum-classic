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
import { Bridge__factory } from './abi/factories/Bridge__factory'
import { Outbox__factory } from './abi/factories/Outbox__factory'
import { TransactionOverrides, BridgeHelper } from './bridge_helpers'

import { ArbSys } from './abi/ArbSys'

const { Zero } = constants







export class Bridge extends L2Bridge {
  l1Bridge: L1Bridge
  walletAddressCache?: string

  constructor(
    erc20BridgeAddress: string,
    arbERC20BridgeAddress: string,
    ethSigner: Signer,
    arbSigner: Signer
  ) {
    super(arbERC20BridgeAddress, arbSigner)

    this.l1Bridge = new L1Bridge(erc20BridgeAddress, ethSigner)
  }
  public updateAllBalances() {
    this.updateAllTokens()
    this.getAndUpdateL1EthBalance()
    this.getAndUpdateL2EthBalance()
  }

  public async updateAllTokens() {
    const l1Tokens = await this.l1Bridge.updateAllL1Tokens()
    const l2Tokens = await this.updateAllL2Tokens()
    return { l1Tokens, l2Tokens }
  }

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

  public async approveToken(
    erc20L1Address: string,
    overrides?: TransactionOverrides
  ) {
    return this.l1Bridge.approveToken(erc20L1Address, overrides)
  }

  public async depositETH(
    value: BigNumber,
    destinationAddress?: string,
    overrides?: TransactionOverrides
  ) {
    return this.l1Bridge.depositETH(value, destinationAddress, overrides)
  }

  public async depositAsERC20(
    erc20L1Address: string,
    amount: BigNumber,
    maxGas: BigNumber,
    gasPriceBid: BigNumber,
    destinationAddress?: string,
    overrides?: TransactionOverrides
  ) {
    // TODO: this will need to (somehow) input the calldata size
    const maxSubmissionPrice = (await this.getTxnSubmissionPrice(Zero))[0]
    return this.l1Bridge.depositAsERC20(
      erc20L1Address,
      amount,
      maxSubmissionPrice,
      maxGas,
      gasPriceBid,
      destinationAddress,
      overrides
    )
  }
  public async depositAsERC777(
    erc20L1Address: string,
    amount: BigNumber,
    maxGas: BigNumber,
    gasPriceBid: BigNumber,
    destinationAddress?: string,
    overrides?: TransactionOverrides
  ) {
    // TODO: this will need to (somehow) input the calldata size
    const maxSubmissionPrice = (await this.getTxnSubmissionPrice(Zero))[0]

    return this.l1Bridge.depositAsERC777(
      erc20L1Address,
      amount,
      maxSubmissionPrice,
      maxGas,
      gasPriceBid,
      destinationAddress,
      overrides
    )
  }
  public getAndUpdateL1TokenData(erc20l1Address: string) {
    return this.l1Bridge.getAndUpdateL1TokenData(erc20l1Address)
  }

  public async getAndUpdateL1EthBalance() {
    return this.l1Bridge.getAndUpdateL1EthBalance()
  }

  public getL2Transaction(l2TransactionHash: string) {
    return BridgeHelper.getL2Transaction(
      l2TransactionHash,
      this.l2Provider
    )
  }

  public getL1Transaction(l1TransactionHash: string) {
    return BridgeHelper.getL1Transaction(l1TransactionHash, this.l1Bridge.l1Provider)
  }

  public calculateL2TransactionHash(
    inboxSequenceNumber: BigNumber,
    l2ChainId?: BigNumber
  ) {
    return BridgeHelper.calculateL2RetryableTransactionHash(
      inboxSequenceNumber,
      l2ChainId || this.l2Provider
    )
  }

  public calculateL2RetryableTransactionHash(
    inboxSequenceNumber: BigNumber,
    l2ChainId?: BigNumber
  ): Promise<string> {
    return BridgeHelper.calculateL2RetryableTransactionHash(
      inboxSequenceNumber,
      l2ChainId || this.l2Provider,
    )
  }

  public async getInboxSeqNumFromContractTransaction(
    l2Transaction: ethers.providers.TransactionReceipt
  ): Promise<Array<BigNumber> | undefined> {
    return BridgeHelper.getInboxSeqNumFromContractTransaction(
      l2Transaction,
      // TODO: we don't need to actually make this query if random address fetches interface
      (await this.l1Bridge.getInbox()).address
    )
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
      l2Transaction, this.l2Provider, ARB_SYS_ADDRESS
    )
  }

  public getDepositTokenEventData(
    l1Transaction: ethers.providers.TransactionReceipt,
    tokenType: 'ERC20' | 'ERC777' = 'ERC20'
  ) {
    return BridgeHelper.getDepositTokenEventData(
      l1Transaction,
      tokenType,
      this.arbTokenBridge.address
    )
  }

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

  public tryOutboxExecute (
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
      activeOutboxAddress,
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
      this.l1Bridge.l1Signer
    )
  }

  public tryGetProofOnce (
    batchNumber: BigNumber,
    indexInBatch: BigNumber
  ) {
    return BridgeHelper.tryGetProofOnce(
      batchNumber,
      indexInBatch,
      this.l2Provider
    )
  }

  public tryGetProof (
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

  public waitUntilOutboxEntryCreated (
    batchNumber: BigNumber,
    activeOutboxAddress: string
  ) {
    return BridgeHelper.waitUntilOutboxEntryCreated(
      batchNumber,
      activeOutboxAddress,
      this.l1Bridge.l1Provider
    )
  }

  public async waitForRetriableReceipt(seqNum: BigNumber) {
    return BridgeHelper.waitForRetriableReceipt(seqNum, this.l2Provider)
  }

  public async getTokenWithdrawEventData(destinationAddress: string) {
    return BridgeHelper.getTokenWithdrawEventData(destinationAddress, this.arbTokenBridge.address, this.l2Provider)
  }

  public async getL2ToL1EventData(destinationAddress: string) {
    return BridgeHelper.getL2ToL1EventData(
      destinationAddress,
      this.l2Provider,
      ARB_SYS_ADDRESS
    )
  }
}
