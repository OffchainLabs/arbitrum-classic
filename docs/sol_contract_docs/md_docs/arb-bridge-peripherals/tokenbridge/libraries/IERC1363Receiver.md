---
title: IERC1363Receiver.sol Spec
---

IERC1363Receiver Interface

Interface for any contract that wants to support transferAndCall or transferFromAndCall
from ERC1363 token contracts as defined in
https://eips.ethereum.org/EIPS/eip-1363

### `onTransferReceived(address operator, address sender, uint256 amount, bytes data) → bytes4` (external)

Handle the receipt of ERC1363 tokens

Any ERC1363 smart contract calls this function on the recipient
after a `transfer` or a `transferFrom`. This function MAY throw to revert and reject the
transfer. Return of other than the magic value MUST result in the
transaction being reverted.
Note: the token contract address is always the message sender.

- `operator`: address The address which called `transferAndCall` or `transferFromAndCall` function

- `sender`: address The address which are token transferred from

- `amount`: uint256 The amount of tokens transferred

- `data`: bytes Additional data with no specified format

**Returns**: unless: throwing
