---
id: Glossary
title: Glossary of Arbitrum Terms
custom_edit_url: https://github.com/OffchainLabs/arbitrum/edit/master/docs/Glossary.md
---

### General

- **Any-Trust Model**: Security model that only requires that any one participating party behave honestly. Any trust model + permissionless participation = trustless (i.e., Arbitrum Rollup).

- **Arbitrum Chain**: A chain running on Ethereum, containing some contracts. Many Arbitrum chains can exist simultaneously.

- **Arbitrum Rollup**: Trustless Arbitrum L2 protocol in which participation is permissionless and underlying layer is used for data availability (e.g., Arbitrum One).

- **Arbitrum Channels**: Trustless L2 protocol in which participation is permissioned, data is kept off chain, and participants update the state via unanimous consent. (Not yet in production).

- **Arbitrum One**: The first Arbitrum chain running on Ethereum mainnet! (Currently in Beta).

- **Arbitrum Sidechains**: Arbitrum protocol in which data is kept off chain and an any-trusted committee is responsible for updating the chain's state via unanimous consent.

- **ArbGas**: Unit for pricing execution on Arbitrum; denominated in the chain's native currency (Ether on Arbitrum One). ArbGas is somewhat analogous to gas on L1 Ethereum, though different factors go into its calculation. ( See [ArbGas and Fees](Inside_Arbitrum.md#arbgas-and-fees).)

- **Arbitrum Full Node**: A party who keeps track of the state of an Arbitrum chain and receives remote procedure calls (RPCs) from clients. Analogous to a non-mining L1 Ethereum node.

- **ArbOS**: Layer 2 "operating system" that trustlessly handles system-level operations; includes the ability to emulate the EVM.

- **AVM**: The Arbitrum Virtual Machine; the environment in which execution takes place in an Arbitrum chain. Tracks the states of all of the contracts on the chain and all of the Eth and tokens deposited into the chain.

- **Chain state**: A particular point in the history of an Arbitrum Chain. A chain state corresponds to a sequence of assertions that have been made, and a verdict about whether each of those assertions was accepted.

- **Client**: A program running on a user's machine, often in the user's browser, that interacts with contracts on an Arbitrum chain and provides a user interface.

- **EthBridge**: A group of contracts running on the Ethereum chain, which act as a record-keeper and rule enforcer for Arbitrum chains.

- **Rollup Protocol**: Protocol for tracking the tree of assertions in an Arbitrum chain their confirmation status.

- **Speed Limit**: Target computation limit for an Arbitrum chain; currently, on Arbitrum One targets 80,000 ArbGas per second.

## Proving Fraud

- **Assertion**: A claim, made by a validator, about what the contracts on a chain will do. An assertion is considered _pending_ until it is confirmed.

- **Challenge**: When two stakers disagree about the correct verdict on an assertion, those stakers can be put in a challenge. The challenge is refereed by the EthBridge. Eventually one staker wins the challenge. The loser forfeits their stake. Half of the loser's stake is given to the winner, and the other half is burned.

- **Confirmation**: The final decision by an Arbitrum chain to accept a node as being a settled part of the chain's history. Once an assertion is confirmed, any L2 to L1 messages (i.e., withdrawals) can be executed.

- **Challenge Period**: Window of time (1 week on Arbitrum One) over which an assertion can be challenged, and after which an assertion can be confirmed.

- **Dissection**: Process by which two challenging parties interactively narrow down their disagreement to a single computational step.

- **One Step Proof**: Final step in a challenge; a single operation of the L2 VM is executed, and the validity of its state transition is verified.

- **Staker**: A party who deposits a stake, in Eth, to vouch for a particular node in an Arbitrum Chain. A party who stakes on a false node can expect to lose their stake. An honest staker can recover their stake once the node they are staked on has been confirmed.

- **Active Validator**: A party who makes staked, disputable assertions about the state of the Arbitrum chain; i.e., proposing state updates or challenging the validity of assertions. (Not to be confused with the Sequencer)

- **Defensive Validator**: A validator that watches the Arbitrum chain and takes action (i.e., stake and challenges) only when and if an invalid assertion occurs.

- **Watchtower Validator**: A validator that never stakes / never takes on chain action, who raises the alarm (by whatever off-chain means it chooses) if it witnesses an invalid assertion.

## Cross Chain Communication

- **Address Alias**: A deterministically generated address to be used on L2 that corresponds to an address on L1 for the purpose of L1 to L2 cross-chain messaging.

- **Fast Exit / Liquidity Exit**: A means by which a user can bypass Arbitrum's challenge period when withdrawing fungible assets (or more generally, executing some "fungible" L2 to L1 operation); a liquidity provider facilitates an atomic swap of the asset on L2 directly to L1.

- **Outbox**: An L1 contract responsible for tracking outgoing (Arbitrum to Ethereum) messages, including withdrawals, which can be executed by users once they are confirmed.

- **Outbox Entry**: A Merkle root of a series of outgoing messages posted over some period of time, stored in the Outbox.

- **Retryable Ticket**: An L1 to L2 cross chain message created by an L1 contract sent to an Arbitrum chain for execution (e.g., a token deposit).

- **Retryable Autoredeem**: The "automatic" execution of a retryable ticket on Arbitrum (using provided ArbGas).

## Token Bridging

- **Arb Token Bridge**: A series of contracts on Ethereum and Arbitrum for trustlessly moving tokens between the L1 and L2.

- **Token Gateway**: A pair of contracts in the token bridge — one on L1, one on L2 — that provide a particular mechanism for handling the transfer of tokens between layers. Token gateways currently active in the bridge are the StandardERC20 Gateway, the CustomERC20 Gateway, and the WETH Gateway.

- **Gateway Router**: Contracts in the token bridge responsible for mapping tokens to their appropriate gateways.

- **Standard Arb-Token**: An L2 token contract deployed via the StandardERC20 gateway; offers basic ERC20 functionality in addition to deposit / withdrawal affordances.

- **Custom Arb-Token**: Any L2 token contract registered to the Arb Token Bridge that isn't a standard arb-token (i.e., a token that uses any gateway other than the StandardERC20 Gateway).

## Transaction Ordering

- **Batch**: A group of L2 transactions posted in a single L1 transaction by the Sequencer.

- **Fair Ordering Algorithm**: BFT algorithm in which a committee comes to consensus on transaction ordering; current single-party Sequencer on Arbitrum one will eventually be replaced by a fair-ordering committee.

- **Forced-Inclusion**: Censorship resistant path for including a message into L2; bypasses any Sequencer involvement.

- **Sequencer**: An entity (currently a single-party on Arbitrum One) given rights to reorder transactions in the Inbox over a small window of time, who can thus give clients sub-blocktime soft confirmations. (not a validator)

- **Sequencer Chain**: An Arbitrum chain that includes Sequencer support (e.g., Arbitrum One).

- **Soft Confirmation**: A semi-trusted promise from the Sequencer to post a user's transaction in the near future; soft-confirmations happen prior to posting on L1, and thus can be given near-instantaneously (i.e., faster than L1 block times)

- **Slow Inbox**: Sequence of L1 initiated message that offer an alternative path for inclusion into the fast Inbox.

- **Fast Inbox**: Contract that holds a sequence of messages sent by clients to the contracts on an Arbitrum Chain; message can be put into the Inbox directly by the Sequencer or indirectly through the slow inbox.
