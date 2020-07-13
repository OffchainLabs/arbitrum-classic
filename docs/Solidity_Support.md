---
id: Solidity_Support
title: Solidity Support
custom_edit_url: https://github.com/OffchainLabs/arbitrum/edit/master/docs/Solidity_Support.md
---

Arbitrum Rollup allows you to deploy a set of Solidity contracts as on a trustless layer 2 sidechain. Arbitrum supports almost all solidity code as expected with a few exceptions.

# Restrictions

Although we support most solidity code, there are a number of restrictions that currently exist.

- Unsupported Solidity Features:

  - `blockhash(uint blockNumber) returns (bytes32)`
  - `block.coinbase`
  - `block.difficulty`
