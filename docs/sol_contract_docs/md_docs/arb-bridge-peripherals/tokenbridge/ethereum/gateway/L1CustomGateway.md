---
title: L1CustomGateway.sol Spec
id: L1CustomGateway
---

Gatway for "custom" bridging functionality

Handles some (but not all!) custom Gateway needs.

### `initialize(address _l1Counterpart, address _l1Router, address _inbox, address _owner)` (public)

### `calculateL2TokenAddress(address l1ERC20) → address` (external)

Calculate the address used when bridging an ERC20 token

this always returns the same as the L1 oracle, but may be out of date.
For example, a custom token may have been registered but not deploy or the contract self destructed.

- `l1ERC20`: address of L1 token

**Returns**: L2: address of a bridged ERC20 token

### `registerTokenToL2(address _l2Address, uint256 _maxGas, uint256 _gasPriceBid, uint256 _maxSubmissionCost) → uint256` (external)

Allows L1 Token contract to trustlessly register its custom L2 counterpart.

- `_l2Address`: counterpart address of L1 token

- `_maxGas`: max gas for L2 retryable exrecution

- `_gasPriceBid`: gas price for L2 retryable ticket

- `_maxSubmissionCost`: base submission cost L2 retryable tick3et

**Returns**: Retryable: ticket ID

### `forceRegisterTokenToL2(address[] _l1Addresses, address[] _l2Addresses, uint256 _maxGas, uint256 _gasPriceBid, uint256 _maxSubmissionCost) → uint256` (external)

Allows owner to force register a custom L1/L2 token pair.

\_l1Addresses[i] counterpart is assumed to be \_l2Addresses[i]

- `_l1Addresses`: array of L1 addresses

- `_l2Addresses`: array of L2 addresses

- `_maxGas`: max gas for L2 retryable exrecution

- `_gasPriceBid`: gas price for L2 retryable ticket

- `_maxSubmissionCost`: base submission cost L2 retryable tick3et

**Returns**: Retryable: ticket ID

### `TokenSet(address l1Address, address l2Address)`
