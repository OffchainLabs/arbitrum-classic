// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.6.11;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

interface IWETH9 is IERC20 {
    function deposit() external payable;

    function withdraw(uint256 _amount) external;
}
