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

function hex32(val: ethers.utils.BigNumber): Uint8Array {
  return ethers.utils.padZeros(ethers.utils.arrayify(val), 32)
}

function encodedAddress(addr: ethers.utils.Arrayish): Uint8Array {
  return ethers.utils.padZeros(ethers.utils.arrayify(addr), 32)
}

function intValueToAddress(value: ArbValue.IntValue): string {
  return ethers.utils.getAddress(
    ethers.utils.hexZeroPad(
      (value as ArbValue.IntValue).bignum.toHexString(),
      20
    )
  )
}

export enum L2MessageCode {
  Transaction = 0,
  ContractTransaction = 1,
  Call = 2,
  TransactionBatch = 3,
}

export class L2Transaction {
  public maxGas: ethers.utils.BigNumber
  public gasPriceBid: ethers.utils.BigNumber
  public sequenceNum: ethers.utils.BigNumber
  public destAddress: string
  public payment: ethers.utils.BigNumber
  public calldata: Uint8Array
  public kind: L2MessageCode.Transaction

  constructor(
    maxGas: ethers.utils.BigNumberish,
    gasPriceBid: ethers.utils.BigNumberish,
    sequenceNum: ethers.utils.BigNumberish,
    destAddress: ethers.utils.Arrayish | undefined,
    payment: ethers.utils.BigNumberish | undefined,
    calldata: ethers.utils.Arrayish | undefined
  ) {
    if (!destAddress) {
      destAddress = '0x'
    }
    if (!calldata) {
      calldata = '0x'
    }
    if (!payment) {
      payment = 0
    }
    this.maxGas = ethers.utils.bigNumberify(maxGas)
    this.gasPriceBid = ethers.utils.bigNumberify(gasPriceBid)
    this.sequenceNum = ethers.utils.bigNumberify(sequenceNum)
    this.destAddress = ethers.utils.hexZeroPad(
      ethers.utils.hexlify(destAddress),
      20
    )
    this.payment = ethers.utils.bigNumberify(payment)
    this.calldata = ethers.utils.arrayify(calldata)
    this.kind = L2MessageCode.Transaction
  }

  static fromData(data: ethers.utils.Arrayish): L2Transaction {
    const bytes = ethers.utils.arrayify(data)
    return new L2Transaction(
      bytes.slice(0, 32),
      bytes.slice(32, 64),
      bytes.slice(64, 96),
      bytes.slice(108, 128),
      bytes.slice(128, 160),
      bytes.slice(160)
    )
  }

  asData(): Uint8Array {
    return ethers.utils.concat([
      hex32(this.maxGas),
      hex32(this.gasPriceBid),
      hex32(this.sequenceNum),
      encodedAddress(this.destAddress),
      hex32(this.payment),
      ethers.utils.arrayify(this.calldata),
    ])
  }

  messageID(sender: string): string {
    return ethers.utils.solidityKeccak256(
      ['address', 'bytes'],
      [sender, this.asData()]
    )
  }

  batchHash(chain: string): string {
    return new ArbValue.TupleValue([
      new ArbValue.IntValue(chain),
      ArbValue.hexToBytestack(this.asData()),
    ]).hash()
  }

  batchData(signature: ethers.utils.Arrayish): string {
    const length = ethers.utils.bigNumberify(this.calldata.length).toHexString()
    const paddedLength = ethers.utils.hexZeroPad(length, 8)
    const data = ethers.utils.concat([paddedLength, this.asData(), signature])
    return ethers.utils.hexlify(data)
  }
}

export class L2Call {
  public maxGas: ethers.utils.BigNumber
  public gasPriceBid: ethers.utils.BigNumber
  public destAddress: string
  public calldata: string
  public kind: L2MessageCode.Call

