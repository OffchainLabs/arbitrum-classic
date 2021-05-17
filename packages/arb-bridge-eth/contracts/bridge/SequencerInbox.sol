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

import "./interfaces/ISequencerInbox.sol";
import "./interfaces/IBridge.sol";
import "../arch/Marshaling.sol";
import "../libraries/Cloneable.sol";

import "./Messages.sol";

contract SequencerInbox is ISequencerInbox, Cloneable {
    uint8 internal constant L2_MSG = 3;
    uint8 internal constant END_OF_BLOCK = 6;

    bytes32[] public override inboxAccs;
    uint256 public override messageCount;

    uint256 totalDelayedMessagesRead;

    IBridge public delayedInbox;
    address public sequencer;
    uint256 public override maxDelayBlocks;
    uint256 public override maxDelaySeconds;

    function initialize(
        IBridge _delayedInbox,
        address _sequencer,
        uint256 _maxDelayBlocks,
        uint256 _maxDelaySeconds
    ) external {
        require(address(delayedInbox) == address(0), "ALREADY_INIT");
        delayedInbox = _delayedInbox;
        sequencer = _sequencer;
        maxDelayBlocks = _maxDelayBlocks;
        maxDelaySeconds = _maxDelaySeconds;
    }

    function getLastDelayedAcc() internal view returns (bytes32) {
        bytes32 acc = 0;
        if (totalDelayedMessagesRead > 0) {
            acc = delayedInbox.inboxAccs(totalDelayedMessagesRead - 1);
        }
        return acc;
    }

    function forceInclusion(
        uint256 _totalDelayedMessagesRead,
        uint8 kind,
        uint256[2] calldata l1BlockAndTimestamp,
        uint256 inboxSeqNum,
        uint256 gasPriceL1,
        address sender,
        bytes32 messageDataHash
    ) external {
        require(_totalDelayedMessagesRead > totalDelayedMessagesRead, "DELAYED_BACKWARDS");
        {
            bytes32 messageHash =
                Messages.messageHash(
                    kind,
                    sender,
                    l1BlockAndTimestamp[0],
                    l1BlockAndTimestamp[1],
                    inboxSeqNum,
                    gasPriceL1,
                    messageDataHash
                );
            require(l1BlockAndTimestamp[0] + maxDelayBlocks < block.number, "MAX_DELAY_BLOCKS");
            require(l1BlockAndTimestamp[1] + maxDelaySeconds < block.timestamp, "MAX_DELAY_TIME");

            bytes32 prevDelayedAcc = 0;
            if (_totalDelayedMessagesRead > 1) {
                prevDelayedAcc = delayedInbox.inboxAccs(_totalDelayedMessagesRead - 2);
            }
            require(
                delayedInbox.inboxAccs(_totalDelayedMessagesRead - 1) ==
                    Messages.addMessageToInbox(prevDelayedAcc, messageHash),
                "DELAYED_ACCUMULATOR"
            );
        }

        uint256 startNum = messageCount;
        bytes32 beforeAcc = 0;
        if (inboxAccs.length > 0) {
            beforeAcc = inboxAccs[inboxAccs.length - 1];
        }

        (bytes32 delayedAcc, bytes32 acc, uint256 count) =
            includeDelayedMessages(
                beforeAcc,
                startNum,
                _totalDelayedMessagesRead,
                block.number,
                block.timestamp
            );
        inboxAccs.push(acc);
        messageCount = count;
        emit DelayedInboxForced(
            startNum,
            beforeAcc,
            count,
            _totalDelayedMessagesRead,
            [acc, delayedAcc],
            inboxAccs.length - 1
        );
    }

    function addSequencerL2BatchFromOrigin(
        bytes calldata transactions,
        uint256[] calldata lengths,
        uint256 l1BlockNumber,
        uint256 timestamp,
        uint256 _totalDelayedMessagesRead,
        bytes32 afterAcc
    ) external {
        // solhint-disable-next-line avoid-tx-origin
        require(msg.sender == tx.origin, "origin only");
        uint256 startNum = messageCount;
        (bytes32 beforeAcc, bytes32 delayedAcc) =
            addSequencerL2BatchImpl(
                transactions,
                lengths,
                l1BlockNumber,
                timestamp,
                _totalDelayedMessagesRead,
                afterAcc
            );
        emit SequencerBatchDeliveredFromOrigin(
            startNum,
            beforeAcc,
            messageCount,
            afterAcc,
            delayedAcc,
            inboxAccs.length - 1
        );
    }

    function addSequencerL2Batch(
        bytes calldata transactions,
        uint256[] calldata lengths,
        uint256 l1BlockNumber,
        uint256 timestamp,
        uint256 _totalDelayedMessagesRead,
        bytes32 afterAcc
    ) external {
        uint256 startNum = messageCount;
        (bytes32 beforeAcc, bytes32 delayedAcc) =
            addSequencerL2BatchImpl(
                transactions,
                lengths,
                l1BlockNumber,
                timestamp,
                _totalDelayedMessagesRead,
                afterAcc
            );
        emit SequencerBatchDelivered(
            startNum,
            beforeAcc,
            messageCount,
            afterAcc,
            transactions,
            lengths,
            l1BlockNumber,
            timestamp,
            _totalDelayedMessagesRead,
            delayedAcc,
            inboxAccs.length - 1
        );
    }

    function addSequencerL2BatchImpl(
        bytes calldata transactions,
        uint256[] calldata lengths,
        uint256 l1BlockNumber,
        uint256 timestamp,
        uint256 _totalDelayedMessagesRead,
        bytes32 afterAcc
    ) private returns (bytes32 beforeAcc, bytes32 delayedAcc) {
        uint256 txCount = lengths.length;
        require(msg.sender == sequencer, "ONLY_SEQUENCER");
        require(l1BlockNumber + maxDelayBlocks >= block.number, "BLOCK_TOO_OLD");
        require(l1BlockNumber <= block.number, "BLOCK_TOO_NEW");
        require(timestamp + maxDelaySeconds >= block.timestamp, "TIME_TOO_OLD");
        require(timestamp <= block.timestamp, "TIME_TOO_NEW");
        require(_totalDelayedMessagesRead >= totalDelayedMessagesRead, "DELAYED_BACKWARDS");
        require(_totalDelayedMessagesRead >= 1, "MUST_DELAYED_INIT");
        require(totalDelayedMessagesRead >= 1 || txCount == 0, "MUST_DELAYED_INIT_START");

        if (inboxAccs.length > 0) {
            beforeAcc = inboxAccs[inboxAccs.length - 1];
        }

        bytes32 prefixHash = keccak256(abi.encodePacked(msg.sender, l1BlockNumber, timestamp));
        (bytes32 acc, uint256 count) = calcL2Batch(transactions, lengths, prefixHash, beforeAcc);

        (delayedAcc, acc, count) = includeDelayedMessages(
            acc,
            count,
            _totalDelayedMessagesRead,
            l1BlockNumber,
            timestamp
        );

        require(count > messageCount, "EMPTY_BATCH");
        inboxAccs.push(acc);
        messageCount = count;

        require(acc == afterAcc, "AFTER_ACC");
    }

    function calcL2Batch(
        bytes memory transactions,
        uint256[] calldata lengths,
        bytes32 prefixHash,
        bytes32 beforeAcc
    ) private view returns (bytes32 acc, uint256 count) {
        uint256 txCount = lengths.length;
        count = messageCount;
        acc = beforeAcc;
        uint256 offset;
        assembly {
            offset := add(transactions, 32)
        }
        for (uint256 i = 0; i < txCount; i++) {
            uint256 length = lengths[i];
            bytes32 messageDataHash;
            assembly {
                messageDataHash := keccak256(offset, length)
            }
            acc = keccak256(abi.encodePacked(acc, count, prefixHash, messageDataHash));
            offset += length;
            count++;
        }
        return (acc, count);
    }

    function includeDelayedMessages(
        bytes32 acc,
        uint256 count,
        uint256 _totalDelayedMessagesRead,
        uint256 l1BlockNumber,
        uint256 timestamp
    )
        private
        returns (
            bytes32,
            bytes32,
            uint256
        )
    {
        bytes32 delayedAcc;
        if (_totalDelayedMessagesRead > totalDelayedMessagesRead) {
            require(_totalDelayedMessagesRead <= delayedInbox.messageCount(), "DELAYED_TOO_FAR");
            delayedAcc = delayedInbox.inboxAccs(_totalDelayedMessagesRead - 1);
            acc = keccak256(
                abi.encodePacked(
                    "Delayed messages:",
                    acc,
                    count,
                    totalDelayedMessagesRead,
                    _totalDelayedMessagesRead,
                    delayedAcc
                )
            );
            count += _totalDelayedMessagesRead - totalDelayedMessagesRead;
            bytes memory emptyBytes;
            acc = keccak256(
                abi.encodePacked(
                    acc,
                    count,
                    keccak256(abi.encodePacked(address(0), l1BlockNumber, timestamp)),
                    keccak256(emptyBytes)
                )
            );
            count++;
            totalDelayedMessagesRead = _totalDelayedMessagesRead;
        }
        return (delayedAcc, acc, count);
    }

    function proveSeqBatchMsgCount(
        bytes calldata proof,
        uint256 offset,
        bytes32 acc
    ) internal pure returns (uint256, uint256) {
        uint256 endCount;

        bytes32 buildingAcc;
        uint256 seqNum;
        bytes32 messageHeaderHash;
        bytes32 messageDataHash;
        (offset, buildingAcc) = Marshaling.deserializeBytes32(proof, offset);
        (offset, seqNum) = Marshaling.deserializeInt(proof, offset);
        (offset, messageHeaderHash) = Marshaling.deserializeBytes32(proof, offset);
        (offset, messageDataHash) = Marshaling.deserializeBytes32(proof, offset);
        buildingAcc = keccak256(
            abi.encodePacked(buildingAcc, seqNum, messageHeaderHash, messageDataHash)
        );
        endCount = seqNum + 1;
        require(buildingAcc == acc, "BATCH_ACC");

        return (offset, endCount);
    }

    function proveBatchContainsSequenceNumber(bytes calldata proof, uint256 inboxCount)
        external
        view
        override
        returns (uint256, bytes32)
    {
        if (inboxCount == 0) {
            return (0, 0);
        }

        (uint256 offset, uint256 seqBatchNum) = Marshaling.deserializeInt(proof, 0);
        uint256 lastBatchCount = 0;
        if (seqBatchNum > 0) {
            (offset, lastBatchCount) = proveSeqBatchMsgCount(
                proof,
                offset,
                inboxAccs[seqBatchNum - 1]
            );
            lastBatchCount++;
        }

        bytes32 seqBatchAcc = inboxAccs[seqBatchNum];
        uint256 thisBatchCount;
        (offset, thisBatchCount) = proveSeqBatchMsgCount(proof, offset, seqBatchAcc);

        require(inboxCount > lastBatchCount, "BATCH_START");
        require(inboxCount <= thisBatchCount, "BATCH_END");

        return (thisBatchCount, seqBatchAcc);
    }
}
