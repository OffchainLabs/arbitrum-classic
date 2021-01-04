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

import "../challenge/IExecutionChallenge.sol";
import "./BisectionChallenge.sol";
import "../challenge/ChallengeUtils.sol";

import "../arch/IOneStepProof.sol";

import "../libraries/MerkleLib.sol";

contract ProgressiveExecutionChallenge is IExecutionChallenge, BisectionChallenge {
    using ChallengeUtils for ChallengeUtils.ExecutionAssertion;

    event OneStepProofCompleted();

    IOneStepProof private executor;

    // Proof was incorrect
    string private constant OSP_PROOF = "OSP_PROOF";

    function connectOneStepProof(address oneStepProof) external {
        executor = IOneStepProof(oneStepProof);
    }

    function hashBisectionAssertion(
        bytes32 machineHash,
        bytes32 inboxAcc,
        bytes32 messageAcc,
        bytes32 logAcc,
        uint256 gasUsed,
        uint256 inboxCount,
        uint256 messageCount,
        uint256 logCount
    ) private pure returns (bytes32) {
        bytes32 stateHash = keccak256(abi.encodePacked(machineHash, inboxAcc, messageAcc, logAcc));
        return keccak256(abi.encodePacked(stateHash, gasUsed, inboxCount, messageCount, logCount));
    }

    // machineFields
    //  initialInbox
    //  initialMessage
    //  initialLog
    function oneStepProofWithMessage(
        bytes32[3] memory _machineFields,
        uint256 _initialGasUsed,
        uint256 _initialInboxCount,
        uint256 _initialMessageCount,
        uint256 _initialLogCount,
        bytes memory _proof,
        uint8 _kind,
        uint256 _blockNumber,
        uint256 _timestamp,
        address _sender,
        uint256 _inboxSeqNum,
        bytes memory _msgData
    ) public onlyOnTurn {
        (uint64 gas, bytes32[5] memory proofFields) = executor.executeStepWithMessage(
            _machineFields,
            _proof,
            _kind,
            _blockNumber,
            _timestamp,
            _sender,
            _inboxSeqNum,
            _msgData
        );

        checkProof(
            gas,
            _machineFields,
            _initialGasUsed,
            _initialInboxCount,
            _initialMessageCount,
            _initialLogCount,
            proofFields
        );
    }

    function oneStepProof(
        bytes32[3] memory _machineFields,
        uint64 _initialGasUsed,
        uint256 _initialInboxCount,
        uint256 _initialMessageCount,
        uint256 _initialLogCount,
        bytes memory _proof
    ) public onlyOnTurn {
        (uint64 gas, bytes32[5] memory proofFields) = executor.executeStep(_machineFields, _proof);

        checkProof(
            gas,
            _machineFields,
            _initialGasUsed,
            _initialInboxCount,
            _initialMessageCount,
            _initialLogCount,
            proofFields
        );
    }

    // fields
    //  initialMachineHash
    //  afterMachineHash
    //  afterInboxHash
    //  afterMessagesHash
    //  afterLogsHash
    function checkProof(
        uint256 gasUsed,
        bytes32[3] memory _machineFields,
        uint256 initialGasUsed,
        uint256 initialInboxCount,
        uint256 initialMessageCount,
        uint256 initialLogCount,
        bytes32[5] memory fields
    ) private {
        // The one step proof already guarantees us that firstMessage and lastMessage
        // are either one or 0 messages apart and the same is true for logs. Therefore
        // we can infer the message count and log count based on whether the fields
        // are equal or not
        bytes32 a2Hash = hashBisectionAssertion(
            fields[1],
            fields[2],
            fields[3],
            fields[4],
            initialGasUsed + gasUsed,
            initialInboxCount + (_machineFields[0] == fields[2] ? 0 : 1),
            initialMessageCount + (_machineFields[1] == fields[3] ? 0 : 1),
            initialLogCount + (_machineFields[2] == fields[4] ? 0 : 1)
        );

        bytes32 a1Hash = hashBisectionAssertion(
            fields[0],
            _machineFields[0],
            _machineFields[1],
            _machineFields[2],
            initialGasUsed,
            initialInboxCount,
            initialMessageCount,
            initialLogCount
        );

        requireMatchesPrevState(ChallengeLib.bisectionChunkHash(1, a1Hash, a2Hash));

        emit OneStepProofCompleted();
        _asserterWin();
    }
}
