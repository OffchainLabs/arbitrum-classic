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

contract BridgeUtils {
    function getCountsAndAccumulators(IBridge[2] calldata bridges)
        external
        view
        returns (uint256[2] memory counts, bytes32[2] memory accs)
    {
        for (uint256 i = 0; i < 2; i++) {
            IBridge bridge = bridges[i];
            uint256 count = bridge.messageCount();
            counts[i] = count;
            if (count > 0) {
                accs[i] = bridge.inboxAccs(count - 1);
            }
        }
    }
}
