---
id: Glossary
title: Glossary of Arbitrum Terms
custom_edit_url: https://github.com/OffchainLabs/arbitrum/edit/master/docs/Glossary.md
---
- Aggregator: A party who keeps track of the state of an Arbitrum chain, receives remote procedure calls (RPCs) from clients, and submits batches of transactions on behalf of users; analogous to a non-mining L1 Ethereum node.
- Assertion: A claim, made by a validator, about what the contracts on a chain will do. An assertion is considered _pending_ until it is confirmed.
- Arbitrum Chain: A chain running on Arbitrum, containing some contracts. Many Arbitrum chains can exist at the same time.
- AVM: The Arbitrum Virtual Machine
- Chain factory: A contract running on Ethereum which, when called, creates a new Arbitrum Chain.
- Chain state: A particular point in the history of an Arbitrum Chain. A chain state corresponds to a sequence of assertions that have been made, and a verdict about whether each of those assertions was accepted.
- Challenge: When two stakers disagree about the correct verdict on an assertion, those stakers can be put in a challenge. The challenge is refereed by the EthBridge. Eventually one staker wins the challenge. The loser forfeits their stake. Half of the loser's stake is given to the winner, and the other half is burned.
- Client: A program running on a user's machine, often in the user's browser, that interacts with contracts on an Arbitrum chain and provides a user interface.
- Confirmation: The final decision by an Arbitrum chain to accept a node as being a settled part of the chain's history. When a node is confirmed, any funds paid out by the chain to the main Ethereum chain are transferred.
- EthBridge: A group of contracts running on the Ethereum chain, which act as a record-keeper and rule enforcer for Arbitrum chains.
- Inbox: Holds a sequence of messages sent by clients to the contracts on an Arbitrum Chain. Every chain's inbox is managed by the on-chain EthBridge. Every message to a chain is timestamped with the Ethereum block number when it was put into the chain's inbox.
- Lockbox: A contract which holds assets (ETH, ERC20 tokens, or ERC721 tokens) a user has withdrawn from the Arbitrum chain, available to the user after their withdrawal is successfully confirmed.
- Staker: A party who deposits a stake, in Eth, to vouch for a particular node in an Arbitrum Chain. A party who stakes on a false node can expect to lose their stake. An honest staker can recover their stake once the node they are staked on has been confirmed.
- Validator: A party who makes staked, disputable assertions about the state of the Arbitrum chain; i.e., proposing state updates or passively monitoring other validator’s assertions and disputing them if they’re invalid.
- Virtual Machine (VM): A program that "runs" on the Arbitrum chain, which tracks the states of all of the contracts on the chain and all of the Eth and tokens deposited into the chain.
