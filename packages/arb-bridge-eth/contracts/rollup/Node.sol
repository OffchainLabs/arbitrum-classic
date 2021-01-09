// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2021, Offchain Labs, Inc.
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

contract Node is Cloneable {
    bytes32 public stateHash;
    bytes32 public challengeHash;
    bytes32 public confirmData;
    uint256 public prev;
    uint256 public deadlineBlock;
    uint256 public stakerCount;
    mapping(address => bool) public stakers;

    address rollup;

    modifier onlyRollup {
        require(msg.sender == rollup, "ROLLUP_ONLY");
        _;
    }

    function initialize(
        address _rollup,
        bytes32 _stateHash,
        bytes32 _challengeHash,
        bytes32 _confirmData,
        uint256 _prev,
        uint256 _deadlineBlock
    ) external {
        rollup = _rollup;
        stateHash = _stateHash;
        challengeHash = _challengeHash;
        confirmData = _confirmData;
        prev = _prev;
        deadlineBlock = _deadlineBlock;
    }

    function destroy() external onlyRollup {
        selfdestruct(msg.sender);
    }

    function addStaker(address staker) external onlyRollup {
        require(!stakers[staker], "ALREADY_STAKED");
        stakers[staker] = true;
        stakerCount++;
    }

    function removeStaker(address staker) external onlyRollup {
        require(stakers[staker], "NOT_STAKED");
        stakers[staker] = false;
        stakerCount--;
    }

    function checkConfirmValid(uint256 totalStakerCount, uint256 latestConfirmed) external view {
        // Verify the block's deadline has passed
        require(block.number >= deadlineBlock, "BEFORE_DEADLINE");

        // Check that prev is latest confirmed
        require(prev == latestConfirmed, "INVALID_PREV");

        // All non-zombie stakers are staked on this node, and no zombie stakers are staked here
        require(stakerCount == totalStakerCount, "NOT_ALL_STAKED");

        // There is at least one non-zombie staker
        require(totalStakerCount > 0, "NO_STAKERS");
    }

    function checkConfirmInvalid(uint256 zombieStakerCount) external view {
        // Verify the block's deadline has passed
        require(block.number >= deadlineBlock, "BEFORE_DEADLINE");

        // Verify that no staker is staked on this node
        require(stakerCount == zombieStakerCount, "HAS_STAKERS");
    }

    function checkConfirmOutOfOrder(uint256 latestConfirmed) external view {
        require(prev != latestConfirmed);
    }
}
