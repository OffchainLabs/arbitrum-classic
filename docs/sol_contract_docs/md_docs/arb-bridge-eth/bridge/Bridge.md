---
title: Bridge.sol Spec
---

### `initialize()` (external)

### `allowedInboxes(address inbox) → bool` (external)

### `allowedOutboxes(address outbox) → bool` (external)

### `deliverMessageToInbox(uint8 kind, address sender, bytes32 messageDataHash) → uint256` (external)

### `executeCall(address destAddr, uint256 amount, bytes data) → bool success, bytes returnData` (external)

### `setInbox(address inbox, bool enabled)` (external)

### `setOutbox(address outbox, bool enabled)` (external)

### `messageCount() → uint256` (external)
