// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2020, Offchain Labs, Inc.
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

import "arb-bridge-eth/contracts/bridge/interfaces/IInbox.sol";
import "arb-bridge-eth/contracts/bridge/interfaces/IOutbox.sol";
import "arb-bridge-eth/contracts/bridge/interfaces/IBridge.sol";

contract L1Buddy {
    IInbox inbox;
    IBridge bridge;
    bool connectedToChain;

    modifier onlyIfConnected {
        require(connectedToChain, "NOT_CONNECTED");
        _;
    }

    modifier onlyL2 {
        require(msg.sender == address(bridge), "ONLY_BRIDGE");
        _;
    }

    constructor(IInbox _inbox) public {
        inbox = _inbox;
        bridge = IBridge(_inbox.bridge());
    }

    function buddyCreated(bool successful) external onlyL2 {
        // This method must be called by the l2 system rather than a contract
        require(l2Sender() == address(0), "ONLY_SYSTEM");
        if (successful) {
            connectedToChain = true;
        }
    }

    function l2Sender() internal view returns (address) {
        return IOutbox(bridge.activeOutbox()).l2ToL1Sender();
    }
}
