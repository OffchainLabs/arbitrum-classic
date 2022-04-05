---
id: Running_Node
title: Running full node for Arbitrum One
sidebar_label: Running a Node
---

### Required Artifacts

- Latest Docker Image: offchainlabs/arb-node:v1.3.0-d994f7d

### Required parameter

- `--l1.url=<Layer 1 Ethereum RPC URL>`
  - Must provide standard Ethereum node RPC endpoint.

### Important ports

- RPC: `8547`
- WebSocket: `8548`
- Sequencer Feed: `9642`

### Putting it all together

- When running docker image, an external volume should be mounted to persist the database across restarts. The mount point should be `/home/user/.arbitrum/mainnet` or `/home/user/.arbitrum/rinkeby` depending on what chain you are connecting to.
- Here is an example of how to run arb-node for mainnet:
  ```
  docker run --rm -it  -v /some/local/dir/arbitrum-mainnet/:/home/user/.arbitrum/mainnet -p 0.0.0.0:8547:8547 -p 0.0.0.0:8548:8548 offchainlabs/arb-node:v1.3.0-d994f7d --l1.url https://l1-node:8545
  ```
- Here is an example of how to run arb-node for rinkeby:
  ```
  docker run --rm -it  -v /some/local/dir/arbitrum-rinkeby/:/home/user/.arbitrum/rinkeby -p 0.0.0.0:8547:8547 -p 0.0.0.0:8548:8548 offchainlabs/arb-node:v1.3.0-d994f7d --l1.url https://l1-rinkeby-node:8545
  ```

### Note on permissions

- The Docker image is configured to run as non-root UID 1000. This means if you are running in Linux and you are getting permission errors when trying to run the docker image, run this command to allow all users to update the persistent folders
  ```
  mkdir /some/local/dir/arbitrum-mainnet
  chmod -fR 777 /some/local/dir/arbitrum-mainnet
  ```
  ```
  mkdir /some/local/dir/arbitrum-rinkeby
  chmod -fR 777 /some/local/dir/arbitrum-rinkeby
  ```

### Optional parameters

- `--feed.input.url=<feed address>`
  - Will default to `https://arb1.arbitrum.io/feed` or `https://rinkeby.arbitrum.io/feed` depending on chain ID reported by ethereum node provided. If running more than a couple nodes, you will want to provide one feed relay per datacenter, see further instructions below.
- `--node.forwarder.target=<sequencer RPC>`
  - Will default to `https://arb1.arbitrum.io/rpc` or `https://rinkeby.arbitrum.io/rpc` depending on chain ID reported by ethereum node provided.
- `--core.cache.timed-expire`
  - Defaults to `20m`, or 20 minutes. Age of oldest blocks to hold in cache so that disk lookups are not required
- `--node.rpc.max-call-gas`
  - Maximum amount of gas that a node will use in call, default is `5000000`
- `--node.rpc.enable-l1-calls`
  - This option enables the ability to request L1 inclusion information about a transaction by including the argument `returnL1InboxBatchInfo` in a `eth_getTransactionReceipt` request
    - Example: `curl http://arbnode -X POST -H "Content-Type: application/json" -d '{"jsonrpc":"2.0","method":"eth_getTransactionReceipt","params": ["txhash", {"returnL1InboxBatchInfo": true}],"id":1}'`
- `--core.checkpoint-gas-frequency`
  - Defaults to `1000000000`. Amount of gas between saving checkpoints to disk. When making archive queries node has to load closest previous checkpoint and then execute up to the requested block. The farther apart the checkpoints, the longer potential execution required. However, saving checkpoints more often slows down the node in general.
- `--node.cache.allow-slow-lookup`
  - When this option is present, will load old blocks from disk if not in memory cache
  - If archive support is desired, recommend using `--node.cache.allow-slow-lookup --core.checkpoint-gas-frequency=156250000`
- `--node.rpc.tracing.enable`
  - Note that you also need to have a database populated with an archive node if you want to trace previous transactions
  - This option enables the ability to call a tracing api which is inspired by the parity tracing API with some differences
    - Example: `curl http://arbnode -X POST -H "Content-Type: application/json" -d '{"jsonrpc":"2.0","method":"arbtrace_call","params":[{"to": "0x6b175474e89094c44da98b954eedeac495271d0f","data": "0x70a082310000000000000000000000006E0d01A76C3Cf4288372a29124A26D4353EE51BE"},["trace"], "latest"],"id":67}'`
  - The `trace_*` methods are renamed to `arbtrace_*`, except `trace_rawTransaction` is not supported
  - Only `trace` type is supported. `vmTrace` and `stateDiff` types are not supported
  - The self-destruct opcode is not included in the trace. To get the list of self-destructed contracts, you can provide the `deletedContracts` parameter to the method

### Arb-Relay

- When running more than one node, you want to run a single arb-relay which can provide a feed for all your nodes.
  The arb-relay is in the same docker image.
- Here is an example of how to run arb-relay for mainnet:
  ```
  docker run --rm -it  -v /some/local/dir/arbitrum-mainnet/:/home/user/.arbitrum/mainnet -p 0.0.0.0:9642:9642 --entrypoint /home/user/go/bin/arb-relay offchainlabs/arb-node:v1.3.0-d994f7d --feed.input.url wss://arb1.arbitrum.io/feed
  ```
- Here is an example of how to run arb-node for mainnet with custom relay:
  ```
  docker run --rm -it  -v /some/local/dir/arbitrum-mainnet/:/home/user/.arbitrum/mainnet -p 0.0.0.0:8547:8547 -p 0.0.0.0:8548:8548 offchainlabs/arb-node:v1.3.0-d994f7d --l1.url https://l1-node:8545 --feed.input.url ws://local-relay-address:9642
  ```
