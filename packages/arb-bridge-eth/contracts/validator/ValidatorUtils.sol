// SPDX-License-Identifier: Apache-2.0

// /*
//  * Copyright 2021, Offchain Labs, Inc.
//  *
//  * Licensed under the Apache License, Version 2.0 (the "License");
//  * you may not use this file except in compliance with the License.
//  * You may obtain a copy of the License at
//  *
//  *    http://www.apache.org/licenses/LICENSE-2.0
//  *
//  * Unless required by applicable law or agreed to in writing, software
//  * distributed under the License is distributed on an "AS IS" BASIS,
//  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  * See the License for the specific language governing permissions and
//  * limitations under the License.
//  */

pragma solidity ^0.6.11;

pragma experimental ABIEncoderV2;

import "../rollup/Rollup.sol";

contract ValidatorUtils {
    enum ConfirmType { NONE, VALID, INVALID }

    enum NodeConflict { NONE, FOUND, INDETERMINATE, INCOMPLETE }

    function getConfig(Rollup rollup)
        external
        view
        returns (
            uint256 challengePeriodBlocks,
            uint256 arbGasSpeedLimitPerBlock,
            uint256 baseStake,
            address stakeToken
        )
    {
        challengePeriodBlocks = rollup.challengePeriodBlocks();
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
        try rollup.checkConfirmValidBefore() returns (INode node) {
            try rollup.checkConfirmValidAfter(node) {
                return (ConfirmType.VALID, 0, address(0));
            } catch {}
        } catch {}

        try
            ValidatorUtils(address(this)).checkRejectableNextNode(
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

    function checkRejectableNextNode(
        Rollup rollup,
        uint256 startNodeOffset,
        uint256 maxNodeCount,
        uint256 startStakerIndex,
        uint256 maxStakerCount
    ) external view returns (uint256, address) {
        rollup.checkUnresolved();
        rollup.checkNoRecentStake();
        uint256 firstUnresolvedNode = rollup.firstUnresolvedNode();
        bool outOfOrder = checkMaybeRejectable(rollup);
        if (outOfOrder) {
            return (0, address(0));
        }
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

    function checkMaybeRejectable(Rollup rollup) private view returns (bool) {
        rollup.checkUnresolved();
        INode node = rollup.getNode(rollup.firstUnresolvedNode());
        bool outOfOrder = node.prev() == rollup.latestConfirmed();
        if (outOfOrder) {
            rollup.checkNoRecentStake();
            // Verify the block's deadline has passed
            require(block.number >= node.deadlineBlock(), "BEFORE_DEADLINE");
            // Verify that no staker is staked on this node
            require(node.stakerCount() == rollup.countStakedZombies(node), "HAS_STAKERS");
        }
        return outOfOrder;
    }

    function refundableStakers(Rollup rollup) external view returns (address[] memory) {
        uint256 stakerCount = rollup.stakerCount();
        address[] memory stakers = new address[](stakerCount);
        uint256 latestConfirmed = rollup.latestConfirmed();
        uint256 index = 0;
        for (uint256 i = 0; i < stakerCount; i++) {
            address staker = rollup.getStakerAddress(i);
            uint256 latestStakedNode = rollup.latestStakedNode(staker);
            if (latestStakedNode <= latestConfirmed) {
                stakers[index] = staker;
                index++;
            }
        }
        assembly {
            mstore(stakers, index)
        }
        return stakers;
    }

    function successorNodes(Rollup rollup, uint256 nodeNum)
        external
        view
        returns (uint256[] memory)
    {
        uint256[] memory nodes = new uint256[](100000);
        uint256 index = 0;
        for (uint256 i = nodeNum + 1; i <= rollup.latestNodeCreated(); i++) {
            INode node = rollup.getNode(i);
            if (node.prev() == nodeNum) {
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

        return
            findRejectableExampleImpl(
                rollup,
                startNodeOffset,
                rollup.latestConfirmed(),
                max,
                getStakers(rollup, startStakerIndex, maxStakerCount)
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
    ) public view returns (address[] memory) {
        uint256 maxStakers = rollup.stakerCount();
        if (startIndex + max < maxStakers) {
            maxStakers = startIndex + max;
        }

        address[] memory stakers = new address[](maxStakers);
        for (uint256 i = 0; i < maxStakers; i++) {
            stakers[i] = rollup.getStakerAddress(startIndex + i);
        }
        return stakers;
    }
}
