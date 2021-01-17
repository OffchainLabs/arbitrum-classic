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

contract RollupCore {
    struct Zombie {
        address stakerAddress;
        uint256 latestStakedNode;
    }

    struct Staker {
        uint256 index;
        uint256 latestStakedNode;
        uint256 amountStaked;
        // currentChallenge is 0 if staker is not in a challenge
        address currentChallenge;
        bool isStaked;
    }

    uint256 private _latestConfirmed;
    uint256 private _firstUnresolvedNode;
    uint256 private _latestNodeCreated;
    uint256 private _lastStakeBlock;
    mapping(uint256 => INode) private _nodes;

    address payable[] private _stakerList;
    mapping(address => Staker) public _stakerMap;

    Zombie[] private _zombies;

    mapping(address => uint256) private _withdrawableFunds;

    function getNode(uint256 i) public view returns (INode) {
        return _nodes[i];
    }

    function getStakerAddress(uint256 i) public view returns (address) {
        return _stakerList[i];
    }

    function isStaked(address staker) public view returns (bool) {
        return _stakerMap[staker].isStaked;
    }

    function latestStakedNode(address staker) public view returns (uint256) {
        return _stakerMap[staker].latestStakedNode;
    }

    function currentChallenge(address staker) public view returns (address) {
        return _stakerMap[staker].currentChallenge;
    }

    function amountStaked(address staker) public view returns (uint256) {
        return _stakerMap[staker].amountStaked;
    }

    function zombieAddress(uint256 i) public view returns (address) {
        return _zombies[i].stakerAddress;
    }

    function zombieLatestStakedNode(uint256 i) public view returns (uint256) {
        return _zombies[i].latestStakedNode;
    }

    function zombieCount() public view returns (uint256) {
        return _zombies.length;
    }

    function withdrawableFunds(address owner) public view returns (uint256) {
        return _withdrawableFunds[owner];
    }

    function firstUnresolvedNode() public view returns (uint256) {
        return _firstUnresolvedNode;
    }

    function latestConfirmed() public view returns (uint256) {
        return _latestConfirmed;
    }

    function latestNodeCreated() public view returns (uint256) {
        return _latestNodeCreated;
    }

    function lastStakeBlock() public view returns (uint256) {
        return _lastStakeBlock;
    }

    function stakerCount() public view returns (uint256) {
        return _stakerList.length;
    }

    function initializeCore(INode initialNode) internal {
        _nodes[0] = initialNode;
        _firstUnresolvedNode = 1;
    }

    function nodeCreated(INode node) internal {
        _latestNodeCreated++;
        _nodes[_latestNodeCreated] = node;
    }

    function updateLatestNodeCreated(uint256 newLatestNodeCreated) internal {
        _latestNodeCreated = newLatestNodeCreated;
    }

    function rejectNextNode() internal {
        destroyNode(_firstUnresolvedNode);
        _firstUnresolvedNode++;
    }

    function confirmNextNode() internal {
        destroyNode(_latestConfirmed);
        _latestConfirmed = _firstUnresolvedNode;
        _firstUnresolvedNode++;
    }

    function createNewStake(address payable stakerAddress, uint256 depositAmount) internal {
        uint256 stakerIndex = _stakerList.length;
        _stakerList.push(stakerAddress);
        _stakerMap[stakerAddress] = Staker(
            stakerIndex,
            _latestConfirmed,
            depositAmount,
            address(0),
            true
        );
        _lastStakeBlock = block.number;
    }

    function inChallenge(address stakerAddress1, address stakerAddress2)
        internal
        view
        returns (address)
    {
        Staker storage staker1 = _stakerMap[stakerAddress1];
        Staker storage staker2 = _stakerMap[stakerAddress2];
        address challenge = staker1.currentChallenge;
        require(challenge == staker2.currentChallenge, "IN_CHAL");
        return challenge;
    }

    function clearChallenge(address stakerAddress) internal {
        Staker storage staker = _stakerMap[stakerAddress];
        staker.currentChallenge = address(0);
    }

    function increaseStakeBy(address stakerAddress, uint256 amountAdded) internal {
        Staker storage staker = _stakerMap[stakerAddress];
        staker.amountStaked += amountAdded;
    }

    function reduceStakeTo(address stakerAddress, uint256 target) internal returns (uint256) {
        Staker storage staker = _stakerMap[stakerAddress];
        uint256 current = staker.amountStaked;
        require(target <= current, "TOO_LITTLE_STAKE");
        uint256 amountWithdrawn = current - target;
        staker.amountStaked = target;
        _withdrawableFunds[stakerAddress] += amountWithdrawn;
        return amountWithdrawn;
    }

    function turnIntoZombie(address stakerAddress) internal {
        Staker storage staker = _stakerMap[stakerAddress];
        _zombies.push(Zombie(stakerAddress, staker.latestStakedNode));
        deleteStaker(staker, stakerAddress);
    }

    function zombieUpdateLatestStakedNode(uint256 i, uint256 latest) internal {
        _zombies[i].latestStakedNode = latest;
    }

    function removeZombie(uint256 i) internal {
        _zombies[i] = _zombies[_zombies.length - 1];
        _zombies.pop();
    }

    function withdrawStaker(address stakerAddress) internal {
        Staker storage staker = _stakerMap[stakerAddress];
        _withdrawableFunds[stakerAddress] += staker.amountStaked;
        deleteStaker(staker, stakerAddress);
    }

    function stakeOnNode(address stakerAddress, uint256 nodeNum) internal {
        Staker storage staker = _stakerMap[stakerAddress];
        INode node = _nodes[nodeNum];
        node.addStaker(stakerAddress);
        staker.latestStakedNode = nodeNum;
    }

    function challengeStarted(
        address staker1,
        address staker2,
        address challenge
    ) internal {
        _stakerMap[staker1].currentChallenge = challenge;
        _stakerMap[staker2].currentChallenge = challenge;
    }

    function withdrawFunds(address owner) internal returns (uint256) {
        uint256 amount = _withdrawableFunds[owner];
        _withdrawableFunds[owner] = 0;
        return amount;
    }

    function deleteStaker(Staker storage staker, address stakerAddress) private {
        uint256 stakerIndex = staker.index;
        _stakerList[stakerIndex] = _stakerList[_stakerList.length - 1];
        _stakerMap[_stakerList[stakerIndex]].index = stakerIndex;
        _stakerList.pop();
        delete _stakerMap[stakerAddress];
    }

    function destroyNode(uint256 nodeNum) private {
        _nodes[nodeNum].destroy();
        _nodes[nodeNum] = INode(0);
    }
}
