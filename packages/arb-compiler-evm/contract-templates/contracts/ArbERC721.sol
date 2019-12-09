/*
 * Copyright 2019, Offchain Labs, Inc.
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

pragma solidity ^0.5.0;

import "./ArbSys.sol";
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";

contract ArbERC721 is ERC721 {
    function adminMint(address account, uint256 tokenId) public {
        require(msg.sender == address(1));
        _mint(account, tokenId);
    }

    function adminBurn(address account, uint256 tokenId) public {
        require(msg.sender == address(1));
        _burn(account, tokenId);
    }
}
