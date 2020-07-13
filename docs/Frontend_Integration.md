---
id: Frontend_Integration
title: Frontend Integration
sidebar_label: Frontend Integration
---

Arbitrum comes with plugins that allow you seamlessly build frontends for Arbitrum base smart contracts using the
same tools you use to build on Ethereum. Most client side smart contract libraries provide an abstraction over direct
blockchain interaction called a provider. Arbitrum provides implementations of these providers, making the libraries
easy to use.

### Javascript

Arbitrum is compatible with both Ethers.js and Web3.js. Providers are created using two parameters: 1) A URL to a validator of the ArbChain 2) An existing Ethereum provider

The validator URL is used for RPC queries requesting off-chain information about the ArbChain. This validator does not need to be trusted, since any client can cheaply verify the correctness of returned results.

#### Web3.js

```js
const ethers = require('ethers')
const ArbProvider = require('arb-provider-ethers').ArbProvider
const ArbEth = require('arb-provider-ethers')
const ProviderBridge = require('arb-ethers-web3-bridge')

const arbProvider = new ArbEth.ArbProvider('http://localhost:1235', ethProvider)
let web3 = new Web3(new ProviderBridge(arbProvider))
```

#### Ethers.js

```js
const ArbProvider = require('arb-provider-ethers').ArbProvider
let provider = new ArbProvider(
  'http://localhost:1235',
  new ethers.providers.Web3Provider(ethProvider)
)
```

### Golang

Arbitrum implements Go-Ethereum's backend interface which allows go-ethereum based go applications to interact with Arbitrum smart contracts. Similar to the javascript interface, this requires a validator URL, a transactor for creating transactions, and a client for connecting to Ethereum.

```golang
import (
    "github.com/ethereum/go-ethereum/common"
    goarbitrum "github.com/offchainlabs/arbitrum/packages/arb-provider-go"
)

client, err := ethclient.Dial(ethURL)
if err != nil {
    return err
}

auth := bind.NewKeyedTransactor(privateKey)
arbBackend, err := goarbitrum.Dial("http://localhost:1235", auth, client)
if err != nil {
    return err
}
testTokenAddress := common.HexToAddress("0x895521964D724c8362A36608AAf09A3D7d0A0445")
testToken, err := NewTestToken(testTokenAddress, arbBackend)
if err != nil {
    return err
}
```
