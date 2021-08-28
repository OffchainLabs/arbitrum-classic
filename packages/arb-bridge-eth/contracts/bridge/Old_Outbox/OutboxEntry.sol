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

import "../../libraries/Cloneable.sol";

contract OutboxEntry is Cloneable {
    address outbox;
    bytes32 public root;
    uint256 public numRemaining;
    mapping(bytes32 => bool) public spentOutput;

    function initialize(bytes32 _root, uint256 _numInBatch) external {
        require(outbox == address(0), "ALREADY_INIT");
        require(root == 0, "ALREADY_INIT");
        require(_root != 0, "BAD_ROOT");
        outbox = msg.sender;
        root = _root;
        numRemaining = _numInBatch;
    }

    function spendOutput(bytes32 _root, bytes32 _id) external returns (uint256) {
        require(msg.sender == outbox, "NOT_FROM_OUTBOX");
        require(!spentOutput[_id], "ALREADY_SPENT");
        require(_root == root, "BAD_ROOT");

        spentOutput[_id] = true;
        numRemaining--;

        return numRemaining;
    }

    function destroy() external {
        require(msg.sender == outbox, "NOT_FROM_OUTBOX");
        safeSelfDestruct(msg.sender);
    }
}
