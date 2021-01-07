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
import { BigNumber } from 'ethers'
import { ContractTransaction } from '@ethersproject/contracts'
import { TransactionReceipt } from '@ethersproject/providers'
import { BytesLike } from '@ethersproject/bytes'
import { use, expect } from 'chai'
import { OneStepProofTester } from '../build/types/OneStepProofTester'
import { BufferProofTester } from '../build/types/BufferProofTester'
import * as fs from 'fs'

const { utils } = ethers

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

async function executeStep(proof: Proof): Promise<ContractTransaction> {
  const proofData = Buffer.from(proof.Proof, 'base64')
  const bufferProofData = Buffer.from(proof.BufferProof || '', 'base64')
  return bufferProofData.length == 0
    ? await ospTester.executeStepTest(
        [
          proof.Assertion.AfterInboxHash,
          proof.Assertion.FirstMessageHash,
          proof.Assertion.FirstLogHash,
        ],
        proofData
      )
    : await ospTester2.executeStepTest(
        [
          proof.Assertion.AfterInboxHash,
          proof.Assertion.FirstMessageHash,
          proof.Assertion.FirstLogHash,
        ],
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
  const files = fs.readdirSync('./test/proofs')
  for (const filename of files) {
    const file = fs.readFileSync('./test/proofs/' + filename)
    const data = JSON.parse(file.toString()) as Proof[]
    describe(`proofs from ${filename}`, function () {
      const receipts: TransactionReceipt[] = []
      const opcodes: number[] = []
      before(async () => {
        for (const proof of data.slice(0, 50)) {
          const proofData = Buffer.from(proof.Proof, 'base64')
          const opcode = proofData[proofData.length - 1]
          if (opcode == 131) {
            // Skip too expensive opcode
            continue
          }
          const tx = await executeStep(proof)
          opcodes.push(opcode)
          receipts.push(await tx.wait())
        }
      })

      it(`should have correct proof`, async function () {
        for (let i = 0; i < receipts.length; i++) {
          const receipt = receipts[i]
          const opcode = opcodes[i]
          const proof = data[i]
          const message = `Opcode ${opcode}`
          const ev = ospTester.interface.parseLog(
            receipt.logs[receipt.logs.length - 1]
          )
          expect(ev.name, message).to.equal('OneStepProofResult')
          const parsedEv = (ev as any) as {
            args: {
              gas: BigNumber
              fields: [BytesLike, BytesLike, BytesLike, BytesLike, BytesLike]
            }
          }
          // console.log("opcode", opcode, fields)
          expect(parsedEv.args.fields[0], message).to.equal(
            utils.hexlify(proof.Assertion.BeforeMachineHash)
          )
          expect(parsedEv.args.fields[1], message).to.equal(
            utils.hexlify(proof.Assertion.AfterMachineHash)
          )
          expect(parsedEv.args.fields[2], message).to.equal(
            utils.hexlify(proof.Assertion.AfterInboxHash)
          )
          expect(parsedEv.args.fields[3], message).to.equal(
            utils.hexlify(proof.Assertion.LastLogHash)
          )
          expect(parsedEv.args.fields[4], message).to.equal(
            utils.hexlify(proof.Assertion.LastMessageHash)
          )
          expect(parsedEv.args.gas, message).to.equal(proof.Assertion.NumGas)
        }
      })

      it(`should have efficient proof [ @skip-on-coverage ]`, async function () {
        for (let i = 0; i < receipts.length; i++) {
          const receipt = receipts[i]
          const opcode = opcodes[i]
          if (receipt.gasUsed.toNumber() > 1000000) {
            console.log(`opcode ${opcode} used ${receipt.gasUsed} gas`)
          }
          expect(receipt.gasUsed.toNumber()).to.be.lessThan(5000000)
        }
      })
    })
  }
})
