---
title: IOutbox.sol Spec
---

### `l2ToL1Sender() → address` (external)

### `l2ToL1Block() → uint256` (external)

### `l2ToL1EthBlock() → uint256` (external)

### `l2ToL1Timestamp() → uint256` (external)

### `processOutgoingMessages(bytes sendsData, uint256[] sendLengths)` (external)

### `OutboxEntryCreated(uint256 batchNum, uint256 outboxIndex, bytes32 outputRoot, uint256 numInBatch)`

### `OutBoxTransactionExecuted(address destAddr, address l2Sender, uint256 outboxIndex, uint256 transactionIndex)`
