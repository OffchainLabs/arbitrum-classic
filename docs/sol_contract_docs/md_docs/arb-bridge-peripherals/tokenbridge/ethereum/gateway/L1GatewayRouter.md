---
title: L1GatewayRouter.sol Spec
---

Handles deposits from Erhereum into Arbitrum. Tokens are routered to their appropriate L1 gateway (Router itself also conforms to the Gateway itnerface).

Router also serves as an L1-L2 token address oracle.

### `onlyOwner()`

### `initialize(address _owner, address _defaultGateway, address _whitelist, address _counterpartGateway, address _inbox)` (public)

### `setDefaultGateway(address newL1DefaultGateway, uint256 _maxGas, uint256 _gasPriceBid, uint256 _maxSubmissionCost) → uint256` (external)

### `setOwner(address newOwner)` (external)

### `setGateway(address _gateway, uint256 _maxGas, uint256 _gasPriceBid, uint256 _maxSubmissionCost) → uint256` (external)

### `setGateways(address[] _token, address[] _gateway, uint256 _maxGas, uint256 _gasPriceBid, uint256 _maxSubmissionCost) → uint256` (external)

### `outboundTransfer(address _token, address _to, uint256 _amount, uint256 _maxGas, uint256 _gasPriceBid, bytes _data) → bytes` (public)
