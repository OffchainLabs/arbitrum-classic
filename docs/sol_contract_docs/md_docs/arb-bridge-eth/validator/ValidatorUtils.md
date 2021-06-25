---
title: ValidatorUtils.sol Spec
id: ValidatorUtils
---

### `getConfig(contract Rollup rollup) → uint256 confirmPeriodBlocks, uint256 extraChallengeTimeBlocks, uint256 arbGasSpeedLimitPerBlock, uint256 baseStake` (external)

### `stakerInfo(contract Rollup rollup, address stakerAddress) → bool isStaked, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge` (external)

### `findStakerConflict(contract Rollup rollup, address staker1, address staker2, uint256 maxDepth) → enum ValidatorUtils.NodeConflict, uint256, uint256` (external)

### `checkDecidableNextNode(contract Rollup rollup) → enum ValidatorUtils.ConfirmType` (external)

### `requireRejectable(contract Rollup rollup) → bool` (external)

### `requireConfirmable(contract Rollup rollup)` (external)

### `refundableStakers(contract Rollup rollup) → address[]` (external)

### `latestStaked(contract Rollup rollup, address staker) → uint256, bytes32` (external)

### `stakedNodes(contract Rollup rollup, address staker) → uint256[]` (external)

### `findNodeConflict(contract Rollup rollup, uint256 node1, uint256 node2, uint256 maxDepth) → enum ValidatorUtils.NodeConflict, uint256, uint256` (public)

### `getStakers(contract Rollup rollup, uint256 startIndex, uint256 max) → address[], bool hasMore` (public)

### `timedOutChallenges(contract Rollup rollup, uint256 startIndex, uint256 max) → contract IChallenge[], bool hasMore` (external)

### `areUnresolvedNodesLinear(contract Rollup rollup) → bool` (external)
