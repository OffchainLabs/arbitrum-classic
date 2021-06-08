// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.6.11;

import "../libraries/IWETH9.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract TestWETH9 is ERC20, IWETH9 {
    constructor(string memory name_, string memory symbol_) public ERC20(name_, symbol_) {}

    function deposit() external payable override {
        _mint(msg.sender, msg.value);
    }

    function withdraw(uint256 _amount) external override {
        _burn(msg.sender, _amount);
        payable(address(msg.sender)).transfer(_amount);
    }
}
