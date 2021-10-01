---
title: ArbAddressTable.sol Spec
id: ArbAddressTable
---

Precompiled contract that exists in every Arbitrum chain at 0x0000000000000000000000000000000000000066.
Allows registering / retrieving addresses at uint indices, saving calldata.

### `register(address addr) → uint256` (external)

Register an address in the address table

- `addr`: address to register

**Returns**: index: of the address (existing index, or newly created index if not already registered)

### `lookup(address addr) → uint256` (external)

- `addr`: address to lookup

**Returns**: index: of an address in the address table (revert if address isn't in the table)

### `addressExists(address addr) → bool` (external)

Check whether an address exists in the address table

- `addr`: address to check for presence in table

**Returns**: true: if address is in table

### `size() → uint256` (external)

**Returns**: size: of address table (= first unused index)

### `lookupIndex(uint256 index) → address` (external)

- `index`: index to lookup address

**Returns**: address: at a given index in address table (revert if index is beyond end of table)

### `decompress(bytes buf, uint256 offset) → address, uint256` (external)

read a compressed address from a bytes buffer

- `buf`: bytes buffer containing an address

- `offset`: offset of target address

**Returns**: resulting: address and updated offset into the buffer (revert if buffer is too short)

### `compress(address addr) → bytes` (external)

compress an address and return the result

- `addr`: address to comppress

**Returns**: compressed: address bytes
