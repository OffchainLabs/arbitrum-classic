---
title: Rollup.sol Spec
---

### `initialize(bytes32 _machineHash, uint256[4] _rollupParams, address _stakeToken, address _owner, bytes _extraConfig, address[6] connectedContracts, address[2] _facets, uint256[2] sequencerInboxParams)` (public)

### `getFacets() → address, address` (public)

Fallback and delegate functions from OZ
https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v3.4.0/contracts/proxy/TransparentUpgradeableProxy.sol
And dispatch pattern from EIP-2535: Diamonds

### `getAdminFacet() → address` (public)

### `getUserFacet() → address` (public)

### `fallback()` (external)

Fallback function that delegates calls to the address returned by `_implementation()`. Will run if no other
function in the contract matches the call data.

### `receive()` (external)

Fallback function that delegates calls to the address returned by `_implementation()`. Will run if call data
is empty.
