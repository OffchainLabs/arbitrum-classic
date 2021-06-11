---
title: ISequencerInbox.sol Spec
---

### `messageCount() → uint256` (external)

### `maxDelayBlocks() → uint256` (external)

### `maxDelaySeconds() → uint256` (external)

### `inboxAccs(uint256 index) → bytes32` (external)

### `proveBatchContainsSequenceNumber(bytes proof, uint256 inboxCount) → uint256, bytes32` (external)

### `SequencerBatchDelivered(uint256 firstMessageNum, bytes32 beforeAcc, uint256 newMessageCount, bytes32 afterAcc, bytes transactions, uint256[] lengths, uint256 l1BlockNumber, uint256 timestamp, uint256 totalDelayedMessagesRead, bytes32 delayedAcc, uint256 seqBatchIndex)`

### `SequencerBatchDeliveredFromOrigin(uint256 firstMessageNum, bytes32 beforeAcc, uint256 newMessageCount, bytes32 afterAcc, bytes32 delayedAcc, uint256 seqBatchIndex)`

### `DelayedInboxForced(uint256 firstMessageNum, bytes32 beforeAcc, uint256 newMessageCount, uint256 totalDelayedMessagesRead, bytes32[2] afterAccAndDelayed, uint256 seqBatchIndex)`
