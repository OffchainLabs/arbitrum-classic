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
        uint assertionIndex,
        uint64 deadline
    );

    event BisectedAssertion(
        address indexed vmAddress,
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
        emit ContinuedChallenge(_challenge.vmAddress, _challenge.players[1], _assertionToChallenge, _challenge.deadline);
    }

    function bisectAssertion(
        Challenge.Data storage _challenge,
        bytes32 _beforeInbox,
        bytes32[] memory _afterHashAndMessageAndLogsBisections,
        uint32 _totalSteps,
        uint64[2] memory _timeBounds
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

        bytes32 fullHash;
        bytes32[] memory bisectionHashes;
        (fullHash, bisectionHashes) = generateBisectionDataImpl(
            BisectAssertionData(
                uint32(_afterHashAndMessageAndLogsBisections.length / 3 - 1),
                _afterHashAndMessageAndLogsBisections,
                _totalSteps,
                _timeBounds,
                _beforeInbox
            )
        );

        require(
            fullHash == _challenge.challengeState,
            "Does not match prev state"
        );

        _challenge.state = Challenge.State.Bisected;
        _challenge.deadline = uint64(block.number) + uint64(_challenge.challengePeriod);
        _challenge.challengeState = MerkleLib.generateRoot(bisectionHashes);

        emit BisectedAssertion(
            _challenge.vmAddress,
            _challenge.players[0],
            _afterHashAndMessageAndLogsBisections,
            _totalSteps,
            _challenge.deadline
        );
    }

    // bisectionFields:
    // beforeHash
    // firstMessageHash
    // firstLogHash

    // then repeated
    //    afterHash
    //    lastMessageHash
    //    lastLogHash

    struct BisectAssertionData {
        uint32 bisectionCount;
        bytes32[] bisectionFields;
        uint32 totalSteps;
        uint64[2] timeBounds;
        bytes32 beforeInbox;
    }

    function generateBisectionDataImpl(
        BisectAssertionData memory _data
    )
        private
        pure
        returns (bytes32 fullHash, bytes32[] memory hashes)
    {
        bytes32 preconditionHash;
        uint32 stepCount = _data.totalSteps / _data.bisectionCount + 1;
        hashes = new bytes32[](_data.bisectionCount);
        uint j;
        for (j = 0; j < _data.bisectionCount; j++) {
            if (j == _data.totalSteps % _data.bisectionCount) {
                stepCount--;
            }
            preconditionHash = ArbProtocol.generatePreconditionHash(
                _data.bisectionFields[j * 3],
                _data.timeBounds,
                _data.beforeInbox
            );
            hashes[j] = keccak256(
                abi.encodePacked(
                    preconditionHash,
                    ArbProtocol.generateAssertionHash(
                        _data.bisectionFields[(j + 1) * 3],
                        stepCount,
                        _data.bisectionFields[j * 3 + 1],
                        _data.bisectionFields[(j + 1) * 3 + 1],
                        _data.bisectionFields[j * 3 + 2],
                        _data.bisectionFields[(j + 1) * 3 + 2]
                    )
                )
            );

            if (j == 0) {
                fullHash = keccak256(
                    abi.encodePacked(
                        preconditionHash,
                        ArbProtocol.generateAssertionHash(
                            _data.bisectionFields[_data.bisectionCount * 3],
                            _data.totalSteps,
                            _data.bisectionFields[1],
                            _data.bisectionFields[_data.bisectionCount * 3 + 1],
                            _data.bisectionFields[2],
                            _data.bisectionFields[_data.bisectionCount * 3 + 2]
                        )
                    )
                );
            }
        }
        return (fullHash, hashes);
    }
}
