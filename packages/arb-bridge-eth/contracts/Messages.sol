// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

pragma solidity ^0.5.11;

import "./arch/Value.sol";

library Messages {
    function messageHash(
        uint8 kind,
        address sender,
        uint256 blockNumber,
        uint256 timestamp,
        uint256 inboxSeqNum,
        bytes32 messageDataHash
    ) internal pure returns (bytes32) {
        return
            keccak256(
                abi.encodePacked(
                    kind,
                    sender,
                    blockNumber,
                    timestamp,
                    inboxSeqNum,
                    messageDataHash
                )
            );
    }

    function messageValue(
        uint8 kind,
        uint256 blockNumber,
        uint256 timestamp,
        address sender,
        uint256 inboxSeqNum,
        bytes memory messageData
    ) internal pure returns (Value.Data memory) {
        Value.Data[] memory tupData = new Value.Data[](6);
        tupData[0] = Value.newInt(uint256(kind));
        tupData[1] = Value.newInt(blockNumber);
        tupData[2] = Value.newInt(timestamp);
        tupData[3] = Value.newInt(uint256(sender));
        tupData[4] = Value.newInt(inboxSeqNum);
        tupData[5] = Value.bytesToBytestackHash(
            messageData,
            0,
            messageData.length
        );
        return Value.newTuple(tupData);
    }

    function addMessageToInbox(bytes32 inbox, bytes32 message)
        internal
        pure
        returns (bytes32)
    {
        return keccak256(abi.encodePacked(inbox, message));
    }

    function addMessageToVMInbox(
        Value.Data memory vmInboxHashValue,
        Value.Data memory message
    ) internal pure returns (Value.Data memory) {
        Value.Data[] memory vals = new Value.Data[](2);
        vals[0] = vmInboxHashValue;
        vals[1] = message;
        return Value.newTuple(vals);
    }
}
