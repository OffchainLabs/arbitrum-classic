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

import "../arch/OneStepProof.sol";

contract OneStepProofTester {
    function validateProof(
        bytes32 beforeHash,
        uint128[4] memory timeBounds,
        bytes32 beforeInbox,
        uint256 beforeInboxValueSize,
        bytes32 afterHash,
        bool didInboxInsn,
        bytes32 firstMessage,
        bytes32 lastMessage,
        bytes32 firstLog,
        bytes32 lastLog,
        uint64 gas,
        bytes memory proof
    ) public pure returns (uint256) {
        return
            OneStepProof.validateProof(
                beforeHash,
                timeBounds,
                beforeInbox,
                beforeInboxValueSize,
                afterHash,
                didInboxInsn,
                firstMessage,
                lastMessage,
                firstLog,
                lastLog,
                gas,
                proof
            );
    }
}
