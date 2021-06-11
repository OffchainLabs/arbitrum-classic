---
title: ISequencerInbox.sol Spec
---

### `setSequencer(address newSequencer)` (external)

### `messageCount() → uint256` (external)

### `maxDelayBlocks() → uint256` (external)

### `maxDelaySeconds() → uint256` (external)

### `inboxAccs(uint256 index) → bytes32` (external)

### `getInboxAccsLength() → uint256` (external)

### `proveBatchContainsSequenceNumber(bytes proof, uint256 inboxCount) → uint256, bytes32` (external)

### `SequencerBatchDelivered(uint256 firstMessageNum, bytes32 beforeAcc, uint256 newMessageCount, bytes32 afterAcc, bytes transactions, uint256[] lengths, uint256[] sectionsMetadata, uint256 seqBatchIndex, address sequencer)`

### `SequencerBatchDeliveredFromOrigin(uint256 firstMessageNum, bytes32 beforeAcc, uint256 newMessageCount, bytes32 afterAcc, uint256 seqBatchIndex)`

### `DelayedInboxForced(uint256 firstMessageNum, bytes32 beforeAcc, uint256 newMessageCount, uint256 totalDelayedMessagesRead, bytes32[2] afterAccAndDelayed, uint256 seqBatchIndex)`

### `SequencerAddressUpdated(address newAddress)`
