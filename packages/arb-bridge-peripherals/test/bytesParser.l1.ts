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

  it('should handle proper uint8s', async function () {
    const sampleSize = 20
    const [min, max] = [ 0, 256 ]
    const randVals = Array(sampleSize + 1)
      .fill(0)
      .map(_ => Math.floor(Math.random() * (max - min) + min))

    const inputs = [...randVals, 0, 1, 254, 255]
    
    const encoded = inputs.map(input =>
      ethers.utils.defaultAbiCoder.encode(['uint8'], [input])
    )

    for(const i in encoded) {
      const input = encoded[i]
      const expectedOutput = inputs[i]
      const [success, actualOutput] = await testBytesParser.bytesToUint8(input)
      expect(success).to.be.true
      expect(actualOutput).to.equal(expectedOutput)
    }

  })

  it('should handle overflow uint8', async function () {
    const input = ethers.utils.defaultAbiCoder.encode(['uint256'], [256])
    const [success, actualOutput] = await testBytesParser.bytesToUint8(input)
    expect(success).to.be.false
    expect(actualOutput).to.equal(0)

    const input2 = ethers.utils.defaultAbiCoder.encode(['uint256'], [257])
    const [success2, actualOutput2] = await testBytesParser.bytesToUint8(input2)
    expect(success2).to.be.false
    expect(actualOutput2).to.equal(0)

    const input3 = ethers.utils.defaultAbiCoder.encode(['uint256'], [30000])
    const [success3, actualOutput3] = await testBytesParser.bytesToUint8(input3)
    expect(success3).to.be.false
    expect(actualOutput3).to.equal(0)

    const input4 = ethers.utils.defaultAbiCoder.encode(['uint256'], [40000002])
    const [success4, actualOutput4] = await testBytesParser.bytesToUint8(input4)
    expect(success4).to.be.false
    expect(actualOutput4).to.equal(0)
  })
})
