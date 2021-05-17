---
title: IArbTokenBridge.sol Spec
---

### `mintFromL1(address l1ERC20, address sender, address dest, uint256 amount, bytes deployData, bytes callHookData)` (external)

Mints tokens in the L2

This function is only callable by the L1 bridge

### `customTokenRegistered(address l1Address, address l2Address)` (external)

Registers a custom ERC20 token implementation to be used when the bridge is minting tokens

This function is only callable by the L1 bridge

### `migrate(address l1ERC20, address sender, address destination, uint256 amount)` (external)

Migrates user balance from an erc20 token to custom token implementation

If a token is bridged before a custom implementation is set users can call this method to migrate to the custom version

### `withdraw(address l1ERC20, address sender, address destination, uint256 amount) → uint256` (external)

Withdraws user funds to the L1

Users need to wait for the rollup's dispute period before triggering their withdrawal in the L1

**Returns**: unique: withdrawal identifier needed to execute the withdrawal in the L1

### `calculateL2TokenAddress(address l1ERC20) → address` (external)

An address oracle that provides users with the L2 address of an L1 token

**Returns**: address: of L2 token that was created using this bridge

### `MintAndCallTriggered(bool success, address sender, address dest, uint256 amount, bytes callHookData)`

### `WithdrawToken(uint256 withdrawalId, address l1Address, uint256 amount, address destination, uint256 exitNum)`

### `TokenCreated(address l1Address, address l2Address)`

### `CustomTokenRegistered(address l1Address, address l2Address)`

### `TokenMinted(address l1Address, address l2Address, address sender, address dest, uint256 amount, bool usedCallHook)`

### `TokenMigrated(address l1Address, address account, uint256 amount)`
