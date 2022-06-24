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

contract BridgeMock is Bridge {
    constructor() public {
        activeOutbox = msg.sender;
    }

    function deliverMessageToInboxTest(
        uint8 kind,
        address sender,
        uint256 blockNumber,
        uint256 blockTimestamp,
        uint256 gasPrice,
        bytes32 messageDataHash
    ) external payable returns (uint256) {
        return
            addMessageToInbox(kind, sender, blockNumber, blockTimestamp, gasPrice, messageDataHash);
    }
}
