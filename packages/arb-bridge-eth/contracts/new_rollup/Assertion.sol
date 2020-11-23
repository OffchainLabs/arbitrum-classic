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

import "../challenge/ChallengeUtils.sol";

library Assertion {
    function hashAssertion(
        bytes32 machineHash,
        bytes32 inboxAcc,
        bytes32 messageAcc,
        bytes32 logAcc,
        uint256 gasUsed,
        uint256 inboxCount,
        uint256 messageCount,
        uint256 logCount
    ) internal pure returns (bytes32) {
        bytes32 innerHash = keccak256(
            abi.encodePacked(machineHash, messageAcc, logAcc, gasUsed, messageCount, logCount)
        );
        return hashAssertion(innerHash, inboxCount, inboxAcc);
    }

    function hashAssertion(
        bytes32 innerHash,
        uint256 inboxCount,
        bytes32 inboxAcc
    ) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(innerHash, inboxCount, inboxAcc));
    }

    function hashInboxNode(
        bytes32 beforeInboxHash,
        bytes32 afterInboxHash,
        uint256 inboxCount
    ) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(beforeInboxHash, afterInboxHash, inboxCount));
    }

    function hashExecutionNode(
        bytes32 beforeAssertionHash,
        bytes32 afterAssertionHash,
        uint256 stepCount
    ) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(beforeAssertionHash, afterAssertionHash, stepCount));
    }
}
