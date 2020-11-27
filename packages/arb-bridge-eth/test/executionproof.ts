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
import { BufferProofTester } from '../build/types/BufferProofTester'
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
  BufferProof: string
}

let ospTester: OneStepProofTester
let ospTester2: BufferProofTester

async function executeStep(proof: Proof) {
  const proofData = Buffer.from(proof.Proof, 'base64')
  const bufferProofData = Buffer.from(proof.BufferProof || '', 'base64')
  return bufferProofData.length == 0
    ? await ospTester.executeStep(
        proof.Assertion.AfterInboxHash,
        proof.Assertion.FirstMessageHash,
        proof.Assertion.FirstLogHash,
        proofData
      )
    : await ospTester2.executeStep(
        proof.Assertion.AfterInboxHash,
        proof.Assertion.FirstMessageHash,
        proof.Assertion.FirstLogHash,
        proofData,
        bufferProofData
      )
}

async function executeTestStep(proof: Proof) {
  const proofData = Buffer.from(proof.Proof, 'base64')
  const bufferProofData = Buffer.from(proof.BufferProof || '', 'base64')
  return bufferProofData.length == 0
    ? await ospTester.executeStepTest(
        proof.Assertion.AfterInboxHash,
        proof.Assertion.FirstMessageHash,
        proof.Assertion.FirstLogHash,
        proofData
      )
    : await ospTester2.executeStepTest(
        proof.Assertion.AfterInboxHash,
        proof.Assertion.FirstMessageHash,
        proof.Assertion.FirstLogHash,
        proofData,
        bufferProofData
      )
}

describe('OneStepProof', function () {
  before(async () => {
    const OneStepProof = await ethers.getContractFactory('OneStepProofTester')
    ospTester = (await OneStepProof.deploy()) as OneStepProofTester
    await ospTester.deployed()

    const BufferProof = await ethers.getContractFactory('BufferProofTester')
    ospTester2 = (await BufferProof.deploy()) as BufferProofTester
    await ospTester2.deployed()
  })
/*
  it.only('brokne', async function () {
    const proof = {"Assertion":{"NumGas":100,"BeforeMachineHash":[22,26,222,129,246,169,171,31,19,48,37,252,54,210,192,111,173,62,58,236,98,13,130,109,227,132,115,15,18,177,54,51],"AfterMachineHash":[139,170,226,20,198,170,140,240,45,156,144,8,70,185,74,125,213,253,48,164,193,77,129,69,50,72,243,206,109,11,186,76],"BeforeInboxHash":[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],"AfterInboxHash":[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],"FirstMessageHash":[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],"LastMessageHash":[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],"MessageCount":0,"FirstLogHash":[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],"LastLogHash":[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],"LogCount":0},
    
    "Proof":"AwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA+gMI6X6CB+5tPg4MxPTfhUByEU+PqvuNkHkk1LYQs55bnZ8+6xg57IWgRvbtpSe8tpRJftP4uYRRFye0BgrCbseWLw2eJ56HigUNkZCKYKPgX1mEve0d9Zlkf+WqeBkvMmKAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHMbmPFQRxGkqQ+7LPxk44erbkP24+O/B+6Wp94OwRUNgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKArw2eJ56HigUNkZCKYKPgX1mEve0d9Zlkf+WqeBkvMmKAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAECep2sgv5Jt1soimVIHajYCrY/nZaomUYUkzxEoOwId/QAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAD/////////////////////////////////////////32tMAGFfldwkmTQHX8xmmUdZa69uBwrICln3na6YqpMvADAKY=",
    "BufferProof":"AQgLEhUAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACkN7NlUi2Ko1gNFqYg4b8hLpryVSEAI9jYvkxYO8+VjYz3E19pyVmYKiS+PFgSkS1QyZJzI7FyzztTE5qyU3R2JB0Co6wbOm+Qiy42lza/CtYwKXiQDbFeN4qQzyCj/fTuOwJ4Cb9wwU2XfyU4YmoGzjHWXs9lBwnnwQuggbgvYScSjQVW3vFQDtU/ETbDwwJhVi3FM+Np9Pf58O2Ic3CsbTLRvfxdrOhZstfFOwXfExpROCoYcntskb8UWQdazxgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHbT4mNtSIG+m6bSPv72mVwhDMPMx/XDmJDSoXKy1TVB4bTLRvfxdrOhZstfFOwXfExpROCoYcntskb8UWQdazxgAr3FRd8r3tAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKQ3s2VSLYqjWA0WpiDhvyEumvJVIQAj2Ni+TFg7z5WNjPcTX2nJWZgqJL48WBKRLVDJknMjsXLPO1MTmrJTdHYkHQKjrBs6b5CLLjaXNr8K1jApeJANsV43ipDPIKP99O47AngJv3DBTZd/JThiagbOMdZez2UHCefBC6CBuC9js1Q7uOOOGvWK+m+25kHBpUbZf4FO9nYpSGvdT0Tni2m0+JjbUiBvpum0j7+9plcIQzDzMf1w5iQ0qFystU1QeAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMVKSit8YiluSHCQV6DmMcfyH66DeJJwutDHfS+7k+62gIQ04PE5S49sCful6xTYvCJt/HVJawWeqqvQ8jwM6ct",
    "Message":null}
    const proofData = Buffer.from(proof.Proof, 'base64')
    const opcode = proofData[proofData.length - 1]
    const bufferProofData = Buffer.from(proof.BufferProof || '', 'base64')
    console.log(await ospTester2.parseProof(bufferProofData))
    const { fields, gas } = await executeStep(proof)
    console.log("opcode", opcode, fields)
  })
*/
  const files = fs.readdirSync('./test/proofs')
  for (const filename of files) {
    const file = fs.readFileSync('./test/proofs/' + filename)
    const data = JSON.parse(file.toString()) as Proof[]
    it(`should handle proofs from ${filename}`, async function () {
      this.timeout(60000)

      for (const proof of data.slice(0, 50)) {
        const proofData = Buffer.from(proof.Proof, 'base64')
        const opcode = proofData[proofData.length - 1]
        if (opcode == 131) {
          // Skip too expensive opcode
          continue
        }
        const { fields, gas } = await executeStep(proof)
        // console.log("opcode", opcode, fields)
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
        const tx = await executeTestStep(proof)
        const receipt = await tx.wait()
        const gas = receipt.gasUsed!.toNumber()
        if (gas > 1000000) {
          console.log(`opcode ${opcode} used ${gas} gas`)
        }
        expect(gas).to.be.lessThan(5000000)
      }
    })
  }
})
