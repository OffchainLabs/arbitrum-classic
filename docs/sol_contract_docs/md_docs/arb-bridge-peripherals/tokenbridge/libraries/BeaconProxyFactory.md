---
title: BeaconProxyFactory.sol Spec
id: BeaconProxyFactory
---

### `initialize(address _beacon)` (external)

### `getSalt(address user, bytes32 userSalt) → bytes32` (public)

### `createProxy(bytes32 userSalt) → address` (external)

### `calculateExpectedAddress(address user, bytes32 userSalt) → address` (public)

### `calculateExpectedAddress(bytes32 salt) → address` (public)
