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

import "@openzeppelin/contracts/access/Ownable.sol";
import "../libraries/Cloneable.sol";

import "./interfaces/IOutboxEntry.sol";

contract OutboxEntry is Ownable, Cloneable, IOutboxEntry {
    bytes32 outputRoot;
    mapping(uint256 => bool) spentOutput;

    function initialize(bytes32 root) external override {
        require(outputRoot != 0, "ALREADY_INIT");
        outputRoot = root;
    }

    function spendOutput(bytes32 calcRoot, uint256 path) external override onlyOwner {
        require(!spentOutput[path], "ALREADY_SPENT");
        require(calcRoot == outputRoot, "BAD_PROOF");
        spentOutput[path] = true;
    }
}
