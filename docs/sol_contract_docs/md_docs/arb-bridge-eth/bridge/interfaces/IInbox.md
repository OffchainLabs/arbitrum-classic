---
title: IInbox.sol Spec
---

### `sendL2Message(bytes messageData) → uint256` (external)

### `sendUnsignedTransaction(uint256 maxGas, uint256 gasPriceBid, uint256 nonce, address destAddr, uint256 amount, bytes data) → uint256` (external)

### `sendContractTransaction(uint256 maxGas, uint256 gasPriceBid, address destAddr, uint256 amount, bytes data) → uint256` (external)

### `sendL1FundedUnsignedTransaction(uint256 maxGas, uint256 gasPriceBid, uint256 nonce, address destAddr, bytes data) → uint256` (external)

### `sendL1FundedContractTransaction(uint256 maxGas, uint256 gasPriceBid, address destAddr, bytes data) → uint256` (external)

### `createRetryableTicket(address destAddr, uint256 arbTxCallValue, uint256 maxSubmissionCost, address submissionRefundAddress, address valueRefundAddress, uint256 maxGas, uint256 gasPriceBid, bytes data) → uint256` (external)

### `depositEth(uint256 maxSubmissionCost) → uint256` (external)

### `bridge() → contract IBridge` (external)
