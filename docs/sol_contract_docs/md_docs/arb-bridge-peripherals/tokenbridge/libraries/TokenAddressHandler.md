---
title: TokenAddressHandler.sol Spec
---

### `isCustomToken(address l1Token) → bool` (public)

### `getCreate2Salt(address l1Token, address l2TemplateERC20) → bytes32` (internal)

### `calculateL2ERC20TokenAddress(address l1Token, address l2TemplateERC20, address l2ArbTokenBridgeAddress, bytes32 cloneableProxyHash) → address` (internal)

### `calculateL2TokenAddress(address l1Token, address l2TemplateERC20, address l2ArbTokenBridgeAddress, bytes32 cloneableProxyHash) → address` (internal)
