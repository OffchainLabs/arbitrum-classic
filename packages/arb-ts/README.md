### Arb-TS

##### Run Integration tests

`yarn test:integration`

Defaults to `kovan4`, for custom network use `--network` flag.

`kovan4` expects env var `DEVNET_PRIVKEY` to be prefunded with at least 0.02 ETH, and env var `INFURA_KEY` to be set.
(see `integration_test/config.ts`)
