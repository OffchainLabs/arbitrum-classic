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

import bre from '@nomiclabs/buidler'
import { Signer, providers, utils } from 'ethers'
import * as chai from 'chai'
import * as chaiAsPromised from 'chai-as-promised'
import { ArbRollup } from '../build/types/ArbRollup'
import { ArbFactory } from '../build/types/ArbFactory'
import { InboxTopChallenge } from '../build/types/InboxTopChallenge'
import { ArbValue } from 'arb-provider-ethers'

import deploy_contracts from '../scripts/deploylib'

chai.use(require('chai-as-promised'))

const { assert, expect } = chai
const ethers = bre.ethers

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

function invalidInboxTopHash(
  lowerHash: string,
  topHash: string,
  chainLength: utils.BigNumberish,
  challengePeriod: utils.BigNumberish
): string {
  return ethers.utils.solidityKeccak256(
    ['bytes32', 'uint256'],
    [inboxTopHash(lowerHash, topHash, chainLength), challengePeriod]
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

function childNodeShortHash(prevNodeHash: string, nodeInnerHash: string) {
  return ethers.utils.solidityKeccak256(
    ['bytes32', 'bytes32'],
    [prevNodeHash, nodeInnerHash]
  )
}

let empty_tuple_hash = ethers.utils.solidityKeccak256(
  ['uint8', 'bytes32', 'uint256'],
  [3, ethers.utils.solidityKeccak256(['uint8'], [0]), 1]
)

let zerobytes32 =
  '0x0000000000000000000000000000000000000000000000000000000000000000'

async function makeEmptyAssertion(
  arb_rollup: ArbRollup,
  vm_state: string,
  num_steps: number,
  start_block: utils.BigNumberish,
  start_time: utils.BigNumberish,
  imported_message_count: utils.BigNumberish,
  read_inbox: boolean
) {
  const start_block_int = ethers.utils.bigNumberify(start_block)
  const start_time_int = ethers.utils.bigNumberify(start_time)
  return arb_rollup.makeAssertion(
    [
      vm_state,
      empty_tuple_hash,
      zerobytes32,
      zerobytes32,
      empty_tuple_hash,
      empty_tuple_hash,
      zerobytes32,
      zerobytes32,
      zerobytes32,
    ],
    0,
    0,
    0,
    num_steps,
    [
      start_block_int,
      start_block_int.add(10),
      start_time_int,
      start_time_int.add(100),
    ],
    imported_message_count,
    read_inbox,
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

  hash() {
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

  outMessagesAcc() {
    return this.outMessages.reduce(
      (acc, val) =>
        ethers.utils.solidityKeccak256(
          ['bytes32', 'bytes32'],
          [acc, val.hash()]
        ),
      zerobytes32
    )
  }

  outLogsAcc() {
    return this.outLogs.reduce(
      (acc, hash) =>
        ethers.utils.solidityKeccak256(['bytes32', 'bytes32'], [acc, hash]),
      zerobytes32
    )
  }

  outMessagesData() {
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

  prevNodeHash() {
    return childNodeHash(
      this.prevPrevNode,
      this.prevDeadline,
      this.prevDataHash,
      this.prevChildType,
      this.prevProtoData.hash()
    )
  }

  deadline() {
    let deadlineTicks = this.blockNumber
      .mul(ethers.utils.bigNumberify(1000))
      .add(grace_period_ticks)
    if (deadlineTicks.lt(this.prevDeadline)) {
      deadlineTicks = this.prevDeadline
    }
    // Should be numArbGas / arbGasSpeedLimitPerTick, but numArbGas is 0 in the test
    return deadlineTicks.add(ethers.utils.bigNumberify(0))
  }

  challengePeriod() {
    // should be plus numArbGas
    return grace_period_ticks.add(ethers.utils.bigNumberify(1000))
  }

  invalidInboxTopDataHash() {
    return inboxTopHash(
      this.claims.afterInboxTop,
      this.inboxValue,
      this.inboxCount.sub(
        this.prevProtoData.inboxCount.add(this.params.importedMessageCount)
      )
    )
  }

  invalidInboxTopChallengeHash() {
    return ethers.utils.solidityKeccak256(
      ['bytes32', 'uint256'],
      [this.invalidInboxTopDataHash(), this.challengePeriod()]
    )
  }

  invalidInboxTopHashInner() {
    return childNodeInnerHash(
      this.deadline(),
      this.invalidInboxTopChallengeHash(),
      0,
      this.prevProtoData.hash()
    )
  }

  invalidInboxTopHash() {
    return childNodeShortHash(
      this.prevNodeHash(),
      this.invalidInboxTopHashInner()
    )
  }

  invalidMessagesHashInner() {
    return childNodeInnerHash(
      this.deadline(),
      invalidMessagesHash(
        this.prevProtoData.inboxTop,
        this.claims.afterInboxTop,
        empty_tuple_hash,
        this.claims.importedMessageSlice,
        this.params.importedMessageCount,
        grace_period_ticks.add(ethers.utils.bigNumberify(1000))
      ),
      1,
      this.prevProtoData.hash()
    )
  }

  invalidMessagesHash() {
    return childNodeShortHash(
      this.prevNodeHash(),
      this.invalidMessagesHashInner()
    )
  }

  updatedProtoData() {
    return new VMProtoData(
      this.claims.executionAssertion.afterState,
      this.claims.afterInboxTop,
      this.prevProtoData.inboxCount.add(this.params.importedMessageCount)
    )
  }

  validDataHash() {
    return ethers.utils.solidityKeccak256(
      ['bytes32', 'bytes32'],
      [
        this.claims.executionAssertion.outMessagesAcc(),
        this.claims.executionAssertion.outLogsAcc(),
      ]
    )
  }

  validHashInner() {
    return childNodeInnerHash(
      this.deadline(),
      this.validDataHash(),
      3,
      this.updatedProtoData().hash()
    )
  }

  validHash() {
    return childNodeShortHash(this.prevNodeHash(), this.validHashInner())
  }
}

async function makeAssertion(
  arb_rollup: ArbRollup,
  prevPrevNode: string,
  prevProtoData: VMProtoData,
  prevDeadline: utils.BigNumberish,
  prevDataHash: string,
  prevChildType: number,
  params: AssertionParams,
  claims: AssertionClaim,
  stakerProof: Array<string>
) {
  let tx = await arb_rollup.makeAssertion(
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
    arb_rollup.interface.parseLog(log)
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

let initial_vm_state =
  '0x9900000000000000000000000000000000000000000000000000000000000000'
let stakeRequirement = 10
let max_execution_steps = 50000
let grace_period_ticks = ethers.utils.bigNumberify(1000)

let arb_rollup: ArbRollup
let challenge: InboxTopChallenge
let assertionInfo: Assertion
let original_node: string
let accounts: Signer[]

describe('ArbRollup', function () {
  before(async function () {
    accounts = await ethers.getSigners()
  })

  it('should initialize', async function () {
    const { ArbFactory } = await deploy_contracts(bre)
    const ArbRollupFactory = await ethers.getContractFactory('ArbFactory')
    const arb_factory = ArbRollupFactory.attach(
      ArbFactory.address
    ) as ArbFactory
    let tx = arb_factory.createRollup(
      initial_vm_state, // vmState
      grace_period_ticks, // gracePeriodTicks
      1000000, // arbGasSpeedLimitPerTick
      max_execution_steps, // maxExecutionSteps
      [20, 1000], // maxTimeBoundsWidth
      stakeRequirement, // stakeRequirement
      await accounts[0].getAddress() // owner
    )
    await expect(tx).to.emit(arb_factory, 'RollupCreated')

    let receipt = await (await tx).wait()
    if (receipt.logs == undefined) {
      throw Error('expected receipt to have logs')
    }

    const logs = receipt.logs.map((log: providers.Log) =>
      arb_factory.interface.parseLog(log)
    )
    const ev = logs[1]
    expect(ev.name).to.equal('RollupCreated')
    let chain_address = ev.values.vmAddress
    const ArbRollup = await ethers.getContractFactory('ArbRollup')
    arb_rollup = ArbRollup.attach(chain_address) as ArbRollup

    original_node = await arb_rollup.latestConfirmed()
    assert.isTrue(
      await arb_rollup.isValidLeaf(original_node),
      'original node should be a leaf'
    )
  })

  it('should fail to assert on invalid leaf', async () => {
    let current_block = await ethers.provider.getBlock('latest')
    await expect(
      makeEmptyAssertion(
        arb_rollup,
        '0x3400000000000000000000000000000000000000000000000000000000000000',
        0,
        current_block.number,
        current_block.timestamp,
        0,
        false
      )
    ).to.be.revertedWith('MAKE_LEAF')
  })

  // it("should fail to assert on halted vm", async () => {
  //   truffleAssert.reverts(makeEmptyAssertion("0x00", 0, 0), "MAKE_RUN");
  // })

  it('should fail to assert over step limit', async () => {
    let current_block = await ethers.provider.getBlock('latest')
    await expect(
      makeEmptyAssertion(
        arb_rollup,
        initial_vm_state,
        max_execution_steps + 1,
        current_block.number,
        current_block.timestamp,
        0,
        false
      )
    ).to.be.revertedWith('MAKE_STEP')
  })

  it('should fail to assert without stake', async () => {
    let current_block = await ethers.provider.getBlock('latest')
    await expect(
      makeEmptyAssertion(
        arb_rollup,
        initial_vm_state,
        0,
        current_block.number,
        current_block.timestamp,
        0,
        false
      )
    ).to.be.revertedWith('INV_STAKER')
  })

  it('should fail to assert outside time bounds', async () => {
    await expect(
      makeEmptyAssertion(
        arb_rollup,
        initial_vm_state,
        0,
        10000,
        10000,
        0,
        false
      )
    ).to.be.revertedWith('MAKE_TIME')
  })

  it('should fail if consuming messages but not reading inbox', async () => {
    let current_block = await ethers.provider.getBlock('latest')
    await expect(
      makeEmptyAssertion(
        arb_rollup,
        initial_vm_state,
        0,
        current_block.number,
        current_block.timestamp,
        10,
        false
      )
    ).to.be.revertedWith('MAKE_MESSAGES')
  })

  it('should fail if reading past lastest inbox message', async () => {
    let current_block = await ethers.provider.getBlock('latest')
    await expect(
      makeEmptyAssertion(
        arb_rollup,
        initial_vm_state,
        0,
        current_block.number,
        current_block.timestamp,
        10,
        true
      )
    ).to.be.revertedWith('MAKE_MESSAGE_CNT')
  })

  it('should create a stake', async () => {
    await expect(arb_rollup.isStaked(await accounts[0].getAddress())).to
      .eventually.be.false
    await expect(
      arb_rollup.connect(accounts[0]).placeStake([], [], {
        value: stakeRequirement,
      })
    ).to.emit(arb_rollup, 'RollupStakeCreated')
    await expect(arb_rollup.isStaked(await accounts[0].getAddress())).to
      .eventually.be.true
  })

  it('should create a second stake', async () => {
    await expect(
      arb_rollup.connect(accounts[1]).placeStake([], [], {
        value: stakeRequirement,
      })
    ).to.emit(arb_rollup, 'RollupStakeCreated')
  })

  it('should make an assertion', async () => {
    assert.isTrue(
      await arb_rollup.isValidLeaf(original_node),
      'latest confirmed should be leaf before asserting'
    )
    let current_block = await ethers.provider.getBlock('latest')
    let prevProtoData = new VMProtoData(
      initial_vm_state,
      empty_tuple_hash,
      ethers.utils.bigNumberify(0)
    )
    let params = new AssertionParams(
      0,
      [
        current_block.number,
        current_block.number + 10,
        current_block.timestamp,
        current_block.timestamp + 100,
      ],
      ethers.utils.bigNumberify(0)
    )
    let claims = new AssertionClaim(
      zerobytes32,
      empty_tuple_hash,
      new ExecutionAssertion(
        '0x8500000000000000000000000000000000000000000000000000000000000000',
        false,
        0,
        [],
        []
      )
    )
    let info = await makeAssertion(
      arb_rollup,
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
      await arb_rollup.isValidLeaf(assertionInfo.prevNodeHash()),
      'original_node confirmed should be removed as leaf'
    )
    assert.isTrue(
      await arb_rollup.isValidLeaf(assertionInfo.invalidInboxTopHash()),
      'invalid inbox top should be leaf'
    )
    assert.isTrue(
      await arb_rollup.isValidLeaf(assertionInfo.invalidMessagesHash()),
      'invalid messages should be leaf'
    )
    // TODO: Check whether invalid execution is leaf
    assert.isTrue(
      await arb_rollup.isValidLeaf(assertionInfo.validHash()),
      'valid child should be leaf'
    )
  })

  it('should allow the second staker to move to conflicting node', async () => {
    await expect(
      arb_rollup
        .connect(accounts[1])
        .moveStake([assertionInfo.invalidInboxTopHashInner()], [])
    )
      .to.emit(arb_rollup, 'RollupStakeMoved')
      .withArgs(
        await accounts[1].getAddress(),
        assertionInfo.invalidInboxTopHash()
      )
  })

  it('should allow the creation of a challenge', async () => {
    let tx_promise = arb_rollup.startChallenge(
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
    let receipt = await (await tx_promise).wait()
    if (receipt.logs === undefined) {
      throw Error('logs must be defined')
    }
    expect(receipt.logs).to.have.lengthOf(2)
    const logs = receipt.logs.map((log: providers.Log) =>
      arb_rollup.interface.parseLog(log)
    )
    let ev = logs[1]
    expect(ev.name).equals('RollupChallengeStarted')
    const challenge_contract = ev.values.challengeContract

    const InboxTopChallenge = await ethers.getContractFactory(
      'InboxTopChallenge'
    )
    challenge = InboxTopChallenge.attach(
      challenge_contract
    ) as InboxTopChallenge
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
      arb_rollup.confirm(
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
    ).to.emit(arb_rollup, 'RollupConfirmed')

    assert.equal(
      await arb_rollup.latestConfirmed(),
      assertionInfo.invalidInboxTopHash(),
      'latest confirmed should now be invalid inbox child'
    )

    assert.isTrue(
      await arb_rollup.isValidLeaf(assertionInfo.invalidInboxTopHash()),
      'invalid inbox top should be leaf'
    )
  })

  it('should prune a leaf', async () => {
    assert.isTrue(
      await arb_rollup.isValidLeaf(assertionInfo.invalidInboxTopHash()),
      'invalid messages should be leaf'
    )
    await expect(
      arb_rollup.pruneLeaves(
        [original_node],
        [assertionInfo.validHashInner()],
        [1],
        [assertionInfo.invalidInboxTopHashInner()],
        [1]
      )
    ).to.emit(arb_rollup, 'RollupPruned')

    assert.isFalse(
      await arb_rollup.isValidLeaf(assertionInfo.validHashInner()),
      'valid node should no longer be leaf'
    )
  })

  it('should assert again', async () => {
    let current_block = await ethers.provider.getBlock('latest')
    let params = new AssertionParams(
      0,
      [
        current_block.number,
        current_block.number + 10,
        current_block.timestamp,
        current_block.timestamp + 100,
      ],
      ethers.utils.bigNumberify(0)
    )
    let claims = new AssertionClaim(
      zerobytes32,
      empty_tuple_hash,
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
        arb_rollup.connect(accounts[1]),
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
      arb_rollup.confirm(
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
    ).to.emit(arb_rollup, 'RollupConfirmed')

    assert.equal(
      await arb_rollup.latestConfirmed(),
      assertionInfo.validHash(),
      'latest confirmed should now be valid child'
    )

    assert.isTrue(
      await arb_rollup.isValidLeaf(assertionInfo.validHash()),
      'valid child should be leaf'
    )
  })

  it('should allow second staker to withdraw', async () => {
    await expect(arb_rollup.connect(accounts[1]).recoverStakeConfirmed([]))
      .to.emit(arb_rollup, 'RollupStakeRefunded')
      .withArgs(await accounts[1].getAddress())
  })
})
