---
id: Public_Testnet
title: Public Testnet Guide
sidebar_label: Public Testnet
---

In order to make it easy for people to get started using Arbitrum Rollup, we've launched our own hosted Arbitrum Rollup chain hosted on the Kovan testnet.

For a general introduction to the Arbitrum Testnet, see our [announcement](https://medium.com/offchainlabs/arbitrum-rollup-testnet-full-featured-and-open-to-all-da3255b562ea).

For a convenient landing page for all your testnet needs, see our [website](https://arbitrum.io/testnet/).

## Connection Information

Hosted Aggregator Node (JSON-RPC Endpoint): [https://kovan5.arbitrum.io/rpc](https://kovan5.arbitrum.io/rpc)

Rollup Chain ID: 144545313136048

Testnet Version: v5

Chain Launched on May 15, 2021

## Connecting to the chain

If you're using metamask, add a custom RPC network to connect to the Arbitrum testnet:

- Network Name: Arbitrum Testnet V5
- New RPC URL: https://kovan5.arbitrum.io/rpc
- ChainID: 144545313136048
- Symbol: ETH
- Block Explorer URL: https://explorer5.arbitrum.io/#/

## Observing transactions

If you'd like to see your transactions in action, check out our [block explorer](https://explorer5.arbitrum.io/#/)!

There you'll be able to see all the transactions being executed in Arbitrum and also see exactly how much Ethereum Gas each transaction uses.

## Bridging Eth and ERC-20 Tokens

In order to deposit and withdraw Eth or tokens, visit https://bridge5.arbitrum.io.

In order to start using the chain, you'll have deposit Eth from kovan so that you can pay for fees in L2. In order to get kovan Eth, use one of the standard faucets from https://github.com/kovan-testnet/faucet.

## Interacting with the chain

Once you've added the Arbitrum Testnet V5 network to metamask, you should be able to interact with the Arbitrum chain just like you would with Ethereum.

The are a couple things to note on the Arbitrum chain.

- Arbitrum uses an EIP-1559-like gas auction system so the gas price you list in your transaction is a bid, but the actual price may be lower
- In order to do a ETH transfer through metamask, you must manually enter a higher gas limit than the default 21,000 gas. 800,000 should work well
- The majority of gas costs paid in the arbitrum chain go to pay for the cost of posting your transaction data to Ethereum

## Deploying your contracts

Deploying your contracts onto the Arbitrum testnet is as easy as changing your RPC endpoint to https://kovan5.arbitrum.io/rpc

For a deeper dive into deploying with truffle see [here](Contract_Deployment.md).

## Porting your frontend

Porting your frontend is just as easy as deploying your contracts. Just take your existing frontend and point it at our RPC endpoint after deploying your contract. For more information and code samples see [here](Frontend_Integration.md)

## V4 Arbitrum Chain

For connecting to our older, "v4" Arbitrum chain (also running on Kovan), use the following:

- Aggregator RPC Endpoint: https://kovan4.arbitrum.io/rpc
- Chain ID: 212984383488152
- Explorer: https://explorer.arbitrum.io/#/

## Kovan Deployment

All contracts are deployed from https://github.com/OffchainLabs/arbitrum/tree/5521d2e8720b8e6724d4dc7a6d7711d72e634e36/packages/arb-bridge-eth/contracts

#### Important Addresses

- Main L1 Rollup Contract: [0x2B0474e5201646fd7d8eEf8522a88376940B1db0](https://kovan.etherscan.io/address/0x2B0474e5201646fd7d8eEf8522a88376940B1db0)
- Ethereum Inbox Contract [0x76bF1345224fE606E2aB38B8E52B83512328A9DF](https://kovan.etherscan.io/address/0x76bF1345224fE606E2aB38B8E52B83512328A9DF)

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

-->
