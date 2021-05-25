---
title: Rollup.sol Spec
---

### `newStake()` (external)

Create a new stake

It is recomended to call stakeOnExistingNode after creating a new stake
so that a griefer doesn't remove your stake by immediately calling returnOldDeposit
/

### `addToDeposit(address stakerAddress)` (external)

Increase the amount staked eth for the given staker

- `stakerAddress`: Address of the staker whose stake is increased
  /

### `withdrawStakerFunds(address payable destination) → uint256` (external)

Withdraw uncomitted funds owned by sender from the rollup chain

- `destination`: Address to transfer the withdrawn funds to
  /
