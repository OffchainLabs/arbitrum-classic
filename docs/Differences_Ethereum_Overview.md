---
id: Differences_Overview
sidebar_label: Overview
title: Overview of Differences with Ethereum
custom_edit_url: https://github.com/OffchainLabs/arbitrum/edit/master/docs/Differences_Ethereum_Overview.md
---

Arbitrum rollups aim to mantain compatibility with Ethereum. Smart contracts are compatible on the bytecode level, but there are certain aspects of the system that work differently to the EVM.

Some of the opcodes have slightly different behaviours, as seen in [Solidity Support](Solidity_Support.md).  
Concepts such as [Time in Arbitrum](Time_in_Arbitrum.md) and [Gas](ArbGas.md) play out differently in Layer 2.  
Other differences are cool [Special Features](Special_Features.md) we squeezed in.

## Ethereum Accounts

### Nonces

Every transaction submitted to Arbitrum will burn a nonce, except if the transaction is formatted incorrectly or does not have the expected nonce.

### L1 to L2 Deposits

Ether can be depositted using two methods: [retryable transactions](L1_L2_Messages.md) or L2 funded by L1 transactions. For end users these behave similarly, but have subtle differences.

When depositing funds it is possible to send ether into a contract address without executing its fallback function - a scenario similar to when contracts _self destruct_, sending funds to a contract address.

## JSON RPC API

The API for Arbitrum aims to be a superset of the [eth spec](https://eth.wiki/json-rpc/API). When interacting with it you can expect all the usual fields, as well as some extra ones used to surface information unique to Arbitrum Rollups.

### Transaction Receipts

Transaction receipts contain the following extra fields

#### L1 Block Number

The Layer 1 block number for the transaction, as specified in [Time in Arbitrum](Time_In_Arbitrum.md).

#### Fee Stats

An object summarizing fee charges for the current transaction. It includes the units used, price paid, and price per unit.

The arbgas charges will vary depending on how a user interacts with Arbitrum, but the following table can be used as a general reference:

| Key           | Cost source                              | Meaning                                                                                                                                   |
| ------------- | ---------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- |
| l1Transaction | Fixed cost of inclusion to Layer 1 inbox | There is a fixed cost of including your transaction into the Layer 1 inbox contract. This is amortised by batching transactions together. |
| l1Calldata    | Layer 1 transaction calldata             | The calldata included in each Layer 1 transaction has a cost associated to it. Aggregators are reimbursed for their costs.                |
| l2Storage     | Layer 2 storage                          | Users are charged whenever a storage slot is written to with a non-zero value.                                                            |
| l2Computation | Layer 2 computation                      | Users are charged per unit of computation used (measured in arbgas).                                                                      |

For more in-depth explanations head over to the [Gas](ArbGas.md) docs page, or the [Inside Arbitrum](Inside_Arbitrum.md#arbgas-and-fees) section on arbgas.

#### Return Data

This includes the data from a smart contract return or the revert reason if you hit an EVM revert statement.

#### Return Code

| Return Code | Meaning                                                                                                                              |
| ----------- | ------------------------------------------------------------------------------------------------------------------------------------ |
| 0           | Transaction success                                                                                                                  |
| 1           | EVM revert                                                                                                                           |
| 2           | Arbitrum is too congested to process your transaction                                                                                |
| 3           | Not enough balance to pay for maxGas at gasPrice                                                                                     |
| 4           | Not enough balance for execution                                                                                                     |
| 5           | Wrong nonce used in transaction                                                                                                      |
| 6           | Transaction was not formatted correctly                                                                                              |
| 7           | Cannot deploy to specified address ( ** defensive code that should never be triggered ** )                                           |
| 8           | Exceeded transaction gas limit                                                                                                       |
| 9           | Amount of ArbGas provided for the tx is less than the amount required to cover L1 costs (the base tx charge plus L1 calldata charge) |
| 10          | Transaction is below the minimum required arbgas                                                                                     |
| 11          | Transaction set an arbgas price that was too low                                                                                     |
| 255         | Unknown failure                                                                                                                      |
