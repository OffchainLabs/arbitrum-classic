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

interface IGlobalPendingInbox {

    event TransactionMessageDelivered(
        address indexed vmSenderId,
        address indexed vmReceiverId,
        address contactAddress,
        uint256 seqNumber,
        uint256 value,
        bytes data
    );

    event EthDepositMessageDelivered(
        address indexed vmReceiverId,
        address sender,
        address destination,
        uint256 value
    );

    event ERC20DepositMessageDelivered(
        address indexed vmReceiverId,
        address sender,
        address destination,
        address tokenAddress,
        uint256 value
    );

    event ERC721DepositMessageDelivered(
        address indexed vmReceiverId,
        address sender,
        address destination,
        address tokenAddress,
        uint256 value
    );

    function getPending() external returns(bytes32, uint);

    function sendMessages(bytes calldata _messages) external;

    function registerForInbox() external;

    function depositEthMessage(
        address _vmAddress, 
        address _destination) external payable;

    function forwardTransactionMessage(
        address _vmAddress,
        address _contractAddress,
        uint256 _seqNumber,
        uint256 _value,
        bytes calldata _data,
        bytes calldata _signature) external;

    function sendTransactionMessage(
        address _vmAddress, 
        address _contractAddress, 
        uint256 _seqNumber,
        uint256 _value,
        bytes calldata _data) external;

    function depositERC20Message(
        address _vmAddress,
        address _tokenContract,
        address _destination,
        uint256 _value) external;

    function depositERC721Message(
        address _vmAddress,
        address _tokenContract,
        address _destination,
        uint256 _value) external;
}
