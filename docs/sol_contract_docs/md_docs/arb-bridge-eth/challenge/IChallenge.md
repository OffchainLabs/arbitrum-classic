---
title: IChallenge.sol Spec
id: IChallenge
---

### `initializeChallenge(contract IOneStepProof[] _executors, address _resultReceiver, bytes32 _executionHash, uint256 _maxMessageCount, address _asserter, address _challenger, uint256 _asserterTimeLeft, uint256 _challengerTimeLeft, contract ISequencerInbox _sequencerBridge, contract IBridge _delayedBridge)` (external)

### `currentResponderTimeLeft() → uint256` (external)

### `lastMoveBlock() → uint256` (external)

### `timeout()` (external)

### `asserter() → address` (external)

### `challenger() → address` (external)

### `clearChallenge()` (external)
