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
import "./libraries/SigUtils.sol";
import "./arch/Protocol.sol";

library Messages {
    uint8 internal constant TRANSACTION_MSG = 0;
    uint8 internal constant ETH_DEPOSIT = 1;
    uint8 internal constant ERC20_DEPOSIT = 2;
    uint8 internal constant ERC721_DEPOSIT = 3;
    uint8 internal constant CONTRACT_TRANSACTION_MSG = 4;
    uint8 internal constant CALL_MSG = 5;
    uint8 internal constant TRANSACTION_BATCH_MSG = 6;

    using Value for Value.Data;
    using BytesLib for bytes;

    function addDeliveredMessageToInbox(bytes32 inbox, bytes32 message)
        internal
        pure
        returns (bytes32)
    {
        return keccak256(abi.encodePacked(inbox, message));
    }

    function addMessageToInbox(
        bytes32 inbox,
        bytes32 messageHash,
        uint256 blockNumber,
        uint256 timestamp,
        uint256 messageNum
    ) internal pure returns (bytes32) {
        return
            addDeliveredMessageToInbox(
                inbox,
                keccak256(
                    abi.encodePacked(
                        messageHash,
                        blockNumber,
                        timestamp,
                        messageNum
                    )
                )
            );
    }

    function addMessageToVMInboxHash(
        Value.Data memory vmInboxHashValue,
        uint256 blockNumber,
        uint256 timestamp,
        uint256 txId,
        Value.Data memory message
    ) internal pure returns (Value.Data memory) {
        Value.Data[] memory tupData = new Value.Data[](4);
        tupData[0] = Value.newInt(blockNumber);
        tupData[1] = Value.newInt(timestamp);
        tupData[2] = Value.newInt(txId);
        tupData[3] = message;

        Value.Data[] memory vals = new Value.Data[](2);
        vals[0] = vmInboxHashValue;
        vals[1] = Value.newTuple(tupData);
        return Value.newTuple(vals);
    }

    function transactionHash(
        address chain,
        address to,
        address from,
        uint256 seqNumber,
        uint256 value,
        bytes32 dataHash
    ) internal pure returns (bytes32) {
        return
            keccak256(
                abi.encodePacked(
                    TRANSACTION_MSG,
                    chain,
                    to,
                    from,
                    seqNumber,
                    value,
                    dataHash
                )
            );
    }

    function transactionReceiptHash(
        address chain,
        address to,
        address from,
        uint256 seqNumber,
        uint256 value,
        bytes32 dataHash
    ) internal pure returns (bytes32) {
        return
            keccak256(
                abi.encodePacked(
                    TRANSACTION_MSG,
                    chain,
                    to,
                    from,
                    seqNumber,
                    value,
                    dataHash
                )
            );
    }

    function transactionMessageValue(
        address chain,
        address to,
        address from,
        uint256 seqNumber,
        uint256 value,
        bytes32 dataHash,
        Value.Data memory dataTuple
    ) internal pure returns (Value.Data memory, bytes32) {
        Value.Data[] memory msgValues = new Value.Data[](4);
        msgValues[0] = Value.newInt(uint256(to));
        msgValues[1] = Value.newInt(seqNumber);
        msgValues[2] = Value.newInt(value);
        msgValues[3] = dataTuple;

        Value.Data memory message = messageOuterLayer(
            TRANSACTION_MSG,
            from,
            Value.newTuple(msgValues)
        );
        bytes32 txHash = transactionReceiptHash(
            chain,
            to,
            from,
            seqNumber,
            value,
            dataHash
        );

        return (message, txHash);
    }

    function transactionBatchHash(bytes memory transactions)
        internal
        pure
        returns (bytes32)
    {
        return keccak256(abi.encodePacked(TRANSACTION_BATCH_MSG, transactions));
    }

    uint256 internal constant TO_OFFSET = 2;
    uint256 internal constant SEQ_OFFSET = 22;
    uint256 internal constant VALUE_OFFSET = 54;
    uint256 internal constant SIG_OFFSET = 86;
    uint256 internal constant DATA_OFFSET = 151;

    function transactionMessageBatchHash(
        bytes32 prevHashImage,
        uint256 prevHashSize,
        address chain,
        bytes memory transactions,
        uint256 blockNum,
        uint256 blockTimestamp
    ) internal pure returns (bytes32) {
        Value.Data memory prevVal = Value.newTuplePreImage(
            prevHashImage,
            prevHashSize
        );
        uint256 transactionsLength = transactions.length;
        uint256 start = 0x00;
        // Continue until we run out of enough data for a tx with no data
        while (start + DATA_OFFSET < transactionsLength) {
            uint16 dataLength = transactions.toUint16(start);
            // Terminate if the input data is shorter than the fixed length + claimed data length
            if (start + DATA_OFFSET + dataLength > transactionsLength) {
                return Value.hash(prevVal);
            }

            (
                Value.Data memory message,
                bytes32 receiptHash,
                bool valid
            ) = transactionMessageBatchHashSingle(start, chain, transactions);

            if (valid) {
                prevVal = addMessageToVMInboxHash(
                    prevVal,
                    blockNum,
                    blockTimestamp,
                    uint256(receiptHash),
                    message
                );
            }
            start += DATA_OFFSET + dataLength;
        }
        return Value.hash(prevVal);
    }

    function transactionMessageBatchHashSingle(
        uint256 start,
        address chain,
        bytes memory transactions
    )
        internal
        pure
        returns (
            Value.Data memory message,
            bytes32 receiptHash,
            bool valid
        )
    {
        bytes32 dataHash = keccak256Subset(
            transactions,
            start + DATA_OFFSET,
            transactions.toUint16(start)
        );

        address from = transactionMessageBatchSingleSender(
            start,
            chain,
            dataHash,
            transactions
        );
        valid = from != address(0);
        if (!valid) {
            // Signature was invalid so we'll ignore this message and keep processing
            return (message, receiptHash, valid);
        }

        Value.Data memory dataVal = Value.bytesToBytestackHash(
            transactions,
            start + DATA_OFFSET,
            transactions.toUint16(start)
        );

        (message, receiptHash) = transactionMessageValue(
            chain,
            transactions.toAddress(start + TO_OFFSET),
            from,
            transactions.toUint(start + SEQ_OFFSET),
            transactions.toUint(start + VALUE_OFFSET),
            dataHash,
            dataVal
        );
        return (message, receiptHash, valid);
    }

    function transactionMessageBatchSingleSender(
        uint256 start,
        address chain,
        bytes32 dataHash,
        bytes memory transactions
    ) internal pure returns (address) {
        return
            SigUtils.recoverAddressFromData(
                keccak256(
                    abi.encodePacked(
                        chain,
                        transactions.toAddress(start + TO_OFFSET),
                        transactions.toUint(start + SEQ_OFFSET),
                        transactions.toUint(start + VALUE_OFFSET),
                        dataHash
                    )
                ),
                transactions,
                start + SIG_OFFSET
            );
    }

    function ethHash(
        address to,
        address from,
        uint256 value
    ) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(ETH_DEPOSIT, to, from, value));
    }

    function ethMessageValue(
        address to,
        address from,
        uint256 value
    ) internal pure returns (Value.Data memory) {
        Value.Data[] memory msgValues = new Value.Data[](2);
        msgValues[0] = Value.newInt(uint256(to));
        msgValues[1] = Value.newInt(value);

        return messageOuterLayer(ETH_DEPOSIT, from, Value.newTuple(msgValues));
    }

    function erc20Hash(
        address to,
        address from,
        address erc20,
        uint256 value
    ) internal pure returns (bytes32) {
        return tokenHash(ERC20_DEPOSIT, to, from, erc20, value);
    }

    function erc20MessageValue(
        address to,
        address from,
        address erc20,
        uint256 value
    ) internal pure returns (Value.Data memory) {
        return tokenMessageValue(ERC20_DEPOSIT, to, from, erc20, value);
    }

    function erc721Hash(
        address to,
        address from,
        address erc721,
        uint256 id
    ) internal pure returns (bytes32) {
        return tokenHash(ERC721_DEPOSIT, to, from, erc721, id);
    }

    function erc721MessageValue(
        address to,
        address from,
        address erc721,
        uint256 id
    ) internal pure returns (Value.Data memory) {
        return tokenMessageValue(ERC721_DEPOSIT, to, from, erc721, id);
    }

    function contractTransactionHash(
        address to,
        address from,
        uint256 value,
        bytes memory data
    ) internal pure returns (bytes32) {
        return
            keccak256(
                abi.encodePacked(
                    CONTRACT_TRANSACTION_MSG,
                    to,
                    from,
                    value,
                    data
                )
            );
    }

    function contractTransactionMessageValue(
        address to,
        address from,
        uint256 value,
        bytes memory data
    ) internal pure returns (Value.Data memory) {
        Value.Data[] memory msgValues = new Value.Data[](3);
        msgValues[0] = Value.newInt(uint256(to));
        msgValues[2] = Value.newInt(value);
        msgValues[3] = Value.bytesToBytestackHash(data);

        return
            messageOuterLayer(
                CONTRACT_TRANSACTION_MSG,
                from,
                Value.newTuple(msgValues)
            );
    }

    function tokenHash(
        uint8 messageType,
        address to,
        address from,
        address tokenContract,
        uint256 value
    ) private pure returns (bytes32) {
        return
            keccak256(
                abi.encodePacked(messageType, to, from, tokenContract, value)
            );
    }

    function tokenMessageValue(
        uint8 messageType,
        address to,
        address from,
        address tokenContract,
        uint256 value
    ) private pure returns (Value.Data memory) {
        Value.Data[] memory msgValues = new Value.Data[](3);
        msgValues[0] = Value.newInt(uint256(tokenContract));
        msgValues[1] = Value.newInt(uint256(to));
        msgValues[2] = Value.newInt(value);

        return messageOuterLayer(messageType, from, Value.newTuple(msgValues));
    }

    function messageOuterLayer(
        uint256 messageType,
        address from,
        Value.Data memory dataTuple
    ) private pure returns (Value.Data memory) {
        Value.Data[] memory msgType = new Value.Data[](3);
        msgType[0] = Value.newInt(messageType);
        msgType[1] = Value.newInt(uint256(from));
        msgType[2] = dataTuple;

        return Value.newTuple(msgType);
    }

    function keccak256Subset(
        bytes memory data,
        uint256 start,
        uint256 length
    ) private pure returns (bytes32 dataHash) {
        assembly {
            dataHash := keccak256(add(add(data, 0x20), start), length)
        }
    }
}
