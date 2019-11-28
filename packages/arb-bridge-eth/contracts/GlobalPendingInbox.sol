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

import "./libraries/ArbValue.sol";
import "./libraries/SigUtils.sol";

import "@openzeppelin/contracts/ownership/Ownable.sol";


contract GlobalPendingInbox is GlobalWallet, IGlobalPendingInbox {
    using SafeMath for uint256;
    using ArbValue for ArbValue.Value;

    address internal constant ETH_ADDRESS = address(0);

    mapping(address => bytes32) pending;

    function pullPendingMessages() external returns(bytes32) {
        bytes32 messages = pending[msg.sender];
        pending[msg.sender] = ArbValue.hashEmptyTuple();
        return messages;
    }

    function sendMessages(bytes calldata _messages) external {
        uint offset = 0;
        bool valid;
        bytes32 messageHash;
        uint256 destination;
        uint256 value;
        uint256 tokenType;
        bytes memory messageData;
        uint totalLength = _messages.length;
        while (offset < totalLength) {
            (
                valid,
                offset,
                messageHash,
                destination,
                value,
                tokenType,
                messageData
            ) = ArbValue.deserializeMessage(_messages, offset);
            if (valid) {
                _sendUnpaidMessage(
                    address(bytes20(bytes32(destination))),
                    bytes21(bytes32(tokenType)),
                    value,
                    msg.sender,
                    messageData
                );
            }
        }
    }

    function registerForInbox() external {
        require(pending[msg.sender] == 0, "Pending must be uninitialized");
        pending[msg.sender] = ArbValue.hashEmptyTuple();
    }

    function sendMessage(
        address _destination,
        bytes21 _tokenType,
        uint256 _amount,
        bytes calldata _data
    )
        external
    {
        _sendUnpaidMessage(
            _destination,
            _tokenType,
            _amount,
            msg.sender,
            _data
        );
    }

    function forwardMessage(
        address _destination,
        bytes21 _tokenType,
        uint256 _amount,
        bytes calldata _data,
        bytes calldata _signature
    )
        external
    {
        address sender = SigUtils.recoverAddress(
            keccak256(
                abi.encodePacked(
                    _destination,
                    ArbValue.deserializeValueHash(_data),
                    _amount,
                    _tokenType
                )
            ),
            _signature
        );

        _sendUnpaidMessage(
            _destination,
            _tokenType,
            _amount,
            sender,
            _data
        );
    }

    function sendEthMessage(address _destination, bytes calldata _data) external payable {
        depositEth(_destination);
        _deliverMessage(
            _destination,
            bytes21(0),
            msg.value,
            msg.sender,
            _data
        );
    }

    function _sendUnpaidMessage(
        address _destination,
        bytes21 _tokenType,
        uint256 _value,
        address _sender,
        bytes memory _data
    )
        private
    {
        bool sent = false;
        if (_tokenType[20] == 0x01) {
            sent = transferNFT(
                _sender,
                _destination,
                address(bytes20(_tokenType)),
                _value
            );
        } else {
            sent = transferToken(
                _sender,
                _destination,
                address(bytes20(_tokenType)),
                _value
            );
        }
        if (sent) {
            _deliverMessage(
                _destination,
                _tokenType,
                _value,
                _sender,
                _data
            );
        }
    }

    function generateSentMessageHash(
        address _dest,
        bytes32 _data,
        bytes21 _tokenType,
        uint256 _value,
        address _sender
    )
        public
        view
        returns (bytes32)
    {

    }

    function _deliverMessage(
        address _destination,
        bytes21 _tokenType,
        uint256 _value,
        address _sender,
        bytes memory _data
    )
        private
    {
        if (pending[_destination] != 0) {
            bytes32 dataHash = ArbValue.deserializeValueHash(_data);
            bytes32 txHash = keccak256(
                abi.encodePacked(
                    _destination,
                    dataHash,
                    _value,
                    _tokenType
                )
            );
            ArbValue.Value[] memory dataValues = new ArbValue.Value[](4);
            dataValues[0] = ArbValue.newHashOnlyValue(dataHash);
            dataValues[1] = ArbValue.newIntValue(block.timestamp);
            dataValues[2] = ArbValue.newIntValue(block.number);
            dataValues[3] = ArbValue.newIntValue(uint(txHash));

            ArbValue.Value[] memory values = new ArbValue.Value[](4);
            values[0] = ArbValue.newTupleValue(dataValues);
            values[1] = ArbValue.newIntValue(uint256(_sender));
            values[2] = ArbValue.newIntValue(_value);
            values[3] = ArbValue.newIntValue(uint256(bytes32(_tokenType)));
            bytes32 messageHash =  ArbValue.newTupleValue(values).hash().hash;

            pending[_destination] = ArbValue.hashTupleValue([
                ArbValue.newIntValue(0),
                ArbValue.newHashOnlyValue(pending[_destination]),
                ArbValue.newHashOnlyValue(messageHash)
            ]);
        }

        emit IGlobalPendingInbox.MessageDelivered(
            _destination,
            _sender,
            _tokenType,
            _value,
            _data
        );
    }
}
