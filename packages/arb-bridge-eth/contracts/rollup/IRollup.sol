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
import "./INodeFactory.sol";
import "../challenge/IChallengeFactory.sol";

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

    function rejectNextNode(uint256 successorWithStake, address stakerAddress) external;

    function rejectNextNodeOutOfOrder() external;

    function confirmNextNode(
        bytes32 logAcc,
        bytes calldata sendsData,
        uint256[] calldata sendLengths
    ) external;

    function newStakeOnExistingNode(
        bytes32 blockHash,
        uint256 blockNumber,
        uint256 nodeNum
    ) external payable;

    function newStakeOnNewNode(
        bytes32 blockHash,
        uint256 blockNumber,
        uint256 nodeNum,
        uint256 prev,
        bytes32[7] calldata assertionBytes32Fields,
        uint256[10] calldata assertionIntFields
    ) external payable;

    function addStakeOnExistingNode(
        bytes32 blockHash,
        uint256 blockNumber,
        uint256 nodeNum
    ) external;

    function addStakeOnNewNode(
        bytes32 blockHash,
        uint256 blockNumber,
        uint256 nodeNum,
        bytes32[7] calldata assertionBytes32Fields,
        uint256[10] calldata assertionIntFields
    ) external;

    function returnOldDeposit(address payable stakerAddress) external;

    function addToDeposit(address stakerAddress) external payable;

    function reduceDeposit(uint256 maxReduction) external;

    function createChallenge(
        address payable staker1Address,
        uint256 nodeNum1,
        address payable staker2Address,
        uint256 nodeNum2,
        bytes32 inboxConsistencyHash,
        bytes32 inboxDeltaHash,
        bytes32 executionHash,
        uint256 executionCheckTime
    ) external;

    function completeChallenge(address winningStaker, address payable losingStaker) external;

    function removeZombie(uint256 zombieNum, uint256 maxNodes) external;

    function removeOldZombies(uint256 startIndex) external;

    // Non mutating calls

    function zombieInfo(uint256 index)
        external
        view
        returns (address stakerAddress, uint256 latestStakedNode);

    function zombieCount() external view returns (uint256);

    function stakerInfo(address staker)
        external
        view
        returns (
            bool isStaked,
            uint256 latestStakedNode,
            uint256 amountStaked,
            address currentChallenge
        );

    function stakerCount() external view returns (uint256);

    function getStakers(uint256 startIndex, uint256 max) external view returns (address[] memory);

    function latestConfirmed() external view returns (uint256);

    function firstUnresolvedNode() external view returns (uint256);

    function latestNodeCreated() external view returns (uint256);

    function nodes(uint256 index) external view returns (INode);

    function lastStakeBlock() external view returns (uint256);

    function stakerList(uint256 index) external view returns (address payable);

    function challengePeriodBlocks() external view returns (uint256);

    function arbGasSpeedLimitPerBlock() external view returns (uint256);

    function baseStake() external view returns (uint256);

    function stakeToken() external view returns (address);

    function challengeFactory() external view returns (IChallengeFactory);

    function nodeFactory() external view returns (INodeFactory);

    function currentRequiredStake() external view returns (uint256);

    function minimumAssertionPeriod() external view returns (uint256);

    function countStakedZombies(INode node) external view returns (uint256);

    function checkNoRecentStake() external view;

    function checkUnresolved() external view;
}
