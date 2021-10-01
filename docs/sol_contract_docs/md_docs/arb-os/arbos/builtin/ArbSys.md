---
title: ArbSys.sol Spec
id: ArbSys
---

Precompiled contract that exists in every Arbitrum chain at address(100), 0x0000000000000000000000000000000000000064. Exposes a variety of system-level functionality.

### `arbOSVersion() → uint256` (external)

Get internal version number identifying an ArbOS build

**Returns**: version: number as int

### `arbBlockNumber() → uint256` (external)

Get Arbitrum block number (distinct from L1 block number; Arbitrum genesis block has block number 0)

**Returns**: block: number as int

### `withdrawEth(address destination) → uint256` (external)

Send given amount of Eth to dest from sender.
This is a convenience function, which is equivalent to calling sendTxToL1 with empty calldataForL1.

- `destination`: recipient address on L1

**Returns**: unique: identifier for this L2-to-L1 transaction.

### `sendTxToL1(address destination, bytes calldataForL1) → uint256` (external)

Send a transaction to L1

- `destination`: recipient address on L1

- `calldataForL1`: (optional) calldata for L1 contract call

**Returns**: a: unique identifier for this L2-to-L1 transaction.

### `getTransactionCount(address account) → uint256` (external)

get the number of transactions issued by the given external account or the account sequence number of the given contract

- `account`: target account

**Returns**: the: number of transactions issued by the given external account or the account sequence number of the given contract

### `getStorageAt(address account, uint256 index) → uint256` (external)

get the value of target L2 storage slot
This function is only callable from address 0 to prevent contracts from being able to call it

- `account`: target account

- `index`: target index of storage slot

**Returns**: stotage: value for the given account at the given index

### `isTopLevelCall() → bool` (external)

check if current call is coming from l1

**Returns**: true: if the caller of this was called directly from L1

### `EthWithdrawal(address destAddr, uint256 amount)`

### `L2ToL1Transaction(address caller, address destination, uint256 uniqueId, uint256 batchNumber, uint256 indexInBatch, uint256 arbBlockNum, uint256 ethBlockNum, uint256 timestamp, uint256 callvalue, bytes data)`
