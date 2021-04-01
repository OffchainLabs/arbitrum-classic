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

import "../bridge/interfaces/IBridge.sol";

contract BridgeCreator {
    function createBridge(
        address rollup,
        address sequencer,
        uint256 sequencerDelayBlocks,
        uint256 sequencerDelaySeconds
    )
        external
        returns (
            Bridge,
            SequencerInbox,
            Inbox,
            RollupEventBridge,
            Outbox
        )
    {
        Bridge delayedBridge = new Bridge();

        SequencerInbox sequencerInbox =
            new SequencerInbox(
                IBridge(delayedBridge),
                sequencer,
                sequencerDelayBlocks,
                sequencerDelaySeconds
            );
        // TODO initialize sequencerInbox
        Inbox inbox = new Inbox(IBridge(delayedBridge));
        RollupEventBridge rollupEventBridge = new RollupEventBridge(address(delayedBridge), rollup);
        delayedBridge.setInbox(address(inbox), true);
        Outbox outbox = new Outbox(rollup, IBridge(delayedBridge));
        delayedBridge.transferOwnership(rollup);
        return (delayedBridge, sequencerInbox, inbox, rollupEventBridge, outbox);
    }
}
