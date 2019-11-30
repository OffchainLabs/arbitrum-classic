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

import "./OneStepProof.sol";
import "./Bisection.sol";
import "./IArbChallenge.sol";

import "../vm/IArbitrumVM.sol";


contract ArbChallenge is IArbChallenge {

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

    event OneStepProofCompleted(
        address asserter,
        bytes proof
    );

    event TimedOutChallenge (
        bool challengerWrong
    );

    event InitiatedChallenge(
        address challenger,
        uint64 deadline
    );

    Challenge.Data challenge;

    function init(
        address vmAddress,
        address[2] calldata _players,
        uint128[2] calldata _escrows,
        uint32 _challengePeriod,
        bytes32 _challengeRoot
    )
        external
    {
        uint64 deadline = uint64(block.number) + uint64(_challengePeriod);
        challenge = Challenge.Data(
            vmAddress,
            _challengeRoot,
            _escrows,
            _players,
            deadline,
            _challengePeriod,
            Challenge.State.Challenged
        );

        emit InitiatedChallenge(
            _players[1],
            deadline
        );
    }

    function bisectAssertion(
        bytes32 _beforeInbox,
        bytes32[] memory _afterHashAndMessageAndLogsBisections,
        uint32 _totalSteps,
        uint64[2] memory _timeBounds
    )
        public
    {
        Bisection.bisectAssertion(
            challenge,
            _beforeInbox,
            _afterHashAndMessageAndLogsBisections,
            _totalSteps,
            _timeBounds
        );
    }

    function continueChallenge(
        uint _assertionToChallenge,
        bytes memory _proof,
        bytes32 _bisectionRoot,
        bytes32 _bisectionHash
    )
        public
    {
        Bisection.continueChallenge(
            challenge,
            _assertionToChallenge,
            _proof,
            _bisectionRoot,
            _bisectionHash
        );
    }

    function oneStepProof(
        bytes32[2] memory _beforeHashAndInbox,
        uint64[2] memory _timeBounds,
        bytes32[5] memory _afterHashAndMessages,
        bytes memory _proof
    )
        public
    {
        OneStepProof.oneStepProof(
            challenge,
            _beforeHashAndInbox,
            _timeBounds,
            _afterHashAndMessages,
            _proof
        );
        emit OneStepProofCompleted(msg.sender, _proof);
        _asserterWin();
    }

    function asserterTimedOut() public {
        require(
            challenge.state == Challenge.State.Challenged,
            "Can only time out asserter if it is their turn"
        );
        require(block.number > challenge.deadline, "Deadline hasn't expired");

        emit TimedOutChallenge(true);
        _challengerWin();
    }

    function challengerTimedOut() public {
        require(
            challenge.state == Challenge.State.Bisected,
            "Can only time out challenger if it is their turn"
        );
        require(block.number > challenge.deadline, "Deadline hasn't expired");
        emit TimedOutChallenge(false);
        _asserterWin();
    }

    function _asserterWin() private {
        IArbitrumVM(challenge.vmAddress).completeChallenge(
            challenge.players,
            [
                challenge.escrows[0] + challenge.escrows[1] / 2,
                0
            ]
        );
        selfdestruct(msg.sender);
    }

    function _challengerWin() private {
        IArbitrumVM(challenge.vmAddress).completeChallenge(
            challenge.players,
            [
                0,
                challenge.escrows[1] + challenge.escrows[0] / 2
            ]
        );
        selfdestruct(msg.sender);
    }
}
