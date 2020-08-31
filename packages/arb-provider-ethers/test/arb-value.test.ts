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
/* eslint-env node, jest */
'use strict'

import { expect } from 'chai'
import * as ethers from 'ethers'

const utils = ethers.utils
import * as arb from '../src/lib/value'
import testCases from './test_cases.json'

// Helper shortcuts
const bn = utils.bigNumberify
const ZEROS_16B = '00000000000000000000000000000000'
const ZEROS_32B = ZEROS_16B + ZEROS_16B
// const ONES_16B = 'ffffffffffffffffffffffffffffffff'
// const ONES_32B = ONES_16B + ONES_16B
const EMPTY_TUPLE_HASH =
  '0x5dfd744753c13a6de9cf0e9e3d0164cdf0b1bbdb5efb3660a684d46477176058'

describe('Constructors', function () {
  const nullHash = '0x' + ZEROS_32B

  it('BasicOp', function () {
    const bop = new arb.BasicOp(10)
    expect(bop.opcode).to.equal(10)
  })

  it('ImmOp', function () {
    const iop = new arb.ImmOp(0x19, new arb.IntValue(utils.bigNumberify(9)))
    expect(iop.opcode).to.equal(0x19)
    expect((iop.value as arb.IntValue).bignum.toNumber()).to.equal(9)
  })

  it('IntValue', function () {
    const iv = new arb.IntValue(utils.bigNumberify(0))
    expect(iv.bignum.toNumber()).to.equal(0)
    expect(iv.typeCode()).to.equal(0)
  })

  it('CodePointValue', function () {
    const cpv = new arb.CodePointValue(new arb.BasicOp(0x60), nullHash)
    expect(cpv.op.opcode).to.equal(0x60)
    expect(cpv.nextHash).to.equal(nullHash)

    // Test BasicOp hash value
    const bopv = new arb.CodePointValue(new arb.BasicOp(0x60), EMPTY_TUPLE_HASH)
    const preCalc =
      '0xe20558cb3c2dbc788aee4091e5596aa3a86530d82bce979e21365703da522301'
    expect(bopv.hash()).to.equal(preCalc)

    // Test ImmOp hash value
    const immop = new arb.ImmOp(0x60, new arb.IntValue(bn(0)))
    const immv = new arb.CodePointValue(immop, EMPTY_TUPLE_HASH)
    const preCalc2 =
      '0x17ae7e9c5b69861db0759631c54ee3a23ecd04bd57d3d3e2b5579c05c5bd0e8b'
    expect(immv.hash()).to.equal(preCalc2)
  })
})

