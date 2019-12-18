/*
 * Copyright 2019, Offchain Labs, Inc.
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

pragma solidity ^0.5.3;

import "../libraries/CloneFactory.sol";

import "./IArbRollup.sol";


contract ArbFactory is CloneFactory {
    event RollupCreated(
        address vmAddress
    );

    address rollupTemplate;
    address globalInboxAddress;
    address challengeFactoryAddress;

    constructor(
        address _rollupTemplate,
        address _globalInboxAddress,
        address _challengeFactoryAddress
    )
        public
    {
        rollupTemplate = _rollupTemplate;
        globalInboxAddress = _globalInboxAddress;
        challengeFactoryAddress = _challengeFactoryAddress;
    }

    function createRollup(
        bytes32 _vmState,
        uint32 _gracePeriod,
        uint32 _maxExecutionSteps,
        uint128 _stakeRequirement,
        address payable _owner
    )
        public
    {
        address clone = createClone(rollupTemplate);
        IArbRollup(clone).init(
            _vmState,
            _gracePeriod,
            _maxExecutionSteps,
            _stakeRequirement,
            _owner,
            challengeFactoryAddress,
            globalInboxAddress
        );
        emit RollupCreated(
            clone
        );
    }
}
