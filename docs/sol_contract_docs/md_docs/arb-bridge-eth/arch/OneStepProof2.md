---
title: OneStepProof2.sol Spec
---

### `makeZeros() → bytes32[]` (internal)

### `keccak1(bytes32 b) → bytes32` (internal)

### `keccak2(bytes32 a, bytes32 b) → bytes32` (internal)

### `get(bytes32 buf, uint256 loc, bytes32[] proof) → bytes32` (internal)

### `checkSize(bytes32 buf, uint256 loc, bytes32[] proof) → bool` (internal)

### `calcHeight(uint256 loc) → uint256` (internal)

### `set(bytes32 buf, uint256 loc, bytes32 v, bytes32[] proof, uint256 nh, bytes32 normal1, bytes32 normal2) → bytes32` (internal)

### `getByte(bytes32 word, uint256 num) → uint256` (internal)

### `setByte(bytes32 word, uint256 num, uint256 b) → bytes32` (internal)

### `setByte(bytes32 word, uint256 num, bytes1 b) → bytes32` (internal)

### `decode(bytes arr, bytes1 _start, bytes1 _end) → bytes32[]` (internal)

### `decodeProof(bytes proof) → struct OneStepProof2.BufferProof` (internal)

### `bytes32FromArray(bytes arr) → uint256` (internal)

### `bytes32FromArray(bytes arr, uint256 offset) → uint256` (internal)

### `bytes32ToArray(bytes32 b) → bytes` (internal)

### `getBuffer8(bytes32 buf, uint256 offset, struct OneStepProof2.BufferProof proof) → uint256` (internal)

### `checkBufferSize(bytes32 buf, uint256 offset, struct OneStepProof2.BufferProof proof) → bool` (internal)

### `getBuffer64(bytes32 buf, uint256 offset, struct OneStepProof2.BufferProof proof) → uint256` (internal)

### `getBuffer256(bytes32 buf, uint256 offset, struct OneStepProof2.BufferProof proof) → uint256` (internal)

### `set(bytes32 buf, uint256 loc, bytes32 v, bytes32[] proof, bytes32[] nproof) → bytes32` (internal)

### `setBuffer8(bytes32 buf, uint256 offset, uint256 b, struct OneStepProof2.BufferProof proof) → bytes32` (internal)

### `setBuffer64(bytes32 buf, uint256 offset, uint256 val, struct OneStepProof2.BufferProof proof) → bytes32` (internal)

### `parseProof(bytes proof) → bytes32[], bytes32[], bytes32[], bytes32[]` (public)

### `setBuffer256(bytes32 buf, uint256 offset, uint256 val, struct OneStepProof2.BufferProof proof) → bytes32` (internal)

### `executeSendInsn(struct OneStepProofCommon.AssertionContext context)` (internal)

### `executeGetBuffer8(struct OneStepProofCommon.AssertionContext context)` (internal)

### `executeGetBuffer64(struct OneStepProofCommon.AssertionContext context)` (internal)

### `executeGetBuffer256(struct OneStepProofCommon.AssertionContext context)` (internal)

### `executeSetBuffer8(struct OneStepProofCommon.AssertionContext context)` (internal)

### `executeSetBuffer64(struct OneStepProofCommon.AssertionContext context)` (internal)

### `executeSetBuffer256(struct OneStepProofCommon.AssertionContext context)` (internal)

### `opInfo(uint256 opCode) → uint256, uint256, uint64, function (struct OneStepProofCommon.AssertionContext) view` (internal)
