/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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
import { Signer, BigNumberish } from 'ethers'
import { ContractTransaction } from '@ethersproject/contracts'
import { assert, expect } from 'chai'
import { ChallengeTester } from '../build/types/ChallengeTester'
import { Challenge } from '../build/types/Challenge'
// import { RollupTester } from '../build/types/RollupTester'
import { initializeAccounts } from './utils'

import { Node, NodeState, Assertion, RollupContract } from './rolluplib'

const initialVmState =
  '0x9900000000000000000000000000000000000000000000000000000000000000'
const zerobytes32 =
  '0x0000000000000000000000000000000000000000000000000000000000000000'

let accounts: Signer[]
let challengeTester: ChallengeTester

describe('Challenge', () => {
  before(async () => {
    accounts = await initializeAccounts()

    const OneStepProof = await ethers.getContractFactory('OneStepProof')
    const osp = await OneStepProof.deploy()
    await osp.deployed()

    const OneStepProof2 = await ethers.getContractFactory('OneStepProof2')
    const osp2 = await OneStepProof2.deploy()
    await osp2.deployed()

    const ChallengeTester = await ethers.getContractFactory('ChallengeTester')
    challengeTester = (await ChallengeTester.deploy(
      osp.address,
      osp2.address
    )) as ChallengeTester
    await challengeTester.deployed()
  })

  let challenge: Challenge
  let challengedNode: Node
  it('should initiate challenge', async function () {
    const block = await ethers.provider.getBlock('latest')
    const prevNodeState = new NodeState(
      block.number,
      0,
      initialVmState,
      zerobytes32,
      0,
      0,
      0,
      1
    )
    const assertion = new Assertion(
      prevNodeState,
      10000000,
      zerobytes32,
      [],
      [],
      []
    )
    challengedNode = new Node(assertion, 10, zerobytes32, 0)
    const fields = challengedNode.challengeFields()
    await challengeTester.startChallenge(
      fields[0],
      fields[1],
      fields[2],
      await accounts[0].getAddress(),
      await accounts[1].getAddress(),
      100,
      100
    )
    const challengeAddress = await challengeTester.challenge()
    const Challenge = await ethers.getContractFactory('Challenge')
    challenge = Challenge.attach(challengeAddress) as Challenge
  })

  it('should bisect execution', async function () {
    const chunks = Array(401).fill(
      challengedNode.assertion.startAssertionHash()
    )
    const tx = await challenge
      .connect(accounts[1])
      .bisectExecution(
        [],
        0,
        0,
        challengedNode.assertion.gasUsed,
        challengedNode.assertion.endAssertionHash(),
        0,
        challengedNode.assertion.startAssertionRestHash(),
        chunks
      )
    const receipt = await tx.wait()
    console.log('Bisection gas used', receipt.gasUsed.toNumber())
  })
})
