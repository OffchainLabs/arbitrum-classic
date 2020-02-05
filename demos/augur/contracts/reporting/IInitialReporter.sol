pragma solidity 0.5.15;

import './IReportingParticipant.sol';
import './IMarket.sol';
import '../libraries/IOwnable.sol';
import '../IAugur.sol';


contract IInitialReporter is IReportingParticipant, IOwnable {
    function initialize(IAugur _augur, IMarket _market, address _designatedReporter) public;
    function report(address _reporter, bytes32 _payoutDistributionHash, uint256[] memory _payoutNumerators, uint256 _initialReportStake) public;
    function designatedReporterShowed() public view returns (bool);
    function initialReporterWasCorrect() public view returns (bool);
    function getDesignatedReporter() public view returns (address);
    function getReportTimestamp() public view returns (uint256);
    function migrateToNewUniverse(address _designatedReporter) public;
    function returnRepFromDisavow() public;
}
