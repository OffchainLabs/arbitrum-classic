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

import "@openzeppelin/contracts-upgradeable/drafts/ERC20PermitUpgradeable.sol";
import "./TransferAndCallToken.sol";

/// @title Arbitrum extended ERC20
/// @notice The recommended ERC20 implementation for Layer 2 tokens
/// @dev This implements the ERC20 standard with transferAndCall extenstion/affordances
contract aeERC20 is ERC20PermitUpgradeable, TransferAndCallToken {
    using AddressUpgradeable for address;

    constructor() public initializer {
        // this is expected to be used as the logic contract behind a proxy
        // override the constructor if you don't wish to use the initialize method
    }

    function _initialize(
        string memory name_,
        string memory symbol_,
        uint8 decimals_
    ) internal initializer {
        __ERC20Permit_init(name_);
        __ERC20_init(name_, symbol_);
        _setupDecimals(decimals_);
    }
}
