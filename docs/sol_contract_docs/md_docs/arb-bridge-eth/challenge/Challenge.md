---
title: Challenge.sol Spec
---

### `onlyOnTurn()`

### `initializeChallenge(contract IOneStepProof[] _executors, address _resultReceiver, bytes32 _executionHash, uint256 _maxMessageCount, address _asserter, address _challenger, uint256 _asserterTimeLeft, uint256 _challengerTimeLeft, contract IBridge _bridge)` (external)

### `bisectExecution(bytes32[] _merkleNodes, uint256 _merkleRoute, uint256 _challengedSegmentStart, uint256 _challengedSegmentLength, bytes32 _oldEndHash, uint256 _gasUsedBefore, bytes32 _assertionRest, bytes32[] _chainHashes)` (external)

Initiate the next round in the bisection by objecting to execution correctness with a bisection
of an execution segment with the same length but a different endpoint. This is either the initial move
or follows another execution objection

- `_merkleNodes`: List of hashes of stubs in the merkle root of segments left by the previous round

- `_merkleRoute`: Bitmap marking whether the path went left or right at each height

- `_challengedSegmentStart`: Offset of the challenged segment into the original challenged segment

- `_challengedSegmentLength`: Number of messages in the challenged segment

- `_oldEndHash`: Hash of the end of the challenged segment. This must be different than the new end since the challenger is disagreeing

- `_gasUsedBefore`: Amount of gas used at the beginning of the challenged segment

- `_assertionRest`: Hash of the rest of the assertion at the beginning of the challenged segment

- `_chainHashes`: Array of intermediate hashes of the challenged segment

### `proveContinuedExecution(bytes32[] _merkleNodes, uint256 _merkleRoute, uint256 _challengedSegmentStart, uint256 _challengedSegmentLength, bytes32 _oldEndHash, uint256 _gasUsedBefore, bytes32 _assertionRest)` (external)

### `oneStepProveExecution(bytes32[] _merkleNodes, uint256 _merkleRoute, uint256 _challengedSegmentStart, uint256 _challengedSegmentLength, bytes32 _oldEndHash, uint256 _initialMessagesRead, bytes32 _initialSendAcc, bytes32 _initialLogAcc, uint256[3] _initialState, bytes _executionProof, bytes _bufferProof, uint8 prover)` (public)

### `timeout()` (external)

### `currentResponder() → address` (public)

### `currentResponderTimeLeft() → uint256` (public)

### `InitiatedChallenge()`

### `Bisected(bytes32 challengeRoot, uint256 challengedSegmentStart, uint256 challengedSegmentLength, bytes32[] chainHashes)`

### `AsserterTimedOut()`

### `ChallengerTimedOut()`

### `OneStepProofCompleted()`

### `ContinuedExecutionProven()`
