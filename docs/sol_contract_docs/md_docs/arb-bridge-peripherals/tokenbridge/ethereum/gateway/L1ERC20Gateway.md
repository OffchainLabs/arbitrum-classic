---
title: L1ERC20Gateway.sol Spec
---

Layer 1 Gateway contract for bridging standard ERC20s

This contract handles token deposits, holds the escrowed tokens on layer 1, and (ultimately) finalizes withdrawals.

Any ERC20 that requires non-standard functionality should use a separate gateway.
Messages to layer 2 use the inbox's createRetryableTicket method.

### `initialize(address _l2Counterpart, address _router, address _inbox, bytes32 _cloneableProxyHash, address _l2BeaconProxyFactory)` (public)

### `postUpgradeInit(bytes32 _cloneableProxyHash, address _l2BeaconProxyFactory)` (external)

### `getOutboundCalldata(address _token, address _from, address _to, uint256 _amount, bytes _data) → bytes outboundCalldata` (public)

### `calculateL2TokenAddress(address l1ERC20) → address` (external)

Calculate the address used when bridging an ERC20 token

this always returns the same as the L1 oracle, but may be out of date.
For example, a custom token may have been registered but not deploy or the contract self destructed.

- `l1ERC20`: address of L1 token

**Returns**: L2: address of a bridged ERC20 token
