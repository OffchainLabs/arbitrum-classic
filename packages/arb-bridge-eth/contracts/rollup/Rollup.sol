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

import "./StakerSet.sol";
import "./ChallengeResultReceiver.sol";

import "./IRollup.sol";
import "./IStakerSet.sol";
import "./INode.sol";
import "./RollupLib.sol";

import "../bridge/IBridge.sol";
import "../interfaces/IERC20.sol";

import "../challenge/ChallengeLib.sol";

contract Rollup is IRollup {
    struct Zombie {
        address stakerAddress;
        uint256 latestStakedNode;
    }

    uint256 public override latestConfirmed;
    uint256 public override firstUnresolvedNode;
    uint256 public override latestNodeCreated;
    mapping(uint256 => INode) public override nodes;
    uint256 public override lastStakeBlock;

    IStakerSet public override stakerSet;

    Zombie[] private zombies;

    // Rollup Config
    uint256 public override challengePeriodBlocks;
    uint256 public override arbGasSpeedLimitPerBlock;
    uint256 public override baseStake;
    address public override stakeToken;

    // Bridge is an IInbox and IOutbox
    IBridge public override bridge;
    IChallengeFactory public override challengeFactory;
    INodeFactory public override nodeFactory;
    ChallengeResultReceiver challengeResultReceiver;

    mapping(address => uint256) public override withdrawableFunds;

    constructor(
        bytes32 _machineHash,
        uint256 _challengePeriodBlocks,
        uint256 _arbGasSpeedLimitPerBlock,
        uint256 _baseStake,
        address _stakeToken,
        address _owner,
        address _bridge,
        address _challengeFactory,
        address _nodeFactory,
        bytes memory _extraConfig
    ) public {
        bridge = IBridge(_bridge);
        bridge.initialize(
            abi.encodePacked(
                uint256(_challengePeriodBlocks),
                uint256(_arbGasSpeedLimitPerBlock),
                uint256(_baseStake),
                bytes32(bytes20(_stakeToken)),
                bytes32(bytes20(_owner)),
                _extraConfig
            )
        );

        challengeFactory = IChallengeFactory(_challengeFactory);
        nodeFactory = INodeFactory(_nodeFactory);
        challengeResultReceiver = new ChallengeResultReceiver();

        bytes32 state =
            RollupLib.nodeStateHash(
                block.number, // block proposed
                0, // total gas used
                _machineHash,
                0, // inbox top
                0, // inbox count
                0, // send count
                0, // log count
                1 // inbox max count includes the initialization message
            );
        INode node =
            INode(
                nodeFactory.createNode(
                    state,
                    0, // challenge hash (not challengeable)
                    0, // confirm data
                    0, // prev node
                    0 // deadline block (not challengeable)
                )
            );
        nodes[0] = node;

        challengePeriodBlocks = _challengePeriodBlocks;
        arbGasSpeedLimitPerBlock = _arbGasSpeedLimitPerBlock;
        baseStake = _baseStake;
        stakeToken = _stakeToken;

        firstUnresolvedNode = 1;

        emit RollupCreated(_machineHash);
    }

    function checkMaybeRejectable() external view override returns (bool) {
        checkUnresolved();
        INode node = nodes[firstUnresolvedNode];
        bool outOfOrder = node.prev() == latestConfirmed;
        if (outOfOrder) {
            checkNoRecentStake();
            // Verify the block's deadline has passed
            require(block.number >= node.deadlineBlock(), "BEFORE_DEADLINE");
            // Verify that no staker is staked on this node
            require(node.stakerCount() == countStakedZombies(node), "HAS_STAKERS");
        }
        return outOfOrder;
    }

    function checkConfirmValid() external view override {
        INode node = checkConfirmValidBefore();
        checkConfirmValidAfter(node);
    }

    function rejectNextNode(uint256 successorWithStake, address stakerAddress) external override {
        checkUnresolved();

        INode node = nodes[firstUnresolvedNode];
        if (node.prev() == latestConfirmed) {
            checkNoRecentStake();
            require(successorWithStake > firstUnresolvedNode, "SUCCESSOR_TO_LOW");
            require(successorWithStake <= latestNodeCreated, "SUCCESSOR_TO_HIGH");
            require(stakerSet.isStaked(stakerAddress), "NOT_STAKED");

            // Confirm that someone is staked on some sibling node
            INode stakedSiblingNode = nodes[successorWithStake];
            // stakedSiblingNode is a child of latestConfirmed
            require(stakedSiblingNode.prev() == latestConfirmed, "BAD_SUCCESSOR");
            // staker is actually staked on stakedSiblingNode
            require(stakedSiblingNode.stakers(stakerAddress), "BAD_STAKER");

            // Verify the block's deadline has passed
            require(block.number >= node.deadlineBlock(), "BEFORE_DEADLINE");

            removeOldZombies(0);

            // Verify that no staker is staked on this node
            require(node.stakerCount() == countStakedZombies(node), "HAS_STAKERS");
        }
        destroyNode(firstUnresolvedNode);
        firstUnresolvedNode++;
    }

    function confirmNextNode(
        bytes32 logAcc,
        bytes calldata sendsData,
        uint256[] calldata sendLengths
    ) external override {
        INode node = checkConfirmValidBefore();
        removeOldZombies(0);
        checkConfirmValidAfter(node);

        bytes32 sendAcc = RollupLib.generateLastMessageHash(sendsData, sendLengths);
        require(node.confirmData() == RollupLib.confirmHash(sendAcc, logAcc), "CONFIRM_DATA");

        bridge.processOutgoingMessages(sendsData, sendLengths);

        destroyNode(latestConfirmed);

        latestConfirmed = firstUnresolvedNode;
        firstUnresolvedNode++;

        emit SentLogs(logAcc);
    }

    function checkConfirmValidBefore() private view returns (INode) {
        checkUnresolved();
        checkNoRecentStake();

        // There is at least one non-zombie staker
        require(stakerSet.count() > 0, "NO_STAKERS");

        INode node = nodes[firstUnresolvedNode];

        // Verify the block's deadline has passed
        require(node.deadlineBlock() <= block.number, "BEFORE_DEADLINE");

        // Check that prev is latest confirmed
        require(node.prev() == latestConfirmed, "INVALID_PREV");

        return node;
    }

    function checkConfirmValidAfter(INode node) private view {
        // All non-zombie stakers are staked on this node
        require(
            node.stakerCount() == stakerSet.count() + countStakedZombies(node),
            "NOT_ALL_STAKED"
        );
    }

    function newStake(uint256 tokenAmount) external payable override {
        // Verify that sender is not already a staker
        require(!stakerSet.isStaked(msg.sender), "ALREADY_STAKED");

        uint256 depositAmount = receiveStakerFunds(tokenAmount);
        require(depositAmount >= currentRequiredStake(), "NOT_ENOUGH_STAKE");

        stakerSet.create(msg.sender, latestConfirmed, depositAmount);
        lastStakeBlock = block.number;
    }

    function withdrawStakerFunds(address payable destination) external override returns (uint256) {
        uint256 amount = withdrawableFunds[destination];
        // Note: This is an unsafe external call and could be used for reentrency
        // This is safe because it occurs after all checks and effects
        sendStakerFunds(destination, amount);
        return amount;
    }

    function stakeOnExistingNode(
        bytes32 blockHash,
        uint256 blockNumber,
        uint256 nodeNum
    ) external override {
        require(blockhash(blockNumber) == blockHash, "invalid known block");
        require(nodeNum >= firstUnresolvedNode && nodeNum <= latestNodeCreated);
        INode node = nodes[nodeNum];
        stakerSet.move(msg.sender, node.prev(), nodeNum);
        node.addStaker(msg.sender);
    }

    function stakeOnNewNode(
        bytes32 blockHash,
        uint256 blockNumber,
        uint256 nodeNum,
        bytes32[7] calldata assertionBytes32Fields,
        uint256[10] calldata assertionIntFields
    ) external override {
        require(blockhash(blockNumber) == blockHash, "invalid known block");
        require(nodeNum == latestNodeCreated + 1, "NODE_NUM");
        RollupLib.Assertion memory assertion =
            RollupLib.decodeAssertion(assertionBytes32Fields, assertionIntFields);
        uint256 prev = stakerSet.moveToNew(msg.sender, nodeNum);
        INode prevNode = nodes[prev];
        // Make sure the previous state is correct against the node being built on
        require(
            RollupLib.beforeNodeStateHash(assertion) == prevNode.stateHash(),
            "PREV_STATE_HASH"
        );

        // inboxMaxCount must be greater than beforeInboxCount since we can't have read past the end of the inbox
        (uint256 inboxMaxCount, bytes32 inboxMaxAcc) = bridge.inboxInfo();
        require(
            assertion.inboxMessagesRead <= inboxMaxCount - assertion.beforeInboxCount,
            "INBOX_PAST_END"
        );

        uint256 prevDeadlineBlock = prevNode.deadlineBlock();
        uint256 timeSinceLastNode = block.number - assertion.beforeProposedBlock;
        uint256 minGasUsed = timeSinceLastNode * arbGasSpeedLimitPerBlock;
        // Verify that assertion meets the minimum Delta time requirement
        require(timeSinceLastNode >= minimumAssertionPeriod(), "TIME_DELTA");

        // Minimum size requirements: each assertion must satisfy either
        require(
            // Consumes at least all inbox messages put into L1 inbox before your prev node’s L1 blocknum
            assertion.inboxMessagesRead >=
                assertion.beforeInboxMaxCount - assertion.beforeInboxCount ||
                // Consumes ArbGas >=100% of speed limit for time since your prev node (based on difference in L1 blocknum)
                assertion.gasUsed >= minGasUsed,
            "TOO_SMALL"
        );

        // Don't allow an assertion to use above a maximum amount of gas
        require(assertion.gasUsed <= minGasUsed * 4, "TOO_LARGE");

        uint256 deadlineBlock = block.number + challengePeriodBlocks;
        if (deadlineBlock < prevDeadlineBlock) {
            deadlineBlock = prevDeadlineBlock;
        }
        uint256 executionCheckTimeBlocks = assertion.gasUsed / arbGasSpeedLimitPerBlock;
        deadlineBlock += executionCheckTimeBlocks;

        INode node =
            INode(
                nodeFactory.createNode(
                    RollupLib.nodeStateHash(assertion, inboxMaxCount),
                    RollupLib.challengeRoot(
                        assertion,
                        inboxMaxCount,
                        inboxMaxAcc,
                        executionCheckTimeBlocks
                    ),
                    RollupLib.confirmHash(assertion),
                    prev,
                    deadlineBlock
                )
            );

        latestNodeCreated++;
        nodes[latestNodeCreated] = node;

        node.addStaker(msg.sender);

        emit NodeCreated(
            latestNodeCreated,
            assertionBytes32Fields,
            assertionIntFields,
            inboxMaxCount,
            inboxMaxAcc
        );
    }

    function returnOldDeposit(address stakerAddress) external override {
        require(stakerSet.unchallengedStaker(stakerAddress), "UNCHAL_STAKER");

        (, uint256 latestStakedNode, uint256 amountStaked, ) = stakerSet.stakerInfo(stakerAddress);
        require(latestStakedNode <= latestConfirmed, "TOO_RECENT");

        stakerSet.remove(stakerAddress);
        withdrawableFunds[stakerAddress] += amountStaked;
    }

    function addToDeposit(address stakerAddress, uint256 tokenAmount) external payable override {
        require(stakerSet.unchallengedStaker(stakerAddress), "UNCHAL_STAKER");
        uint256 additionalStake = receiveStakerFunds(tokenAmount);
        stakerSet.increaseStake(stakerAddress, additionalStake);
    }

    function reduceDeposit(uint256 target) external override {
        require(stakerSet.unchallengedStaker(msg.sender), "UNCHAL_STAKER");
        uint256 currentRequired = currentRequiredStake();
        if (target > currentRequired) {
            target = currentRequired;
        }
        uint256 stakeReleased = stakerSet.reduceStakeTarget(msg.sender, target);
        withdrawableFunds[msg.sender] += stakeReleased;
    }

    // nodeFields
    //  inboxConsistencyHash
    //  inboxDeltaHash
    //  executionHash
    function createChallenge(
        address payable[2] calldata stakers,
        uint256[2] calldata nodeNums,
        bytes32[3] calldata nodeFields,
        uint256 executionCheckTime
    ) external override {
        require(nodeNums[0] < nodeNums[1], "WRONG_ORDER");
        require(nodeNums[1] <= latestNodeCreated, "NOT_PROPOSED");
        require(latestConfirmed < nodeNums[0], "ALREADY_CONFIRMED");

        INode node1 = nodes[nodeNums[0]];
        INode node2 = nodes[nodeNums[1]];

        require(node1.prev() == node2.prev(), "DIFF_PREV");

        require(stakerSet.unchallengedStaker(stakers[0]), "UNCHAL_STAKER");
        require(stakerSet.unchallengedStaker(stakers[1]), "UNCHAL_STAKER");

        require(node1.stakers(stakers[0]), "STAKER1_NOT_STAKED");
        require(node2.stakers(stakers[1]), "STAKER2_NOT_STAKED");

        require(
            node1.challengeHash() ==
                ChallengeLib.challengeRootHash(
                    nodeFields[0],
                    nodeFields[1],
                    nodeFields[2],
                    executionCheckTime
                ),
            "CHAL_HASH"
        );

        // Start a challenge between staker1 and staker2. Staker1 will defend the correctness of node1, and staker2 will challenge it.
        address challengeAddress =
            challengeFactory.createChallenge(
                address(challengeResultReceiver),
                nodeFields[0],
                nodeFields[1],
                nodeFields[2],
                executionCheckTime,
                stakers[0],
                stakers[1],
                challengePeriodBlocks
            );

        stakerSet.setChallenge(stakers[0], challengeAddress);
        stakerSet.setChallenge(stakers[1], challengeAddress);

        emit RollupChallengeStarted(challengeAddress, stakers[0], stakers[1], nodeNums[0]);
    }

    function completeChallenge(
        address challengeContract,
        address winningStaker,
        address losingStaker
    ) external override {
        require(msg.sender == address(challengeResultReceiver), "WRONG_SENDER");

        (, , uint256 winnerStake, address winnerChallenge) = stakerSet.stakerInfo(losingStaker);
        (, uint256 loserLatestStakedNode, uint256 loserStake, address loserChallenge) =
            stakerSet.stakerInfo(losingStaker);

        // Only the challenge contract can declare winners and losers
        require(winnerChallenge == challengeContract);
        require(loserChallenge == challengeContract);

        if (loserStake > winnerStake) {
            uint256 extraLoserStake = loserStake - winnerStake;
            withdrawableFunds[losingStaker] += extraLoserStake;
            loserStake -= extraLoserStake;
        }

        uint256 amountWon = loserStake / 2;
        loserStake -= amountWon;

        stakerSet.increaseStake(winningStaker, amountWon);
        stakerSet.setChallenge(winningStaker, address(0));

        // TODO: deposit extra loserStake into ArbOS

        zombies.push(Zombie(losingStaker, loserLatestStakedNode));
        stakerSet.remove(losingStaker);
    }

    function removeZombie(uint256 zombieNum, uint256 maxNodes) external override {
        require(zombieNum <= zombies.length, "NO_SUCH_ZOMBIE");
        Zombie storage zombie = zombies[zombieNum];
        uint256 latestStakedNode = zombie.latestStakedNode;
        uint256 nodesRemoved = 0;
        while (latestStakedNode > firstUnresolvedNode && nodesRemoved < maxNodes) {
            INode node = nodes[latestStakedNode];
            node.removeStaker(zombie.stakerAddress);
            latestStakedNode = node.prev();
            nodesRemoved++;
        }
        if (latestStakedNode < firstUnresolvedNode) {
            zombies[zombieNum] = zombies[zombies.length - 1];
            zombies.pop();
        } else {
            zombie.latestStakedNode = latestStakedNode;
        }
    }

    function zombieInfo(uint256 index)
        external
        view
        override
        returns (address stakerAddress, uint256 latestStakedNode)
    {
        Zombie storage zombie = zombies[index];
        return (zombie.stakerAddress, zombie.latestStakedNode);
    }

    function zombieCount() external view override returns (uint256) {
        return zombies.length;
    }

    function removeOldZombies(uint256 startIndex) public override {
        uint256 numZombies = zombies.length;
        for (uint256 i = startIndex; i < numZombies; i++) {
            Zombie storage zombie = zombies[i];
            while (zombie.latestStakedNode < firstUnresolvedNode) {
                zombies[i] = zombies[numZombies - 1];
                zombies.pop();
                numZombies--;
                if (i >= numZombies) {
                    return;
                }
                zombie = zombies[i];
            }
        }
    }

    function currentRequiredStake() public view override returns (uint256) {
        uint256 MAX_INT = 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff;
        uint256 latestConfirmedDeadline = nodes[latestConfirmed].deadlineBlock();
        if (block.number < latestConfirmedDeadline) {
            return baseStake;
        }
        uint256 latestConfirmedAge = block.number - latestConfirmedDeadline;
        uint256 challengePeriodsPassed = latestConfirmedAge / challengePeriodBlocks;
        if (challengePeriodsPassed > 255) {
            challengePeriodsPassed = 255;
        }
        uint256 multiplier = 2**challengePeriodsPassed - 1;
        if (multiplier == 0) {
            multiplier = 1;
        }

        if (multiplier > MAX_INT / baseStake) {
            return MAX_INT;
        }

        return baseStake * multiplier;
    }

    function minimumAssertionPeriod() public view override returns (uint256) {
        return challengePeriodBlocks / 10;
    }

    function countStakedZombies(INode node) public view override returns (uint256) {
        uint256 numZombies = zombies.length;
        uint256 stakedZombieCount = 0;
        for (uint256 i = 0; i < numZombies; i++) {
            Zombie storage zombie = zombies[i];
            if (node.stakers(zombie.stakerAddress)) {
                stakedZombieCount++;
            }
        }
        return stakedZombieCount;
    }

    function checkNoRecentStake() public view override {
        // No stake has been placed during the last challengePeriod blocks
        require(block.number - lastStakeBlock >= challengePeriodBlocks, "RECENT_STAKE");
    }

    function checkUnresolved() public view override {
        require(
            firstUnresolvedNode > latestConfirmed && firstUnresolvedNode <= latestNodeCreated,
            "NO_UNRESOLVED"
        );
    }

    function receiveStakerFunds(uint256 tokenAmount) private returns (uint256) {
        if (stakeToken == address(0)) {
            require(tokenAmount == 0, "BAD_STK_TYPE");
            return msg.value;
        } else {
            require(msg.value == 0, "BAD_STK_TYPE");
            require(
                IERC20(stakeToken).transferFrom(msg.sender, address(this), tokenAmount),
                "TRANSFER_FAILED"
            );
            return tokenAmount;
        }
    }

    function sendStakerFunds(address payable destination, uint256 amount) private {
        if (amount == 0) {
            return;
        }
        if (stakeToken == address(0)) {
            destination.transfer(amount);
        } else {
            require(IERC20(stakeToken).transfer(destination, amount), "TRANSFER_FAILED");
        }
    }

    function destroyNode(uint256 nodeNum) private {
        nodes[nodeNum].destroy();
        nodes[nodeNum] = INode(0);
    }
}
