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
        uint assertionIndex,
        uint64 deadline
    );

    event BisectedAssertion(
        bytes32[] afterHashAndMessageAndLogsBisections,
        uint32 totalSteps,
        uint64 deadline
    );

    // Can only continue challenge in response to bisection
    string constant CON_STATE = "CON_STATE";

    // Incorrect previous state
    string constant CON_PREV = "CON_PREV";

    // deadline expired
    string constant CON_DEADLINE = "CON_DEADLINE";

    // Only original challenger can continue challenge
    string constant CON_SENDER = "CON_SENDER";

    // Invalid assertion selected
    string constant CON_PROOF = "CON_PROOF";

    // Can only bisect assertion in response to a challenge
    string constant BIS_STATE = "BIS_STATE";

    // Incorrect previous state
    string constant BIS_PREV = "BIS_PREV";

    // deadline expired
    string constant BIS_DEADLINE = "BIS_DEADLINE";

    // Only original asserter can continue bisect
    string constant BIS_SENDER = "BIS_SENDER";

    function continueChallenge(
        Challenge.Data storage _challenge,
        uint _assertionToChallenge,
        bytes memory _proof,
        bytes32 _bisectionRoot,
        bytes32 _bisectionHash
    )
        public
    {
        require(_challenge.state == Challenge.State.Bisected, CON_STATE);
        require(_bisectionRoot == _challenge.challengeState, CON_PREV);
        require(block.number <= _challenge.deadline, CON_DEADLINE);
        require(msg.sender == _challenge.players[1], CON_SENDER);
        require(
            MerkleLib.verifyProof(
                _proof,
                _bisectionRoot,
                _bisectionHash,
                _assertionToChallenge + 1
            ),
            CON_PROOF
        );

        _challenge.state = Challenge.State.Challenged;
        uint64 deadline = uint64(block.number) + uint64(_challenge.challengePeriod);
        _challenge.deadline = deadline;
        _challenge.challengeState = _bisectionHash;
        emit ContinuedChallenge(_assertionToChallenge, deadline);
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
        require(Challenge.State.Challenged == _challenge.state, BIS_STATE);
        require(block.number <= _challenge.deadline, BIS_DEADLINE);
        require(msg.sender == _challenge.players[0], BIS_SENDER);

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
            BIS_PREV
        );

        uint32 bisectionCount = uint32(dataLength / 3 - 1);
        bytes32[] memory hashes = new bytes32[](bisectionCount);
        hashes[0] = keccak256(
            abi.encodePacked(
                _preData,
                _bisectionFields[0],
                ArbProtocol.generateAssertionHash(
                    _bisectionFields[3],
                    _totalSteps / bisectionCount + _totalSteps%bisectionCount,
                    _bisectionFields[1],
                    _bisectionFields[4],
                    _bisectionFields[2],
                    _bisectionFields[5]
                )
            )
        );
        for (uint j = 1; j < bisectionCount; j++) {
            hashes[j] = keccak256(
                abi.encodePacked(
                    _preData,
                    _bisectionFields[j * 3],
                    ArbProtocol.generateAssertionHash(
                        _bisectionFields[(j + 1) * 3],
                        _totalSteps / bisectionCount,
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
        uint64 deadline = uint64(block.number) + uint64(_challenge.challengePeriod);
        _challenge.deadline = deadline;

        emit BisectedAssertion(
            _bisectionFields,
            _totalSteps,
            deadline
        );
    }
}
