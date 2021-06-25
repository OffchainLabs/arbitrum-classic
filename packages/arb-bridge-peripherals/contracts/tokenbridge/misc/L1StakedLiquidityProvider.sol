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
import "@openzeppelin/contracts/access/Ownable.sol";

import "./ConfirmRoots.sol";
import { IExitLiquidityProvider } from "./L1PassiveFastExitManager.sol";

import "arb-bridge-eth/contracts/libraries/MerkleLib.sol";

contract StakedLiquidityProvider is Ownable, IExitLiquidityProvider {
    uint256 internal constant SendType_sendTxToL1 = 0;
    uint256 public constant fee_div = 100;

    address tokenBridge;
    ConfirmRoots confirmRoots;
    Rollup rollup;

    address trustedStaker;

    constructor(
        address _tokenBridge,
        address _confirmRoots,
        address _trustedStaker
    ) public {
        tokenBridge = _tokenBridge;
        confirmRoots = ConfirmRoots(_confirmRoots);
        rollup = confirmRoots.rollup();
        trustedStaker = _trustedStaker;
    }

    function withdrawLiquidity(
        address dest,
        address erc20,
        uint256 amount
    ) external onlyOwner {
        require(IERC20(erc20).transfer(dest, amount), "INSUFFICIENT_LIQUIDITIY");
    }

    function requestLiquidity(
        address dest,
        address erc20,
        uint256 amount,
        uint256 exitNum,
        bytes calldata liquidityProof
    ) external override returns (bytes memory) {
        require(msg.sender == tokenBridge, "NOT_BRIDGE");
        (
            uint256 nodeNum,
            bytes32[] memory withdrawProof,
            uint256 merklePath,
            uint256 l2Block,
            uint256 l2Timestamp
        ) = abi.decode(liquidityProof, (uint256, bytes32[], uint256, uint256, uint256));

        {
            bytes32 userTx = userTxHash(exitNum, dest, erc20, amount, l2Block, l2Timestamp);
            bytes32 confirmRoot =
                MerkleLib.calculateRoot(
                    withdrawProof,
                    merklePath,
                    keccak256(abi.encodePacked(userTx))
                );
            require(confirmRoots.confirmRoots(confirmRoot, nodeNum), "INVALID_ROOT");
            require(rollup.getNode(nodeNum).stakers(trustedStaker), "NOT_TRUSTED");
        }

        uint256 fee = amount / fee_div;
        require(IERC20(erc20).transfer(dest, amount - fee), "INSUFFICIENT_LIQUIDITIY");
        return "";
    }

    function userTxHash(
        uint256 exitNum,
        address dest,
        address erc20,
        uint256 amount,
        uint256 l2Block,
        uint256 l2Timestamp
    ) private view returns (bytes32) {
        bytes memory data =
            abi.encodeWithSignature(
                "withdrawFromL2(uint256,address,address,uint256)",
                exitNum,
                dest,
                erc20,
                amount
            );
        return
            keccak256(
                abi.encodePacked(
                    SendType_sendTxToL1,
                    uint256(uint160(bytes20(msg.sender))),
                    uint256(uint160(bytes20(msg.sender))),
                    l2Block,
                    l2Timestamp,
                    uint256(0),
                    data
                )
            );
    }
}
