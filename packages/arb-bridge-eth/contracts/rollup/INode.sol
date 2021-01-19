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

interface INode {
    function initialize(
        address _rollup,
        bytes32 _stateHash,
        bytes32 _challengeHash,
        bytes32 _confirmData,
        uint256 _prev,
        uint256 _deadlineBlock
    ) external;

    function destroy() external;

    function addStaker(address staker) external;

    function removeStaker(address staker) external;

    function childCreated() external;

    function stateHash() external view returns (bytes32);

    function challengeHash() external view returns (bytes32);

    function confirmData() external view returns (bytes32);

    function prev() external view returns (uint256);

    function deadlineBlock() external view returns (uint256);

    function stakerCount() external view returns (uint256);

    function stakers(address staker) external view returns (bool);

    function firstChildBlock() external view returns (uint256);

    function requirePastDeadline() external view;

    function requireRejectExample(uint256 latestConfirmed, address stakerAddress) external view;
}
