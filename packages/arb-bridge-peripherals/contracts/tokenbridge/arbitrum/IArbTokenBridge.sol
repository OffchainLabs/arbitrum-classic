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
        uint256 id,
        address indexed l1Address,
        uint256 indexed amount,
        address indexed destination,
        uint256 exitNum
    );

    event TokenCreated(
        address indexed l1Address,
        address indexed l2Address,
        StandardTokenType indexed tokenType
    );

    event TokenMinted(
        address l1Address,
        address indexed l2Address,
        StandardTokenType tokenType,
        address indexed sender,
        address indexed dest,
        uint256 amount,
        bool usedCallHook
    );

    event TokenMigrated(
        address indexed from,
        address indexed to,
        address indexed account,
        uint256 amount
    );

    function mintFromL1(
        address l1ERC20,
        address sender,
        StandardTokenType tokenType,
        address dest,
        uint256 amount,
        bytes calldata deployData,
        bytes calldata callHookData
    ) external;

    function withdraw(
        address l1ERC20,
        address destination,
        uint256 amount
    ) external returns (uint256);

    function migrate(
        address l1ERC20,
        address target,
        address account,
        uint256 amount
    ) external;

    function customTokenRegistered(address l1Address, address l2Address) external;

    function calculateL2TokenAddress(address l1ERC20) external view returns (address);
}
