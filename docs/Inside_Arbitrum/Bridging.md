---
id: Bridging
title: Inside Arbitrum: Bridging Between L1 and L2
sidebar_label: L1-L2 Bridging
---
We have already covered how users interact with L2 contracts--they submit transactions by putting messages into the chain’s inbox, or having a full node aggregator do so on their behalf. Let’s talk about how contract interact between L1 and L2--how an L1 contract calls an L2 contract, and vice versa.

The L1 and L2 chains run asynchronously from each other, so it is not possible to make a cross-chain call that produces a result within the same transaction as the caller. Instead, cross-chain calls must be asynchronous, meaning that the caller submits the call at some point in time, and the call runs later. As a consequence, a cross-chain contract-to-contract call can never produce a result that is available to the calling contract (except for acknowledgement that the call was successfully submitted for later execution).

## L1 contracts can submit L2 transactions

An L1 contract can submit an L2 transaction, just like a user would, by calling the EthBridge. This L2 transaction will run later, producing results that will not be available to the L1 caller. The transaction will execute at L2, at the same Ethereum block number and timestamp as the L1 caller, but the L1 caller won’t be able to see any results from the L2 transaction.

The advantage of this method is that it is simple and has relatively low latency. The disadvantage, compared to the other method we’ll describe soon, is that the L2 transaction might revert if the L1 caller doesn’t get the L2 gas price and max gas amount right. Because the L1 caller can’t see the result of its L2 transaction, it can’t be absolutely sure that its L2 transaction will succeed.

This would introduce a serious a problem for certain types of L1 to L2 interactions. Consider a transaction that includes depositing a token on L1 to be made available at some address on L2. If the L1 side succeeds, but the L2 side reverts, you've just sent some tokens to the L1 inbox contract that are unrecoverable on either L2 or L1. Not good.

## L1 to L2 ticket-based transactions

Fortunately, we have another method for L1 to L2 calls, which is more robust against gas-related failures, that uses a ticket-based system. The idea is that an L1 contract can submit a “pre-packaged” transaction and immediately receive a “ticketID” that identifies that submission. Later, anyone can call a special pre-compiled contract at L2, providing the ticketID, to try redeeming the ticket and executing the transaction.

The pre-packaged transaction includes the sender’s address, a destination address, a callvalue, and calldata. All of this is saved, and the callvalue is deducted from the sender’s account and (logically) attached to the pre-packaged transaction.

If the redemption succeeds, the transaction is done, a receipt is issued for it, and the ticketID is canceled and can’t be used again. If the redemption fails, for example because the packaged transaction fails, the redemption reports failure and the ticketID remains available for redemption.

As an option (and what we expect to be the default), the original submitter can try to redeem their submitted transaction immediately, at the time of its submission, in the hope that this redemption will succeed. Thus, i.e., our "token deposit" example above should, in the happy, common case, still only require a single signature from the user. If this instant redemption fails, the ticketID will still exist which others can redeem later.

Submitting a transaction in this way carries a price in ETH which the submitter must pay, which varies based on the calldata size of the transaction. Once submitted, the ticket is valid for about a week. It will remain valid, as long as someone pays weekly rent to keep it alive. If the rent goes unpaid and the ticket has not been redeemed, it is deleted.

When the ticket is redeemed, it runs with sender and origin equal to the original submitter, and with the destination, callvalue, and calldata the submitter provided at the time of submission. (The submitter can specify an address to which the callvalue will be refunded if the transaction is dropped for lack of rent without ever being redeemed.)

This rent-based mechanism is a bit more cumbersome than direct L1 to L2 transactions, but it has the advantage that the submission cost is predictable and the ticket will always be available for redemption if the submission cost is paid. As long as there is some user who is willing to redeem the ticket (and pay rent if needed), the L2 transaction will eventually be able to execute and will not be silently dropped.

## L2 to L1 ticket-based calls

Calls from L2 to L1 operate in a similar way, with a ticket-based system. An L2 contract can call a method of the precompiled ArbSys contract, to send a transaction to L1. When the execution of the L2 transaction containing the submission is confirmed at L1 (some days later), a ticket is created in the EthBridge. That ticket can be renewed by anyone who calls a certain EthBridge method and submits the ticketID. 

These tickets have unlimited lifetime, until they’re successfully redeemed. No rent is required, as the costs are covered by network fees that are collected elsewhere in Arbitrum.