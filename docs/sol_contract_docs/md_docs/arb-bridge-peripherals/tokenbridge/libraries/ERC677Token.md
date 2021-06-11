---
title: ERC677Token.sol Spec
---

### `transferAndCall(address _to, uint256 _value, bytes _data) → bool success` (public)

transfer token to a contract address with additional data if the recipient is a contact.

- `_to`: The address to transfer to.

- `_value`: The amount to be transferred.

- `_data`: The extra data to be passed to the receiving contract.
