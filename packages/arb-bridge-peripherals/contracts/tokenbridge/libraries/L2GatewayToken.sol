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

import "./aeERC20.sol";
import "./BytesParser.sol";
import "../arbitrum/IArbToken.sol";

/**
 * @title Standard (i.e., non-custom) contract used as a base for different L2 Gateways
 */
abstract contract L2GatewayToken is aeERC20, IArbToken {
    address public l2Gateway;
    address public override l1Address;

    modifier onlyGateway {
        require(msg.sender == l2Gateway, "ONLY_GATEWAY");
        _;
    }

    /**
     * @notice initialize the token
     * @dev the L2 bridge assumes this does not fail or revert
     * @param name_ ERC20 token name
     * @param symbol_ ERC20 token symbol
     * @param decimals_ ERC20 decimals
     * @param l2Gateway_ L2 gateway this token communicates with
     * @param l1Counterpart_ L1 address of ERC20
     */
    function _initialize(
        string memory name_,
        string memory symbol_,
        uint8 decimals_,
        address l2Gateway_,
        address l1Counterpart_
    ) internal virtual {
        require(l2Gateway_ != address(0), "INVALID_GATEWAY");
        require(l2Gateway == address(0), "ALREADY_INIT");
        l2Gateway = l2Gateway_;
        l1Address = l1Counterpart_;

        aeERC20._initialize(name_, symbol_, decimals_);
    }

    /**
     * @notice Mint tokens on L2. Callable path is L1Gateway depositToken (which handles L1 escrow), which triggers L2Gateway, which calls this
     * @param account recipient of tokens
     * @param amount amount of tokens minted
     */
    function bridgeMint(address account, uint256 amount) external virtual override onlyGateway {
        _mint(account, amount);
    }

    /**
     * @notice Burn tokens on L2.
     * @dev only the token bridge can call this
     * @param account owner of tokens
     * @param amount amount of tokens burnt
     */
    function bridgeBurn(address account, uint256 amount) external virtual override onlyGateway {
        _burn(account, amount);
    }
}
