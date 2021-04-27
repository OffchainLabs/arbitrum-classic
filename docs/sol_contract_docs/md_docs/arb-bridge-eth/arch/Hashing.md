---
title: Hashing.sol Spec
---

### `keccak1(bytes32 b) → bytes32` (internal)

### `keccak2(bytes32 a, bytes32 b) → bytes32` (internal)

### `bytes32FromArray(bytes arr, uint256 offset) → uint256` (internal)

### `merkleRoot(bytes data, uint256 startOffset, uint256 dataLength, bool pack) → bytes32, bool` (internal)

### `roundUpToPow2(uint256 len) → uint256` (internal)

### `bytesToBufferHash(bytes buf, uint256 startOffset, uint256 length) → bytes32` (internal)

### `hashInt(uint256 val) → bytes32` (internal)

### `hashCodePoint(struct Value.CodePoint cp) → bytes32` (internal)

### `hashTuplePreImage(bytes32 innerHash, uint256 valueSize) → bytes32` (internal)

### `hash(struct Value.Data val) → bytes32` (internal)

### `getTuplePreImage(struct Value.Data[] vals) → struct Value.Data` (internal)
