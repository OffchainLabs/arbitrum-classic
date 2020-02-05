pragma solidity 0.5.15;
pragma experimental ABIEncoderV2;

import '../IAugur.sol';
import '../IAugurCreationDataGetter.sol';
import '../reporting/IMarket.sol';
import '../libraries/token/IERC20.sol';
import '../libraries/math/SafeMathUint256.sol';
import '../reporting/IReportingParticipant.sol';
import '../reporting/IDisputeWindow.sol';
import '../trading/IFillOrder.sol';
import '../trading/IOrders.sol';


/**
 * @title Hot Loading
 * @notice A Utility contract for pulling market data in a single call intended for fast rendering of market presentation before full platform data can be accessed
 */
contract HotLoading {
    using SafeMathUint256 for uint256;

    enum ReportingState {
        PreReporting,
        DesignatedReporting,
        OpenReporting,
        CrowdsourcingDispute,
        AwaitingNextWindow,
        AwaitingFinalization,
        Finalized,
        Forking,
        AwaitingForkMigration
    }

    struct DisputeWindowData {
        address disputeWindow;
        uint256 startTime;
        uint256 endTime;
        uint256 purchased;
        uint256 fees;
    }

    struct MarketData {
        string extraInfo;
        address marketCreator;
        address owner;
        bytes32[] outcomes;
        IMarket.MarketType marketType;
        int256[] displayPrices;
        address designatedReporter;
        ReportingState reportingState;
        uint256 disputeRound;
        uint256[] winningPayout;
        uint256 volume;
        uint256 openInterest;
        uint256[] lastTradedPrices;
        address universe;
        uint256 numTicks;
        uint256 feeDivisor;
        uint256 affiliateFeeDivisor;
        uint256 endTime;
        uint256 numOutcomes;
        uint256 validityBond;
        uint256 reportingFeeDivisor;
        uint256[] outcomeVolumes;
    }

    function getMarketData(IAugur _augur, IMarket _market, IFillOrder _fillOrder, IOrders _orders) external view returns (MarketData memory _marketData) {
        IAugurCreationDataGetter.MarketCreationData memory _marketCreationData = IAugurCreationDataGetter(address(_augur)).getMarketCreationData(_market);
        _marketData.extraInfo = _marketCreationData.extraInfo;
        _marketData.marketCreator = _marketCreationData.marketCreator;
        _marketData.outcomes = _marketCreationData.outcomes;
        _marketData.displayPrices = _marketCreationData.displayPrices;
        _marketData.marketType = _marketCreationData.marketType;
        _marketData.owner = _market.getOwner();
        _marketData.validityBond = _market.getValidityBondAttoCash();
        _marketData.numOutcomes = _market.getNumberOfOutcomes();
        _marketData.endTime = _market.getEndTime();
        _marketData.affiliateFeeDivisor = _market.affiliateFeeDivisor();
        _marketData.feeDivisor = _market.getMarketCreatorSettlementFeeDivisor();
        _marketData.numTicks = _market.getNumTicks();
        _marketData.universe = address(_market.getUniverse());
        _marketData.disputeRound = _market.getNumParticipants() - 1;
        IInitialReporter _initialReporter = _market.getInitialReporter();
        _marketData.designatedReporter = _initialReporter.getDesignatedReporter();
        _marketData.openInterest = _market.getOpenInterest();
        _marketData.volume = _fillOrder.getMarketVolume(_market);
        _marketData.lastTradedPrices = getLastTradedPrices(_market, _marketData.numOutcomes, _orders);
        _marketData.reportingState = getReportingState(_augur, _market);
        _marketData.reportingFeeDivisor = _market.getUniverse().getReportingFeeDivisor();
        _marketData.outcomeVolumes = _fillOrder.getMarketOutcomeValues(_market);
        if (_marketData.reportingState == ReportingState.Finalized) {
            IReportingParticipant _winningReportingParticipant = _market.getWinningReportingParticipant();
            _marketData.winningPayout = _winningReportingParticipant.getPayoutNumerators();
        }
    }

    function getLastTradedPrices(IMarket _market, uint256 _numOutcomes, IOrders _orders) public view returns (uint256[] memory _lastTradedPrices) {
        _lastTradedPrices = new uint256[](_numOutcomes);
        for (uint256 _outcome = 0; _outcome < _numOutcomes; _outcome++) {
            _lastTradedPrices[_outcome] = _orders.getLastOutcomePrice(_market, _outcome);
        }
    }

    function getReportingState(IAugur _augur, IMarket _market) public view returns (ReportingState) {
        if (_market.isFinalized()) {
            return ReportingState.Finalized;
        }

        IUniverse _universe = _market.getUniverse();
        IMarket _forkingMarket = _universe.getForkingMarket();
        if (_forkingMarket != IMarket(0)) {
            if (_forkingMarket == _market) {
                return ReportingState.Forking;
            }
            return ReportingState.AwaitingForkMigration;
        }

        IDisputeWindow _disputeWindow = _market.getDisputeWindow();
        uint256 _curTime = _augur.getTimestamp();

        if (_disputeWindow != IDisputeWindow(0)) {
            if (_curTime > _disputeWindow.getEndTime()) {
                return ReportingState.AwaitingFinalization;
            }
            if (_market.getDisputePacingOn() && _curTime < _disputeWindow.getStartTime()) {
                return ReportingState.AwaitingNextWindow;
            }
            return ReportingState.CrowdsourcingDispute;
        }

        if (_curTime > _market.getDesignatedReportingEndTime()) {
            return ReportingState.OpenReporting;
        }

        if (_curTime > _market.getEndTime()) {
            return ReportingState.DesignatedReporting;
        }

        return ReportingState.PreReporting;
    }

    function getCurrentDisputeWindowData(IAugur _augur, IUniverse _universe) external view returns (DisputeWindowData memory _disputeWindowData) {
        IDisputeWindow _disputeWindow = _universe.getCurrentDisputeWindow(false);
        if (_disputeWindow == IDisputeWindow(0)) {
            (uint256 _startTime, uint256 _duration) = _universe.getDisputeWindowStartTimeAndDuration(_augur.getTimestamp(), false);
            _disputeWindowData.startTime = _startTime;
            _disputeWindowData.endTime = _startTime + _duration;
            return _disputeWindowData;
        }
        _disputeWindowData.disputeWindow = address(_disputeWindow);
        _disputeWindowData.startTime = _disputeWindow.getStartTime();
        _disputeWindowData.endTime = _disputeWindow.getEndTime();
        _disputeWindowData.purchased = _disputeWindow.totalSupply();
        _disputeWindowData.fees = IERC20(_augur.lookup('Cash')).balanceOf(address(_disputeWindow));
    }

    function getTotalValidityBonds(IMarket[] calldata _markets) external view returns (uint256 _totalValidityBonds) {
        for (uint256 _i = 0; _i < _markets.length; _i++) {
            IMarket _market = _markets[_i];
            _totalValidityBonds += _market.getValidityBondAttoCash();
        }
    }
}