// SPDX-License-Identifier: Apache-2.0

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

pragma solidity ^0.5.11;

library ChallengeUtils {
    uint256 public constant INVALID_INBOX_TOP_TYPE = 0;
    uint256 public constant INVALID_MESSAGES_TYPE = 1;
    uint256 public constant INVALID_EXECUTION_TYPE = 2;
    uint256 public constant VALID_CHILD_TYPE = 3;

    function getInvalidInboxType() internal pure returns (uint256) {
        return INVALID_INBOX_TOP_TYPE;
    }

    function getInvalidMsgsType() internal pure returns (uint256) {
        return INVALID_MESSAGES_TYPE;
    }

    function getInvalidExType() internal pure returns (uint256) {
        return INVALID_EXECUTION_TYPE;
    }

    function getValidChildType() internal pure returns (uint256) {
        return VALID_CHILD_TYPE;
    }

    function inboxTopHash(
        bytes32 _lowerHash,
        bytes32 _topHash,
        uint256 _chainLength
    ) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(_lowerHash, _topHash, _chainLength));
    }

    function messagesHash(
        bytes32 _lowerHashA,
        bytes32 _topHashA,
        bytes32 _lowerHashB,
        bytes32 _topHashB,
        uint256 _chainLength
    ) internal pure returns (bytes32) {
        return
            keccak256(
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
        uint64 _numSteps,
        bytes32 _beforeHash,
        bytes32 _beforeInboxHash,
        bytes32 _assertionHash
    ) internal pure returns (bytes32) {
        return
            keccak256(
                abi.encodePacked(
                    _numSteps,
                    _beforeHash,
                    _beforeInboxHash,
                    _assertionHash
                )
            );
    }
}
