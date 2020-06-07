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
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/* eslint-env node */
'use strict'

import { EVMCode, EVMResult, TxMessage, processLog } from './message'
import { ArbClient, AVMProof } from './client'
import { AggregatorClient } from './aggregator'
import * as ArbValue from './value'
import { ArbWallet } from './wallet'
import { Contract } from './contract'
import * as Hashing from './hashing'

import * as ethers from 'ethers'

const promisePoller = require('promise-poller').default

import { ArbRollupFactory } from './abi/ArbRollupFactory'
import { ArbRollup } from './abi/ArbRollup'

import { GlobalInboxFactory } from './abi/GlobalInboxFactory'
import { GlobalInbox } from './abi/GlobalInbox'

import { ArbSysFactory } from './abi/ArbSysFactory'
import { ArbSys } from './abi/ArbSys'

import { ArbInfoFactory } from './abi/ArbInfoFactory'

// EthBridge event names
const EB_EVENT_VMC = 'VMCreated'
const EB_EVENT_CDA = 'RollupAsserted'
const TransactionMessageDelivered = 'TransactionMessageDelivered'
const EthDepositMessageDelivered = 'EthDepositMessageDelivered'
const ERC20DepositMessageDelivered = 'ERC20DepositMessageDelivered'
const ERC721DepositMessageDelivered = 'ERC721DepositMessageDelivered'

const ARB_SYS_ADDRESS = '0x0000000000000000000000000000000000000064'
const ARB_INFO_ADDRESS = '0x0000000000000000000000000000000000000065'

interface PossibleMessageResult {
  val: ArbValue.Value
  evmVal: EVMResult
  proof?: AVMProof
}

