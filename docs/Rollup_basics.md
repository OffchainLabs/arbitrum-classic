---
id: Rollup_basics
title: Arbitrum Rollup Basics
sidebar_label: Arbitrum Rollup Basics
---

This document explains the basic concepts behind Arbitrum Rollup and provides links for learning more about the various system components..

## Arbitrum Rollup Chains (ArbChains)

An ArbChain is an environment where smart contracts can run. ArbChains run off-chain, giving them higher performance
and lower cost than on-chain Ethereum contracts.
Despite running off-chain, ArbChains are _trustless_, meaning that any
party can force an ArbChain to execute correctly according to the code of its contracts.

You can make as many ArbChains as you want. Each ArbChain provides the same interface as the Ethereum blockchain, and contracts can be deployed on the chain at anytime using the `CREATE` and `CREATE2` opcode. You can also deploy contracts by sending transactions to address zero on the chain, just as you would on Ethereum.
> Note: Earlier versions of Arbitrum did not fully support dynamic contract launching and required code for all contracts to be provided when initially launching the chain. We have now removed this limitation.


Applications that reside on the same ArbChain can make synchronous calls to one another exactly as they would on Ethereum.
Typically, you'll want to group dApps together into a single ArbChain if they need to interact and compose with one another. Applications that do not require synchronous interaction with others can be deployed on their own ArbChain.

Every ArbChain has a set of Validators, who monitor the chain to ensure that it executes correctly.
Validators will deposit currency stakes, which they will lose if they behave dishonestly. Validating an Arbitrum Rollup chain is permissionless; anyone who puts down a stake can serve as a validator.

An ArbChain contains a set of contracts.
Over time, you can launch new contracts in an ArbChain, and contracts can self-destruct, exactly as they do on the Ethereum blockchain.
Arbitrum supports EVM, and therefore supports Solidity contracts. See [Solidity Support](Solidity_Support.md) for a list of the few Solidity features which Arbitrum does not support.


## How to make an ArbChain


Suppose you’re a developer who has written a dApp for the Ethereum platform. Arbitrum interoperates with Ethereum, so you can launch your dapp on an Arbitrum Rollup chain and get better speed and scalability. 

You’re starting with a dapp—or you’re planning to develop one—that’s made up of some contracts written in Solidity, along with a browser-based front end.  Here’s how to use Arbitrum with your dapp. Arbitrum supports dynamic launching of contracts on deployed chains just as in Ethereum by sending transactions to address zero or by using the `CREATE` and `CREATE2` opcodes. Arbitrum is fully compatible with standard build tools (e.g. Truffle, Buidler). You can either launch your dApp on a new Rollup or deploy it to an existing one. Here, we describe the process for launching a new chain.

First, you’ll want to identify an initial set of validators for your chain. We’ll talk later about how you might choose validators, and why people might want to validate a chain. Of course, validators will be able to come and go at will once the chain is going.  

When you’re ready to launch your chain, you send an Ethereum transaction to the EthBridge—the component that connects Arbitrum to Ethereum—telling it to launch your chain on Arbitrum Rollup, and specifying [chain parameters](Chain_parameters.md).  The EthBridge will start an Ethereum contract to manage your chain, and some parameters will be recorded on the main Ethereum chain.

Your Rollup chain is now up and running on Arbitrum. You can deploy contracts to it, by sending the same transaction that you would send to Ethereum if you wanted to deploy there. Users of your dapp can launch your existing front-end interface in their browsers. The front end will automatically interact with the running chain using Arbitrum’s front-end plug-ins for web3, ethers, or go-ethereum.

Thanks to the EthBridge, you and your users can interact with your chain, or send ether or any other Ethereum-based token to contracts on your chain, and the dapp can send ether or tokens to any Ethereum user or to any contract on Arbitrum or Ethereum. You do this by depositing funds into your Arbitrum wallet, which is managed by the EthBridge.

## Validators for an ArbChain

The actual work of monitoring an ArbChain and advancing its contracts is done by validators.
Each ArbChain has its own validators, and anyone can be a validator for any ArbChain.

The software needed to be a validator is available for free from Offchain Labs, in source code or packaged as a Docker image.

A validator keeps track of the state of the ArbChain it is validating.
If another validator tries to do something dishonest, an honest validator will challenge the dishonest action.
The challenge will be refereed by an on-chain Ethereum contract, and eventually the honest party will win the challenge and take part of the dishonest validator's stake.

The Arbitrum protocol ensures that an honest party can always win challenges.
This deters dishonesty, and it ensures that any one honest party can force correct behavior, even if everyone else is dishonest.
That's what makes ArbChains trustless.

For more information on how validators advance a chain and resolve disputes, see [Progress and Dispute Resolution](Dispute_Resolution.md).

### Who will validate your ArbChain?

At this point, you might be wondering who will be validators of your ArbChain.
Arbitrum allows anyone to validate any ArbChain, but who will step up and do it?

Although validation is permissionless and anyone _can_ validate, it's important to make sure that each ArbChain always has at least one honest validator that is paying attention. As a dApp developer, you would probably want to validate your dApp's ArbChain, to help ensure its correctness and show confidence in your dApp.

For developers who do not want to validate themselves, you'll be able to hire a validator-as-a-service to validate your ArbChain once Arbitrum is released for production.  
You might want to hire more than one, if you or your users worry that one might misbehave. As long as any validator behaves honestly, the Arbitrum protocol guarantees that all contracts in the ArbChain will run correctly according to their code. 

Some users will want to validate as well, in order to protect their interest in the dapp's correct execution. Users may also hire validators that they trust to ensure the chains correct operation. A healthy ArbChain with an engaged userbase will likely have a variety of validators.



## How clients interact with contracts

Arbitrum client programs can use the same client frameworks as on Ethereum, including web3, ethers, and go-ethereum.
Offchain Labs provides [plug-ins for these frameworks](Frontend_Integration.md), to make them work with Arbitrum.
The details in this section are all handled by the plug-ins, so you don't have to worry about them.
But read on, if you're interested in how things work.

When a client wants to make a call to a contract in an ArbChain, the client creates an Arbitrum message,
and makes an Ethereum call to the on-chain Inbox contract, to insert the message into the ArbChain's inbox.

Inside the ArbChain, the program running on the chain's Arbitrum Virtual Machine will consume messages from the chain's inbox in order.
It will interpret each message as a call to a contract, and it will execute that call, updating the contract's state.
At the end of the call, the program will emit an Arbitrum log item, which will become visible to the client program.
The client program will then read the call's result from the log item.


## The Arbitrum Virtual Machine

Although Arbitrum supports EVM, under the hood it runs the Arbitrum Virtual Machine (AVM). The AVM is optimized for allowing fast progress in the optimistic case while maintaining the ability to efficiently resolve disputes. To learn more, you can read a detailed overview of the [AVM Design Rationale](AVM_Design.md) as well as the [AVM Specification](AVM_Specification.md), a lower level description of the semantics of the AVM architecture.

## ArbOS

ArbOS, the Arbitrum operating system sits atop of the AVM and is responsible for isolating untrusted contracts from one another, tracking and limiting their resource usage using [ArbGas](ArbGas.md), and managing the economic model that collects fees from users to fund the operation of a chain's validators. ArbOS gives Arbitrum a great deal of flexibility by offloading work that would have been done in the L1 smart contract into cheaper L2 code. To learn more, see the section on [ArbOS](ArbOS.md).


## Other topics

There are some other topics that have documentation pages of their own:

- [Ethereum interoperability](Ethereum_Interoperability.md)
- [Finality and confirmation in Arbitrum](Finality.md)
