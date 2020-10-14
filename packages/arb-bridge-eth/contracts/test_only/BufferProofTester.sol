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

import "../arch/OneStepProof2.sol";

contract BufferProofTester is OneStepProof2 {
    event BufferProofTestEvent();

    function testGet(bytes32 buf, uint loc, bytes32[] memory proof) public pure returns (bytes32) {
        return get(buf, loc, proof);
    }

    function testSet(bytes32 buf, uint loc, bytes32 v, bytes32[] memory proof, uint nh, bytes32 normal1, bytes32 normal2) public pure returns (bytes32) {
        return set(buf, loc, v, proof, nh, normal1, normal2);
    }

    function executeStepTest(
        bytes32 inboxAcc,
        bytes32 messagesAcc,
        bytes32 logsAcc,
        bytes calldata proof,
        bytes calldata bproof
    ) external {
        AssertionContext memory context = initializeExecutionContext(
            inboxAcc,
            messagesAcc,
            logsAcc,
            proof,
            bproof
        );

        executeOp(context);
        emit BufferProofTestEvent();
    }
}
