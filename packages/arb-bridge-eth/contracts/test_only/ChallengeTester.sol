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

import "../challenge/IChallengeFactory.sol";
import "../rollup/IRollup.sol";

contract ChallengeTester is IRollup {
    IChallengeFactory private challengeFactory;

    constructor(address challengeFactory_) public {
        challengeFactory = IChallengeFactory(challengeFactory_);
    }

    /* solhint-disable-next-line no-unused-vars */
    function completeChallenge(address, address payable) external override {
        return;
    }

    function startChallenge(
        bytes32 inboxConsistencyHash,
        bytes32 inboxDeltaHash,
        bytes32 executionHash,
        uint256 executionCheckTimeBlocks,
        address payable asserter,
        address payable challenger,
        uint256 challengePeriodBlocks
    ) public {
        challengeFactory.createChallenge(
            inboxConsistencyHash,
            inboxDeltaHash,
            executionHash,
            executionCheckTimeBlocks,
            asserter,
            challenger,
            challengePeriodBlocks
        );
    }
}
