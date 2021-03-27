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

import "./Messages.sol";
import "../libraries/MerkleLib.sol";

contract SequencerInbox {
    uint8 internal constant L2_MSG = 3;
    bytes32[] public inboxAccs;
    uint256 totalMessages;
    address public sequencer;
    uint256 maxBlock;
    uint256 maxTimestamp;
    uint256 maxDelayBlocks;
    uint256 maxDelaySeconds;

    function verifyGroups(
        uint256[] memory l1BlockNumbers,
        uint256[] memory timestamps,
        uint256[] calldata groupSizes
    ) private view returns (uint256) {
        uint256 count = l1BlockNumbers.length;
        require(timestamps.length == count, "LENGTH_MISMATCH");
        require(groupSizes.length == count, "LENGTH_MISMATCH");
        require(groupSizes.length > 0, "MUST_HAVE_GROUP");
        uint256 prevBlock = maxBlock;
        uint256 prevTimestamp = maxTimestamp;
        uint256 total = 0;
        for (uint256 i = 0; i < count; i++) {
            uint256 groupSize = groupSizes[i];
            require(groupSize > 0, "NEED_NONZERO_GROUP");
            uint256 l1BlockNumber = l1BlockNumbers[i];
            uint256 timestamp = timestamps[i];
            require(l1BlockNumber + maxDelayBlocks >= block.number, "BLOCK_TOO_OLD");
            require(l1BlockNumber <= block.number, "BLOCK_TOO_NEW");
            require(l1BlockNumber >= prevBlock, "BLOCK_WENT_BACK");
            require(timestamp + maxDelaySeconds >= block.timestamp, "TIME_TOO_OLD");
            require(timestamp <= block.timestamp, "TIME_TOO_NEW");
            require(timestamp >= prevTimestamp, "TIME_WENT_BACK");
            prevBlock = l1BlockNumber;
            prevTimestamp = timestamp;
            total += groupSize;
        }
        return total;
    }

    function addSequencerL2Batch(
        bytes calldata transactions,
        uint256[] calldata lengths,
        uint256[] calldata groupSizes,
        uint256[] calldata l1BlockNumbers,
        uint256[] calldata timestamps
    ) external returns (uint256 batchCount) {
        require(msg.sender == sequencer, "ONLY_SEQUENCER");
        require(groupSizes.length < 50, "TOO_MANY_GROUPS");
        require(
            verifyGroups(l1BlockNumbers, timestamps, groupSizes) == lengths.length,
            "BAD_GROUP_SIZES"
        );
        uint256 offset = 0;
        bytes32[] memory hashes = new bytes32[](lengths.length);
        // Use return value to store variable temporarily to fix stack size issue
        batchCount = totalMessages;
        for (uint256 i = 0; i < lengths.length; i++) {
            bytes32 transactionHash = keccak256(bytes(transactions[offset:offset + lengths[i]]));
            hashes[i] = keccak256(abi.encodePacked(batchCount + i, transactionHash));
            offset += lengths[i];
        }
        bytes32 batchHash =
            keccak256(
                abi.encode(MerkleLib.generateRoot(hashes), groupSizes, l1BlockNumbers, timestamps)
            );
        batchCount = inboxAccs.length;
        bytes32 prevAcc = 0;
        if (batchCount > 0) {
            prevAcc = inboxAccs[batchCount - 1];
        }
        inboxAccs.push(Messages.addMessageToInbox(prevAcc, batchHash));
        maxBlock = l1BlockNumbers[l1BlockNumbers.length - 1];
        maxTimestamp = timestamps[timestamps.length - 1];
        totalMessages += lengths.length;
        return batchCount;
    }

    function messageCount() external view returns (uint256) {
        return inboxAccs.length;
    }
}
