---
title: IRollup.sol Spec
---

### `initialize(bytes32 _machineHash, uint256 _confirmPeriodBlocks, uint256 _extraChallengeTimeBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, bytes _extraConfig, address[6] connectedContracts)` (external)

### `completeChallenge(address winningStaker, address losingStaker)` (external)

### `returnOldDeposit(address stakerAddress)` (external)

### `RollupCreated(bytes32 machineHash)`

### `NodeCreated(uint256 nodeNum, bytes32 parentNodeHash, bytes32 nodeHash, bytes32 executionHash, uint256 inboxMaxCount, bytes32 afterInboxAcc, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields)`

### `NodeConfirmed(uint256 nodeNum, bytes32 afterSendAcc, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount)`

### `NodeRejected(uint256 nodeNum)`

### `RollupChallengeStarted(address challengeContract, address asserter, address challenger, uint256 challengedNode)`
