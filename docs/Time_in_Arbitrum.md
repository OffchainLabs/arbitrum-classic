---
id: Time_in_Arbitrum
title: Time in Arbitrum
sidebar_label: Time in Arbitrum
---

As in Ethereum, Arbitrum clients submit transactions they want to see happen, and the system (usually) executes those transactions at some later time.
In Arbitrum Rollup, clients submit transactions by posting messages to the Ethereum chain, either directly or through an aggregator.
These messages are put into the chain's _inbox_.
Every message is timestamped with the time it was put into the inbox.

Messages in an ArbChain's inbox are processed in order. Generally, some time will elapse between the time when a message is put into the inbox (and timestamped) and the time when the contract processes the message and carries out the transaction requested by the message.

Generally Arbitrum will execute a transaction based on the time it was submitted into the inbox.

If your Solidity contract, running on Arbitrum Rollup, accesses `block.number` or `block.timestamp`, this will give you the block or timestamp that the message entered the inbox respectively. These values will only ever increase since they are marked by Ethereum in order.
