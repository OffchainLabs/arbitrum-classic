---
title: StandardArbERC20.sol Spec
---

Standard (i.e., non-custom) contract deployed by L2Gateway.sol as L2 ERC20. Includes standard ERC20 interface plus additional methods for deposits/withdraws

### `bridgeInit(address _l1Address, bytes _data)` (public)

initialize the token

the L2 bridge assumes this does not fail or revert

- `_l1Address`: L1 address of ERC20

- `_data`: encoded symbol/name/decimal data for initial deploy
