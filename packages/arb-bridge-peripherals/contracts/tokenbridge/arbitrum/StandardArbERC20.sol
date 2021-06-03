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
import "arb-bridge-eth/contracts/libraries/Cloneable.sol";
import "./IArbStandardToken.sol";
import "../libraries/BytesParser.sol";
import "./router/L2Gateway.sol";

/**
 * @title Standard (i.e., non-custom) contract deployed by ArbTokenBridge.sol as L2 ERC20. Includes standard ERC20 interface plus additional methods for deposits/withdraws
 */
contract StandardArbERC20 is aeERC20, Cloneable, IArbStandardToken {
    address public gatewayAddress;
    address public override l1Address;

    modifier onlyGateway {
        require(msg.sender == address(gatewayAddress), "ONLY_GATEWAY");
        _;
    }

    /**
     * @notice initialize the token
     * @dev the L2 bridge assumes this does not fail or revert
     * @param _l1Address L1 address of ERC20
     * @param _data encoded symbol/name/decimal data for initial deploy
     */
    function bridgeInit(address _l1Address, bytes memory _data) external override {
        require(address(l1Address) == address(0), "Already inited");
        gatewayAddress = msg.sender;
        l1Address = _l1Address;

        (bytes memory name, bytes memory symbol, bytes memory decimals) =
            abi.decode(_data, (bytes, bytes, bytes));
        // what if decode reverts? shouldn't as this is encoded by L1 contract

        aeERC20.initialize(
            BytesParserWithDefault.toString(name, ""),
            BytesParserWithDefault.toString(symbol, ""),
            BytesParserWithDefault.toUint8(decimals, 18)
        );
    }

    /**
     * @notice Mint tokens on L2. Callable path is EthErc20Bridge.depositToken (which handles L1 escrow), which triggers ArbTokenBridge.mintFromL1, which calls this
     * @param account recipient of tokens
     * @param amount amount of tokens minted
     */
    function bridgeMint(address account, uint256 amount) external override onlyGateway {
        _mint(account, amount);
    }

    /**
     * @notice Burn tokens on L2.
     * @dev only the token bridge can call this
     * @param account owner of tokens
     * @param amount amount of tokens burnt
     */
    function bridgeBurn(address account, uint256 amount) external override onlyGateway {
        _burn(account, amount);
    }
}
