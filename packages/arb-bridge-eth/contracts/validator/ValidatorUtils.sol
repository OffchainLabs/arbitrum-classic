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

    struct FindStakerConflictFrame {
        bool[] compatible;
        uint256[] ourChildren;
        uint256 latestStakedNode;
    }

    function findStakerConflict(
        Rollup rollup,
        address staker,
        uint256 startIndex,
        uint256 max
    )
        external
        view
        returns (
            address,
            uint256,
            uint256,
            bool
        )
    {
        uint256 latestConfirmed = rollup.latestConfirmed();
        uint256 node = rollup.latestStakedNode(staker);
        FindStakerConflictFrame memory frame;
        frame.latestStakedNode = node;
        frame.compatible = new bool[](rollup.latestNodeCreated() - latestConfirmed + 1);
        frame.ourChildren = new uint256[](node - latestConfirmed + 1);
        while (node > latestConfirmed) {
            frame.compatible[node - latestConfirmed] = true;
            uint256 newNode = rollup.getNode(node).prev();
            frame.ourChildren[newNode - latestConfirmed] = node;
            node = newNode;
        }
        frame.compatible[0] = true;
        (address[] memory stakers, bool hasMore) = getStakers(rollup, startIndex, max);
        for (uint256 i = 0; i < stakers.length; i++) {
            address otherStaker = stakers[i];
            if (rollup.currentChallenge(otherStaker) != address(0)) {
                continue;
            }
            node = rollup.latestStakedNode(otherStaker);
            if (frame.compatible[node - latestConfirmed]) {
                continue;
            }
            uint256 otherStakerChild;
            while (!frame.compatible[node - latestConfirmed]) {
                // We won't revisit this node for this staker,
                // and if we make it to another staker,
                // this staker's path must be compatible.
                frame.compatible[node - latestConfirmed] = true;
                otherStakerChild = node;
                node = rollup.getNode(node).prev();
            }
            if (node < frame.latestStakedNode) {
                return (
                    otherStaker,
                    frame.ourChildren[node - latestConfirmed],
                    otherStakerChild,
                    false
                );
            }
        }
        return (address(0), 0, 0, hasMore);
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
        rollup.requireUnresolvedExists();
        INode node = rollup.getNode(rollup.firstUnresolvedNode());
        bool inOrder = node.prev() == rollup.latestConfirmed();
        if (inOrder) {
            // Verify the block's deadline has passed
            require(block.number >= node.deadlineBlock(), "BEFORE_DEADLINE");
            rollup.getNode(node.prev()).requirePastChildConfirmDeadline();

            // Verify that no staker is staked on this node
            require(node.stakerCount() == rollup.countStakedZombies(node), "HAS_STAKERS");
        }
        return inOrder;
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
}
