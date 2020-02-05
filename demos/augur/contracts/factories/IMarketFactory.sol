pragma solidity 0.5.15;

import '../IAugur.sol';
import '../reporting/IUniverse.sol';
import '../reporting/IMarket.sol';


contract IMarketFactory {
    function createMarket(IAugur _augur, uint256 _endTime, uint256 _feePerCashInAttoCash, IAffiliateValidator _affiliateValidator, uint256 _affiliateFeeDivisor, address _designatedReporterAddress, address _sender, uint256 _numOutcomes, uint256 _numTicks) public returns (IMarket _market);
}
