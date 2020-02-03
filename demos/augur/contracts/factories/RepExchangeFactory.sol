pragma solidity 0.5.15;

import 'ROOT/libraries/CloneFactory.sol';
import 'ROOT/IAugur.sol';
import 'ROOT/reporting/Universe.sol';
import 'ROOT/ISimpleDex.sol';
import 'ROOT/IRepExchange.sol';


/**
 * @title RepExchange Factory
 * @notice A Factory contract to create Rep Exchange contracts
 */
contract RepExchangeFactory is CloneFactory {
    function createRepExchange(IAugur _augur, address _repTokenAddress) public returns (ISimpleDex) {
        address _exchange = createClone(_augur.lookup("RepExchange"));
        IRepExchange(_exchange).initialize(address(_augur), _repTokenAddress);
        return ISimpleDex(_exchange);
    }
}
