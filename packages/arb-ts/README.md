# Arb-Ts

Typescript library for client-side interactions with Arbitrum. Arb-ts provides common helper functionaliy as well access to the underlying smart contract interfaces.

Below is an overview of the Arb-Ts functionlity. See the [tutorials](https://github.com/OffchainLabs/arbitrum-tutorials) for examples of how to use these classes.

### Bridging assets

Arb-Ts can be used to bridge assets to/from the rollup chain.The following asset bridgers are currently available:

- EthBridger
- Erc20Bridger

All asset bridgers have the following methods:

- **deposit** - moves assets from the L1 to the L2
- **depositEstimateGas** - estimates the gas required to do the deposit
- **withdraw** - moves assets from the L2 to the L1
- **withdrawEstimateGas** - estimate the gas required to do the withdrawal
  Which accept different parameters depending on the asset bridger type

### Cross chain messages

When assets are moved by the L1 and L2 cross chain messages are sent. The lifecycles of these messages are encapsulated in the classes `L1ToL2Message` and `L2ToL1Message`. These objects are commonly created from the receipts of transactions that send cross chain messages. A cross chain message will eventually result in a transaction being executed on the destination chain, and these message classes provide the ability to wait for that finalizing transaction to occur.

### Networks

Arb-Ts comes pre-configured for the Mainnet and Rinkeby, and their Arbitrum counterparts. However the networks functionlity can be used register networks for custom Arbitrum instances. Most of the classes in Arb-Ts depend on network objects so this must be configured before using other Arb-Ts functionlity.

### Inbox tools

As part of normal operation the Arbitrum sequencer will messages into the rollup chain. However, if the sequencer is unavailable and not posting batches, the inbox tools can be used to force the inclusion of transactions into the rollup chain.

### Utils

- **EventFetcher** - A utility to provide typing for the fetching of events
- **MultiCaller** - A utility for executing multiple calls as part of a single RPC request. This can be useful for reducing round trips.
- **constants** - A list of useful Arbitrum related constants

### Run Integration tests

`yarn test:integration`

Defaults to `rinkArby`, for custom network use `--network` flag.

`rinkArby` expects env var `DEVNET_PRIVKEY` to be prefunded with at least 0.02 ETH, and env var `INFURA_KEY` to be set.
(see `integration_test/config.ts`)

### Bridge A Standard Token

Bridging new a token to L2 (i.e., deploying a new token contract) through the standard gateway is done by simply depositing a token that hasn't yet been bridged. This repo includes a script to trigger this initial deposit/deployment:

1. clone `arbitrum` monorepo

1. `git submodule update --init --recursive`

1. `yarn install` (from root)

1. `cd packages/arb-ts`

1. Set `PRIVKEY` environmental variable (you can use .env) to the key of the account from which you'll be deploying (account should have some balance of the token you're bridging).

1. Set MAINNET_RPC environmental variable to L1 RPC endpoint (i.e., https://mainnet.infura.io/v3/my-infura-key)

1. `yarn bridgeStandardToken`

Required CL params:
`networkID`:number — Chain ID of L1 network
`l1TokenAddress`:string — address of L1 token to be bridged

Ex:
`yarn bridgeStandardToken --networkID 4 --l1TokenAddress 0xdf032bc4b9dc2782bb09352007d4c57b75160b15 --amount 3`
