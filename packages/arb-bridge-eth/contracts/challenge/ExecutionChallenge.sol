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

import "./BisectionChallenge.sol";
import "./ChallengeUtils.sol";

import "../arch/OneStepProof.sol";

import "../libraries/MerkleLib.sol";

contract ExecutionChallenge is BisectionChallenge {
    using ChallengeUtils for ChallengeUtils.ExecutionAssertion;
    using Hashing for Value.Data;

    event BisectedAssertion(
        bytes32[] machineHashes,
        uint32 inboxInsnIndex,
        bytes32[] messageAccs,
        bytes32[] logAccs,
        uint64[] outCounts,
        uint64[] gases,
        uint64 totalSteps,
        uint256 deadlineTicks
    );

    event OneStepProofCompleted();

    // Incorrect previous state
    string private constant BIS_INPLEN = "BIS_INPLEN";
    // Proof was incorrect
    string private constant OSP_PROOF = "OSP_PROOF";

    struct BisectAssertionData {
        bytes32 beforeInbox;
        bytes32[] machineHashes;
        uint32 inboxInsnIndex;
        bytes32[] messageAccs;
        bytes32[] logAccs;
        uint64[] outCounts;
        uint64[] gases;
        uint64 totalSteps;
    }

    // @param inboxInsnIndex is 0 if the assertion didn't include an inbox instruction, and otherwise the index of the segment including it plus 1
    function bisectAssertion(
        bytes32 _beforeInbox,
        bytes32[] memory _machineHashes,
        uint32 inboxInsnIndex,
        bytes32[] memory _messageAccs,
        bytes32[] memory _logAccs,
        uint64[] memory _outCounts,
        uint64[] memory _gases,
        uint64 _totalSteps
    ) public asserterAction {
        BisectAssertionData memory bisection = BisectAssertionData(
            _beforeInbox,
            _machineHashes,
            inboxInsnIndex,
            _messageAccs,
            _logAccs,
            _outCounts,
            _gases,
            _totalSteps
        );
        _bisectAssertion(bisection);
    }

    function _checkBisectionPrecondition(BisectAssertionData memory _data)
        private
        view
    {
        uint256 bisectionCount = _data.machineHashes.length - 1;
        require(bisectionCount + 1 == _data.messageAccs.length, BIS_INPLEN);
        require(bisectionCount + 1 == _data.logAccs.length, BIS_INPLEN);
        require(bisectionCount == _data.gases.length, BIS_INPLEN);
        require(bisectionCount * 2 == _data.outCounts.length, BIS_INPLEN);
        uint64 totalGas = 0;
        uint64 totalMessageCount = 0;
        uint64 totalLogCount = 0;
        for (uint256 i = 0; i < bisectionCount; i++) {
            totalGas += _data.gases[i];
            totalMessageCount += _data.outCounts[i];
            totalLogCount += _data.outCounts[bisectionCount + i];
        }

        requireMatchesPrevState(
            ChallengeUtils
                .ExecutionAssertion(
                _data
                    .totalSteps,
                _data.machineHashes[0],
                _data
                    .beforeInbox,
                _data.machineHashes[bisectionCount],
                _data.inboxInsnIndex > 0 ? true : false,
                totalGas,
                _data.messageAccs[0],
                _data.messageAccs[bisectionCount],
                totalMessageCount,
                _data.logAccs[0],
                _data.logAccs[bisectionCount],
                totalLogCount
            )
                .hash()
        );
    }

    function _generateBisectionHash(
        BisectAssertionData memory data,
        uint32 stepCount,
        uint256 bisectionCount,
        uint256 i
    ) private pure returns (bytes32) {
        return
            ChallengeUtils
                .ExecutionAssertion(
                stepCount,
                data.machineHashes[i],
                data
                    .beforeInbox,
                data.machineHashes[i + 1],
                data.inboxInsnIndex == i + 1,
                data.gases[i],
                data.messageAccs[i],
                data.messageAccs[i + 1],
                data.outCounts[i],
                data.logAccs[i],
                data.logAccs[i + 1],
                data.outCounts[bisectionCount + i]
            )
                .hash();
    }

    function _emitBisectionEvent(BisectAssertionData memory data) private {
        emit BisectedAssertion(
            data.machineHashes,
            data.inboxInsnIndex,
            data.messageAccs,
            data.logAccs,
            data.outCounts,
            data.gases,
            data.totalSteps,
            deadlineTicks
        );
    }

    function _bisectAssertion(BisectAssertionData memory _data) private {
        uint256 bisectionCount = _data.machineHashes.length - 1;
        _checkBisectionPrecondition(_data);
        bytes32[] memory hashes = new bytes32[](bisectionCount);
        hashes[0] = _generateBisectionHash(
            _data,
            uint32(firstSegmentSize(uint256(_data.totalSteps), bisectionCount)),
            bisectionCount,
            0
        );
        for (uint256 i = 1; i < bisectionCount; i++) {
            // If the previous segment called inbox, set the inbox to the empty tuple
            if (_data.inboxInsnIndex == i) {
                _data.beforeInbox = Value.newEmptyTuple().hash();
            }
            hashes[i] = _generateBisectionHash(
                _data,
                uint32(
                    otherSegmentSize(uint256(_data.totalSteps), bisectionCount)
                ),
                bisectionCount,
                i
            );
        }

        commitToSegment(hashes);
        asserterResponded();

        _emitBisectionEvent(_data);
    }

    function oneStepProof(
        bytes32 _beforeHash,
        bytes32 _beforeInbox,
        uint256 _beforeInboxValueSize,
        bytes32 _afterHash,
        bool _didInboxInsns,
        bytes32 _firstMessage,
        bytes32 _lastMessage,
        bytes32 _firstLog,
        bytes32 _lastLog,
        uint64 _gas,
        bytes memory _proof
    ) public asserterAction {
        bytes32 beforeInbox = Value
            .newTuplePreImage(_beforeInbox, _beforeInboxValueSize)
            .hash();
        // The one step proof already guarantees us that _firstMessage and _lastMessage
        // are either one or 0 messages apart and the same is true for logs. Therefore
        // we can infer the message count and log count based on whether the fields
        // are equal or not
        ChallengeUtils.ExecutionAssertion memory assertion = ChallengeUtils
            .ExecutionAssertion(
            1,
            _beforeHash,
            beforeInbox,
            _afterHash,
            _didInboxInsns,
            _gas,
            _firstMessage,
            _lastMessage,
            _firstMessage == _lastMessage ? 0 : 1,
            _firstLog,
            _lastLog,
            _firstLog == _lastLog ? 0 : 1
        );
        require(
            _firstMessage == _lastMessage || _firstLog == _lastLog,
            "sent both logs and messages"
        );
        requireMatchesPrevState(assertion.hash());
        OneStepProof.AssertionContext memory context = OneStepProof
            .validateProof(
            _beforeHash,
            _beforeInbox,
            _beforeInboxValueSize,
            _didInboxInsns,
            _firstMessage,
            _lastMessage,
            _firstLog,
            _lastLog,
            _gas,
            _proof
        );

        require(
            Machine.hash(context.machine) == _afterHash,
            "Proof had non matching end state"
        );

        // require(
        //     _data.afterHash == endMachine.hash(),
        //     string(abi.encodePacked("Proof had non matching end state: ", endMachine.toString(),
        //     " afterHash = ", DebugPrint.bytes32string(_data.afterHash), "\nendMachine = ", DebugPrint.bytes32string(endMachine.hash())))
        // );

        emit OneStepProofCompleted();
        _asserterWin();
    }
}
