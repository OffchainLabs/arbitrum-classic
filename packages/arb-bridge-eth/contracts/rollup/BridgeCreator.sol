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
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/proxy/ProxyAdmin.sol";
import "../libraries/Whitelist.sol";

contract BridgeCreator is Ownable {
    Bridge public delayedBridgeTemplate;
    SequencerInbox public sequencerInboxTemplate;
    Inbox public inboxTemplate;
    RollupEventBridge public rollupEventBridgeTemplate;
    Outbox public outboxTemplate;

    event TemplatesUpdated();

    constructor() public Ownable() {
        delayedBridgeTemplate = new Bridge();
        sequencerInboxTemplate = new SequencerInbox();
        inboxTemplate = new Inbox();
        rollupEventBridgeTemplate = new RollupEventBridge();
        outboxTemplate = new Outbox();
    }

    function updateTemplates(
        address _delayedBridgeTemplate,
        address _sequencerInboxTemplate,
        address _inboxTemplate,
        address _rollupEventBridgeTemplate,
        address _outboxTemplate
    ) external onlyOwner {
        delayedBridgeTemplate = Bridge(_delayedBridgeTemplate);
        sequencerInboxTemplate = SequencerInbox(_sequencerInboxTemplate);
        inboxTemplate = Inbox(_inboxTemplate);
        rollupEventBridgeTemplate = RollupEventBridge(_rollupEventBridgeTemplate);
        outboxTemplate = Outbox(_outboxTemplate);

        emit TemplatesUpdated();
    }

    struct CreateBridgeFrame {
        ProxyAdmin admin;
        Bridge delayedBridge;
        SequencerInbox sequencerInbox;
        Inbox inbox;
        RollupEventBridge rollupEventBridge;
        Outbox outbox;
        Whitelist whitelist;
    }

    function createBridge(
        address adminProxy,
        address rollup,
        address sequencer
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
        CreateBridgeFrame memory frame;
        {
            frame.delayedBridge = Bridge(
                address(
                    new TransparentUpgradeableProxy(address(delayedBridgeTemplate), adminProxy, "")
                )
            );
            frame.sequencerInbox = SequencerInbox(
                address(
                    new TransparentUpgradeableProxy(address(sequencerInboxTemplate), adminProxy, "")
                )
            );
            frame.inbox = Inbox(
                address(new TransparentUpgradeableProxy(address(inboxTemplate), adminProxy, ""))
            );
            frame.rollupEventBridge = RollupEventBridge(
                address(
                    new TransparentUpgradeableProxy(
                        address(rollupEventBridgeTemplate),
                        adminProxy,
                        ""
                    )
                )
            );
            frame.outbox = Outbox(
                address(new TransparentUpgradeableProxy(address(outboxTemplate), adminProxy, ""))
            );
            frame.whitelist = new Whitelist();
        }

        frame.delayedBridge.initialize();
        frame.sequencerInbox.initialize(IBridge(frame.delayedBridge), sequencer, rollup);
        frame.inbox.initialize(IBridge(frame.delayedBridge), address(frame.whitelist));
        frame.rollupEventBridge.initialize(address(frame.delayedBridge), rollup);
        frame.outbox.initialize(rollup, IBridge(frame.delayedBridge));

        frame.delayedBridge.setInbox(address(frame.inbox), true);
        frame.delayedBridge.transferOwnership(rollup);

        frame.whitelist.setOwner(rollup);

        return (
            frame.delayedBridge,
            frame.sequencerInbox,
            frame.inbox,
            frame.rollupEventBridge,
            frame.outbox
        );
    }
}