describe('TupleValue', function () {
  it('Empty', function () {
    const emptyTuple = new arb.TupleValue([])
    expect(emptyTuple.contents).to.deep.equal([])
    expect(emptyTuple.typeCode()).to.equal(3)
    expect(emptyTuple.hash()).to.equal(EMPTY_TUPLE_HASH)
  })

  it('Two Tuple', function () {
    const twoTuple = new arb.TupleValue([
      new arb.IntValue(utils.bigNumberify(0)),
      new arb.IntValue(utils.bigNumberify(100)),
    ])
    expect((twoTuple.get(0) as arb.IntValue).bignum.isZero()).to.equal(true)
    expect(
      (twoTuple.get(1) as arb.IntValue).bignum.eq(utils.bigNumberify(100))
    ).to.equal(true)
    expect(twoTuple.typeCode()).to.equal(3 + 2)
  })

  it('Largest Tuple', function () {
    const mtsv = new arb.TupleValue(
      Array(arb.MAX_TUPLE_SIZE).fill(new arb.TupleValue([]))
    )
    expect(mtsv.contents.length).to.equal(arb.MAX_TUPLE_SIZE)
    for (let i = 0; i < arb.MAX_TUPLE_SIZE; i++) {
      expect((mtsv.get(i) as arb.TupleValue).contents.length).to.equal(0)
    }
    expect(mtsv.typeCode()).to.equal(3 + arb.MAX_TUPLE_SIZE)
    // Pre calculated hash
    const p =
      '0x50d4ffda29a6d6324f06a725d7df50f897b34d245b2cff697281c437e454c777'
    expect(mtsv.hash()).to.equal(p)
  })

  it('Greater than MAX_TUPLE_SIZE', function () {
    expect(() => new arb.TupleValue(Array(arb.MAX_TUPLE_SIZE + 1))).to.throw(
      'Error TupleValue: illegal size ' + (arb.MAX_TUPLE_SIZE + 1)
    )
  })

  it('get and set', function () {
    const emptyTuple = new arb.TupleValue(
      Array(arb.MAX_TUPLE_SIZE).fill(new arb.TupleValue([]))
    )
    let t = emptyTuple
    // set
    for (let i = 0; i < arb.MAX_TUPLE_SIZE; i++) {
      t = t.set(i, new arb.IntValue(utils.bigNumberify(i)))
      expect((t.contents[i] as arb.IntValue).bignum.toNumber()).to.equal(i)
    }

    expect(() => t.set(arb.MAX_TUPLE_SIZE, new arb.TupleValue([]))).to.throw(
      'Error TupleValue set: index out of bounds ' + arb.MAX_TUPLE_SIZE
    )
    expect(() => t.set(-1, new arb.TupleValue([]))).to.throw(
      'Error TupleValue set: index out of bounds ' + -1
    )

    // get
    for (let i = 0; i < arb.MAX_TUPLE_SIZE; i++) {
      expect((t.get(i) as arb.IntValue).bignum.toNumber()).to.equal(i)
    }

    expect(() => t.get(arb.MAX_TUPLE_SIZE)).to.throw(
      'Error TupleValue get: index out of bounds ' + arb.MAX_TUPLE_SIZE
    )
    expect(() => t.get(-1)).to.throw(
      'Error TupleValue get: index out of bounds ' + -1
    )
  })
})

// Marshaled sizes as hexstrings
const M_INT_VALUE_SIZE = 1 + 32
const M_CODE_POINT_SIZE = 1 + 1 + 1 + 0 + 32 // Without val
const M_TUPLE_SIZE = 1 + 0 // Without other vals

