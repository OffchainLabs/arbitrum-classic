---
title: NodeInterface.sol Spec
---

Interface for providing Outbox proof data
@notice This contract doesn't exist on-chain. Instead it is a virtual interface accessible at 0x00000000000000000000000000000000000000C8
This is a cute trick to allow an Arbitrum node to provide data without us having to implement an additional RPC )

### `lookupMessageBatchProof(uint256 batchNum, uint64 index) → bytes32[] proof, uint256 path, address l2Sender, address l1Dest, uint256 l2Block, uint256 l1Block, uint256 timestamp, uint256 amount, bytes calldataForL1` (external)

Returns the proof necessary to redeem a message

- `batchNum`: index of outbox entry (i.e., outgoing messages Merkle root) in array of outbox entries

- `index`: index of outgoing message in outbox entry

**Returns**: proof: Merkle proof of message inclusion in outbox entry

**Returns**: path: Index of message in outbox entry

**Returns**: l2Sender: sender if original message (i.e., caller of ArbSys.sendTxToL1)

**Returns**: l1Dest: destination address for L1 contract call

**Returns**: l2Block: l2 block number at which sendTxToL1 call was made

**Returns**: l1Block: l1 block number at which sendTxToL1 call was made

**Returns**: timestamp: l2 Timestamp at which sendTxToL1 call was made

**Returns**: amount: value in L1 message in wei

**Returns**: calldataForL1: abi-encoded L1 message data

### `estimateRetryableTicket(address sender, uint256 deposit, address destAddr, uint256 l2CallValue, uint256 maxSubmissionCost, address excessFeeRefundAddress, address callValueRefundAddress, uint256 maxGas, uint256 gasPriceBid, bytes data) → uint256, uint256` (external)

Estimate the cost of putting a message in the L2 inbox that is reexecuted

- `sender`: sender of the L1 and L2 transaction

- `deposit`: amount to deposit to sender in L2

- `destAddr`: destination L2 contract address

- `l2CallValue`: call value for retryable L2 message

- `maxSubmissionCost`: Max gas deducted from user's L2 balance to cover base submission fee

- `excessFeeRefundAddress`: maxgas x gasprice - execution cost gets credited here on L2 balance

- `callValueRefundAddress`: l2Callvalue gets credited here on L2 if retryable txn times out or gets cancelled

- `maxGas`: Max gas deducted from user's L2 balance to cover L2 execution

- `gasPriceBid`: price bid for L2 execution

- `data`: ABI encoded data of L2 message

**Returns**: gas: used, and gas price to execute this transaction
