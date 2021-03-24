---
id: EVM_Compatibility
title: Inside Arbitrum: EVM Compatibility
sidebar_label: EVM Compatibility
---
ArbOS provides functionality to emulate the execution of an Ethereum-compatible chain. It tracks accounts, executes contract code, and handles the details of creating and running EVM code. Clients can submit EVM transactions, including the ones that deploy contracts, and ArbOS ensures that the submitted transactions run in a compatible manner.

## The account table

ArbOS maintains an account table, which keeps track of the state of every account in the emulated Ethereum chain. The table entry for an account contains the account’s balance, its nonce, its code and storage (if it is a contract), and some other information associated with Arbitrum-specific features. An account’s entry in the table is initialized the first time anything happens in the account on the Arbitrum chain.

## Translating EVM code to run on AVM

EVM code can’t run directly on the AVM architecture, so ArbOS has to translate EVM code into equivalent AVM code in order to run it. This is done within ArbOS to ensure that it is trustless.

(Some old versions of Arbitrum used a separate compiler to translate EVM to AVM code, but that had significant disadvantages in both security and functionality, so we switched to built-in translation in ArbOS.)

ArbOS takes an EVM contract’s program and translates it into an AVM code segment that has equivalent functionality. Some instructions can be translated directly; for example an EVM add instruction translates into a single AVM add instruction. Other instructions translate into a call to a library provided by ArbOS. For example, the EVM *CREATE2* instruction, which creates a new contract whose address is computed in a special way, translates into a call to an ArbOS function called *evmOp_create2*.

To deploy a new contract, ArbOS takes the submitted EVM constructor code, translates it into AVM code for the constructor, and runs that code. If it completes successfully, ArbOS takes the return data, interprets it as EVM code, translates that into AVM code, and then installs that AVM code into the account table as the code for the now-deployed contract. Future calls to the contract will jump to the beginning of that AVM code.

## More EVM emulation details

When an EVM transaction is running, ArbOS keeps an EVM Call Strack, a stack of EVM Call Frames which represent the nested calls inside of the transaction. Each EVM Call Frame records the data for one level of call. When an inner call returns or reverts, ArbOS cleans up its call frame, and either propagates the effects of the call (if the call returned) or discards them (if the call reverted).

The EVM call family of instructions (call, delegatecall, etc.) invoke ArbOS library functions that create new EVM Call Frames according to the EVM-specified behavior of the call type that was made, including propagation of calldata and gas. Any gas left over after a call is returned to the caller, as on Ethereum.

Certain errors in executing EVM, such as stack underflows, will trigger a corresponding error during AVM emulation. ArbOS’s error handler detects that the error occurred while emulating a certain EVM call, and reverts that call accordingly, returning control to its caller or providing a transaction receipt as appropriate. ArbOS distinguishes out-of-gas errors from other errors by looking at the ArbGasRemaining register, which is automatically set to MaxUint256 if an out-of-gas error occurred.

The result of these mechanisms is that EVM code can run compatibly on Arbitrum.

There are three main differences between EVM execution on Arbitrum compared to Ethereum. 

First, the two EVM instructions DIFFICULTY and COINBASE, which don’t make sense on an L2 chain, return fixed constant values. 

Second, the EVM instruction BLOCKHASH returns a pseudorandom value that is a digest of the chain’s history, but not the same hash value that Ethereum would return at the same block number. 

Third, Arbitrum uses the ArbGas system so everything related to gas is denominated in ArbGas, including the gas costs of operations, and the results of gas-related instructions like GASLIMIT and GASPRICE.

## Deploying EVM contracts

On Ethereum, an EVM contract is deployed by sending a transaction to the null account, with calldata consisting of the contract’s constructor. Ethereum runs the constructor, and if the constructor succeeds, its return data is set up as the code for the new contract. 

Arbitrum uses a similar pattern. To deploy a new EVM contract, ArbOS takes the submitted EVM constructor code, translates it into AVM code for the constructor, and runs that AVM code. If it completes successfully, ArbOS takes the return data, interprets it as EVM code, translates that into AVM code, and then installs that AVM code into the account table as the code for the now-deployed contract. Future calls to the contract will jump to the beginning of that AVM code.

## Transaction receipts

When an EVM transaction finishes running, whether or not it succeeded, ArbOS issues a transaction receipt, using the AVM *log* instruction. The receipt is detected by Arbitrum nodes, which can use it to provide results to users according to the standard Ethereum RPC API.