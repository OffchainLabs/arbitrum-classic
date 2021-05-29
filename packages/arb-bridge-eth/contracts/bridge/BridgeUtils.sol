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

import "./interfaces/IBridge.sol";
import "./interfaces/ISequencerInbox.sol";

contract BridgeUtils {
    function getCountsAndAccumulators(IBridge delayedBridge, ISequencerInbox sequencerInbox)
        external
        view
        returns (uint256[2] memory counts, bytes32[2] memory accs)
    {
        uint256 delayedCount = delayedBridge.messageCount();
        if (delayedCount > 0) {
            counts[0] = delayedCount;
            accs[0] = delayedBridge.inboxAccs(delayedCount - 1);
        }
        uint256 sequencerCount = sequencerInbox.getInboxAccsLength();
        if (sequencerCount > 0) {
            counts[1] = sequencerCount;
            accs[1] = sequencerInbox.inboxAccs(sequencerCount - 1);
        }
    }
}
