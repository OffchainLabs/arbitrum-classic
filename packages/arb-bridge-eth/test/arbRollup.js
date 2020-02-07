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

const ArbRollup = artifacts.require("ArbRollup");
const ChallengeFactory = artifacts.require("ChallengeFactory");
const GlobalInbox = artifacts.require("GlobalInbox");

const { expectEvent, expectRevert } = require("@openzeppelin/test-helpers");

function inboxTopHash(lowerHash, topHash, chainLength) {
  return web3.utils.soliditySha3(
    { t: "bytes32", v: lowerHash },
    { t: "bytes32", v: topHash },
    { t: "uint256", v: chainLength }
  );
}

function invalidInboxTopHash(lowerHash, topHash, chainLength, challengePeriod) {
  return web3.utils.soliditySha3(
    { t: "bytes32", v: inboxTopHash(lowerHash, topHash, chainLength) },
    { t: "uint256", v: challengePeriod }
  );
}

function messagesHash(lowerHashA, topHashA, lowerHashB, topHashB, chainLength) {
  return web3.utils.soliditySha3(
    { t: "bytes32", v: lowerHashA },
    { t: "bytes32", v: topHashA },
    { t: "bytes32", v: lowerHashB },
    { t: "bytes32", v: topHashB },
    { t: "uint256", v: chainLength }
  );
}

function invalidMessagesHash(
  lowerHashA,
  topHashA,
  lowerHashB,
  topHashB,
  chainLength,
  challengePeriod
) {
  return web3.utils.soliditySha3(
    {
      t: "bytes32",
      v: messagesHash(lowerHashA, topHashA, lowerHashB, topHashB, chainLength)
    },
    { t: "uint256", v: challengePeriod }
  );
}

function childNodeInnerHash(
  deadlineTicks,
  nodeDataHash,
  childType,
  vmProtoStateHash
) {
  return web3.utils.soliditySha3(
    { t: "bytes32", v: vmProtoStateHash },
    { t: "uint256", v: deadlineTicks },
    { t: "bytes32", v: nodeDataHash },
    { t: "uint256", v: childType }
  );
}

function childNodeHash(
  prevNodeHash,
  deadlineTicks,
  nodeDataHash,
  childType,
  vmProtoStateHash
) {
  return web3.utils.soliditySha3(
    { t: "bytes32", v: prevNodeHash },
    {
      t: "bytes32",
      v: childNodeInnerHash(
        deadlineTicks,
        nodeDataHash,
        childType,
        vmProtoStateHash
      )
    }
  );
}

function childNodeShortHash(prevNodeHash, nodeInnerHash) {
  return web3.utils.soliditySha3(
    { t: "bytes32", v: prevNodeHash },
    { t: "bytes32", v: nodeInnerHash }
  );
}

let empty_tuple_hash = web3.utils.soliditySha3({ t: "uint8", v: 3 });

async function makeEmptyAssertion(
  vm_state,
  num_steps,
  start_block,
  imported_message_count,
  read_inbox
) {
  return arb_rollup.makeAssertion(
    [
      vm_state,
      empty_tuple_hash,
      "0x00",
      "0x00",
      empty_tuple_hash,
      empty_tuple_hash,
      "0x00",
      "0x00",
      "0x00"
    ],
    0,
    0,
    0,
    num_steps,
    [start_block, start_block + 10],
    imported_message_count,
    read_inbox,
    0,
    []
  );
}

class VMProtoData {
  constructor(machineHash, inboxTop, inboxCount) {
    this.machineHash = machineHash;
    this.inboxTop = inboxTop;
    this.inboxCount = inboxCount;
  }

  hash() {
    return web3.utils.soliditySha3(
      { t: "bytes32", v: this.machineHash },
      { t: "bytes32", v: this.inboxTop },
      { t: "uint256", v: this.inboxCount }
    );
  }
}

class AssertionParams {
  constructor(numSteps, timeBounds, importedMessageCount) {
    this.numSteps = numSteps;
    this.timeBounds = timeBounds;
    this.importedMessageCount = importedMessageCount;
  }
}

class ExecutionAssertion {
  constructor(afterState, didReadInbox, numGas, outMessagesAcc, outLogsAcc) {
    this.afterState = afterState;
    this.didReadInbox = didReadInbox;
    this.numGas = numGas;
    this.outMessagesAcc = outMessagesAcc;
    this.outLogsAcc = outLogsAcc;
  }
}

class AssertionClaim {
  constructor(afterInboxTop, importedMessageSlice, executionAssertion) {
    this.afterInboxTop = afterInboxTop;
    this.importedMessageSlice = importedMessageSlice;
    this.executionAssertion = executionAssertion;
  }
}

