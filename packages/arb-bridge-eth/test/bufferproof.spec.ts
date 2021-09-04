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
import { BufferProofTester } from '../build/types/BufferProofTester'

const { utils } = ethers

let ospTester: BufferProofTester

type bytes32 = string

function keccak1(a: bytes32): bytes32 {
  return utils.solidityKeccak256(['bytes32'], [a])
}

function keccak2(a: bytes32, b: bytes32): bytes32 {
  return utils.solidityKeccak256(['bytes32', 'bytes32'], [a, b])
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

function makeProof(
  arr: bytes32[],
  offset: number,
  sz: number,
  loc: number
): bytes32[] {
  if (sz === 1) {
    return [arr[loc]]
  } else if (loc < offset + sz / 2) {
    const proof = makeProof(arr, offset, sz / 2, loc)
    const hash = merkleHash(arr, offset + sz / 2, sz / 2)
    return proof.concat([hash])
  } else {
    const proof = makeProof(arr, offset + sz / 2, sz / 2, loc)
    const hash = merkleHash(arr, offset, sz / 2)
    return proof.concat([hash])
  }
}

function normalizationProof(arr: bytes32[], sz: number) {
  if (sz == 1) {
    return {
      left: merkleHash(arr, 0, 1),
      right: merkleHash(arr, 0, 1),
      height: 0,
    }
  }
  return {
    left: merkleHash(arr, 0, sz / 2),
    right: merkleHash(arr, sz / 2, sz / 2),
    height: makeProof(arr, 0, sz, 0).length,
  }
}

function elem(a: number): bytes32 {
  return '0x' + a.toString(16).padStart(64, '0')
}

export function makeZeros(): bytes32[] {
  const zeros: bytes32[] = []
  let acc = keccak1(elem(0))
  for (let i = 0; i < 64; i++) {
    zeros.push(acc)
    acc = keccak2(zeros[i], zeros[i])
  }
  return zeros
}

export function fromBytes(buf: Buffer): bytes32[] {
  const str = buf.toString('hex')
  const res = []
  for (let i = 0; i < buf.length / 32; i++) {
    res.push('0x' + str.substr(i * 64, 64))
  }
  return res
}

function testArray1() {
  const arr: bytes32[] = []
  for (let i = 0; i < 31; i++) {
    arr.push(elem(i))
  }
  arr.push(elem(0))
  return arr
}

function testArray2() {
  const arr: bytes32[] = []
  for (let i = 0; i < 31; i++) {
    arr.push(elem(i))
  }
  arr.push('0xffffffff00000000000000000000000000000000000000000000000000000000')
  return arr
}

function testArray3() {
  const arr: bytes32[] = []
  for (let i = 0; i < 17; i++) {
    arr.push(elem(i))
  }
  for (let i = 17; i < 32; i++) {
    arr.push(elem(0))
  }
  return arr
}

describe('BufferProof', function () {
  before(async () => {
    const OneStepProof = await ethers.getContractFactory('BufferProofTester')
    ospTester = (await OneStepProof.deploy()) as BufferProofTester
    await ospTester.deployed()
  })

  describe('#get', function () {
    it('should work with correct proof', async () => {
      const arr: bytes32[] = []
      for (let i = 0; i < 32; i++) {
        arr.push(elem(i))
      }
      const buf = merkleHash(arr, 0, 32)
      const proof = makeProof(arr, 0, 32, 23)
      const res = await ospTester.testGet(buf, 23, proof)
      expect(res).to.equal(elem(23))
    })

    it('should be filled with zeros', async () => {
      const arr: bytes32[] = []
      for (let i = 0; i < 32; i++) {
        arr.push(elem(i))
      }
      const buf = merkleHash(arr, 0, 32)
      const proof = makeProof(arr, 0, 32, 230 % 32)
      const res = await ospTester.testGet(buf, 230, proof)
      expect(res).to.equal(elem(0))
    })
  })

  describe('#set', function () {
    it('should work with correct proof', async () => {
      const arr: bytes32[] = []
      for (let i = 0; i < 32; i++) {
        arr.push(elem(i))
      }
      const buf = merkleHash(arr, 0, 32)
      const proof = makeProof(arr, 0, 32, 23)
      arr[23] = elem(10)
      const nproof = normalizationProof(arr, 32)
      const res = await ospTester.testSet(
        buf,
        23,
        elem(10),
        proof,
        nproof.height,
        nproof.left,
        nproof.right
      )
      expect(res).to.equal(merkleHash(arr, 0, 32))
    })

    it('extending should work', async () => {
      const arr: bytes32[] = [elem(1)]
      const buf = merkleHash(arr, 0, 1)
      const proof = makeProof(arr, 0, 1, 0)
      const narr: bytes32[] = [elem(1), elem(2)]
      const nproof = normalizationProof(narr, 2)
      const res = await ospTester.testSet(
        buf,
        1,
        elem(2),
        proof,
        nproof.height,
        nproof.left,
        nproof.right
      )
      expect(res).to.equal(merkleHash(narr, 0, 2))
    })

    it('extending more should work', async () => {
      const arr: bytes32[] = [elem(1)]
      const buf = merkleHash(arr, 0, 1)
      const proof = makeProof(arr, 0, 1, 0)
      const narr: bytes32[] = [
        elem(1),
        elem(0),
        elem(0),
        elem(0),
        elem(0),
        elem(2),
        elem(0),
        elem(0),
      ]
      const nproof = normalizationProof(narr, 8)
      const res = await ospTester.testSet(
        buf,
        5,
        elem(2),
        proof,
        nproof.height,
        nproof.left,
        nproof.right
      )
      expect(res).to.equal(merkleHash(narr, 0, 8))
    })

    it('shrinking should work', async () => {
      const arr: bytes32[] = [
        elem(1),
        elem(0),
        elem(0),
        elem(0),
        elem(0),
        elem(2),
        elem(0),
        elem(0),
      ]
      const narr: bytes32[] = [elem(1)]
      const buf = merkleHash(arr, 0, 8)
      const proof = makeProof(arr, 0, 8, 5)
      const nproof = normalizationProof(narr, 1)
      const res = await ospTester.testSet(
        buf,
        5,
        elem(0),
        proof,
        nproof.height,
        nproof.left,
        nproof.right
      )
      expect(res).to.equal(merkleHash(narr, 0, 1))
    })
  })

  describe('#checkSize', function () {
    it('should work when actual size is larger', async () => {
      const arr = testArray1()
      const buf = merkleHash(arr, 0, 32)
      const proof = makeProof(arr, 0, 32, 23)
      const res = await ospTester.testCheckSize(buf, 23 * 32 + 12, proof)
      expect(res).to.equal(false)
    })
    it('should work when proof length is too small', async () => {
      const arr = testArray1()
      const buf = merkleHash(arr, 0, 32)
      const proof = makeProof(arr, 0, 32, 230 % 32)
      const res = await ospTester.testCheckSize(buf, 230 * 32 + 12, proof)
      expect(res).to.equal(true)
    })
    it('should work with the exact size', async () => {
      const arr = testArray1()
      const buf = merkleHash(arr, 0, 32)
      const proof = makeProof(arr, 0, 32, 31)
      const res = await ospTester.testCheckSize(buf, 31 * 32, proof)
      expect(res).to.equal(true)
    })
    it('should work inside words', async () => {
      const arr = testArray2()
      const buf = merkleHash(arr, 0, 32)
      const proof = makeProof(arr, 0, 32, 31)
      const res = await ospTester.testCheckSize(buf, 31 * 32, proof)
      expect(res).to.equal(false)
      const res2 = await ospTester.testCheckSize(buf, 31 * 32 + 4, proof)
      expect(res2).to.equal(true)
    })
    it('should work with empty right subtrees', async () => {
      const arr = testArray3()
      const buf = merkleHash(arr, 0, 32)
      const proof = makeProof(arr, 0, 32, 17)
      const res = await ospTester.testCheckSize(buf, 17 * 32, proof)
      expect(res).to.equal(true)
    })
    it('should not work with non-empty right subtrees', async () => {
      const arr = testArray3()
      arr[31] = elem(123)
      const buf = merkleHash(arr, 0, 32)
      const proof = makeProof(arr, 0, 32, 17)
      const res = await ospTester.testCheckSize(buf, 17 * 32, proof)
      expect(res).to.equal(false)
    })
  })
})
