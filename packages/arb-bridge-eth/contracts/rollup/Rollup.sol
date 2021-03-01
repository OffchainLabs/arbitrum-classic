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

import "./RollupCore.sol";
import "@openzeppelin/contracts/utils/Pausable.sol";
import "@openzeppelin/contracts/proxy/ProxyAdmin.sol";
import "./RollupEventBridge.sol";

import "./IRollup.sol";
import "./INode.sol";
import "./INodeFactory.sol";
import "../challenge/IChallengeFactory.sol";
import "../bridge/interfaces/IBridge.sol";
import "../bridge/interfaces/IOutbox.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import "../bridge/Messages.sol";
import "./RollupLib.sol";
import "../libraries/Cloneable.sol";

contract Rollup is Cloneable, RollupCore, Pausable, IRollup {
    // TODO: Configure this value based on the cost of sends
    uint8 internal constant MAX_SEND_COUNT = 100;

    // A little over 15 minutes
    uint256 public constant minimumAssertionPeriod = 75;

    // Rollup Config
    uint256 public confirmPeriodBlocks;
    uint256 public extraChallengeTimeBlocks;
    uint256 public arbGasSpeedLimitPerBlock;
    uint256 public baseStake;
    address public stakeToken;

    // Bridge is an IInbox and IOutbox
    IBridge public bridge;
    IOutbox public outbox;
    RollupEventBridge public rollupEventBridge;
    IChallengeFactory public challengeFactory;
    INodeFactory public nodeFactory;
    address public owner;
    ProxyAdmin public admin;

    uint256 latestNodeToTruncateTo;
    uint256 nextStakerToTruncate;
    bool truncating;

    modifier onlyOwner {
        require(msg.sender == owner, "ONLY_OWNER");
        _;
    }

    // connectedContracts = [admin, bridge, outbox, rollupEventBridge, challengeFactory, nodeFactory]
    function initialize(
        bytes32 _machineHash,
        uint256 _confirmPeriodBlocks,
        uint256 _extraChallengeTimeBlocks,
        uint256 _arbGasSpeedLimitPerBlock,
        uint256 _baseStake,
        address _stakeToken,
        address _owner,
        bytes calldata _extraConfig,
        address[6] calldata connectedContracts
    ) external override {
        require(confirmPeriodBlocks == 0, "ALREADY_INIT");
        require(_confirmPeriodBlocks != 0, "BAD_CONF_PERIOD");

        bridge = IBridge(connectedContracts[1]);
        outbox = IOutbox(connectedContracts[2]);
        bridge.setOutbox(connectedContracts[2], true);
        rollupEventBridge = RollupEventBridge(connectedContracts[3]);
        bridge.setInbox(connectedContracts[3], true);

        rollupEventBridge.rollupInitialized(
            _confirmPeriodBlocks,
            _extraChallengeTimeBlocks,
            _arbGasSpeedLimitPerBlock,
            _baseStake,
            _stakeToken,
            _owner,
            _extraConfig
        );

        challengeFactory = IChallengeFactory(connectedContracts[4]);
        nodeFactory = INodeFactory(connectedContracts[5]);

        INode node = createInitialNode(_machineHash);
        initializeCore(node);

        confirmPeriodBlocks = _confirmPeriodBlocks;
        extraChallengeTimeBlocks = _extraChallengeTimeBlocks;
        arbGasSpeedLimitPerBlock = _arbGasSpeedLimitPerBlock;
        baseStake = _baseStake;
        stakeToken = _stakeToken;
        owner = _owner;
        admin = ProxyAdmin(connectedContracts[0]);

        emit RollupCreated(_machineHash);
    }

    function createInitialNode(bytes32 _machineHash) private returns (INode) {
        bytes32 state =
            RollupLib.nodeStateHash(
                block.number, // block proposed
                0, // total gas used
                _machineHash,
                0, // inbox count
                0, // send count
                0, // log count
                1 // Initialization message already in inbox
            );
        return
            INode(
                nodeFactory.createNode(
                    state,
                    0, // challenge hash (not challengeable)
                    0, // confirm data
                    0, // prev node
                    block.number // deadline block (not challengeable)
                )
            );
    }

    /**
     * @notice Add a contract authorized to put messages into this rollup's inbox
     * @param _outbox Outbox contract to add
     */
    function setOutbox(IOutbox _outbox) external onlyOwner {
        outbox = _outbox;
        bridge.setOutbox(address(_outbox), true);
    }

    /**
     * @notice Disable an old outbox from interacting with the bridge
     * @param _outbox Outbox contract to remove
     */
    function removeOldOutbox(address _outbox) external onlyOwner {
        require(_outbox != address(outbox), "CUR_OUTBOX");
        bridge.setOutbox(_outbox, false);
    }

    /**
     * @notice Enable or disable an inbox contract
     * @param _inbox Inbox contract to add or remove
     * @param _enabled New status of inbox
     */
    function setInbox(address _inbox, bool _enabled) external onlyOwner {
        bridge.setInbox(address(_inbox), _enabled);
    }

    /**
     * @notice Switch over to a new implementation of the rollup
     * @param _newRollup New implementation contract
     */
    function upgradeImplementation(address _newRollup) external onlyOwner {
        address currentAddress = address(this);
        admin.upgrade(TransparentUpgradeableProxy(payable(currentAddress)), _newRollup);
    }

    /**
     * @notice Switch over to a new implementation of the rollup
     * @param _newRollup New implementation contract
     * @param _data Data to call the new rollup implementation with
     */
    function upgradeImplementationAndCall(address _newRollup, bytes calldata _data)
        external
        onlyOwner
    {
        address currentAddress = address(this);
        admin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(currentAddress)),
            _newRollup,
            _data
        );
    }

    /**
     * @notice Pause interaction with the rollup contract
     */
    function pause() external onlyOwner {
        _pause();
    }

    /**
     * @notice Resume interaction with the rollup contract
     */
    function resume() external onlyOwner {
        require(!truncating, "STILL_TRUNCATING");
        _unpause();
    }

    /**
     * @notice Begin the process of trunacting the chain back to the given node
     * @dev maxItems is used to make sure this doesn't exceed the max gas cost
     * @param newLatestNodeCreated Index that we want to be the latest unresolved node
     * @param maxItems Maximum number of items to eliminate to eliminate
     */
    function beginTruncatingNodes(uint256 newLatestNodeCreated, uint256 maxItems)
        external
        onlyOwner
        whenPaused
    {
        require(!truncating, "ALREADY_TRUNCATING");
        require(newLatestNodeCreated < latestNodeCreated(), "TOO_NEW");
        require(newLatestNodeCreated >= firstUnresolvedNode() - 1, "TOO_OLD");
        latestNodeToTruncateTo = newLatestNodeCreated;
        truncating = true;
        continueTruncatingNodes(maxItems);
    }

    /**
     * @notice Continue the process of trunacting the chain back to the given node
     * @dev maxItems is used to make sure this doesn't exceed the max gas cost
     * @param maxItems Maximum number of items to eliminate to eliminate
     */
    function continueTruncatingNodes(uint256 maxItems) public onlyOwner whenPaused {
        require(truncating, "NOT_TRUNCATING");
        uint256 target = latestNodeToTruncateTo;

        uint256 stakerIndex = nextStakerToTruncate;
        uint256 stakers = stakerCount();
        while (maxItems > 0 && stakerIndex < stakers) {
            address stakerAddress = getStakerAddress(stakerIndex);
            uint256 latestStakedNode = latestStakedNode(stakerAddress);
            while (maxItems > 0 && latestStakedNode > target) {
                INode node = getNode(latestStakedNode);
                latestStakedNode = node.prev();
                maxItems--;
            }
            stakerUpdateLatestStakedNode(stakerAddress, latestStakedNode);

            if (latestStakedNode > target) {
                nextStakerToTruncate = stakerIndex;
                return;
            }
            stakerIndex++;
        }
        nextStakerToTruncate = stakerIndex;

        uint256 latest;
        for (latest = latestNodeCreated(); maxItems > 0 && latest > target; latest--) {
            INode node = getNode(latest);
            node.destroy();
            maxItems--;
        }
        updateLatestNodeCreated(latest);
        if (latest == target) {
            latestNodeToTruncateTo = 0;
            nextStakerToTruncate = 0;
            truncating = false;
        }
    }

    /**
     * @notice Reject the next unresolved node
     * @param successorWithStake Example sibling node
     * @param stakerAddress Example staker staked on sibling
     */
    function rejectNextNode(uint256 successorWithStake, address stakerAddress)
        external
        whenNotPaused
    {
        requireUnresolvedExists();
        uint256 latest = latestConfirmed();
        uint256 firstUnresolved = firstUnresolvedNode();
        INode node = getNode(firstUnresolved);
        if (node.prev() == latest) {
            requireUnresolved(successorWithStake);
            require(isStaked(stakerAddress), "NOT_STAKED");

            // Confirm that someone is staked on some sibling node
            getNode(successorWithStake).requireRejectExample(latest, stakerAddress);

            // Verify the block's deadline has passed
            node.requirePastDeadline();

            getNode(latest).requirePastChildConfirmDeadline();

            removeOldZombies(0);

            // Verify that no staker is staked on this node
            require(node.stakerCount() == countStakedZombies(node), "HAS_STAKERS");
        }
        rejectNextNode();
        rollupEventBridge.nodeRejected(firstUnresolved);
    }

    /**
     * @notice Confirm the next unresolved node
     * @param logAcc Accumulator of the AVM logs in the confirmed node
     * @param sendsData Concatenated data of the sends included in the confirmed node
     * @param sendLengths Lengths of the included sends
     */
    function confirmNextNode(
        bytes32 logAcc,
        bytes calldata sendsData,
        uint256[] calldata sendLengths
    ) external whenNotPaused {
        requireUnresolvedExists();

        // There is at least one non-zombie staker
        require(stakerCount() > 0, "NO_STAKERS");

        uint256 firstUnresolved = firstUnresolvedNode();
        INode node = getNode(firstUnresolved);

        // Verify the block's deadline has passed
        node.requirePastDeadline();

        // Check that prev is latest confirmed
        require(node.prev() == latestConfirmed(), "INVALID_PREV");

        getNode(latestConfirmed()).requirePastChildConfirmDeadline();

        removeOldZombies(0);

        // All non-zombie stakers are staked on this node
        require(
            node.stakerCount() == stakerCount().add(countStakedZombies(node)),
            "NOT_ALL_STAKED"
        );

        bytes32 sendAcc = RollupLib.generateLastMessageHash(sendsData, sendLengths);
        require(node.confirmData() == RollupLib.confirmHash(sendAcc, logAcc), "CONFIRM_DATA");

        outbox.processOutgoingMessages(sendsData, sendLengths);

        confirmNextNode();

        rollupEventBridge.nodeConfirmed(firstUnresolved);

        emit SentLogs(logAcc);
    }

    /**
     * @notice Create a new stake
     * @param tokenAmount If staking in something other than eth, this is the amount of tokens staked, otherwise 0
     */
    function newStake(uint256 tokenAmount) external payable whenNotPaused {
        // Verify that sender is not already a staker
        require(!isStaked(msg.sender), "ALREADY_STAKED");

        uint256 depositAmount = receiveStakerFunds(tokenAmount);
        require(depositAmount >= currentRequiredStake(), "NOT_ENOUGH_STAKE");

        createNewStake(msg.sender, depositAmount);

        rollupEventBridge.stakeCreated(msg.sender, latestConfirmed());
    }

    /**
     * @notice Withdraw uncomitted funds owned by sender from the rollup chain
     * @param destination Address to transfer the withdrawn funds to
     */
    function withdrawStakerFunds(address payable destination)
        external
        whenNotPaused
        returns (uint256)
    {
        uint256 amount = withdrawFunds(msg.sender);
        // Note: This is an unsafe external call and could be used for reentrency
        // This is safe because it occurs after all checks and effects
        sendStakerFunds(destination, amount);
        return amount;
    }

    /**
     * @notice Move stake onto an existing node
     * @param nodeNum Inbox of the node to move stake to. This must by a child of the node the staker is currently staked on
     * @param nodeHash Node hash of nodeNum (protects against reorgs)
     */
    function stakeOnExistingNode(uint256 nodeNum, bytes32 nodeHash) external whenNotPaused {
        require(isStaked(msg.sender), "NOT_STAKED");

        require(getNodeHash(nodeNum) == nodeHash, "NODE_REORG");
        require(nodeNum >= firstUnresolvedNode() && nodeNum <= latestNodeCreated());
        INode node = getNode(nodeNum);
        require(latestStakedNode(msg.sender) == node.prev(), "NOT_STAKED_PREV");
        stakeOnNode(msg.sender, nodeNum, confirmPeriodBlocks);
    }

    /**
     * @notice Move stake onto a new node
     * @param expectedLastHash The hash of the latest sibling if it exists or else the parent (protects against reorgs)
     * @param expectedInboxHash The expected inbox accumulator hash after the assertion (protects against reorgs)
     * @param assertionBytes32Fields Assertion data for creating
     * @param assertionIntFields Assertion data for creating
     */
    function stakeOnNewNode(
        bytes32 expectedLastHash,
        bytes32 expectedInboxHash,
        bytes32[4] calldata assertionBytes32Fields,
        uint256[10] calldata assertionIntFields
    ) external whenNotPaused {
        require(isStaked(msg.sender), "NOT_STAKED");

        uint256 nodeNum = latestNodeCreated() + 1;
        uint256 deadlineBlock;
        uint256 inboxMaxCount;
        bytes32 afterInboxHash = 0;
        INode node;
        bytes32 executionHash;
        INode prevNode = getNode(latestStakedNode(msg.sender));
        {
            RollupLib.Assertion memory assertion =
                RollupLib.decodeAssertion(assertionBytes32Fields, assertionIntFields);
            executionHash = RollupLib.executionHash(assertion);
            // Make sure the previous state is correct against the node being built on
            require(
                RollupLib.beforeNodeStateHash(assertion) == prevNode.stateHash(),
                "PREV_STATE_HASH"
            );

            uint256 baseTime = prevNode.deadlineBlock().sub(assertion.beforeProposedBlock);
            require(
                prevNode.firstChildBlock() == 0 ||
                    block.number < baseTime.add(prevNode.firstChildBlock()),
                "NO_NEW_CHILDREN"
            );

            uint256 timeSinceLastNode = block.number.sub(assertion.beforeProposedBlock);
            // Verify that assertion meets the minimum Delta time requirement
            require(timeSinceLastNode >= minimumAssertionPeriod, "TIME_DELTA");

            // Minimum size requirements: each assertion must satisfy either
            require(
                // Consumes at least all inbox messages put into L1 inbox before your prev node’s L1 blocknum
                assertion.inboxMessagesRead >=
                    assertion.beforeInboxMaxCount.sub(assertion.beforeInboxCount) ||
                    // Consumes ArbGas >=100% of speed limit for time since your prev node (based on difference in L1 blocknum)
                    assertion.gasUsed >= timeSinceLastNode.mul(arbGasSpeedLimitPerBlock) ||
                    assertion.sendCount == MAX_SEND_COUNT,
                "TOO_SMALL"
            );

            // Don't allow an assertion to use above a maximum amount of gas
            require(
                assertion.gasUsed <= timeSinceLastNode.mul(arbGasSpeedLimitPerBlock).mul(4),
                "TOO_LARGE"
            );

            {
                // Set deadline rounding up to the nearest block
                uint256 checkTime =
                    assertion.gasUsed.add(arbGasSpeedLimitPerBlock.sub(1)).div(
                        arbGasSpeedLimitPerBlock
                    );
                deadlineBlock = max(block.number.add(confirmPeriodBlocks), prevNode.deadlineBlock())
                    .add(checkTime);
            }

            rollupEventBridge.nodeCreated(
                nodeNum,
                latestStakedNode(msg.sender),
                deadlineBlock,
                msg.sender
            );

            inboxMaxCount = bridge.messageCount();
            // Ensure that the assertion doesn't read past the end of the current inbox
            uint256 afterInboxCount = assertion.beforeInboxCount.add(assertion.inboxMessagesRead);
            require(afterInboxCount <= inboxMaxCount, "INBOX_PAST_END");
            if (afterInboxCount > 0) {
                afterInboxHash = bridge.inboxAccs(afterInboxCount - 1);
            }

            node = INode(
                nodeFactory.createNode(
                    RollupLib.nodeStateHash(assertion, inboxMaxCount),
                    RollupLib.challengeRoot(assertion, executionHash, block.number),
                    RollupLib.confirmHash(assertion),
                    latestStakedNode(msg.sender),
                    deadlineBlock
                )
            );
            prevNode.childCreated(nodeNum);
        }

        {
            bytes32 lastHash;
            uint256 latestSibling = prevNode.latestChildNumber();
            bool hasSibling = latestSibling > 0;
            if (hasSibling) {
                lastHash = getNodeHash(prevNode.latestChildNumber());
            } else {
                lastHash = getNodeHash(node.prev());
            }
            require(lastHash == expectedLastHash, "UNEXPECTED_LAST");
            require(afterInboxHash == expectedInboxHash, "UNEXPECTED_INBOX");
            bytes32 nodeHash =
                RollupLib.nodeHash(hasSibling, lastHash, executionHash, afterInboxHash);
            nodeCreated(node, nodeHash);
        }
        stakeOnNode(msg.sender, nodeNum, confirmPeriodBlocks);

        emit NodeCreated(
            nodeNum,
            getNodeHash(node.prev()),
            getNodeHash(nodeNum),
            executionHash,
            inboxMaxCount,
            afterInboxHash,
            assertionBytes32Fields,
            assertionIntFields
        );
    }

    /**
     * @notice Refund a staker that is currently staked on or before the latest confirmed node
     * @param stakerAddress Address of the staker whose stake is refunded
     */
    function returnOldDeposit(address stakerAddress) external override whenNotPaused {
        require(latestStakedNode(stakerAddress) <= latestConfirmed(), "TOO_RECENT");
        requireUnchallengedStaker(stakerAddress);
        withdrawStaker(stakerAddress);
    }

    /**
     * @notice Increase the amount staked for the given staker
     * @param stakerAddress Address of the staker whose stake is increased
     * @param tokenAmount If staking in something other than eth, this is the amount of tokens staked, otherwise 0
     */
    function addToDeposit(address stakerAddress, uint256 tokenAmount)
        external
        payable
        whenNotPaused
    {
        requireUnchallengedStaker(stakerAddress);
        increaseStakeBy(stakerAddress, receiveStakerFunds(tokenAmount));
    }

    /**
     * @notice Reduce the amount staked for the sender
     * @param target Target amount of stake for the staker. If this is below the current minimum, it will be set to minimum instead
     */
    function reduceDeposit(uint256 target) external whenNotPaused {
        requireUnchallengedStaker(msg.sender);
        uint256 currentRequired = currentRequiredStake();
        if (target < currentRequired) {
            target = currentRequired;
        }
        reduceStakeTo(msg.sender, target);
    }

    /**
     * @notice Start a challenge between the given stakers over the node created by the first staker assuming that the two are staked on conflicting nodes
     * @param stakers Stakers engaged in the challenge. The first staker should be staked on the first node
     * @param nodeNums Nodes of the stakers engaged in the challenge. The first node should be the earliest and is the one challenged
     * @param executionHashes Challenge related data for the two nodes
     * @param proposedTimes Times that the two nodes were proposed
     * @param maxMessageCounts Total number of messages consumed by the two nodes
     */
    function createChallenge(
        address payable[2] calldata stakers,
        uint256[2] calldata nodeNums,
        bytes32[2] calldata executionHashes,
        uint256[2] calldata proposedTimes,
        uint256[2] calldata maxMessageCounts
    ) external whenNotPaused {
        require(nodeNums[0] < nodeNums[1], "WRONG_ORDER");
        require(nodeNums[1] <= latestNodeCreated(), "NOT_PROPOSED");
        require(latestConfirmed() < nodeNums[0], "ALREADY_CONFIRMED");

        INode node1 = getNode(nodeNums[0]);
        INode node2 = getNode(nodeNums[1]);

        require(node1.prev() == node2.prev(), "DIFF_PREV");

        requireUnchallengedStaker(stakers[0]);
        requireUnchallengedStaker(stakers[1]);

        require(node1.stakers(stakers[0]), "STAKER1_NOT_STAKED");
        require(node2.stakers(stakers[1]), "STAKER2_NOT_STAKED");

        require(
            node1.challengeHash() ==
                RollupLib.challengeRootHash(
                    executionHashes[0],
                    proposedTimes[0],
                    maxMessageCounts[0]
                ),
            "CHAL_HASH"
        );

        require(
            node2.challengeHash() ==
                RollupLib.challengeRootHash(
                    executionHashes[1],
                    proposedTimes[1],
                    maxMessageCounts[1]
                ),
            "CHAL_HASH"
        );

        uint256 commonEndTime =
            node1.deadlineBlock().sub(proposedTimes[0]).add(extraChallengeTimeBlocks).add(
                getNode(node1.prev()).firstChildBlock()
            );
        // Start a challenge between staker1 and staker2. Staker1 will defend the correctness of node1, and staker2 will challenge it.
        // We must ensure that the challenge time left never underflows by restricting when nodes can be created
        address challengeAddress =
            challengeFactory.createChallenge(
                address(this),
                executionHashes[0],
                maxMessageCounts[0],
                stakers[0],
                stakers[1],
                commonEndTime.sub(proposedTimes[0]),
                commonEndTime.sub(proposedTimes[1]),
                bridge
            );

        challengeStarted(stakers[0], stakers[1], challengeAddress);

        emit RollupChallengeStarted(challengeAddress, stakers[0], stakers[1], nodeNums[0]);
    }

    /**
     * @notice Inform the rollup that the challenge between the given stakers is completed
     * @dev completeChallenge isn't pausable since in flight challenges should be allowed to complete or else they could be forced to timeout
     * @param winningStaker Address of the winning staker
     * @param losingStaker Address of the losing staker
     */
    function completeChallenge(address winningStaker, address losingStaker) external override {
        // Only the challenge contract can declare winners and losers
        require(msg.sender == inChallenge(winningStaker, losingStaker), "WRONG_SENDER");

        uint256 remainingLoserStake = amountStaked(losingStaker);
        uint256 winnerStake = amountStaked(winningStaker);
        if (remainingLoserStake > winnerStake) {
            remainingLoserStake = remainingLoserStake.sub(reduceStakeTo(losingStaker, winnerStake));
        }

        uint256 amountWon = remainingLoserStake / 2;
        increaseStakeBy(winningStaker, amountWon);
        remainingLoserStake = remainingLoserStake.sub(amountWon);
        clearChallenge(winningStaker);

        increaseStakeBy(owner, remainingLoserStake);
        turnIntoZombie(losingStaker);
    }

    /**
     * @notice Remove the given zombie from nodes it is staked on, moving backwords from the latest node it is staked on
     * @param zombieNum Index of the zombie to remove
     * @param maxNodes Maximum number of nodes to remove the zombie from (to limit the cost of this transaction)
     */
    function removeZombie(uint256 zombieNum, uint256 maxNodes) external whenNotPaused {
        require(zombieNum <= zombieCount(), "NO_SUCH_ZOMBIE");
        address zombieStakerAddress = zombieAddress(zombieNum);
        uint256 latestStakedNode = zombieLatestStakedNode(zombieNum);
        uint256 nodesRemoved = 0;
        uint256 firstUnresolved = firstUnresolvedNode();
        while (latestStakedNode >= firstUnresolved && nodesRemoved < maxNodes) {
            INode node = getNode(latestStakedNode);
            node.removeStaker(zombieStakerAddress);
            latestStakedNode = node.prev();
            nodesRemoved++;
        }
        if (latestStakedNode < firstUnresolved) {
            removeZombie(zombieNum);
        } else {
            zombieUpdateLatestStakedNode(zombieNum, latestStakedNode);
        }
    }

    /**
     * @notice Remove any zombies whose latest stake is earlier than the first unresolved node
     * @param startIndex Index in the zombie list to start removing zombies from (to limit the cost of this transaction)
     */
    function removeOldZombies(uint256 startIndex) public {
        uint256 currentZombieCount = zombieCount();
        uint256 firstUnresolved = firstUnresolvedNode();
        for (uint256 i = startIndex; i < currentZombieCount; i++) {
            while (zombieLatestStakedNode(i) < firstUnresolved) {
                removeZombie(i);
                currentZombieCount--;
                if (i >= currentZombieCount) {
                    return;
                }
            }
        }
    }

    /**
     * @notice Calculate the current amount of funds required to place a new stake in the rollup
     * @dev If the stake requirement get's too high, this function may stop reverting due to overflow, but
     * that only blocks operations that should be blocked anyway
     * @return The current minimum stake requirement
     */
    function currentRequiredStake() public view returns (uint256) {
        // If there are no unresolved nodes, then you can use the base stake
        uint256 firstUnresolvedNodeNum = firstUnresolvedNode();
        if (firstUnresolvedNodeNum - 1 == latestNodeCreated()) {
            return baseStake;
        }
        INode firstUnresolved = getNode(firstUnresolvedNodeNum);

        uint256 firstUnresolvedDeadline = firstUnresolved.deadlineBlock();
        if (block.number < firstUnresolvedDeadline) {
            return baseStake;
        }
        uint24[10] memory numerators =
            [1, 122971, 128977, 80017, 207329, 114243, 314252, 129988, 224562, 162163];
        uint24[10] memory denominators =
            [1, 114736, 112281, 64994, 157126, 80782, 207329, 80017, 128977, 86901];
        uint256 firstUnresolvedAge = block.number.sub(firstUnresolvedDeadline);
        uint256 periodsPassed = firstUnresolvedAge.mul(10).div(confirmPeriodBlocks);
        // Overflow check
        if (periodsPassed.div(10) >= 255) {
            return type(uint256).max;
        }
        uint256 baseMultiplier = 2**periodsPassed.div(10);
        uint256 withNumerator = baseMultiplier * numerators[periodsPassed % 10];
        // Overflow check
        if (withNumerator / baseMultiplier != numerators[periodsPassed % 10]) {
            return type(uint256).max;
        }
        uint256 multiplier = withNumerator.div(denominators[periodsPassed % 10]);
        if (multiplier == 0) {
            multiplier = 1;
        }
        uint256 fullStake = baseStake * multiplier;
        // Overflow check
        if (fullStake / baseStake != multiplier) {
            return type(uint256).max;
        }
        return fullStake;
    }

    /**
     * @notice Calculate the number of zombies staked on the given node
     *
     * @dev This function could be uncallable if there are too many zombies. However,
     * removeZombie and removeOldZombies can be used to remove any zombies that exist
     * so that this will then be callable
     *
     * @param node The node on which to count staked zombies
     * @return The number of zombies staked on the node
     */
    function countStakedZombies(INode node) public view returns (uint256) {
        uint256 currentZombieCount = zombieCount();
        uint256 stakedZombieCount = 0;
        for (uint256 i = 0; i < currentZombieCount; i++) {
            if (node.stakers(zombieAddress(i))) {
                stakedZombieCount++;
            }
        }
        return stakedZombieCount;
    }

    /**
     * @notice Verify that there are some number of nodes still unresolved
     */
    function requireUnresolvedExists() public view {
        uint256 firstUnresolved = firstUnresolvedNode();
        require(
            firstUnresolved > latestConfirmed() && firstUnresolved <= latestNodeCreated(),
            "NO_UNRESOLVED"
        );
    }

    function requireUnresolved(uint256 nodeNum) public view {
        require(nodeNum >= firstUnresolvedNode(), "ALREADY_DECIDED");
        require(nodeNum <= latestNodeCreated(), "DOESNT_EXIST");
    }

    /**
     * @notice Ensure that funds are properly received
     * @param tokenAmount If staking in something other than eth, this is the amount of tokens to transfer, otherwise 0
     * @return Amount of funds that have been received by the rollup
     */
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

    /**
     * @notice Send funds to the given address, if staking is eth, transfer eth, otherwise transfer tokens
     * @param destination Address to tranfer funds to
     * @param amount Amount of funds to transfer
     */
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

    /**
     * @notice Verify that the given address is staked and not actively in a challenge
     * @param stakerAddress Address to check
     */
    function requireUnchallengedStaker(address stakerAddress) private view {
        require(isStaked(stakerAddress), "NOT_STAKED");
        require(currentChallenge(stakerAddress) == address(0), "IN_CHAL");
    }
}
