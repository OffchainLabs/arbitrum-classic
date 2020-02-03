pragma solidity 0.5.15;

import 'ROOT/reporting/IUniverse.sol';


contract IUniverseFactory {
    function createUniverse(IUniverse _parentUniverse, bytes32 _parentPayoutDistributionHash, uint256[] memory _payoutNumerators) public returns (IUniverse);
}
