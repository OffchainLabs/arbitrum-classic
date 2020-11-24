---
id: Dispute_Resolution
title: Making Progress and Dispute Resolution
sidebar_label: Progress and Dispute Resolution
---

## How Arbitrum contracts make progress: the common case

A contract running on Arbitrum is inside an Arbitrum VM, which is associated with a particular Arbitrum chain. Inside a chain’s VM, [ArbOS](ArbOS.md) picks up messages from the chain’s inbox, dispatches them to the appropriate contract, and publishes results. So contracts make progress when their chain’s VM makes progress.

To advance a chain’s VM, any validator of the Arbitrum chain can make a disputable assertion: a claim, submitted to the Ethereum chain via the EthBridge, that asserts what the correct next state hash and actions of the VM are. After this disputable assertion, the other validators have a time window to dispute the assertion if they believe it incorrect. If the assertion is correct, it’s almost certain that nobody will dispute it as this will run counter to their incentives and incur a significant penalty. After the dispute window closes, the EthBridge will accept the assertion as correct, allowing the VM to make progress.

## How a chain advances: the malicious case

The challenging case arises when a dishonest validator tries to cheat. There are two things a dishonest validator might try to do. First, they might try to corrupt the execution of the chain by making a disputable assertion that is false in the hopes that the EthBridge will accept that assertion. Second, they might let somebody else make a truthful disputable assertion and then start a bogus dispute about it.

Arbitrum deters misbehavior by requiring validators who make or challenge an assertion to put down a currency stake that is held by the EthBridge. If a validator gets caught cheating, it will forfeit its stake. (Half of the cheater’s stake goes to the other party in the dispute. The other half is burned. Validators who haven’t been caught cheating can get their stake back, once the challenge period for the thing they’re staked on has passed.)

Arbitrum’s dispute resolution protocol is designed to resolve disputes very efficiently, by identifying a cheating party while requiring a minimum of on-chain Ethereum activity. The dispute plays out as a contest between an asserter (who claims their disputable assertion is correct) and a challenger (who claims that the same assertion is wrong), with the EthBridge acting as referee.

## Dispute resolution

The dispute resolution protocol goes in two stages. First, the players narrow down their disagreement, using an on-chain bisection protocol, until they are disagreeing about a single step of the VM’s computation. Then the asserter sends a one-step proof—a proof of correctness for the execution of a single VM instruction—to the EthBridge which checks that tiny proof.

The bisection protocol starts when the asserter claims that a VM starting with state-hash X can execute N instructions, resulting in state-hash Y, and the challenger responds that this is false.
The asserter is required to divide their assertion into K assertions, each involving the execution of N/K instructions, which fit together to give the initial assertion. Then the challenger has to pick one of those K smaller assertions to challenge. If either player fails to act within a specified time limit measured in a number of blocks, they lose. After one round of K-way division, the size of the dispute has been cut to N/K steps.

The process continues with further divisions. After a logarithmic number of divisions (logarithmic in the number of instructions executed, N), the dispute has been narrowed to a single step: the dispute will be over whether or not a VM starting with state-hash Y can execute a single instruction to get to the state-hash Z.

At this point the asserter has to give a one-step proof to the EthBridge, containing the information that the EthBridge needs to quickly verify that the one-step assertion is correct. Or perhaps the asserter will fail to provide a valid one-step proof. Either way, the dispute is resolved.

![img](assets/Arbitrum_dispute.png)

We won’t go into the details of one-step proofs here. Suffice it to say that [Arbitrum’s custom VM architecture](AVM_Design.md) comes into play here, making the one-step proof small (a few hundred bytes at most) and quick to check (costing about 90,000 Ethereum gas, or about one dollar at current prices).

By making disputes relatively cheap to resolve, and imposing a substantial penalty on the loser, Arbitrum strongly disincentivizes attempts to cheat, but even if a dispute occurs this doesn’t impose a huge on-chain impact. In the common case, validators will agree and progress will occur off-chain.
