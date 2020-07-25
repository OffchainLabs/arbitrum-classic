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
import { OneStepProofTester } from '../build/types/OneStepProofTester'
import * as fs from 'fs'

use(chaiAsPromised)

interface Assertion {
  AfterHash: number[]
  DidInboxInsn: boolean
  NumGas: number
  FirstMessageHash: number[]
  LastMessageHash: number[]
  FirstLogHash: number[]
  LastLogHash: number[]
}

interface Proof {
  BeforeHash: number[]
  Assertion: Assertion
  InboxInner: number[]
  InboxSize: number
  Proof: string
}

let ospTester: OneStepProofTester

describe('OneStepProof', async () => {
  before(async () => {
    const OneStepProofTester = await ethers.getContractFactory(
      'OneStepProofTester'
    )
    ospTester = (await OneStepProofTester.deploy()) as OneStepProofTester
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
        const afterHash = await ospTester.validateProof(
          proof.BeforeHash,
          proof.InboxInner,
          proof.InboxSize,
          proof.Assertion.DidInboxInsn,
          proof.Assertion.FirstMessageHash,
          proof.Assertion.LastMessageHash,
          proof.Assertion.FirstLogHash,
          proof.Assertion.LastLogHash,
          proof.Assertion.NumGas,
          Buffer.from(proof.Proof, 'base64')
        )
        expect(afterHash).to.equal(utils.hexlify(proof.Assertion.AfterHash))
        i++
      }
    }).timeout(20000)
  }
})
