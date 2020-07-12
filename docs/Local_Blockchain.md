---
id: Local_Blockchain
title: Running on Local Blockchain
sidebar_label: Local Blockchain
---

## Launching a Local Blockchain

To build a docker image hosting a local test blockchain docker image with Arbitrum smart contracts already deployed, run:

```bash
yarn docker:build:geth
```

To start the local blockchain inside the Arbitrum monorepo, run:

```bash
yarn docker:geth
```

Note that stopping and restarting the client will lose all blockchain state.

## Launching the chain

To set up a local rollup chain using the Arbitrum geth docker image with 1 or more validators, run the following from the root arbitrum repo.

```bash
yarn demo:initialize [--validatorcount N]
```

Running the `demo:initialize` command will perform two main tasks 1) Launch an Arbitrum Rollup chain on the local testnet 2) Create a `validator-states` folder. This folder contains pre-seeded wallets for the created validators which are prepared for launch. It serves as a lightweight simulation of an enviroment where the validators are running on multiple machines.

## Deploying your validators

To launch a set of docker images containing your validators, run:

```bash
yarn demo:deploy
```

Unlike the blockchain docker image, the validators can be stopped and restarted without losing any state.
