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

/* eslint-env node, mocha */

import { ethers } from '@nomiclabs/buidler'
import { assert } from 'chai'
import { ValueTester } from '../build/types/ValueTester'
import { ArbValue } from 'arb-provider-ethers'

import testCases from './test_cases.json'

const testVal =
  '0x5345325345325345325345325345325345325345325345325345325345325435'

describe('Value', () => {
  let valueTester: ValueTester

  before(async () => {
    const ValueTester = await ethers.getContractFactory('ValueTester')
    valueTester = (await ValueTester.deploy()) as ValueTester
    await valueTester.deployed()
  })

  it('should initialize', async () => {
    const val = new ArbValue.IntValue(100)
    const res = await valueTester.deserializeHash(ArbValue.marshal(val), 0)
    assert.isTrue(res['0'], "value didn't deserialize correctly")
    assert.equal(val.hash(), res['2'], 'value hashes incorrectly')
  })

  for (let i = 0; i < testCases.length; i++) {
    it(testCases[i].name, async () => {
      const expectedHash = testCases[i].hash
      const res = await valueTester.deserializeHash(
        '0x' + testCases[i].proof_value,
        0
      )
      assert.isTrue(res['0'], "value didn't deserialize correctly")
      assert.equal('0x' + expectedHash, res['2'], 'value hashes incorrectly')
    })
  }

  it('should parse erc20 message', async () => {
    const val = new ArbValue.TupleValue([
      new ArbValue.IntValue(2),
      new ArbValue.IntValue(
        '1454660323771124265538360532739934987166685588469'
      ),
      new ArbValue.TupleValue([
        new ArbValue.IntValue(
          '641988807973089174688456409219473706566398216120'
        ),
        new ArbValue.IntValue(
          '1454660323771124265538360532739934987166685588469'
        ),
        new ArbValue.IntValue(1543),
      ]),
    ])

    const valData = ArbValue.marshal(val)
    const res = await valueTester.deserializeMessageData(valData, 0)
    const offset = res['1'].toNumber()
    assert.isTrue(res['0'], "value didn't deserialize correctly")
    assert.equal(res['2'].toNumber(), 2, 'Incorrect message type')
    assert.equal(
      res['3'],
      '0xFeCd3992654bFC565c3aFc6C4d7b14dCe603EbF5',
      'Incorrect sender'
    )

    const res2 = await valueTester.getERCTokenMsgData(valData, offset)
    assert.isTrue(res2['0'], "value didn't deserialize correctly")
    assert.equal(
      res2['2'],
      '0x7073c616a8A3F277Ea4511fCe9EBB2656a1b87B8',
      'Incorrect token contract'
    )
    assert.equal(
      res2['3'],
      '0xFeCd3992654bFC565c3aFc6C4d7b14dCe603EbF5',
      'Incorrect dest'
    )
    assert.equal(res2['4'].toNumber(), 1543, 'Incorrect value')
  })
  it('should parse eth message', async () => {
    const val = new ArbValue.TupleValue([
      new ArbValue.IntValue(2),
      new ArbValue.IntValue(
        '1454660323771124265538360532739934987166685588469'
      ),
      new ArbValue.TupleValue([
        new ArbValue.IntValue(
          '1454660323771124265538360532739934987166685588469'
        ),
        new ArbValue.IntValue(1543),
      ]),
    ])

    const valData = ArbValue.marshal(val)
    const res = await valueTester.deserializeMessageData(valData, 0)
    const offset = res['1'].toNumber()
    assert.isTrue(res['0'], "value didn't deserialize correctly")
    assert.equal(res['2'].toNumber(), 2, 'Incorrect message type')
    assert.equal(
      res['3'],
      '0xFeCd3992654bFC565c3aFc6C4d7b14dCe603EbF5',
      'Incorrect sender'
    )

    const res2 = await valueTester.getEthMsgData(valData, offset)
    assert.isTrue(res2['0'], "value didn't deserialize correctly")
    assert.equal(
      res2['2'],
      '0xFeCd3992654bFC565c3aFc6C4d7b14dCe603EbF5',
      'Incorrect dest'
    )
    assert.equal(res2['3'].toNumber(), 1543, 'Incorrect value')
  })

  it('should properly calculate bytestack hash 32 bytes', async () => {
    const ethVal = await valueTester.bytesToBytestackHash(testVal)
    const jsVal = ArbValue.hexToBytestack(testVal).hash()
    assert.equal(ethVal, jsVal)
  })

  it('should properly calculate bytestack hash 64 bytes', async () => {
    const ethVal = await valueTester.bytesToBytestackHash(
      testVal + testVal.slice(2)
    )
    const jsVal = ArbValue.hexToBytestack(testVal + testVal.slice(2)).hash()
    assert.equal(ethVal, jsVal)
  })

  it('should properly calculate bytestack hash 16 bytes', async () => {
    const ethVal = await valueTester.bytesToBytestackHash(testVal.slice(0, 34))
    const jsVal = ArbValue.hexToBytestack(testVal.slice(0, 34)).hash()
    assert.equal(ethVal, jsVal)
  })

  it('should properly calculate bytestack hash 19 bytes', async () => {
    const ethVal = await valueTester.bytesToBytestackHash(testVal.slice(0, 40))
    const jsVal = ArbValue.hexToBytestack(testVal.slice(0, 40)).hash()
    assert.equal(ethVal, jsVal)
  })

  it('should properly convert bytestack to bytes', async () => {
    const bytestack = ArbValue.hexToBytestack(testVal.slice(0, 40))
    const bytestackData = ethers.utils.hexlify(ArbValue.marshal(bytestack))
    const ethVal = await valueTester.bytestackToBytes(bytestackData)
    assert.equal(ethVal, testVal.slice(0, 40))
  })
})
