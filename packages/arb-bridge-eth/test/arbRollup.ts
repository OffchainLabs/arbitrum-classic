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
import { ethers, deployments, run } from 'hardhat'
import { Signer, BigNumberish } from 'ethers'
import { ContractTransaction } from '@ethersproject/contracts'
import { TransactionResponse } from '@ethersproject/providers'
import { assert, expect } from 'chai'
import { Rollup } from '../build/types/Rollup'
import { Node as NodeCon } from '../build/types/Node'
import { RollupCreatorNoProxy } from '../build/types/RollupCreatorNoProxy'
import { RollupCreatorNoProxy__factory } from '../build/types/factories/RollupCreatorNoProxy__factory'
import { Challenge } from '../build/types/Challenge'
// import { RollupTester } from '../build/types/RollupTester'
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
const stakeRequirement = 10
const stakeToken = '0x0000000000000000000000000000000000000000'
const confirmationPeriodBlocks = 100
const arbGasSpeedLimitPerBlock = 1000000
const minimumAssertionPeriod = 75
const sequencerDelayBlocks = 15
const sequencerDelaySeconds = 900

let rollup: RollupContract
let challenge: Challenge
// let rollupTester: RollupTester
// let assertionInfo: Assertion
let accounts: Signer[]

async function createRollup(): Promise<{
  rollupCon: Rollup
  blockCreated: number
}> {
  const ChallengeFactory = await deployments.get('ChallengeFactory')
  const RollupCreatorNoProxy = (await ethers.getContractFactory(
    'RollupCreatorNoProxy'
  )) as RollupCreatorNoProxy__factory
  const rollupCreator = await RollupCreatorNoProxy.deploy(
    ChallengeFactory.address,
    initialVmState,
    confirmationPeriodBlocks,
    0,
    arbGasSpeedLimitPerBlock,
    stakeRequirement,
    stakeToken,
    await accounts[0].getAddress(), // owner
    await accounts[1].getAddress(), // sequencer
    sequencerDelayBlocks,
    sequencerDelaySeconds,
    '0x'
  )

  const receipt = await (rollupCreator.deployTransaction as TransactionResponse).wait()
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

async function makeSimpleNode(
  rollup: RollupContract,
  parentNode: Node,
  prevNode?: Node
): Promise<{ tx: ContractTransaction; node: Node }> {
  const block = await ethers.provider.getBlock('latest')
  const challengedAssertion = makeSimpleAssertion(
    parentNode.afterState,
    (block.number - parentNode.afterState.proposedBlock + 1) *
      arbGasSpeedLimitPerBlock
  )
  const { tx, node, event } = await rollup.stakeOnNewNode(
    parentNode,
    challengedAssertion,
    zerobytes32,
    '0x',
    prevNode
  )
  assert.equal(event.nodeHash, node.nodeHash)
  assert.equal(event.executionHash, node.executionHash())
  return { tx, node }
}

let prevNode: Node

describe('ArbRollup', () => {
  it('should deploy contracts', async function () {
    accounts = await initializeAccounts()

    await run('deploy', { tags: 'test' })
  })

  it('should initialize', async function () {
    const { rollupCon, blockCreated } = await createRollup()
    rollup = new RollupContract(rollupCon)
    const originalNode = await rollup.latestConfirmed()
    const nodeAddress = await rollup.getNode(originalNode)

    const NodeContract = await ethers.getContractFactory('Node')
    const node = NodeContract.attach(nodeAddress) as NodeCon

    const newState = new NodeState(
      new ExecutionState(0, initialVmState, 0, 0, 0, zerobytes32, zerobytes32),
      blockCreated,
      1
    )

    const initialExecState = new ExecutionState(
      0,
      initialVmState,
      0,
      0,
      0,
      zerobytes32,
      zerobytes32
    )
    const initialNodeState = new NodeState(initialExecState, blockCreated, 1)
    const initialAssertion = new Assertion(
      initialNodeState,
      0,
      initialVmState,
      [],
      [],
      []
    )
    prevNode = new Node(initialAssertion, blockCreated, 1, zerobytes32)

    assert.equal(
      await node.stateHash(),
      prevNode.afterState.hash(),
      'initial confirmed node should have set initial state'
    )
  })

  it('should place stake', async function () {
    const stake = await rollup.currentRequiredStake()
    await rollup.newStake({ value: stake })
  })

  it('should place stake on new node', async function () {
    await tryAdvanceChain(minimumAssertionPeriod)
    const { node } = await makeSimpleNode(rollup, prevNode)
    prevNode = node
  })

  it('should let a new staker place on existing node', async function () {
    await rollup.connect(accounts[1]).newStake({ value: 10 })

    await rollup.connect(accounts[1]).stakeOnExistingNode(1, prevNode.nodeHash)
  })

  it('should move stake to a new node', async function () {
    await tryAdvanceChain(minimumAssertionPeriod)
    const { node } = await makeSimpleNode(rollup, prevNode)
    prevNode = node
  })

  it('should let the second staker place on the new node', async function () {
    await rollup.connect(accounts[1]).stakeOnExistingNode(2, prevNode.nodeHash)
  })

  it('should confirm node', async function () {
    await tryAdvanceChain(confirmationPeriodBlocks * 2)
    await rollup.confirmNextNode(zerobytes32, 0, [], zerobytes32, 0)
  })

  it('should confirm next node', async function () {
    await tryAdvanceChain(minimumAssertionPeriod)
    await rollup.confirmNextNode(zerobytes32, 0, [], zerobytes32, 0)
  })

  let challengedNode: Node
  let validNode: Node
  it('should let the first staker make another node', async function () {
    await tryAdvanceChain(minimumAssertionPeriod)
    const { node } = await makeSimpleNode(rollup, prevNode)
    challengedNode = node
    validNode = node
  })

  let challengerNode: Node
  it('should let the second staker make a conflicting node', async function () {
    await tryAdvanceChain(minimumAssertionPeriod)
    const { node } = await makeSimpleNode(
      rollup.connect(accounts[1]),
      prevNode,
      validNode
    )
    challengerNode = node
  })

  it('should fail to confirm first staker node', async function () {
    await tryAdvanceChain(
      confirmationPeriodBlocks +
        challengedNode.checkTime(arbGasSpeedLimitPerBlock)
    )
    await expect(
      rollup.confirmNextNode(zerobytes32, 0, [], zerobytes32, 0)
    ).to.be.revertedWith('NOT_ALL_STAKED')
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
    const { node } = await makeSimpleNode(rollup, validNode)
    challengedNode = node
  })

  it('new staker should make a conflicting node', async function () {
    const stake = await rollup.currentRequiredStake()
    await rollup.connect(accounts[2]).newStake({ value: stake })

    await rollup.connect(accounts[2]).stakeOnExistingNode(3, validNode.nodeHash)

    const { node } = await makeSimpleNode(
      rollup.connect(accounts[2]),
      validNode,
      challengedNode
    )
    challengerNode = node
  })

  it('asserter should win via timeout', async function () {
    await tryAdvanceChain(
      confirmationPeriodBlocks +
        challengedNode.checkTime(arbGasSpeedLimitPerBlock) +
        1
    )
    await challenge.timeout()
  })

  it('confirm first staker node', async function () {
    await rollup.confirmNextNode(zerobytes32, 0, [], zerobytes32, 0)
  })

  it('should reject out of order second node', async function () {
    await rollup.rejectNextNode(stakeToken)
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
    chunks[0] = challengedNode.beforeState.execState.challengeHash()

    await challenge
      .connect(accounts[2])
      .bisectExecution(
        [],
        0,
        challengedNode.beforeState.execState.gasUsed,
        ethers.BigNumber.from(challengedNode.afterState.execState.gasUsed).sub(
          challengedNode.beforeState.execState.gasUsed
        ),
        challengedNode.afterState.execState.challengeHash(),
        challengedNode.beforeState.execState.gasUsed,
        challengedNode.beforeState.execState.challengeRestHash(),
        chunks
      )
  })

  it('challenger should win via timeout', async function () {
    await tryAdvanceChain(
      confirmationPeriodBlocks +
        challengedNode.checkTime(arbGasSpeedLimitPerBlock) +
        1
    )
    await challenge.timeout()
  })

  it('should reject out of order second node', async function () {
    await rollup.rejectNextNode(await accounts[2].getAddress())
  })

  it('confirm next node', async function () {
    await tryAdvanceChain(confirmationPeriodBlocks)
    await rollup.confirmNextNode(zerobytes32, 0, [], zerobytes32, 0)
  })

  it('can add stake', async function () {
    await rollup
      .connect(accounts[2])
      .addToDeposit(await accounts[2].getAddress(), { value: 5 })
  })

  it('can reduce stake', async function () {
    await rollup.connect(accounts[2]).reduceDeposit(5)
  })

  it('returns stake to staker', async function () {
    await rollup.returnOldDeposit(await accounts[2].getAddress())
  })
})
