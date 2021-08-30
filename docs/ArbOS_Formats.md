---
id: ArbOS_Formats
title: ArbOS message / log formats
sidebar_label: ArbOS Data Formats
---

This specification describes the format of messages used for communication between clients, the EthBridge, and ArbOS. It includes incoming messages that are sent to ArbOS, outgoing messages that are emitted by ArbOS and recorded on the L1 chain, and log items emitted by ArbOS.

In this specification, all integers are big-endian. Uint means an unsigned integer. Unless otherwise specified, uints are 256 bits.

Hashing uses Ethereum's keccak256 algorithm, unless otherwise specified.

## Chain ID

Every Arbitrum chain has a 48-bit chain ID, which is the low-order 48 bits of the L1 Ethereum address of the chain's EthBridge contract.

## Incoming messages

Incoming messages are put into a chain's EthBridge-managed inbox, and received by a chain's instance of ArbOS.

An incoming message is a 6-tuple:

- message type (uint)
- L1 block number (uint): L1 block number when this message was inserted into the inbox
- L1 timestamp (uint): timestamp of L1 block when this message was inserted into the inbox
- Sender (address encoded as uint)
- RequestID: 0 for the first message inserted into the inbox; otherwise 1 + the requestID of the previous message inserted into the inbox
- Size of type-specific data (uint)
- Type-specific data: (buffer)

The L1 block number and/or L1 timestamp fields can be set to zero. Zero values in these fields will be replaced, by ArbOS, with the value of the same field in the previous message. If there was no previous message, ArbOS will leave these values as zero. (Note that the EthBridge will never create messages with zeroed block number or timestamp fields. The treatment of zero block number and timestamp values exists only as a convenience for use in private executions of a chain.)

Each message type is associated with rules, imposed by the Arbitrum protocol, regarding which properties the EthBridge must verify before sending a specific message type. These rules are not described here because they are not a part of the data format.

##### Message type 0: Eth deposit

[This is no longer supported.]

##### Message type 1: ERC20 deposit

[This is no longer supported.]

##### Message type 2: ERC721 deposit

[This is no longer supported.]

##### Message type 3: L2 message

This message type is initiated by a client, via a transaction to the EthBridge. Its purpose is to deliver to ArbOS an L2 data payload which the EthBridge does not need to understand. The EthBridge simply passes on the type-specific data uninterpreted. ArbOS will parse and validate the L2 data.

Details of L2 message subtypes and formats are listed in a separate section below.

##### Message type 4: chain initialization message

This message type is initiated by the EthBridge, as part of the creation of a new L2 chain, in order to convey parameters of the chain to ArbOS. It must only be sent as the first message in the inbox of a new chain.

Type-specific data:

- challenge period, in seconds (uint)
- ArbGas speed limit, in ArbGas per second (uint)
- maximum number of execution steps allowed in an assertion (uint)
- minimum stake requirement, in Wei (uint)
- address of the staking token, or zero if staking in ETH (address encoded as uint)
- address of the chain's owner (address encoded as uint)
- option data

Option data consists of a sequence of zero or more chunks. ArbOS will ignore a chunk if it does not know how to handle that chunk's option ID.

Each chunk is:

- option ID (64-bit uint)
- option payload length (64-bit uint)
- option payload

At present. the following options are supported:

- [Option 0 is currently unused]
- [Option 1 is currently unused]
- Option 2: set charging parameters:
  - speed limit per second (uint);
  - L1 gas per L2 tx (uint);
  - ArbGas per L2 tx (uint);
  - L1 gas per L2 calldata unit (uint) [a non-zero calldata byte is 16 units; a zero calldata byte is 4 units];
  - ArbGas per L2 calldata unit (uint);
  - L1 gas per storage unit allocated (uint);
  - ratio of L1 gas price to base ArbGas price;
  - network fee recipient (address encoded as uint);
  - congestion fee recipient (address encoded as uint)
- Option 3: set default aggregator
  - Default aggregator address (address encoded as uint)

All other options are ignored at present.

##### Message type 5: buddy contract creation

[This is no longer supported.]

##### Message type 6: end of Arbitrum block

A message of this type directs ArbOS to end the current Arbitrum block and start a new one. All integer or address fields (other than the message type) should be zero, and the buffer should be empty.

**Message type 7: L2 transaction funded by L1**

