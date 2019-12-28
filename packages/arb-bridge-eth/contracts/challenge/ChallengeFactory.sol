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

    function generateNonce(address asserter, address challenger) public view returns(uint) {
        return uint(keccak256(abi.encodePacked(
            asserter,
            challenger,
            msg.sender
        )));
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
        return address(bytes20(keccak256(abi.encodePacked(
            byte(0xff),
            address(this),
            generateNonce(asserter, challenger),
            codeHash
        ))));
    }

    function createMessagesChallenge(
        address payable _asserter,
        address payable _challenger,
        uint32 _challengePeriod,
        bytes32 _bottomHash,
        bytes32 _topHash,
        bytes32 _segmentHash,
        uint32 _chainLength
    )
        external
        returns(address)
    {
        address clone = create2Clone(
            messagesChallengeTemplate,
            generateNonce(address(_asserter), address(_challenger))
        );
        IMessagesChallenge(clone).init(
            msg.sender,
            _asserter,
            _challenger,
            _challengePeriod,
            _bottomHash,
            _topHash,
            _segmentHash,
            _chainLength
        );
        return address(clone);
    }

    function createPendingTopChallenge(
        address payable _asserter,
        address payable _challenger,
        uint32 _challengePeriod,
        bytes32 _topHash,
        bytes32 _lowerHash
    )
        external
        returns(address)
    {
        address clone = create2Clone(
            pendingTopChallengeTemplate,
            generateNonce(address(_asserter), address(_challenger))
        );
        IPendingTopChallenge(clone).init(
            msg.sender,
            _asserter,
            _challenger,
            _challengePeriod,
            _topHash,
            _lowerHash
        );
        return address(clone);
    }

    function createExecutionChallenge(
        address payable _asserter,
        address payable _challenger,
        uint32 _challengePeriod,
        bytes32 _beforeHash,
        bytes32 _beforeInbox,
        uint64[2] calldata _timeBounds,
        bytes32 _assertionHash
    )
        external
        returns(address)
    {
        address clone = create2Clone(
            executionChallengeTemplate,
            generateNonce(address(_asserter), address(_challenger))
        );
        IExecutionChallenge(clone).init(
            msg.sender,
            _asserter,
            _challenger,
            _challengePeriod,
            _beforeHash,
            _beforeInbox,
            _timeBounds,
            _assertionHash
        );
        return address(clone);
    }
}
