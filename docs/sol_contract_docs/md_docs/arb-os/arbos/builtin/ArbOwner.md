---
title: ArbOwner.sol Spec
id: ArbOwner
---

### `giveOwnership(address newOwnerAddr)` (external)

### `addToReserveFunds()` (external)

### `setFeesEnabled(bool enabled)` (external)

### `getFeeRecipients() → address, address` (external)

### `setFeeRecipients(address netFeeRecipient, address congestionFeeRecipient)` (external)

### `setFairGasPriceSender(address addr)` (external)

### `setGasAccountingParams(uint256 speedLimitPerBlock, uint256 gasPoolMax, uint256 maxTxGasLimit)` (external)

### `setSecondsPerSend(uint256 blocksPerSend)` (external)

### `deployContract(bytes constructorData, address deemedSender, uint256 deemedNonce) → address` (external)

### `startCodeUpload()` (external)

### `continueCodeUpload(bytes marshalledCode)` (external)

### `getUploadedCodeHash() → bytes32` (external)

### `finishCodeUploadAsArbosUpgrade(bytes32 requiredCodeHash)` (external)

### `finishCodeUploadAsPluggable(uint256 id, bool keepState)` (external)

### `bindAddressToPluggable(address addr, uint256 pluggableId)` (external)

### `getTotalOfEthBalances() → uint256` (external)
