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
    function executeStep(
        bytes32 beforeInbox,
        uint256 beforeInboxValueSize,
        bytes32 firstMessage,
        bytes32 firstLog,
        bytes memory proof
    )
        public
        pure
        returns (
            bytes32 startHash,
            bytes32 endHash,
            bytes32 logAcc,
            bytes32 messageAcc,
            uint64 gas,
            bool didInboxInsn
        )
    {
        Value.Data memory inbox = Value.newTuplePreImage(
            beforeInbox,
            beforeInboxValueSize
        );
        (
            OneStepProof.AssertionContext memory context,
            uint8 opcode
        ) = OneStepProof.initializeExecutionContext(
            inbox,
            firstMessage,
            firstLog,
            proof
        );

        OneStepProof.executeOp(context, opcode);
        return (
            Machine.hash(context.startMachine),
            Machine.hash(context.afterMachine),
            context.logAcc,
            context.messageAcc,
            context.gas,
            context.didInboxInsn
        );
    }
}
