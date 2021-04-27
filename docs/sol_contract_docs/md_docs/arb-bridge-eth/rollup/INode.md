---
title: INode.sol Spec
---

### `initialize(address _rollup, bytes32 _stateHash, bytes32 _challengeHash, bytes32 _confirmData, uint256 _prev, uint256 _deadlineBlock)` (external)

### `destroy()` (external)

### `addStaker(address staker) → uint256` (external)

### `removeStaker(address staker)` (external)

### `childCreated(uint256)` (external)

### `newChildConfirmDeadline(uint256 deadline)` (external)

### `stateHash() → bytes32` (external)

### `challengeHash() → bytes32` (external)

### `confirmData() → bytes32` (external)

### `prev() → uint256` (external)

### `deadlineBlock() → uint256` (external)

### `noChildConfirmedBeforeBlock() → uint256` (external)

### `stakerCount() → uint256` (external)

### `stakers(address staker) → bool` (external)

### `firstChildBlock() → uint256` (external)

### `latestChildNumber() → uint256` (external)

### `requirePastDeadline()` (external)

### `requirePastChildConfirmDeadline()` (external)
