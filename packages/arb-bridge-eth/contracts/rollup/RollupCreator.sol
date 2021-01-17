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

import "@openzeppelin/contracts/proxy/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/TransparentUpgradeableProxy.sol";
import "../bridge/Bridge.sol";
import "../bridge/Inbox.sol";
import "../bridge/Outbox.sol";

import "./IRollup.sol";
import "../bridge/interfaces/IBridge.sol";

contract RollupCreator is Ownable {
    event RollupCreated(address rollupAddress);

    address rollupTemplate;
    address challengeFactory;
    address nodeFactory;

    function setTemplates(
        address _rollupTemplate,
        address _challengeFactory,
        address _nodeFactory
    ) external onlyOwner {
        rollupTemplate = _rollupTemplate;
        challengeFactory = _challengeFactory;
        nodeFactory = _nodeFactory;
    }

    struct CreateRollupFrame {
        ProxyAdmin admin;
        Bridge bridge;
        Inbox inbox;
        Outbox outbox;
        TransparentUpgradeableProxy rollup;
    }

    // After this setup:
    // Rollup should be the owner of bridge
    // Rollup should be the owner of it's upgrade admin
    // Bridge should have a single inbox and outbox
    function createRollup(
        bytes32 _machineHash,
        uint256 _challengePeriodBlocks,
        uint256 _arbGasSpeedLimitPerBlock,
        uint256 _baseStake,
        address _stakeToken,
        address _owner,
        bytes memory _extraConfig
    ) public returns (IRollup) {
        CreateRollupFrame memory frame;
        frame.admin = new ProxyAdmin();
        frame.rollup = new TransparentUpgradeableProxy(rollupTemplate, address(frame.admin), "");

        frame.bridge = new Bridge();
        frame.inbox = new Inbox(IBridge(frame.bridge));
        frame.bridge.setInbox(address(frame.inbox), true);
        frame.outbox = new Outbox(address(frame.rollup), IBridge(frame.bridge));
        frame.bridge.setOutbox(address(frame.outbox), true);

        frame.bridge.transferOwnership(address(frame.rollup));
        frame.admin.transferOwnership(address(frame.rollup));
        IRollup(address(frame.rollup)).initialize(
            address(frame.outbox),
            _machineHash,
            _challengePeriodBlocks,
            _arbGasSpeedLimitPerBlock,
            _baseStake,
            _stakeToken,
            _owner,
            address(frame.bridge),
            challengeFactory,
            nodeFactory,
            _extraConfig,
            address(frame.admin)
        );
        emit RollupCreated(address(frame.rollup));
        return IRollup(address(frame.rollup));
    }
}
