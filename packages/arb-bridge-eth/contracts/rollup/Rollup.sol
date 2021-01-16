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

contract Rollup is Pausable, IRollup {
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

    uint8 internal constant INITIALIZATION_MSG_TYPE = 4;

    uint256 public override latestConfirmed;
    uint256 public override firstUnresolvedNode;
    uint256 public override latestNodeCreated;
    mapping(uint256 => INode) public override nodes;
    uint256 public override lastStakeBlock;

    address payable[] public override stakerList;
    mapping(address => Staker) public stakerMap;

    Zombie[] private zombies;

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

    mapping(address => uint256) public override withdrawableFunds;

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
        nodes[0] = node;

        challengePeriodBlocks = _challengePeriodBlocks;
        arbGasSpeedLimitPerBlock = _arbGasSpeedLimitPerBlock;
        baseStake = _baseStake;
        stakeToken = _stakeToken;
        owner = _owner;
        admin = ProxyAdmin(_admin);

        firstUnresolvedNode = 1;

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

    function ownerPause() external onlyOwner {
        _pause();
    }

    function ownerResume() external onlyOwner {
        _unpause();
    }

    function rejectNextNode(uint256 successorWithStake, address stakerAddress)
        external
        override
        whenNotPaused
    {
        checkUnresolved();

        INode node = nodes[firstUnresolvedNode];
        if (node.prev() == latestConfirmed) {
            checkNoRecentStake();
            require(successorWithStake > firstUnresolvedNode, "SUCCESSOR_TO_LOW");
            require(successorWithStake <= latestNodeCreated, "SUCCESSOR_TO_HIGH");
            require(stakerMap[stakerAddress].isStaked, "NOT_STAKED");

            // Confirm that someone is staked on some sibling node
            nodes[successorWithStake].checkRejectExample(latestConfirmed, stakerAddress);

            // Verify the block's deadline has passed
            node.checkPastDeadline();

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
    ) external override whenNotPaused {
        INode node = checkConfirmValidBefore();
        removeOldZombies(0);
        checkConfirmValidAfter(node);

        bytes32 sendAcc = RollupLib.generateLastMessageHash(sendsData, sendLengths);
        require(node.confirmData() == RollupLib.confirmHash(sendAcc, logAcc), "CONFIRM_DATA");

        outbox.processOutgoingMessages(sendsData, sendLengths);

        destroyNode(latestConfirmed);

        latestConfirmed = firstUnresolvedNode;
        firstUnresolvedNode++;

        emit SentLogs(logAcc);
    }

    function newStake(uint256 tokenAmount) external payable override whenNotPaused {
        // Verify that sender is not already a staker
        require(!stakerMap[msg.sender].isStaked, "ALREADY_STAKED");

        uint256 depositAmount = receiveStakerFunds(tokenAmount);
        require(depositAmount >= currentRequiredStake(), "NOT_ENOUGH_STAKE");

        uint256 stakerIndex = stakerList.length;
        stakerList.push(msg.sender);
        stakerMap[msg.sender] = Staker(
            stakerIndex,
            latestConfirmed,
            depositAmount,
            address(0),
            true
        );
        lastStakeBlock = block.number;
    }

    function withdrawStakerFunds(address payable destination)
        external
        override
        whenNotPaused
        returns (uint256)
    {
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
    ) external override whenNotPaused {
        Staker storage staker = stakerMap[msg.sender];
        require(staker.isStaked, "NOT_STAKED");

        require(blockhash(blockNumber) == blockHash, "invalid known block");
        require(nodeNum >= firstUnresolvedNode && nodeNum <= latestNodeCreated);
        INode node = nodes[nodeNum];
        require(staker.latestStakedNode == node.prev(), "NOT_STAKED_PREV");
        node.addStaker(msg.sender);
        staker.latestStakedNode = nodeNum;
    }

    function stakeOnNewNode(
        bytes32 blockHash,
        uint256 blockNumber,
        uint256 nodeNum,
        bytes32[7] calldata assertionBytes32Fields,
        uint256[10] calldata assertionIntFields
    ) external override whenNotPaused {
        Staker storage staker = stakerMap[msg.sender];
        require(staker.isStaked, "NOT_STAKED");
        require(blockhash(blockNumber) == blockHash, "invalid known block");
        require(nodeNum == latestNodeCreated + 1, "NODE_NUM");
        RollupLib.Assertion memory assertion =
            RollupLib.decodeAssertion(assertionBytes32Fields, assertionIntFields);
        uint256 prev = staker.latestStakedNode;
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
        staker.latestStakedNode = latestNodeCreated;

        emit NodeCreated(
            latestNodeCreated,
            assertionBytes32Fields,
            assertionIntFields,
            inboxMaxCount,
            inboxMaxAcc
        );
    }

    function returnOldDeposit(address stakerAddress) external override whenNotPaused {
        Staker storage staker = stakerMap[stakerAddress];
        require(staker.latestStakedNode <= latestConfirmed, "TOO_RECENT");
        checkUnchallengedStaker(staker);
        uint256 amountStaked = staker.amountStaked;
        deleteStaker(staker);
        withdrawableFunds[stakerAddress] += amountStaked;
    }

    function addToDeposit(address stakerAddress, uint256 tokenAmount)
        external
        payable
        override
        whenNotPaused
    {
        Staker storage staker = stakerMap[stakerAddress];
        checkUnchallengedStaker(staker);
        staker.amountStaked += receiveStakerFunds(tokenAmount);
    }

    function reduceDeposit(uint256 maxReduction) external override whenNotPaused {
        Staker storage staker = stakerMap[msg.sender];
        checkUnchallengedStaker(staker);
        uint256 currentRequired = currentRequiredStake();
        require(staker.amountStaked > currentRequired);
        uint256 withdrawAmount = staker.amountStaked - currentRequired;
        // Cap withdrawAmount at maxReduction
        if (withdrawAmount > maxReduction) {
            withdrawAmount = maxReduction;
        }
        staker.amountStaked -= withdrawAmount;
        withdrawableFunds[msg.sender] += withdrawAmount;
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
        require(nodeNums[1] <= latestNodeCreated, "NOT_PROPOSED");
        require(latestConfirmed < nodeNums[0], "ALREADY_CONFIRMED");

        INode node1 = nodes[nodeNums[0]];
        INode node2 = nodes[nodeNums[1]];

        require(node1.prev() == node2.prev(), "DIFF_PREV");

        checkUnchallengedStaker(stakerMap[stakers[0]]);
        checkUnchallengedStaker(stakerMap[stakers[1]]);

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

        stakerMap[stakers[0]].currentChallenge = challengeAddress;
        stakerMap[stakers[1]].currentChallenge = challengeAddress;

        emit RollupChallengeStarted(challengeAddress, stakers[0], stakers[1], nodeNums[0]);
    }

    // completeChallenge isn't pausable since in flight challenges should be allowed to complete or else they
    // could be forced to timeout
    function completeChallenge(address winningStaker, address losingStaker) external override {
        Staker storage winner = stakerMap[winningStaker];
        Staker storage loser = stakerMap[losingStaker];
        require(winner.currentChallenge == loser.currentChallenge, "SAME_CHAL");
        // Only the challenge contract can declare winners and losers
        require(msg.sender == address(winner.currentChallenge), "WRONG_SENDER");

        uint256 loserStake = loser.amountStaked;

        if (loserStake > winner.amountStaked) {
            uint256 extraLoserStake = loserStake - winner.amountStaked;
            withdrawableFunds[losingStaker] += extraLoserStake;
            loserStake -= extraLoserStake;
        }

        uint256 amountWon = loserStake / 2;
        winner.amountStaked += amountWon;
        loserStake -= amountWon;
        winner.currentChallenge = address(0);

        // TODO: deposit extra loserStake into ArbOS

        zombies.push(Zombie(losingStaker, loser.latestStakedNode));
        deleteStaker(loser);
    }

    function removeZombie(uint256 zombieNum, uint256 maxNodes) external override whenNotPaused {
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

    function stakerCount() external view override returns (uint256) {
        return stakerList.length;
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

    function deleteStaker(Staker storage staker) private {
        uint256 stakerIndex = staker.index;
        address stakerAddress = stakerList[stakerIndex];
        stakerList[stakerIndex] = stakerList[stakerList.length - 1];
        stakerMap[stakerList[stakerIndex]].index = stakerIndex;
        stakerList.pop();
        delete stakerMap[stakerAddress];
    }

    function checkUnchallengedStaker(Staker storage staker) private view {
        require(staker.isStaked, "NOT_STAKED");
        require(staker.currentChallenge == address(0), "IN_CHAL");
    }

    function checkConfirmValidBefore() public view override returns (INode) {
        checkUnresolved();
        checkNoRecentStake();

        // There is at least one non-zombie staker
        require(stakerList.length > 0, "NO_STAKERS");

        INode node = nodes[firstUnresolvedNode];

        // Verify the block's deadline has passed
        node.checkPastDeadline();

        // Check that prev is latest confirmed
        require(node.prev() == latestConfirmed, "INVALID_PREV");

        return node;
    }

    function checkConfirmValidAfter(INode node) public view override {
        // All non-zombie stakers are staked on this node
        require(
            node.stakerCount() == stakerList.length + countStakedZombies(node),
            "NOT_ALL_STAKED"
        );
    }
}
