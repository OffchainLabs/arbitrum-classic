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

import "../libraries/aeERC20.sol";
import "@openzeppelin/contracts-upgradeable/drafts/ERC20PermitUpgradeable.sol";

contract TestERC20 is aeERC20 {
    constructor() public {
        aeERC20._initialize("IntArbTestToken", "IARB", uint8(18));
    }

    function mint() external {
        _mint(msg.sender, 50000000);
    }
}

// test token code inspired from maker
contract Bytes32ERC20 {
    mapping(address => uint256) public balanceOf;
    mapping(address => mapping(address => uint256)) public allowance;

    function transfer(address dst, uint256 wad) public returns (bool) {
        return transferFrom(msg.sender, dst, wad);
    }

    function transferFrom(
        address src,
        address dst,
        uint256 wad
    ) public returns (bool) {
        if (src != msg.sender) {
            allowance[src][msg.sender] = allowance[src][msg.sender] - wad;
        }

        balanceOf[src] = balanceOf[src] - wad;
        balanceOf[dst] = balanceOf[dst] + wad;

        return true;
    }

    function approve(address guy, uint256 wad) public returns (bool) {
        allowance[msg.sender][guy] = wad;
        return true;
    }

    function mint() public {
        balanceOf[msg.sender] += 1 ether;
    }
}

contract Bytes32ERC20WithMetadata is Bytes32ERC20 {
    bytes32 public name = 0x4d616b6572000000000000000000000000000000000000000000000000000000;
    bytes32 public symbol = 0x4d4b520000000000000000000000000000000000000000000000000000000000;
    // TODO: what if this overflows?
    uint8 public decimals = 18;

    // no totalSupply field
    // uint256 public totalSupply;
}
