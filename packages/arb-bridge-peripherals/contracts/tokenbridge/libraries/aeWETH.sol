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

import "./L2GatewayToken.sol";
import "./IWETH9.sol";

/// @title Arbitrum extended WETH
contract aeWETH is L2GatewayToken, IWETH9 {
    function initialize(
        string memory _name,
        string memory _symbol,
        uint8 _decimals,
        address _l2Gateway,
        address _l1Address
    ) external {
        L2GatewayToken._initialize(_name, _symbol, _decimals, _l2Gateway, _l1Address);
    }

    function bridgeMint(address account, uint256 amount) external virtual override {
        revert("NO_BRIDGE_MINT");
    }

    function bridgeBurn(address account, uint256 amount) external virtual override {
        // can be used to allow bridge burn so users can withdraw in a single tx
        // instead use transferAndCall or permit
        revert("NO_BRIDGE_BURN");
    }

    function deposit() external payable override {
        _deposit(msg.sender);
    }

    function depositTo(address account) external payable {
        _deposit(account);
    }

    function withdraw(uint256 amount) external override {
        _withdraw(msg.sender, amount);
    }

    function withdrawTo(address account, uint256 amount) external {
        _withdraw(account, amount);
    }

    function _withdraw(address account, uint256 amount) internal {
        _burn(msg.sender, amount);
        payable(account).transfer(amount);
    }

    function _deposit(address account) internal {
        _mint(account, msg.value);
    }

    receive() external payable {
        _deposit(msg.sender);
    }
}