class Assertion {
  constructor(
    blockNumber,
    inboxValue,
    inboxCount,
    prevPrevNode,
    prevProtoData,
    prevDeadline,
    prevDataHash,
    prevChildType,
    params,
    claims
  ) {
    this.blockNumber = blockNumber;
    this.inboxValue = inboxValue;
    this.inboxCount = inboxCount;

    this.prevPrevNode = prevPrevNode;
    this.prevProtoData = prevProtoData;
    this.prevDeadline = prevDeadline;
    this.prevDataHash = prevDataHash;
    this.prevChildType = prevChildType;
    this.params = params;
    this.claims = claims;
  }

  prevNodeHash() {
    return childNodeHash(
      this.prevPrevNode,
      this.prevDeadline,
      this.prevDataHash,
      this.prevChildType,
      this.prevProtoData.hash()
    );
  }

  deadline() {
    return 13000 * this.blockNumber + grace_period_ticks;
  }

  invalidInboxTopHashInner() {
    return childNodeInnerHash(
      this.deadline(),
      invalidInboxTopHash(
        this.claims.afterInboxTop,
        this.inboxValue,
        this.inboxCount -
          (this.prevPrevNode.inboxCount + this.params.importedMessageCount),
        grace_period_ticks + 13000
      ),
      0,
      this.prevProtoData.hash()
    );
  }

  invalidInboxTopHash() {
    return childNodeShortHash(
      this.prevNodeHash(),
      this.invalidInboxTopHashInner()
    );
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
        grace_period_ticks + 13000
      ),
      1,
      this.prevProtoData.hash()
    );
  }

  invalidMessagesHash() {
    return childNodeShortHash(
      this.prevNodeHash(),
      this.invalidMessagesHashInner()
    );
  }

  updatedProtoData() {
    return new VMProtoData(
      this.claims.executionAssertion.afterState,
      this.claims.afterInboxTop,
      this.prevProtoData.inboxCount + this.params.importedMessageCount
    );
  }

  validDataHash() {
    return web3.utils.soliditySha3(
      { t: "bytes32", v: this.claims.executionAssertion.outMessagesAcc },
      { t: "bytes32", v: this.claims.executionAssertion.outLogsAcc }
    );
  }

  validHashInner() {
    return childNodeInnerHash(
      this.deadline(),
      this.validDataHash(),
      3,
      this.updatedProtoData().hash()
    );
  }

  validHash() {
    return childNodeShortHash(this.prevNodeHash(), this.validHashInner());
  }
}

