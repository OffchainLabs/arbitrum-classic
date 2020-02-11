---
id: Rinkeby
title: Running on Rinkeby
sidebar_label: Rinkeby
---

## Launching the chain

To set up a rollup chain the the Arbitrum Rollup Deployment [tool](https://developer.offchainlabs.com/tools/deployment/).

The deployment tool will assist you in configuring and launching your Rollup Chain. After you've successfully launched your chain, it will give you the chain's Ethereum address.

## Initializing the validators

```bash
./scripts/initialize_validators.py [contract_path] [validator_count]
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
