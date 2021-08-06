---
id: Developer_Quickstart
title: Arbitrum Developer Quickstart
sidebar_label: Developer Quickstart
custom_edit_url:https://github.com/OffchainLabs/arbitrum/edit/master/docs/Developer_Quickstart.md
---

Arbitrum is a suite of Ethereum scaling solutions that enables high-throughput, low cost smart contracts while remaining trustlessly secure. Arbitrum has three modes: AnyTrust Channels, AnyTrust Sidechains, and Arbitrum Rollup. The following documentation describes how to use Arbitrum Rollup, which is currently live on testnet. Whether you're a developer that just wants to start building or you're curious into digging deeper into the internals of Arbitrum and how it works, this site is the right place for you.

### How does Arbitrum work?

If you're looking to discover how Arbitrum works, the best place to begin is by the [Rollups basics](Rollup_basics.md) section, which gives a high level overview of Arbitrum's internals. From there, you can jump into more detailed explainers on various components of the system.

### How Can I Start Building

If you want to get started using Arbitrum with no setup required, check out our [public testnet](Public_Testnet.md) running on top of Rinkeby.

### How Can I Develop Locally

The very first step to start building with Arbitrum is [installing](Installation.md) Arbitrum and its dependencies. Next, you'll need to deploy an Arbitrum chain on an L1 blockchain. You can follow the [local testnet guide](Local_Blockchain.md) for a quickstart walkthrough deployment of an Arbitrum Rollup chain on the local testnet.

Note that Arbitrum chains support dynamic launching of contracts, so you don't need to setup an Arbitrum chain for each application you build, and indeed you may deploy your contracts on a testnet chain which you did not launch. The benefits of having multiple applications on the same Arbitrum Rollup chain is that they'll be able to interact synchronously, just as they would if they were launched directly on Ethereum.

Once you have deployed Arbitrum, you can [build and run the demo app](#hello-arbitrum) or [deploy your own contracts](Contract_Deployment.md).

**Want to learn more? Check out the** [**open source code**](https://github.com/offchainlabs/arbitrum)**. Join the team on** [**Discord**](https://discord.gg/ZpZuw7p)**.**

## Setup Local Geth and Rollup Blockchain

See [Local Blockchain Setup](Local_Blockchain.md).

## Hello, Arbitrum 

Check out our Arbitrum Interop Quickstart repo here! https://github.com/OffchainLabs/arbitrum-tutorials

This monorepo will help you get started with building on Arbitrum. It provides various simple demos showing and explaining how to interact with Arbitrum — deploying and using contracts directly on L2, moving Ether and tokens betweens L1 and L2, and more.

We show how you can use broadly supported Ethereum ecosystem tooling (Hardhat, Ethers-js, etc.) as well as our special [arb-ts](https://github.com/OffchainLabs/arbitrum/tree/master/packages/arb-ts) library for convenience.

### What's included?

#### :white_check_mark: Basics

- 🐹 [Pet Shop DApp](./packages/demo-dapp-pet-shop/) (L2 only)
- 🗳 [Election DApp](./packages/demo-dapp-election/) (L2 only)

#### :white_check_mark: Moving Stuff around

- ⤴️ 🔹 [Deposit Ether](./packages/eth_deposit/)
- ⤵️ 🔹 [Withdraw Ether](./packages/eth_withdraw/)
- ⤴️ 💸 [Deposit Token](./packages/token_deposit/)
- ⤵️ 💸 [Withdraw token](./packages/token_withdraw/)

#### :white_check_mark: General Interop

- 🤝 [Greeter](./packages/greeter/) (L1 to L2)
- 📤 [Outbox](./packages/outbox-execute/) (L2 to L1)

#### :white_check_mark: Advanced Features

- ®️ [Arb Address Table](./packages/address_table/)

<p align="center">
  <img width="350" height="100" src= "https://arbitrum.io/wp-content/uploads/2021/01/cropped-Arbitrum_Horizontal-Logo-Full-color-White-background-scaled-1.jpg" />
</p>
