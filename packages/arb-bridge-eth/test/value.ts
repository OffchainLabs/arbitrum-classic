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
import { assert, expect } from 'chai'
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
    assert.equal(val.hash(), res['1'], 'value hashes incorrectly')
  })

  for (let i = 0; i < testCases.length; i++) {
    it(testCases[i].name, async () => {
      const expectedHash = testCases[i].hash
      const res = await valueTester.deserializeHash(
        '0x' + testCases[i].proof_value,
        0
      )
      assert.equal('0x' + expectedHash, res['1'], 'value hashes incorrectly')
    })
  }

  const cases = [
    testVal.slice(0, 34),
    testVal.slice(0, 40),
    testVal,
    testVal + testVal.slice(2),
  ]

  cases.forEach(data => {
    const dataLength = ethers.utils.hexDataLength(data)
    it(
      'should properly calculate bytestack hash of length ' + dataLength,
      async () => {
        const ethVal = await valueTester.bytesToBytestackHash(
          data,
          0,
          dataLength
        )
        const jsVal = ArbValue.hexToBytestack(data).hash()
        assert.equal(ethVal, jsVal)
      }
    )
  })

  it('should properly convert bytestack to bytes', async () => {
    const bytestack = ArbValue.hexToBytestack(testVal.slice(0, 40))
    const marshalled = ArbValue.marshal(bytestack)
    const bytestackData = ethers.utils.hexlify(marshalled)
    const { 0: valid, 1: offset, 2: data } = await valueTester.bytestackToBytes(
      bytestackData,
      0
    )
    assert.isTrue(valid)
    expect(offset).to.equal(marshalled.length)
    assert.equal(data, testVal.slice(0, 40))
  })
})
