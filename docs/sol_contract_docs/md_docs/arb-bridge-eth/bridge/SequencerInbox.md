---
title: SequencerInbox.sol Spec
---

### `initialize(contract IBridge _delayedInbox, address _sequencer, uint256 _maxDelayBlocks, uint256 _maxDelaySeconds)` (external)

### `forceInclusion(uint256 _totalDelayedMessagesRead, uint8 kind, uint256[2] l1BlockAndTimestamp, uint256 inboxSeqNum, uint256 gasPriceL1, address sender, bytes32 messageDataHash)` (external)

### `addSequencerL2BatchFromOrigin(bytes transactions, uint256[] lengths, uint256 l1BlockNumber, uint256 timestamp, uint256 _totalDelayedMessagesRead, bytes32 afterAcc)` (external)

### `addSequencerL2Batch(bytes transactions, uint256[] lengths, uint256 l1BlockNumber, uint256 timestamp, uint256 _totalDelayedMessagesRead, bytes32 afterAcc)` (external)

### `proveBatchContainsSequenceNumber(bytes proof, uint256 inboxCount) → uint256, bytes32` (external)
