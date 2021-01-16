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

pragma solidity ^0.6.11;

import "./OutboxCore.sol";
import "./OutboxEntry.sol";

import "./IOutbox.sol";

import "./Messages.sol";
import "../libraries/MerkleLib.sol";
import "../libraries/BytesLib.sol";

contract Outbox is OutboxCore, IOutbox {
    using BytesLib for bytes;

    uint8 internal constant MSG_ROOT = 0;

    // Note, this variables are set and then wiped during a single transaction. Therefore their values don't need to be maintained, and their slots will be empty outside of transactions
    address private _l2ToL1Sender;
    uint128 private _l2ToL1Block;
    uint128 private _l2ToL1Timestamp;

    function l2ToL1Sender() external view override returns (address) {
        return _l2ToL1Sender;
    }

    function l2ToL1Block() external view override returns (uint256) {
        return uint256(_l2ToL1Sender);
    }

    function l2ToL1Timestamp() external view override returns (uint256) {
        return uint256(_l2ToL1Sender);
    }

    function _processOutgoingMessages(bytes memory sendsData, uint256[] calldata sendLengths)
        internal
    {
        // If we've reached here, we've already confirmed that sum(sendLengths) == sendsData.length
        uint256 messageCount = sendLengths.length;
        uint256 offset = 0;
        for (uint256 i = 0; i < messageCount; i++) {
            // Otherwise we have an unsupported message type and we skip the message
            if (uint8(sendsData[offset]) == MSG_ROOT) {
                bytes32 outputRoot = sendsData.toBytes32(offset + 1);
                outboxes.push(new OutboxEntry(outputRoot));
            }
            offset += sendLengths[i];
        }
    }

    function executeTransaction(
        uint256 outboxIndex,
        bytes calldata proof,
        uint256 index,
        address l2Sender,
        address destAddr,
        uint256 l2Block,
        uint256 l2Timestamp,
        uint256 amount,
        bytes calldata calldataForL1
    ) external override {
        bytes32 userTx =
            keccak256(
                abi.encodePacked(
                    uint256(uint160(bytes20(l2Sender))),
                    uint256(uint160(bytes20(destAddr))),
                    l2Block,
                    l2Timestamp,
                    amount,
                    calldataForL1
                )
            );

        spendOutput(outboxIndex, proof, index, userTx);

        _l2ToL1Sender = l2Sender;
        _l2ToL1Block = uint128(l2Block);
        _l2ToL1Timestamp = uint128(l2Timestamp);

        (bool success, ) = destAddr.call{ value: amount }(calldataForL1);
        require(success);

        _l2ToL1Sender = address(0);
        _l2ToL1Block = 0;
        _l2ToL1Timestamp = 0;
    }

    function spendOutput(
        uint256 outboxIndex,
        bytes memory proof,
        uint256 path,
        bytes32 item
    ) private {
        // Hash the leaf an extra time to prove it's a leaf
        (bytes32 calcRoot, ) =
            MerkleLib.verifyMerkleProof(proof, keccak256(abi.encodePacked(item)), path);
        outboxes[outboxIndex].spendOutput(calcRoot, path);
    }
}
