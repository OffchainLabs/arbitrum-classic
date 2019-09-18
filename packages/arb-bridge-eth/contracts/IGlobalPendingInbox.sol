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

    event MessageDelivered(
        address indexed vmId,
        address sender,
        bytes21 tokenType,
        uint256 value,
        bytes data
    );

    function pullPendingMessages() external returns(bytes32);

    function sendMessages(
        bytes21[] calldata _tokenTypes,
        bytes calldata _messageData,
        uint16[] calldata _tokenTypeNum,
        uint256[] calldata _amounts,
        address[] calldata _destinations
    )
        external;

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

    // This function assumes that tokenTypes and amounts are valid and in canonical
    // order. The pair (tokenType, amount) are sorted in ascending order with
    // tokenTypes as the primary key and amount as the secondary key
    // Token type only allows repeats for NFTs and amounts disallow repeats for NFTs
    function hasFunds(
        address _owner,
        bytes21[] calldata _tokenTypes,
        uint256[] calldata _amounts
    )
        external
        view
        returns(bool);
}
