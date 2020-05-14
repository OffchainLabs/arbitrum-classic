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

import "bytes/contracts/BytesLib.sol";

contract GlobalInbox is GlobalEthWallet, GlobalFTWallet, GlobalNFTWallet, IGlobalInbox, PaymentRecords {

    uint8 internal constant TRANSACTION_MSG = 0;
    uint8 internal constant ETH_DEPOSIT = 1;
    uint8 internal constant ERC20_DEPOSIT = 2;
    uint8 internal constant ERC721_DEPOSIT = 3;

    using BytesLib for bytes;
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
        uint256 totalLength = _messages.length;


        uint256 currentNode = 0;
        uint256 currentIndex = 0;

        while (offset < totalLength && currentNode < nodeHashes.length) {
            if(messageCounts[currentNode] == 0){

                currentNode += 1;
                currentIndex = 0;
            } else {
                (   valid,
                    offset,
                    messageType,
                    sender
                ) = Value.deserializeMessageData(_messages, offset);
                if (!valid) {
                    break;
                }
                (valid, offset) = sendDeserializedMsg(nodeHashes[currentNode], currentIndex, _messages, offset, messageType);
                if (!valid) {
                    break;
                }

                currentIndex += 1;
                if(currentIndex >= messageCounts[currentNode]){
                    currentNode += 1;
                    currentIndex = 0;
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

    function deliverTransactionBatch(
        address _chain,
        address[] memory _tos,
        uint256[] memory _seqNumbers,
        uint256[] memory _values,
        uint256[] memory _messageLengths,
        bytes memory _data,
        bytes memory _signatures
    )
        public
    {
        uint256 messageCount = _tos.length;
        uint256 dataOffset = 0;
        require(_seqNumbers.length == messageCount, "wrong input length");
        require(_values.length == messageCount, "wrong input length");
        require(_messageLengths.length == messageCount, "wrong input length");

        Inbox storage inbox = inboxes[_chain];
        bytes32 inboxVal = inbox.value;

        for (uint256 i = 0; i < messageCount; i++) {
            uint256 messageLength = _messageLengths[i];
            bytes memory messageData = _data.slice(dataOffset, messageLength);
            dataOffset += messageLength;

            bytes32 messageHash = deliverTransactionSingle(
                _chain,
                _tos[i],
                _seqNumbers[i],
                _values[i],
                messageData,
                _signatures,
                i * 65
            );

            inboxVal = Protocol.addMessageToInbox(inboxVal, messageHash);
        }

        inbox.value = inboxVal;
        inbox.count += messageCount;
    }

    function deliverTransactionSingle(
        address _chain,
        address _to,
        uint256 _seqNumber,
        uint256 _value,
        bytes memory _messageData,
        bytes memory _signatures,
        uint256 signatureOffset
    )
        private
        returns(bytes32)
    {
        address from = SigUtils.recoverAddress(
            keccak256(
                abi.encodePacked(
                    _chain,
                    _to,
                    _seqNumber,
                    _value,
                    _messageData
                )
            ),
            _signatures,
            signatureOffset
        );

        bytes32 messageHash = Messages.transactionHash(
            _chain,
            _to,
            from,
            _seqNumber,
            _value,
            _messageData,
            block.number
        );


        emit IGlobalInbox.TransactionMessageDelivered(
            _chain,
            _to,
            from,
            _seqNumber,
            _value,
            _messageData
        );

        return messageHash;
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
            _data,
            block.number
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
        uint256 messageNum = inboxes[_chain].count + 1;
        bytes32 messageHash = Messages.ethHash(
            _to,
            _from,
            _value,
            block.number,
            messageNum
        );

        _deliverMessage(_chain, messageHash);

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
        uint256 messageNum = inboxes[_chain].count + 1;
        bytes32 messageHash = Messages.erc20Hash(
            _to,
            _from,
            _erc20,
            _value,
            block.number,
            messageNum
        );

        _deliverMessage(_chain, messageHash);

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
        uint256 messageNum = inboxes[_chain].count + 1;
        bytes32 messageHash = Messages.erc721Hash(
            _to,
            _from,
            _erc721,
            _id,
            block.number,
            messageNum
        );

        _deliverMessage(_chain, messageHash);

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
        uint256 messageNum = inboxes[_chain].count + 1;
        bytes32 messageHash = Messages.contractTransactionHash(
            _to,
            _from,
            _value,
            _data,
            block.number,
            messageNum
        );

        _deliverMessage(_chain, messageHash);

        emit IGlobalInbox.ContractTransactionMessageDelivered(
            _chain,
            _to,
            _from,
            _value,
            _data,
            messageNum
        );
    }

    function _deliverMessage(address _chain, bytes32 _messageHash) private {
        Inbox storage inbox = inboxes[_chain];
        inbox.value = Protocol.addMessageToInbox(inbox.value, _messageHash);
        inbox.count++;
    }
}
