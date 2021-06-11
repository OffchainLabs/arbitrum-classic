---
title: L2ERC20Gateway.sol Spec
---

### `initialize(address _l1Counterpart, address _router, address _beaconProxyFactory)` (public)

### `postUpgradeInit(address _beaconProxyFactory)` (external)

### `calculateL2TokenAddress(address l1ERC20) → address` (external)

Calculate the address used when bridging an ERC20 token

this always returns the same as the L1 oracle, but may be out of date.
For example, a custom token may have been registered but not deploy or the contract self destructed.

- `l1ERC20`: address of L1 token

**Returns**: L2: address of a bridged ERC20 token

### `cloneableProxyHash() → bytes32` (public)

### `getUserSalt(address l1ERC20) → bytes32` (public)
