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

pragma solidity ^0.5.17;

contract Receiver {
    uint256 public test = 5;
    Receiver2 public other;

    constructor(address otherReciver) public {
        other = Receiver2(otherReciver);
    }

    function mutate() external payable {
        test = 6;

        other.mutate();
    }
}

contract Receiver2 {
    uint256 public test = 7;

    function mutate() external payable {
        test = 8;
    }
}
