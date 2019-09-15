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

import "../libraries/MerkleLib.sol";
import "../libraries/ArbProtocol.sol";


library Bisection {
    enum State {
        NoChallenge,
        Challenged,
        Bisected
    }

    struct Challenge {
        address vmAddress;
        // After bisection this is an array of all sub-assertions
        // After a challenge, the first assertion is the main assertion
        bytes32 challengeState;

        uint128[2] escrows;

        address[2] players;

        uint64 deadline;

        // The current deadline at which the challenge timeouts and a winner is
        // declared. This deadline resets at each step in the challenge
        uint32 challengePeriod;

        State state;
    }

    event ContinuedChallenge (
        address indexed vmAddress,
        address challenger,
        uint assertionIndex
    );

    event BisectedAssertion(
        address indexed vmAddress,
        address bisecter,
        bytes32[] afterHashAndMessageAndLogsBisections,
        uint32 totalSteps,
        uint256[] totalMessageAmounts
    );

    function continueChallenge(
        Challenge storage _challenge,
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

        _challenge.state = Bisection.State.Challenged;
        _challenge.deadline = uint64(block.number) + uint64(_challenge.challengePeriod);
        _challenge.challengeState = _bisectionHash;
        emit ContinuedChallenge(_challenge.vmAddress, _challenge.players[1], _assertionToChallenge);
    }

    // fields
    // _beforeHash
    // _beforeInbox
    function bisectAssertion(
        Challenge storage _challenge,
        bytes32[2] memory _fields,
        bytes32[] memory _afterHashAndMessageAndLogsBisections,
        uint256[] memory _totalMessageAmounts,
        uint32 _totalSteps,
        uint64[2] memory _timeBounds,
        bytes21[] memory _tokenTypes,
        uint256[] memory _beforeBalances
    )
        public
    {
        require(
            _challenge.state == Bisection.State.Challenged,
            "Can only bisect assertion in response to a challenge"
        );
        require(
            _tokenTypes.length == 0 ||
            (_totalMessageAmounts.length % _tokenTypes.length == 0),
            "Incorrect input length"
        );
        require(
            _tokenTypes.length == 0 ||
            ((_afterHashAndMessageAndLogsBisections.length - 1) / 3 ==
            _totalMessageAmounts.length / _tokenTypes.length),
            "Incorrect input length"
        );
        require(_tokenTypes.length == _beforeBalances.length, "Incorrect input length");

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
                uint32((_afterHashAndMessageAndLogsBisections.length - 1) / 3),
                _afterHashAndMessageAndLogsBisections,
                _totalMessageAmounts,
                _totalSteps,
                _fields[0],
                _timeBounds,
                _tokenTypes,
                _beforeBalances,
                _fields[1]
            )
        );

        require(
            fullHash == _challenge.challengeState,
            "Does not match prev state"
        );

        _challenge.state = Bisection.State.Bisected;
        _challenge.deadline = uint64(block.number) + uint64(_challenge.challengePeriod);
        _challenge.challengeState = MerkleLib.generateRoot(bisectionHashes);

        emit BisectedAssertion(
            _challenge.vmAddress,
            _challenge.players[0],
            _afterHashAndMessageAndLogsBisections,
            _totalSteps,
            _totalMessageAmounts
        );
    }

    // bisectionFields:
    // afterHash
    // Message Bisections
    // Logs Bisections

    struct BisectAssertionData {
        uint32 bisectionCount;
        bytes32[] bisectionFields;
        uint256[] totalMessageAmounts;
        uint32 totalSteps;
        bytes32 beforeHash;
        uint64[2] timeBounds;
        bytes21[] tokenTypes;
        uint256[] beforeBalances;
        bytes32 beforeInbox;
    }

    struct GenerateBisectionHashesImplFrame {
        bytes32 beforeHash;
        bytes32 preconditionHash;
        bytes32 fullHash;
        bytes32[] hashes;
        uint256[] coinAmounts;
        uint32 stepCount;
    }

    function generateBisectionDataImpl(
        BisectAssertionData memory _data
    )
        private
        pure
        returns (bytes32, bytes32[] memory)
    {
        GenerateBisectionHashesImplFrame memory frame;
        frame.hashes = new bytes32[](_data.bisectionCount);
        frame.stepCount = _data.totalSteps / _data.bisectionCount + 1;
        uint i;
        uint j;
        frame.beforeHash = _data.beforeHash;
        for (j = 0; j < _data.bisectionCount; j++) {
            if (j == _data.totalSteps % _data.bisectionCount) {
                frame.stepCount--;
            }
            frame.coinAmounts = new uint256[](_data.tokenTypes.length);
            for (i = 0; i < _data.tokenTypes.length; i++) {
                frame.coinAmounts[i] += _data.totalMessageAmounts[j * _data.tokenTypes.length + i];
            }
            frame.preconditionHash = ArbProtocol.generatePreconditionHash(
                frame.beforeHash,
                _data.timeBounds,
                _data.beforeInbox,
                _data.tokenTypes,
                _data.beforeBalances
            );
            for (i = 0; i < _data.tokenTypes.length; i++) {
                _data.beforeBalances[i] -= frame.coinAmounts[i];
            }
            frame.hashes[j] = keccak256(
                abi.encodePacked(
                    frame.preconditionHash,
                    ArbProtocol.generateAssertionHash(
                        _data.bisectionFields[j],
                        frame.stepCount,
                        _data.bisectionFields[_data.bisectionCount + j],
                        _data.bisectionFields[_data.bisectionCount + j + 1],
                        _data.bisectionFields[_data.bisectionCount * 2 + 1 + j],
                        _data.bisectionFields[_data.bisectionCount * 2 + 2 + j],
                        frame.coinAmounts
                    )
                )
            );
            frame.beforeHash = _data.bisectionFields[j];

            if (j == 0) {
                frame.coinAmounts = new uint256[](_data.tokenTypes.length);
                for (i = 0; i < _data.totalMessageAmounts.length; i++) {
                    frame.coinAmounts[i % _data.tokenTypes.length] += _data.totalMessageAmounts[i];
                }
                frame.fullHash = keccak256(
                    abi.encodePacked(
                        frame.preconditionHash,
                        ArbProtocol.generateAssertionHash(
                            _data.bisectionFields[_data.bisectionCount - 1],
                            _data.totalSteps,
                            _data.bisectionFields[_data.bisectionCount],
                            _data.bisectionFields[_data.bisectionCount * 2],
                            _data.bisectionFields[_data.bisectionCount * 2 + 1],
                            _data.bisectionFields[_data.bisectionFields.length - 1],
                            frame.coinAmounts
                        )
                    )
                );
            }
        }
        return (frame.fullHash, frame.hashes);
    }
}
