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

contract GlobalInbox is
    GlobalEthWallet,
    GlobalFTWallet,
    GlobalNFTWallet,
    IGlobalInbox,
    PaymentRecords // solhint-disable-next-line bracket-align
{
    uint8 internal constant ETH_TRANSFER = 0;
    uint8 internal constant ERC20_TRANSFER = 1;
    uint8 internal constant ERC721_TRANSFER = 2;
    uint8 internal constant L2_MSG = 3;

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
        Messages.OutgoingMessage memory message;

        uint256 nodeCount = nodeHashes.length;
        for (uint256 i = 0; i < nodeCount; i++) {
            for (uint256 j = 0; j < messageCounts[i]; j++) {
                (valid, offset, message) = Messages.unmarshalOutgoingMessage(
                    _messages,
                    offset
                );
                if (!valid) {
                    return;
                }
                sendDeserializedMsg(nodeHashes[i], j, message);
            }
        }
    }

    function sendDeserializedMsg(
        bytes32 nodeHash,
        uint256 messageIndex,
        Messages.OutgoingMessage memory message
    ) private {
        if (message.kind == ETH_TRANSFER) {
            (bool valid, Messages.EthMessage memory eth) = Messages
                .parseEthMessage(message.data);
            if (valid) {
                address paymentOwner = getPaymentOwner(
                    eth.dest,
                    nodeHash,
                    messageIndex
                );
                transferEth(msg.sender, paymentOwner, eth.value);
                deletePayment(eth.dest, nodeHash, messageIndex);
            }
        } else if (message.kind == ERC20_TRANSFER) {
            (bool valid, Messages.ERC20Message memory erc20) = Messages
                .parseERC20Message(message.data);
            if (valid) {
                address paymentOwner = getPaymentOwner(
                    erc20.dest,
                    nodeHash,
                    messageIndex
                );
                transferERC20(
                    msg.sender,
                    paymentOwner,
                    erc20.token,
                    erc20.value
                );
                deletePayment(erc20.dest, nodeHash, messageIndex);
            }
        } else if (message.kind == ERC721_TRANSFER) {
            (bool valid, Messages.ERC721Message memory erc721) = Messages
                .parseERC721Message(message.data);
            if (valid) {
                address paymentOwner = getPaymentOwner(
                    erc721.dest,
                    nodeHash,
                    messageIndex
                );
                transferNFT(msg.sender, paymentOwner, erc721.token, erc721.id);
                deletePayment(erc721.dest, nodeHash, messageIndex);
            }
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
            ETH_TRANSFER,
            msg.sender,
            abi.encodePacked(bytes32(bytes20(_to)), msg.value)
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
            ERC20_TRANSFER,
            msg.sender,
            abi.encodePacked(
                bytes32(bytes20(_erc20)),
                bytes32(bytes20(_to)),
                _value
            )
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
            ERC721_TRANSFER,
            msg.sender,
            abi.encodePacked(
                bytes32(bytes20(_erc721)),
                bytes32(bytes20(_to)),
                _id
            )
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
