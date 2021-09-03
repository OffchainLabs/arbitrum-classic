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

import "../rollup/Rollup.sol";
import "../challenge/IChallenge.sol";

contract ValidatorUtils {
    enum ConfirmType {
        NONE,
        VALID,
        INVALID
    }

    enum NodeConflict {
        NONE,
        FOUND,
        INDETERMINATE,
        INCOMPLETE
    }

    function getConfig(Rollup rollup)
        external
        view
        returns (
            uint256 confirmPeriodBlocks,
            uint256 extraChallengeTimeBlocks,
            uint256 avmGasSpeedLimitPerBlock,
            uint256 baseStake
        )
    {
        confirmPeriodBlocks = rollup.confirmPeriodBlocks();
        extraChallengeTimeBlocks = rollup.extraChallengeTimeBlocks();
        avmGasSpeedLimitPerBlock = rollup.avmGasSpeedLimitPerBlock();
        baseStake = rollup.baseStake();
    }

    function stakerInfo(Rollup rollup, address stakerAddress)
        external
        view
        returns (
            bool isStaked,
            uint256 latestStakedNode,
            uint256 amountStaked,
            address currentChallenge
        )
    {
        return (
            rollup.isStaked(stakerAddress),
            rollup.latestStakedNode(stakerAddress),
            rollup.amountStaked(stakerAddress),
            rollup.currentChallenge(stakerAddress)
        );
    }

    function findStakerConflict(
        Rollup rollup,
        address staker1,
        address staker2,
        uint256 maxDepth
    )
        external
        view
        returns (
            NodeConflict,
            uint256,
            uint256
        )
    {
        uint256 staker1NodeNum = rollup.latestStakedNode(staker1);
        uint256 staker2NodeNum = rollup.latestStakedNode(staker2);
        return findNodeConflict(rollup, staker1NodeNum, staker2NodeNum, maxDepth);
    }

    function checkDecidableNextNode(Rollup rollup) external view returns (ConfirmType) {
        try ValidatorUtils(address(this)).requireConfirmable(rollup) {
            return ConfirmType.VALID;
        } catch {}

        try ValidatorUtils(address(this)).requireRejectable(rollup) {
            return ConfirmType.INVALID;
        } catch {
            return ConfirmType.NONE;
        }
    }

    function requireRejectable(Rollup rollup) external view returns (bool) {
        IRollupUser(address(rollup)).requireUnresolvedExists();
        INode node = rollup.getNode(rollup.firstUnresolvedNode());
        bool inOrder = node.prev() == rollup.latestConfirmed();
        if (inOrder) {
            // Verify the block's deadline has passed
            require(block.number >= node.deadlineBlock(), "BEFORE_DEADLINE");
            rollup.getNode(node.prev()).requirePastChildConfirmDeadline();

            // Verify that no staker is staked on this node
            require(
                node.stakerCount() == IRollupUser(address(rollup)).countStakedZombies(node),
                "HAS_STAKERS"
            );
        }
        return inOrder;
    }

    function requireConfirmable(Rollup rollup) external view {
        IRollupUser(address(rollup)).requireUnresolvedExists();

        uint256 stakerCount = rollup.stakerCount();
        // There is at least one non-zombie staker
        require(stakerCount > 0, "NO_STAKERS");

        uint256 firstUnresolved = rollup.firstUnresolvedNode();
        INode node = rollup.getNode(firstUnresolved);

        // Verify the block's deadline has passed
        node.requirePastDeadline();
        rollup.getNode(node.prev()).requirePastChildConfirmDeadline();

        // Check that prev is latest confirmed
        require(node.prev() == rollup.latestConfirmed(), "INVALID_PREV");
        require(
            node.stakerCount() ==
                stakerCount + IRollupUser(address(rollup)).countStakedZombies(node),
            "NOT_ALL_STAKED"
        );
    }

    function refundableStakers(Rollup rollup) external view returns (address[] memory) {
        uint256 stakerCount = rollup.stakerCount();
        address[] memory stakers = new address[](stakerCount);
        uint256 latestConfirmed = rollup.latestConfirmed();
        uint256 index = 0;
        for (uint256 i = 0; i < stakerCount; i++) {
            address staker = rollup.getStakerAddress(i);
            uint256 latestStakedNode = rollup.latestStakedNode(staker);
            if (
                latestStakedNode <= latestConfirmed && rollup.currentChallenge(staker) == address(0)
            ) {
                stakers[index] = staker;
                index++;
            }
        }
        assembly {
            mstore(stakers, index)
        }
        return stakers;
    }

    function latestStaked(Rollup rollup, address staker) external view returns (uint256, bytes32) {
        uint256 node = rollup.latestStakedNode(staker);
        if (node == 0) {
            node = rollup.latestConfirmed();
        }
        bytes32 acc = rollup.getNodeHash(node);
        return (node, acc);
    }

    function stakedNodes(Rollup rollup, address staker) external view returns (uint256[] memory) {
        uint256[] memory nodes = new uint256[](100000);
        uint256 index = 0;
        for (uint256 i = rollup.latestConfirmed(); i <= rollup.latestNodeCreated(); i++) {
            INode node = rollup.getNode(i);
            if (node.stakers(staker)) {
                nodes[index] = i;
                index++;
            }
        }
        // Shrink array down to real size
        assembly {
            mstore(nodes, index)
        }
        return nodes;
    }

    function findNodeConflict(
        Rollup rollup,
        uint256 node1,
        uint256 node2,
        uint256 maxDepth
    )
        public
        view
        returns (
            NodeConflict,
            uint256,
            uint256
        )
    {
        uint256 firstUnresolvedNode = rollup.firstUnresolvedNode();
        uint256 node1Prev = rollup.getNode(node1).prev();
        uint256 node2Prev = rollup.getNode(node2).prev();

        for (uint256 i = 0; i < maxDepth; i++) {
            if (node1 == node2) {
                return (NodeConflict.NONE, node1, node2);
            }
            if (node1Prev == node2Prev) {
                return (NodeConflict.FOUND, node1, node2);
            }
            if (node1Prev < firstUnresolvedNode && node2Prev < firstUnresolvedNode) {
                return (NodeConflict.INDETERMINATE, 0, 0);
            }
            if (node1Prev < node2Prev) {
                node2 = node2Prev;
                node2Prev = rollup.getNode(node2).prev();
            } else {
                node1 = node1Prev;
                node1Prev = rollup.getNode(node1).prev();
            }
        }
        return (NodeConflict.INCOMPLETE, node1, node2);
    }

    function getStakers(
        Rollup rollup,
        uint256 startIndex,
        uint256 max
    ) public view returns (address[] memory, bool hasMore) {
        uint256 maxStakers = rollup.stakerCount();
        if (startIndex + max <= maxStakers) {
            maxStakers = startIndex + max;
            hasMore = true;
        }

        address[] memory stakers = new address[](maxStakers);
        for (uint256 i = 0; i < maxStakers; i++) {
            stakers[i] = rollup.getStakerAddress(startIndex + i);
        }
        return (stakers, hasMore);
    }

    function timedOutChallenges(
        Rollup rollup,
        uint256 startIndex,
        uint256 max
    ) external view returns (IChallenge[] memory, bool hasMore) {
        (address[] memory stakers, bool hasMoreStakers) = getStakers(rollup, startIndex, max);
        IChallenge[] memory challenges = new IChallenge[](stakers.length);
        uint256 index = 0;
        for (uint256 i = 0; i < stakers.length; i++) {
            address staker = stakers[i];
            address challengeAddr = rollup.currentChallenge(staker);
            if (challengeAddr != address(0)) {
                IChallenge challenge = IChallenge(challengeAddr);
                uint256 timeSinceLastMove = block.number - challenge.lastMoveBlock();
                if (
                    timeSinceLastMove > challenge.currentResponderTimeLeft() &&
                    challenge.asserter() == staker
                ) {
                    challenges[index] = IChallenge(challenge);
                    index++;
                }
            }
        }
        // Shrink array down to real size
        assembly {
            mstore(challenges, index)
        }
        return (challenges, hasMoreStakers);
    }

    // Worst case runtime of O(depth), as it terminates if it switches paths.
    function areUnresolvedNodesLinear(Rollup rollup) external view returns (bool) {
        uint256 end = rollup.latestNodeCreated();
        for (uint256 i = rollup.firstUnresolvedNode(); i <= end; i++) {
            if (i > 0 && rollup.getNode(i).prev() != i - 1) {
                return false;
            }
        }
        return true;
    }
}
