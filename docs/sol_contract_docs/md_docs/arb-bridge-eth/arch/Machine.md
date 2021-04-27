---
title: Machine.sol Spec
---

### `addStackVal(struct Value.Data stackValHash, struct Value.Data valHash) → struct Value.Data` (internal)

### `toString(struct Machine.Data machine) → string` (internal)

### `setErrorStop(struct Machine.Data machine)` (internal)

### `setHalt(struct Machine.Data machine)` (internal)

### `addDataStackValue(struct Machine.Data machine, struct Value.Data val)` (internal)

### `addAuxStackValue(struct Machine.Data machine, struct Value.Data val)` (internal)

### `addDataStackInt(struct Machine.Data machine, uint256 val)` (internal)

### `hash(struct Machine.Data machine) → bytes32` (internal)

### `clone(struct Machine.Data machine) → struct Machine.Data` (internal)

### `deserializeMachine(bytes data, uint256 offset) → uint256, struct Machine.Data` (internal)
