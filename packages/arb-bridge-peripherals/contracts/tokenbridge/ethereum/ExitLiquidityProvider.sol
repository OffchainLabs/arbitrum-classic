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

import "arb-bridge-eth/contracts/rollup/Rollup.sol";
import "./ConfirmRoots.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract ExitLiquidityProvider is Ownable {
    ConfirmRoots confirmRoots;
    Rollup rollup;
    uint256 public constant fee_div = 100;
    address trustedStaker;

    constructor(address _confirmRoots) public {
        confirmRoots = ConfirmRoots(_confirmRoots);
    }

    function withdrawLiquidity(
        address dest,
        address erc20,
        uint256 amount
    ) external onlyOwner {
        require(IERC20(erc20).transfer(dest, amount), "INSUFFICIENT_LIQUIDITIY");
    }

    function requestLiquidity(
        bytes32 confirmRoot,
        address dest,
        address erc20,
        uint256 amount,
        bytes calldata liquidityProof
    ) external {
        uint256 nodeNum = abi.decode(liquidityProof, (uint256));
        require(confirmRoots.confirmRoots(confirmRoot, nodeNum), "INVALID_ROOT");
        require(rollup.getNode(nodeNum).stakers(trustedStaker), "NOT_TRUSTED");
        uint256 fee = amount / fee_div;
        require(IERC20(erc20).transfer(dest, amount - fee), "INSUFFICIENT_LIQUIDITIY");
    }
}
