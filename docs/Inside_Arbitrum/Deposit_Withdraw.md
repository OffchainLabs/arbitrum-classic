---
id: Deposit_Withdraw
title: Inside Arbitrum: Deposits and Withdrawals
sidebar_label: Deposits and Withdrawals
---

## Depositing and withdrawing ETH

The cross-chain call mechanisms described in the Bridging section carry callvalue, which is transported from L1 to L2, or L2 to L1, as part of the cross-chain call. This mechanism can be used to transfer funds between L1 and L2 by doing no-op calls that carry callvalue, in much the same way that Ethereum implements account-to-account transfers via dummy calls that carry callvalue.

Recall that cross-chain calls use a ticket system, where a contract on one chain does a call to create a ticket on one chain that is redeemable on the other chain. In the case of a deposit or withdrawal, the party who is receiving the funds will always be willing and eager to redeem the ticket. The deposit/withdrawal functionality in the Offchain Labs bridging SDK will automatically redeem the tickets.

## Depositing and withdrawing tokens

The Arbitrum team provides a default token bridge, which allows ERC 20, ERC 721, and [others?] tokens to be moved between the L1 and L2 chains. We’ll describe the functionality in terms of ERC 20 tokens, to simplify the explanation. It should be obvious how these mechanisms extend to different token types.

The default token bridge doesn’t have any superpowers--any programmer could create a bridge with the same functionality. We recommend that you don’t, because multiple bridges could confuse users. If you think you need some functionality that the default token bridge doesn’t provide, please come and talk to the Arbitrum team.

The default token bridge uses a buddy contract arrangement, involving a single L1 token bridge contract, and an L2 contract for each token address. 

To deposit a token from L1 to L2, a user calls the deposit method of the L1 token bridge contract, transferring that token to the L1 token bridge. The L1 token bridge holds the token and triggers a transaction at L2 that will cause the L2 contract managing that token type to mint the token at L2 and credit it to the original caller’s account (or another address specified by the original caller). The contract now lives at L2.

Eventually, the token’s owner (whoever that is at the time) may direct that the token be withdrawn back to L1. This causes the L2 contract managing the token type to burn the token at L2 and do an L2-to-L1 call back to the L1 token bridge contract, directing it to release one of the tokens it is holding, and pay that token to the L2 owner’s address back on L1 (or to another address specified by the L2 owner).

The last piece to cover is the addresses used by the default token bridge. The default token bridge contract at L1 lives at a well-known L1 address. An L2 contract is deployed at the same address on the Arbitrum chain, to serve as a “buddy” to the L1 token bridge contract. Both contracts trust calls from their buddy on the other chain, allowing them to coordinate the paired lock/mint operations needed for deposit and the paired burn/unlock operations needed for withdrawal. 

When the L2 token bridge contract receives a deposit from L1, it first checks to see whether there is already an L2 ERC20 token contract instantiated to handle that token type. If not, it creates one. This creation uses the CREATE2 opcode, allowing the token bridge contract to create the token at a predictable address. The address will be a known function of the L1 token address.

Note that this bridging scheme causes a token to use different addresses at L1 and L2. If the L1 address of the token is A, then the L2 address will be F(A) for a known invertible function F. Software at L2 will need to be aware of this address mapping. The L2 token bridge contract provides a method to compute this mapping.

(A previous version of Arbitrum used the identity function as F, meaning that the L2 token contract was put at the same address as the L1 token contract occupied on L1. This was convenient when it worked, but it is not always possible to put the L2 contract at a specified address, even with special support from ArbOS to bend the normal rules of contract address assignment. In particular, it was possible for a contract to already be deployed at the same address, leading to potential trouble. The current scheme avoids that problem, due to the security properties of CREATE2, with the tradeoff that token addresses must be mapped using the F function.)