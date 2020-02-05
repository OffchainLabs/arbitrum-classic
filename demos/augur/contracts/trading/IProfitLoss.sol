pragma solidity 0.5.15;


import '../IAugur.sol';
import '../reporting/IMarket.sol';


contract IProfitLoss {
    function initialize(IAugur _augur) public;
    function recordFrozenFundChange(IUniverse _universe, IMarket _market, address _account, uint256 _outcome, int256 _frozenFundDelta) public returns (bool);
    function adjustTraderProfitForFees(IMarket _market, address _trader, uint256 _outcome, uint256 _fees) public returns (bool);
    function recordTrade(IUniverse _universe, IMarket _market, address _longAddress, address _shortAddress, uint256 _outcome, int256 _amount, int256 _price, uint256 _numLongTokens, uint256 _numShortTokens, uint256 _numLongShares, uint256 _numShortShares) public returns (bool);
    function recordClaim(IMarket _market, address _account, uint256[] memory _outcomeFees) public returns (bool);
}
