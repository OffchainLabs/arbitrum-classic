---
id: Rinkeby
title: Running on Rinkeby
sidebar_label: Rinkeby
---

## Launching the chain

To set up a rollup chain the the Arbitrum Rollup Deployment [tool](https://developer.offchainlabs.com/tools/deployment/).

The deployment tool will assist you in configuring and launching your Rollup Chain. After you've successfully launched your chain, it will give you the chain's Ethereum address.

## Initializing the validators

Run the following command to initialize a set of validators for a given Rollup contract. The `contract_path` should be to the same file you submitted to the deployment tool, the `rollup_address` should be the address output by the tool after a successful deployment, and the `eth_url` should be a RPC or WebSocket connection to an Ethereum node.

You can select to have one or more validators run on your machine for testing purposes.

```bash
./scripts/initialize_validators.py [contract_path] [validator_count] [rollup_address] [eth_url]
```

Note that this command executes the `initialize_validators.py` script and if you are not in the root of the `arbitrum`
directory, you may need to use a different path.

Running the `initialize_validators` command will create a `validator-states` folder with `validator_count` subfolders, each configured with a validator for the given Rollup contract.. It serves as a lightweight simulation of an enviroment where the validators are running on multiple machines.

## Deploying your validators

To launch a set of docker images containing your validators, run:

```bash
./scripts/arb_deploy.py validator-states --password=[password]
```

`arb_deploy.py` takes a path to the validator-states created in the previous step. Unlike the blockchain docker image, the validators can be stopped and restarted without losing any state. The password argument is used to secure the validator keystore. On the first deployment you set the password to any value, and on later deployments you must resubmit the same password.
