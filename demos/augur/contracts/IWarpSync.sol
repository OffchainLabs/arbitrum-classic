pragma solidity 0.5.15;

import './reporting/IUniverse.sol';
import './reporting/IMarket.sol';


contract IWarpSync {
    function markets(address _universe) external returns (IMarket);
    function notifyMarketFinalized() public;
}