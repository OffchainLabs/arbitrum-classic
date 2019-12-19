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
import "./IMessagesChallenge.sol";
import "./IPendingTopChallenge.sol";
import "./IExecutionChallenge.sol";


contract ChallengeFactory is CloneFactory, IChallengeFactory {

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

    function createMessagesChallenge(
        address _asserter,
        uint _asserterIndex,
        address _challenger,
        uint _challengerIndex,
        uint32 _challengePeriod,
        bytes32 _bottomHash,
        bytes32 _topHash,
        bytes32 _segmentHash,
        uint32 _chainLength
    )
        external
        returns(address)
    {
        address clone = createClone(messagesChallengeTemplate);
        IMessagesChallenge(clone).init(
            msg.sender,
            _asserter,
            _asserterIndex,
            _challenger,
            _challengerIndex,
            _challengePeriod,
            _bottomHash,
            _topHash,
            _segmentHash,
            _chainLength
        );
        return address(clone);
    }

    function createPendingTopChallenge(
        address _asserter,
        uint _asserterIndex,
        address _challenger,
        uint _challengerIndex,
        uint32 _challengePeriod,
        bytes32 _topHash,
        bytes32 _lowerHash
    )
        external
        returns(address)
    {
        address clone = createClone(pendingTopChallengeTemplate);
        IPendingTopChallenge(clone).init(
            msg.sender,
            _asserter,
            _asserterIndex,
            _challenger,
            _challengerIndex,
            _challengePeriod,
            _topHash,
            _lowerHash
        );
        return address(clone);
    }

    function createExecutionChallenge(
        address _asserter,
        uint _asserterIndex,
        address _challenger,
        uint _challengerIndex,
        uint32 _challengePeriod,
        bytes32 _beforeHash,
        bytes32 _beforeInbox,
        uint64[2] calldata _timeBounds,
        bytes32 _assertionHash
    )
        external
        returns(address)
    {
        address clone = createClone(executionChallengeTemplate);
        IExecutionChallenge(clone).init(
            msg.sender,
            _asserter,
            _asserterIndex,
            _challenger,
            _challengerIndex,
            _challengePeriod,
            _beforeHash,
            _beforeInbox,
            _timeBounds,
            _assertionHash
        );
        return address(clone);
    }
}
