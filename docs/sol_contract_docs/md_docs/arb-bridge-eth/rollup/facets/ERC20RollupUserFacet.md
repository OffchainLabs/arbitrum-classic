---
title: ERC20RollupUserFacet.sol Spec
---

### `initialize(address _stakeToken)` (public)

### `newStake(uint256 tokenAmount)` (external)

Create a new stake

It is recomended to call stakeOnExistingNode after creating a new stake
so that a griefer doesn't remove your stake by immediately calling returnOldDeposit

- `tokenAmount`: the amount of tokens staked
  /

### `addToDeposit(address stakerAddress, uint256 tokenAmount)` (external)

Increase the amount staked tokens for the given staker

- `stakerAddress`: Address of the staker whose stake is increased

- `tokenAmount`: the amount of tokens staked
  /

### `withdrawStakerFunds(address payable destination) → uint256` (external)

Withdraw uncomitted funds owned by sender from the rollup chain

- `destination`: Address to transfer the withdrawn funds to
  /
