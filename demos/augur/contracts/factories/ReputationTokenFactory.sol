pragma solidity 0.5.15;


import 'ROOT/IAugur.sol';
import 'ROOT/reporting/IUniverse.sol';
import 'ROOT/reporting/IV2ReputationToken.sol';
import 'ROOT/reporting/ReputationToken.sol';
import 'ROOT/factories/IReputationTokenFactory.sol';


/**
 * @title Reputation Token Factory
 * @notice A Factory contract to create Reputation Token contracts
 * @dev Only meant to be used by the universe corresponding to the token. This creates a normal contract rather than a delegate. As there shouldn't be many REP tokens in existence this will save on gas.
 */
contract ReputationTokenFactory is IReputationTokenFactory {
    function createReputationToken(IAugur _augur, IUniverse _parentUniverse) public returns (IV2ReputationToken) {
        IUniverse _universe = IUniverse(msg.sender);
        return IV2ReputationToken(new ReputationToken(_augur, _universe, _parentUniverse));
    }
}
