---
title: IArbStandardToken.sol Spec
---

### `bridgeInit(address _l1Address, bytes _data)` (external)

initialize the token

the L2 bridge assumes this does not fail or revert

- `_l1Address`: L1 address of ERC20

- `_data`: encoded symbol/name/decimal data for initial deploy

### `migrate(address destination, uint256 amount)` (external)

Migrate tokens from to a custom token contract; this should only happen/matter
if a standard ERC20 is deployed for an L1 custom contract before the L2 custom contract gets registered

- `destination`: destination address

- `amount`: amount of tokens withdrawn
