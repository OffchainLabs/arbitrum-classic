---
title: ITokenGateway.sol Spec
---

### `outboundTransfer(address _token, address _to, uint256 _amount, uint256 _maxGas, uint256 _gasPriceBid, bytes _data) → bytes` (external)

### `finalizeInboundTransfer(address _token, address _from, address _to, uint256 _amount, bytes _data) → bytes` (external)

### `calculateL2TokenAddress(address l1ERC20) → address` (external)

### `OutboundTransferInitiated(address token, address _from, address _to, uint256 _transferId, uint256 _amount, bytes _data)`

### `InboundTransferFinalized(address token, address _from, address _to, uint256 _transferId, uint256 _amount, bytes _data)`
