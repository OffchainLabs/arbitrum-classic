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

import "../libraries/CloneFactory.sol";

import "./IChallengeFactory.sol";
import "./IChallenge.sol";

contract ChallengeFactory is CloneFactory, IChallengeFactory {
    ICloneable public challengeTemplate;
    address public oneStepProofAddress;

    constructor(address _challengeTemplate, address _oneStepProofAddress) public {
        challengeTemplate = ICloneable(_challengeTemplate);
        oneStepProofAddress = _oneStepProofAddress;
    }

    function createChallenge(
        bytes32 _inboxConsistencyHash,
        bytes32 _inboxDeltaHash,
        bytes32 _executionHash,
        uint256 _executionCheckTimeBlocks,
        address payable _asserter,
        address payable _challenger,
        uint256 _challengePeriodBlocks
    ) external returns (address) {
        address clone = createClone(challengeTemplate);
        IChallenge(clone).initializeChallenge(
            oneStepProofAddress,
            msg.sender,
            _inboxConsistencyHash,
            _inboxDeltaHash,
            _executionHash,
            _executionCheckTimeBlocks,
            _asserter,
            _challenger,
            _challengePeriodBlocks
        );
        return address(clone);
    }
}
