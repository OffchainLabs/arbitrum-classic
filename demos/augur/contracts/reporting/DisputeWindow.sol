// Copyright (C) 2015 Forecast Foundation OU, full GPL notice in LICENSE

pragma solidity 0.5.15;


import './IDisputeWindow.sol';
import '../libraries/Initializable.sol';
import './IUniverse.sol';
import './IReputationToken.sol';
import './IMarket.sol';
import '../ICash.sol';
import '../factories/MarketFactory.sol';
import '../libraries/math/SafeMathUint256.sol';
import './IDisputeWindow.sol';
import '../libraries/token/VariableSupplyToken.sol';
import '../IAugur.sol';
import '../CashSender.sol';


/**
 * @title Dispute Window
 * @notice A contract used to encapsulate a window of time in which markets can be disputed as well as the pot where reporting fees are collected and distributed.
 */
contract DisputeWindow is Initializable, VariableSupplyToken, IDisputeWindow, CashSender {
    using SafeMathUint256 for uint256;

    uint256 public invalidMarketsTotal;
    uint256 public validityBondTotal;

    uint256 public incorrectDesignatedReportTotal;
    uint256 public initialReportBondTotal;

    uint256 public designatedReportNoShowsTotal;
    uint256 public designatedReporterNoShowBondTotal;

    IAugur public augur;
    IUniverse public universe;
    ICash public cash;
    address public buyParticipationTokens;
    uint256 private startTime;
    bool public participationTokensEnabled;

    uint256 public windowId;
    uint256 public duration;

    function initialize(IAugur _augur, IUniverse _universe, uint256 _disputeWindowId, bool _participationTokensEnabled, uint256 _duration, uint256 _startTime) public beforeInitialized {
        endInitialization();
        augur = _augur;
        universe = _universe;
        duration = _duration;
        windowId = _disputeWindowId;
        cash = ICash(_augur.lookup("Cash"));
        buyParticipationTokens = _augur.lookup("BuyParticipationTokens");
        require(cash != ICash(0));
        require(buyParticipationTokens != address(0));
        startTime = _startTime;
        participationTokensEnabled = _participationTokensEnabled;

        initializeCashSender(_augur.lookup("DaiVat"), address(cash));
    }

    function onMarketFinalized() public {
        IMarket _market = IMarket(msg.sender);
        require(universe.isContainerForMarket(_market));

        uint256 _currentValidityBond = universe.getOrCacheValidityBond();
        uint256 _currentInitialReportBond = universe.getOrCacheDesignatedReportStake();
        uint256 _currentNoShowBond = universe.getOrCacheDesignatedReportNoShowBond();

        uint256 _validityBond = _market.getValidityBondAttoCash();
        uint256 _repBond = _market.getInitialReporter().getSize();

        IInitialReporter _initialReporter = _market.getInitialReporter();

        if (_validityBond >= _currentValidityBond / 2) {
            validityBondTotal = validityBondTotal.add(_validityBond);
            if (_market.isFinalizedAsInvalid()) {
                invalidMarketsTotal = invalidMarketsTotal.add(_validityBond);
            }
        }

        if (_repBond >= _currentInitialReportBond / 2) {
            initialReportBondTotal = initialReportBondTotal.add(_repBond);
            if (!_initialReporter.initialReporterWasCorrect()) {
                incorrectDesignatedReportTotal = incorrectDesignatedReportTotal.add(_repBond);
            }
        }

        if (_repBond >= _currentNoShowBond / 2) {
            designatedReporterNoShowBondTotal = designatedReporterNoShowBondTotal.add(_repBond);
            if (!_initialReporter.designatedReporterShowed()) {
                designatedReportNoShowsTotal = designatedReportNoShowsTotal.add(_repBond);
            }
        }
    }

    /**
     * @notice Buy tokens which can be redeemed for reporting fees
     * @param _attotokens The number of tokens to purchase
     * @return bool True
     */
    function buy(uint256 _attotokens) public returns (bool) {
        buyInternal(msg.sender, _attotokens);
        return true;
    }

    function trustedBuy(address _buyer, uint256 _attotokens) public returns (bool) {
        require(msg.sender == buyParticipationTokens);
        buyInternal(_buyer, _attotokens);
        return true;
    }

    function buyInternal(address _buyer, uint256 _attotokens) private {
        require(participationTokensEnabled, "Cannot buy Participation tokens in initial market dispute windows");
        require(_attotokens > 0, "DisputeWindow.buy: amount cannot be 0");
        require(isActive(), "DisputeWindow.buy: window is not active");
        require(!universe.isForking(), "DisputeWindow.buy: universe is forking");
        getReputationToken().trustedDisputeWindowTransfer(_buyer, address(this), _attotokens);
        mint(_buyer, _attotokens);
    }

    /**
     * @notice Redeem tokens for reporting fees
     * @param _account The account to redeem tokens for
     * @return bool True
     */
    function redeem(address _account) public returns (bool) {
        require(isOver() || universe.isForking(), "DisputeWindow.redeem: window is not over");

        uint256 _attoParticipationTokens = balances[_account];

        if (_attoParticipationTokens == 0) {
            return true;
        }

        uint256 _cashBalance = cashBalance(address(this));

        // Burn tokens and send back REP
        uint256 _supply = totalSupply;
        burn(_account, _attoParticipationTokens);
        require(getReputationToken().transfer(_account, _attoParticipationTokens));

        uint256 _feePayoutShare = 0;
        if (_cashBalance != 0) {
            // Pay out fees
            _feePayoutShare = _cashBalance.mul(_attoParticipationTokens).div(_supply);
            cashTransfer(_account, _feePayoutShare);
        }

        augur.logParticipationTokensRedeemed(universe, _account, _attoParticipationTokens, _feePayoutShare);
        return true;
    }

    function getTypeName() public view returns (bytes32) {
        return "DisputeWindow";
    }

    /**
     * @return The universe associated with this window
     */
    function getUniverse() public view returns (IUniverse) {
        return universe;
    }

    /**
     * @return The reputation token associated with this window
     */
    function getReputationToken() public view returns (IReputationToken) {
        return universe.getReputationToken();
    }

    /**
     * @return When the window begins as a uint256 timestamp
     */
    function getStartTime() public view returns (uint256) {
        return startTime;
    }

    /**
     * @return When the window ends as a uint256 timestamp
     */
    function getEndTime() public view returns (uint256) {
        return getStartTime().add(duration);
    }

    function getWindowId() public view returns (uint256) {
        return windowId;
    }

    /**
     * @return Bool indicating if the window has begun and has not yet ended
     */
    function isActive() public view returns (bool) {
        if (augur.getTimestamp() <= getStartTime()) {
            return false;
        }
        if (augur.getTimestamp() >= getEndTime()) {
            return false;
        }
        return true;
    }

    /**
     * @return Bool indicating if the window has ended
     */
    function isOver() public view returns (bool) {
        return augur.getTimestamp() >= getEndTime();
    }

    function onTokenTransfer(address _from, address _to, uint256 _value) internal {
        augur.logParticipationTokensTransferred(universe, _from, _to, _value, balances[_from], balances[_to]);
    }

    function onMint(address _target, uint256 _amount) internal {
        augur.logParticipationTokensMinted(universe, _target, _amount, totalSupply, balances[_target]);
    }

    function onBurn(address _target, uint256 _amount) internal {
        augur.logParticipationTokensBurned(universe, _target, _amount, totalSupply, balances[_target]);
    }
}
