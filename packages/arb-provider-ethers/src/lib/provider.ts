/*
 * Copyright 2019-2020, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on afn "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/* eslint-env node */
'use strict'

import {
  Result,
  ResultCode,
  Log,
  MessageCode,
  L2MessageCode,
  L2Call,
  L2Message,
} from './message'
import { ArbClient, AVMProof, NodeInfo } from './client'
import { AggregatorClient } from './aggregator'
import * as ArbValue from './value'
import { ArbWallet } from './wallet'

import * as ethers from 'ethers'

import promisePoller from 'promise-poller'

import { ArbRollupFactory } from './abi/ArbRollupFactory'
import { ArbRollup } from './abi/ArbRollup'

import { GlobalInboxFactory } from './abi/GlobalInboxFactory'
import { GlobalInbox } from './abi/GlobalInbox'

import { ArbSysFactory } from './abi/ArbSysFactory'
import { ArbSys } from './abi/ArbSys'

import { ArbInfoFactory } from './abi/ArbInfoFactory'

// EthBridge event names
const EB_EVENT_CDA = 'RollupAsserted'
const MessageDelivered = 'MessageDelivered'
const MessageDeliveredFromOrigin = 'MessageDeliveredFromOrigin'

const ARB_SYS_ADDRESS = '0x0000000000000000000000000000000000000064'
const ARB_INFO_ADDRESS = '0x0000000000000000000000000000000000000065'

interface PossibleMessageResult {
  val: ArbValue.Value
  result: Result
  nodeInfo?: NodeInfo
  proof?: AVMProof
}

interface MessageResult {
  result: Result
  txHash: string
  nodeInfo: NodeInfo
}

interface VerifyMessageResult {
  value: ArbValue.Value
}

interface Message {
  to: string
  sequenceNum: ethers.utils.BigNumberish
  value: ethers.utils.BigNumberish
  data: string
  signature: string
  pubkey: string
}

export class ArbProvider extends ethers.providers.BaseProvider {
  public chainId: number
  public ethProvider: ethers.providers.JsonRpcProvider
  public client: ArbClient
  public aggregator?: AggregatorClient
  public chainAddress: string

  private arbRollupCache?: ArbRollup
  private globalInboxCache?: GlobalInbox
  private validatorAddressesCache?: string[]
  private latestLocation?: NodeInfo

  constructor(
    chainAddress: string,
    validatorUrl: string,
    provider: ethers.providers.JsonRpcProvider,
    aggregatorUrl?: string
  ) {
    const chainId = ethers.utils
      .bigNumberify(ethers.utils.hexDataSlice(chainAddress, 12))
      .toNumber()
    super(chainId)
    this.chainAddress = chainAddress
    this.chainId = chainId
    this.ethProvider = provider
    this.client = new ArbClient(validatorUrl)

    if (aggregatorUrl) {
      this.aggregator = new AggregatorClient(aggregatorUrl)
    }
  }

  public arbRollupConn(): ArbRollup {
    if (!this.arbRollupCache) {
      const arbRollup = ArbRollupFactory.connect(
        this.chainAddress,
        this.ethProvider
      )
      this.arbRollupCache = arbRollup
      return arbRollup
    }
    return this.arbRollupCache
  }

  public async globalInboxConn(): Promise<GlobalInbox> {
    if (!this.globalInboxCache) {
      const arbRollup = this.arbRollupConn()
      const globalInboxAddress = await arbRollup.globalInbox()
      const globalInbox = GlobalInboxFactory.connect(
        globalInboxAddress,
        this.ethProvider
      )
      this.globalInboxCache = globalInbox
      return globalInbox
    }
    return this.globalInboxCache
  }

  public getArbSys(): ArbSys {
    return ArbSysFactory.connect(ARB_SYS_ADDRESS, this)
  }

  public getSigner(index: number): ArbWallet {
    return new ArbWallet(this.ethProvider.getSigner(index), this)
  }

