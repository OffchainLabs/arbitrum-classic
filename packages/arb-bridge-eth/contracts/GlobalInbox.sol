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

pragma solidity ^0.5.3;

import "./GlobalEthWallet.sol";
import "./GlobalFTWallet.sol";
import "./GlobalNFTWallet.sol";
import "./IGlobalInbox.sol";
import "./Messages.sol";
import "./PaymentRecords.sol";

import "./arch/Protocol.sol";
import "./arch/Value.sol";

import "./libraries/SigUtils.sol";
import "./libraries/BytesLib.sol";

contract GlobalInbox is GlobalEthWallet, GlobalFTWallet, GlobalNFTWallet, IGlobalInbox, PaymentRecords {
    uint8 internal constant TRANSACTION_MSG = 0;
    uint8 internal constant ETH_DEPOSIT = 1;
    uint8 internal constant ERC20_DEPOSIT = 2;
    uint8 internal constant ERC721_DEPOSIT = 3;

    uint8 internal constant TRANSACTION_BATCH_MSG = 6;

    using Value for Value.Data;

    address internal constant ETH_ADDRESS = address(0);

    struct Inbox {
        bytes32 value;
        uint256 count;
    }

    mapping(address => Inbox) inboxes;

    function getInbox(address account) external view returns(bytes32, uint) {
        Inbox storage inbox = inboxes[account];
        return (inbox.value, inbox.count);
    }

    function sendMessages(bytes calldata _messages, uint[] calldata messageCounts, bytes32[] calldata nodeHashes) external {
        bool valid;
        uint256 offset = 0;
        uint256 messageType;
        address sender;

        uint256 nodeCount = nodeHashes.length;
        for (uint256 i = 0; i < nodeCount; i++) {
            for (uint256 j = 0; j < messageCounts[i]; j++) {
                (   valid,
                    offset,
                    messageType,
                    sender
                ) = Value.deserializeMessageData(_messages, offset);
                if (!valid) {
                    return;
                }
                (valid, offset) = sendDeserializedMsg(nodeHashes[i], j, _messages, offset, messageType);
                if (!valid) {
                    return;
                }
            }
        }
    }

    function sendDeserializedMsg(
        bytes32 nodeHash,
        uint256 messageIndex,
        bytes memory _messages,
        uint256 startOffset,
        uint256 messageType
    )
        private
        returns(
            bool, // valid
            uint256 // offset
        )
    {
        if (messageType == ETH_DEPOSIT) {
            (
                bool valid,
                uint256 offset,
                address to,
                uint256 value
            ) = Value.getEthMsgData(_messages, startOffset);

            if (!valid) {
                return (false, startOffset);
            }

            address paymentOwner = getPaymentOwner(to, nodeHash, messageIndex);
            transferEth(msg.sender, paymentOwner, value);
            delete paymentMap[nodeHash][messageIndex][to];

            return (true, offset);
        } else if (messageType == ERC20_DEPOSIT) {
            (
                bool valid,
                uint256 offset,
                address erc20,
                address to,
                uint256 value
            ) = Value.getERCTokenMsgData(_messages, startOffset);
            if (!valid) {
                return (false, startOffset);
            }

            address paymentOwner = getPaymentOwner(to, nodeHash, messageIndex);
            transferERC20(msg.sender, paymentOwner, erc20, value);
            delete paymentMap[nodeHash][messageIndex][to];

            return (true, offset);
        } else if (messageType == ERC721_DEPOSIT) {
            (
                bool valid,
                uint256 offset,
                address erc721,
                address to,
                uint256 value
            ) = Value.getERCTokenMsgData(_messages, startOffset);
            if (!valid) {
                return (false, startOffset);
            }

            address paymentOwner = getPaymentOwner(to, nodeHash, messageIndex);
            transferNFT(msg.sender, paymentOwner, erc721, value);
            delete paymentMap[nodeHash][messageIndex][to];

            return (true, offset);
        } else {
            return (false, startOffset);
        }
    }

    function sendTransactionMessage(
        address _chain,
        address _to,
        uint256 _seqNumber,
        uint256 _value,
        bytes calldata _data
    )
        external
    {
        _deliverTransactionMessage(
            _chain,
            _to,
            msg.sender,
            _seqNumber,
            _value,
            _data
        );
    }

    function depositEthMessage(address _chain, address _to) external payable {
        depositEth(_chain);

        _deliverEthMessage(
            _chain,
            _to,
            msg.sender,
            msg.value
        );
    }

    function depositERC20Message(
        address _chain,
        address _to,
        address _erc20,
        uint256 _value
    )
        external
    {
        depositERC20(_erc20, _chain, _value);

        _deliverERC20TokenMessage(
            _chain,
            _to,
            msg.sender,
            _erc20,
            _value
        );
    }

    function depositERC721Message(
        address _chain,
        address _to,
        address _erc721,
        uint256 _id
    )
        external
    {
        depositERC721(_erc721, _chain, _id);

        _deliverERC721TokenMessage(
            _chain,
            _to,
            msg.sender,
            _erc721,
            _id
        );
    }

    function forwardContractTransactionMessage(
        address _to,
        address _from,
        uint256 _value,
        bytes calldata _data
    )
        external
    {
        _deliverContractTransactionMessage(
            msg.sender,
            _to,
            _from,
            _value,
            _data
        );
    }

    function forwardEthMessage(address _to, address _from) external payable {
        depositEth(msg.sender);

        _deliverEthMessage(
            msg.sender,
            _to,
            _from,
            msg.value
        );
    }

    // // Transaction format
    // //   tx length bytes(32 bytes)
    // //   to (20 bytes)
    // //   seqNumber (32 bytes)
    // //   value (32 bytes)
    // //   signature (65 bytes)
    // //   data (arbitrary length)


    function deliverTransactionBatch(
        address chain,
        bytes calldata transactions
    )
        external
    {
        require(msg.sender == tx.origin, "origin only");
        bytes32 messageHash = keccak256(
            abi.encodePacked(
                TRANSACTION_BATCH_MSG,
                transactions
            )
        );

        _deliverMessage(chain, messageHash);

        emit TransactionMessageBatchDelivered(chain);
    }

    function _deliverTransactionMessage(
        address _chain,
        address _to,
        address _from,
        uint256 _seqNumber,
        uint256 _value,
        bytes memory _data
    )
        private
    {
        bytes32 messageHash = Messages.transactionHash(
            _chain,
            _to,
            _from,
            _seqNumber,
            _value,
            keccak256(_data)
        );

        _deliverMessage(_chain, messageHash);

        emit IGlobalInbox.TransactionMessageDelivered(
            _chain,
            _to,
            _from,
            _seqNumber,
            _value,
            _data
        );
    }

    function _deliverEthMessage(
        address _chain,
        address _to,
        address _from,
        uint256 _value
    )
        private
    {
        bytes32 messageHash = Messages.ethHash(
            _to,
            _from,
            _value
        );

        uint256 messageNum = _deliverMessage(_chain, messageHash);

        emit IGlobalInbox.EthDepositMessageDelivered(
            _chain,
            _to,
            msg.sender,
            msg.value,
            messageNum
        );
    }

    function _deliverERC20TokenMessage(
        address _chain,
        address _to,
        address _from,
        address _erc20,
        uint256 _value
    )
        private
    {
        bytes32 messageHash = Messages.erc20Hash(
            _to,
            _from,
            _erc20,
            _value
        );

        uint256 messageNum = _deliverMessage(_chain, messageHash);

        emit IGlobalInbox.ERC20DepositMessageDelivered(
            _chain,
            _to,
            _from,
            _erc20,
            _value,
            messageNum
        );
    }

    function _deliverERC721TokenMessage(
        address _chain,
        address _to,
        address _from,
        address _erc721,
        uint256 _id
    )
        private
    {
        bytes32 messageHash = Messages.erc721Hash(
            _to,
            _from,
            _erc721,
            _id
        );

        uint256 messageNum = _deliverMessage(_chain, messageHash);

        emit IGlobalInbox.ERC721DepositMessageDelivered(
            _chain,
            _to,
            _from,
            _erc721,
            _id,
            messageNum
        );
    }

    function _deliverContractTransactionMessage(
        address _chain,
        address _to,
        address _from,
        uint256 _value,
        bytes memory _data
    )
        private
    {
        bytes32 messageHash = Messages.contractTransactionHash(
            _to,
            _from,
            _value,
            _data
        );

        uint256 messageNum = _deliverMessage(_chain, messageHash);

        emit IGlobalInbox.ContractTransactionMessageDelivered(
            _chain,
            _to,
            _from,
            _value,
            _data,
            messageNum
        );
    }

    function _deliverMessage(address _chain, bytes32 _messageHash) private returns(uint256) {
        Inbox storage inbox = inboxes[_chain];
        uint256 updatedCount = inbox.count + 1;
        inbox.value = Messages.addMessageToInbox(
            inbox.value,
            _messageHash,
            block.number,
            block.timestamp,
            updatedCount
        );
        inbox.count = updatedCount;
        return updatedCount;
    }
}
