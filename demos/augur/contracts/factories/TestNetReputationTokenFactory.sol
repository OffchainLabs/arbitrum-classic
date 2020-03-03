pragma solidity 0.5.15;

import '../IAugur.sol';
import '../reporting/IUniverse.sol';
import '../reporting/IV2ReputationToken.sol';
import '../TestNetReputationToken.sol';
import '../factories/IReputationTokenFactory.sol';
import '../libraries/CloneFactory.sol';


/**
 * @title TestNet Reputation Token Factory
 * @notice A Factory contract to create TestNet Reputation Token contracts
 * @dev Only meant for use in Testing environments. Only meant to be used by the universe corresponding to the token. This creates a normal contract rather than a delegate. As there shouldn't be many REP tokens in existance this will save on gas.
 */
contract TestNetReputationTokenFactory is CloneFactory, IReputationTokenFactory {
    function createReputationToken(IAugur _augur, IUniverse _parentUniverse) public returns (IV2ReputationToken) {
        address newContractAddress = createNewContract();
        TestNetReputationToken _testNetRepToken = TestNetReputationToken(newContractAddress);

        IUniverse _universe = IUniverse(msg.sender);
        _testNetRepToken.initialize(_augur, _universe, _parentUniverse);
        IV2ReputationToken _reputationToken = IV2ReputationToken(_testNetRepToken);
        
        return _reputationToken;
    }
}
