---
title: L2ArbitrumGateway.sol Spec
---

Common interface for gatways on Arbitrum messaging to L1.

### `gasReserveIfCallRevert() → uint256` (public)

### `getOutboundCalldata(address _token, address _from, address _to, uint256 _amount, bytes _data) → bytes outboundCalldata` (public)

### `outboundTransfer(address _l1Token, address _to, uint256 _amount, bytes _data) → bytes` (public)

### `outboundTransfer(address _l1Token, address _to, uint256 _amount, uint256 _maxGas, uint256 _gasPriceBid, bytes _data) → bytes` (public)

Initiates a token withdrawal from Arbitrum to Ethereum

- `_l1Token`: l1 address of token

- `_to`: destination address

- `_amount`: amount of tokens withdrawn

- `_maxGas`: max gas provided for outbox execution market (todo)

- `_gasPriceBid`: provided for outbox execution market (todo)
  @ @return encoded unique identifier for withdrawal

### `finalizeInboundTransfer(address _token, address _from, address _to, uint256 _amount, bytes _data) → bytes` (external)

Mint on L2 upon L1 deposit.
If token not yet deployed and symbol/name/decimal data is included, deploys StandardArbERC20

Callable only by the L1ERC20Gateway.outboundTransfer method. For initial deployments of a token the L1 L1ERC20Gateway
is expected to include the deployData. If not a L1 withdrawal is automatically triggered for the user

- `_token`: L1 address of ERC20

- `_from`: account that initiated the deposit in the L1

- `_to`: account to be credited with the tokens in the L2 (can be the user's L2 account or a contract)

- `_amount`: token amount to be minted to the user

- `_data`: encoded symbol/name/decimal data for deploy, in addition to any additional callhook data
