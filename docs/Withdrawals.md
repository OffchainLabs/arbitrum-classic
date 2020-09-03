---
id: Withdrawals
title: Withdrawing funds from Arbitrum
sidebar_label: Withdrawals
---

As discussed in the section on [finality](Finality.md), because Arbitrum's execution happens optimistically, the Ethereum blockchain cannot immediately confirm the correct state and must wait for the challenge window to expire (or until all challenges are resolved). 

When it comes to execution within the Rollup this does not pose any problem or add any delay. The [Arbitrum Rollup protocol](Rollup_Protocol.md) manages a tree of assertions, and allows validators to pipeline execution by continuing to build the tree even before all nodes are confirmed. This means that an honest validator can continue to advance the state of the machine with confidence (and the ability to enforce) that eventually Ethereum will recognize the honest branch as the correct and valid one. And although it will take some time for Ethereum to recognize which branch is correct, anyone that is validating the chain will immediately know this. Anyone that is validating the chain will immediately know which branch is correct and therefore which branch will eventually be accepted by the protocol.

The one part of the protocol that is affected by the confirmation delay is withdrawals. Since Arbitrum cannot undo a withdrawal once it has released funds from the L2, the system cannot allow funds to be withdrawn from the ArbChain until it has been confirmed on the Ethereum chain that the withdrawal is valid. 

## Liquidity Exits

Recall that although it takes time for the EthBridge to confirm a withdrawal, any validating node can immediately detect whether a withdrawal is legitimate or not. This creates an opportunity for liquidity providers to lend money for a fee against a pending withdrawal. The user will be able to access their funds (minus the fee) right away, and the liquidity provider will be paid back once the system has confirmed and processed the withdrawal.

## NFTs and Messages

Liquidity exits work well for fungible tokens. However for non-fungible tokens for which a liquidity provider cannot lend an equivalent substitute, withdrawing will still incur the delay for system confirmation. Simialrly, when one wants the ArbChain to post a message to the L1 (e.g. that will be processed as an asynchronous call by another contract), they will also incur the confirmation delay.
