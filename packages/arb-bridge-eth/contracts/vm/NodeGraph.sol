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

pragma solidity ^0.5.3;

import "./RollupUtils.sol";
import "./VM.sol";
import "../IGlobalPendingInbox.sol";

import "../challenge/ChallengeUtils.sol";
import "../challenge/ChallengeType.sol";

import "../arch/Value.sol";
import "../arch/Protocol.sol";

import "../libraries/RollupTime.sol";


contract NodeGraph is ChallengeType {

    using SafeMath for uint256;

    // invalid leaf
    string constant MAKE_LEAF = "MAKE_LEAF";
    // Can only disputable assert if machine is not errored or halted
    string constant MAKE_RUN = "MAKE_RUN";
    // Tried to execute too many steps
    string constant MAKE_STEP = "MAKE_STEP";
    // Precondition: not within time bounds
    string constant MAKE_TIME = "MAKE_TIME";

    uint256 constant VALID_CHILD_TYPE = 3;
    uint256 constant MAX_CHILD_TYPE = 3;

    // Fields
    //   prevLeafHash
    //   afterPendingTop
    //   importedMessagesSlice
    //   afterVMHash
    //   messagesAccHash
    //   logsAccHash

    event RollupAsserted(
        bytes32[6] fields,
        uint256 importedMessageCount,
        uint128[2] timeBoundsBlocks,
        bool didInboxInsn,
        uint32 numSteps,
        uint64 numArbGas
    );

    event RollupConfirmed(bytes32 nodeHash);

    IGlobalPendingInbox public globalInbox;
    VM.Params vmParams;
    mapping (bytes32 => bool) private leaves;
    bytes32 private latestConfirmedPriv;

    event RollupPruned(bytes32 leaf);

    struct MakeAssertionData {
        bytes32 beforeVMHash;
        bytes32 beforeInboxHash;
        bytes32 beforePendingTop;
        uint256 beforePendingCount;

        bytes32 prevPrevLeafHash;
        uint256 prevDeadlineTicks;
        bytes32 prevDisputableNodeHash;
        uint32  prevChildType;

        uint32 numSteps;
        uint128[2] timeBoundsBlocks;
        uint256 afterPendingCount;

        bytes32 afterPendingTop;

        bytes32 importedMessagesSlice;

        bytes32 afterVMHash;
        bool didInboxInsn;
        uint64 numArbGas;
        bytes32 messagesAccHash;
        bytes32 logsAccHash;
    }

    function pruneLeaf(
        bytes32 _leaf,
        bytes32 from,
        bytes32[] calldata leafProof,
        bytes32[] calldata latestConfirmedProof
    )
        external
    {
        require(isValidLeaf(_leaf), "invalid leaf");
        require(
            RollupUtils.isConflict(
                from,
                _leaf,
                latestConfirmed(),
                leafProof,
                latestConfirmedProof
            ),
            "Invalid conflict proof"
        );
        delete leaves[_leaf];

        emit RollupPruned(_leaf);
    }


    function init(
        bytes32 _vmState,
        uint128 _gracePeriodTicks,
        uint128 _arbGasSpeedLimitPerTick,
        uint32 _maxExecutionSteps,
        address _globalInboxAddress
    )
        internal
    {
        globalInbox = IGlobalPendingInbox(_globalInboxAddress);

        globalInbox.registerForInbox();

        // VM protocol state
        bytes32 vmProtoStateHash = RollupUtils.protoStateHash(_vmState, Value.hashEmptyTuple(), Value.hashEmptyTuple(), 0);
        bytes32 initialNode = RollupUtils.childNodeHash(
            0,
            0,
            0,
            0,
            vmProtoStateHash
        );
        latestConfirmedPriv = initialNode;
        leaves[initialNode] = true;

        // VM parameters
        vmParams.gracePeriodTicks = _gracePeriodTicks;
        vmParams.arbGasSpeedLimitPerTick = _arbGasSpeedLimitPerTick;
        vmParams.maxExecutionSteps = _maxExecutionSteps;
    }

    function _computeDeadline(
        uint256 checkTimeTicks,
        uint256 gracePeriodTicks,
        uint256 prevDeadlineTicks
    ) private view returns(uint256) {
        uint256 deadlineTicks = RollupTime.blocksToTicks(block.number) + gracePeriodTicks ;
        if (deadlineTicks >= prevDeadlineTicks) {
            return deadlineTicks + checkTimeTicks;
        } else {
            return prevDeadlineTicks + checkTimeTicks;
        }
    }

    function makeAssertion(MakeAssertionData memory data) internal returns(bytes32, bytes32) {
        bytes32 vmProtoHashBefore = RollupUtils.protoStateHash(
            data.beforeVMHash,
            data.beforeInboxHash,
            data.beforePendingTop,
            data.beforePendingCount
        );
        bytes32 prevLeaf = RollupUtils.childNodeHash(
            data.prevPrevLeafHash,
            data.prevDeadlineTicks,
            data.prevDisputableNodeHash,
            data.prevChildType,
            vmProtoHashBefore
        );
        require(isValidLeaf(prevLeaf), MAKE_LEAF);
        require(!VM.isErrored(data.beforeVMHash) && !VM.isHalted(data.beforeVMHash), MAKE_RUN);
        require(data.numSteps <= vmParams.maxExecutionSteps, MAKE_STEP);
        require(VM.withinTimeBounds(data.timeBoundsBlocks), MAKE_TIME);

        uint256 deadlineTicks = _computeDeadline(
            data.numArbGas / vmParams.arbGasSpeedLimitPerTick,
            vmParams.gracePeriodTicks,
            data.prevDeadlineTicks
        );
        (bytes32 pendingValue, uint256 pendingCount) = globalInbox.getPending();
        bytes32 execBeforeInboxHash = Protocol.addMessagesToInbox(data.beforeInboxHash, data.importedMessagesSlice);

        leaves[generateInvalidPendingTopLeaf(
            data,
            prevLeaf,
            deadlineTicks,
            pendingValue,
            pendingCount,
            vmProtoHashBefore
        )] = true;
        leaves[generateInvalidMessagesLeaf(
            data,
            prevLeaf,
            deadlineTicks,
            vmProtoHashBefore
        )] = true;
        leaves[generateInvalidExecutionLeaf(
            data,
            prevLeaf,
            deadlineTicks,
            execBeforeInboxHash,
            vmProtoHashBefore
        )] = true;
        if (data.didInboxInsn) {
            execBeforeInboxHash = Value.hashEmptyTuple();
        }
        bytes32 validHash = generateValidLeaf(
            data,
            prevLeaf,
            deadlineTicks,
            execBeforeInboxHash
        );
        leaves[validHash] = true;
        delete leaves[prevLeaf];

        emit RollupAsserted(
            [
                prevLeaf,
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
        return (prevLeaf, validHash);
    }

    function confirmNode(bytes32 to) internal {
        latestConfirmedPriv = to;
        emit RollupConfirmed(to);
    }

    function latestConfirmed() internal view returns (bytes32) {
        return latestConfirmedPriv;
    }

    function isValidLeaf(bytes32 leaf) internal view returns(bool) {
        return leaves[leaf];
    }

    function generateInvalidPendingTopLeaf(
        MakeAssertionData memory data,
        bytes32 prevLeaf,
        uint256 deadlineTicks,
        bytes32 pendingValue,
        uint256 pendingCount,
        bytes32 vmProtoHashBefore
    )
        private
        view
        returns(bytes32)
    {
        return RollupUtils.childNodeHash(
            prevLeaf,
            deadlineTicks,
            keccak256(
                abi.encodePacked(
                    ChallengeUtils.pendingTopHash(
                        data.afterPendingTop,
                        pendingValue,
                        pendingCount.sub(data.afterPendingCount)
                    ),
                    vmParams.gracePeriodTicks + RollupTime.blocksToTicks(1)
                )
            ),
            INVALID_PENDING_TOP_TYPE,
            vmProtoHashBefore
        );
    }

    function generateInvalidMessagesLeaf(
        MakeAssertionData memory data,
        bytes32 prevLeaf,
        uint256 deadlineTicks,
        bytes32 vmProtoHashBefore
    )
        private
        view
        returns(bytes32)
    {
        return RollupUtils.childNodeHash(
            prevLeaf,
            deadlineTicks,
            keccak256(
                abi.encodePacked(
                    ChallengeUtils.messagesHash(
                        data.beforePendingTop,
                        data.afterPendingTop,
                        0x00,
                        data.importedMessagesSlice,
                        data.afterPendingCount.sub(data.beforePendingCount)
                    ),
                    vmParams.gracePeriodTicks + RollupTime.blocksToTicks(1)
                )
            ),
            INVALID_MESSAGES_TYPE,
            vmProtoHashBefore
        );
    }

    function generateInvalidExecutionLeaf(
        MakeAssertionData memory data,
        bytes32 prevLeaf,
        uint256 deadlineTicks,
        bytes32 execBeforeInboxHash,
        bytes32 vmProtoHashBefore
    )
        private
        view
        returns(bytes32)
    {
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
        return RollupUtils.childNodeHash(
            prevLeaf,
            deadlineTicks,
            keccak256(
                abi.encodePacked(
                    ChallengeUtils.executionHash(
                        Protocol.generatePreconditionHash(
                             data.beforeVMHash,
                             data.timeBoundsBlocks,
                             execBeforeInboxHash
                        ),
                        assertionHash
                    ),
                    vmParams.gracePeriodTicks + data.numArbGas / vmParams.arbGasSpeedLimitPerTick
                )
            ),
            INVALID_EXECUTION_TYPE,
            vmProtoHashBefore
        );
    }

    function generateValidLeaf(
        MakeAssertionData memory data,
        bytes32 prevLeaf,
        uint256 deadlineTicks,
        bytes32 execBeforeInboxHash
    )
        private
        pure
        returns(bytes32)
    {
        return RollupUtils.childNodeHash(
            prevLeaf,
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
    }
}
