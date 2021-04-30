---
title: StandardArbERC20.sol Spec
---

Standard (i.e., non-custom) contract deployed by ArbTokenBridge.sol as L2 ERC20. Includes standard ERC20 interface plus additional methods for deposits/withdraws

### `onlyBridge()`

### `bridgeInit(address _l1Address, bytes _data)` (external)

initialize the token

the L2 bridge assumes this does not fail or revert

- `_l1Address`: L1 address of ERC20

- `_data`: encoded symbol/name/decimal data for initial deploy

### `bridgeMint(address account, uint256 amount)` (external)

Mint tokens on L2. Callable path is EthErc20Bridge.depositToken (which handles L1 escrow), which triggers ArbTokenBridge.mintFromL1, which calls this

- `account`: recipient of tokens

- `amount`: amount of tokens minted

### `bridgeBurn(address account, uint256 amount)` (external)

Burn tokens on L2.

only the token bridge can call this

- `account`: owner of tokens

- `amount`: amount of tokens burnt

### `withdraw(address account, uint256 amount)` (external)

Initiates a token withdrawal

- `account`: destination address

- `amount`: amount of tokens withdrawn

### `migrate(address account, uint256 amount)` (external)

Migrate tokens from to a custom token contract; this should only happen/matter if a standard ERC20 is deployed for an L1 custom contract before the L2 custom contract gets registered

- `account`: destination address

- `amount`: amount of tokens withdrawn
