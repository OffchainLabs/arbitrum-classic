---
title: ArbTokenBridge.sol Spec
---

### `onlyEthPair()`

This ensures that a method can only be called from the L1 pair of this contract

### `initialize(address _l1Pair, address _templateERC20)` (external)

Initialize L2 bridge

- `_l1Pair`: Address of L1 side of token bridge (EthERC20Bridge.sol)

- `_templateERC20`: Address of template ERC20 (i.e, StandardArbERC20.sol). Used for salt in computing L2 address.

### `mintAndCall(contract IArbToken token, uint256 amount, address sender, address dest, bytes data)` (external)

this function can only be callable by the bridge itself

This method is inspired by EIP 677/1363 for calls to be executed after minting.
A reserve amount of gas is always kept in case this call reverts or uses up all gas.
The reserve is the amount of gas needed to catch the revert and do the necessary alternative logic.

### `mintFromL1(address l1ERC20, address sender, address dest, uint256 amount, bytes deployData, bytes callHookData)` (external)

Mint on L2 upon L1 deposit.
If token not yet deployed and symbol/name/decimal data is included, deploys StandardArbERC20
If minting a custom token whose L2 counterpart hasn't yet been deployed/registered (!) deploys a temporary StandardArbERC20 that can later be migrated to custom token.

Callable only by the EthERC20Bridge.depositToken function. For initial deployments of a token the L1 EthERC20Bridge
is expected to include the deployData. If not a L1 withdrawal is automatically triggered for the user

- `l1ERC20`: L1 address of ERC20

- `sender`: account that initiated the deposit in the L1

- `dest`: account to be credited with the tokens in the L2 (can be the user's L2 account or a contract)

- `amount`: token amount to be minted to the user

- `deployData`: encoded symbol/name/decimal data for initial deploy

- `callHookData`: optional data for external call upon minting

### `customTokenRegistered(address l1Address, address l2Address)` (external)

Sets the L1 / L2 custom token pairing; called from the L1 via EthErc20Bridge.registerCustomL2Token

this doesn't check if the L2 token is actually deployed - this way the L1 and L2 address oracles are
always consistent. The necessary existence checks are done before interacting with the tokens.

- `l1Address`: Address of L1 custom token implementation

- `l2Address`: Address of L2 custom token implementation

### `withdraw(address l1ERC20, address sender, address destination, uint256 amount) → uint256` (external)

send a withdraw message to the L1 outbox

this call is initiated by the token, ie StandardArbERC20.withdraw or WhateverCustomToken.whateverWithdrawMethod

- `l1ERC20`: L1 address of ERC20

- `destination`: the account to be credited with the tokens

- `amount`: token amount to be withdrawn

### `migrate(address l1ERC20, address sender, address destination, uint256 amount)` (external)

If a token is bridged as a StandardArbERC20 before a custom implementation is set,
users can call this method via StandardArbERC20.migrate to migrate to the custom version

- `l1ERC20`: L1 address of ERC20

- `sender`: the account that called the migration

- `destination`: the account to be credited with the tokens

- `amount`: token amount to be migrated

### `calculateL2TokenAddress(address l1ERC20) → address` (public)

Calculate the address used when bridging an ERC20 token

this always returns the same as the L1 oracle, but may be out of date.
For example, a custom token may have been registered but not deploy or the contract self destructed.

- `l1ERC20`: address of L1 token

**Returns**: L2: address of a bridged ERC20 token

### `calculateL2ERC20TokenAddress(address l1ERC20) → address` (public)

Calculate the address used when bridging an ERC20 token

If there is a custom token registered with the bridge, this address won't be used.

- `l1ERC20`: address of L1 token

**Returns**: L2: address of ERC20 tokens deployed by this bridge

### `getBeacon() → address` (external)

utility function used in ClonableBeaconProxy.

this method makes it possible to use ClonableBeaconProxy.creationCode without encoding constructor parameters

**Returns**: the: token logic to be used in a proxy contract.
