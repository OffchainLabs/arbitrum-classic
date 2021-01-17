// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2021, Offchain Labs, Inc.
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

pragma solidity ^0.6.11;

interface IInbox {
    event InboxMessageDelivered(uint256 indexed messageNum, uint8 kind, address sender, bytes data);

    event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum, uint8 kind, address sender);

    event BuddyContractPair(address indexed sender);

    function sendL2Message(bytes calldata messageData) external;

    function depositEthMessage(address to) external payable;

    function deployL2ContractPair(
        uint256 maxGas,
        uint256 gasPriceBid,
        uint256 payment,
        bytes calldata contractData
    ) external;
}
