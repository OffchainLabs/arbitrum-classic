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
        uint64 totalSteps,
        uint256 deadlineTicks
    );

    event OneStepProofCompleted();

    // Incorrect previous state
    string constant BIS_INPLEN = "BIS_INPLEN";
    // Proof was incorrect
    string constant OSP_PROOF = "OSP_PROOF";

    struct BisectAssertionData {
        bytes32 beforeInbox;
        uint256 beforeInboxValueSize;
        uint128[2] timeBoundsBlocks;
        bytes32[] machineHashes;
        bool[] didInboxInsns;
        bytes32[] messageAccs;
        bytes32[] logAccs;
        uint64[] gases;
        uint64 totalSteps;
    }

    function bisectAssertion(
        bytes32 _beforeInbox,
        uint256 _beforeInboxValueSize,
        uint128[2] memory _timeBoundsBlocks,
        bytes32[] memory _machineHashes,
        bool[] memory _didInboxInsns,
        bytes32[] memory _messageAccs,
        bytes32[] memory _logAccs,
        uint64[] memory _gases,
        uint64 _totalSteps
    )
        public
        asserterAction
    {
        _bisectAssertion(
            BisectAssertionData(
                _beforeInbox,
                _beforeInboxValueSize,
                _timeBoundsBlocks,
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

        bytes32 beforeInbox = Value.hashTuplePreImage(_data.beforeInbox, _data.beforeInboxValueSize);

        bytes32 preconditionHash = Protocol.generatePreconditionHash(
             _data.machineHashes[0],
             _data.timeBoundsBlocks,
             beforeInbox
        );
        bytes32 assertionHash = Protocol.generateAssertionHash(
            _data.machineHashes[bisectionCount],
            everDidInboxInsn,
            totalGas,
            _data.messageAccs[0],
            _data.messageAccs[bisectionCount],
            _data.logAccs[0],
            _data.logAccs[bisectionCount]
        );

        requireMatchesPrevState(
            ChallengeUtils.executionHash(_data.totalSteps, preconditionHash, assertionHash)
        );

        bytes32[] memory hashes = new bytes32[](bisectionCount);
        assertionHash = Protocol.generateAssertionHash(
            _data.machineHashes[1],
            _data.didInboxInsns[0],
            _data.gases[0],
            _data.messageAccs[0],
            _data.messageAccs[1],
            _data.logAccs[0],
            _data.logAccs[1]
        );
        hashes[0] = ChallengeUtils.executionHash(
            uint32(firstSegmentSize(uint(_data.totalSteps), bisectionCount)),
            Protocol.generatePreconditionHash(
                 _data.machineHashes[0],
                 _data.timeBoundsBlocks,
                 beforeInbox
            ),
            assertionHash
        );

        for (uint256 i = 1; i < bisectionCount; i++) {
            if (_data.didInboxInsns[i-1]) {
                beforeInbox = Value.hashEmptyTuple();
            }
            assertionHash = Protocol.generateAssertionHash(
                _data.machineHashes[i + 1],
                _data.didInboxInsns[i],
                _data.gases[i],
                _data.messageAccs[i],
                _data.messageAccs[i + 1],
                _data.logAccs[i],
                _data.logAccs[i + 1]
            );
            hashes[i] = ChallengeUtils.executionHash(
                uint32(otherSegmentSize(uint(_data.totalSteps), bisectionCount)),
                Protocol.generatePreconditionHash(
                     _data.machineHashes[i],
                     _data.timeBoundsBlocks,
                     beforeInbox
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

    function oneStepProof(
        bytes32 _beforeHash,
        bytes32 _beforeInbox,
        uint256 _beforeInboxValueSize,
        uint128[2] memory _timeBoundsBlocks,
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
        setPreCondition(_beforeHash,
                        _timeBoundsBlocks,
                        _beforeInbox,
                        _beforeInboxValueSize,
                        _afterHash,
                        _didInboxInsns,
                        _gas,
                        _firstMessage,
                        _lastMessage,
                        _firstLog,
                        _lastLog);

        uint256 correctProof = OneStepProof.validateProof(
            _beforeHash,
            _timeBoundsBlocks,
            _beforeInbox,
            _beforeInboxValueSize,
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

    function setPreCondition(
        bytes32 _beforeHash,
        uint128[2] memory _timeBoundsBlocks,
        bytes32 _beforeInbox,
        uint256 _beforeInboxValueSize,
        bytes32 _afterHash,
        bool _didInboxInsns,
        uint64  _gas,
        bytes32 _firstMessage,
        bytes32 _lastMessage,
        bytes32 _firstLog,
        bytes32 _lastLog) internal {

        bytes32 beforeInbox = Value.hashTuplePreImage(_beforeInbox, _beforeInboxValueSize);
        bytes32 precondition = Protocol.generatePreconditionHash(
             _beforeHash,
             _timeBoundsBlocks,
            beforeInbox
        );
        requireMatchesPrevState(
            ChallengeUtils.executionHash(
                1,
                precondition,
                Protocol.generateAssertionHash(
                    _afterHash,
                    _didInboxInsns,
                    _gas,
                    _firstMessage,
                    _lastMessage,
                    _firstLog,
                    _lastLog
                )
            )
        );
    }

    function resolveChallengeAsserterWon() internal {
        IStaking(vmAddress).resolveChallenge(asserter, challenger, INVALID_EXECUTION_TYPE);
    }

    function resolveChallengeChallengerWon() internal {
        IStaking(vmAddress).resolveChallenge(challenger, asserter, INVALID_EXECUTION_TYPE);
    }
}
