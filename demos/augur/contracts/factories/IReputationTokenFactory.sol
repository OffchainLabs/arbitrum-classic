pragma solidity 0.5.15;

import '../IAugur.sol';
import '../reporting/IUniverse.sol';
import '../reporting/IV2ReputationToken.sol';


contract IReputationTokenFactory {
    function createReputationToken(IAugur _augur, IUniverse _parentUniverse) public returns (IV2ReputationToken);
}
 