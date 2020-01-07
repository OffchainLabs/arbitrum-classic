const ArbRollup = artifacts.require("ArbRollup");
const ChallengeFactory = artifacts.require("ChallengeFactory");
const GlobalPendingInbox = artifacts.require("GlobalPendingInbox");

const { expectEvent, expectRevert } = require("@openzeppelin/test-helpers");

function pendingTopHash(lowerHash, topHash, chainLength) {
  return web3.utils.soliditySha3(
    { t: "bytes32", v: lowerHash },
    { t: "bytes32", v: topHash },
    { t: "uint256", v: chainLength }
  );
}

function invalidPendingTopHash(
  lowerHash,
  topHash,
  chainLength,
  challengePeriod
) {
  return web3.utils.soliditySha3(
    { t: "bytes32", v: pendingTopHash(lowerHash, topHash, chainLength) },
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

function validHash(messagesAcc, logsAcc) {
  return web3.utils.soliditySha3(
    { t: "bytes32", v: messagesAcc },
    { t: "bytes32", v: logsAcc }
  );
}

function protoStateHash(machineHash, pendingTop, pendingCount) {
  return web3.utils.soliditySha3(
    { t: "bytes32", v: machineHash },
    { t: "bytes32", v: pendingTop },
    { t: "uint256", v: pendingCount }
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
      v: web3.utils.soliditySha3(
        { t: "bytes32", v: vmProtoStateHash },
        { t: "uint256", v: deadlineTicks },
        { t: "bytes32", v: nodeDataHash },
        { t: "uint256", v: childType }
      )
    }
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

let initial_vm_state = "0x99";
let stakeRequirement = 10;
let max_execution_steps = 50000;
let grace_period_ticks = 10000;

var arb_rollup;
var deadline;

contract("ArbRollup", accounts => {
  it("should initialize", async () => {
    let challenge_factory = await ChallengeFactory.deployed();
    let global_inbox = await GlobalPendingInbox.deployed();
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

  it("should fail if reading past lastest pending", async () => {
    let current_block = await web3.eth.getBlock("latest");
    await expectRevert(
      makeEmptyAssertion(initial_vm_state, 0, current_block.number, 10, true),
      "MAKE_MESSAGE_CNT"
    );
  });

  it("should create a stake", async () => {
    let latest = await arb_rollup.latestConfirmed();
    await expectEvent(
      await arb_rollup.placeStake(latest, latest, [], [], {
        from: accounts[0],
        value: stakeRequirement
      }),
      "RollupStakeCreated"
    );
  });

  it("should make an assertion", async () => {
    let latest = await arb_rollup.latestConfirmed();
    assert.isTrue(
      await arb_rollup.isValidLeaf(latest),
      "latest confirmed should be leaf before asserting"
    );
    let current_block = await web3.eth.getBlock("latest");
    let tx = await makeEmptyAssertion(
      initial_vm_state,
      0,
      current_block.number,
      0,
      false
    );
    let test = await expectEvent(tx, "RollupAsserted");
    deadline = 13000 * tx.receipt.blockNumber + grace_period_ticks;

    let invalid_pending_top_hash_val = childNodeHash(
      latest,
      deadline,
      invalidPendingTopHash(
        empty_tuple_hash,
        empty_tuple_hash,
        0,
        grace_period_ticks + 13000
      ),
      0,
      protoStateHash(initial_vm_state, empty_tuple_hash, 0)
    );
    let invalid_messages_hash_val = childNodeHash(
      latest,
      deadline,
      invalidMessagesHash(
        empty_tuple_hash,
        empty_tuple_hash,
        empty_tuple_hash,
        empty_tuple_hash,
        0,
        grace_period_ticks + 13000
      ),
      1,
      protoStateHash(initial_vm_state, empty_tuple_hash, 0)
    );
    let valid_child_hash = childNodeHash(
      latest,
      deadline,
      validHash("0x00", "0x00"),
      3,
      protoStateHash("0x00", empty_tuple_hash, 0)
    );
    assert.isFalse(
      await arb_rollup.isValidLeaf(latest),
      "latest confirmed should be removed as leaf"
    );
    assert.isTrue(
      await arb_rollup.isValidLeaf(invalid_pending_top_hash_val),
      "invalid pending top should be leaf"
    );
    assert.isTrue(
      await arb_rollup.isValidLeaf(invalid_messages_hash_val),
      "invalid messages should be leaf"
    );
    // TODO: Check whether invalid execution is leaf
    assert.isTrue(
      await arb_rollup.isValidLeaf(valid_child_hash),
      "valid child should be leaf"
    );
  });

  it("should confirm an assertion", async () => {
    let latest = await arb_rollup.latestConfirmed();

    let current_block = await web3.eth.getBlock("latest");
    await expectEvent(
      await arb_rollup.confirmValid(
        deadline,
        "0x",
        "0x00",
        protoStateHash("0x00", empty_tuple_hash, 0),
        [accounts[0]],
        [],
        [0, 0]
      ),
      "RollupConfirmed"
    );
  });
});
