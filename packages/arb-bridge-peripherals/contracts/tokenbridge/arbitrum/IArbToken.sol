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

interface IArbToken {
    function initialize(
        address _bridge,
        address _l1Address,
        uint8 _decimals
    ) external;

    function bridgeMint(address account, uint256 amount, bytes memory data) external;

    function withdraw(address destination, uint256 amount) external;

    /// @dev This function is optional. If it's not enabled, this data won't be updatable based on its paired L1 contract
    function updateInfo(string calldata newName, string calldata newSymbol, uint8 newDecimals) external;
}
