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

pragma solidity >=0.4.21 <0.7.0;

contract Simple {
    uint256 x;
    uint256 public y;
    event TestEvent(uint256 value);

    constructor() public payable {
        y = msg.value;
        emit TestEvent(msg.value);
    }

    receive() external payable {
        require(false, "no deposits");
    }

    function exists() external payable returns (uint256) {
        x = 5;
        emit TestEvent(msg.value);
        return 10;
    }

    function reverts() external payable {
        require(false, "this is a test");
    }

    function acceptPayment() external payable {}

    function rejectPayment() external {}

    function nestedCall(uint256 value) external {
        address(this).call{ value: value }("");
    }
}
