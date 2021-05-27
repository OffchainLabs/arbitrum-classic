### Arb-TS

Typescript library for doing Arbitrum stuff, particular stuff like interacting between L1 and L2. Uses [ethers-js](https://docs.ethers.io/v5/single-page/).

### Quickstart

```ts
const ethProvider = new providers.JsonRpcProvider(ethRPC)
const arbProvider = new providers.JsonRpcProvider(arbRPC)

const connectedL1Wallet = new Wallet(myPrivateKey, ethProvider)
const connectedL2Wallet = new Wallet(myPrivateKey, arbProvider)

const bridge = new Bridge(
  erc20BridgeAddress,
  arbTokenBridgeAddress,
  connectedL1Wallet,
  connectedL2Wallet
)

bridge.depositEth(parseEther('32'))
```

See [integration tests](https://github.com/OffchainLabs/arbitrum/blob/develop/packages/arb-ts/integration_test/arb-bridge.test.ts) for sample usage.

#### Run Integration tests

`yarn test:integration`

Defaults to `kovan4`, for custom network use `--network` flag.

`kovan4` expects env var `DEVNET_PRIVKEY` to be prefunded with at least 0.02 ETH, and env var `INFURA_KEY` to be set.
(see `integration_test/config.ts`)
