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
import { ArbRollup } from '../build/types/ArbRollup'
import { ArbFactory } from '../build/types/ArbFactory'
import { RollupTester } from '../build/types/RollupTester'
import { InboxTopChallenge } from '../build/types/InboxTopChallenge'
import { ArbValue } from 'arb-provider-ethers'
import deploy_contracts from '../scripts/deploy'

chai.use(chaiAsPromised)

const { ethers } = bre
const { assert, expect } = chai

const initialVmState =
  '0x9900000000000000000000000000000000000000000000000000000000000000'
const stakeRequirement = 10
const maxExecutionSteps = 50000
const gracePeriodTicks = ethers.utils.bigNumberify(1000)

function inboxTopHash(
  lowerHash: string,
  topHash: string,
  chainLength: utils.BigNumberish
): string {
  return ethers.utils.solidityKeccak256(
    ['bytes32', 'bytes32', 'uint256'],
    [lowerHash, topHash, chainLength]
  )
}

function childNodeInnerHash(
  deadlineTicks: utils.BigNumberish,
  nodeDataHash: string,
  childType: number,
  vmProtoStateHash: string
): string {
  return ethers.utils.solidityKeccak256(
    ['bytes32', 'uint256', 'bytes32', 'uint256'],
    [vmProtoStateHash, deadlineTicks, nodeDataHash, childType]
  )
}

function childNodeHash(
  prevNodeHash: string,
  deadlineTicks: utils.BigNumberish,
  nodeDataHash: string,
  childType: number,
  vmProtoStateHash: string
): string {
  return ethers.utils.solidityKeccak256(
    ['bytes32', 'bytes32'],
    [
      prevNodeHash,
      childNodeInnerHash(
        deadlineTicks,
        nodeDataHash,
        childType,
        vmProtoStateHash
      ),
    ]
  )
}

function childNodeShortHash(
  prevNodeHash: string,
  nodeInnerHash: string
): string {
  return ethers.utils.solidityKeccak256(
    ['bytes32', 'bytes32'],
    [prevNodeHash, nodeInnerHash]
  )
}

const zerobytes32 =
  '0x0000000000000000000000000000000000000000000000000000000000000000'

async function makeEmptyAssertion(
  arbRollup: ArbRollup,
  vmState: string,
  numSteps: number,
  importedMessageCount: utils.BigNumberish
): Promise<ContractTransaction> {
  const block = await ethers.provider.getBlock('latest')
  return arbRollup.makeAssertion(
    [
      vmState,
      zerobytes32,
      zerobytes32,
      zerobytes32,
      zerobytes32,
      zerobytes32,
      zerobytes32,
      zerobytes32,
    ],
    [0, 0, importedMessageCount, 0, 0],
    block.hash,
    block.number,
    0,
    0,
    0,
    numSteps,
    0,
    []
  )
}

class VMProtoData {
  public inboxCount: utils.BigNumber
  public messageCount: utils.BigNumber
  public logCount: utils.BigNumber
  constructor(
    public machineHash: string,
    public inboxTop: string,
    inboxCount: utils.BigNumberish,
    messageCount: utils.BigNumberish,
    logCount: utils.BigNumberish
  ) {
    this.inboxCount = ethers.utils.bigNumberify(inboxCount)
    this.messageCount = ethers.utils.bigNumberify(messageCount)
    this.logCount = ethers.utils.bigNumberify(logCount)
  }

  hash(): string {
    return ethers.utils.solidityKeccak256(
      ['bytes32', 'bytes32', 'uint256', 'uint256', 'uint256'],
      [
        this.machineHash,
        this.inboxTop,
        this.inboxCount,
        this.messageCount,
        this.logCount,
      ]
    )
  }
}

class AssertionParams {
  constructor(
    public numSteps: number,
    public importedMessageCount: utils.BigNumberish
  ) {}
}

class ExecutionAssertion {
  constructor(
    public afterState: string,
    public numGas: number,
    public outMessages: ArbValue.Value[],
    public outLogs: ArbValue.Value[]
  ) {}

  outMessagesAcc(): string {
    return this.outMessages.reduce(
      (acc, val) =>
        ethers.utils.solidityKeccak256(
          ['bytes32', 'bytes32'],
          [acc, val.hash()]
        ),
      zerobytes32
    )
  }

