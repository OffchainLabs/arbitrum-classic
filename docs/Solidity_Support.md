---
id: Solidity_Support
title: Solidity Support
custom_edit_url: https://github.com/OffchainLabs/arbitrum/edit/master/docs/Solidity_Support.md
---

Arbitrum Rollup supports EVM transactions, and therefore allows you to trustlessly deploy Solidity contracts (as well as Vyper or any other language that compile to EVM). Arbitrum supports almost all Solidity code as expected with a few exceptions that we detail below.

# Differences from Solidity on Ethereum

Although Arbitrum supports Solidity code, there are differences in the effects of a few operations, including language features that don't make much sense in the Layer 2 context.

- `tx.gasprice` will return 1
- `blockhash(x)` will always return zero
- `block.coinbase` will return zero
- `block.difficulty` will return the constant 2500000000000000
- `block.gaslimit` will return the block's ArbGas limit
- `gasleft` will return the amount of ArbGas remaining

# Time

Arbitrum supports `block.number` and `block.timestamp`. For the semantics of these features in the Arbitrum context, please see [Time in Arbitrum](Time_in_Arbitrum.md).
