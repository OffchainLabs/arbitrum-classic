---
title: L1ERC20Gateway.sol Spec
id: L1ERC20Gateway
---

Layer 1 Gateway contract for bridging standard ERC20s

This contract handles token deposits, holds the escrowed tokens on layer 1, and (ultimately) finalizes withdrawals.

Any ERC20 that requires non-standard functionality should use a separate gateway.
Messages to layer 2 use the inbox's createRetryableTicket method.

### `initialize(address _l2Counterpart, address _router, address _inbox, bytes32 _cloneableProxyHash, address _l2BeaconProxyFactory)` (public)

### `getOutboundCalldata(address _token, address _from, address _to, uint256 _amount, bytes _data) → bytes outboundCalldata` (public)
