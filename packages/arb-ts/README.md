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

#### Byte Serializing Solidity Arguments Schema

Arb-ts includes methods for [serializing parameters](https://developer.offchainlabs.com/docs/special_features#parameter-byte-serialization) for a solidity method into a single byte array to minimize calldata. It uses the following schema:

#### address[]:

| field         | size (bytes)       | Description                                                             |
| ------------- | ------------------ | ----------------------------------------------------------------------- |
| length        | 1                  | Size of array                                                           |
| is-registered | 1                  | 1 = all registered, 0 = not all registered                              |
| addresses     | 4 or 20 (x length) | If is registered, left-padded 4-byte integers; otherwise, eth addresses |

#### non-address[]:

| field  | size (bytes) | Description              |
| ------ | ------------ | ------------------------ |
| length | 1            | Size of array            |
| items  | (variable)   | All items (concatenated) |

#### address:

| field         | size (bytes) | Description                                                       |
| ------------- | ------------ | ----------------------------------------------------------------- |
| is-registered | 1            | 1 = registered, 0 = not registered                                |
| address       | 4 or 20      | If registered, left-padded 4-byte integer; otherwise, eth address |
