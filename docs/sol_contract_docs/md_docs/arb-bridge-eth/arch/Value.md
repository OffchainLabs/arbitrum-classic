---
title: Value.sol Spec
---

### `tupleTypeCode() → uint8` (internal)

### `tuplePreImageTypeCode() → uint8` (internal)

### `intTypeCode() → uint8` (internal)

### `bufferTypeCode() → uint8` (internal)

### `codePointTypeCode() → uint8` (internal)

### `valueTypeCode() → uint8` (internal)

### `hashOnlyTypeCode() → uint8` (internal)

### `isValidTupleSize(uint256 size) → bool` (internal)

### `typeCodeVal(struct Value.Data val) → struct Value.Data` (internal)

### `valLength(struct Value.Data val) → uint8` (internal)

### `isInt(struct Value.Data val) → bool` (internal)

### `isInt64(struct Value.Data val) → bool` (internal)

### `isCodePoint(struct Value.Data val) → bool` (internal)

### `isTuple(struct Value.Data val) → bool` (internal)

### `isBuffer(struct Value.Data val) → bool` (internal)

### `newEmptyTuple() → struct Value.Data` (internal)

### `newBoolean(bool val) → struct Value.Data` (internal)

### `newInt(uint256 _val) → struct Value.Data` (internal)

### `newHashedValue(bytes32 valueHash, uint256 valueSize) → struct Value.Data` (internal)

### `newTuple(struct Value.Data[] _val) → struct Value.Data` (internal)

### `newTuplePreImage(bytes32 preImageHash, uint256 size) → struct Value.Data` (internal)

### `newCodePoint(uint8 opCode, bytes32 nextHash) → struct Value.Data` (internal)

### `newCodePoint(uint8 opCode, bytes32 nextHash, struct Value.Data immediate) → struct Value.Data` (internal)

### `newBuffer(bytes32 bufHash) → struct Value.Data` (internal)
