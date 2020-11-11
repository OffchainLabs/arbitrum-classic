// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.5.17;

import "./Node.sol";

contract Rollup {
    struct Staker {
        uint256 latestStakedNode;
        uint256 amountStaked;
        // currentChallenge is 0 if staker is not in a challenge
        address currentChallenge;
        bool isZombie;
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

        nodes[firstUnresolvedNode].confirmInvalid();

        discardUnresolvedNode();
    }

    // If the node previous to this one is not the latest confirmed, we can reject immediately
    function rejectNextNodeOutOfOrder() external {
        nodes[firstUnresolvedNode].confirmOutOfOrder(latestConfirmed);

        discardUnresolvedNode();
    }

    function confirmNextNode() external {
        // No stake has been placed during the last challengePeriod blocks
        require(block.number - lastStakeBlock >= challengePeriod);

        nodes[firstUnresolvedNode].confirmValid(stakerCount, latestConfirmed);
        discardUnresolvedNode();
    }

    function newStakeOnExistingNode(uint256 nodeNum) external payable {
        verifyCanStake();
        Node node = nodes[nodeNum];
        require(node.prev() == latestConfirmed);
        addStaker(nodeNum, node);
    }

    function newStakeOnNewNode(
        uint256 nodeNum,
        uint256 prev /* assertion data */
    ) external payable {
        verifyCanStake();
        require(nodeNum == latestNodeCreated + 1);
        require(prev == latestConfirmed);

        // TODO: Verify that the preconditions of assertion are consistent with the postconditions of prev
        // TODO: Verify that assertion meets the minimum size requirement
        // TODO: Verify that assertion meets the minimum Delta time requirement
        Node node = new Node(
            0, // TODO: assertion hash
            latestConfirmed,
            block.number,
            0, // TODO: deadline block
            1
        );
        addStaker(nodeNum, node);
        nodes[nodeNum] = node;
        latestNodeCreated++;
    }

    function addStakeOnExistingNode(uint256 nodeNum) external {
        Staker storage staker = stakers[msg.sender];
        require(!staker.isZombie);
        Node node = nodes[nodeNum];
        require(staker.latestStakedNode == node.prev());
        node.addStaker(msg.sender);
        staker.latestStakedNode = nodeNum;
    }

    function addStakeOnNewNode(uint256 nodeNum) external {
        require(nodeNum == latestNodeCreated + 1);
        Staker storage staker = stakers[msg.sender];
        require(!staker.isZombie);

        // TODO: Verify that the preconditions of assertion are consistent with the postconditions of prev
        // TODO: Verify that assertion meets the minimum size requirement
        // TODO: Verify that assertion meets the minimum Delta time requirement

        Node node = new Node(
            0, // TODO: assertion hash
            staker.latestStakedNode,
            block.number,
            0, // TODO: deadline block
            1
        );
        node.addStaker(msg.sender);
        nodes[nodeNum] = node;
        staker.latestStakedNode = nodeNum;
    }

    function returnOldDeposit(address payable stakerAddress) external {
        Staker memory staker = stakers[stakerAddress];
        require(!staker.isZombie);
        require(staker.latestStakedNode <= latestConfirmed);
        require(staker.currentChallenge == address(0));

        delete stakers[stakerAddress];
        // Staker could force transfer to revert. We may want to allow funds to be withdrawn separately
        stakerAddress.transfer(staker.amountStaked);
    }

    function addToDeposit() external payable {
        Staker memory staker = stakers[msg.sender];
        require(!staker.isZombie);
        require(staker.currentChallenge == address(0));
        staker.amountStaked += msg.value;
    }

    function reduceDeposit(uint256 maxReduction) external {
        Staker memory staker = stakers[msg.sender];
        require(!staker.isZombie);
        require(staker.currentChallenge == address(0));
        uint256 currentRequired = currentRequiredStake();
        require(staker.amountStaked > currentRequired);
        uint256 withdrawAmount = staker.amountStaked - currentRequired;
        // Cap withdrawAmount at maxReduction
        if (withdrawAmount > maxReduction) {
            withdrawAmount = maxReduction;
        }
        msg.sender.transfer(withdrawAmount);
    }

    function removeZombieStake(uint256 nodeNum, address stakerAddress) external {
        require(stakers[stakerAddress].isZombie);
        require(nodeNum >= firstUnresolvedNode);
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

        require(!staker1.isZombie);
        require(staker1.currentChallenge == address(0));
        require(node1.stakers(staker1Address));

        require(!staker2.isZombie);
        require(staker2.currentChallenge == address(0));
        require(node2.stakers(staker2Address));

        require(node1.prev() == node2.prev());
        require(latestConfirmed < nodeNum1);
        require(nodeNum1 < nodeNum2);

        // Start a challenge between staker1 and staker2. Staker1 will defend the correctness of node1, and staker2 will challenge it.
        // TODO: How to we want to handle the two challenge types

        // TODO: Actually launch challenge
        address challengeAddress = address(0);
        staker1.currentChallenge = challengeAddress;
        staker2.currentChallenge = challengeAddress;
    }

    function completeChallenge(address winningStaker, address losingStaker) external {
        Staker storage winner = stakers[winningStaker];
        Staker storage loser = stakers[losingStaker];

        // Only the challenge contract can declare winners and losers
        require(winner.currentChallenge == msg.sender);
        require(loser.currentChallenge == msg.sender);

        uint256 winnerPrize = 0;
        if (winner.amountStaked > loser.amountStaked) {
            winner.amountStaked += loser.amountStaked / 2;
        } else {
            winner.amountStaked += winner.amountStaked / 2;

            winnerPrize = winner.amountStaked / 2;
        }

        loser.amountStaked = 0;
        loser.isZombie = true;
    }

    function currentRequiredStake() public view returns (uint256) {
        uint256 latestConfirmedAge = block.number - nodes[latestConfirmed].deadlineBlock();
        uint256 challengePeriodsPassed = latestConfirmedAge / challengePeriod;
        if (challengePeriodsPassed > 255) {
            challengePeriodsPassed = 255;
        }
        uint256 multiplier = 2**challengePeriodsPassed - 1;
        if (multiplier == 0) {
            multiplier = 1;
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
        // TODO: Can we assume that all stakers are staked on the latest confirmed?
        require(!nodes[latestConfirmed].stakers(msg.sender));
        require(msg.value >= currentRequiredStake());
    }

    function addStaker(uint256 nodeNum, Node node) private {
        require(stakers[msg.sender].latestStakedNode == 0, "ALREADY_STAKED");
        stakers[msg.sender] = Staker(nodeNum, msg.value, address(0), false);
        node.addStaker(msg.sender);
    }
}
