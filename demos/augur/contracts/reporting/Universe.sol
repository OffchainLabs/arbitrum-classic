pragma solidity 0.5.15;


import 'ROOT/reporting/IUniverse.sol';
import 'ROOT/factories/IReputationTokenFactory.sol';
import 'ROOT/factories/IDisputeWindowFactory.sol';
import 'ROOT/factories/IMarketFactory.sol';
import 'ROOT/factories/IOICashFactory.sol';
import 'ROOT/reporting/IMarket.sol';
import 'ROOT/reporting/IV2ReputationToken.sol';
import 'ROOT/reporting/IDisputeWindow.sol';
import 'ROOT/reporting/Reporting.sol';
import 'ROOT/libraries/math/SafeMathUint256.sol';
import 'ROOT/ICash.sol';
import 'ROOT/reporting/IOICash.sol';
import 'ROOT/reporting/IAffiliateValidator.sol';
import 'ROOT/external/IDaiVat.sol';
import 'ROOT/external/IDaiPot.sol';
import 'ROOT/external/IDaiJoin.sol';
import 'ROOT/utility/IFormulas.sol';
import 'ROOT/IAugur.sol';
import 'ROOT/CashSender.sol';
import 'ROOT/IRepExchange.sol';
import 'ROOT/factories/IRepExchangeFactory.sol';


/**
 * @title Universe
 * @notice A Universe encapsulates a whole instance of Augur. In the event of a fork in a Universe it will split into child Universes which each represent a different version of the truth with respect to how the forking market should resolve.
 */
