---
title: OneStepProofCommon.sol Spec
---

### `executeStep(contract IBridge bridge, uint256 initialMessagesRead, bytes32[2] accs, bytes proof, bytes bproof) → uint64 gas, uint256 totalMessagesRead, bytes32[4] fields` (external)

### `executeStepDebug(contract IBridge bridge, uint256 initialMessagesRead, bytes32[2] accs, bytes proof, bytes bproof) → string startMachine, string afterMachine` (external)

### `returnContext(struct OneStepProofCommon.AssertionContext context) → uint64 gas, uint256 totalMessagesRead, bytes32[4] fields` (internal)

### `popVal(struct OneStepProofCommon.ValueStack stack) → struct Value.Data` (internal)

### `pushVal(struct OneStepProofCommon.ValueStack stack, struct Value.Data val)` (internal)

### `handleError(struct OneStepProofCommon.AssertionContext context)` (internal)

### `deductGas(struct OneStepProofCommon.AssertionContext context, uint64 amount) → bool` (internal)

### `handleOpcodeError(struct OneStepProofCommon.AssertionContext context)` (internal)

### `initializeExecutionContext(uint256 initialMessagesRead, bytes32[2] accs, bytes proof, bytes bproof, contract IBridge bridge) → struct OneStepProofCommon.AssertionContext` (internal)

### `executeOp(struct OneStepProofCommon.AssertionContext context)` (internal)

### `opInfo(uint256 opCode) → uint256, uint256, uint64, function (struct OneStepProofCommon.AssertionContext) view` (internal)
