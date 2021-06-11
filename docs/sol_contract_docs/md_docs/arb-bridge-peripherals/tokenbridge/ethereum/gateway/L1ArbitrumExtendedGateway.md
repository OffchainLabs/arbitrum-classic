---
title: L1ArbitrumExtendedGateway.sol Spec
---

### `transferExitAndCall(uint256 _exitNum, address _initialDestination, address _newDestination, bytes _data)` (external)

Allows a user to redirect their right to claim a withdrawal to another address.

This method also allows you to make an arbitrary call after the transfer, similar to ERC677.
This does not change the original data that will be triggered with the withdrawal's external call.
The exit receiver is the one to

- `_exitNum`: Sequentially increasing exit counter determined by the L2 bridge

- `_initialDestination`: address the L2 withdrawal call initially set as the destination.

- `_newDestination`: address the L1 will now call instead of the previously set destination

- `_data`: optional data for external call upon transfering the exit

### `getCurrentDestination(uint256 _exitNum, address _initialDestination) → address` (public)

### `encodeWithdrawal(uint256 _exitNum, address _initialDestination) → bytes32` (public)

### `WithdrawRedirected(address from, address to, uint256 exitNum, bytes data, bool madeExternalCall)`
