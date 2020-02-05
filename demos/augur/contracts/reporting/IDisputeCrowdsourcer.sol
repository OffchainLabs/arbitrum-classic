pragma solidity 0.5.15;

import './IReportingParticipant.sol';
import './IDisputeWindow.sol';
import '../libraries/token/IERC20.sol';
import '../IAugur.sol';


contract IDisputeCrowdsourcer is IReportingParticipant, IERC20 {
    function initialize(IAugur _augur, IMarket market, uint256 _size, bytes32 _payoutDistributionHash, uint256[] memory _payoutNumerators, uint256 _crowdsourcerGeneration) public;
    function contribute(address _participant, uint256 _amount, bool _overload) public returns (uint256);
    function setSize(uint256 _size) public;
    function getRemainingToFill() public view returns (uint256);
    function correctSize() public returns (bool);
    function getCrowdsourcerGeneration() public view returns (uint256);
}
