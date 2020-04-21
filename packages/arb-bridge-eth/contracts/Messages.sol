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


library Messages {
    uint8 internal constant TRANSACTION_MSG = 0;
    uint8 internal constant ETH_DEPOSIT = 1;
    uint8 internal constant ERC20_DEPOSIT = 2;
    uint8 internal constant ERC721_DEPOSIT = 3;
    uint8 internal constant CONTRACT_TRANSACTION_MSG = 4;

    using Value for Value.Data;

    function transactionHash(
        address chain,
        address to,
        address from,
        uint256 seqNumber,
        uint256 value,
        bytes memory data,
        uint256 blockNumber
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
                data,
                blockNumber
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
        uint256 blockNumber
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
                data
            )
        );

        Value.Data[] memory msgValues = new Value.Data[](4);
        msgValues[0] = Value.newInt(uint256(to));
        msgValues[1] = Value.newInt(seqNumber);
        msgValues[2] = Value.newInt(value);
        msgValues[3] = Value.bytesToBytestackHash(data);

        Value.Data[] memory msgType = new Value.Data[](3);
        msgType[0] = Value.newInt(TRANSACTION_MSG);
        msgType[1] = Value.newInt(uint256(from));
        msgType[2] = Value.newTuple(msgValues);

        Value.Data[] memory tup_data = new Value.Data[](3);
        tup_data[0] = Value.newInt(blockNumber);
        tup_data[1] = Value.newInt(uint256(txHash));
        tup_data[2] = Value.newTuple(msgType);
        Value.Data memory tuple = Value.newTuple(tup_data);

        return Value.hashTuple(tuple);
    }

    function transactionMessage(
        address chain,
        address to,
        address from,
        uint256 seqNumber,
        uint256 value,
        bytes memory data,
        uint256 blockNumber
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
                data
            )
        );

        Value.Data[] memory msgValues = new Value.Data[](4);
        msgValues[0] = Value.newInt(uint256(to));
        msgValues[1] = Value.newInt(seqNumber);
        msgValues[2] = Value.newInt(value);
        msgValues[3] = Value.bytesToBytestackHash(data);

        Value.Data[] memory msgType = new Value.Data[](3);
        msgType[0] = Value.newInt(TRANSACTION_MSG);
        msgType[1] = Value.newInt(uint256(from));
        msgType[2] = Value.newTuple(msgValues);

        Value.Data[] memory tup_data = new Value.Data[](3);
        tup_data[0] = Value.newInt(blockNumber);
        tup_data[1] = Value.newInt(uint256(txHash));
        tup_data[2] = Value.newTuple(msgType);

        return Value.newTuple(tup_data);
    }

    function ethHash(
        address to,
        address from,
        uint256 value,
        uint256 blockNumber,
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
                messageNum
            )
        );
    }

    function ethMessageHash(
        address to,
        address from,
        uint256 value,
        uint256 blockNumber,
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

        Value.Data[] memory ethMsg = new Value.Data[](3);
        ethMsg[0] = Value.newInt(blockNumber);
        ethMsg[1] = Value.newInt(messageNum);
        ethMsg[2] = Value.newTuple(msgType);

        return Value.newTuple(ethMsg).hash().hash;
    }

    function ethMessageValue(
        address to,
        address from,
        uint256 value,
        uint256 blockNumber,
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

        Value.Data[] memory ethMsg = new Value.Data[](3);
        ethMsg[0] = Value.newInt(blockNumber);
        ethMsg[1] = Value.newInt(messageNum);
        ethMsg[2] = Value.newTuple(msgType);

        return Value.newTuple(ethMsg);
    }

    function erc20Hash(
        address to,
        address from,
        address erc20,
        uint256 value,
        uint256 blockNumber,
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
            messageNum
        );
    }

    function erc20MessageValue(
        address to,
        address from,
        address erc20,
        uint256 value,
        uint256 blockNumber,
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
            messageNum
        );
    }

    function erc721Hash(
        address to,
        address from,
        address erc721,
        uint256 id,
        uint256 blockNumber,
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
            messageNum
        );
    }

    function erc721MessageValue(
        address to,
        address from,
        address erc721,
        uint256 id,
        uint256 blockNumber,
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
            messageNum
        );
    }

    function contractTransactionHash(
        address to,
        address from,
        uint256 value,
        bytes memory data,
        uint256 blockNumber,
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
                messageNum
            )
        );
    }

    function contractTransactionMessageValue(
        address to,
        address from,
        uint256 value,
        bytes memory data,
        uint256 blockNumber,
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

        Value.Data[] memory tup_data = new Value.Data[](3);
        tup_data[0] = Value.newInt(blockNumber);
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
                messageNum
            )
        );
    }

    function tokenMessageValue(
        uint8 messageType,
        address to,
        address from,
        address tokenContract,
        uint256 value,
        uint256 blockNumber,
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

        Value.Data[] memory ercTokenMsg = new Value.Data[](3);
        ercTokenMsg[0] = Value.newInt(blockNumber);
        ercTokenMsg[1] = Value.newInt(messageNum);
        ercTokenMsg[2] = Value.newTuple(msgType);

        return  Value.newTuple(ercTokenMsg);
    }
}
