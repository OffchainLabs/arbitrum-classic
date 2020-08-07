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
    uint8 internal constant INITIALIZATION_MSG = 4;

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

    /**
     * @notice Process a set of marshalled messages confirmed by a rollup chain
     * @dev messageCounts and nodeHashes are used to uniquely identify messages in conjunction with PaymentRecords
     * @param messages Contiguously marshaled messages from a set of assertions
     * @param messageCounts Number of messages in each assertion confirmed
     * @param nodeHashes Hash of each node that has been confirmed
     */
    function sendMessages(
        bytes calldata messages,
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
                    messages,
                    offset
                );
                if (!valid) {
                    return;
                }
                sendDeserializedMsg(nodeHashes[i], j, message);
            }
        }
    }

    /**
     * @notice Send a generic L2 message to a given Arbitrum Rollup chain
     * @dev This method is an optimization to avoid having to emit the entirety of the messageData in a log. Instead validators are expected to be able to parse the data from the transaction's input
     * @param chain Address of the rollup chain that the ETH is deposited into
     * @param messageData Data of the message being sent
     */
    function sendL2MessageFromOrigin(address chain, bytes calldata messageData)
        external
    {
        // solhint-disable-next-line avoid-tx-origin
        require(msg.sender == tx.origin, "origin only");
        uint256 inboxSeqNum = _deliverMessageImpl(
            chain,
            L2_MSG,
            msg.sender,
            keccak256(messageData)
        );
        emit IGlobalInbox.MessageDeliveredFromOrigin(
            chain,
            L2_MSG,
            msg.sender,
            inboxSeqNum
        );
    }

    /**
     * @notice Send a generic L2 message to a given Arbitrum Rollup chain
     * @dev This method can be used to send any type of message that doesn't require L1 validation
     * @param chain Address of the rollup chain that the ETH is deposited into
     * @param messageData Data of the message being sent
     */
    function sendL2Message(address chain, bytes calldata messageData) external {
        _deliverMessage(chain, L2_MSG, msg.sender, messageData);
    }

    /**
     * @notice Send a generic L2 message to a given Arbitrum Rollup chain
     * @dev This method can be used to send any type of message that doesn't require L1 validation
     * @param messageData Data of the message being sent
     */
    function sendInitializationMessage(bytes calldata messageData) external {
        _deliverMessage(
            msg.sender,
            INITIALIZATION_MSG,
            msg.sender,
            messageData
        );
    }

    /**
     * @notice Deposits ETH into a given Arbitrum Rollup chain
     * @dev This method is payable and will deposit all value it is called with
     * @param chain Address of the rollup chain that the ETH is deposited into
     * @param to Address on the rollup chain that will receive the ETH
     */
    function depositEthMessage(address chain, address to) external payable {
        depositEth(chain);
        _deliverMessage(
            chain,
            ETH_TRANSFER,
            msg.sender,
            abi.encodePacked(uint256(uint160(bytes20(to))), msg.value)
        );
    }

    /**
     * @notice Deposits an ERC20 token into a given Arbitrum Rollup chain
     * @dev This method requires approving this contract for transfers
     * @param chain Address of the rollup chain that the token is deposited into
     * @param erc20 L1 address of the token being deposited
     * @param to Address on the rollup chain that will receive the tokens
     * @param value Quantity of tokens being deposited
     */
    function depositERC20Message(
        address chain,
        address erc20,
        address to,
        uint256 value
    ) external {
        depositERC20(erc20, chain, value);
        _deliverMessage(
            chain,
            ERC20_TRANSFER,
            msg.sender,
            abi.encodePacked(
                uint256(uint160(bytes20(erc20))),
                uint256(uint160(bytes20(to))),
                value
            )
        );
    }

    /**
     * @notice Deposits an ERC721 token into a given Arbitrum Rollup chain
     * @dev This method requires approving this contract for transfers
     * @param chain Address of the rollup chain that the token is deposited into
     * @param erc721 L1 address of the token being deposited
     * @param to Address on the rollup chain that will receive the token
     * @param id ID of the token being deposited
     */
    function depositERC721Message(
        address chain,
        address erc721,
        address to,
        uint256 id
    ) external {
        depositERC721(erc721, chain, id);
        _deliverMessage(
            chain,
            ERC721_TRANSFER,
            msg.sender,
            abi.encodePacked(
                uint256(uint160(bytes20(erc721))),
                uint256(uint160(bytes20(to))),
                id
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
}
