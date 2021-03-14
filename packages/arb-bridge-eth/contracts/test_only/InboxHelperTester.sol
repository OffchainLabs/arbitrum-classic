// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2012, Offchain Labs, Inc.
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

import "../bridge/InboxHelper.sol";

contract InboxHelperTester {
    function chainId(address rollup) external pure returns (uint256) {
        return InboxHelper.chainId(rollup);
    }

    function requestID(uint256 messageNum, address rollup) external pure returns (bytes32) {
        return InboxHelper.requestID(messageNum, rollup);
    }

    function retryableTicketID(uint256 messageNum, address rollup) external pure returns (bytes32) {
        return InboxHelper.retryableTicketID(messageNum, rollup);
    }
}
