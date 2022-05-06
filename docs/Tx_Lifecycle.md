---
id: Tx_Lifecycle
title: Transaction Call Lifecycle
sidebar_label: Transaction Lifecycle
---

There are a number of different phases that a transaction goes through before a user can consider it confirmed, starting with guaranteeing transaction ordering and ending with guaranteeing transaction execution. We start with the point at which the user submits the transaction to the sequencer (possibly forwarded through another node).

- **A)** The sequencer has ordered and confirmed the transaction but has not yet posted it to the L1 chain in a batch. At this stage,the transaction can be considered final if the user is willing to trust the sequencer, but a malicious sequencer could violate this finality. In the future we plan on adding a layer of crypto-economic security via bonding and slashing to punish a sequencer who equivocates and violates this trust.
  <br>

- **B)** The sequencer has posted a batch containing the transaction on the L1 chain. At this point, the sequencer has no special power (assuming the batch isn’t reorged out of the L1 chain) and no longer needs to be trusted at all. For users who don't want to rely on trusting the sequencer, their transaction is now as confirmed as the L1 batch transaction which included it.

    <br>
    Once the ordering is guaranteed, the result of the transaction is fully guaranteed assuming that any one honest validator will force the protocol to execute correctly.
    <br>    <br>

- **C)** A validator creates an assertion that asserts the result of the user’s transaction; note that the validator has no power to censor/exclude your transaction (i.e., they are forced to include transactions that are next on the queue) and has no power to reorder. Other validators can also stake on that assertion. At that point if the user trusts at least any one of the validators (or is themself a validator), the transaction can be considered fully trusted.  
   <br>
- **D)** The 7 day challenge period ends and the assertion is confirmed. At this point the result of the transaction is fully locked in on the L1 chain.

---

Most users will only need to wait for confirmation by the sequencer when the outcome is guaranteed (assuming the sequencer doesn't violate its promise). Initially we’ll be running the sequencer, and over time we’ll shift over to a more decentralized sequencer with slashing so the risk here to users is low.

For users that want a higher level of security, they can wait until the sequencer posts a batch, which will [typically](https://arbiscan.io/batches) happen every few minutes. The time between batches just needs to be enough to amortize the relatively low constant costs of posting a batch. It's likely that exchanges supporting direct withdrawals will want to wait for some number of L1 confirmations on a batch transaction before releasing funds, similarly to how they wait for a certain number of confirmations on L1 chains.

Additionally, users always have the option of avoiding reliance on the sequencer altogether by posting transactions directly on-chain themselves (i.e., skipping steps A and B above). In this scenario, their transaction will be queued up for validators to include in an assertion once the sequencer posts its next batch or a protocol enforced time-period elapses (set to 24 hours for mainnet beta), whichever comes first. Users should rarely, if ever, need to go this route to include their transaction, but the recourse to do so ensures that Arbitrum’s censorship resistance properties are ensured even if the sequencer goes offline.

After the transaction ordering is locked in, the result of your transaction is fully deterministic. Assuming there's a single honest validator protecting the correctness of the rollup, you can consider your transaction settled at this point. The only important remaining event to occur is waiting for the assertion to confirm which will allow for the finalization of any withdrawals that were processed in the assertion.
