# Time in Arbitrum Rollup

As on Ethereum, Arbitrum doesn't use real wall-clock time, but instead uses block numbers, which correspond to block numbers on the main Ethereum chain.

Also as on Ethereum, clients submit transactions they want to see happen, and the system (usually) executes those transactions at some later time.
In Arbitrum Rollup, clients submit transactions by posting messages to the Ethereum chain.
These messages are put into the chain's *pending inbox*.
Every message is timestamped with the time it was put into the pending inbox.

Messages in an ArbChain's pending inbox are processed in order.
Generally, some time will elapse between the time when a message is put into the pending inbox (and timestamped) and the time when the contract processes the message and carries out the transaction requested by the message.

If your Solidity contract, running on Arbitrum Rollup, accesses ``block.number``, this will give you the ArbChain's current time. 
Because ArbChains operate off-chain, an ArbChain's time might differ a little bit from the current Ethereum block time.
The ArbChain's clock will never get ahead of the Ethereum block time, but the ArbChain's clock might "run a bit slow". 
In particular, it might be as many as 20 blocks behind the true Ethereum time.

As you would expect, the block time given by `block.number` will never decrease; it can only increase.

If you want to know when the current message (the one that requested the transaction you're currently running) was timestamped, you can get it in Solidity by writing:

    ArbSys(address(100)).currentMessageTime()

At any time, you can get an upper bound on the true Ethereum block time by using this Solidity call:

    ArbSys(address(100)).timeUpperBound()

This will never be more than 20 blocks ahead of ``block.number``.
The real Ethereum block number will always be between the ArbChain's ``block.number`` and what timeUpperBound() would return (inclusive).

Bear in mind that the ArbChain might "freeze" at any time, so the upper bound is only a real upper bound at the moment that you make that call.
After the call, the Ethereum time might have advanced.

## Use Cases

Most developers will be fine using ``block.number``.
If you don't mind a little bit of time lag, you can keep on using ``block.number`` just like you would on Ethereum.

### Setting deadlines for user response
An exception is a use case like an auction, where the contract emits an event, and you want to give people out there in userland N block times to respond.
To do that, you'll want to "start the clock ticking" by using code like this:

    deadline = ArbSys(address(100)).timeUpperBound() + N;
    emit StartAuction(..., deadline);

You'll presumably have an ``endAuction`` call, which can only be called by a transaction submitted to the pending inbox after the deadline.
To ensure that, you can use code like this:

    function endAuction(...) public {
        require(ArbSys(address(100)).currentMessageTime() > deadline, "Auction has not ended yet");
