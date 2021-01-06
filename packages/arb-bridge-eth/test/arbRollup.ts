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
import bre from '@nomiclabs/buidler'
import { Signer, ContractTransaction, providers, utils } from 'ethers'
import * as chai from 'chai'
import chaiAsPromised from 'chai-as-promised'
import { Rollup } from '../build/types/Rollup'
import { Node } from '../build/types/Node'
import { RollupCreator } from '../build/types/RollupCreator'
import { Challenge } from '../build/types/Challenge'
// import { RollupTester } from '../build/types/RollupTester'
import { ArbValue } from 'arb-provider-ethers'
import deploy_contracts from '../scripts/deploy'
import { initializeAccounts } from './utils'

chai.use(chaiAsPromised)

const { ethers } = bre
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
let rollup: Rollup
let challenge: Challenge
// let rollupTester: RollupTester
// let assertionInfo: Assertion
let originalNode: string
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

  const logs = receipt.logs.map((log: providers.Log) =>
    rollupCreator.interface.parseLog(log)
  )
  const ev = logs[logs.length - 1]
  expect(ev.name).to.equal('RollupCreated')
  const chainAddress = ev.values.rollupAddress
  const Rollup = await ethers.getContractFactory('Rollup')
  return {
    rollupCon: Rollup.attach(chainAddress) as Rollup,
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

class NodeState {
  constructor(
    public proposedBlock: utils.BigNumberish,
    public stepsRun: utils.BigNumberish,
    public machineHash: string,
    public inboxTop: string,
    public inboxCount: utils.BigNumberish,
    public sendCount: utils.BigNumberish,
    public logCount: utils.BigNumberish,
    public inboxMaxCount: utils.BigNumberish
  ) {}

  hash(): string {
    return ethers.utils.solidityKeccak256(
      [
        'uint256',
        'uint256',
        'bytes32',
        'bytes32',
        'uint256',
        'uint256',
        'uint256',
        'uint256',
      ],
      [
        this.proposedBlock,
        this.stepsRun,
        this.machineHash,
        this.inboxTop,
        this.inboxCount,
        this.sendCount,
        this.logCount,
        this.inboxMaxCount,
      ]
    )
  }
}

function buildAccumulator(base: string, hashes: string[]): string {
  let acc = base
  for (const h of hashes) {
    acc = ethers.utils.solidityKeccak256(['bytes32', 'bytes32'], [acc, h])
  }
  return acc
}

class Assertion {
  public inboxDelta: string
  public inboxMessagesRead: utils.BigNumberish
  public sendAcc: string
  public sendCount: utils.BigNumberish
  public logAcc: string
  public logCount: utils.BigNumberish
  public afterInboxHash: string

  constructor(
    public prevNodeState: NodeState,
    public stepsExecuted: utils.BigNumberish,
    public gasUsed: utils.BigNumberish,
    public afterMachineHash: string,
    messages: string[],
    sends: string[],
    logs: string[]
  ) {
    this.inboxDelta = buildAccumulator(zerobytes32, messages.reverse())
    this.inboxMessagesRead = messages.length
    this.afterInboxHash = buildAccumulator(prevNodeState.inboxTop, messages)

    this.sendAcc = buildAccumulator(zerobytes32, sends)
    this.sendCount = sends.length

    this.logAcc = buildAccumulator(zerobytes32, logs)
    this.logCount = logs.length
  }

  bytes32Fields(): string[] {
    return [
      this.prevNodeState.machineHash,
      this.prevNodeState.inboxTop,
      this.inboxDelta,
      this.sendAcc,
      this.logAcc,
      this.afterInboxHash,
      this.afterMachineHash,
    ]
  }

  intFields(): utils.BigNumberish[] {
    return [
      this.prevNodeState.proposedBlock,
      this.prevNodeState.stepsRun,
      this.prevNodeState.inboxCount,
      this.prevNodeState.sendCount,
      this.prevNodeState.logCount,
      this.prevNodeState.inboxMaxCount,
      this.stepsExecuted,
      this.inboxMessagesRead,
      this.gasUsed,
      this.sendCount,
      this.logCount,
    ]
  }
}

let initialNodeState: NodeState

describe('ArbRollup', () => {
  it('should deploy contracts', async function () {
    accounts = await initializeAccounts()
    const { RollupCreator } = await deploy_contracts(bre)
    rollupCreator = RollupCreator as RollupCreator

    // const RollupTester = await ethers.getContractFactory('RollupTester')
    // rollupTester = (await RollupTester.deploy()) as RollupTester
    // await rollupTester.deployed()
  })

  // it('should not be able to shut down the template', async () => {
  //   const template = await arbFactory.rollupTemplate()
  //   const ArbRollup = await ethers.getContractFactory('ArbRollup')
  //   const templateRollup = ArbRollup.attach(template) as ArbRollup
  //   await templateRollup.init(
  //     initialVmState, // vmState
  //     gracePeriodTicks, // gracePeriodTicks
  //     1000000, // arbGasSpeedLimitPerTick
  //     maxExecutionSteps, // maxExecutionSteps
  //     stakeRequirement, // stakeRequirement
  //     stakeToken,
  //     await accounts[0].getAddress(), // owner
  //     await arbFactory.challengeFactoryAddress(),
  //     await arbFactory.globalInboxAddress(),
  //     '0x'
  //   )
  //   await expect(templateRollup.owner()).to.eventually.equal(
  //     await accounts[0].getAddress()
  //   )
  //   await expect(templateRollup.ownerShutdown()).to.be.revertedWith('NOT_CLONE')
  // })

  // it('should be able to shut down a clone', async () => {
  //   const rollup = await createRollup()
  //   await expect(rollup.owner()).to.eventually.equal(
  //     await accounts[0].getAddress()
  //   )
  //   await rollup.ownerShutdown()
  // })

  it('should initialize', async function () {
    const { rollupCon, blockCreated } = await createRollup()
    rollup = rollupCon
    originalNode = await rollup.latestConfirmed()
    const nodeAddress = await rollup.nodes(originalNode)

    const Node = await ethers.getContractFactory('Node')
    const node = Node.attach(nodeAddress) as Node

    initialNodeState = new NodeState(
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
      initialNodeState.hash(),
      'initial confirmed node should have set initial state'
    )
  })

  it('it should place stake on new node', async function () {
    const block = await ethers.provider.getBlock('latest')
    await tryAdvanceChain(challengePeriodBlocks / 10)
    const assertion = new Assertion(
      initialNodeState,
      100,
      100,
      zerobytes32,
      [],
      [],
      []
    )
    await rollup.newStakeOnNewNode(
      block.hash,
      block.number,
      1, // new node num
      0, // prev node num
      assertion.bytes32Fields(),
      assertion.intFields(),
      { value: 10 }
    )
  })

  it('it should confirm node', async function () {
    await tryAdvanceChain(challengePeriodBlocks)
    await rollup.confirmNextNode(zerobytes32, '0x', [])
  })

  // it('should fail to assert on invalid leaf', async () => {
  //   await expect(
  //     makeEmptyAssertion(
  //       arbRollup,
  //       '0x3400000000000000000000000000000000000000000000000000000000000000',
  //       0,
  //       0
  //     )
  //   ).to.be.revertedWith('MAKE_LEAF')
  // })

  // it("should fail to assert on halted vm", async () => {
  //   truffleAssert.reverts(makeEmptyAssertion("0x00", 0, 0), "MAKE_RUN");
  // })

  // it('should fail to assert over step limit', async () => {
  //   await expect(
  //     makeEmptyAssertion(arbRollup, initialVmState, maxExecutionSteps + 1, 0)
  //   ).to.be.revertedWith('MAKE_STEP')
  // })

  // it('should fail to assert without stake', async () => {
  //   await expect(
  //     makeEmptyAssertion(arbRollup, initialVmState, 0, 0)
  //   ).to.be.revertedWith('INV_STAKER')
  // })

  // it('should fail if reading past lastest inbox message', async () => {
  //   await expect(
  //     makeEmptyAssertion(arbRollup, initialVmState, 0, 10)
  //   ).to.be.revertedWith('MAKE_MESSAGE_CNT')
  // })

  // it('should create a stake', async () => {
  //   await expect(arbRollup.isStaked(await accounts[0].getAddress())).to
  //     .eventually.be.false
  //   await expect(
  //     arbRollup.connect(accounts[0]).placeStake([], [], {
  //       value: stakeRequirement,
  //     })
  //   ).to.emit(arbRollup, 'RollupStakeCreated')
  //   await expect(arbRollup.isStaked(await accounts[0].getAddress())).to
  //     .eventually.be.true
  // })

  // it('should create a second stake', async () => {
  //   await expect(
  //     arbRollup.connect(accounts[1]).placeStake([], [], {
  //       value: stakeRequirement,
  //     })
  //   ).to.emit(arbRollup, 'RollupStakeCreated')
  // })

  // it('should make an assertion', async () => {
  //   assert.isTrue(
  //     await arbRollup.isValidLeaf(originalNode),
  //     'latest confirmed should be leaf before asserting'
  //   )
  //   const prevProtoData = new VMProtoData(initialVmState, zerobytes32, 0, 0, 0)
  //   const params = new AssertionParams(0, ethers.utils.bigNumberify(0))
  //   const execAssertion = new ExecutionAssertion(
  //     '0x8500000000000000000000000000000000000000000000000000000000000000',
  //     zerobytes32,
  //     0,
  //     [],
  //     []
  //   )
  //   const block = await ethers.provider.getBlock('latest')
  //   const info = await makeAssertion(
  //     arbRollup,
  //     zerobytes32,
  //     prevProtoData,
  //     0,
  //     zerobytes32,
  //     0,
  //     params,
  //     execAssertion,
  //     [],
  //     block.hash,
  //     block.number
  //   )

  //   assertionInfo = info.assertion

  //   assert.isFalse(
  //     await arbRollup.isValidLeaf(assertionInfo.prevNodeHash()),
  //     'originalNode confirmed should be removed as leaf'
  //   )
  //   assert.isTrue(
  //     await arbRollup.isValidLeaf(assertionInfo.invalidInboxTopHash()),
  //     'invalid inbox top should be leaf'
  //   )
  //   // TODO: Check whether invalid execution is leaf
  //   assert.isTrue(
  //     await arbRollup.isValidLeaf(assertionInfo.validHash()),
  //     'valid child should be leaf'
  //   )
  // })

  // it('should allow the second staker to move to conflicting node', async () => {
  //   await expect(
  //     arbRollup
  //       .connect(accounts[1])
  //       .moveStake([assertionInfo.invalidInboxTopHashInner()], [])
  //   )
  //     .to.emit(arbRollup, 'RollupStakeMoved')
  //     .withArgs(
  //       await accounts[1].getAddress(),
  //       assertionInfo.invalidInboxTopHash()
  //     )
  // })

  // it('should allow the creation of a challenge', async () => {
  //   const txPromise = arbRollup.startChallenge(
  //     await accounts[0].getAddress(),
  //     await accounts[1].getAddress(),
  //     assertionInfo.prevNodeHash(),
  //     assertionInfo.deadline(),
  //     [2, 0],
  //     [
  //       assertionInfo.updatedProtoData().hash(),
  //       assertionInfo.prevProtoData.hash(),
  //     ],
  //     [],
  //     [],
  //     assertionInfo.validDataHash(),
  //     assertionInfo.invalidInboxTopDataHash(),
  //     assertionInfo.challengePeriod()
  //   )
  //   const receipt = await (await txPromise).wait()
  //   if (receipt.logs === undefined) {
  //     throw Error('logs must be defined')
  //   }
  //   expect(receipt.logs).to.have.lengthOf(2)
  //   const logs = receipt.logs.map((log: providers.Log) =>
  //     arbRollup.interface.parseLog(log)
  //   )
  //   const ev = logs[1]
  //   expect(ev.name).equals('RollupChallengeStarted')
  //   const challengeContract = ev.values.challengeContract

  //   const InboxTopChallenge = await ethers.getContractFactory(
  //     'InboxTopChallenge'
  //   )
  //   challenge = InboxTopChallenge.attach(challengeContract) as InboxTopChallenge
  // })

  // it('should timeout the challenge', async () => {
  //   await tryAdvanceChain(3)
  //   await challenge.timeoutChallenge()
  // })

  // it('should confirm invalid inbox top node', async () => {
  //   await tryAdvanceChain(3)
  //   await expect(
  //     arbRollup.confirm(
  //       assertionInfo.prevProtoData.hash(),
  //       assertionInfo.prevProtoData.messageCount,
  //       [0],
  //       [assertionInfo.deadline()],
  //       [assertionInfo.invalidInboxTopChallengeHash()],
  //       [],
  //       [],
  //       [],
  //       '0x',
  //       [await accounts[1].getAddress()].sort(),
  //       [],
  //       [0, 0]
  //     )
  //   ).to.emit(arbRollup, 'RollupConfirmed')

  //   assert.equal(
  //     await arbRollup.latestConfirmed(),
  //     assertionInfo.invalidInboxTopHash(),
  //     'latest confirmed should now be invalid inbox child'
  //   )

  //   assert.isTrue(
  //     await arbRollup.isValidLeaf(assertionInfo.invalidInboxTopHash()),
  //     'invalid inbox top should be leaf'
  //   )
  // })

  // it('should prune a leaf', async () => {
  //   assert.isTrue(
  //     await arbRollup.isValidLeaf(assertionInfo.validHash()),
  //     'valid node should be leaf'
  //   )
  //   await expect(
  //     arbRollup.pruneLeaves(
  //       [originalNode],
  //       [assertionInfo.validHashInner()],
  //       [1],
  //       [assertionInfo.invalidInboxTopHashInner()],
  //       [1]
  //     )
  //   ).to.emit(arbRollup, 'RollupPruned')

  //   assert.isFalse(
  //     await arbRollup.isValidLeaf(assertionInfo.validHashInner()),
  //     'valid node should no longer be leaf'
  //   )
  // })

  // it('should assert again', async () => {
  //   const params = new AssertionParams(0, ethers.utils.bigNumberify(0))
  //   const execAssertion = new ExecutionAssertion(
  //     zerobytes32,
  //     zerobytes32,
  //     0,
  //     [new ArbValue.TupleValue([new ArbValue.IntValue(10)])],
  //     []
  //   )

  //   const { prevLeaf, vmProtoHashBefore } = await computePrevLeaf(
  //     rollupTester,
  //     assertionInfo.prevNodeHash(),
  //     assertionInfo.prevProtoData,
  //     assertionInfo.deadline(),
  //     assertionInfo.invalidInboxTopChallengeHash(),
  //     0,
  //     params,
  //     execAssertion
  //   )

  //   expect(vmProtoHashBefore, 'wrong vmProtoHashBefore').to.equal(
  //     assertionInfo.prevProtoData.hash()
  //   )
  //   expect(prevLeaf, 'wrong prevLeaf').to.equal(
  //     assertionInfo.invalidInboxTopHash()
  //   )

  //   assert.isTrue(
  //     await arbRollup.isValidLeaf(prevLeaf),
  //     'invalid inbox node should be leaf'
  //   )

  //   const block = await ethers.provider.getBlock('latest')
  //   assertionInfo = (
  //     await makeAssertion(
  //       arbRollup.connect(accounts[1]),
  //       assertionInfo.prevNodeHash(),
  //       assertionInfo.prevProtoData,
  //       assertionInfo.deadline(),
  //       assertionInfo.invalidInboxTopChallengeHash(),
  //       0,
  //       params,
  //       execAssertion,
  //       [],
  //       block.hash,
  //       block.number
  //     )
  //   ).assertion
  // })

  // it('should confirm valid node', async () => {
  //   await tryAdvanceChain(3)
  //   const { validNodeHashes, lastNodeHash } = await rollupTester.confirm(
  //     await arbRollup.latestConfirmed(),
  //     assertionInfo.prevProtoData.hash(),
  //     assertionInfo.prevProtoData.messageCount,
  //     [2],
  //     [assertionInfo.deadline()],
  //     [],
  //     [assertionInfo.assertion.outLogsAcc()],
  //     [assertionInfo.updatedProtoData().hash()],
  //     [assertionInfo.assertion.outMessages.length],
  //     assertionInfo.assertion.outMessagesData()
  //   )

  //   expect(validNodeHashes.length).to.equal(1)
  //   expect(validNodeHashes[0]).to.equal(assertionInfo.validHash())

  //   assert.equal(
  //     lastNodeHash,
  //     assertionInfo.validHash(),
  //     'calculated last node should be the valid node'
  //   )

  //   await expect(
  //     arbRollup.confirm(
  //       assertionInfo.prevProtoData.hash(),
  //       assertionInfo.prevProtoData.messageCount,
  //       [2],
  //       [assertionInfo.deadline()],
  //       [],
  //       [assertionInfo.assertion.outLogsAcc()],
  //       [assertionInfo.updatedProtoData().hash()],
  //       [assertionInfo.assertion.outMessages.length],
  //       assertionInfo.assertion.outMessagesData(),
  //       [await accounts[1].getAddress()].sort(),
  //       [],
  //       [0, 0]
  //     )
  //   ).to.emit(arbRollup, 'RollupConfirmed')

  //   assert.equal(
  //     await arbRollup.latestConfirmed(),
  //     assertionInfo.validHash(),
  //     'latest confirmed should now be valid child'
  //   )

  //   assert.isTrue(
  //     await arbRollup.isValidLeaf(assertionInfo.validHash()),
  //     'valid child should be leaf'
  //   )
  // })

  // it('should allow second staker to withdraw', async () => {
  //   await expect(arbRollup.connect(accounts[1]).recoverStakeConfirmed([]))
  //     .to.emit(arbRollup, 'RollupStakeRefunded')
  //     .withArgs(await accounts[1].getAddress())
  // })
})
