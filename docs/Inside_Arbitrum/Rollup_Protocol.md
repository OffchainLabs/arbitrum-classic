---
id: Rollup_Protocol
title: Inside Arbitrum: The Rollup Protocol
sidebar_label: Rollup Protocol
---
Before diving into the rollup protocol, there are two things we need to cover.

First, *if you’re an Arbitrum user or developer, you don’t need to understand the rollup protocol*. You don’t ever need to think about it, unless you want to. Your relationship with it can be like a train passenger’s relationship with the train’s engine: you know it exists, you rely on it to keep working, but you don’t spend your time monitoring it or studying its internals. 

You’re welcome to study, observe, and even participate in the rollup protocol, but you don’t need to, and most people won’t. So if you’re a typical train passenger who just wants to read or talk to your neighbor, you can skip right to the next section of this document. If not, read on!

The second thing to understand about the rollup protocol is *that the protocol doesn’t decide the results of transactions, it only confirms the results*. The results are uniquely determined by the sequence of messages in the chain’s inbox. So once your transaction message is in the chain’s inbox, its result is knowable--and Arbitrum nodes will report your transaction to be done. The role of the rollup protocol is to confirm transaction results that, as far as Arbitrum users are concerned, already exist. (This is why Arbitrum users can ignore the rollup protocol.)

You might wonder why we need the rollup protocol. If everyone knows the results of transactions already, why bother confirming them? The protocol exists for two reasons. First, somebody might lie about a result, and we need a definitive, trustless way to tell who is lying. Second, Ethereum doesn’t know the results. The whole point of a Layer 2 scaling system is to run transactions without Ethereum needing to do all of the work--and indeed Arbitrum can go fast enough that Ethereum couldn’t hope to monitor every Arbitrum transaction. But once a result is confirmed, Ethereum knows about it and can rely on it.

With those preliminaries behind us, let’s jump into the details of the rollup protocol.

The parties who participate in the protocol are called *validators*. Anyone can be a validator. Some validators will choose to be stakers--they will place an ETH deposit which they’ll be able to recover if they’re not caught cheating. These roles are permissionless: anyone can be a validator or a staker.

The key security property of the rollup protocol is the *AnyTrust Guarantee*, which says that any one honest validator can force the correct execution of the chain to be confirmed. This means that execution of an Arbitrum chain is as trustless as Ethereum. You, and you alone (or someone you hire) can force your transactions to be processed correctly. And that is true no matter how many malicious people are trying to stop you.

## The Rollup Chain

The rollup protocol tracks a chain of rollup blocks. These are separate from Ethereum blocks. You can think of the rollup blocks as forming a separate chain, which the Arbitrum rollup protocol manages and oversees.

Validators can propose rollup blocks. New rollup blocks will be *unresolved* at first. Eventually every rollup block will be *resolved*, by being either *confirmed* or *rejected*. The confirmed blocks make up the confirmed history of the chain.

Each rollup block contains:

- rollup block number
- predecessor block number: rollup block number of the last block before this one that is (claimed to be) correct
- how much computation the chain has done in its history (measured in ArbGas)
- how many inbox messages have been consumed in the chain’s history
- a hash of the outputs produced over the chain’s history
- a hash of the chain state.



Except for the rollup block number, the contents of the block are all just claims by the block’s proposer. Arbitrum doesn’t know at first whether any of these fields are correct. If all of these fields are correct, the protocol should eventually confirm the block. If one or more of these fields are incorrect, the protocol should eventually reject the block.

A block is implicitly claiming that its predecessor block is correct. This implies, transitively, that a block implicitly claims the correctness of a complete history of the chain: a sequence of ancestor blocks that reaches all the way back to the birth of the chain.

A block is also implicitly claiming that its older siblings (older blocks with the same predecessor), if there are any, are incorrect. If two blocks are siblings, and the older sibling is correct, then the younger sibling is considered incorrect, even if everything else in the younger sibling is true.

The block is assigned a deadline, which says how long other validators have to respond to it. If you’re a validator, and you agree that a rollup block is correct, you don’t need to do anything. If you disagree with a rollup block, you can post another block with a different result, and you’ll probably end up in a challenge against the first block’s staker. (More on challenges below.)

n the normal case, the rollup chain will look like this:

