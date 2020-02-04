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
    // Imported messages without reading them
    string constant MAKE_MESSAGES = "MAKE_MESSAGES";
    // Tried to import more messages than exist in pending inbox
    string constant MAKE_MESSAGE_CNT = "MAKE_MESSAGE_CNT";

    string constant PRUNE_LEAF = "PRUNE_LEAF";
    string constant PRUNE_PROOFLEN = "PRUNE_PROOFLEN";
    string constant PRUNE_CONFLICT = "PRUNE_CONFLICT";

    uint256 constant VALID_CHILD_TYPE = 3;
    uint256 constant MAX_CHILD_TYPE = 3;

    // Fields
    //  prevLeaf
    //  pendingValue
    //  afterPendingTop
    //  importedMessagesSlice
    //  afterVMHash
    //  messagesAccHash
    //  logsAccHash

    event RollupAsserted(
        bytes32[7] fields,
        uint256 pendingCount,
        uint256 importedMessageCount,
        uint128[2] timeBoundsBlocks,
        uint64 numArbGas,
        uint64 numSteps,
        bool didInboxInsn
    );

    event RollupConfirmed(bytes32 nodeHash);

    event RollupPruned(bytes32 leaf);

    event RollupCreated(bytes32 initVMHash);

    IGlobalPendingInbox public globalInbox;
    VM.Params public vmParams;
    mapping (bytes32 => bool) private leaves;
    bytes32 private latestConfirmedPriv;

    struct MakeAssertionData {
        bytes32 beforeVMHash;
        bytes32 beforePendingTop;
        uint256 beforePendingCount;

        bytes32 prevPrevLeafHash;
        uint256 prevDeadlineTicks;
        bytes32 prevDataHash;
        uint32  prevChildType;

        uint64 numSteps;
        uint128[2] timeBoundsBlocks;
        uint256 importedMessageCount;

        bytes32 afterPendingTop;

        bytes32 importedMessagesSlice;

        bytes32 afterVMHash;
        bool didInboxInsn;
        uint64 numArbGas;
        bytes32 messagesAccHash;
        bytes32 logsAccHash;
    }

    function pruneLeaves(
        bytes32[] memory fromNodes,
        bytes32[] memory leafProofs,
        uint256[] memory leafProofLengths,
        bytes32[] memory latestConfProofs,
        uint256[] memory latestConfirmedProofLengths
    )
        public
    {
        uint pruneCount = fromNodes.length;

        require(
            leafProofLengths.length == pruneCount &&
            latestConfirmedProofLengths.length == pruneCount,
            "input length mistmatch"
        );
        uint256 prevLeafOffset = 0;
        uint256 prevConfOffset = 0;
        for (uint256 i = 0; i < pruneCount; i++) {
            bytes32 from = fromNodes[i];
            require(leafProofLengths[i] > 0 && latestConfirmedProofLengths[i] > 0, PRUNE_PROOFLEN);
            uint256 nextLeafOffset = prevLeafOffset + leafProofLengths[i];
            uint256 nextConfOffset = prevConfOffset + latestConfirmedProofLengths[i];

            // If the function call was produced valid at any point, either all these checks will pass or all will fail
            require(
                leafProofs[prevLeafOffset] != latestConfProofs[prevConfOffset] &&
                RollupUtils.calculatePathOffset(from, latestConfProofs, prevConfOffset, nextConfOffset) == latestConfirmed(),
                PRUNE_CONFLICT
            );

            bytes32 leaf = RollupUtils.calculatePathOffset(from, leafProofs, prevLeafOffset, nextLeafOffset);
            if (isValidLeaf(leaf)) {
                delete leaves[leaf];
                emit RollupPruned(leaf);
            }

            prevLeafOffset = nextLeafOffset;
            prevConfOffset = nextConfOffset;
        }
    }

    function latestConfirmed() public view returns (bytes32) {
        return latestConfirmedPriv;
    }

    function isValidLeaf(bytes32 leaf) public view returns(bool) {
        return leaves[leaf];
    }

    function init(
        bytes32 _vmState,
        uint128 _gracePeriodTicks,
        uint128 _arbGasSpeedLimitPerTick,
        uint64 _maxExecutionSteps,
        uint64 _maxTimeBoundsWidth,
        address _globalInboxAddress
    )
        internal
    {
        globalInbox = IGlobalPendingInbox(_globalInboxAddress);

        // VM protocol state
        bytes32 vmProtoStateHash = RollupUtils.protoStateHash(
            _vmState,
            Value.hashEmptyTuple(),
            0
        );
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
        vmParams.maxTimeBoundsWidth = _maxTimeBoundsWidth;

        emit RollupCreated(_vmState);
    }

    function makeAssertion(MakeAssertionData memory data) internal returns(bytes32, bytes32) {
        bytes32 vmProtoHashBefore = RollupUtils.protoStateHash(
            data.beforeVMHash,
            data.beforePendingTop,
            data.beforePendingCount
        );
        bytes32 prevLeaf = RollupUtils.childNodeHash(
            data.prevPrevLeafHash,
            data.prevDeadlineTicks,
            data.prevDataHash,
            data.prevChildType,
            vmProtoHashBefore
        );
        require(isValidLeaf(prevLeaf), MAKE_LEAF);
        require(!VM.isErrored(data.beforeVMHash) && !VM.isHalted(data.beforeVMHash), MAKE_RUN);
        require(data.numSteps <= vmParams.maxExecutionSteps, MAKE_STEP);
        require(data.timeBoundsBlocks[1] <= data.timeBoundsBlocks[0]+vmParams.maxTimeBoundsWidth);
        require(VM.withinTimeBounds(data.timeBoundsBlocks), MAKE_TIME);
        require(data.importedMessageCount == 0 || data.didInboxInsn, MAKE_MESSAGES);

        (bytes32 pendingValue, uint256 pendingCount) = globalInbox.getPending();
        require(data.importedMessageCount <= pendingCount.sub(data.beforePendingCount), MAKE_MESSAGE_CNT);

        uint256 gracePeriodTicks = vmParams.gracePeriodTicks;
        uint256 checkTimeTicks = data.numArbGas / vmParams.arbGasSpeedLimitPerTick;
        uint256 deadlineTicks = RollupTime.blocksToTicks(block.number) + gracePeriodTicks;
        if (deadlineTicks < data.prevDeadlineTicks) {
            deadlineTicks = data.prevDeadlineTicks;
        }
        deadlineTicks += checkTimeTicks;

        bytes32 invalidPending = generateInvalidPendingTopLeaf(
            data,
            prevLeaf,
            deadlineTicks,
            pendingValue,
            pendingCount,
            vmProtoHashBefore,
            gracePeriodTicks
        );
        bytes32 invalidMessages = generateInvalidMessagesLeaf(
            data,
            prevLeaf,
            deadlineTicks,
            vmProtoHashBefore,
            gracePeriodTicks
        );
        bytes32 invalidExec = generateInvalidExecutionLeaf(
            data,
            prevLeaf,
            deadlineTicks,
            vmProtoHashBefore,
            gracePeriodTicks,
            checkTimeTicks
        );
        bytes32 validHash = generateValidLeaf(
            data,
            prevLeaf,
            deadlineTicks
        );

        leaves[invalidPending] = true;
        leaves[invalidMessages] = true;
        leaves[invalidExec] = true;
        leaves[validHash] = true;
        delete leaves[prevLeaf];

        emitAssertedEvent(data, prevLeaf, pendingValue, pendingCount);
        return (prevLeaf, validHash);
    }

    function confirmNode(bytes32 to) internal {
        latestConfirmedPriv = to;
        emit RollupConfirmed(to);
    }

    function emitAssertedEvent(MakeAssertionData memory data, bytes32 prevLeaf, bytes32 pendingValue, uint256 pendingCount) private {
        emit RollupAsserted(
            [
                prevLeaf,
                pendingValue,
                data.afterPendingTop,
                data.importedMessagesSlice,
                data.afterVMHash,
                data.messagesAccHash,
                data.logsAccHash
            ],
            pendingCount,
            data.importedMessageCount,
            data.timeBoundsBlocks,
            data.numArbGas,
            data.numSteps,
            data.didInboxInsn
        );
    }

    function generateInvalidPendingTopLeaf(
        MakeAssertionData memory data,
        bytes32 prevLeaf,
        uint256 deadlineTicks,
        bytes32 pendingValue,
        uint256 pendingCount,
        bytes32 vmProtoHashBefore,
        uint256 gracePeriodTicks
    )
        private
        pure
        returns(bytes32)
    {
        bytes32 challengeHash = ChallengeUtils.pendingTopHash(
            data.afterPendingTop,
            pendingValue,
            pendingCount - (data.beforePendingCount + data.importedMessageCount)
        );
        return RollupUtils.childNodeHash(
            prevLeaf,
            deadlineTicks,
            RollupUtils.challengeDataHash(
                challengeHash,
                gracePeriodTicks + RollupTime.blocksToTicks(1)
            ),
            INVALID_PENDING_TOP_TYPE,
            vmProtoHashBefore
        );
    }

    function generateInvalidMessagesLeaf(
        MakeAssertionData memory data,
        bytes32 prevLeaf,
        uint256 deadlineTicks,
        bytes32 vmProtoHashBefore,
        uint256 gracePeriodTicks
    )
        private
        pure
        returns(bytes32)
    {
        bytes32 challengeHash = ChallengeUtils.messagesHash(
            data.beforePendingTop,
            data.afterPendingTop,
            Value.hashEmptyTuple(),
            data.importedMessagesSlice,
            data.importedMessageCount
        );
        bytes32 nodeDataHash = RollupUtils.challengeDataHash(
            challengeHash,
            gracePeriodTicks + RollupTime.blocksToTicks(1)
        );
        return RollupUtils.childNodeHash(
            prevLeaf,
            deadlineTicks,
            nodeDataHash,
            INVALID_MESSAGES_TYPE,
            vmProtoHashBefore
        );
    }

    function generateInvalidExecutionLeaf(
        MakeAssertionData memory data,
        bytes32 prevLeaf,
        uint256 deadlineTicks,
        bytes32 vmProtoHashBefore,
        uint256 gracePeriodTicks,
        uint256 checkTimeTicks
    )
        private
        pure
        returns(bytes32)
    {
        bytes32 preconditionHash = Protocol.generatePreconditionHash(
             data.beforeVMHash,
             data.timeBoundsBlocks,
             data.importedMessagesSlice
        );
        bytes32 assertionHash = Protocol.generateAssertionHash(
            data.afterVMHash,
            data.didInboxInsn,
            data.numArbGas,
            0x00,
            data.messagesAccHash,
            0x00,
            data.logsAccHash
        );
        bytes32 executionHash = ChallengeUtils.executionHash(
            data.numSteps,
            preconditionHash,
            assertionHash
        );
        return RollupUtils.childNodeHash(
            prevLeaf,
            deadlineTicks,
            RollupUtils.challengeDataHash(
                executionHash,
                gracePeriodTicks + checkTimeTicks
            ),
            INVALID_EXECUTION_TYPE,
            vmProtoHashBefore
        );
    }

    function generateValidLeaf(
        MakeAssertionData memory data,
        bytes32 prevLeaf,
        uint256 deadlineTicks
    )
        private
        pure
        returns(bytes32)
    {
        return RollupUtils.childNodeHash(
            prevLeaf,
            deadlineTicks,
            RollupUtils.validDataHash(
                data.messagesAccHash,
                data.logsAccHash
            ),
            VALID_CHILD_TYPE,
            RollupUtils.protoStateHash(
                data.afterVMHash,
                data.afterPendingTop,
                data.beforePendingCount + data.importedMessageCount
            )
        );
    }
}
