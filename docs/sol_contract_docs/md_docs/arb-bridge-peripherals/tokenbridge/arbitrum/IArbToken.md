---
title: IArbToken.sol Spec
---

### `bridgeMint(address account, uint256 amount)` (external)

should increase token supply by amount, and should (probably) only be callable by the L1 bridge.

### `bridgeBurn(address account, uint256 amount)` (external)

should decrease token supply by amount, and should (probably) only be callable by the L1 bridge.

### `withdraw(address account, uint256 amount)` (external)

withdraw user tokens from L2 to the L1

### `l1Address() → address` (external)

**Returns**: address: of layer 1 token
