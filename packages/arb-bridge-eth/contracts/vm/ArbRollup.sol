/*
 * Copyright 2019, Offchain Labs, Inc.
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

pragma solidity ^0.5.3;

import "./Leaves.sol";

import "./VM.sol";
import "./IArbRollup.sol";
import "../libraries/RollupTime.sol";

import "../IGlobalPendingInbox.sol";

import "../arch/Value.sol";


contract ArbRollup is Leaves, IArbRollup {
    // invalid leaf
    string constant MAKE_LEAF = "MAKE_LEAF";
    // Can only disputable assert if machine is not errored or halted
    string constant MAKE_RUN = "MAKE_RUN";
    // Tried to execute too many steps
    string constant MAKE_STEP = "MAKE_STEP";
    // Precondition: not within time bounds
    string constant MAKE_TIME = "MAKE_TIME";
    // invalid staker location proof
    string constant MAKE_STAKER_PROOF = "MAKE_STAKER_PROOF";

    // must include proof for all stakers
    string constant CONF_COUNT = "CONF_COUNT";
    // Stakers must be ordered
    string constant CONF_ORDER = "CONF_ORDER";
    // at least one active staker disagrees
    string constant CONF_STAKER_PROOF = "CONF_STAKER_PROOF";
    // Type is not invalid
    string constant CONF_INV_TYPE = "CONF_INV_TYPE";
    // There must be at least one staker
    string constant CONF_HAS_STAKER = "CONF_HAS_STAKER";


    // Only callable by owner
    string constant ONLY_OWNER = "ONLY_OWNER";

    using SafeMath for uint256;

    IGlobalPendingInbox public globalInbox;

    address   owner;
    VM.Params vmParams;


    // Fields
    //   prevLeafHash
    //   afterPendingTop
    //   importedMesssagesSlice
    //   afterVMHash
    //   messagesAccHash
    //   logsAccHash

    event RollupAsserted(
        bytes32[6] fields,
        uint importedMessageCount,
        uint64[2] timeBoundsBlocks,
        bool didInboxInsn,
        uint32 numSteps,
        uint64 numArbGas
    );

    event RollupConfirmed(bytes32 nodeHash);

    event ConfirmedAssertion(
        bytes32 logsAccHash
    );

    struct MakeAssertionData {
        bytes32 beforeVMHash;
        bytes32 beforeInboxHash;
        bytes32 beforePendingTop;
        uint beforePendingCount;
        bytes32 prevPrevLeafHash;
        bytes32 prevDisputableNodeHash;
        bytes32[] stakerProof;
        bytes32 afterPendingTop;
        uint afterPendingCount;
        bytes32 importedMessagesSlice;
        bytes32 afterVMHash;
        bool didInboxInsn;
        bytes32 afterInboxHash;
        bytes32 messagesAccHash;
        bytes32 logsAccHash;
        uint32 numSteps;
        uint64 numArbGas;
        uint128[2] timeBoundsBlocks;
        uint256 prevDeadlineTicks;
        uint32  prevChildType;
    }

    function init(
        bytes32 _vmState,
        uint128 _gracePeriodTicks,
        uint128 _arbGasSpeedLimitPerTick,
        uint32 _maxExecutionSteps,
        uint128 _stakeRequirement,
        address payable _owner,
        address _challengeFactoryAddress,
        address _globalInboxAddress
    )
        external
    {
        Leaves.init(_vmState, _stakeRequirement, _challengeFactoryAddress);

        globalInbox = IGlobalPendingInbox(_globalInboxAddress);

        globalInbox.registerForInbox();
        owner = _owner;

        // VM parameters
        vmParams.gracePeriodTicks = _gracePeriodTicks;
        vmParams.arbGasSpeedLimitPerTick = _arbGasSpeedLimitPerTick;
        vmParams.maxExecutionSteps = _maxExecutionSteps;
        vmParams.pendingInboxHash = Value.hashEmptyTuple();
    }

    // fields
    //  beforeVMHash
    //  beforeInboxHash
    //  beforePendingTop
    //  prevPrevLeafHash
    //  prevDisputableNodeHash
    //  afterPendingTop
    //  importedMessagesSlice
    //  afterVMHash
    //  afterInboxHash
    //  messagesAccHash
    //  logsAccHash

    function makeAssertion(
        bytes32[11] calldata _fields,
        bytes32[] calldata _stakerProof,
        uint _beforePendingCount,
        uint _afterPendingCount,
        bool _didInboxInsn,
        uint32 _numSteps,
        uint64 _numArbGas,
        uint128[2] calldata _timeBoundsBlocks,
        uint256 _prevDeadlineTicks,
        uint32 _prevChildType
    )
        external
    {
        return _makeAssertion(
            MakeAssertionData(
                _fields[0],
                _fields[1],
                _fields[2],
                _beforePendingCount,
                _fields[3],
                _fields[4],
                _stakerProof,
                _fields[5],
                _afterPendingCount,
                _fields[6],
                _fields[7],
                _didInboxInsn,
                _fields[8],
                _fields[9],
                _fields[10],
                _numSteps,
                _numArbGas,
                _timeBoundsBlocks,
                _prevDeadlineTicks,
                _prevChildType
            )
        );
    }

    function confirmValid(
        uint deadlineTicks,
        bytes calldata _messages,
        bytes32 logsAcc,
        bytes32 vmProtoStateHash,
        address[] calldata stakerAddresses,
        bytes32[] calldata stakerProofs,
        uint[]  calldata stakerProofOffsets
    )
        external
    {
        _confirmNode(
            deadlineTicks,
            RollupUtils.validNodeHash(
                Protocol.generateLastMessageHash(_messages),
                logsAcc
            ),
            VALID_CHILD_TYPE,
            vmProtoStateHash,
            stakerAddresses,
            stakerProofs,
            stakerProofOffsets
        );

        globalInbox.sendMessages(_messages);

        emit ConfirmedAssertion(
            logsAcc
        );
    }

    function confirmInvalid(
        uint    deadlineTicks,
        bytes32 challengeNodeData,
        uint    branch,
        bytes32 vmProtoStateHash,
        address[] calldata stakerAddresses,
        bytes32[] calldata stakerProofs,
        uint[]  calldata stakerProofOffsets
    )
        external
    {
        require(branch < VALID_CHILD_TYPE, CONF_INV_TYPE);
        _confirmNode(
            deadlineTicks,
            challengeNodeData,
            branch,
            vmProtoStateHash,
            stakerAddresses,
            stakerProofs,
            stakerProofOffsets
        );
    }

    modifier onlyOwner() {
        require(msg.sender == owner, ONLY_OWNER);
        _;
    }

/*    function activateVM() external onlyOwner {
        if (vm.state == VM.State.Uninitialized) {
            vm.state = VM.State.Waiting;
        }
    }

    function ownerShutdown() external onlyOwner {
        _shutdown();
    }
    */

    struct MakeAssertionFrame {
        bytes32 vmProtoHashBefore;
        bytes32 prevLeaf;
        bytes32 pendingValue;
        uint pendingCount;
    }

    function _makeAssertion(MakeAssertionData memory data) private {
        MakeAssertionFrame memory frame;
        frame.vmProtoHashBefore = RollupUtils.protoStateHash(
            data.beforeVMHash,
            data.beforeInboxHash,
            data.beforePendingTop,
            data.beforePendingCount
        );
        frame.prevLeaf = RollupUtils.childNodeHash(
            data.prevPrevLeafHash,
            data.prevDeadlineTicks,
            data.prevDisputableNodeHash,
            data.prevChildType,
            frame.vmProtoHashBefore
        );
        require(isValidLeaf(frame.prevLeaf), MAKE_LEAF);
        require(!VM.isErrored(data.beforeVMHash) && !VM.isHalted(data.beforeVMHash), MAKE_RUN);
        require(data.numSteps <= vmParams.maxExecutionSteps, MAKE_STEP);
        require(withinTimeBounds(data.timeBoundsBlocks), MAKE_TIME);

        Staker storage staker = getValidStaker(msg.sender);
        require(RollupUtils.isPath(staker.location, frame.prevLeaf, data.stakerProof), MAKE_STAKER_PROOF);

        uint deadlineTicks = RollupTime.blocksToTicks(uint128(block.number)) + vmParams.gracePeriodTicks + data.numArbGas/vmParams.arbGasSpeedLimitPerTick;
        (frame.pendingValue, frame.pendingCount) = globalInbox.getPending();
        bytes32[] memory leaves = new bytes32[](MAX_CHILD_TYPE);
        leaves[INVALID_PENDING_TOP_TYPE] = RollupUtils.childNodeHash(
            frame.prevLeaf,
            deadlineTicks,
            ChallengeUtils.pendingTopHash(
                frame.pendingValue,
                data.afterPendingTop,
                frame.pendingCount.sub(data.afterPendingCount)
            ),
            INVALID_PENDING_TOP_TYPE,
            frame.vmProtoHashBefore
        );
        leaves[INVALID_MESSAGES_TYPE] = RollupUtils.childNodeHash(
            frame.prevLeaf,
            deadlineTicks,
            ChallengeUtils.messagesHash(
                data.beforePendingTop,
                data.afterPendingTop,
                0x00,
                data.importedMessagesSlice,
                data.afterPendingCount.sub(data.beforePendingCount)
            ),
            INVALID_MESSAGES_TYPE,
            frame.vmProtoHashBefore
        );
        bytes32 assertionHash = Protocol.generateAssertionHash(
            data.afterVMHash,
            data.didInboxInsn,
            data.numSteps,
            data.numArbGas,
            0x00,
            data.messagesAccHash,
            0x00,
            data.logsAccHash
        );
        bytes32 execBeforeInboxHash = Protocol.addMessagesToInbox(data.beforeInboxHash, data.importedMessagesSlice);
        leaves[INVALID_EXECUTION_TYPE] = RollupUtils.childNodeHash(
            frame.prevLeaf,
            deadlineTicks,
            ChallengeUtils.executionHash(
                Protocol.generatePreconditionHash(
                     data.beforeVMHash,
                     data.timeBoundsBlocks,
                     execBeforeInboxHash
                ),
                assertionHash
            ),
            INVALID_EXECUTION_TYPE,
            frame.vmProtoHashBefore
        );
        if (data.didInboxInsn) {
            execBeforeInboxHash = Value.hashEmptyTuple();
        }
        leaves[VALID_CHILD_TYPE] = RollupUtils.childNodeHash(
            frame.prevLeaf,
            deadlineTicks,
            RollupUtils.validNodeHash(
                data.messagesAccHash,
                data.logsAccHash
            ),
            VALID_CHILD_TYPE,
            RollupUtils.protoStateHash(
                data.afterVMHash,
                execBeforeInboxHash,
                data.afterPendingTop,
                data.afterPendingCount
            )
        );
        splitLeaf(frame.prevLeaf, leaves);
        staker.location = leaves[VALID_CHILD_TYPE];

        emit RollupAsserted(
            [
                frame.prevLeaf,
                data.afterPendingTop,
                data.importedMessagesSlice,
                data.afterVMHash,
                data.messagesAccHash,
                data.logsAccHash
            ],
            data.afterPendingCount.sub(data.beforePendingCount),
            data.timeBoundsBlocks,
            data.didInboxInsn,
            data.numSteps,
            data.numArbGas
        );
    }

    function _confirmNode(
        uint deadline,
        bytes32 nodeDataHash,
        uint branch,
        bytes32 vmProtoStateHash,
        address[] memory stakerAddresses,
        bytes32[] memory stakerProofs,
        uint[]  memory stakerProofOffsets
    )
        private
    {
        uint _stakerCount = stakerAddresses.length;
        require(_stakerCount == getStakerCount(), CONF_COUNT);
        bytes32 to = RollupUtils.childNodeHash(
            latestConfirmed(),
            deadline,
            nodeDataHash,
            branch,
            vmProtoStateHash
        );
        bytes20 prevStaker = 0x00;
        bool hasStaker = false;
        for (uint i = 0; i < _stakerCount; i++) {
            address stakerAddress = stakerAddresses[i];
            require(bytes20(stakerAddress) > prevStaker, CONF_ORDER);
            Staker storage staker = getValidStaker(stakerAddress);
            if (staker.creationTime >= deadline) {
                require(
                    RollupUtils.isPathOffset(
                        to,
                        staker.location,
                        stakerProofs,
                        stakerProofOffsets[i],
                        stakerProofOffsets[i+1]
                    ),
                    CONF_STAKER_PROOF
                );
                hasStaker = true;
            }
            prevStaker = bytes20(stakerAddress);
        }
        require(hasStaker, CONF_HAS_STAKER);

        updateLatestConfirmed(to);

        emit RollupConfirmed(to);
    }

    function withinTimeBounds(uint128[2] memory _timeBoundsBlocks) private view returns (bool) {
        return block.number >= _timeBoundsBlocks[0] && block.number <= _timeBoundsBlocks[1];
    }
}
