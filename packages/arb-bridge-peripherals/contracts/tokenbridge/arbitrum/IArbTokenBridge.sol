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

import "../ethereum/IEthERC20Bridge.sol";

interface IArbTokenBridge {
    event MintAndCallTriggered(
        bool success,
        address indexed sender,
        address indexed dest,
        uint256 amount,
        bytes callHookData
    );

    event WithdrawToken(
        uint256 withdrawalId,
        address indexed l1Address,
        uint256 amount,
        address indexed destination,
        uint256 indexed exitNum
    );

    event TokenCreated(address indexed l1Address, address indexed l2Address);

    event CustomTokenRegistered(address indexed l1Address, address indexed l2Address);

    event TokenMinted(
        address l1Address,
        address indexed l2Address,
        address indexed sender,
        address indexed dest,
        uint256 amount,
        bool usedCallHook
    );

    event TokenMigrated(address indexed l1Address, address indexed account, uint256 amount);

    // The following functions are only callable by the L1 bridge

    /// @notice Mints tokens in the L2
    /// @dev This function is only callable by the L1 bridge
    function mintFromL1(
        address l1ERC20,
        address sender,
        address dest,
        uint256 amount,
        bytes calldata deployData,
        bytes calldata callHookData
    ) external;

    /// @notice Registers a custom ERC20 token implementation to be used when the bridge is minting tokens
    /// @dev This function is only callable by the L1 bridge
    function customTokenRegistered(address l1Address, address l2Address) external;

    // The following functions are only callable by users in the L2

    /// @notice Migrates user balance from an erc20 token to custom token implementation
    /// @dev If a token is bridged before a custom implementation is set users can call this method to migrate to the custom version
    function migrate(
        address l1ERC20,
        address account,
        uint256 amount
    ) external;

    /// @notice Withdraws user funds to the L1
    /// @dev Users need to wait for the rollup's dispute period before triggering their withdrawal in the L1
    /// @return unique withdrawal identifier needed to execute the withdrawal in the L1
    function withdraw(
        address l1ERC20,
        address destination,
        uint256 amount
    ) external returns (uint256);

    /// @notice An address oracle that provides users with the L2 address of an L1 token
    /// @return address of L2 token that was created using this bridge
    function calculateL2TokenAddress(address l1ERC20) external view returns (address);
}
