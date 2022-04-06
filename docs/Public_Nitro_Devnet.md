---
id: Public_Nitro_Devnet
title: Public Nitro Devnet Guide
sidebar_label: Public Nitro Devnet
---

In order to make it easy for people to get started using Arbitrum Nitro, we've launched our own hosted Arbitrum Nitro Rollup chain hosted on the Goerli testnet.

For a general introduction to the Arbitrum Nitro, see our [announcement](https://medium.com/offchainlabs/arbitrum-nitro-sneak-preview-44550d9054f5).

As we approach mainnet launch and incorporate feedback, we will likely reset the devnet a few times, so please plan accordingly.

## Connecting to the chain

If you're using Metamask, add a custom RPC network to connect to the Arbitrum devnet:

- Network Name: Arbitrum Nitro Devnet
- RPC URL: https://nitro-devnet.arbitrum.io/rpc
- ChainID: 421612
- Symbol: ETH
- Block Explorer URL: https://nitro-devnet-explorer.arbitrum.io

- Retryable Dashboard: http://retryable-tx-panel-nitro.arbitrum.io/
- Token Bridge: https://nitro-devnet-bridge.arbitrum.io/

## Observing transactions

If you'd like to see your transactions in action, check out our [block explorer](https://nitro-devnet-explorer.arbitrum.io)!

There you'll be able to see all the transactions being executed in Arbitrum and also see exactly how much Ethereum Gas each transaction uses.

## Bridging Eth and ERC-20 Tokens

In order to deposit and withdraw Eth or tokens, visit https://nitro-devnet-bridge.arbitrum.io/.

You can get ETH to use in the Arbitrum Nitro Devnet by visiting the [twitter faucet](https://twitter.com/intent/tweet?text=ok%20I%20need%20@arbitrum%20to%20give%20me%20Nitro%20devnet%20gas.%20like%20VERY%20SOON.%20I%20cant%20take%20this,%20I%E2%80%99ve%20been%20waiting%20for%20@nitro_devnet%20release.%20I%20just%20want%20to%20start%20developing.%20but%20I%20need%20the%20gas%20IN%20MY%20WALLET%20NOW.%20can%20devs%20DO%20SOMETHING??%20%20SEND%20HERE:%200x-your-eth-address-here).

You may also deposit Eth from Goerli so that you can pay for fees in L2. In order to get Goerli Eth, use one of the standard faucets.

## Interacting with the chain

Once you've added the Arbitrum Nitro Devnet network to Metamask, you should be able to interact with the Arbitrum chain just like you would with Ethereum.

The are a couple things to note on the Arbitrum chain.

- Arbitrum uses an EIP-1559-like gas auction system so the gas price you list in your transaction is a bid, but the actual price may be lower
- The majority of gas costs paid in the arbitrum chain go to pay for the cost of posting your transaction data to Ethereum

## Deploying your contracts

Deploying your contracts onto the Arbitrum Devnet is as easy as changing your RPC endpoint to https://nitro-devnet.arbitrum.io/rpc

For a deeper dive into deploying with truffle see [here](Contract_Deployment.md).

## Porting your frontend

Porting your frontend is just as easy as deploying your contracts. Just take your existing frontend and point it at our RPC endpoint after deploying your contract. For more information and code samples see [here](Frontend_Integration.md)
