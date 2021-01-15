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
import "../bridge/IBridge.sol";
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

    event SentLogs(bytes32 logsAccHash);

    // Section: Node decisions

    function rejectNextNode(uint256 successorWithStake, address stakerAddress) external;

    function confirmNextNode(
        bytes32 logAcc,
        bytes calldata sendsData,
        uint256[] calldata sendLengths
    ) external;

    // Section: Staking amount changing

    function newStake(uint256 tokenAmount) external payable;

    function withdrawStakerFunds(address payable destination) external returns (uint256);

    function reduceDeposit(uint256 maxReduction) external;

    function returnOldDeposit(address stakerAddress) external;

    function addToDeposit(address stakerAddress, uint256 tokenAmount) external payable;

    // Section: Stake movement

    function stakeOnExistingNode(
        bytes32 blockHash,
        uint256 blockNumber,
        uint256 nodeNum
    ) external;

    function stakeOnNewNode(
        bytes32 blockHash,
        uint256 blockNumber,
        uint256 nodeNum,
        bytes32[7] calldata assertionBytes32Fields,
        uint256[10] calldata assertionIntFields
    ) external;

    // Section: Challenges

    // nodeFields
    //  inboxConsistencyHash
    //  inboxDeltaHash
    //  executionHash
    function createChallenge(
        address payable[2] calldata stakers,
        uint256[2] calldata nodeNums,
        bytes32[3] calldata nodeFields,
        uint256 executionCheckTime
    ) external;

    function completeChallenge(address winningStaker, address losingStaker) external;

    // Section: Zombie cleanup

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

    function withdrawableFunds(address owner) external view returns (uint256);

    function checkMaybeRejectable() external view returns (bool);

    function checkConfirmValid() external view;

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

    function bridge() external view returns (IBridge);

    function challengeFactory() external view returns (IChallengeFactory);

    function nodeFactory() external view returns (INodeFactory);

    function currentRequiredStake() external view returns (uint256);

    function minimumAssertionPeriod() external view returns (uint256);

    function countStakedZombies(INode node) external view returns (uint256);

    function checkNoRecentStake() external view;

    function checkUnresolved() external view;
}
