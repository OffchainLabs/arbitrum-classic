---
id: Frontend_Integration
title: Frontend Integration
sidebar_label: Frontend Integration
---

Arbitrum comes with tools to make front-end integration as seamless as possible for Ethereum web developers.

Arbitrum aggregators support the [Ethereum JSON-RPC API](https://eth.wiki/json-rpc/API); thus, popular Ethereum libraries for interacting with the Ethereum chain can be used for Arbitrum interactions with little to no modifications.

For Ethereum/Arbitrum "bridge" functionality — methods that involve communicating between the L1 Ethereum chain and the L2 Arbitrum chain (i.e., depositing and withdrawing assets), we provide our own libraries for convenience.

## Arbitrum Integration

#### Ethers.js (recommended)

Ethers-js can be used to interact with an Arbitrum chain exactly as one would use it to interact with Ethereum ([see docs](https://docs.ethers.io/v5/)); simply instantiate a provider connected to an Arbitrum aggregator.

I.e., with MetaMask already connected to an Arbitrum aggregator via a custom RPC url:

```ts
import * as ethers from 'ethers'

const arbProvider = new ethers.providers.Web3Provider(window.ethereum)
```

Or instantiate a provider directly via a aggregator's URL

```ts
import * as ethers from 'ethers'
const arbProvider = new ethers.providers.JsonRpcProvider(
  'http://ArbAggregatorUrl.com'
)
```

#### Web3.js

Likewise, a Web3 provider can be instantiated directly via an Arbitrum aggregator url:
https://web3js.readthedocs.io/en/v1.2.11/index.html

```ts
import * as Web3 from "web3";
var arbWeb3Provider = new Web3('http://ArbAggregatorUrl.com);

```

## Arbitrum / Ethereum Bridge

Arbitrum offers two options for accessing the Arbitrum/Ethereum bridge methods:

#### 1. React Hook Bridge SDK

**Installation**:

```
yarn add token-bridge-sdk
```

**Usage**:

```ts
import { useArbTokenBridge } from 'arb-token-bridge'

const App = () => {
  const bridge = useArbTokenBridge(
    ethProvider,
    arbProvider,
    rollupAddress,
    ethSigner,
    arbSigner
  )
}
```

See [token-bridge-sdk documentation](https://bridgedocs.offchainlabs.com) for full API.

#### 2. Arb Provider Ethers Bridge (minimal)

**Installation**:

```
yarn add arb-provider-ethers
```

**Usage**:

```ts
import { L1Bridge, withdrawEth } from 'arb-provider-ethers'

/* instatiate an ethers-js signer (https://docs.ethers.io/v5/api/signer/) for your ethereum provider*/
const ethSigner = ethProvider.getSigner(0)
/* instatiate an ethers-js signer for your arbitrum provider*/
const arbSigner = arbProvider.getSigner(0)

const l1Bridge = new L1Bridge(ethSigner, '0xArbChainAddress')

// deposit eth:
const txnResponse = await l1Bridge.depositETH('0xwalletaddress', weiValue)

//withdraw eth
const txnResponse = await withdrawEth(arbSigner, weiValue)
```

See [l1Bridge](https://github.com/OffchainLabs/arbitrum/blob/develop/packages/arb-provider-ethers/src/lib/l1bridge.ts) and [l2Bridge](https://github.com/OffchainLabs/arbitrum/blob/develop/packages/arb-provider-ethers/src/lib/l2bridge.ts) for available methods (documentation coming soon.)
