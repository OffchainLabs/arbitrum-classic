---
id: ArbGas_Fees
title: Inside Arbitrum: ArbGas and Fees
sidebar_label: ArbGas and Fees
---
ArbGas is used by Arbitrum to track the cost of execution on an Arbitrum chain. It is similar in concept to Ethereum gas, in the sense that every Arbitrum Virtual Machine instruction has an ArbGas cost, and the cost of a computation is the sum of the ArbGas costs of the instructions in it.

ArbGas is not directly comparable to Ethereum gas. Arbitrum does not have a hard ArbGas limit, and in general an Arbitrum chain can consume many more ArbGas units per second of computation, compared to the number of Ethereum gas units in Ethereum's gas limit. Developers and users should think of ArbGas as much more plentiful and much cheaper per unit than Ethereum gas.

## Why ArbGas?

One of the design principles of the Arbitrum Virtual Machine (AVM) is that every instruction should take a predictable amount of time to validate, prove, and proof-check. As a corollary, we want a way to count or estimate the time required to validate any computation.

There are two reasons for this. First, we want to ensure that proof-checking has a predictable cost, so we can predict how much L1 gas is needed by the EthBridge and ensure that the EthBridge will never come anywhere close to the L1 gas limit.

Second, accurate estimation of validation time is important to maximize the throughput of a rollup chain, because it allows us to set the chain's speed limit safely.

## ArbGas in rollup blocks

Every rollup block includes an amount of ArbGas used by the computation so far, which implies an amount used since the predecessor rollup block. Like everything else in the rollup block, this value is only a claim made by the staker who created the block, and the block will be defeatable in a challenge if the claim is wrong.

Although the ArbGas value in a rollup block might not be correct, it can be used reliably as a limit on how much computation is required to validate the block. This is true because a validator who is checking the block can cut off their computation after that much ArbGas has been consumed; if that much ArbGas has been consumed without reaching the end of the rollup block, then the rollup block must be wrong and the checker can safely challenge it. For this reason, the rollup protocol can safely use the ArbGas claim in a rollup block, minus the amount in the predecessor block, as an upper bound on the time required to validate the rollup block’s correctness.

A rollup block can safely be challenged even if the ArbGas usage is the only aspect of the block that is false. When a claim is bisected, the claims will include (claimed) ArbGas usages, which must sum to the ArbGas usage of the parent claim. It follows that if the parent's ArbGas claim is false, at least one of the sub-claims must make a false ArbGas claim. So a challenger who knows that the parent's ArbGas claim is false will always be able to find a claim that has a false ArbGas claim.

Eventually the dispute will get down to a single AVM instruction, with a claim about that instruction's ArbGas usage. One-step proof verification checks that this claim is correct. So a false ArbGas claim in an assertion can be pursued all the way down to a single instruction with a false ArbGas claim that will be detected by the one-step proof verification in the EthBridge.

## ArbGas accounting in the AVM

The AVM architecture also does ArbGas accounting internally, using a machine register called ArbGasRemaining, which is a 256-bit unsigned integer that behaves as follows:

* The register is initially set to MaxUint256.
* Immediately before any instruction executes, that instruction's ArbGas cost is subtracted from the register. If this would make the register's value less than zero, an error is generated and the register's value is set to MaxUint256. (The error causes a control transfer as specified in the AVM specification.)
* A special instruction can be used to read the register's value.
* Another special instruction can be used to set the register to any desired value.

This mechanism allows ArbOS to control and account for the ArbGas usage of application code. ArbOS can limit an application call's use of ArbGas to N units by setting the register to N before calling the application, then catching the out-of-ArbGas error if it is generated. At the beginning of the runtime's error-handler, ArbOS would read the ArbGasRemaining register, then set the register to MaxUint256 to ensure that the error-handler could not run out of ArbGas. If the read of the register gave a value close to MaxInt256, then it must be the case that the application generated an out-of-ArbGas error. (It could be the case that the application generates a different error while a small amount of ArbGas remains, then an out-of-ArbGas error occurs at the very beginning of the error-handler. In this case, the second error would set the ArbGasRemaining to MaxInt256 and throw control back to the beginning of the error-handler, causing the error-handler to conclude that an out-of-ArbGas error was caused by the application. This is a reasonable behavior which we will consider to be correct.)

