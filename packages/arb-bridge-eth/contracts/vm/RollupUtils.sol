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
    string constant ARG_LENGTHS = "ARG_LENGTHS";
    string constant NODE_MATCH = "NODE_MATCH";

    function protoStateHash(
        bytes32 machineHash,
        bytes32 pendingTop,
        uint256 pendingCount
    )
        internal
        pure
        returns(bytes32)
    {
        return keccak256(
            abi.encodePacked(
                machineHash,
                pendingTop,
                pendingCount
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

    function verifyMultipathProof(
        bytes32 from,
        bytes32[] memory to,
        uint64[] memory startingPoints,
        bytes32[] memory proofs,
        uint64[] memory proofLengths,
        uint64[] memory permutation
    ) public pure returns(bool) {
        require(startingPoints.length == proofLengths.length, ARG_LENGTHS);
        require(to.length == permutation.length, ARG_LENGTHS);

        bytes32[] memory proven = new bytes32[](1+startingPoints.length);

        proven[0] = from;
        uint64 proofOffset = 0;
        for (uint64 i=0; i<startingPoints.length; i++) {
            proven[i+1] = calculatePathOffset(
                proven[startingPoints[i]],
                proofs, 
                proofOffset, 
                proofOffset+proofLengths[i]
            );
            proofOffset += proofLengths[i];
        }

        for (uint64 i=0; i<startingPoints.length; i++) {
            require(proven[permutation[i]] == to[i], NODE_MATCH);
        }
    }
}
