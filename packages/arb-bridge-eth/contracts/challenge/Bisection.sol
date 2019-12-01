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
        address challenger,
        uint assertionIndex,
        uint64 deadline
    );

    event BisectedAssertion(
        address bisecter,
        bytes32[] afterHashAndMessageAndLogsBisections,
        uint32 totalSteps,
        uint64 deadline
    );

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
        emit ContinuedChallenge(_challenge.players[1], _assertionToChallenge, _challenge.deadline);
    }

    // bisectionFields:
    // beforeHash
    // firstMessageHash
    // firstLogHash

    // then repeated
    //    afterHash
    //    lastMessageHash
    //    lastLogHash

    function bisectAssertion(
        Challenge.Data storage _challenge,
        bytes32 _preData,
        bytes32[] memory _bisectionFields,
        uint32 _totalSteps
    )
        public
    {
        require(
            _challenge.state == Challenge.State.Challenged,
            "Can only bisect assertion in response to a challenge"
        );

        require(
            block.number <= _challenge.deadline,
            "Challenge deadline expired"
        );
        require(
            msg.sender == _challenge.players[0],
            "Only orignal asserter can bisect"
        );

        uint dataLength = _bisectionFields.length;

        require(
            keccak256(
                abi.encodePacked(
                    _preData,
                    _bisectionFields[0],
                    ArbProtocol.generateAssertionHash(
                        _bisectionFields[dataLength - 3],
                        _totalSteps,
                        _bisectionFields[1],
                        _bisectionFields[dataLength - 2],
                        _bisectionFields[2],
                        _bisectionFields[dataLength - 1]
                    )
                )
            ) == _challenge.challengeState,
            "Does not match prev state"
        );

        uint32 bisectionCount = uint32(dataLength / 3 - 1);
        bytes32[] memory hashes = new bytes32[](bisectionCount);
        uint32 stepCount = _totalSteps / bisectionCount + 1;
        for (uint j = 0; j < bisectionCount; j++) {
            if (j == _totalSteps % bisectionCount) {
                stepCount--;
            }
            hashes[j] = keccak256(
                abi.encodePacked(
                    _preData,
                    _bisectionFields[j * 3],
                    ArbProtocol.generateAssertionHash(
                        _bisectionFields[(j + 1) * 3],
                        stepCount,
                        _bisectionFields[j * 3 + 1],
                        _bisectionFields[(j + 1) * 3 + 1],
                        _bisectionFields[j * 3 + 2],
                        _bisectionFields[(j + 1) * 3 + 2]
                    )
                )
            );
        }
        _challenge.challengeState = MerkleLib.generateRoot(hashes);
        _challenge.state = Challenge.State.Bisected;
        _challenge.deadline = uint64(block.number) + uint64(_challenge.challengePeriod);

        emit BisectedAssertion(
            _challenge.players[0],
            _bisectionFields,
            _totalSteps,
            _challenge.deadline
        );
    }
}
