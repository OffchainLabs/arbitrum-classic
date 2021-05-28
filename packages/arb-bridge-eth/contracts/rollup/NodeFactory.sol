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

import "./Node.sol";
import "./INodeFactory.sol";

import "@openzeppelin/contracts/proxy/BeaconProxy.sol";
import "@openzeppelin/contracts/proxy/UpgradeableBeacon.sol";

contract NodeFactory is INodeFactory {
    UpgradeableBeacon public beacon;

    constructor() public {
        address templateContract = address(new Node());
        beacon = new UpgradeableBeacon(templateContract);
        beacon.transferOwnership(msg.sender);
    }

    function createNode(
        bytes32 _stateHash,
        bytes32 _challengeHash,
        bytes32 _confirmData,
        uint256 _prev,
        uint256 _deadlineBlock
    ) external override returns (address) {
        address clone = address(new BeaconProxy(address(beacon), ""));
        Node(clone).initialize(
            msg.sender,
            _stateHash,
            _challengeHash,
            _confirmData,
            _prev,
            _deadlineBlock
        );
        return address(clone);
    }
}
