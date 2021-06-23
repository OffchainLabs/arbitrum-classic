// SPDX-License-Identifier: MIT
pragma solidity >0.6.0 <0.8.0;

import "../arbitrum/IArbToken.sol";
import "./IWETH9.sol";

interface IaeWETH is IArbToken, IWETH9 {
    function transferToGateway(address _from, uint256 amount) external;
}