  private async getArbTxId(
    ethReceipt: ethers.providers.TransactionReceipt
  ): Promise<string | null> {
    const globalInbox = await this.globalInboxConn()
    if (ethReceipt.logs) {
      const logs = ethReceipt.logs.map(log =>
        globalInbox.interface.parseLog(log)
      )
      for (const log of logs) {
        if (!log) {
          continue
        }
        if (log.name == MessageDelivered) {
          if (log.values.kind == MessageCode.L2) {
            const msg = L2Message.fromData(log.values.data)
            if (msg.message.kind == L2MessageCode.Transaction) {
              return msg.message.messageID(log.values.sender)
            }
          }
          return log.values.inboxSeqNum
        }
      }
    }
    return null
  }

  public async getPaymentMessage(
    assertedNodeHash: string,
    messageIndex: string
  ): Promise<VerifyMessageResult | null> {
    const results = await this.client.getOutputMessage(
      assertedNodeHash,
      messageIndex
    )

    if (results != null) {
      return {
        value: results.outputMsg,
      }
    }

    return null
  }

  public async getMessageResult(txHash: string): Promise<MessageResult | null> {
    // TODO: Make sure that there can be no collision between arbitrum transaction hashes and
    // Ethereum transaction hashes so that an attacker cannot fool the client into accepting a
    // false transction
    let arbTxHash: string
    const ethReceipt = await this.ethProvider.getTransactionReceipt(txHash)
    if (ethReceipt) {
      // If this receipt exists, it must've been an L1 hash
      const arbTxId = await this.getArbTxId(ethReceipt)
      if (!arbTxId) {
        // If the Ethereum transaction wasn't actually a message send, the input data was bad
        return null
      }
      arbTxHash = arbTxId
    } else {
      arbTxHash = txHash
    }

    const rawResult = await this.client.getMessageResult(arbTxHash)

    if (!rawResult) {
      return null
    }
    const { val, nodeInfo, proof } = rawResult

    const result = Result.fromValue(val)

    const txHashCheck = result.incoming.messageID()

    // Check txHashCheck matches txHash
    if (arbTxHash !== txHashCheck) {
      throw Error(
        'txHash did not match its queried transaction ' +
          arbTxHash +
          ' ' +
          txHashCheck
      )
    }

    if (proof === undefined || nodeInfo.l1TxHash === undefined) {
      return {
        result,
        txHash: arbTxHash,
        nodeInfo,
      }
    }

    // Step 1: prove that val is in logPostHash
    if (!this.processLogsProof(val, proof)) {
      throw Error('Failed to prove val is in logPostHash')
    }

    // Step 2: prove that logPostHash is in assertion and assertion is valid
    const { confirmations } = await this.processConfirmedDisputableAssertion(
      nodeInfo,
      proof
    )

    return {
      result,
      txHash: arbTxHash,
      nodeInfo: {
        ...nodeInfo,
        l1Confirmations: confirmations,
      },
    }
  }

