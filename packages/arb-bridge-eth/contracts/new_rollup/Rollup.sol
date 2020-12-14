// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.5.17;

import "./Node.sol";
import "./Assertion.sol";

import "../inbox/IGlobalInbox.sol";
import "../rollup/RollupUtils.sol";

contract Rollup {
    event SentLogs(bytes32 logsAccHash);

    struct Staker {
        uint256 latestStakedNode;
        uint256 amountStaked;
        // currentChallenge is 0 if staker is not in a challenge
        address currentChallenge;
        bool isZombie;
        bool isStaked;
    }

    uint256 latestConfirmed;
    uint256 firstUnresolvedNode;
    uint256 latestNodeCreated;
    Node[] nodes;
    uint256 lastStakeBlock;
    uint256 stakerCount;
    mapping(address => Staker) stakers;

    uint256 baseStake;
    uint256 challengePeriod;

    IGlobalInbox public globalInbox;

    constructor(bytes32 machineHash) public {
        bytes32 assertionHash = Assertion.hashAssertion(machineHash, 0, 0, 0, 0, 0, 0, 0);

        bytes32 inboxNodeHash = Assertion.hashInboxNode(0, 0, 0);

        bytes32 executionNodeHash = Assertion.hashExecutionNode(assertionHash, assertionHash, 0);

        Node node = new Node(
            inboxNodeHash,
            executionNodeHash,
            latestConfirmed,
            block.number,
            0, // TODO: deadline block
            0
        );
        nodes[0] = node;
    }

    function rejectNextNode(uint256 successorWithStake, address stakerAddress) external {
        // No stake has been placed during the last challengePeriod blocks
        require(block.number - lastStakeBlock >= challengePeriod);

        require(!stakers[stakerAddress].isZombie);

        // Confirm that someone is staked on some sibling node
        Node stakedSiblingNode = nodes[successorWithStake];
        // stakedSiblingNode is a child of latestConfirmed
        require(stakedSiblingNode.prev() == latestConfirmed);
        // staker is actually staked on stakedSiblingNode
        require(stakedSiblingNode.stakers(stakerAddress));

        Node node = nodes[firstUnresolvedNode];
        node.confirmInvalid();
        discardUnresolvedNode();
        node.destroy();
    }

    // If the node previous to this one is not the latest confirmed, we can reject immediately
    function rejectNextNodeOutOfOrder() external {
        Node node = nodes[firstUnresolvedNode];
        node.confirmOutOfOrder(latestConfirmed);
        discardUnresolvedNode();
        node.destroy();
    }

    function confirmNextNode(
        bytes32 logAcc,
        bytes calldata messages,
        uint256 beforeSendCount,
        uint256 sendCount
    ) external {
        // No stake has been placed during the last challengePeriod blocks
        require(block.number - lastStakeBlock >= challengePeriod);
        Node node = nodes[firstUnresolvedNode];
        node.confirmValid(stakerCount, latestConfirmed);

        (bytes32 lastMsgHash, ) = RollupUtils.generateLastMessageHash(messages, 0, sendCount);

        bytes32 confirmData = keccak256(
            abi.encodePacked(lastMsgHash, logAcc, beforeSendCount, sendCount)
        );

        // TODO: check that confirmData matches up with node

        // Send all messages is a single batch
        globalInbox.sendMessages(messages, beforeSendCount, beforeSendCount + sendCount);

        emit SentLogs(logAcc);

        discardUnresolvedNode();
        node.destroy();
    }

    function newStakeOnExistingNode(
        bytes32 blockHash,
        uint256 blockNumber,
        uint256 nodeNum
    ) external payable {
        require(blockhash(blockNumber) == blockHash, "invalid known block");
        verifyCanStake();
        // Must stake on valid node
        checkValidNodeNumForStake(nodeNum);
        Node node = nodes[nodeNum];
        require(node.prev() == latestConfirmed);
        addStaker(nodeNum, node);
    }

    /*

    // Precondition Data
    bytes32 beforeInboxAcc;
    uint256 beforeInboxCount;

    bytes32 topInboxAcc;
    uint256 topInboxCount;

    bytes32 beforeStateHash;

    // Postcondition Data

    bytes32 afterInboxAcc;
    uint256 afterInboxCount;

    bytes32 afterStateHash

    uint256 totalSteps;
    */

    function newStakeOnNewNode(
        bytes32 blockHash,
        uint256 blockNumber,
        uint256 nodeNum,
        uint256 prev,
        bytes32 beforeInboxHash,
        uint256 beforeInboxCount,
        bytes32 beforeStateHash,
        bytes32 afterInboxHash,
        uint256 afterInboxCount,
        bytes32 afterStateHash,
        uint256 numSteps
    ) external payable {
        require(blockhash(blockNumber) == blockHash, "invalid known block");
        verifyCanStake();
        require(nodeNum == latestNodeCreated + 1);
        require(prev == latestConfirmed);
        require(afterInboxCount >= beforeInboxCount);

        (bytes32 inboxValue, uint256 inboxCount) = globalInbox.getInbox(address(this));

        require(afterInboxCount <= inboxCount);

        bytes32 inboxNodeHash = Assertion.hashInboxNode(
            afterInboxHash,
            inboxValue,
            inboxCount - afterInboxCount
        );

        bytes32 executionNodeHash = Assertion.hashExecutionNode(
            Assertion.hashAssertion(beforeStateHash, beforeInboxCount, beforeInboxHash),
            Assertion.hashAssertion(afterStateHash, afterInboxCount, afterInboxHash),
            numSteps
        );

        // TODO: Verify that assertion meets the minimum size requirement
        // TODO: Verify that assertion meets the minimum Delta time requirement
        Node node = new Node(
            inboxNodeHash,
            executionNodeHash,
            latestConfirmed,
            block.number,
            0, // TODO: deadline block
            0
        );
        addStaker(nodeNum, node);
        nodes[nodeNum] = node;
        latestNodeCreated++;
    }

    function addStakeOnExistingNode(
        bytes32 blockHash,
        uint256 blockNumber,
        uint256 nodeNum
    ) external {
        require(blockhash(blockNumber) == blockHash, "invalid known block");
        checkValidNodeNumForStake(nodeNum);
        Staker storage staker = stakers[msg.sender];
        require(!staker.isZombie);
        Node node = nodes[nodeNum];
        require(staker.latestStakedNode == node.prev());
        node.addStaker(msg.sender);
        staker.latestStakedNode = nodeNum;
    }

    function addStakeOnNewNode(
        bytes32 blockHash,
        uint256 blockNumber,
        uint256 nodeNum
    ) external {
        require(blockhash(blockNumber) == blockHash, "invalid known block");
        require(nodeNum == latestNodeCreated + 1);
        Staker storage staker = stakers[msg.sender];
        require(!staker.isZombie);

        // TODO: Verify that the preconditions of assertion are consistent with the postconditions of prev
        // TODO: Verify that assertion meets the minimum size requirement
        // TODO: Verify that assertion meets the minimum Delta time requirement

        Node node = new Node(
            0,
            0, // TODO: assertion hash
            staker.latestStakedNode,
            block.number,
            0, // TODO: deadline block
            0
        );
        node.addStaker(msg.sender);
        nodes[nodeNum] = node;
        staker.latestStakedNode = nodeNum;
        latestNodeCreated++;
    }

    function returnOldDeposit(address payable stakerAddress) external {
        Staker storage staker = stakers[stakerAddress];
        checkUnchallengedStaker(staker);
        require(staker.latestStakedNode <= latestConfirmed);

        delete stakers[stakerAddress];
        // TODO: Staker could force transfer to revert. We may want to allow funds to be withdrawn separately
        stakerAddress.transfer(staker.amountStaked);
    }

    function addToDeposit() external payable {
        Staker storage staker = stakers[msg.sender];
        checkUnchallengedStaker(staker);
        staker.amountStaked += msg.value;
    }

    function reduceDeposit(uint256 maxReduction) external {
        Staker storage staker = stakers[msg.sender];
        checkUnchallengedStaker(staker);
        uint256 currentRequired = currentRequiredStake();
        require(staker.amountStaked > currentRequired);
        uint256 withdrawAmount = staker.amountStaked - currentRequired;
        // Cap withdrawAmount at maxReduction
        if (withdrawAmount > maxReduction) {
            withdrawAmount = maxReduction;
        }
        msg.sender.transfer(withdrawAmount);
    }

    function removeZombieStaker(uint256 nodeNum, address stakerAddress) external {
        require(stakers[stakerAddress].isZombie);
        nodes[nodeNum].removeStaker(stakerAddress);
    }

    function createChallenge(
        address staker1Address,
        uint256 nodeNum1,
        address staker2Address,
        uint256 nodeNum2
    ) external {
        Staker storage staker1 = stakers[staker1Address];
        Staker storage staker2 = stakers[staker2Address];
        Node node1 = nodes[nodeNum1];
        Node node2 = nodes[nodeNum2];

        checkUnchallengedStaker(staker1);
        require(node1.stakers(staker1Address));

        checkUnchallengedStaker(staker2);
        require(node2.stakers(staker2Address));

        require(node1.prev() == node2.prev());
        require(latestConfirmed < nodeNum1);
        require(nodeNum1 < nodeNum2);
        require(nodeNum2 <= latestNodeCreated);

        // Start a challenge between staker1 and staker2. Staker1 will defend the correctness of node1, and staker2 will challenge it.
        // TODO: How to we want to handle the two challenge types

        // TODO: Actually launch challenge
        address challengeAddress = address(0);
        staker1.currentChallenge = challengeAddress;
        staker2.currentChallenge = challengeAddress;
    }

    function completeChallenge(address winningStaker, address payable losingStaker) external {
        Staker storage winner = stakers[winningStaker];
        Staker storage loser = stakers[losingStaker];

        // Only the challenge contract can declare winners and losers
        require(winner.currentChallenge == msg.sender);
        require(loser.currentChallenge == msg.sender);

        if (loser.amountStaked > winner.amountStaked) {
            uint256 extraLoserStake = loser.amountStaked - winner.amountStaked;
            // TODO: unsafe to transfer to the loser directly
            losingStaker.transfer(extraLoserStake);
            loser.amountStaked -= extraLoserStake;
        }

        winner.amountStaked += loser.amountStaked / 2;

        // TODO: deposit extra loser stake into ArbOS

        loser.amountStaked = 0;
        loser.isZombie = true;
        winner.currentChallenge = address(0);
        loser.currentChallenge = address(0);
    }

    function currentRequiredStake() public view returns (uint256) {
        uint256 MAX_INT = 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff;

        uint256 latestConfirmedAge = block.number - nodes[latestConfirmed].deadlineBlock();
        uint256 challengePeriodsPassed = latestConfirmedAge / challengePeriod;
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

    function discardUnresolvedNode() private {
        // node can be discarded
        nodes[firstUnresolvedNode] = Node(0);
        firstUnresolvedNode++;
    }

    function verifyCanStake() private {
        // Verify that sender is not already a staker
        require(!stakers[msg.sender].isStaked);
        require(msg.value >= currentRequiredStake());
    }

    function addStaker(uint256 nodeNum, Node node) private {
        require(!stakers[msg.sender].isStaked, "ALREADY_STAKED");
        stakers[msg.sender] = Staker(nodeNum, msg.value, address(0), false, true);
        node.addStaker(msg.sender);
    }

    function checkValidNodeNumForStake(uint256 nodeNum) private view {
        require(nodeNum >= firstUnresolvedNode && nodeNum <= latestNodeCreated);
    }

    function checkUnchallengedStaker(Staker storage staker) private view {
        require(!staker.isZombie);
        require(staker.currentChallenge == address(0));
    }
}
