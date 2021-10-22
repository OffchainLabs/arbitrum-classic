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

/* eslint-env node, mocha */
import { ethers } from 'hardhat'
import { expect } from 'chai'
import { TestBytesParser } from '../build/types/TestBytesParser'

const sample = 'abcdefghijklmnopqrstuvwxyz'.repeat(10)

describe('Bridge peripherals layer 1', () => {
  let testBytesParser: TestBytesParser

  before(async function () {
    const TestBytesParser = await ethers.getContractFactory('TestBytesParser')
    testBytesParser = await TestBytesParser.deploy()
  })

  it('should fail on empty data', async function () {
    const [success, res] = await testBytesParser.bytesToString('0x')
    expect(success).to.be.false
    expect(res).to.equal('')
  })

  it('should handle bytes32 string name', async function () {
    for (let i = 0; i < 32; i++) {
      const testString = sample.substr(0, i)
      const encoded = ethers.utils.formatBytes32String(testString)
      const [success, res] = await testBytesParser.bytesToString(encoded)
      expect(success).to.be.true
      expect(res).to.equal(testString)
    }
  })

  it('should fail on non-terminated bytes32', async function () {
    const testString = sample.substr(0, 10)
    let encoded = ethers.utils.formatBytes32String(testString)
    encoded = encoded.substr(0, encoded.length - 1) + '1'
    const [success, res] = await testBytesParser.bytesToString(encoded)
    expect(success).to.be.false
    expect(res).to.equal('')
  })

  it('should handle strings', async function () {
    const strings = ['', 'a', sample]
    for (const testString of strings) {
      const encoded = ethers.utils.defaultAbiCoder.encode(
        ['string'],
        [testString]
      )
      const [success, res] = await testBytesParser.bytesToString(encoded)
      expect(success).to.be.true
      expect(res).to.equal(testString)
    }
  })

  it('should revert on invalid inputs', async function () {
    const data = [
      '0x01',
      '0x0000000000000000000000000000000000000020',
      ethers.utils.randomBytes(64),
      ethers.utils.randomBytes(128),
    ]
    for (const testString of data) {
      await expect(testBytesParser.bytesToString(testString)).to.be.reverted
    }
  })
})
