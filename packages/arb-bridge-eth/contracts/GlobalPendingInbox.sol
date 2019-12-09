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
    uint8 internal constant DATA_MSGTYPE = 0;
    uint8 internal constant ETH_MSGTYPE = 1;
    uint8 internal constant ERC20_MSGTYPE = 2;
    uint8 internal constant ERC721_MSGTYPE = 3;

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
        uint256 destination;
        uint256 value;
        uint256 tokenContract;
        bytes memory messageData;
        uint totalLength = _messages.length;

        while (offset < totalLength) {

            uint256 messageType = _messages[offset]; //check

            if(messageType == DATA_MSGTYPE){

                (   valid,
                    offset,
                    messageHash,
                    destination,
                    messageData
                ) = Value.deserializeDataMessage(_messages, offset);
            
                if(valid){
                    _deliverDataMessage(
                        msg.sender,
                        address(bytes20(bytes32(destination))),
                        messageData
                    );
                }

            }else if(messageType == ETH_MSGTYPE){

                (   valid,
                    offset,
                    messageHash,
                    destination,
                    value,
                    messageData
                ) = Value.deserializeEthMessage(_messages, offset);

                if(valid){
                    transferEthMessage(destination, value, messageData);
                }   

            }else{

                (   valid,
                    offset,
                    messageHash,
                    destination,
                    value,
                    tokenAddress
                ) = Value.deserializeERCTokenMessage(_messages, offset);

                if(valid){

                    if(messageType == ERC20_MSGTYPE){
                        transferERC20Message(
                            tokenAddress,
                            msg.sender, 
                            destination, 
                            value);
                    }else{
                        ransferERC721Message(
                            tokenAddress,
                            msg.sender, 
                            destination, 
                            value);
                    }
                }
            }
        }
    }

    function forwardMessage(
        address _destination,
        bytes calldata _data,
        bytes calldata _signature) external
    {
        address sender = SigUtils.recoverAddress(
            keccak256(
                abi.encodePacked(
                    _destination,
                    ArbValue.deserializeValueHash(_data)
                )
            ),
            _signature
        );

        _deliverDataMessage(
            sender,
            _destination,
            _data
        );
    }

    function sendMessage(address _destination, bytes calldata _data) external
    {
        _deliverDataMessage(
            msg.sender,
            _destination,
            _data
        );
    }

    function forwardERC20Message(
        address _tokenContract,
        address _sender,
        address _destination,
        uint256 _value,
        bytes calldata _signature) external
    {
        address sender = SigUtils.recoverAddress(
            keccak256(
                abi.encodePacked(
                    _tokenContract,
                    _sender,
                    _destination,
                    _value
                )
            ),
            _signature
        );

        transferERC20Message(_tokenContract, sender, _destination, _value);
    }

    function forwardERC721Message(
        address _tokenContract,
        address _sender,
        address _destination,
        uint256 _value,
        bytes calldata _signature) external
    {
        address sender = SigUtils.recoverAddress(
            keccak256(
                abi.encodePacked(
                    _tokenContract,
                    _sender,
                    _destination,
                    _value
                )
            ),
            _signature
        );

        transferERC721Message(_tokenContract, sender, _destination, _value);
    }

    function transferERC20Message(
        address _tokenContract,
        address _sender,
        address _destination,
        uint256 _value) external
    {
        transferToken(
            _sender,
            _destination,
            _tokenContract),
            _value
        );

        _deliverERCTokenMessage(
            _sender,
            _destination,
            ERC20_MSGTYPE,
            _tokenContract,
            _value
        );
    }

    function transferERC721Message(
        address _tokenContract,
        address _sender,
        address _destination,
        uint256 _value) external
    {
        transferNFT(
            _sender,
            _destination,
            _tokenContract,
            _value
        );

        _deliverERCTokenMessage(
            _sender,
            _destination,
            ERC721_MSGTYPE,
            _tokenContract,
            _value
        );
    }  

    function despositEthMessage(address _destination, bytes calldata _data) external payable 
    {
        depositEth(_destination);
        
        _deliverEthMessage(
            _destination,
            msg.value,
            msg.sender,
            _data
        );
    }

    function transferEthMessage(
        address _destination, 
        uint256 _value, 
        bytes calldata _data) external
    {
        transferEth(_destination, _value);

        _deliverEthMessage(
            _destination,
            _value,
            msg.sender,
            _data
        );
    }

    function _deliverDataMessage(
        address _sender,
        address _destination,
        bytes memory _data) private
    {
        if (pending[_destination] != 0) 
        {
            bytes32 dataHash = ArbValue.deserializeValueHash(_data);
            bytes32 txHash = keccak256(
                abi.encodePacked(
                    _sender,
                    _destination,
                    dataHash,
                    pending[_destination]
                )
            );

            ArbValue.Value[] memory msgValues = new ArbValue.Value[](1);
            msgValues[0] = ArbValue.newHashOnly(dataHash);

            ArbValue.Value[] memory values = new ArbValue.Value[](5);
            values[0] = ArbValue.newIntValue(DATA_MSGTYPE);
            values[1] = ArbValue.newIntValue(uint256(txHash));
            values[2] = ArbValue.newIntValue(block.timestamp);
            values[3] = ArbValue.newIntValue(block.number);
            values[4] = ArbValue.newIntValue(uint256(_sender));

            bytes32 messageHash =  ArbValue.newTupleValue(values).hash().hash;

            _deliverMessage(_destination, messageHash);
        }

        emit IGlobalPendingInbox.DataMessageDelivered(_destination, msg.sender, _data);
    }

    function _deliverEthMessage(
        address _destination,
        uint256 _value,
        address _sender,
        bytes memory _data) private
    {
        if (pending[_destination] != 0) 
        {
            bytes32 dataHash = ArbValue.deserializeValueHash(_data);
            bytes32 txHash = keccak256(
                abi.encodePacked(
                    _sender,
                    _destination,
                    dataHash,
                    _value,
                    pending[_destination]
                )
            );

            ArbValue.Value[] memory msgValues = new ArbValue.Value[](2);
            msgValues[0] = ArbValue.newIntValue(_value);
            msgValues[1] = ArbValue.newHashOnly(dataHash);

            ArbValue.Value[] memory values = new ArbValue.Value[](6);
            values[0] = ArbValue.newIntValue(ETH_MSGTYPE);
            values[1] = ArbValue.newIntValue(uint256(txHash));
            values[2] = ArbValue.newIntValue(block.timestamp);
            values[3] = ArbValue.newIntValue(block.number);
            values[4] = ArbValue.newIntValue(uint256(_sender));
            values[5] = ArbValue.newTupleValue(msgValues);

            bytes32 messageHash =  ArbValue.newTupleValue(values).hash().hash;

            _deliverMessage(_destination, messageHash);
        }

        emit IGlobalPendingInbox.EthMessageDelivered(_destination, _sender, _value, _data);
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

            ArbValue.Value[] memory msgValues = new ArbValue.Value[](2);
            msgValues[0] = ArbValue.newIntValue(uint256(_contractAddress));
            msgValues[1] = ArbValue.newIntValue(_value);

            ArbValue.Value[] memory values = new ArbValue.Value[](6);
            values[0] = ArbValue.newIntValue(_messageType);
            values[1] = ArbValue.newIntValue(uint256(txHash));
            values[2] = ArbValue.newIntValue(block.timestamp);
            values[3] = ArbValue.newIntValue(block.number);
            values[4] = ArbValue.newIntValue(uint256(_sender));
            values[5] = ArbValue.newTupleValue(msgValues);

            bytes32 messageHash =  ArbValue.newTupleValue(values).hash().hash;

            _deliverMessage(_destination, messageHash);
        }

        emit IGlobalPendingInbox.ERCTokenMessageDelivered(_destination, _sender, _messageType, _contractAddress, _value);
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
