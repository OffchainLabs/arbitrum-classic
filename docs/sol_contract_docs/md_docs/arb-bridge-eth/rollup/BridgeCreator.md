---
title: BridgeCreator.sol Spec
---

### `updateTemplates(address _delayedBridgeTemplate, address _sequencerInboxTemplate, address _inboxTemplate, address _rollupEventBridgeTemplate, address _outboxTemplate)` (external)

### `createBridge(address adminProxy, address rollup, address sequencer, uint256 sequencerDelayBlocks, uint256 sequencerDelaySeconds) → contract Bridge, contract SequencerInbox, contract Inbox, contract RollupEventBridge, contract Outbox` (external)

### `TemplatesUpdated()`
