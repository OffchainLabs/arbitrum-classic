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

import "./InboxCore.sol";
import "./IInbox.sol";
import "./Messages.sol";

contract Inbox is InboxCore, IInbox {
    uint8 internal constant ETH_TRANSFER = 0;
    uint8 internal constant L2_MSG = 3;
    uint8 internal constant INITIALIZATION_MSG = 4;
    uint8 internal constant L2_CONTRACT_PAIR = 5;

    /**
     * @notice Send a generic L2 message to the chain
     * @dev This method is an optimization to avoid having to emit the entirety of the messageData in a log. Instead validators are expected to be able to parse the data from the transaction's input
     * @param messageData Data of the message being sent
     */
    function sendL2MessageFromOrigin(bytes calldata messageData) external {
        // solhint-disable-next-line avoid-tx-origin
        require(msg.sender == tx.origin, "origin only");
        (uint256 msgNum, bytes32 beforeInboxAcc) =
            deliverMessageToInbox(
                Messages.messageHash(
                    L2_MSG,
                    msg.sender,
                    block.number,
                    block.timestamp, // solhint-disable-line not-rely-on-time
                    inboxMaxCount,
                    keccak256(messageData)
                )
            );
        emit MessageDeliveredFromOrigin(msgNum, beforeInboxAcc, L2_MSG, msg.sender);
    }

    /**
     * @notice Send a generic L2 message to the chain
     * @dev This method can be used to send any type of message that doesn't require L1 validation
     * @param messageData Data of the message being sent
     */
    function sendL2Message(bytes calldata messageData) external override {
        _deliverMessage(L2_MSG, msg.sender, messageData);
    }

    function deployL2ContractPair(
        uint256 maxGas,
        uint256 gasPriceBid,
        uint256 payment,
        bytes calldata contractData
    ) external override {
        require(isContract(msg.sender), "must be called by contract");
        _deliverMessage(
            L2_CONTRACT_PAIR,
            msg.sender,
            abi.encodePacked(maxGas, gasPriceBid, payment, contractData)
        );
        emit BuddyContractPair(msg.sender);
    }

    function sendInitializationMessage(bytes memory messageData) internal {
        _deliverMessage(INITIALIZATION_MSG, address(this), messageData);
    }

    /**
     * @notice Deposits ETH into the chain
     * @dev This method is payable and will deposit all value it is called with
     * @param to Address on the chain that will receive the ETH
     */
    function depositEthMessage(address to) external payable override {
        _deliverMessage(
            ETH_TRANSFER,
            msg.sender,
            abi.encodePacked(uint256(uint160(bytes20(to))), msg.value)
        );
    }

    function _deliverMessage(
        uint8 _kind,
        address _sender,
        bytes memory _messageData
    ) private {
        (uint256 msgNum, bytes32 beforeInboxAcc) =
            deliverMessageToInbox(
                Messages.messageHash(
                    _kind,
                    _sender,
                    block.number,
                    block.timestamp, // solhint-disable-line not-rely-on-time
                    inboxMaxCount,
                    keccak256(_messageData)
                )
            );
        emit MessageDelivered(msgNum, beforeInboxAcc, _kind, _sender, _messageData);
    }

    // Implementation taken from OpenZeppelin (https://github.com/OpenZeppelin/openzeppelin-contracts/blob/release-v3.1.0/contracts/utils/Address.sol)
    function isContract(address account) private view returns (bool) {
        // According to EIP-1052, 0x0 is the value returned for not-yet created accounts
        // and 0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470 is returned
        // for accounts without code, i.e. `keccak256('')`
        bytes32 codehash;

        bytes32 accountHash = 0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470;
        // solhint-disable-next-line no-inline-assembly
        assembly {
            codehash := extcodehash(account)
        }
        return (codehash != accountHash && codehash != 0x0);
    }
}
