---
id: Public_Testnet
title: Public Testnet Guide
sidebar_label: Public Testnet
---

In order to make it easy for people to get started using Arbitrum Rollup, we've launched our own hosted Arbitrum Rollup chain hosted on the Rinkeby testnet.

For a general introduction to the Arbitrum Testnet, see our [announcement](https://medium.com/offchainlabs/arbitrum-rollup-testnet-full-featured-and-open-to-all-da3255b562ea).

## Connecting to the chain

If you're using Metamask, add a custom RPC network to connect to the Arbitrum testnet:

- Network Name: Arbitrum Testnet
- RPC URL: https://rinkeby.arbitrum.io/rpc
- ChainID: 421611
- Symbol: ETH
- Block Explorer URL: https://testnet.arbiscan.io/

## Observing transactions

If you'd like to see your transactions in action, check out our [block explorer](https://testnet.arbiscan.io/)!

There you'll be able to see all the transactions being executed in Arbitrum and also see exactly how much Ethereum Gas each transaction uses.

## Bridging Eth and ERC-20 Tokens

In order to deposit and withdraw Eth or tokens, visit https://bridge.arbitrum.io.

In order to start using the chain, you'll have deposit Eth from Rinkeby so that you can pay for fees in L2. In order to get Rinkeby Eth, use one of the standard faucets from https://faucet.rinkeby.io/

## Interacting with the chain

Once you've added the Arbitrum Rinkeby Testnet network to Metamask, you should be able to interact with the Arbitrum chain just like you would with Ethereum.

The are a couple things to note on the Arbitrum chain.

- Arbitrum uses an EIP-1559-like gas auction system so the gas price you list in your transaction is a bid, but the actual price may be lower
- The majority of gas costs paid in the arbitrum chain go to pay for the cost of posting your transaction data to Ethereum

## Deploying your contracts

Deploying your contracts onto the Arbitrum testnet is as easy as changing your RPC endpoint to https://rinkeby.arbitrum.io/rpc

For a deeper dive into deploying with truffle see [here](Contract_Deployment.md).

## Porting your frontend

Porting your frontend is just as easy as deploying your contracts. Just take your existing frontend and point it at our RPC endpoint after deploying your contract. For more information and code samples see [here](Frontend_Integration.md)

## Rinkeby Deployment

All contracts are deployed from https://github.com/OffchainLabs/arbitrum/tree/69c58d6b33c4dfb7d8293ccfdcb1675798201b7e/packages/arb-bridge-eth/contracts

### Important Addresses

For a list of useful contract addresses, look at the list [here](Useful_Addresses.md).
