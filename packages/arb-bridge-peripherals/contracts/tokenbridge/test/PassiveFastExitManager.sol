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

import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

import "../ethereum/IEthERC20Bridge.sol";
import "../ethereum/IExitLiquidityProvider.sol";

contract PassiveFastExitManager is IExitTransferCallReceiver {
    address bridge;

    function setBridge(address _bridge) external {
        bridge = _bridge;
    }

    modifier onlyBridge {
        require(msg.sender == bridge, "ONLY_BRIDGE");
        _;
    }

    function onExitTransfered(
        address sender,
        uint256 amount,
        address erc20,
        bytes calldata data
    ) external override onlyBridge returns (bytes4) {
        (
            address initialDestination,
            uint256 exitNum,
            uint256 maxFee,
            address liquidityProvider,
            bytes memory liquidityProof,
            bytes memory spareData
        ) = abi.decode(data, (address, uint256, uint256, address, bytes, bytes));

        {
            uint256 balancePrior = IERC20(erc20).balanceOf(sender);

            // Liquidity provider is responsible for validating if this is a valid exit
            IExitLiquidityProvider(liquidityProvider).requestLiquidity(
                initialDestination,
                erc20,
                amount,
                exitNum,
                liquidityProof
            );

            uint256 balancePost = IERC20(erc20).balanceOf(sender);

            // User must be sent at least (amount - maxFee) or execution reverts
            require(
                SafeMath.sub(balancePost, balancePrior) >= SafeMath.sub(amount, maxFee),
                "User did not get credited with enough tokens"
            );
        }

        IEthERC20Bridge(bridge).transferExitAndCall(
            initialDestination,
            erc20,
            amount,
            exitNum,
            liquidityProvider,
            spareData
        );
        return PassiveFastExitManager.onExitTransfered.selector;
    }
}
