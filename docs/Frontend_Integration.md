---
id: Frontend_Integration
title: Frontend Integration
sidebar_label: Frontend Integration
---

Arbitrum comes with tools to make front-end integration as seamless as possible for Ethereum web developers.

Arbitrum nodes support the [Ethereum JSON-RPC API](https://eth.wiki/json-rpc/API); thus, popular Ethereum libraries for interacting with the Ethereum chain can be used for Arbitrum interactions with little-to-no modifications.

For Ethereum/Arbitrum "bridge" functionality — methods that involve communicating between the L1 Ethereum chain and the L2 Arbitrum chain (i.e., depositing and withdrawing assets), we provide our own libraries for convenience.

## Demos

See our [Tutorials](https://github.com/OffchainLabs/arbitrum-tutorials) repo for client-side integration demos.

## Arbitrum Integration

#### Ethers.js

[`ethers`](https://www.npmjs.com/package/ethers) can be used to interact with an Arbitrum chain exactly as one would use it to interact with Ethereum ([see docs](https://docs.ethers.io/v5/)); simply instantiate a provider connected to an Arbitrum node.

I.e., with MetaMask already connected to an Arbitrum node via a custom RPC url:

```ts
import * as ethers from 'ethers'

const arbProvider = new ethers.providers.Web3Provider(window.ethereum)
```

Or instantiate a provider directly via an Arbitrum node's URL

```ts
import * as ethers from 'ethers'
const arbProvider = new ethers.providers.JsonRpcProvider(
  'http://ArbNodeUrl.com'
)
```

#### Web3.js

Likewise, a Web3 provider can be instantiated directly via an Arbitrum node url:
https://web3js.readthedocs.io/en/v1.2.11/index.html

```ts
import * as Web3 from 'web3'
var arbWeb3Provider = new Web3('http://ArbNodeUrl.com')
```

## Arbitrum / Ethereum Bridge

Accessing bridging methods can be done via our `@arbitrum/sdk` library, or by simply connecting to the relevant contracts directly.

#### @arbitrum/sdk

**Installation**:

```
yarn add @arbitrum/sdk
```

**Usage (with Ethers wallets/providers)**:

```ts
import { getL2Network, EthBridger } from '@arbitrum/sdk'

const l2Network = await getL2Network(
  l2ChainID /** <-- chain id of target Arbitrum chain */
)
const ethBridger = new EthBridger(l2Network)
```

