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

import "./Challenge.sol";
import "./IChallengeFactory.sol";
import "../libraries/CloneFactory.sol";

contract ChallengeFactory is CloneFactory, IChallengeFactory {
    ICloneable public challengeTemplate;
    address public oneStepProofAddress;
    address public oneStepProof2Address;

    constructor(address _oneStepProofAddress, address _oneStepProof2Address) public {
        challengeTemplate = ICloneable(new Challenge());
        oneStepProofAddress = _oneStepProofAddress;
        oneStepProof2Address = _oneStepProof2Address;
    }

    function createChallenge(
        bytes32 _inboxConsistencyHash,
        bytes32 _inboxDeltaHash,
        bytes32 _executionHash,
        uint256 _executionCheckTimeBlocks,
        address _asserter,
        address _challenger,
        uint256 _challengePeriodBlocks
    ) external override returns (address) {
        address clone = createClone(challengeTemplate);
        IChallenge(clone).initializeChallenge(
            oneStepProofAddress,
            oneStepProof2Address,
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
