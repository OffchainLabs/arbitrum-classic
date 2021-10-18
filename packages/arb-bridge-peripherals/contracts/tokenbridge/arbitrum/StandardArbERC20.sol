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

import "arb-bridge-eth/contracts/libraries/Cloneable.sol";
import "../libraries/L2GatewayToken.sol";
import "../libraries/BytesParser.sol";
import "./IArbToken.sol";

/**
 * @title Standard (i.e., non-custom) contract deployed by L2Gateway.sol as L2 ERC20. Includes standard ERC20 interface plus additional methods for deposits/withdraws
 */
contract StandardArbERC20 is IArbToken, L2GatewayToken, Cloneable {
    struct ERC20Getters {
        bool ignoreDecimals;
        bool ignoreName;
        bool ignoreSymbol;
    }
    ERC20Getters private availableGetters;

    /**
     * @notice initialize the token
     * @dev the L2 bridge assumes this does not fail or revert
     * @param _l1Address L1 address of ERC20
     * @param _data encoded symbol/name/decimal data for initial deploy
     */
    function bridgeInit(address _l1Address, bytes memory _data) public virtual {
        (bytes memory name_, bytes memory symbol_, bytes memory decimals_) = abi.decode(
            _data,
            (bytes, bytes, bytes)
        );
        // what if decode reverts? shouldn't as this is encoded by L1 contract

        /*
         *  if parsing fails, the type's default value gets assigned
         *  the parsing can fail for different reasons:
         *      1. method not available in L1 (empty input)
         *      2. data type is encoded differently in the L1 (trying to abi decode the wrong data type)
         *  currently (1) returns a parser fails and (2) reverts as there is no `abi.tryDecode`
         *  https://github.com/ethereum/solidity/issues/10381
         */

        (bool parseNameSuccess, string memory parsedName) = BytesParser.toString(name_);
        (bool parseSymbolSuccess, string memory parsedSymbol) = BytesParser.toString(symbol_);
        (bool parseDecimalSuccess, uint8 parsedDecimals) = BytesParser.toUint8(decimals_);

        L2GatewayToken._initialize(
            parsedName,
            parsedSymbol,
            parsedDecimals,
            msg.sender, // _l2Gateway,
            _l1Address // _l1Counterpart
        );

        // here we assume that (2) would have reverted, so if the parser failed its because the getter isn't available in the L1.
        // instead of storing on a struct, we could instead set a magic number, at something like `type(uint8).max` or random string
        // to be more general we instead use an extra storage slot
        availableGetters = ERC20Getters({
            ignoreName: !parseNameSuccess,
            ignoreSymbol: !parseSymbolSuccess,
            ignoreDecimals: !parseDecimalSuccess
        });
    }

    function decimals() public view override returns (uint8) {
        // no revert message just as in the L1 if you called and the function is not implemented
        if (availableGetters.ignoreDecimals) revert();
        return super.decimals();
    }

    function name() public view override returns (string memory) {
        // no revert message just as in the L1 if you called and the function is not implemented
        if (availableGetters.ignoreName) revert();
        return super.name();
    }

    function symbol() public view override returns (string memory) {
        // no revert message just as in the L1 if you called and the function is not implemented
        if (availableGetters.ignoreSymbol) revert();
        return super.symbol();
    }
}
