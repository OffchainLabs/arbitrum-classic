pragma solidity >=0.4.21 <0.6.0;

import "./ArbSys.sol";

contract Test {
  function sendERC20(address tokenAddress, uint256 amount) public {
    ArbSys(address(0x01)).sendERC20(
      msg.sender,
      tokenAddress,
      amount
    );
  }

  function sendERC721(address tokenAddress, uint256 id) public {
    ArbSys(address(0x01)).sendERC721(
      msg.sender,
      tokenAddress,
      id
    );
  }

  function deposit() payable public {}
}
