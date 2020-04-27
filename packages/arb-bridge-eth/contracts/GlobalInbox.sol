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

import "./arch/Protocol.sol";
import "./arch/Value.sol";

import "./libraries/SigUtils.sol";

contract GlobalInbox is GlobalEthWallet, GlobalFTWallet, GlobalNFTWallet, IGlobalInbox {

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

    function sendMessages(bytes calldata _messages) external {
        bool valid;
        uint256 offset = 0;
        uint256 messageType;
        address sender;
        uint256 totalLength = _messages.length;

        while (offset < totalLength) {
            (
                valid,
                offset,
                messageType,
                sender
            ) = Value.deserializeMessageData(_messages, offset);
            if (!valid) {
                break;
            }
            (valid, offset) = sendDeserializedMsg(_messages, offset, messageType);
            if (!valid) {
                break;
            }
        }
    }

    function sendDeserializedMsg(
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
            transferEth(msg.sender, to, value);
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
            transferERC20(msg.sender, to, erc20, value);
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
            transferNFT(msg.sender, to, erc721, value);
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
        uint256 messageCount = tos.length;
        require(seqNumbers.length == messageCount, "wrong input length");
        require(values.length == messageCount, "wrong input length");
        require(dataLengths.length == messageCount, "wrong input length");

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
            keccak256(_data),
            block.number,
            block.timestamp
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
            block.timestamp,
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
            block.timestamp,
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
            block.timestamp,
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
            block.timestamp,
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
