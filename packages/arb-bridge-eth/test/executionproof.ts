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
import { use, expect, assert } from 'chai'
import { OneStepProofTester } from '../build/types/OneStepProofTester'
import { IOneStepProof } from '../build/types/IOneStepProof'
import { Bridge } from '../build/types/Bridge'
import * as fs from 'fs'

const { utils } = ethers

interface ExecutionCut {
  GasUsed: number
  TotalMessagesRead: BytesLike
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
let executors: IOneStepProof[]
let bridge: Bridge

function getProver(op: number) {
  if ((op >= 0xa1 && op <= 0xa6) || op == 0x70) {
    return 1
  } else if (op >= 0x20 && op <= 0x24) {
    return 2
  } else {
    return 0
  }
}

describe('OneStepProof', function () {
  before(async () => {
    const OneStepProofTester = await ethers.getContractFactory(
      'OneStepProofTester'
    )
    ospTester = (await OneStepProofTester.deploy()) as OneStepProofTester
    await ospTester.deployed()

    const OneStepProof = await ethers.getContractFactory('OneStepProof')
    const osp1 = (await OneStepProof.deploy()) as IOneStepProof
    await osp1.deployed()

    const OneStepProof2 = await ethers.getContractFactory('OneStepProof2')
    const osp2 = (await OneStepProof2.deploy()) as IOneStepProof
    await osp2.deployed()

    const OneStepProofHash = await ethers.getContractFactory('OneStepProofHash')
    const osp3 = (await OneStepProofHash.deploy()) as IOneStepProof
    await osp3.deployed()

    executors = [osp1, osp2, osp3]

    const Bridge = await ethers.getContractFactory('Bridge')
    bridge = (await Bridge.deploy()) as Bridge
    await bridge.deployed()
  })
  const files = fs.readdirSync('./test/proofs')
  for (const filename of files) {
    if (!filename.endsWith('json')) {
      continue
    }
    if (
      filename == 'opcodetestecops.mexe-proofs.json' ||
      filename == 'opcodeecpairing.mexe-proofs.json'
    ) {
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
      this.timeout(100000)
      const receipts: TransactionReceipt[] = []
      const opcodes: number[] = []
      it(`should execute steps`, async function () {
        for (const proof of data.slice(0, 50)) {
          const proofData = ethers.utils.arrayify(proof.Proof)
          const opcode = proofData[0]
          if (opcode == 131) {
            // Skip too expensive opcode
            continue
          }
          const prover = getProver(opcode)
          try {
            const tx = await ospTester.executeStepTest(
              executors[prover].address,
              bridge.address,
              proof.BeforeCut.TotalMessagesRead,
              [proof.BeforeCut.SendAcc, proof.BeforeCut.LogAcc],
              proof.Proof,
              proof.BufferProof
            )

            const receipt = await tx.wait()
            receipts.push(receipt)
            opcodes.push(opcode)
          } catch (e) {
            assert.fail(`Failed to generate proof ${opcode}, ${prover}`)
          }
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
              totalMessagesRead: BigNumber
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
            utils.hexlify(proof.AfterCut.SendAcc)
          )
          expect(parsedEv.args.fields[3], message).to.equal(
            utils.hexlify(proof.AfterCut.LogAcc)
          )
          expect(parsedEv.args.totalMessagesRead, message).to.equal(
            BigNumber.from(proof.AfterCut.TotalMessagesRead)
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
