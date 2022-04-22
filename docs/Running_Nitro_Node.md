---
id: Running_Nitro_Node
title: Running full Nitro node for Arbitrum Goerli DevNet
sidebar_label: Running a Nitro Node
---

### Required Artifacts

- Latest Docker Image: `offchainlabs/nitro-node:v2.0.0-alpha.2`

### Required parameter

- `--l1.url=<Layer 1 Ethereum RPC URL>`
  - Must provide standard Ethereum node RPC endpoint.

### Important ports

- RPC: `8547`
- WebSocket: `8548`
- Sequencer Feed: `9642`

### Putting it all together

- When running docker image, an external volume should be mounted to persist the database across restarts. The mount point should be `/home/user/.arbitrum/goerli`.
- Here is an example of how to run nitro-node for goerli:

  ```
  docker run --rm -it  -v /some/local/dir/arbitrum-goerli/:/home/user/.arbitrum/goerli -p 0.0.0.0:8547:8547 -p 0.0.0.0:8548:8548 offchainlabs/nitro-node:v2.0.0-alpha.2 --l1.url https://l1-goerli-node:8545
  ```

  - Note that if you are running L1 node on localhost, you may need to add `--network host` right after `docker run` to use docker host-based networking

### Note on permissions

- The Docker image is configured to run as non-root UID 1000. This means if you are running in Linux or OSX and you are getting permission errors when trying to run the docker image, run this command to allow all users to update the persistent folders
  ```
  mkdir /some/local/dir/arbitrum-goerli
  chmod -fR 777 /some/local/dir/arbitrum-goerli
  ```

### Optional parameters

- `--http.api`
  - APIs offered over the HTTP-RPC interface (default `net,web3,eth`)
  - Add `debug` to enable tracing
- `--http.corsdomain`
  - Comma separated list of domains from which to accept cross origin requests (browser enforced)
- `--http.vhosts`
  - Comma separated list of virtual hostnames from which to accept requests (server enforced). Accepts `*` wildcard (default `localhost`)
- `--node.archive`
  - Retain past block state
- `--node.feed.input.url=<feed address>`
  - Defaults to `wss://nitro-devnet.arbitrum.io/feed`. If running more than a couple nodes, you will want to provide one feed relay per datacenter, see further instructions below.
- `--node.forwarder.target=<sequencer RPC>`
  - Defaults to `https://nitro-devnet.arbitrum.io/rpc`
- `--node.rpc.evm-timeout`
  - Defaults to `5s`, timeout used for `eth_call` (0 == no timeout)
- `--node.rpc.gas-cap`
  - Defaults to `50000000`, cap on computation gas that can be used in `eth_call`/`estimateGas` (0 = no cap)
- `--node.rpc.tx-fee-cap`
  - Defaults to `1`, cap on transaction fee (in ether) that can be sent via the RPC APIs (0 = no cap)
- `--node.type`
  - Defaults to `forwarder`
  - More information on running validator is provided below

### Arb-Relay

- When running more than one node, you want to run a single arb-relay which can provide a feed for all your nodes.
  The arb-relay is in the same docker image.
- Here is an example of how to run nitro-relay for goerli:
  ```
  docker run --rm -it  -p 0.0.0.0:9642:9642 --entrypoint relay offchainlabs/nitro-node:v2.0.0-alpha.2 --node.feed.input.url wss://nitro-devnet.arbitrum.io/feed
  ```
- Here is an example of how to run nitro-node for goerli with custom relay:
  ```
  docker run --rm -it  -p 0.0.0.0:8547:8547 -p 0.0.0.0:8548:8548 offchainlabs/nitro-node:v2.0.0-alpha.2 --l1.url https://l1-goeri-node:8545 --feed.input.url ws://local-relay-address:9642
  ```
