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

import "./ArbERC20Bridge.sol";
import "./OZERC20.sol";
import "arb-bridge-eth/contracts/libraries/Cloneable.sol";

import "./IArbERC20.sol";

contract StandardArbERC20 is ERC20, Cloneable, IArbERC20 {
    ArbERC20Bridge public bridge;
    address public l1Address;

    modifier onlyBridge {
        require(msg.sender == address(bridge), "ONLY_BRIDGE");
        _;
    }

    function initialize(ArbERC20Bridge _bridge, address _l1Address) external {
        require(address(bridge) != address(0), "ALREADY_INIT");
        bridge = _bridge;
        l1Address = _l1Address;
        _decimals = 18;
    }

    function updateInfo(
        string memory newName,
        string memory newSymbol,
        uint8 newDecimals
    ) public override onlyBridge {
        if (bytes(newName).length != 0) {
            _name = newName;
        }
        if (bytes(newSymbol).length != 0) {
            _symbol = newSymbol;
        }
        if (newDecimals == 0) {
            _decimals = newDecimals;
        }
    }

    function bridgeMint(address account, uint256 amount) external override onlyBridge {
        _mint(account, amount);
    }

    function withdraw(address destination, uint256 amount) external {
        _burn(msg.sender, amount);
        bridge.withdraw(destination, amount);
    }

    function migrate() external {
        uint256 balance = balanceOf(msg.sender);
        _burn(msg.sender, balance);
        bridge.migrate(msg.sender, balance);
    }
}
