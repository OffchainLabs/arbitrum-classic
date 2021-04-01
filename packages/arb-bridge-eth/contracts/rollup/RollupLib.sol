// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2021, Offchain Labs, Inc.
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

pragma solidity ^0.6.11;

import "../challenge/ChallengeLib.sol";
import "../bridge/interfaces/ISequencerInbox.sol";
import "../arch/Marshaling.sol";
import "./INode.sol";

library RollupLib {
    struct Config {
        bytes32 machineHash;
        uint256 confirmPeriodBlocks;
        uint256 extraChallengeTimeBlocks;
        uint256 arbGasSpeedLimitPerBlock;
        uint256 baseStake;
        address stakeToken;
        address owner;
        address sequencer;
        uint256 sequencerDelayBlocks;
        uint256 sequencerDelaySeconds;
        bytes extraConfig;
    }

    struct ExecutionState {
        uint256 gasUsed;
        bytes32 machineHash;
        uint256 inboxCount;
        uint256 sendCount;
        uint256 logCount;
        bytes32 sendAcc;
        bytes32 logAcc;
        uint256 proposedBlock;
        uint256 inboxMaxCount;
    }

    function stateHash(ExecutionState memory execState) internal pure returns (bytes32) {
        return
            keccak256(
                abi.encodePacked(
                    execState.gasUsed,
                    execState.machineHash,
                    execState.inboxCount,
                    execState.sendCount,
                    execState.logCount,
                    execState.sendAcc,
                    execState.logAcc,
                    execState.proposedBlock,
                    execState.inboxMaxCount
                )
            );
    }

    struct Assertion {
        ExecutionState beforeState;
        ExecutionState afterState;
    }

    function decodeExecutionState(
        bytes32[3] memory bytes32Fields,
        uint256[4] memory intFields,
        uint256 proposedBlock,
        uint256 inboxMaxCount
    ) internal pure returns (ExecutionState memory) {
        return
            ExecutionState(
                intFields[0],
                bytes32Fields[0],
                intFields[1],
                intFields[2],
                intFields[3],
                bytes32Fields[1],
                bytes32Fields[2],
                proposedBlock,
                inboxMaxCount
            );
    }

    function decodeAssertion(
        bytes32[3][2] memory bytes32Fields,
        uint256[4][2] memory intFields,
        uint256 beforeProposedBlock,
        uint256 beforeInboxMaxCount,
        uint256 inboxMaxCount
    ) internal view returns (Assertion memory) {
        return
            Assertion(
                decodeExecutionState(
                    bytes32Fields[0],
                    intFields[0],
                    beforeProposedBlock,
                    beforeInboxMaxCount
                ),
                decodeExecutionState(bytes32Fields[1], intFields[1], block.number, inboxMaxCount)
            );
    }

    function executionStateChallengeHash(ExecutionState memory state)
        internal
        pure
        returns (bytes32)
    {
        return
            ChallengeLib.assertionHash(
                state.gasUsed,
                ChallengeLib.assertionRestHash(
                    state.inboxCount,
                    state.machineHash,
                    state.sendAcc,
                    state.sendCount,
                    state.logAcc,
                    state.logCount
                )
            );
    }

    function executionHash(Assertion memory assertion) internal pure returns (bytes32) {
        return
            ChallengeLib.bisectionChunkHash(
                assertion.beforeState.gasUsed,
                assertion.afterState.gasUsed - assertion.beforeState.gasUsed,
                RollupLib.executionStateChallengeHash(assertion.beforeState),
                RollupLib.executionStateChallengeHash(assertion.afterState)
            );
    }

    function challengeRoot(
        Assertion memory assertion,
        bytes32 assertionExecHash,
        uint256 blockProposed
    ) internal pure returns (bytes32) {
        return challengeRootHash(assertionExecHash, blockProposed, assertion.afterState.inboxCount);
    }

    function challengeRootHash(
        bytes32 execution,
        uint256 proposedTime,
        uint256 maxMessageCount
    ) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(execution, proposedTime, maxMessageCount));
    }

    function confirmHash(Assertion memory assertion) internal pure returns (bytes32) {
        return
            confirmHash(
                assertion.beforeState.sendAcc,
                assertion.afterState.sendAcc,
                assertion.afterState.logAcc,
                assertion.afterState.sendCount,
                assertion.afterState.logCount
            );
    }

    function confirmHash(
        bytes32 beforeSendAcc,
        bytes32 afterSendAcc,
        bytes32 afterLogAcc,
        uint256 afterSendCount,
        uint256 afterLogCount
    ) internal pure returns (bytes32) {
        return
            keccak256(
                abi.encodePacked(
                    beforeSendAcc,
                    afterSendAcc,
                    afterSendCount,
                    afterLogAcc,
                    afterLogCount
                )
            );
    }

    function feedAccumulator(
        bytes memory messageData,
        uint256[] memory messageLengths,
        bytes32 beforeAcc
    ) internal pure returns (bytes32) {
        uint256 offset = 0;
        uint256 messageCount = messageLengths.length;
        uint256 dataLength = messageData.length;
        bytes32 messageAcc = beforeAcc;
        for (uint256 i = 0; i < messageCount; i++) {
            uint256 messageLength = messageLengths[i];
            require(offset + messageLength <= dataLength, "DATA_OVERRUN");
            bytes32 messageHash;
            assembly {
                messageHash := keccak256(add(messageData, add(offset, 32)), messageLength)
            }
            messageAcc = keccak256(abi.encodePacked(messageAcc, messageHash));
            offset += messageLength;
        }
        require(offset == dataLength, "DATA_LENGTH");
        return messageAcc;
    }

    function nodeHash(
        bool hasSibling,
        bytes32 lastHash,
        bytes32 assertionExecHash,
        bytes32 inboxAcc
    ) internal pure returns (bytes32) {
        uint8 hasSiblingInt = hasSibling ? 1 : 0;
        return keccak256(abi.encodePacked(hasSiblingInt, lastHash, assertionExecHash, inboxAcc));
    }

    function nodeAccumulator(bytes32 prevAcc, bytes32 newNodeHash) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(prevAcc, newNodeHash));
    }

    function proveSeqBatchMsgCount(
        bytes calldata proof,
        uint256 offset,
        bytes32 acc
    ) internal pure returns (uint256, uint256) {
        uint256 messageCount;

        bytes32 buildingAcc;
        (offset, buildingAcc) = Marshaling.deserializeBytes32(proof, offset);
        uint8 isDelayed = uint8(proof[offset]);
        offset++;
        require(isDelayed == 0 || isDelayed == 1, "IS_DELAYED_NUM");
        if (isDelayed == 0) {
            uint256 seqNum;
            bytes32 messageHash;
            (offset, seqNum) = Marshaling.deserializeInt(proof, offset);
            (offset, messageHash) = Marshaling.deserializeBytes32(proof, offset);
            buildingAcc = keccak256(
                abi.encodePacked("Sequencer message:", buildingAcc, seqNum, messageHash)
            );
            messageCount = seqNum + 1;
        } else {
            uint256 firstSequencerSeqNum;
            uint256 delayedStart;
            uint256 delayedEnd;
            bytes32 delayedEndAcc;
            (offset, firstSequencerSeqNum) = Marshaling.deserializeInt(proof, offset);
            (offset, delayedStart) = Marshaling.deserializeInt(proof, offset);
            (offset, delayedEnd) = Marshaling.deserializeInt(proof, offset);
            (offset, delayedEndAcc) = Marshaling.deserializeBytes32(proof, offset);
            buildingAcc = keccak256(
                abi.encodePacked(
                    "Delayed messages:",
                    buildingAcc,
                    firstSequencerSeqNum,
                    delayedStart,
                    delayedEnd,
                    delayedEndAcc
                )
            );
            messageCount = delayedEnd - delayedStart + firstSequencerSeqNum;
        }
        require(buildingAcc == acc, "BATCH_ACC");

        return (offset, messageCount);
    }

    function proveSequencerBatchContains(
        ISequencerInbox sequencerBridge,
        bytes calldata proof,
        uint256 inboxCount
    ) internal view returns (bytes32) {
        if (inboxCount == 0) {
            return 0;
        }

        (uint256 offset, uint256 seqBatchNum) = Marshaling.deserializeInt(proof, 0);
        uint256 lastBatchCount = 0;
        if (seqBatchNum > 0) {
            (offset, lastBatchCount) = proveSeqBatchMsgCount(
                proof,
                offset,
                sequencerBridge.inboxAccs(seqBatchNum - 1)
            );
            lastBatchCount++;
        }

        bytes32 seqBatchAcc = sequencerBridge.inboxAccs(seqBatchNum);
        uint256 thisBatchCount;
        (offset, thisBatchCount) = proveSeqBatchMsgCount(proof, offset, seqBatchAcc);

        require(inboxCount > lastBatchCount, "BATCH_START");
        require(inboxCount <= thisBatchCount, "BATCH_END");

        return seqBatchAcc;
    }
}
