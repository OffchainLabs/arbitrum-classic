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
import { L2Bridge } from './l2Bridge'
import { BridgeHelper } from './bridge_helpers'
import { PayableOverrides } from '@ethersproject/contracts'
import { NODE_INTERFACE_ADDRESS } from './precompile_addresses'
import { NodeInterface__factory } from './abi/factories/NodeInterface__factory'
import { L1ERC20Gateway__factory } from './abi/factories/L1ERC20Gateway__factory'
import networks from './networks'

interface RetryableGasArgs {
  maxSubmissionPrice?: BigNumber
  maxGas?: BigNumber
  gasPriceBid?: BigNumber
  maxSubmissionPriceIncreaseRatio?: BigNumber
}

/**
 * Main class for accessing token bridge methods; inherits methods from {@link L1Bridge} and {@link L2Bridge}
 */
export class Bridge {
  l1Bridge: L1Bridge
  l2Bridge: L2Bridge
  walletAddressCache?: string
  outboxAddressCache?: string

  private constructor(l1BridgeObj: L1Bridge, l2BridgeObj: L2Bridge) {
    this.l1Bridge = l1BridgeObj
    this.l2Bridge = l2BridgeObj
  }

  public updateAllBalances() {
    this.updateAllTokens()
    this.getAndUpdateL1EthBalance()
    this.getAndUpdateL2EthBalance()
  }

  static async init(
    ethSigner: Signer,
    arbSigner: Signer,
    l1GatewayRouterAddress?: string,
    l2GatewayRouterAddress?: string
  ) {
    if (!ethSigner.provider || !arbSigner.provider) {
      throw new Error('Signer needs a provider')
    }

    const ethSignerAddress = await ethSigner.getAddress()
    const arbSignerAddress = await arbSigner.getAddress()
    const l1ChainId = await ethSigner.getChainId()
    const l2ChainId = await arbSigner.getChainId()

    if (ethSignerAddress !== arbSignerAddress) {
      throw new Error('L1 & L2 wallets must be of the same address')
    }

    const l1Network = networks[l1ChainId]
    const l2Network = networks[l2ChainId]

    if (l1Network) {
      if (l1Network.isArbitrum)
        throw new Error('Connected to an Arbitrum networks as the L1...')
      l1GatewayRouterAddress = l1Network.tokenBridge.l1GatewayRouter
    } else if (!l1GatewayRouterAddress) {
      throw new Error(
        'Network not in config, and no l1GatewayRouter Address provided'
      )
    }

    if (l2Network) {
      if (!l2Network.isArbitrum)
        throw new Error('Connected to an L1 network as the L2...')
      l2GatewayRouterAddress = l2Network.tokenBridge.l2GatewayRouter
    } else if (!l2GatewayRouterAddress) {
      throw new Error(
        'Network not in config, and no l2GatewayRouter address provided'
      )
    }

    if (l1Network && l2Network) {
      if (l1Network.partnerChainID !== l2Network.chainID)
        throw new Error('L1 and L2 networks are not connected')
    }

    // check routers are deployed
    const l1RouterCode = await ethSigner.provider.getCode(
      l1GatewayRouterAddress
    )
    if (l1RouterCode === '0x') {
      throw new Error(`No code deployed to ${l1GatewayRouterAddress} in the L1`)
    }

    const l2RouterCode = await arbSigner.provider.getCode(
      l2GatewayRouterAddress
    )
    if (l2RouterCode === '0x') {
      throw new Error(`No code deployed to ${l2GatewayRouterAddress} in the L2`)
    }

    const l1BridgeObj = new L1Bridge(l1GatewayRouterAddress, ethSigner)
    const l2BridgeObj = new L2Bridge(l2GatewayRouterAddress, arbSigner)

    return new Bridge(l1BridgeObj, l2BridgeObj)
  }

