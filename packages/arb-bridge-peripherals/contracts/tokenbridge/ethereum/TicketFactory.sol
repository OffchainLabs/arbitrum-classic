// SPDX-License-Identifier: MIT
pragma solidity ^0.6.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";

contract TicketFactory is ERC721 {
  constructor() ERC721("Arb Bridge Withdrawal Ticket", "WDRW") public {}

  function createId(
    address bridge, 
    address token, 
    address owner, 
    uint256 exitNum
  ) public pure returns (uint256) {
    return uint256(keccak256(abi.encodePacked(bridge, token, owner, exitNum)));
  }

  function mint(
    address token, 
    address owner, 
    uint256 exitNum, 
    address recipient
  ) external returns (uint256 id) {
    id = createId(msg.sender, token, owner, exitNum);
    _mint(recipient, id);
  }

  function exists(uint256 tokenId) public view returns (bool) {
    return _exists(tokenId);
  }

  function exists(
    address bridge, 
    address token, 
    address owner, 
    uint256 exitNum
  ) public view returns (bool) {
    return _exists(createId(bridge, token, owner, exitNum));
  }

  function burn(
    address token, 
    address owner, 
    uint256 exitNum
  ) external returns (address tokenOwner) {
    uint256 id = createId(msg.sender, token, owner, exitNum);
    tokenOwner = ownerOf(id);
    _burn(id);
  }
}
