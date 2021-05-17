---
title: Node.sol Spec
---

### `onlyRollup()`

### `initialize(address _rollup, bytes32 _stateHash, bytes32 _challengeHash, bytes32 _confirmData, uint256 _prev, uint256 _deadlineBlock)` (external)

Mark the given staker as staked on this node

- `_rollup`: Initial value of rollup

- `_stateHash`: Initial value of stateHash

- `_challengeHash`: Initial value of challengeHash

- `_confirmData`: Initial value of confirmData

- `_prev`: Initial value of prev

- `_deadlineBlock`: Initial value of deadlineBlock

### `destroy()` (external)

Destroy this node

### `addStaker(address staker) → uint256` (external)

Mark the given staker as staked on this node

- `staker`: Address of the staker to mark

**Returns**: The: number of stakers after adding this one

### `removeStaker(address staker)` (external)

Remove the given staker from this node

- `staker`: Address of the staker to remove

### `childCreated(uint256 number)` (external)

### `newChildConfirmDeadline(uint256 deadline)` (external)

### `requirePastDeadline()` (external)

Check whether the current block number has met or passed the node's deadline

### `requirePastChildConfirmDeadline()` (external)

Check whether the current block number has met or passed deadline for children of this node to be confirmed
