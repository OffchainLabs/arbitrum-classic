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

## Deployed Contracts

### Protocol (L1)

|                     | Goerli                                                                                                                |
| ------------------- | --------------------------------------------------------------------------------------------------------------------- |
| Rollup              | [0x767CfF8D8de386d7cbe91DbD39675132ba2f5967](https://goerli.etherscan.io/address/0x767CfF8D8de386d7cbe91DbD39675132ba2f5967) |
| Delayed Inbox       | [0x1fdbbcc914e84af593884bf8e8dd6877c29035a2](https://goerli.etherscan.io/address/0x1fdbbcc914e84af593884bf8e8dd6877c29035a2) |
| Sequencer Inbox     | [0xb32f4257e05c56c53d46bbec9e85770eb52425d6](https://goerli.etherscan.io/address/0xb32f4257e05c56c53d46bbec9e85770eb52425d6) |
| Bridge              | [0x9903a892da86c1e04522d63b08e5514a921e81df](https://goerli.etherscan.io/address/0x9903a892da86c1e04522d63b08e5514a921e81df) |
| Outbox              | [0xFDF2B11347dA17326BAF30bbcd3F4b09c4719584](https://goerli.etherscan.io/address/0xFDF2B11347dA17326BAF30bbcd3F4b09c4719584) |

### Token Bridge

**IMPORTANT**: _Do **not** simply transfer tokens or Ether to any of the addresses below; it will result in loss of funds._

_Users should only interact with the token bridge via dapp interfaces like https://nitro-devnet-bridge.arbitrum.io.

|                       | Goerli / Nitro Devnet                                                                                                 |
| --------------------- | --------------------------------------------------------------------------------------------------------------------- |
| L1 Gateway Router     | [0x8BDFa67ace22cE2BFb2fFebe72f0c91CDA694d4b](https://goerli.etherscan.io/address/0x8BDFa67ace22cE2BFb2fFebe72f0c91CDA694d4b) |
| L2 Gateway Router     | [0xC502Ded1EE1d616B43F7f20Ebde83Be1A275ca3c](https://nitro-devnet-explorer.arbitrum.io/address/0xC502Ded1EE1d616B43F7f20Ebde83Be1A275ca3c)  |
| L1 ERC20 Gateway      | [0x6336C4e811b2f7D17d45b6241Fd47F2E11621Ffb](https://goerli.etherscan.io/address/0x6336C4e811b2f7D17d45b6241Fd47F2E11621Ffb) |
| L2 ERC20 Gateway      | [0xf298434ffE691400b932f4b14B436f451F4CED76](https://nitro-devnet-explorer.arbitrum.io/address/0xf298434ffE691400b932f4b14B436f451F4CED76)  |
| L1 Arb-Custom Gateway | [0x23D4e0D7Cb7AE7CF745E82262B17eb46535Ae819](https://goerli.etherscan.io/address/0x23D4e0D7Cb7AE7CF745E82262B17eb46535Ae819) |
| L2 Arb-Custom Gateway | [0x7AC493f26EF26904E52fE46C8DaEE247b9A556B8](https://nitro-devnet-explorer.arbitrum.io/address/0x7AC493f26EF26904E52fE46C8DaEE247b9A556B8)  |
| L1 Weth Gateway       | [0x64bfF696bE6a087A81936b9a2489624015381be4](https://goerli.etherscan.io/address/0x64bfF696bE6a087A81936b9a2489624015381be4) |
| L2 Weth Gateway       | [0xf10c7CAA33A3360f60053Bc1081980f62567505F](https://nitro-devnet-explorer.arbitrum.io/address/0xf10c7CAA33A3360f60053Bc1081980f62567505F)  |
| L1 Weth               | [0xb4fbf271143f4fbf7b91a5ded31805e42b2208d6](https://goerli.etherscan.io/address/0xb4fbf271143f4fbf7b91a5ded31805e42b2208d6) |
| L2 Weth               | [0x96CfA560e7332DebA750e330fb6f59E2269f40Dd](https://nitro-devnet-explorer.arbitrum.io/address/0x96CfA560e7332DebA750e330fb6f59E2269f40Dd)  |

### Arbitrum Precompiles (L2, same on all Arb-chains)

|                  | Address                                                                                                              |
| ---------------- | -------------------------------------------------------------------------------------------------------------------- |
| ArbSys           | [0x0000000000000000000000000000000000000064](https://nitro-devnet-explorer.arbitrum.io/address/0x0000000000000000000000000000000000000064) |
| ArbRetryableTx   | [0x000000000000000000000000000000000000006E](https://nitro-devnet-explorer.arbitrum.io/address/0x000000000000000000000000000000000000006E) |
| ArbGasInfo       | [0x000000000000000000000000000000000000006C](https://nitro-devnet-explorer.arbitrum.io/address/0x000000000000000000000000000000000000006C) |
| ArbAddressTable  | [0x0000000000000000000000000000000000000066](https://nitro-devnet-explorer.arbitrum.io/address/0x0000000000000000000000000000000000000066) |
| ArbStatistics    | [0x000000000000000000000000000000000000006F](https://nitro-devnet-explorer.arbitrum.io/address/0x000000000000000000000000000000000000006F) |
| NodeInterface    | [0x00000000000000000000000000000000000000C8](https://nitro-devnet-explorer.arbitrum.io/address/0x00000000000000000000000000000000000000C8) |
| ArbBLS           | [0x0000000000000000000000000000000000000067](https://nitro-devnet-explorer.arbitrum.io/address/0x0000000000000000000000000000000000000067) |
| ArbInfo          | [0x0000000000000000000000000000000000000065](https://nitro-devnet-explorer.arbitrum.io/address/0x0000000000000000000000000000000000000065) |
| ArbAggregator    | [0x000000000000000000000000000000000000006D](https://nitro-devnet-explorer.arbitrum.io/address/0x000000000000000000000000000000000000006D) |
| ArbFunctionTable | [0x0000000000000000000000000000000000000068](https://nitro-devnet-explorer.arbitrum.io/address/0x0000000000000000000000000000000000000068) |

### Misc

|              | Goerli / Nitro Devnet                                                                                                | 
| ------------ | -------------------------------------------------------------------------------------------------------------------- | 
| L2 Multicall | [0x1068dbfcc13f3a22fcAe684943AFA43cc66fA689](https://nitro-devnet-explorer.arbitrum.io/address/0x1068dbfcc13f3a22fcAe684943AFA43cc66fA689) | 
