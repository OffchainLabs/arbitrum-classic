pragma solidity 0.5.15;

import 'ROOT/IAugur.sol';
import 'ROOT/reporting/IDisputeCrowdsourcer.sol';


contract IDisputeCrowdsourcerFactory {
    function createDisputeCrowdsourcer(IAugur _augur, uint256 _size, bytes32 _payoutDistributionHash, uint256[] memory _payoutNumerators, uint256 _crowdsourcerGeneration) public returns (IDisputeCrowdsourcer);
}
