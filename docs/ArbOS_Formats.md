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
- Type-specific data: (byte array)

The L1 block number and/or L1 timestamp fields can be set to zero. Zero values in these fields will be replaced, by ArbOS, with the value of the same field in the previous message. If there was no previous message, ArbOS will leave these values as zero. (Note that the EthBridge will never create messages with zeroed block number or timestamp fields. The treatment of zero block number and timestamp values exists only as a convenience for use in private executions of a chain.)

Each message type is associated with rules, imposed by the Arbitrum protocol, regarding which properties the EthBridge must verify before sending a specific message type. These rules are not described here because they are not a part of the data format.

##### Message type 0: Eth deposit

It represents a transfer of Eth to an account on the L2 chain.

Type-specific data:

- L2 address to receive the Eth (address encoded as uint)
- number of Wei (uint)

##### Message type 1: ERC20 deposit

It represents a transfer of ERC20 tokens to an account on the L2 chain.

Type-specific data:

- address of the ERC20 token (address encoded as uint)
- L2 address to receive the tokens (address encoded as uint)
- number of Wei (uint)

##### Message type 2: ERC721 deposit

This message type must be initiated by the EthBridge. It represents a transfer of an ERC721 token to an account on the L2 chain.

Type-specific data:

- address of the ERC721 token (address encoded as uint)
- L2 address to receive the tokens (address encoded as uint)
- token identifier (uint)

##### Message type 3: L2 message

This message type is initiated by a client, via a transaction to the EthBridge. Its purpose is to deliver to ArbOS an L2 data payload which the EthBridge does not need to understand. The EthBridge simply passes on the type-specific data uninterpreted. ArbOS will parse and validate the L2 data.

Details of L2 message subtypes and formats are listed in a separate section below.

##### Message type 4: chain initialization message

This message type is initiated by the EthBridge, as part of the creation of a new L2 chain, in order to convey parameters of the chain to ArbOS. It must only be sent as the first message in the inbox of a new chain.

Type-specific data:

- challenge period, in milliseconds (uint)
- ArbGas speed limit, in ArbGas per second (uint)
- maximum number of execution steps allowed in an assertion (uint)
- minimum stake requirement, in Wei (uint)
- address of the staking token, or zero if staking in ETH (address encoded as uint)
- address of the chain's owner (address encoded as uint)
- option data

Option data consists of a sequence of zero or more chunks. ArbOS will ignore chunk IDs that it does not understand.

Each chunk is:

- option ID (64-bit uint)
- option payload length (64-bit uint)
- option payload

At present, no options are supported.

##### Message type 5: buddy contract creation

This message type is initiated by a call from an L1 contract to the EthBridge. The EthBridge must check that the call came from a contract, and reject it otherwise.

This message type allows an L1 contract to deploy an L2 contract at an L2 address that is equal to the contract's L1 address. This L2 deploy will be exactly like any ordinary L2 deploy, except for how the address of the deployed L2 contract is determined.

Type-specific data:

- maximum ArbGas to use (uint)
- ArbGas price bid, in wei (uint)
- Eth payment, in wei (uint)
- constructor code and data, encoded per Ethereum ABI (bytes)

##### Message type 6: reserved

This message type is reserved for internal use by ArbOS. It should never appear in the inbox.

##### Message type 7: Eth deposit transaction

This message type first transfers the given amount of Eth to the sender's account. Then it executes the given transaction.

Type-specific data:

- Unsigned Transaction or Contract transaction as defined in the L2 message section below

## L2 messages

As noted above, an L2 message is one type of incoming message that can be put into an L2 chain's inbox. The purpose of an L2 message is to convey information, typically a transaction request, to ArbOS. The EthBridge does not examine or interpret the contents of an L2 message.

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
- callvalue, in wei (uint)
- calldata (bytes)

**Subtype 3: L2 message batch** has subtype-specific data consisting of a sequence of one or more items, where each item consists of:

- L2 message length (64-bit uint)
- L2 message (byte array)

The L2 messages in a batch will be separated, and treated as if each had arrived separately, in the order in which they appear in the batch.

The enclosed L2 message may not have subtype 5 (sequencer batch). All other subtypes are allowed.

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

**Subtype 7: compressed signed tx** has subtype specific data of:

- sequence number (RLP-encoded uint)
- ArbGas price bid, in wei (RLP-encoded uint)
- ArbGas limit (RLP-encoded uint)
- destination [compressed address]
- callvalue [compressed amount]
- calldata [bytes]
- r [32 byte uint, big-endian]
- s [32 byte uint, big-endian]
- v low-order byte of signature v

Here, the compressed address is one of three things:

- an RLP encoding of a null value interpreted as the null address
- an RLP-encoding of a value 1-19 bytes in length, which is interpreted as an index into the IndexedAddressTable
- an RLP-encoding of a 20-byte value which is interpreted as an address

Compressed amount is encoded in scientific notation as A\*10^B:

- A is an RLP-encoded uint
- If A > 0, include B as a byte, otherwise do nothing

## Logs

ArbOS emits two types of log items: transaction receipts and block summaries.

### Tx receipts

ArbOS emits one log item for each transaction request it receives. A transaction request is an L2 message of subtype 0, 1, or 4, which might arrive in its own incoming message or might be part of a message batch or sequencer batch. Regardless, each individual request will cause its own separate tx receipt to be emitted.

