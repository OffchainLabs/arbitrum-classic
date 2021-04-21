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

However, the Arbitrum protocol also offers ways passing of messages between the layer 1 and layer 2 chains.

The most common use-case for direct inter-chain communication is depositing and withdrawing assets; this, however, is only one specific application of generalized cross-chain contract calls that Arbitrum supports. This page covers the generalized protocol; for further explanation, see [Inside Arbitrum: Bridging](Inside_Arbitrum.md#bridging).

## Ethereum to Arbitrum: Retryable Tickets

#### Explanation

Arbitrum offers several ways for an Ethereum transaction to send a message to Arbitrum ([see "L2 Messages"](Data_Formats.md)); however, the generally recommended method to use for direct L1 to L2 communication is via retryable tickets.

The idea is the following: a layer 1 transaction is put in the Inbox with instructions to submit a transaction to L2 (including calldata, callvalue, and gas info) in such a way that if it doesn't execute successfully, it gets put into an L2 "retry buffer." This means that for a period of time (likely on the scale of the chain's dispute window, so roughly one week), anybody can attempt to "redeem" the the L2 transaction ticket by re-executing it.

The rationale here is to account for cases like the following: say we want a transaction that lets a user deposit a token onto Arbitrum; this will entail escrowing some tokens in a contract on L1, and sending a message to mint the same amount of tokens on L2. Suppose that the L1 transaction succeeds but the L2 message reverts due to insufficient gas. In a naive implementation, this would be a serious problem — the user has simply transferred tokens a contract and received nothing on L2; those tokens are stuck in the contract indefinitely. With retryable tickets, however, the user (or some other benevolent bystander) has a week to simply re-execute the L2 message with sufficient gas.

#### Retryable Tickets API

A convenience method for creating retryable tickets is exposed in `Inbox.createRetryableTicket`:

```sol
    /**
    @notice Put an message in the L2 inbox that can be re-executed for some fixed amount of time if it reverts
    * @dev all msg.value will deposited to callValueRefundAddress on L2
    * @param destAddr destination L2 contract address
    * @param l2CallValue call value for retryable L2 message
    * @param  maxSubmissionCost Max gas deducted from user's L2 balance to cover base submission fee
    * @param excessFeeRefundAddress maxGas x gasprice - execution cost gets credited here on L2 balance
    * @param callValueRefundAddress l2Callvalue gets credited here on L2 if retryable txn times out or gets cancelled
    * @param maxGas Max gas deducted from user's L2 balance to cover L2 execution
    * @param gasPriceBid price bid for L2 execution
    * @param data ABI encoded data of L2 message
    * @return unique id for retryable transaction (keccak256(requestID, uint(0) )
     */
    function createRetryableTicket(
        address destAddr,
        uint256 l2CallValue,
        uint256 maxSubmissionCost,
        address excessFeeRefundAddress,
        address callValueRefundAddress,
        uint256 maxGas,
        uint256 gasPriceBid,
        bytes calldata data
    ) external payable override returns (uint256)
```

