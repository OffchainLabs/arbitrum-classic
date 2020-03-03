pragma solidity 0.5.15;


import '../IAugur.sol';
import '../libraries/CloneFactory.sol';
import '../reporting/IUniverse.sol';
import '../reporting/IV2ReputationToken.sol';
import '../reporting/ReputationToken.sol';
import '../factories/IReputationTokenFactory.sol';


/**
 * @title Reputation Token Factory
 * @notice A Factory contract to create Reputation Token contracts
 * @dev Only meant to be used by the universe corresponding to the token. This creates a normal contract rather than a delegate. As there shouldn't be many REP tokens in existence this will save on gas.
 */
contract ReputationTokenFactory is CloneFactory, IReputationTokenFactory {
    function createReputationToken(IAugur _augur, IUniverse _parentUniverse) public returns (IV2ReputationToken) {
    	address newContractAddress = createNewContract();
        ReputationToken _reputationToken = ReputationToken(newContractAddress);

        IUniverse _universe = IUniverse(msg.sender);
        _reputationToken.initializeRepToken(_augur, _universe, _parentUniverse);

        return IV2ReputationToken(_reputationToken);
    }
}
