---
id: AVM_Specification
title: Arbitrum VM informal architecture spec
sidebar_label: Arbitrum VM Specification
---

This document describes informally the semantics of the Arbitrum VM architecture. This version is simplified, compared to the one described in the Arbitrum paper published in USENIX Security 2018. This version is also tailored to more closely match the Ethereum data formats and instruction set, to ease implementation of an EVM-to-AVM translator.

We expect that some implementations will optimize to save time and space, provided that the result is equivalent to what is described here.

## Values

A Value can have the following types:

-   Integer: a 256-bit integer;
-   Codepoint: a Value that represents a point in the executable code (implemented as a triple (pc, operation, nextcodehash);
-   Tuple: an array of up to 8 slots, numbered consecutively from zero, with each slot containing a Value.

Note that integer values are not explicitly specified as signed or unsigned. Where the distinction matters, this specification states which operations treat values as signed and which treat them as unsigned. (Where there is no difference, this spec does not specify signed vs. unsigned.)

The special value None refers to a 0-tuple (a tuple containing 0 elements).

Note that the slots of a Tuple may in turn contain other Tuples. As a result, a Tuple may represent an arbitrary multi-level tree (more precisely, a distributed acyclic graph) of values.

Note that all Values are immutable. As a consequence of this, it is impossible to make a Tuple that contains itself either directly or indirectly (reason: if Tuple A contains Tuple B, then A must have been created after B).

Because Tuples are immutable and form an acyclic structure, it is possible for an implementation to represent a Tuple as a pointer to a small array object, and to use reference counting to manage the lifetime of Tuples.

## Operations

There are two types of operations:

-   BasicOp: A simple opcode
-   ImmediateOp: A opcode along with a Value

Execution of an operation behaves as follows:

-   BasicOp: Execute the opcode according to the instruction definition included below
-   ImmediateOp: First push the included Value to the stack. Then execute the opcode as described below

## Stacks

A stack represents a (possibly empty) pushdown stack of Values. Stacks support Push and Pop operations, which have the expected semantics. Doing a Pop operation on a stack that is empty causes the VM to enter the Error state. A stack is represented as either None (for an empty stack) or the 2-tuple [topValue, restOfStack], where restOfStack in turn is a stack.

## Marshaling

Marshaling is an invertible mapping from a Value to an array of bytes. Outside the VM, Values are typically represented in Marshaled form.

Unmarshaling is the inverse operation, which turns an array of bytes into a Value. This document does not specify the Unmarshaling operation, except to say that for all Values V, Unmarshal(Marshal(V)) yields V. If a byte-array A could not possibly have been made by marshaling, then attempting to compute unmarshal(A) raises an Error.

An Integer marshals to

-   the byte 0
-   32-byte big-endian representation of the value.

A Codepoint marshals to:

-   byte val of 1
-   8-byte big-endian representation of the pc.
-   the marshalling of the operation
-   32-byte nextHash value

A BasicOp marshals to:

-   byte val of 0
-   1-byte representation of the opcode

A ImmediateOp marshals to:

-   byte val of 1
-   1-byte representation of the opcode
-   Marshalled value

A Tuple marshals to

-   byte val of (3 + the number of items in the tuple)
-   the concatenation of the results of marshaling the Value in each slot of the Tuple.

## Hashing Values

The Hash of an Integer is the Keccak-256 of the 32-byte big endian encoding of the integer, encoded as an Integer in big-endian fashion.

The Hash of a Tuple is computed by concatenating the byte value (3 + number of items in the tuple) followed by the Hashes of the items in tuple, and computing the Keccak-256 of the result, encoded as an Integer in big-endian fashion.

The Hash of a Codepoint containing a BasicOp is computed by concatenating the byte value 1, followed by the 1-byte representation of the opcode, followed by the 32-byte nextHash value, and computing the Keccak-256 of the result, encoded as an Integer in big-endian fashion.

The Hash of a Codepoint containing an ImmediateOp is computed by concatenating the byte value 1, followed by the 1-byte representation of the opcode, followed by the 32-byte Hash of the immediate value, followed by the 32-byte nextHash value, and computing the Keccak-256 of the result, encoded as an Integer in big-endian fashion.

## Virtual Machine State

The state of a VM is either a special state Halted, or a special state ErrorStop, or an extensive state.

An extensive state of a VM contains the following:

-   Current Codepoint: a Codepoint that represents the current point of execution;
-   Data Stack: a Stack that is the primary working area for computation;
-   Aux Stack: a Stack that provides auxiliary storage;
-   Register: a mutable storage cell that can hold a single Value;
-   Static: an immutable Value that is initialized when the VM is created;
-   Error Codepoint: a Codepoint that is meant to be used in response to an Error.

When a VM is initialized, it is in an extensive state. The Data Stack, Aux Stack, Register, and Error Codepoint are initialized to None, None, None, and the Codepoint (0, 0, 0), respectively. The entity that creates the VM supplies the initial values of the Current Codepoint and Static.

## Hashing a VM State

If a VM is in the Halted state, its state hash is the Integer 0.

If a VM is in the ErrorStop state, its state hash is the Integer 1.

If a VM is in an extensive state, its state hash is computed by concatenating the hash of the Instruction Stack, the hash of the Data Stack, the hash of the Aux Stack, the hash of the Register, the hash of the Static, and the hash of the Error Codepoint, hashing the result using Keccak-256.

## The Runtime Environment

The Runtime Environment is a VM’s interface to the outside world. Certain instructions are described as interacting with the Runtime Environment.

The implementation of the Runtime Environment will vary in different scenarios. In standalone AVM emulators, the Runtime Environment might be controlled by command-line options or a configuration file. In production uses of Arbitrum, the Runtime Environment will be provided by an Arbitrum Validator, and will be specified by Preconditions that form part of Assertions in the Arbitrum Protocol.

The runtime environment will supply values for:

-   lower and upper bounds on the time (Integers)
-   a Value that is the contents of VM’s Inbox
-   a reference to a “Wallet” object

Implementers and developers can assume that the Runtime Environment will satisfy the "time consistency" property: if the gettime instruction returns a lower bound of L, then all later executions of gettime will return an upper bound greater than or equal to L. This is equivalent to assuming that there is a hidden, non-decreasing time value that is always consistent with the bounds returned by gettime.

## Inbox and Messages

Each VM has an Inbox, which is supplied by the Runtime Environment. At any time the inbox holds a single value representing a sequence of messages. An Inbox contains either:

-   None, indicating that the Inbox is empty and contains no messages; or
-   a 2-tuple [b, v], where b is an Inbox and v is a value, indicating that the inbox contains the sequence of messages in b followed by the message v.

A VM can receive the contents of its Inbox by using the inbox instruction (described below), which pushes the current Inbox contents onto the Data Stack and empties the Inbox.

Messages may arrive in the Inbox at any time (but not during the execution of an inbox instruction). Any newly arriving messages will be in sequence after the messages that were previously in the Inbox.

## Errors

Certain conditions are said to “raise an Error”. If an Error is raised when the Error Codepoint is (0,0,0), then the VM enters the ErrorStop state. If an Error is raised when the Error Codepoint is not (0,0,0), then the Current Codepoint is set equal to the Error Codepoint.

## Blocking

Some instructions are said to “block” until some condition is true. This means that if the condition is false, the instruction cannot complete execution and the VM must stay in the state it was in before the instruction. If the condition is true, the VM can execute the instruction. The result is as if the VM stopped and waited until the condition became true.

## ArbGas

Every instruction consumes some amount of Arbitrum Gas, also known as ArbGas. (ArbGas is not the same as Ethereum Gas.)

The ArbGas costs of instructions might change in the future.

There is no charge for ArbGas and there is no limit on the amount of ArbGas that a VM can use. The purpose of ArbGas is to allow clients and validators to monitor and control how much computational work a VM is requiring. For example, it is likely that features will be added to allow validators to limit the amount of ArbGas that any one call to a VM or contract can consume.

## Instructions

A VM that is in the Halted or ErrorStop state cannot execute instructions.

A VM that is in an extensive state can execute an instruction. To execute an instruction, the VM gets the opcode from the Current Codepoint.

If the opcode is not a value that is an opcode of the AVM instruction set, then an Error is raised. Otherwise the VM carries out the semantics of the opcode.

The instructions are as follows:

| Opcode                                     | Nickname   | Semantics                                                                                                                                                                                                                                                                                                      | ArbGas cost |
| ------------------------------------------ | ---------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------- |
| 00s: Arithmetic Operations                 |            | &nbsp;                                                                                                                                                                                                                                                                                                         |
| 0x01                                       | add        | Pop two values (A, B) off the Data Stack. If A and B are both Integers, Push the value A+B (truncated to 256 bits) onto the Data Stack. Otherwise, raise an Error.                                                                                                                                             | 3           |
| 0x02                                       | mul        | Same as add, except multiply rather than add.                                                                                                                                                                                                                                                                  | 3           |
| 0x03                                       | sub        | Same as add, except subtract rather than add.                                                                                                                                                                                                                                                                  | 3           |
| 0x04                                       | div        | Pop two values (A, B) off the Data Stack. If A and B are both Integers and B is non-zero, Push A/B (unsigned integer divide) onto the Data Stack. Otherwise, raise an Error.                                                                                                                                   | 4           |
| 0x05                                       | sdiv       | Pop two values (A, B) off the Data Stack. If A and B are both Integers and B is non-zero, Push A/B (signed integer divide) onto the Data Stack. Otherwise, raise an Error.                                                                                                                                     | 7           |
| 0x06                                       | mod        | Same as div, except compute (A mod B), treating A and B as unsigned integers.                                                                                                                                                                                                                                  | 4           |
| 0x07                                       | smod       | Same as sdiv, except compute (A mod B), treating A and B as signed integers, rather than doing integer divide. The result is defined to be equal to sgn(A)\*(abs(A)%abs(B)), where sgn(x) is 1,0, or -1, if x is positive, zero, or negative, respectively, and abs is absolute value.                         | 7           |
| 0x08                                       | addmod     | Pop three values (A, B, C) off the Data Stack. If A, B, and C are all Integers, and C is not zero, Treating A, B, and C as unsigned integers, Push the value (A+B) % C (calculated without 256-bit truncation until end) onto the Data Stack. Otherwise, raise an Error.                                       | 4           |
| 0x09                                       | mulmod     | Same as addmod, except multiply rather than add.                                                                                                                                                                                                                                                               | 4           |
| 0x0a                                       | exp        | Same as add, except exponentiate rather than add.                                                                                                                                                                                                                                                              | 25          |
| &nbsp;                                     |            | &nbsp;                                                                                                                                                                                                                                                                                                         |
| 10s: Comparison & Bitwise Logic Operations |            | &nbsp;                                                                                                                                                                                                                                                                                                         |
| 0x10                                       | lt         | Pop two values (A,B) off the Data Stack. If A and B are both Integers, then (treating A and B as unsigned integers, if A<B, push 1 on the Data Stack; otherwise push 0 on the Data Stack). Otherwise, raise an Error.                                                                                          | 2           |
| 0x11                                       | gt         | Same as lt, except greater than rather than less than                                                                                                                                                                                                                                                          | 2           |
| 0x12                                       | slt        | Pop two values (A,B) off the Data Stack. If A and B are both Integers, then (treating A and B as signed integers, if A<B, push 1 on the Data Stack; otherwise push 0 on the Data Stack). Otherwise, raise an Error.                                                                                            | 2           |
| 0x13                                       | sgt        | Same as slt, except greater than rather than less than                                                                                                                                                                                                                                                         | 2           |
| 0x14                                       | eq         | Pop two values (A, B) off the Data Stack. If A and B have different types, raise an Error. Otherwise if A and B are equal by value, Push 1 on the Data Stack. (Two Tuples are equal by value if they have the same number of slots and are equal by value in every slot.) Otherwise, Push 0 on the Data Stack. | 2           |
| 0x15                                       | iszero     | If A is the Integer 0, push 1 onto the Data Stack. Otherwise, if A is a non-zero Integer, push 0 onto the Data Stack. Otherwise (i.e., if A is not an Integer), raise an Error.                                                                                                                                | 1           |
| 0x16                                       | and        | Pop two values (A, B) off the Data Stack. If A and B are both Integers, then push the bitwise and of A and B on the Data Stack. Otherwise, raise an Error.                                                                                                                                                     | 2           |
| 0x17                                       | or         | Same as and, except bitwise or rather than bitwise and                                                                                                                                                                                                                                                         | 2           |
| 0x18                                       | xor        | Same as and, except bitwise xor rather than bitwise and                                                                                                                                                                                                                                                        | 2           |
| 0x19                                       | not        | Pop one value (A) off the Data Stack. If A is an Integer, then push the bitwise negation of A on the Data Stack. Otherwise, raise an Error.                                                                                                                                                                    | 1           |
| 0x1a                                       | byte       | Pop two values (A, B) off the Data Stack. If A and B are both Integers, (if B<32 (interpreting B as unsigned), then push the B’th byte of A onto the Data Stack, otherwise push Integer 0 onto the Data Stack). Otherwise, raise an Error.                                                                     | 4           |
| 0x1b                                       | signextend | Pop two values (A, B) off the Data Stack. If A and B are both Integers, (if B<31 (interpreting B as unsigned), sign extend A from (B + 1) \* 8 bits to 256 bits and Push the result onto the Data Stack; otherwise push A onto the Data Stack). Otherwise, raise an Error.                                     | 7           |
| &nbsp; | | &nbsp; |
| 20s: Hashing | | &nbsp; |
| 0x20 | hash | Pop a Value (A) off of the Data Stack. Push Hash(A) onto the Data Stack. | 40 |
| 0x21 | type | Pop a Value (A) off of the Data Stack. If A is an Integer, Push Integer 0 onto the Data Stack. Otherwise, if A is a Codepoint, Push Integer 1 onto the Data Stack. Otherwise (A is a Tuple), push Integer 3 onto the Data Stack. | 3 |
| 0x22 | ethhash2 | Pop two Values (A,B) off of the Data Stack. If A and B are both Integers, convert (big-endian) each of A and B into length-32 byte arrays, concatenate the two into a 64-byte array, compute the Ethereum hash of that byte-array, convert the result into an Integer (big-endian), and Push the resulting Integer onto the Data Stack. Otherwise, raise an Error. | 40 |
| &nbsp; | | &nbsp; |
| 30s: Stack, Memory, Storage and Flow Operations | | &nbsp; |
| 0x30 | pop | Pop one value off of the Data Stack, and discard that value. | 1 |
| 0x31 | spush | Push a copy of Static onto the Data Stack. | 1 |
| 0x32 | rpush | Push a copy of Register onto the Data Stack. | 1 |
| 0x33 | rset | Pop a Value (A) off of the Data Stack. Set Register to A. | 2 |
| 0x34 | jump | Pop one value (A) off of the Data Stack. If A is a Codepoint, set the Instruction Stack to A. Otherwise raise an Error. | 4 |
| 0x35 | cjump | Pop two values (A, B) off of the Data Stack. If A is not a Codepoint or B is not an Integer, raise an Error. Otherwise (If B is zero, do Nothing. Otherwise set the Instruction Stack to A.). | 4 |
| 0x36 | stackempty | If the Data Stack is empty, Push 1 on the Data Stack, otherwise Push 0 on the Data Stack. | 2 |
| 0x37 | pcpush | Push the Codepoint of the currently executing operation to the Data Stack. | 1 |
| 0x38 | auxpush | Pop one value off the Data Stack, and push it to the Aux Stack. | 1 |
| 0x39 | auxpop | Pop one value off the Aux Stack, and push it to the Data Stack. | 1 |
| 0x3a | auxstackempty | If the Aux Stack is empty, Push 1 on the Data Stack, otherwise Push 0 on the Data Stack. | 2 |
| 0x3b | nop | Do nothing. | 1 |
| 0x3c | errpush | Push a copy of the Error Codepoint onto the Data Stack. | 1 |
| 0x3d | errset | Pop a Value (A) off of the Data Stack. Set the Error Codepoint to A. | 1 |
| &nbsp; | | &nbsp; |
| 40s: Duplication and Exchange Operations | | &nbsp; |
| 0x40 | dup0 | Pop one value (A) off of the Data Stack. Push A onto the Data Stack. Push A onto the Data Stack. | 1 |
| 0x41 | dup1 | Pop two values (A,B) off the Data Stack. Push B,A,B onto the Data Stack, in that order. | 1 |
| 0x42 | dup2 | Pop three values (A,B,C) off the Data Stack. Push C,B,A,C onto the Data Stack, in that order. | 1 |
| 0x43 | swap1 | Pop two values (A,B) off the Data Stack. Push A onto the Data Stack. Push B onto the Data Stack. | 1 |
| 0x44 | swap2 | Pop three values (A,B,C) off the data Stack. Push A,B,C onto the Data Stack, in that order. | 1 |
| &nbsp; | | &nbsp; |
| 50s: Tuple Operations | | &nbsp; |
| 0x50 | tget | Pop two values (A,B) off the Data Stack. If B is a Tuple, and A is an integer, and A>=0 and A is less than length(B), then Push the value in the A_th slot of B onto the Data Stack. Otherwise raise an Error. | 2 |
| 0x51 | tset | Pop three values (A,B,C) off of the Data Stack. If B is a Tuple, and A is an Integer, and A>=0, and A is less than length(B), then create a new Tuple that is identical to B, except that slot A has been set to C, then Push the new Tuple onto the Data Stack. Otherwise, raise an Error. | 15 |
| 0x52 | tlen | Pop a value (A) off the Data Stack. If A is a Tuple, push the length of A (i.e. the number of slots in A) onto the Data Stack. Otherwise, raise an Error. | 2 |
| &nbsp; | | &nbsp; |
| 60s: Logging Operations | | &nbsp; |
| 0x60 | breakpoint | In an AVM emulator, return control to the Runtime Environment. | 100 |
| 0x61 | log | Pop a Value (A) off the Data Stack, and convey A to the Runtime Environment as a log event. | 100 |
| &nbsp; | | &nbsp; |
| 70s: System operations | | &nbsp; |
| 0x70 | send | Pop a Value (A) off the Data Stack. Tell the Runtime Environment to publish A as an outgoing message of this VM. | 100 |
| 0x71 | gettime | Push a 2-tuple [mintime, maxtime] onto the Data Stack, where mintime and maxtime are (Integer) lower and upper bounds on the current time, as supplied by the Runtime Environment. | 40 |
| 0x72 | inbox | Pop a Value (A) off of the Data Stack. Block until either (a) the VM's inbox is non-empty, or (b) the lower bound on current time (as would be returned by the gettime instruction) is greater than or equal to A. Then push the Inbox contents, as supplied by the Runtime Environment, onto the Data Stack, and set the inbox to None. | 40 |
| 0x73 | error | Raise an Error. | 5 |
| 0x74 | halt | Enter the Halted state. | 10 |
