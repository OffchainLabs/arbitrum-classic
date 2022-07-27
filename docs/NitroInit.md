---
id: NitroInit
title: Initializing Nitro network from Legacy
sidebar_label: Initializing Nitro network from Legacy
---

Nitro software has the option to initialize a chain with state imported from a legacy arbitrum node. Rinkeby Testnet, and later Arbitrum One, will use this option. Entire state of the blockchain will be preserved when it is upgraded from legacy to nitro. Users have multiple options how to import the state into nitro node.

## Option 1: Init from a seed database

The --init.url option for nitro can accept either local path (prefixed with "file:") or remote URL where an archive containing already-initialized database. This is the simplest, fastest initialization option.

## Option 2: Download exported state and create genesis block in nitro

In this method - nitro node will initialize state and create the genesis block based on data we provide.

### History

History is exported from legacy node in two folders, each containing a geth database:

- **l2chaindata/ancient** cotains all block-headers, transactions and receipts executed in the legacy node. Nitro node uses the history to be able to answer simple requests, like eth_getTransactionReceipt, from the legacy blockchain. The last block in the chain is the only one that affects the genesis block (timestamp is copied from the last block, and parentHash is taken from the last block's blockHash).
- **classic-msg**. This data does not impact consensus and is optional. It allows a nitro node to provide the information required when redeeming a withdrawal made on the classic node.

### State

Blockchain state exported as a series of json files. State read from these json files is composed into the state-root of the genesis block. Using json format allows easier verification of the state using scripts, or manual state initialization for local or test blockchains.

### Running node initialization

- Note: Import state requires more resources than running a normal nitro node. Make sure you are using a powerful enough machine.
- Place l2chaindata and classic-msg (optional) directories in nitro's instance directory - e.g. ${HOME}/.arbitrum/rinkeby-nitro/
- Launch the node with argument `--init.import-file=/path/to/state/index.json`
- If the system is under extreme memory load, try using `--init.accounts-per-sync`. This will write partial state to hard-disk during initialization. A reasonable value to try would be 100000. A system with constrained memory might require a lower value.
- You can add argument `--init.then-quit` to make node quit after init is done

## Option 3: export data from your own classic-node and import it to nitro

- Launch node with the option `--node.rpc.nitroexport.enable=true` note: this is only recommended for nodes with no public/external interfaces. All exportes data will be written to directory "nitroexport" under the legacy instance directory - e.g. ${HOME}/.arbitrum/rinkeby/nitroexport. Export data can be used to initialize a nito node as detailed in option 2.
- Wait for the node to sync up till the end of the classical blockchain.
- Issue rpc call `arb_exportHistory` with parameter "latest" to start exporting blocks. This will return immediately.
- Optional: Issue rpc call `arb_exportOutbox` with parameter "0xffffffffffffffff" to start exporting classical outbox data. This will return immediately.
- Issue rpc call `arb_exportState` with parameter "latest" to export state. Unless disconnected - this will only return after state export is done. State will be created in a separate subdir under the output directory.
- Use `arb_exportHistoryStatus` and `arb_exportOutboxStatus` to see how they progress. These call will return an error if one was encountered during export, and number of block / outbox batch exported otherwise.

Important note: exporting the state on your own classic node should produce the same state as using files supplied by offchain labs (e.g. the same genesis blockhash). However, files themselves will not necessarily be identical. For example - state export is done in parallel, so entries in the files could appear in a different order.
