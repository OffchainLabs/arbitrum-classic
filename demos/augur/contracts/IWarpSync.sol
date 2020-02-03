pragma solidity 0.5.15;

import 'ROOT/reporting/IUniverse.sol';
import 'ROOT/reporting/IMarket.sol';


contract IWarpSync {
    function markets(address _universe) external returns (IMarket);
    function notifyMarketFinalized() public;
}