If the application code returns control to the runtime without generating an out-of-ArbGas error, the runtime can read the ArbGasRemaining register and subtract to determine how much ArbGas the application call used. This can be charged to the application's account.

The runtime can safely ignore the ArbGas accounting mechanism. If the special instructions are never used, the register will be set to MaxInt256, and will decrease but in practice will never get to zero, so no error will ever be generated.

The translator that turns EVM code into equivalent AVM code will never generate the instruction that sets the ArbGasRemaining register, so untrusted code cannot manipulate its own gas allocation.

## The Speed Limit

The security of Arbitrum chains depends on the assumption that when one validator creates a rollup block, other validators will check it and issue a challenge if it is wrong. This requires that the other validators have the time and resources to check each rollup block in time to issue a timely challenge. The Arbitrum protocol takes this into account in setting deadlines for rollup blocks.

This sets an effective speed limit on execution of an Arbitrum VM: in the long run the VM cannot make progress faster than a validator can emulate its execution. If rollup blocks are published at a rate faster than the speed limit, their deadlines will get farther and farther in the future. Due to the limit, enforced by the rollup protocol contracts, on how far in the future a deadline can be, this will eventually cause new rollup blocks to be slowed down, thereby enforcing the effective speed limit.

Being able to set the speed limit accurately depends on being able to estimate the time required to validate an AVM computation, with some accuracy. Any uncertainty in estimating validation time will force us to set the speed limit lower, to be safe. And we do not want to set the speed limit lower, so we try to enable accurate estimation.

## Fees

User transactions pay fees, to cover the cost of operating the chain. These fees are assessed and collected by ArbOS at L2. They are denominated in ETH. Some of the fees are immediately paid to the aggregator who submitted the transaction (if there is one, and subject to some limitations discussed below). The rest go into a network fee pool that is used to pay service providers who help to enable the chain’s secure operation, such as validators.

Fees are charged for four resources that a transaction can use:

* *L2 tx*: a base fee for each L2 transaction, to cover the cost of servicing a transaction
* *L1 calldata*: a fee per units of L1 calldata directly attributable to the transaction (each non-zero byte of calldata is 16 units, and each zero byte of calldata is 4 units)
* *computation*: a fee per unit of ArbGas used by the transaction
* *storage*: a fee per location of AVM storage, based on the net increase in EVM storage due to the transaction

Each of these four resources has a price, which may vary over time. The resource prices, which are denominated in ETH (more precisely, in wei), are set as follows:

### Prices for L2 tx and L1 calldata

The prices of the first two resources, L2 tx and L1 calldata, depend on the L1 gas price. ArbOS can’t directly see the L1 gas price, so it estimates the L1 gas prices, as a weighted average of recent gas prices actually paid by aggregators on a specific list of addresses. ArbOS also keeps a running average of (1 / batchSize) for recent transaction batches. 

* The base price of an L2 tx is CP/B, where C is the L1 cost required for an aggregator to submit an empty batch, P is the estimated L1 gas price, and (1/B) is the average (1/batchSize). 
* The base price of a unit of L1 calldata is just the estimated L1 gas price.

These base prices ensure that the amount collected is equal to the aggregator’s cost of submitting the transaction in a batch, assuming a typical batch size.

If the transaction was submitted by an aggregator, ArbOS collects these base fees for L2 tx and L1 calldata, and credits that amount immediately to the aggregator. ArbOS also adds a 15% markup and deposits those funds into the network fee account, to help cover overhead and other chain costs. If the transaction was not submitted by an aggregator, ArbOS collects only the 15% portion and credits that to the network fee account.