interface MessageResult {
  evmVal: EVMResult
  txHash: string
  validNodeHash: string
  onChainTxHash: string
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
  private vmIdCache?: string

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
      const vmID = await this.client.getVmID()
      const arbRollup = ArbRollupFactory.connect(vmID, this.ethProvider)
      this.arbRollupCache = arbRollup
      return arbRollup
    }
    return this.arbRollupCache
  }

  public async chainAddress(): Promise<string> {
    return this.client.getVmID()
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

  // public async getValidatorAddresses(): Promise<string[]> {
  //     if (!this.validatorAddressesCache) {
  //         const ArbRollup = await this.arbChannelConn();
  //         const validators = await this.client.getValidatorList();
  //         const isValidators = await arbChannel.isValidatorList(validators);
  //         if (!isValidators) {
  //             throw new Error('Incorrect validator list');
  //         }

  //         // Cache the set of lowercase validator addresses (without "0x")
  //         this.validatorAddressesCache = validators.map((addr: string) => addr.toLowerCase().slice(2)).sort();
  //         return this.validatorAddressesCache;
  //     }
  //     return this.validatorAddressesCache;
  // }

  // public async verifyUnanimousSignatures(
  //     assertionHash: ethers.utils.Arrayish,
  //     validatorSigs: string[],
  // ): Promise<void> {
  //     const validatorAddresses = await this.getValidatorAddresses();
  //     if (validatorAddresses.length !== validatorSigs.length) {
  //         throw Error('Expected: ' + validatorAddresses.length + ' signatures.\nReceived: ' + validatorSigs.length);
  //     }

  //     const addresses = validatorSigs
  //         .map(sig =>
  //             ethers.utils
  //                 .verifyMessage(ethers.utils.arrayify(assertionHash), sig)
  //                 .toLowerCase()
  //                 .slice(2),
  //         )
  //         .sort();

  //     for (let i = 0; i < validatorAddresses.length; i++) {
  //         if (validatorAddresses[i] !== addresses[i]) {
  //             throw Error('Invalid signature');
  //         }
  //     }
  // }

  public async getVmID(): Promise<string> {
    if (!this.vmIdCache) {
      const vmId = await this.client.getVmID()
      // Guard against race condition
      if (!this.vmIdCache) {
        this.vmIdCache = vmId
      }
      return vmId
    }
    return this.vmIdCache
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
        if (log.name == TransactionMessageDelivered) {
          const vmId = await this.getVmID()
          return Hashing.calculateTransactionHash(
            vmId,
            log.values.to,
            log.values.from,
            log.values.seqNumber,
            log.values.value,
            log.values.data
          )
        } else if (
          log.name == EthDepositMessageDelivered ||
          log.name == ERC20DepositMessageDelivered ||
          log.name == ERC721DepositMessageDelivered
        ) {
          return ethers.utils.hexZeroPad(
            log.values.messageNum.toHexString(),
            32
          )
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

  private async getRawMessageResult(
    txHash: string
  ): Promise<PossibleMessageResult | null> {
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

    const result = await this.client.getMessageResult(arbTxHash)

    if (!result) {
      return null
    }
    const { val, proof } = result

    const evmVal = processLog(val as ArbValue.TupleValue)

    const txHashCheck = evmVal.message.txHash

    // Check txHashCheck matches txHash
    if (txHash !== txHashCheck) {
      throw Error(
        'txHash did not match its queried transaction ' +
          arbTxHash +
          ' ' +
          txHashCheck
      )
    }

    return {
      val,
      evmVal,
      proof,
    }
  }

  public async getMessageResult(txHash: string): Promise<MessageResult | null> {
    const result = await this.getRawMessageResult(txHash)
    if (!result) {
      return null
    }
    const { val, evmVal, proof } = result

    // If we weren't passed a proof, there's nothing left to do.
    // Higher level code will have to choose whether to accept the
    // result.
    if (proof === undefined) {
      return null
    }

    // Step 1: prove that val is in logPostHash
    if (!this.processLogsProof(val, proof)) {
      throw Error('Failed to prove val is in logPostHash')
    }

    // Step 2: prove that logPostHash is in assertion and assertion is valid
    const validNodeHash = await this.processConfirmedDisputableAssertion(proof)

    return {
      evmVal,
      txHash: evmVal.message.txHash,
      validNodeHash,
      onChainTxHash: proof.onChainTxHash,
    }
  }

  // This should return a Promise (and may throw errors)
  // method is the method name (e.g. getBalance) and params is an
  // object with normalized values passed in, depending on the method
  public async perform(method: string, params: any): Promise<any> {
    // console.log('perform', method, params);
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
        let status = 0
        let logs: ethers.providers.Log[] = []
        if (
          result.evmVal.returnType === EVMCode.Return ||
          result.evmVal.returnType === EVMCode.Stop
        ) {
          status = 1
          logs = result.evmVal.logs
        }

        const txReceipt: ethers.providers.TransactionReceipt = {
          blockHash: result.txHash,
          blockNumber: result.evmVal.message.blockNumber.toNumber(),
          confirmations: 1000,
          cumulativeGasUsed: ethers.utils.bigNumberify(1),
          from: result.evmVal.message.sender,
          gasUsed: ethers.utils.bigNumberify(1),
          logs,
          status,
          to: result.evmVal.message.message.getDest(),
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
          const txMessage = result.evmVal.message.message as TxMessage
          const tx = {
            blockHash: result.txHash,
            blockNumber: result.evmVal.message.blockNumber.toNumber(),
            confirmations: 1000,
            data: ethers.utils.hexlify(txMessage.data),
            from: result.evmVal.message.sender,
            gasLimit: ethers.utils.bigNumberify(1),
            gasPrice: ethers.utils.bigNumberify(1),
            hash: result.txHash,
            nonce: 0,
            to: txMessage.getDest(),
            value: txMessage.amount,
            chainId: 123456789,
          } as ethers.providers.TransactionResponse
          return this.ethProvider._wrapTransaction(tx)
        }
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
    }
    const forwardResponse = this.ethProvider.perform(method, params)
    // console.log('Forwarding query to provider', method, forwardResponse);
    return forwardResponse
  }

  public async call(
    transaction: ethers.providers.TransactionRequest,
    blockTag?: ethers.providers.BlockTag | Promise<ethers.providers.BlockTag>
  ): Promise<string> {
    if (!transaction.to) {
      throw Error('Cannot create call without a destination')
    }
    const to = await transaction.to
    let from = await transaction.from
    const rawData = await transaction.data
    if (!from) {
      from = '0x1000000000000000000000000000000000000000'
    }
    let data = '0x'
    if (rawData) {
      data = ethers.utils.hexlify(rawData)
    }
    const resultData = await this.client.call(to, from, data)
    return ethers.utils.hexlify(resultData)
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

  // partialHash: hexString
  // logPostHash: hexString
  // validatorSigs: []hexString
  // Throws error if assertionHash is not signed by all validators
  // private async processUnanimousAssertion(
  //     partialHash: string,
  //     logPostHash: string,
  //     validatorSigs: string[],
  // ): Promise<void> {
  //     const vmId = await this.getVmID();
  //     const assertionHash = ethers.utils.solidityKeccak256(
  //         ['address', 'bytes32', 'bytes32'],
  //         [vmId, partialHash, logPostHash],
  //     );
  //     await this.verifyUnanimousSignatures(assertionHash, validatorSigs);
  // }

  // logPostHash: hexString
  // onChainTxHash: hexString
  // Returns the valid node hash if assertionHash is logged by the onChainTxHash
  private async processConfirmedDisputableAssertion(
    proof: AVMProof
  ): Promise<string> {
    const receipt = await this.ethProvider.waitForTransaction(
      proof.onChainTxHash
    )
    if (!receipt.logs) {
      throw Error('RollupAsserted tx had no logs')
    }
    const arbRollup = await this.arbRollupConn()
    const events = receipt.logs.map(l => arbRollup.interface.parseLog(l))
    // DisputableAssertion Event
    const eventIndex = events.findIndex(event => event.name === EB_EVENT_CDA)
    if (eventIndex == -1) {
      throw Error(
        'RollupAsserted ' + proof.onChainTxHash + ' not found on chain'
      )
    }

    const rawLog = receipt.logs[eventIndex]
    const cda = events[eventIndex]
    const vmId = await this.getVmID()
    // Check correct VM
    if (rawLog.address.toLowerCase() !== vmId.toLowerCase()) {
      throw Error(
        'RollupAsserted Event is from a different address: ' +
          rawLog.address +
          '\nExpected address: ' +
          vmId
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

    // DisputableAssertion is correct
    // TODO: must wait for finality (past the re-org period)

    return cda.values.fields[7]
  }
}
