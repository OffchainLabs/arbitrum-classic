// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.6.11;

import "../libraries/IWETH9.sol";

interface IWETH9L2 is IWETH9 {
    event Approval(address indexed src, address indexed guy, uint256 wad);
    event Transfer(address indexed src, address indexed dst, uint256 wad);
    event Deposit(address indexed dst, uint256 wad);
    event Withdrawal(address indexed src, uint256 wad);

    function balanceOf(address guy) external view returns (uint256);

    function allowance(address guy) external view returns (uint256);

    function totalSupply() external view returns (uint256);

    function approve(address guy, uint256 wad) external returns (bool);

    function transfer(address dst, uint256 wad) external returns (bool);

    function transferFrom(
        address src,
        address dst,
        uint256 wad
    ) external returns (bool);
}
