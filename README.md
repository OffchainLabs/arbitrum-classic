<p align="center"><img src="https://offchainlabs.com/c79291eee1a8e736eebd9a2c708dbe44.png" width="600"></p>

# Arbitrum Monorepo

[![CircleCI](https://circleci.com/gh/OffchainLabs/arbitrum.svg?style=svg)](https://circleci.com/gh/OffchainLabs/arbitrum) [![codecov](https://codecov.io/gh/OffchainLabs/arbitrum/branch/master/graph/badge.svg)](https://codecov.io/gh/OffchainLabs/arbitrum)

Arbitrum is a Layer 2 cryptocurrency platform that makes smart contracts scalable, fast, and private. Arbitrum interoperates closely with Ethereum, so Ethereum developers can easily cross-compile their contracts to run on Arbitrum. Arbitrum achieves these goals through a unique combination of incentives, network protocol design, and virtual machine architecture.

Want to learn more? Join the team on [Discord](https://discord.gg/ZpZuw7p), follow the [developer guide](https://developer.offchainlabs.com), and read the [white paper](https://offchainlabs.com/arbitrum.pdf)!

Arbitrum technologies are patent pending. This repository is offered under the Apache 2.0 license. See LICENSE for details.

## Current Status

#### Arbitrum is currently Alpha software and should not be used in production environments.

## Quickstart

Clone the monorepo to get started:

```bash
git clone -b master --depth=1 -c advice.detachedHead=false https://github.com/OffchainLabs/arbitrum.git
cd arbitrum
yarn
yarn install:deps
cd demos/pet-shop
truffle migrate --reset --compile-all --network arbitrum
cd ../..
./scripts/arb_deploy.py demos/pet-shop/contract.ao 3

# Start the frontend in another session:
cd demos/pet-shop && yarn start
```

[Next](https://developer.offchainlabs.com/docs/Developer_Quickstart/#use-the-dapp), setup Metamask with a new
account and the mnemonic:

```
jar deny prosper gasp flush glass core corn alarm treat leg smart
```

or use your existing account with the private key derived from the previous mnemonic:

```
0x41a9550a0ae23fd52f3b99acab194db2e4474262db64dfd46807bca9e061e211
```

You will also need to switch networks to `http://localhost:7545` to connect to the Ganache started by the `arb_deploy.py`
script before adopting your pets.
