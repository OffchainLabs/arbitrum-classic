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
import "./IExecutionChallenge.sol";

import "../arch/OneStepProof.sol";
import "../arch/Protocol.sol";

import "../libraries/MerkleLib.sol";


contract ExecutionChallenge is BisectionChallenge, IExecutionChallenge {

    event BisectedAssertion(
        bytes32[] machineHashes,
        bytes32[] messageAccs,
        bytes32[] logAccs,
        uint64[]  gases,
        uint32 totalSteps,
        uint64 deadline
    );

    event OneStepProofCompleted();

    // Incorrect previous state
    string constant BIS_INPLEN = "BIS_INPLEN";
    // Proof was incorrect
    string constant OSP_PROOF = "OSP_PROOF";


    function init(
        address _vmAddress,
        address payable _asserter,
        address payable _challenger,
        uint32 _challengePeriod,
        bytes32 _beforeHash,
        bytes32 _beforeInbox,
        uint64[2] calldata _timeBounds,
        bytes32 _assertionHash
    )
        external
    {
        BisectionChallenge.initializeBisection(
            _vmAddress,
            _asserter,
            _challenger,
            _challengePeriod,
            encodeSegment(
                keccak256(
                    abi.encodePacked(
                        _timeBounds[0],
                        _timeBounds[1],
                        _beforeInbox
                    )
                ),
                _beforeHash,
                _assertionHash
            )
        );
    }

    struct BisectAssertionData {
        bytes32 preData;
        bytes32[] machineHashes;
        bytes32[] messageAccs;
        bytes32[] logAccs;
        uint64[] gases;
        uint32 totalSteps;
    }

    function bisectAssertion(
        bytes32 _preData,
        bytes32[] memory _machineHashes,
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
                _preData,
                _machineHashes,
                _messageAccs,
                _logAccs,
                _gases,
                _totalSteps
            )
        );
    }

    function _bisectAssertion(BisectAssertionData memory _data) private {
        uint bisectionCount = _data.machineHashes.length - 1;
        require(bisectionCount + 1 == _data.messageAccs.length, BIS_INPLEN);
        require(bisectionCount + 1 == _data.logAccs.length, BIS_INPLEN);
        require(bisectionCount == _data.gases.length, BIS_INPLEN);

        uint64 totalGas = 0;
        for (uint i = 0; i < bisectionCount; i++) {
            totalGas += _data.gases[i];
        }

        requireMatchesPrevState(
            encodeSegment(
                _data.preData,
                _data.machineHashes[0],
                Protocol.generateAssertionHash(
                    _data.machineHashes[bisectionCount],
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
        hashes[0] = encodeSegment(
            _data.preData,
            _data.machineHashes[0],
            Protocol.generateAssertionHash(
                _data.machineHashes[1],
                firstSegmentSize(_data.totalSteps, bisectionCount),
                _data.gases[0],
                _data.messageAccs[0],
                _data.messageAccs[1],
                _data.logAccs[0],
                _data.logAccs[1]
            )
        );
        for (uint i = 1; i < bisectionCount; i++) {
            hashes[i] = encodeSegment(
                _data.preData,
                _data.machineHashes[i],
                Protocol.generateAssertionHash(
                    _data.machineHashes[i + 1],
                    otherSegmentSize(_data.totalSteps, bisectionCount),
                    _data.gases[i],
                    _data.messageAccs[i],
                    _data.messageAccs[i + 1],
                    _data.logAccs[i],
                    _data.logAccs[i + 1]
                )
            );
        }

        commitToSegment(hashes);
        asserterResponded();

        emit BisectedAssertion(
            _data.machineHashes,
            _data.messageAccs,
            _data.logAccs,
            _data.gases,
            _data.totalSteps,
            deadline
        );
    }

    function oneStepProof(
        bytes32 _beforeHash,
        bytes32 _beforeInbox,
        uint64[2] memory _timeBounds,
        bytes32 _afterHash,
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
            encodeSegment(
                keccak256(
                    abi.encodePacked(
                        _timeBounds[0],
                        _timeBounds[1],
                        _beforeInbox
                    )
                ),
                _beforeHash,
                Protocol.generateAssertionHash(
                    _afterHash,
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

    function encodeSegment(
        bytes32 _preData,
        bytes32 _beforeHash,
        bytes32 _assertionHash

    )
        private
        pure
        returns(bytes32)
    {
        return keccak256(
            abi.encodePacked(
                _preData,
                _beforeHash,
                _assertionHash
            )
        );
    }
}
