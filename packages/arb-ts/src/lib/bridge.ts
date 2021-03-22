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
import { ArbSys } from './abi/ArbSys'

const { Zero } = constants

// TODO: can we import these interfaces directly from typechain?
export interface L2ToL1EventResult {
  caller: string
  destination: string
  uniqueId: BigNumber
  batchNumber: BigNumber
  indexInBatch: BigNumber
  arbBlockNum: BigNumber
  ethBlockNum: BigNumber
  timestamp: string
  callvalue: BigNumber
  data: string
}

interface BuddyDeployEventResult {
  _sender: string
  _contract: string
  withdrawalId: BigNumber
  success: boolean
}

interface WithdrawTokenEventResult {
  id: BigNumber
  l1Address: string
  amount: BigNumber
  destination: string
  exitNum: BigNumber
}
import { Bridge__factory } from './abi/factories/Bridge__Factory'
import { Outbox__factory } from './abi/factories/Outbox__Factory'

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
    console.warn('init bridge.ts')

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
    // TODO: this will need to (somehow) input the calldata size
    const maxSubmissionPrice = (await this.getTxnSubmissionPrice(Zero))[0]
    return this.l1Bridge.depositAsERC20(
      erc20L1Address,
      amount,
      maxSubmissionPrice,
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
    // TODO: this will need to (somehow) input the calldata size
    const maxSubmissionPrice = (await this.getTxnSubmissionPrice(Zero))[0]

    return this.l1Bridge.depositAsERC777(
      erc20L1Address,
      amount,
      maxSubmissionPrice,
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

  public async getL2Transaction(l2TransactionHash: string) {
    const txReceipt = await this.l2Provider.getTransactionReceipt(
      l2TransactionHash
    )

    if (!txReceipt) throw new Error("Can't find L2 transaction receipt?")

    return txReceipt
  }

  public async getL1Transaction(l1TransactionHash: string) {
    const txReceipt = await this.l1Bridge.l1Provider.getTransactionReceipt(
      l1TransactionHash
    )

    if (!txReceipt) throw new Error("Can't find L2 transaction receipt?")

    return txReceipt
  }

  public async calculateL2TransactionHash(
    inboxSequenceNumber: BigNumber,
    l2ChainId?: BigNumber
  ) {
    if (!l2ChainId)
      l2ChainId = BigNumber.from((await this.l2Provider.getNetwork()).chainId)

    return ethers.utils.keccak256(
      ethers.utils.concat([
        ethers.utils.zeroPad(l2ChainId.toHexString(), 32),
        ethers.utils.zeroPad(inboxSequenceNumber.toHexString(), 32),
      ])
    )
  }

  public async getInboxSeqNumFromContractTransaction(
    l2Transaction: ethers.providers.TransactionReceipt
  ): Promise<Array<BigNumber> | undefined> {
    const Inbox = await this.l1Bridge.getInbox()
    const iface = Inbox.interface
    const messageDelivered = iface.getEvent('InboxMessageDelivered')
    const messageDeliveredFromOrigin = iface.getEvent(
      'InboxMessageDeliveredFromOrigin'
    )

    const eventTopics = {
      InboxMessageDelivered: iface.getEventTopic(messageDelivered),
      InboxMessageDeliveredFromOrigin: iface.getEventTopic(
        messageDeliveredFromOrigin
      ),
    }

    const logs = l2Transaction.logs.filter(
      log =>
        log.topics[0] === eventTopics.InboxMessageDelivered ||
        log.topics[0] === eventTopics.InboxMessageDeliveredFromOrigin
    )

    if (logs.length === 0) return undefined
    return logs.map(log => BigNumber.from(log.topics[1]))
  }

  public getBuddyDeployInL2Transaction(
    l2Transaction: ethers.providers.TransactionReceipt
  ) {
    const iface = new ethers.utils.Interface([
      `event Deployed(address indexed _sender, address indexed _contract, uint256 indexed withdrawalId, bool _success)`,
    ])
    const DeployedEvent = iface.getEvent('Deployed')
    const eventTopic = iface.getEventTopic(DeployedEvent)
    const logs = l2Transaction.logs.filter(log => log.topics[0] === eventTopic)
    return logs.map(
      log => (iface.parseLog(log).args as unknown) as BuddyDeployEventResult
    )
  }

  public async getWithdrawalsInL2Transaction(
    l2Transaction: ethers.providers.TransactionReceipt
  ): Promise<Array<L2ToL1EventResult>> {
    const iface = this.arbSys.interface
    const l2ToL1Event = iface.getEvent('L2ToL1Transaction')
    const eventTopic = iface.getEventTopic(l2ToL1Event)

    const logs = l2Transaction.logs.filter(log => log.topics[0] === eventTopic)

    return logs.map(
      log => (iface.parseLog(log).args as unknown) as L2ToL1EventResult
    )
  }

  public async triggerL2ToL1Transaction(
    batchNumber: BigNumber,
    indexInBatch: BigNumber,
    singleAttempt = false
  ) {
    console.log('going to get proof')
    let res: {
      proof: string[]
      path: BigNumber
      l2Sender: string
      l1Dest: string
      l2Block: BigNumber
      l1Block: BigNumber
      timestamp: BigNumber
      amount: BigNumber
      calldataForL1: string
    }

    if (singleAttempt) {
      const _res = await this.tryGetProofOnce(batchNumber, indexInBatch)
      if (_res === null) {
        console.warn('Proof not found')
        return
      }
      res = _res
    }
    res = await this.tryGetProof(batchNumber, indexInBatch)
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
    } = res

    console.log('got proof')

    const inbox = await this.l1Bridge.getInbox()
    const bridgeAddress = await inbox.bridge()
    const bridge = await Bridge__factory.connect(
      bridgeAddress,
      this.l1Bridge.l1Provider
    )

    const activeOutbox = await bridge.allowedOutboxList(0)
    try {
      // index 1 should not exist
      await bridge.allowedOutboxList(1)
      console.error('There is more than 1 outbox registered with the bridge?!')
    } catch (e) {
      // this should fail!
      console.log('All is good')
    }

    const outboxExecuteTransactionReceipt = await this.tryOutboxExecute(
      activeOutbox,
      batchNumber,
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

  public tryOutboxExecute = async (
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
    calldataForL1: string,
    retryDelay = 500
  ): Promise<ContractReceipt> => {
    const outbox = Outbox__factory.connect(
      activeOutboxAddress,
      this.l1Bridge.l1Signer
    )
    await this.waitUntilOutboxEntryCreated(batchNumber, activeOutboxAddress)
    try {
      // TODO: wait until assertion is confirmed before execute
      // We can predict and print number of missing blocks
      // if not challenged
      const outboxExecute = await outbox.executeTransaction(
        batchNumber,
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
        retryDelay
      )
    }
  }

  public tryGetProofOnce = async (
    batchNumber: BigNumber,
    indexInBatch: BigNumber,
    retryDelay = 500
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
  } | null> => {
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
        console.log('Withdrawal detected, but batch not created yet.')
      } else {
        console.log("Withdrawal proof didn't work. Not sure why")
        console.log(e)
      }
    }
    return null
  }

  public tryGetProof = async (
    batchNumber: BigNumber,
    indexInBatch: BigNumber,
    retryDelay = 500
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

  private wait = (ms: number) => new Promise(res => setTimeout(res, ms))

  public waitUntilOutboxEntryCreated = async (
    batchNumber: BigNumber,
    activeOutboxAddress: string,
    retryDelay = 500
  ): Promise<string> => {
    try {
      // if outbox entry not created yet, this reads from array out of bounds
      const expectedEntry = await this.getOutboxEntry(
        batchNumber,
        activeOutboxAddress
      )
      console.log('Found entry index!')
      return expectedEntry
    } catch (e) {
      console.log("can't find entry, lets wait a bit?")
      if (e.message === 'invalid opcode: opcode 0xfe not defined') {
        console.log('Array out of bounds, wait until the entry is posted')
      } else {
        console.log(e)
        console.log(e.message)
      }
      await this.wait(retryDelay)
      console.log('Starting new attempt')
      return this.waitUntilOutboxEntryCreated(
        batchNumber,
        activeOutboxAddress,
        retryDelay
      )
    }
  }

  private getOutboxEntry = async (
    batchNumber: BigNumber,
    outboxAddress: string
  ): Promise<string> => {
    const iface = new ethers.utils.Interface([
      'function outboxes(uint256) public view returns (address)',
    ])
    const outbox = new ethers.Contract(outboxAddress, iface).connect(
      this.l1Bridge.l1Provider
    )
    return outbox.outboxes(batchNumber)
  }

  public async getTokenWithdrawEventData(destinationAddress: string) {
    const iface = this.arbTokenBridge.interface
    const tokenWithdrawEvent = iface.getEvent('WithdrawToken')
    const tokenWithdrawTopic = iface.getEventTopic(tokenWithdrawEvent)

    const topics = [
      tokenWithdrawTopic,
      null,
      null,
      ethers.utils.hexZeroPad(destinationAddress, 32),
    ]

    const logs = await this.l2Provider.getLogs({
      address: this.arbTokenBridge.address,
      // @ts-ignore
      topics,
      fromBlock: 0,
      toBlock: 'latest',
    })

    return logs.map(
      log => (iface.parseLog(log).args as unknown) as WithdrawTokenEventResult
    )
  }

  public async getL2ToL1EventData(destinationAddress: string) {
    const iface = this.arbSys.interface
    const l2ToL1TransactionEvent = iface.getEvent('L2ToL1Transaction')
    const l2ToL1TransactionTopic = iface.getEventTopic(l2ToL1TransactionEvent)

    const topics = [
      l2ToL1TransactionTopic,
      ethers.utils.hexZeroPad(destinationAddress, 32),
    ]

    const logs = await this.l2Provider.getLogs({
      address: ARB_SYS_ADDRESS,
      topics,
      fromBlock: 0,
      toBlock: 'latest',
    })

    return logs.map(
      log => (iface.parseLog(log).args as unknown) as L2ToL1EventResult
    )
  }
}
