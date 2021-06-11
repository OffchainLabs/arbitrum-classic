---
title: L2GatewayRouter.sol Spec
---

Handles withdrawals from Ethereum into Arbitrum. Tokens are routered to their appropriate L2 gateway (Router itself also conforms to the Gateway interface).

Router also serves as an L2-L1 token address oracle.

### `initialize(address _counterpartGateway, address _defaultGateway)` (public)

### `setGateway(address[] _l1Token, address[] _gateway)` (external)

### `outboundTransfer(address _l1Token, address _to, uint256 _amount, bytes _data) → bytes` (public)

### `setDefaultGateway(address newL2DefaultGateway)` (external)
