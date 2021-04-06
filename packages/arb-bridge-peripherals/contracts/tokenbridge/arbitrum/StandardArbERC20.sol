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

// import "./open-zeppelin/OZERC20.sol";
import "../libraries/aeERC20.sol";
import "arb-bridge-eth/contracts/libraries/Cloneable.sol";
import "./IArbToken.sol";
import "./ArbTokenBridge.sol";
import "../libraries/BytesParser.sol";

contract StandardArbERC20 is aeERC20, Cloneable, IArbToken {
    ArbTokenBridge public bridge;
    address public l1Address;

    modifier onlyBridge {
        require(msg.sender == address(bridge), "ONLY_BRIDGE");
        _;
    }

    function bridgeInit(address _l1Address, bytes memory _data) external override returns (bool) {
        require(address(l1Address) == address(0), "Already inited");
        bridge = ArbTokenBridge(msg.sender);
        l1Address = _l1Address;

        (bytes memory name, bytes memory symbol, bytes memory decimals) =
            abi.decode(_data, (bytes, bytes, bytes));
        // what if decode reverts? shouldn't as this is encoded by L1 contract

        aeERC20.initialize(
            BytesParserWithDefault.toString(name, ""),
            BytesParserWithDefault.toString(symbol, ""),
            BytesParserWithDefault.toUint8(decimals, 18)
        );
        return true;
    }

    function bridgeMint(
        address account,
        uint256 amount,
        bytes memory data
    ) external override onlyBridge {
        _mint(account, amount);
    }

    function withdraw(address destination, uint256 amount) external {
        _burn(msg.sender, amount);
        bridge.withdraw(l1Address, destination, amount);
    }

    function migrate(
        uint256 amount,
        address target,
        bytes memory data
    ) external {
        _burn(msg.sender, amount);
        // migrating from 20 to 777, so allow data
        bridge.migrate(l1Address, target, msg.sender, amount, data);
    }
}
