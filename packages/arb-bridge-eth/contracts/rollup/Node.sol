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

import "./INode.sol";
import "../libraries/Cloneable.sol";
import "../libraries/SafeMath.sol";

contract Node is Cloneable, INode {
    using SafeMath for uint256;

    /// @notice Hash of the state of the chain as of this node
    bytes32 public override stateHash;

    /// @notice Hash of the data that can be challenged
    bytes32 public override challengeHash;

    /// @notice Hash of the data that will be committed if this node is confirmed
    bytes32 public override confirmData;

    /// @notice Index of the node previous to this one
    uint256 public override prev;

    /// @notice Deadline at which this node can be confirmed
    uint256 public override deadlineBlock;

    /// @notice Deadline at which a child of this node can be confirmed
    uint256 public override noChildConfirmedBeforeBlock;

    /// @notice Number of stakers staked on this node. This includes real stakers and zombies
    uint256 public override stakerCount;

    /// @notice Mapping of the stakers staked on this node with true if they are staked. This includes real stakers and zombies
    mapping(address => bool) public override stakers;

    /// @notice Address of the rollup contract to which this node belongs
    address public rollup;

    /// @notice This value starts at zero and is set to a value when the first child is created. After that it is constant until the node is destroyed
    uint256 public override firstChildBlock;

    modifier onlyRollup {
        require(msg.sender == rollup, "ROLLUP_ONLY");
        _;
    }

    /**
     * @notice Mark the given staker as staked on this node
     * @param _rollup Initial value of rollup
     * @param _stateHash Initial value of stateHash
     * @param _challengeHash Initial value of challengeHash
     * @param _confirmData Initial value of confirmData
     * @param _prev Initial value of prev
     * @param _deadlineBlock Initial value of deadlineBlock
     */
    function initialize(
        address _rollup,
        bytes32 _stateHash,
        bytes32 _challengeHash,
        bytes32 _confirmData,
        uint256 _prev,
        uint256 _deadlineBlock
    ) external override {
        rollup = _rollup;
        stateHash = _stateHash;
        challengeHash = _challengeHash;
        confirmData = _confirmData;
        prev = _prev;
        deadlineBlock = _deadlineBlock;
        noChildConfirmedBeforeBlock = _deadlineBlock;
    }

    /**
     * @notice Destroy this node
     */
    function destroy() external override onlyRollup {
        selfdestruct(msg.sender);
    }

    /**
     * @notice Mark the given staker as staked on this node
     * @param staker Address of the staker to mark
     * @return The number of stakers after adding this one
     */
    function addStaker(address staker) external override onlyRollup returns (uint256) {
        require(!stakers[staker], "ALREADY_STAKED");
        stakers[staker] = true;
        stakerCount++;
        return stakerCount;
    }

    /**
     * @notice Remove the given staker from this node
     * @param staker Address of the staker to remove
     */
    function removeStaker(address staker) external override onlyRollup {
        require(stakers[staker], "NOT_STAKED");
        stakers[staker] = false;
        stakerCount--;
    }

    function childCreated() external override onlyRollup {
        if (firstChildBlock == 0) {
            firstChildBlock = block.number;
        }
    }

    function newChildConfirmDeadline(uint256 deadline) external override onlyRollup {
        noChildConfirmedBeforeBlock = deadline;
    }

    /**
     * @notice Check whether the current block number has met or passed the node's deadline
     */
    function requirePastDeadline() external view override {
        require(block.number >= deadlineBlock, "BEFORE_DEADLINE");
    }

    /**
     * @notice Check whether the current block number has met or passed deadline for children of this node to be confirmed
     */
    function requirePastChildConfirmDeadline() external view override {
        require(block.number >= noChildConfirmedBeforeBlock, "CHILD_TOO_RECENT");
    }

    /**
     * @notice Check whether the current block number has met or passed the node's deadline
     * @param latestConfirmed Address of the node that should be this node's prev
     * @param stakerAddress Address on the staker that should be staked on this node
     */
    function requireRejectExample(uint256 latestConfirmed, address stakerAddress)
        external
        view
        override
    {
        require(prev == latestConfirmed, "BAD_SUCCESSOR");
        require(stakers[stakerAddress], "BAD_STAKER");
    }
}
