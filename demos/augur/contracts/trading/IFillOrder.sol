pragma solidity 0.5.15;


import './Order.sol';


contract IFillOrder {
    function publicFillOrder(bytes32 _orderId, uint256 _amountFillerWants, bytes32 _tradeGroupId, bytes32 _fingerprint) external returns (uint256);
    function fillOrder(address _filler, bytes32 _orderId, uint256 _amountFillerWants, bytes32 tradeGroupId, bytes32 _fingerprint) external returns (uint256);
    function fillZeroXOrder(IMarket _market, uint256 _outcome, uint256 _price, Order.Types _orderType, address _creator, uint256 _amount, bytes32 _fingerprint, bytes32 _tradeGroupId, address _filler) external returns (uint256, uint256);
    function getMarketOutcomeValues(IMarket _market) public view returns (uint256[] memory);
    function getMarketVolume(IMarket _market) public view returns (uint256);
}
