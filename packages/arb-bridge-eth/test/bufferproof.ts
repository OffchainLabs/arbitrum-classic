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

import { ethers } from '@nomiclabs/buidler'
import { utils } from 'ethers'
import { use, expect } from 'chai'
import chaiAsPromised from 'chai-as-promised'
import { BufferProofTester } from '../build/types/BufferProofTester'

use(chaiAsPromised)

let ospTester: BufferProofTester

type bytes32 = string

function keccak1(a: bytes32) : bytes32 {
  return utils.solidityKeccak256(['bytes32'], [a]);
}

function keccak2(a: bytes32, b: bytes32) : bytes32 {
  return utils.solidityKeccak256(['bytes32', 'bytes32'], [a, b]);
}

function merkleHash(arr: bytes32[], offset: number, sz: number) : bytes32 {
  if (sz === 1) {
    return keccak1(arr[offset])
  } else {
    return keccak2(merkleHash(arr, offset, sz/2), merkleHash(arr, offset+sz/2, sz/2))
  }
}

function makeProof(arr: bytes32[], offset: number, sz: number, loc: number) : bytes32[] {
  if (sz === 1) {
    return [arr[loc]]
  } else if (loc < offset + sz/2) {
    const proof = makeProof(arr, offset, sz/2, loc)
    return proof.concat([merkleHash(arr, offset+sz/2, sz/2)])
  } else {
    const proof = makeProof(arr, offset+sz/2, sz/2, loc)
    return proof.concat([merkleHash(arr, offset, sz/2)])
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
    left: merkleHash(arr, 0, sz/2),
    right: merkleHash(arr, sz/2, sz/2),
    height: makeProof(arr, 0, sz, 0).length,
  }
}

function elem(a: number): bytes32 {
  return '0x' + a.toString(16).padStart(64, '0')
}

function makeZeros(): bytes32[] {
  let zeros : bytes32[] = []
  let acc = keccak1(elem(0))
  for (let i = 0; i < 64; i++) {
    zeros.push(acc)
    acc = keccak2(zeros[i], zeros[i])
  }
  return zeros
}


describe('BufferProof', function () {
  before(async () => {
    const OneStepProof = await ethers.getContractFactory('BufferProofTester')
    ospTester = (await OneStepProof.deploy()) as BufferProofTester
    await ospTester.deployed()
  })

  describe('#get', function () {
    it('should work with correct proof', async () => {
      let arr : bytes32[] = []
      for (let i = 0; i < 32; i++) {
        arr.push(elem(i))
      }
      let buf = merkleHash(arr, 0, 32)
      let proof = makeProof(arr, 0, 32, 23)
      // console.log(proof)
      const res = await ospTester.testGet(buf, 23, proof)
      // console.log(res)
      expect(res).to.equal(elem(23))
    })

    it('should be filled with zeros', async () => {
      let arr : bytes32[] = []
      for (let i = 0; i < 32; i++) {
        arr.push(elem(i))
      }
      let buf = merkleHash(arr, 0, 32)
      let proof = makeProof(arr, 0, 32, 230%32)
      // console.log(proof)
      const res = await ospTester.testGet(buf, 230, proof)
      // console.log(res)
      expect(res).to.equal(elem(0))
    })
  })

  describe('#set', function () {
    it('should work with correct proof', async () => {
      let arr : bytes32[] = []
      for (let i = 0; i < 32; i++) {
        arr.push(elem(i))
      }
      let buf = merkleHash(arr, 0, 32)
      let proof = makeProof(arr, 0, 32, 23)
      console.log("proof length", proof.length)
      arr[23] = elem(10)
      const nproof = normalizationProof(arr, 32)
      const res = await ospTester.testSet(buf, 23, elem(10), proof, nproof.height, nproof.left, nproof.right)
      // console.log(res, merkleHash(arr, 0, 32))
      expect(res).to.equal(merkleHash(arr, 0, 32))
    })

    it('extending should work', async () => {
      let arr : bytes32[] = [elem(1)]
      let buf = merkleHash(arr, 0, 1)
      let proof = makeProof(arr, 0, 1, 0)
      let narr : bytes32[] = [elem(1), elem(2)]
      const nproof = normalizationProof(narr, 2)
      const res = await ospTester.testSet(buf, 1, elem(2), proof, nproof.height, nproof.left, nproof.right)
      // console.log(res, merkleHash(arr, 0, 32))
      expect(res).to.equal(merkleHash(narr, 0, 2))
    })

    it('extending more should work', async () => {
      let arr : bytes32[] = [elem(1)]
      let buf = merkleHash(arr, 0, 1)
      let proof = makeProof(arr, 0, 1, 0)
      let narr : bytes32[] = [elem(1), elem(0), elem(0), elem(0), elem(0), elem(2), elem(0), elem(0)]
      const nproof = normalizationProof(narr, 8)
      const res = await ospTester.testSet(buf, 5, elem(2), proof, nproof.height, nproof.left, nproof.right)
      // console.log(res, merkleHash(arr, 0, 32))
      expect(res).to.equal(merkleHash(narr, 0, 8))
    })

    it('shrinking should work', async () => {
      let arr : bytes32[] = [elem(1), elem(0), elem(0), elem(0), elem(0), elem(2), elem(0), elem(0)]
      let narr : bytes32[] = [elem(1)]
      let buf = merkleHash(arr, 0, 8)
      let narr2 : bytes32[] = [elem(1), elem(0), elem(0), elem(0), elem(0), elem(0), elem(0), elem(0)]
      let proof = makeProof(arr, 0, 8, 5)
      const nproof = normalizationProof(narr, 1)
      const res = await ospTester.testSet(buf, 5, elem(0), proof, nproof.height, nproof.left, nproof.right)
      // console.log(res, merkleHash(arr, 0, 32))
      expect(res).to.equal(merkleHash(narr, 0, 1))
    })
  })

})