  constructor(
    maxGas: ethers.utils.BigNumberish | undefined,
    gasPriceBid: ethers.utils.BigNumberish | undefined,
    destAddress: ethers.utils.Arrayish | undefined,
    calldata: ethers.utils.Arrayish | undefined
  ) {
    if (!maxGas) {
      maxGas = 0
    }
    if (!gasPriceBid) {
      gasPriceBid = 0
    }
    if (!destAddress) {
      destAddress = ethers.utils.hexZeroPad('0x', 20)
    }
    if (!calldata) {
      calldata = '0x'
    }
    this.maxGas = ethers.utils.bigNumberify(maxGas)
    this.gasPriceBid = ethers.utils.bigNumberify(gasPriceBid)
    this.destAddress = ethers.utils.hexlify(destAddress)
    this.calldata = ethers.utils.hexlify(calldata)
    this.kind = L2MessageCode.Call
  }

  static fromData(data: ethers.utils.Arrayish): L2Call {
    const bytes = ethers.utils.arrayify(data)
    return new L2Call(
      bytes.slice(0, 32),
      bytes.slice(32, 64),
      bytes.slice(64, 96),
      bytes.slice(96)
    )
  }

  asData(): Uint8Array {
    return ethers.utils.concat([
      hex32(this.maxGas),
      hex32(this.gasPriceBid),
      encodedAddress(this.destAddress),
      this.calldata,
    ])
  }
}

export type L2SubMessage = L2Transaction | L2Call

export enum MessageCode {
  Eth = 0,
  ERC20 = 1,
  ERC721 = 2,
  L2 = 3,
}

export class EthMessage {
  public kind: MessageCode.Eth
  public dest: ethers.utils.Arrayish
  public value: ethers.utils.BigNumber

  constructor(dest: string, value: ethers.utils.BigNumberish) {
    this.kind = MessageCode.Eth
    this.dest = dest
    this.value = ethers.utils.bigNumberify(value)
  }

  static fromData(data: ethers.utils.Arrayish): EthMessage {
    const bytes = ethers.utils.arrayify(data)
    return new EthMessage(
      ethers.utils.hexlify(bytes.slice(12, 32)),
      bytes.slice(32, 64)
    )
  }

  asData(): Uint8Array {
    return ethers.utils.concat([encodedAddress(this.dest), hex32(this.value)])
  }
}

export class ERC20Message {
  public kind: MessageCode.ERC20
  public tokenAddress: string
  public dest: string
  public value: ethers.utils.BigNumber

  constructor(
    tokenAddress: ethers.utils.Arrayish,
    dest: ethers.utils.Arrayish,
    value: ethers.utils.BigNumberish
  ) {
    this.kind = MessageCode.ERC20
    this.tokenAddress = ethers.utils.hexlify(tokenAddress)
    this.dest = ethers.utils.hexlify(dest)
    this.value = ethers.utils.bigNumberify(value)
  }

  static fromData(data: ethers.utils.Arrayish): ERC20Message {
    const bytes = ethers.utils.arrayify(data)
    return new ERC20Message(
      bytes.slice(12, 32),
      bytes.slice(44, 64),
      bytes.slice(64, 96)
    )
  }

  asData(): Uint8Array {
    return ethers.utils.concat([
      encodedAddress(this.tokenAddress),
      encodedAddress(this.dest),
      hex32(this.value),
    ])
  }
}

export class ERC721Message {
  public kind: MessageCode.ERC721
  public tokenAddress: string
  public dest: string
  public id: ethers.utils.BigNumber

  constructor(
    tokenAddress: ethers.utils.Arrayish,
    dest: ethers.utils.Arrayish,
    id: ethers.utils.BigNumberish
  ) {
    this.kind = MessageCode.ERC721
    this.tokenAddress = ethers.utils.hexlify(tokenAddress)
    this.dest = ethers.utils.hexlify(dest)
    this.id = ethers.utils.bigNumberify(id)
  }

  static fromData(data: ethers.utils.Arrayish): ERC721Message {
    const bytes = ethers.utils.arrayify(data)
    return new ERC721Message(
      bytes.slice(12, 32),
      bytes.slice(44, 64),
      bytes.slice(64, 96)
    )
  }

