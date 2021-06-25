---
title: L2GatewayToken.sol Spec
id: L2GatewayToken
---

Standard (i.e., non-custom) contract used as a base for different L2 Gateways

### `onlyGateway()`

### `bridgeMint(address account, uint256 amount)` (external)

Mint tokens on L2. Callable path is L1Gateway depositToken (which handles L1 escrow), which triggers L2Gateway, which calls this

- `account`: recipient of tokens

- `amount`: amount of tokens minted

### `bridgeBurn(address account, uint256 amount)` (external)

Burn tokens on L2.

only the token bridge can call this

- `account`: owner of tokens

- `amount`: amount of tokens burnt