  // This should return a Promise (and may throw errors)
  // method is the method name (e.g. getBalance) and params is an
  // object with normalized values passed in, depending on the method
  /* eslint-disable no-alert, @typescript-eslint/no-explicit-any */
  public async perform(method: string, params: any): Promise<any> {
    // console.log('perform', method, params)
    switch (method) {
      case 'getCode': {
        if (
          params.address == ARB_SYS_ADDRESS ||
          params.address == ARB_INFO_ADDRESS
        ) {
          return '0x100'
        }
        const arbInfo = ArbInfoFactory.connect(ARB_INFO_ADDRESS, this)
        return arbInfo.getCode(params.address)
      }
      case 'getTransactionCount': {
        const arbsys = this.getArbSys()
        const count = await arbsys.getTransactionCount(params.address)
        return count.toNumber()
      }
      case 'getTransactionReceipt': {
        const result = await this.getMessageResult(params.transactionHash)
        if (!result) {
          return null
        }

        const currentBlockNum = await this.ethProvider.getBlockNumber()
        const messageBlockNum = result.result.incoming.blockNumber.toNumber()
        const confirmations = currentBlockNum - messageBlockNum + 1
        const block = await this.ethProvider.getBlock(messageBlockNum)

        const incoming = result.result.incoming
        if (
          incoming.msg.kind != MessageCode.L2 ||
          incoming.msg.message.kind != L2MessageCode.Transaction
        ) {
          throw Error(
            'Can only call getTransaction on an L2 Transaction message'
          )
        }
        const l2tx = incoming.msg.message

        let contractAddress = undefined
        if (ethers.utils.hexStripZeros(l2tx.destAddress) == '0x0') {
          contractAddress = ethers.utils.hexlify(
            result.result.returnData.slice(12)
          )
        }

        let status = 0
        const logs: ethers.providers.Log[] = []
        if (result.result.resultCode === ResultCode.Return) {
          status = 1
          let logIndex = 0
          for (const log of result.result.logs) {
            logs.push({
              ...log,
              transactionLogIndex: 0,
              transactionIndex: 0,
              blockNumber: messageBlockNum,
              transactionHash: result.txHash,
              logIndex,
              blockHash: block.hash,
            })
            logIndex++
          }
        }

        const txReceipt: ethers.providers.TransactionReceipt = {
          blockHash: block.hash,
          blockNumber: messageBlockNum,
          contractAddress: contractAddress,
          confirmations: confirmations,
          cumulativeGasUsed: result.result.gasUsed,
          from: result.result.incoming.sender,
          gasUsed: result.result.gasUsed,
          logs,
          status,
          to: l2tx.destAddress,
          transactionHash: result.txHash,
          transactionIndex: 0,
          byzantium: true,
        }
        return txReceipt
      }
      case 'getTransaction': {
        const getMessageRequest = async (): Promise<ethers.providers.TransactionResponse | null> => {
          const result = await this.getMessageResult(params.transactionHash)
          if (!result) {
            return null
          }
          const incoming = result.result.incoming
          if (
            incoming.msg.kind != MessageCode.L2 ||
            incoming.msg.message.kind != L2MessageCode.Transaction
          ) {
            throw Error(
              'Can only call getTransaction on an L2 Transaction message'
            )
          }
          const l2tx = incoming.msg.message
          const tx: ethers.utils.Transaction = {
            data: ethers.utils.hexlify(l2tx.calldata),
            from: result.result.incoming.sender,
            gasLimit: ethers.utils.bigNumberify(1),
            gasPrice: ethers.utils.bigNumberify(1),
            hash: result.txHash,
            nonce: l2tx.sequenceNum.toNumber(),
            to: l2tx.destAddress,
            value: l2tx.payment,
            chainId: 123456789,
          }
          const response = this.ethProvider._wrapTransaction(tx)
          if (result.nodeInfo === undefined) {
            return response
          }
          return {
            ...response,
            blockHash: result.nodeInfo.nodeHash,
            blockNumber: result.nodeInfo.nodeHeight,
            confirmations: 1000,
          }
        }
        /* eslint-disable no-alert, @typescript-eslint/no-explicit-any */
        return promisePoller({
          interval: 100,
          shouldContinue: (reason?: any, value?: any): boolean => {
            if (reason) {
              return true
            } else if (value) {
              return false
            } else {
              return true
            }
          },
          taskFn: getMessageRequest,
        })
      }
      case 'getLogs': {
        return this.client.findLogs(params.filter)
      }
      case 'getBalance': {
        const arbInfo = ArbInfoFactory.connect(ARB_INFO_ADDRESS, this)
        return arbInfo.getBalance(params.address)
      }
      case 'getBlockNumber': {
        const location = await this.client.getLatestPendingNodeLocation()
        if (location) {
          if (
            this.latestLocation &&
            (this.latestLocation.nodeHeight !== location.nodeHeight ||
              this.latestLocation.nodeHash !== location.nodeHash)
          ) {
            this.resetEventsBlock(location.nodeHeight)
          }
          this.latestLocation = location
        }
        return this.ethProvider.getBlockNumber()
      }
      case 'estimateGas': {
        const tx: ethers.providers.TransactionRequest = params.transaction
        const result = await this.callImpl(tx)
        return result.gasUsed
      }
    }
    // console.log('Forwarding query to provider', method, params);
    return await this.ethProvider.perform(method, params)
  }

