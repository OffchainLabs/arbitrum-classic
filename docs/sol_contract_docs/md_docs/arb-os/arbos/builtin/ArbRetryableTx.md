---
title: ArbRetryableTx.sol Spec
id: ArbRetryableTx
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

### `getTimeout(bytes32 ticketId) → uint256` (external)

Return the timestamp when ticketId will age out, or zero if ticketId does not exist.
The timestamp could be in the past, because aged-out tickets might not be discarded immediately.

- `ticketId`: unique ticket identifier

**Returns**: timestamp: for ticket's deadline

### `getSubmissionPrice(uint256 calldataSize) → uint256, uint256` (external)

Return the price, in wei, of submitting a new retryable tx with a given calldata size.

- `calldataSize`: call data size to get price of (in wei)

**Returns**: Price: is guaranteed not to change until nextUpdateTimestamp.

### `getKeepalivePrice(bytes32 ticketId) → uint256, uint256` (external)

Return the price, in wei, of extending the lifetime of ticketId by an additional lifetime period. Revert if ticketId doesn't exist.

- `ticketId`: unique ticket identifier

**Returns**: Price: is guaranteed not to change until nextUpdateTimestamp.

### `keepalive(bytes32 ticketId) → uint256` (external)

Deposits callvalue into the sender's L2 account, then adds one lifetime period to the life of ticketId.
If successful, emits LifetimeExtended event.
Revert if ticketId does not exist, or if the timeout of ticketId is already at least one lifetime period in the future, or if the sender has insufficient funds (after the deposit).

- `ticketId`: unique ticket identifier

**Returns**: New: timeout of ticketId.

### `getBeneficiary(bytes32 ticketId) → address` (external)

Return the beneficiary of ticketId.
Revert if ticketId doesn't exist.

- `ticketId`: unique ticket identifier

**Returns**: address: of beneficiary for ticket

### `cancel(bytes32 ticketId)` (external)

Cancel ticketId and refund its callvalue to its beneficiary.
Revert if ticketId doesn't exist, or if called by anyone other than ticketId's beneficiary.

- `ticketId`: unique ticket identifier

### `TicketCreated(bytes32 ticketId)`

### `LifetimeExtended(bytes32 ticketId, uint256 newTimeout)`

### `Redeemed(bytes32 ticketId)`

### `Canceled(bytes32 ticketId)`
