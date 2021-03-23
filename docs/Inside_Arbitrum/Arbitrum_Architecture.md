---
id: Arbitrum_Architecture
title: Inside Arbitrum: Arbitrum Architecture
sidebar_label: Arbitrum Architecture
---
This diagram shows the basic architecture of Arbitrum.

![img](https://lh5.googleusercontent.com/1qwGMCrLQjJMv9zhWIUYkQXoDR2IksU5IzcSUPNJ5pWkY81pCvr7WkTf4-sb41cVohcnL-i6y8M1LU8v-4RXT_fdOsaMuLXnjwerSuKTQdHE-Hrvf4qBhRQ2r7qjxuAi3mk3hgkh)

On the left we have users and the service providers who help them connect to the chain(s) of their choice. On the right we have the Arbitrum system itself, built in layers on top of Ethereum. 

We’ll work our way up on the right side to describe how the Arbitrum stack works, then we’ll talk about what happens on the left side to connect users to it.

On the bottom right is good old **Ethereum**. Arbitrum builds on Ethereum and inherits its security from Ethereum.

On top of Ethereum is the **EthBridge**, a set of Ethereum contracts that manage an Arbitrum chain. The EthBridge referees the Arbitrum rollup protocol, which ensures that the layers above it operate correctly. (More on the rollup protocol below in the XXX section.) The EthBridge also maintains the chain’s inbox and outbox, allowing people and contracts to send transaction messages to the chain, and to observe and use the outputs of those transactions. Users, L1 Ethereum contracts, and Arbitrum nodes make calls to the EthBridge contracts to interact with the Arbitrum chain. (More on that below in the XXX section.)

The horizontal layer boundary above the EthBridge is labeled **AVM Architecture**, because what the EthBridge provides to the layer above it is an Arbitrum Virtual Machine, which can execute a computer program that reads inputs and produces outputs. This is the most important interface in Arbitrum, because it divides Layer 1 from Layer 2--it divides the Layer 1 components that provide the inbox/execution/outbox abstraction from the Layer 2 components that use that abstraction.

![img](https://lh4.googleusercontent.com/qwf_aYyB1AfX9s-_PQysOmPNtWB164_qA6isj3NhkDnmcro6J75f6MC2_AjlN60lpSkSw6DtZwNfrt13F3E_G8jdvjeWHX8EophDA2oUM0mEpPVeTlMbsjUCMmztEM0WvDpyWZ6R)

The next layer up is **ArbOS**. This is a software program, written by Offchain Labs, that runs on the Arbitrum Virtual Machine, and serves as a record-keeper, traffic cop, and enforcer for the execution of smart contracts on the Arbitrum chain. It’s called ArbOS because it plays a role like a (lightweight version of) the operating system on a laptop or phone--it’s the program that starts up first and that manages the execution of all other code on the chain. Importantly, ArbOS runs entirely at Layer 2, off of the Ethereum chain, so it can take advantage of the scalability and low cost of Layer 2 computation.

The horizontal layer boundary above ArbOS is called **EVM compatibility** because ArbOS provides an Ethereum Virtual Machine compatible execution environment for smart contracts. That is, you can send ArbOS the EVM code for a contract, in the same way you would send that contract to Ethereum, and ArbOS will load the contract and enable it to service transactions, just like on Ethereum. ArbOS takes care of the details of compatibility, so the smart contract programmer can just write their code like they would on Ethereum (or often, just take existing Ethereum contracts and redeploy them).

At the top of the stack--the upper right portion of the diagram--are **EVM contracts** which have been deployed to the Arbitrum chain by developers, and which execute transactions that are submitted to the chain.

That’s the right hand side of the diagram, which provides the Arbitrum chain functionality. Now let’s turn to the left side, which more directly supports users.

On the lower left are standard **Ethereum nodes**, which are used to interact with the Ethereum chain.
Just above that are **Arbitrum nodes**. As the name suggests, these are used to interact with Arbitrum. They support the same API as Ethereum nodes, so they work well with existing Ethereum tools -- you can point your Ethereum-compatible wallet or tools at an Arbitrum node and they’ll be able to talk to each other. Just like on Ethereum, anyone can run an Arbitrum node, but many people will choose instead to rely on a node run by someone else.

Some Arbitrum nodes service user requests, and others choose to serve only as validators, which work to ensure the correctness of the Arbitrum chain. (See section XXX below for details on validators.)

Last, but certainly not least, we see **users** on the upper left. Users use wallets, dapp front ends, and other tools to interact with Arbitrum. Because Arbitrum nodes support the same API as Ethereum, users don’t need entirely new tooling and developers don’t need to rewrite their dapps.