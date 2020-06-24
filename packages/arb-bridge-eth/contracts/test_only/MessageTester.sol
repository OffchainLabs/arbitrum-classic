// SPDX-License-Identifier: Apache-2.0

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

pragma solidity ^0.5.11;

import "../Messages.sol";

contract MessageTester {
    uint8 internal constant TRANSACTION_BATCH_MSG = 6;

    function transactionHash(
        address chain,
        address to,
        address from,
        uint256 seqNumber,
        uint256 value,
        bytes memory data
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
            keccak256(data)
        );
    }

    function transactionMessageHash(
        address chain,
        address to,
        address from,
        uint256 seqNumber,
        uint256 value,
        bytes memory data
    )
        public
        pure
        returns(bytes32, bytes32)
    {
        (Value.Data memory tuple, bytes32 receiptHash) = Messages.transactionMessageValue(
            chain,
            to,
            from,
            seqNumber,
            value,
            keccak256(data),
            Value.bytesToBytestackHash(data, 0, data.length)
        );

        return (Value.hash(tuple), receiptHash);
    }

    function transactionBatchHash(
        bytes memory transactions
    )
        public
        pure
        returns(bytes32)
    {
        return Messages.transactionBatchHash(transactions);
    }

    function transactionMessageBatchHashSingle(
        uint256 start,
        address chain,
        bytes memory transactions
    )
        public
        pure
        returns(bytes32, bytes32, bool)
    {
        (Value.Data memory message, bytes32 receiptHash, bool valid) = Messages.transactionMessageBatchHashSingle(
            start,
            chain,
            transactions
        );
        return (Value.hash(message), receiptHash, valid);
    }

    function transactionMessageBatchSingleSender(
        uint256 start,
        address chain,
        bytes32 dataHash,
        bytes memory transactions
    )
        public
        pure
        returns(address)
    {
        return Messages.transactionMessageBatchSingleSender(
            start,
            chain,
            dataHash,
            transactions
        );
    }

    function transactionMessageBatchHash(
        bytes32 prev,
        uint256 prevSize,
        address chain,
        bytes memory transactions,
        uint256 blockNum,
        uint256 blockTimestamp
    )
        public
        pure
        returns(bytes32)
    {
        return Messages.transactionMessageBatchHash(
            prev,
            prevSize,
            chain,
            transactions,
            blockNum,
            blockTimestamp
        );
    }

    function ethHash(
        address to,
        address from,
        uint256 value
    )
        public
        pure
        returns(bytes32)
    {
        return Messages.ethHash(
            to,
            from,
            value
        );
    }

    function ethMessageHash(
        address to,
        address from,
        uint256 value
    )
        public
        pure
        returns(bytes32)
    {
        return Value.hash(Messages.ethMessageValue(
            to,
            from,
            value
        ));
    }

    function erc20Hash(
        address to,
        address from,
        address erc20,
        uint256 value
    )
        public
        pure
        returns(bytes32)
    {
        return Messages.erc20Hash(
            to,
            from,
            erc20,
            value
        );
    }

    function erc20MessageHash(
        address to,
        address from,
        address erc20,
        uint256 value
    )
        public
        pure
        returns(bytes32)
    {
        return Value.hash(Messages.erc20MessageValue(
            to,
            from,
            erc20,
            value
        ));
    }

    function erc721Hash(
        address to,
        address from,
        address erc721,
        uint256 id
    )
        public
        pure
        returns(bytes32)
    {
        return Messages.erc721Hash(
            to,
            from,
            erc721,
            id
        );
    }

    function erc721MessageHash(
        address to,
        address from,
        address erc721,
        uint256 id
    )
        public
        pure
        returns(bytes32)
    {
        return Value.hash(Messages.erc721MessageValue(
            to,
            from,
            erc721,
            id
        ));
    }

    function addMessageToInbox(
        bytes32 inboxHash,
        bytes32 messageHash,
        uint256 blockNumber,
        uint256 timestamp,
        uint256 messageNum
    )
        public
        pure
        returns(bytes32)
    {
        return Messages.addMessageToInbox(
            inboxHash,
            messageHash,
            blockNumber,
            timestamp,
            messageNum
        );
    }

    function addMessageToVMInboxHash(
        bytes32 inboxTuplePreimage,
        uint256 inboxTupleSize,
        uint256 blockNumber,
        uint256 timestamp,
        uint256 txId,
        bytes32 messageTuplePreimage,
        uint256 messageTupleSize
    )
        public
        pure
        returns(bytes32)
    {
        return Value.hash(
            Messages.addMessageToVMInboxHash(
                Value.newTuplePreImage(inboxTuplePreimage, inboxTupleSize),
                blockNumber,
                timestamp,
                txId,
                Value.newTuplePreImage(messageTuplePreimage, messageTupleSize)
            )
        );
    }
}
