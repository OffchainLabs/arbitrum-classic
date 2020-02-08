---
id: Tx_Lifecycle
title: How a transaction call executes on Arbitrum
sidebar_label: Lifecycle of a Tx call
---

When a client makes a call to a contract running on Arbitrum, and gets a result back, here is what happens inside Arbitrum:

1. Client software submits a transaction to Ethereum specifying the Arbitrum transaction call it wants to make
2. The client's transaction gets into an Ethereum block. This puts the call into the Arbitrum chain's Inbox.
3. One of the chain's validators sees the call arriving in the Inbox and builds an assertion in which the Arbitrum Chain receives the call message and executes the call (probably in a batch with a bunch of other call messages).
4. The validator submits an Ethereum transaction containing the assertion
5. The assertion transaction appears in an Ethereum block and the chain's validators see it.
6. The chain's validators, having verified the correctness of the assertion, extract the transaction result.
7. A validator serves out the transaction result to the client.
