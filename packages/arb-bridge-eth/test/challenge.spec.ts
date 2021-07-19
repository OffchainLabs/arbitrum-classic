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
import { Bridge } from '../build/types/Bridge'
import { SequencerInbox } from '../build/types/SequencerInbox'
import { initializeAccounts } from './utils'

import {
  Node,
  ExecutionState,
  NodeState,
  Assertion,
  RollupContract,
} from './rolluplib'

const initialVmState =
  '0x9900000000000000000000000000000000000000000000000000000000000000'
const zerobytes32 =
  '0x0000000000000000000000000000000000000000000000000000000000000000'

let accounts: Signer[]
let challengeTester: ChallengeTester
let bridge: Bridge
let sequencerInbox: SequencerInbox

describe('Challenge', () => {
  before(async () => {
    accounts = await initializeAccounts()

    const OneStepProof = await ethers.getContractFactory('OneStepProof')
    const osp = await OneStepProof.deploy()
    await osp.deployed()

    const OneStepProof2 = await ethers.getContractFactory('OneStepProof2')
    const osp2 = await OneStepProof2.deploy()
    await osp2.deployed()

    const OneStepProof3 = await ethers.getContractFactory('OneStepProofHash')
    const osp3 = await OneStepProof3.deploy()
    await osp3.deployed()

    const ChallengeTester = await ethers.getContractFactory('ChallengeTester')
    challengeTester = (await ChallengeTester.deploy([
      osp.address,
      osp2.address,
      osp3.address,
    ])) as ChallengeTester
    await challengeTester.deployed()

    const Bridge = await ethers.getContractFactory('Bridge')
    bridge = (await Bridge.deploy()) as Bridge
    await bridge.deployed()
    await bridge.initialize()

    const RollupMock = await ethers.getContractFactory('RollupMock')
    const rollupMock = await RollupMock.deploy()
    await rollupMock.deployed()
    await rollupMock.setMock(15, 900)

    const SequencerInbox = await ethers.getContractFactory('SequencerInbox')
    sequencerInbox = (await SequencerInbox.deploy()) as SequencerInbox
    await sequencerInbox.deployed()
    await sequencerInbox.initialize(
      bridge.address,
      await accounts[0].getAddress(),
      rollupMock.address
    )
  })

  let challenge: Challenge
  let challengedNode: Node
  it('should initiate challenge', async function () {
    const block = await ethers.provider.getBlock('latest')

    const prevNodeState = new NodeState(
      new ExecutionState(0, initialVmState, 0, 0, 0, zerobytes32, zerobytes32),
      block.number,
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
    challengedNode = new Node(assertion, 10, 0, zerobytes32)
    await challengeTester.startChallenge(
      challengedNode.executionHash(),
      challengedNode.afterState.execState.inboxCount,
      await accounts[0].getAddress(),
      await accounts[1].getAddress(),
      100,
      100,
      sequencerInbox.address,
      bridge.address
    )
    const challengeAddress = await challengeTester.challenge()
    const Challenge = await ethers.getContractFactory('Challenge')
    challenge = Challenge.attach(challengeAddress) as Challenge
  })

  it('should bisect execution', async function () {
    const chunks = Array(401).fill(
      challengedNode.beforeState.execState.challengeHash()
    )
    const tx = await challenge
      .connect(accounts[1])
      .bisectExecution(
        [],
        0,
        0,
        challengedNode.gasUsed(),
        challengedNode.afterState.execState.challengeHash(),
        0,
        challengedNode.beforeState.execState.challengeRestHash(),
        chunks
      )
    const receipt = await tx.wait()
    console.log('Bisection gas used', receipt.gasUsed.toNumber())
  })
})
