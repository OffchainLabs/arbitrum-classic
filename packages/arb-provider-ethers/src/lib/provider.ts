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
import { ArbClient } from './client'
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

export interface AVMProof {
  logPreHash: string
  logValHashes: string[]
}

interface MessageResult {
  result: Result
  txIndex: number
  startLogIndex: number
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

  private arbRollupCache?: ArbRollup
  private globalInboxCache?: GlobalInbox
  private validatorAddressesCache?: string[]
  private chainAddressCache?: string

  constructor(
    validatorUrl: string,
    provider: ethers.providers.JsonRpcProvider,
    aggregatorUrl?: string
  ) {
    super(123456789)
    this.chainId = 123456789
    this.ethProvider = provider
    this.client = new ArbClient(validatorUrl)

    if (aggregatorUrl) {
      this.aggregator = new AggregatorClient(aggregatorUrl)
    }
  }

  public async arbRollupConn(): Promise<ArbRollup> {
    if (!this.arbRollupCache) {
      const vmID = await this.chainAddress()
      const arbRollup = ArbRollupFactory.connect(vmID, this.ethProvider)
      this.arbRollupCache = arbRollup
      return arbRollup
    }
    return this.arbRollupCache
  }

  public async globalInboxConn(): Promise<GlobalInbox> {
    if (!this.globalInboxCache) {
      const arbRollup = await this.arbRollupConn()
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

  public async chainAddress(): Promise<string> {
    if (!this.chainAddressCache) {
      const chainAddress = await this.client.getChainAddress()
      // Guard against race condition
      if (!this.chainAddressCache) {
        this.chainAddressCache = chainAddress
      }
      return chainAddress
    }
    return this.chainAddressCache
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

  public async getPaymentMessage(index: number): Promise<VerifyMessageResult> {
    const results = await this.client.getOutputMessage(index)
    return {
      value: results.outputMsg,
    }
  }

  public async getMessageResult(txHash: string): Promise<MessageResult> {
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
        throw Error("txHash wasn't an ethereum or arbitrum transaction")
      }
      arbTxHash = arbTxId
    } else {
      arbTxHash = txHash
    }

    const { log, txIndex, startLogIndex } = await this.client.getRequestResult(
      arbTxHash
    )

    const result = Result.fromValue(log)

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

    // Optionally check if the log was actually included in an assertion
    // const validateLogs = false
    // if (validateLogs) {
    //   const assertionTxHash = ''
    //   let proof: AVMProof
    //   await this.verifyDisputableAssertion(assertionTxHash, log, proof)
    // }

    return {
      result,
      txIndex,
      startLogIndex,
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
        const { result, txIndex, startLogIndex } = await this.getMessageResult(
          params.transactionHash
        )

        const currentBlockNum = await this.ethProvider.getBlockNumber()
        const messageBlockNum = result.incoming.blockNumber.toNumber()
        const confirmations = currentBlockNum - messageBlockNum + 1
        const block = await this.ethProvider.getBlock(messageBlockNum)

        const incoming = result.incoming
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
          contractAddress = ethers.utils.hexlify(result.returnData.slice(12))
        }

        let status = 0
        const logs: ethers.providers.Log[] = []
        if (result.resultCode === ResultCode.Return) {
          status = 1
          let logIndex = startLogIndex
          for (const log of result.logs) {
            logs.push({
              ...log,
              transactionIndex: txIndex,
              blockNumber: messageBlockNum,
              transactionHash: incoming.messageID(),
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
          cumulativeGasUsed: result.gasUsed,
          from: incoming.sender,
          gasUsed: result.gasUsed,
          logs,
          status,
          to: l2tx.destAddress,
          transactionHash: incoming.messageID(),
          transactionIndex: txIndex,
          byzantium: true,
        }
        return txReceipt
      }
      case 'getTransaction': {
        const getMessageRequest = async (): Promise<
          ethers.providers.TransactionResponse
        > => {
          const {
            result,
            txIndex,
            startLogIndex,
          } = await this.getMessageResult(params.transactionHash)
          const incoming = result.incoming
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
            from: incoming.sender,
            gasLimit: ethers.utils.bigNumberify(1),
            gasPrice: ethers.utils.bigNumberify(1),
            hash: incoming.messageID(),
            nonce: l2tx.sequenceNum.toNumber(),
            to: l2tx.destAddress,
            value: l2tx.payment,
            chainId: 123456789,
          }
          const response = this.ethProvider._wrapTransaction(tx)
          const blockNumber = incoming.blockNumber.toNumber()
          const block = await this.ethProvider.getBlock(blockNumber)
          return {
            ...response,
            blockHash: block.hash,
            blockNumber: blockNumber,
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
        return this.client.getBlockCount()
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

  // Returns the valid node hash if assertionHash is logged by the onChainTxHash
  private async verifyDisputableAssertion(
    assertionTxHash: string,
    value: ArbValue.Value,
    proof: AVMProof
  ): Promise<void> {
    const receipt = await this.ethProvider.waitForTransaction(assertionTxHash)
    if (!receipt.logs) {
      throw Error('RollupAsserted tx had no logs')
    }
    const arbRollup = await this.arbRollupConn()
    const events = receipt.logs.map(l => arbRollup.interface.parseLog(l))
    // DisputableAssertion Event
    const eventIndex = events.findIndex(event => event.name === EB_EVENT_CDA)
    if (eventIndex == -1) {
      throw Error('RollupAsserted ' + assertionTxHash + ' not found on chain')
    }

    const rawLog = receipt.logs[eventIndex]
    const cda = events[eventIndex]
    const vmId = await this.chainAddress()
    // Check correct VM
    if (rawLog.address.toLowerCase() !== vmId.toLowerCase()) {
      throw Error(
        'RollupAsserted Event is from a different address: ' +
          rawLog.address +
          '\nExpected address: ' +
          vmId
      )
    }

    const startHash = ethers.utils.solidityKeccak256(
      ['bytes32', 'bytes32'],
      [proof.logPreHash, value.hash()]
    )
    const logPostHash = proof.logValHashes.reduce(
      (acc, hash) =>
        ethers.utils.solidityKeccak256(['bytes32', 'bytes32'], [acc, hash]),
      startHash
    )

    // Check correct logs hash
    if (cda.values.fields[6] !== logPostHash) {
      throw Error(
        'RollupAsserted Event on-chain logPostHash is: ' +
          cda.values.fields[6] +
          '\nExpected: ' +
          logPostHash
      )
    }
  }
}
