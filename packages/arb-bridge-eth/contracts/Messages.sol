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

pragma solidity ^0.5.3;

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

    function transactionHash(
        address chain,
        address to,
        address from,
        uint256 seqNumber,
        uint256 value,
        bytes32 dataHash,
        uint256 blockNumber,
        uint256 timestamp
    )
        internal
        pure
        returns(bytes32)
    {
        return keccak256(
            abi.encodePacked(
                TRANSACTION_MSG,
                chain,
                to,
                from,
                seqNumber,
                value,
                dataHash,
                blockNumber,
                timestamp
            )
        );
    }


    function transactionMessageHash(
        address chain,
        address to,
        address from,
        uint256 seqNumber,
        uint256 value,
        bytes memory data,
        uint256 blockNumber,
        uint256 blockTimestamp
    )
        internal
        pure
        returns(bytes32)
    {
        return transactionMessageHash(
            chain,
            to,
            from,
            seqNumber,
            value,
            keccak256(data),
            Value.bytesToBytestackHash(data, 0, data.length),
            blockNumber,
            blockTimestamp
        );
    }

    function transactionMessageHash(
        address chain,
        address to,
        address from,
        uint256 seqNumber,
        uint256 value,
        bytes32 dataHash,
        bytes32 dataTupleHash,
        uint256 blockNumber,
        uint256 timestamp
    )
        internal
        pure
        returns(bytes32)
    {
        bytes32 txHash = keccak256(
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
        Value.Data[] memory msgValues = new Value.Data[](4);
        msgValues[0] = Value.newInt(uint256(to));
        msgValues[1] = Value.newInt(seqNumber);
        msgValues[2] = Value.newInt(value);
        msgValues[3] = Value.newHashOnly(dataTupleHash);

        Value.Data[] memory msgType = new Value.Data[](3);
        msgType[0] = Value.newInt(TRANSACTION_MSG);
        msgType[1] = Value.newInt(uint256(from));
        msgType[2] = Value.newTuple(msgValues);

        return Value.hashTuple([
            Value.newInt(blockNumber),
            Value.newInt(timestamp),
            Value.newInt(uint256(txHash)),
            Value.newTuple(msgType)
        ]);
    }

    function transactionBatchHash(
        bytes memory transactions,
        uint256 blockNum,
        uint256 blockTimestamp
    )
        internal
        pure
        returns(bytes32)
    {
        return keccak256(
            abi.encodePacked(
                TRANSACTION_BATCH_MSG,
                transactions,
                blockNum,
                blockTimestamp
            )
        );
    }

    uint256 internal constant TO_OFFSET = 2;
    uint256 internal constant SEQ_OFFSET = 22;
    uint256 internal constant VALUE_OFFSET = 54;
    uint256 internal constant SIG_OFFSET = 86;
    uint256 internal constant DATA_OFFSET = 151;

    function transactionMessageBatchHash(
        bytes32 prev,
        address chain,
        bytes memory transactions,
        uint256 blockNum,
        uint256 blockTimestamp
    )
        internal
        pure
        returns(bytes32)
    {
        uint256 transactionsLength = transactions.length;
        uint256 start = 0x00;
        // Continue until we run out of enough data for a tx with no data
        while (start + DATA_OFFSET < transactionsLength) {
            uint16 dataLength = transactions.toUint16(start);
            // Terminate if the input data is shorter than the fixed length + claimed data length
            if (start + DATA_OFFSET + dataLength > transactionsLength) {
                return prev;
            }

            bytes32 messageHash = transactionMessageBatchHashSingle(
                start,
                chain,
                transactions,
                blockNum,
                blockTimestamp
            );

            prev = Protocol.addMessageToVMInbox(prev, messageHash);
            start += DATA_OFFSET + dataLength;
        }
        return prev;
    }

    function keccak256Subset(bytes memory data, uint256 start, uint256 length) internal pure returns(bytes32 dataHash) {
        assembly {
            dataHash := keccak256(add(add(data, 0x20), start), length)
        }
    }

    function transactionMessageBatchHashSingle(
        uint256 start,
        address chain,
        bytes memory transactions,
        uint256 blockNum,
        uint256 blockTimestamp
    )
        internal
        pure
        returns(bytes32)
    {
        bytes32 dataHash = keccak256Subset(transactions, start + DATA_OFFSET, transactions.toUint16(start));

        address from = transactionMessageBatchSingleSender(
            start,
            chain,
            dataHash,
            transactions
        );
        require(from != address(0), "invalid sig");

        bytes32 dataTupHash = Value.bytesToBytestackHash(transactions, start + DATA_OFFSET, transactions.toUint16(start));

        return transactionMessageHash(
            chain,
            transactions.toAddress(start + TO_OFFSET),
            from,
            transactions.toUint(start + SEQ_OFFSET),
            transactions.toUint(start + VALUE_OFFSET),
            dataHash,
            dataTupHash,
            blockNum,
            blockTimestamp
        );
    }

    function transactionMessageBatchSingleSender(
        uint256 start,
        address chain,
        bytes32 dataHash,
        bytes memory transactions
    )
        internal
        pure
        returns(address)
    {
        return SigUtils.recoverAddressFromData(
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
        uint256 value,
        uint256 blockNumber,
        uint256 timestamp,
        uint256 messageNum
    )
        internal
        pure
        returns(bytes32)
    {
        return keccak256(
            abi.encodePacked(
                ETH_DEPOSIT,
                to,
                from,
                value,
                blockNumber,
                timestamp,
                messageNum
            )
        );
    }

    function ethMessageHash(
        address to,
        address from,
        uint256 value,
        uint256 blockNumber,
        uint256 timestamp,
        uint256 messageNum
    )
        internal
        pure
        returns(bytes32)
    {
        Value.Data[] memory msgValues = new Value.Data[](2);
        msgValues[0] = Value.newInt(uint256(to));
        msgValues[1] = Value.newInt(value);

        Value.Data[] memory msgType = new Value.Data[](3);
        msgType[0] = Value.newInt(ETH_DEPOSIT);
        msgType[1] = Value.newInt(uint256(from));
        msgType[2] = Value.newTuple(msgValues);

        Value.Data[] memory ethMsg = new Value.Data[](4);
        ethMsg[0] = Value.newInt(blockNumber);
        ethMsg[1] = Value.newInt(timestamp);
        ethMsg[2] = Value.newInt(messageNum);
        ethMsg[3] = Value.newTuple(msgType);

        return Value.newTuple(ethMsg).hash().hash;
    }

    function erc20Hash(
        address to,
        address from,
        address erc20,
        uint256 value,
        uint256 blockNumber,
        uint256 timestamp,
        uint256 messageNum
    )
        internal
        pure
        returns(bytes32)
    {
        return tokenHash(
            ERC20_DEPOSIT,
            to,
            from,
            erc20,
            value,
            blockNumber,
            timestamp,
            messageNum
        );
    }

    function erc20MessageHash(
        address to,
        address from,
        address erc20,
        uint256 value,
        uint256 blockNumber,
        uint256 timestamp,
        uint256 messageNum
    )
        internal
        pure
        returns(bytes32)
    {
        return tokenMessageHash(
            ERC20_DEPOSIT,
            to,
            from,
            erc20,
            value,
            blockNumber,
            timestamp,
            messageNum
        );
    }

    function erc721Hash(
        address to,
        address from,
        address erc721,
        uint256 id,
        uint256 blockNumber,
        uint256 timestamp,
        uint256 messageNum
    )
        internal
        pure
        returns(bytes32)
    {
        return tokenHash(
            ERC721_DEPOSIT,
            to,
            from,
            erc721,
            id,
            blockNumber,
            timestamp,
            messageNum
        );
    }

    function erc721MessageHash(
        address to,
        address from,
        address erc721,
        uint256 id,
        uint256 blockNumber,
        uint256 timestamp,
        uint256 messageNum
    )
        internal
        pure
        returns(bytes32)
    {
        return tokenMessageHash(
            ERC721_DEPOSIT,
            to,
            from,
            erc721,
            id,
            blockNumber,
            timestamp,
            messageNum
        );
    }

    function contractTransactionHash(
        address to,
        address from,
        uint256 value,
        bytes memory data,
        uint256 blockNumber,
        uint256 timestamp,
        uint256 messageNum
    )
        internal
        pure
        returns(bytes32)
    {
        return keccak256(
            abi.encodePacked(
                CONTRACT_TRANSACTION_MSG,
                to,
                from,
                value,
                data,
                blockNumber,
                timestamp,
                messageNum
            )
        );
    }

    function contractTransactionMessageHash(
        address to,
        address from,
        uint256 value,
        bytes memory data,
        uint256 blockNumber,
        uint256 timestamp,
        uint256 messageNum
    )
        internal
        pure
        returns(bytes32)
    {
        bytes32 dataHash = Value.bytesToBytestackHash(data, 0, data.length);
        Value.Data[] memory msgValues = new Value.Data[](3);
        msgValues[0] = Value.newInt(uint256(to));
        msgValues[2] = Value.newInt(value);
        msgValues[3] = Value.newHashOnly(dataHash);

        Value.Data[] memory msgType = new Value.Data[](3);
        msgType[0] = Value.newInt(CONTRACT_TRANSACTION_MSG);
        msgType[1] = Value.newInt(uint256(from));
        msgType[2] = Value.newTuple(msgValues);

        return Value.hashTuple([
            Value.newInt(blockNumber),
            Value.newInt(timestamp),
            Value.newInt(messageNum),
            Value.newTuple(msgType)
        ]);
    }

    function tokenHash(
        uint8 messageType,
        address to,
        address from,
        address tokenContract,
        uint256 value,
        uint256 blockNumber,
        uint256 timestamp,
        uint256 messageNum
    )
        private
        pure
        returns(bytes32)
    {
        return keccak256(
            abi.encodePacked(
                messageType,
                to,
                from,
                tokenContract,
                value,
                blockNumber,
                timestamp,
                messageNum
            )
        );
    }

    function tokenMessageHash(
        uint8 messageType,
        address to,
        address from,
        address tokenContract,
        uint256 value,
        uint256 blockNumber,
        uint256 timestamp,
        uint256 messageNum
    )
        private
        pure
        returns(bytes32)
    {
        Value.Data[] memory msgValues = new Value.Data[](3);
        msgValues[0] = Value.newInt(uint256(tokenContract));
        msgValues[1] = Value.newInt(uint256(to));
        msgValues[2] = Value.newInt(value);

        Value.Data[] memory msgType = new Value.Data[](3);
        msgType[0] = Value.newInt(messageType);
        msgType[1] = Value.newInt(uint256(from));
        msgType[2] = Value.newTuple(msgValues);

        Value.Data[] memory ercTokenMsg = new Value.Data[](4);
        ercTokenMsg[0] = Value.newInt(blockNumber);
        ercTokenMsg[1] = Value.newInt(timestamp);
        ercTokenMsg[2] = Value.newInt(messageNum);
        ercTokenMsg[3] = Value.newTuple(msgType);

        return  Value.newTuple(ercTokenMsg).hash().hash;
    }
}
