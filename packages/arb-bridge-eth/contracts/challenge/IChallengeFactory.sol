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


interface IChallengeFactory {

    function createMessagesChallenge(
        address payable _asserter,
        address payable _challenger,
        uint32 _challengePeriod,
        bytes32 _messagesHash
    )
        external
        returns(address);

    function createPendingTopChallenge(
        address payable _asserter,
        address payable _challenger,
        uint32 _challengePeriod,
        bytes32 _pendingTopHash
    )
        external
        returns(address);

    function createExecutionChallenge(
        address payable _asserter,
        address payable _challenger,
        uint32 _challengePeriod,
        bytes32 _executionHash
    )
        external
        returns(address);

    function generateCloneAddress(
        address asserter,
        address challenger,
        bytes32 codeHash
    )
        external
        view
        returns(address);
}
