---
title: IERC1363Spender.sol Spec
---

IERC1363Spender Interface

Interface for any contract that wants to support approveAndCall
from ERC1363 token contracts as defined in
https://eips.ethereum.org/EIPS/eip-1363

### `onApprovalReceived(address sender, uint256 amount, bytes data) → bytes4` (external)

Handle the approval of ERC1363 tokens

Any ERC1363 smart contract calls this function on the recipient
after an `approve`. This function MAY throw to revert and reject the
approval. Return of other than the magic value MUST result in the
transaction being reverted.
Note: the token contract address is always the message sender.

- `sender`: address The address which called `approveAndCall` function

- `amount`: uint256 The amount of tokens to be spent

- `data`: bytes Additional data with no specified format

**Returns**: unless: throwing
