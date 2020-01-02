/*
 * Copyright 2019, Offchain Labs, Inc.
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

pragma solidity ^0.5.3;

import "./BisectionChallenge.sol";
import "./ChallengeUtils.sol";

import "../arch/OneStepProof.sol";
import "../arch/Protocol.sol";

import "../libraries/MerkleLib.sol";


contract ExecutionChallenge is BisectionChallenge {

    event BisectedAssertion(
        bytes32[] machineHashes,
        bool[] didInboxInsns,
        bytes32[] messageAccs,
        bytes32[] logAccs,
        uint64[] gases,
        uint32 totalSteps,
        uint256 deadlineTicks
    );

    event OneStepProofCompleted();

    // Incorrect previous state
    string constant BIS_INPLEN = "BIS_INPLEN";
    // Proof was incorrect
    string constant OSP_PROOF = "OSP_PROOF";

    struct BisectAssertionData {
        bytes32 beforeInbox;
        uint64[2] timeBounds;
        bytes32[] machineHashes;
        bool[] didInboxInsns;
        bytes32[] messageAccs;
        bytes32[] logAccs;
        uint64[] gases;
        uint32 totalSteps;
    }

    function bisectAssertion(
        bytes32 _beforeInbox,
        uint64[2] memory _timeBounds,
        bytes32[] memory _machineHashes,
        bool[] memory _didInboxInsns,
        bytes32[] memory _messageAccs,
        bytes32[] memory _logAccs,
        uint64[] memory _gases,
        uint32 _totalSteps
    )
        public
        asserterAction
    {
        _bisectAssertion(
            BisectAssertionData(
                _beforeInbox,
                _timeBounds,
                _machineHashes,
                _didInboxInsns,
                _messageAccs,
                _logAccs,
                _gases,
                _totalSteps
            )
        );
    }

    function oneStepProof(
        bytes32 _beforeHash,
        bytes32 _beforeInbox,
        uint64[2] memory _timeBounds,
        bytes32 _afterHash,
        bool _didInboxInsns,
        bytes32 _firstMessage,
        bytes32 _lastMessage,
        bytes32 _firstLog,
        bytes32 _lastLog,
        uint64  _gas,
        bytes memory _proof
    )
        public
        asserterAction
    {
        requireMatchesPrevState(
            ChallengeUtils.executionHash(
                Protocol.generatePreconditionHash(
                     _beforeHash,
                     _timeBounds,
                     _beforeInbox
                ),
                Protocol.generateAssertionHash(
                    _afterHash,
                    _didInboxInsns,
                    1,
                    _gas,
                    _firstMessage,
                    _lastMessage,
                    _firstLog,
                    _lastLog
                )
            )
        );

        uint correctProof = OneStepProof.validateProof(
            _beforeHash,
            _timeBounds,
            _beforeInbox,
            _afterHash,
            _didInboxInsns,
            _firstMessage,
            _lastMessage,
            _firstLog,
            _lastLog,
            _gas,
            _proof
        );

        require(correctProof == 0, OSP_PROOF);
        emit OneStepProofCompleted();
        _asserterWin();
    }

    function _bisectAssertion(BisectAssertionData memory _data) private {
        uint bisectionCount = _data.machineHashes.length - 1;
        require(bisectionCount == _data.didInboxInsns.length, BIS_INPLEN);
        require(bisectionCount + 1 == _data.messageAccs.length, BIS_INPLEN);
        require(bisectionCount + 1 == _data.logAccs.length, BIS_INPLEN);
        require(bisectionCount == _data.gases.length, BIS_INPLEN);

        uint64 totalGas = 0;
        bool everDidInboxInsn = false;
        for (uint i = 0; i < bisectionCount; i++) {
            totalGas += _data.gases[i];
            everDidInboxInsn = everDidInboxInsn || _data.didInboxInsns[i];
        }

        requireMatchesPrevState(
            ChallengeUtils.executionHash(
                Protocol.generatePreconditionHash(
                     _data.machineHashes[0],
                     _data.timeBounds,
                     _data.beforeInbox
                ),
                Protocol.generateAssertionHash(
                    _data.machineHashes[bisectionCount],
                    everDidInboxInsn,
                    _data.totalSteps,
                    totalGas,
                    _data.messageAccs[0],
                    _data.messageAccs[bisectionCount],
                    _data.logAccs[0],
                    _data.logAccs[bisectionCount]
                )
            )
        );

        bytes32[] memory hashes = new bytes32[](bisectionCount);
        hashes[0] = ChallengeUtils.executionHash(
            Protocol.generatePreconditionHash(
                 _data.machineHashes[0],
                 _data.timeBounds,
                 _data.beforeInbox
            ),
            Protocol.generateAssertionHash(
                _data.machineHashes[1],
                _data.didInboxInsns[0],
                uint32(firstSegmentSize(uint(_data.totalSteps), bisectionCount)),
                _data.gases[0],
                _data.messageAccs[0],
                _data.messageAccs[1],
                _data.logAccs[0],
                _data.logAccs[1]
            )
        );
        bytes32 assertionHash;
        for (uint i = 1; i < bisectionCount; i++) {
            if (_data.didInboxInsns[i-1]) {
                _data.beforeInbox = Value.hashEmptyTuple();
            }
            assertionHash = Protocol.generateAssertionHash(
                _data.machineHashes[i + 1],
                _data.didInboxInsns[i],
                uint32(otherSegmentSize(uint(_data.totalSteps), bisectionCount)),
                _data.gases[i],
                _data.messageAccs[i],
                _data.messageAccs[i + 1],
                _data.logAccs[i],
                _data.logAccs[i + 1]
            );
            hashes[i] = ChallengeUtils.executionHash(
                Protocol.generatePreconditionHash(
                     _data.machineHashes[i],
                     _data.timeBounds,
                     _data.beforeInbox
                ),
                assertionHash
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
}
