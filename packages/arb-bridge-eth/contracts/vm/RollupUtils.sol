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

import "../arch/Protocol.sol";
import "../libraries/RollupTime.sol";


library RollupUtils {

    uint256 constant VALID_CHILD_TYPE = 3;

    event ConfirmedValidAssertion(
        bytes32 indexed nodeHash
    );

    struct ConfirmData {
        bytes32 initalProtoStateHash;
        uint256[] branches;
        uint256[] deadlineTicks;
        bytes32[] challengeNodeData;
        bytes32[] logsAcc;
        bytes32[] vmProtoStateHashes;
        uint256[] messageCounts;
        bytes messages;
    }

    function confirm(
        RollupUtils.ConfirmData memory data,
        bytes32 confNode
    )
        internal
        pure
        returns(bytes32[] memory validNodeHashes, bytes32 lastNode)
    {
        uint256 nodeCount = data.branches.length;
        _verifyDataLength(data);
        uint256 validNum = 0;
        uint256 invalidNum = 0;
        uint256 messagesOffset = 0;

        validNodeHashes = new bytes32[](nodeCount);

        bytes32 vmProtoStateHash = data.initalProtoStateHash;

        for (uint256 i = 0; i < nodeCount; i++) {
            uint256 branchType = data.branches[i];
            bytes32 nodeDataHash;
            if (branchType == VALID_CHILD_TYPE) {
                (messagesOffset, nodeDataHash, vmProtoStateHash) = processValidNode(data, validNum, messagesOffset);
                validNum++;
            } else {
                nodeDataHash = data.challengeNodeData[invalidNum];
                invalidNum++;
            }

            confNode = childNodeHash(
                confNode,
                data.deadlineTicks[i],
                nodeDataHash,
                branchType,
                vmProtoStateHash
            );

            if (branchType == VALID_CHILD_TYPE) {
                validNodeHashes[validNum - 1] = confNode;
            }
        }
        return (validNodeHashes, confNode);
    }

    function generateLastMessageHash(bytes memory messages, uint256 startOffset, uint256 count) internal pure returns (bytes32, uint256) {
        bool valid;
        bytes32 hashVal = 0x00;
        Value.Data memory messageVal;
        uint256 offset = startOffset;
        for (uint256 i = 0; i < count; i++) {
            (valid, offset, messageVal) = Value.deserialize(messages, offset);
            require(valid, "Invalid output message");
            hashVal = keccak256(abi.encodePacked(hashVal, Value.hash(messageVal)));
        }
        return (hashVal, offset);
    }

    function processValidNode(
        RollupUtils.ConfirmData memory data,
        uint256 validNum,
        uint256 startOffset
    )
        internal
        pure
        returns(uint256, bytes32, bytes32)
    {
        (bytes32 lastMsgHash, uint256 messagesOffset) = generateLastMessageHash(
            data.messages,
            startOffset,
            data.messageCounts[validNum]
        );
        bytes32 nodeDataHash = validDataHash(
            lastMsgHash,
            data.logsAcc[validNum]
        );
        bytes32 vmProtoStateHash = data.vmProtoStateHashes[validNum];
        return (messagesOffset, nodeDataHash, vmProtoStateHash);
    }

    function _verifyDataLength(RollupUtils.ConfirmData memory data) private pure{
        uint256 nodeCount = data.branches.length;
        uint256 validNodeCount = data.messageCounts.length;
        require(data.vmProtoStateHashes.length == validNodeCount);
        require(data.logsAcc.length == validNodeCount);
        require(data.deadlineTicks.length == nodeCount);
        require(data.challengeNodeData.length == nodeCount - validNodeCount);
    }

    function protoStateHash(
        bytes32 machineHash,
        bytes32 inboxTop,
        uint256 inboxCount
    )
        internal
        pure
        returns(bytes32)
    {
        return keccak256(
            abi.encodePacked(
                machineHash,
                inboxTop,
                inboxCount
            )
        );
    }

    function validDataHash(
        bytes32 messagesAcc,
        bytes32 logsAcc
    )
        internal
        pure
        returns(bytes32)
    {
        return keccak256(
            abi.encodePacked(
                messagesAcc,
                logsAcc
            )
        );
    }

    function challengeDataHash(
        bytes32 challenge,
        uint256 challengePeriod
    )
        internal
        pure
        returns(bytes32)
    {
        return keccak256(
            abi.encodePacked(
                challenge,
                challengePeriod
            )
        );
    }

    function childNodeHash(
        bytes32 prevNodeHash,
        uint256 deadlineTicks,
        bytes32 nodeDataHash,
        uint256 childType,
        bytes32 vmProtoStateHash
    )
        internal
        pure
        returns(bytes32)
    {
        return keccak256(
            abi.encodePacked(
                prevNodeHash,
                keccak256(
                    abi.encodePacked(
                        vmProtoStateHash,
                        deadlineTicks,
                        nodeDataHash,
                        childType
                    )
                )
            )
        );
    }

    function calculatePath(
        bytes32 from,
        bytes32[] memory proof
    )
        internal
        pure
        returns(bytes32)
    {
        return calculatePathOffset(
            from,
            proof,
            0,
            proof.length
        );
    }

    function calculatePathOffset(
        bytes32 from,
        bytes32[] memory proof,
        uint256 start,
        uint256 end
    )
        internal
        pure
        returns(bytes32)
    {
        bytes32 node = from;
        for (uint256 i = start; i<end; i++) {
            node = keccak256(abi.encodePacked(node, proof[i]));
        }
        return node;
    }
}
