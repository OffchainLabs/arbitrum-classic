// SPDX-License-Identifier: Apache-2.0

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

pragma solidity ^0.6.11;

import "../challenge/ChallengeFactory.sol";

contract ChallengeTester {
    address public challenge;
    bool public challengeCompleted;
    address public winner;
    address public loser;
    ChallengeFactory public challengeFactory;
    uint256 public challengeExecutionBisectionDegree = 400;

    constructor(IOneStepProof[] memory _executors) public {
        challengeFactory = new ChallengeFactory(_executors);
    }

    /* solhint-disable-next-line no-unused-vars */
    function completeChallenge(address _winner, address payable _loser) external {
        winner = _winner;
        loser = _loser;
        challengeCompleted = true;
    }

    function startChallenge(
        bytes32 executionHash,
        uint256 maxMessageCount,
        address payable asserter,
        address payable challenger,
        uint256 asserterTimeLeft,
        uint256 challengerTimeLeft,
        ISequencerInbox sequencerBridge,
        IBridge delayedBridge
    ) public {
        challenge = challengeFactory.createChallenge(
            address(this),
            executionHash,
            maxMessageCount,
            asserter,
            challenger,
            asserterTimeLeft,
            challengerTimeLeft,
            sequencerBridge,
            delayedBridge
        );
    }
}
