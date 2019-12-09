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

    event DataMessageDelivered(
        address indexed vmId,
        address sender,
        bytes data
    );

    event ERCTokenMessageDelivered(
        address indexed vmId,
        address sender,
        uint256 messageType,
        address tokenAddress,
        uint256 value
    );

    event EthMessageDelivered(
        address indexed vmId,
        address sender,
        uint256 value,
        bytes data
    );

    function pullPendingMessages() external returns(bytes32);

    function sendMessages(bytes calldata _messages) external;

    function registerForInbox() external;

    function sendMessage(
        address _destination,
        bytes21 _tokenType,
        uint256 _amount,
        bytes calldata _data
    )
        external;

    function forwardMessage(
        address _destination,
        bytes21 _tokenType,
        uint256 _amount,
        bytes calldata _data,
        bytes calldata _signature
    )
        external;

    function sendEthMessage(address _destination, bytes calldata _data) external payable;
}
