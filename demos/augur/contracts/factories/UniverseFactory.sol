pragma solidity 0.5.15;

import 'ROOT/IAugur.sol';
import 'ROOT/reporting/Universe.sol';


/**
 * @title Universe Factory
 * @notice A Factory contract to create Universe contracts
 * @dev Only meant to be used by the Augur contract. This creates a normal contract rather than a delegate. As there shouldn't be many Universes in existance this will save on gas.
 */
contract UniverseFactory {
    function createUniverse(IUniverse _parentUniverse, bytes32 _parentPayoutDistributionHash, uint256[] memory _payoutNumerators) public returns (IUniverse) {
        IAugur _augur = IAugur(msg.sender);
        return IUniverse(new Universe(_augur, _parentUniverse, _parentPayoutDistributionHash, _payoutNumerators));
    }
}
