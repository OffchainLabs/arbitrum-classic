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
import "../rollup/Rollup.sol";
import "../validator/IGasRefunder.sol";

import "./Messages.sol";

interface OldRollup {
    function sequencerInboxMaxDelayBlocks() external view returns (uint256);

    function sequencerInboxMaxDelaySeconds() external view returns (uint256);
}

contract SequencerInbox is ISequencerInbox, Cloneable {
    // Sequencer-Inbox state accumulator
    bytes32[] public override inboxAccs;

    // Number of messages included in the sequencer-inbox; tracked seperately from inboxAccs since multiple messages can be included in a single inboxAcc update (i.e., many messages in a batch, many batches in a single inboxAccs update, etc)
    uint256 public override messageCount;

    // count of messages read from the delayedInbox
    uint256 public totalDelayedMessagesRead;

    IBridge public delayedInbox;
    address private deprecatedSequencer;
    address public rollup;
    mapping(address => bool) public override isSequencer;

    // Window in which only the Sequencer can update the Inbox; this delay is what allows the Sequencer to give receipts with sub-blocktime latency.
    uint256 public override maxDelayBlocks;
    uint256 public override maxDelaySeconds;

    function initialize(
        IBridge _delayedInbox,
        address _sequencer,
        address _rollup
    ) external {
        require(address(delayedInbox) == address(0), "ALREADY_INIT");
        delayedInbox = _delayedInbox;
        isSequencer[_sequencer] = true;
        rollup = _rollup;
        // it is assumed that maxDelayBlocks and maxDelaySeconds are set by the rollup
    }

    function postUpgradeInit() external {
        // it is assumed the sequencer inbox contract is behind a Proxy controlled by a
        // proxy admin. this function can only be called by the proxy admin contract
        address proxyAdmin = ProxyUtil.getProxyAdmin();
        require(msg.sender == proxyAdmin, "NOT_FROM_ADMIN");

        // the sequencer inbox needs to query the old rollup interface since it will be upgraded first
        OldRollup _rollup = OldRollup(rollup);

        maxDelayBlocks = _rollup.sequencerInboxMaxDelayBlocks();
        maxDelaySeconds = _rollup.sequencerInboxMaxDelaySeconds();

        isSequencer[deprecatedSequencer] = true;
    }

    /// @notice DEPRECATED - use isSequencer instead
    function sequencer() external view override returns (address) {
        return deprecatedSequencer;
    }

    function setIsSequencer(address addr, bool newIsSequencer) external override {
        require(msg.sender == rollup, "ONLY_ROLLUP");
        isSequencer[addr] = newIsSequencer;
        emit IsSequencerUpdated(addr, newIsSequencer);
    }

    function setMaxDelay(uint256 newMaxDelayBlocks, uint256 newMaxDelaySeconds) external override {
        require(msg.sender == rollup, "ONLY_ROLLUP");
        maxDelayBlocks = newMaxDelayBlocks;
        maxDelaySeconds = newMaxDelaySeconds;
        emit MaxDelayUpdated(newMaxDelayBlocks, newMaxDelaySeconds);
    }

    /**
     * @notice Move messages from the delayed inbox into the Sequencer inbox. Callable by any address. Necessary iff Sequencer hasn't included them before delay period expired.
     */

    function forceInclusion(
        uint256 _totalDelayedMessagesRead,
        uint8 kind,
        uint256[2] calldata l1BlockAndTimestamp,
        uint256 inboxSeqNum,
        uint256 gasPriceL1,
        address sender,
        bytes32 messageDataHash,
        bytes32 delayedAcc
    ) external {
        require(_totalDelayedMessagesRead > totalDelayedMessagesRead, "DELAYED_BACKWARDS");
        {
            bytes32 messageHash = Messages.messageHash(
                kind,
                sender,
                l1BlockAndTimestamp[0],
                l1BlockAndTimestamp[1],
                inboxSeqNum,
                gasPriceL1,
                messageDataHash
            );
            // Can only force-include after the Sequencer-only window has expired.
            require(l1BlockAndTimestamp[0] + maxDelayBlocks < block.number, "MAX_DELAY_BLOCKS");
            require(l1BlockAndTimestamp[1] + maxDelaySeconds < block.timestamp, "MAX_DELAY_TIME");

            // Verify that message hash represents the last message sequence of delayed message to be included
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

        (bytes32 acc, uint256 count) = includeDelayedMessages(
            beforeAcc,
            startNum,
            _totalDelayedMessagesRead,
            block.number,
            block.timestamp,
            delayedAcc
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
        uint256[] calldata sectionsMetadata,
        bytes32 afterAcc
    ) external {
        // solhint-disable-next-line avoid-tx-origin
        require(msg.sender == tx.origin, "origin only");
        uint256 startNum = messageCount;
        bytes32 beforeAcc = addSequencerL2BatchImpl(
            transactions,
            lengths,
            sectionsMetadata,
            afterAcc
        );
        emit SequencerBatchDeliveredFromOrigin(
            startNum,
            beforeAcc,
            messageCount,
            afterAcc,
            inboxAccs.length - 1
        );
    }

    function addSequencerL2BatchFromOriginWithGasRefunder(
        bytes calldata transactions,
        uint256[] calldata lengths,
        uint256[] calldata sectionsMetadata,
        bytes32 afterAcc,
        IGasRefunder gasRefunder
    ) external {
        // solhint-disable-next-line avoid-tx-origin
        require(msg.sender == tx.origin, "origin only");

        // Add in a lower bound of calldata cost to the gas estimate
        // startGasLeft = gasleft() + 4 * calldatasize()
        uint256 startGasLeft;
        assembly {
            startGasLeft := mul(calldatasize(), 4)
        }
        startGasLeft += gasleft();

        uint256 startNum = messageCount;
        bytes32 beforeAcc = addSequencerL2BatchImpl(
            transactions,
            lengths,
            sectionsMetadata,
            afterAcc
        );
        emit SequencerBatchDeliveredFromOrigin(
            startNum,
            beforeAcc,
            messageCount,
            afterAcc,
            inboxAccs.length - 1
        );

        if (gasRefunder != IGasRefunder(0)) {
            gasRefunder.onGasSpent(msg.sender, startGasLeft - gasleft());
        }
    }

    /**
     * @notice Sequencer adds a batch to inbox.
     * @param transactions concatenated bytes of L2 messages
     * @param lengths length of each txn in transctions (for parsing)
     * @param sectionsMetadata Each consists of [numItems, l1BlockNumber, l1Timestamp, newTotalDelayedMessagesRead, newDelayedAcc]
     * @param afterAcc Expected inbox hash after batch is added
     * @dev sectionsMetadata lets the sequencer delineate new l1Block numbers and l1Timestamps within a given batch; this lets the sequencer minimize the number of batches created (and thus amortizing cost) while still giving timely receipts
     */
    function addSequencerL2Batch(
        bytes calldata transactions,
        uint256[] calldata lengths,
        uint256[] calldata sectionsMetadata,
        bytes32 afterAcc
    ) external {
        uint256 startNum = messageCount;
        bytes32 beforeAcc = addSequencerL2BatchImpl(
            transactions,
            lengths,
            sectionsMetadata,
            afterAcc
        );
        emit SequencerBatchDelivered(
            startNum,
            beforeAcc,
            messageCount,
            afterAcc,
            transactions,
            lengths,
            sectionsMetadata,
            inboxAccs.length - 1,
            msg.sender
        );
    }

    function addSequencerL2BatchImpl(
        bytes memory transactions,
        uint256[] calldata lengths,
        uint256[] calldata sectionsMetadata,
        bytes32 afterAcc
    ) private returns (bytes32 beforeAcc) {
        require(isSequencer[msg.sender], "ONLY_SEQUENCER");

        if (inboxAccs.length > 0) {
            beforeAcc = inboxAccs[inboxAccs.length - 1];
        }

        uint256 runningCount = messageCount;
        bytes32 runningAcc = beforeAcc;
        uint256 processedItems = 0;
        uint256 dataOffset;
        assembly {
            dataOffset := add(transactions, 32)
        }
        for (uint256 i = 0; i + 5 <= sectionsMetadata.length; i += 5) {
            // Each metadata section consists of:
            // [numItems, l1BlockNumber, l1Timestamp, newTotalDelayedMessagesRead, newDelayedAcc]
            {
                uint256 l1BlockNumber = sectionsMetadata[i + 1];
                require(l1BlockNumber + maxDelayBlocks >= block.number, "BLOCK_TOO_OLD");
                require(l1BlockNumber <= block.number, "BLOCK_TOO_NEW");
            }
            {
                uint256 l1Timestamp = sectionsMetadata[i + 2];
                require(l1Timestamp + maxDelaySeconds >= block.timestamp, "TIME_TOO_OLD");
                require(l1Timestamp <= block.timestamp, "TIME_TOO_NEW");
            }

            {
                bytes32 prefixHash = keccak256(
                    abi.encodePacked(msg.sender, sectionsMetadata[i + 1], sectionsMetadata[i + 2])
                );
                uint256 numItems = sectionsMetadata[i];
                (runningAcc, runningCount, dataOffset) = calcL2Batch(
                    dataOffset,
                    lengths,
                    processedItems,
                    numItems,
                    prefixHash,
                    runningCount,
                    runningAcc
                );
                processedItems += numItems;
            }

            uint256 newTotalDelayedMessagesRead = sectionsMetadata[i + 3];
            require(newTotalDelayedMessagesRead >= totalDelayedMessagesRead, "DELAYED_BACKWARDS");
            require(newTotalDelayedMessagesRead >= 1, "MUST_DELAYED_INIT");
            require(
                totalDelayedMessagesRead >= 1 || sectionsMetadata[i] == 0,
                "MUST_DELAYED_INIT_START"
            );
            // Sequencer decides how many messages (if any) to include from the delayed inbox
            if (newTotalDelayedMessagesRead > totalDelayedMessagesRead) {
                (runningAcc, runningCount) = includeDelayedMessages(
                    runningAcc,
                    runningCount,
                    newTotalDelayedMessagesRead,
                    sectionsMetadata[i + 1], // block number
                    sectionsMetadata[i + 2], // timestamp
                    bytes32(sectionsMetadata[i + 4]) // delayed accumulator
                );
            }
        }

        uint256 startOffset;
        assembly {
            startOffset := add(transactions, 32)
        }
        require(dataOffset >= startOffset, "OFFSET_OVERFLOW");
        require(dataOffset <= startOffset + transactions.length, "TRANSACTIONS_OVERRUN");

        require(runningCount > messageCount, "EMPTY_BATCH");
        inboxAccs.push(runningAcc);
        messageCount = runningCount;

        require(runningAcc == afterAcc, "AFTER_ACC");
    }

    function calcL2Batch(
        uint256 beforeOffset,
        uint256[] calldata lengths,
        uint256 lengthsOffset,
        uint256 itemCount,
        bytes32 prefixHash,
        uint256 beforeCount,
        bytes32 beforeAcc
    )
        private
        pure
        returns (
            bytes32 acc,
            uint256 count,
            uint256 offset
        )
    {
        offset = beforeOffset;
        count = beforeCount;
        acc = beforeAcc;
        itemCount += lengthsOffset;
        for (uint256 i = lengthsOffset; i < itemCount; i++) {
            uint256 length = lengths[i];
            bytes32 messageDataHash;
            assembly {
                messageDataHash := keccak256(offset, length)
            }
            acc = keccak256(abi.encodePacked(acc, count, prefixHash, messageDataHash));
            offset += length;
            count++;
        }
        return (acc, count, offset);
    }

    // Precondition: _totalDelayedMessagesRead > totalDelayedMessagesRead
    function includeDelayedMessages(
        bytes32 acc,
        uint256 count,
        uint256 _totalDelayedMessagesRead,
        uint256 l1BlockNumber,
        uint256 timestamp,
        bytes32 delayedAcc
    ) private returns (bytes32, uint256) {
        require(_totalDelayedMessagesRead <= delayedInbox.messageCount(), "DELAYED_TOO_FAR");
        require(delayedAcc == delayedInbox.inboxAccs(_totalDelayedMessagesRead - 1), "DELAYED_ACC");
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
        return (acc, count);
    }

    /**
     * @notice Prove message count as of provided inbox state hash
     * @param proof proof data
     * @param offset offset for parsing proof data
     * @param inboxAcc target inbox state hash
     */
    function proveSeqBatchMsgCount(
        bytes calldata proof,
        uint256 offset,
        bytes32 inboxAcc
    ) internal pure returns (uint256, uint256) {
        uint256 endMessageCount;

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
        endMessageCount = seqNum + 1;
        require(buildingAcc == inboxAcc, "BATCH_ACC");

        return (offset, endMessageCount);
    }

    /**
     * @notice Show that given messageCount falls inside of some batch and prove/return inboxAcc state. This is used to ensure that the creation of new nodes are replay protected to the state of the inbox, thereby ensuring their validity/invalidy can't be modified upon reorging the inbox contents.
     * @dev (wrapper in leiu of proveBatchContainsSequenceNumber for sementics)
     * @return (message count at end of target batch, inbox hash as of target batch)
     */
    function proveInboxContainsMessage(bytes calldata proof, uint256 _messageCount)
        external
        view
        override
        returns (uint256, bytes32)
    {
        return proveInboxContainsMessageImp(proof, _messageCount);
    }

    // deprecated in favor of proveInboxContainsMessage
    function proveBatchContainsSequenceNumber(bytes calldata proof, uint256 _messageCount)
        external
        view
        returns (uint256, bytes32)
    {
        return proveInboxContainsMessageImp(proof, _messageCount);
    }

    function proveInboxContainsMessageImp(bytes calldata proof, uint256 _messageCount)
        internal
        view
        returns (uint256, bytes32)
    {
        if (_messageCount == 0) {
            return (0, 0);
        }

        (uint256 offset, uint256 targetInboxStateIndex) = Marshaling.deserializeInt(proof, 0);

        uint256 messageCountAsOfPreviousInboxState = 0;
        if (targetInboxStateIndex > 0) {
            (offset, messageCountAsOfPreviousInboxState) = proveSeqBatchMsgCount(
                proof,
                offset,
                inboxAccs[targetInboxStateIndex - 1]
            );
        }

        bytes32 targetInboxState = inboxAccs[targetInboxStateIndex];
        uint256 messageCountAsOfTargetInboxState;
        (offset, messageCountAsOfTargetInboxState) = proveSeqBatchMsgCount(
            proof,
            offset,
            targetInboxState
        );

        require(_messageCount > messageCountAsOfPreviousInboxState, "BATCH_START");
        require(_messageCount <= messageCountAsOfTargetInboxState, "BATCH_END");

        return (messageCountAsOfTargetInboxState, targetInboxState);
    }

    function getInboxAccsLength() external view override returns (uint256) {
        return inboxAccs.length;
    }
}
