/*
 * Copyright 2020, Offchain Labs, Inc.
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
import { HashingTester } from '../build/types/HashingTester'

const { utils } = ethers

let hashTester: HashingTester

type bytes32 = string

function keccak1(a: bytes32): bytes32 {
  return utils.solidityKeccak256(['bytes32'], [a])
}

function keccak2(a: bytes32, b: bytes32): bytes32 {
  return utils.solidityKeccak256(['bytes32', 'bytes32'], [a, b])
}

function elem(a: number): bytes32 {
  return '0x' + a.toString(16).padStart(64, '0')
}

function emptyBuffer(num: number) {
  const arr: bytes32[] = []
  for (let i = 0; i < num; i++) {
    arr.push(elem(0))
  }
  return arr
}

function bufferToBytes(lst: bytes32[]) {
  return '0x' + lst.map(a => a.substr(2)).join('')
}

function merkleHash(arr: bytes32[], offset: number, sz: number): bytes32 {
  if (sz === 1) {
    return keccak1(arr[offset])
  } else {
    const h1 = merkleHash(arr, offset, sz / 2)
    const h2 = merkleHash(arr, offset + sz / 2, sz / 2)
    return keccak2(h1, h2)
  }
}

describe('Hashing', function () {
  before(async () => {
    const Hashing = await ethers.getContractFactory('HashingTester')
    hashTester = (await Hashing.deploy()) as HashingTester
    await hashTester.deployed()
  })

  describe('#merkleRoot', function () {
    it('should work with empty buffer', async () => {
      expect(await hashTester.testMerkleHash('0x')).to.equal(
        '0x290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563'
      )
    })
    it('should work with small buffer', async () => {
      expect(await hashTester.testMerkleHash('0x00000000')).to.equal(
        '0x290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563'
      )
    })
    it('should work with 40 byte buffer', async () => {
      expect(
        await hashTester.testMerkleHash(
          '0x00000000000000000000000000000000000000000000000000000000000000000000000000000000'
        )
      ).to.equal(
        '0x290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563'
      )
    })
    it('should work with 80 byte buffer', async () => {
      expect(
        await hashTester.testMerkleHash(
          '0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000'
        )
      ).to.equal(
        '0x290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563'
      )
    })

    it('should work with 128 byte buffer', async () => {
      expect(
        await hashTester.testMerkleHash(bufferToBytes(emptyBuffer(4)))
      ).to.equal(
        '0x290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563'
      )
    })

    it('should work with 65 byte buffer', async () => {
      const buf = [
        elem(0),
        elem(0),
        '0x0100000000000000000000000000000000000000000000000000000000000000',
        elem(0),
      ]
      expect(
        await hashTester.testMerkleHash(
          '0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001'
        )
      ).to.equal(merkleHash(buf, 0, 4))
    })

    it('should work with buffers with single word', async () => {
      for (let i = 16; i < 32; i++) {
        const buf = emptyBuffer(32)
        buf[i] = elem(1)
        expect(await hashTester.testMerkleHash(bufferToBytes(buf))).to.equal(
          merkleHash(buf, 0, 32)
        )
      }

      for (let i = 2; i < 4; i++) {
        const buf = emptyBuffer(4)
        buf[i] = elem(1)
        expect(await hashTester.testMerkleHash(bufferToBytes(buf))).to.equal(
          merkleHash(buf, 0, 4)
        )
      }
    })
  })
})