This message type encodes an L2 transaction that is funded by calldata provided at L1. The type-specific data must be the same as an L2 message of subtype 0 or 1.

**Message type 8: Rollup protocol event**

[This is not yet documented.]

**Message type 9: Send tx to retry buffer**

This message type encodes and delivers an L2 transaction; if gas is provided, it will be executed immediately. If no gas is provided or the execution reverts, it will be placed in the L2 retry buffer.

Type-specific data:

- destination address (address encoded as uint)
- callvalue, in wei (uint)
- L1-to-L2 deposit, in wei (uint)
- maximum submission cost (uint)
- credit-back address (address encoded as uint)
- beneficiary (address encoded as uint)
- max gas for immediate redemption request (uint) [if zero, don't try to redeem immediately]
- gas price for immediate redemption request (uint)
- calldata size (uint)
- calldata

If this message is properly formatted, the L1-to-L2 deposit amount will be credited to the sender's L2 ETH account.

Then, if the caller's L2 balance (after the L1-to-L2 deposit has occurred) is at least callvalue+maximumSubmissionCost, the transaction will proceed.

- _callvalue+maximumSubmissionCost_ will be deducted from the caller's L2 ETH account,
- _currentSubmissionCost_ will be collected by ArbOS as a submission fee,
- _maximumSubmissionCost-currentSubmissionCost_ will be credited to the creditBack address's L2 account,
- a retryable tx will be created and held in ArbOS's buffer, containing _callvalue_, with the specified beneficiary [if the retryable tx times out or is canceled by the beneficiary, the callvalue will be refunded to the beneficiary]
- ArbOS will emit a transaction receipt reporting success, with 32 bytes of return data: a transaction ID for the retryable tx (uint) which will be equal to keccak256(submissionID, uint(0) ), where submissionID is the requestID of this message
- if "the max gas for immediate redemption request" field is non-zero, and the sender has at least (maxgas x gasprice) in its ETH account, then
  - (maxgas x gasprice) is transferred from the sender's ETH account to the credit-back address's ETH account,
  - a transaction is immediately created, as if the sender had called RetryableTx.redeem(retryableTxId), with callvalue 0 and the specified maxgas and gasprice, and with transaction ID keccak256(submissionID, uint(1)). The credit-back address pays for the transaction's gas.

Otherwise, ArbOS will emit a transaction receipt reporting a failure code, with no return data.

**Message type 10: L2 batch, out-of-band processing for gas estimation**

This message type can't be sent to the chain but is only used by Arbitrum nodes to estimate gas usage of submitted transactions.

The type-specific data consists of:

- 3 (byte)
- aggregator address (address encoded as uint)
- limit on computation ArbGas (uint)
- an L2 message of subtype 7

This executes the enclosed L2 message, with the following deviations from the normal semantics.

- The aggregator address specified in the message is considered to be the aggregator that submitted this message.
- The transaction signature on the L2 message is ignored, and the enclosed transaction is treated as if it carried a valid signature by the sender of this L1 message.
- The ArbGas used for computation will be limited by the supplied limit. (The supplied limit will be ignored if it is larger than the ordinary limit that applies to all transactions.)

## L2 messages

As noted above, an L2 message is one type of incoming message that can be put into an L2 chain's inbox. The purpose of an L2 message is to convey information, typically a transaction request, to ArbOS. Except where specified here, the EthBridge does not examine or interpret the contents of an L2 message.

An L2 message consists of:

- an L2 message subtype (byte)
- subtype-specific data (byte array)

**Subtype 0: unsigned tx from user** has subtype-specific data of:

- ArbGas limit (uint)
- ArbGas price bid, in wei (uint)
- sequence number (uint)
- destination address (uint)
- callvalue, in wei (uint)
- calldata (bytes)

**Subtype 1: tx from contract** has subtype-specific data of:

- ArbGas limit (uint)
- ArbGas price bid, in wei (uint)
- destination address (uint)
- callvalue, in wei (uint)
- calldata (bytes)

**Subtype 2: non-mutating call** has subtype-specific data of:

- ArbGas limit (uint)
- ArbGas price bid, in wei (uint)
- destination address (uint)
- calldata (bytes)

**Subtype 3: L2 message batch** has subtype-specific data consisting of a sequence of one or more items, where each item consists of:

- L2 message length (64-bit uint)
- L2 message (byte array)

The L2 messages in a batch will be separated, and treated as if each had arrived separately, in the order in which they appear in the batch.

The enclosed L2 message may have any valid subtype.

**Subtype 4: signed tx from user** has subtype-specific data that is identical to the standard Ethereum encoded transaction format. The subtype-specific data consists of an RLP-encoded list containing:

- ArbGas limit (RLP-encoded uint)
- ArbGas price bid, in wei (RLP-encoded uint)
- sequence number (RLP-encoded uint)
- destination address (RLP-encoded address)
- callvalue, in wei (RLP-encoded uint)
- calldata (RLP-encoded byte array)
- v (RLP-encoded uint)
- r (RLP-encoded uint)
- s (RLP-encoded uint)

Here v, r, and s comprise an EIP-155 compliant ECDSA signature by the transaction's sender, based on the L2 chain's chainID.

The destination address is encoded consistently with Ethereum: a zero address is encoded as an empty byte array, and any other value is encoded as an array of 20 bytes.

**Subtype 5** is reserved for future use.

**Subtype 6: heartbeat message** has no subtype-specific data. This message has no effect, except to notify ArbOS that the block number and timestamp in the enclosing L1 message has been reached. ArbOS merely notes the block number and timestamp, then discards the message.

**Subtype 7: signed compressed transaction** encodes a signed, compressed transaction. The subtype-specific data is defined by the compression format, which is documented elsewhere.

**Subtype 8: BLS signed message batch** encodes a batch of transactions, which are signed using a BLS aggregate signature. The subtype-specific data is:

- number of messages in batch (RLP-encoded uint)
- BLS signature (2 uints)
- for each message: compressed message data (in compressed format that is described elsewhere)

## Logs

ArbOS emits three types of log items: transaction receipts, block summaries, and outgoing message contents.

### Tx receipts

ArbOS emits one log item for each transaction request it receives. A transaction request is an L2 message of subtype 0, 1, 4, or 7, which might arrive in its own incoming message or might be part of a batch. Regardless, each individual request will cause its own separate tx receipt to be emitted.

ArbOS will make its best effort to emit a tx receipt for each transaction request received, regardless of whether the transaction succeeds or fails; but this will not be possible for certain kinds of erroneous requests.

A tx receipt log item consists of:

- 0 (uint)
- incoming request info, an 8-tuple consisting of:
  - 3 (uint)
  - L1 block number (uint)
  - Arbitrum block number (uint)
  - L2 timestamp (uint)
  - address of sender (address represented as uint)
  - requestID (uint) [described below]
  - a 2-tuple consisting of:
    - size of L2 message (uint)
    - contents of L2 message for the request (buffer)
  - a 3-tuple consisting of:
    - provenance info
    - aggregator info
    - whether message was artificially injected via sideload (boolean encoded as uint 0 or 1)
- tx result info, a 3-tuple consisting of:
  - return code (uint) [described below]
  - returndata (2-tuple of size (uint) and contents (buffer))
  - EVM logs [format described below]
- ArbGas info, as 2-tuple consisting of:
  - ArbGas used (uint)
  - ArbGas price paid, in wei (uint)
- cumulative info in Arbitrum block, a 3-tuple consisting of:
  - ArbGas used in current Arbitrum block including this tx (uint)
  - index of this tx within this Arbitrum block (uint)
  - number of EVM logs emitted in this Arbitrum block before this tx (uint)
- fee information for the transaction, a 4-tuple consisting of:
  - a 4-tuple of prices of:
    - L2 transaction (uint)
    - L1 calldata bytes (uint)
    - L2 storage (uint)
    - L2 computation (uint)
  - a 4-tuple of units used of:
    - L2 transaction (uint, will always be 1)
    - L1 calldata bytes (uint)
    - L2 storage (uint)
    - L2 computation (uint)
  - a 4-tuple of wei paid for: [might not equal product of units and price, e.g. if user has insufficient funds to pay, or no aggregator was reimbursed]
    - L2 transaction (uint)
    - L1 calldata bytes (uint)
    - L2 storage (uint)
    - L2 computation (uint)
  - address of aggregator that was reimbursed (or zero if there wasn't one) (address encoded as uint)

Possible return codes are:
0: tx returned (success)
1: tx reverted
2: tx dropped due to L2 congestion
3: insufficient funds to pay for ArbGas
4: insufficient balance for callvalue
5: [no longer used (previously: bad sequence number, superseded by 14 and 15)]
6: message format error
7: cannot deploy at requested address
8: exceeded tx gas limit
9: insufficient gas to cover L1 charges
10: below minimum gas for a tx
11: gas price too low
12: no gas to auto-redeem retryable ticket
13: sender not permitted
14: sequence number too low
15: sequence number too high
255: unknown error

EVM logs are formatted as an EVM value, as a linked list in reverse order, such as this: (_log3_, (_log2_, (_log1_, (_log0_, () ) ) ) ). In this example there are four EVM log items, with the first one being _log0_ and the last being _log3_. Each EVM log is structured as an AVM tuple _(address, (dataSize, dataBuffer), topic0, topic1, ...)_, with as many topics as are present in that particular EVM log item.

##### Request IDs

A requestID is a uint that uniquely identifies a transaction request.

For a signed transaction, the requestID is the same value that Ethereum would use for the same transaction: the hash of the RLP-encoded transaction data (which is the subtype-specific data for subtype 4).

For an unsigned transaction that is an L2 message of subtype 0, the requestID is computed as:
hash(
sender address (as uint),
hash (
chainID (uint),
hash(subtype-specific data)
)
)

For other transactions, the requestID is computed from incoming message contents as follows. An incoming message is assigned a requestID of hash(chainID, inboxSeqNum), where inboxSeqNum is the value N such that this is the Nth message that has ever arrived in the chain's inbox. If the incoming message includes a batch, the K'th item in the batch is assigned a requestID of hash(requestID of batch, K). If batches are nested, this rule is applied recursively.

It is infeasible to find two distinct requests that have the same requestID. This is true because requestIDs are the output of a collision-free hash function, and it is not possible to create two distinct requests that will have the same input to the hash function. Signed transaction IDs cannot collide with the other types, because the other types' hash preimages both start with a zero byte (because sender address and chainID are zero-filled in the most-significant byte of a big-endian value) and the RLP encoding of a list cannot start with a zero byte. The other two types cannot have the same hash preimage because subtype-0 messages use a hash output as their second word, which with overwhelming probability will be too large to be feasible as the sequence number or batch index that occupies the same position in the default request ID scheme.

When an incoming message is included through the delayed inbox, the inbox sequence number gets adjusted so it doesn't overlap with the sequencer's inbox. For these messages, you can calculate the request ID by masking the high order bit, as follows: inboxSeqNum | (1 << 255).

### Block summary

A block summary is emitted at the end of every L1 block that contains any L2 transactions. No summary is emitted for a block that has no L2 activity.

A block summary item consists of:

- 1 (uint)
- Arbitrum block number (uint)
- timestamp (uint)
- current ArbGas limit per block (uint)
- statistics for this block: 5-tuple of
  - total ArbGas used (uint)
  - number of transactions (uint)
  - number of EVM logs (uint)
  - number of AVM logs (uint)
  - number of sends (uint)
- statistics for the chain since it was created (same format as previous item)
- gas accounting summary: 5-tuple of:
  - current ArbGas price in wei (uint)
  - size of current ArbGas pool (uint)
  - reserve funds in wei (uint)
  - total wei paid to validators over all time (uint)
  - address receiving validator payments (address encoded as uint)
- previous block number that had a block summary, or 0 if this is the first block to have a block summary
- Ethereum block number

### Outgoing message contents

ArbOS supports sending messages from Arbitrum contracts to L1. The contents of each message are emitted as a log item; and a Merkle root covering a batch of outgoing messages will later be published to the L1 as an Arbitrum Send.

The log item to publish the contents of a message consists of:

- 2 (uint)
- outgoing message batch number (uint)
- index within outgoing message batch (uint)
- size of message (uint)
- message contents (buffer)

## Arbitrum Sends

Arbitrum Sends are values emitted by an Arbitrum chain that are recorded at L1. A send consists of a sequence of bytes, encoded as a pair: size (uint) and contents (buffer).

The contents of a Send consist of:

- Send type (byte)
- Type-specific data

Currently only one type is supported: an L2-to-L1 call, which has send type 3. Its type-specific data consists of:

- L2 caller (address encoded as uint)
- L1 destination (address encoded as uint)
- Arbitrum block number (uint)
- Ethereum block number (uint)
- timestamp (uint)
- callvalue (uint)
- calldata (sequence of bytes)
