// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2020, Offchain Labs, Inc.
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

contract Challenge {
    // Default value must be Uninitialized
    enum Kind { Uninitialized, InboxConsistency, InboxDelta, Execution, StoppedShort }

    Kind kind;
    bytes32 challengeRoot;

    function chooseChallengeType(
        Kind _kind,
        bytes32 _inboxConsistencyHash,
        bytes32 _inboxDeltaHash,
        bytes32 _executionHash
    ) external {
        require(kind == Kind.Uninitialized);

        if (kind == Kind.InboxConsistency) {
            challengeRoot = _inboxConsistencyHash;
        } else if (kind == Kind.InboxDelta) {
            challengeRoot = _inboxDeltaHash;
        } else if (kind == Kind.Execution) {
            challengeRoot = _executionHash;
        } else if (kind == Kind.StoppedShort) {
            challengeRoot = _executionHash;
        } else {
            require(false, "invalid kind");
        }
    }

    function executionStoppedShort(
        bytes32 _inboxConsistencyHash,
        bytes32 _inboxDeltaHash,
        bytes32 _executionHash
    ) external {}
}