  asData(): Uint8Array {
    return ethers.utils.concat([
      encodedAddress(this.tokenAddress),
      encodedAddress(this.dest),
      hex32(this.id),
    ])
  }
}

function l2SubMessageFromData(data: ethers.utils.Arrayish): L2SubMessage {
  const bytes = ethers.utils.arrayify(data)
  const kind = bytes[0]
  switch (kind) {
    case L2MessageCode.Transaction:
      return L2Transaction.fromData(bytes.slice(1))
    case L2MessageCode.Call:
      return L2Call.fromData(bytes.slice(1))
    default:
      throw Error('invalid L2 message type ' + kind)
  }
}

export class L2Message {
  public kind: MessageCode.L2

  constructor(public message: L2SubMessage) {
    this.kind = MessageCode.L2
  }

  static fromData(data: ethers.utils.Arrayish): L2Message {
    return new L2Message(l2SubMessageFromData(data))
  }

  asData(): Uint8Array {
    return ethers.utils.concat([[this.message.kind], this.message.asData()])
  }
}

export type Message = EthMessage | ERC20Message | ERC721Message | L2Message

function newMessageFromData(
  kind: MessageCode,
  data: ethers.utils.Arrayish
): Message {
  switch (kind) {
    case MessageCode.Eth:
      return EthMessage.fromData(data)
    case MessageCode.ERC20:
      return ERC20Message.fromData(data)
    case MessageCode.ERC721:
      return ERC721Message.fromData(data)
    case MessageCode.L2:
      return L2Message.fromData(data)
    default:
      throw 'Invalid arb message type'
  }
}

export class IncomingMessage {
  public msg: Message
  public blockNumber: ethers.utils.BigNumber
  public timestamp: ethers.utils.BigNumber
  public sender: string
  public inboxSeqNum: ethers.utils.BigNumber
  constructor(
    msg: Message,
    blockNumber: ethers.utils.BigNumberish,
    timestamp: ethers.utils.BigNumberish,
    sender: string,
    inboxSeqNum: ethers.utils.BigNumberish
  ) {
    this.msg = msg
    this.blockNumber = ethers.utils.bigNumberify(blockNumber)
    this.timestamp = ethers.utils.bigNumberify(timestamp)
    this.sender = sender
    this.inboxSeqNum = ethers.utils.bigNumberify(inboxSeqNum)
  }

  static fromValue(val: ArbValue.Value): IncomingMessage {
    const tup = val as ArbValue.TupleValue
    const kind = (tup.get(0) as ArbValue.IntValue).bignum.toNumber()
    const data = ArbValue.bytestackToBytes(tup.get(5) as ArbValue.TupleValue)

    return new IncomingMessage(
      newMessageFromData(kind, data),
      (tup.get(1) as ArbValue.IntValue).bignum,
      (tup.get(2) as ArbValue.IntValue).bignum,
      intValueToAddress(tup.get(3) as ArbValue.IntValue),
      (tup.get(4) as ArbValue.IntValue).bignum
    )
  }

  asValue(): ArbValue.Value {
    return new ArbValue.TupleValue([
      new ArbValue.IntValue(this.msg.kind),
      new ArbValue.IntValue(this.blockNumber),
      new ArbValue.IntValue(this.timestamp),
      new ArbValue.IntValue(this.sender),
      new ArbValue.IntValue(this.inboxSeqNum),
      ArbValue.hexToBytestack(this.msg.asData()),
    ])
  }

  commitmentHash(): string {
    return ethers.utils.solidityKeccak256(
      ['uint8', 'address', 'uint256', 'uint256', 'uint256', 'bytes32'],
      [
        this.msg.kind,
        this.sender,
        this.blockNumber,
        this.timestamp,
        this.inboxSeqNum,
        ethers.utils.keccak256(this.msg.asData()),
      ]
    )
  }

