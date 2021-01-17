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

import "./Inbox.sol";
import "./Outbox.sol";

import "./interfaces/IBridge.sol";

contract Bridge is Ownable, IBridge {
    mapping(address => bool) public override allowedInboxes;
    mapping(address => bool) public override allowedOutboxes;

    address public override activeOutbox;

    bytes32 private inboxMaxAcc;
    uint256 private inboxMaxCount;

    function deliverMessageToInbox(
        uint8 kind,
        address sender,
        bytes32 messageDataHash
    ) external payable override returns (uint256) {
        require(allowedInboxes[msg.sender], "NOT_FROM_INBOX");
        uint256 count = inboxMaxCount;
        bytes32 inboxAcc = inboxMaxAcc;
        bytes32 messageHash =
            Messages.messageHash(
                kind,
                sender,
                block.number,
                block.timestamp, // solhint-disable-line not-rely-on-time
                count,
                messageDataHash
            );
        inboxMaxAcc = Messages.addMessageToInbox(inboxAcc, messageHash);
        inboxMaxCount = count + 1;
        emit MessageDelivered(count, inboxAcc, msg.sender, kind, sender, messageDataHash);
        return count;
    }

    function executeCall(
        address destAddr,
        uint256 amount,
        bytes calldata data
    ) external override returns (bool success, bytes memory returnData) {
        require(allowedOutboxes[msg.sender], "NOT_FROM_OUTBOX");
        activeOutbox = msg.sender;
        (success, returnData) = destAddr.call{ value: amount }(data);
        activeOutbox = address(0);
    }

    function setInbox(address inbox, bool enabled) external override onlyOwner {
        allowedInboxes[inbox] = enabled;
    }

    function setOutbox(address inbox, bool enabled) external override onlyOwner {
        allowedOutboxes[inbox] = enabled;
    }

    function inboxInfo() external view override returns (uint256, bytes32) {
        return (inboxMaxCount, inboxMaxAcc);
    }
}
