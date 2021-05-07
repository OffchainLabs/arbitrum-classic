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

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "../ethereum/IExitLiquidityProvider.sol";

contract FastExitMock is IExitLiquidityProvider {
    uint256 fee = 0;

    function setFee(uint256 _fee) external {
        fee = _fee;
    }

    function requestLiquidity(
        address dest,
        address erc20,
        uint256 amount,
        uint256 exitNum,
        bytes calldata liquidityProof
    ) external override {
        IERC20(erc20).transfer(dest, amount - fee);
    }
}
