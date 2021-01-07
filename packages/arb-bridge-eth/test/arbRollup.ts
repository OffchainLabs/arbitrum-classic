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
import {
  Signer,
  ContractTransaction,
  BigNumberish,
  providers,
  utils,
} from 'ethers'
import * as chai from 'chai'
import * as chaiAsPromised from 'chai-as-promised'
import { Rollup } from '../build/types/Rollup'
import { Node } from '../build/types/Node'
import { RollupCreator } from '../build/types/RollupCreator'
import { Challenge } from '../build/types/Challenge'
// import { RollupTester } from '../build/types/RollupTester'
import { ArbValue } from 'arb-provider-ethers'
import deploy_contracts from '../scripts/deploy'
import { initializeAccounts } from './utils'

import { NodeState, Assertion, RollupContract } from './rolluplib'

chai.use(chaiAsPromised)

const { assert, expect } = chai

const initialVmState =
  '0x9900000000000000000000000000000000000000000000000000000000000000'
const zerobytes32 =
  '0x0000000000000000000000000000000000000000000000000000000000000000'
const stakeRequirement = 10
const stakeToken = '0x0000000000000000000000000000000000000000'
const challengePeriodBlocks = 100
const arbGasSpeedLimitPerBlock = 1000000

let rollupCreator: RollupCreator
let rollup: RollupContract
let challenge: Challenge
// let rollupTester: RollupTester
// let assertionInfo: Assertion
let accounts: Signer[]

async function createRollup(): Promise<{
  rollupCon: Rollup
  blockCreated: number
}> {
  const tx = rollupCreator.createRollup(
    initialVmState,
    challengePeriodBlocks,
    arbGasSpeedLimitPerBlock,
    stakeRequirement,
    stakeToken,
    await accounts[0].getAddress(), // owner
    '0x'
  )
  await expect(tx).to.emit(rollupCreator, 'RollupCreated')

  const receipt = await (await tx).wait()
  if (receipt.logs == undefined) {
    throw Error('expected receipt to have logs')
  }

  const ev = rollupCreator.interface.parseLog(
    receipt.logs[receipt.logs.length - 1]
  )
  expect(ev.name).to.equal('RollupCreated')
  const parsedEv = (ev as any) as { args: { rollupAddress: string } }
  const Rollup = await ethers.getContractFactory('Rollup')
  return {
    rollupCon: Rollup.attach(parsedEv.args.rollupAddress) as Rollup,
    blockCreated: receipt.blockNumber!,
  }
}

async function tryAdvanceChain(blocks: number): Promise<void> {
  try {
    for (let i = 0; i < blocks; i++) {
      await ethers.provider.send('evm_mine', [])
    }
  } catch (e) {
    // EVM mine failed. Try advancing the chain by sending txes if the node
    // is in dev mode and mints blocks when txes are sent
    for (let i = 0; i < blocks; i++) {
      const tx = await accounts[0].sendTransaction({
        value: 0,
        to: await accounts[0].getAddress(),
      })
      await tx.wait()
    }
  }
}

function makeSimpleAssertion(
  prevNodeState: NodeState,
  numGas: BigNumberish
): Assertion {
  return new Assertion(prevNodeState, 100, numGas, zerobytes32, [], [], [])
}

let prevNodeState: NodeState

