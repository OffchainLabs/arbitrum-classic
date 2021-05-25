---
title: RollupCreator.sol Spec
---

### `setTemplates(contract BridgeCreator _bridgeCreator, contract ICloneable _rollupTemplate, address _challengeFactory, address _nodeFactory)` (external)

### `createRollup(bytes32 _machineHash, uint256 _confirmPeriodBlocks, uint256 _extraChallengeTimeBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, address _sequencer, uint256 _sequencerDelayBlocks, uint256 _sequencerDelaySeconds, bytes _extraConfig) → contract IRollup` (external)

### `RollupCreated(address rollupAddress, address inboxAddress, address adminProxy)`

### `TemplatesUpdated()`
