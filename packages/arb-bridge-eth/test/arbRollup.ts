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
import { Rollup } from '../build/types/Rollup'
import { Node as NodeCon } from '../build/types/Node'
import { RollupCreator } from '../build/types/RollupCreator'
import { Challenge } from '../build/types/Challenge'
// import { RollupTester } from '../build/types/RollupTester'
import deploy_contracts from '../scripts/deploy'
import { initializeAccounts } from './utils'

import { Node, NodeState, Assertion, RollupContract } from './rolluplib'

const initialVmState =
  '0x9900000000000000000000000000000000000000000000000000000000000000'
const zerobytes32 =
  '0x0000000000000000000000000000000000000000000000000000000000000000'
const stakeRequirement = 10
const stakeToken = '0x0000000000000000000000000000000000000000'
const confirmationPeriodBlocks = 100
const arbGasSpeedLimitPerBlock = 1000000
const minimumAssertionPeriod = 75

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
  const tx = rollupCreator.createRollupNoProxy(
    initialVmState,
    confirmationPeriodBlocks,
    0,
    arbGasSpeedLimitPerBlock,
    stakeRequirement,
    stakeToken,
    await accounts[0].getAddress(), // owner
    '0x'
  )

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
  gasUsed: BigNumberish
): Assertion {
  return new Assertion(prevNodeState, gasUsed, zerobytes32, [], [], [])
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
    const nodeAddress = await rollup.getNode(originalNode)

    const Node = await ethers.getContractFactory('Node')
    const node = Node.attach(nodeAddress) as NodeCon

    prevNodeState = new NodeState(
      blockCreated,
      0,
      initialVmState,
      zerobytes32,
      0,
      0,
      0,
      1
    )

    assert.equal(
      await node.stateHash(),
      prevNodeState.hash(),
      'initial confirmed node should have set initial state'
    )
  })

  it('should place stake', async function () {
    const stake = await rollup.currentRequiredStake()
    await rollup.newStake(0, { value: stake })
  })

  it('should place stake on new node', async function () {
    await tryAdvanceChain(minimumAssertionPeriod)
    const block = await ethers.provider.getBlock('latest')
    const assertion = makeSimpleAssertion(
      prevNodeState,
      (block.number - prevNodeState.proposedBlock + 1) *
        arbGasSpeedLimitPerBlock
    )

    const { node } = await rollup.stakeOnNewNode(block, assertion, 1)
    prevNodeState = node.afterNodeState()
  })

  it('should let a new staker place on existing node', async function () {
    const block = await ethers.provider.getBlock('latest')
    await rollup.connect(accounts[1]).newStake(0, { value: 10 })

    await rollup.connect(accounts[1]).stakeOnExistingNode(block, 1)
  })

  it('should move stake to a new node', async function () {
    await tryAdvanceChain(minimumAssertionPeriod)
    const block = await ethers.provider.getBlock('latest')
    const assertion = makeSimpleAssertion(
      prevNodeState,
      (block.number - prevNodeState.proposedBlock + 1) *
        arbGasSpeedLimitPerBlock
    )
    const { node } = await rollup.stakeOnNewNode(block, assertion, 2)
    prevNodeState = node.afterNodeState()
  })

  it('should let the second staker place on the new node', async function () {
    const block = await ethers.provider.getBlock('latest')
    await rollup.connect(accounts[1]).stakeOnExistingNode(block, 2)
  })

  it('should confirm node', async function () {
    await tryAdvanceChain(confirmationPeriodBlocks * 2)
    await rollup.confirmNextNode(zerobytes32, [])
  })

  it('should confirm next node', async function () {
    await tryAdvanceChain(minimumAssertionPeriod)
    await rollup.confirmNextNode(zerobytes32, [])
  })

  let challengedNode: Node
  let validNodeState: NodeState
  it('should let the first staker make another node', async function () {
    await tryAdvanceChain(minimumAssertionPeriod)
    const block = await ethers.provider.getBlock('latest')
    const challengedAssertion = makeSimpleAssertion(
      prevNodeState,
      (block.number - prevNodeState.proposedBlock + 1) *
        arbGasSpeedLimitPerBlock
    )
    const { node } = await rollup.stakeOnNewNode(block, challengedAssertion, 3)
    challengedNode = node
    validNodeState = node.afterNodeState()
  })

  let challengerNode: Node
  it('should let the second staker make a conflicting node', async function () {
    await tryAdvanceChain(minimumAssertionPeriod)
    const block = await ethers.provider.getBlock('latest')
    const assertion = makeSimpleAssertion(
      prevNodeState,
      (block.number - prevNodeState.proposedBlock + 1) *
        arbGasSpeedLimitPerBlock
    )
    const { node } = await rollup
      .connect(accounts[1])
      .stakeOnNewNode(block, assertion, 4)
    challengerNode = node
  })

  it('should fail to confirm first staker node', async function () {
    await tryAdvanceChain(
      confirmationPeriodBlocks +
        challengedNode.assertion.checkTime(arbGasSpeedLimitPerBlock)
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
      challengedNode,
      challengerNode
    )
    const receipt = await (await tx).wait()
    const ev = rollup.rollup.interface.parseLog(
      receipt.logs![receipt.logs!.length - 1]
    )
    expect(ev.name).to.equal('RollupChallengeStarted')
    const parsedEv = (ev as any) as { args: { challengeContract: string } }
    const Challenge = await ethers.getContractFactory('Challenge')
    challenge = Challenge.attach(parsedEv.args.challengeContract) as Challenge
  })

  it('should make a new node', async function () {
    const block = await ethers.provider.getBlock('latest')
    const challengedAssertion = makeSimpleAssertion(
      validNodeState,
      (block.number - validNodeState.proposedBlock + 1) *
        arbGasSpeedLimitPerBlock
    )
    const { node } = await rollup.stakeOnNewNode(block, challengedAssertion, 5)
    challengedNode = node
  })

  it('new staker should make a conflicting node', async function () {
    const block = await ethers.provider.getBlock('latest')
    const assertion = makeSimpleAssertion(
      validNodeState,
      (block.number - validNodeState.proposedBlock + 10) *
        arbGasSpeedLimitPerBlock
    )
    const stake = await rollup.currentRequiredStake()
    await rollup.connect(accounts[2]).newStake(0, { value: stake })

    await rollup.connect(accounts[2]).stakeOnExistingNode(block, 3)

    const { node } = await rollup
      .connect(accounts[2])
      .stakeOnNewNode(block, assertion, 6)
    challengerNode = node
  })

  it('asserter should win via timeout', async function () {
    await tryAdvanceChain(
      confirmationPeriodBlocks +
        challengedNode.assertion.checkTime(arbGasSpeedLimitPerBlock) +
        1
    )
    await challenge.timeout()
  })

  it('confirm first staker node', async function () {
    await rollup.confirmNextNode(zerobytes32, [])
  })

  it('should reject out of order second node', async function () {
    await rollup.rejectNextNode(0, stakeToken)
  })

  it('should initiate another challenge', async function () {
    const tx = rollup.createChallenge(
      await accounts[0].getAddress(),
      5,
      await accounts[2].getAddress(),
      6,
      challengedNode,
      challengerNode
    )
    const receipt = await (await tx).wait()
    const ev = rollup.rollup.interface.parseLog(
      receipt.logs![receipt.logs!.length - 1]
    )
    expect(ev.name).to.equal('RollupChallengeStarted')
    const parsedEv = (ev as any) as { args: { challengeContract: string } }
    const Challenge = await ethers.getContractFactory('Challenge')
    challenge = Challenge.attach(parsedEv.args.challengeContract) as Challenge
  })

  it('challenger should reply in challenge', async function () {
    const chunks = Array(401).fill(zerobytes32)
    chunks[0] = challengedNode.assertion.startAssertionHash()

    await challenge
      .connect(accounts[2])
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
  })

  it('challenger should win via timeout', async function () {
    await tryAdvanceChain(
      confirmationPeriodBlocks +
        challengedNode.assertion.checkTime(arbGasSpeedLimitPerBlock) +
        1
    )
    await challenge.timeout()
  })

  it('should reject out of order second node', async function () {
    await rollup.rejectNextNode(6, await accounts[2].getAddress())
  })

  it('confirm next node', async function () {
    await tryAdvanceChain(confirmationPeriodBlocks)
    await rollup.confirmNextNode(zerobytes32, [])
  })

  it('can add stake', async function () {
    await rollup
      .connect(accounts[2])
      .addToDeposit(await accounts[2].getAddress(), 0, { value: 5 })
  })

  it('can reduce stake', async function () {
    await rollup.connect(accounts[2]).reduceDeposit(5)
  })

  it('returns stake to staker', async function () {
    await rollup.returnOldDeposit(await accounts[2].getAddress())
  })
})
