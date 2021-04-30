---
title: IBridge.sol Spec
---

### `deliverMessageToInbox(uint8 kind, address sender, bytes32 messageDataHash) → uint256` (external)

### `executeCall(address destAddr, uint256 amount, bytes data) → bool success, bytes returnData` (external)

### `setInbox(address inbox, bool enabled)` (external)

### `setOutbox(address inbox, bool enabled)` (external)

### `activeOutbox() → address` (external)

### `allowedInboxes(address inbox) → bool` (external)

### `allowedOutboxes(address outbox) → bool` (external)

### `inboxAccs(uint256 index) → bytes32` (external)

### `messageCount() → uint256` (external)

### `MessageDelivered(uint256 messageIndex, bytes32 beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash)`
