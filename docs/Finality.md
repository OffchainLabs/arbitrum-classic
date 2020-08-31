---
id: Finality
title: Confirmation and finality in Arbitrum Rollup
sidebar_label: Confirmation and finality
---

As in Ethereum, transactions in Arbitrum don't become final right away.
If there are malicious validators, they can make conflicting claims about which behaviors of an ArbChain are correct,
and the system will take some time to sort out who is correct.

In effect, each validator is vouching for a particular future history of the ArbChain.
If you're a client, and your transaction appears in one proposed history and not in another, then you need to know which history the system will choose.

Clients need to wait until they can have enough confidence in a transaction's acceptance before assuming that it is written in stone.

There are three models of confirmation that a client can choose to rely on.

_In Arbitrum Rollup Beta 1, clients use the second option, relying on the validator they're connected to._  
In later versions, the client software will be able to choose which model to rely on.

### On-chain confirmation of your transaction

The strongest type of confirmation is on-chain confirmation.
This is when an ArbChain's on-chain manager records the transaction's confirmation.
Once the confirmation is solidly recorded on-chain, it will never be undone.

On-chain confirmation is typically slower than the other types.

### Validators staked on your transaction

The ArbChain's validators will all place stakes on proposed future histories of the chain.
Arbitrum enforces two rules regarding those stakes:

- If a validator is staked on a correct future history, it can force that history to be confirmed by the system eventually.
- If a validator is staked on a history that is not confirmed, that validator will lose its stake.

You can choose to treat a transaction as final when a validator of your choice is staked on a history containing that transaction.

If you are a validator, then you can rely on yourself.
If your transaction is part of the correct future, you can be sure the transaction will be confirmed.
Of course, you can always become a validator, if you want to make sure your correct transactions will be confirmed.

### Validators or observers vouching for the transaction, without staking

The weakest form of evidence for a pending transaction is that people you trust say it is on the correct future history.
If they aren't staked on that future, they have nothing to lose if they're lying, except possibly their reputation.
But if you strongly trust them, that might be enough.


## Withdrawals

One aspect of the system that is closely related to finality is the ability to withdraw funds from the ArbChain back onto Ethereum. This is discussed fully in the [withdrawals](Withdrawals.md) section of the documentation.
