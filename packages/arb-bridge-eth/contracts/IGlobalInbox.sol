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

interface IGlobalInbox {

    event TransactionMessageDelivered(
        address indexed chain,
        address indexed to,
        address indexed from,
        uint256 seqNumber,
        uint256 value,
        bytes data
    );

    event EthDepositMessageDelivered(
        address indexed chain,
        address indexed to,
        address indexed from,
        uint256 value,
        uint256 messageNum
    );

    event ERC20DepositMessageDelivered(
        address indexed chain,
        address indexed to,
        address indexed from,
        address erc20,
        uint256 value,
        uint256 messageNum
    );

    event ERC721DepositMessageDelivered(
        address indexed chain,
        address indexed to,
        address indexed from,
        address erc721,
        uint256 id,
        uint256 messageNum
    );

    event ContractTransactionMessageDelivered(
        address indexed chain,
        address indexed to,
        address indexed from,
        uint256 value,
        bytes data,
        uint256 messageNum
    );

    function getInbox(address account) external view returns(bytes32, uint);

    function sendMessages(bytes calldata _messages, uint[] calldata messageCounts, bytes32[] calldata nodeHashes) external;

    function depositEthMessage(address _chain, address _to) external payable;

    function sendTransactionMessage(
        address _chain,
        address _to,
        uint256 _seqNumber,
        uint256 _value,
        bytes calldata _data
    )
        external;

    function depositERC20Message(
        address _chain,
        address _to,
        address _erc20,
        uint256 _value
    )
        external;

    function depositERC721Message(
        address _chain,
        address _to,
        address _erc721,
        uint256 _value
    )
        external;

    // msg.sender is the chain receiving the message
    function forwardContractTransactionMessage(
        address _to,
        address _from,
        uint256 _value,
        bytes calldata _data
    )
        external;

    // msg.sender is the chain receiving the message
    function forwardEthMessage(address _to, address _from) external payable;
}
