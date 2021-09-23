---
id: L1_L2_Messages
title: Messaging Between Layers
sidebar_label: Messaging Between Layers
---

## Standard Arbitrum Transactions: Calls from clients

Standard, client-generated transaction calls on the Arbitrum chain are sent through the EthBridge using Inbox.sendL2Message:

```solidity
function sendL2Message(address chain, bytes calldata messageData) external;
```

Generally calls will come in batches from an aggregator as described in [Transaction Lifecycle](Tx_Lifecycle.md).

However, the Arbitrum protocol also offers ways of passing messages between the layer 1 and layer 2 chains.

The most common use-case for direct inter-chain communication is depositing and withdrawing assets; this, however, is only one specific application of generalized cross-chain contract calls that Arbitrum supports. This page covers the generalized protocol; for further explanation, see [Inside Arbitrum: Bridging](Inside_Arbitrum.md#bridging).

## Ethereum to Arbitrum: Retryable Tickets

Retryable tickets are the Arbitrum protocol’s canonical method for passing generalized messages from Ethereum to Arbitrum. A retryable ticket is an L2 message encoded and delivered by L1; if gas is provided, it will be executed immediately. If no gas is provided or the execution reverts, it will be placed in the L2 retry buffer, where any user can re-execute for some fixed period (roughly one week).

### Motivation

Retryable tickets are designed to gracefully handle various potentially tricky aspects of cross-chain messaging:

- **Overpaying for L2 Gas:** An L1 contract has to supply gas for the L2 transaction’s execution; if the L1 side overpays, this begets the questions of what to do with this excess Ether.

- **L2 transaction reversion — breaking atomicity:** It is vital for many use-cases of L1 to L2 transactions that the state updates on both layers are atomic, i.e., if the L1 side succeeds, there is assurance that the L2 will eventually succeed as well. The canonical example here is a token deposit: on L1, tokens are escrowed in some contract, and a message is sent to L2 to mint some corresponding tokens. If the L1 side succeeds and the L2 reverts, the user has simply lost their tokens (i.e., donated them to the bridge contract).

- **L2 transaction reversion — handling L2 callvalue:** The L1 side must supply the Ether for the callvalue of the L2 transaction (by depositing it); if the L2 transaction reverts, this begets the questions of what to do with this in-limbo Ether.

Retryable tickets handle all these things (and handle them well!)

### Transaction Types / Terminology

| Txn Type           | Description                                                                                                                                           | Appearance                                                                                                                     | Tx ID                                                                  |
| ------------------ | ----------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------ | ---------------------------------------------------------------------- |
| Retryable Ticket   | Quasi-transaction that sits in the retry buffer and has a lifetime over which it can be executed, i.e., “redeemed.”                                   | Emitted when message is included; will succeed if user supplies sufficient ETH to cover base-fee + callvalue, otherwise fails. | _keccak256(zeroPad(l2ChainId), zeroPad(bitFlipedinboxSequenceNumber))_ |
| Redemption Txn     | Transaction that results a retryable ticket being successfully redeemed; looks like a normal L2 transaction.                                          | Emitted after a retryable ticket is successfully redeemed, either user-initiated or via an auto-redeem.                        | _keccak256(zeroPad(retryable-ticket-id), 0)_                           |
| Auto-Redeem Record | Quasi-transaction ArbOS creates automatically which attempts to redeem a retryable ticket immediately when it is submitted using the ArbGas provided. | Attempted / emitted iff gas\*gas-price > 0. If it fails, retryable ticket stays in the retry buffer.                           | \_keccak256(zeroPad(retryable-ticket-id), 1)                           |

### Retryable Tickets Contract API

A retryable ticket is created by calling [Inbox.createRetryableTicket](./sol_contract_docs/md_docs/arb-bridge-eth/bridge/Inbox.md). Other operations involving retryable tickets are preformed via the [ArbRetryableTicket](./sol_contract_docs/md_docs/arb-os/arbos/builtin/ArbRetryableTx.md) precompile interface at L2 address 0x000000000000000000000000000000000000006E. (This contract won't need to be used under normal circumstances.)

### Parameters:

There are a total of 10 parameters that the L1 must pass to the L2 when creating a retryable ticket.

5 of them have to do with allocating ETH/Gas:

- **DepositValue:** Total ETH deposited from L1 to L2.
- **CallValue:** Call-value for L2 transaction.
- **GasPrice:** L2 Gas price bid for immediate L2 execution attempt (queryable via standard eth\*gasPrice RPC)
- **MaxGas:** Gas limit for immediate L2 execution attempt (can be estimated via \_NodeInterface.estimateRetryableTicket\*)
- **MaxSubmissionCost:** Amount of ETH allocated to pay for the base submission fee. The base submission fee is a parameter unique to retryable transactions; the user is charged the base submission fee to cover the storage costs of keeping their ticket’s calldata in the retry buffer. (current base submission fee is queryable via `ArbRetryableTx.getSubmissionPrice`)

Intuitively: if a user does not desire immediate redemption, they should provide a DepositValue of at least `CallValue + MaxSubmissionCost`. If they do desire immediate execution, they should provide a DepositValue of at least
`CallValue + MaxSubmissionCost + (GasPrice x MaxGas).`

### Other Parameters

- **Destination Address:** Address from which transaction will be initiated on L2.
- **Credit-Back Address:** Address to which all excess gas is credited on L2; i.e., excess ETH for base submission cost (`MaxSubmissionCost - ActualSubmissionCostPaid`) and excess ETH provided for L2 execution (` (GasPrice x MaxGas) - ActualETHSpentInExecution`).
- **Beneficiary:** Address to which CallValue will be credited to on L2 if the retryable ticket times out or is cancelled. The Beneficiary is also the address with the right to cancel a Retryable Ticket (if the ticket hasn’t been redeemed yet).
  `Calldata:` data encoding the L2 contract call.
  `Calldata Size:` CallData size.

### Important Note About Base Submission Fee

If an L1 transaction underpays for a retryable ticket's base submission free, the retryable ticket creation on L2 simply fails. Given that this potentially breaks the atomicity of the L1 / L2 transactions, applications should avoid this scenario. The current base submission fee returned by `ArbRetryableTx.getSubmissionPrice` increases once every 24 period by at most 50% of its current value. Since any amount overpaid will be credited to the `credit-back-address`, it is highly recommended that applications judiciously overpay relative to the current price.

In a future release, the base submission fee will be calculated using the 1559 `BASE_FEE` and collected directly at L1; underpayment will simply result in the L1 transaction reverting, thus avoiding the complications above entirely.

### Retryable Transaction Lifecycle:

When a retryable ticket is initiated from the L1, the following things take place:

- DepositValue is credited to the sender’s account on L2.
  - If DepositValue is less than MaxSubmissionCost + Callvalue, the Retryable Ticket fails.
- Submission fee is collected: submission fee is deducted from the sender’s L2 account; MaxSubmissionCost - submission fee is credited to Credit-Back Address.

- Callvalue is deducted from sender’s L2 account and a Retryable Ticket is successfully created.
  - If the sender has insufficient funds to cover the callvalue (even after the DepositValue has been credited), the Retryable Ticket fails.
- If MaxGas and MaxPrice are both > 0, an immediate redemption will be attempted:

- MaxGas x MaxPrice is credited to the Credit-Back Address.
- The retryable ticket is automatically executed —i.e., the transaction encoded in Calldata with Callvalue — with gas provided by the Credit-Back Address.
- If it succeeds, a successful Immediate-Redeem Txn is emitted along with a successful Redemption Transaction. Any excess gas remains in the Credit-Back Address.
- If it reverts, a failed Immediate-Redeem Txn is emitted, and the Retryable Ticket is placed in the retry-buffer.
  ...Otherwise, the Retryable Ticket goes straight into the retry buffer, and no Immediate-Redeem Txn is ever emitted.

Any user can redeem a Retryable Ticket in the retry buffer (before it expires) by calling ArbRetryableTicket.redeem(Redemption-TxID). The user provides gas the way they would a normal L2 transaction.

If the Retryable Ticket is cancelled or expires before it is redeemed, Callvalue is credited to Beneficiary.

### Depositing ETH via Retryables

Currently, the canonical method for depositing ETH into Arbitrum is to create a retryable ticket using the [Inbox.depositEth](./sol_contract_docs/md_docs/arb-bridge-eth/bridge/Inbox.md) method. A Retryable Ticket is created with 0 Callvalue, 0 MaxGas, 0 GasPrice, and empty Calldata. The DepositValue credited to the sender’s account in step 1 simply remains there.

The Retryable Ticket gets put in the retry buffer and can in theory be redeemed, but redeeming is a no-op.

Beyond the superfluous ticket creation, this is suboptimal in that the base submission fee is deducted from the amount deposited, so the user will see a (slight) discrepancy between the amount sent to be deposited and ultimate amount credited in their L2 address. A special message type Taylor-made for ETH deposits that handles them more cleanly will be exposed soon.

### Address Aliasing

When a retryable ticket is executed on L2, the sender's address —i.e., that which is returned by `msg.sender` — will _not_ simply be the address of the contract on L1 that initiated the message; rather it will be the contract's "L2 Alias." A contract address's L2 alias is its value increased by the hex value `0x1111000000000000000000000000000000001111`:

```
L2_Alias = L1_Contract_ Address + 0x1111000000000000000000000000000000001111
```

The Arbitrum protocol's usage of L2 Aliases for L1-to-L2 messages prevents cross-chain exploits that would otherwise be possible if we simply reused L1 contact addresses.

If for some reason you need to compute the L1 address from an L2 alias on chain, you can use our `AddressAliasHelper` library:

```sol
    modifier onlyFromMyL1Contract() override {
        require(AddressAliasHelper.undoL1ToL2Alias(msg.sender) === myL1ContractAddress, "ONLY_COUNTERPART_CONTRACT");
        _;
    }
```

_Of note: the (highly) recommended convenience methods `Inbox.depositEth` and `Inbox.createRetryableTicket` handle subtracting the offset for the credit back address and beneficiary address, and using the these reverse-aliased values as the inputs for the create retryable tickets. This way, the beneficiary and credit back addresses as received on L2 will be equivalent to those provided as the L1 params._

#### Directly Redeeming / Cancelling Retryables

In the normal, happy case, a retryable ticket is automatically redeemed and executed when it arrives on L2; in this case only a single user-action (publishing the transaction on L1 that creates a retryable ticket) is required. If the transaction is not auto-redeemed (either because no L2 gas was provided or because the L2 retryable ticket reverts for whatever reason), the transaction can now either be redeemed via a signed L2 transaction (probably desireable) or cancelled.

To redeem the retryable, any account can call `ArbRetryableTx.redeem(redemption-txn-id)`. To cancel, the beneficiary address can call `ArbRetryableTx.cancel(redemption-txn-id)`.

See the [ArbRetryableTx](./sol_contract_docs/md_docs/arb-os/arbos/builtin/ArbRetryableTx.md) interface for other related methods, and [arb-ts](https://arb-ts-docs.netlify.app/) for convenience methods around using retryables.

### Demo

See [Greeter](https://github.com/OffchainLabs/arbitrum-tutorials/tree/master/packages/greeter) for a simple demo illustrating the use of retryable tickets.

## Arbitrum to Ethereum

### Explanation

L2 to L1 messages work similar to L1 to L2 messages, but in reverse: an L2 transaction is published with an L1 message as encoded data, to be executed later.

A key difference, however, is that in the L2 to L1 direction, a user must wait for the dispute period to pass between publishing their messages and actually executing it on L1; this is a direct consequence of the security model of Optimistic Rollups (see [finalty](Finality.md).) Additionally, unlike retyable tickets, outgoing messages have no upper bounded timeout; once the dispute window passes, they can be executed at any point. No rush.

### L2 to L1 Messages Lifecycle

The lifecycle of sending a message from layer 2 to layer 1 can be broken down into roughly 4 steps, only 2 which (at most!) require the end user to publish transactions.

**1. Publish L2 to L1 transaction (Arbitrum transaction)**

A client initiates the process by publishing a message on L2 via `ArbSys.sendTxToL1`.

**2. Outbox entry gets created**

After the Arbitrum chain advances some set amount of time, ArbOS gathers all outgoing messages, Merklizes them, and publishes the root as an OutboxEntry in the chain's outbox. Note that this happens "automatically"; i.e., it requires no additional action from the user.

**3. Client gets Merkle proof of outgoing message**

After the Outbox entry is published on the L1 chain, the user (or anybody) can compute the Merkle proof of inclusion of their outgoing message. They do this by calling `NodeInterface.lookupMessageBatchProof`:

```sol

/** @title Interface for providing Outbox proof data
 *  @notice This contract doesn't exist on-chain. Instead it is a virtual interface accessible at 0x00000000000000000000000000000000000000C8
 * This is a cute trick to allow an Arbitrum node to provide data without us having to implement an additional RPC )
 */

interface NodeInterface {
    /**
    * @notice Returns the proof necessary to redeem a message
    * @param batchNum index of outbox entry (i.e., outgoing messages Merkle root) in array of outbox entries
    * @param index index of outgoing message in outbox entry
    * @return (
        * proof: Merkle proof of message inclusion in outbox entry
        * path: Index of message in outbox entry
        * l2Sender: sender if original message (i.e., caller of ArbSys.sendTxToL1)
        * l1Dest: destination address for L1 contract call
        * l2Block l2 block number at which sendTxToL1 call was made
        * l1Block l1 block number at which sendTxToL1 call was made
        * timestamp l2 Timestamp at which sendTxToL1 call was made
        * amouunt value in L1 message in wei
        * calldataForL1 abi-encoded L1 message data
        *
    */
    function lookupMessageBatchProof(uint256 batchNum, uint64 index)
        external
        view
        returns (
            bytes32[] memory proof,
            uint256 path,
            address l2Sender,
            address l1Dest,
            uint256 l2Block,
            uint256 l1Block,
            uint256 timestamp,
            uint256 amount,
            bytes memory calldataForL1
        );
}

```

**4. The user executes the L1 message (Ethereum Transaction)**

Anytime after the dispute window passes, any user can execute the L1 message by calling Outbox.executeTransaction; if it reverts, it can be re-executed any number of times and with no upper time-bound:

```sol
 /**
    * @notice Executes a messages in an Outbox entry. Reverts if dispute period hasn't expired and
    * @param outboxIndex Index of OutboxEntry in outboxes array
    * @param proof Merkle proof of message inclusion in outbox entry
    * @param index Index of message in outbox entry
    * @param l2Sender sender if original message (i.e., caller of ArbSys.sendTxToL1)
    * @param destAddr destination address for L1 contract call
    * @param l2Block l2 block number at which sendTxToL1 call was made
    * @param l1Block l1 block number at which sendTxToL1 call was made
    * @param l2Timestamp l2 Timestamp at which sendTxToL1 call was made
    * @param amount value in L1 message in wei
    * @param calldataForL1 abi-encoded L1 message data
     */
    function executeTransaction(
        uint256 outboxIndex,
        bytes32[] calldata proof,
        uint256 index,
        address l2Sender,
        address destAddr,
        uint256 l2Block,
        uint256 l1Block,
        uint256 l2Timestamp,
        uint256 amount,
        bytes calldata calldataForL1
    )
```

### Demo

See [outbox-execute](https://github.com/OffchainLabs/arbitrum-tutorials/tree/master/packages/outbox-execute) for client side interactions with the outbox using the [arb-ts](https://arb-ts-docs.netlify.app/) library.

See also [integration tests](https://github.com/OffchainLabs/arbitrum/tree/master/packages/arb-ts/integration_test) and our [Token Bridge UI](https://github.com/OffchainLabs/arb-token-bridge).
