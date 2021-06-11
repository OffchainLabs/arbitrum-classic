---
title: ICustomToken.sol Spec
id: ICustomToken
---

Minimum expected interface for L1 custom token (see TestCustomTokenL1.sol for an example implementation)

### `registerTokenOnL2(address l2CustomTokenAddress, uint256 maxSubmissionCost, uint256 maxGas, uint256 gasPriceBid)` (external)

Should make an external call to EthERC20Bridge.registerCustomL2Token

### `transferFrom(address sender, address recipient, uint256 amount) → bool` (external)

### `balanceOf(address account) → uint256` (external)
