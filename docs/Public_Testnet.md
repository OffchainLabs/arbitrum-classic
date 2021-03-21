---
id: Public_Testnet
title: Public Testnet Guide
sidebar_label: Public Testnet
---

In order to make it easy for people to get started using Arbitrum Rollup, we've launched our own hosted Arbitrum Rollup chain hosted on the Kovan testnet.

For a general introduction to the Arbitrum Testnet, see our [announcement](https://medium.com/offchainlabs/arbitrum-rollup-testnet-full-featured-and-open-to-all-da3255b562ea).

For a convenient landing page for all your testnet needs, see our [website](https://arbitrum.io/testnet/).

## Connection Information

Hosted Aggregator Node (JSON-RPC Endpoint): [https://kovan4.arbitrum.io/rpc](https://kovan4.arbitrum.io/rpc)

Rollup Chain ID: 212984383488152

Testnet Version: v4

Chain Launched on March 19, 2021

## Deploying your contracts

Deploying your contracts onto the Arbitrum testnet is as easy as changing your RPC endpoint to https://kovan4.arbitrum.io/rpc

For a deeper dive into deploying with truffle see [here](Contract_Deployment.md).

## Porting your frontend

Porting your frontend is just as easy as deploying your contracts. Just take your existing frontend and point it at our RPC endpoint after deploying your contract. For more information and code samples see [here](Frontend_Integration.md)

If you're using metamask, add a custom RPC network to connect to the Arbitrum testnet:

- Network Name: Arbitrum Testnet V3
- New RPC URL: https://kovan4.arbitrum.io/rpc
- ChainID (Optional): 212984383488152
- Symbol: ETH
- Block Explorer URL: https://explorer.arbitrum.io/#/

## Observing transactions

If you'd like to see your transactions in action, check out our [block explorer](https://explorer.arbitrum.io/#/)!

There you'll be able to see all the transactions being executed in Arbitrum and also see exactly how much Ethereum Gas each transaction uses.

<!--
## Running your own node

We're running an aggregator and validator for our testnet, so you don't have to run any of your own infrastructure. However Arbitrum Rollup is totally decentralized, so if you'd like to run your own infrastructure you can avoid our servers entirely.

The very first step to start building with Arbitrum is [installing](Installation.md). After that you can initialize your local setup by running:

```bash
yarn prod:initialize  0x2e8aF9f74046D3E55202Fcfb893348316B142230 https://kovan.infura.io/v3/YOUR_INFURA_API_ID
```

Running the `prod:initialize` command will create a `arbitrum/rollups/ 0x3B493fD1731528531471Cd18ea2f29f1463D6514` folder with two subfolders, one configured for an aggregator, and the other a validator.

To deploy the validator and aggregator, run

```bash
yarn deploy:validators  0x2e8aF9f74046D3E55202Fcfb893348316B142230 --password=[password]
```

Upon deploying a validator, you'll be asked to deposit the staking requirement, 1 Kovan ETH.

The password argument is used to secure the validator keystore. On the first deployment you set the password to any value, and on later deployments you must resubmit the same password.
-->
## Kovan Deployment

All contracts are deployed from https://github.com/OffchainLabs/arbitrum/tree/v0.7.2/packages/arb-bridge-eth/contracts

#### Important Addresses

- Main L1 Rollup Contract: [0x2e8aF9f74046D3E55202Fcfb893348316B142230](https://kovan.etherscan.io/address/0x2e8aF9f74046D3E55202Fcfb893348316B142230)
- Our Hosted Aggregator Address: [0xa300a724d86564615763c58f579248e0d7d08d36](https://kovan.etherscan.io/address/0xa300a724d86564615763c58f579248e0d7d08d36)
- Our Hosted Validator Address: [0xa300bb7096beb6ca8142bdc7dfa7123726131ef8](https://kovan.etherscan.io/address/0xa300bb7096beb6ca8142bdc7dfa7123726131ef8)

#### Rollup Contracts

- [ArbFactory](https://github.com/OffchainLabs/arbitrum/blob/v0.7.2/packages/arb-bridge-eth/contracts/rollup/ArbFactory.sol) - [0x1818CA53Ed3dCf11c5992865388BC37E1e9DD312](https://kovan.etherscan.io/address/0x1818CA53Ed3dCf11c5992865388BC37E1e9DD312)
- [ArbRollup](https://github.com/OffchainLabs/arbitrum/blob/v0.7.2/packages/arb-bridge-eth/contracts/rollup/ArbRollup.sol) - [0x731E9FA93b3b9Fa26f1308544c7651a9D5813Fbc](https://kovan.etherscan.io/address/0x731E9FA93b3b9Fa26f1308544c7651a9D5813Fbc) (Template contract)

#### Inbox Contract

- [GlobalInbox](https://github.com/OffchainLabs/arbitrum/blob/v0.7.2/packages/arb-bridge-eth/contracts/inbox/GlobalInbox.sol) - [0xE681857DEfE8b454244e701BA63EfAa078d7eA85](https://kovan.etherscan.io/address/0xE681857DEfE8b454244e701BA63EfAa078d7eA85)

#### Fraud Proofs Contracts

- [ChallengeFactory](https://github.com/OffchainLabs/arbitrum/blob/v0.7.2/packages/arb-bridge-eth/contracts/challenge/ChallengeFactory.sol) - [0x7064443c714996fbF3e783d5347C528E2E1Ad877](https://kovan.etherscan.io/address/0x7064443c714996fbF3e783d5347C528E2E1Ad877)
- [InboxTopChallenge](https://github.com/OffchainLabs/arbitrum/blob/v0.7.2/packages/arb-bridge-eth/contracts/challenge/InboxTopChallenge.sol) - [0x2C3927aCC3AB8051BAd86B80b9d2D9A1dE5999BD](https://kovan.etherscan.io/address/0x2C3927aCC3AB8051BAd86B80b9d2D9A1dE5999BD) (Template contract)
- [ExecutionChallenge](https://github.com/OffchainLabs/arbitrum/blob/v0.7.2/packages/arb-bridge-eth/contracts/challenge/ExecutionChallenge.sol) - [0x356e19929FCb4973c131d558300E3E353cb8e1C9](https://kovan.etherscan.io/address/0x356e19929FCb4973c131d558300E3E353cb8e1C9) (Template contract)
- [OneStepProof](https://github.com/OffchainLabs/arbitrum/blob/v0.7.2/packages/arb-bridge-eth/contracts/arch/OneStepProof.sol) - [0x082D26eeAa348A7C02291cd1948c66a79fc80aAD](https://kovan.etherscan.io/address/0x082D26eeAa348A7C02291cd1948c66a79fc80aAD)

## V2 Arbitrum Chain

For connecting to our older, "v2" Arbitrum chain (also running on Kovan), use the following:

- Aggregator RPC Endpoint: https://kovan2.arbitrum.io/rpc
- Chain ID: 152709604825713
- Rollup Contract Address: [0xC34Fd04E698dB75f8381BFA7298e8Ae379bFDA71](https://kovan.etherscan.io/address/0xC34Fd04E698dB75f8381BFA7298e8Ae379bFDA71)
