---
title: L1CustomGateway.sol Spec
id: L1CustomGateway
---

Gatway for "custom" bridging functionality

Handles some (but not all!) custom Gateway needs.

### `initialize(address _l1Counterpart, address _l1Router, address _inbox, address _owner)` (public)

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