See [EthErc20Bridge.depositToken](https://github.com/OffchainLabs/arbitrum/blob/5bd9a456a780582715a62affd887d35e2eb138b0/packages/arb-bridge-peripherals/contracts/tokenbridge/ethereum/EthERC20Bridge.sol#L267) for example usage.

Additionally, a precompiled `ArbRetryableTx` contract exists in every Arbitrum chain at address `0x000000000000000000000000000000000000006E`, which exposes methods relevant to retryable transactions:

```sol

pragma solidity >=0.4.21 <0.7.0;

/**
* @title precompiled contract in every Arbitrum chain for retryable transaction related data retrieval and interactions. Exists at 0x000000000000000000000000000000000000006E
*/
interface ArbRetryableTx {

    /**
    * @notice Redeem a redeemable tx.
    * Revert if called by an L2 contract, or if txId does not exist, or if txId reverts.
    * If this returns, txId has been completed and is no longer available for redemption.
    * If this reverts, txId is still available for redemption (until it times out or is canceled).
    @param txId unique identifier of retryabale message: keccak256(requestID, uint(0) )
     */
    function redeem(bytes32 txId) external;

    /**
    * @notice Return the minimum lifetime of redeemable txn.
    * @return lifetime in seconds
    */
    function getLifetime() external view returns(uint);

    /**
    * @notice Return the timestamp when txId will age out, or zero if txId does not exist.
    * The timestamp could be in the past, because aged-out txs might not be discarded immediately.
    * @param txId unique identifier of retryabale message: keccak256(requestID, uint(0) )
    * @return timestamp for txn's deadline
    */
    function getTimeout(bytes32 txId) external view returns(uint);

    /**
    * @notice Return the price, in wei, of submitting a new retryable tx with a given calldata size.
    * @param calldataSize call data size to get price of (in wei)
    * @return (price, nextUpdateTimestamp). Price is guaranteed not to change until nextUpdateTimestamp.
    */
    function getSubmissionPrice(uint calldataSize) external view returns (uint, uint);

    /**
     * @notice Return the price, in wei, of extending the lifetime of txId by an additional lifetime period. Revert if txId doesn't exist.
     * @param txId  unique identifier of retryabale message: keccak256(requestID, uint(0) )
     * @return (price, nextUpdateTimestamp). Price is guaranteed not to change until nextUpdateTimestamp.
    */
    function getKeepalivePrice(bytes32 txId) external view returns(uint, uint);

    /**
    @notice Deposits callvalue into the sender's L2 account, then adds one lifetime period to the life of txId.
    * If successful, emits LifetimeExtended event.
    * Revert if txId does not exist, or if the timeout of txId is already at least one lifetime in the future, or if the sender has insufficient funds (after the deposit).
    * @param txId unique identifier of retryabale message: keccak256(requestID, uint(0) )
    * @return New timeout of txId.
    */
    function keepalive(bytes32 txId) external payable returns(uint);

    /**
    * @notice Return the beneficiary of txId.
    * Revert if txId doesn't exist.
    * @param txId unique identifier of retryabale message: keccak256(requestID, uint(0) )
    * @return address of beneficiary for transaction
    */
    function getBeneficiary(bytes32 txId) external view returns (address);

    /**
    @notice Cancel txId and refund its callvalue to its beneficiary.
    * Revert if txId doesn't exist, or if called by anyone other than txId's beneficiary.
    @param txId unique identifier of retryabale message: keccak256(requestID, uint(0) )
    */
    function cancel(bytes32 txId) external;

    event LifetimeExtended(bytes32 indexed txId, uint newTimeout);
    event Redeemed(bytes32 indexed txId);
    event Canceled(bytes32 indexed txId);
}


```

This ArbRetryableTx interface is instantiated and exposed `bridge` class of [arb-ts](https://github.com/OffchainLabs/arbitrum/tree/master/packages/arb-ts), i.e.,

```ts
myBridge.ArbRetryableTx.redeem('mytxid')
```

## Arbitrum to Ethereum

### Explanation

L2 to L1 messages work similar to L1 to L2 messages, but in reverse: an L2 transaction is published with an L1 message as encoded data, to be executed later.

A key difference, however, is that in the L2 to L1 direction, a user must wait for the dispute period to pass between publishing their messages and actually executing it on L1; this is a direct consequence of the security model of Optimistic Rollups (see [finalty](Finality.md).) Additionally, unlike retyable tickets, outgoing messages have no upper bounded timeout; once the dispute window passes, they can be executed at any point. No rush.

### L2 to L1 Messages Lifecycle && API

The lifecycle of sending a message from layer 2 to layer 1 can be broken down into roughly 4 steps, only 2 which (at most!) require the end user to publish transactions.

**1. Publish L2 to L1 transaction (Arbitrum transaction)**

A client initiates the process by publishing a message on L2 via `ArbSys.sendTxToL1` (see [ArbSys](Arbsys.md), and see [ArbTokenBridge.\_withdraw](https://github.com/OffchainLabs/arbitrum/blob/5bd9a456a780582715a62affd887d35e2eb138b0/packages/arb-bridge-peripherals/contracts/tokenbridge/arbitrum/ArbTokenBridge.sol#L256) for example usage.

**2. Outbox entry gets created**

After the Arbitrum chain advances some set amount of time, ArbOS gathers all outgoing messages, Merklizes them, and publishes the root as an [OutboxEntry](https://github.com/OffchainLabs/arbitrum/blob/master/packages/arb-bridge-eth/contracts/bridge/OutboxEntry.sol) in the chain's outbox. Note that this happens "automatically"; i.e., it requires no additional action from the user.

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

Note that convenience methods for the steps outlined here are provided in the [arb-ts](https://github.com/OffchainLabs/arbitrum/tree/master/packages/arb-ts) client side library.

For relevant example usage, see [integration tests](https://github.com/OffchainLabs/arbitrum/blob/master/packages/arb-ts/integration_test/arb-bridge.test.ts.md) and our [Token Bridge UI](https://github.com/OffchainLabs/arb-token-bridge).

TODO: execution market?
