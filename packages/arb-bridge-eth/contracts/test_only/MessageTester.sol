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

import "../bridge/Messages.sol";

contract MessageTester {
    function messageHash(
        uint8 messageType,
        address sender,
        uint256 blockNumber,
        uint256 timestamp,
        uint256 inboxSeqNum,
        uint256 gasPriceL1,
        bytes32 messageDataHash
    ) public pure returns (bytes32) {
        return
            Messages.messageHash(
                messageType,
                sender,
                blockNumber,
                timestamp,
                inboxSeqNum,
                gasPriceL1,
                messageDataHash
            );
    }

    function addMessageToInbox(bytes32 inbox, bytes32 message) public pure returns (bytes32) {
        return Messages.addMessageToInbox(inbox, message);
    }
}
