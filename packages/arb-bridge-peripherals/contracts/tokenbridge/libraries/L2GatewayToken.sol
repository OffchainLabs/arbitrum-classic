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
     * @param _name ERC20 token name
     * @param _symbol ERC20 token symbol
     * @param _decimals ERC20 decimals
     * @param _l2Gateway L2 gateway this token communicates with
     * @param _l1Counterpart L1 address of ERC20
     */
    function initialize(
        string memory _name,
        string memory _symbol,
        uint8 _decimals,
        address _l2Gateway,
        address _l1Counterpart
    ) internal virtual {
        require(_l2Gateway != address(0), "INVALID_GATEWAY");
        require(l2Gateway == address(0), "ALREADY_INIT");
        l2Gateway = _l2Gateway;
        l1Address = _l1Counterpart;

        aeERC20.initialize(_name, _symbol, _decimals);
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
