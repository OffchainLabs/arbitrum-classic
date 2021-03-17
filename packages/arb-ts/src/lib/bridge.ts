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
import { Signer, BigNumber, ethers, ContractReceipt } from 'ethers'
import { L1Bridge } from './l1Bridge'
import { L2Bridge } from './l2Bridge'
import { BridgeFactory } from './abi/BridgeFactory'
import { OutboxFactory } from './abi/OutboxFactory'

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

  public async approveToken(erc20L1Address: string) {
    return this.l1Bridge.approveToken(erc20L1Address)
  }

  public async depositETH(value: BigNumber, destinationAddress?: string) {
    return this.l1Bridge.depositETH(value, destinationAddress)
  }

  public async depositAsERC20(
    erc20L1Address: string,
    amount: BigNumber,
    maxGas: BigNumber,
    gasPriceBid: BigNumber,
    destinationAddress?: string
  ) {
    return this.l1Bridge.depositAsERC20(
      erc20L1Address,
      amount,
      maxGas,
      gasPriceBid,
      destinationAddress
    )
  }
  public async depositAsERC777(
    erc20L1Address: string,
    amount: BigNumber,
    maxGas: BigNumber,
    gasPriceBid: BigNumber,
    destinationAddress?: string
  ) {
    return this.l1Bridge.depositAsERC777(
      erc20L1Address,
      amount,
      maxGas,
      gasPriceBid,
      destinationAddress
    )
  }
  public getAndUpdateL1TokenData(erc20l1Address: string) {
    return this.l1Bridge.getAndUpdateL1TokenData(erc20l1Address)
  }

  public async getAndUpdateL1EthBalance() {
    return this.l1Bridge.getAndUpdateL1EthBalance()
  }

  public async triggerL2ToL1Transaction(l2TransactionHash: string) {
    const l2ToL1Event = this.arbSys.interface.getEvent('L2ToL1Transaction')
    const eventTopic = this.arbSys.interface.getEventTopic(l2ToL1Event)

    const txReceipt = await this.l2Provider.getTransactionReceipt(
      l2TransactionHash
    )

    const logs = txReceipt.logs.filter(log => log.topics[0] === eventTopic)

    if (logs.length !== 1)
      throw new Error('Not exactly 1 log emitted of L2 to L1 tx?')

    const log = this.arbSys.interface.parseLog(logs[0])
    const {
      caller,
      destination,
      uniqueId,
      batchNumber,
      indexInBatch,
      arbBlockNum,
      ethBlockNum,
      timestamp: l2LogTimestamp,
      callvalue,
      data,
    } = log.args

    console.log('going to get proof')
    const {
      proof,
      path,
      l2Sender,
      l1Dest,
      l2Block,
      l1Block,
      timestamp: proofTimestamp,
      amount,
      calldataForL1,
    } = await this.tryGetProof(batchNumber, indexInBatch)

    console.log('got proof')

    const inbox = await this.l1Bridge.getInbox()
    const bridgeAddress = await inbox.bridge()
    const bridge = await BridgeFactory.connect(
      bridgeAddress,
      this.l1Bridge.l1Provider
    )

    const outboxIndex = BigNumber.from(0)
    const activeOutbox = await bridge.allowedOutboxList(outboxIndex)
    try {
      // index 1 should not exist
      await bridge.allowedOutboxList(1)
      console.error("There is more than 1 outbox registered with the bridge?!")
    } catch(e) {
      // this should fail!
      console.log("All is good")
    }

    const outboxExecuteTransactionReceipt = await this.tryOutboxExecute(
      activeOutbox,
      outboxIndex,
      proof,
      path,
      l2Sender,
      l1Dest,
      l2Block,
      l1Block,
      proofTimestamp,
      amount,
      calldataForL1
    )
    return outboxExecuteTransactionReceipt
  }

  private wait = (ms: number) => new Promise(res => setTimeout(res, ms))
  public tryOutboxExecute = async (
    activeOutboxAddress: string,
    outboxNumber: BigNumber,
    proof: Array<string>,
    path: BigNumber,
    l2Sender: string,
    l1Dest: string,
    l2Block: BigNumber,
    l1Block: BigNumber,
    timestamp: BigNumber,
    amount: BigNumber,
    calldataForL1: string,
    retryDelay: number = 500
  ): Promise<ContractReceipt> => {
    const outbox = await OutboxFactory.connect(
      activeOutboxAddress,
      this.l1Bridge.l1Provider
    ).connect(this.l1Bridge.l1Signer)

    try {
      // TODO: wait until assertion is confirmed before execute
      // We can predict and print number of missing blocks
      // if not challenged
      const outboxExecute = await outbox.executeTransaction(
        outboxNumber,
        proof,
        path,
        l2Sender,
        l1Dest,
        l2Block,
        l1Block,
        timestamp,
        amount,
        calldataForL1
      )
      console.log(`Transaction hash: ${outboxExecute.hash}`)
      console.log('Waiting for receipt')
      const receipt = await outboxExecute.wait()
      console.log('Receipt emitted')
      return receipt
    } catch (e) {
      console.log('failed to execute tx')
      console.log(e)
      console.log('Waiting for delay before retrying')
      // TODO: should exponential backoff?
      await this.wait(retryDelay)
      console.log('Retrying now')
      return this.tryOutboxExecute(
        activeOutboxAddress,
        outboxNumber,
        proof,
        path,
        l2Sender,
        l1Dest,
        l2Block,
        l1Block,
        timestamp,
        amount,
        calldataForL1,
        retryDelay
      )
    }
  }

  public tryGetProof = async (
    batchNumber: BigNumber,
    indexInBatch: BigNumber,
    retryDelay: number = 500
  ): Promise<{
    proof: Array<string>
    path: BigNumber
    l2Sender: string
    l1Dest: string
    l2Block: BigNumber
    l1Block: BigNumber
    timestamp: BigNumber
    amount: BigNumber
    calldataForL1: string
  }> => {
    const nodeInterfaceAddress = '0x00000000000000000000000000000000000000C8'

    const contractInterface = new ethers.utils.Interface([
      `function lookupMessageBatchProof(uint256 batchNum, uint64 index)
          external
          view
          returns (
              bytes32[] memory proof,
              uint256 path,
              address l2Sender,
              address l1Dest,
              uint256 l2Block,
              uint256 l1Block,
              uint256 timestamp,
              uint256 amount,
              bytes memory calldataForL1
          )`,
    ])
    const nodeInterface = new ethers.Contract(
      nodeInterfaceAddress,
      contractInterface
    ).connect(this.l2Signer.provider!)

    try {
      const res = await nodeInterface.callStatic.lookupMessageBatchProof(
        batchNumber,
        indexInBatch
      )
      return res
    } catch (e) {
      const expectedError = "batch doesn't exist"
      if (
        e &&
        e.error &&
        e.error.message &&
        e.error.message === expectedError
      ) {
        console.log(
          'Withdrawal detected, but batch not created yet. Going to wait a bit.'
        )
      } else {
        console.log("Withdrawal proof didn't work. Not sure why")
        console.log(e)
        console.log('Going to try again after waiting')
      }
      await this.wait(retryDelay)
      console.log('New attempt starting')
      // TODO: should exponential backoff?
      return this.tryGetProof(batchNumber, indexInBatch, retryDelay)
    }
  }
}