In order for an aggregator to be reimbursed for submitting a transaction, three things must be true:

1. The transaction’s nonce is correct. [This prevents an aggregator from resubmitting a transaction to collect multiple reimbursements for it.]
2. The transaction’s sender has ETH in its L2 account to pay the fees.
3. The aggregator is listed as the sender’s “preferred aggregator”. (ArbOS records a preferred aggregator for each account, with the default being a specific aggregator address, which is operated by Offchain Labs on the flagship Arbitrum chain.) [This prevents an aggregator from front-running another aggregator’s batches to steal its reimbursements.]

If these conditions are not all met, the transaction is treated as not submitted by an aggregator so only the network fee portion of the fee is collected.

### Price for storage

Transactions are charged based on the net increase they cause in the total amount of AVM contract storage that exists. Decreases in contract storage do not receive a credit. Each storage location costs 2000 times the estimated L1 gas price. This means that storage costs about 10% as much as it does on the L1 chain.

### Price for ArbGas

ArbGas pricing depends on a minimum price, and a congestion pricing mechanism. 

The minimum ArbGas price is set equal to the estimated L1 gas price divided by 10,000. The price of ArbGas will never go lower than this.

The price will rise above the minimum if the chain is starting to get congested. The idea is similar to Ethereum, to deal with a risk of overload by raising the price of gas enough that demand will meet supply. The mechanism is inspired by Ethereum’s proposed EIP-1559, which uses a base price that, at the beginning of each block, is multiplied by a factor between 7/8 and 9/8, depending on how busy the chain seems to be. When the demand seems to be exceeding supply, the factor will be more than 1; when supply exceeds demand, the factor will be less than 1. (But the price never goes below the minimum.)

The automatic adjustment mechanism depends on an “ArbGas pool” that is tracked by ArbOS. The ArbGas pool has a maximum capacity that is equal to 60 seconds of full-speed computation at the chain’s speed limit. ArbGas used by transactions is subtracted from the gas pool, and at the beginning of each new Ethereum block, ArbGas is added to the pool corresponding to full-speed execution for the number of timestamp seconds since the last block (subject to the maximum pool capacity).

After adding the new gas to the pool, if the new gas pool size is G, the current ArbGas price is multiplied by (1350S - G) / (1200S) where S is the ArbGas speed limit of the chain. This quantity will be 7/8 when G = 60S (the maximum gas pool size) and it will be 9/8 when G = 0.

#### The congestion limit

The gas pool is also used to limit the amount of gas available to transactions at a particular timestamp. In particular, if a transaction could require more gas than is available in the gas pool, that transaction is rejected without being executed, returning a transaction receipt indicating that the transaction was dropped due to congestion. This prevents the chain from building up a backlog that is very long--if transactions are being submitted consistently faster than they can be validated, eventually the gas pool will become empty and transactions will be dropped. Meanwhile the transaction price will be escalating, so that the price mechanism will be able to bring supply and demand back into alignment.

#### ArbGas accounting and the second-price auction

As on Ethereum, Arbitrum transactions submit a maximum gas amount (here, “maxgas” for short) and a gas price bid (here, “gasbid” for short). A transaction will use up to its maxgas amount of gas, or will revert if more gas would be needed. 

On Ethereum, the gas price paid by a transaction is equal to its gasbid. Arbitrum, by contrast, treats the gasbid as the maximum amount the transaction is willing to pay for gas. The actual price paid is equal to the current Arbitrum ArbGas price, whatever that is, as long as it is less than or equal to the transaction’s gasbid. If the transaction’s gasbid is less than the current Arbitrum gas price, the transaction is dropped and a transaction receipt issued, saying that the transaction’s gasbid was too low.

So Arbitrum transactions shouldn’t try to “game” their gasbid by trying to match it too closely to the current ArbGas price. Instead, transactions should set their gasbid equal to the maximum price they’re willing to pay for ArbGas, with the caveat that the sender must have at least gasbid*maxgas in its L2 ETH account.