---
id: Public_Nitro_Devnet
title: Public Nitro Testnet Guide
sidebar_label: Public Nitro Rollup Testnet
---


***Note: Frormer Nitro Devnet chain 421612 has been deprecated in favor of chain 421613:***

In order to make it easy for people to get started using Arbitrum Nitro, we've launched our own hosted Arbitrum Nitro Rollup chain hosted on the Goerli testnet.

For a general introduction to the Arbitrum Nitro, see our [announcement](https://medium.com/offchainlabs/arbitrum-nitro-sneak-preview-44550d9054f5).

As we approach mainnet launch and incorporate feedback, we will likely reset the devnet a few times, so please plan accordingly.

## Connecting to the chain

Connect your wallet to the devnet; if [your wallet](https://portal.arbitrum.one/#wallets) requires it, add the Arbitrum Nitro Devnet as a custom RPC:

- Network Name: Arbitrum Nitro Devnet
- RPC URL: goerli-rollup.arbitrum.io/rpc 
- ChainID: 421613
- Symbol: ETH
- Block Explorer URL: https://goerli-rollup-explorer.arbitrum.io/

- Retryable Dashboard: http://retryable-tx-panel-nitro.arbitrum.io/
- Token Bridge: https://bridge.arbitrum.io/

## Observing transactions

If you'd like to see your transactions in action, check out our [block explorer](https://goerli-rollup-explorer.arbitrum.io/)!

There you'll be able to see all the transactions being executed in Arbitrum and also see exactly how much Ethereum Gas each transaction uses.

## Bridging Eth and ERC-20 Tokens

In order to deposit and withdraw Eth or tokens, visit https://bridge.arbitrum.io/.

You can get ETH to use in the Arbitrum Nitro Devnet by visiting the [twitter faucet](https://twitter.com/intent/tweet?text=ok%20I%20need%20@arbitrum%20to%20give%20me%20Nitro%20devnet%20gas.%20like%20VERY%20SOON.%20I%20cant%20take%20this,%20I%E2%80%99ve%20been%20waiting%20for%20@nitro_devnet%20release.%20I%20just%20want%20to%20start%20developing.%20but%20I%20need%20the%20gas%20IN%20MY%20WALLET%20NOW.%20can%20devs%20DO%20SOMETHING??%20%20SEND%20HERE:%200x-your-eth-address-here).

You may also deposit Eth from Goerli so that you can pay for fees in L2. In order to get Goerli Eth, use one of the standard faucets.

## Interacting with the chain

Once your wallet is connect to the Arbitrum Nitro Devnet, you should be able to interact with the chain just like you would with Ethereum.

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
| Rollup              | [0x45e5cAea8768F42B385A366D3551Ad1e0cbFAb17](https://goerli.etherscan.io/address/0x45e5cAea8768F42B385A366D3551Ad1e0cbFAb17) |
| Delayed Inbox       | [0x6BEbC4925716945D46F0Ec336D5C2564F419682C](https://goerli.etherscan.io/address/0x6BEbC4925716945D46F0Ec336D5C2564F419682C) |
| Sequencer Inbox     | [0x0484A87B144745A2E5b7c359552119B6EA2917A9](https://goerli.etherscan.io/address/0x0484A87B144745A2E5b7c359552119B6EA2917A9) |
| Bridge              | [0xaf4159A80B6Cc41ED517DB1c453d1Ef5C2e4dB72](https://goerli.etherscan.io/address/0xaf4159A80B6Cc41ED517DB1c453d1Ef5C2e4dB72) |
| Outbox              | [0x45Af9Ed1D03703e480CE7d328fB684bb67DA5049](https://goerli.etherscan.io/address/0x45Af9Ed1D03703e480CE7d328fB684bb67DA5049) |
| OneStepProver0      | [0xD7422f07fe48f6e82E40587feb2acaE1451f08A6](https://goerli.etherscan.io/address/0xD7422f07fe48f6e82E40587feb2acaE1451f08A6) |
| OneStepProverMemory | [0x9221854E95283670E58738805a2d20405d17682E](https://goerli.etherscan.io/address/0x9221854E95283670E58738805a2d20405d17682E) |
| OneStepProverMath   | [0xFe18aB9B105a8C13Fbd67a0DaCb1C70e84Bb5d5E](https://goerli.etherscan.io/address/0xFe18aB9B105a8C13Fbd67a0DaCb1C70e84Bb5d5E) |
| OneStepProverHostIo | [0x5518772ddb8e65416c6572E28BE58dAfc8A3834c](https://goerli.etherscan.io/address/0x5518772ddb8e65416c6572E28BE58dAfc8A3834c) |
| OneStepProofEntry   | [0xe46a0585C3Cb05AaE200161534Af1aE5Dff61294](https://goerli.etherscan.io/address/0xe46a0585C3Cb05AaE200161534Af1aE5Dff61294) |

### Token Bridge

**IMPORTANT**: _Do **not** simply transfer tokens or Ether to any of the addresses below; it will result in loss of funds._

_Users should only interact with the token bridge via dapp interfaces like https://bridge.arbitrum.io_.

|                       | Goerli / Nitro Devnet                                                                                                 |
| --------------------- | --------------------------------------------------------------------------------------------------------------------- |
| L1 Gateway Router     | [0x4c7708168395aEa569453Fc36862D2ffcDaC588c](https://goerli.etherscan.io/address/0x4c7708168395aEa569453Fc36862D2ffcDaC588c) |
| L2 Gateway Router     | [0xE5B9d8d42d656d1DcB8065A6c012FE3780246041](https://goerli-rollup-explorer.arbitrum.io/address/0xE5B9d8d42d656d1DcB8065A6c012FE3780246041)  |
| L1 ERC20 Gateway      | [0x715D99480b77A8d9D603638e593a539E21345FdF](https://goerli.etherscan.io/address/0x715D99480b77A8d9D603638e593a539E21345FdF) |
| L2 ERC20 Gateway      | [0x2eC7Bc552CE8E51f098325D2FcF0d3b9d3d2A9a2](https://goerli-rollup-explorer.arbitrum.io/address/0x2eC7Bc552CE8E51f098325D2FcF0d3b9d3d2A9a2)  |
| L1 Arb-Custom Gateway | [0x9fDD1C4E4AA24EEc1d913FABea925594a20d43C7](https://goerli.etherscan.io/address/0x9fDD1C4E4AA24EEc1d913FABea925594a20d43C7) |
| L2 Arb-Custom Gateway | [0x8b6990830cF135318f75182487A4D7698549C717](https://goerli-rollup-explorer.arbitrum.io/address/0x8b6990830cF135318f75182487A4D7698549C717)  |
| L1 Weth Gateway       | [0x6e244cD02BBB8a6dbd7F626f05B2ef82151Ab502](https://goerli.etherscan.io/address/0x6e244cD02BBB8a6dbd7F626f05B2ef82151Ab502) |
| L2 Weth Gateway       | [0xf9F2e89c8347BD96742Cc07095dee490e64301d6](https://goerli-rollup-explorer.arbitrum.io/address/0xf9F2e89c8347BD96742Cc07095dee490e64301d6)  |
| L1 Weth               | [0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6](https://goerli.etherscan.io/address/0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6) |
| L2 Weth               | [0xe39Ab88f8A4777030A534146A9Ca3B52bd5D43A3](https://goerli-rollup-explorer.arbitrum.io/address/0xe39Ab88f8A4777030A534146A9Ca3B52bd5D43A3)  |

### Arbitrum Precompiles (L2, same on all Arb-chains)

|                  | Address                                                                                                              |
| ---------------- | -------------------------------------------------------------------------------------------------------------------- |
| ArbSys           | [0x0000000000000000000000000000000000000064](https://goerli-rollup-explorer.arbitrum.io/address/0x0000000000000000000000000000000000000064) |
| ArbRetryableTx   | [0x000000000000000000000000000000000000006E](https://goerli-rollup-explorer.arbitrum.io/address/0x000000000000000000000000000000000000006E) |
| ArbGasInfo       | [0x000000000000000000000000000000000000006C](https://goerli-rollup-explorer.arbitrum.io/address/0x000000000000000000000000000000000000006C) |
| ArbAddressTable  | [0x0000000000000000000000000000000000000066](https://goerli-rollup-explorer.arbitrum.io/address/0x0000000000000000000000000000000000000066) |
| ArbStatistics    | [0x000000000000000000000000000000000000006F](https://goerli-rollup-explorer.arbitrum.io/address/0x000000000000000000000000000000000000006F) |
| NodeInterface    | [0x00000000000000000000000000000000000000C8](https://goerli-rollup-explorer.arbitrum.io/address/0x00000000000000000000000000000000000000C8) |
| ArbBLS           | [0x0000000000000000000000000000000000000067](https://goerli-rollup-explorer.arbitrum.io/address/0x0000000000000000000000000000000000000067) |
| ArbInfo          | [0x0000000000000000000000000000000000000065](https://goerli-rollup-explorer.arbitrum.io/address/0x0000000000000000000000000000000000000065) |
| ArbAggregator    | [0x000000000000000000000000000000000000006D](https://goerli-rollup-explorer.arbitrum.io/address/0x000000000000000000000000000000000000006D) |
| ArbFunctionTable | [0x0000000000000000000000000000000000000068](https://goerli-rollup-explorer.arbitrum.io/address/0x0000000000000000000000000000000000000068) |

### Misc

|              | Nitro Devnet                                                                                                | 
| ------------ | -------------------------------------------------------------------------------------------------------------------- | 
| L2 Multicall | [0x108B25170319f38DbED14cA9716C54E5D1FF4623](https://goerli-rollup-explorer.arbitrum.io/address/0x108B25170319f38DbED14cA9716C54E5D1FF4623) | 
