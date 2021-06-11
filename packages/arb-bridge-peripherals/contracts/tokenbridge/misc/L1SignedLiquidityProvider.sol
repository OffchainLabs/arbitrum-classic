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
import "arb-bridge-eth/contracts/libraries/MerkleLib.sol";

import "./ConfirmRoots.sol";
import { IExitLiquidityProvider } from "./L1PassiveFastExitManager.sol";

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/cryptography/ECDSA.sol";

contract L1SignedLiquidityProvider is Ownable, IExitLiquidityProvider {
    uint256 public constant fee_div = 100;
    address tokenBridge;
    address signer;

    constructor(address _tokenBridge, address _signer) public {
        tokenBridge = _tokenBridge;
        signer = _signer;
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
    ) external override {
        require(msg.sender == tokenBridge, "NOT_BRIDGE");
        bytes32 withdrawData = keccak256(abi.encodePacked(exitNum, dest, erc20, amount));
        require(ECDSA.recover(withdrawData, liquidityProof) == signer, "BAD_SIG");
        uint256 fee = amount / fee_div;
        require(IERC20(erc20).transfer(dest, amount - fee), "INSUFFICIENT_LIQUIDITIY");
    }
}
