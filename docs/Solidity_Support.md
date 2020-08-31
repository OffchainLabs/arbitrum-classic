---
id: Solidity_Support
title: Solidity Support
custom_edit_url: https://github.com/OffchainLabs/arbitrum/edit/master/docs/Solidity_Support.md
---

Arbitrum Rollup supports EVM transactions, and therefore allows you to trustlessly deploy Solidity contracts (as well as Vyper or any other language that compile to EVM). Arbitrum supports almost all Solidity code as expected with a few exceptions that we detail below.

# Restrictions

Although we support most Solidity code, we have a few restrictions, including language features that don't make much sense in the Layer 2 context.

- Unsupported Solidity Features:

  - `blockhash(uint blockNumber) returns (bytes32)`
  - `block.coinbase`
  - `block.difficulty`
  
  
  # Time
  
  Abitrum supports `block.number` and `block.timestamp`. For the semantics of these features in the Arbitrum context, please see [Time in Arbitrum](Time_in_Arbitrum.md).
