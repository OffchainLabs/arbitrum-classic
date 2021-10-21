### Arb-TS

Typescript library for doing Arbitrum stuff, particular stuff like interacting between L1 and L2. Uses [ethers](https://docs.ethers.io/v5/single-page/).

### Quickstart

```ts
const ethProvider = new providers.JsonRpcProvider(ethRPC)
const arbProvider = new providers.JsonRpcProvider(arbRPC)

const connectedL1Wallet = new Wallet(myPrivateKey, ethProvider)
const connectedL2Wallet = new Wallet(myPrivateKey, arbProvider)

const bridge = await Bridge.init(
  connectedL1Wallet,
  connectedL2Wallet
  l1GatewayRouter,
  l2GatewayRouter
)

bridge.depositEth(parseEther('32'))
```

See [integration tests](https://github.com/OffchainLabs/arbitrum/blob/develop/packages/arb-ts/integration_test/arb-bridge.test.ts) for sample usage.

#### Run Integration tests

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

1. Set `PRIVKEY` environmental variable to the key of the account from which you'll be deploying (account should have some balance of the token you're bridging).

1. `yarn bridgeStandardToken`

Required CL params:
`networkID`:number — Chain ID of L1 network
`l1TokenAddress`:string — address of L1 token to be bridged
`amount`:number — token amount to bridge for initial deposit. Raw amount will be used (i.e., not adjusted for decimals) (any non-zero amount will work)

Ex:
`yarn bridgeStandardToken --networkID 4 --l1TokenAddress 0xdf032bc4b9dc2782bb09352007d4c57b75160b15 --amount 3`
