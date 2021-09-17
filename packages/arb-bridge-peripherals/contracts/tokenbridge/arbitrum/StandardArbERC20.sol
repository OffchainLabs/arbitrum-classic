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

        L2GatewayToken._initialize(
            BytesParserWithDefault.toString(name_, ""),
            BytesParserWithDefault.toString(symbol_, ""),
            BytesParserWithDefault.toUint8(decimals_, 18),
            msg.sender, // _l2Gateway,
            _l1Address // _l1Counterpart
        );
    }

    // values generated with https://hardhat.org/plugins/hardhat-storage-layout.html
    // by running `yarn hardhat storage-slots`
    uint256 constant INITIALIZED_STORAGE_SLOT = 0;

    function getInitialized() internal view returns (bool initialized) {
        assembly {
            initialized := sload(INITIALIZED_STORAGE_SLOT)
        }
    }

    function resetInitialized() internal {
        bool FALSE = false;
        assembly {
            sstore(INITIALIZED_STORAGE_SLOT, FALSE)
        }
    }

    function isEqualString(string memory a, string memory b) internal pure returns (bool) {
        return keccak256(abi.encode(a)) == keccak256(abi.encode(b));
    }

    /// @notice this is a one time use function intended to fix the name/symbol of the maker token
    function updateInfo() external {
        // this can only be triggered for the maker token at 0x2e9a6Df78E42a30712c10a9Dc4b1C8656f8F2879
        require(address(this) == 0x2e9a6Df78E42a30712c10a9Dc4b1C8656f8F2879, "NOT_MKR_TOKEN");

        string
            memory expectedOldName = "0x4d616b6572000000000000000000000000000000000000000000000000000000";
        string
            memory expectedOldSymbol = "0x4d4b520000000000000000000000000000000000000000000000000000000000";

        // validate info wasn't already updated
        require(isEqualString(expectedOldName, this.name()), "NAME_ALREADY_UPDATE");
        require(isEqualString(expectedOldSymbol, this.symbol()), "SYMBOL_ALREADY_UPDATE");

        string memory newExpectedName = "Maker";
        string memory newExpectedSymbol = "MKR";

        require(getInitialized(), "NOT_INITIALIZED");
        resetInitialized();
        require(!getInitialized(), "RESET_INITIALIZED_FAIL");

        L2GatewayToken._initialize(
            newExpectedName,
            newExpectedSymbol,
            this.decimals(),
            l2Gateway, // _l2Gateway,
            l1Address // _l1Counterpart
        );
        require(getInitialized(), "NOT_INITIALIZED_AFTER");

        // verify new values were correctly set
        require(isEqualString(newExpectedName, this.name()), "NAME_NOT_UPDATED");
        require(isEqualString(newExpectedSymbol, this.symbol()), "SYMBOL_NOT_UPDATED");
    }
}
