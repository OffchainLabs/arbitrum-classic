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
/* eslint-env browser */
'use strict'

import * as ArbValue from './value'
import * as ethers from 'ethers'

export enum EVMCode {
  Revert = 0,
  Invalid = 1,
  Return = 2,
  Stop = 3,
  BadSequenceCode = 4,
}

function intValueToAddress(value: ArbValue.IntValue): string {
  return ethers.utils.getAddress(
    ethers.utils.hexZeroPad(
      (value as ArbValue.IntValue).bignum.toHexString(),
      20
    )
  )
}

function logValToLog(
  val: ArbValue.Value,
  index: number,
  orig: EthBridgeMessage
): ethers.providers.Log {
  const value = val as ArbValue.TupleValue
  return {
    blockNumber: orig.blockNumber.toNumber(),
    blockHash: orig.txHash,
    transactionIndex: 0,
    removed: false,
    transactionLogIndex: index,
    address: ethers.utils.hexlify((value.get(0) as ArbValue.IntValue).bignum),
    data: ethers.utils.hexlify(
      ArbValue.bytestackToBytes(value.get(1) as ArbValue.TupleValue)
    ),
    topics: value.contents
      .slice(2)
      .map(rawTopic =>
        ethers.utils.hexZeroPad(
          ethers.utils.hexlify((rawTopic as ArbValue.IntValue).bignum),
          32
        )
      ),
    transactionHash: orig.txHash,
    logIndex: index,
  }
}

function stackValueToList(value: ArbValue.TupleValue): ArbValue.Value[] {
  const values = []
  while (value.contents.length !== 0) {
    values.push(value.get(1))
    value = value.get(0) as ArbValue.TupleValue
  }
  return values
}

export class TxCall {
  public to: string
  public data: Uint8Array

  constructor(value: ArbValue.TupleValue) {
    this.to = intValueToAddress(value.get(0) as ArbValue.IntValue)
    this.data = ArbValue.bytestackToBytes(value.get(1) as ArbValue.TupleValue)
  }

  getDest(): string {
    return this.to
  }
}

export class TxMessage {
  public to: string
  public sequenceNum: ethers.utils.BigNumber
  public amount: ethers.utils.BigNumber
  public data: Uint8Array

  constructor(value: ArbValue.TupleValue) {
    this.to = intValueToAddress(value.get(0) as ArbValue.IntValue)
    this.sequenceNum = (value.get(1) as ArbValue.IntValue).bignum
    this.amount = (value.get(2) as ArbValue.IntValue).bignum
    this.data = ArbValue.bytestackToBytes(value.get(3) as ArbValue.TupleValue)
  }

  getDest(): string {
    return this.to
  }
}

export class ContractTxMessage {
  public to: string
  public amount: ethers.utils.BigNumber
  public data: Uint8Array

  constructor(value: ArbValue.TupleValue) {
    this.to = intValueToAddress(value.get(0) as ArbValue.IntValue)
    this.amount = (value.get(1) as ArbValue.IntValue).bignum
    this.data = ArbValue.bytestackToBytes(value.get(2) as ArbValue.TupleValue)
  }

  getDest(): string {
    return this.to
  }
}

class EthTransferMessage {
  public dest: string
  public amount: ethers.utils.BigNumber

  constructor(value: ArbValue.TupleValue) {
    this.dest = intValueToAddress(value.get(0) as ArbValue.IntValue)
    this.amount = (value.get(1) as ArbValue.IntValue).bignum
  }

  getDest(): string {
    return this.dest
  }
}
class TokenTransferMessage {
  public tokenAddress: string
  public dest: string
  public amount: ethers.utils.BigNumber

  constructor(value: ArbValue.TupleValue) {
    this.tokenAddress = intValueToAddress(value.get(0) as ArbValue.IntValue)
    this.dest = intValueToAddress(value.get(1) as ArbValue.IntValue)
    this.amount = (value.get(2) as ArbValue.IntValue).bignum
  }

  getDest(): string {
    return this.dest
  }
}

function unmarshalExecutedMessage(
  typecode: number,
  message: ArbValue.TupleValue
): ArbMessage {
  switch (typecode) {
    case 0:
      return new TxMessage(message)
    case 1:
      return new EthTransferMessage(message)
    case 2:
      return new TokenTransferMessage(message)
    case 3:
      return new TokenTransferMessage(message)
    case 4:
      return new ContractTxMessage(message)
    case 5:
      return new TxCall(message)
    default:
      throw 'Invalid arb message type'
  }
}

