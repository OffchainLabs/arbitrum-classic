---
id: Security_Considerations
title: Security Considerations & Gotchas
sidebar_label: Security Considerations
---

Ethereum dApp developers looking to build on Arbitrum will generally find the experience, tooling, and protocol rationale to be very familiar; that said, there are a number of important considerations and "gotchas" that developers should be aware of. For many smart contract applications, none of the things below will apply, but devs are advised to do due diligence:

- **Different EVM/Soldity behavior**: Arbitrum supports every EVM opcode, which means any Solidity (or Vyper, Yuul, etc.) you write will compile and run on Arbitrum. Certain operations, however, have different behavior on Arbitrum than on Ethereum; for a full list, see [Solidity Support](Solidity_Support.md). Also Arbitrum supports most — but not all — Ethereum [precompiles.](Differences_Overview.md)

- **Timing Assumptions with block numbers and timestamps**: Timing assumptions one may make on layer 1 about `block.timestamp` and `block.number` (i.e., that a new block will be produced every 15 seconds on average) won't hold on Arbitrum. The the rate at which new "L2 blocks" are produced / delineated isn't reliably predictable; `block.timestamp` is a better way to measure time on Arbitrum, but even so, relying on either `block.timestamp` or `block.number` on L2 to measure time should only be done over large time scales. For more, see [Block Numbers and Time](Time_in_Arbitrum.md).

- **L1-to-L2 Transactions' Address Aliases**: "Retryable tickets" are a special transaction type for sending a message from L1 to L2; if you plan on using them, it is highly recommended you read through their [documentation](L1_L2_Messages.md) and carefully test your contracts' behavior before deploying on mainnet. Of particular note: in L2 message on the receiving side, `msg.sender` not return the L1 contract, but rather it's [address alias](L1_L2_Messages.md#address-aliasing).

- **L1-to-L2 Tranasaction's Ticket Creation Failure**: If you underpay for the base submission cost when attempting to create a retryable ticket, the ticket on L2 will simply not get created despite the L1 transaction being confirmed; this could — depending on what your cross-chain app is doing — be very bad. The current base submission cost can be queried directly from an Arbitrum node, and it updates once every 24 hour period; for a given update, it can increase by at most 50% of its previous value. Any amount you overpay will not be lost; it will be collected at your specified address on L2. To be safe, we highly recommend you overpay. (In a future release, the base submission cost will collected directly on L1, avoiding failure mode described above entirely.)

  If you're utilizing retryable tickets in your dApp and anything in the above paragraph didn't make sense to you (or even if it did), then oh boy should you _really_ read the [documentation](L1_L2_Messages.md).

- **Hard Coded Gas Values**: [ArbGas](ArbGas.md) is denominated differently than Ethereum L1 gas; thus, a contract with a hard coded value that works on L1 may break if it is deployed unmodified to L2; hard coded gas values should be adjusted (or better yet, removed entirely if possible (it's probably possible; really now, why are you hard-coding a gas value?))

- **ETH Deposit Special Behavior**: Retryable tickets are leveraged in a special way to handle deposits of Ether from L1 onto Arbitrum; if your application will be using Ether deposits directly, it's worth understanding the [details of their design](L1_L2_Messages). Additionally, the ability to credit Ether to an L2 account without publishing a signed transaction on L2 yields some behavior not possible on L1, namely, the ability to credit a contract with Ether without triggering its [receive fallback function](https://docs.soliditylang.org/en/v0.6.2/contracts.html#receive-ether-function).

- **Block Hashes For Randomness**: Arbitrum's L2 block hashes should not be relied on as a secure source of randomness.

If you have any questions, don't hesitate to engage with us and our community on our [Discord](https://discord.gg/ZpZuw7p)
