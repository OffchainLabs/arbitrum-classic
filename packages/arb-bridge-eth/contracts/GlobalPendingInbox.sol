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
    uint8 internal constant ETH_WITHDRAWAL = 4;
    uint8 internal constant ERC20_WITHDRAWAL = 5;
    uint8 internal constant ERC721_WITHDRAWAL = 6;

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
        uint offset = 0;
        bool valid;
        uint256 messageType;
        bytes32 messageHash;
        uint256 sender;
        uint256 tokenContract;
        bytes memory messageData;
        // must convert Data to bytes for transaction messageData
        Data[] memory data;
        uint totalLength = _messages.length;

        while (offset < totalLength) {
            (
                valid, 
                offset, 
                messageHash, 
                messageType, 
                sender, 
                data) = Value.deserializeMessage(_messages, offset);

            if(valid){
                if(messageType == TRANSACTION_MSG)
                {
                    (   valid,
                        destination,
                        seq_number,
                        value,
                        messageData
                    ) = Value.getTransactionMsgData(data);
                
                    if(valid){
                        sendTransactionMessage(
                            address(bytes20(bytes32(destination))),
                            seq_number,
                            value,
                            messageData
                        );
                    }
                }else if(messageType == ETH_DEPOSIT)
                {
                    (   valid,
                        destination,
                        value
                    ) = Value.getEthMsgData(data);

                    if(valid){
                        depositEthMessage(
                            address(bytes20(bytes32(destination))), 
                            value);
                    }   

                }else if(messageType == ETH_WITHDRAWAL)
                {
                    (   valid,
                        destination,
                        value
                    ) = Value.getEthMsgData(data);

                    if(valid){
                        withdrawEthMessage(
                            address(bytes20(bytes32(destination))), 
                            value);
                    }   

                }else if(messageType == ERC20_DEPOSIT)
                {
                    (   valid,
                        tokenContract,
                        destination,
                        value
                    ) = Value.getERCTokenMsgData(data);

                    if(valid){
                        depositERC20Message(
                            address(bytes20(bytes32(tokenContract))), 
                            address(bytes20(bytes32(destination))), 
                            value);
                    }   

                }else if(messageType == ERC20_WITHDRAWAL)
                {
                    (   valid,
                        tokenContract,
                        destination,
                        value
                    ) = Value.getERCTokenMsgData(data);

                    if(valid){
                        withdrawERC20Message(
                            address(bytes20(bytes32(tokenContract))), 
                            address(bytes20(bytes32(destination))), 
                            value);
                    }   

                }else if(messageType == ERC721_DEPOSIT)
                {
                    (   valid,
                        tokenContract,
                        destination,
                        value
                    ) = Value.getERCTokenMsgData(data);

                    if(valid){
                        depositERC721Message(
                            address(bytes20(bytes32(tokenContract))), 
                            address(bytes20(bytes32(destination))), 
                            value);
                    }   

                }else if(messageType == ERC721_WITHDRAWAL)
                {
                    (   valid,
                        tokenContract,
                        destination,
                        value
                    ) = Value.getERCTokenMsgData(data);

                    if(valid){
                        withdrawERC721Message(
                            address(bytes20(bytes32(tokenContract))), 
                            address(bytes20(bytes32(destination))), 
                            value);
                    }
                }

            }
        }
    }

    function forwardTransactionMessage(
        address _destination,
        uint256 _seq_number,
        uint256 _value,
        bytes calldata _data,
        bytes calldata _signature) external
    {
        address sender = SigUtils.recoverAddress(
            keccak256(
                abi.encodePacked(
                    _destination,
                    _seq_number,
                    _value,
                    ArbValue.deserializeValueHash(_data)
                )
            ),
            _signature
        );

        _deliverTransactionMessage(
            sender,
            _destination,
            _seq_number,
            _value,
            _data
        );
    }

    function forwardDepositERC20Message(
        address _tokenContract,
        address _destination,
        uint256 _value,
        bytes calldata _signature) external
    {
        address sender = SigUtils.recoverAddress(
            keccak256(
                abi.encodePacked(
                    _tokenContract,
                    _destination,
                    _value
                )
            ),
            _signature
        );

        depositERC20(_tokenContract, sender, _destination, _value);

        _deliverERCTokenMessage(
            sender,
            _destination,
            ERC20_DEPOSIT,
            _tokenContract,
            _value);
    }

    function forwardDepositERC721Message(
        address _tokenContract,
        address _destination,
        uint256 _value,
        bytes calldata _signature) external
    {
        address sender = SigUtils.recoverAddress(
            keccak256(
                abi.encodePacked(
                    _tokenContract,
                    _destination,
                    _value
                )
            ),
            _signature
        );

        depositERC721(_tokenContract, sender, _destination, _value);

        _deliverERCTokenMessage(
            sender,
            _destination,
            ERC721_DEPOSIT,
            _tokenContract,
            _value);
    }

    function sendTransactionMessage(
        address _destination, 
        uint256 _seq_number,
        uint256 _value,
        bytes calldata _data) external
    {
        _deliverTransactionMessage(
            msg.sender,
            _destination,
            _seq_number,
            _value,
            _data
        );
    }

    function despositEthMessage(address _destination) external payable 
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

    function despositEthMessage(address _destination, uint256 _value) external 
    {
        transferEth(_destination, _value);
        
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

    function withdrawEthMessage(address _destination, uint256 _value) external 
    {
        transferEth(_destination, _value);

        _deliverEthMessage(
            msg.sender,
            _destination,
            ETH_WITHDRAWAL,
            _value
        );

        emit IGlobalPendingInbox.EthWithdrawMessageDelivered(
            msg.sender,
            _destination,
            msg.value);
    }

    function depositERC20Message(
        address _tokenContract,
        address _destination,
        uint256 _value) external
    {
        depositERC20(_tokenContract, _destination, _value);

        _deliverERCTokenMessage(
            msg.sender,
            _destination,
            ERC20_DEPOSIT,
            _tokenContract,
            _value);
    }

    function withdrawERC20Message(
        address _tokenContract,
        address _destination,
        uint256 _value) external
    {
        withdrawERC20(_tokenContract, _destination, _value);

        _deliverERCTokenMessage(
            msg.sender,
            _destination,
            ERC20_WITHDRAWAL,
            _tokenContract,
            _value);
    }

    function depositERC721Message(
        address _tokenContract,
        address _destination,
        uint256 _value) external
    {
        depositERC721(_tokenContract, _destination, _value);

        _deliverERCTokenMessage(
            msg.sender,
            _destination,
            ERC721_DEPOSIT,
            _tokenContract,
            _value);
    }

    function withdrawERC721Message(
        address _tokenContract,
        address _destination,
        uint256 _value) external
    {
        withdrawERC721(_tokenContract, _destination, _value);

        _deliverERCTokenMessage(
            msg.sender,
            _destination,
            ERC721_WITHDRAWAL,
            _tokenContract,
            _value);
    }

    function _deliverTransactionMessage(
        address _sender,
        address _destination,
        uint256 _seq_number,
        uint256 _value,
        bytes memory _data) private
    {
        if (pending[_destination] != 0) 
        {
            bytes32 dataHash = ArbValue.deserializeValueHash(_data);
            bytes32 txHash = keccak256(
                abi.encodePacked(
                    _sender,
                    _destination,
                    _seq_number,
                    dataHash,
                    pending[_destination]
                )
            );

            ArbValue.Value[] memory msgValues = new ArbValue.Value[](4);
            msgValues[0] = ArbValue.newIntValue(uint256(_destination));
            msgValues[1] = ArbValue.newIntValue(_seq_number);
            msgValues[2] = ArbValue.newIntValue(_value);
            msgValues[3] = ArbValue.newHashOnly(dataHash);

            ArbValue.Value[] memory msgType = new ArbValue.Value[](3);
            msgType[0] = ArbValue.newIntValue(TRANSACTION_MSG);
            msgType[1] = ArbValue.newIntValue(uint256(_sender));
            msgType[2] = ArbValue.newTupleValue(msgValues);

            ArbValue.Value[] memory dataMsg = new ArbValue.Value[](4);
            dataMsg[0] = ArbValue.newIntValue(block.timestamp);
            dataMsg[1] = ArbValue.newIntValue(block.number);
            dataMsg[2] = ArbValue.newIntValue(uint256(txHash));
            dataMsg[3] = ArbValue.newTupleValue(msgType);

            bytes32 messageHash =  ArbValue.newTupleValue(dataMsg).hash().hash;

            _deliverMessage(_destination, messageHash);
        }

        emit IGlobalPendingInbox.TransactionMessageDelivered(
            _sender,
            _destination,
            _seq_number.
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

            ArbValue.Value[] memory msgValues = new ArbValue.Value[](3);
            msgValues[0] = ArbValue.newIntValue(uint256(_destination));
            msgValues[1] = ArbValue.newIntValue(_value);

            ArbValue.Value[] memory msgType = new ArbValue.Value[](3);
            msgType[0] = ArbValue.newIntValue(_messageType);
            msgType[1] = ArbValue.newIntValue(uint256(_sender));
            msgType[2] = ArbValue.newTupleValue(msgValues);

            ArbValue.Value[] memory ethMsg = new ArbValue.Value[](4);
            ethMsg[0] = ArbValue.newIntValue(block.timestamp);
            ethMsg[1] = ArbValue.newIntValue(block.number);
            ethMsg[2] = ArbValue.newIntValue(uint256(txHash));
            ethMsg[3] = ArbValue.newTupleValue(msgType);

            bytes32 messageHash =  ArbValue.newTupleValue(ethMsg).hash().hash;

            _deliverMessage(_destination, messageHash);
        }
    }

    function _deliverERCTokenMessage(
        address _sender,
        address _destination,
        uint256 _messageType,
        address _contractAddress,
        uint256 _value) private
    {
        if (pending[_destination] != 0) 
        {
            bytes32 txHash = keccak256(
                abi.encodePacked(
                    _sender,
                    _destination,
                    _contractAddress,
                    _value,
                    pending[_destination]
                )
            );

            ArbValue.Value[] memory msgValues = new ArbValue.Value[](3);
            msgValues[0] = ArbValue.newIntValue(uint256(_contractAddress));
            msgValues[1] = ArbValue.newIntValue(uint256(_destination));
            msgValues[2] = ArbValue.newIntValue(_value);

            ArbValue.Value[] memory msgType = new ArbValue.Value[](3);
            msgType[0] = ArbValue.newIntValue(_messageType);
            msgType[1] = ArbValue.newIntValue(uint256(_sender));
            msgType[2] = ArbValue.newTupleValue(msgValues);

            ArbValue.Value[] memory ercTokenMsg = new ArbValue.Value[](4);
            ercTokenMsg[0] = ArbValue.newIntValue(block.timestamp);
            ercTokenMsg[1] = ArbValue.newIntValue(block.number);
            ercTokenMsg[2] = ArbValue.newIntValue(uint256(txHash));
            ercTokenMsg[3] = ArbValue.newTupleValue(msgType);

            bytes32 messageHash =  ArbValue.newTupleValue(ercTokenMsg).hash().hash;

            _deliverMessage(_destination, messageHash);
        }

        if(_messageType == ERC20_DEPOSIT){

            emit IGlobalPendingInbox.DepositERC20MessageDelivered(
                _destination,
                msg.sender,
                _tokenContract,
                _value);

        }else if(_messageType == ERC20_WITHDRAWAL){

            emit IGlobalPendingInbox.WithdrawERC20MessageDelivered(
                msg.sender,
                _destination,
                _tokenContract,
                _value);

        }else if(_messageType == ERC721_DEPOSIT){

            emit IGlobalPendingInbox.DepositERC721MessageDelivered(
                _destination,
                msg.sender,
                _tokenContract,
                _value);

        }else if(_messageType == ERC721_WITHDRAWAL){

            emit IGlobalPendingInbox.WithdrawERC721MessageDelivered(
                msg.sender,
                _destination,
                _tokenContract,
                _value);
        }
    }

    function _deliverMessage(
        address _destination,
        bytes32 messageHash) private 
    {
        pending[_destination] = ArbValue.hashTupleValue([
            ArbValue.newIntValue(0),
            ArbValue.newHashOnlyValue(pending[_destination]),
            ArbValue.newHashOnlyValue(messageHash)
        ]);
    }
}
