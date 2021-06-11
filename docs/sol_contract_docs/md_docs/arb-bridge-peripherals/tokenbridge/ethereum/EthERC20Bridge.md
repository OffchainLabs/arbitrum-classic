---
title: EthERC20Bridge.sol Spec
---

Layer 1 contract for bridging ERC20s and custom fungible tokens

This contract handles token deposits, holds the escrowed tokens on layer 1, and (ulimately) finalizes withdrawals.

Custom tokens that are sufficiently "weird," (i.e., dynamic supply adjustment, say) should use their own, custom bridge.
All messages to layer 2 use the inbox's createRetryableTicket method.

### `onlyL2Address()`

This ensures that a method can only be called from the L2 pair of this contract

### `initialize(address _inbox, address _l2TemplateERC20, address _l2ArbTokenBridgeAddress)` (external)

Initialize L1 bridge

- `_inbox`: Address of Arbitrum chain's L1 Inbox.sol contract used to submit transactions to the L2

- `_l2TemplateERC20`: Address of template ERC20 (i.e, StandardArbERC20.sol). Used for salt in computing L2 address.

- `_l2ArbTokenBridgeAddress`: Address of L2 side of token bridge (ArbTokenBridge.sol)

### `registerCustomL2Token(address l2CustomTokenAddress, uint256 maxSubmissionCost, uint256 maxGas, uint256 gasPriceBid, address refundAddress) → uint256` (external)

Called by a custom token on L1 to register with a previously deployed custom token on L2.
The L1 contract should conform to ICustomToken.sol; L2 contract should conform to IArbCustomToken.sol.

If the L2 side hasn't yet been deployed, a safe, temporary fallback scenario will take place
(see ArbTokenBridge.customTokenRegistered). But please, save yourself and the trouble, and just deploy the L2 contract first.

- `l2CustomTokenAddress`: L2 address of previously deployed custom token contract

- `maxSubmissionCost`: Max gas deducted from user's L2 balance to cover base submission fee

- `maxGas`: Max gas deducted from user's L2 balance to cover L2 execution

- `gasPriceBid`: Gas price for L2 execution

- `refundAddress`: Address to refund overbid for maxSubmissionCost and/or maxGas\*gasPriceBid execution

### `transferExitAndCall(address initialDestination, address erc20, uint256 amount, uint256 exitNum, address to, bytes data)` (external)

Allows a user to redirect their right to claim a withdrawal to another address

This method also allows you to make an arbitrary call after the transfer, similar to ERC677

- `initialDestination`: address the L2 withdrawal call initially set as the destination.

- `erc20`: L1 token address

- `amount`: token amount (should match amount in previously-initiated withdrawal)

- `exitNum`: Sequentially increasing exit counter determined by the L2 bridge

- `data`: optional data for external call upon transfering the exit

### `withdrawFromL2(uint256 exitNum, address erc20, address initialDestination, uint256 amount)` (external)

Finalizes a withdraw via Outbox message; callable only by ArbTokenBridge.\_withdraw

- `exitNum`: Sequentially increasing exit counter determined by the L2 bridge

- `erc20`: L1 address of token being withdrawn from

- `initialDestination`: address the L2 withdrawal call initially set as the destination.

- `amount`: Token amount being withdrawn

### `getDepositCalldata(address erc20, address sender, address destination, uint256 amount, bytes callHookData) → bool isDeployed, bytes depositCalldata` (public)

Utility method that allows you to get the calldata to be submitted to the L2 for a token deposit

- `erc20`: L1 address of ERC20

- `sender`: account initiating the L1 deposit

- `destination`: account to be credited with the tokens in the L2 (can be the user's L2 account or a contract)

- `amount`: Token Amount

- `callHookData`: optional data for external call upon minting

**Returns**: isDeployed: if token has already been deployed to the L2

**Returns**: depositCalldata: calldata submitted to the L2

### `deposit(address erc20, address destination, uint256 amount, uint256 maxSubmissionCost, uint256 maxGas, uint256 gasPriceBid, bytes callHookData) → uint256 seqNum, uint256 depositCalldataLength` (external)

Deposit standard or custom ERC20 token. If L2 side hasn't been deployed yet, includes name/symbol/decimals data for initial L2 deploy.

- `erc20`: L1 address of ERC20

- `destination`: account to be credited with the tokens in the L2 (can be the user's L2 account or a contract)

- `amount`: Token Amount

- `maxSubmissionCost`: Max gas deducted from user's L2 balance to cover base submission fee

- `maxGas`: Max gas deducted from user's L2 balance to cover L2 execution

- `gasPriceBid`: Gas price for L2 execution

- `callHookData`: optional data for external call upon minting

**Returns**: seqNum: ticket ID used to redeem the retryable transaction in the L2

**Returns**: depositCalldataLength: length of calldata submitted to the L2

### `calculateL2TokenAddress(address erc20) → address` (public)

Calculate the address used when bridging an ERC20 token

this always returns the same as the L@ oracle, but may be out of date.
For example, a custom token may have been registered but not deploy or the contract self destructed.

- `erc20`: address of L1 token

**Returns**: L2: address of a bridged ERC20 token
