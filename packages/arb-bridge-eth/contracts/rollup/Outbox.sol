// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2019-2021, Offchain Labs, Inc.
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

import "./Messages.sol";

import "../libraries/MerkleLib.sol";
import "../libraries/BytesLib.sol";

contract Outbox {
    using BytesLib for bytes;

    uint8 internal constant MSG_ROOT = 0;

    OutboxEntry[] outboxes;

    function processOutgoingMessages(bytes memory messageData, uint256[] memory messageLengths)
        internal
    {
        // If we've reached here, we've already confirmed that sum(messageLengths) == messageData.length
        uint256 messageCount = messageLengths.length;
        uint256 offset = 0;
        for (uint256 i = 0; i < messageCount; i++) {
            // Otherwise we have an unsupported message type and we skip the message
            if (uint8(messageData[offset]) == MSG_ROOT) {
                bytes32 outputRoot = messageData.toBytes32(offset + 1);
                outboxes.push(new OutboxEntry(outputRoot));
            }
            offset += messageLengths[i];
        }
    }

    function executeTransaction(
        uint256 outboxIndex,
        bytes calldata _proof,
        uint256 _index,
        address destAddr,
        uint256 amount,
        bytes calldata calldataForL1
    ) external {
        bytes32 userTx = keccak256(
            abi.encodePacked(uint256(uint160(bytes20(destAddr))), amount, calldataForL1)
        );

        spendOutput(outboxIndex, _proof, _index, userTx);

        (bool success, ) = destAddr.call.value(amount)(calldataForL1);
        require(success);
    }

    function spendOutput(
        uint256 outboxIndex,
        bytes memory proof,
        uint256 index,
        bytes32 item
    ) private {
        // Flip the first bit to prove this is a leaf
        item = item ^ bytes32(uint256(1));
        (bytes32 calcRoot, ) = MerkleLib.verifyMerkleProof(proof, item, index + 1);
        outboxes[outboxIndex].spendOutput(calcRoot, index);
    }
}

contract OutboxEntry {
    address rollup;
    bytes32 outputRoot;
    mapping(uint256 => bool) spentOutput;

    constructor(bytes32 root) public {
        rollup = msg.sender;
        outputRoot = root;
    }

    function spendOutput(bytes32 calcRoot, uint256 index) external {
        require(msg.sender == rollup);
        require(!spentOutput[index]);
        require(calcRoot == outputRoot);
        spentOutput[index] = true;
    }
}
