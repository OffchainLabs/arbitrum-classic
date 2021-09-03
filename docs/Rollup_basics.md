---
id: Rollup_basics
title: Arbitrum Rollup Basics
sidebar_label: Arbitrum Rollup Basics
---

This document explains the basic concepts behind Arbitrum Rollup, an Optimistic Rollup protocol designed for the Ethereum blockchain, and provides links for learning more about the various system components.

## Intro to Rollups

Rollups are one of the hot ideas in smart contract scalability these days. The idea has been around for some time, but several groups, including ours at Offchain Labs, have been promoting it lately. Before discussing Arbitrum Rollup and its unique characteristics, let’s zoom out and talk about what rollup is and what are the common features that any rollup will have.

Rollups are a general approach to scaling open contracts, that is, contracts that everyone can see and interact with. In a rollup, transactions are written on Ethereum (as calldata), but the actual computation and storage of the contract are done off-chain. Somebody (a validator) posts on-chain an assertion (also known as a Rollup block) about what the contract will do — a list of actions taken by the contract, such as payments made, and a cryptographic hash of its state after the contract has executed the calls that have already been posted on-chain. We can think of the assertion as “rolling up” all of the calls and their results into a single on-chain transaction.

Where rollup systems differ is in how they ensure that the assertions are correct.

## Optimistic Rollup

Optimistic Rollup refers to a type of rollup that is optimistic in the sense that when an assertion is posted, it does not contain an accompanying proof guaranteeing its validity. Instead, when the assertion is posted on-chain, the validator making that assertion posts a bond, and there is a time window in which anyone can post their own bond and challenge the assertion, if they think it’s wrong. This is sometimes called a “fraud proof”. If the asserter is wrong, they will lose their bond. If the challenge period expires with no successful challenges, the assertion is accepted and becomes final.

All Optimistic Rollups are interactive--as the asserter first posts an optimistic assertion and another validator can trigger the dispute resolution process by responding with a challenge. However, we can further classify Optimistic Rollups by how many rounds of interaction are required to resolve disputes. In multi-round interactive rollup, like Arbitrum Rollup, there is an initial challenge window during which a challenger can post a bond and claim that the assertion was wrong. What follows is a back-and-forth interactive protocol between the asserter and the challenger, with an on-chain contract acting as a referee for the protocol. In the end the referee determines that one party made a false claim, and punishes that party by taking their bond. The idea is to minimize the amount of on-chain work to resolve the dispute by using an interactive protocol between the two disputants to narrow down the dispute as far as possible before the on-chain referee has to evaluate evidence about the contract’s behavior “on the merits”.

Compared to other Rollup approaches, Arbitrum Rollup's design shines in that the amount of data on chain is quite low, and it can support arbitrary EVM smart contracts and works with all Ethereum developer tooling. In the rest of this page, we'll focus less on how Arbitrum achieves its properties and more on the developer and user experience of Arbitrum Rollup.

## Arbitrum Rollup

An Arbitrum Rollup chain is a super scaled Layer 2 (L2) chain. Like all Rollups, the Arbitrum Rollup chain is built on top of and secured by the Ethereum blockchain, and all transaction data is logged on Ethereum. From a user and developer perspective, interacting with Arbitrum feels exactly like interacting with Ethereum. Arbitrum supports the same RPC interface as Ethereum, supports all EVM languages, and natively supports all Ethereum tooling without any special adapters. The only way in which an Arbitrum Rollup chain does not resemble Ethereum is the cost: transactions on Arbitrum cost a small fraction of what they would if run natively on Ethereum.

Porting contracts from Ethereum to Arbitrum is fast and simple; there's no need to change any code or download any new software. Arbitrum has full support for the EVM just like Ethereum. This means that all smart contract languages that work with Ethereum (e.g. all versions of Solidity, Vyper Yul), also work natively with Arbitrum. See [Solidity Support](Solidity_Support.md) for detailed compatibility information. Similarly all standard frontend Ethereum tooling (e.g. Truffle, Hardhat, The Graph, ethers.js) also work natively with Arbitrum. See [Frontend Integration](Frontend_Integration.md) for more detail. and natively supports all Ethereum tooling.

Although developers and users don't need to download any custom software to deploy contracts and interact with the Arbitrum Rollup chain, some users may want to validate the chain for themselves. When using Arbitrum Rollup, your security is guaranteed by the fact that any single honest user can guarantee that the system runs correctly. Validating the Arbitrum chain is fully permissionless; all you need to do is download the Arbitrum Validator node software and point your node at the chain. To issue or dispute an assertion, you simply need to place a stake that you'll get back after the claim is resolved (assuming you acted honestly).

In short, Arbitrum enables you to interact with and deploy smart contracts at a fraction of the cost of using Ethereum natively, and using all the same tooling you use to interact with Ethereum today without compromising on security or decentralization. No custom tooling is required to use the chain, but anyone can elect to validate the chain.

## Executing and Securing the Chain

