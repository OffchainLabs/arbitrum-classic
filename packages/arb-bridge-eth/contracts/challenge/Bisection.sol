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

import "./Challenge.sol";

import "../libraries/MerkleLib.sol";
import "../libraries/ArbProtocol.sol";


library Bisection {

    event ContinuedChallenge (
        address indexed vmAddress,
        address challenger,
        uint assertionIndex
    );

    event BisectedAssertion(
        address indexed vmAddress,
        address bisecter,
        bytes32 preconditionHash,
        bytes32[] bisectionHashes,
        uint32 numSteps
    );

    event BisectedAssertionOther(
        address indexed vmAddress,
        address bisecter,
        bytes32[] bisectionHashes,
        uint32 numSteps,
        uint256[] prevOutputValues
    );

    uint internal constant SPLIT_COUNT = uint(6);

    function continueChallenge(
        Challenge.Data storage _challenge,
        uint _assertionToChallenge,
        bytes memory _proof,
        bytes32 _bisectionRoot,
        bytes32 _bisectionHash
    )
        public
    {
        require(
            _bisectionRoot == _challenge.challengeState,
            "continueChallenge: Incorrect previous state"
        );
        require(
            block.number <= _challenge.deadline,
            "Challenge deadline expired"
        );
        require(
            msg.sender == _challenge.players[1],
            "Only original challenger can continue challenge"
        );
        require(
            MerkleLib.verifyProof(
                _proof,
                _bisectionRoot,
                _bisectionHash,
                _assertionToChallenge + 1
            ),
            "Invalid assertion selected"
        );

        _challenge.state = Challenge.State.Challenged;
        _challenge.deadline = uint64(block.number) + uint64(_challenge.challengePeriod);
        _challenge.challengeState = _bisectionHash;
        emit ContinuedChallenge(_challenge.vmAddress, _challenge.players[1], _assertionToChallenge);
    }

    function bisectAssertionFirst(
        Challenge.Data storage _challenge,
        uint32 _numSteps,
        bytes32 _preconditionHash,
        bytes32 _assertionHash,
        bytes32[] memory _bisectionHashes
    )
        public
    {
        validateBisection(_challenge);

        require(
            keccak256(
                abi.encodePacked(
                    _preconditionHash,
                    _assertionHash,
                    _numSteps
                )
            ) == _challenge.challengeState,
            "Bisector incorrectly revealed bisection segments"
        );

        executeBisection(
            _challenge,
            _preconditionHash,
            _assertionHash,
            _numSteps,
            _bisectionHashes
        );

        emit BisectedAssertion(
            _challenge.vmAddress,
            _challenge.players[0],
            _preconditionHash,
            _bisectionHashes,
            _numSteps
        );
    }

    // fields
    //   _beforeHash
    //   _beforeInbox
    //   _firstMessageHash
    //   _firstLogHash
    //   _a1AfterHash
    //   _a1LastMessageHash
    //   _a1LastLogHash
    //   _a2AfterHash
    //   _a2LastMessageHash
    //   _a2LastLogHash
    function bisectAssertionOther(
        Challenge.Data storage _challenge,
        bytes32[10] memory _fields,
        uint64[2] memory _timeBounds,
        bytes21[] memory _tokenTypes,
        uint256[] memory _beforeBalances,
        uint32 _a1NumSteps,
        uint256[] memory _a1OutputValues,
        uint32 _a2NumSteps,
        uint256[] memory _a2OutputValues,
        bytes32[] memory _bisectionHashes
    )
        public
    {
        _bisectAssertionOther(
            _challenge,
            Challenge.BisectOtherData(
                _fields[0],
                _fields[1],
                _timeBounds,
                _tokenTypes,
                _beforeBalances,
                _fields[2],
                _fields[3],
                _fields[4],
                _a1NumSteps,
                _fields[5],
                _fields[6],
                _a1OutputValues,
                _fields[7],
                _a2NumSteps,
                _fields[8],
                _fields[9],
                _a2OutputValues
            ),
            _bisectionHashes
        );
    }

    function _bisectAssertionOther(
        Challenge.Data storage _challenge,
        Challenge.BisectOtherData memory _data,
        bytes32[] memory _bisectionHashes
    )
        private
    {
        validateBisection(_challenge);

        Challenge.validateBisectionOther(_challenge, _data);

        uint32 numSteps = _data.a2NumSteps - _data.a1NumSteps;

        bytes32 newPre = ArbProtocol.generatePreconditionHash(
            _data.a1AfterHash,
            _data.timeBounds,
            _data.beforeInbox,
            _data.tokenTypes,
            _data.beforeBalances
        );
        executeBisection(
            _challenge,
            newPre,
            ArbProtocol.generateAssertionHash(
                _data.a2AfterHash,
                numSteps,
                _data.firstMessageHash,
                _data.a2LastLogHash,
                _data.firstLogHash,
                _data.a2LastLogHash,
                _data.a2OutputValues
            ),
            numSteps,
            _bisectionHashes
        );

        emit BisectedAssertionOther(
            _challenge.vmAddress,
            _challenge.players[0],
            _bisectionHashes,
            numSteps,
            _data.a1OutputValues
        );
    }

    function validateBisection(Challenge.Data storage _challenge) private view {
        require(
            block.number <= _challenge.deadline,
            "Challenge deadline expired"
        );
        require(
            msg.sender == _challenge.players[0],
            "Only orignal asserter can bisect"
        );

        require(
            _challenge.state == Challenge.State.Challenged,
            "Can only bisect assertion in response to a challenge"
        );
    }

    function executeBisection(
        Challenge.Data storage _challenge,
        bytes32 _preconditionHash,
        bytes32 _assertionHash,
        uint32 _totalSteps,
        bytes32[] memory _bisectionHashes
    )
        private
    {
        uint bisectionCount = _bisectionHashes.length + 1;

        require(
            bisectionCount == SPLIT_COUNT ||
            (_totalSteps < SPLIT_COUNT && bisectionCount == _totalSteps),
            "Incorrect bisection count"
        );

        uint32 stepCount = _totalSteps / uint32(bisectionCount);

        if (_totalSteps % bisectionCount > 0) {
            stepCount++;
        }

        bytes32[] memory hashes = new bytes32[](bisectionCount + 1);
        hashes[0] = keccak256(
            abi.encodePacked(
                _preconditionHash,
                _bisectionHashes[0],
                stepCount
            )
        );
        for (uint i = 1; i < bisectionCount - 1; i++) {
            if (i == _totalSteps % bisectionCount) {
                stepCount--;
            }
            hashes[i] = keccak256(
                abi.encodePacked(
                    _preconditionHash,
                    _bisectionHashes[i - 1],
                    _bisectionHashes[i],
                    stepCount
                )
            );
        }
        if (bisectionCount - 1 == _totalSteps % bisectionCount) {
            stepCount--;
        }
        hashes[bisectionCount] = keccak256(
            abi.encodePacked(
                _preconditionHash,
                _bisectionHashes[bisectionCount - 1],
                _assertionHash,
                stepCount
            )
        );

        _challenge.state = Challenge.State.Bisected;
        _challenge.deadline = uint64(block.number) + uint64(_challenge.challengePeriod);
        _challenge.challengeState = MerkleLib.generateRoot(hashes);
    }
}
