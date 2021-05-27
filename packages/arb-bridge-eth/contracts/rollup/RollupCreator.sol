// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2021, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

pragma solidity ^0.6.11;

import "../bridge/Bridge.sol";
import "../bridge/SequencerInbox.sol";
import "../bridge/Inbox.sol";
import "../bridge/Outbox.sol";
import "./RollupEventBridge.sol";
import "./BridgeCreator.sol";

import "@openzeppelin/contracts/proxy/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

import "./Rollup.sol";
import "./facets/RollupUser.sol";
import "./facets/RollupAdmin.sol";
import "../bridge/interfaces/IBridge.sol";

import "./RollupLib.sol";
import "../libraries/CloneFactory.sol";
import "../libraries/ICloneable.sol";

contract RollupCreator is Ownable, CloneFactory {
    event RollupCreated(address indexed rollupAddress, address inboxAddress, address adminProxy);
    event TemplatesUpdated();

    BridgeCreator public bridgeCreator;
    ICloneable public rollupTemplate;
    address public challengeFactory;
    address public nodeFactory;
    address public rollupAdminFacet;
    address public rollupUserFacet;

    constructor() public Ownable() {}

    function setTemplates(
        BridgeCreator _bridgeCreator,
        ICloneable _rollupTemplate,
        address _challengeFactory,
        address _nodeFactory,
        address _rollupAdminFacet,
        address _rollupUserFacet
    ) external onlyOwner {
        bridgeCreator = _bridgeCreator;
        rollupTemplate = _rollupTemplate;
        challengeFactory = _challengeFactory;
        nodeFactory = _nodeFactory;
        rollupAdminFacet = _rollupAdminFacet;
        rollupUserFacet = _rollupUserFacet;
        emit TemplatesUpdated();
    }

    function createRollup(
        bytes32 _machineHash,
        uint256 _confirmPeriodBlocks,
        uint256 _extraChallengeTimeBlocks,
        uint256 _arbGasSpeedLimitPerBlock,
        uint256 _baseStake,
        address _stakeToken,
        address _owner,
        address _sequencer,
        uint256 _sequencerDelayBlocks,
        uint256 _sequencerDelaySeconds,
        bytes calldata _extraConfig
    ) external returns (address) {
        return
            createRollup(
                RollupLib.Config(
                    _machineHash,
                    _confirmPeriodBlocks,
                    _extraChallengeTimeBlocks,
                    _arbGasSpeedLimitPerBlock,
                    _baseStake,
                    _stakeToken,
                    _owner,
                    _sequencer,
                    _sequencerDelayBlocks,
                    _sequencerDelaySeconds,
                    _extraConfig
                )
            );
    }

    struct CreateRollupFrame {
        ProxyAdmin admin;
        Bridge delayedBridge;
        SequencerInbox sequencerInbox;
        Inbox inbox;
        RollupEventBridge rollupEventBridge;
        Outbox outbox;
        address rollup;
    }

    // After this setup:
    // Rollup should be the owner of bridge
    // RollupOwner should be the owner of Rollup's ProxyAdmin
    // RollupOwner should be the owner of Rollup
    // Bridge should have a single inbox and outbox
    function createRollup(RollupLib.Config memory config) private returns (address) {
        CreateRollupFrame memory frame;
        frame.admin = new ProxyAdmin();
        frame.rollup = address(
            new TransparentUpgradeableProxy(address(rollupTemplate), address(frame.admin), "")
        );

        (
            frame.delayedBridge,
            frame.sequencerInbox,
            frame.inbox,
            frame.rollupEventBridge,
            frame.outbox
        ) = bridgeCreator.createBridge(
            address(frame.admin),
            frame.rollup,
            config.sequencer,
            config.sequencerDelayBlocks,
            config.sequencerDelaySeconds
        );

        frame.admin.transferOwnership(config.owner);
        Rollup(payable(frame.rollup)).initialize(
            config.machineHash,
            config.confirmPeriodBlocks,
            config.extraChallengeTimeBlocks,
            config.arbGasSpeedLimitPerBlock,
            config.baseStake,
            config.stakeToken,
            config.owner,
            config.extraConfig,
            [
                address(frame.delayedBridge),
                address(frame.sequencerInbox),
                address(frame.outbox),
                address(frame.rollupEventBridge),
                challengeFactory,
                nodeFactory
            ],
            [rollupAdminFacet, rollupUserFacet]
        );
        emit RollupCreated(frame.rollup, address(frame.inbox), address(frame.admin));
        return frame.rollup;
    }
}
