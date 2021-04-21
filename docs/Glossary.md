---
id: Glossary
title: Glossary of Arbitrum Terms
custom_edit_url: https://github.com/OffchainLabs/arbitrum/edit/master/docs/Glossary.md
---

- ArbOS: Layer 2 "operating system" that trustlessly handles system-level operations
- ArbGas: Unit for pricing execution on Arbitrum; denominated in Ether, Arbitrum's native currency. ArgGas is somewhat analogous to gas on L1 Ethereum, though different factors go into its calculation. [ArbGas and Fees](Inside_Arbitrum.md)
- Arbitrum Full Node: A party who keeps track of the state of an Arbitrum chain and receives remote procedure calls (RPCs) from clients. Analogous to a non-mining L1 Ethereum node.
- Aggregator: An Arbitrum Full Node that also receives transactions from users and submits them in batches on their behalf.
- Assertion: A claim, made by a validator, about what the contracts on a chain will do. An assertion is considered _pending_ until it is confirmed.
- Arbitrum Chain: A chain running on Arbitrum, containing some contracts. Many Arbitrum chains can exist at the same time.
- AVM: The Arbitrum Virtual Machine
- Token Bridge: A series of contracts on Ethereum and Arbitrum for trustlessly moving tokens between the L1 and L2.
- Chain factory: A contract running on Ethereum which, when called, creates a new Arbitrum Chain.
- Chain state: A particular point in the history of an Arbitrum Chain. A chain state corresponds to a sequence of assertions that have been made, and a verdict about whether each of those assertions was accepted.
- Challenge: When two stakers disagree about the correct verdict on an assertion, those stakers can be put in a challenge. The challenge is refereed by the EthBridge. Eventually one staker wins the challenge. The loser forfeits their stake. Half of the loser's stake is given to the winner, and the other half is burned.
- Client: A program running on a user's machine, often in the user's browser, that interacts with contracts on an Arbitrum chain and provides a user interface.
- Confirmation: The final decision by an Arbitrum chain to accept a node as being a settled part of the chain's history. When a node is confirmed, any funds paid out by the chain to the main Ethereum chain are transferred.
- EthBridge: A group of contracts running on the Ethereum chain, which act as a record-keeper and rule enforcer for Arbitrum chains.
- Inbox: Holds a sequence of messages sent by clients to the contracts on an Arbitrum Chain. Every chain's inbox is managed by the on-chain EthBridge.
- Outbox: An L1 contract responsible for tracking outgoing (Arbitrum to Ethereum) messages, including withdrawals, which can be executed by users once they are confirmed.
- Outbox Entry: A Merkle root of a series of outgoing messages posted over some period of time, stored in the Outbox.
- Staker: A party who deposits a stake, in Eth, to vouch for a particular node in an Arbitrum Chain. A party who stakes on a false node can expect to lose their stake. An honest staker can recover their stake once the node they are staked on has been confirmed.
- Sequencer: an entity given rights to reorder transactions in the Inbox over a small window of time, who can thus give clients sub-blocktime soft confirmations.
- Sequencer Chain: An Arbitrum chain that includes Sequencer support.
- Validator: A party who makes staked, disputable assertions about the state of the Arbitrum chain; i.e., proposing state updates or passively monitoring other validator’s assertions and disputing them if they’re invalid.
- Virtual Machine (VM): A program that "runs" on the Arbitrum chain, which tracks the states of all of the contracts on the chain and all of the Eth and tokens deposited into the chain.
