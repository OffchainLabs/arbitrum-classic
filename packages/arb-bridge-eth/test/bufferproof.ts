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

function merkleHash(arr: bytes32[], offset: number, sz: number) {
  
}

describe('OneStepProof', function () {
  before(async () => {
    const OneStepProof = await ethers.getContractFactory('BufferProofTester')
    ospTester = (await OneStepProof.deploy()) as BufferProofTester
    await ospTester.deployed()
  })

  describe('#get', function () {

    it('should work with correct proof', async () => {

    })

  })

})
