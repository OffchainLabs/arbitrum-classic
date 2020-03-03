pragma solidity 0.5.15;

import '../IAugur.sol';
import '../reporting/Universe.sol';
import '../libraries/CloneFactory.sol';


/**
 * @title Universe Factory
 * @notice A Factory contract to create Universe contracts
 * @dev Only meant to be used by the Augur contract. This creates a normal contract rather than a delegate. As there shouldn't be many Universes in existance this will save on gas.
 */
contract UniverseFactory is CloneFactory {
    function createUniverse(IUniverse _parentUniverse, bytes32 _parentPayoutDistributionHash, uint256[] memory _payoutNumerators) public returns (IUniverse) {
        address newContractAddress = createNewContract();
        Universe _universe = Universe(newContractAddress);

        IAugur _augur = IAugur(msg.sender);
        _universe.initializeUniverse(_augur, _parentUniverse, _parentPayoutDistributionHash, _payoutNumerators);

        return IUniverse(_universe);
    }
}
