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
import "./ArbProtocol.sol";
import "./OneStepProof.sol";
import "./Bisection.sol";


contract ChallengeManager is IChallengeManager {

    event ContinuedChallenge (
        bytes32 indexed vmId,
        address challenger,
        uint assertionIndex
    );

    event BisectedAssertion(
        bytes32 indexed vmId,
        address bisecter,
        bytes32[] afterHashAndMessageAndLogsBisections,
        uint32 totalSteps,
        uint256[] totalMessageAmounts
    );

    event OneStepProofCompleted(
        bytes32 indexed vmId,
        address asserter,
        bytes proof
    );

    event TimedOutChallenge (
        bytes32 indexed vmId,
        bool challengerWrong
    );

    IVMTracker vmTracker;
    mapping(bytes32 => Bisection.Challenge) challenges;

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

        challenges[_vmId] = Bisection.Challenge(
            _vmId,
            _challengeRoot,
            _escrows,
            _players,
            uint64(block.number) + uint64(_challengePeriod),
            _challengePeriod,
            Bisection.State.Challenged
        );
    }

    // fields
    // _beforeHash
    // _beforeInbox
    function bisectAssertion(
        bytes32 _challengeId,
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
        Bisection.Challenge storage challenge = challenges[_challengeId];
        Bisection.bisectAssertion(
            challenge,
            _fields,
            _afterHashAndMessageAndLogsBisections,
            _totalMessageAmounts,
            _totalSteps,
            _timeBounds,
            _tokenTypes,
            _beforeBalances
        );
    }

    function continueChallenge(
        bytes32 _vmId,
        uint _assertionToChallenge,
        bytes memory _proof,
        bytes32 _bisectionRoot,
        bytes32 _bisectionHash
    )
        public
    {
        Bisection.Challenge storage challenge = challenges[_vmId];
        Bisection.continueChallenge(
            challenge,
            _assertionToChallenge,
            _proof,
            _bisectionRoot,
            _bisectionHash
        );
    }

    function oneStepProof(
        bytes32 _vmId,
        bytes32[2] memory _beforeHashAndInbox,
        uint64[2] memory _timeBounds,
        bytes21[] memory _tokenTypes,
        uint256[] memory _beforeBalances,
        bytes32[5] memory _afterHashAndMessages,
        uint256[] memory _amounts,
        bytes memory _proof
    )
        public
    {
        Bisection.Challenge storage challenge = challenges[_vmId];
        require(
            challenge.state == Bisection.State.Challenged,
            "Can only one step proof following a single step challenge"
        );
        require(block.number <= challenge.deadline, "One step proof missed deadline");

        require(
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

    function asserterTimedOut(bytes32 _vmId) public {
        Bisection.Challenge storage challenge = challenges[_vmId];
        require(
            challenge.state == Bisection.State.Challenged,
            "Can only time out asserter if it is their turn"
        );
        require(block.number > challenge.deadline, "Deadline hasn't expired");

        _challengerWin(_vmId, challenge);

        emit TimedOutChallenge(_vmId, true);
    }

    function challengerTimedOut(bytes32 _vmId) public {
        Bisection.Challenge storage challenge = challenges[_vmId];
        require(
            challenge.state == Bisection.State.Bisected,
            "Can only time out challenger if it is their turn"
        );
        require(block.number > challenge.deadline, "Deadline hasn't expired");

        _asserterWin(_vmId, challenge);

        emit TimedOutChallenge(_vmId, false);
    }

    function _asserterWin(bytes32 _vmId, Bisection.Challenge storage challenge) private {
        vmTracker.completeChallenge(
            _vmId,
            challenge.players,
            [
                challenge.escrows[0] + challenge.escrows[1] / 2,
                0
            ]
        );
    }

    function _challengerWin(bytes32 _vmId, Bisection.Challenge storage challenge) private {
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
