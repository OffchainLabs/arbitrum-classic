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

import "../libraries/CloneFactory.sol";

import "./IChallengeFactory.sol";
import "./IBisectionChallenge.sol";
import "./ChallengeType.sol";


contract ChallengeFactory is CloneFactory, ChallengeType, IChallengeFactory {

    // Invalid challenge type
    string constant INVALID_TYPE = "INVALID_TYPE";

    address public messagesChallengeTemplate;
    address public pendingTopChallengeTemplate;
    address public executionChallengeTemplate;

    constructor(
        address _messagesChallengeTemplate,
        address _pendingTopChallengeTemplate,
        address _executionChallengeTemplate
    ) public {
        messagesChallengeTemplate = _messagesChallengeTemplate;
        pendingTopChallengeTemplate = _pendingTopChallengeTemplate;
        executionChallengeTemplate = _executionChallengeTemplate;
    }

    function generateCloneAddress(
        address asserter,
        address challenger,
        bytes32 codeHash
    )
        external
        view
        returns(address)
    {
        return address(
            bytes20(
                keccak256(
                    abi.encodePacked(
                        byte(0xff),
                        address(this),
                        generateNonce(asserter, challenger),
                        codeHash
                    )
                )
            )
        );
    }

    function createChallenge(
        address payable _asserter,
        address payable _challenger,
        uint256 _challengePeriodTicks,
        bytes32 _challengeHash,
        uint challengeType
    )
        external
        returns(address)
    {
        address challengeTemplate;
        if (challengeType == INVALID_PENDING_TOP_TYPE) {
            challengeTemplate = pendingTopChallengeTemplate;
        } else if (challengeType == INVALID_MESSAGES_TYPE) {
            challengeTemplate = messagesChallengeTemplate;
        } else if (challengeType == INVALID_EXECUTION_TYPE) {
            challengeTemplate = executionChallengeTemplate;
        } else {
            require(false, INVALID_TYPE);
        }
        address clone = create2Clone(
            challengeTemplate,
            generateNonce(address(_asserter), address(_challenger))
        );
        IBisectionChallenge(clone).initializeBisection(
            msg.sender,
            _asserter,
            _challenger,
            _challengePeriodTicks,
            _challengeHash
        );
        return address(clone);
    }

    function generateNonce(address asserter, address challenger) public view returns(uint) {
        return uint(
            keccak256(
                abi.encodePacked(
                    asserter,
                    challenger,
                    msg.sender
                )
            )
        );
    }
}
