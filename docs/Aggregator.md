---
id: Aggregator
title: Aggregators in Arbitrum
sidebar_label: Aggregators
---

An aggregator plays the same role that a node plays in Ethereum. Client software can do remote procedure calls (RPCs) to an aggregator, using the standard API, to interact with an Arbitrum chain. The aggregator will then make calls to the EthBridge and produce transactions results to the client, just an an Ethereum node would.

Most clients will use an aggregator to submit their transactions to an Arbitrum chain, although this is not required. There is no limit on how many aggregators can exist, nor on who can be an aggregator.

To improve efficiency, aggregators will usually package together multiple client transactions into a single message to be submitted to the Arbitrum chain.
