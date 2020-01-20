# Confirmation and Finality in Arbitrum Rollup

As in Ethereum, transactions in Arbitrum don't become final right away. 
Validators can make conflicting claims about which behaviors of an ArbChain are correct, 
and the system will take some time to sort out who is correct.

In effect, each validator is vouching for a particular future history of the ArbChain.
If you're a client, and your transaction appears in one proposed history and not in another, then you need to know which history the system will choose.

Clients need to wait until they can have enough confidence in a transaction's acceptance before assuming that it is written in stone.

In Arbitrum there are several flavors of confirmation that a client can choose to rely on.

### On-chain confirmation

The strongest type of confirmation is on-chain confirmation. 
This is when an ArbChain's on-chain manager records the transaction's confirmation. 
Once the confirmation is solidly recorded on-chain, it will never be undone.

On-chain confirmation is typically slower than the other types.

### Validators' stakes

The ArbChain's validators will have place stakes on proposed future histories of the chain. 
Arbitrum enforces two rules regarding those stakes:
* If a validator is staked on a correct future history, it can force that history to be chosen by the system eventually.
* If a validator is staked on a history that is not chosen, it will lose its stake.

### If you are a validator

If you're a validator, you will know which of the pending claims about the ArbChain's future is correct.
And if you put a stake on the correct future, you can force its eventual on-chain confirmation.
If anybody tries to dispute your correctness, you'll be able to take their stakes.

So if you're a validator, and your transaction is part of the correct future, you can be sure the transaction will be confirmed.
Of course, you can always become a validator, if you want to make sure your correct transactions will be confirmed

### Other validator(s) staked on the transaction

If some validator is staked on a history that includes your transaction, then you know that that validator is standing behind the transaction.
The validator has staked its money on your transaction: it will definitely lose its stake if your transaction isn't eventually confirmed on-chain.

If multiple validators are staked on your transaction, so much the better.
Either your transaction will be confirmed, or they will all lose their stakes.

For many clients, this will be enough--they'll be satisfied to rely on the staked positions of validators.

### People vouching for the transaction

The weakest form of evidence for a pending transaction is that people you trust say it is on the correct future history.
If they aren't staked on that future, they have nothing to lose if they're lying, except possibly their reputation.
But if you strongly trust them, that might be enough.
