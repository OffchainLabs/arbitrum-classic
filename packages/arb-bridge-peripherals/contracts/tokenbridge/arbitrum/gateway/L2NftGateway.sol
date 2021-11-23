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

import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";
import "arb-bridge-eth/contracts/libraries/AddressAliasHelper.sol";
import "../../ethereum/gateway/L1NftGateway.sol";
import "../L2ArbitrumMessenger.sol";

/**
 * @title Common interface for L1 and L2 Gateway Routers
 */
contract L2NftGateway is L2ArbitrumMessenger, IERC721Receiver {
    address public counterpartGateway;
    address public router;

    function initialize(address _counterpartGateway, address _router) public {
        require(counterpartGateway == address(0), "ALREADY_INIT");
        require(_counterpartGateway != address(0), "BAD_COUNTERPART");
        require(_router != address(0), "BAD_ROUTER");
        counterpartGateway = _counterpartGateway;
        router = _router;
    }

    function finalizeDeposit(
        address token,
        uint256 tokenId,
        address from,
        address to,
        bytes calldata data
    ) external {
        require(msg.sender == AddressAliasHelper.applyL1ToL2Alias(counterpartGateway));
        // TODO: check if token deployed
    }

    function onERC721Received(
        address operator,
        address from,
        uint256 tokenId,
        bytes calldata data
    ) external override returns (bytes4) {
        // TODO: trigger a withdrawal
        revert("NOT_IMPLEMENTED");
        // return this.onERC721Received.selector;
    }

    /// @dev this is useful if you want to claim airdrops
    function newEscrowAddress(address token, uint256 tokenId) internal returns (address) {
        // TODO: create2 for a separate address that holds escrow
        // TODO: can we allow this to happen after the initial deposit? if someone sponsors the gas, I don't see why not
        // TODO: can this new escrow thing be called by anyone to be used?
        // TODO: what if user transfers straight to newEscrowAddress
        // TODO: can we trigger auto deposit when airdrops to this addr? depends on base fee
        return getEscrowAddress(token, tokenId);
    }

    function getEscrowAddress(address token, uint256 tokenId) public view returns (address) {
        // TODO: implement
        // TODO: should be able to query before this gets deployed
        return address(0);
    }

    function updateBaseUriToL2() external {
        // TODO: should we also allow update name/symbol?
    }

    function updateTokenUriToL2() external {
        // TODO: take in batch
    }

    function getOutboundCalldata(
        address l1Token,
        uint256 tokenId,
        address from,
        address to,
        bytes calldata data
    ) public view virtual returns (bytes memory) {
        // TODO: query the 721 for "name" / "symbol" / "uri"
        return
            abi.encodeWithSelector(
                // TODO: use correct selector
                L1NftGateway.deposit.selector,
                l1Token,
                tokenId,
                from,
                to,
                data
            );
    }

    function calculateL2TokenAddress(address l1ERC721) public view virtual returns (address) {
        // TODO: implement address oracle
        revert("NOT_IMPLEMENTED");
    }
}
