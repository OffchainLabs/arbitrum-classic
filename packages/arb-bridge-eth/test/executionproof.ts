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
import { OneStepProofTesterFactory } from '../build/types/OneStepProofTesterFactory'
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

let ospTester: OneStepProofTester

describe('OneStepProof', function () {
  before(async () => {
    const UtilLibrary = await ethers.getContractFactory('MerkleUtil')
    const utilLibrary = await UtilLibrary.deploy()
    const link = {
      ["__$dadff4c8e57a85477fa98436c23c3d6d3b$__"]: utilLibrary.address
    }
    const OneStepProof = new OneStepProofTesterFactory(link, (await ethers.signers())[0])
    // await ethers.getContractFactory('OneStepProofTester')
    ospTester = (await OneStepProof.deploy()) as OneStepProofTester
    await ospTester.deployed()
  })

  const files = fs.readdirSync('./test/proofs')
  for (const filename of files) {
    const file = fs.readFileSync('./test/proofs/' + filename)
    const data = JSON.parse(file.toString()) as Proof[]
    it(`should handle proofs from ${filename}`, async function () {
      this.timeout(60000)

      for (const proof of data.slice(0, 25)) {
        const proofData = Buffer.from(proof.Proof, 'base64')
        const opcode = proofData[proofData.length - 1]
        if (opcode == 131) {
          // Skip too expensive opcode
          continue
        }
        const { fields, gas } = await ospTester.executeStep(
          proof.Assertion.AfterInboxHash,
          proof.Assertion.FirstMessageHash,
          proof.Assertion.FirstLogHash,
          proofData
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
      }
    })

    it(`efficiently run proofs from ${filename} [ @skip-on-coverage ]`, async function () {
      this.timeout(60000)

      for (const proof of data.slice(0, 25)) {
        const proofData = Buffer.from(proof.Proof, 'base64')
        const opcode = proofData[proofData.length - 1]
        const tx = await ospTester.executeStepTest(
          proof.Assertion.AfterInboxHash,
          proof.Assertion.FirstMessageHash,
          proof.Assertion.FirstLogHash,
          proofData
        )
        const receipt = await tx.wait()
        const gas = receipt.gasUsed!.toNumber()
        if (gas > 1000000) {
          console.log(`opcode ${opcode} used ${gas} gas`)
        }
        expect(gas).to.be.lessThan(2500000)
      }
    })
  }
})
