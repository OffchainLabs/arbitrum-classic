---
title: TokenGateway.sol Spec
---

### `onlyCounterpartGateway()`

### `onlyRouter()`

### `calculateL2TokenAddress(address l1ERC20) → address` (external)

Calculate the address used when bridging an ERC20 token

this always returns the same as the L1 oracle, but may be out of date.
For example, a custom token may have been registered but not deploy or the contract self destructed.

- `l1ERC20`: address of L1 token

**Returns**: L2: address of a bridged ERC20 token

### `outboundTransfer(address _token, address _to, uint256 _amount, uint256 _maxGas, uint256 _gasPriceBid, bytes _data) → bytes` (external)

### `getOutboundCalldata(address _token, address _from, address _to, uint256 _amount, bytes _data) → bytes` (public)

### `finalizeInboundTransfer(address _token, address _from, address _to, uint256 _amount, bytes _data) → bytes` (external)
