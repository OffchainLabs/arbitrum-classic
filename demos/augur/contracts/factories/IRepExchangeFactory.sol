pragma solidity 0.5.15;

import '../IAugur.sol';
import '../ISimpleDex.sol';


contract IRepExchangeFactory {
    function createRepExchange(IAugur _augur, address _repTokenAddress) public returns (ISimpleDex);
}
