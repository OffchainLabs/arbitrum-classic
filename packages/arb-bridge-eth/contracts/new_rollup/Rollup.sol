// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.5.17;

contract Rollup {
    struct Node {
        bytes32 assertionHash;
        uint256 prev;
        uint256 proposedBlock;
        uint256 deadlineBlock;
        uint256 stakerCount;
        mapping(address => bool) stakers;
    }

    struct Staker {
        uint256 latestStakedNode;
        uint256 amountStaked;
        // currentChallenge is 0 if staker is not in a challenge
        address currentChallenge;
    }

    uint256 latestConfirmed;
    uint256 firstUnresolvedNode;
    uint256 highestNode;
    Node[] nodes;
    uint256 lastStakeBlock;
    uint256 stakerCount;
    mapping(address => Staker) stakers;

    uint256 baseStake;
    uint256 challengePeriod;

    function rejectNextNode(uint256 successorWithStake, address staker) external {
        verifyRejectable(successorWithStake, staker);
        discardUnresolvedNode();
    }

    function confirmNextNode() external {
        Node storage node = nodes[firstUnresolvedNode];
        // Verify the blocks deadline has passed
        require(node.deadlineBlock <= block.number);

        // All non-zombie stakers are staked on this node, and no zombie stakers are staked here
        require(stakerCount == node.stakerCount);
        // There is at least one non-zombie staker
        require(stakerCount > 0);

        // No stake has been placed during the last challengePeriod blocks
        require(block.number - lastStakeBlock >= challengePeriod);

        discardUnresolvedNode();
    }

    function newStakeOnExistingNode(uint256 nodeNum) external payable {
        verifyCanStake();
        Node storage node = nodes[nodeNum];
        require(node.prev == latestConfirmed);
        addStaker(nodeNum, node);
    }

    function newStakeOnNewNode(
        uint256 nodeNum,
        uint256 prev /* assertion data */
    ) external payable {
        verifyCanStake();
        require(nodeNum == highestNode + 1);
        require(prev == latestConfirmed);

        // TODO: Verify that the preconditions of assertion are consistent with the postconditions of prev
        // TODO: Verify that assertion meets the minimum size requirement
        // TODO: Verify that assertion meets the minimum Delta time requirement
        nodes[nodeNum] = Node(
            0, // TODO: assertion hash
            latestConfirmed,
            block.number,
            0, // TODO: deadline block
            1
        );
        Node storage node = nodes[nodeNum];
        addStaker(nodeNum, node);
        highestNode++;
    }

    function addStakeOnExistingNode(uint256 nodeNum) external {
        Staker storage staker = stakers[msg.sender];
        // TODO: Verify that caller is a non-zombie staker
        Node storage node = nodes[nodeNum];
        require(staker.latestStakedNode == node.prev);
        node.stakers[msg.sender] = true;
        staker.latestStakedNode = nodeNum;
    }

    function addStakeOnNewNode(uint256 nodeNum) external {
        require(nodeNum == highestNode + 1);
        Staker storage staker = stakers[msg.sender];
        // TODO: Verify that caller is a non-zombie staker

        // TODO: Verify that the preconditions of assertion are consistent with the postconditions of prev
        // TODO: Verify that assertion meets the minimum size requirement
        // TODO: Verify that assertion meets the minimum Delta time requirement

        nodes[nodeNum] = Node(
            0, // TODO: assertion hash
            staker.latestStakedNode,
            block.number,
            0, // TODO: deadline block
            1
        );
        Node storage node = nodes[nodeNum];
        node.stakers[msg.sender] = true;
        staker.latestStakedNode = nodeNum;
    }

    function returnOldDeposit(address payable stakerAddress) external {
        Staker memory staker = stakers[stakerAddress];
        // Verify that staker is a non-zombie staker
        require(staker.latestStakedNode <= latestConfirmed);
        require(staker.currentChallenge == address(0));

        delete stakers[stakerAddress];
        // Staker could force transfer to revert. We may want to allow funds to be withdrawn separately
        stakerAddress.transfer(staker.amountStaked);
    }

    function addToDeposit() external payable {
        Staker memory staker = stakers[msg.sender];
        // Verify that staker is a non-zombie staker
        require(staker.currentChallenge == address(0));
        staker.amountStaked += msg.value;
    }

    function reduceDeposit(uint256 maxReduction) external {
        Staker memory staker = stakers[msg.sender];
        // Verify that staker is a non-zombie staker
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
        // TODO: Verify that zombieStaker is a zombie
        require(nodeNum >= firstUnresolvedNode);
        Node storage node = nodes[nodeNum];
        node.stakers[stakerAddress] = false;
    }

    function createChallenge(
        address staker1Address,
        uint256 nodeNum1,
        address staker2Address,
        uint256 nodeNum2
    ) external {
        Staker storage staker1 = stakers[staker1Address];
        Staker storage staker2 = stakers[staker2Address];
        Node storage node1 = nodes[nodeNum1];
        Node storage node2 = nodes[nodeNum2];
        // TODO: Verify that staker1 is not a zombie
        // TODO: Verify that staker2 is not a zombie
        require(staker1.currentChallenge == address(0));
        require(node1.stakers[staker1Address]);
        require(staker2.currentChallenge == address(0));
        require(node2.stakers[staker2Address]);

        require(node1.prev == node2.prev);
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
    }

    function currentRequiredStake() public view returns (uint256) {
        uint256 latestConfirmedAge = block.number - nodes[latestConfirmed].deadlineBlock;
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
        firstUnresolvedNode++;
    }

    function verifyRejectable(uint256 successorWithStake, address stakerAddress) private view {
        Node storage node = nodes[firstUnresolvedNode];

        // If the node previous to this one is not the latest confirmed, we can reject immediately
        if (node.prev != latestConfirmed) {
            return;
        }

        // Verify the blocks deadline has passed
        require(node.deadlineBlock <= block.number);

        // No stake has been placed during the last challengePeriod blocks
        require(block.number - lastStakeBlock >= challengePeriod);

        // Verify that no staker is staked on this node
        require(node.stakerCount == 0);

        Staker storage staker = stakers[stakerAddress];
        // TODO: verify staker is not a zombie

        Node storage stakedSiblingNode = nodes[successorWithStake];
        // stakedSiblingNode is a child of latestConfirmed
        require(stakedSiblingNode.prev == latestConfirmed);
        // staker is actually staked on stakedSiblingNode
        require(stakedSiblingNode.stakers[stakerAddress]);
    }

    function verifyCanStake() private {
        // Verify that sender is not already a staker
        // TODO: Can we assume that all stakers are staked on the latest confirmed?
        require(!nodes[latestConfirmed].stakers[msg.sender]);
        require(msg.value >= currentRequiredStake());
    }

    function addStaker(uint256 nodeNum, Node storage node) private {
        stakers[msg.sender] = Staker(nodeNum, msg.value, address(0));
        node.stakers[msg.sender] = true;
    }
}
