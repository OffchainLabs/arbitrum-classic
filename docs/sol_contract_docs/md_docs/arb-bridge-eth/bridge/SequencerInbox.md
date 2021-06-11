---
title: SequencerInbox.sol Spec
---

### `initialize(contract IBridge _delayedInbox, address _sequencer, address _rollup)` (external)

### `setSequencer(address newSequencer)` (external)

### `maxDelayBlocks() → uint256` (public)

### `maxDelaySeconds() → uint256` (public)

### `forceInclusion(uint256 _totalDelayedMessagesRead, uint8 kind, uint256[2] l1BlockAndTimestamp, uint256 inboxSeqNum, uint256 gasPriceL1, address sender, bytes32 messageDataHash, bytes32 delayedAcc)` (external)

### `addSequencerL2BatchFromOrigin(bytes transactions, uint256[] lengths, uint256[] sectionsMetadata, bytes32 afterAcc)` (external)

### `addSequencerL2Batch(bytes transactions, uint256[] lengths, uint256[] sectionsMetadata, bytes32 afterAcc)` (external)

### `proveBatchContainsSequenceNumber(bytes proof, uint256 inboxCount) → uint256, bytes32` (external)

### `getInboxAccsLength() → uint256` (external)
