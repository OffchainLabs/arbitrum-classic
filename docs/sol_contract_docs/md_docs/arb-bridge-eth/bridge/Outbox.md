---
title: Outbox.sol Spec
---

### `initialize(address _rollup, contract IBridge _bridge)` (external)

### `l2ToL1Sender() → address` (external)

When l2ToL1Sender returns a nonzero address, the message was originated by an L2 account
When the return value is zero, that means this is a system message

### `l2ToL1Block() → uint256` (external)

### `l2ToL1EthBlock() → uint256` (external)

### `l2ToL1Timestamp() → uint256` (external)

### `processOutgoingMessages(bytes sendsData, uint256[] sendLengths)` (external)

### `executeTransaction(uint256 outboxIndex, bytes32[] proof, uint256 index, address l2Sender, address destAddr, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 amount, bytes calldataForL1)` (external)

Executes a messages in an Outbox entry. Reverts if dispute period hasn't expired and

- `outboxIndex`: Index of OutboxEntry in outboxes array

- `proof`: Merkle proof of message inclusion in outbox entry

- `index`: Index of message in outbox entry

- `l2Sender`: sender if original message (i.e., caller of ArbSys.sendTxToL1)

- `destAddr`: destination address for L1 contract call

- `l2Block`: l2 block number at which sendTxToL1 call was made

- `l1Block`: l1 block number at which sendTxToL1 call was made

- `l2Timestamp`: l2 Timestamp at which sendTxToL1 call was made

- `amount`: value in L1 message in wei

- `calldataForL1`: abi-encoded L1 message data

### `calculateItemHash(address l2Sender, address destAddr, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 amount, bytes calldataForL1) → bytes32` (public)

### `calculateMerkleRoot(bytes32[] proof, uint256 path, bytes32 item) → bytes32` (public)

### `outboxesLength() → uint256` (public)