  private async callImpl(
    transaction: ethers.providers.TransactionRequest,
    blockTag?: ethers.providers.BlockTag | Promise<ethers.providers.BlockTag>
  ): Promise<Result> {
    const from = await transaction.from
    const tx = new L2Call(
      await transaction.gasLimit,
      await transaction.gasPrice,
      await transaction.to,
      await transaction.data
    )

    const callLatest = (): Promise<ArbValue.Value> => {
      return this.client.pendingCall(tx, from)
    }

    let resultVal: ArbValue.Value
    const tag = await blockTag
    if (tag) {
      if (tag == 'pending') {
        resultVal = await this.client.pendingCall(tx, from)
      } else if (tag == 'latest') {
        resultVal = await callLatest()
      } else {
        throw Error('Invalid block tag')
      }
    } else {
      resultVal = await callLatest()
    }
    return Result.fromValue(resultVal)
  }

  public async call(
    transaction: ethers.providers.TransactionRequest,
    blockTag?: ethers.providers.BlockTag | Promise<ethers.providers.BlockTag>
  ): Promise<string> {
    const result = await this.callImpl(transaction, blockTag)
    if (result.resultCode != ResultCode.Return) {
      throw new Error('Call was reverted')
    }
    return ethers.utils.hexlify(result.returnData)
  }

  // value: *Value
  // Returns true if the hash of value is in logPostHash and false otherwise
  private processLogsProof(value: ArbValue.Value, proof: AVMProof): boolean {
    const startHash = ethers.utils.solidityKeccak256(
      ['bytes32', 'bytes32'],
      [proof.logPreHash, value.hash()]
    )
    const checkHash = proof.logValHashes.reduce(
      (acc, hash) =>
        ethers.utils.solidityKeccak256(['bytes32', 'bytes32'], [acc, hash]),
      startHash
    )

    return proof.logPostHash === checkHash
  }

  // logPostHash: hexString
  // onChainTxHash: hexString
  // Returns the valid node hash if assertionHash is logged by the onChainTxHash
  private async processConfirmedDisputableAssertion(
    nodeInfo: NodeInfo,
    proof: AVMProof
  ): Promise<ethers.providers.TransactionReceipt> {
    if (nodeInfo.l1TxHash === undefined) {
      throw Error("node doesn't exist on chain")
    }
    const receipt = await this.ethProvider.waitForTransaction(nodeInfo.l1TxHash)
    if (!receipt.logs) {
      throw Error('RollupAsserted tx had no logs')
    }
    const arbRollup = this.arbRollupConn()
    const events = receipt.logs.map(l => arbRollup.interface.parseLog(l))
    // DisputableAssertion Event
    const eventIndex = events.findIndex(event => event.name === EB_EVENT_CDA)
    if (eventIndex == -1) {
      throw Error('RollupAsserted ' + nodeInfo.l1TxHash + ' not found on chain')
    }

    const rawLog = receipt.logs[eventIndex]
    const cda = events[eventIndex]
    // Check correct VM
    if (rawLog.address.toLowerCase() !== this.chainAddress.toLowerCase()) {
      throw Error(
        'RollupAsserted Event is from a different address: ' +
          rawLog.address +
          '\nExpected address: ' +
          this.chainAddress
      )
    }

    // Check correct logs hash
    if (cda.values.fields[6] !== proof.logPostHash) {
      throw Error(
        'RollupAsserted Event on-chain logPostHash is: ' +
          cda.values.fields[6] +
          '\nExpected: ' +
          proof.logPostHash
      )
    }

    if (cda.values.fields[7] != nodeInfo.nodeHash) {
      throw Error(
        'RollupAsserted Event on-chain nodeHash is: ' +
          cda.values.fields[7] +
          '\nExpected: ' +
          nodeInfo.nodeHash
      )
    }

    return receipt
  }
}