  outLogsAcc(): string {
    return this.outLogs.reduce(
      (acc, val) =>
        ethers.utils.solidityKeccak256(
          ['bytes32', 'bytes32'],
          [acc, val.hash()]
        ),
      zerobytes32
    )
  }

  outMessagesData(): Uint8Array {
    return this.outMessages.reduce(
      (acc, val) => ethers.utils.concat([acc, ArbValue.marshal(val)]),
      new Uint8Array()
    )
  }
}

class AssertionClaim {
  constructor(
    public afterInboxTop: string,
    public executionAssertion: ExecutionAssertion
  ) {}
}

class Assertion {
  public blockNumber: utils.BigNumber
  public inboxCount: utils.BigNumber
  public prevDeadline: utils.BigNumber

  constructor(
    blockNumber: utils.BigNumberish,
    public inboxValue: string,
    inboxCount: utils.BigNumberish,
    public prevPrevNode: string,
    public prevProtoData: VMProtoData,
    prevDeadline: utils.BigNumberish,
    public prevDataHash: string,
    public prevChildType: number,
    public params: AssertionParams,
    public claims: AssertionClaim
  ) {
    this.blockNumber = ethers.utils.bigNumberify(blockNumber)
    this.inboxCount = ethers.utils.bigNumberify(inboxCount)
    this.prevDeadline = ethers.utils.bigNumberify(prevDeadline)
  }

  prevNodeHash(): string {
    return childNodeHash(
      this.prevPrevNode,
      this.prevDeadline,
      this.prevDataHash,
      this.prevChildType,
      this.prevProtoData.hash()
    )
  }

  deadline(): utils.BigNumber {
    let deadlineTicks = this.blockNumber
      .mul(ethers.utils.bigNumberify(1000))
      .add(gracePeriodTicks)
    if (deadlineTicks.lt(this.prevDeadline)) {
      deadlineTicks = this.prevDeadline
    }
    // Should be numArbGas / arbGasSpeedLimitPerTick, but numArbGas is 0 in the test
    return deadlineTicks.add(ethers.utils.bigNumberify(0))
  }

  challengePeriod(): utils.BigNumber {
    // should be plus numArbGas
    return gracePeriodTicks.add(ethers.utils.bigNumberify(1000))
  }

  invalidInboxTopDataHash(): string {
    return inboxTopHash(
      this.claims.afterInboxTop,
      this.inboxValue,
      this.inboxCount.sub(
        this.prevProtoData.inboxCount.add(this.params.importedMessageCount)
      )
    )
  }

  invalidInboxTopChallengeHash(): string {
    return ethers.utils.solidityKeccak256(
      ['bytes32', 'uint256'],
      [this.invalidInboxTopDataHash(), this.challengePeriod()]
    )
  }

  invalidInboxTopHashInner(): string {
    return childNodeInnerHash(
      this.deadline(),
      this.invalidInboxTopChallengeHash(),
      0,
      this.prevProtoData.hash()
    )
  }

  invalidInboxTopHash(): string {
    return childNodeShortHash(
      this.prevNodeHash(),
      this.invalidInboxTopHashInner()
    )
  }

  updatedProtoData(): VMProtoData {
    return new VMProtoData(
      this.claims.executionAssertion.afterState,
      this.claims.afterInboxTop,
      this.prevProtoData.inboxCount.add(this.params.importedMessageCount),
      this.prevProtoData.messageCount.add(
        this.claims.executionAssertion.outMessages.length
      ),
      this.prevProtoData.logCount.add(
        this.claims.executionAssertion.outLogs.length
      )
    )
  }

  validDataHash(): string {
    return ethers.utils.solidityKeccak256(
      ['bytes32', 'bytes32'],
      [
        this.claims.executionAssertion.outMessagesAcc(),
        this.claims.executionAssertion.outLogsAcc(),
      ]
    )
  }

  validHashInner(): string {
    return childNodeInnerHash(
      this.deadline(),
      this.validDataHash(),
      2,
      this.updatedProtoData().hash()
    )
  }

  validHash(): string {
    return childNodeShortHash(this.prevNodeHash(), this.validHashInner())
  }
}