describe('ArbRollup', () => {
  it('should deploy contracts', async function () {
    accounts = await initializeAccounts()
    const { RollupCreator } = await deploy_contracts()
    rollupCreator = RollupCreator as RollupCreator

    // const RollupTester = await ethers.getContractFactory('RollupTester')
    // rollupTester = (await RollupTester.deploy()) as RollupTester
    // await rollupTester.deployed()
  })

  it('should initialize', async function () {
    const { rollupCon, blockCreated } = await createRollup()
    rollup = new RollupContract(rollupCon)
    const originalNode = await rollup.latestConfirmed()
    const nodeAddress = await rollup.nodes(originalNode)

    const Node = await ethers.getContractFactory('Node')
    const node = Node.attach(nodeAddress) as Node

    prevNodeState = new NodeState(
      blockCreated,
      0,
      initialVmState,
      zerobytes32,
      0,
      0,
      0,
      0
    )

    assert.equal(
      await node.stateHash(),
      prevNodeState.hash(),
      'initial confirmed node should have set initial state'
    )
  })

  it('should place stake on new node', async function () {
    const block = await ethers.provider.getBlock('latest')
    await tryAdvanceChain(challengePeriodBlocks / 10)
    const assertion = makeSimpleAssertion(prevNodeState, 100)
    const tx = await rollup.newStakeOnNewNode(block, assertion, 1, 0, {
      value: 10,
    })

    const receipt = await tx.wait()
    prevNodeState = assertion.createdNodeState(receipt.blockNumber!, 1)
  })

  it('should let a new staker place on existing node', async function () {
    const block = await ethers.provider.getBlock('latest')
    await rollup
      .connect(accounts[1])
      .newStakeOnExistingNode(block, 1, { value: 10 })
  })

  it('should move stake to a new node', async function () {
    await tryAdvanceChain(challengePeriodBlocks / 10)
    const block = await ethers.provider.getBlock('latest')
    const assertion = makeSimpleAssertion(
      prevNodeState,
      (block.number - prevNodeState.proposedBlock + 1) *
        arbGasSpeedLimitPerBlock
    )
    const tx = await rollup.addStakeOnNewNode(block, assertion, 2)
    const receipt = await tx.wait()
    prevNodeState = assertion.createdNodeState(receipt.blockNumber!, 1)
  })

  it('should let the second staker place on the new node', async function () {
    const block = await ethers.provider.getBlock('latest')
    await rollup.connect(accounts[1]).addStakeOnExistingNode(block, 2)
  })

  it('should confirm node', async function () {
    await tryAdvanceChain(challengePeriodBlocks)
    await rollup.confirmNextNode(zerobytes32, [])
  })

  it('should confirm next node', async function () {
    await tryAdvanceChain(challengePeriodBlocks / 10)
    await rollup.confirmNextNode(zerobytes32, [])
  })

  let challengedAssertion: Assertion
  let validNodeState: NodeState
  it('should let the first staker make another node', async function () {
    await tryAdvanceChain(challengePeriodBlocks / 10)
    const block = await ethers.provider.getBlock('latest')
    challengedAssertion = makeSimpleAssertion(
      prevNodeState,
      (block.number - prevNodeState.proposedBlock + 1) *
        arbGasSpeedLimitPerBlock
    )
    const tx = await rollup.addStakeOnNewNode(block, challengedAssertion, 3)
    const receipt = await tx.wait()
    validNodeState = challengedAssertion.createdNodeState(
      receipt.blockNumber!,
      1
    )
  })

  it('should let the second staker make a conflicting node', async function () {
    await tryAdvanceChain(challengePeriodBlocks / 10)
    const block = await ethers.provider.getBlock('latest')
    const assertion = makeSimpleAssertion(
      prevNodeState,
      (block.number - prevNodeState.proposedBlock + 1) *
        arbGasSpeedLimitPerBlock
    )
    await rollup.connect(accounts[1]).addStakeOnNewNode(block, assertion, 4)
  })

  it('should fail to confirm first staker node', async function () {
    await tryAdvanceChain(
      challengePeriodBlocks +
        challengedAssertion.checkTime(arbGasSpeedLimitPerBlock)
    )
    await expect(rollup.confirmNextNode(zerobytes32, [])).to.be.revertedWith(
      'NOT_ALL_STAKED'
    )
  })

  let challenge: Challenge
  it('should initiate a challenge', async function () {
    const tx = rollup.createChallenge(
      await accounts[0].getAddress(),
      3,
      await accounts[1].getAddress(),
      4,
      challengedAssertion,
      await rollup.inboxMaxValue(),
      1
    )
    expect(tx).to.emit(rollup, 'RollupChallengeStarted')
    const receipt = await (await tx).wait()
    const ev = rollup.rollup.interface.parseLog(
      receipt.logs![receipt.logs!.length - 1]
    )
    expect(ev.name).to.equal('RollupChallengeStarted')
    const parsedEv = (ev as any) as { args: { challengeContract: string } }
    const Challenge = await ethers.getContractFactory('Challenge')
    challenge = Challenge.attach(parsedEv.args.challengeContract) as Challenge
  })

  it('should win via timeout', async function () {
    await tryAdvanceChain(
      challengePeriodBlocks +
        challengedAssertion.checkTime(arbGasSpeedLimitPerBlock) +
        1
    )
    await challenge.timeout()
  })

  it('confirm first staker node', async function () {
    await rollup.confirmNextNode(zerobytes32, [])
  })

  it('should reject out of order second node', async function () {
    await rollup.rejectNextNodeOutOfOrder()
  })

  it('should make a new node', async function () {
    await tryAdvanceChain(challengePeriodBlocks / 10)
    const block = await ethers.provider.getBlock('latest')
    challengedAssertion = makeSimpleAssertion(
      validNodeState,
      (block.number - validNodeState.proposedBlock + 1) *
        arbGasSpeedLimitPerBlock
    )
    await rollup.addStakeOnNewNode(block, challengedAssertion, 5)
  })

  it('new staker should make a conflicting node', async function () {
    await tryAdvanceChain(challengePeriodBlocks / 10)
    const block = await ethers.provider.getBlock('latest')
    const assertion = makeSimpleAssertion(
      validNodeState,
      (block.number - validNodeState.proposedBlock + 1) *
        arbGasSpeedLimitPerBlock
    )
    const stake = await rollup.currentRequiredStake()
    await rollup
      .connect(accounts[2])
      .newStakeOnNewNode(block, assertion, 6, 3, { value: stake })
  })

  it('should initiate another challenge', async function () {
    const tx = rollup.createChallenge(
      await accounts[0].getAddress(),
      5,
      await accounts[2].getAddress(),
      6,
      challengedAssertion,
      await rollup.inboxMaxValue(),
      1
    )
    expect(tx).to.emit(rollup, 'RollupChallengeStarted')
    const receipt = await (await tx).wait()
    const ev = rollup.rollup.interface.parseLog(
      receipt.logs![receipt.logs!.length - 1]
    )
    expect(ev.name).to.equal('RollupChallengeStarted')
    const parsedEv = (ev as any) as { args: { challengeContract: string } }
    const Challenge = await ethers.getContractFactory('Challenge')
    challenge = Challenge.attach(parsedEv.args.challengeContract) as Challenge
  })
})
