---
title: IRollupAdmin.sol Spec
---

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

### `pause()` (external)

Pause interaction with the rollup contract

### `resume()` (external)

Resume interaction with the rollup contract

### `setFacets(address newAdminFacet, address newUserFacet)` (external)

Set the addresses of rollup logic facets called

- `newAdminFacet`: address of logic that owner of rollup calls

- `newUserFacet`: ddress of logic that user of rollup calls

### `setValidator(address[] _validator, bool[] _val)` (external)

Set the addresses of the validator whitelist

It is expected that both arrays are same length, and validator at
position i corresponds to the value at position i

- `_validator`: addresses to set in the whitelist

- `_val`: value to set in the whitelist for corresponding address

### `setOwner(address newOwner)` (external)

Set a new owner address for the rollup

- `newOwner`: address of new rollup owner

### `setMinimumAssertionPeriod(uint256 newPeriod)` (external)

Set minimum assertion period for the rollup

- `newPeriod`: new minimum period for assertions

### `setConfirmPeriodBlocks(uint256 newConfirmPeriod)` (external)

Set number of blocks until a node is considered confirmed

- `newConfirmPeriod`: new number of blocks until a node is confirmed

### `setExtraChallengeTimeBlocks(uint256 newExtraTimeBlocks)` (external)

Set number of extra blocks after a challenge

- `newExtraTimeBlocks`: new number of blocks

### `setArbGasSpeedLimitPerBlock(uint256 newArbGasSpeedLimitPerBlock)` (external)

Set speed limit per block

- `newArbGasSpeedLimitPerBlock`: maximum arbgas to be used per block

### `setBaseStake(uint256 newBaseStake)` (external)

Set base stake required for an assertion

- `newBaseStake`: maximum arbgas to be used per block

### `setStakeToken(address newStakeToken)` (external)

Set the token used for stake, where address(0) == eth

Before changing the base stake token, you might need to change the
implementation of the Rollup User facet!

- `newStakeToken`: address of token used for staking

### `setSequencerInboxMaxDelayBlocks(uint256 newSequencerInboxMaxDelayBlocks)` (external)

Set max delay in blocks for sequencer inbox

- `newSequencerInboxMaxDelayBlocks`: max number of blocks

### `setSequencerInboxMaxDelaySeconds(uint256 newSequencerInboxMaxDelaySeconds)` (external)

Set max delay in seconds for sequencer inbox

- `newSequencerInboxMaxDelaySeconds`: max number of seconds

### `setChallengeExecutionBisectionDegree(uint256 newChallengeExecutionBisectionDegree)` (external)

Set execution bisection degree

- `newChallengeExecutionBisectionDegree`: execution bisection degree

### `updateWhitelistConsumers(address whitelist, address newWhitelist, address[] targets)` (external)

Updates a whitelist address for its consumers

setting the newWhitelist to address(0) disables it for consumers

- `whitelist`: old whitelist to be deprecated

- `newWhitelist`: new whitelist to be used

- `targets`: whitelist consumers to be triggered

### `setWhitelistEntries(address whitelist, address[] user, bool[] val)` (external)

Updates a whitelist's entries

user at position i will be assigned value i

- `whitelist`: whitelist to be updated

- `user`: users to be updated in the whitelist

- `val`: if user is or not allowed in the whitelist

### `setSequencer(address newSequencer)` (external)

Updates a sequencer address at the sequencer inbox

- `newSequencer`: new sequencer address to be used

### `upgradeBeacon(address beacon, address newImplementation)` (external)

Upgrades the implementation of a beacon controlled by the rollup

- `beacon`: address of beacon to be upgraded

- `newImplementation`: new address of implementation