![img](https://lh3.googleusercontent.com/vv118kJMXj76PG6J-Jv4BC9KTpe72mdfD1uWoqhKXvKKfPWHW6wMMCvJ9KKQx_VXIw34XfzT4yfyNVtQVstYRczLk6kLKvBv8Pbl-0MjSzGxz1Z_8T5Y_6UcDMWpy7_D9PxQYKdT)

On the left, representing an earlier part of the chain’s history, we have confirmed rollup blocks. These have been fully accepted and recorded by the EthBridge. The newest of the confirmed blocks, block 94, is called the “latest confirmed block.” On the right, we see a set of newer proposed rollup blocks. The EthBridge can’t yet confirm or reject them, because their deadlines haven’t run out yet. The oldest block whose fate has yet to be determined, block 95, is called the “first unresolved block.”

Notice that a proposed block can build on an earlier proposed block. This allows validators to continue proposing blocks without needing to wait for the EthBridge to confirm the previous one. Normally, all of the proposed blocks will be valid, so they will all eventually be accepted.

Here’s another example of what the chain state might look like, if several validators are being malicious. It’s a contrived example, designed to illustrate a variety of cases that can come up in the protocol, all smashed into a single scenario.

![img](https://lh3.googleusercontent.com/IKBNeX9IVAD5Vom8vqYER4CEZhTecJJrp51ddlEGYiZrdV6y9zaG0Ip8HuKgfJ-eS9_TN_C2I0EPl-7H5ITRgSQqJONnSE7X0P62sRbGoiv_shmijBxsVDJL9RhWbyDjs2lKxU-M)

There’s a lot going on here, so let’s unpack it. 

* Block 100 has been confirmed. 
* Block 101 claimed to be a correct successor to block 100, but 101 was rejected (hence the X drawn in it). 
* Block 102 was eventually confirmed as the correct successor to 100.
* Block 103 was confirmed and is now the latest confirmed block.
* Block 104 was proposed as a successor to block 103, and 105 was proposed as a successor to 104. 104 was rejected as incorrect, and as a consequence 105 was rejected because its predecessor was rejected.
* Block 106 is unresolved. It claims to be a correct successor to block 103 but the protocol hasn’t yet decided whether to confirm or reject it. It is the first unresolved block.
* Blocks 107 and 108 claim to chain from 106. They are also unresolved. If 106 is rejected, they will be automatically rejected too.
* Block 109 disagrees with block 106, because they both claim the same predecessor. At least one of them will eventually be rejected, but the protocol hasn’t yet resolved them.
* Block 110 claims to follow 109. It is unresolved. If 109 is rejected, 110 will be automatically rejected too.
* Block 111 claims to follow 105. 111 will inevitably be rejected because its predecessor has already been rejected. But it hasn’t been rejected yet, because the protocol resolves blocks in block number order, so the protocol will have to resolve 106 through 110, in order, before it can resolve 111. After 110 has been resolved, 111 can be rejected immediately.

Again: this sort of thing is very unlikely in practice. In this diagram, at least four parties must have staked on wrong nodes, and when the dust settles at least four parties will have lost their stakes. The protocol handles these cases correctly, of course, but they’re rare corner cases. This diagram is designed to illustrate the variety of situations that are possible in principle, and how the protocol would deal with them.

## Staking 

At any given time, some validators will be stakers, and some will not. Stakers deposit funds that are held by the EthBridge and will be confiscated if the staker loses a challenge. Currently all chains accept stakes in ETH.

A single stake can cover a chain of rollup blocks. Every staker is staked on the latest confirmed block; and if you’re staked on a block, you can also stake on one successor of that block. So you might be staked on a sequence of blocks that represent a single coherent claim about the correct history of the chain. A single stake suffices to commit you to that sequence of blocks.

In order to create a new rollup block, you must be a staker, and you must already be staked on the predecessor of the new block you’re creating. The stake requirement for block creation ensures that anyone who creates a new block has something to lose if that block is eventually rejected. 

The EthBridge keeps track of the current required stake amount. Normally this will equal the base stake amount, which is a parameter of the Arbitrum chain. But if the chain has been slow to make progress lately, the required stake will increase, as described in more detail below.

The rules for staking are as follows:

* If you’re not staked, you can stake on the latest confirmed rollup block. When doing this, you deposit with the EthBridge the current minimum stake amount.
* If you’re staked on a rollup block, you can also add your stake to any one successor of that block. (The EthBridge tracks the maximum rollup block number you’re staked on, and lets you add your stake to any successor of that block, updating your maximum to that successor.) This doesn’t require you to place a new stake.
  * A special case of adding your stake to a successor block is when you create a new rollup block as a successor to a block you’re already staked on.
* If you’re staked only on the latest confirmed block (and possibly earlier blocks), you or anyone else can ask to have your stake refunded. Your staked funds will be returned to you, and you will no longer be a staker. 
* If you lose a challenge, your stake is removed from all blocks and you forfeit your staked funds.

Notice that once you are staked on a rollup block, there is no way to unstake. You are committed to that block. Eventually one of two things will happen: that block will be confirmed, or you will lose your stake. The only way to get your stake back is to wait until all of the rollup blocks you are staked on are confirmed.

### Setting the current minimum stake amount

One detail we deferred earlier is how the current minimum stake amount is set. Normally, this is just equal to the base stake amount, which is a parameter of the Arbitrum chain. However, if the chain has been slow to make progress in confirming blocks, the stake requirement will escalate temporarily. Specifically, the base stake amount is multiplied by a factor that is exponential in the time since the deadline of the first unresolved node passed. This ensures that if malicious parties are placing false stakes to try to delay progress (despite the fact that they’re losing those stakes), the stake requirement goes up so that the cost of such a delay attack increases exponentially. As block resolution starts advancing again, the stake requirement will go back down.

## Rules for Confirming or Rejecting Rollup Blocks

The rules for resolving rollup blocks are fairly simple.

The first unresolved block can be confirmed if:

* the block’s predecessor is the latest confirmed block, and
* the block’s deadline has passed, and
* there is at least one staker, and 
* all stakers are staked on the block.

The first unresolved block can be rejected if:

* the block’s predecessor has been rejected, or
* all of the following are true:
  * the block’s deadline has passed, and
  * there is at least one staker, and
  * no staker is staked on the block.

A consequence of these rules is that once the first unresolved block’s deadline has passed (and assuming there is at least one staker staked on something other than the latest confirmed block), the only way the block can be unresolvable is if at least one staker is staked on it and at least one staker is staked on a different block with the same predecessor. If this happens, the two stakers are disagreeing about which block is correct. It’s time for a challenge, to resolve the disagreement.

## Challenges

Suppose the rollup chain looks like this:

![img](https://lh4.googleusercontent.com/kAZY9H73dqcHvboFDby9nrtbYZrbsHCYtE5X9NIZQsvcz58vV0WUWUq1xsYKzYWQSc1nPZ8W86LLX0lD3y-ctEaG2ISa2Wpz2pYxTzW09P1UvqSDuoqkHlGDYLLMTzLqX4rlP8Ca)

Blocks 93 and 95 are sibling blocks (they both have 92 as predecessor). Alice is staked on 93 and Bob is staked on 95. 

At this point we know that Alice and Bob disagree about the correctness of block 93, with Alice committed to 93 being correct and Bob committed to 93 being incorrect. (Bob is staked on 95, and 95 claims that 92 is the last correct block before it, which implies that 93 would be incorrect.) 

Whenever two stakers are staked on sibling blocks, and neither of those stakers is already in a challenge, anyone can start a challenge between the two. The rollup protocol will record the challenge and referee it, eventually declaring a winner and confiscating the loser’s stake. The loser will be removed as a staker.

The challenge is a game in which Alice and Bob alternate moves, with an Ethereum contract as the referee. Alice, the defender, moves first.

The game will operate in two phases: dissection, followed by one-step proof. Dissection will narrow down the size of the dispute until it is a dispute about just one instruction of execution. Then the one-step proof will determine who is right about that one instruction.

We’ll describe the dissection part of the protocol twice. First, we’ll give a simplified version which is easier to understand but less efficient. Then we’ll describe how the real version differs from the simplified one.

### Dissection Protocol: Simplified Version

Alice is defending the claim that starting with the state in the predecessor block, the state of the Virtual Machine can advance to the state specified in block A. Essentially she is claiming that the Virtual Machine can execute N instructions, and that that execution will consume M inbox messages and transform the hash of outputs from H’ to H.

Alice’s first move requires her to dissect her make claims about intermediate states between the beginning (0 instructions executed) and the end (N instructions executed). So we require Alice to divide her claim in half, and post the state at the half-way point, after N/2 instructions have been executed.

Now Alice has effectively bisected her N-step assertion into two (N/2)-step assertions. Bob has to point to one of those two half-size assertions and claim it is wrong.

At this point we’re effectively back in the original situation: Alice having made an assertion that Bob disagrees with. But we have cut the size of the assertion in half, from N to N/2. We can apply the same method again, with Alice bisecting and Bob choosing one of the halves, to reduce the size to N/4. And we can continue bisecting, so that after a logarithmic number of rounds Alice and Bob will be disagreeing about a single step of execution. That’s where the dissection phase of the protocol ends, and Alice must make a one-step proof which will be checked by the EthBridge.

### Why Dissection Correctly Identifies a Cheater

Before talking about the complexities of the real challenge protocol, let’s stop to understand why the simplified version of the protocol is correct. Here correctness means two things: (1) if Alice’s initial claim is correct, Alice can always win the challenge, and (2) if Alice’s initial claim is incorrect, Bob can always win the challenge.

To prove (1), observe that if Alice’s initial claim is correct, she can offer a truthful midpoint claim, and both of the implied half-size claims will be correct. So whichever half Bob objects to, Alice will again be in the position of defending a correct claim. At each stage of the protocol, Alice will be defending a correct claim. At the end, Alice will have a correct one-step claim to prove, so that claim will be provable and Alice can win the challenge.

To prove (2), observe that if Alice’s initial claim is incorrect, this can only be because her claimed endpoint after N steps is incorrect. Now when Alice offers her midpoint state claim, that midpoint claim is either correct or incorrect. If it’s incorrect, then Bob can challenge Alice’s first-half claim, which will be incorrect. If Alice’s midpoint state claim is correct, then her second-half claim must be incorrect, so Bob can challenge that. So whatever Alice does, Bob will be able to challenge an incorrect half-size claim. At each stage of the protocol, Bob can identify an incorrect claim to challenge. At the end, Alice will have an incorrect one-step claim to prove, which she will be unable to do, so Bob can win the challenge.

(If you’re a stickler for mathematical precision, it should be clear how these arguments can be turned into proofs by induction on N.)

### The Real Dissection Protocol

The real dissection protocol is conceptually similar to the simplified one described above, but with several changes that improve efficiency or deal with necessary corner cases. Here is a list of the differences.

**K-way dissection:** Rather than dividing a claim into two segments of size N/2, we divide it into K segments of size N/K. This requires posting K-1 intermediate claims, at points evenly spaced through the claimed execution. This reduces the number of rounds by a factor of log(K)/log(2).

**Answer a dissection with a dissection:** Rather than having each round of the protocol require two moves, where Alice dissects and Bob chooses a segment to challenge, we instead require Bob, in challenging a segment, to post his own claimed endpoint state for that segment (which must differ from Alice’s) as well as his own dissection of his version of the segment. Alice will then respond by identifying a subsegment, posting an alternative endpoint for that segment, and dissecting it. This reduces the number of moves in the game by an additional factor of 2, because the size is cut by a factor of K for every move, rather than for every two moves.

**Dissect by ArbGas, Not By Steps**: Rather than measuring the size of an interval by the number of instructions executed, we measure by the amount of ArbGas it consumes. This allows the protocol to precisely determine how long it will take a validator to check the correctness of a claim (because validator checking time is proportional to ArbGas used) which allows the protocol to more precisely set deadlines for rollup nodes. The drawback of this is that different instructions require different amounts of ArbGas, so we can no longer assume that a segment of execution can end on an exact N/K step boundary. The real protocol allows a claim to end at the first instruction boundary at or after its target endpoint; and the correctness argument accounts for the possibility that parties will lie about where the correct endpoint of a segment is.

**Deal With the Empty-Inbox Case**: The real AVM can’t always execute N ArbGas units without getting stuck. The machine might halt, or it might have to wait because its inbox is exhausted so it can’t go on until more messages arrive. So Bob must be allowed to respond to Alice’s claim of N units of execution by claiming that N steps are not possible. The real protocol thus allows any response (but not the initial claim) to claim a special end state that means essentially that the specified amount of execution is not possible under the current conditions.

**Time Limits:** Each player is given a time allowance. The total time a player uses for all of their moves must be less than the time allowance, or they lose the game. Think of the time allowance as being about a week. 

It should be clear that these changes don’t affect the basic correctness of the challenge protocol. They do, however, improve its efficiency and enable it to handle all of the cases that can come up in practice.

### Efficiency

The challenge protocol is designed so that the dispute can be resolved with a minimum of work required by the EthBridge in its role as referee. When it is Alice’s move, the EthBridge only needs to keep track of the time Alice uses, and ensure that her move does include 99 intermediate points as required. The EthBridge doesn’t need to pay attention to whether those claims are correct in any way; it only needs to know whether Alice’s move “has the right shape”. 

The only point where the EthBridge needs to evaluate a move “on the merits” is at the one-step proof, where it needs to look at Alice’s proof and determine whether the proof that was provided does indeed establish that the virtual machine moves from the before state to the claimed after state after one step of computation. We’ll discuss the details of one-step proofs below in the Arbitrum Virtual Machine section.