Arbitrum Rollup chains are secured by the Multi-Round Interactive Optimistic Rollup protocol first [published](https://www.usenix.org/conference/usenixsecurity18/presentation/kalodner) in Usenix Security conference in 2018. Any user can submit an assertion about the execution of the Rollup chain. After that assertion is submitted to Ethereum, a challenge period begins during which any other user can challenge the correctness of that assertion. After a challenge has been initiated, the dispute is mediated by Ethereum and it's guaranteed that an honest user will always win a challenge. To incentivize honest operation, validators place bonds that they will forfeit in the event that they lose a dispute. For more details on the protocol see the [protocol design](Rollup_Protocol.md) page.

## Submitting Transactions

Recall that the key features of Rollups is that they log all transaction data on chain. All transactions executed on the Arbitrum Rollup chain are submitted to an Inbox smart contract running on Ethereum. The execution of the rollup chain is based entirely on the transactions submitted by the inbox, and so anyone monitoring the inbox can know the correct state of the Arbitrum chain by simply executing the transactions from that contract. Despite the fact that executing a transaction on Arbitrum requires an Ethereum transaction, Arbitrum transactions use only a small fraction of the gas of an equivalent Ethereum transaction since only the raw transaction data goes to Ethereum but execution and contract storage happen off-chain. Arbitrum Rollup further provides a suite of compression tools that allow for further minimizing the amount of data that needs to be logged on the Ethereum blockchain.

## Aggregating Transactions

While users can submit their transactions to the inbox contract directly, they may elect to go through an aggregator, a node which collects transactions and submits them in batches to Ethereum.

There are two primary benefits to using aggregators. First, Ethereum transactions have a minimum cost that is significantly larger than a typical Arbitrum transaction. Submitting Arbitrum transactions individually will incur a sizable overhead. By submitting batches of Arbitrum transactions as a single Ethereum transaction, this minimum cost gets amortized over all transactions in the batch and substantially lowers gas costs. Second, transacting through an aggregator makes it possible for users to use the Rollup chain without the need to directly transact with or hold assets on Layer 1 Ethereum.

To learn more about Arbitrum aggregators, click [here](https://developer.offchainlabs.com/docs/aggregator).

## Throughput

Since all Arbitrum transaction data is posted to Ethereum, the cost per transactions as well as the number of transactions that Arbitrum can support per second is limited by the amount of data that can be posted to Ethereum during that time. For this reason, optimizing the compression of transactions to minimize what needs to be posted on-chain is critical for reducing costs and increasing throughput.

All transactions have associated metadata that needs to be stored on-chain, but some transactions will themselves carry custom calldata payloads. For this reason, the throughput of a chain is highly linked to the specific transactions being benchmarked. For simple transfer transactions that do not carry their own calldata, our benchmarks have shown that Arbitrum will allow for up to 4,500 transfer transactions per second. The Arbitrum block explorer gives detailed cost benchmarks for all transactions posted on the Arbitrum Rollup testnet chain, and is a very useful tool for benchmarking specific workloads.

## Cross contract communication

Contracts that are deployed on the Arbitrum Rollup chain can make synchronous calls to one another exactly as they would on Ethereum. A natural follow-up is how contracts on the Arbitrum Rollup chain interact with contracts on Ethereum. Cross chain and cross shard interoperability is a hard problem, but since rollups are rooted in Ethereum, interoperability can be relatively smooth, albeit asynchronous.

It's important to differentiate between calls from Ethereum to Arbitrum and calls from Arbitrum to Ethereum. Ethereum contracts can send transactions into Arbitrum which arrive quickly. However general transactions from Arbitrum to Ethereum are slower since they need to wait for the challenge period to expire before being confirmed. Luckily there are elegant solutions that allow users to quickly withdraw fungible assets from Arbitrum to Ethereum. See the docs entry on [withdrawals](Withdrawals.md) for more information on this.

Arbitrum Rollup contains a Token Bridge SDK and a Token Bridge user interface that serves as an easy way for trustlessly transferring assets between Arbitrum and Ethereum. Aside from support for transferring Ethereum-native assets to Arbitrum, Arbitrum also includes functionality that enables creating Arbitrum-native tokens that are minted on the Arbitrum Rollup chain, which can be subsequently transferred to the Ethereum blockchain.

## The Arbitrum Virtual Machine

Although Arbitrum supports EVM, under the hood it runs the Arbitrum Virtual Machine (AVM). The AVM is never exposed to developers or users, so if you're just interested in how to use Arbitrum, you can safely ignore it. But if you're curious about the inner workings of Arbitrum and how it achieves its scalability, read on.
The AVM is optimized for allowing fast progress in the optimistic case while maintaining the ability to efficiently resolve disputes. To learn more, you can read a detailed overview of the [AVM Design Rationale](AVM_Design.md) as well as the [AVM Specification](AVM_Specification.md), a lower level description of the semantics of the AVM architecture.

## ArbOS

ArbOS, the Arbitrum operating system, sits on top of the AVM and is responsible for isolating untrusted contracts from one another, tracking and limiting their resource usage using [ArbGas](ArbGas.md), and managing the economic model that collects fees from users to fund the operation of a chain's validators. ArbOS gives Arbitrum a great deal of flexibility by offloading work that would have been done in the L1 smart contract into cheaper L2 code. To learn more, see the section on [ArbOS](ArbOS.md).
