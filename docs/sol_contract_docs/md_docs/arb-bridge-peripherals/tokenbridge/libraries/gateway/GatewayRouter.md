---
title: GatewayRouter.sol Spec
---

Common interface for L1 and L2 Gateway Routers

### `finalizeInboundTransfer(address _token, address _from, address _to, uint256 _amount, bytes _data) → bytes` (external)

### `outboundTransfer(address _token, address _to, uint256 _amount, uint256 _maxGas, uint256 _gasPriceBid, bytes _data) → bytes` (public)

### `getOutboundCalldata(address _token, address _from, address _to, uint256 _amount, bytes _data) → bytes` (public)

### `getGateway(address _token) → address gateway` (public)

### `TransferRouted(address token, address _userFrom, address _userTo, address gateway)`

### `GatewaySet(address l1Token, address gateway)`

### `DefaultGatewayUpdated(address newDefaultGateway)`
