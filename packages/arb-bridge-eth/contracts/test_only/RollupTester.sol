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

pragma solidity ^0.5.3;

import "../vm/RollupUtils.sol";

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
    )
        public
        pure
        returns(bytes32[] memory validNodeHashes, bytes32 lastNode)
    {
        return RollupUtils.confirm(
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

    function generateLastMessageHash(bytes memory messages, uint256 startOffset, uint256 length) public pure returns (bytes32, uint) {
        return RollupUtils.generateLastMessageHash(messages, startOffset, length);
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
        returns(uint256, bytes32, bytes32)
    {
        return RollupUtils.processValidNode(
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
}
