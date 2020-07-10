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

contract GlobalInbox is
    GlobalEthWallet,
    GlobalFTWallet,
    GlobalNFTWallet,
    IGlobalInbox,
    PaymentRecords // solhint-disable-next-line bracket-align
{
    uint8 internal constant ETH_DEPOSIT = 0;
    uint8 internal constant ERC20_DEPOSIT = 1;
    uint8 internal constant ERC721_DEPOSIT = 2;
    uint8 internal constant L2_MSG = 3;

    uint8 internal constant TRANSACTION_BATCH_MSG = 6;

    using Value for Value.Data;

    address internal constant ETH_ADDRESS = address(0);

    struct Inbox {
        bytes32 value;
        uint256 count;
    }

    mapping(address => Inbox) private inboxes;

    function getInbox(address account)
        external
        view
        returns (bytes32, uint256)
    {
        Inbox storage inbox = inboxes[account];
        return (inbox.value, inbox.count);
    }

    function sendMessages(
        bytes calldata _messages,
        uint256[] calldata messageCounts,
        bytes32[] calldata nodeHashes
    ) external {
        bool valid;
        uint256 offset = 0;
        uint256 messageType;
        address sender;

        uint256 nodeCount = nodeHashes.length;
        for (uint256 i = 0; i < nodeCount; i++) {
            for (uint256 j = 0; j < messageCounts[i]; j++) {
                (valid, offset, messageType, sender) = Value
                    .deserializeMessageData(_messages, offset);
                if (!valid) {
                    return;
                }
                (valid, offset) = sendDeserializedMsg(
                    nodeHashes[i],
                    j,
                    _messages,
                    offset,
                    messageType
                );
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
        returns (
            bool, // valid
            uint256 // offset
        )
    {
        if (messageType == ETH_DEPOSIT) {
            (bool valid, uint256 offset, address to, uint256 value) = Value
                .getEthMsgData(_messages, startOffset);

            if (!valid) {
                return (false, startOffset);
            }

            address paymentOwner = getPaymentOwner(to, nodeHash, messageIndex);
            transferEth(msg.sender, paymentOwner, value);
            deletePayment(to, nodeHash, messageIndex);

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
            deletePayment(to, nodeHash, messageIndex);

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
            deletePayment(to, nodeHash, messageIndex);

            return (true, offset);
        } else {
            return (false, startOffset);
        }
    }

    function sendL2MessageFromOrigin(
        address _chain,
        bytes calldata _messageData
    ) external {
        // solhint-disable-next-line avoid-tx-origin
        require(msg.sender == tx.origin, "origin only");
        uint256 inboxSeqNum = _deliverMessageImpl(
            _chain,
            L2_MSG,
            msg.sender,
            keccak256(_messageData)
        );
        emit IGlobalInbox.MessageDeliveredFromOrigin(
            _chain,
            L2_MSG,
            msg.sender,
            inboxSeqNum
        );
    }

    function sendL2Message(address _chain, bytes calldata _messageData)
        external
    {
        _deliverMessage(_chain, L2_MSG, msg.sender, _messageData);
    }

    function depositEthMessage(address _chain, address _to) external payable {
        depositEth(_chain);
        _deliverMessage(
            _chain,
            ETH_DEPOSIT,
            msg.sender,
            abi.encodePacked(_to, msg.value)
        );
    }

    function depositERC20Message(
        address _chain,
        address _erc20,
        address _to,
        uint256 _value
    ) external {
        depositERC20(_erc20, _chain, _value);
        _deliverMessage(
            _chain,
            ERC20_DEPOSIT,
            msg.sender,
            abi.encodePacked(_erc20, _to, _value)
        );
    }

    function depositERC721Message(
        address _chain,
        address _erc721,
        address _to,
        uint256 _id
    ) external {
        depositERC721(_erc721, _chain, _id);
        _deliverMessage(
            _chain,
            ERC721_DEPOSIT,
            msg.sender,
            abi.encodePacked(_erc721, _to, _id)
        );
    }

    function _deliverMessage(
        address _chain,
        uint8 _kind,
        address _sender,
        bytes memory _messageData
    ) private {
        uint256 inboxSeqNum = _deliverMessageImpl(
            _chain,
            _kind,
            _sender,
            keccak256(_messageData)
        );
        emit IGlobalInbox.MessageDelivered(
            _chain,
            _kind,
            _sender,
            inboxSeqNum,
            _messageData
        );
    }

    function _deliverMessageImpl(
        address _chain,
        uint8 _kind,
        address _sender,
        bytes32 _messageDataHash
    ) private returns (uint256) {
        Inbox storage inbox = inboxes[_chain];
        uint256 updatedCount = inbox.count + 1;
        bytes32 messageHash = Messages.messageHash(
            _kind,
            _sender,
            block.number,
            block.timestamp, // solhint-disable-line not-rely-on-time
            updatedCount,
            _messageDataHash
        );
        inbox.value = Messages.addMessageToInbox(inbox.value, messageHash);
        inbox.count = updatedCount;
        return updatedCount;
    }
}
