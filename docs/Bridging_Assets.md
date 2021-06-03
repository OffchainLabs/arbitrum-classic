---
id: Bridging_Assets
title: Token Bridging
sidebar_label: Token Bridging
---

The Arbitrum protocol's ability to [pass messages between L1 and L2](L1_L2_Messages.md) can be leveraged to trustlessly move assets from Ethereum to an Arbitrum chain and back. Any asset / asset type can in principle be bridged, including Ether, ERC20 tokens, ERC-721 tokens, etc.

## Depositing And Withdrawing Ether

To move Ether from Ethereum onto the Arbitrum chain, you execute a deposit transaction via `Inbox.depositEth`. This transfers funds to the Bridge contract on the L1 and credits the same funds to you inside the Arbitrum chain at the specified address.

```sol
function depositEth(address destAddr) external payable override returns (uint256)
```

As far as Ethereum knows, all deposited funds are held by Arbitrum's Bridge contract.

Withdrawing ether can be done using the [ArbSys](Arbsys.md) withdrawEth method:

```sol
ArbSys(100).withdrawEth{ value: 2300000 }(destAddress)
```

Upon withdrawing, the Ether balance is burnt on the Arbitrum side, and will later be made available on the Ethereum side.

`ArbSys.withdrawEth` is actually a convenience function is which is equivalent to calling `ArbSys.sendTxToL1` with empty calldataForL1. Like any other `sendTxToL1` call, it will require an additional call to `Outbox.executeTransaction` on L1 after the dispute period elapses for the user to finalize claiming their funds on L1 (see ["L2 to L1 Messages Lifecycle && API"](L1_L2_Messages.md)). Once the withdrawal is executed from the Outbox, the user's Ether balance will be credited on L1.

## Bridging ERC20 Tokens

### Overview

The Arbitrum protocol itself technically has no native notion of any token standards, and gives no built-in advantage or special recognition to any particular token bridge. Described here is the "Canonical Bridge," which we at Offchain Labs implemented, and which should be the primary bridge most users and applications use; it is (effectively) a DApp with contracts on both Ethereum and Arbitrum that leverages Arbitrum's cross-chain message passing system to achieve basic desired token-bridging functionality. We recommend that you use it!

### Design Rationale

_Our design and thinking has been influenced by many in the Ethereum community, including [this proposal](https://ethereum-magicians.org/t/outlining-a-standard-interface-for-cross-domain-erc20-transfers/6151) from Maurelian & Ben Jones,
work with David Mihal, and feedback from many projects building on Arbitrum, too numerous to mention!_

_We use the term "Gateway" (a contract for facilitating cross-domain transfers) as per the proposal linked above_

Three core goals motivate the design of our bridging system:

####1. Custom "Gateway" functionality

For many ERC20 tokens, standard bridging functionality is sufficient, which entails the following: a token contract on Ethereum is associated with a "paired" token contract on Arbitrum. Depositing a token entails escrowing some amount of the token in an L1 bridge contract, and minting the same amount at the paired token contract on L2. On L2, the paired contract behaves much like a normal ERC20 token contract. Withdrawing entails burning some amount of the token in the L2 contract, which then can later be claimed from the L1 bridge contract.

Many tokens, however, require custom Gateway systems which are hard to generalize. E.g.,

- Tokens which accrue interest to their holders need to ensure that interest is dispersed properly across layers, and doesn't simply accrue to the bridge contracts
- Our cross-domain WETH implementations requires tokens be wrapped and unwrappeded as they move accross layers.
- Etc.

Thus, our bridge architecture must allow for new, custom Gateways to be dynamically added over time.

####2. Canonical L2 Representation Per L1 Token Contract
...having multiple custom Gateways is well and good, but we also want to avoid a situation in which a single L1 token that uses our bridging system can be represented at multiple addresses/contracts on L2 (as this adds significant friction and confusion for users and developers). Thus, we need a way to track which L1 token uses which gateway, and in turn, to have a canonical address oracle that maps the tokens addresses across the Ethereum and Arbitrum domains.

####3. Domain Agnostic
[This post](<(https://ethereum-magicians.org/t/outlining-a-standard-interface-for-cross-domain-erc20-transfers/6151)>) convinced us thinking about this early on is important; while here we are focused on bridging assets between Ethereum L1 and a single Arbitrum chain, we expect that overtime, Gateways will be developed to transfer assets between any combination of Rollups, Shards, and other L1s. Thus, we follow domain-neutral semantics like "outBoundTransfer" over things like "deposit" and "withdraw", and ensure that common interfaces are sufficiently extensible to support custom (i.e., domain-specific) functionality.

### Canonical Token Bridge Implementation

With that, we can outline our Token Bridge implementation.

On Ethereum L1 resides a single `GatewayRouter` contract; this contract is responsible for mapping the addresses of tokens on L1 to its corresponding L1 `TokenGateway` contract; a token can opt-in to one and only one `TokenGateway`.

Each L1 `TokenGateway` corresponds to exactly one L2 `TokenGateway`; this pair of contracts together handles the cross-domain asset transfers.

Important to note is that GatewayRouter itself conforms to the `TokenGateway` interface. Crucially, _all L1 to L2 messages are initiated via the GatewayRouter_. I.e., the gateway router forwards a token's Outgoing Message to its corresponding L1Gateway.

#### Standard Arb-ERC20 Bridging

To help illustrate what this looks like in practice, let's go through the steps of what depositing and withdrawing `SomeERC20Token` via our Standard ERC20 gateway looks like. Here, we're assuming that `SomeERC20Token` has already opted in to the Standard ERC20 Gateway.

#### Deposits

1. A user calls `GatewayRouter.outBoundTransfer` with `SomeERC20Token`'s L1 address as a parameter.
2. `GatewayRouter` looks up `SomeERC20Token`'s canonical gateway, and finds that it's the Standard ERC20 gateway (the `L1ERC20Gateway` contract).
3. `GatewayRouter` calls `L1ERC20Gateway.outBoundTransfer`, forwarding the appropriate parameters.
4. `L1ERC20Gateway` escrows tokens, and creates a retryable ticket to trigger `L2ERC20Gateway`'s `finalizeInboundTransfer` method on L2.
5. `finalizeInboundTransfer` mints the appropriate amount of tokens at the `arbSomeERC20Token` contract on L2.

#### Withdrawals

1. On Arbitrum, a user calls `L2ERC20Gateway.outBoundTransfer`.
2. This burns tokens on L2, and calls ArbSys with an encoded message to `L1ERC20Gateway.finalizeInboundTransfer`, which will be eventually executed on L1.
3. After the dispute window expires and the assertion with the user's transaction is confirmed, a user can call `Outbox.executeTransaction`, which in turn calls the encoded `L1ERC20Gateway.finalizeInboundTransfer` message, releasing the user's tokens from escrow.

Note that in the system described above, one pair of Gateway contracts handles the bridging of many ERC20s; other custom Gateways may well be made to connect a single L1 contract to an L2 counterpart (i.e., the Arb-WETH bridge alluded to earlier). Ultimately, the Gateway interface is flexible enough to support many different use-cases.

## Arbitrum-Native ERC20 Tokens

It is (of course) possible to deploy an ERC20 token contract directly to Arbitrum, i.e., with no layer 1 counterpart. Such a token functions normally within Arbitrum, but simply can't be withdrawn to layer 1. (In principle, enabling L2 native tokens to be withdrawn to an L1 contract could be possible via a similar mechanism our Canonical Token Bridge uses with the layers flipped, i.e., an "anti-bridge," but such functionality isn't currently supported.)