export type ArbMessage =
  | TxCall
  | TxMessage
  | EthTransferMessage
  | TokenTransferMessage
  | TxCall

class EthBridgeMessage {
  public blockNumber: ethers.utils.BigNumber
  public timestamp: ethers.utils.BigNumber
  public txHash: string
  public sender: string
  public message: ArbMessage
  public calldataHash: string

  constructor(value: ArbValue.TupleValue) {
    this.blockNumber = (value.get(0) as ArbValue.IntValue).bignum
    this.timestamp = (value.get(1) as ArbValue.IntValue).bignum
    this.txHash = ethers.utils.hexZeroPad(
      (value.get(2) as ArbValue.IntValue).bignum.toHexString(),
      32
    )
    const restVal = value.get(3) as ArbValue.TupleValue
    const typecode = (restVal.get(0) as ArbValue.IntValue).bignum.toNumber()
    this.sender = intValueToAddress(restVal.get(1) as ArbValue.IntValue)

    this.message = unmarshalExecutedMessage(
      typecode,
      restVal.get(2) as ArbValue.TupleValue
    )
    this.calldataHash = restVal.hash()
  }
}

export type EVMResult =
  | EVMReturn
  | EVMRevert
  | EVMStop
  | EVMBadSequenceCode
  | EVMInvalid

export class EVMReturn {
  public message: EthBridgeMessage
  public data: Uint8Array
  public logs: ethers.providers.Log[]
  public returnType: EVMCode.Return

  constructor(value: ArbValue.TupleValue) {
    this.message = new EthBridgeMessage(value.get(0) as ArbValue.TupleValue)
    this.data = ArbValue.bytestackToBytes(value.get(2) as ArbValue.TupleValue)
    this.logs = stackValueToList(value.get(1) as ArbValue.TupleValue).map(
      (val, index) => {
        return logValToLog(val, index, this.message)
      }
    )
    this.returnType = EVMCode.Return
  }
}

export class EVMRevert {
  public message: EthBridgeMessage
  public data: Uint8Array
  public returnType: EVMCode.Revert

  constructor(value: ArbValue.TupleValue) {
    this.message = new EthBridgeMessage(value.get(0) as ArbValue.TupleValue)
    this.data = ArbValue.bytestackToBytes(value.get(2) as ArbValue.TupleValue)
    this.returnType = EVMCode.Revert
  }
}

export class EVMStop {
  public message: EthBridgeMessage
  public logs: ethers.providers.Log[]
  public returnType: EVMCode.Stop

  constructor(value: ArbValue.TupleValue) {
    this.message = new EthBridgeMessage(value.get(0) as ArbValue.TupleValue)
    this.logs = stackValueToList(value.get(1) as ArbValue.TupleValue).map(
      (val, index) => {
        return logValToLog(val, index, this.message)
      }
    )
    this.returnType = EVMCode.Stop
  }
}

export class EVMBadSequenceCode {
  public message: EthBridgeMessage
  public returnType: EVMCode.BadSequenceCode

  constructor(value: ArbValue.TupleValue) {
    this.message = new EthBridgeMessage(value.get(0) as ArbValue.TupleValue)
    this.returnType = EVMCode.BadSequenceCode
  }
}

export class EVMInvalid {
  public message: EthBridgeMessage
  public returnType: EVMCode.Invalid

  constructor(value: ArbValue.TupleValue) {
    this.message = new EthBridgeMessage(value.get(0) as ArbValue.TupleValue)
    this.returnType = EVMCode.Invalid
  }
}

export function processLog(value: ArbValue.TupleValue): EVMResult {
  const returnCode = value.get(3) as ArbValue.IntValue
  switch (returnCode.bignum.toNumber()) {
    case EVMCode.Return:
      return new EVMReturn(value)
    case EVMCode.Revert:
      return new EVMRevert(value)
    case EVMCode.Stop:
      return new EVMStop(value)
    case EVMCode.BadSequenceCode:
      return new EVMBadSequenceCode(value)
    case EVMCode.Invalid:
      return new EVMInvalid(value)
    default:
      throw Error('processLogs Invalid EVM return code')
  }
}
