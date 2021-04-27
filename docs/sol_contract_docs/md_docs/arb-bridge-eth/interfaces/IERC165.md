---
title: IERC165.sol Spec
---

### `supportsInterface(bytes4 interfaceID) → bool` (external)

Query if a contract implements an interface

Interface identification is specified in ERC-165. This function
uses less than 30,000 gas.

- `interfaceID`: The interface identifier, as specified in ERC-165

**Returns**: if: the contract implements `interfaceID` and
`interfaceID` is not 0xffffffff, `false` otherwise
