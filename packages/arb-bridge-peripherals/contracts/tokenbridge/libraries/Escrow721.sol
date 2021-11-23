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
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/utils/Create2.sol";

library Escrow721Handler {
    // TODO: be careful, these break through upgrades
    bytes constant bytecode = type(Escrow721).creationCode;
    bytes32 constant bytecodeHash = keccak256(type(Escrow721).creationCode);

    function getCreate2EscrowAddress(address l1Token, uint256 tokenId)
        internal
        view
        returns (address)
    {
        bytes32 salt = getSalt(l1Token, tokenId);
        // TODO: this address oracle breaks with upgrades as bytecodeHash
        // is calculated during compile time
        return Create2.computeAddress(salt, bytecodeHash);
    }

    function create2Deploy(address l1Token, uint256 tokenId) internal returns (address) {
        // "The pair (contract address, uint256 tokenId) [...] globally unique"
        // ~ https://eips.ethereum.org/EIPS/eip-721#rationale
        bytes32 salt = getSalt(l1Token, tokenId);
        return Create2.deploy(0, salt, bytecode);
    }

    function getSalt(address l1Token, uint256 tokenId) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(l1Token, tokenId));
    }
}

contract Escrow721 is IERC721Receiver {
    address public owner;

    constructor() public {
        owner = msg.sender;
    }

    function onERC721Received(
        address operator,
        address from,
        uint256 tokenId,
        bytes calldata data
    ) external override returns (bytes4) {
        // this won't get triggered as this contract will be `code == 0`
        // during deployment tx
        return this.onERC721Received.selector;
    }

    function bridgeAirdrop() external {
        // TODO: allow bridging of tokens airdropped to escrow
        // this will be claimable by the current L2 owner of the
        // Nft `(contract address, uint256 tokenId)` that is encoded into
        // this escrow's create2 salt
        revert("NOT_IMPLEMENTED");
    }

    function requestEscrow(
        address user,
        address tokenAddr,
        uint256 tokenId
    ) external returns (bool) {
        require(msg.sender == owner, "NOT_OWNER");
        require(IERC165(tokenAddr).supportsInterface(0x80ac58cd), "165_INTERFACE_NOT_DETECTED");
        // TODO: should we check if it was transfered here before getting created?
        IERC721(tokenAddr).safeTransferFrom(user, address(this), tokenId);
    }

    function releaseEscrow(
        address to,
        address tokenAddr,
        uint256 tokenId,
        bytes calldata data
    ) external returns (bool) {
        require(msg.sender == owner, "NOT_OWNER");
        IERC721(tokenAddr).safeTransferFrom(address(this), to, tokenId, data);
        return true;
    }
}
