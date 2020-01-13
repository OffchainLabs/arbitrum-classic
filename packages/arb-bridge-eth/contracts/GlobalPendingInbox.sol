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

    mapping(address => bytes32) pending;

    function pullPendingMessages() external returns(bytes32) {
        bytes32 messages = pending[msg.sender];
        pending[msg.sender] = Value.hashEmptyTuple();
        return messages;
    }

    function registerForInbox() external {
        require(pending[msg.sender] == 0, "Pending must be uninitialized");
        pending[msg.sender] = Value.hashEmptyTuple();
    }

    function sendMessages(bytes calldata _messages) external {
        bool valid;
        uint offset = 0;
        bytes32 messageHash;
        uint256 messageType;
        uint256 sender;
        bytes memory messageData;
        uint totalLength = _messages.length;

        while (offset < totalLength) {
            (
                valid, 
                offset, 
                messageHash, 
                messageType, 
                sender, 
                messageData) = Value.deserializeMessage(_messages, offset);

            if(valid){
                sendDeserializedMsgs(messageData, messageType);
            }
        }
    }

    function sendDeserializedMsgs(bytes memory messageData, uint256 messageType) private {
        if(messageType == TRANSACTION_MSG)
        {
            bool valid;
            uint256 vmAddress;
            uint256 contractAddress;
            uint256 seqNumber;
            uint256 value;
            (   valid,
                vmAddress,
                contractAddress,
                seqNumber,
                value,
                messageData
            ) = Value.getTransactionMsgData(messageData);
        
            if(valid){
                _deliverTransactionMessage(
                    msg.sender,
                    address(bytes20(bytes32(vmAddress))),
                    address(bytes20(bytes32(contractAddress))),
                    seqNumber,
                    value,
                    messageData
                );
            }
        }else if(messageType == ETH_DEPOSIT)
        {
            bool valid;
            uint256 destination;
            uint256 value;

            (   valid,
                destination,
                value
            ) = Value.getEthMsgData(messageData);

            if(valid){
                depositEthMessage(
                    address(bytes20(bytes32(destination))), 
                    value);
            }   

        }else if(messageType == ERC20_DEPOSIT)
        {
            bool valid;
            uint256 destination;
            uint256 tokenContract;
            uint256 value;

            (   valid,
                tokenContract,
                destination,
                value
            ) = Value.getERCTokenMsgData(messageData);

            if(valid){
                depositERC20Message(
                    address(bytes20(bytes32(destination))), //wrong
                    address(bytes20(bytes32(tokenContract))), 
                    address(bytes20(bytes32(destination))), 
                    value);
            }   

        }else if(messageType == ERC721_DEPOSIT)
        {
            bool valid;
            uint256 destination;
            uint256 tokenContract;
            uint256 value;

            (   valid,
                tokenContract,
                destination,
                value
            ) = Value.getERCTokenMsgData(messageData);

            if(valid){
                depositERC721Message(
                    address(bytes20(bytes32(destination))), //wrong
                    address(bytes20(bytes32(tokenContract))), 
                    address(bytes20(bytes32(destination))), 
                    value);
            }   

        }

    }

    function forwardTransactionMessage(
        address _vmAddress,
        address _contractAddress, 
        uint256 _seqNumber,
        uint256 _value,
        bytes calldata _data,
        bytes calldata _signature) external
    {
        address sender = SigUtils.recoverAddress(
            keccak256(
                abi.encodePacked(
                    _vmAddress,
                    _contractAddress,
                    _seqNumber,
                    _value,
                    Value.deserializeHashed(_data)
                )
            ),
            _signature
        );

        _deliverTransactionMessage(
            sender,
            _vmAddress,
            _contractAddress,
            _seqNumber,
            _value,
            _data
        );
    }

    function sendTransactionMessage(
        address _vmAddress,
        address _contractAddress, 
        uint256 _seqNumber,
        uint256 _value,
        bytes calldata _data) external
    {
        _deliverTransactionMessage(
            msg.sender,
            _vmAddress,
            _contractAddress,
            _seqNumber,
            _value,
            _data
        );
    }

    function depositEthMessage(address _destination) external payable 
    {
        depositEth(_destination);
        
        _deliverEthMessage(
            msg.sender,
            _destination,
            ETH_DEPOSIT,
            msg.value
        );

        emit IGlobalPendingInbox.EthDepositMessageDelivered(
            _destination,
            msg.sender,
            msg.value);
    }

    function depositEthMessage(address payable _destination, uint256 _value) public
    {
        transferEth(_destination, _value);
        
        _deliverEthMessage(
            msg.sender,
            _destination,
            ETH_DEPOSIT,
            _value
        );

        emit IGlobalPendingInbox.EthDepositMessageDelivered(
            _destination,
            msg.sender,
            _value);
    }

    function depositERC20Message(
        address _vmAddress,
        address _tokenContract,
        address _destination,
        uint256 _value
    )
        public
    {
        depositERC20(_tokenContract, _vmAddress, _value);

        _deliverERCTokenMessage(
            _vmAddress,
            msg.sender,
            _destination,
            ERC20_DEPOSIT,
            _tokenContract,
            _value);
    }

    function depositERC721Message(
        address _vmAddress,
        address _tokenContract,
        address _destination,
        uint256 _value) public
    {
        depositERC721(_tokenContract, _vmAddress, _value);

        _deliverERCTokenMessage(
            _vmAddress,
            msg.sender,
            _destination,
            ERC721_DEPOSIT,
            _tokenContract,
            _value);
    }

    function _deliverTransactionMessage(
        address _sender,
        address _vmAddress,
        address _contractAddress,
        uint256 _seqNumber,
        uint256 _value,
        bytes memory _data) private
    {
        if (pending[_vmAddress] != 0) 
        {
            bytes32 dataHash = Value.deserializeHashed(_data);
            bytes32 txHash = keccak256(
                abi.encodePacked(
                    _sender,
                    _contractAddress,
                    _seqNumber,
                    dataHash,
                    pending[_vmAddress]
                )
            );

            Value.Data[] memory msgValues = new Value.Data[](4);
            msgValues[0] = Value.newInt(uint256(_contractAddress));
            msgValues[1] = Value.newInt(_seqNumber);
            msgValues[2] = Value.newInt(_value);
            msgValues[3] = Value.newHashOnly(dataHash);

            Value.Data[] memory msgType = new Value.Data[](3);
            msgType[0] = Value.newInt(TRANSACTION_MSG);
            msgType[1] = Value.newInt(uint256(_sender));
            msgType[2] = Value.newTuple(msgValues);

            Value.Data[] memory dataMsg = new Value.Data[](4);
            dataMsg[0] = Value.newInt(block.timestamp);
            dataMsg[1] = Value.newInt(block.number);
            dataMsg[2] = Value.newInt(uint256(txHash));
            dataMsg[3] = Value.newTuple(msgType);

            bytes32 messageHash =  Value.newTuple(dataMsg).hash().hash;

            _deliverMessage(_vmAddress, messageHash);
        }

        emit IGlobalPendingInbox.TransactionMessageDelivered(
            _sender,
            _vmAddress,
            _contractAddress,
            _seqNumber,
            _value,
            _data);
    }

    function _deliverEthMessage(
        address _sender,
        address _destination,
        uint256 _messageType,
        uint256 _value) private
    {
        if (pending[_destination] != 0) 
        {
            bytes32 txHash = keccak256(
                abi.encodePacked(
                    _sender,
                    _destination,
                    _value,
                    pending[_destination]
                )
            );

            Value.Data[] memory msgValues = new Value.Data[](3);
            msgValues[0] = Value.newInt(uint256(_destination));
            msgValues[1] = Value.newInt(_value);

            Value.Data[] memory msgType = new Value.Data[](3);
            msgType[0] = Value.newInt(_messageType);
            msgType[1] = Value.newInt(uint256(_sender));
            msgType[2] = Value.newTuple(msgValues);

            Value.Data[] memory ethMsg = new Value.Data[](4);
            ethMsg[0] = Value.newInt(block.timestamp);
            ethMsg[1] = Value.newInt(block.number);
            ethMsg[2] = Value.newInt(uint256(txHash));
            ethMsg[3] = Value.newTuple(msgType);

            bytes32 messageHash =  Value.newTuple(ethMsg).hash().hash;

            _deliverMessage(_destination, messageHash);
        }
    }

    function _deliverERCTokenMessage(
        address _vmAddress,
        address _sender,
        address _destination,
        uint256 _messageType,
        address _tokenContract,
        uint256 _value) private
    {
        if (pending[_vmAddress] != 0) 
        {
            bytes32 txHash = keccak256(
                abi.encodePacked(
                    _sender,
                    _destination,
                    _tokenContract,
                    _value
                )
            );

            Value.Data[] memory msgValues = new Value.Data[](3);
            msgValues[0] = Value.newInt(uint256(_tokenContract));
            msgValues[1] = Value.newInt(uint256(_destination));
            msgValues[2] = Value.newInt(_value);

            Value.Data[] memory msgType = new Value.Data[](3);
            msgType[0] = Value.newInt(_messageType);
            msgType[1] = Value.newInt(uint256(_sender));
            msgType[2] = Value.newTuple(msgValues);

            Value.Data[] memory ercTokenMsg = new Value.Data[](4);
            ercTokenMsg[0] = Value.newInt(block.timestamp);
            ercTokenMsg[1] = Value.newInt(block.number);
            ercTokenMsg[2] = Value.newInt(uint(txHash));
            ercTokenMsg[3] = Value.newTuple(msgType);

            bytes32 messageHash =  Value.newTuple(ercTokenMsg).hash().hash;

            _deliverMessage(_vmAddress, messageHash);

            if(_messageType == ERC20_DEPOSIT){

            emit IGlobalPendingInbox.DepositERC20MessageDelivered(
                _vmAddress,
                msg.sender,
                _tokenContract,
                _value);

            }
        }

        if(_messageType == ERC20_DEPOSIT){

            emit IGlobalPendingInbox.DepositERC20MessageDelivered(
                _destination,
                msg.sender,
                _tokenContract,
                _value);

        }else if(_messageType == ERC721_DEPOSIT){

            emit IGlobalPendingInbox.DepositERC721MessageDelivered(
                _destination,
                msg.sender,
                _tokenContract,
                _value);
        }
    }

    function _deliverMessage(
        address _vmAddress,
        bytes32 _messageHash) private 
    {
        pending[_vmAddress] = Value.hashTuple([
            Value.newInt(0),
            Value.newHashOnly(pending[_vmAddress]),
            Value.newHashOnly(_messageHash)
        ]);
    }
}
