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

pragma solidity ^0.6.11;

import "../arch/OneStepProof2.sol";

contract BufferProofTester is OneStepProof2 {
    event OneStepProofResult(uint64 gas, uint256 totalMessagesRead, bytes32[4] fields);

    function testGet(
        bytes32 buf,
        uint256 loc,
        bytes32[] memory proof
    ) public pure returns (bytes32) {
        return get(buf, loc, proof);
    }

    function testCheckSize(
        bytes32 buf,
        uint256 offset,
        bytes32[] memory proof
    ) public pure returns (bool) {
        bytes32 w = get(buf, offset / 32, proof);
        for (uint256 i = offset % 32; i < 32; i++) {
            if (getByte(w, i) != 0) return false;
        }
        return checkSize(buf, offset / 32, proof);
    }

    function testSet(
        bytes32 buf,
        uint256 loc,
        bytes32 v,
        bytes32[] memory proof,
        uint256 nh,
        bytes32 normal1,
        bytes32 normal2
    ) public pure returns (bytes32) {
        return set(buf, loc, v, proof, nh, normal1, normal2);
    }

    function executeStepTest(
        uint256 initialMessagesRead,
        bytes32 initialSendAcc,
        bytes32 initialLogAcc,
        bytes calldata proof,
        bytes calldata bproof
    ) external {
        (uint64 gas, uint256 totalMessagesRead, bytes32[4] memory fields) =
            OneStepProof2(address(this)).executeStep(
                IBridge(0),
                initialMessagesRead,
                [initialSendAcc, initialLogAcc],
                proof,
                bproof
            );
        emit OneStepProofResult(gas, totalMessagesRead, fields);
    }
}
