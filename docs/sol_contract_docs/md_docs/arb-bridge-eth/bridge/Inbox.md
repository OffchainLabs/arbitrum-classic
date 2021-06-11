---
title: Inbox.sol Spec
---

### `initialize(contract IBridge _bridge)` (external)

### `sendL2MessageFromOrigin(bytes messageData) → uint256` (external)

Send a generic L2 message to the chain

This method is an optimization to avoid having to emit the entirety of the messageData in a log. Instead validators are expected to be able to parse the data from the transaction's input

- `messageData`: Data of the message being sent

### `sendL2Message(bytes messageData) → uint256` (external)

Send a generic L2 message to the chain

This method can be used to send any type of message that doesn't require L1 validation

- `messageData`: Data of the message being sent

### `sendL1FundedUnsignedTransaction(uint256 maxGas, uint256 gasPriceBid, uint256 nonce, address destAddr, bytes data) → uint256` (external)

### `sendL1FundedContractTransaction(uint256 maxGas, uint256 gasPriceBid, address destAddr, bytes data) → uint256` (external)

### `sendUnsignedTransaction(uint256 maxGas, uint256 gasPriceBid, uint256 nonce, address destAddr, uint256 amount, bytes data) → uint256` (external)

### `sendContractTransaction(uint256 maxGas, uint256 gasPriceBid, address destAddr, uint256 amount, bytes data) → uint256` (external)

### `depositEth(address destAddr) → uint256` (external)

### `depositEthRetryable(address destAddr, uint256 maxSubmissionCost, uint256 maxGas, uint256 maxGasPrice) → uint256` (external)

### `createRetryableTicket(address destAddr, uint256 l2CallValue, uint256 maxSubmissionCost, address excessFeeRefundAddress, address callValueRefundAddress, uint256 maxGas, uint256 gasPriceBid, bytes data) → uint256` (external)

Put an message in the L2 inbox that can be reexecuted for some fixed amount of time if it reverts

all msg.value will deposited to callValueRefundAddress on L2

- `destAddr`: destination L2 contract address

- `l2CallValue`: call value for retryable L2 message

- `maxSubmissionCost`: Max gas deducted from user's L2 balance to cover base submission fee

- `excessFeeRefundAddress`: maxgas x gasprice - execution cost gets credited here on L2 balance

- `callValueRefundAddress`: l2Callvalue gets credited here on L2 if retryable txn times out or gets cancelled

- `maxGas`: Max gas deducted from user's L2 balance to cover L2 execution

- `gasPriceBid`: price bid for L2 execution

- `data`: ABI encoded data of L2 message

**Returns**: unique: id for retryable transaction (keccak256(requestID, uint(0) )