ArbOS will make its best effort to emit a tx receipt for each transaction request received, regardless of whether the transaction succeeds or fails; but this will not be possible for certain kinds of erroneous requests.

A tx receipt log item consists of:

- 0 (uint)
- incoming request info consisting of:
  - 3 (uint)
  - L1 block number (uint)
  - L2 timestamp (uint)
  - address of sender (address represented as uint)
  - requestID (uint). [described below]
  - L2 message for the request (byte array)
- tx result info consisting of:
  - return code (uint) [described below]
  - returndata (byte array)
  - EVM logs [format described below]
- ArbGas info consisting of:
  - ArbGas used (uint)
  - ArbGas price paid, in wei (uint)
- cumulative info in L1 block consisting of:
  - ArbGas used in current L1 block including this tx (uint)
  - index of this tx within this L1 block (uint)
  - number of EVM logs emitted in this L1 block before this tx (uint)

Possible return codes are:
0: tx returned (success)
1: tx reverted
2: tx dropped due to L2 congestion
3: insufficient funds to pay for ArbGas
4: insufficient balance
5: bad sequence number
6: message format error
255: unknown error

EVM logs are formatted as an EVM value, as a linked list in reverse order, such as this: (_log3_, (_log2_, (_log1_, (_log0_, () ) ) ) ). In this example there are four EVM log items, with the first one being _log0_ and the last being _log3_. Each EVM log is structured as an AVM tuple _(address, marshalledData, topic0, topic1, ...)_, with as many topics as are present in that particular EVM log item.

##### Request IDs

A requestID is a uint that uniquely identifies a transaction request.

For a signed transaction, the requestID is the same value that Ethereum would use for the same transaction: the hash of the RLP-encoded transaction data (which is the subtype-specific data for subtype 4).

For an unsigned transaction that is an L2 message of subtype 0, the requestID is computed as:
hash(
sender address (as uint),
hash (
chainID (uint),
MarshalledDataHash(subtype-specific data)
)
)

For other transactions, the requestID is computed from incoming message contents as follows. An incoming message is assigned a requestID of hash(chainID, inboxSeqNum), where inboxSeqNum is the value N such that this is the Nth message that has ever arrived in the chain's inbox. If the incoming message includes a batch, the K'th item in the batch is assigned a requestID of hash(requestID of batch, K). If batches are nested, this rule is applied recursively.

It is infeasible to find two distinct requests that have the same requestID. This is true because requestIDs are the output of a collision-free hash function, and it is not possible to create two distinct requests that will have the same input to the hash function. Signed transaction IDs cannot collide with the other types, because the other types' hash preimages both start with a zero byte (because sender address and chainID are zero-filled in the most-significant byte of a big-endian value) and the RLP encoding of a list cannot start with a zero byte. The other two types cannot have the same hash preimage because subtype-0 messages use a hash output as their second word, which with overwhelming probability will be too large to be feasible as the sequence number or batch index that occupies the same position in the default request ID scheme.

##### MarshalledData and the MarshalledDataHash algorithm

The MarshalledData format is a way to encode an arbitrary-size byte array as a set of nested tuples, in a format easily consumable by an Arbitrum VM. The MarshalledData representation of a bytearray is the result of this pseudocode:

function marshal(ba: ByteArray) -> MarshalledBytes {
let nwords = (ba.size + 31) / 32;
let words = ();
let i = 0;
while (i < nwords) {
words = (ba[32*i .. 32*(i+1)], words); // zero-fill any bytes beyond end of ba
i = i + 1;
}
return (ba.size, words);
}

(In the above code, _ba[x..y]_ extracts bytes _x_ through _y-1_ inclusive from the byte array _ba_, and converts them in big-endian fashion to an unsigned integer. If _y_ is greater than the size of _ba_, any bytes beyond the end of _ba_ are implicitly zero-filled.)

The MarshalledDataHash algorithm computes a collision-free hash of a MarshalledData structure. It is defined by this pseudocode:

function marshalledDataHash(md: MarshalledData) -> uint256 {
let (size, contents) = md;
let ret = size;
while (contents != () ) {
ret = hash(ret, contents[0]);
contents = contents[1];
}
return ret;
}

## Outgoing messages

Outgoing messages reflect actions that require action at L1 or that need to be specifically visible to L1 contracts.

An outgoing message consists of:

- outgoing message type (uint)
- sender (address)
- type-specific data (byte array)

There are four outgoing message types.

**Type 0: Eth Withdrawal** is sent when an Eth withdrawal operation succeeds. It has type-specific data of:

- destination address (address encoded as uint)
- amount, in wei (uint)

**Type 1: ERC20 Withdrawal** is sent when an ERC20 withdrawal operation succeeds. It has type-specific data of:

- token address (address encoded as uint),
- destination address (address encoded as uint),
- amount (uint)

**Type 2: ERC721 Withdrawal** is sent when an ERC721 withdrawal operation succeeds. It has type-specific data of:

- token address (address encoded as uint),
- destination address (address encoded as uint),
- token ID (uint)

**Type 5: Buddy contract notification** is sent when a buddy contract deploy operation has concluded, whether or not the operation succeeded. It has one byte of type-specific data, which is 1 if the buddy contract was successfully created, or 0 otherwise.
