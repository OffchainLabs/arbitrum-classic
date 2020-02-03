This document explains the basic concepts behind Arbitrum Rollup.

## Arbitrum Rollup Chains (ArbChains)

An ArbChain is an environment where smart contracts can run. ArbChains run off-chain, giving them higher performance
and lower cost than on-chain Ethereum contracts. 
Despite running off-chain, ArbChains are *trustless*, meaning that any
party can force an ArbChain to execute correctly according to the code of its contracts.

You can make as many ArbChains as you want. 
Typically, you'll want one ArbChain for each dapp you want to create.

Every ArbChain has a set of Validators, who monitor the chain to ensure that it executes correctly.
Validators will deposit currency stakes, which they will lose if they behave dishonestly.

An ArbChain contains a set of contracts. 
Over time, you can start new contracts in an ArbChain, and contracts can self-destruct.
There is one limitation, though: you have to provide the code for a contract when you initially build the chain. 
You can start as many instances of those contracts as you want, 
but you can't introduce new contract code into a chain after the chain has started.
(This limitation will probably be lifted in future versions.)

## How to make an ArbChain

There are two steps required to start an ArbChain.

First, you compile your Solidity contracts, using the open-source Arbitrum compiler. 
This will bundle your contracts together, along with some runtime and "glue" code, 
into an Arbitrum executable file (or ".ao file"). 
The .ao file is a single program, to be run on the Arbitrum Virtual Machine (AVM) architecture, 
which will manage and run the contracts and wallets needed for your chain. 
Think of it as an ArbChain in a box.

Next, you start your ArbChain by making a call to an Ethereum contract called the EthBridge.  
The EthBridge will start an Ethereum contract to manage your chain, 
and some parameters will be recorded on the main Ethereum chain.

Now your ArbChain is ready to go.

## Validators for an ArbChain

The actual work of monitoring an ArbChain and advancing its contracts is done by validators. 
Each ArbChain has its own validators, and anyone can be a validator for any ArbChain.

The software needed to be a validator is available for free from OffChain Labs, in source code or packaged as a Docker image.

A validator keeps track of the state of the ArbChain it is validating.
If another validator tries to do something dishonest, an honest validator will challenge the dishonest action.
The challenge will be refereed by an on-chain Ethereum contract, and eventually the honest party will win the challenge and take part of the dishonest validator's stake.

The Arbitrum protocol ensures that an honest party can always win challenges. 
This deters dishonesty, and it ensures that any one honest party can force correct behavior, even if everyone else is dishonest.
That's what makes ArbChains trustless.

### Who will validate your ArbChain?

At this point, you might be wondering who will be validators of your ArbChain. 
Arbitrum allows anyone to validate any ArbChain, but who will step up and do it?

As a dapp developer, you would probably want to validate your dapp's ArbChain, to help ensure its correctness and show confidence in your dapp.
Some users will want to validate as well, in order to protect their interest in the dapp's correct execution.

Once Arbitrum is released for production, you'll be able to hire a validator-as-a-service to validate your ArbChain.
You might want to hire more than one, if you or your users worry that one might misbehave.

## How clients interact with contracts

Arbitrum client programs can use the same client frameworks as on Ethereum, including web3, ethers, and go-ethereum.
Offchain Labs provides plug-ins for these frameworks, to make them work with Arbitrum. 
The details in this section are all handled by the plug-ins, so you don't have to worry about them.
But read on, if you're interested in how things work.

When a client wants to make a call to a contract in an ArbChain, the client creates an Arbitrum message, 
and makes an Ethereum call to the on-chain Pending Inbox contract, to insert the message into the ArbChain's pending inbox.

Inside the ArbChain, the program running on the chain's Arbitrum Virtual Machine will consume messages from the chain's
pending inbox in order.
It will interpret each message as a call to a contract, and it will execute that call, updating the contract's state.
At the end of the call, the program will emit an Arbitrum log item, which will become visible to the client program.
The client program will then read the call's result from the log item.

## Other topics

There are some other topics that have documentation pages of their own:
* [Eth and tokens in Arbitrum](Payments_and_tokens.md)
* [Finality and confirmation in Arbitrum](Finality.md)
