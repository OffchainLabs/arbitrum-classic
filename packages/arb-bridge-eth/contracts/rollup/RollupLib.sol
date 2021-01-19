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

library RollupLib {
    function nodeStateHash(
        uint256 proposedBlock,
        uint256 totalGasUsed,
        bytes32 machineHash,
        bytes32 inboxTop,
        uint256 inboxCount,
        uint256 totalMessageCount,
        uint256 totalLogCount,
        uint256 inboxMaxCount
    ) internal pure returns (bytes32) {
        return
            keccak256(
                abi.encodePacked(
                    proposedBlock,
                    totalGasUsed,
                    machineHash,
                    inboxTop,
                    inboxCount,
                    totalMessageCount,
                    totalLogCount,
                    inboxMaxCount
                )
            );
    }

    struct Assertion {
        uint256 beforeProposedBlock;
        uint256 beforeTotalGasUsed;
        bytes32 beforeMachineHash;
        bytes32 beforeInboxHash;
        uint256 beforeInboxCount;
        uint256 beforeTotalSendCount;
        uint256 beforeTotalLogCount;
        uint256 beforeInboxMaxCount;
        bytes32 inboxDelta;
        uint256 inboxMessagesRead;
        uint256 gasUsed;
        bytes32 sendAcc;
        uint256 sendCount;
        bytes32 logAcc;
        uint256 logCount;
        bytes32 afterInboxHash;
        bytes32 afterMachineHash;
    }

    function decodeAssertion(bytes32[7] memory bytes32Fields, uint256[10] memory intFields)
        internal
        pure
        returns (Assertion memory)
    {
        return
            Assertion(
                intFields[0],
                intFields[1],
                bytes32Fields[0],
                bytes32Fields[1],
                intFields[2],
                intFields[3],
                intFields[4],
                intFields[5],
                bytes32Fields[2],
                intFields[6],
                intFields[7],
                bytes32Fields[3],
                intFields[8],
                bytes32Fields[4],
                intFields[9],
                bytes32Fields[5],
                bytes32Fields[6]
            );
    }

    function beforeNodeStateHash(Assertion memory assertion) internal pure returns (bytes32) {
        return
            nodeStateHash(
                assertion.beforeProposedBlock,
                assertion.beforeTotalGasUsed,
                assertion.beforeMachineHash,
                assertion.beforeInboxHash,
                assertion.beforeInboxCount,
                assertion.beforeTotalSendCount,
                assertion.beforeTotalLogCount,
                assertion.beforeInboxMaxCount
            );
    }

    function nodeStateHash(Assertion memory assertion, uint256 inboxMaxCount)
        internal
        view
        returns (bytes32)
    {
        return
            nodeStateHash(
                block.number,
                assertion.beforeTotalGasUsed + assertion.gasUsed,
                assertion.afterMachineHash,
                assertion.afterInboxHash,
                assertion.beforeInboxCount + assertion.inboxMessagesRead,
                assertion.beforeTotalSendCount + assertion.sendCount,
                assertion.beforeTotalLogCount + assertion.logCount,
                inboxMaxCount
            );
    }

    function executionHash(Assertion memory assertion) private pure returns (bytes32) {
        return
            ChallengeLib.bisectionChunkHash(
                0,
                assertion.gasUsed,
                ChallengeLib.assertionHash(
                    0,
                    ChallengeLib.assertionRestHash(
                        assertion.inboxDelta,
                        assertion.beforeMachineHash,
                        0,
                        0,
                        0,
                        0
                    )
                ),
                ChallengeLib.assertionHash(
                    assertion.gasUsed,
                    ChallengeLib.assertionRestHash(
                        0,
                        assertion.afterMachineHash,
                        assertion.sendAcc,
                        assertion.sendCount,
                        assertion.logAcc,
                        assertion.logCount
                    )
                )
            );
    }

    function challengeRoot(
        Assertion memory assertion,
        uint256 inboxTopCount,
        bytes32 inboxTopHash,
        uint256 blockProposed
    ) internal pure returns (bytes32) {
        bytes32 inboxConsistencyHash =
            ChallengeLib.bisectionChunkHash(
                0,
                inboxTopCount - assertion.beforeInboxCount - assertion.inboxMessagesRead,
                inboxTopHash,
                assertion.afterInboxHash
            );

        bytes32 inboxDeltaHash =
            ChallengeLib.bisectionChunkHash(
                0,
                assertion.inboxMessagesRead,
                ChallengeLib.inboxDeltaHash(assertion.afterInboxHash, 0),
                ChallengeLib.inboxDeltaHash(assertion.beforeInboxHash, assertion.inboxDelta)
            );

        return
            challengeRootHash(
                inboxConsistencyHash,
                inboxDeltaHash,
                executionHash(assertion),
                blockProposed
            );
    }

    function challengeRootHash(
        bytes32 inboxConsistency,
        bytes32 inboxDelta,
        bytes32 execution,
        uint256 proposedTime
    ) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(inboxConsistency, inboxDelta, execution, proposedTime));
    }

    function confirmHash(Assertion memory assertion) internal pure returns (bytes32) {
        return confirmHash(assertion.sendAcc, assertion.logAcc);
    }

    function confirmHash(bytes32 sendAcc, bytes32 logAcc) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(sendAcc, logAcc));
    }

    function generateLastMessageHash(bytes memory messageData, uint256[] memory messageLengths)
        internal
        pure
        returns (bytes32)
    {
        uint256 offset = 0;
        uint256 messageCount = messageLengths.length;
        uint256 dataLength = messageData.length;
        bytes32 messageAcc = 0;
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
}
