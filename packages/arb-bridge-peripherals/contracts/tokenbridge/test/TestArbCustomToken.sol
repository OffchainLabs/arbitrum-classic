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

import "../arbitrum/IArbToken.sol";
import "../libraries/aeERC20.sol";

contract TestArbCustomToken is aeERC20, IArbToken {
    address public l2Gateway;
    address public override l1Address;

    modifier onlyGateway() {
        require(msg.sender == l2Gateway, "ONLY_l2GATEWAY");
        _;
    }

    constructor(address _l2Gateway, address _l1Address) public {
        l2Gateway = _l2Gateway;
        l1Address = _l1Address;
        aeERC20._initialize("TestCustomToken", "CARB", uint8(18));
    }

    function someWackyCustomStuff() public {}

    function bridgeMint(address account, uint256 amount) external override onlyGateway {
        _mint(account, amount);
    }

    function bridgeBurn(address account, uint256 amount) external override onlyGateway {
        _burn(account, amount);
    }
}
