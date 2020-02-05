pragma solidity 0.5.15;


import './Order.sol';
import '../reporting/IMarket.sol';


contract ITrade {
    function publicTrade(Order.TradeDirections _direction, IMarket _market, uint256 _outcome, uint256 _amount, uint256 _price, bytes32 _betterOrderId, bytes32 _worseOrderId, uint256 _tradeGroupID) external returns (bytes32);
    function publicFillBestOrder(Order.TradeDirections _direction, IMarket _market, uint256 _outcome, uint256 _amount, uint256 _price, uint256 _tradeGroupID) external returns (uint256);
    function trade(address _sender, Order.TradeDirections _direction, IMarket _market, uint256 _outcome, uint256 _amount, uint256 _price, bytes32 _betterOrderId, bytes32 _worseOrderId, uint256 _tradeGroupID) internal returns (bytes32);
    function fillBestOrder(address _sender, Order.TradeDirections _direction, IMarket _market, uint256 _outcome, uint256 _amount, uint256 _price, uint256 _tradeGroupID) internal returns (uint256);
}
