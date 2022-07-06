---
id: AnyTrust
title: AnyTrust Chains
sidebar_label: AnyTrust Chains
---

AnyTrust chains are an Arbitrum chain type, distinct from Arbitrum Rollup chains. The Arbitrum One mainnet chain (chain ID 42161) is — and will always be — a rollup chain; soon, Offchain Labs will be opening up the first public AnyTrust chain testnet (followed by mainnet!)

The fundamental tradeoff between Rollup and AnyTrust is decentralization vs. transaction costs: Rollup chains inherit their security directly from layer 1 without introducing new trust assumptions, whereas AnyTrust chains introduce their own security assumption, and are thus able to charge users lower fees.

**Basics**

The key changes from Rollup is in where and how the chain's data is stored.

To add transactions to a Rollup chain, the transaction data must be posted on the L1 as calldata. This directly leverages the security assumptions of Ethereum itself to guarantee that the chain's data is available to any party, which, in turn, means anyone can actively validate the chain. This property, combined with the fact and any one honest validator can force correct execution of the chain, means that rollup chains are trustless.

AnyTrust changes this in the following ways:

1. Transaction data is managed off-chain
1. A chain has a Data Availability Committee comprised of a fixed list of named parties
1. Data availability — and in turn correctness of the chain — requires some specified number of those committee members to be honest

For example, there might be 20 members of the committee, and an assumption that at least two of the 20 are honest.

The core advantage of AnyTrust is that in the normal case, transactions are much cheaper than Rollup.

**K-of-N honest**

A useful shorthand is to talk about an AnyTrust chain as making, say, a 2-of-20 honest assumption. That means there is a committee of size 20 and security relies on at least 2 of the committee members being honest. In general, different AnyTrust chains could have different values of K and N.

If K-of-N are honest, then it follows that anything that is vouched for by a “quorum” of N+1-K committee members must be correct. I.e., For 2-of-20, a quorum is any 19 of the 20 members, and anything that 19 members vouch for must be correct. The logic here is that there can’t be more than 18 dishonest committee members, so any quorum of 19 must have an honest party; if a quorum says that something is true, an honest party must be saying that thing is true, so that thing must be true.

The security downside of AnyTrust is that if there is a quorum that is entirely evil, it can destroy the security of the chain. So if a chain assumes that 2-of-20 are honest, but actually only one is honest, then there are 19 malicious committee who can form an evil quorum and steal everything. That’s why we choose K to be a small number like 2.

**Lowering cost using offchain data availability**

In Rollup, we put all transaction data on the Ethereum chain as calldata to ensure that everyone can get the transaction data if they need it. That’s the biggest cost of operating a rollup chain.

In AnyTrust, if a quorum says that they are storing some transaction data and will provide it to others on demand, then we don’t need to put that data on Ethereum, because the AnyTrust honesty assumption is that there is an honest committee member who will provide the data to anyone who needs it.

That’s how AnyTrust lowers the cost of L2 transactions.

**Fallback to rollup**

What if there isn’t a cooperative quorum? Then an AnyTrust chain operates as a regular Rollup. The switch between “quorum mode” and “rollup mode” happens seamlessly, in both directions.

So basically, an AnyTrust chain will behave differently depending on how many committee members are honest and participating:

- **Turbo Mode: 19 out of 20 Honest** If a quorum (e.g., 19 out of 20) is honest and participating, the chain operates in “turbo mode”; data stays off-chain, enabling with low transaction cost.
- **Rollup Mode: 2 out of 20 Honest** If there is not an honest participating quorum, but the K-of-N (e.g. 2-of-20) honesty assumption holds, the chain operates in “rollup mode” with exactly the characteristics of a rollup chain.
- **Failure Mode: 19 out of 20 Evil** If the honesty assumption doesn’t hold and there is an evil quorum, then the chain loses security. No guarantees are possible if this happens.

**How it works**

One of the nicest aspects of AnyTrust is how easy it is to implement once you have a working rollup system; it requires only slight modifications to the Arbitrum Rollup codebase.

Basically, in Rollup, the L1 inbox contract ensures that the hash of some data can be put into the chain’s inbox only if that data is on the L1 chain. AnyTrust changes this rule so a hash can be put in the inbox if the data is on the L1 chain _or_ a quorum has signed a promise to provide the data. That’s all that needs to change in the Rollup protocol. There are also some changes to the validator software but they’re not too complicated.

**Choosing the committee**

Obviously it matters who the members of the Data Availability Committee of an AnyTrust chain are. Unlike a Rollup, where data availability is ensured by construction, an AnyTrust chain needs to explicitly choose the parties responsible for data availability. Choosing the committee is essentially a governance question. At minimum, the identities of the committee members, or the process by which they are elected, should be public information, so that users can gain some degree of assurance that a sufficient portion of them are trustworthy.  
