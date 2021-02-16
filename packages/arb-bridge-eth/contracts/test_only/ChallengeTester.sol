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

import "../challenge/Challenge.sol";

contract ChallengeTester {
    address private oneStepProofAddress;
    address private oneStepProof2Address;
    address public challenge;
    bool public challengeCompleted;
    address public winner;
    address public loser;

    constructor(address _oneStepProofAddress, address _oneStepProof2Address) public {
        oneStepProofAddress = _oneStepProofAddress;
        oneStepProof2Address = _oneStepProof2Address;
    }

    /* solhint-disable-next-line no-unused-vars */
    function completeChallenge(address _winner, address payable _loser) external {
        winner = _winner;
        loser = _loser;
        challengeCompleted = true;
    }

    function startChallenge(
        bytes32 inboxConsistencyHash,
        bytes32 inboxDeltaHash,
        bytes32 executionHash,
        address payable asserter,
        address payable challenger,
        uint256 asserterTimeLeft,
        uint256 challengerTimeLeft
    ) public {
        Challenge chal = new Challenge();
        chal.initializeChallenge(
            oneStepProofAddress,
            oneStepProof2Address,
            address(this),
            inboxConsistencyHash,
            inboxDeltaHash,
            executionHash,
            asserter,
            challenger,
            asserterTimeLeft,
            challengerTimeLeft
        );
        challenge = address(chal);
    }
}
