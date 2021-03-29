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

import "./interfaces/ISequencerInbox.sol";

contract SequencerInbox is ISequencerInbox {
    uint8 internal constant L2_MSG = 3;
    bytes32[] public override inboxAccs;
    uint256 public override messageCount;
    address public sequencer;
    uint256 public override maxDelayBlocks;
    uint256 public override maxDelaySeconds;

    function addSequencerL2Batch(
        bytes calldata transactions,
        uint256[] calldata lengths,
        uint256 l1BlockNumber,
        uint256 timestamp
    ) external returns (uint256) {
        require(msg.sender == sequencer, "ONLY_SEQUENCER");
        require(l1BlockNumber + maxDelayBlocks >= block.number, "BLOCK_TOO_OLD");
        require(l1BlockNumber <= block.number, "BLOCK_TOO_NEW");
        require(timestamp + maxDelaySeconds >= block.timestamp, "TIME_TOO_OLD");
        require(timestamp <= block.timestamp, "TIME_TOO_NEW");

        bytes32 prevAcc = 0;
        if (inboxAccs.length > 0) {
            prevAcc = inboxAccs[inboxAccs.length - 1];
        }
        uint256 offset = 0;
        uint256 count = messageCount;
        for (uint256 i = 0; i < lengths.length; i++) {
            bytes32 messageDataHash = keccak256(bytes(transactions[offset:offset + lengths[i]]));
            bytes32 messageHash =
                Messages.messageHash(
                    L2_MSG,
                    msg.sender,
                    l1BlockNumber,
                    timestamp, // solhint-disable-line not-rely-on-time
                    count,
                    tx.gasprice,
                    messageDataHash
                );
            prevAcc = Messages.addMessageToInbox(prevAcc, messageHash);
            offset += lengths[i];
            count++;
        }
        inboxAccs.push(prevAcc);
        messageCount = count;
        return count;
    }
}
