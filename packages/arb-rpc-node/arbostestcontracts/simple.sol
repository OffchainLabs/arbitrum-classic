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

contract ComplexConstructorCon {
    constructor(bytes32 salt) public payable {
        Simple(msg.sender).exists();
        new ComplexConstructorCon2{ salt: salt, value: msg.value / 2 }(654);
        Simple(msg.sender).nestedCall(54);
    }

    receive() external payable {}

    function getVal() external returns (uint256) {
        return 20;
    }
}

contract ComplexConstructorCon2 {
    constructor(uint256 val) public payable {
        msg.sender.transfer(msg.value / 2);
    }

    function getVal() external returns (uint256) {
        return 20;
    }
}

contract Reverter {
    constructor() public {
        require(false, "Intentional revert");
    }
}

contract Simple {
    uint256 x;
    uint256 public y;
    uint256[] array;

    event TestEvent(uint256 value, address sender);

    constructor() public payable {
        y = msg.value;
        emit TestEvent(msg.value, msg.sender);
    }

    receive() external payable {
        require(false, "no deposits");
    }

    function exists() external payable returns (uint256) {
        x = 5;
        emit TestEvent(msg.value, msg.sender);
        return 10;
    }

    function arrayPush() external payable returns (uint256) {
        array.push(x+1);
        emit TestEvent(msg.value, msg.sender);
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

    function nestedCall2(uint256 value, address dest) external returns (bytes memory) {
        Simple(this).exists();
        (, bytes memory data) = address(dest).call{ value: value }("");
        return data;
    }

    function crossCall(Simple con) external returns (uint256) {
        return con.exists() + 1;
    }

    function trace(uint256 arg) external payable returns (uint256) {
        ComplexConstructorCon con = new ComplexConstructorCon{ value: msg.value / 2 }("0x43254");
        try new Reverter() {} catch {}
        return con.getVal();
    }

    event Variable(bool[10] _variable);

    function debug() external {
        bool[10] memory seen;

        emit Variable(seen);
    }
}
