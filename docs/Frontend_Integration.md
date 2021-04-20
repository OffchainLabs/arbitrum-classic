---
id: Frontend_Integration
title: Frontend Integration
sidebar_label: Frontend Integration
---

Arbitrum comes with tools to make front-end integration as seamless as possible for Ethereum web developers.

Arbitrum nodes support the [Ethereum JSON-RPC API](https://eth.wiki/json-rpc/API); thus, popular Ethereum libraries for interacting with the Ethereum chain can be used for Arbitrum interactions with little-to-no modifications.

For Ethereum/Arbitrum "bridge" functionality — methods that involve communicating between the L1 Ethereum chain and the L2 Arbitrum chain (i.e., depositing and withdrawing assets), we provide our own libraries for convenience.

## Arbitrum Integration

#### Ethers.js

Ethers-js can be used to interact with an Arbitrum chain exactly as one would use it to interact with Ethereum ([see docs](https://docs.ethers.io/v5/)); simply instantiate a provider connected to an Arbitrum node.

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

The bridging-related contract addresses for our Kovan4 testnet:

```json
L1:
EthErc20Bridge.sol:  0x2948ac43e4AfF448f6af0F7a11F18Bb6062dd271
Inbox.sol:           0xD71d47AD1b63981E9dB8e4A78C0b30170da8a601

L2:
ArbTokenBridge.sol:  0x64b92d4f02cE1b4BDE2D16B6eAEe521E27f28e07
ArbSys.sol:          0x0000000000000000000000000000000000000064
ArbRetryableTx.sol:  0x000000000000000000000000000000000000006E

```

Accessing bridging methods can be done via our `arb-ts` library, or by simply connecting to the relevant contracts directly.

#### 1. arb-ts

**Installation**:

```
yarn add arb-ts ethers-js
```

**Usage (with Ethers-js wallets/providers)**:

```ts
import { providers, Wallet } from 'ethers'
import { Bridge } from 'arb-ts'

const l1Provider = new providers.JsonRpcProvider('http://EthNodeUrl.com')
const l2Provider = new providers.JsonRpcProvider('http://ArbNodeUrl.com')

const l1Signer = new Wallet('0xmyprivatekey!!', l1Provider)
const l2Signer = new Wallet('0xmyprivatekey!!', l2Provider)

const bridge = new Bridge(
  '0xL1EthErc20BridgeAddress',
  '0xl2ArbTokenBridgeAddress',
  l1Signer,
  l2Signer
)
```

See [tests](https://github.com/OffchainLabs/arbitrum/blob/develop/packages/arb-ts/integration_test/arb-bridge.test.ts) for sample usage. Full arb-ts API documentation coming soon.

(Note that we've deprecated the old `arb-provider-ethers` library; arb-ts is the recommended replacement)

#### 2. Alternative: Direct Contract Instantiation

Typechain interfaces for all contracts listed above are made available via arb-ts. i.e.,

```ts
import { ArbSys__factory } from 'arb-ts'

const arbSys = ArbSys__factory.connect(ARB_SYS_ADDRESS, l2Signer)

arbSys.withdrawEth('0xmyaddress')
```