  messageID(): string {
    if (this.msg.kind == MessageCode.L2) {
      const l2message = this.msg
      if (l2message.message.kind == L2MessageCode.Transaction) {
        return l2message.message.messageID(this.sender)
      }
    }
    return this.inboxSeqNum.toHexString()
  }
}

export class OutgoingMessage {
  constructor(public msg: Message, public sender: string) {}

  asValue(): ArbValue.Value {
    return new ArbValue.TupleValue([
      new ArbValue.IntValue(this.msg.kind),
      new ArbValue.IntValue(this.sender),
      ArbValue.hexToBytestack(this.msg.asData()),
    ])
  }
}

export enum ResultCode {
  Return = 0,
  Revert = 1,
  Congestion = 2,
  InsufficientGasFunds = 3,
  InsufficientTxFunds = 4,
  BadSequenceCode = 5,
  InvalidMessageFormatCode = 6,
  UnknownErrorCode = 255,
}

export class Log {
  constructor(
    public address: string,
    public topics: string[],
    public data: string
  ) {}

  static fromValue(val: ArbValue.Value): Log {
    console.log('raw log', val)
    const tup = val as ArbValue.TupleValue
    const topics = tup.contents
      .slice(2)
      .map(rawTopic =>
        ethers.utils.hexZeroPad(
          ethers.utils.hexlify((rawTopic as ArbValue.IntValue).bignum),
          32
        )
      )
    return new Log(
      intValueToAddress(tup.get(0) as ArbValue.IntValue),
      topics,
      ethers.utils.hexlify(
        ArbValue.bytestackToBytes(tup.get(1) as ArbValue.TupleValue)
      )
    )
  }

  asValue(): ArbValue.Value {
    const values: ArbValue.Value[] = []
    values.push(new ArbValue.IntValue(this.address))
    values.push(ArbValue.hexToBytestack(this.data))
    for (const topic of this.topics) {
      values.push(new ArbValue.IntValue(topic))
    }
    return new ArbValue.TupleValue(values)
  }
}

export class Result {
  public incoming: IncomingMessage
  public resultCode: ResultCode
  public returnData: Uint8Array
  public logs: Log[]
  public gasUsed: ethers.utils.BigNumber
  public gasPrice: ethers.utils.BigNumber

  constructor(
    incoming: IncomingMessage,
    resultCode: ResultCode,
    returnData: Uint8Array,
    logs: Log[],
    gasUsed: ethers.utils.BigNumberish,
    gasPrice: ethers.utils.BigNumberish
  ) {
    this.incoming = incoming
    this.resultCode = resultCode
    this.returnData = returnData
    this.logs = logs
    this.gasUsed = ethers.utils.bigNumberify(gasUsed)
    this.gasPrice = ethers.utils.bigNumberify(gasPrice)
  }

  static fromValue(val: ArbValue.Value): Result {
    const tup = val as ArbValue.TupleValue
    const incoming = IncomingMessage.fromValue(tup.get(0))
    const logs = stackValueToList(tup.get(3) as ArbValue.TupleValue).map(val =>
      Log.fromValue(val)
    )
    return new Result(
      incoming,
      (tup.get(1) as ArbValue.IntValue).bignum.toNumber(),
      ArbValue.bytestackToBytes(tup.get(2) as ArbValue.TupleValue),
      logs,
      (tup.get(4) as ArbValue.IntValue).bignum.toNumber(),
      (tup.get(5) as ArbValue.IntValue).bignum.toNumber()
    )
  }

  asValue(): ArbValue.Value {
    return new ArbValue.TupleValue([
      this.incoming.asValue(),
      new ArbValue.IntValue(this.resultCode),
      ArbValue.hexToBytestack(this.returnData),
      new ArbValue.IntValue(this.gasUsed),
      new ArbValue.IntValue(this.gasPrice),
    ])
  }
}

function stackValueToList(value: ArbValue.TupleValue): ArbValue.Value[] {
  const values = []
  while (value.contents.length !== 0) {
    values.push(value.get(0))
    value = value.get(1) as ArbValue.TupleValue
  }
  return values
}
