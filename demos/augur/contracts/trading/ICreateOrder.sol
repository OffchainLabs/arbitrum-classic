pragma solidity 0.5.15;


import './Order.sol';
import '../reporting/IMarket.sol';


contract ICreateOrder {
    function publicCreateOrder(Order.Types, uint256, uint256, IMarket, uint256, bytes32, bytes32, bytes32) external returns (bytes32);
    function createOrder(address, Order.Types, uint256, uint256, IMarket, uint256, bytes32, bytes32, bytes32) external returns (bytes32);
}
