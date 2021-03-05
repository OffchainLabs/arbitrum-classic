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

import "arb-bridge-eth/contracts/bridge/interfaces/IInbox.sol";
import "arb-bridge-eth/contracts/bridge/interfaces/IOutbox.sol";
import "arb-bridge-eth/contracts/bridge/interfaces/IBridge.sol";

import "../ethereum/L1Buddy.sol";
import "../util/BuddyUtil.sol";

contract TestConstructorBuddy is L1Buddy {
    address public l2Buddy;

    constructor(
        address _inbox,
        address _l2Deployer,
        uint256 _maxGas,
        uint256 _gasPrice,
        bytes memory _deployCode
    )
        L1Buddy(_inbox, _l2Deployer)
        public
    {
        L1Buddy.initiateBuddyDeploy(_maxGas, _gasPrice, _deployCode);
    }

    function handleDeploySuccess() internal override {
        l2Buddy = BuddyUtil.calculateL2Address(
            address(L1Buddy.l2Deployer),
            address(this),
            L1Buddy.codeHash
        );
        // this deletes the codehash from state!
        L1Buddy.handleDeploySuccess();
    }

    function handleDeployFail() internal override {}
}

contract TestBuddy is L1Buddy {
    address public l2Buddy;

    constructor(
        address _inbox,
        address _l2Deployer
    )
        L1Buddy(_inbox, _l2Deployer)
        public
    {}

    function handleDeploySuccess() internal override {
        l2Buddy = BuddyUtil.calculateL2Address(
            address(L1Buddy.l2Deployer),
            address(this),
            L1Buddy.codeHash
        );
        // this deletes the codehash from state!
        L1Buddy.handleDeploySuccess();
    }
    function handleDeployFail() internal override {}
}
