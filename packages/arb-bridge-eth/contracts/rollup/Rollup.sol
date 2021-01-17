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

import "./IRollup.sol";
import "./INode.sol";
import "../bridge/interfaces/IBridge.sol";
import "../bridge/interfaces/IOutbox.sol";
import "../bridge/interfaces/IBridge.sol";
import "../interfaces/IERC20.sol";

import "../bridge/Messages.sol";
import "./RollupLib.sol";
import "../challenge/ChallengeLib.sol";

contract Rollup is RollupCore, Pausable, IRollup {
    uint8 internal constant INITIALIZATION_MSG_TYPE = 4;

    // Rollup Config
    uint256 public override challengePeriodBlocks;
    uint256 public override arbGasSpeedLimitPerBlock;
    uint256 public override baseStake;
    address public override stakeToken;

    // Bridge is an IInbox and IOutbox
    IBridge public override bridge;
    IOutbox public override outbox;
    IChallengeFactory public override challengeFactory;
    INodeFactory public override nodeFactory;
    address public owner;
    ProxyAdmin admin;

    modifier onlyOwner {
        require(msg.sender == owner, "ONLY_OWNER");
        _;
    }

    function initialize(
        IOutbox _outbox,
        bytes32 _machineHash,
        uint256 _challengePeriodBlocks,
        uint256 _arbGasSpeedLimitPerBlock,
        uint256 _baseStake,
        address _stakeToken,
        address _owner,
        IBridge _bridge,
        address _challengeFactory,
        address _nodeFactory,
        bytes memory _extraConfig,
        address _admin
    ) external override {
        bridge = _bridge;
        bridge.setInbox(address(this), true);
        outbox = _outbox;
        bytes32 initMsgHash =
            keccak256(
                abi.encodePacked(
                    uint256(_challengePeriodBlocks),
                    uint256(_arbGasSpeedLimitPerBlock),
                    uint256(_baseStake),
                    bytes32(bytes20(_stakeToken)),
                    bytes32(bytes20(_owner)),
                    _extraConfig
                )
            );
        bridge.deliverMessageToInbox(
            Messages.messageHash(
                INITIALIZATION_MSG_TYPE,
                address(this),
                block.number,
                block.timestamp, // solhint-disable-line not-rely-on-time
                0,
                initMsgHash
            )
        );

        challengeFactory = IChallengeFactory(_challengeFactory);
        nodeFactory = INodeFactory(_nodeFactory);

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
        initializeCore(node);

        challengePeriodBlocks = _challengePeriodBlocks;
        arbGasSpeedLimitPerBlock = _arbGasSpeedLimitPerBlock;
        baseStake = _baseStake;
        stakeToken = _stakeToken;
        owner = _owner;
        admin = ProxyAdmin(_admin);

        emit RollupCreated(_machineHash);
    }

    function addOutbox(IOutbox _outbox) external onlyOwner {
        outbox = _outbox;
        bridge.setOutbox(address(_outbox), true);
    }

    function addInbox(address _inbox) external onlyOwner {
        bridge.setInbox(address(_inbox), true);
    }

    function upgradeImplementation(address _newRollup) external onlyOwner {
        address currentAddress = address(this);
        admin.upgrade(TransparentUpgradeableProxy(payable(currentAddress)), _newRollup);
    }

    function upgradeImplementationAndCall(address _newRollup, bytes calldata data)
        external
        onlyOwner
    {
        address currentAddress = address(this);
        admin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(currentAddress)),
            _newRollup,
            data
        );
    }

    function pause() external onlyOwner {
        _pause();
    }

    function resume() external onlyOwner {
        _unpause();
    }

    function truncateNodes(uint256 _latestNodeCreated) external onlyOwner {
        uint256 oldLatestNodeCreated = latestNodeCreated();
        require(_latestNodeCreated < oldLatestNodeCreated, "TOO_NEW");
        require(_latestNodeCreated >= firstUnresolvedNode() - 1, "TOO_OLD");

        for (uint256 i = oldLatestNodeCreated; i >= _latestNodeCreated; i--) {
            INode node = getNode(i);
            node.destroy();
        }
        updateLatestNodeCreated(_latestNodeCreated);
    }

    function rejectNextNode(uint256 successorWithStake, address stakerAddress)
        external
        override
        whenNotPaused
    {
        checkUnresolved();
        uint256 latest = latestConfirmed();
        uint256 firstUnresolved = firstUnresolvedNode();
        INode node = getNode(firstUnresolved);
        if (node.prev() == latest) {
            checkNoRecentStake();
            require(successorWithStake > firstUnresolved, "SUCCESSOR_TO_LOW");
            require(successorWithStake <= latestNodeCreated(), "SUCCESSOR_TO_HIGH");
            require(isStaked(stakerAddress), "NOT_STAKED");

            // Confirm that someone is staked on some sibling node
            getNode(successorWithStake).checkRejectExample(latest, stakerAddress);

            // Verify the block's deadline has passed
            node.checkPastDeadline();

            removeOldZombies(0);

            // Verify that no staker is staked on this node
            require(node.stakerCount() == countStakedZombies(node), "HAS_STAKERS");
        }
        rejectNextNode();
    }

    function confirmNextNode(
        bytes32 logAcc,
        bytes calldata sendsData,
        uint256[] calldata sendLengths
    ) external override whenNotPaused {
        INode node = checkConfirmValidBefore();
        removeOldZombies(0);
        checkConfirmValidAfter(node);

        bytes32 sendAcc = RollupLib.generateLastMessageHash(sendsData, sendLengths);
        require(node.confirmData() == RollupLib.confirmHash(sendAcc, logAcc), "CONFIRM_DATA");

        outbox.processOutgoingMessages(sendsData, sendLengths);

        confirmNextNode();

        emit SentLogs(logAcc);
    }

    function newStake(uint256 tokenAmount) external payable override whenNotPaused {
        // Verify that sender is not already a staker
        require(!isStaked(msg.sender), "ALREADY_STAKED");

        uint256 depositAmount = receiveStakerFunds(tokenAmount);
        require(depositAmount >= currentRequiredStake(), "NOT_ENOUGH_STAKE");

        createNewStake(msg.sender, depositAmount);
    }

    function withdrawStakerFunds(address payable destination)
        external
        override
        whenNotPaused
        returns (uint256)
    {
        uint256 amount = withdrawFunds(msg.sender);
        // Note: This is an unsafe external call and could be used for reentrency
        // This is safe because it occurs after all checks and effects
        sendStakerFunds(destination, amount);
        return amount;
    }

    function stakeOnExistingNode(
        bytes32 blockHash,
        uint256 blockNumber,
        uint256 nodeNum
    ) external override whenNotPaused {
        require(isStaked(msg.sender), "NOT_STAKED");

        require(blockhash(blockNumber) == blockHash, "invalid known block");
        require(nodeNum >= firstUnresolvedNode() && nodeNum <= latestNodeCreated());
        INode node = getNode(nodeNum);
        require(latestStakedNode(msg.sender) == node.prev(), "NOT_STAKED_PREV");
        stakeOnNode(msg.sender, nodeNum);
    }

    function stakeOnNewNode(
        bytes32 blockHash,
        uint256 blockNumber,
        uint256 nodeNum,
        bytes32[7] calldata assertionBytes32Fields,
        uint256[10] calldata assertionIntFields
    ) external override whenNotPaused {
        require(isStaked(msg.sender), "NOT_STAKED");
        require(blockhash(blockNumber) == blockHash, "invalid known block");
        require(nodeNum == latestNodeCreated() + 1, "NODE_NUM");
        RollupLib.Assertion memory assertion =
            RollupLib.decodeAssertion(assertionBytes32Fields, assertionIntFields);
        INode prevNode = getNode(latestStakedNode(msg.sender));
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

        uint256 timeSinceLastNode = block.number - assertion.beforeProposedBlock;
        // Verify that assertion meets the minimum Delta time requirement
        require(timeSinceLastNode >= minimumAssertionPeriod(), "TIME_DELTA");

        // Minimum size requirements: each assertion must satisfy either
        require(
            // Consumes at least all inbox messages put into L1 inbox before your prev node’s L1 blocknum
            assertion.inboxMessagesRead >=
                assertion.beforeInboxMaxCount - assertion.beforeInboxCount ||
                // Consumes ArbGas >=100% of speed limit for time since your prev node (based on difference in L1 blocknum)
                assertion.gasUsed >= timeSinceLastNode * arbGasSpeedLimitPerBlock,
            "TOO_SMALL"
        );

        // Don't allow an assertion to use above a maximum amount of gas
        require(assertion.gasUsed <= timeSinceLastNode * arbGasSpeedLimitPerBlock * 4, "TOO_LARGE");

        uint256 deadlineBlock = block.number + challengePeriodBlocks;
        uint256 prevDeadlineBlock = prevNode.deadlineBlock();
        if (deadlineBlock < prevDeadlineBlock) {
            deadlineBlock = prevDeadlineBlock;
        }
        deadlineBlock += assertion.gasUsed / arbGasSpeedLimitPerBlock;

        INode node =
            INode(
                nodeFactory.createNode(
                    RollupLib.nodeStateHash(assertion, inboxMaxCount),
                    RollupLib.challengeRoot(
                        assertion,
                        inboxMaxCount,
                        inboxMaxAcc,
                        assertion.gasUsed / arbGasSpeedLimitPerBlock
                    ),
                    RollupLib.confirmHash(assertion),
                    latestStakedNode(msg.sender),
                    deadlineBlock
                )
            );

        nodeCreated(node);
        stakeOnNode(msg.sender, nodeNum);

        emit NodeCreated(
            nodeNum,
            assertionBytes32Fields,
            assertionIntFields,
            inboxMaxCount,
            inboxMaxAcc
        );
    }

    function returnOldDeposit(address stakerAddress) external override whenNotPaused {
        require(latestStakedNode(stakerAddress) <= latestConfirmed(), "TOO_RECENT");
        checkUnchallengedStaker(stakerAddress);
        withdrawStaker(stakerAddress);
    }

    function addToDeposit(address stakerAddress, uint256 tokenAmount)
        external
        payable
        override
        whenNotPaused
    {
        checkUnchallengedStaker(stakerAddress);
        increaseStakeBy(stakerAddress, receiveStakerFunds(tokenAmount));
    }

    function reduceDeposit(uint256 target) external override whenNotPaused {
        checkUnchallengedStaker(msg.sender);
        uint256 currentRequired = currentRequiredStake();
        if (target > currentRequired) {
            target = currentRequired;
        }
        reduceStakeTo(msg.sender, target);
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
    ) external override whenNotPaused {
        require(nodeNums[0] < nodeNums[1], "WRONG_ORDER");
        require(nodeNums[1] <= latestNodeCreated(), "NOT_PROPOSED");
        require(latestConfirmed() < nodeNums[0], "ALREADY_CONFIRMED");

        INode node1 = getNode(nodeNums[0]);
        INode node2 = getNode(nodeNums[1]);

        require(node1.prev() == node2.prev(), "DIFF_PREV");

        checkUnchallengedStaker(stakers[0]);
        checkUnchallengedStaker(stakers[1]);

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
                address(this),
                nodeFields[0],
                nodeFields[1],
                nodeFields[2],
                executionCheckTime,
                stakers[0],
                stakers[1],
                challengePeriodBlocks
            );

        challengeStarted(stakers[0], stakers[1], challengeAddress);

        emit RollupChallengeStarted(challengeAddress, stakers[0], stakers[1], nodeNums[0]);
    }

    // completeChallenge isn't pausable since in flight challenges should be allowed to complete or else they
    // could be forced to timeout
    function completeChallenge(address winningStaker, address losingStaker) external override {
        // Only the challenge contract can declare winners and losers
        require(msg.sender == inChallenge(winningStaker, losingStaker), "WRONG_SENDER");

        uint256 loserStake = amountStaked(losingStaker);
        uint256 winnerStake = amountStaked(winningStaker);
        if (loserStake > winnerStake) {
            loserStake -= reduceStakeTo(losingStaker, winnerStake);
        }

        uint256 amountWon = loserStake / 2;
        increaseStakeBy(winningStaker, amountWon);
        loserStake -= amountWon;
        clearChallenge(winningStaker);

        // TODO: deposit extra loserStake into ArbOS
        turnIntoZombie(losingStaker);
    }

    function removeZombie(uint256 zombieNum, uint256 maxNodes) external override whenNotPaused {
        require(zombieNum <= zombieCount(), "NO_SUCH_ZOMBIE");
        address zombieStakerAddress = zombieAddress(zombieNum);
        uint256 latestStakedNode = zombieLatestStakedNode(zombieNum);
        uint256 nodesRemoved = 0;
        uint256 firstUnresolved = firstUnresolvedNode();
        while (latestStakedNode > firstUnresolved && nodesRemoved < maxNodes) {
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

    function removeOldZombies(uint256 startIndex) public override {
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

    function currentRequiredStake() public view override returns (uint256) {
        uint256 MAX_INT = 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff;
        uint256 latestConfirmedDeadline = getNode(latestConfirmed()).deadlineBlock();
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
        uint256 currentZombieCount = zombieCount();
        uint256 stakedZombieCount = 0;
        for (uint256 i = 0; i < currentZombieCount; i++) {
            if (node.stakers(zombieAddress(i))) {
                stakedZombieCount++;
            }
        }
        return stakedZombieCount;
    }

    function checkNoRecentStake() public view override {
        // No stake has been placed during the last challengePeriod blocks
        require(block.number - lastStakeBlock() >= challengePeriodBlocks, "RECENT_STAKE");
    }

    function checkUnresolved() public view override {
        uint256 firstUnresolved = firstUnresolvedNode();
        require(
            firstUnresolved > latestConfirmed() && firstUnresolved <= latestNodeCreated(),
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

    function checkUnchallengedStaker(address stakerAddress) private view {
        require(isStaked(stakerAddress), "NOT_STAKED");
        require(currentChallenge(stakerAddress) == address(0), "IN_CHAL");
    }

    function checkConfirmValidBefore() public view override returns (INode) {
        checkUnresolved();
        checkNoRecentStake();

        // There is at least one non-zombie staker
        require(stakerCount() > 0, "NO_STAKERS");

        uint256 firstUnresolved = firstUnresolvedNode();
        INode node = getNode(firstUnresolved);

        // Verify the block's deadline has passed
        node.checkPastDeadline();

        // Check that prev is latest confirmed
        require(node.prev() == latestConfirmed(), "INVALID_PREV");

        return node;
    }

    function checkConfirmValidAfter(INode node) public view override {
        // All non-zombie stakers are staked on this node
        require(node.stakerCount() == stakerCount() + countStakedZombies(node), "NOT_ALL_STAKED");
    }
}
