---
title: RollupCore.sol Spec
---

### `getNode(uint256 nodeNum) → contract INode` (public)

Get the address of the Node contract for the given node

- `nodeNum`: Index of the node

**Returns**: Address: of the Node contract

### `getStakerAddress(uint256 stakerNum) → address` (public)

Get the address of the staker at the given index

- `stakerNum`: Index of the staker

**Returns**: Address: of the staker

### `isStaked(address staker) → bool` (public)

Check whether the given staker is staked

- `staker`: Staker address to check

**Returns**: True: or False for whether the staker was staked

### `latestStakedNode(address staker) → uint256` (public)

Get the latest staked node of the given staker

- `staker`: Staker address to lookup

**Returns**: Latest: node staked of the staker

### `currentChallenge(address staker) → address` (public)

Get the current challenge of the given staker

- `staker`: Staker address to lookup

**Returns**: Current: challenge of the staker

### `amountStaked(address staker) → uint256` (public)

Get the amount staked of the given staker

- `staker`: Staker address to lookup

**Returns**: Amount: staked of the staker

### `zombieAddress(uint256 zombieNum) → address` (public)

Get the original staker address of the zombie at the given index

- `zombieNum`: Index of the zombie to lookup

**Returns**: Original: staker address of the zombie

### `zombieLatestStakedNode(uint256 zombieNum) → uint256` (public)

Get Latest node that the given zombie at the given index is staked on

- `zombieNum`: Index of the zombie to lookup

**Returns**: Latest: node that the given zombie is staked on

### `zombieCount() → uint256` (public)

**Returns**: Current: number of un-removed zombies

### `withdrawableFunds(address owner) → uint256` (public)

Get the amount of funds withdrawable by the given address

- `owner`: Address to check the funds of

**Returns**: Amount: of funds withdrawable by owner

### `firstUnresolvedNode() → uint256` (public)

If all nodes have been resolved, this will be latestNodeCreated + 1

**Returns**: Index: of the first unresolved node

### `latestConfirmed() → uint256` (public)

**Returns**: Index: of the latest confirmed node

### `latestNodeCreated() → uint256` (public)

**Returns**: Index: of the latest rollup node created

### `lastStakeBlock() → uint256` (public)

**Returns**: Ethereum: block that the most recent stake was created

### `stakerCount() → uint256` (public)

**Returns**: Number: of active stakers currently staked

### `initializeCore(contract INode initialNode)` (internal)

Initialize the core with an initial node

- `initialNode`: Initial node to start the chain with

### `nodeCreated(contract INode node, bytes32 nodeHash)` (internal)

React to a new node being created by storing it an incrementing the latest node counter

- `node`: Node that was newly created

- `nodeHash`: The hash of said node

### `getNodeHash(uint256 index) → bytes32` (public)

**Returns**: Node: hash as of this node number

### `updateLatestNodeCreated(uint256 newLatestNodeCreated)` (internal)

Update the latest node created

- `newLatestNodeCreated`: New value for the latest node created

### `rejectNextNode()` (internal)

Reject the next unresolved node

### `confirmNextNode()` (internal)

Confirm the next unresolved node

### `createNewStake(address payable stakerAddress, uint256 depositAmount)` (internal)

Create a new stake

- `stakerAddress`: Address of the new staker

- `depositAmount`: Stake amount of the new staker

### `inChallenge(address stakerAddress1, address stakerAddress2) → address` (internal)

Check to see whether the two stakers are in the same challenge

- `stakerAddress1`: Address of the first staker

- `stakerAddress2`: Address of the second staker

**Returns**: Address: of the challenge that the two stakers are in

### `clearChallenge(address stakerAddress)` (internal)

Make the given staker as not being in a challenge

- `stakerAddress`: Address of the staker to remove from a challenge

### `challengeStarted(address staker1, address staker2, address challenge)` (internal)

Mark both the given stakers as engaged in the challenge

- `staker1`: Address of the first staker

- `staker2`: Address of the second staker

- `challenge`: Address of the challenge both stakers are now in

### `increaseStakeBy(address stakerAddress, uint256 amountAdded)` (internal)

Add to the stake of the given staker by the given amount

- `stakerAddress`: Address of the staker to increase the stake of

- `amountAdded`: Amount of stake to add to the staker

### `reduceStakeTo(address stakerAddress, uint256 target) → uint256` (internal)

Reduce the stake of the given staker to the given target

- `stakerAddress`: Address of the staker to reduce the stake of

- `target`: Amount of stake to leave with the staker

**Returns**: Amount: of value released from the stake

### `turnIntoZombie(address stakerAddress)` (internal)

Remove the given staker and turn them into a zombie

- `stakerAddress`: Address of the staker to remove

### `zombieUpdateLatestStakedNode(uint256 zombieNum, uint256 latest)` (internal)

Update the latest staked node of the zombie at the given index

- `zombieNum`: Index of the zombie to move

- `latest`: New latest node the zombie is staked on

### `removeZombie(uint256 zombieNum)` (internal)

Remove the zombie at the given index

- `zombieNum`: Index of the zombie to remove

### `withdrawStaker(address stakerAddress)` (internal)

Remove the given staker and return their stake

- `stakerAddress`: Address of the staker withdrawing their stake

### `stakerUpdateLatestStakedNode(address stakerAddress, uint256 latest)` (internal)

Update the latest staked node of the staker at the given addresss

- `stakerAddress`: Address of the staker to move

- `latest`: New latest node the staker is staked on

### `stakeOnNode(address stakerAddress, uint256 nodeNum, uint256 confirmPeriodBlocks) → uint256` (internal)

Advance the given staker to the given node

- `stakerAddress`: Address of the staker adding their stake

- `nodeNum`: Index of the node to stake on

### `withdrawFunds(address owner) → uint256` (internal)

Clear the withdrawable funds for the given address

- `owner`: Address of the account to remove funds from

**Returns**: Amount: of funds removed from account

### `max(uint256 a, uint256 b) → uint256` (internal)
