---
title: Rollup.sol Spec
---

### `onlyOwner()`

### `initialize(bytes32 _machineHash, uint256 _confirmPeriodBlocks, uint256 _extraChallengeTimeBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, bytes _extraConfig, address[6] connectedContracts)` (external)

### `setOutbox(contract IOutbox _outbox)` (external)

Add a contract authorized to put messages into this rollup's inbox

- `_outbox`: Outbox contract to add

### `removeOldOutbox(address _outbox)` (external)

Disable an old outbox from interacting with the bridge

- `_outbox`: Outbox contract to remove

### `setInbox(address _inbox, bool _enabled)` (external)

Enable or disable an inbox contract

- `_inbox`: Inbox contract to add or remove

- `_enabled`: New status of inbox

### `upgradeImplementation(address _newRollup)` (external)

Switch over to a new implementation of the rollup

- `_newRollup`: New implementation contract

### `upgradeImplementationAndCall(address _newRollup, bytes _data)` (external)

Switch over to a new implementation of the rollup

- `_newRollup`: New implementation contract

- `_data`: Data to call the new rollup implementation with

### `pause()` (external)

Pause interaction with the rollup contract

### `resume()` (external)

Resume interaction with the rollup contract

### `beginTruncatingNodes(uint256 newLatestNodeCreated, uint256 maxItems)` (external)

Begin the process of trunacting the chain back to the given node

maxItems is used to make sure this doesn't exceed the max gas cost

- `newLatestNodeCreated`: Index that we want to be the latest unresolved node

- `maxItems`: Maximum number of items to eliminate to eliminate

### `continueTruncatingNodes(uint256 maxItems)` (public)

Continue the process of trunacting the chain back to the given node

maxItems is used to make sure this doesn't exceed the max gas cost

- `maxItems`: Maximum number of items to eliminate to eliminate

### `rejectNextNode(address stakerAddress)` (external)

Reject the next unresolved node

- `stakerAddress`: Example staker staked on sibling

### `confirmNextNode(bytes32 beforeSendAcc, bytes sendsData, uint256[] sendLengths, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount)` (external)

Confirm the next unresolved node

- `beforeSendAcc`: Accumulator of the AVM sends from the beginning of time up to the end of the previous confirmed node

- `sendsData`: Concatenated data of the sends included in the confirmed node

- `sendLengths`: Lengths of the included sends

- `afterSendCount`: Total number of AVM sends emitted from the beginning of time after this node is confirmed

- `afterLogAcc`: Accumulator of the AVM logs from the beginning of time up to the end of this node

- `afterLogCount`: Total number of AVM logs emitted from the beginning of time after this node is confirmed

### `newStake(uint256 tokenAmount)` (external)

Create a new stake

- `tokenAmount`: If staking in something other than eth, this is the amount of tokens staked, otherwise 0

### `withdrawStakerFunds(address payable destination) → uint256` (external)

Withdraw uncomitted funds owned by sender from the rollup chain

- `destination`: Address to transfer the withdrawn funds to

### `stakeOnExistingNode(uint256 nodeNum, bytes32 nodeHash)` (external)

Move stake onto an existing node

- `nodeNum`: Inbox of the node to move stake to. This must by a child of the node the staker is currently staked on

- `nodeHash`: Node hash of nodeNum (protects against reorgs)

### `stakeOnNewNode(bytes32 expectedNodeHash, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields, uint256 beforeProposedBlock, uint256 beforeInboxMaxCount)` (external)

Move stake onto a new node

- `expectedNodeHash`: The hash of the node being created (protects against reorgs)

- `assertionBytes32Fields`: Assertion data for creating

- `assertionIntFields`: Assertion data for creating

### `returnOldDeposit(address stakerAddress)` (external)

Refund a staker that is currently staked on or before the latest confirmed node

- `stakerAddress`: Address of the staker whose stake is refunded
  /

### `addToDeposit(address stakerAddress, uint256 tokenAmount)` (external)

Increase the amount staked for the given staker

- `stakerAddress`: Address of the staker whose stake is increased

- `tokenAmount`: If staking in something other than eth, this is the amount of tokens staked, otherwise 0
  /

### `reduceDeposit(uint256 target)` (external)

Reduce the amount staked for the sender

- `target`: Target amount of stake for the staker. If this is below the current minimum, it will be set to minimum instead
  /

### `createChallenge(address payable[2] stakers, uint256[2] nodeNums, bytes32[2] executionHashes, uint256[2] proposedTimes, uint256[2] maxMessageCounts)` (external)

Start a challenge between the given stakers over the node created by the first staker assuming that the two are staked on conflicting nodes

- `stakers`: Stakers engaged in the challenge. The first staker should be staked on the first node

- `nodeNums`: Nodes of the stakers engaged in the challenge. The first node should be the earliest and is the one challenged

- `executionHashes`: Challenge related data for the two nodes

- `proposedTimes`: Times that the two nodes were proposed

- `maxMessageCounts`: Total number of messages consumed by the two nodes
  /

### `completeChallenge(address winningStaker, address losingStaker)` (external)

Inform the rollup that the challenge between the given stakers is completed

completeChallenge isn't pausable since in flight challenges should be allowed to complete or else they could be forced to timeout

- `winningStaker`: Address of the winning staker

- `losingStaker`: Address of the losing staker
  /

### `removeZombie(uint256 zombieNum, uint256 maxNodes)` (external)

Remove the given zombie from nodes it is staked on, moving backwords from the latest node it is staked on

- `zombieNum`: Index of the zombie to remove

- `maxNodes`: Maximum number of nodes to remove the zombie from (to limit the cost of this transaction)
  /

### `removeOldZombies(uint256 startIndex)` (public)

Remove any zombies whose latest stake is earlier than the first unresolved node

- `startIndex`: Index in the zombie list to start removing zombies from (to limit the cost of this transaction)
  /

### `currentRequiredStake() → uint256` (public)

Calculate the current amount of funds required to place a new stake in the rollup

If the stake requirement get's too high, this function may stop reverting due to overflow, but
that only blocks operations that should be blocked anyway

**Returns**: The: current minimum stake requirement
/

### `countStakedZombies(contract INode node) → uint256` (public)

Calculate the number of zombies staked on the given node

This function could be uncallable if there are too many zombies. However,
removeZombie and removeOldZombies can be used to remove any zombies that exist
so that this will then be callable

- `node`: The node on which to count staked zombies

**Returns**: The: number of zombies staked on the node
/

### `requireUnresolvedExists()` (public)

Verify that there are some number of nodes still unresolved
/

### `requireUnresolved(uint256 nodeNum)` (public)
