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

contract ConstructorCallback {
    event TestEvent(uint256 dataLength);
    event TestEvent2(address dataLength);

    constructor() public payable {
        emit TestEvent(msg.data.length);
        ConstructorCallback2(msg.sender).test2();
    }

    function test(address data) external {
        emit TestEvent2(data);
    }
}

contract ConstructorCallback2 {
    event TestEvent3(bool indexed success, bytes returnData);

    function test() external payable {
        new ConstructorCallback();
    }

    function test2() external payable {
        (bool success, bytes memory returnData) = address(msg.sender).call(
            abi.encodeWithSelector(ConstructorCallback.test.selector, msg.sender)
        );
        emit TestEvent3(success, returnData);
    }
}
