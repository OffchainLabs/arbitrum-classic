pragma solidity 0.5.15;

import '../libraries/CloneFactory.sol';
import '../IAugur.sol';
import '../reporting/Universe.sol';
import '../ISimpleDex.sol';
import '../IRepExchange.sol';


/**
 * @title RepExchange Factory
 * @notice A Factory contract to create Rep Exchange contracts
 */
contract RepExchangeFactory is CloneFactory {
    function createRepExchange(IAugur _augur, address _repTokenAddress) public returns (ISimpleDex) {
    	address newContractAddress = createNewContract();
        IRepExchange(newContractAddress).initialize(address(_augur), _repTokenAddress);
        return ISimpleDex(newContractAddress);
    }
}
