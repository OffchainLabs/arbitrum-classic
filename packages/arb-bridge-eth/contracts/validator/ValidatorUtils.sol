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

import "../rollup/Rollup.sol";
import "../rollup/Node.sol";

contract ValidatorUtils {
    enum ConfirmType { NONE, VALID, OUT_OF_ORDER, INVALID }

    function refundStakers(Rollup rollup, address payable[] calldata stakers) external {
        uint256 stakerCount = stakers.length;
        for (uint256 i = 0; i < stakerCount; i++) {
            try rollup.returnOldDeposit(stakers[i]) {} catch {}
        }
    }

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
        try rollup.checkUnresolved() {} catch {
            return (ConfirmType.NONE, 0, address(0));
        }
        uint256 latestConfirmed = rollup.latestConfirmed();
        uint256 firstUnresolvedNode = rollup.firstUnresolvedNode();
        Node currentUnresolved = rollup.nodes(firstUnresolvedNode);
        try currentUnresolved.checkConfirmOutOfOrder(latestConfirmed) {
            return (ConfirmType.OUT_OF_ORDER, 0, address(0));
        } catch {}
        try rollup.checkNoRecentStake() {} catch {
            return (ConfirmType.NONE, 0, address(0));
        }
        uint256 zombieCount = rollup.countStakedZombies(currentUnresolved);
        try
            currentUnresolved.checkConfirmValid(rollup.stakerCount() + zombieCount, latestConfirmed)
        {
            return (ConfirmType.VALID, 0, address(0));
        } catch {}
        try currentUnresolved.checkConfirmInvalid(zombieCount) {} catch {
            return (ConfirmType.NONE, 0, address(0));
        }
        // Node might be invalid
        (bool found, uint256 successorWithStake, address stakerAddress) =
            findRejectableExample(
                rollup,
                startNodeOffset,
                maxNodeCount,
                startStakerIndex,
                maxStakerCount
            );
        if (!found) {
            return (ConfirmType.NONE, 0, address(0));
        }
        return (ConfirmType.INVALID, successorWithStake, stakerAddress);
    }

    function checkConfirmableNextNode(Rollup rollup) external view {
        rollup.checkUnresolved();
        rollup.checkNoRecentStake();
        uint256 firstUnresolvedNode = rollup.firstUnresolvedNode();
        uint256 latestConfirmed = rollup.latestConfirmed();
        uint256 stakerCount = rollup.stakerCount();
        Node currentUnresolved = rollup.nodes(firstUnresolvedNode);
        uint256 zombieCount = rollup.countStakedZombies(currentUnresolved);
        currentUnresolved.checkConfirmValid(stakerCount + zombieCount, latestConfirmed);
    }

    function checkRejectableOutOfOrder(Rollup rollup) external view {
        rollup.checkUnresolved();
        uint256 latestConfirmed = rollup.latestConfirmed();
        uint256 firstUnresolvedNode = rollup.firstUnresolvedNode();
        Node currentUnresolved = rollup.nodes(firstUnresolvedNode);
        currentUnresolved.checkConfirmOutOfOrder(latestConfirmed);
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
        Node currentUnresolved = rollup.nodes(firstUnresolvedNode);
        currentUnresolved.checkConfirmInvalid(rollup.countStakedZombies(currentUnresolved));
        (bool found, uint256 successorWithStake, address stakerAddress) =
            findRejectableExample(
                rollup,
                startNodeOffset,
                maxNodeCount,
                startStakerIndex,
                maxStakerCount
            );
        require(found, "NO_EXAMPLE");
        return (successorWithStake, stakerAddress);
    }

    function refundableStakers(Rollup rollup) external view returns (address[] memory) {
        uint256 stakerCount = rollup.stakerCount();
        address[] memory stakers = new address[](stakerCount);
        uint256 latestConfirmed = rollup.latestConfirmed();
        uint256 index = 0;
        for (uint256 i = 0; i < stakerCount; i++) {
            address staker = rollup.stakerList(i);
            (, uint256 latestStakedNode, , , ) = rollup.stakerMap(staker);
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
            Node node = rollup.nodes(i);
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
            Node node = rollup.nodes(i);
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
        uint256 firstUnresolvedNode = rollup.firstUnresolvedNode();
        address[] memory stakers = rollup.getStakers(startStakerIndex, maxStakerCount);
        uint256 latestNodeCreated = rollup.latestNodeCreated();
        if (firstUnresolvedNode + 1 + startNodeOffset > latestNodeCreated) {
            return (false, 0, address(0));
        }
        uint256 max = latestNodeCreated - (firstUnresolvedNode + startNodeOffset);
        if (max > maxNodeCount) {
            max = maxNodeCount;
        }
        return findRejectableExampleImpl(rollup, max, startNodeOffset, stakers);
    }

    function findRejectableExampleImpl(
        Rollup rollup,
        uint256 max,
        uint256 startNodeOffset,
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
        uint256 firstUnresolvedNode = rollup.firstUnresolvedNode();
        uint256 latestConfirmed = rollup.latestConfirmed();
        uint256 stakerCount = stakers.length;

        for (uint256 i = 0; i <= max; i++) {
            uint256 nodeIndex = firstUnresolvedNode + 1 + startNodeOffset + i;
            Node node = rollup.nodes(nodeIndex);
            if (node.prev() != latestConfirmed) {
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
}
