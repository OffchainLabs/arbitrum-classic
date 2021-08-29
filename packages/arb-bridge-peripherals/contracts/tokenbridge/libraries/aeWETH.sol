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
        string memory name_,
        string memory symbol_,
        uint8 decimals_,
        address l2Gateway_,
        address l1Address_
    ) external {
        L2GatewayToken._initialize(name_, symbol_, decimals_, l2Gateway_, l1Address_);
    }

    function bridgeMint(address account, uint256 amount) external virtual override {
        // we want weth to always be fully collaterized
        revert("NO_BRIDGE_MINT");
    }

    function bridgeBurn(address account, uint256 amount) external virtual override onlyGateway {
        _burn(account, amount);
        (bool success, ) = msg.sender.call{ value: amount }("");
        require(success, "FAIL_TRANSFER");
    }

    function deposit() external payable override {
        depositTo(msg.sender);
    }

    function withdraw(uint256 amount) external override {
        withdrawTo(msg.sender, amount);
    }

    function depositTo(address account) public payable {
        _mint(account, msg.value);
    }

    function withdrawTo(address account, uint256 amount) public {
        _burn(msg.sender, amount);
        (bool success, ) = account.call{ value: amount }("");
        require(success, "FAIL_TRANSFER");
    }

    receive() external payable {
        depositTo(msg.sender);
    }
}
