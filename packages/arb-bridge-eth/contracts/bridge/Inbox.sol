// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2019-2021, Offchain Labs, Inc.
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

import "./interfaces/IInbox.sol";
import "./interfaces/IBridge.sol";

import "./Messages.sol";

contract Inbox is IInbox {
    uint8 internal constant ETH_TRANSFER = 0;
    uint8 internal constant L2_MSG = 3;
    uint8 internal constant L1MessageType_L2FundedByL1 = 7;

    uint8 internal constant L2MessageType_unsignedEOATx = 0;
    uint8 internal constant L2MessageType_unsignedContractTx = 1;

    IBridge public override bridge;

    constructor(IBridge _bridge) public {
        bridge = _bridge;
    }

    /**
     * @notice Send a generic L2 message to the chain
     * @dev This method is an optimization to avoid having to emit the entirety of the messageData in a log. Instead validators are expected to be able to parse the data from the transaction's input
     * @param messageData Data of the message being sent
     */
    function sendL2MessageFromOrigin(bytes calldata messageData) external {
        // solhint-disable-next-line avoid-tx-origin
        require(msg.sender == tx.origin, "origin only");
        uint256 msgNum = deliverToBridge(L2_MSG, msg.sender, keccak256(messageData));
        emit InboxMessageDeliveredFromOrigin(msgNum);
    }

    /**
     * @notice Send a generic L2 message to the chain
     * @dev This method can be used to send any type of message that doesn't require L1 validation
     * @param messageData Data of the message being sent
     */
    function sendL2Message(bytes calldata messageData) external override {
        uint256 msgNum = deliverToBridge(L2_MSG, msg.sender, keccak256(messageData));
        emit InboxMessageDelivered(msgNum, messageData);
    }

    function sendL1FundedUnsignedTransaction(
        uint256 maxGas,
        uint256 gasPriceBid,
        uint256 nonce,
        address destAddr,
        bytes calldata data
    ) external payable override {
        _deliverMessage(
            L1MessageType_L2FundedByL1,
            msg.sender,
            abi.encodePacked(
                L2MessageType_unsignedEOATx,
                maxGas,
                gasPriceBid,
                nonce,
                uint256(uint160(bytes20(destAddr))),
                msg.value,
                data
            )
        );
    }

    function sendL1FundedContractTransaction(
        uint256 maxGas,
        uint256 gasPriceBid,
        address destAddr,
        bytes calldata data
    ) external payable override {
        _deliverMessage(
            L1MessageType_L2FundedByL1,
            msg.sender,
            abi.encodePacked(
                L2MessageType_unsignedContractTx,
                maxGas,
                gasPriceBid,
                uint256(uint160(bytes20(destAddr))),
                msg.value,
                data
            )
        );
    }

    function sendUnsignedTransaction(
        uint256 maxGas,
        uint256 gasPriceBid,
        uint256 nonce,
        address destAddr,
        uint256 amount,
        bytes calldata data
    ) external override {
        _deliverMessage(
            L2_MSG,
            msg.sender,
            abi.encodePacked(
                L2MessageType_unsignedEOATx,
                maxGas,
                gasPriceBid,
                nonce,
                uint256(uint160(bytes20(destAddr))),
                amount,
                data
            )
        );
    }

    function sendContractTransaction(
        uint256 maxGas,
        uint256 gasPriceBid,
        address destAddr,
        uint256 amount,
        bytes calldata data
    ) external override {
        _deliverMessage(
            L2_MSG,
            msg.sender,
            abi.encodePacked(
                L2MessageType_unsignedContractTx,
                maxGas,
                gasPriceBid,
                uint256(uint160(bytes20(destAddr))),
                amount,
                data
            )
        );
    }

    function depositEth(address destAddr) external payable override {
        _deliverMessage(
            L1MessageType_L2FundedByL1,
            destAddr,
            abi.encodePacked(
                L2MessageType_unsignedContractTx,
                uint256(0),
                uint256(0),
                uint256(uint160(bytes20(destAddr))),
                msg.value
            )
        );
    }

    function _deliverMessage(
        uint8 _kind,
        address _sender,
        bytes memory _messageData
    ) private {
        uint256 msgNum = deliverToBridge(_kind, _sender, keccak256(_messageData));
        emit InboxMessageDelivered(msgNum, _messageData);
    }

    function deliverToBridge(
        uint8 kind,
        address sender,
        bytes32 messageDataHash
    ) private returns (uint256) {
        return bridge.deliverMessageToInbox{ value: msg.value }(kind, sender, messageDataHash);
    }
}
