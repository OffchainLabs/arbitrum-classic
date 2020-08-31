---
id: Tx_Lifecycle
title: How a transaction call executes on Arbitrum
sidebar_label: Lifecycle of a Tx call
---

When a client makes a call to a contract running on Arbitrum, and gets a result back, here is what happens inside Arbitrum:

1. Client software make a remote procedure call (RPC) to an Arbitrum aggregator, asking the aggregator to submit a transaction.
2. The aggregator optionally packages this user's transaction together with user others' transactions. The user's transaction will get packaged into some Arbitrum message.
3. The aggregator calls the EthBridge to send the Arbitrum message to ArbOS on the Arbitrum chain. 
4. The message gets into an Ethereum block. This puts the message into the Arbitrum chain's Inbox.  At this point the result of the transaction is fully determined.
5. The aggregator sends the transaction result back to the client.

At this point the result of the transaction has been fully determined, but further steps need to occur to get the result recognized and fully confirmed.

1. One of the chain's validators sees the call arrive in the Inbox and builds an assertion in which ArbOS, running on the Arbitrum Chain, receives the call message and dispatches the call to the called contract's code, the contract's code executes, and ArbOS emits a transaction receipt.
2. The validator submits to the EthBridge an Ethereum transaction containing the assertion.
3. The assertion transaction appears in an Ethereum block and the chain's validators see it and validate that it is correct.
4. Eventually the assertion is accepted by the Arbitrum Rollup Protocol.  At this point any withdrawals or other L1 side-effects of the transaction occur at L1.
