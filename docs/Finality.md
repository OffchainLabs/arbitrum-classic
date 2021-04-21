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

_In the current Arbitrum beta version, clients use the third option by default, relying on the aggregator they're connected to._  
In later versions, the client software will be able to choose which model to rely on.

### On-chain confirmation of your transaction

The strongest type of confirmation is on-chain confirmation.
This is when an ArbChain's on-chain EthBridge records the transaction's confirmation.
Once the confirmation is solidly confirmed on-chain, it will never be undone.

On-chain confirmation is typically slower than the other types.

### Validators staked on your transaction

The ArbChain's validators will place stakes on proposed future histories of the chain.
Arbitrum enforces two rules regarding those stakes:

- If a validator is staked on a correct future history, it can force that history to be confirmed by the system eventually.
- If a validator is staked on a history that is not confirmed, that validator will lose its stake.

You can choose to treat a transaction as final when a validator or validators who you trust are staked on a history containing that transaction.

If you are a validator, then you can rely on yourself.
If your transaction is part of the correct future, you can be sure the transaction will be confirmed.
Of course, you can always become a validator, if you want to make sure your correct transactions will be confirmed.

### Validators or aggregators vouching for the transaction, without staking

The weakest form of evidence for a pending transaction is that someone you trust says it is on the correct future history.
If they aren't staked on that future, they will not directly be penalized if they're lying, except possibly their reputation.
But if you strongly trust them, that might be enough.

### Finality in Sequencer Mode

For an Arbitrum chain in Sequencer mode, a Sequencer can provide fast, semi-trusted off-chain confirmations to users. Even a malicious Sequencer will _not_ be able to get away with an invalid state update (just like any other validator); a Sequencer can, however, fail to include a transaction after if it was promised, or reorder transactions over a short window of time. Cryptoeconomic penalties for Sequencer misbehavior can be enforced (i.e., a Sequencer posts a bond that gets slashed if it violates its promises), but these security mechanisms run orthogonal to those of Arbitrum itself.

Once the Sequencer publishes a user's transaction on chain, the transactions is secured by the Arbitrum protocol, and a client accepts confirmation according to one the three models outlined above. See [Sequencer Mode](Inside_Arbitrum#sequencer-mode.md) for more info.

## Withdrawals

One aspect of the system that is closely related to finality is the ability to withdraw funds from the ArbChain back onto Ethereum. This is discussed fully in the [withdrawals](Withdrawals.md) section of the documentation.
