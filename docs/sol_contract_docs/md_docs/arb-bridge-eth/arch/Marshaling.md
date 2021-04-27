---
title: Marshaling.sol Spec
---

### `deserializeMessage(bytes data, uint256 startOffset) → bool, uint256, address, uint8, bytes` (internal)

### `deserializeRawMessage(bytes data, uint256 startOffset) → bool, uint256, bytes` (internal)

### `deserializeHashPreImage(bytes data, uint256 startOffset) → uint256 offset, struct Value.Data value` (internal)

### `deserializeInt(bytes data, uint256 startOffset) → uint256, uint256` (internal)

### `deserializeBytes32(bytes data, uint256 startOffset) → uint256, bytes32` (internal)

### `deserializeCodePoint(bytes data, uint256 startOffset) → uint256, struct Value.Data` (internal)

### `deserializeTuple(uint8 memberCount, bytes data, uint256 startOffset) → uint256, struct Value.Data[]` (internal)

### `deserialize(bytes data, uint256 startOffset) → uint256, struct Value.Data` (internal)
