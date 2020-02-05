pragma solidity 0.5.15;


import '../libraries/ITyped.sol';
import './IUniverse.sol';
import './IMarket.sol';
import './IReputationToken.sol';
import '../ICash.sol';
import '../IAugur.sol';
import '../libraries/token/IERC20.sol';


contract IDisputeWindow is ITyped, IERC20 {
    function invalidMarketsTotal() external view returns (uint256);
    function validityBondTotal() external view returns (uint256);

    function incorrectDesignatedReportTotal() external view returns (uint256);
    function initialReportBondTotal() external view returns (uint256);

    function designatedReportNoShowsTotal() external view returns (uint256);
    function designatedReporterNoShowBondTotal() external view returns (uint256);

    function initialize(IAugur _augur, IUniverse _universe, uint256 _disputeWindowId, bool _participationTokensEnabled, uint256 _duration, uint256 _startTime) public;
    function trustedBuy(address _buyer, uint256 _attotokens) public returns (bool);
    function getUniverse() public view returns (IUniverse);
    function getReputationToken() public view returns (IReputationToken);
    function getStartTime() public view returns (uint256);
    function getEndTime() public view returns (uint256);
    function getWindowId() public view returns (uint256);
    function isActive() public view returns (bool);
    function isOver() public view returns (bool);
    function onMarketFinalized() public;
    function redeem(address _account) public returns (bool);
}
