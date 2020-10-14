---
id: Public_Testnet
title: Public Testnet Guide
sidebar_label: Public Testnet
---

In order to make it easy for people to get started using Arbitrum Rollup, we've launched our own hosted Arbitrum Rollup chain hosted on the Kovan testnet.

For a general introduction to the Arbitrum Testnet, see our [announcement](https://TODO.com).

For a convinient landing page for all your testnet needs, see our [website](https://arbitrum.io/testnet/).

## Connection Information

Hosted Aggregator Node (JSON-RPC Endpoint): [https://node.offchainlabs.com:8547](https://node.offchainlabs.com:8547)

Rollup Chain ID: 215728282823301

## Deploying your contracts

Deploying your contracts onto the Arbitrum testnet is as easy as changing your RPC endpoint to https://node.offchainlabs.com:8547

For a deeper dive into deploying with truffle see [here](Contract_Deployment.md).

## Porting your frontend

Porting your frontend is just as easy as deploying your contracts. Just take your existing frontend and point it at our RPC endpoint after deploying your contract. For more information and code samples see [here](Frontend_Integration.md)

## Observing transactions

If you'd like to see your transactions in action, check out our [block explorer](https://explorer.offchainlabs.com/#/chain/kovan-alpha)!

There you'll be able to see all the transactions being executed in Arbitrum and also see exactly how much Ethereum Gas each transaction uses.

## Running your own node

We're running an aggregator and validator for our testnet, so you don't have to run any of your own infrastructure. However Arbitrum Rollup is totally decentralized, so if you'd like to run your own infrastructure you can avoid our servers entirely.

The very first step to start building with Arbitrum is [installing](Installation.md). After that you can initialize your local setup by running:
```bash
yarn prod:initialize 0x175c0b09453cbb44fb7f56ba5638c43427aa6a85 https://kovan.infura.io/v3/INSERT_API_ID
```

Running the `prod:initialize` command will create a `arbitrum/rollups/0x175c0b09453cbb44fb7f56ba5638c43427aa6a85` folder two subfolders, one configured for an aggregator, and the other a validator.

To deploy the validator and aggregator, run

```bash
yarn deploy:validators 0x175c0b09453cbb44fb7f56ba5638c43427aa6a85 --password=[password]
```

The password argument is used to secure the validator keystore. On the first deployment you set the password to any value, and on later deployments you must resubmit the same password.

## Kovan Deployment

All contracts are deployed from https://github.com/OffchainLabs/arbitrum/tree/v0.7.2/packages/arb-bridge-eth/contracts

#### Important Addresses

- Main L1 Rollup Contract: [0x175c0b09453cbb44fb7f56ba5638c43427aa6a85](https://kovan.etherscan.io/address/0x175c0b09453cbb44fb7f56ba5638c43427aa6a85)
- Hosted Aggregator Address: [0x1d143638962dc93c52c4053a3dcce71cccb30bd3](https://kovan.etherscan.io/address/0x1d143638962dc93c52c4053a3dcce71cccb30bd3)
- Hosted Validator Address: [0x705c33d9364dd570bc2998a1a1e788221c14d2da](https://kovan.etherscan.io/address/0x705c33d9364dd570bc2998a1a1e788221c14d2da)

#### Rollup Contracts
- [ArbFactory](https://github.com/OffchainLabs/arbitrum/blob/v0.7.2/packages/arb-bridge-eth/contracts/rollup/ArbFactory.sol) - [0xee1250962014364aCf506061E66e78e65b8bCEEC](https://kovan.etherscan.io/address/0xee1250962014364aCf506061E66e78e65b8bCEEC)
- [ArbRollup](https://github.com/OffchainLabs/arbitrum/blob/v0.7.2/packages/arb-bridge-eth/contracts/rollup/ArbRollup.sol) - [0x5c1351258f436dA83f37D6A46424225A08914bd5](https://kovan.etherscan.io/address/0x5c1351258f436dA83f37D6A46424225A08914bd5) (Template contract)

#### Inbox Contract
- [GlobalInbox](https://github.com/OffchainLabs/arbitrum/blob/v0.7.2/packages/arb-bridge-eth/contracts/inbox/GlobalInbox.sol) - [0xE681857DEfE8b454244e701BA63EfAa078d7eA85](https://kovan.etherscan.io/address/0xE681857DEfE8b454244e701BA63EfAa078d7eA85)

#### Fraud Proofs Contracts
- [ChallengeFactory](https://github.com/OffchainLabs/arbitrum/blob/v0.7.2/packages/arb-bridge-eth/contracts/challenge/ChallengeFactory.sol) - [0x7064443c714996fbF3e783d5347C528E2E1Ad877](https://kovan.etherscan.io/address/0x7064443c714996fbF3e783d5347C528E2E1Ad877)
- [InboxTopChallenge](https://github.com/OffchainLabs/arbitrum/blob/v0.7.2/packages/arb-bridge-eth/contracts/challenge/InboxTopChallenge.sol) - [0x2C3927aCC3AB8051BAd86B80b9d2D9A1dE5999BD](https://kovan.etherscan.io/address/0x2C3927aCC3AB8051BAd86B80b9d2D9A1dE5999BD) (Template contract)
- [ExecutionChallenge](https://github.com/OffchainLabs/arbitrum/blob/v0.7.2/packages/arb-bridge-eth/contracts/challenge/ExecutionChallenge.sol) - [0x356e19929FCb4973c131d558300E3E353cb8e1C9](https://kovan.etherscan.io/address/0x356e19929FCb4973c131d558300E3E353cb8e1C9) (Template contract)
- [OneStepProof](https://github.com/OffchainLabs/arbitrum/blob/v0.7.2/packages/arb-bridge-eth/contracts/arch/OneStepProof.sol) - [0x082D26eeAa348A7C02291cd1948c66a79fc80aAD](https://kovan.etherscan.io/address/0x082D26eeAa348A7C02291cd1948c66a79fc80aAD)
