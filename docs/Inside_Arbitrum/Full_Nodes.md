---
id: Full_Nodes
title: Inside Arbitrum: Full Nodes
sidebar_label: Full Nodes
---
As the name suggests, full nodes in Arbitrum play the same role that full nodes play in Ethereum: they know the state of the chain and they provide an API that others can use to interact with the chain.

Arbitrum full nodes operate ["above the line"](Above_or_Below_the_Line.md.md), meaning that they don’t worry about the rollup protocol but simply treat their Arbitrum chain as a machine consuming inputs to produce outputs. A full node has a built-in AVM emulator that allows it to do this.

## Batching transactions: full node as an aggregator

One important role of a full node is serving as an aggregator, which means that the full node receives a set of signed transactions from users and assembles them into a batch which it submits to the chain’s inbox as a single unit. Submitting transactions in batches is more efficient than one-by-one submission, because each submission is an L1 transaction to the EthBridge, and Ethereum imposes a base cost of 21,000 L1 gas per transaction. Submitting a batch allows that fixed cost (and some other fixed costs) to be amortized across a larger group of user transactions. That said, submitting a batch is permissionless, so any user can, say, submit a single transaction "batch" if the need arises; Arbitrum thereby inherits the same censorship resistance as the Ethereum.

## Compressing transactions

In addition to batching transactions, full nodes can also compress transactions to save on L1 calldata space. Compression is transparent to users and to contracts running on Arbitrum--they see only normal uncompressed transactions. The process works like this: the user submits a normal transaction to a full node; the full node compresses the transactions and submits it to the chain; ArbOS receives the transaction and uncompresses it to recover the original transaction; ArbOS verifies the signature on the transaction and processes the transaction.

Compression can make the “header information” in a transaction much smaller. Full nodes typically use both compression and batching when submitting transactions. 

## Aggregator costs and fees

An aggregator that submits transactions on behalf of users will have costs due to the L1 transactions it uses to submit those transactions to the Inbox contract. Arbitrum includes a facility to collect fees from users, to reimburse the aggregator for its costs. This is detailed in the [ArbGas and Fees](ArbGas_and_Fees.md) section.