describe('Marshaling', function () {
  it('marshal and unmarshal IntValue', function () {
    for (const i of [0, 1, 100, '0x9271342394932492394']) {
      const iv = new arb.IntValue(bn(i))
      const marshaledBytes = arb.marshal(iv)
      expect(marshaledBytes.length).to.equal(M_INT_VALUE_SIZE)
      const unmarshaledValue = arb.unmarshal(marshaledBytes)
      expect((unmarshaledValue as arb.IntValue).bignum.eq(bn(i))).to.equal(true)
    }

    // Test that negative IntValues throw on marshal
    expect(() => arb.marshal(new arb.IntValue(bn(-1)))).to.throw(
      'Error marshaling IntValue: negative values not supported'
    )

    // Test without "0x"
    const iv = new arb.IntValue(bn(99))
    const marshaledBytes = arb.marshal(iv)
    expect(marshaledBytes.length).to.equal(M_INT_VALUE_SIZE)
    const unmarshaledValue = arb.unmarshal(marshaledBytes)
    expect((unmarshaledValue as arb.IntValue).bignum.eq(bn(99))).to.equal(true)
  })

  it('marshal and unmarshal CodePointValue', function () {
    const op = new arb.BasicOp(10)
    const nextHash = '0x' + ZEROS_32B
    const basicTCV = new arb.CodePointValue(op, nextHash)
    const marshaledBytes = arb.marshal(basicTCV)
    expect(marshaledBytes.length).to.equal(M_CODE_POINT_SIZE)
    const revValue = arb.unmarshal(marshaledBytes) as arb.CodePointValue
    expect(revValue.op.opcode).to.equal(op.opcode)
    expect(revValue.nextHash).to.equal(nextHash)
    expect(revValue.toString()).to.equal(basicTCV.toString())

    const iv = new arb.IntValue(bn(60))
    expect(arb.marshal(iv).length).to.equal(M_INT_VALUE_SIZE)
    const immTCV = new arb.CodePointValue(new arb.ImmOp(0x19, iv), nextHash)
    const mb = arb.marshal(immTCV)
    expect(mb.length).to.equal(M_CODE_POINT_SIZE + M_INT_VALUE_SIZE)
    const revImmValue = arb.unmarshal(mb) as arb.CodePointValue
    expect(revImmValue.op.opcode).to.equal(0x19)
    expect(
      ((revImmValue.op as arb.ImmOp).value as arb.IntValue).bignum.toNumber()
    ).to.equal(60)
    expect(revImmValue.nextHash).to.equal(nextHash)
    expect(revImmValue.toString()).to.equal(immTCV.toString())
  })

  it('marshal and unmarshal TupleValue', function () {
    // Empty Tuple
    const etv = new arb.TupleValue([])
    const etvm = arb.marshal(etv)
    expect(etvm.length).to.equal(M_TUPLE_SIZE)
    const etvRev = arb.unmarshal(etvm)
    expect(etvRev.toString()).to.equal(etv.toString())

    // Full Tuple of Empty Tuple"s
    const ftv = new arb.TupleValue(Array(8).fill(new arb.TupleValue([])))
    const ftvm = arb.marshal(ftv)
    expect(ftvm.length).to.equal(M_TUPLE_SIZE + M_TUPLE_SIZE * 8)
    const ftvRev = arb.unmarshal(ftvm)
    expect(ftvRev.toString()).to.equal(ftv.toString())

    // Full Tuple of IntValue"s
    const fitv = new arb.TupleValue(Array(8).fill(new arb.IntValue(bn(0))))
    const fitvm = arb.marshal(fitv)
    expect(fitvm.length).to.equal(M_TUPLE_SIZE + M_INT_VALUE_SIZE * 8)
    const fitvRev = arb.unmarshal(fitvm) as arb.TupleValue
    expect(fitvRev.toString()).to.equal(fitv.toString())
    expect((fitvRev.get(0) as arb.IntValue).bignum.toNumber()).to.equal(0)
    expect((fitvRev.get(7) as arb.IntValue).bignum.toNumber()).to.equal(0)
    expect(() => fitvRev.get(8)).to.throw(
      'Error TupleValue get: index out of bounds 8'
    )
  })

  it('illegal inputs', function () {
    // Illegal Value
    expect(() => arb.unmarshal('0x99')).to.throw(
      'Error unmarshaling value no such TYPE: 99'
    )

    const [tyCodePoint, erroneousOpTy] = ['0x01', 'FF']
    expect(() => arb.unmarshal(tyCodePoint + erroneousOpTy)).to.throw(
      'Error unmarshalOp no such immCount: 255'
    )

    expect(() => arb.unmarshal('0x01')).to.throw(
      'Error extracting bytes: Uint8Array is too short'
    )
  })
})

describe('Integration', function () {
  it('hexToBytestack and bytestackToBytes', function () {
    // Create test value
    const messageBytes =
      '0x0b030b03030303030303005ce0c8f1e004fe36aa260ecd02c68ca0c6dea5a4acdfe0b8b10d7b526360046b0b0303030303030300781371cb80a394c637cebf3e3d48a268a44ad21cd68239afb3c3a37196d582c10b030303030303030032edc9a1000000000000000000000000000000000000000000000000000000000303030090e130e5da79003b67479a3ed2caf5585e93ae6771de6cdec6d7641bd2e60180'
    const bytestack = arb.hexToBytestack(messageBytes)
    const messageBytes2 = arb.bytestackToBytes(bytestack)
    expect(ethers.utils.hexlify(messageBytes2)).to.equal(messageBytes)
  })
})

describe('test_cases.json', function () {
  for (let i = 0; i < testCases.length; i++) {
    it(testCases[i].name, function () {
      const expectedHash = testCases[i].hash
      const value = arb.unmarshal('0x' + testCases[i].value)
      const hash = value.hash().slice(2)
      if (hash !== expectedHash) {
        console.log(value.toString())
      }
      expect(hash).to.equal(expectedHash)
    })
  }
})
