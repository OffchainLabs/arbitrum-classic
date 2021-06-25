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
import "../bridge/Inbox.sol";
import "../bridge/Outbox.sol";
import "../bridge/SequencerInbox.sol";
import "../rollup/RollupEventBridge.sol";
import "../rollup/Rollup.sol";
import "../rollup/NodeFactory.sol";

import "../rollup/Rollup.sol";
import "../bridge/interfaces/IBridge.sol";

import "../rollup/RollupLib.sol";
import "../rollup/facets/RollupUser.sol";
import "../rollup/facets/RollupAdmin.sol";

import "../libraries/Whitelist.sol";

contract RollupCreatorNoProxy {
    event RollupCreated(address rollupAddress, Inbox inbox);

    constructor(
        address _challengeFactory,
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
        bytes memory _extraConfig
    ) public {
        RollupLib.Config memory config =
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
            );

        createRollupNoProxy(config, _challengeFactory);
        selfdestruct(msg.sender);
    }

    struct CreateRollupFrame {
        Bridge delayedBridge;
        SequencerInbox sequencerInbox;
        Inbox inbox;
        RollupEventBridge rollupEventBridge;
        Outbox outbox;
        address rollup;
    }

    struct CreateBridgeFrame {
        Bridge delayedBridge;
        SequencerInbox sequencerInbox;
        Inbox inbox;
        RollupEventBridge rollupEventBridge;
        Outbox outbox;
        Whitelist whitelist;
    }

    function createBridge(address rollup, address sequencer)
        private
        returns (
            Bridge,
            SequencerInbox,
            Inbox,
            RollupEventBridge,
            Outbox
        )
    {
        CreateBridgeFrame memory frame;
        {
            frame.delayedBridge = new Bridge();
            frame.sequencerInbox = new SequencerInbox();
            frame.inbox = new Inbox();
            frame.rollupEventBridge = new RollupEventBridge();
            frame.outbox = new Outbox();
            // frame.whitelist = new Whitelist();
        }

        frame.delayedBridge.initialize();
        frame.sequencerInbox.initialize(IBridge(frame.delayedBridge), sequencer, rollup);
        frame.inbox.initialize(IBridge(frame.delayedBridge), address(0));
        frame.rollupEventBridge.initialize(address(frame.delayedBridge), rollup);
        frame.outbox.initialize(rollup, IBridge(frame.delayedBridge));

        frame.delayedBridge.setInbox(address(frame.inbox), true);
        frame.delayedBridge.transferOwnership(rollup);

        // frame.whitelist.setOwner(rollup);

        return (
            frame.delayedBridge,
            frame.sequencerInbox,
            frame.inbox,
            frame.rollupEventBridge,
            frame.outbox
        );
    }

    function createRollupNoProxy(RollupLib.Config memory config, address challengeFactory)
        private
        returns (address)
    {
        CreateRollupFrame memory frame;
        frame.rollup = address(new Rollup());
        (
            frame.delayedBridge,
            frame.sequencerInbox,
            frame.inbox,
            frame.rollupEventBridge,
            frame.outbox
        ) = createBridge(frame.rollup, config.sequencer);

        Rollup(payable(frame.rollup)).initialize(
            config.machineHash,
            [
                config.confirmPeriodBlocks,
                config.extraChallengeTimeBlocks,
                config.arbGasSpeedLimitPerBlock,
                config.baseStake
            ],
            config.stakeToken,
            config.owner,
            config.extraConfig,
            [
                address(frame.delayedBridge),
                address(frame.sequencerInbox),
                address(frame.outbox),
                address(frame.rollupEventBridge),
                challengeFactory,
                address(new NodeFactory())
            ],
            [address(new RollupAdminFacet()), address(new RollupUserFacet())],
            [config.sequencerDelayBlocks, config.sequencerDelaySeconds]
        );
        emit RollupCreated(frame.rollup, frame.inbox);
        return frame.rollup;
    }
}
