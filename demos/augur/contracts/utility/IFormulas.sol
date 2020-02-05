pragma solidity 0.5.15;

import '../reporting/IDisputeWindow.sol';


contract IFormulas {
    function calculateFloatingValue(uint256 _totalBad, uint256 _total, uint256 _targetDivisor, uint256 _previousValue, uint256 _floor) public pure returns (uint256 _newValue);
    function calculateValidityBond(IDisputeWindow  _previousDisputeWindow, uint256 _previousValidityBondInAttoCash) public view returns (uint256);
    function calculateDesignatedReportStake(IDisputeWindow  _previousDisputeWindow, uint256 _previousDesignatedReportStakeInAttoRep, uint256 _initialReportMinValue) public view returns (uint256);
    function calculateDesignatedReportNoShowBond(IDisputeWindow  _previousDisputeWindow, uint256 _previousDesignatedReportNoShowBondInAttoRep, uint256 _initialReportMinValue) public view returns (uint256);
}