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
        bytes32 inboxHash,
        bytes32 pendingTop
    )
        internal
        pure
        returns(bytes32)
    {
        return keccak256(abi.encodePacked(
            machineHash,
            inboxHash,
            pendingTop
        ));
    }

    function disputableNodeHash(
        uint deadline,
        bytes32 preconditionHash,
        bytes32 pendingAssertion,
        bytes32 importedAssertion,
        bytes32 assertionHash
    )
        internal
        pure
        returns(bytes32)
    {
        return keccak256(
            abi.encodePacked(
                deadline,
                preconditionHash,
                pendingAssertion,
                importedAssertion,
                assertionHash
            )
        );
    }

    function childNodeHash(
        bytes32 prevNodeHash,
        bytes32 disputableNodeHashVal,
        uint    childType,
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
                        disputableNodeHashVal,
                        childType,
                        vmProtoStateHash
                    )
                )
            )
        );
    }

    function importedAssertionHash(bytes32 beforePendingTop, uint32 messageCount, bytes32 messagesSlice) internal pure returns(bytes32) {
        return keccak256(
            abi.encodePacked(
                beforePendingTop,
                messageCount,
                messagesSlice
            )
        );
    }

    function pendingAssertionHash(bytes32 afterPendingTop, bytes32 currentPending) internal pure returns(bytes32) {
        return keccak256(
            abi.encodePacked(
                afterPendingTop,
                currentPending
            )
        );
    }

    function isPath(
        bytes32 from,
        bytes32 to,
        bytes32[] memory proof
    )
        internal
        pure
        returns(bool)
    {
        return isPathOffset(
            from,
            to,
            proof,
            0,
            proof.length
        );
    }

    function isPathOffset(
        bytes32 from,
        bytes32 to,
        bytes32[] memory proof,
        uint start,
        uint end
    )
        internal
        pure
        returns(bool)
    {
        bytes32 node = from;
        for (uint i = start; i<end; i++) {
            node = keccak256(abi.encodePacked(node, proof[i]));
        }
        return (node==to);
    }

    function isInPath(
        bytes32 from,
        bytes32 middle,
        bytes32 to,
        bytes32[] memory proof1,
        bytes32[] memory proof2
    )
        internal
        pure
        returns(bool)
    {
        return isPath(from, middle, proof1) &&
               isPath(middle, to, proof2);
    }

    function isConflict(
        bytes32 from,
        bytes32 to1,
        bytes32 to2,
        bytes32[] memory proof1,
        bytes32[] memory proof2
    )
        internal
        pure
        returns(bool)
    {
        return (proof1[0] != proof2[0]) &&
            isPath(from, to1, proof1) &&
            isPath(from, to2, proof2);
    }
}
