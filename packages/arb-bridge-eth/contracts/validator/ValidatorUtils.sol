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
    enum ConfirmType { NONE, VALID, INVALID }

    enum NodeConflict { NONE, FOUND, INDETERMINATE, INCOMPLETE }

    function getConfig(Rollup rollup)
        external
        view
        returns (
            uint256 confirmPeriodBlocks,
            uint256 extraChallengeTimeBlocks,
            uint256 arbGasSpeedLimitPerBlock,
            uint256 baseStake,
            address stakeToken
        )
    {
        confirmPeriodBlocks = rollup.confirmPeriodBlocks();
        extraChallengeTimeBlocks = rollup.extraChallengeTimeBlocks();
        arbGasSpeedLimitPerBlock = rollup.arbGasSpeedLimitPerBlock();
        baseStake = rollup.baseStake();
        stakeToken = rollup.stakeToken();
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

    function checkDecidableNextNode(
        Rollup rollup,
        uint256 startNodeOffset,
        uint256 maxNodeCount,
        uint256 startStakerIndex,
        uint256 maxStakerCount
    )
        external
        view
        returns (
            ConfirmType,
            uint256,
            address
        )
    {
        try ValidatorUtils(address(this)).requireConfirmable(rollup) {
            return (ConfirmType.VALID, 0, address(0));
        } catch {}

        try
            ValidatorUtils(address(this)).requireRejectableNextNode(
                rollup,
                startNodeOffset,
                maxNodeCount,
                startStakerIndex,
                maxStakerCount
            )
        returns (uint256 successorWithStake, address stakerAddress) {
            return (ConfirmType.INVALID, successorWithStake, stakerAddress);
        } catch {
            return (ConfirmType.NONE, 0, address(0));
        }
    }

    function requireRejectableNextNode(
        Rollup rollup,
        uint256 startNodeOffset,
        uint256 maxNodeCount,
        uint256 startStakerIndex,
        uint256 maxStakerCount
    ) external view returns (uint256, address) {
        bool outOfOrder = requireMaybeRejectable(rollup);
        if (outOfOrder) {
            return (0, address(0));
        }
        uint256 firstUnresolvedNode = rollup.firstUnresolvedNode();
        (bool found, uint256 successorWithStake, address stakerAddress) =
            findRejectableExample(
                rollup,
                firstUnresolvedNode + 1 + startNodeOffset,
                maxNodeCount,
                startStakerIndex,
                maxStakerCount
            );
        require(found, "NO_EXAMPLE");
        return (successorWithStake, stakerAddress);
    }

    function requireMaybeRejectable(Rollup rollup) private view returns (bool) {
        rollup.requireUnresolvedExists();
        INode node = rollup.getNode(rollup.firstUnresolvedNode());
        bool outOfOrder = node.prev() == rollup.latestConfirmed();
        if (outOfOrder) {
            // Verify the block's deadline has passed
            require(block.number >= node.deadlineBlock(), "BEFORE_DEADLINE");
            rollup.getNode(node.prev()).requirePastChildConfirmDeadline();

            // Verify that no staker is staked on this node
            require(node.stakerCount() == rollup.countStakedZombies(node), "HAS_STAKERS");
        }
        return outOfOrder;
    }

    function requireConfirmable(Rollup rollup) external view {
        rollup.requireUnresolvedExists();

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
            node.stakerCount() == stakerCount + rollup.countStakedZombies(node),
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
            if (node1Prev < firstUnresolvedNode || node2Prev < firstUnresolvedNode) {
                return (NodeConflict.INDETERMINATE, 0, 0);
            }
            if (node1 < node2) {
                node2 = node2Prev;
                node2Prev = rollup.getNode(node2).prev();
            } else {
                node1 = node1Prev;
                node1Prev = rollup.getNode(node1).prev();
            }
        }
        return (NodeConflict.INCOMPLETE, node1, node2);
    }

    function findRejectableExample(
        Rollup rollup,
        uint256 startNodeOffset,
        uint256 maxNodeCount,
        uint256 startStakerIndex,
        uint256 maxStakerCount
    )
        private
        view
        returns (
            bool found,
            uint256 successorWithStake,
            address stakerAddress
        )
    {
        uint256 latestNodeCreated = rollup.latestNodeCreated();
        if (startNodeOffset > latestNodeCreated) {
            return (false, 0, address(0));
        }
        uint256 max = latestNodeCreated - startNodeOffset;
        if (max > maxNodeCount) {
            max = maxNodeCount;
        }

        (address[] memory stakers, ) = getStakers(rollup, startStakerIndex, maxStakerCount);
        return
            findRejectableExampleImpl(
                rollup,
                startNodeOffset,
                rollup.latestConfirmed(),
                max,
                stakers
            );
    }

    function findRejectableExampleImpl(
        Rollup rollup,
        uint256 firstNodeToCheck,
        uint256 prev,
        uint256 max,
        address[] memory stakers
    )
        private
        view
        returns (
            bool,
            uint256,
            address
        )
    {
        uint256 stakerCount = stakers.length;
        for (uint256 i = 0; i <= max; i++) {
            uint256 nodeIndex = firstNodeToCheck + i;
            INode node = rollup.getNode(nodeIndex);
            if (node.prev() != prev) {
                continue;
            }
            for (uint256 j = 0; j < stakerCount; j++) {
                if (node.stakers(stakers[j])) {
                    return (true, nodeIndex, stakers[j]);
                }
            }
        }
        return (false, 0, address(0));
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
                if (timeSinceLastMove > challenge.currentResponderTimeLeft()) {
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
}
