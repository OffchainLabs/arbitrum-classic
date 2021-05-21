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

If you want to get started using Arbitrum with no setup required, check out our [public testnet](Public_Testnet.md) running on top of Kovan.

### How Can I Develop Locally

The very first step to start building with Arbitrum is [installing](Installation.md) Arbitrum and its dependencies. Next, you'll need to deploy an Arbitrum chain on an L1 blockchain. Arbitrum Rollup supports deployment both on a [local testnet](Local_Blockchain.md) and on the [Rinkeby Testnet](Rinkeby.md). The following quickstart walks through deployment of an Arbitrum Rollup chain on the local testnet.

Note that Arbitrum chains support dynamic launching of contracts, so you don't need to setup an Arbitrum chain for each application you build, and indeed you may deploy your contracts on a testnet chain which you did not launch. The benefits of having multiple applications on the same Arbitrum Rollup chain is that they'll be able to interact synchronously, just as they would if they were launched directly on Ethereum.

Once you have deployed Arbitrum, you can [build and run the demo app](#hello-arbitrum) or [deploy your own contracts](Contract_Deployment.md).

**Want to learn more? Check out the** [**open source code**](https://github.com/offchainlabs/arbitrum)**. Join the team on** [**Discord**](https://discord.gg/ZpZuw7p)**.**

## Setup Local Geth and Rollup Blockchain

See [Local Blockchain Setup](Local_Blockchain.md).

## Hello, Arbitrum

Now you'll deploy and run a demo dApp on Arbitrum. The dApp is based on
a simple Pet Shop dApp that is used in a Truffle tutorial.

First clone the pet-shop demo dApp and install dependencies:

```bash
git clone https://github.com/OffchainLabs/demo-dapp-pet-shop
cd demo-dapp-pet-shop
yarn
```

### Deployment

Deploy contracts to Arbitrum :

```bash
truffle migrate --network arbitrum
```

### Use the dApp

1. Install [Metamask](https://metamask.io/)

   > Once Metamask is installed, open it and select
   > `Import Account` and enter one of the following pre-funded private keys
   >
   > ```
   > 0x979f020f6f6f71577c09db93ba944c89945f10fade64cfc7eb26137d5816fb76
   > 0xd26a199ae5b6bed1992439d1840f7cb400d0a55a0c9f796fa67d7c571fbb180e
   > 0xaf5c2984cb1e2f668ae3fd5bbfe0471f68417efd012493538dcd42692299155b
   > 0x9af1e691e3db692cc9cad4e87b6490e099eb291e3b434a0d3f014dfd2bb747cc
   > 0x27e926925fb5903ee038c894d9880f74d3dd6518e23ab5e5651de93327c7dffa
   > 0xe4b33c0bb790b88f2463facaf86ae7c17cbdab41187e69ddde8cc1c1fda7c9ab
   > ```

2) Select local arbitrum network in Metamask

   - Go back to Metamask or click the extension icon
   - Select `Main Ethereum Network` top right hand side
   - Choose `Custom RPC`
   - Enter `Local Arbitrum` as the network name
   - Enter `http://127.0.0.1:8547` as the RPC url
   - Press the save button
   - Metamask should now have an Local Arbitrum account holding ETH

3) Launch the front-end

   ```bash
   yarn start
   ```

   The browser will open to [localhost:8080](http://localhost:8080)

   In the popup window that appears, select `Connect`

4) Adopt some pets

   The pet shop dApp should now be running in your browser. Choose a pet or two
   and click the adopt button to adopt your new animal friend(s).

### Summary

If you want to try another dApp run, deploy the solidity contracts and launch the frontend

```bash
git clone https://github.com/OffchainLabs/demo-dapp-election
cd demo-dapp-election
yarn
truffle migrate --network arbitrum
yarn start
```
