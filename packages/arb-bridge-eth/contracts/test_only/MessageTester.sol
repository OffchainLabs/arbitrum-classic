/*
 * Copyright 2020, Offchain Labs, Inc.
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

import "../Messages.sol";

contract MessageTester {
    uint8 internal constant TRANSACTION_BATCH_MSG = 6;

    function transactionHash(
        address chain,
        address to,
        address from,
        uint256 seqNumber,
        uint256 value,
        bytes memory data,
        uint256 blockNumber,
        uint256 timestamp
    )
        public
        pure
        returns(bytes32)
    {
        return Messages.transactionHash(
            chain,
            to,
            from,
            seqNumber,
            value,
            keccak256(data),
            blockNumber,
            timestamp
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
        uint256 timestamp
    )
        public
        pure
        returns(bytes32)
    {
        return Messages.transactionMessageHash(
            chain,
            to,
            from,
            seqNumber,
            value,
            data,
            blockNumber,
            timestamp
        );
    }

    function transactionBatchHash(
        address chain,
        address[] memory tos,
        uint256[] memory seqNumbers,
        uint256[] memory values,
        uint32[] memory dataLengths,
        bytes memory data,
        bytes memory signatures,
        uint256 blockNumber,
        uint256 timestamp
    )
        public
        pure
        returns(bytes32)
    {
        return Messages.transactionBatchHash(
            chain,
            tos,
            seqNumbers,
            values,
            dataLengths,
            data,
            signatures,
            blockNumber,
            timestamp
        );
    }

    event TransactionBatchHashRet(
        bytes32 messageHash
    );

    function transactionBatchHash2(
        address chain,
        address[] calldata tos,
        uint256[] calldata seqNumbers,
        uint256[] calldata values,
        uint32[] calldata dataLengths,
        bytes calldata /* data */,
        bytes calldata /* signatures */

    )
        external
    {
        bytes32 messageHash;
        assembly {
            let ptr := mload(0x40)
            mstore8(ptr, TRANSACTION_BATCH_MSG)
            ptr := add(ptr, 1)
            calldatacopy(ptr, 4, sub(calldatasize, 4))
            ptr := add(ptr, sub(calldatasize, 4))
            mstore(ptr, number)
            ptr := add(ptr, 32)
            mstore(ptr, timestamp)
            ptr := add(ptr, 32)
            messageHash := keccak256(mload(0x40), sub(ptr, mload(0x40)))
        }
        emit TransactionBatchHashRet(messageHash);
    }

    function transactionMessageBatchHash(
        bytes32 prev,
        address chain,
        address[] memory tos,
        uint256[] memory seqNumbers,
        uint256[] memory values,
        uint32[] memory dataLengths,
        bytes memory data,
        bytes memory signatures,
        uint256 blockNumber,
        uint256 timestamp
    )
        public
        pure
        returns(bytes32)
    {
        return Messages.transactionMessageBatchHash(
            prev,
            chain,
            tos,
            seqNumbers,
            values,
            dataLengths,
            data,
            signatures,
            [blockNumber, timestamp]
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
        public
        pure
        returns(bytes32)
    {
        return Messages.ethHash(
            to,
            from,
            value,
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
        public
        pure
        returns(bytes32)
    {
        return Messages.ethMessageHash(
            to,
            from,
            value,
            blockNumber,
            timestamp,
            messageNum
        );
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
        public
        pure
        returns(bytes32)
    {
        return Messages.erc20Hash(
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
        public
        pure
        returns(bytes32)
    {
        return Messages.erc20MessageHash(
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
        public
        pure
        returns(bytes32)
    {
        return Messages.erc721Hash(
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
        public
        pure
        returns(bytes32)
    {
        return Messages.erc721MessageHash(
            to,
            from,
            erc721,
            id,
            blockNumber,
            timestamp,
            messageNum
        );
    }
}
