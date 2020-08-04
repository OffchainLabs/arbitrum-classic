// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2020, Offchain Labs, Inc.
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
pragma solidity ^0.5.11;

import "../rollup/RollupUtils.sol";
import "../rollup/NodeGraphUtils.sol";

contract RollupTester {
    function confirm(
        bytes32 confNode,
        bytes32 initalProtoStateHash,
        uint256[] memory branches,
        uint256[] memory deadlineTicks,
        bytes32[] memory challengeNodeData,
        bytes32[] memory logsAcc,
        bytes32[] memory vmProtoStateHashes,
        uint256[] memory messageCounts,
        bytes memory messages
    ) public pure returns (bytes32[] memory validNodeHashes, bytes32 lastNode) {
        return
            RollupUtils.confirm(
                RollupUtils.ConfirmData(
                    initalProtoStateHash,
                    branches,
                    deadlineTicks,
                    challengeNodeData,
                    logsAcc,
                    vmProtoStateHashes,
                    messageCounts,
                    messages
                ),
                confNode
            );
    }

    function generateLastMessageHash(
        bytes memory messages,
        uint256 startOffset,
        uint256 length
    ) public pure returns (bytes32, uint256) {
        return
            RollupUtils.generateLastMessageHash(messages, startOffset, length);
    }

    function processValidNode(
        bytes32 initalProtoStateHash,
        uint256[] memory branches,
        uint256[] memory deadlineTicks,
        bytes32[] memory challengeNodeData,
        bytes32[] memory logsAcc,
        bytes32[] memory vmProtoStateHashes,
        uint256[] memory messageCounts,
        bytes memory messages,
        uint256 validNum,
        uint256 startOffset
    )
        public
        pure
        returns (
            uint256,
            bytes32,
            bytes32
        )
    {
        return
            RollupUtils.processValidNode(
                RollupUtils.ConfirmData(
                    initalProtoStateHash,
                    branches,
                    deadlineTicks,
                    challengeNodeData,
                    logsAcc,
                    vmProtoStateHashes,
                    messageCounts,
                    messages
                ),
                validNum,
                startOffset
            );
    }

    function calculateLeafFromPath(bytes32 from, bytes32[] memory proof)
        public
        pure
        returns (bytes32)
    {
        return RollupUtils.calculateLeafFromPath(from, proof);
    }

    function childNodeHash(
        bytes32 prevNodeHash,
        uint256 deadlineTicks,
        bytes32 nodeDataHash,
        uint256 childType,
        bytes32 vmProtoStateHash
    ) public pure returns (bytes32) {
        return
            RollupUtils.childNodeHash(
                prevNodeHash,
                deadlineTicks,
                nodeDataHash,
                childType,
                vmProtoStateHash
            );
    }

    function computeProtoHashBefore(
        bytes32 machineHash,
        bytes32 inboxTop,
        uint256 inboxCount,
        uint256 messageCount,
        uint256 logCount
    ) public pure returns (bytes32) {
        return
            RollupUtils.protoStateHash(
                machineHash,
                inboxTop,
                inboxCount,
                messageCount,
                logCount
            );
    }

    function computePrevLeaf(
        bytes32[9] memory fields,
        uint256[5] memory fields2,
        uint32 prevChildType,
        uint64 numSteps,
        uint64 numArbGas,
        uint64 messageCount,
        uint64 logCount
    ) public pure returns (bytes32, bytes32) {
        NodeGraphUtils.AssertionData memory assertData = NodeGraphUtils
            .makeAssertion(
            fields,
            fields2,
            prevChildType,
            numSteps,
            numArbGas,
            messageCount,
            logCount
        );

        return NodeGraphUtils.computePrevLeaf(assertData);
    }

    function generateInvalidInboxTopLeaf(
        uint256[4] memory invalidInboxData,
        bytes32[9] memory fields,
        uint256[5] memory fields2,
        uint32 prevChildType,
        uint64 numSteps,
        uint64 numArbGas,
        uint64 messageCount,
        uint64 logCount
    ) public pure returns (bytes32) {
        NodeGraphUtils.AssertionData memory assertData = NodeGraphUtils
            .makeAssertion(
            fields,
            fields2,
            prevChildType,
            numSteps,
            numArbGas,
            messageCount,
            logCount
        );

        return _generateInvalidInboxTopLeaf(assertData, invalidInboxData);
    }

    function _generateInvalidInboxTopLeaf(
        NodeGraphUtils.AssertionData memory assertData,
        uint256[4] memory invalidInboxData
    ) private pure returns (bytes32) {
        (bytes32 prevLeaf, bytes32 vmProtoHashBefore) = NodeGraphUtils
            .computePrevLeaf(assertData);

        return
            NodeGraphUtils.generateInvalidInboxTopLeaf(
                assertData,
                prevLeaf,
                invalidInboxData[3],
                bytes32(invalidInboxData[0]),
                invalidInboxData[1],
                vmProtoHashBefore,
                invalidInboxData[2]
            );
    }

    function generateInvalidMessagesLeaf(
        uint256 gracePeriodTicks,
        uint256 deadlineTicks,
        bytes32[9] memory fields,
        uint256[5] memory fields2,
        uint32 prevChildType,
        uint64 numSteps,
        uint64 numArbGas,
        uint64 messageCount,
        uint64 logCount
    ) public pure returns (bytes32) {
        NodeGraphUtils.AssertionData memory assertData = NodeGraphUtils
            .makeAssertion(
            fields,
            fields2,
            prevChildType,
            numSteps,
            numArbGas,
            messageCount,
            logCount
        );

        return
            _generateInvalidMessagesLeaf(
                assertData,
                gracePeriodTicks,
                deadlineTicks
            );
    }

    function _generateInvalidMessagesLeaf(
        NodeGraphUtils.AssertionData memory assertData,
        uint256 gracePeriodTicks,
        uint256 deadlineTicks
    ) private pure returns (bytes32) {
        (bytes32 prevLeaf, bytes32 vmProtoHashBefore) = NodeGraphUtils
            .computePrevLeaf(assertData);

        return
            NodeGraphUtils.generateInvalidMessagesLeaf(
                assertData,
                prevLeaf,
                deadlineTicks,
                vmProtoHashBefore,
                gracePeriodTicks
            );
    }

    function generateInvalidExecutionLeaf(
        uint256 gracePeriodTicks,
        uint256 checkTimeTicks,
        uint256 deadlineTicks,
        bytes32[9] memory fields,
        uint256[5] memory fields2,
        uint32 prevChildType,
        uint64 numSteps,
        uint64 numArbGas,
        uint64 messageCount,
        uint64 logCount
    ) public pure returns (bytes32) {
        NodeGraphUtils.AssertionData memory assertData = NodeGraphUtils
            .makeAssertion(
            fields,
            fields2,
            prevChildType,
            numSteps,
            numArbGas,
            messageCount,
            logCount
        );

        return
            _generateInvalidExecutionLeaf(
                assertData,
                gracePeriodTicks,
                checkTimeTicks,
                deadlineTicks
            );
    }

    function _generateInvalidExecutionLeaf(
        NodeGraphUtils.AssertionData memory assertData,
        uint256 gracePeriodTicks,
        uint256 checkTimeTicks,
        uint256 deadlineTicks
    ) private pure returns (bytes32) {
        (bytes32 prevLeaf, bytes32 vmProtoHashBefore) = NodeGraphUtils
            .computePrevLeaf(assertData);

        return
            NodeGraphUtils.generateInvalidExecutionLeaf(
                assertData,
                prevLeaf,
                deadlineTicks,
                vmProtoHashBefore,
                gracePeriodTicks,
                checkTimeTicks
            );
    }

    function generateValidLeaf(
        uint256 deadlineTicks,
        bytes32[9] memory fields,
        uint256[5] memory fields2,
        uint32 prevChildType,
        uint64 numSteps,
        uint64 numArbGas,
        uint64 messageCount,
        uint64 logCount
    ) public pure returns (bytes32) {
        NodeGraphUtils.AssertionData memory assertData = NodeGraphUtils
            .makeAssertion(
            fields,
            fields2,
            prevChildType,
            numSteps,
            numArbGas,
            messageCount,
            logCount
        );

        return _generateValidLeaf(assertData, deadlineTicks);
    }

    function _generateValidLeaf(
        NodeGraphUtils.AssertionData memory assertData,
        uint256 deadlineTicks
    ) private pure returns (bytes32) {
        (bytes32 prevLeaf, ) = NodeGraphUtils.computePrevLeaf(assertData);

        return
            NodeGraphUtils.generateValidLeaf(
                assertData,
                prevLeaf,
                deadlineTicks
            );
    }
}
