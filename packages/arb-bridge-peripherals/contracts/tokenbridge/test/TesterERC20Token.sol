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

contract TesterERC20Token {
    uint256 public decimals;
    bytes32 public name;
    bytes32 public symbol;

    constructor(
        uint256 _decimals,
        bytes32 _name,
        bytes32 _symbol
    ) public {
        decimals = _decimals;
        name = _name;
        symbol = _symbol;
    }
}

contract TesterERC20TokenNoMetadata {
    constructor() public {}
}
