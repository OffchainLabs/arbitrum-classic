---
id: Time_in_Arbitrum
title: Time in Arbitrum
sidebar_label: Time in Arbitrum
---

Arbitrum supports both block numbers which correspond to block numbers on the main Ethereum chain as well as an approximation of the current timestamp,

Also as on Ethereum, clients submit transactions they want to see happen, and the system (usually) executes those transactions at some later time.
In Arbitrum Rollup, clients submit transactions by posting messages to the Ethereum chain, either directly or through an aggregator.
These messages are put into the chain's _inbox_.
Every message is timestamped with the time it was put into the inbox.

Messages in an ArbChain's inbox are processed in order.
Generally, some time will elapse between the time when a message is put into the inbox (and timestamped) and the time when the contract processes the message and carries out the transaction requested by the message.

If your Solidity contract, running on Arbitrum Rollup, accesses `block.number` or `block.timestamp`, this will give you the ArbChain's current block or timestamp respectively.
Because ArbChains operate off-chain, an ArbChain's block number or timestamp might differ a little bit from the current Ethereum block.
The ArbChain's clock will never get ahead of the Ethereum's, but the ArbChain's clock might "run a bit slow".
A rollup chain can configure the maximum amount its clock can differ, but generally it will stay within 20 blocks or 10 minutes
In particular, a rollup chain can configure how far behind its clock may fall, but generally it will stay within 20 blocks or 10 minutes.

As you would expect, the results given by `block.number` and `block.timestamp` will never decrease; they can only increase.

If you want to know when the current message (the one that requested the transaction you're currently running) arrived, you can get it by the following Soldiity calls:

    ArbSys(address(100)).currentMessageBlock()
    ArbSys(address(100)).currentMessageTimestamp()

At any time, you can get an upper bound on the true Ethereum block number or timestamp by using these Solidity calls:

    ArbSys(address(100)).blockUpperBound()
    ArbSys(address(100)).timestampUpperBound()

This will have the same accurancy guarantee from the ArbChain's configuration described above. For example the `blockUpperBound()` can be configured to never be more than 20 blocks ahead of the current Ethereum block number.
The real Ethereum block number and timestamp will always be between the ArbChain's `block.number`/`block.timestamp` and what `blockUpperBound()`/`timestampUpperBound()` would return (inclusive).

Bear in mind that the ArbChain might "freeze" at any time, so the upper bound is only a real upper bound at the moment that you make that call.
After the call, the Ethereum time might have advanced.

## Use Cases

Most developers will be fine using `block.number` and `block.timestamp`.
If you don't mind a little bit of time lag, you can keep on using `block.number` and `block.timestamp` just like you would on Ethereum.

### Setting deadlines for user response

An exception is a use case like an auction, where the contract emits an event, and you want to give people out there in userland N hours to respond.
To do that, you'll want to "start the clock ticking" by using code like this:

    deadline = ArbSys(address(100)).timestampUpperBound() + N hours;
    emit StartAuction(..., deadline);

You'll presumably have an `endAuction` call, which can only be called by a transaction submitted to the inbox after the deadline.
To ensure that, you can use code like this:

    function endAuction(...) public {
        require(ArbSys(address(100)).timestampUpperBound() > deadline, "Auction has not ended yet");
