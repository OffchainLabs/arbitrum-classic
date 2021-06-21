// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.6.11;

interface IWETH9 {
    function deposit() external payable;

    function withdraw(uint256 _amount) external;
}
