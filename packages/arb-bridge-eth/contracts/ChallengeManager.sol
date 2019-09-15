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

    event OneStepProofCompleted(
        address indexed vmAddress,
        address asserter,
        bytes proof
    );

    event TimedOutChallenge (
        address indexed vmAddress,
        bool challengerWrong
    );

    mapping(address => Bisection.Challenge) challenges;

    function initiateChallenge(
        address[2] calldata _players,
        uint128[2] calldata _escrows,
        uint32 _challengePeriod,
        bytes32 _challengeRoot
    )
        external
    {
        require(challenges[msg.sender].challengeState == 0x00, "There must be no existing challenge");

        challenges[msg.sender] = Bisection.Challenge(
            msg.sender,
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
        address _challengeId,
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
        address _vmAddress,
        uint _assertionToChallenge,
        bytes memory _proof,
        bytes32 _bisectionRoot,
        bytes32 _bisectionHash
    )
        public
    {
        Bisection.Challenge storage challenge = challenges[_vmAddress];
        Bisection.continueChallenge(
            challenge,
            _assertionToChallenge,
            _proof,
            _bisectionRoot,
            _bisectionHash
        );
    }

    function oneStepProof(
        address _vmAddress,
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
        Bisection.Challenge storage challenge = challenges[_vmAddress];
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
        _asserterWin(challenge);
        emit OneStepProofCompleted(_vmAddress, msg.sender, _proof);
    }

    function asserterTimedOut(address _vmAddress) public {
        Bisection.Challenge storage challenge = challenges[_vmAddress];
        require(
            challenge.state == Bisection.State.Challenged,
            "Can only time out asserter if it is their turn"
        );
        require(block.number > challenge.deadline, "Deadline hasn't expired");

        _challengerWin(challenge);

        emit TimedOutChallenge(_vmAddress, true);
    }

    function challengerTimedOut(address _vmAddress) public {
        Bisection.Challenge storage challenge = challenges[_vmAddress];
        require(
            challenge.state == Bisection.State.Bisected,
            "Can only time out challenger if it is their turn"
        );
        require(block.number > challenge.deadline, "Deadline hasn't expired");

        _asserterWin(challenge);

        emit TimedOutChallenge(_vmAddress, false);
    }

    function _asserterWin(Bisection.Challenge storage challenge) private {
        IVMTracker(challenge.vmAddress).completeChallenge(
            challenge.players,
            [
                challenge.escrows[0] + challenge.escrows[1] / 2,
                0
            ]
        );
    }

    function _challengerWin(Bisection.Challenge storage challenge) private {
        IVMTracker(challenge.vmAddress).completeChallenge(
            challenge.players,
            [
                0,
                challenge.escrows[1] + challenge.escrows[0] / 2
            ]
        );
    }
}
