---
title: IEthERC20Bridge.sol Spec
---

### `hasTriedDeploy(address erc20) → bool` (external)

### `registerCustomL2Token(address l2CustomTokenAddress, uint256 maxSubmissionCost, uint256 maxGas, uint256 gasPriceBid, address refundAddress) → uint256` (external)

### `fastWithdrawalFromL2(address liquidityProvider, bytes liquidityProof, address initialDestination, address erc20, uint256 amount, uint256 exitNum, uint256 maxFee)` (external)

### `withdrawFromL2(uint256 exitNum, address erc20, address initialDestination, uint256 amount)` (external)

### `deposit(address erc20, address destination, uint256 amount, uint256 maxSubmissionCost, uint256 maxGas, uint256 gasPriceBid, bytes callHookData) → uint256` (external)

### `calculateL2TokenAddress(address erc20) → address` (external)

### `ActivateCustomToken(uint256 seqNum, address l1Address, address l2Address)`

### `DeployToken(uint256 seqNum, address l1Address)`

### `WithdrawRedirected(address user, address liquidityProvider, address erc20, uint256 amount, uint256 exitNum)`

### `WithdrawExecuted(address initialDestination, address destination, address erc20, uint256 amount, uint256 exitNum)`

### `DepositToken(address destination, address sender, uint256 seqNum, uint256 value, address tokenAddress)`
