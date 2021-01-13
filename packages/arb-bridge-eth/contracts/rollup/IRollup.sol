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

interface IRollup {
    event RollupCreated(bytes32 machineHash);

    event NodeCreated(
        uint256 indexed nodeNum,
        bytes32[7] assertionBytes32Fields,
        uint256[10] assertionIntFields,
        uint256 inboxMaxCount,
        bytes32 inboxMaxHash
    );

    event RollupChallengeStarted(
        address indexed challengeContract,
        address asserter,
        address challenger,
        uint256 challengedNode
    );

    struct Staker {
        uint256 index;
        uint256 latestStakedNode;
        uint256 amountStaked;
        // currentChallenge is 0 if staker is not in a challenge
        address currentChallenge;
        bool isStaked;
    }

    function latestConfirmed() external view returns (uint256);

    function firstUnresolvedNode() external view returns (uint256);

    function latestNodeCreated() external view returns (uint256);

    function nodes(uint256 index) external view returns (INode);

    function lastStakeBlock() external view returns (uint256);

    function stakerList(uint256 index) external view returns (address payable);

    // function stakerMap(address staker) external view returns(Staker memory);

    // Zombie[] zombies;

    // // Rollup Config
    // uint256 public challengePeriodBlocks;
    // uint256 public arbGasSpeedLimitPerBlock;
    // uint256 public baseStake;
    // address public stakeToken;

    // IChallengeFactory public challengeFactory;
    // INodeFactory public nodeFactory;

    function completeChallenge(address winningStaker, address payable losingStaker) external;
}