async function makeAssertion(
  arbRollup: ArbRollup,
  prevPrevNode: string,
  prevProtoData: VMProtoData,
  prevDeadline: utils.BigNumberish,
  prevDataHash: string,
  prevChildType: number,
  params: AssertionParams,
  claims: AssertionClaim,
  stakerProof: Array<string>,
  knownValidBlockHash: string,
  knownValidBlockHeight: number
): Promise<{ receipt: providers.TransactionReceipt; assertion: Assertion }> {
  const fields1 = [
    prevProtoData.machineHash,
    claims.executionAssertion.afterState,
    prevProtoData.inboxTop,
    claims.afterInboxTop,
    claims.executionAssertion.outMessagesAcc(),
    claims.executionAssertion.outLogsAcc(),
    prevPrevNode,
    prevDataHash,
  ]
  const fields2 = [
    prevProtoData.inboxCount,
    prevDeadline,
    params.importedMessageCount,
    prevProtoData.messageCount,
    prevProtoData.logCount,
  ]
  const tx = await arbRollup.makeAssertion(
    fields1,
    fields2,
    knownValidBlockHash,
    knownValidBlockHeight,
    claims.executionAssertion.outMessages.length,
    claims.executionAssertion.outLogs.length,
    prevChildType,
    params.numSteps,
    claims.executionAssertion.numGas,
    stakerProof
  )

  const receipt = await tx.wait()
  if (receipt.blockNumber == undefined || receipt.logs == undefined) {
    throw Error('expected receipt to have block number and logs')
  }

  const logs = receipt.logs.map((log: providers.Log) =>
    arbRollup.interface.parseLog(log)
  )
  const assertion = new Assertion(
    receipt.blockNumber,
    logs[0].values.fields[1],
    logs[0].values.inboxCount,
    prevPrevNode,
    prevProtoData,
    prevDeadline,
    prevDataHash,
    prevChildType,
    params,
    claims
  )

  return {
    receipt: receipt,
    assertion,
  }
}

let arbFactory: ArbFactory
let arbRollup: ArbRollup
let challenge: InboxTopChallenge
let rollupTester: RollupTester
let assertionInfo: Assertion
let originalNode: string
let accounts: Signer[]

