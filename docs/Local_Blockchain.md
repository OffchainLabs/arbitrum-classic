---
id: Local_Blockchain
title: Running on Local Blockchain
sidebar_label: Local Blockchain
---

## Basics

To run Arbitrum locally, you need several things:

* A local Ethereum blockchain (the L1)
* A set of Arbitrum smart contracts deployed on that L1
* The Arbitrum blockchain itself (the L2)
* And, something that interacts with the L1 and the L2, in this case, one or more Arbitrum validators.   

## Launching a Local Ethereum Blockchain (the L1)

To *build* a docker image hosting a local Geth blockchain with Arbitrum smart contracts (the `eth-bridge`) already deployed, run:

```bash
yarn docker:build:geth
```

To *start* this local Geth, run:

```bash
yarn docker:geth
```

Note that stopping and restarting the client will lose all blockchain state. At this point, you should see typical Geth INFO updates scrolling in your terminal. Geth is now happy running at localhost:7545.  

## Configuring your local Arbitrum chain (the L2)

To set up a local rollup chain using the Arbitrum geth docker image you created above with 1 or more validators, open a second terminal window and run the following from the root arbitrum repo:

```bash
yarn demo:initialize [--validatorcount N=1]
```

This command will fail if Geth is not up yet and/or if there were problems with the Hardhat deployment of all the bridge smart contracts. Specifically, you will not have `bridge_eth_addresses.json`, and thus, your L2 will not be able to talk to the bridge on the L1. Running `demo:initialize` will *configure* an Arbitrum L2 rollup on your local machine - see the contents of `/rollups/local/` for more information on what `demo:initialize` sets up. Among other things, this folder contains pre-seeded wallets for the created validators which are prepared for launch. It serves as a lightweight simulation of an enviroment where the validators are running on multiple machines.

## Firing up the Arbitrum L2 and Deploying your validator(s)

To *launch* the L2 node and run the validators, run:

```bash
yarn demo:deploy
```

Unlike the blockchain docker image, the validators can be stopped and restarted without losing any state. At this point, you will have a local Geth running in your first terminal window, and `arbitrum_arb-validator1_1` and `arbitrum_arb-node_1` in terminal two. You can now watch the L1, the L2, and the validator(s) interact and execute a variety of predefined transactions. 
