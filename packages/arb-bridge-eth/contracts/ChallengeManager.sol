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

import "./IChallengeManager.sol";
import "./IVMTracker.sol";
import "./MerkleLib.sol";
import "./ArbProtocol.sol";
import "./OneStepProof.sol";


contract ChallengeManager is IChallengeManager {
    enum ChallengeState {
        NoChallenge,
        Challenged,
        Bisected
    }

    struct Challenge {
        // After bisection this is an array of all sub-assertions
        // After a challenge, the first assertion is the main assertion
        bytes32 challengeState;

        uint128[2] escrows;

        address[2] players;

        // The current deadline at which the challenge timeouts and a winner is
        // declared. This deadline resets at each step in the challenge
        uint32 challengePeriod;
    }

    IVMTracker vmTracker;
    mapping(bytes32 => Challenge) challenges;

    constructor(IVMTracker _vmTracker) public {
        vmTracker = _vmTracker;
    }

    function initiateChallenge(
        bytes32 _vmId,
        address[2] calldata _players,
        uint128[2] calldata _escrows,
        uint32 _challengePeriod,
        bytes32 _challengeRoot
    )
        external
    {
        require(msg.sender == address(vmTracker), "Challenge must be forwarded from main contract");
        require(challenges[_vmId].challengeState == 0x00, "There must be no existing challenge");

        challenges[_vmId] = Challenge(
            keccak256(
                abi.encodePacked(
                    _challengeRoot,
                    ChallengeState.Challenged,
                    uint64(block.number) + uint64(_challengePeriod)
                )
            ),
            _escrows,
            _players,
            _challengePeriod
        );
    }

    event BisectedAssertion(
        bytes32 indexed vmId,
        address bisecter,
        bytes32[] afterHashAndMessageAndLogsBisections,
        uint32 totalSteps,
        uint256[] totalMessageAmounts
    );

    // fields
    // _vmId
    // _beforeHash
    // _beforeInbox
    function bisectAssertion(
        bytes32[3] memory _fields,
        bytes32[] memory _afterHashAndMessageAndLogsBisections,
        uint256[] memory _totalMessageAmounts,
        uint32 _totalSteps,
        uint64[2] memory _timeBounds,
        bytes21[] memory _tokenTypes,
        uint256[] memory _beforeBalances,
        uint64 _deadline
    )
        public
    {
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

        bytes32 fullHash;
        bytes32[] memory bisectionHashes;
        (fullHash, bisectionHashes) = generateBisectionDataImpl(
            BisectAssertionData(
                uint32((_afterHashAndMessageAndLogsBisections.length - 1) / 3),
                _afterHashAndMessageAndLogsBisections,
                _totalMessageAmounts,
                _totalSteps,
                _fields[1],
                _timeBounds,
                _tokenTypes,
                _beforeBalances,
                _fields[2]
            )
        );

        Challenge storage challenge = challenges[_fields[0]];
        require(
            keccak256(
                abi.encodePacked(
                    fullHash,
                    ChallengeState.Challenged,
                    _deadline
                )
            ) == challenge.challengeState, "Does not match prev state"
        );

        require(block.number <= _deadline, "Challenge deadline expired");
        require(msg.sender == challenge.players[0], "Only orignal asserter can bisect");

        challenge.challengeState = keccak256(
            abi.encodePacked(
                MerkleLib.generateRoot(bisectionHashes),
                ChallengeState.Bisected,
                uint64(block.number) + uint64(challenge.challengePeriod)
            )
        );

        emit BisectedAssertion(
            _fields[0],
            challenge.players[0],
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

    event ContinuedChallenge (
        bytes32 indexed vmId,
        address challenger,
        uint assertionIndex
    );

    function continueChallenge(
        bytes32 _vmId,
        uint _assertionToChallenge,
        bytes memory _proof,
        uint64 _deadline,
        bytes32 _bisectionRoot,
        bytes32 _bisectionHash
    )
        public
    {
        Challenge storage challenge = challenges[_vmId];
        require(
            keccak256(
                abi.encodePacked(
                    _bisectionRoot,
                    ChallengeState.Bisected,
                    _deadline
                )
            ) == challenge.challengeState, "continueChallenge: Incorrect previous state"
        );
        require(block.number <= _deadline, "Challenge deadline expired");
        require(msg.sender == challenge.players[1], "Only original challenger can continue challenge");
        require(
            MerkleLib.verifyProof(
                _proof,
                _bisectionRoot,
                _bisectionHash,
                _assertionToChallenge + 1
            ),
            "Invalid assertion selected"
        );

        challenge.challengeState = keccak256(
            abi.encodePacked(
                _bisectionHash,
                ChallengeState.Challenged,
                uint64(block.number) + uint64(challenge.challengePeriod)
            )
        );
        emit ContinuedChallenge(_vmId, challenge.players[1], _assertionToChallenge);
    }

    event OneStepProofCompleted(
        bytes32 indexed vmId,
        address asserter,
        bytes proof
    );

    event OneStepProofDebug(bytes32 indexed vmId, bytes32[10] proofData);

    function oneStepProof(
        bytes32 _vmId,
        bytes32[2] memory _beforeHashAndInbox,
        uint64[2] memory _timeBounds,
        bytes21[] memory _tokenTypes,
        uint256[] memory _beforeBalances,
        bytes32[5] memory _afterHashAndMessages,
        uint256[] memory _amounts,
        bytes memory _proof,
        uint64 _deadline
    )
        public
    {
        Challenge storage challenge = challenges[_vmId];
        require(block.number <= _deadline, "One step proof missed deadline");

        require(
            keccak256(
                abi.encodePacked(
                    keccak256(
                        abi.encodePacked(
                            ArbProtocol.generatePreconditionHash(
                                _beforeHashAndInbox[0],
                                _timeBounds,
                                _beforeHashAndInbox[1],
                                _tokenTypes,
                                _beforeBalances
                            ),
                            ArbProtocol.generateAssertionHash(
                                _afterHashAndMessages[0],
                                1,
                                _afterHashAndMessages[1],
                                _afterHashAndMessages[2],
                                _afterHashAndMessages[3],
                                _afterHashAndMessages[4],
                                _amounts
                            )
                        )
                    ),
                ChallengeState.Challenged,
                _deadline
                )
            ) == challenge.challengeState,
            "One step proof with invalid prev state"
        );

        uint correctProof = OneStepProof.validateProof(
            [
                _beforeHashAndInbox[0],
                _beforeHashAndInbox[1],
                _afterHashAndMessages[0],
                _afterHashAndMessages[1],
                _afterHashAndMessages[2],
                _afterHashAndMessages[3],
                _afterHashAndMessages[4]
            ],
            _timeBounds,
            _tokenTypes,
            _beforeBalances,
            _amounts,
            _proof
        );

        require(correctProof == 0, "Proof was incorrect");
        _asserterWin(_vmId, challenge);
        emit OneStepProofCompleted(_vmId, msg.sender, _proof);
    }

    event TimedOutChallenge (
        bytes32 indexed vmId,
        bool challengerWrong
    );

    function asserterTimedOut(bytes32 _vmId, bytes32 _rootHash, uint64 _deadline) public {
        Challenge storage challenge = challenges[_vmId];
        require(
            keccak256(
                abi.encodePacked(
                    _rootHash,
                    ChallengeState.Challenged,
                    _deadline
                )
            ) == challenge.challengeState,
            "Incorrect previous state"
        );
        require(block.number > _deadline, "Deadline hasn't expired");

        _challengerWin(_vmId, challenge);

        emit TimedOutChallenge(_vmId, true);
    }

    function challengerTimedOut(bytes32 _vmId, bytes32 _rootHash, uint64 _deadline) public {
        Challenge storage challenge = challenges[_vmId];
        require(
            keccak256(
                abi.encodePacked(
                    _rootHash,
                    ChallengeState.Bisected,
                    _deadline
                )
            ) == challenge.challengeState,
            "Incorrect previous state"
        );
        require(block.number > _deadline, "Deadline hasn't expired");

        _asserterWin(_vmId, challenge);

        emit TimedOutChallenge(_vmId, false);
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

    function _asserterWin(bytes32 _vmId, Challenge storage challenge) private {
        vmTracker.completeChallenge(
            _vmId,
            challenge.players,
            [
                challenge.escrows[0] + challenge.escrows[1] / 2,
                0
            ]
        );
    }

    function _challengerWin(bytes32 _vmId, Challenge storage challenge) private {
        vmTracker.completeChallenge(
            _vmId,
            challenge.players,
            [
                0,
                challenge.escrows[1] + challenge.escrows[0] / 2
            ]
        );
    }
}
