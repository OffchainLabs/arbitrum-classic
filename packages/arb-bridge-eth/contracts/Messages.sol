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

    function deliveredMessageHash(
        bytes32 messageHash,
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
                messageHash,
                blockNumber,
                timestamp,
                messageNum
            )
        );
    }

    function transactionHash(
        address chain,
        address to,
        address from,
        uint256 seqNumber,
        uint256 value,
        bytes32 dataHash,
        uint256 blockNumber,
        uint256 timestamp,
        uint256 messageNum
    )
        internal
        pure
        returns(bytes32)
    {
        return deliveredMessageHash(
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
            ),
            blockNumber,
            timestamp,
            messageNum
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
        Value.Data memory tuple =  transactionMessage(
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

        return Value.hashTuple(tuple);
    }

    function transactionMessage(
        address chain,
        address to,
        address from,
        uint256 seqNumber,
        uint256 value,
        bytes32 dataHash,
        Value.Data memory dataTuple,
        uint256 blockNumber,
        uint256 timestamp
    )
        internal
        pure
        returns(Value.Data memory)
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
        msgValues[3] = dataTuple;

        Value.Data[] memory msgType = new Value.Data[](3);
        msgType[0] = Value.newInt(TRANSACTION_MSG);
        msgType[1] = Value.newInt(uint256(from));
        msgType[2] = Value.newTuple(msgValues);

        Value.Data[] memory tup_data = new Value.Data[](4);
        tup_data[0] = Value.newInt(blockNumber);
        tup_data[1] = Value.newInt(timestamp);
        tup_data[2] = Value.newInt(uint256(txHash));
        tup_data[3] = Value.newTuple(msgType);

        return Value.newTuple(tup_data);
    }

    function transactionBatchHash(
        bytes memory transactions,
        uint256 blockNum,
        uint256 blockTimestamp,
        uint256 messageNum
    )
        internal
        pure
        returns(bytes32)
    {
        return deliveredMessageHash(
            keccak256(
                abi.encodePacked(
                    TRANSACTION_BATCH_MSG,
                    transactions
                )
            ),
            blockNum,
            blockTimestamp,
            messageNum
        );
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
    )
        internal
        pure
        returns(bytes32)
    {
        Value.Data memory prevVal = Value.newTuplePreImage(prevHashImage, prevHashSize);
        uint256 transactionsLength = transactions.length;
        uint256 start = 0x00;
        // Continue until we run out of enough data for a tx with no data
        while (start + DATA_OFFSET < transactionsLength) {
            uint16 dataLength = transactions.toUint16(start);
            // Terminate if the input data is shorter than the fixed length + claimed data length
            if (start + DATA_OFFSET + dataLength > transactionsLength) {
                return Value.hash(prevVal);
            }

            Value.Data memory message = transactionMessageBatchHashSingle(
                start,
                chain,
                transactions,
                blockNum,
                blockTimestamp
            );

            prevVal = Protocol.addMessageToVMInboxHash(prevVal, message);
            start += DATA_OFFSET + dataLength;
        }
        return Value.hash(prevVal);
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
        returns(Value.Data memory)
    {
        bytes32 dataHash = keccak256Subset(transactions, start + DATA_OFFSET, transactions.toUint16(start));

        address from = transactionMessageBatchSingleSender(
            start,
            chain,
            dataHash,
            transactions
        );
        require(from != address(0), "invalid sig");

        Value.Data memory dataTup = Value.bytesToBytestackHash(
                transactions, 
                start + DATA_OFFSET, 
                transactions.toUint16(start)
            );

        return transactionMessage(
            chain,
            transactions.toAddress(start + TO_OFFSET),
            from,
            transactions.toUint(start + SEQ_OFFSET),
            transactions.toUint(start + VALUE_OFFSET),
            dataHash,
            dataTup,
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
        return deliveredMessageHash(
            keccak256(
                abi.encodePacked(
                    ETH_DEPOSIT,
                    to,
                    from,
                    value
                )
            ),
            blockNumber,
            timestamp,
            messageNum
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
        Value.Data memory tuple = ethMessageValue(
            to, 
            from, 
            value, 
            blockNumber, 
            timestamp,
            messageNum);
        
        return Value.hashTuple(tuple);
    }

    function ethMessageValue(
        address to,
        address from,
        uint256 value,
        uint256 blockNumber,
        uint256 timestamp,
        uint256 messageNum
    )
        internal
        pure
        returns(Value.Data memory)
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

        return Value.newTuple(ethMsg);
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

    function erc20MessageValue(
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
        returns(Value.Data memory)
    {
        return tokenMessageValue(
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

    function erc721MessageValue(
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
        returns(Value.Data memory)
    {
        return tokenMessageValue(
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
        return deliveredMessageHash(
            keccak256(
                abi.encodePacked(
                    CONTRACT_TRANSACTION_MSG,
                    to,
                    from,
                    value,
                    data
                )
            ),
            blockNumber,
            timestamp,
            messageNum
        );
    }

    function contractTransactionMessageValue(
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
        returns(Value.Data memory)
    {
        Value.Data[] memory msgValues = new Value.Data[](3);
        msgValues[0] = Value.newInt(uint256(to));
        msgValues[2] = Value.newInt(value);
        msgValues[3] = Value.bytesToBytestackHash(data);

        Value.Data[] memory msgType = new Value.Data[](3);
        msgType[0] = Value.newInt(CONTRACT_TRANSACTION_MSG);
        msgType[1] = Value.newInt(uint256(from));
        msgType[2] = Value.newTuple(msgValues);

        Value.Data[] memory tup_data = new Value.Data[](4);
        tup_data[0] = Value.newInt(blockNumber);
        tup_data[1] = Value.newInt(timestamp);
        tup_data[1] = Value.newInt(messageNum);
        tup_data[2] = Value.newTuple(msgType);

        return Value.newTuple(tup_data);
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
        return deliveredMessageHash(
            keccak256(
                abi.encodePacked(
                    messageType,
                    to,
                    from,
                    tokenContract,
                    value
                )
            ),
            blockNumber,
            timestamp,
            messageNum
        );
    }

    function tokenMessageValue(
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
        returns(Value.Data memory)
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

        return  Value.newTuple(ercTokenMsg);
    }
}
