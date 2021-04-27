---
title: ChallengeLib.sol Spec
---

### `firstSegmentSize(uint256 totalCount, uint256 bisectionCount) → uint256` (internal)

### `otherSegmentSize(uint256 totalCount, uint256 bisectionCount) → uint256` (internal)

### `bisectionChunkHash(uint256 _segmentStart, uint256 _segmentLength, bytes32 _startHash, bytes32 _endHash) → bytes32` (internal)

### `inboxDeltaHash(bytes32 _inboxAcc, bytes32 _deltaAcc) → bytes32` (internal)

### `assertionHash(uint256 _arbGasUsed, bytes32 _restHash) → bytes32` (internal)

### `assertionRestHash(uint256 _totalMessagesRead, bytes32 _machineState, bytes32 _sendAcc, uint256 _sendCount, bytes32 _logAcc, uint256 _logCount) → bytes32` (internal)
