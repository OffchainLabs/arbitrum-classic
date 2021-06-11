---
title: L1ArbitrumGateway.sol Spec
id: L1ArbitrumGateway
---

Common interface for gatways on L1 messaging to Arbitrum.

### `finalizeInboundTransfer(address _token, address _from, address _to, uint256 _amount, bytes _data) → bytes` (external)

Finalizes a withdrawal via Outbox message; callable only by L2Gateway.outboundTransfer

- `_token`: L1 address of token being withdrawn from

- `_from`: initiator of withdrawal

- `_to`: address the L2 withdrawal call set as the destination.

- `_amount`: Token amount being withdrawn

- `_data`: encoded exitNum (Sequentially increasing exit counter determined by the L2Gateway) and additinal hook data

### `gasReserveIfCallRevert() → uint256` (public)

### `getCurrentDestination(uint256 _exitNum, address _initialDestination) → address` (public)

### `parseInboundData(bytes _data) → uint256 _exitNum, bytes _extraData` (public)

### `outboundTransfer(address _l1Token, address _to, uint256 _amount, uint256 _maxGas, uint256 _gasPriceBid, bytes _data) → bytes res` (external)

Deposit ERC20 token from Ethereum into Arbitrum. If L2 side hasn't been deployed yet, includes name/symbol/decimals data for initial L2 deploy. Initiate by GatewayRouter.

- `_l1Token`: L1 address of ERC20

- `_to`: account to be credited with the tokens in the L2 (can be the user's L2 account or a contract)

- `_amount`: Token Amount

- `_maxGas`: Max gas deducted from user's L2 balance to cover L2 execution

- `_gasPriceBid`: Gas price for L2 execution

- `_data`: encoded data from router and user

**Returns**: res: abi encoded inbox sequence number

### `getOutboundCalldata(address _l1Token, address _from, address _to, uint256 _amount, bytes _data) → bytes outboundCalldata` (public)
