---
id: AVM
title: Inside Arbitrum: The AVM Architecture
sidebar_label: AVM Architecture
---
The Arbitrum Virtual Machine (AVM) is the interface between the Layer 1 and Layer 2 parts of Arbitrum. Layer 1 *provides* the AVM interface and ensures correct execution of the virtual machine. Layer 2 *runs on* the AVM virtual machine and provides the functionality to deploy and run contracts, track balances, and all of the things a smart-contract-enabled blockchain needs to do.

**Every Arbitrum chain has a single AVM** which does all of the computation and maintains all of the storage for everything that happens on the chain. Unlike some other systems which have a separate “VM” for each contract, Arbitrum uses a single virtual machine for the whole chain, much like Ethereum. The management of multiple contracts on an Arbitrum chain is done by software that runs on top of the AVM. 

At its core, a chain’s VM executes in this simple model, consuming messages from its inbox, changing its state, and producing outputs.

 ![img](https://lh4.googleusercontent.com/qwf_aYyB1AfX9s-_PQysOmPNtWB164_qA6isj3NhkDnmcro6J75f6MC2_AjlN60lpSkSw6DtZwNfrt13F3E_G8jdvjeWHX8EophDA2oUM0mEpPVeTlMbsjUCMmztEM0WvDpyWZ6R)

The starting point for the AVM design is the Ethereum Virtual Machine (EVM). Because Arbitrum aims to efficiently execute programs written or compiled for EVM, the AVM uses many aspects of EVM unchanged. For example, AVM adopts EVM's basic integer datatype (a 256-bit big-endian unsigned integer), as well as the instructions that operate on EVM integers.

## Why AVM differs from EVM

Differences between AVM and EVM are motivated by the needs of Arbitrum's Layer 2 protocol and Arbitrum's use of a interactive proving to resolve disputes.

### Execution vs. proving

Arbitrum, unlike EVM and similar architectures, supports both execution (advancing the state of a computation by executing it, which is always done off-chain in Arbitrum) and proving (convincing an L1 contract or other trusted party that a claim about execution is correct). EVM-based systems resolve disputes by re-executing the disputed code, whereas Arbitrum relies on a more efficient challenge protocol that leads to an eventual proof.

One nice consequence of separating execution from proving -- and never needing to re-execute blocks of code on an L1 chain -- is that we can optimize execution and proving for the different environments they’ll be used in. Execution is optimized for speed in a local, trusted environment, because local execution is the common case. Proving, on the other hand, will be needed less often but must still be efficient enough to be viable even on the busy Ethereum L1 chain. Proof-checking will rarely be needed, but proving must always be possible. The logical separation of execution from proving allows execution speed to be optimized more aggressively in the common case where proving turns out not to be needed.

### Requirements from ArbOS

Another difference in requirements is that Arbitrum uses ArbOS, an "operating system" that runs at Layer 2. ArbOS controls the execution of separate contracts to isolate them from each other and track their resource usage. In support of this, the AVM includes instructions to support saving and restoring the machine's stack, managing machine registers that track resource usage, and receiving messages from external callers. These instructions are used by ArbOS itself, but ArbOS ensures that they will never appear in untrusted code.

Supporting these functions in Layer 2 trusted software, rather than building them in to the L1-enforced rules of the architecture as Ethereum does, offers significant advantages in cost because these operations can benefit from the lower cost of computation and storage at Layer 2, instead of having to manage those resources as part of the Layer 1 EthBridge contract. Having a trusted operating system at Layer 2 also has significant advantages in flexibility, because Layer 2 code is easier to evolve, or to customize for a particular chain, than a Layer-1 enforced VM architecture would be.

The use of a Layer 2 trusted operating system does require some support in the architecture, for example to allow the OS to limit and track resource usage by contracts.

### Supporting Merkleization

Any Layer 2 protocol that relies on assertions and dispute resolution (which includes at least all rollup protocols) must define a rule for Merkle-hashing the full state of the virtual machine so that claims about parts of the state can be efficiently made to the base layer. That rule must be part of the architecture specification because it is relied upon in resolving disputes. It must also be reasonably efficient for validators to maintain the Merkle hash and/or recompute it when needed. This affects how the architecture structures its memory, for example. Any storage structure that is large and mutable will be relatively expensive to Merkleize, and a specific algorithm for Merkleizing it would need to be part of the architecture specification.

The AVM architecture responds to this challenge by having only bounded-size, immutable memory objects ("Tuples"), which can include other Tuples by reference. Tuples cannot be modified in-place but there is an instruction to copy a Tuple with a modification. This allows the construction of tree structures which can behave like a large flat memory. Applications can use functionalities such as large flat arrays, key-value stores, and so on, by accessing libraries that use Tuples internally.

The semantics of Tuples make it impossible to create cyclic structures of Tuples, so an AVM implementation can safely manage Tuples by using reference-counted, immutable structures. The hash of each Tuple value need only be computed once, because the contents are immutable.

### Codepoints: Optimizing code for proving

The conventional organization of code is to store a linear array of instructions, and keep a program counter pointing to the next instruction that will be executed. With this conventional approach, proving the result of one instruction of execution requires logarithmic time and space, because a Merkle proof must be presented to prove which instruction is at the current program counter.

The AVM does this more efficiently, by separating execution from proving. Execution uses the standard abstraction of an array of instructions indexed by a program counter, but proving uses an equivalent CodePoint construct that allows proving and proof-checking to be done in constant time and space. The CodePoint for the instruction at some PC value is the pair (opcode at PC, Hash(CodePoint at PC+1)). (If there is no CodePoint at PC+1, then zero is used instead.)

For proving purposes, the "program counter" is replaced by a "current CodePoint hash" value, which is part of the machine state. The preimage of this hash will contain the current opcode, and the hash of the following codepoint, which is everything a proof verifier needs to verify what the opcode is and what the current CodePoint hash value will be after the instruction, if the instruction isn't a Jump.

All jump instructions use jump destinations that are CodePoints, so a proof about execution of a jump instruction also has immediately at hand not only the PC that is being jumped to, but also what the contents of the "current CodePoint hash" register will be after the jump executes. In every case, proof and verification requires constant time and space.

In normal execution (when proving is not required), implementations will typically just use PC values as on a conventional architecture. However, when a one-step proof is needed, the prover can use a pre-built lookup table to get the CodePoint hashes corresponding to any relevant PCs.

### Creating code at runtime

Code is added to an AVM in two ways. First, some code is created when the AVM starts running. This code is read in from an AVM executable file (a .mexe file) and preloaded by the AVM emulator.

Second, the AVM has three instructions to create new CodePoints: one that makes a new Error CodePoint, and two that make new CodePoints (one for a CodePoint with an immediate value and one for a CodePoint without) given an opcode, possibly an immediate value, and a next CodePoint. These are used by ArbOS when translating EVM code for execution. (For more details on this, see the [ArbOS](ArbOS.md) section.)

### Getting messages from the Inbox

The *inbox* instruction consumes the next message from the VM’s inbox and pushes it onto the Data Stack. If all messages in the inbox have been consumed already, the inbox instruction blocks--the VM cannot complete the inbox instruction, nor can it do anything else, until a message arrives and the inbox instruction can complete. If the inbox has been completely consumed, any purported one-step proof of executing the inbox instruction will be rejected.The *inboxpeek* instruction does not consume a message from the inbox but simply reports whether or not the first unconsumed message in the inbox is at a specified block number. If there are no unconsumed messages in the inbox, *inboxpeek* blocks until there is one.

### Producing outputs

The AVM has two instructions that can produce outputs: *send* and *log*. Both are hashed into the output hash accumulator that records the (hash of) the VM’s outputs, but *send* causes its value to be recorded as calldata on the L1 chain, while *log* does not. This means that outputs produced with send will be visible to L1 contracts, while those produced with log will not. Of course, sends are more expensive than logs.
A useful design pattern is for a sequence of values to be produced as logs, and then a Merkle hash of those values to be produced as a single send. That allows an L1 contract to see the Merkle hash of the full sequence of outputs, so that it can verify the individual values when it sees them. ArbOS uses this design pattern, as described below.

### ArbGas and gas tracking

The AVM has a notion of ArbGas, which is like gas on Ethereum. ArbGas measures the cost of executing an instruction, based on how long it will take a validator to execute it. Every AVM instruction has an ArbGas cost. 

Arbitrum instructions have different gas costs than their Ethereum counterparts, for two reasons. First, the relative costs of executing Instruction A versus Instruction B can be different on a Layer 2 system versus on Ethereum. For example, storage accesses can be cheaper on Arbitrum relative to add instructions. ArbGas costs are based on the relative cost on Arbitrum.

The AVM architecture has a machine register called ArbGas Remaining. Before executing any instruction, the ArbGas cost of that instruction is deducted from ArbGas Remaining. If this would underflow the register (indicating that the execution is “out of ArbGas”) a hard error is generated and the ArbGasRemaining register is set to MaxUint256. 

The AVM has instructions to get and set the ArbGasRemaining register, which ArbOS uses to limit and count the ArbGas used by user contracts.

For information on ArbGas prices and other fee-related matters, see the Fees section.

### Error handling

Error conditions can arise in AVM execution in several ways, including stack underflows, ArbGas exhaustion, and type errors such as trying to jump to a value that is not a CodePoint. 

The AVM architecture has an Error CodePoint register that can be read and written by special instructions. When an error occurs, the Next CodePoint register is set equal to the Error CodePoint register, essentially jumping to the specified error handler. 