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


library RollupUtils {
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
