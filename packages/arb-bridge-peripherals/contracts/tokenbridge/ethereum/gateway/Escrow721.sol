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

contract Escrow721 is IERC721Receiver {
    address owner;

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
