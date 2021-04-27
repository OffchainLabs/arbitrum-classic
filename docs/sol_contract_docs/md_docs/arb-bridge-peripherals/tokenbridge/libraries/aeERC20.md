---
title: aeERC20.sol Spec
---

Arbitrum extended ERC20

The recommended ERC20 implementation for Layer 2 tokens

This implements the ERC20 standard with extensions to improve UX (ERC1363 & ERC2612)

### `initialize(string name, string symbol, uint8 decimals)` (public)

### `supportsInterface(bytes4 interfaceId) → bool` (public)

See {IERC165-supportsInterface}.

### `transferAndCall(address recipient, uint256 amount) → bool` (public)

Transfer tokens to a specified address and then execute a callback on recipient.

- `recipient`: The address to transfer to.

- `amount`: The amount to be transferred.

**Returns**: A: boolean that indicates if the operation was successful.

### `transferAndCall(address recipient, uint256 amount, bytes data) → bool` (public)

Transfer tokens to a specified address and then execute a callback on recipient.

- `recipient`: The address to transfer to

- `amount`: The amount to be transferred

- `data`: Additional data with no specified format

**Returns**: A: boolean that indicates if the operation was successful.

### `transferFromAndCall(address sender, address recipient, uint256 amount) → bool` (public)

Transfer tokens from one address to another and then execute a callback on recipient.

- `sender`: The address which you want to send tokens from

- `recipient`: The address which you want to transfer to

- `amount`: The amount of tokens to be transferred

**Returns**: A: boolean that indicates if the operation was successful.

### `transferFromAndCall(address sender, address recipient, uint256 amount, bytes data) → bool` (public)

Transfer tokens from one address to another and then execute a callback on recipient.

- `sender`: The address which you want to send tokens from

- `recipient`: The address which you want to transfer to

- `amount`: The amount of tokens to be transferred

- `data`: Additional data with no specified format

**Returns**: A: boolean that indicates if the operation was successful.

### `approveAndCall(address spender, uint256 amount) → bool` (public)

Approve spender to transfer tokens and then execute a callback on recipient.

- `spender`: The address allowed to transfer to

- `amount`: The amount allowed to be transferred

**Returns**: A: boolean that indicates if the operation was successful.

### `approveAndCall(address spender, uint256 amount, bytes data) → bool` (public)

Approve spender to transfer tokens and then execute a callback on recipient.

- `spender`: The address allowed to transfer to.

- `amount`: The amount allowed to be transferred.

- `data`: Additional data with no specified format.

**Returns**: A: boolean that indicates if the operation was successful.

### `_checkAndCallTransfer(address sender, address recipient, uint256 amount, bytes data) → bool` (internal)

Internal function to invoke `onTransferReceived` on a target address
The call is not executed if the target address is not a contract

- `sender`: address Representing the previous owner of the given token value

- `recipient`: address Target address that will receive the tokens

- `amount`: uint256 The amount mount of tokens to be transferred

- `data`: bytes Optional data to send along with the call

**Returns**: whether: the call correctly returned the expected magic value

### `_checkAndCallApprove(address spender, uint256 amount, bytes data) → bool` (internal)

Internal function to invoke `onApprovalReceived` on a target address
The call is not executed if the target address is not a contract

- `spender`: address The address which will spend the funds

- `amount`: uint256 The amount of tokens to be spent

- `data`: bytes Optional data to send along with the call

**Returns**: whether: the call correctly returned the expected magic value
