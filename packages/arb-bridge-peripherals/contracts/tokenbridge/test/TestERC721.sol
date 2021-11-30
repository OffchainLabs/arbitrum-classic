// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.6.11;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";

contract TestERC721 is ERC721 {
    constructor(string memory name_, string memory symbol_) public ERC721(name_, symbol_) {}

    function mint(
        address to,
        uint256 tokenId,
        bytes calldata _data
    ) external {
        // TODO: what if we try minting on the L2 and this reverts?
        _safeMint(to, tokenId, _data);
    }
}