async function makeAssertion(
  prevPrevNode,
  prevProtoData,
  prevDeadline,
  prevDataHash,
  prevChildType,
  params,
  claims,
  stakerProof
) {
  let receipt = await arb_rollup.makeAssertion(
    [
      prevProtoData.machineHash,
      prevProtoData.inboxTop,
      prevPrevNode,
      prevDataHash,
      claims.afterInboxTop,
      claims.importedMessageSlice,
      claims.executionAssertion.afterState,
      claims.executionAssertion.outMessagesAcc,
      claims.executionAssertion.outLogsAcc
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
  );

  return {
    receipt: receipt,
    assertion: new Assertion(
      receipt.receipt.blockNumber,
      receipt.logs[0].args.inboxValue,
      0,
      prevPrevNode,
      prevProtoData,
      prevDeadline,
      prevDataHash,
      prevChildType,
      params,
      claims
    )
  };
}

let initial_vm_state = "0x99";
let stakeRequirement = 10;
let max_execution_steps = 50000;
let grace_period_ticks = 10000;

var arb_rollup;
var assertionInfo;

contract("ArbRollup", accounts => {
  it("should initialize", async () => {
    let challenge_factory = await ChallengeFactory.deployed();
    let global_inbox = await GlobalInbox.deployed();
    arb_rollup = await ArbRollup.new();
    await arb_rollup.init(
      initial_vm_state, // vmState
      grace_period_ticks, // gracePeriodTicks
      1000000, // arbGasSpeedLimitPerTick
      max_execution_steps, // maxExecutionSteps
      stakeRequirement, // stakeRequirement
      accounts[0], // owner
      challenge_factory.address,
      global_inbox.address
    );

    original_node = await arb_rollup.latestConfirmed();
  });

  it("should fail to assert on invalid leaf", async () => {
    let current_block = await web3.eth.getBlock("latest");
    await expectRevert(
      makeEmptyAssertion("0x34", 0, current_block.number, 0),
      "MAKE_LEAF"
    );
  });

  // it("should fail to assert on halted vm", async () => {
  //   truffleAssert.reverts(makeEmptyAssertion("0x00", 0, 0), "MAKE_RUN");
  // })

  it("should fail to assert over step limit", async () => {
    let current_block = await web3.eth.getBlock("latest");
    await expectRevert(
      makeEmptyAssertion(
        initial_vm_state,
        max_execution_steps + 1,
        current_block.number,
        0,
        false
      ),
      "MAKE_STEP"
    );
  });

  it("should fail to assert without stake", async () => {
    let current_block = await web3.eth.getBlock("latest");
    await expectRevert(
      makeEmptyAssertion(initial_vm_state, 0, current_block.number, 0, false),
      "INV_STAKER"
    );
  });

  it("should fail to assert outside time bounds", async () => {
    await expectRevert(
      makeEmptyAssertion(initial_vm_state, 0, 10000, 0, false),
      "MAKE_TIME"
    );
  });

  it("should fail if consuming messages but not reading inbox", async () => {
    let current_block = await web3.eth.getBlock("latest");
    await expectRevert(
      makeEmptyAssertion(initial_vm_state, 0, current_block.number, 10, false),
      "MAKE_MESSAGES"
    );
  });

  it("should fail if reading past lastest inbox message", async () => {
    let current_block = await web3.eth.getBlock("latest");
    await expectRevert(
      makeEmptyAssertion(initial_vm_state, 0, current_block.number, 10, true),
      "MAKE_MESSAGE_CNT"
    );
  });

  it("should create a stake", async () => {
    let receipt = await arb_rollup.placeStake([], [], {
      from: accounts[0],
      value: stakeRequirement
    });
    await expectEvent(receipt, "RollupStakeCreated");
    console.log("placeStake gas used:", receipt.receipt.gasUsed);
  });

  it("should make an assertion", async () => {
    assert.isTrue(
      await arb_rollup.isValidLeaf(original_node),
      "latest confirmed should be leaf before asserting"
    );
    let current_block = await web3.eth.getBlock("latest");
    let prevProtoData = new VMProtoData(initial_vm_state, empty_tuple_hash, 0);
    let params = new AssertionParams(
      0,
      [current_block.number, current_block.number + 10],
      0
    );
    let claims = new AssertionClaim(
      "0x00",
      empty_tuple_hash,
      new ExecutionAssertion("0x85", false, 0, "0x00", "0x00")
    );
    let info = await makeAssertion(
      "0x00",
      prevProtoData,
      0,
      "0x00",
      0,
      params,
      claims,
      []
    );

    assertionInfo = info.assertion;

    assert.isFalse(
      await arb_rollup.isValidLeaf(assertionInfo.prevNodeHash()),
      "original_node confirmed should be removed as leaf"
    );
    assert.isTrue(
      await arb_rollup.isValidLeaf(assertionInfo.invalidInboxTopHash()),
      "invalid inbox top should be leaf"
    );
    assert.isTrue(
      await arb_rollup.isValidLeaf(assertionInfo.invalidMessagesHash()),
      "invalid messages should be leaf"
    );
    // TODO: Check whether invalid execution is leaf
    assert.isTrue(
      await arb_rollup.isValidLeaf(assertionInfo.validHash()),
      "valid child should be leaf"
    );

    console.log("makeAssertion gas used:", info.receipt.receipt.gasUsed);
  });

  it("should confirm an assertion", async () => {
    let receipt = await arb_rollup.confirmValid(
      assertionInfo.deadline(),
      "0x",
      assertionInfo.claims.executionAssertion.outLogsAcc,
      assertionInfo.updatedProtoData().hash(),
      [accounts[0]],
      [],
      [0, 0]
    );
    await expectEvent(receipt, "RollupConfirmed");

    assert.equal(
      await arb_rollup.latestConfirmed(),
      assertionInfo.validHash(),
      "latest confirmed should now be valid child"
    );

    assert.isTrue(
      await arb_rollup.isValidLeaf(assertionInfo.validHash()),
      "invalid inbox top should be leaf"
    );

    console.log("confirmValid gas used:", receipt.receipt.gasUsed);
  });

  it("should prune a leaf", async () => {
    assert.isTrue(
      await arb_rollup.isValidLeaf(assertionInfo.invalidInboxTopHash()),
      "invalid messages should be leaf"
    );
    let receipt = await arb_rollup.pruneLeaf(
      original_node,
      [assertionInfo.invalidInboxTopHashInner()],
      [assertionInfo.validHashInner()]
    );
    await expectEvent(receipt, "RollupPruned");
    assert.isFalse(
      await arb_rollup.isValidLeaf(assertionInfo.invalidInboxTopHash()),
      "invalid messages should be leaf"
    );
    console.log("pruneLeaf gas used:", receipt.receipt.gasUsed);
  });

  it("should assert again", async () => {
    let current_block = await web3.eth.getBlock("latest");
    let params = new AssertionParams(
      0,
      [current_block.number, current_block.number + 10],
      0
    );
    let claims = new AssertionClaim(
      "0x00",
      empty_tuple_hash,
      new ExecutionAssertion("0x00", false, 0, "0x00", "0x00")
    );

    let info = await makeAssertion(
      assertionInfo.prevNodeHash(),
      assertionInfo.updatedProtoData(),
      assertionInfo.deadline(),
      assertionInfo.validDataHash(),
      3,
      params,
      claims,
      []
    );

    console.log("makeAssertion gas used:", info.receipt.receipt.gasUsed);
  });
});
