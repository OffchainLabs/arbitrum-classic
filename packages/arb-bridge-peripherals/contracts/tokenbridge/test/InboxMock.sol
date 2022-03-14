// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2020, Offchain Labs, Inc.
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

import "arb-bridge-eth/contracts/bridge/interfaces/IOutbox.sol";
import "arb-bridge-eth/contracts/bridge/interfaces/IBridge.sol";
import "arb-bridge-eth/contracts/bridge/interfaces/IInbox.sol";

contract InboxMock {
    address l2ToL1SenderMock = address(0);

    event TicketData(uint256 maxSubmissionCost);
    event RefundAddresses(address excessFeeRefundAddress,address callValueRefundAddress);

    function createRetryableTicket(
        address, /* destAddr */
        uint256, /* l2CallValue */
        uint256 maxSubmissionCost,
        address excessFeeRefundAddress,
        address callValueRefundAddress,
        uint256, /* maxGas */
        uint256, /* gasPriceBid */
        bytes calldata /* data */
    ) external payable returns (uint256) {
        emit TicketData(maxSubmissionCost);
        emit RefundAddresses(excessFeeRefundAddress,callValueRefundAddress);
        return 0;
    }

    function bridge() external view returns (IBridge) {
        return IBridge(address(this));
    }

    function activeOutbox() external view returns (address) {
        return address(this);
    }

    function setL2ToL1Sender(address sender) external {
        l2ToL1SenderMock = sender;
    }

    function l2ToL1Sender() external view returns (address) {
        return l2ToL1SenderMock;
    }
}
