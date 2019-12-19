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
        returns(address);

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
        returns(address);

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
        returns(address);
}
