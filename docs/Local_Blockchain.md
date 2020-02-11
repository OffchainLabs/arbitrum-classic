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

To set up a local rollup chain using the Arbitrum geth docker image described in the installation instructions, run:

```bash
./scripts/setup_rollup.py [contract_path] [validator_count]
```

This command will initialize your rollup chain and create a `validator-states` folder with configuration
information for `validator_count` validators prepared to validator it.

Note that this command executes the `setup_rollup.py` script and if you are not in the root of the `arbitrum`
directory, you may need to use a different path.

Running the `setup_rollup` command will perform two main tasks 1) Launch an Arbitrum Rollup chain on the local testnet 2) Create a `validator-states` folder. This folder contains pre-seeded wallets for the created validators. It serves as a lightweight simulation of an enviroment where the validators are running on multiple machines.

## Deploying your validators

To launch a set of docker images containing your validators, run:

```bash
./scripts/arb_deploy.py validator-states
```

`arb_deploy.py` takes a path to the validator-states created in the previous step. Unlike the blockchain docker image, the validators can be stopped and restarted without losing any state.
