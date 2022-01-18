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

import "@openzeppelin/contracts/math/SafeMath.sol";

struct NodeFixed {
    // Hash of the state of the chain as of this node
    bytes32 stateHash;
    // Hash of the data that can be challenged
    bytes32 challengeHash;
    // Hash of the data that will be committed if this node is confirmed
    bytes32 confirmData;
}

struct NodeMutable {
    // CHRIS: do we really want this? eek
    // Index of the node previous to this one
    uint256 prevNum;
    // CHRIS: note about why we have this here for now even though
    // CHRIS: it's immutable - look at moving it back out in further refactoring
    // Deadline at which this node can be confirmed
    uint256 deadlineBlock;
    // Deadline at which a child of this node can be confirmed
    uint256 noChildConfirmedBeforeBlock;
    // Number of stakers staked on this node. This includes real stakers and zombies
    uint256 stakerCount;
    // This value starts at zero and is set to a value when the first child is created. After that it is constant until the node is destroyed or the owner destroys pending nodes
    uint256 firstChildBlock;
    // The number of the latest child of this node to be created
    uint256 latestChildNumber;
}

/**
 * @notice Utility functions for Node
 */
library NodeLib {
    using SafeMath for uint256;

    // CHRIS: comments and names
    function initFixed(
        bytes32 _stateHash,
        bytes32 _challengeHash,
        bytes32 _confirmData
    ) internal pure returns (NodeFixed memory) {
        NodeFixed memory node;
        node.stateHash = _stateHash;
        node.challengeHash = _challengeHash;
        node.confirmData = _confirmData;
        return node;
    }

    // CHRIS: comments and names
    function initMutable(uint256 _prevNum, uint256 _deadlineBlock)
        internal
        pure
        returns (NodeMutable memory)
    {
        NodeMutable memory node;
        node.prevNum = _prevNum;
        node.noChildConfirmedBeforeBlock = _deadlineBlock;
        node.deadlineBlock = _deadlineBlock;
        return node;
    }

    function nodeFixedHash(NodeFixed memory self) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(self.stateHash, self.challengeHash, self.confirmData));
    }

    /**
     * @notice Update child properties
     * @param number The child number to set
     */
    function childCreated(NodeMutable storage self, uint256 number) internal {
        if (self.firstChildBlock == 0) {
            self.firstChildBlock = block.number;
        }
        self.latestChildNumber = number;
    }

    /**
     * @notice Update the child confirmed deadline
     * @param deadline The new deadline to set
     */
    function newChildConfirmDeadline(NodeMutable storage self, uint256 deadline) internal {
        self.noChildConfirmedBeforeBlock = deadline;
    }

    /**
     * @notice Check whether the current block number has met or passed the node's deadline
     */
    function requirePastDeadline(NodeMutable memory self) internal view {
        require(block.number >= self.deadlineBlock, "BEFORE_DEADLINE");
    }

    /**
     * @notice Check whether the current block number has met or passed deadline for children of this node to be confirmed
     */
    function requirePastChildConfirmDeadline(NodeMutable memory self) internal view {
        require(block.number >= self.noChildConfirmedBeforeBlock, "CHILD_TOO_RECENT");
    }
}
