---
title: ArbRetryableTx.sol Spec
---

precompiled contract in every Arbitrum chain for retryable transaction related data retrieval and interactions. Exists at 0x000000000000000000000000000000000000006E

### `redeem(bytes32 txId)` (external)

Redeem a redeemable tx.
Revert if called by an L2 contract, or if txId does not exist, or if txId reverts.
If this returns, txId has been completed and is no longer available for redemption.
If this reverts, txId is still available for redemption (until it times out or is canceled).
@param txId unique identifier of retryable message: keccak256(keccak256(ArbchainId, inbox-sequence-number), uint(0) )

### `getLifetime() → uint256` (external)

Return the minimum lifetime of redeemable txn.

**Returns**: lifetime: in seconds

### `getTimeout(bytes32 txId) → uint256` (external)

Return the timestamp when txId will age out, or zero if txId does not exist.
The timestamp could be in the past, because aged-out txs might not be discarded immediately.

- `txId`: unique identifier of retryable message: keccak256(keccak256(ArbchainId, inbox-sequence-number), uint(0) )

**Returns**: timestamp: for txn's deadline

### `getSubmissionPrice(uint256 calldataSize) → uint256, uint256` (external)

Return the price, in wei, of submitting a new retryable tx with a given calldata size.

- `calldataSize`: call data size to get price of (in wei)

**Returns**: Price: is guaranteed not to change until nextUpdateTimestamp.

### `getKeepalivePrice(bytes32 txId) → uint256, uint256` (external)

Return the price, in wei, of extending the lifetime of txId by an additional lifetime period. Revert if txId doesn't exist.

- `txId`: unique identifier of retryable message: keccak256(keccak256(ArbchainId, inbox-sequence-number), uint(0) )

**Returns**: Price: is guaranteed not to change until nextUpdateTimestamp.

### `keepalive(bytes32 txId) → uint256` (external)

Deposits callvalue into the sender's L2 account, then adds one lifetime period to the life of txId.
If successful, emits LifetimeExtended event.
Revert if txId does not exist, or if the timeout of txId is already at least one lifetime in the future, or if the sender has insufficient funds (after the deposit).

- `txId`: unique identifier of retryable message: keccak256(keccak256(ArbchainId, inbox-sequence-number), uint(0) )

**Returns**: New: timeout of txId.

### `getBeneficiary(bytes32 txId) → address` (external)

Return the beneficiary of txId.
Revert if txId doesn't exist.

- `txId`: unique identifier of retryable message: keccak256(keccak256(ArbchainId, inbox-sequence-number), uint(0) )

**Returns**: address: of beneficiary for transaction

### `cancel(bytes32 txId)` (external)

Cancel txId and refund its callvalue to its beneficiary.
Revert if txId doesn't exist, or if called by anyone other than txId's beneficiary.
@param txId unique identifier of retryable message: keccak256(keccak256(ArbchainId, inbox-sequence-number), uint(0) )

### `LifetimeExtended(bytes32 txId, uint256 newTimeout)`

### `Redeemed(bytes32 txId)`

### `Canceled(bytes32 txId)`
