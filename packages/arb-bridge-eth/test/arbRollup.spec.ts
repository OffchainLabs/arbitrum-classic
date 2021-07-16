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
import { Signer, BigNumberish, Contract, BytesLike } from 'ethers'
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
  nodeHash,
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
let rollupAdmin: Contract
let challenge: Challenge
// let rollupTester: RollupTester
// let assertionInfo: Assertion
let accounts: Signer[]

async function createRollup(): Promise<{
  rollupCon: Rollup
  blockCreated: number
}> {
  const rollupConfig: [
    BytesLike,
    BigNumberish,
    BigNumberish,
    BigNumberish,
    BigNumberish,
    string,
    string,
    string,
    BigNumberish,
    BigNumberish,
    BytesLike
  ] = [
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
    '0x',
  ]

  let receipt
  let rollupCreator

  if (process.env['ROLLUP_DEBUG'] === '1') {
    // this deploys the rollup contracts without proxies to facilitate debugging
    const ChallengeFactory = await deployments.get('ChallengeFactory')
    const RollupCreatorNoProxy = (await ethers.getContractFactory(
      'RollupCreatorNoProxy'
    )) as RollupCreatorNoProxy__factory
    rollupCreator = await RollupCreatorNoProxy.deploy(
      ChallengeFactory.address,
      ...rollupConfig
    )
    receipt = await rollupCreator.deployTransaction.wait()
  } else {
    rollupCreator = await ethers.getContractAt(
      'RollupCreator',
      (await deployments.get('RollupCreator')).address
    )
    const createRollupTx = await rollupCreator.createRollup(...rollupConfig)
    receipt = await createRollupTx.wait()
  }

  if (!receipt.logs) {
    throw Error('expected receipt to have logs')
  }

  const ev = rollupCreator.interface.parseLog(
    receipt.logs[receipt.logs.length - 1]
  )
  expect(ev.name).to.equal('RollupCreated')
  const parsedEv = (ev as any) as { args: { rollupAddress: string } }

  const Rollup = (await ethers.getContractFactory('RollupUserFacet')).connect(
    accounts[8]
  )
  const RollupAdmin = (
    await ethers.getContractFactory('RollupAdminFacet')
  ).connect(accounts[0])

  rollupAdmin = RollupAdmin.attach(parsedEv.args.rollupAddress)
  await rollupAdmin.setValidator(
    [
      await accounts[1].getAddress(),
      await accounts[2].getAddress(),
      await accounts[8].getAddress(),
    ],
    [true, true, true]
  )

  const rollupCon = Rollup.attach(parsedEv.args.rollupAddress) as Rollup

  return {
    rollupCon: rollupCon,
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
      await accounts[8].getAddress(),
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
      await accounts[8].getAddress(),
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

  it('should pause the contracts then resume', async function () {
    const prevIsPaused = await rollup.rollup.paused()
    expect(prevIsPaused).to.equal(false)

    await rollupAdmin.pause()

    const postIsPaused = await rollup.rollup.paused()
    expect(postIsPaused).to.equal(true)

    await expect(
      rollup
        .connect(accounts[2])
        .addToDeposit(await accounts[2].getAddress(), { value: 5 })
    ).to.be.revertedWith('Pausable: paused')

    await rollupAdmin.resume()
  })

  it('should allow admin to truncate nodes', async function () {
    const prevLatestConfirmed = await rollup.rollup.latestConfirmed()
    expect(prevLatestConfirmed.toNumber()).to.equal(6)
    // prevNode is prevLatestConfirmed
    prevNode = challengerNode

    const stake = await rollup.currentRequiredStake()

    await rollup.newStake({ value: stake })
    const { node: node1 } = await makeSimpleNode(rollup, prevNode)
    const node1Num = await rollup.rollup.latestNodeCreated()

    await tryAdvanceChain(minimumAssertionPeriod)

    await rollup.connect(accounts[1]).newStake({ value: stake })
    const { node: node2 } = await makeSimpleNode(
      rollup.connect(accounts[1]),
      prevNode,
      node1
    )
    const node2Num = await rollup.rollup.latestNodeCreated()

    const tx = await rollup.createChallenge(
      await accounts[8].getAddress(),
      node1Num,
      await accounts[1].getAddress(),
      node2Num,
      node1,
      node2
    )
    const receipt = await tx.wait()
    const ev = rollup.rollup.interface.parseLog(
      receipt.logs![receipt.logs!.length - 1]
    )
    expect(ev.name).to.equal('RollupChallengeStarted')
    const parsedEv = (ev as any) as { args: { challengeContract: string } }
    const Challenge = await ethers.getContractFactory('Challenge')
    challenge = Challenge.attach(parsedEv.args.challengeContract) as Challenge

    const preCode = await ethers.provider.getCode(challenge.address)
    expect(preCode).to.not.equal('0x')

    await expect(
      rollupAdmin.forceResolveChallenge(
        [await accounts[8].getAddress()],
        [await accounts[1].getAddress()]
      )
    ).to.be.revertedWith('Pausable: not paused')

    await expect(
      rollup.createChallenge(
        await accounts[8].getAddress(),
        node1Num,
        await accounts[1].getAddress(),
        node2Num,
        node1,
        node2
      )
    ).to.be.revertedWith('IN_CHAL')

    await rollupAdmin.pause()

    await rollupAdmin.forceResolveChallenge(
      [await accounts[8].getAddress()],
      [await accounts[1].getAddress()]
    )

    // challenge should have been destroyed
    const postCode = await ethers.provider.getCode(challenge.address)
    expect(postCode).to.equal('0x')

    const challengeA = await rollupAdmin.currentChallenge(
      await accounts[8].getAddress()
    )
    const challengeB = await rollupAdmin.currentChallenge(
      await accounts[1].getAddress()
    )
    const ZERO_ADDR = '0x0000000000000000000000000000000000000000'

    expect(challengeA).to.equal(ZERO_ADDR)
    expect(challengeB).to.equal(ZERO_ADDR)

    await rollupAdmin.forceRefundStaker([
      await accounts[8].getAddress(),
      await accounts[1].getAddress(),
    ])

    const block = await ethers.provider.getBlock('latest')
    const assertion = makeSimpleAssertion(
      prevNode.afterState,
      (block.number - prevNode.afterState.proposedBlock + 1) *
        arbGasSpeedLimitPerBlock
    )

    const hasSibling = true
    const newNodeHash = async () =>
      nodeHash(
        hasSibling,
        await rollup.rollup.getNodeHash(
          await rollup.rollup.latestNodeCreated()
        ),
        assertion.executionHash(),
        zerobytes32
      )

    await rollupAdmin.forceCreateNode(
      await newNodeHash(),
      assertion.bytes32Fields(),
      assertion.intFields(),
      prevNode.afterState.proposedBlock,
      prevNode.afterState.inboxMaxCount,
      prevLatestConfirmed,
      1,
      zerobytes32
    )
    const adminNodeNum = await rollup.rollup.latestNodeCreated()
    const midLatestConfirmed = await rollup.rollup.latestConfirmed()
    expect(midLatestConfirmed.toNumber()).to.equal(6)

    expect(adminNodeNum.toNumber()).to.equal(node2Num.toNumber() + 1)

    await rollupAdmin.forceCreateNode(
      await newNodeHash(),
      assertion.bytes32Fields(),
      assertion.intFields(),
      prevNode.afterState.proposedBlock,
      prevNode.afterState.inboxMaxCount,
      prevLatestConfirmed,
      1,
      zerobytes32
    )
    const postLatestCreated = await rollup.rollup.latestNodeCreated()

    const sends: Array<BytesLike> = []
    const messageData = ethers.utils.concat(sends)
    const messageLengths = sends.map(msg => msg.length)

    await rollupAdmin.forceConfirmNode(
      adminNodeNum,
      zerobytes32,
      messageData,
      messageLengths,
      0,
      zerobytes32,
      0
    )

    const postLatestConfirmed = await rollup.rollup.latestConfirmed()
    expect(postLatestCreated).to.equal(adminNodeNum.add(1))
    expect(postLatestConfirmed).to.equal(adminNodeNum)

    await rollupAdmin.resume()
  })
})
