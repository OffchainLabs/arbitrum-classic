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

import "./StandardArbERC20.sol";
import "arb-bridge-eth/contracts/libraries/CloneFactory.sol";
import "arb-bridge-eth/contracts/libraries/ICloneable.sol";
import "arbos-contracts/contracts/ArbSys.sol";

contract ArbERC20Bridge is CloneFactory {
    /// @notice This mapping is from L1 address to L2 address
    mapping(address => address) public l1ToL2;

    /// @notice This mapping is from L1 address to L2 address
    mapping(address => address) public l2ToL1;

    uint256 exitNum;

    ICloneable templateERC20;

    constructor() public {
        templateERC20 = new StandardArbERC20();
    }

    function mintFromL1(
        address erc20,
        address account,
        uint256 amount
    ) external {
        // This ensures that this method can only be called from the L1 pair of this contract
        require(tx.origin == address(this));
        StandardArbERC20 token = ensureTokenExists(erc20);
        token.mint(account, amount);
    }

    function updateTokenInfo(
        address erc20,
        string calldata name,
        string calldata symbol,
        uint8 decimals
    ) external payable {
        require(tx.origin == address(this));
        StandardArbERC20 token = ensureTokenExists(erc20);
        token.updateInfo(name, symbol, decimals);
    }

    function withdraw(address destination, uint256 amount) external {
        address l1Contract = l2ToL1[msg.sender];
        require(l1Contract != address(0), "NOT_FROM_TOKEN");
        ArbSys(100).sendTxToL1(
            address(this),
            abi.encodeWithSignature(
                "withdrawFromL2(uint256,address,address,uint256)",
                exitNum,
                l1Contract,
                destination,
                amount
            )
        );
        exitNum++;
    }

    function ensureTokenExists(address l1ERC20) private returns (StandardArbERC20) {
        address l2Contract = l1ToL2[l1ERC20];
        if (l2Contract == address(0)) {
            l2Contract = createClone(templateERC20);
            l1ToL2[l1ERC20] = l2Contract;
            l2ToL1[l2Contract] = l1ERC20;
            StandardArbERC20(l2Contract).initialize(this, l1ERC20);
        }
        return StandardArbERC20(l2Contract);
    }
}
