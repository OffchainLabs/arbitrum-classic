// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2019-2021, Offchain Labs, Inc.
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

import "./Challenge.sol";
import "./IChallengeFactory.sol";
import "../libraries/CloneFactory.sol";

contract ChallengeFactory is CloneFactory, IChallengeFactory {
    ICloneable public challengeTemplate;
    IOneStepProof[] public executors;

    constructor(IOneStepProof[] memory _executors) public {
        challengeTemplate = ICloneable(new Challenge());
        executors = _executors;
    }

    function createChallenge(
        address _resultReceiver,
        bytes32 _executionHash,
        uint256 _maxMessageCount,
        address _asserter,
        address _challenger,
        uint256 _asserterTimeLeft,
        uint256 _challengerTimeLeft,
        ISequencerInbox _sequencerBridge,
        IBridge _delayedBridge
    ) external override returns (address) {
        address clone = createClone(challengeTemplate);
        IChallenge(clone).initializeChallenge(
            executors,
            _resultReceiver,
            _executionHash,
            _maxMessageCount,
            _asserter,
            _challenger,
            _asserterTimeLeft,
            _challengerTimeLeft,
            _sequencerBridge,
            _delayedBridge
        );
        return address(clone);
    }
}
