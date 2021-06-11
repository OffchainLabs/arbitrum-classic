---
title: RollupEventBridge.sol Spec
---

### `onlyRollup()`

### `initialize(address _bridge, address _rollup)` (external)

### `rollupInitialized(uint256 confirmPeriodBlocks, uint256 extraChallengeTimeBlocks, uint256 arbGasSpeedLimitPerBlock, uint256 baseStake, address stakeToken, address owner, bytes extraConfig)` (external)

### `nodeCreated(uint256 nodeNum, uint256 prev, uint256 deadline, address asserter)` (external)

### `nodeConfirmed(uint256 nodeNum)` (external)

### `nodeRejected(uint256 nodeNum)` (external)

### `stakeCreated(address staker, uint256 nodeNum)` (external)

### `claimNode(uint256 nodeNum, address staker)` (external)
