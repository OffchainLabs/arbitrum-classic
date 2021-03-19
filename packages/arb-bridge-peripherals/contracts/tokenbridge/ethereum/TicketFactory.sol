// SPDX-License-Identifier: MIT
pragma solidity ^0.6.0;

import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";

contract TicketFactory is ERC1155 {
  constructor() ERC1155("") public {}

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
    _mint(recipient, id, 1, '');
  }

  function holdsTicket(address holder, uint256 tokenId) public view returns (bool) {
    return balanceOf(holder, tokenId) == 1;
  }

  function holdsTicket(
    address holder,
    address bridge, 
    address token, 
    address owner, 
    uint256 exitNum
  ) public view returns (bool) {
    return balanceOf(holder, createId(bridge, token, owner, exitNum)) == 1;
  }

  function burn(
    address token, 
    address owner, 
    uint256 exitNum
  ) external returns (address tokenOwner) {
    uint256 id = createId(msg.sender, token, owner, exitNum);
    tokenOwner = ownerOf(id);
    _burn(id, 1);
  }
}
