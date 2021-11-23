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

import "@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol";

contract StandardArbERC721 is ERC721Upgradeable {
    address public gateway;
    address public l1Address;
    // TODO: would be cool if we could instead inline these values
    // TODO: do we have risk of overwriting the random slot of proxy admin?
    mapping(uint256 => bytes) private _rawTokenUri;
    bytes private rawName;
    bytes private rawSymbol;

    function bridgeInit(
        address _l1Address,
        bytes calldata _rawName,
        bytes calldata _rawSymbol
    ) external {
        gateway = msg.sender;
        l1Address = _l1Address;
        rawName = _rawName;
        rawSymbol = _rawSymbol;
    }

    modifier onlyGateway() {
        require(msg.sender == gateway, "NOT_GATEWAY");
        _;
    }

    function mint(
        address to,
        uint256 tokenId,
        bytes calldata tokenUri
    ) external onlyGateway {
        _mint(to, tokenId);
    }

    function burn(uint256 tokenId) external onlyGateway {
        _burn(tokenId);
    }

    function tokenURI(uint256 tokenId) public view override returns (string memory) {
        bytes memory data = _rawTokenUri[tokenId];
        // TODO: write a unit test for this
        assembly {
            return(data, mload(data))
        }
    }

    function name() public view override returns (string memory) {
        bytes memory data = rawName;
        assembly {
            return(data, mload(data))
        }
    }

    function symbol() public view override returns (string memory) {
        bytes memory data = rawSymbol;
        assembly {
            return(data, mload(data))
        }
    }
}
