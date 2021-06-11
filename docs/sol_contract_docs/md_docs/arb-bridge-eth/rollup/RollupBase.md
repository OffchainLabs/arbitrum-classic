---
title: RollupBase.sol Spec
---

### `RollupCreated(bytes32 machineHash)`

### `NodeCreated(uint256 nodeNum, bytes32 parentNodeHash, bytes32 nodeHash, bytes32 executionHash, uint256 inboxMaxCount, uint256 afterInboxBatchEndCount, bytes32 afterInboxBatchAcc, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields)`

### `NodeConfirmed(uint256 nodeNum, bytes32 afterSendAcc, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount)`

### `NodeRejected(uint256 nodeNum)`

### `RollupChallengeStarted(address challengeContract, address asserter, address challenger, uint256 challengedNode)`

### `StakerReassigned(address staker, uint256 newNode)`

### `NodesDestroyed(uint256 startNode, uint256 endNode)`

### `OwnerFunctionCalled(uint256 id)`