contract Universe is IUniverse, CashSender {
    using SafeMathUint256 for uint256;

    uint256 public creationTime;
    mapping(address => uint256) public marketBalance;

    IAugur public augur;
    IUniverse private parentUniverse;
    IFormulas public formulas;
    IShareToken public shareToken;
    bytes32 private parentPayoutDistributionHash;
    uint256[] public payoutNumerators;
    IV2ReputationToken private reputationToken;
    IOICash public openInterestCash;
    IMarket private forkingMarket;
    bytes32 private tentativeWinningChildUniversePayoutDistributionHash;
    uint256 private forkEndTime;
    uint256 private forkReputationGoal;
    uint256 private disputeThresholdForFork;
    uint256 private disputeThresholdForDisputePacing;
    uint256 private initialReportMinValue;
    mapping(uint256 => IDisputeWindow) public disputeWindows;
    mapping(address => bool) private markets;
    mapping(bytes32 => IUniverse) private childUniverses;
    uint256 private openInterestInAttoCash;
    IMarketFactory public marketFactory;
    IDisputeWindowFactory public disputeWindowFactory;

    mapping (address => uint256) public validityBondInAttoCash;
    mapping (address => uint256) public designatedReportStakeInAttoRep;
    mapping (address => uint256) public designatedReportNoShowBondInAttoRep;
    uint256 public previousValidityBondInAttoCash;
    uint256 public previousDesignatedReportStakeInAttoRep;
    uint256 public previousDesignatedReportNoShowBondInAttoRep;

    mapping (address => uint256) public shareSettlementFeeDivisor;
    uint256 public previousReportingFeeDivisor;

    uint256 constant public INITIAL_WINDOW_ID_BUFFER = 365 days * 10 ** 8;
    uint256 constant public DEFAULT_NUM_OUTCOMES = 2;
    uint256 constant public DEFAULT_NUM_TICKS = 100;

    // DAI / DSR specific
    uint256 public totalBalance;
    ICash public cash;
    IDaiVat public daiVat;
    IDaiPot public daiPot;
    IDaiJoin public daiJoin;

    IRepExchange public repExchange;

    uint256 constant public RAY = 10 ** 27;

    constructor(IAugur _augur, IUniverse _parentUniverse, bytes32 _parentPayoutDistributionHash, uint256[] memory _payoutNumerators) public {
        augur = _augur;
        creationTime = _augur.getTimestamp();
        parentUniverse = _parentUniverse;
        parentPayoutDistributionHash = _parentPayoutDistributionHash;
        payoutNumerators = _payoutNumerators;
        reputationToken = IReputationTokenFactory(augur.lookup("ReputationTokenFactory")).createReputationToken(augur, parentUniverse);
        marketFactory = IMarketFactory(augur.lookup("MarketFactory"));
        disputeWindowFactory = IDisputeWindowFactory(augur.lookup("DisputeWindowFactory"));
        openInterestCash = IOICashFactory(augur.lookup("OICashFactory")).createOICash(augur);
        shareToken = IShareToken(augur.lookup("ShareToken"));
        repExchange = IRepExchange(address(IRepExchangeFactory(augur.lookup("RepExchangeFactory")).createRepExchange(augur, address(reputationToken))));
        updateForkValues();
        formulas = IFormulas(augur.lookup("Formulas"));
        cash = ICash(augur.lookup("Cash"));
        daiVat = IDaiVat(augur.lookup("DaiVat"));
        daiPot = IDaiPot(augur.lookup("DaiPot"));
        daiJoin = IDaiJoin(augur.lookup("DaiJoin"));
        assertContractsNotZero();
        daiVat.hope(address(daiPot));
        daiVat.hope(address(daiJoin));
        cash.approve(address(daiJoin), 2 ** 256 - 1);
        daiVat.hope(address(augur));

        initializeCashSender(address(daiVat), address(cash));
    }

    function assertContractsNotZero() private view {
        require(marketFactory != IMarketFactory(0));
        require(disputeWindowFactory != IDisputeWindowFactory(0));
        require(shareToken != IShareToken(0));
        require(formulas != IFormulas(0));
        require(cash != ICash(0));
        require(daiVat != IDaiVat(0));
        require(daiPot != IDaiPot(0));
        require(daiJoin != IDaiJoin(0));
    }

    function fork() public returns (bool) {
        updateForkValues();
        require(!isForking());
        require(isContainerForMarket(IMarket(msg.sender)));
        forkingMarket = IMarket(msg.sender);
        forkEndTime = augur.getTimestamp().add(Reporting.getForkDurationSeconds());
        augur.logUniverseForked(forkingMarket);
        return true;
    }

    function updateForkValues() public returns (bool) {
        require(!isForking());
        uint256 _totalRepSupply = reputationToken.getTotalTheoreticalSupply();
        forkReputationGoal = _totalRepSupply.div(2); // 50% of REP migrating results in a victory in a fork
        disputeThresholdForFork = _totalRepSupply.div(Reporting.getForkThresholdDivisor()); // 2.5% of the total rep supply
        initialReportMinValue = disputeThresholdForFork.div(3).div(2**(Reporting.getMaximumDisputeRounds()-2)).add(1); // This value will result in a maximum 20 round dispute sequence
        disputeThresholdForDisputePacing = disputeThresholdForFork.div(2**(Reporting.getMinimumSlowRounds()+1)); // Disputes begin normal pacing once there are 8 rounds remaining in the fastest case to fork. The "last" round is the one that causes a fork and requires no time so the exponent here is 9 to provide for that many rounds actually occurring.
        return true;
    }

    function getPayoutNumerator(uint256 _outcome) public view returns (uint256) {
        return payoutNumerators[_outcome];
    }

    function getWinningChildPayoutNumerator(uint256 _outcome) public view returns (uint256) {
        return getWinningChildUniverse().getPayoutNumerator(_outcome);
    }

    /**
     * @return This Universe's parent universe or 0x0 if this is the Genesis universe
     */
    function getParentUniverse() public view returns (IUniverse) {
        return parentUniverse;
    }

    /**
     * @return The Bytes32 payout distribution hash of the parent universe or 0x0 if this is the Genesis universe
     */
    function getParentPayoutDistributionHash() public view returns (bytes32) {
        return parentPayoutDistributionHash;
    }

    /**
     * @return The REP token associated with this Universe
     */
    function getReputationToken() public view returns (IV2ReputationToken) {
        return reputationToken;
    }

    /**
     * @return The market in this universe that has triggered a fork if there is one
     */
    function getForkingMarket() public view returns (IMarket) {
        return forkingMarket;
    }

    /**
     * @return The end of the window to migrate REP for the fork if one has occurred in this Universe
     */
    function getForkEndTime() public view returns (uint256) {
        return forkEndTime;
    }

    /**
     * @return The amount of REP migrated into a child universe needed to win a fork
     */
    function getForkReputationGoal() public view returns (uint256) {
        return forkReputationGoal;
    }

    /**
     * @return The amount of REP in a single bond that will trigger a fork if filled
     */
    function getDisputeThresholdForFork() public view returns (uint256) {
        return disputeThresholdForFork;
    }

    /**
     * @return The amount of REP in a single bond that will trigger slow dispute rounds for a market
     */
    function getDisputeThresholdForDisputePacing() public view returns (uint256) {
        return disputeThresholdForDisputePacing;
    }

    /**
     * @return The minimum size of the initial report bond
     */
    function getInitialReportMinValue() public view returns (uint256) {
        return initialReportMinValue;
    }

    /**
     * @return The payout array associated with this universe if it is a child universe from a fork
     */
    function getPayoutNumerators() public view returns (uint256[] memory) {
        return payoutNumerators;
    }

    /**
     * @param _disputeWindowId The id for a dispute window.
     * @return The dispute window for the associated id
     */
    function getDisputeWindow(uint256 _disputeWindowId) public view returns (IDisputeWindow) {
        return disputeWindows[_disputeWindowId];
    }

    /**
     * @return a Bool indicating if this Universe is forking or has forked
     */
    function isForking() public view returns (bool) {
        return forkingMarket != IMarket(0);
    }

    function isForkingMarket() public view returns (bool) {
        return forkingMarket == IMarket(msg.sender);
    }

    /**
     * @param _parentPayoutDistributionHash The payout distribution hash associated with a child Universe to get
     * @return a Universe contract
     */
    function getChildUniverse(bytes32 _parentPayoutDistributionHash) public view returns (IUniverse) {
        return childUniverses[_parentPayoutDistributionHash];
    }

    /**
     * @param _timestamp The timestamp of the desired dispute window
     * @param _initial Bool indicating if the window is an initial dispute window or a standard dispute window
     * @return The id of the specified dispute window
     */
    function getDisputeWindowId(uint256 _timestamp, bool _initial) public view returns (uint256) {
        uint256 _windowId = _timestamp.sub(Reporting.getDisputeWindowBufferSeconds()).div(getDisputeRoundDurationInSeconds(_initial));
        if (_initial) {
            _windowId = _windowId.add(INITIAL_WINDOW_ID_BUFFER);
        }
        return _windowId;
    }

    /**
     * @param _initial Bool indicating if the window is an initial dispute window or a standard dispute window
     * @return The duration in seconds for a dispute window
     */
    function getDisputeRoundDurationInSeconds(bool _initial) public view returns (uint256) {
        return _initial ? Reporting.getInitialDisputeRoundDurationSeconds() : Reporting.getDisputeRoundDurationSeconds();
    }

    /**
     * @param _timestamp The timestamp of the desired dispute window
     * @param _initial Bool indicating if the window is an initial dispute window or a standard dispute window
     * @return The dispute window for the specified params
     */
    function getOrCreateDisputeWindowByTimestamp(uint256 _timestamp, bool _initial) public returns (IDisputeWindow) {
        uint256 _windowId = getDisputeWindowId(_timestamp, _initial);
        if (disputeWindows[_windowId] == IDisputeWindow(0)) {
            (uint256 _startTime, uint256 _duration) = getDisputeWindowStartTimeAndDuration(_timestamp, _initial);
            IDisputeWindow _disputeWindow = disputeWindowFactory.createDisputeWindow(augur, _windowId, _duration, _startTime, !_initial);
            disputeWindows[_windowId] = _disputeWindow;
            augur.logDisputeWindowCreated(_disputeWindow, _windowId, _initial);
        }
        return disputeWindows[_windowId];
    }

    function getDisputeWindowStartTimeAndDuration(uint256 _timestamp, bool _initial) public view returns (uint256 _startTime, uint256 _duration) {
        _duration = getDisputeRoundDurationInSeconds(_initial);
        uint256 _buffer = Reporting.getDisputeWindowBufferSeconds();
        _startTime = _timestamp.sub(_buffer).div(_duration).mul(_duration).add(_buffer);
    }

    /**
     * @param _timestamp The timestamp of the desired dispute window
     * @param _initial Bool indicating if the window is an initial dispute window or a standard dispute window
     * @return The dispute window for the specified params if it exists
     */
    function getDisputeWindowByTimestamp(uint256 _timestamp, bool _initial) public view returns (IDisputeWindow) {
        uint256 _windowId = getDisputeWindowId(_timestamp, _initial);
        return disputeWindows[_windowId];
    }

    /**
     * @param _initial Bool indicating if the window is an initial dispute window or a standard dispute window
     * @return The dispute window before the previous one
     */
    function getOrCreatePreviousPreviousDisputeWindow(bool _initial) public returns (IDisputeWindow) {
        return getOrCreateDisputeWindowByTimestamp(augur.getTimestamp().sub(getDisputeRoundDurationInSeconds(_initial).mul(2)), _initial);
    }

    /**
     * @param _initial Bool indicating if the window is an initial dispute window or a standard dispute window
     * @return The dispute window before the current one
     */
    function getOrCreatePreviousDisputeWindow(bool _initial) public returns (IDisputeWindow) {
        return getOrCreateDisputeWindowByTimestamp(augur.getTimestamp().sub(getDisputeRoundDurationInSeconds(_initial)), _initial);
    }

    /**
     * @param _initial Bool indicating if the window is an initial dispute window or a standard dispute window
     * @return The current dispute window
     */
    function getOrCreateCurrentDisputeWindow(bool _initial) public returns (IDisputeWindow) {
        return getOrCreateDisputeWindowByTimestamp(augur.getTimestamp(), _initial);
    }

    /**
     * @param _initial Bool indicating if the window is an initial dispute window or a standard dispute window
     * @return The current dispute window if it exists
     */
    function getCurrentDisputeWindow(bool _initial) public view returns (IDisputeWindow) {
        return getDisputeWindowByTimestamp(augur.getTimestamp(), _initial);
    }

    /**
     * @param _initial Bool indicating if the window is an initial dispute window or a standard dispute window
     * @return The dispute window after the current one
     */
    function getOrCreateNextDisputeWindow(bool _initial) public returns (IDisputeWindow) {
        return getOrCreateDisputeWindowByTimestamp(augur.getTimestamp().add(getDisputeRoundDurationInSeconds(_initial)), _initial);
    }

    /**
     * @param _parentPayoutNumerators Array of payouts for each outcome associated with the desired child Universe
     * @return The specified Universe
     */
    function createChildUniverse(uint256[] memory _parentPayoutNumerators) public returns (IUniverse) {
        bytes32 _parentPayoutDistributionHash = forkingMarket.derivePayoutDistributionHash(_parentPayoutNumerators);
        IUniverse _childUniverse = getChildUniverse(_parentPayoutDistributionHash);
        if (_childUniverse == IUniverse(0)) {
            _childUniverse = augur.createChildUniverse(_parentPayoutDistributionHash, _parentPayoutNumerators);
            childUniverses[_parentPayoutDistributionHash] = _childUniverse;
        }
        return _childUniverse;
    }

    function updateTentativeWinningChildUniverse(bytes32 _parentPayoutDistributionHash) public returns (bool) {
        IUniverse _tentativeWinningUniverse = getChildUniverse(tentativeWinningChildUniversePayoutDistributionHash);
        IUniverse _updatedUniverse = getChildUniverse(_parentPayoutDistributionHash);
        uint256 _currentTentativeWinningChildUniverseRepMigrated = 0;
        if (_tentativeWinningUniverse != IUniverse(0)) {
            _currentTentativeWinningChildUniverseRepMigrated = _tentativeWinningUniverse.getReputationToken().getTotalMigrated();
        }
        uint256 _updatedUniverseRepMigrated = _updatedUniverse.getReputationToken().getTotalMigrated();
        if (_updatedUniverseRepMigrated > _currentTentativeWinningChildUniverseRepMigrated) {
            tentativeWinningChildUniversePayoutDistributionHash = _parentPayoutDistributionHash;
        }
        if (_updatedUniverseRepMigrated >= forkReputationGoal) {
            forkingMarket.finalize();
        }
        return true;
    }

    /**
     * @return The child Universe which won in a fork if one exists
     */
    function getWinningChildUniverse() public view returns (IUniverse) {
        require(isForking());
        require(tentativeWinningChildUniversePayoutDistributionHash != bytes32(0));
        IUniverse _tentativeWinningUniverse = getChildUniverse(tentativeWinningChildUniversePayoutDistributionHash);
        uint256 _winningAmount = _tentativeWinningUniverse.getReputationToken().getTotalMigrated();
        require(_winningAmount >= forkReputationGoal || augur.getTimestamp() > forkEndTime);
        return _tentativeWinningUniverse;
    }

    function isContainerForDisputeWindow(IDisputeWindow _shadyDisputeWindow) public view returns (bool) {
        uint256 _disputeWindowId = _shadyDisputeWindow.getWindowId();
        IDisputeWindow _legitDisputeWindow = disputeWindows[_disputeWindowId];
        return _shadyDisputeWindow == _legitDisputeWindow;
    }

    function isContainerForMarket(IMarket _shadyMarket) public view returns (bool) {
        return markets[address(_shadyMarket)];
    }

    function migrateMarketOut(IUniverse _destinationUniverse) public returns (bool) {
        IMarket _market = IMarket(msg.sender);
        require(isContainerForMarket(_market));
        markets[msg.sender] = false;
        uint256 _cashBalance = marketBalance[address(msg.sender)];
        uint256 _marketOI = shareToken.totalSupplyForMarketOutcome(_market, 0).mul(_market.getNumTicks());
        withdraw(address(this), _cashBalance, msg.sender);
        openInterestInAttoCash = openInterestInAttoCash.sub(_marketOI);
        cash.approve(address(augur), _cashBalance);
        _destinationUniverse.migrateMarketIn(_market, _cashBalance, _marketOI);
        return true;
    }

    function migrateMarketIn(IMarket _market, uint256 _cashBalance, uint256 _marketOI) public returns (bool) {
        require(address(parentUniverse) == msg.sender);
        markets[address(_market)] = true;
        deposit(address(msg.sender), _cashBalance, address(_market));
        openInterestInAttoCash = openInterestInAttoCash.add(_marketOI);
        augur.logMarketMigrated(_market, parentUniverse);
        return true;
    }

    function isContainerForReportingParticipant(IReportingParticipant _shadyReportingParticipant) public view returns (bool) {
        IMarket _shadyMarket = _shadyReportingParticipant.getMarket();
        if (_shadyMarket == IMarket(0) || !isContainerForMarket(_shadyMarket)) {
            return false;
        }
        return _shadyMarket.isContainerForReportingParticipant(_shadyReportingParticipant);
    }

    function isParentOf(IUniverse _shadyChild) public view returns (bool) {
        bytes32 _parentPayoutDistributionHash = _shadyChild.getParentPayoutDistributionHash();
        return getChildUniverse(_parentPayoutDistributionHash) == _shadyChild;
    }

    function decrementOpenInterest(uint256 _amount) public returns (bool) {
        require(msg.sender == address(shareToken));
        openInterestInAttoCash = openInterestInAttoCash.sub(_amount);
        return true;
    }

    function decrementOpenInterestFromMarket(IMarket _market) public returns (bool) {
        require(isContainerForMarket(IMarket(msg.sender)));
        uint256 _amount = shareToken.totalSupplyForMarketOutcome(_market, 0).mul(_market.getNumTicks());
        openInterestInAttoCash = openInterestInAttoCash.sub(_amount);
        return true;
    }

    function incrementOpenInterest(uint256 _amount) public returns (bool) {
        require(msg.sender == address(shareToken));
        openInterestInAttoCash = openInterestInAttoCash.add(_amount);
        return true;
    }

    /**
     * @return The total amount of Cash in the system which is at risk (Held in escrow for Shares)
     */
    function getOpenInterestInAttoCash() public view returns (uint256) {
        return openInterestInAttoCash;
    }

    /**
     * @return The OI Cash contract
     */
    function isOpenInterestCash(address _address) public view returns (bool) {
        return _address == address(openInterestCash);
    }

    /**
     * @return The Market Cap of this Universe's REP
     */
    function pokeRepMarketCapInAttoCash() public returns (uint256) {
        uint256 _attoCashPerRep = repExchange.pokePrice();
        return getRepMarketCapInAttoCashInternal(_attoCashPerRep);
    }

    function getRepMarketCapInAttoCashInternal(uint256 _attoCashPerRep ) private view returns (uint256) {
        return reputationToken.totalSupply().mul(_attoCashPerRep).div(10 ** 18);
    }

    /**
     * @return The Target Market Cap of this Universe's REP for use in calculating the Reporting Fee
     */
    function getTargetRepMarketCapInAttoCash() public view returns (uint256) {
        // Target MCAP = OI * TARGET_MULTIPLIER
        uint256 _totalOI = openInterestCash.totalSupply().add(getOpenInterestInAttoCash());
        return _totalOI.mul(Reporting.getTargetRepMarketCapMultiplier());
    }

    /**
     * @return The Validity bond required for this dispute window
     */
    function getOrCacheValidityBond() public returns (uint256) {
        IDisputeWindow _disputeWindow = getOrCreateCurrentDisputeWindow(false);
        IDisputeWindow  _previousDisputeWindow = getOrCreatePreviousPreviousDisputeWindow(false);
        uint256 _currentValidityBondInAttoCash = validityBondInAttoCash[address(_disputeWindow)];
        if (_currentValidityBondInAttoCash != 0) {
            return _currentValidityBondInAttoCash;
        }
        if (previousValidityBondInAttoCash == 0) {
            previousValidityBondInAttoCash = Reporting.getDefaultValidityBond();
        }
        _currentValidityBondInAttoCash = formulas.calculateValidityBond(_previousDisputeWindow, previousValidityBondInAttoCash);
        validityBondInAttoCash[address(_disputeWindow)] = _currentValidityBondInAttoCash;
        previousValidityBondInAttoCash = _currentValidityBondInAttoCash;
        augur.logValidityBondChanged(_currentValidityBondInAttoCash);
        return _currentValidityBondInAttoCash;
    }

    /**
     * @return The Designated Report stake for this dispute window
     */
    function getOrCacheDesignatedReportStake() public returns (uint256) {
        updateForkValues();
        IDisputeWindow _disputeWindow = getOrCreateCurrentDisputeWindow(false);
        IDisputeWindow _previousDisputeWindow = getOrCreatePreviousPreviousDisputeWindow(false);
        uint256 _currentDesignatedReportStakeInAttoRep = designatedReportStakeInAttoRep[address(_disputeWindow)];
        if (_currentDesignatedReportStakeInAttoRep != 0) {
            return _currentDesignatedReportStakeInAttoRep;
        }
        if (previousDesignatedReportStakeInAttoRep == 0) {
            previousDesignatedReportStakeInAttoRep = initialReportMinValue;
        }
        _currentDesignatedReportStakeInAttoRep = formulas.calculateDesignatedReportStake(_previousDisputeWindow, previousDesignatedReportStakeInAttoRep, initialReportMinValue);
        designatedReportStakeInAttoRep[address(_disputeWindow)] = _currentDesignatedReportStakeInAttoRep;
        previousDesignatedReportStakeInAttoRep = _currentDesignatedReportStakeInAttoRep;
        augur.logDesignatedReportStakeChanged(_currentDesignatedReportStakeInAttoRep);
        return _currentDesignatedReportStakeInAttoRep;
    }

    /**
     * @return The No Show bond for this dispute window
     */
    function getOrCacheDesignatedReportNoShowBond() public returns (uint256) {
        IDisputeWindow _disputeWindow = getOrCreateCurrentDisputeWindow(false);
        IDisputeWindow _previousDisputeWindow = getOrCreatePreviousPreviousDisputeWindow(false);
        uint256 _currentDesignatedReportNoShowBondInAttoRep = designatedReportNoShowBondInAttoRep[address(_disputeWindow)];
        if (_currentDesignatedReportNoShowBondInAttoRep != 0) {
            return _currentDesignatedReportNoShowBondInAttoRep;
        }
        if (previousDesignatedReportNoShowBondInAttoRep == 0) {
            previousDesignatedReportNoShowBondInAttoRep = initialReportMinValue;
        }
        _currentDesignatedReportNoShowBondInAttoRep = formulas.calculateDesignatedReportNoShowBond(_previousDisputeWindow, previousDesignatedReportNoShowBondInAttoRep, initialReportMinValue);
        designatedReportNoShowBondInAttoRep[address(_disputeWindow)] = _currentDesignatedReportNoShowBondInAttoRep;
        previousDesignatedReportNoShowBondInAttoRep = _currentDesignatedReportNoShowBondInAttoRep;
        augur.logNoShowBondChanged(_currentDesignatedReportNoShowBondInAttoRep);
        return _currentDesignatedReportNoShowBondInAttoRep;
    }

    /**
     * @return The required REP bond for market creation this dispute window
     */
    function getOrCacheMarketRepBond() public returns (uint256) {
        return getOrCacheDesignatedReportNoShowBond().max(getOrCacheDesignatedReportStake());
    }

    /**
     * @dev this should be used in contracts so that the fee is actually set
     * @return The reporting fee for this dispute window
     */
    function getOrCacheReportingFeeDivisor() public returns (uint256) {
        IDisputeWindow _disputeWindow = getOrCreateCurrentDisputeWindow(false);
        uint256 _currentFeeDivisor = shareSettlementFeeDivisor[address(_disputeWindow)];
        if (_currentFeeDivisor != 0) {
            return _currentFeeDivisor;
        }

        _currentFeeDivisor = calculateReportingFeeDivisorInternal();

        shareSettlementFeeDivisor[address(_disputeWindow)] = _currentFeeDivisor;
        previousReportingFeeDivisor = _currentFeeDivisor;
        augur.logReportingFeeChanged(_currentFeeDivisor);
        return _currentFeeDivisor;
    }

    /**
     * @dev this should be used for estimation purposes as it is a view and does not actually freeze or recalculate the rate
     * @return The reporting fee for this dispute window
     */
    function getReportingFeeDivisor() public view returns (uint256) {
        IDisputeWindow _disputeWindow = getCurrentDisputeWindow(false);
        uint256 _currentFeeDivisor = shareSettlementFeeDivisor[address(_disputeWindow)];
        if (_currentFeeDivisor != 0) {
            return _currentFeeDivisor;
        }

        if (previousReportingFeeDivisor == 0) {
            return Reporting.getDefaultReportingFeeDivisor();
        }

        return previousReportingFeeDivisor;
    }

    function calculateReportingFeeDivisorInternal() private returns (uint256) {
        uint256 _repMarketCapInAttoCash = pokeRepMarketCapInAttoCash();
        uint256 _targetRepMarketCapInAttoCash = getTargetRepMarketCapInAttoCash();
        uint256 _reportingFeeDivisor = 0;
        if (previousReportingFeeDivisor == 0 || _targetRepMarketCapInAttoCash == 0) {
            _reportingFeeDivisor = Reporting.getDefaultReportingFeeDivisor();
        } else {
            _reportingFeeDivisor = previousReportingFeeDivisor.mul(_repMarketCapInAttoCash).div(_targetRepMarketCapInAttoCash);
        }

        _reportingFeeDivisor = _reportingFeeDivisor
            .max(Reporting.getMinimumReportingFeeDivisor())
            .min(Reporting.getMaximumReportingFeeDivisor());

        return _reportingFeeDivisor;
    }

    /**
     * @notice Create a Yes / No Market
     * @param _endTime The time at which the event should be reported on. This should be safely after the event outcome is known and verifiable
     * @param _feePerCashInAttoCash The market creator fee specified as the attoCash to be taken from every 1 Cash which is received during settlement
     * @param _affiliateValidator Optional contract which validate the referrer for any attempt at distributing affiliate fees
     * @param _affiliateFeeDivisor The percentage of market creator fees which is designated for affiliates specified as a divisor (4 would mean that 25% of market creator fees may go toward affiliates)
     * @param _designatedReporterAddress The address which will provide the initial report on the market
     * @param _extraInfo Additional info about the market in JSON format.
     * @return The created Market
     */
    function createYesNoMarket(uint256 _endTime, uint256 _feePerCashInAttoCash, IAffiliateValidator _affiliateValidator, uint256 _affiliateFeeDivisor, address _designatedReporterAddress, string memory _extraInfo) public returns (IMarket _newMarket) {
        _newMarket = createMarketInternal(_endTime, _feePerCashInAttoCash, _affiliateValidator, _affiliateFeeDivisor, _designatedReporterAddress, msg.sender, DEFAULT_NUM_OUTCOMES, DEFAULT_NUM_TICKS);
        augur.onYesNoMarketCreated(_endTime, _extraInfo, _newMarket, msg.sender, _designatedReporterAddress, _feePerCashInAttoCash);
        return _newMarket;
    }

    /**
     * @notice Create a Categorical Market
     * @param _endTime The time at which the event should be reported on. This should be safely after the event outcome is known and verifiable
     * @param _feePerCashInAttoCash The market creator fee specified as the attoCash to be taken from every 1 Cash which is received during settlement
     * @param _affiliateValidator Optional contract which validate the referrer for any attempt at distributing affiliate fees
     * @param _affiliateFeeDivisor The percentage of market creator fees which is designated for affiliates specified as a divisor (4 would mean that 25% of market creator fees may go toward affiliates)
     * @param _designatedReporterAddress The address which will provide the initial report on the market
     * @param _outcomes Array of outcome labels / descriptions
     * @param _extraInfo Additional info about the market in JSON format.
     * @return The created Market
     */
    function createCategoricalMarket(uint256 _endTime, uint256 _feePerCashInAttoCash, IAffiliateValidator _affiliateValidator, uint256 _affiliateFeeDivisor, address _designatedReporterAddress, bytes32[] memory _outcomes, string memory _extraInfo) public returns (IMarket _newMarket) {
        _newMarket = createMarketInternal(_endTime, _feePerCashInAttoCash, _affiliateValidator, _affiliateFeeDivisor, _designatedReporterAddress, msg.sender, uint256(_outcomes.length), DEFAULT_NUM_TICKS);
        augur.onCategoricalMarketCreated(_endTime, _extraInfo, _newMarket, msg.sender, _designatedReporterAddress, _feePerCashInAttoCash, _outcomes);
        return _newMarket;
    }

    /**
     * @notice Create a Scalar Market
     * @param _endTime The time at which the event should be reported on. This should be safely after the event outcome is known and verifiable
     * @param _feePerCashInAttoCash The market creator fee specified as the attoCash to be taken from every 1 Cash which is received during settlement
     * @param _affiliateValidator Optional contract which validate the referrer for any attempt at distributing affiliate fees
     * @param _affiliateFeeDivisor The percentage of market creator fees which is designated for affiliates specified as a divisor (4 would mean that 25% of market creator fees may go toward affiliates)
     * @param _designatedReporterAddress The address which will provide the initial report on the market
     * @param _prices 2 element Array comprising a min price and max price in atto units in order to support decimal values. For example if the display range should be between .1 and .5 the prices should be 10**17 and 5 * 10 ** 17 respectively
     * @param _numTicks The number of ticks for the market. This controls the valid price range. Assume a market with min/maxPrices of 0 and 10**18. A numTicks of 100 would mean that the available valid display prices would be .01 to .99 with step size .01. Similarly a numTicks of 10 would be .1 to .9 with a step size of .1.
     * @param _extraInfo Additional info about the market in JSON format.
     * @return The created Market
     */
    function createScalarMarket(uint256 _endTime, uint256 _feePerCashInAttoCash, IAffiliateValidator _affiliateValidator, uint256 _affiliateFeeDivisor, address _designatedReporterAddress, int256[] memory _prices, uint256 _numTicks, string memory _extraInfo) public returns (IMarket _newMarket) {
        _newMarket = createMarketInternal(_endTime, _feePerCashInAttoCash, _affiliateValidator, _affiliateFeeDivisor, _designatedReporterAddress, msg.sender, DEFAULT_NUM_OUTCOMES, _numTicks);
        augur.onScalarMarketCreated(_endTime, _extraInfo, _newMarket, msg.sender, _designatedReporterAddress, _feePerCashInAttoCash, _prices, _numTicks);
        return _newMarket;
    }

    function createMarketInternal(uint256 _endTime, uint256 _feePerCashInAttoCash, IAffiliateValidator _affiliateValidator, uint256 _affiliateFeeDivisor, address _designatedReporterAddress, address _sender, uint256 _numOutcomes, uint256 _numTicks) private returns (IMarket _newMarket) {
        reputationToken.trustedUniverseTransfer(_sender, address(marketFactory), getOrCacheMarketRepBond());
        _newMarket = marketFactory.createMarket(augur, _endTime, _feePerCashInAttoCash, _affiliateValidator, _affiliateFeeDivisor, _designatedReporterAddress, _sender, _numOutcomes, _numTicks);
        markets[address(_newMarket)] = true;
        shareToken.initializeMarket(_newMarket, _numOutcomes + 1, _numTicks); // To account for Invalid
        return _newMarket;
    }

    function saveDaiInDSR(uint256 _amount) private returns (bool) {
        daiJoin.join(address(this), _amount);
        uint256 _chi = daiPot.drip();
        uint256 _sDaiAmount = _amount.mul(RAY) / _chi; // sDai may be lower than the full amount joined above. This means the VAT may have some dust and we'll be saving less than intended by a dust amount
        daiPot.join(_sDaiAmount);
        return true;
    }

    function withdrawDaiFromDSR(uint256 _amount) private returns (bool) {
        uint256 _chi = daiPot.drip();
        uint256 _sDaiAmount = _amount.mul(RAY) / _chi; // sDai may be lower than the amount needed to retrieve `amount` from the VAT. We cover for this rounding error below
        if (_sDaiAmount.mul(_chi) < _amount.mul(RAY)) {
            _sDaiAmount += 1;
        }
        _sDaiAmount = _sDaiAmount.min(daiPot.pie(address(this))); // Never try to draw more than the balance in the pot. If we have less than needed we _must_ have enough already in the VAT provided no negative interest was ever applied
        withdrawSDaiFromDSR(_sDaiAmount);
        return true;
    }

    function withdrawSDaiFromDSR(uint256 _sDaiAmount) private returns (bool) {
        daiPot.exit(_sDaiAmount);
        if (daiJoin.live() == 1) {
            daiJoin.exit(address(this), daiVat.dai(address(this)).div(RAY));
        }
        return true;
    }

    function deposit(address _sender, uint256 _amount, address _market) public returns (bool) {
        require(augur.isTrustedSender(msg.sender) || msg.sender == _sender || msg.sender == address(openInterestCash));
        augur.trustedCashTransfer(_sender, address(this), _amount);
        totalBalance = totalBalance.add(_amount);
        marketBalance[_market] = marketBalance[_market].add(_amount);
        saveDaiInDSR(_amount);
        return true;
    }

    function withdraw(address _recipient, uint256 _amount, address _market) public returns (bool) {
        if (_amount == 0) {
            return true;
        }
        require(augur.isTrustedSender(msg.sender) || augur.isKnownMarket(IMarket(msg.sender)) || msg.sender == address(openInterestCash));
        totalBalance = totalBalance.sub(_amount);
        marketBalance[_market] = marketBalance[_market].sub(_amount);
        withdrawDaiFromDSR(_amount);
        cashTransfer(_recipient, _amount);
        return true;
    }

    function sweepInterest() public returns (bool) {
        uint256 _extraCash = 0;
        uint256 _chi = daiPot.drip();
        withdrawSDaiFromDSR(daiPot.pie(address(this))); // Pull out all funds
        saveDaiInDSR(totalBalance); // Put the required funds back in savings
        _extraCash = cashBalance(address(this));
        // The amount in the DSR pot and VAT must cover our totalBalance of Dai
        assert(daiPot.pie(address(this)).mul(_chi).add(daiVat.dai(address(this))) >= totalBalance.mul(RAY));
        cashTransfer(address(getOrCreateNextDisputeWindow(false)), _extraCash);
        return true;
    }
}
