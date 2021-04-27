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

### `fastWithdrawalFromL2(address liquidityProvider, bytes liquidityProof, address initialDestination, address erc20, uint256 amount, uint256 exitNum, uint256 maxFee)` (external)

Allows a user to redirect their right to claim a withdrawal to a liquidityProvider, in exchange for a fee.

This method expects the liquidityProvider to verify the liquidityProof, but it ensures the withdrawer's balance
is appropriately updated. It is otherwise agnostic to the details of IExitLiquidityProvider.requestLiquidity.

- `liquidityProvider`: address of an IExitLiquidityProvider

- `liquidityProof`: encoded data required by the liquidityProvider in order to validate a fast withdrawal.

- `initialDestination`: address the L2 withdrawal call initially set as the destination.

- `erc20`: L1 token address

- `amount`: token amount (should match amount in previously-initiated withdrawal)

- `exitNum`: Sequentially increasing exit counter determined by the L2 bridge

- `maxFee`: max mount of erc20 token user will pay for fast exit

### `withdrawFromL2(uint256 exitNum, address erc20, address initialDestination, uint256 amount)` (external)

Finalizes a withdraw via Outbox message; callable only by ArbTokenBridge.\_withdraw

- `exitNum`: Sequentially increasing exit counter determined by the L2 bridge

- `erc20`: L1 address of token being withdrawn from

- `initialDestination`: address the L2 withdrawal call initially set as the destination.

- `amount`: Token amount being withdrawn

### `callStatic(address targetContract, bytes4 targetFunction) → bytes` (internal)

utility function used to perform external read-only calls.

the result is returned even if the call failed, the L2 is expected to
identify and deal with this.

**Returns**: result: bytes, even if the call failed.

### `depositToken(address erc20, address sender, address destination, uint256 amount, struct EthERC20Bridge.RetryableTxParams retryableParams, bytes deployData, bytes callHookData) → uint256` (internal)

internal function used to escrow tokens, then trigger their minting in the L2

- `erc20`: L1 token address

- `sender`: account that initiated the deposit in the L1

- `destination`: account to be credited with the tokens in the L2 (can be the user's L2 account or a contract)

- `amount`: token amount to be minted to the user

- `retryableParams`: params for inbox's createRetryableTicket

- `deployData`: encoded symbol/name/decimal data for initial deploy

- `callHookData`: optional data for external call upon minting

**Returns**: ticket: ID used to redeem the retryable transaction in the L2

### `deposit(address erc20, address destination, uint256 amount, uint256 maxSubmissionCost, uint256 maxGas, uint256 gasPriceBid, bytes callHookData) → uint256` (external)

Deposit standard or custom ERC20 token. If L2 side hasn't been deployed yet, includes name/symbol/decimals data for initial L2 deploy.

- `erc20`: L1 address of ERC20

- `destination`: account to be credited with the tokens in the L2 (can be the user's L2 account or a contract)

- `amount`: Token Amount

- `maxSubmissionCost`: Max gas deducted from user's L2 balance to cover base submission fee

- `maxGas`: Max gas deducted from user's L2 balance to cover L2 execution

- `gasPriceBid`: Gas price for L2 execution

- `callHookData`: optional data for external call upon minting

**Returns**: ticket: ID used to redeem the retryable transaction in the L2

### `encodeWithdrawal(uint256 exitNum, address initialDestination, address erc20, uint256 amount) → bytes32` (internal)

Output unique identifier for a token withdrawal. Used for tracking fast exits.

- `exitNum`: Sequentially increasing exit counter

- `initialDestination`: address for tokens before/unless otherwise redirected (via, i.e., a fast-withdrawal)

- `erc20`: L1 address of token being withdrawn

- `amount`: amount of token being withdrawn

**Returns**: bytes: hash uniquely identifying withdrawal

### `calculateL2TokenAddress(address erc20) → address` (public)

Calculate the address used when bridging an ERC20 token

this always returns the same as the L@ oracle, but may be out of date.
For example, a custom token may have been registered but not deploy or the contract self destructed.

- `erc20`: address of L1 token

**Returns**: L2: address of a bridged ERC20 token
