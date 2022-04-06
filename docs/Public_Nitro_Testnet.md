---
id: Public_Nitro_Testnet
title: Public Nitro Testnet Guide
sidebar_label: Public Nitro Testnet
---

In order to make it easy for people to get started using Arbitrum Nitro, we've launched our own hosted Arbitrum Nitro Rollup chain hosted on the Goerli testnet.

For a general introduction to the Arbitrum Nitro, see our [announcement](https://medium.com/offchainlabs/arbitrum-nitro-sneak-preview-44550d9054f5).

## Connecting to the chain

If you're using Metamask, add a custom RPC network to connect to the Arbitrum testnet:

- Network Name: Arbitrum Nitro Testnet
- RPC URL: INSERT_RPC_URL
- ChainID: 421612
- Symbol: ETH
- Block Explorer URL: INSERT_EXPLORER_URL

## Observing transactions

If you'd like to see your transactions in action, check out our [block explorer](INSERT_EXPLORER_URL)!

There you'll be able to see all the transactions being executed in Arbitrum and also see exactly how much Ethereum Gas each transaction uses.

## Bridging Eth and ERC-20 Tokens

In order to deposit and withdraw Eth or tokens, visit INSERT_BRIDGE_URL.

You can get ETH to use in the Arbitrum Nitro testnet by visiting the [twitter faucet](INSERT_FAUCET_URL).

You may also deposit Eth from Goerli so that you can pay for fees in L2. In order to get Goerli Eth, use one of the standard faucets.

## Interacting with the chain

Once you've added the Arbitrum Nitro Testnet network to Metamask, you should be able to interact with the Arbitrum chain just like you would with Ethereum.

The are a couple things to note on the Arbitrum chain.

- Arbitrum uses an EIP-1559-like gas auction system so the gas price you list in your transaction is a bid, but the actual price may be lower
- The majority of gas costs paid in the arbitrum chain go to pay for the cost of posting your transaction data to Ethereum

## Deploying your contracts

Deploying your contracts onto the Arbitrum testnet is as easy as changing your RPC endpoint to INSERT_RPC_URL

For a deeper dive into deploying with truffle see [here](Contract_Deployment.md).

## Porting your frontend

Porting your frontend is just as easy as deploying your contracts. Just take your existing frontend and point it at our RPC endpoint after deploying your contract. For more information and code samples see [here](Frontend_Integration.md)
