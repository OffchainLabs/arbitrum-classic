---
title: RetryableTicketCreator.sol Spec
id: RetryableTicketCreator
---

### `createRetryableTicket(address destAddr, uint256 l2CallValue, uint256 maxSubmissionCost, address excessFeeRefundAddress, address callValueRefundAddress, uint256 maxGas, uint256 gasPriceBid, bytes data)` (external)

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
