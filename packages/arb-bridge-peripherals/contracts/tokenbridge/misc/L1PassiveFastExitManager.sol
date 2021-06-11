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

import "../ethereum/gateway/L1ArbitrumExtendedGateway.sol";

interface IExitLiquidityProvider {
    function requestLiquidity(
        address dest,
        address erc20,
        uint256 amount,
        uint256 exitNum,
        bytes calldata liquidityProof
    ) external;
}

contract L1PassiveFastExitManager is ITradeableExitReceiver {
    address bridge;

    function setBridge(address _bridge) external {
        bridge = _bridge;
    }

    modifier onlyBridge {
        require(msg.sender == bridge, "ONLY_BRIDGE");
        _;
    }

    struct ExitDataFrame {
        address initialDestination;
        uint256 maxFee;
        address liquidityProvider;
        uint256 amount;
        address erc20;
        bytes liquidityProof;
        bytes spareData;
    }

    function onExitTransfer(
        address sender,
        uint256 exitNum,
        bytes calldata data
    ) external override onlyBridge returns (bool) {
        ExitDataFrame memory frame;
        {
            (
                frame.initialDestination,
                frame.maxFee,
                frame.liquidityProvider,
                frame.amount,
                frame.erc20,
                frame.liquidityProof,
                frame.spareData
            ) = abi.decode(data, (address, uint256, address, uint256, address, bytes, bytes));
        }

        {
            uint256 balancePrior;
            {
                balancePrior = IERC20(frame.erc20).balanceOf(sender);
            }

            // Liquidity provider is responsible for validating if this is a valid exit
            IExitLiquidityProvider(frame.liquidityProvider).requestLiquidity(
                frame.initialDestination,
                frame.erc20,
                frame.amount,
                exitNum,
                frame.liquidityProof
            );

            uint256 balancePost = IERC20(frame.erc20).balanceOf(sender);

            // User must be sent at least (amount - maxFee) or execution reverts
            require(
                SafeMath.sub(balancePost, balancePrior) >= SafeMath.sub(frame.amount, frame.maxFee),
                "User did not get credited with enough tokens"
            );
        }

        L1ArbitrumExtendedGateway(bridge).transferExitAndCall(
            exitNum,
            frame.initialDestination,
            frame.liquidityProvider,
            frame.spareData
        );
        return true;
    }
}
