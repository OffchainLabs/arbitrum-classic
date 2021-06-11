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

### `isZombie(address staker) → bool` (public)

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

### `getNodeHash(uint256 index) → bytes32` (public)

**Returns**: Node: hash as of this node number
