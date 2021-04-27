---
title: RollupLib.sol Spec
---

### `stateHash(struct RollupLib.ExecutionState execState) → bytes32` (internal)

### `decodeExecutionState(bytes32[3] bytes32Fields, uint256[4] intFields, uint256 proposedBlock, uint256 inboxMaxCount) → struct RollupLib.ExecutionState` (internal)

### `decodeAssertion(bytes32[3][2] bytes32Fields, uint256[4][2] intFields, uint256 beforeProposedBlock, uint256 beforeInboxMaxCount, uint256 inboxMaxCount) → struct RollupLib.Assertion` (internal)

### `executionStateChallengeHash(struct RollupLib.ExecutionState state) → bytes32` (internal)

### `executionHash(struct RollupLib.Assertion assertion) → bytes32` (internal)

### `challengeRoot(struct RollupLib.Assertion assertion, bytes32 assertionExecHash, uint256 blockProposed) → bytes32` (internal)

### `challengeRootHash(bytes32 execution, uint256 proposedTime, uint256 maxMessageCount) → bytes32` (internal)

### `confirmHash(struct RollupLib.Assertion assertion) → bytes32` (internal)

### `confirmHash(bytes32 beforeSendAcc, bytes32 afterSendAcc, bytes32 afterLogAcc, uint256 afterSendCount, uint256 afterLogCount) → bytes32` (internal)

### `feedAccumulator(bytes messageData, uint256[] messageLengths, bytes32 beforeAcc) → bytes32` (internal)

### `nodeHash(bool hasSibling, bytes32 lastHash, bytes32 assertionExecHash, bytes32 inboxAcc) → bytes32` (internal)

### `nodeAccumulator(bytes32 prevAcc, bytes32 newNodeHash) → bytes32` (internal)
