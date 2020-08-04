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
    // fields
    //  startHash
    //  endHash
    //  inboxAcc
    //  logAcc
    //  messageAcc

    function execute(OneStepProof.AssertionContext memory context)
        private
        pure
        returns (bytes32[5] memory fields, uint64 gas)
    {
        OneStepProof.executeOp(context);
        return (
            [
                Machine.hash(context.startMachine),
                Machine.hash(context.afterMachine),
                context.inboxAcc,
                context.logAcc,
                context.messageAcc
            ],
            context.gas
        );
    }

    function executeStep(
        bytes32 firstInbox,
        bytes32 firstMessage,
        bytes32 firstLog,
        bytes memory proof
    ) public pure returns (bytes32[5] memory fields, uint64 gas) {
        OneStepProof.AssertionContext memory context = OneStepProof
            .initializeExecutionContext(
            firstInbox,
            firstMessage,
            firstLog,
            proof
        );
        return execute(context);
    }

    function executeInboxStep(
        bytes32 firstInbox,
        bytes32 firstMessage,
        bytes32 firstLog,
        bytes memory proof,
        uint8 kind,
        uint256 blockNumber,
        uint256 timestamp,
        address sender,
        uint256 inboxSeqNum,
        bytes memory msgData
    ) public pure returns (bytes32[5] memory fields, uint64 gas) {
        OneStepProof.AssertionContext memory context = OneStepProof
            .initializeInboxExecutionContext(
            firstInbox,
            firstMessage,
            firstLog,
            proof,
            kind,
            blockNumber,
            timestamp,
            sender,
            inboxSeqNum,
            msgData
        );
        return execute(context);
    }
}
