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

import "./ChallengeImpl.sol";
import "./IArbChallenge.sol";

import "../vm/IArbitrumVM.sol";


contract ArbChallenge is IArbChallenge {

    event InitiatedChallenge(
        uint64 deadline
    );

    event ContinuedChallenge (
        uint assertionIndex,
        uint64 deadline
    );

    event BisectedAssertion(
        bytes32[] machineHashes,
        bytes32[] messageAccs,
        bytes32[] logAccs,
        uint32 totalSteps,
        uint64 deadline
    );

    event OneStepProofCompleted();

    event TimedOutChallenge (
        bool challengerWrong
    );

    Challenge.Data challenge;

    function init(
        address _vmAddress,
        address[2] calldata _players,
        uint128[2] calldata _escrows,
        uint32 _challengePeriod,
        bytes32 _beforeHash,
        bytes32 _beforeInbox,
        uint64[2] calldata _timeBounds,
        bytes32 _assertionHash
    )
        external
    {
        ChallengeImpl.initializeChallenge(
            challenge,
            _vmAddress,
            _players,
            _escrows,
            _challengePeriod,
            _beforeHash,
            _beforeInbox,
            _timeBounds,
            _assertionHash
        );
    }

    function bisectAssertion(
        bytes32 _preData,
        bytes32[] memory _machineHashes,
        bytes32[] memory _messageAccs,
        bytes32[] memory _logAccs,
        uint32 _totalSteps
    )
        public
    {
        ChallengeImpl.bisectAssertion(
            challenge,
            _preData,
            _machineHashes,
            _messageAccs,
            _logAccs,
            _totalSteps
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
        ChallengeImpl.continueChallenge(
            challenge,
            _assertionToChallenge,
            _proof,
            _bisectionRoot,
            _bisectionHash
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
        bytes memory _proof
    )
        public
    {
        ChallengeImpl.oneStepProof(
            challenge,
            _beforeHash,
            _beforeInbox,
            _timeBounds,
            _afterHash,
            _firstMessage,
            _lastMessage,
            _firstLog,
            _lastLog,
            _proof
        );
        emit OneStepProofCompleted();
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