  /**
   * Update state of all tracked tokens (balance, allowance), etc. and returns state
   */
  public async updateAllTokens() {
    const l1Tokens = await this.l1Bridge.updateAllL1Tokens()
    const l2Tokens = await this.l2Bridge.updateAllL2Tokens()
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

  get l1GatewayRouter() {
    return this.l1Bridge.l1GatewayRouter
  }

  defaultL1Gateway() {
    return this.l1Bridge.getDefaultL1Gateway()
  }

  /**
   * Set allowance for L1 router contract
   */
  public async approveToken(
    erc20L1Address: string,
    overrides?: PayableOverrides
  ) {
    return this.l1Bridge.approveToken(erc20L1Address, overrides)
  }

  /**
   * Deposit ether from L1 to L2.
   */
  public async depositETH(
    value: BigNumber,
    _maxSubmissionPriceIncreaseRatio?: BigNumber,
    overrides?: PayableOverrides
  ) {
    const maxSubmissionPriceIncreaseRatio =
      _maxSubmissionPriceIncreaseRatio || BigNumber.from(13)

    const maxSubmissionPrice = (await this.l2Bridge.getTxnSubmissionPrice(0))[0]
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
    const l1ChainId = await this.l1Bridge.l1Signer.getChainId()
    const { l1WethGateway: l1WethGatewayAddress } = networks[
      l1ChainId
    ].tokenBridge

    const gasPriceBid =
      retryableGasArgs.gasPriceBid ||
      (await this.l2Bridge.l2Provider.getGasPrice())

    const sender = await this.l1Bridge.l1Signer.getAddress()

    const expectedL1GatewayAddress = await this.l1Bridge.getGatewayAddress(
      erc20L1Address
    )

    let estimateGasCallValue = constants.Zero

    if (l1WethGatewayAddress === expectedL1GatewayAddress) {
      // forwarded deposited eth as call value for weth deposit
      estimateGasCallValue = amount
    }

    const l1Gateway = L1ERC20Gateway__factory.connect(
      expectedL1GatewayAddress,
      this.l1Bridge.l1Provider
    )

    const depositCalldata = await l1Gateway.getOutboundCalldata(
      erc20L1Address,
      sender,
      destinationAddress ? destinationAddress : sender,
      amount,
      '0x'
    )

    const maxSubmissionPriceIncreaseRatio =
      retryableGasArgs.maxSubmissionPriceIncreaseRatio || BigNumber.from(13)

    const maxSubmissionPrice = (
      await this.l2Bridge.getTxnSubmissionPrice(depositCalldata.length - 2)
    )[0]
      .mul(maxSubmissionPriceIncreaseRatio)
      .div(BigNumber.from(10))

    const nodeInterface = NodeInterface__factory.connect(
      NODE_INTERFACE_ADDRESS,
      this.l2Bridge.l2Provider
    )

    const l2Dest = await l1Gateway.counterpartGateway()

    const maxGas = (
      await nodeInterface.estimateRetryableTicket(
        expectedL1GatewayAddress,
        ethers.utils.parseEther('0.05').add(estimateGasCallValue),
        l2Dest,
        estimateGasCallValue,
        maxSubmissionPrice,
        sender,
        sender,
        0,
        0,
        depositCalldata
      )
    )[0]
    console.log('DONE ESTIMATING GAS')

    // calculate required forwarding gas
    let ethDeposit = overrides && (await overrides.value)
    if (!ethDeposit || BigNumber.from(ethDeposit).isZero()) {
      ethDeposit = await maxSubmissionPrice.add(gasPriceBid.mul(maxGas))
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

  public async getAndUpdateL2TokenData(erc20l1Address: string) {
    const l2TokenAddress = await this.l1Bridge.getERC20L2Address(erc20l1Address)
    return this.l2Bridge.getAndUpdateL2TokenData(erc20l1Address, l2TokenAddress)
  }

  public async getAndUpdateL1EthBalance() {
    return this.l1Bridge.getAndUpdateL1EthBalance()
  }

  public async getAndUpdateL2EthBalance() {
    return this.l2Bridge.getAndUpdateL2EthBalance()
  }

  public getL2Transaction(l2TransactionHash: string) {
    return BridgeHelper.getL2Transaction(
      l2TransactionHash,
      this.l2Bridge.l2Provider
    )
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
      l2ChainId || this.l2Bridge.l2Provider
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
      l2ChainId || this.l2Bridge.l2Provider
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
      l2ChainId || this.l2Bridge.l2Provider
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

  public async redeemRetryableTicket(
    l1Transaction: string | ContractReceipt,
    waitTimeForL2Receipt = 900000, // 15 minutes
    overrides?: PayableOverrides
  ) {
    if (typeof l1Transaction == 'string') {
      l1Transaction = await this.getL1Transaction(l1Transaction)
    }
    const inboxSeqNum = await this.getInboxSeqNumFromContractTransaction(
      l1Transaction
    )
    if (!inboxSeqNum) throw new Error('Inbox not triggered')

    const l2TxnHash = await this.calculateL2TransactionHash(inboxSeqNum[0])
    console.log('waiting for retryable ticket...', l2TxnHash)

    const l2Txn = await this.l2Bridge.l2Provider.waitForTransaction(
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
      this.l2Bridge.l2Provider
    )
    console.log('Redeeming retryable ticket:', retryHash)
    return this.l2Bridge.arbRetryableTx.redeem(retryHash)
  }

  public async cancelRetryableTicket(
    l1Transaction: string | ContractReceipt,
    waitTimeForL2Receipt = 900000, // 15 minutes
    overrides?: PayableOverrides
  ) {
    if (typeof l1Transaction == 'string') {
      l1Transaction = await this.getL1Transaction(l1Transaction)
    }
    const inboxSeqNum = await this.getInboxSeqNumFromContractTransaction(
      l1Transaction
    )
    if (!inboxSeqNum) throw new Error('Inbox not triggered')

    const l2TxnHash = await this.calculateL2TransactionHash(inboxSeqNum[0])
    console.log('waiting for retryable ticket...', l2TxnHash)

    const l2Txn = await this.l2Bridge.l2Provider.waitForTransaction(
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
    const redemptionTxHash = await BridgeHelper.calculateL2RetryableTransactionHash(
      inboxSeqNum[0],
      this.l2Bridge.l2Provider
    )
    const redemptionRec = await this.l1Bridge.l1Provider.getTransactionReceipt(
      redemptionTxHash
    )
    if (redemptionRec && redemptionRec.status === 1) {
      throw new Error(
        `Can't cancel retryable, it's already been redeemed: ${redemptionTxHash}`
      )
    }

    return this.l2Bridge.arbRetryableTx.cancel(redemptionTxHash)
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
      this.l2Bridge.l2Provider
    )
  }

  public async getDepositTokenEventData(
    l1Transaction: ethers.providers.TransactionReceipt
  ) {
    const defaultGatewayAddress = (await this.l1Bridge.getDefaultL1Gateway())
      .address
    return BridgeHelper.getDepositTokenEventData(
      l1Transaction,
      defaultGatewayAddress
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
      this.l2Bridge.l2Provider,
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
      this.l2Bridge.l2Provider
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
      this.l2Bridge.l2Provider,
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
  public async waitForRetryableReceipt(seqNum: BigNumber) {
    return BridgeHelper.waitForRetryableReceipt(
      seqNum,
      this.l2Bridge.l2Provider
    )
  }

  public async getTokenWithdrawEventData(
    l1TokenAddress: string,
    destinationAddress: string
  ) {
    const l2ERC20Gateway = await this.l2Bridge.l2GatewayRouter.getGateway(
      l1TokenAddress
    )

    return BridgeHelper.getTokenWithdrawEventData(
      destinationAddress,
      l2ERC20Gateway,
      this.l2Bridge.l2Provider
    )
  }

  public async getL2ToL1EventData(destinationAddress: string) {
    return BridgeHelper.getL2ToL1EventData(
      destinationAddress,
      this.l2Bridge.l2Provider
    )
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
      this.l2Bridge.l2Provider
    )
  }

  public async getERC20L2Address(erc20L1Address: string) {
    return this.l1Bridge.getERC20L2Address(erc20L1Address)
  }

  public async withdrawETH(
    value: BigNumber,
    destinationAddress?: string,
    overrides?: PayableOverrides
  ) {
    return this.l2Bridge.withdrawETH(value, destinationAddress, overrides)
  }

  public async withdrawERC20(
    erc20l1Address: string,
    amount: BigNumber,
    destinationAddress?: string,
    overrides: PayableOverrides = {}
  ) {
    return this.l2Bridge.withdrawERC20(
      erc20l1Address,
      amount,
      destinationAddress,
      overrides
    )
  }

  public isWhiteListed(address: string, whiteListAddress: string) {
    return BridgeHelper.isWhiteListed(
      address,
      whiteListAddress,
      this.l1Bridge.l1Provider
    )
  }

  public async setGateways(
    tokenAddresses: string[],
    gatewayAddresses: string[]
  ) {
    const gasPriceBid = await this.l2Bridge.l2Provider.getGasPrice()

    const maxSubmissionPrice = (
      await this.l2Bridge.getTxnSubmissionPrice(
        // 20 per address, 100 as buffer/ estimate for any additional calldata
        100 + 20 * (tokenAddresses.length + gatewayAddresses.length)
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
  public async getL1GatewaySetEventData() {
    const l1ChainId = await this.l1Bridge.l1Signer.getChainId()
    const l1GatewayRouterAddress =
      networks[l1ChainId].tokenBridge.l1GatewayRouter
    return BridgeHelper.getGatewaySetEventData(
      l1GatewayRouterAddress,
      this.l1Bridge.l1Provider
    )
  }

  public async getL2GatewaySetEventData() {
    const l1ChainId = await this.l1Bridge.l1Signer.getChainId()
    const l2GatewayRouterAddress =
      networks[l1ChainId].tokenBridge.l2GatewayRouter
    return BridgeHelper.getGatewaySetEventData(
      l2GatewayRouterAddress,
      this.l2Bridge.l2Provider
    )
  }
}
