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

import { ethers, deployments } from '@nomiclabs/buidler'
import { Signer, ContractTransaction, providers, utils } from 'ethers'
import * as chai from 'chai'
import chaiAsPromised from 'chai-as-promised'
import { ArbRollup } from '../build/types/ArbRollup'
import { ArbFactory } from '../build/types/ArbFactory'
import { InboxTopChallenge } from '../build/types/InboxTopChallenge'
import { ArbValue } from 'arb-provider-ethers'

chai.use(chaiAsPromised)

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

function messagesHash(
  lowerHashA: string,
  topHashA: string,
  lowerHashB: string,
  topHashB: string,
  chainLength: utils.BigNumberish
): string {
  return ethers.utils.solidityKeccak256(
    ['bytes32', 'bytes32', 'bytes32', 'bytes32', 'uint256'],
    [lowerHashA, topHashA, lowerHashB, topHashB, chainLength]
  )
}

function invalidMessagesHash(
  lowerHashA: string,
  topHashA: string,
  lowerHashB: string,
  topHashB: string,
  chainLength: utils.BigNumberish,
  challengePeriod: utils.BigNumberish
): string {
  return ethers.utils.solidityKeccak256(
    ['bytes32', 'uint256'],
    [
      messagesHash(lowerHashA, topHashA, lowerHashB, topHashB, chainLength),
      challengePeriod,
    ]
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

const emptyTupleHash = ethers.utils.solidityKeccak256(
  ['uint8', 'bytes32', 'uint256'],
  [3, ethers.utils.solidityKeccak256(['uint8'], [0]), 1]
)

const zerobytes32 =
  '0x0000000000000000000000000000000000000000000000000000000000000000'

async function makeEmptyAssertion(
  arbRollup: ArbRollup,
  vmState: string,
  numSteps: number,
  startBlock: utils.BigNumberish,
  startTime: utils.BigNumberish,
  importedMessageCount: utils.BigNumberish,
  readInbox: boolean
): Promise<ContractTransaction> {
  const startBlockInt = ethers.utils.bigNumberify(startBlock)
  const startTimeInt = ethers.utils.bigNumberify(startTime)
  return arbRollup.makeAssertion(
    [
      vmState,
      emptyTupleHash,
      zerobytes32,
      zerobytes32,
      emptyTupleHash,
      emptyTupleHash,
      zerobytes32,
      zerobytes32,
      zerobytes32,
    ],
    0,
    0,
    0,
    numSteps,
    [startBlockInt, startBlockInt.add(10), startTimeInt, startTimeInt.add(100)],
    importedMessageCount,
    readInbox,
    0,
    []
  )
}

class VMProtoData {
  public inboxCount: utils.BigNumber
  constructor(
    public machineHash: string,
    public inboxTop: string,
    inboxCount: utils.BigNumberish
  ) {
    this.inboxCount = ethers.utils.bigNumberify(inboxCount)
  }

  hash(): string {
    return ethers.utils.solidityKeccak256(
      ['bytes32', 'bytes32', 'uint256'],
      [this.machineHash, this.inboxTop, this.inboxCount]
    )
  }
}

class AssertionParams {
  constructor(
    public numSteps: number,
    public timeBounds: utils.BigNumberish[],
    public importedMessageCount: utils.BigNumberish
  ) {}
}

class ExecutionAssertion {
  constructor(
    public afterState: string,
    public didReadInbox: boolean,
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
      (acc, hash) =>
        ethers.utils.solidityKeccak256(['bytes32', 'bytes32'], [acc, hash]),
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
    public importedMessageSlice: string,
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

  invalidMessagesHashInner(): string {
    return childNodeInnerHash(
      this.deadline(),
      invalidMessagesHash(
        this.prevProtoData.inboxTop,
        this.claims.afterInboxTop,
        emptyTupleHash,
        this.claims.importedMessageSlice,
        this.params.importedMessageCount,
        gracePeriodTicks.add(ethers.utils.bigNumberify(1000))
      ),
      1,
      this.prevProtoData.hash()
    )
  }

  invalidMessagesHash(): string {
    return childNodeShortHash(
      this.prevNodeHash(),
      this.invalidMessagesHashInner()
    )
  }

  updatedProtoData(): VMProtoData {
    return new VMProtoData(
      this.claims.executionAssertion.afterState,
      this.claims.afterInboxTop,
      this.prevProtoData.inboxCount.add(this.params.importedMessageCount)
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
      3,
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
  stakerProof: Array<string>
): Promise<{ receipt: providers.TransactionReceipt; assertion: Assertion }> {
  const tx = await arbRollup.makeAssertion(
    [
      prevProtoData.machineHash,
      prevProtoData.inboxTop,
      prevPrevNode,
      prevDataHash,
      claims.afterInboxTop,
      claims.importedMessageSlice,
      claims.executionAssertion.afterState,
      claims.executionAssertion.outMessagesAcc(),
      claims.executionAssertion.outLogsAcc(),
    ],
    prevProtoData.inboxCount,
    prevDeadline,
    prevChildType,
    params.numSteps,
    params.timeBounds,
    params.importedMessageCount,
    claims.executionAssertion.didReadInbox,
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
  return {
    receipt: receipt,
    assertion: new Assertion(
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
    ),
  }
}

let arbRollup: ArbRollup
let challenge: InboxTopChallenge
let assertionInfo: Assertion
let originalNode: string
let accounts: Signer[]

describe('ArbRollup', function () {
  before(async function () {
    accounts = await ethers.getSigners()
    await deployments.fixture()
  })

  it('should initialize', async function () {
    const arbFactoryDeployment = await deployments.get('ArbFactory')
    const ArbRollupFactory = await ethers.getContractFactory('ArbFactory')
    const arbFactory = ArbRollupFactory.attach(
      arbFactoryDeployment.address
    ) as ArbFactory
    const tx = arbFactory.createRollup(
      initialVmState, // vmState
      gracePeriodTicks, // gracePeriodTicks
      1000000, // arbGasSpeedLimitPerTick
      maxExecutionSteps, // maxExecutionSteps
      [20, 1000], // maxTimeBoundsWidth
      stakeRequirement, // stakeRequirement
      await accounts[0].getAddress() // owner
    )
    await expect(tx).to.emit(arbFactory, 'RollupCreated')

    const receipt = await (await tx).wait()
    if (receipt.logs == undefined) {
      throw Error('expected receipt to have logs')
    }

    const logs = receipt.logs.map((log: providers.Log) =>
      arbFactory.interface.parseLog(log)
    )
    const ev = logs[1]
    expect(ev.name).to.equal('RollupCreated')
    const chainAddress = ev.values.vmAddress
    const ArbRollup = await ethers.getContractFactory('ArbRollup')
    arbRollup = ArbRollup.attach(chainAddress) as ArbRollup

    originalNode = await arbRollup.latestConfirmed()
    assert.isTrue(
      await arbRollup.isValidLeaf(originalNode),
      'original node should be a leaf'
    )
  })

  it('should fail to assert on invalid leaf', async () => {
    const currentBlock = await ethers.provider.getBlock('latest')
    await expect(
      makeEmptyAssertion(
        arbRollup,
        '0x3400000000000000000000000000000000000000000000000000000000000000',
        0,
        currentBlock.number,
        currentBlock.timestamp,
        0,
        false
      )
    ).to.be.revertedWith('MAKE_LEAF')
  })

  // it("should fail to assert on halted vm", async () => {
  //   truffleAssert.reverts(makeEmptyAssertion("0x00", 0, 0), "MAKE_RUN");
  // })

  it('should fail to assert over step limit', async () => {
    const currentBlock = await ethers.provider.getBlock('latest')
    await expect(
      makeEmptyAssertion(
        arbRollup,
        initialVmState,
        maxExecutionSteps + 1,
        currentBlock.number,
        currentBlock.timestamp,
        0,
        false
      )
    ).to.be.revertedWith('MAKE_STEP')
  })

  it('should fail to assert without stake', async () => {
    const currentBlock = await ethers.provider.getBlock('latest')
    await expect(
      makeEmptyAssertion(
        arbRollup,
        initialVmState,
        0,
        currentBlock.number,
        currentBlock.timestamp,
        0,
        false
      )
    ).to.be.revertedWith('INV_STAKER')
  })

  it('should fail to assert outside time bounds', async () => {
    await expect(
      makeEmptyAssertion(arbRollup, initialVmState, 0, 10000, 10000, 0, false)
    ).to.be.revertedWith('MAKE_TIME')
  })

  it('should fail if consuming messages but not reading inbox', async () => {
    const currentBlock = await ethers.provider.getBlock('latest')
    await expect(
      makeEmptyAssertion(
        arbRollup,
        initialVmState,
        0,
        currentBlock.number,
        currentBlock.timestamp,
        10,
        false
      )
    ).to.be.revertedWith('MAKE_MESSAGES')
  })

  it('should fail if reading past lastest inbox message', async () => {
    const currentBlock = await ethers.provider.getBlock('latest')
    await expect(
      makeEmptyAssertion(
        arbRollup,
        initialVmState,
        0,
        currentBlock.number,
        currentBlock.timestamp,
        10,
        true
      )
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
    const currentBlock = await ethers.provider.getBlock('latest')
    const prevProtoData = new VMProtoData(
      initialVmState,
      emptyTupleHash,
      ethers.utils.bigNumberify(0)
    )
    const params = new AssertionParams(
      0,
      [
        currentBlock.number,
        currentBlock.number + 10,
        currentBlock.timestamp,
        currentBlock.timestamp + 100,
      ],
      ethers.utils.bigNumberify(0)
    )
    const claims = new AssertionClaim(
      zerobytes32,
      emptyTupleHash,
      new ExecutionAssertion(
        '0x8500000000000000000000000000000000000000000000000000000000000000',
        false,
        0,
        [],
        []
      )
    )
    const info = await makeAssertion(
      arbRollup,
      zerobytes32,
      prevProtoData,
      0,
      zerobytes32,
      0,
      params,
      claims,
      []
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
    assert.isTrue(
      await arbRollup.isValidLeaf(assertionInfo.invalidMessagesHash()),
      'invalid messages should be leaf'
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
      [3, 0],
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
      await arbRollup.isValidLeaf(assertionInfo.invalidInboxTopHash()),
      'invalid messages should be leaf'
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
    const currentBlock = await ethers.provider.getBlock('latest')
    const params = new AssertionParams(
      0,
      [
        currentBlock.number,
        currentBlock.number + 10,
        currentBlock.timestamp,
        currentBlock.timestamp + 100,
      ],
      ethers.utils.bigNumberify(0)
    )
    const claims = new AssertionClaim(
      zerobytes32,
      emptyTupleHash,
      new ExecutionAssertion(
        zerobytes32,
        false,
        0,
        [new ArbValue.TupleValue([new ArbValue.IntValue(10)])],
        []
      )
    )

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
        []
      )
    ).assertion
  })

  it('should confirm valid node', async () => {
    await ethers.provider.send('evm_mine', [])
    await ethers.provider.send('evm_mine', [])
    await ethers.provider.send('evm_mine', [])

    await expect(
      arbRollup.confirm(
        assertionInfo.prevProtoData.hash(),
        [3],
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
