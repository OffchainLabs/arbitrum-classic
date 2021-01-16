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

import "./IStakerSet.sol";

contract StakerSet is IStakerSet {
    struct Staker {
        uint256 index;
        uint256 latestStakedNode;
        uint256 amountStaked;
        // currentChallenge is 0 if staker is not in a challenge
        address currentChallenge;
        bool isStaked;
    }

    address public rollup;

    address[] public override stakerList;
    mapping(address => Staker) public stakerMap;

    modifier onlyRollup {
        require(msg.sender == rollup, "ONLY_ROLLUP");
        _;
    }

    function create(
        address stakerAddress,
        uint256 latestConfirmed,
        uint256 amount
    ) external override onlyRollup {
        require(msg.sender == rollup, "ONLY_ROLLUP");

        uint256 stakerIndex = stakerList.length;
        stakerList.push(stakerAddress);
        stakerMap[stakerAddress] = Staker(stakerIndex, latestConfirmed, amount, address(0), true);
    }

    function increaseStake(address stakerAddress, uint256 amount) external override onlyRollup {
        Staker storage staker = getStaker(stakerAddress);
        staker.amountStaked += amount;
    }

    function setChallenge(address winningStaker, address challenge) external override onlyRollup {
        Staker storage staker = getStaker(winningStaker);
        staker.currentChallenge = challenge;
    }

    function reduceStakeTarget(address stakerAddress, uint256 target)
        external
        override
        onlyRollup
        returns (uint256)
    {
        Staker storage staker = getStaker(stakerAddress);
        uint256 balance = staker.amountStaked;
        require(target <= balance, "TOO_LITTLE");
        uint256 amountWithdrawn = balance - target;
        staker.amountStaked = target;
        return amountWithdrawn;
    }

    function move(
        address stakerAddress,
        uint256 prevNode,
        uint256 newNode
    ) external override onlyRollup {
        Staker storage staker = getStaker(stakerAddress);
        require(staker.latestStakedNode == prevNode, "NOT_STAKED_PREV");
        staker.latestStakedNode = newNode;
    }

    function moveToNew(address stakerAddress, uint256 newNode)
        external
        override
        onlyRollup
        returns (uint256)
    {
        Staker storage staker = getStaker(stakerAddress);
        uint256 prev = staker.latestStakedNode;
        staker.latestStakedNode = newNode;
        return prev;
    }

    function remove(address stakerAddress) external override onlyRollup {
        Staker storage staker = stakerMap[stakerAddress];
        uint256 stakerIndex = staker.index;
        stakerList[stakerIndex] = stakerList[stakerList.length - 1];
        stakerMap[stakerList[stakerIndex]].index = stakerIndex;
        stakerList.pop();
        delete stakerMap[stakerAddress];
    }

    function stakerInfo(address stakerAddress)
        external
        view
        override
        returns (
            bool isStaked,
            uint256 latestStakedNode,
            uint256 amountStaked,
            address currentChallenge
        )
    {
        Staker storage staker = stakerMap[stakerAddress];
        return (
            staker.isStaked,
            staker.latestStakedNode,
            staker.amountStaked,
            staker.currentChallenge
        );
    }

    function count() external view override returns (uint256) {
        return stakerList.length;
    }

    function isStaked(address stakerAddress) external view override returns (bool) {
        return stakerMap[stakerAddress].isStaked;
    }

    function unchallengedStaker(address stakerAddress) external view override returns (bool) {
        Staker storage staker = stakerMap[stakerAddress];
        return staker.isStaked && staker.currentChallenge == address(0);
    }

    function getStakers(uint256 startIndex, uint256 max)
        external
        view
        override
        returns (address[] memory)
    {
        uint256 maxStakers = stakerList.length;
        if (startIndex + max < maxStakers) {
            maxStakers = startIndex + max;
        }

        address[] memory stakers = new address[](maxStakers);
        for (uint256 i = 0; i < maxStakers; i++) {
            stakers[i] = stakerList[startIndex + i];
        }
        return stakers;
    }

    function getStaker(address stakerAddress) private view returns (Staker storage) {
        Staker storage staker = stakerMap[stakerAddress];
        require(staker.isStaked, "NOT_STAKED");
        return staker;
    }
}
