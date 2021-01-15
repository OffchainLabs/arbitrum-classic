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

import "./Rollup.sol";
import "../bridge/IBridgeFactory.sol";

contract RollupCreator {
    event RollupCreated(address rollupAddress);

    IBridgeFactory bridgeFactory;
    address challengeFactory;
    address nodeFactory;

    constructor(
        address _bridgeFactory,
        address _challengeFactory,
        address _nodeFactory
    ) public {
        bridgeFactory = IBridgeFactory(_bridgeFactory);
        challengeFactory = _challengeFactory;
        nodeFactory = _nodeFactory;
    }

    function createRollup(
        bytes32 _machineHash,
        uint256 _challengePeriodBlocks,
        uint256 _arbGasSpeedLimitPerBlock,
        uint256 _baseStake,
        address _stakeToken,
        address _owner,
        bytes calldata _extraConfig
    ) external returns (Rollup) {
        Rollup rollup =
            new Rollup(
                _machineHash,
                _challengePeriodBlocks,
                _arbGasSpeedLimitPerBlock,
                _baseStake,
                _stakeToken,
                _owner,
                bridgeFactory.newBridge(),
                challengeFactory,
                nodeFactory,
                _extraConfig
            );
        emit RollupCreated(address(rollup));
        return rollup;
    }
}
