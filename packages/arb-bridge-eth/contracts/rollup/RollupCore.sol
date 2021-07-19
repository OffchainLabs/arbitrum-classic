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
import "./IRollupCore.sol";
import "./RollupLib.sol";
import "./INodeFactory.sol";
import "./RollupEventBridge.sol";
import "../bridge/interfaces/ISequencerInbox.sol";

import "@openzeppelin/contracts/math/SafeMath.sol";

contract RollupCore is IRollupCore {
    using SafeMath for uint256;

    // Stakers become Zombies after losing a challenge
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
    mapping(uint256 => bytes32) private _nodeHashes;

    address payable[] private _stakerList;
    mapping(address => Staker) public override _stakerMap;

    Zombie[] private _zombies;

    mapping(address => uint256) private _withdrawableFunds;

    /**
     * @notice Get the address of the Node contract for the given node
     * @param nodeNum Index of the node
     * @return Address of the Node contract
     */
    function getNode(uint256 nodeNum) public view override returns (INode) {
        return _nodes[nodeNum];
    }

    /**
     * @notice Get the address of the staker at the given index
     * @param stakerNum Index of the staker
     * @return Address of the staker
     */
    function getStakerAddress(uint256 stakerNum) external view override returns (address) {
        return _stakerList[stakerNum];
    }

    /**
     * @notice Check whether the given staker is staked
     * @param staker Staker address to check
     * @return True or False for whether the staker was staked
     */
    function isStaked(address staker) public view override returns (bool) {
        return _stakerMap[staker].isStaked;
    }

    /**
     * @notice Get the latest staked node of the given staker
     * @param staker Staker address to lookup
     * @return Latest node staked of the staker
     */
    function latestStakedNode(address staker) public view override returns (uint256) {
        return _stakerMap[staker].latestStakedNode;
    }

    /**
     * @notice Get the current challenge of the given staker
     * @param staker Staker address to lookup
     * @return Current challenge of the staker
     */
    function currentChallenge(address staker) public view override returns (address) {
        return _stakerMap[staker].currentChallenge;
    }

    /**
     * @notice Get the amount staked of the given staker
     * @param staker Staker address to lookup
     * @return Amount staked of the staker
     */
    function amountStaked(address staker) public view override returns (uint256) {
        return _stakerMap[staker].amountStaked;
    }

    /**
     * @notice Get the original staker address of the zombie at the given index
     * @param zombieNum Index of the zombie to lookup
     * @return Original staker address of the zombie
     */
    function zombieAddress(uint256 zombieNum) public view override returns (address) {
        return _zombies[zombieNum].stakerAddress;
    }

    /**
     * @notice Get Latest node that the given zombie at the given index is staked on
     * @param zombieNum Index of the zombie to lookup
     * @return Latest node that the given zombie is staked on
     */
    function zombieLatestStakedNode(uint256 zombieNum) public view override returns (uint256) {
        return _zombies[zombieNum].latestStakedNode;
    }

    /// @return Current number of un-removed zombies
    function zombieCount() public view override returns (uint256) {
        return _zombies.length;
    }

    function isZombie(address staker) public view override returns (bool) {
        for (uint256 i = 0; i < _zombies.length; i++) {
            if (staker == _zombies[i].stakerAddress) {
                return true;
            }
        }
        return false;
    }

    /**
     * @notice Get the amount of funds withdrawable by the given address
     * @param owner Address to check the funds of
     * @return Amount of funds withdrawable by owner
     */
    function withdrawableFunds(address owner) external view override returns (uint256) {
        return _withdrawableFunds[owner];
    }

    /**
     * @return Index of the first unresolved node
     * @dev If all nodes have been resolved, this will be latestNodeCreated + 1
     */
    function firstUnresolvedNode() public view override returns (uint256) {
        return _firstUnresolvedNode;
    }

    /// @return Index of the latest confirmed node
    function latestConfirmed() public view override returns (uint256) {
        return _latestConfirmed;
    }

    /// @return Index of the latest rollup node created
    function latestNodeCreated() public view override returns (uint256) {
        return _latestNodeCreated;
    }

    /// @return Ethereum block that the most recent stake was created
    function lastStakeBlock() external view override returns (uint256) {
        return _lastStakeBlock;
    }

    /// @return Number of active stakers currently staked
    function stakerCount() public view override returns (uint256) {
        return _stakerList.length;
    }

    /**
     * @notice Initialize the core with an initial node
     * @param initialNode Initial node to start the chain with
     */
    function initializeCore(INode initialNode) internal {
        _nodes[0] = initialNode;
        _firstUnresolvedNode = 1;
    }

    /**
     * @notice React to a new node being created by storing it an incrementing the latest node counter
     * @param node Node that was newly created
     * @param nodeHash The hash of said node
     */
    function nodeCreated(INode node, bytes32 nodeHash) internal {
        _latestNodeCreated++;
        _nodes[_latestNodeCreated] = node;
        _nodeHashes[_latestNodeCreated] = nodeHash;
    }

    /// @return Node hash as of this node number
    function getNodeHash(uint256 index) public view override returns (bytes32) {
        return _nodeHashes[index];
    }

    /// @notice Reject the next unresolved node
    function _rejectNextNode() internal {
        destroyNode(_firstUnresolvedNode);
        _firstUnresolvedNode++;
    }

    /// @notice Confirm the next unresolved node
    function confirmNextNode(
        bytes32 beforeSendAcc,
        bytes calldata sendsData,
        uint256[] calldata sendLengths,
        uint256 afterSendCount,
        bytes32 afterLogAcc,
        uint256 afterLogCount,
        IOutbox outbox,
        RollupEventBridge rollupEventBridge
    ) internal {
        confirmNode(
            _firstUnresolvedNode,
            beforeSendAcc,
            sendsData,
            sendLengths,
            afterSendCount,
            afterLogAcc,
            afterLogCount,
            outbox,
            rollupEventBridge
        );
    }

    function confirmNode(
        uint256 nodeNum,
        bytes32 beforeSendAcc,
        bytes calldata sendsData,
        uint256[] calldata sendLengths,
        uint256 afterSendCount,
        bytes32 afterLogAcc,
        uint256 afterLogCount,
        IOutbox outbox,
        RollupEventBridge rollupEventBridge
    ) internal {
        bytes32 afterSendAcc = RollupLib.feedAccumulator(sendsData, sendLengths, beforeSendAcc);

        INode node = getNode(nodeNum);
        // Authenticate data against node's confirm data pre-image
        require(
            node.confirmData() ==
                RollupLib.confirmHash(
                    beforeSendAcc,
                    afterSendAcc,
                    afterLogAcc,
                    afterSendCount,
                    afterLogCount
                ),
            "CONFIRM_DATA"
        );

        // trusted external call to outbox
        outbox.processOutgoingMessages(sendsData, sendLengths);

        destroyNode(_latestConfirmed);
        _latestConfirmed = nodeNum;
        _firstUnresolvedNode = nodeNum + 1;

        rollupEventBridge.nodeConfirmed(nodeNum);
        emit NodeConfirmed(nodeNum, afterSendAcc, afterSendCount, afterLogAcc, afterLogCount);
    }

    /**
     * @notice Create a new stake at latest confirmed node
     * @param stakerAddress Address of the new staker
     * @param depositAmount Stake amount of the new staker
     */
    function createNewStake(address payable stakerAddress, uint256 depositAmount) internal {
        uint256 stakerIndex = _stakerList.length;
        _stakerList.push(stakerAddress);
        _stakerMap[stakerAddress] = Staker(
            stakerIndex,
            _latestConfirmed,
            depositAmount,
            address(0), // new staker is not in challenge
            true
        );
        _lastStakeBlock = block.number;
    }

    /**
     * @notice Check to see whether the two stakers are in the same challenge
     * @param stakerAddress1 Address of the first staker
     * @param stakerAddress2 Address of the second staker
     * @return Address of the challenge that the two stakers are in
     */
    function inChallenge(address stakerAddress1, address stakerAddress2)
        internal
        view
        returns (address)
    {
        Staker storage staker1 = _stakerMap[stakerAddress1];
        Staker storage staker2 = _stakerMap[stakerAddress2];
        address challenge = staker1.currentChallenge;
        require(challenge != address(0), "NO_CHAL");
        require(challenge == staker2.currentChallenge, "DIFF_IN_CHAL");
        return challenge;
    }

    /**
     * @notice Make the given staker as not being in a challenge
     * @param stakerAddress Address of the staker to remove from a challenge
     */
    function clearChallenge(address stakerAddress) internal {
        Staker storage staker = _stakerMap[stakerAddress];
        staker.currentChallenge = address(0);
    }

    /**
     * @notice Mark both the given stakers as engaged in the challenge
     * @param staker1 Address of the first staker
     * @param staker2 Address of the second staker
     * @param challenge Address of the challenge both stakers are now in
     */
    function challengeStarted(
        address staker1,
        address staker2,
        address challenge
    ) internal {
        _stakerMap[staker1].currentChallenge = challenge;
        _stakerMap[staker2].currentChallenge = challenge;
    }

    /**
     * @notice Add to the stake of the given staker by the given amount
     * @param stakerAddress Address of the staker to increase the stake of
     * @param amountAdded Amount of stake to add to the staker
     */
    function increaseStakeBy(address stakerAddress, uint256 amountAdded) internal {
        Staker storage staker = _stakerMap[stakerAddress];
        staker.amountStaked = staker.amountStaked.add(amountAdded);
    }

    /**
     * @notice Reduce the stake of the given staker to the given target
     * @param stakerAddress Address of the staker to reduce the stake of
     * @param target Amount of stake to leave with the staker
     * @return Amount of value released from the stake
     */
    function reduceStakeTo(address stakerAddress, uint256 target) internal returns (uint256) {
        Staker storage staker = _stakerMap[stakerAddress];
        uint256 current = staker.amountStaked;
        require(target <= current, "TOO_LITTLE_STAKE");
        uint256 amountWithdrawn = current.sub(target);
        staker.amountStaked = target;
        _withdrawableFunds[stakerAddress] = _withdrawableFunds[stakerAddress].add(amountWithdrawn);
        return amountWithdrawn;
    }

    /**
     * @notice Remove the given staker and turn them into a zombie
     * @param stakerAddress Address of the staker to remove
     */
    function turnIntoZombie(address stakerAddress) internal {
        Staker storage staker = _stakerMap[stakerAddress];
        _zombies.push(Zombie(stakerAddress, staker.latestStakedNode));
        deleteStaker(stakerAddress);
    }

    /**
     * @notice Update the latest staked node of the zombie at the given index
     * @param zombieNum Index of the zombie to move
     * @param latest New latest node the zombie is staked on
     */
    function zombieUpdateLatestStakedNode(uint256 zombieNum, uint256 latest) internal {
        _zombies[zombieNum].latestStakedNode = latest;
    }

    /**
     * @notice Remove the zombie at the given index
     * @param zombieNum Index of the zombie to remove
     */
    function removeZombie(uint256 zombieNum) internal {
        _zombies[zombieNum] = _zombies[_zombies.length - 1];
        _zombies.pop();
    }

    /**
     * @notice Remove the given staker and return their stake
     * @param stakerAddress Address of the staker withdrawing their stake
     */
    function withdrawStaker(address stakerAddress) internal {
        Staker storage staker = _stakerMap[stakerAddress];
        _withdrawableFunds[stakerAddress] = _withdrawableFunds[stakerAddress].add(
            staker.amountStaked
        );
        deleteStaker(stakerAddress);
    }

    /**
     * @notice Advance the given staker to the given node
     * @param stakerAddress Address of the staker adding their stake
     * @param nodeNum Index of the node to stake on
     */
    function stakeOnNode(
        address stakerAddress,
        uint256 nodeNum,
        uint256 confirmPeriodBlocks
    ) internal {
        Staker storage staker = _stakerMap[stakerAddress];
        INode node = _nodes[nodeNum];
        uint256 newStakerCount = node.addStaker(stakerAddress);
        staker.latestStakedNode = nodeNum;
        if (newStakerCount == 1) {
            INode parent = _nodes[node.prev()];
            parent.newChildConfirmDeadline(block.number.add(confirmPeriodBlocks));
        }
    }

    /**
     * @notice Clear the withdrawable funds for the given address
     * @param owner Address of the account to remove funds from
     * @return Amount of funds removed from account
     */
    function withdrawFunds(address owner) internal returns (uint256) {
        uint256 amount = _withdrawableFunds[owner];
        _withdrawableFunds[owner] = 0;
        return amount;
    }

    /**
     * @notice Increase the withdrawable funds for the given address
     * @param owner Address of the account to add withdrawable funds to
     */
    function increaseWithdrawableFunds(address owner, uint256 amount) internal {
        _withdrawableFunds[owner] = _withdrawableFunds[owner].add(amount);
    }

    /**
     * @notice Remove the given staker
     * @param stakerAddress Address of the staker to remove
     */
    function deleteStaker(address stakerAddress) private {
        Staker storage staker = _stakerMap[stakerAddress];
        uint256 stakerIndex = staker.index;
        _stakerList[stakerIndex] = _stakerList[_stakerList.length - 1];
        _stakerMap[_stakerList[stakerIndex]].index = stakerIndex;
        _stakerList.pop();
        delete _stakerMap[stakerAddress];
    }

    /**
     * @notice Destroy the given node and clear out its address
     * @param nodeNum Index of the node to remove
     */
    function destroyNode(uint256 nodeNum) internal {
        _nodes[nodeNum].destroy();
        _nodes[nodeNum] = INode(0);
    }

    function nodeDeadline(
        uint256 arbGasSpeedLimitPerBlock,
        uint256 gasUsed,
        uint256 confirmPeriodBlocks,
        INode prevNode
    ) internal view returns (uint256 deadlineBlock) {
        // Set deadline rounding up to the nearest block
        uint256 checkTime =
            gasUsed.add(arbGasSpeedLimitPerBlock.sub(1)).div(arbGasSpeedLimitPerBlock);

        deadlineBlock = max(block.number.add(confirmPeriodBlocks), prevNode.deadlineBlock()).add(
            checkTime
        );

        uint256 olderSibling = prevNode.latestChildNumber();
        if (olderSibling != 0) {
            deadlineBlock = max(deadlineBlock, getNode(olderSibling).deadlineBlock());
        }
        return deadlineBlock;
    }

    function max(uint256 a, uint256 b) internal pure returns (uint256) {
        return a > b ? a : b;
    }

    struct StakeOnNewNodeFrame {
        uint256 currentInboxSize;
        INode node;
        bytes32 executionHash;
        INode prevNode;
        bytes32 lastHash;
        bool hasSibling;
        uint256 deadlineBlock;
        uint256 gasUsed;
        uint256 sequencerBatchEnd;
        bytes32 sequencerBatchAcc;
    }

    struct CreateNodeDataFrame {
        uint256 prevNode;
        uint256 confirmPeriodBlocks;
        uint256 arbGasSpeedLimitPerBlock;
        ISequencerInbox sequencerInbox;
        RollupEventBridge rollupEventBridge;
        INodeFactory nodeFactory;
    }

    // TODO: Configure this value based on the cost of sends and add test
    uint8 internal constant MAX_SEND_COUNT = 100;

    function createNewNode(
        RollupLib.Assertion memory assertion,
        bytes32[3][2] calldata assertionBytes32Fields,
        uint256[4][2] calldata assertionIntFields,
        bytes calldata sequencerBatchProof,
        CreateNodeDataFrame memory inputDataFrame,
        bytes32 expectedNodeHash
    ) internal returns (bytes32 newNodeHash) {
        StakeOnNewNodeFrame memory memoryFrame;
        {
            // validate data
            memoryFrame.gasUsed = RollupLib.assertionGasUsed(assertion);
            memoryFrame.prevNode = getNode(inputDataFrame.prevNode);
            memoryFrame.currentInboxSize = inputDataFrame.sequencerInbox.messageCount();

            // Make sure the previous state is correct against the node being built on
            require(
                RollupLib.stateHash(assertion.beforeState) == memoryFrame.prevNode.stateHash(),
                "PREV_STATE_HASH"
            );

            // Ensure that the assertion doesn't read past the end of the current inbox
            require(
                assertion.afterState.inboxCount <= memoryFrame.currentInboxSize,
                "INBOX_PAST_END"
            );
            // Insure inbox tip after assertion is included in a sequencer-inbox batch and return inbox acc; this gives replay protection against the state of the inbox
            (memoryFrame.sequencerBatchEnd, memoryFrame.sequencerBatchAcc) = inputDataFrame
                .sequencerInbox
                .proveInboxContainsMessage(sequencerBatchProof, assertion.afterState.inboxCount);
        }

        {
            memoryFrame.executionHash = RollupLib.executionHash(assertion);

            memoryFrame.deadlineBlock = nodeDeadline(
                inputDataFrame.arbGasSpeedLimitPerBlock,
                memoryFrame.gasUsed,
                inputDataFrame.confirmPeriodBlocks,
                memoryFrame.prevNode
            );

            memoryFrame.hasSibling = memoryFrame.prevNode.latestChildNumber() > 0;
            // here we don't use ternacy operator to remain compatible with slither
            if (memoryFrame.hasSibling) {
                memoryFrame.lastHash = getNodeHash(memoryFrame.prevNode.latestChildNumber());
            } else {
                memoryFrame.lastHash = getNodeHash(inputDataFrame.prevNode);
            }

            memoryFrame.node = INode(
                inputDataFrame.nodeFactory.createNode(
                    RollupLib.stateHash(assertion.afterState),
                    RollupLib.challengeRoot(assertion, memoryFrame.executionHash, block.number),
                    RollupLib.confirmHash(assertion),
                    inputDataFrame.prevNode,
                    memoryFrame.deadlineBlock
                )
            );
        }

        {
            uint256 nodeNum = latestNodeCreated() + 1;
            memoryFrame.prevNode.childCreated(nodeNum);

            newNodeHash = RollupLib.nodeHash(
                memoryFrame.hasSibling,
                memoryFrame.lastHash,
                memoryFrame.executionHash,
                memoryFrame.sequencerBatchAcc
            );
            require(newNodeHash == expectedNodeHash, "UNEXPECTED_NODE_HASH");

            nodeCreated(memoryFrame.node, newNodeHash);
            inputDataFrame.rollupEventBridge.nodeCreated(
                nodeNum,
                inputDataFrame.prevNode,
                memoryFrame.deadlineBlock,
                msg.sender
            );
        }

        emit NodeCreated(
            latestNodeCreated(),
            getNodeHash(inputDataFrame.prevNode),
            newNodeHash,
            memoryFrame.executionHash,
            memoryFrame.currentInboxSize,
            memoryFrame.sequencerBatchEnd,
            memoryFrame.sequencerBatchAcc,
            assertionBytes32Fields,
            assertionIntFields
        );

        return newNodeHash;
    }
}
