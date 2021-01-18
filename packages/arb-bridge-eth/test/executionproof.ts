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

interface ExecutionCut {
  GasUsed: number
  InboxDelta: BytesLike
  MachineState: BytesLike
  SendAcc: BytesLike
  SendCount: BytesLike
  LogAcc: BytesLike
  LogCount: BytesLike
}

interface Proof {
  BeforeCut: ExecutionCut
  AfterCut: ExecutionCut
  Proof: string
  BufferProof: string
}

let ospTester: OneStepProofTester
let ospTester2: BufferProofTester

async function executeStep(proof: Proof): Promise<ContractTransaction> {
  const proofData = Buffer.from(proof.Proof, 'base64')
  const bufferProofData = Buffer.from(proof.BufferProof || '', 'base64')
  const machineFields: [BytesLike, BytesLike, BytesLike] = [
    proof.BeforeCut.InboxDelta,
    proof.BeforeCut.SendAcc,
    proof.BeforeCut.LogAcc,
  ]
  return bufferProofData.length == 0
    ? await ospTester.executeStepTest(machineFields, proofData)
    : await ospTester2.executeStepTest(
        machineFields,
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
    if (!filename.endsWith('json')) {
      continue
    }
    const file = fs.readFileSync('./test/proofs/' + filename)
    let data: Proof[]
    try {
      data = JSON.parse(file.toString()) as Proof[]
    } catch (e) {
      console.log(`Failed to load ${file}`)
      throw e
    }

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

          let message = `Opcode ${opcode}`
          if (i > 0) {
            message = `Opcode ${opcode.toString(16)}, Prev Opcode ${opcodes[
              i - 1
            ].toString(16)}`
          }
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
            utils.hexlify(proof.BeforeCut.MachineState)
          )
          expect(parsedEv.args.fields[1], message).to.equal(
            utils.hexlify(proof.AfterCut.MachineState)
          )
          expect(parsedEv.args.fields[2], message).to.equal(
            utils.hexlify(proof.AfterCut.InboxDelta)
          )
          expect(parsedEv.args.fields[3], message).to.equal(
            utils.hexlify(proof.AfterCut.SendAcc)
          )
          expect(parsedEv.args.fields[4], message).to.equal(
            utils.hexlify(proof.AfterCut.LogAcc)
          )
          expect(parsedEv.args.gas, message).to.equal(
            proof.AfterCut.GasUsed - proof.BeforeCut.GasUsed
          )
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
