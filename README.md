<p align="center"><img src="https://offchainlabs.com/c79291eee1a8e736eebd9a2c708dbe44.png" width="600"></p>

# Arbitrum Monorepo

[![CircleCI](https://circleci.com/gh/OffchainLabs/arbitrum.svg?style=svg)](https://circleci.com/gh/OffchainLabs/arbitrum) [![codecov](https://codecov.io/gh/OffchainLabs/arbitrum/branch/master/graph/badge.svg)](https://codecov.io/gh/OffchainLabs/arbitrum)

Arbitrum is a Layer 2 cryptocurrency platform that makes smart contracts scalable, fast, and private. Arbitrum interoperates closely with Ethereum, so Ethereum developers can easily cross-compile their contracts to run on Arbitrum. Arbitrum achieves these goals through a unique combination of incentives, network protocol design, and virtual machine architecture. Arbitrum has three modes: channels, AnyTrust sidechains, and rollup. Channels and sidechains provide the AnyTrust Guarantee which ensures that the code will run correctly as long as any validator is honest.

Want to learn more? Join the team on [Discord](https://discord.gg/ZpZuw7p), follow the [developer guide](https://developer.offchainlabs.com), and read the [white paper](https://offchainlabs.com/arbitrum.pdf)!

Arbitrum technologies are patent pending. This repository is offered under the Apache 2.0 license. See LICENSE for details.

## Current Status

#### Arbitrum is currently Alpha software and should not be used in production environments.

## Quickstart

Follow [the guide](https://developer.offchainlabs.com/docs/Developer_Quickstart/) on our developer site to build a demo dapp on Arbitrum.

## Local setup

MacOS dependencies

```bash
brew tap ethereum/ethereum
brew install ethereum boost libtool autoconf automake gmp cmake
```

Start local Geth instance and deploy contracts

```bash
yarn run:local:geth
```

Install the validator

```bash
yarn install:validator
yarn demo:initialize
```

Run the validator, when prompted input the password "`pass`"

```bash
arb-validator validate -rpc -blocktime=2 rollups/local/validator0  http://localhost:7545 [rollup address]
```
