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
import "../challenge/BisectionChallenge.sol";
import "../challenge/ChallengeUtils.sol";

import "../arch/IOneStepProof.sol";

import "../libraries/MerkleLib.sol";

contract ProgressiveExecutionChallenge is IExecutionChallenge, BisectionChallenge {
    using ChallengeUtils for ChallengeUtils.ExecutionAssertion;

    event BisectedAssertion(bytes32[] assertionHashes, uint256 deadlineTicks);

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
        uint64 gasUsed,
        uint64 messageCount,
        uint64 logCount
    ) private pure returns (bytes32) {
        return
            keccak256(
                abi.encodePacked(
                    machineHash,
                    inboxAcc,
                    messageAcc,
                    logAcc,
                    gasUsed,
                    messageCount,
                    logCount
                )
            );
    }

    function bisectAssertion(
        bytes32 _a1Hash,
        bytes32 _a2Hash,
        uint64 _totalSteps,
        bytes32[] calldata _bisectionHashes
    ) external {
        requireMatchesPrevState(keccak256(abi.encodePacked(_a1Hash, _a2Hash, _totalSteps)));

        uint256 innerCuts = _bisectionHashes.length;
        uint256 totalCuts = innerCuts + 2;
        bytes32[] memory hashes = new bytes32[](totalCuts);
        hashes[0] = keccak256(
            abi.encodePacked(
                _a1Hash,
                _bisectionHashes[0],
                uint64(firstSegmentSize(uint256(_totalSteps), totalCuts))
            )
        );

        uint64 otherStepCount = uint64(otherSegmentSize(uint256(_totalSteps), totalCuts));
        for (uint256 i = 0; i < innerCuts; i++) {
            hashes[i] = keccak256(
                abi.encodePacked(_bisectionHashes[i - 1], _bisectionHashes[i], otherStepCount)
            );
        }
        hashes[totalCuts - 1] = keccak256(
            abi.encodePacked(_bisectionHashes[innerCuts - 1], _a2Hash, otherStepCount)
        );

        commitToSegment(hashes);
        asserterResponded();
    }

    // machineFields
    //  initialInbox
    //  initialMessage
    //  initialLog
    function oneStepProofWithMessage(
        bytes32[3] memory _machineFields,
        uint64 _initialGasUsed,
        uint64 _initialMessageCount,
        uint64 _initialLogCount,
        bytes memory _proof,
        uint8 _kind,
        uint256 _blockNumber,
        uint256 _timestamp,
        address _sender,
        uint256 _inboxSeqNum,
        bytes memory _msgData
    ) public asserterAction {
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
            _initialMessageCount,
            _initialLogCount,
            proofFields
        );
    }

    function oneStepProof(
        bytes32[3] memory _machineFields,
        uint64 _initialGasUsed,
        uint64 _initialMessageCount,
        uint64 _initialLogCount,
        bytes memory _proof
    ) public asserterAction {
        (uint64 gas, bytes32[5] memory proofFields) = executor.executeStep(_machineFields, _proof);

        checkProof(
            gas,
            _machineFields,
            _initialGasUsed,
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
        uint64 gasUsed,
        bytes32[3] memory _machineFields,
        uint64 initialGasUsed,
        uint64 initialMessageCount,
        uint64 initialLogCount,
        bytes32[5] memory fields
    ) private {
        bytes32 a1Hash = hashBisectionAssertion(
            fields[0],
            _machineFields[0],
            _machineFields[1],
            _machineFields[2],
            initialGasUsed,
            initialMessageCount,
            initialLogCount
        );

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
            initialMessageCount + (_machineFields[1] == fields[3] ? 0 : 1),
            initialLogCount + (_machineFields[2] == fields[4] ? 0 : 1)
        );

        requireMatchesPrevState(keccak256(abi.encodePacked(a1Hash, a2Hash, uint64(1))));

        emit OneStepProofCompleted();
        _asserterWin();
    }
}
