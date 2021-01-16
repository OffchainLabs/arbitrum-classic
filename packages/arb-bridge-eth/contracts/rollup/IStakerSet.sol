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

interface IStakerSet {
    // Mutating functions (only callable by owner)

    function create(
        address stakerAddress,
        uint256 latestConfirmed,
        uint256 amount
    ) external;

    function increaseStake(address stakerAddress, uint256 amount) external;

    function setChallenge(address winningStaker, address challenge) external;

    function reduceStakeTarget(address stakerAddress, uint256 target) external returns (uint256);

    function move(
        address stakerAddress,
        uint256 prevNode,
        uint256 newNode
    ) external;

    function moveToNew(address stakerAddress, uint256 newNode) external returns (uint256);

    function remove(address stakerAddress) external;

    // View functions

    function stakerList(uint256 index) external view returns (address);

    function stakerInfo(address stakerAddress)
        external
        view
        returns (
            bool isStaked,
            uint256 latestStakedNode,
            uint256 amountStaked,
            address currentChallenge
        );

    function count() external view returns (uint256);

    function isStaked(address stakerAddress) external view returns (bool);

    function unchallengedStaker(address stakerAddress) external view returns (bool);

    function getStakers(uint256 startIndex, uint256 max) external view returns (address[] memory);
}