describe('ArbRollup', async () => {
  it('should deploy contracts', async function () {
    accounts = await ethers.getSigners()
    const { ArbFactory } = await deploy_contracts(bre)
    arbFactory = ArbFactory as ArbFactory

    const RollupTester = await ethers.getContractFactory('RollupTester')
    rollupTester = (await RollupTester.deploy()) as RollupTester
    await rollupTester.deployed()
  })

  it('should initialize', async function () {
    const tx = arbFactory.createRollup(
      initialVmState, // vmState
      gracePeriodTicks, // gracePeriodTicks
      1000000, // arbGasSpeedLimitPerTick
      maxExecutionSteps, // maxExecutionSteps
      stakeRequirement, // stakeRequirement
      await accounts[0].getAddress(), // owner
      '0x'
    )
    await expect(tx).to.emit(arbFactory, 'RollupCreated')

    const receipt = await (await tx).wait()
    if (receipt.logs == undefined) {
      throw Error('expected receipt to have logs')
    }

    const logs = receipt.logs.map((log: providers.Log) =>
      arbFactory.interface.parseLog(log)
    )
    const ev = logs[2]
    expect(ev.name).to.equal('RollupCreated')
    const chainAddress = ev.values.rollupAddress
    const ArbRollup = await ethers.getContractFactory('ArbRollup')
    arbRollup = ArbRollup.attach(chainAddress) as ArbRollup

    originalNode = await arbRollup.latestConfirmed()
    assert.isTrue(
      await arbRollup.isValidLeaf(originalNode),
      'original node should be a leaf'
    )
  })

  it('should fail to assert on invalid leaf', async () => {
    await expect(
      makeEmptyAssertion(
        arbRollup,
        '0x3400000000000000000000000000000000000000000000000000000000000000',
        0,
        0
      )
    ).to.be.revertedWith('MAKE_LEAF')
  })

  // it("should fail to assert on halted vm", async () => {
  //   truffleAssert.reverts(makeEmptyAssertion("0x00", 0, 0), "MAKE_RUN");
  // })

  it('should fail to assert over step limit', async () => {
    await expect(
      makeEmptyAssertion(arbRollup, initialVmState, maxExecutionSteps + 1, 0)
    ).to.be.revertedWith('MAKE_STEP')
  })

  it('should fail to assert without stake', async () => {
    await expect(
      makeEmptyAssertion(arbRollup, initialVmState, 0, 0)
    ).to.be.revertedWith('INV_STAKER')
  })

  it('should fail if reading past lastest inbox message', async () => {
    await expect(
      makeEmptyAssertion(arbRollup, initialVmState, 0, 10)
    ).to.be.revertedWith('MAKE_MESSAGE_CNT')
  })

  it('should create a stake', async () => {
    await expect(arbRollup.isStaked(await accounts[0].getAddress())).to
      .eventually.be.false
    await expect(
      arbRollup.connect(accounts[0]).placeStake([], [], {
        value: stakeRequirement,
      })
    ).to.emit(arbRollup, 'RollupStakeCreated')
    await expect(arbRollup.isStaked(await accounts[0].getAddress())).to
      .eventually.be.true
  })

  it('should create a second stake', async () => {
    await expect(
      arbRollup.connect(accounts[1]).placeStake([], [], {
        value: stakeRequirement,
      })
    ).to.emit(arbRollup, 'RollupStakeCreated')
  })

  it('should make an assertion', async () => {
    assert.isTrue(
      await arbRollup.isValidLeaf(originalNode),
      'latest confirmed should be leaf before asserting'
    )
    const prevProtoData = new VMProtoData(initialVmState, zerobytes32, 0, 0, 0)
    const params = new AssertionParams(0, ethers.utils.bigNumberify(0))
    const claims = new AssertionClaim(
      zerobytes32,
      new ExecutionAssertion(
        '0x8500000000000000000000000000000000000000000000000000000000000000',
        0,
        [],
        []
      )
    )
    const block = await ethers.provider.getBlock('latest')
    const info = await makeAssertion(
      arbRollup,
      zerobytes32,
      prevProtoData,
      0,
      zerobytes32,
      0,
      params,
      claims,
      [],
      block.hash,
      block.number
    )

    assertionInfo = info.assertion

    assert.isFalse(
      await arbRollup.isValidLeaf(assertionInfo.prevNodeHash()),
      'originalNode confirmed should be removed as leaf'
    )
    assert.isTrue(
      await arbRollup.isValidLeaf(assertionInfo.invalidInboxTopHash()),
      'invalid inbox top should be leaf'
    )
    // TODO: Check whether invalid execution is leaf
    assert.isTrue(
      await arbRollup.isValidLeaf(assertionInfo.validHash()),
      'valid child should be leaf'
    )
  })

  it('should allow the second staker to move to conflicting node', async () => {
    await expect(
      arbRollup
        .connect(accounts[1])
        .moveStake([assertionInfo.invalidInboxTopHashInner()], [])
    )
      .to.emit(arbRollup, 'RollupStakeMoved')
      .withArgs(
        await accounts[1].getAddress(),
        assertionInfo.invalidInboxTopHash()
      )
  })

  it('should allow the creation of a challenge', async () => {
    const txPromise = arbRollup.startChallenge(
      await accounts[0].getAddress(),
      await accounts[1].getAddress(),
      assertionInfo.prevNodeHash(),
      assertionInfo.deadline(),
      [2, 0],
      [
        assertionInfo.updatedProtoData().hash(),
        assertionInfo.prevProtoData.hash(),
      ],
      [],
      [],
      assertionInfo.validDataHash(),
      assertionInfo.invalidInboxTopDataHash(),
      assertionInfo.challengePeriod()
    )
    const receipt = await (await txPromise).wait()
    if (receipt.logs === undefined) {
      throw Error('logs must be defined')
    }
    expect(receipt.logs).to.have.lengthOf(2)
    const logs = receipt.logs.map((log: providers.Log) =>
      arbRollup.interface.parseLog(log)
    )
    const ev = logs[1]
    expect(ev.name).equals('RollupChallengeStarted')
    const challengeContract = ev.values.challengeContract

    const InboxTopChallenge = await ethers.getContractFactory(
      'InboxTopChallenge'
    )
    challenge = InboxTopChallenge.attach(challengeContract) as InboxTopChallenge
  })

  it('should timeout the challenge', async () => {
    await ethers.provider.send('evm_mine', [])
    await ethers.provider.send('evm_mine', [])
    await ethers.provider.send('evm_mine', [])

    await challenge.timeoutChallenge()
  })

  it('should confirm invalid inbox top node', async () => {
    await ethers.provider.send('evm_mine', [])
    await ethers.provider.send('evm_mine', [])
    await ethers.provider.send('evm_mine', [])

    await expect(
      arbRollup.confirm(
        assertionInfo.prevProtoData.hash(),
        [0],
        [assertionInfo.deadline()],
        [assertionInfo.invalidInboxTopChallengeHash()],
        [],
        [],
        [],
        '0x',
        [await accounts[1].getAddress()].sort(),
        [],
        [0, 0]
      )
    ).to.emit(arbRollup, 'RollupConfirmed')

    assert.equal(
      await arbRollup.latestConfirmed(),
      assertionInfo.invalidInboxTopHash(),
      'latest confirmed should now be invalid inbox child'
    )

    assert.isTrue(
      await arbRollup.isValidLeaf(assertionInfo.invalidInboxTopHash()),
      'invalid inbox top should be leaf'
    )
  })

  it('should prune a leaf', async () => {
    assert.isTrue(
      await arbRollup.isValidLeaf(assertionInfo.validHash()),
      'valid node should be leaf'
    )
    await expect(
      arbRollup.pruneLeaves(
        [originalNode],
        [assertionInfo.validHashInner()],
        [1],
        [assertionInfo.invalidInboxTopHashInner()],
        [1]
      )
    ).to.emit(arbRollup, 'RollupPruned')

    assert.isFalse(
      await arbRollup.isValidLeaf(assertionInfo.validHashInner()),
      'valid node should no longer be leaf'
    )
  })

  it('should assert again', async () => {
    const params = new AssertionParams(0, ethers.utils.bigNumberify(0))
    const claims = new AssertionClaim(
      zerobytes32,
      new ExecutionAssertion(
        zerobytes32,
        0,
        [new ArbValue.TupleValue([new ArbValue.IntValue(10)])],
        []
      )
    )

    const block = await ethers.provider.getBlock('latest')
    assertionInfo = (
      await makeAssertion(
        arbRollup.connect(accounts[1]),
        assertionInfo.prevNodeHash(),
        assertionInfo.prevProtoData,
        assertionInfo.deadline(),
        assertionInfo.invalidInboxTopChallengeHash(),
        0,
        params,
        claims,
        [],
        block.hash,
        block.number
      )
    ).assertion
  })

  it('should confirm valid node', async () => {
    await ethers.provider.send('evm_mine', [])
    await ethers.provider.send('evm_mine', [])
    await ethers.provider.send('evm_mine', [])

    const { validNodeHashes, lastNode } = await rollupTester.confirm(
      await arbRollup.latestConfirmed(),
      assertionInfo.prevProtoData.hash(),
      [2],
      [assertionInfo.deadline()],
      [],
      [assertionInfo.claims.executionAssertion.outLogsAcc()],
      [assertionInfo.updatedProtoData().hash()],
      [assertionInfo.claims.executionAssertion.outMessages.length],
      assertionInfo.claims.executionAssertion.outMessagesData()
    )

    expect(validNodeHashes.length).to.equal(1)
    expect(validNodeHashes[0]).to.equal(assertionInfo.validHash())

    assert.equal(
      lastNode,
      assertionInfo.validHash(),
      'calculated last node should be the valid node'
    )

    await expect(
      arbRollup.confirm(
        assertionInfo.prevProtoData.hash(),
        [2],
        [assertionInfo.deadline()],
        [],
        [assertionInfo.claims.executionAssertion.outLogsAcc()],
        [assertionInfo.updatedProtoData().hash()],
        [assertionInfo.claims.executionAssertion.outMessages.length],
        assertionInfo.claims.executionAssertion.outMessagesData(),
        [await accounts[1].getAddress()].sort(),
        [],
        [0, 0]
      )
    ).to.emit(arbRollup, 'RollupConfirmed')

    assert.equal(
      await arbRollup.latestConfirmed(),
      assertionInfo.validHash(),
      'latest confirmed should now be valid child'
    )

    assert.isTrue(
      await arbRollup.isValidLeaf(assertionInfo.validHash()),
      'valid child should be leaf'
    )
  })

  it('should allow second staker to withdraw', async () => {
    await expect(arbRollup.connect(accounts[1]).recoverStakeConfirmed([]))
      .to.emit(arbRollup, 'RollupStakeRefunded')
      .withArgs(await accounts[1].getAddress())
  })
})
