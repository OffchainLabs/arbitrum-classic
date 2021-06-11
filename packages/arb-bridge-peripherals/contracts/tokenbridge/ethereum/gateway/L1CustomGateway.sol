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

import "./L1ArbitrumExtendedGateway.sol";
import "../../arbitrum/gateway/L2CustomGateway.sol";
import "@openzeppelin/contracts/utils/Address.sol";

/**
 * @title Gatway for "custom" bridging functionality
 * @notice Handles some (but not all!) custom Gateway needs.
 */
contract L1CustomGateway is L1ArbitrumExtendedGateway {
    using Address for address;
    // stores addresses of L2 tokens to be used
    mapping(address => address) public l1ToL2Token;

    event TokenSet(address indexed l1Address, address indexed l2Address);

    function initialize(
        address _l1Counterpart,
        address _l1Router,
        address _inbox
    ) public virtual {
        L1ArbitrumExtendedGateway._initialize(_l1Counterpart, _l1Router, _inbox);
    }

    /**
     * @notice Calculate the address used when bridging an ERC20 token
     * @dev this always returns the same as the L1 oracle, but may be out of date.
     * For example, a custom token may have been registered but not deploy or the contract self destructed.
     * @param l1ERC20 address of L1 token
     * @return L2 address of a bridged ERC20 token
     */
    function _calculateL2TokenAddress(address l1ERC20)
        internal
        view
        virtual
        override
        returns (address)
    {
        return l1ToL2Token[l1ERC20];
    }

    /**
     * @notice Calculate the address used when bridging an ERC20 token
     * @dev this always returns the same as the L1 oracle, but may be out of date.
     * For example, a custom token may have been registered but not deploy or the contract self destructed.
     * @param l1ERC20 address of L1 token
     * @return L2 address of a bridged ERC20 token
     */
    function calculateL2TokenAddress(address l1ERC20)
        external
        view
        virtual
        override
        onlyRouter
        returns (address)
    {
        // will revert if not called by router
        return _calculateL2TokenAddress(l1ERC20);
    }

    /**
     * @notice Allows L1 Token contract to trustlessly register its custom L2 counterpart.

     * @param l2Address address of L1 token
     * @param _maxGas max gas for L2 retryable exrecution 
     * @param _gasPriceBid gas price for L2 retryable ticket 
     * @param  _maxSubmissionCost base submission cost  L2 retryable tick3et 
     * @return Retryable ticket ID
     */
    function registerTokenToL2(
        address l2Address,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        uint256 _maxSubmissionCost
    ) external virtual returns (uint256) {
        require(address(msg.sender).isContract(), "MUST_BE_CONTRACT");
        l1ToL2Token[msg.sender] = l2Address;

        bytes memory _data =
            abi.encodeWithSelector(
                L2CustomGateway.registerTokenFromL1.selector,
                msg.sender,
                l2Address
            );

        return sendTxToL2(msg.sender, 0, _maxSubmissionCost, _maxGas, _gasPriceBid, _data);
    }
}
