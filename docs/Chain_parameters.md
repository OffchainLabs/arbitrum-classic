---
id: Chain_parameters
title: Setting parameters for your Arbitrum Chain
sidebar_label: Setting ArbChain parameters
---

When you create an ArbChain, you have to specify some parameters for the chain.
This document describes these parameters and recommends how to set them.

Note: If you run the validator (`arb-validator` or `arb_deploy.py`) without any options regarding parameters, you will get the recommended default parameter values.

The parameters are:

-   stake requirement: This specifies how many Eth a validator must deposit as a stake.
    Making this bigger strengthens deterrence against malicious behavior by validators,
    but it also increases the amount of capital that validators have to lock up.
-   grace period: This is the time period allowed for one validator to challenge an assertion made by another validator.
    Making this bigger makes certain attacks against the system more difficult,
    but it also slows down confirmation of transactions.
-   max execution steps: This is the maximum number of AVM instructions that can be executed in a single assertion.
    Making this bigger increases the efficiency of the protocol (with diminishing returns as it gets really big),
    but it can also slow down confirmation of transactions.
-   speed limit: This is the maximum amount of ArbGas consumption that can occur per second of real time.
    Making this bigger lets the contracts in your ArbChain run faster,
    but making it too big risks overwhelming validators who can't keep up.

## Recommended parameters for production testing

For executing in a production-like setting, we suggest the following parameters:

-   stake requirement: 2% of total value at stake in your ArbChain
-   grace period: 180 minutes
-   max execution steps: 1,000,000,000
-   speed limit: 100,000,000 ArbGas per second

(We say "production-like setting" because we do not currently recommend use of Arbitrum Rollup in a true
production setting on the main chain.)

## Recommended parameters for debugging

If you're debugging your dapp, you care more about fast turnaround than security, so we suggest the following parameters:

-   stake requirement: 0.1 Eth
-   grace period: 10 minutes
-   max execution steps: 1,000,000,000
-   speed limit: 20,000,000 ArbGas per second
