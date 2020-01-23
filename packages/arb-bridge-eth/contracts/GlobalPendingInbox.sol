/*
 * Copyright 2019, Offchain Labs, Inc.
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

import "./GlobalWallet.sol";
import "./IGlobalPendingInbox.sol";

import "./arch/Protocol.sol";
import "./arch/Value.sol";

import "./libraries/SigUtils.sol";

import "@openzeppelin/contracts/ownership/Ownable.sol";

contract GlobalPendingInbox is GlobalWallet, IGlobalPendingInbox {

    uint8 internal constant TRANSACTION_MSG = 0;
    uint8 internal constant ETH_DEPOSIT = 1;
    uint8 internal constant ERC20_DEPOSIT = 2;
    uint8 internal constant ERC721_DEPOSIT = 3;

    using SafeMath for uint256;
    using Value for Value.Data;

    address internal constant ETH_ADDRESS = address(0);

    struct PendingInbox {
        bytes32 value;
        uint256 count;
    }

    mapping(address => PendingInbox) pending;

    function getPending() external returns(bytes32, uint) {
        PendingInbox storage pendingInbox = pending[msg.sender];
        return (pendingInbox.value, pendingInbox.count);
    }

    function registerForInbox() external {
        require(pending[msg.sender].value == 0, "Pending must be uninitialized");
        pending[msg.sender].value = Value.hashEmptyTuple();
    }

    address inboxAddress = 0xCAAd408788C192979384768DD5bE04eC1b3787dA;

    function sendMessages(bytes calldata _messages) external {
        bool valid;
        uint offset = 0;
        bytes32 messageHash;
        uint256 messageType;
        uint256 sender;
        bytes memory messageData;
        uint256 totalLength = _messages.length;

        emit AssertionEvent(inboxAddress, false, 0 , msg.sender, "event1");
        emit AssertionEvent(msg.sender, false, 0 , msg.sender, "event1");

        while (offset < totalLength) {
            (
                valid,
                offset,
                messageHash,
                messageType,
                sender,
                messageData) = Value.deserializeMessage(_messages, offset);

            if (valid) {
                sendDeserializedMsg(messageData, messageType);
            }
        }
    }

    function sendDeserializedMsg(bytes memory messageData, uint256 messageType) private {
        if (messageType == ETH_DEPOSIT) {
            (
                bool valid,
                uint256 destination,
                uint256 value
            ) = Value.getEthMsgData(messageData);

            emit AssertionEvent(inboxAddress, valid, messageType, address(bytes20(bytes32(destination))), "event2");
            emit AssertionEvent(msg.sender, valid, messageType, address(bytes20(bytes32(destination))), "event2");

            if (valid) {
                require(
                    transferEth(
                        msg.sender,
                        address(bytes20(bytes32(destination))),
                        value
                    )
                );
            }

        } else if (messageType == ERC20_DEPOSIT) {
            (
                bool valid,
                uint256 tokenContract,
                uint256 destination,
                uint256 value
            ) = Value.getERCTokenMsgData(messageData);

            emit AssertionEvent(inboxAddress, valid, messageType, address(bytes20(bytes32(destination))), "event3");
            emit AssertionEvent(msg.sender, valid, messageType, address(bytes20(bytes32(destination))), "event3");

            if (valid) {
                require(
                    transferERC20(
                        msg.sender,
                        address(bytes20(bytes32(destination))),
                        address(bytes20(bytes32(tokenContract))),
                        value
                    )
                );
            }

        } else if (messageType == ERC721_DEPOSIT) {
            (
                bool valid,
                uint256 tokenContract,
                uint256 destination,
                uint256 value
            ) = Value.getERCTokenMsgData(messageData);

            emit AssertionEvent(inboxAddress, valid, messageType, address(bytes20(bytes32(destination))), "event4");

            if (valid) {
                require(
                    transferNFT(
                        msg.sender,
                        address(bytes20(bytes32(destination))),
                        address(bytes20(bytes32(tokenContract))),
                        value
                    )
                );
            }
        }
    }

    function forwardTransactionMessage(
        address _chain,
        address _to,
        uint256 _seqNumber,
        uint256 _value,
        bytes calldata _data,
        bytes calldata _signature
    )
        external
    {
        address sender = SigUtils.recoverAddress(
            keccak256(
                abi.encodePacked(
                    _chain,
                    _to,
                    _seqNumber,
                    _value,
                    Value.deserializeHashed(_data)
                )
            ),
            _signature
        );

        _deliverTransactionMessage(
            _chain,
            _to,
            sender,
            _seqNumber,
            _value,
            _data
        );
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
            msg.sender,
            _to,
            msg.value
        );
    }

    function depositERC20Message(
        address _chain,
        address _to,
        address _erc20,
        uint256 _value
    )
        public
    {
        depositERC20(_erc20, _chain, _value);

        _deliverERC20TokenMessage(
            _chain,
            msg.sender,
            _to,
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
            msg.sender,
            _to,
            _erc721,
            _id
        );
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
        PendingInbox storage pendingInbox = pending[_chain];

        if (pendingInbox.value != 0) {
            bytes32 dataHash = Value.deserializeHashed(_data);
            bytes32 txHash = keccak256(
                abi.encodePacked(
                    TRANSACTION_MSG,
                    _chain,
                    _to,
                    _from,
                    _seqNumber,
                    _value,
                    dataHash
                )
            );

            Value.Data[] memory msgValues = new Value.Data[](4);
            msgValues[0] = Value.newInt(uint256(_to));
            msgValues[1] = Value.newInt(_seqNumber);
            msgValues[2] = Value.newInt(_value);
            msgValues[3] = Value.newHashOnly(dataHash);

            Value.Data[] memory msgType = new Value.Data[](3);
            msgType[0] = Value.newInt(TRANSACTION_MSG);
            msgType[1] = Value.newInt(uint256(_from));
            msgType[2] = Value.newTuple(msgValues);

            Value.Data[] memory dataMsg = new Value.Data[](3);
            dataMsg[0] = Value.newInt(block.number);
            dataMsg[1] = Value.newInt(uint256(txHash));
            dataMsg[2] = Value.newTuple(msgType);

            bytes32 messageHash =  Value.newTuple(dataMsg).hash().hash;

            _deliverMessage(_chain, messageHash);

            emit IGlobalPendingInbox.TransactionMessageDelivered(
                _chain,
                _to,
                _from,
                _seqNumber,
                _value,
                _data
            );
        }
    }

    function _deliverEthMessage(
        address _chain,
        address _from,
        address _to,
        uint256 _value
    )
        private
    {
        if (pending[_chain].value != 0)
        {
            bytes32 txHash = keccak256(
                abi.encodePacked(
                    ETH_DEPOSIT,
                    _chain,
                    _to,
                    _from,
                    _value
                )
            );

            Value.Data[] memory msgValues = new Value.Data[](3);
            msgValues[0] = Value.newInt(uint256(_to));
            msgValues[1] = Value.newInt(_value);

            Value.Data[] memory msgType = new Value.Data[](3);
            msgType[0] = Value.newInt(ETH_DEPOSIT);
            msgType[1] = Value.newInt(uint256(_from));
            msgType[2] = Value.newTuple(msgValues);

            Value.Data[] memory ethMsg = new Value.Data[](3);
            ethMsg[0] = Value.newInt(block.number);
            ethMsg[1] = Value.newInt(uint256(txHash));
            ethMsg[2] = Value.newTuple(msgType);

            bytes32 messageHash =  Value.newTuple(ethMsg).hash().hash;

            _deliverMessage(_chain, messageHash);

            emit IGlobalPendingInbox.EthDepositMessageDelivered(
                _chain,
                _to,
                msg.sender,
                msg.value
            );
        }
    }

    function _deliverERC20TokenMessage(
        address _chain,
        address _from,
        address _to,
        address _erc20,
        uint256 _value
    )
        private
    {
        if (pending[_chain].value != 0)
        {
            bytes32 messageHash = _tokenMessageHash(
                ERC20_DEPOSIT,
                _chain,
                _from,
                _to,
                _erc20,
                _value
            );

            _deliverMessage(_chain, messageHash);

            emit IGlobalPendingInbox.ERC20DepositMessageDelivered(
                _chain,
                _to,
                _from,
                _erc20,
                _value
            );
        }
    }

    function _deliverERC721TokenMessage(
        address _chain,
        address _from,
        address _to,
        address _erc721,
        uint256 _id
    )
        private
    {
        if (pending[_chain].value != 0)
        {
            bytes32 messageHash = _tokenMessageHash(
                ERC721_DEPOSIT,
                _chain,
                _from,
                _to,
                _erc721,
                _id
            );

            _deliverMessage(_chain, messageHash);

            emit IGlobalPendingInbox.ERC721DepositMessageDelivered(
                _chain,
                _to,
                _from,
                _erc721,
                _id
            );
        }
    }

    function _tokenMessageHash(
        uint8 _messageType,
        address _chain,
        address _from,
        address _to,
        address _tokenContract,
        uint256 _value
    )
        private
        view
        returns(bytes32)
    {
        bytes32 txHash = keccak256(
            abi.encodePacked(
                _messageType,
                _chain,
                _from,
                _to,
                _tokenContract,
                _value
            )
        );

        Value.Data[] memory msgValues = new Value.Data[](3);
        msgValues[0] = Value.newInt(uint256(_tokenContract));
        msgValues[1] = Value.newInt(uint256(_to));
        msgValues[2] = Value.newInt(_value);

        Value.Data[] memory msgType = new Value.Data[](3);
        msgType[0] = Value.newInt(_messageType);
        msgType[1] = Value.newInt(uint256(_from));
        msgType[2] = Value.newTuple(msgValues);

        Value.Data[] memory ercTokenMsg = new Value.Data[](3);
        ercTokenMsg[0] = Value.newInt(block.number);
        ercTokenMsg[1] = Value.newInt(uint(txHash));
        ercTokenMsg[2] = Value.newTuple(msgType);

        return  Value.newTuple(ercTokenMsg).hash().hash;
    }

    function _deliverMessage(address _chain, bytes32 _messageHash) private {
        PendingInbox storage pendingInbox = pending[_chain];

        pendingInbox.value = Protocol.addMessageToPending(
            pendingInbox.value,
            _messageHash
        );

        pendingInbox.count++;
    }
}
