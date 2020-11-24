# arb-provider-ethers

Arbitrum browser provider for ethers.

Arbitrum technologies are patent pending. This repository is offered under the Apache 2.0 license. See LICENSE for details.

## Testing

Run with `yarn test --coverage`

View coverage results with `open coverage/lcov-report/index.html` or `xdg-open coverage/lcov-report/index.html`

## Byte Serializing Solidity Arguements Schema

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
