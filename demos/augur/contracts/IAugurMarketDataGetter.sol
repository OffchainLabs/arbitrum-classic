pragma solidity 0.5.15;

import 'ROOT/reporting/IMarket.sol';


contract IAugurMarketDataGetter {
    function getMarketType(IMarket _market) public view returns (IMarket.MarketType _marketType);
    function getMarketOutcomes(IMarket _market) public view returns (bytes32[] memory _outcomes);
}
