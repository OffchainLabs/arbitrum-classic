---
id: Solidity_Support
title: Solidity Support
custom_edit_url: https://github.com/OffchainLabs/arbitrum/edit/master/docs/Solidity_Support.md
---

Arbitrum Rollup supports EVM transactions, and therefore allows you to trustlessly deploy Solidity contracts (as well as Vyper or any other language that compile to EVM). Arbitrum supports almost all Solidity code as expected with a few exceptions that we detail below.

# Differences from Solidity on Ethereum

Although Arbitrum supports Solidity code, there are differences in the effects of a few operations, including language features that don't make much sense in the Layer 2 context.

- `tx.gasprice` returns the user’s ArbGas price bid
- `blockhash(x)` returns Arbitrum blockhash, a value deterministically generated from the state of the inbox. Note that the Arbitrum blockhash is _not_ dependent on the L1 blockhash, and thus should not be assumed to have any economic security (as a randomness seed, say)
- `block.coinbase` returns zero
- `block.difficulty` returns the constant 2500000000000000
- `block.gaslimit` returns the block's ArbGas limit
- `gasleft` returns the amount of ArbGas remaining
- `block.number` on a non-Sequencer Arbitrum chain, returns the L1 block number at which the transaction was submitted to the inbox; on a Sequencer Arbitrum chain, returns an "estimate" of the L1 block number at which the Sequencer received the transaction (see [Time in Arbitrum](Time_in_Arbitrum.md))
- `msg.sender` works the same way it does on Ethereum for normal L2-to-L2 transactions; for L1-to-L2 "retryable ticket" transactions, it will return the L2 address alias of the L1 contract that triggered the message. See [retryable ticket address aliasing](L1_L2_Messages.md#address-aliasing) for more.

# Time

Arbitrum supports `block.number` and `block.timestamp`. For the semantics of these features in the Arbitrum context, please see [Time in Arbitrum](Time_in_Arbitrum.md).
