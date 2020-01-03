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


library ChallengeUtils {

    function pendingTopHash(
        bytes32 _lowerHash,
        bytes32 _topHash,
        uint256 _chainLength
    )
        internal
        pure
        returns(bytes32)
    {
        return keccak256(
            abi.encodePacked(
                _topHash,
                _lowerHash,
                _chainLength
            )
        );
    }

    function messagesHash(
        bytes32 _lowerHashA,
        bytes32 _topHashA,
        bytes32 _lowerHashB,
        bytes32 _topHashB,
        uint256 _chainLength
    )
        internal
        pure
        returns(bytes32)
    {
        return keccak256(
            abi.encodePacked(
                _lowerHashA,
                _topHashA,
                _lowerHashB,
                _topHashB,
                _chainLength
            )
        );
    }

    function executionHash(
        bytes32 _preconditionHash,
        bytes32 _assertionHash
    )
        internal
        pure
        returns(bytes32)
    {
        return keccak256(
            abi.encodePacked(
                _preconditionHash,
                _assertionHash
            )
        );
    }
}
