---
id: Time_in_Arbitrum
title: Block Numbers and Time In Arbitrum
sidebar_label: Block Numbers and Time
---

As in Ethereum, Arbitrum clients submit transactions they want to see happen, and the system (usually) executes those transactions at some later time.
In Arbitrum Rollup, clients submit transactions by posting messages to the Ethereum chain, either directly or through an aggregator, or on a Sequencer chain, though a Sequencer.
These messages are put into the chain's _inbox_.

Messages in an ArbChain's inbox are processed in order. Generally, some time will elapse between the time when a message is put into the inbox (and timestamped) and the time when the contract processes the message and carries out the transaction requested by the message.

## Block Numbers: Arbitrum vs. Ethereum

Arbitrum blocks are assigned their own block numbers, distinct from Ethereum's L1 block numbers.

A single Ethereum block could include within it multiple Arbitrum blocks (if, say, the Arbitrum chain is getting heavy activity); however, an Arbitrum block cannot span across multiple Ethereum blocks. Thus, any given Arbitrum transaction is associated with exactly one Ethereum block and one Arbitrum block.

## Ethereum Block Numbers Within Arbitrum

On a [Sequencer](Inside_Arbitrum.md#sequencer-mode) chain, accessing block numbers within an Arbitrum smart contract (i.e., `block.number` in Solidity) will return a value _close to_ (but not necessarily exactly) the L1 block number at which the Sequencer received the transaction. On a non-sequencer chain, `block.number` will return the L1 block number at which the transaction was inserted into the L1 inbox.

```sol
// some Arbitrum contract:
block.number // => returns L1 block number ("ish")
```

As a general rule, any timing assumptions a contract makes about block numbers and timestamps should be considered generally reliable in the longer term (i.e., on the order of at least several hours) but unreliable in the shorter term (minutes). (It so happens these are generally the same assumptions one should operate under when using block numbers directly on Ethereum!)

## Arbitrum Block Numbers

Arbitrum blocks have their own block numbers, starting at 0 at the Arbitrum genesis block and updating sequentially.

ArbOS and the Sequencer are responsible for delineating when one Arbitrum block ends and the next one begins; one should expect to see Arbitrum blocks produced at a relatively steady rate.

A client that queries an Arbitrum node's RPC interface (for, ie., transaction receipts) will receive the transaction's Arbitrum block number as the standard block number field. The L1 block number will also be included in the added `l1BlockNumber field`.

```ts
const txnReceipt = await arbitrumProvider.getTransactionReceipt('0x...')
/** 
    txnReceipt.blockNumber => Arbitrum block number
    txnReceipt.l1BlockNumber => L1 block number ("ish")
*/
```

The Arbitrum block number can also be retrieved within an Arbitrum contract via [ArbSys](./sol_contract_docs/md_docs/arb-os/arbos/builtin/ArbSys.md):

```sol
 ArbSys(100).arbBlockNumber() // returns Arbitrum block number
```

## Case Study: Multicall

The Multicall contract offers a great case study for the difference between L1 and L2 block numbers.

The [canonical implementation](https://github.com/makerdao/multicall/) of Multicall returns the value of `block.number`. If attempting to use out-of-the-box, some applications might face unintended behaviour.

You can find a deployed version of the adapted Multicall2 at [0x842eC2c7D803033Edf55E478F461FC547Bc54EB2](https://arbiscan.io/address/0x842eC2c7D803033Edf55E478F461FC547Bc54EB2#code).

By default the `getBlockNumber`, `tryBlockAndAggregate`, and `aggregate` functions return the L2 block number. This allows you to use this value to compare your state against the tip of the chain.

The `getL1BlockNumber` function can be queried if applications need to surface the L1 block number.
