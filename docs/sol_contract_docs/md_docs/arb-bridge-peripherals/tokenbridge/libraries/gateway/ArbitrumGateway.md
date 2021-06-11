---
title: ArbitrumGateway.sol Spec
---

### `outboundTransfer(address _l1Token, address _to, uint256 _amount, uint256 _maxGas, uint256 _gasPriceBid, bytes _data) → bytes` (external)

### `finalizeInboundTransfer(address _token, address _from, address _to, uint256 _amount, bytes _data) → bytes` (external)

### `inboundEscrowAndCall(address _l2Address, uint256 _amount, address _from, address _to, bytes _data)` (external)

### `gasReserveIfCallRevert() → uint256` (public)

### `TransferAndCallTriggered(bool success, address _from, address _to, uint256 _amount, bytes callHookData)`
