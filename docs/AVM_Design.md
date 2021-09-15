---
id: AVM_Design
title: AVM design rationale
sidebar_label: AVM design rationale
---

This document outlines the design rationale for the Arbitrum Virtual Machine (AVM) architecture.

The starting point for the AVM design is the Ethereum Virtual Machine (EVM). Because Arbitrum aims to efficiently execute programs written or compiled for EVM, Arbitrum uses many aspects of EVM unchanged. For example, AVM adopts EVM's basic integer datatype (a 256-bit big-endian unsigned integer), as well as the instructions that operate on EVM integers.

Differences between AVM and EVM are motivated by the needs of Arbitrum's Layer 2 protocol and Arbitrum's use of a multi-round challenge protocol to resolve disputes.

## Execution vs. proving

Arbitrum, unlike EVM and similar architectures, needs to support both execution (advancing the state of a computation by local emulation) and proving (convincing an L1 contract or other trusted party that a claim about execution is correct). EVM-based systems resolve disputes by re-executing the disputed code, whereas Arbitrum relies on a challenge protocol that leads to an eventual proof.

We want execution to be optimized for speed in a local, trusted environment, because local execution is the common case. Proving, on the other hand, will be needed less often but must still be efficient enough to be viable even in a congested L1 system. The system will rarely need to prove, but it always needs to be _prepared_ to prove. The logical separation of execution from proving allows execution speed to be optimized more aggressively in the common case where proving turns out not to be needed.

## ArbOS

Another difference in requirements is that Arbitrum uses ArbOS, a "operating system" that runs at Layer 2. ArbOS controls the execution of separate contracts to isolate them from each other and track their resource usage.

Supporting these functions in Layer 2 trusted software, rather than building them in to the L1-enforced rules of the architecture as Ethereum does, offers significant advantages in cost because these operations can benefit from the lower cost of computation and storage at Layer 2, instead of having to manage those resources as part of the Layer 1 EthBridge contract. Having a trusted operating system at Layer 2 also has significant advantages in flexibility, because Layer 2 code is easier to evolve, or to customize for a particular chain, than a Layer-1 enforced VM architecture would be.

The use of a Layer 2 trusted operating system does require some support in the architecture, for example to allow the OS to limit and track resource usage by contracts.

## Supporting Merkleization

Any Layer 2 protocol that relies on assertions and dispute resolution (which includes at least all rollup protocols) must define a rule for Merkle-hashing the full state of the virtual machine. That rule must be part of the architecture definition because it is relied upon in resolving disputes.

It must also be reasonably efficient to maintain the Merkle hash and/or recompute it when needed. This affects how the architecture structures its memory, for example. Any storage structure that is large and mutable will be relatively expensive to Merkleize, and an algorithm for Merkleizing it must be part of the architecture specification.

The AVM architecture responds to this challenge by having only bounded-size, immutable memory objects ("Tuples"), which can include other Tuples by reference. Tuples cannot be modified in-place but there is an instruction to copy a Tuple with a modification. This allows the construction of tree structures which can behave like a large flat memory. Applications can use functionalities such as large flat arrays, key-value stores, and so on, by accessing libraries that use Tuples internally.

The semantics of Tuples make it impossible to create cyclic structures of Tuples, so an AVM implementation can safely manage Tuples by using reference-counted, immutable structures. The hash of each Tuple value need only be computed once, because the contents are immutable.

## Code organization: Codepoints

The conventional organization of code is to store a linear array of instructions, and keep a program counter pointing to the next instruction that will be executed. With this conventional approach, proving an instruction of execution requires logarithmic time and space, because a Merkle proof must be presented to prove which instruction is under the current PC.

The AVM uses this conventional approach for execution, but it adds a feature that makes proving and proof-checking require constant time and space. The CodePoint for the instruction at some PC value is the pair (instruction at PC, Hash(CodePoint at PC+1)). (If there is no CodePoint at PC+1, then zero is used instead.)

For proving purposes, the "program counter" is replaced by a "current CodePoint hash" value, which is part of the machine state. The preimage of this hash will contain the current instruction, and the hash of the following codepoint, which is everything the verifier needs to verify what the instruction is and what the current CodePoint hash value will be after the instruction, if the instruction isn't a Jump.

All jump instructions use jump destinations that are CodePoints, so a proof about execution of a jump instruction also has immediately at hand not only the PC that is being jumped to, but also what the contents of the "current CodePoint hash" register will be after the jump executes. In every case, proof and verification requires constant time and space.

In normal execution (when proving is not required), implementations will typically just use PC values as on a conventional architecture. However, when a proof is needed, the prover can use a lookup table to get the CodePoint hashes corresponding to any relevant PCs.

## Support for ArbOS

ArbOS, running at Layer 2, isolates untrusted programs from each other, tracks and limits their resource usage, and manages the economic model that collects fees from users to fund the operation of a chain's validators.

In support of this, the AVM includes instructions to support saving and restoring the machine's stack, managing machine registers that track resource usage, and receives messages from external callers. These instructions are used by ArbOS itself, but ArbOS ensures that they will never appear in untrusted code.
