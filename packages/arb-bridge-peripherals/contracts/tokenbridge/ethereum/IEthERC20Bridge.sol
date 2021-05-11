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

interface IEthERC20Bridge {
    event ActivateCustomToken(uint256 indexed seqNum, address indexed l1Address, address l2Address);

    event DeployToken(uint256 indexed seqNum, address indexed l1Address);

    event WithdrawRedirected(
        address indexed user,
        address indexed liquidityProvider,
        address erc20,
        uint256 amount,
        uint256 indexed exitNum
    );

    event WithdrawExecuted(
        address indexed initialDestination,
        address indexed destination,
        address erc20,
        uint256 amount,
        uint256 indexed exitNum
    );

    event DepositToken(
        address indexed destination,
        address sender,
        uint256 indexed seqNum,
        uint256 value,
        address indexed tokenAddress
    );

    function hasTriedDeploy(address erc20) external view returns (bool);

    function registerCustomL2Token(
        address l2CustomTokenAddress,
        uint256 maxSubmissionCost,
        uint256 maxGas,
        uint256 gasPriceBid,
        address refundAddress
    ) external payable returns (uint256);

    function fastWithdrawalFromL2(
        address liquidityProvider,
        bytes memory liquidityProof,
        address initialDestination,
        address erc20,
        uint256 amount,
        uint256 exitNum,
        uint256 maxFee
    ) external;

    function withdrawFromL2(
        uint256 exitNum,
        address erc20,
        address initialDestination,
        uint256 amount
    ) external;

    function deposit(
        address erc20,
        address destination,
        uint256 amount,
        uint256 maxSubmissionCost,
        uint256 maxGas,
        uint256 gasPriceBid,
        bytes calldata callHookData
    ) external payable returns (uint256, uint256);

    function calculateL2TokenAddress(address erc20) external view returns (address);
}
