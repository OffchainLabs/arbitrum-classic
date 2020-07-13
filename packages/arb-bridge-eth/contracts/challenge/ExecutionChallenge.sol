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
    using Hashing for Value.Data;

    event BisectedAssertion(
        bytes32[] machineHashes,
        bool[] didInboxInsns,
        bytes32[] messageAccs,
        bytes32[] logAccs,
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
        bool[] didInboxInsns;
        bytes32[] messageAccs;
        bytes32[] logAccs;
        uint64[] gases;
        uint64 totalSteps;
    }

    function bisectAssertion(
        bytes32 _beforeInbox,
        bytes32[] calldata _machineHashes,
        bool[] calldata _didInboxInsns,
        bytes32[] calldata _messageAccs,
        bytes32[] calldata _logAccs,
        uint64[] calldata _gases,
        uint64 _totalSteps
    ) external asserterAction {
        _bisectAssertion(
            BisectAssertionData(
                _beforeInbox,
                _machineHashes,
                _didInboxInsns,
                _messageAccs,
                _logAccs,
                _gases,
                _totalSteps
            )
        );
    }

    function _bisectAssertion(BisectAssertionData memory _data) private {
        uint256 bisectionCount = _data.machineHashes.length - 1;
        require(bisectionCount == _data.didInboxInsns.length, BIS_INPLEN);
        require(bisectionCount + 1 == _data.messageAccs.length, BIS_INPLEN);
        require(bisectionCount + 1 == _data.logAccs.length, BIS_INPLEN);
        require(bisectionCount == _data.gases.length, BIS_INPLEN);

        uint64 totalGas = 0;
        bool everDidInboxInsn = false;
        for (uint256 i = 0; i < bisectionCount; i++) {
            totalGas += _data.gases[i];
            everDidInboxInsn = everDidInboxInsn || _data.didInboxInsns[i];
        }

        requireMatchesPrevState(
            ChallengeUtils.executionHash(
                _data.totalSteps,
                _data.machineHashes[0],
                _data.beforeInbox,
                _data.machineHashes[bisectionCount],
                everDidInboxInsn,
                totalGas,
                _data.messageAccs[0],
                _data.messageAccs[bisectionCount],
                _data.logAccs[0],
                _data.logAccs[bisectionCount]
            )
        );

        bytes32[] memory hashes = new bytes32[](bisectionCount);
        hashes[0] = ChallengeUtils.executionHash(
            uint32(firstSegmentSize(uint256(_data.totalSteps), bisectionCount)),
            _data.machineHashes[0],
            _data.beforeInbox,
            _data.machineHashes[1],
            _data.didInboxInsns[0],
            _data.gases[0],
            _data.messageAccs[0],
            _data.messageAccs[1],
            _data.logAccs[0],
            _data.logAccs[1]
        );

        for (uint256 i = 1; i < bisectionCount; i++) {
            if (_data.didInboxInsns[i - 1]) {
                _data.beforeInbox = Value.newNone().hash();
            }
            hashes[i] = ChallengeUtils.executionHash(
                uint32(
                    otherSegmentSize(uint256(_data.totalSteps), bisectionCount)
                ),
                _data.machineHashes[i],
                _data.beforeInbox,
                _data.machineHashes[i + 1],
                _data.didInboxInsns[i],
                _data.gases[i],
                _data.messageAccs[i],
                _data.messageAccs[i + 1],
                _data.logAccs[i],
                _data.logAccs[i + 1]
            );
        }

        commitToSegment(hashes);
        asserterResponded();

        emit BisectedAssertion(
            _data.machineHashes,
            _data.didInboxInsns,
            _data.messageAccs,
            _data.logAccs,
            _data.gases,
            _data.totalSteps,
            deadlineTicks
        );
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
        verifyPreCondition(
            _beforeHash,
            _beforeInbox,
            _beforeInboxValueSize,
            _afterHash,
            _didInboxInsns,
            _gas,
            _firstMessage,
            _lastMessage,
            _firstLog,
            _lastLog
        );

        Machine.Data memory endMachine = OneStepProof.validateProof(
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
            Machine.hash(endMachine) == _afterHash,
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

    function verifyPreCondition(
        bytes32 _beforeHash,
        bytes32 _beforeInbox,
        uint256 _beforeInboxValueSize,
        bytes32 _afterHash,
        bool _didInboxInsns,
        uint64 _gas,
        bytes32 _firstMessage,
        bytes32 _lastMessage,
        bytes32 _firstLog,
        bytes32 _lastLog
    ) internal view {
        bytes32 beforeInbox = Value
            .newTuplePreImage(_beforeInbox, _beforeInboxValueSize)
            .hash();
        requireMatchesPrevState(
            ChallengeUtils.executionHash(
                1,
                _beforeHash,
                beforeInbox,
                _afterHash,
                _didInboxInsns,
                _gas,
                _firstMessage,
                _lastMessage,
                _firstLog,
                _lastLog
            )
        );
    }
}
