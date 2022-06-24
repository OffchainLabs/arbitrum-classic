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

pragma solidity >=0.4.21 <0.9.0;

contract ComplexConstructorCon {
    constructor(bytes32 salt) payable {
        Simple(payable(msg.sender)).exists();
        new ComplexConstructorCon2{ salt: salt, value: msg.value / 2 }(654);
        Simple(payable(msg.sender)).nestedCall(54);
    }

    receive() external payable {}

    function getVal() external returns (uint256) {
        return 20;
    }
}

contract ComplexConstructorCon2 {
    constructor(uint256 val) payable {
        payable(msg.sender).transfer(msg.value / 2);
    }

    function getVal() external returns (uint256) {
        return 20;
    }
}

contract Reverter {
    constructor() {
        require(false, "Intentional revert");
    }
}

contract Destroyer1 {
    constructor() {
        selfdestruct(payable(msg.sender));
    }
}

contract Destroyer2 {
    function destroy() public {
        selfdestruct(payable(msg.sender));
    }

    function test1() public payable returns (uint256) {
        return 10;
    }

    function test2(address to) public payable returns (uint256) {
        (bool success, bytes memory data) = to.delegatecall(
            abi.encodeWithSelector(this.test1.selector)
        );
        require(success);
        uint256 ret = abi.decode(data, (uint256));
        return ret;
    }

    function test3(address to) public payable returns (uint256 c) {
        bytes memory input = abi.encodeWithSelector(this.test1.selector);
        uint256 inputLength = input.length;
        bool success;
        assembly {
            let d := add(input, 32)
            let ret := mload(0x40)
            success := callcode(gas(), to, 0, d, inputLength, ret, 0x20)
            c := mload(ret)
        }
        require(success);
        return c;
    }

    function test4(address to) public payable returns (uint256) {
        return Destroyer2(to).test1();
    }
}

contract Simple {
    uint256 x;
    uint256 public y;
    uint256[] array;

    event TestEvent(uint256 value, address sender);

    constructor() payable {
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
        array.push(x + 1);
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
        new Destroyer1();
        Destroyer2 test = new Destroyer2();
        Destroyer2 test2 = new Destroyer2();

        bytes memory input1 = abi.encodeWithSelector(Destroyer2.test2.selector, test2);
        bytes memory input2 = abi.encodeWithSelector(Destroyer2.test3.selector, test2);
        uint256 inputLength = input1.length;
        address to = address(test);
        assembly {
            let d := add(input1, 32)
            let ret := mload(0x40)
            let success := callcode(gas(), to, 0, d, inputLength, ret, 0x20)
            mstore(0x40, add(ret, 0x44))
        }
        assembly {
            let d := add(input2, 32)
            let ret := mload(0x40)
            let success := callcode(gas(), to, 0, d, inputLength, ret, 0x20)
            mstore(0x40, add(ret, 0x44))
        }
        to.delegatecall(input1);
        to.delegatecall(input2);
        to.delegatecall(abi.encodeWithSelector(Destroyer2.test4.selector, test2));
        test.test2(address(test2));
        return con.getVal();
    }

    event Variable(bool[10] _variable);

    function debug() external {
        bool[10] memory seen;

        emit Variable(seen);
    }
}
