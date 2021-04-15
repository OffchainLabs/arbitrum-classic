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

import "@openzeppelin/contracts-upgradeable/utils/Create2Upgradeable.sol";

abstract contract TokenAddressHandler {
    mapping(address => address) public customL2Token;

    function isCustomToken(address l1Token) public view returns (bool) {
        return customL2Token[l1Token] != address(0);
    }

    function getCreate2Salt(address l1Token, address l2TemplateERC20)
        internal
        pure
        returns (bytes32)
    {
        return keccak256(abi.encodePacked(l1Token, l2TemplateERC20));
    }

    function calculateL2ERC20TokenAddress(
        address l1Token,
        address l2TemplateERC20,
        address l2ArbTokenBridgeAddress,
        bytes32 cloneableProxyHash
    ) internal view returns (address) {
        bytes32 salt = getCreate2Salt(l1Token, l2TemplateERC20);
        return Create2Upgradeable.computeAddress(salt, cloneableProxyHash, l2ArbTokenBridgeAddress);
    }

    function calculateL2TokenAddress(
        address l1Token,
        address l2TemplateERC20,
        address l2ArbTokenBridgeAddress,
        bytes32 cloneableProxyHash
    ) internal view returns (address) {
        address customTokenAddress = customL2Token[l1Token];

        if (customTokenAddress != address(0)) {
            return customTokenAddress;
        } else {
            return
                calculateL2ERC20TokenAddress(
                    l1Token,
                    l2TemplateERC20,
                    l2ArbTokenBridgeAddress,
                    cloneableProxyHash
                );
        }
    }
}
