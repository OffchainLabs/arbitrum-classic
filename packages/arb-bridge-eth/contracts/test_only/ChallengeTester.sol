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

pragma solidity ^0.5.11;

import "../challenge/IChallengeFactory.sol";
import "../vm/IStaking.sol";

contract ChallengeTester is IStaking {
    IChallengeFactory private challengeFactory;

    constructor(address challengeFactory_) public {
        challengeFactory = IChallengeFactory(challengeFactory_);
    }

    function resolveChallenge(address payable, address) external {
        return;
    }

    function startChallenge(
        address payable asserterAddress,
        address payable challengerAddress,
        uint128 challengerPeriodTicks,
        bytes32 challengerDataHash,
        uint256 challengeType
    ) public {
        challengeFactory.createChallenge(
            asserterAddress,
            challengerAddress,
            challengerPeriodTicks,
            challengerDataHash,
            challengeType
        );
    }
}
