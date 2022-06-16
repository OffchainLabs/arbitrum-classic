---
id: Chain_parameters
title: Setting parameters for your Arbitrum Chain
sidebar_label: Setting ArbChain parameters
---

When you create an ArbChain, you have to specify some parameters for the chain.
This document describes these parameters and recommends how to set them.

Note: If you run the validator (`arb-node --node.type=validator`) without any options regarding parameters, you will get the recommended default parameter values.

The parameters are:

- stake requirement: This specifies how much currency a validator must deposit as a stake.
  Making this bigger strengthens deterrence against malicious behavior by validators,
  but it also increases the amount of capital that validators have to lock up.
- grace period: This is the time period allowed for one validator to challenge an assertion made by another validator.
  Making this bigger makes certain attacks against the system more difficult,
  but it also slows down confirmation of transactions.
- speed limit: This controls how fast computations on the chain can go, to make sure that every validator can keep up with
  computations. This is essentially the speed of the slowest validator you expect to have. The speed is scaled so that 1.0
  is the speed of a typical developer laptop. For example, if you are using half of a laptop's capacity to run the
  validator, you should set the speed limit to 0.5.
- max assertion size: This is the largest amount of computation that you want to include in one assertion. If the chain is
  running at 100% capacity, the updated state of the chain will be posted to the Ethereum chain this often.

## Recommended parameters

For executing in a production-like setting, we suggest the following parameters:

- stake requirement: 1 Eth, or 2% of total value in chain, whichever is more
- grace period: 360 minutes
- speed limit: 1.0
- max assertion size: 50 seconds

(We say "production-like setting" because we currently **strongly recommend against** using Arbitrum Rollup in a true
production setting on the main chain.)

If you're debugging your dApp, you care more about fast turnaround than security, so we suggest the following parameters:

- stake requirement: 0.1 Eth
- grace period: 10 minutes
- speed limit: 0.2
- max assertion size: 15 seconds

The presets on the chain launcher page will follow these recommendations.

**To learn more about how to choose parameters and the rationale behind the default parameters, see our Medium post on [optimizing challenge periods](https://medium.com/offchainlabs/optimizing-challenge-periods-in-rollup-b61378c87277).**
