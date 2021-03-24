---
id: EthBridge
title: Inside Arbitrum: The EthBridge
sidebar_label: The EthBridge
---
The EthBridge is a set of Ethereum contracts that manage an Arbitrum chain. The EthBridge keeps track of the chain’s inbox contents, the hash of the chain’s state, and information about the outputs. The EthBridge is the ultimate source of authority about what is going on in the Arbitrum chain. 

The EthBridge is the foundation that Arbitrum’s security is built on. The EthBridge runs on Ethereum, so it is transparent and executes trustlessly. 

The *Inbox* contract manages the chain’s inbox. Inbox keeps track of the (hash of) every message in the inbox. Calling one of the send* methods of Inbox will insert a message into the Arbitrum chain’s inbox. 

The Inbox contract makes sure that certain information in incoming messages is accurate: that the sender is correctly recorded, and that the Ethereum block number and timestamp are correctly recorded in the message.

Unsurprisingly, there is also an *Outbox* contract, which manages outputs of the chain. When a rollup block is confirmed, the outputs produced in that rollup block are put into the outbox. How outputs end up being reflected on Ethereum is detailed in the [Bridging](Bridging.md) section.

The Rollup contract and its friends are responsible for managing the rollup protocol. They track the state of the Arbitrum chain: the rollup blocks that have been proposed, accepted, and/or rejected, and who has staked on which rollup nodes. The Challenge contract and its friends are responsible for tracking and resolving any disputes between validators about which rollup blocks are correct. The functionality of Rollup, Challenge, and their friends will be detailed below in the Rollup Protocol section.