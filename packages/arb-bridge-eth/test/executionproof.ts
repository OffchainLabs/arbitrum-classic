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
import { OneStepProof } from '../build/types/OneStepProof'
import * as fs from 'fs'

use(chaiAsPromised)

interface Assertion {
  NumGas: number
  BeforeMachineHash: number[]
  AfterMachineHash: number[]
  BeforeInboxHash: number[]
  AfterInboxHash: number[]
  FirstMessageHash: number[]
  LastMessageHash: number[]
  FirstLogHash: number[]
  LastLogHash: number[]
}

interface Proof {
  Assertion: Assertion
  Proof: string
}

let ospTester: OneStepProof

describe('OneStepProof', async () => {
  before(async () => {
    const OneStepProof = await ethers.getContractFactory('OneStepProof')
    ospTester = (await OneStepProof.deploy()) as OneStepProof
    await ospTester.deployed()
  })

  const files = fs.readdirSync('./test/proofs')
  for (const filename of files) {
    const file = fs.readFileSync('./test/proofs/' + filename)
    const data = JSON.parse(file.toString()) as Proof[]
    it(`handle proofs from ${filename}`, async () => {
      let i = 0
      for (const proof of data) {
        if (i > 25) {
          // Some tests are too big to run every case
          return
        }
        const { fields, gas } = await ospTester.executeStep(
          proof.Assertion.AfterInboxHash,
          proof.Assertion.FirstMessageHash,
          proof.Assertion.FirstLogHash,
          Buffer.from(proof.Proof, 'base64')
        )
        expect(fields[0]).to.equal(
          utils.hexlify(proof.Assertion.BeforeMachineHash)
        )
        expect(fields[1]).to.equal(
          utils.hexlify(proof.Assertion.AfterMachineHash)
        )
        expect(fields[2]).to.equal(
          utils.hexlify(proof.Assertion.AfterInboxHash)
        )
        expect(fields[3]).to.equal(utils.hexlify(proof.Assertion.LastLogHash))
        expect(fields[4]).to.equal(
          utils.hexlify(proof.Assertion.LastMessageHash)
        )
        expect(gas).to.equal(proof.Assertion.NumGas)

        i++
      }
    }).timeout(20000)
  }
})
