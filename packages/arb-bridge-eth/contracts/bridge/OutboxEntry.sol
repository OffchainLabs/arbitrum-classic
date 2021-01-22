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

import "../libraries/Cloneable.sol";

import "./interfaces/IOutbox.sol";
import "./interfaces/IBridge.sol";

contract OutboxEntry is Cloneable {
    IBridge bridge;
    bytes32 public root;
    uint256 public numRemaining;
    mapping(uint256 => bool) public spentOutput;

    function initialize(
        IBridge _bridge,
        bytes32 _root,
        uint256 _numInBatch
    ) external {
        require(root == 0, "ALREADY_INIT");
        require(_root != 0, "BAD_ROOT");
        bridge = _bridge;
        root = _root;
        numRemaining = _numInBatch;
    }

    function spendOutput(bytes32 _root, uint256 _id) external {
        require(msg.sender == address(bridge), "ONLY_BRIDGE");
        // Make sure this call was generated as a system call from the bridge rather than an L2 call
        require(IOutbox(bridge.activeOutbox()).l2ToL1Sender() == address(0), "ONLY_SYSTEM");
        require(!spentOutput[_id], "ALREADY_SPENT");
        require(_root == root, "BAD_ROOT");
        spentOutput[_id] = true;
        numRemaining--;
        if (numRemaining == 0) {
            safeSelfDestruct(msg.sender);
        }
    }
}
