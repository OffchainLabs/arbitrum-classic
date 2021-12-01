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
pragma experimental ABIEncoderV2;

import "./INode.sol";
import "@openzeppelin/contracts/math/SafeMath.sol";

struct Node {
    /// @notice Hash of the state of the chain as of this node
    bytes32 stateHash;

    /// @notice Hash of the data that can be challenged
    bytes32 challengeHash;

    /// @notice Hash of the data that will be committed if this node is confirmed
    bytes32 confirmData;

    /// @notice Index of the node previous to this one
    uint256 prev;

    /// @notice Deadline at which this node can be confirmed
    uint256 deadlineBlock;

    /// @notice Deadline at which a child of this node can be confirmed
    uint256 noChildConfirmedBeforeBlock;

    /// @notice Number of stakers staked on this node. This includes real stakers and zombies
    uint256 stakerCount;

    /// @notice Mapping of the stakers staked on this node with true if they are staked. This includes real stakers and zombies
    mapping(address => bool) stakers;

    /// @notice Address of the rollup contract to which this node belongs
    address rollup;

    /// @notice This value starts at zero and is set to a value when the first child is created. After that it is constant until the node is destroyed or the owner destroys pending nodes
    uint256 firstChildBlock;

    /// @notice The number of the latest child of this node to be created
    uint256 latestChildNumber;
}

library NodeOps {
    using SafeMath for uint256;

    // CHRIS: no longer necessary
    // modifier onlyRollup() {
    //     require(msg.sender == rollup, "ROLLUP_ONLY");
    //     _;
    // }

    /**
     * @notice Mark the given staker as staked on this node
     * @param _stateHash Initial value of stateHash
     * @param _challengeHash Initial value of challengeHash
     * @param _confirmData Initial value of confirmData
     * @param _prev Initial value of prev
     * @param _deadlineBlock Initial value of deadlineBlock
     */
    function initialize(
        bytes32 _stateHash,
        bytes32 _challengeHash,
        bytes32 _confirmData,
        uint256 _prev,
        uint256 _deadlineBlock
    ) internal returns (Node memory) {
        // CHRIS: remove this function?
        Node memory node;
        node.stateHash = _stateHash;
        node.challengeHash = _challengeHash;
        node.confirmData = _confirmData;
        node.prev = _prev;
        node.deadlineBlock = _deadlineBlock;
        node.noChildConfirmedBeforeBlock = _deadlineBlock;
        return node;
    }

    // CHRIS: updates to docs re Node contract

    // destroy by removing from the mapping? we used to keep the address there
    // CHRIS:check usages of destroy again
    /**
     * @notice Destroy this node
     */
    // function destroy() external override onlyRollup {
    //     safeSelfDestruct(msg.sender);
    // }


    // CHRIS: shouldn't all the Node be passed by storage instead of memory?
    // CHRIS: we want to update in place don't we? check the passing/copying semantics

    // CHRIS: check usages of each of these functions - they were protected before so we need to make sure they're not publicly accessible

    /**
     * @notice Mark the given staker as staked on this node
     * @param staker Address of the staker to mark
     * @return The number of stakers after adding this one
     */
    function addStaker(Node storage self, address staker) internal returns (uint256) {
        require(!self.stakers[staker], "ALREADY_STAKED");
        self.stakers[staker] = true;
        self.stakerCount++;
        return self.stakerCount;
    }

    // CHRIS: all the functions in here should be internal

    /**
     * @notice Remove the given staker from this node
     * @param staker Address of the staker to remove
     */
    function removeStaker(Node storage self, address staker) internal {
        require(self.stakers[staker], "NOT_STAKED");
        self.stakers[staker] = false;
        self.stakerCount--;
    }

    function childCreated(Node memory self, uint256 number) internal {
        if (self.firstChildBlock == 0) {
            self.firstChildBlock = block.number;
        }
        self.latestChildNumber = number;
    }

    // CHRIS: we should have docs on each of these
    // CHRIS: should we have tests on any of these?

    function newChildConfirmDeadline(Node storage self, uint256 deadline) internal {
        self.noChildConfirmedBeforeBlock = deadline;
    }

    // CHRIS: update to view/pure where applicable?

    /**
     * @notice Check whether the current block number has met or passed the node's deadline
     */
    function requirePastDeadline(Node memory self) internal view {
        require(block.number >= self.deadlineBlock, "BEFORE_DEADLINE");
    }

    /**
     * @notice Check whether the current block number has met or passed deadline for children of this node to be confirmed
     */
    function requirePastChildConfirmDeadline(Node memory self) internal view {
        require(block.number >= self.noChildConfirmedBeforeBlock, "CHILD_TOO_RECENT");
    }
}
