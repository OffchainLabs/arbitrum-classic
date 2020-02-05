pragma solidity 0.5.15;
pragma experimental ABIEncoderV2;

import './IAugur.sol';
import './IAugurCreationDataGetter.sol';
import './libraries/token/IERC20.sol';
import './libraries/math/SafeMathUint256.sol';
import './factories/IUniverseFactory.sol';
import './reporting/IUniverse.sol';
import './reporting/IMarket.sol';
import './reporting/IDisputeWindow.sol';
import './reporting/IReputationToken.sol';
import './reporting/IReportingParticipant.sol';
import './reporting/IDisputeCrowdsourcer.sol';
import './reporting/IShareToken.sol';
import './trading/IOrders.sol';
import './trading/Order.sol';
import './reporting/Reporting.sol';
import './libraries/ContractExists.sol';
import './ITime.sol';
import './CashSender.sol';
import './reporting/IAffiliates.sol';


// Centralized approval authority and event emissions

/**
 * @title Augur
 * @notice The core global contract of the Augur platform. Provides a contract registry and and authority on which contracts should be trusted.
 */
contract Augur is IAugur, IAugurCreationDataGetter, CashSender {
    using SafeMathUint256 for uint256;
    using ContractExists for address;

    enum TokenType {
        ReputationToken,
        DisputeCrowdsourcer,
        ParticipationToken
    }

    event MarketCreated(IUniverse indexed universe, uint256 endTime, string extraInfo, IMarket market, address indexed marketCreator, address designatedReporter, uint256 feePerCashInAttoCash, int256[] prices, IMarket.MarketType marketType, uint256 numTicks, bytes32[] outcomes, uint256 noShowBond, uint256 timestamp);
    event InitialReportSubmitted(address indexed universe, address indexed reporter, address indexed market, address initialReporter, uint256 amountStaked, bool isDesignatedReporter, uint256[] payoutNumerators, string description, uint256 nextWindowStartTime, uint256 nextWindowEndTime, uint256 timestamp);
    event DisputeCrowdsourcerCreated(address indexed universe, address indexed market, address disputeCrowdsourcer, uint256[] payoutNumerators, uint256 size, uint256 disputeRound);
    event DisputeCrowdsourcerContribution(address indexed universe, address indexed reporter, address indexed market, address disputeCrowdsourcer, uint256 amountStaked, string description, uint256[] payoutNumerators, uint256 currentStake, uint256 stakeRemaining, uint256 disputeRound, uint256 timestamp);
    event DisputeCrowdsourcerCompleted(address indexed universe, address indexed market, address disputeCrowdsourcer, uint256[] payoutNumerators, uint256 nextWindowStartTime, uint256 nextWindowEndTime, bool pacingOn, uint256 totalRepStakedInPayout, uint256 totalRepStakedInMarket, uint256 disputeRound, uint256 timestamp);
    event InitialReporterRedeemed(address indexed universe, address indexed reporter, address indexed market, address initialReporter, uint256 amountRedeemed, uint256 repReceived, uint256[] payoutNumerators, uint256 timestamp);
    event DisputeCrowdsourcerRedeemed(address indexed universe, address indexed reporter, address indexed market, address disputeCrowdsourcer, uint256 amountRedeemed, uint256 repReceived, uint256[] payoutNumerators, uint256 timestamp);
    event ReportingParticipantDisavowed(address indexed universe, address indexed market, address reportingParticipant);
    event MarketParticipantsDisavowed(address indexed universe, address indexed market);
    event MarketFinalized(address indexed universe, address indexed market, uint256 timestamp, uint256[] winningPayoutNumerators);
    event MarketMigrated(address indexed market, address indexed originalUniverse, address indexed newUniverse);
    event UniverseForked(address indexed universe, IMarket forkingMarket);
    event UniverseCreated(address indexed parentUniverse, address indexed childUniverse, uint256[] payoutNumerators, uint256 creationTimestamp);
    event CompleteSetsPurchased(address indexed universe, address indexed market, address indexed account, uint256 numCompleteSets, uint256 timestamp);
    event CompleteSetsSold(address indexed universe, address indexed market, address indexed account, uint256 numCompleteSets, uint256 fees, uint256 timestamp);
    event TradingProceedsClaimed(address indexed universe, address indexed sender, address market, uint256 outcome, uint256 numShares, uint256 numPayoutTokens, uint256 fees, uint256 timestamp);
    event TokensTransferred(address indexed universe, address token, address indexed from, address indexed to, uint256 value, TokenType tokenType, address market);
    event TokensMinted(address indexed universe, address indexed token, address indexed target, uint256 amount, TokenType tokenType, address market, uint256 totalSupply);
    event TokensBurned(address indexed universe, address indexed token, address indexed target, uint256 amount, TokenType tokenType, address market, uint256 totalSupply);
    event TokenBalanceChanged(address indexed universe, address indexed owner, address token, TokenType tokenType, address market, uint256 balance, uint256 outcome);
    event DisputeWindowCreated(address indexed universe, address disputeWindow, uint256 startTime, uint256 endTime, uint256 id, bool initial);
    event InitialReporterTransferred(address indexed universe, address indexed market, address from, address to);
    event MarketTransferred(address indexed universe, address indexed market, address from, address to);
    event MarketOIChanged(address indexed universe, address indexed market, uint256 marketOI);
    event ParticipationTokensRedeemed(address indexed universe, address indexed disputeWindow, address indexed account, uint256 attoParticipationTokens, uint256 feePayoutShare, uint256 timestamp);
    event TimestampSet(uint256 newTimestamp);
    event ValidityBondChanged(address indexed universe, uint256 validityBond);
    event DesignatedReportStakeChanged(address indexed universe, uint256 designatedReportStake);
    event NoShowBondChanged(address indexed universe, uint256 noShowBond);
    event ReportingFeeChanged(address indexed universe, uint256 reportingFee);
    event ShareTokenBalanceChanged(address indexed universe, address indexed account, address indexed market, uint256 outcome, uint256 balance);
    event MarketRepBondTransferred(address indexed universe, address market, address from, address to);
    event WarpSyncDataUpdated(address indexed universe, uint256 warpSyncHash, uint256 marketEndTime);

    event RegisterContract(address contractAddress, bytes32 key);
    event FinishDeployment();

    mapping(address => bool) private markets;
    mapping(address => bool) private universes;
    mapping(address => bool) private crowdsourcers;
    mapping(address => bool) private trustedSender;

    mapping(address => MarketCreationData) private marketCreationData;

    address public uploader;
    mapping(bytes32 => address) private registry;

    ITime public time;
    IUniverse public genesisUniverse;

    uint256 public forkCounter;
    mapping (address => uint256) universeForkIndex;

    uint256 public upgradeTimestamp = Reporting.getInitialUpgradeTimestamp();

    int256 private constant DEFAULT_MIN_PRICE = 0;
    int256 private constant DEFAULT_MAX_PRICE = 1 ether;
    uint256 private constant RECOMMENDED_TRADE_INTERVAL = 10;
    uint256 private constant DEFAULT_RECOMMENDED_TRADE_INTERVAL = 10 ** 17;

    modifier onlyUploader() {
        require(msg.sender == uploader);
        _;
    }

    constructor() public {
        uploader = msg.sender;
    }

    //
    // Registry
    //

    function registerContract(bytes32 _key, address _address) public onlyUploader returns (bool) {
        require(registry[_key] == address(0), "Augur.registerContract: key has already been used in registry");
        require(_address.exists());
        registry[_key] = _address;
        if (_key == "ShareToken" || _key == "MarketFactory" || _key == "EthExchange") {
            trustedSender[_address] = true;
        } else if (_key == "Time") {
            time = ITime(_address);
        } else if (_key == "DaiVat") {
            vat = IDaiVat(_address);
        } else if (_key == "Cash") {
            cash = ICash(_address);
        }
        emit RegisterContract(_address, _key);
        return true;
    }

    /**
     * @notice Find the contract address for a particular key
     * @param _key The key to lookup
     * @return the address of the registered contract if one exists for the given key
     */
    function lookup(bytes32 _key) public view returns (address) {
        return registry[_key];
    }

    function finishDeployment() public onlyUploader returns (bool) {
        uploader = address(1);
        emit FinishDeployment();
        return true;
    }

    //
    // Universe
    //

    function createGenesisUniverse() public onlyUploader returns (IUniverse) {
        require(genesisUniverse == IUniverse(0));
        genesisUniverse = createUniverse(IUniverse(0), bytes32(0), new uint256[](0));
        return genesisUniverse;
    }

    function createChildUniverse(bytes32 _parentPayoutDistributionHash, uint256[] memory _parentPayoutNumerators) public returns (IUniverse) {
        IUniverse _parentUniverse = getAndValidateUniverse(msg.sender);
        return createUniverse(_parentUniverse, _parentPayoutDistributionHash, _parentPayoutNumerators);
    }

    function createUniverse(IUniverse _parentUniverse, bytes32 _parentPayoutDistributionHash, uint256[] memory _parentPayoutNumerators) private returns (IUniverse) {
        IUniverseFactory _universeFactory = IUniverseFactory(registry["UniverseFactory"]);
        IUniverse _newUniverse = _universeFactory.createUniverse(_parentUniverse, _parentPayoutDistributionHash, _parentPayoutNumerators);
        universes[address(_newUniverse)] = true;
        trustedSender[address(_newUniverse)] = true;
        trustedSender[address(_newUniverse.repExchange())] = true;
        emit UniverseCreated(address(_parentUniverse), address(_newUniverse), _parentPayoutNumerators, getTimestamp());
        return _newUniverse;
    }

    function isKnownUniverse(IUniverse _universe) public view returns (bool) {
        return universes[address(_universe)];
    }

    function getUniverseForkIndex(IUniverse _universe) public view returns (uint256) {
        return universeForkIndex[address(_universe)];
    }

    //
    // Crowdsourcers
    //

    function isKnownCrowdsourcer(IDisputeCrowdsourcer _crowdsourcer) public view returns (bool) {
        return crowdsourcers[address(_crowdsourcer)];
    }

    function disputeCrowdsourcerCreated(IUniverse _universe, address _market, address _disputeCrowdsourcer, uint256[] memory _payoutNumerators, uint256 _size, uint256 _disputeRound) public returns (bool) {
        require(isKnownUniverse(_universe));
        require(_universe.isContainerForMarket(IMarket(msg.sender)));
        crowdsourcers[_disputeCrowdsourcer] = true;
        emit DisputeCrowdsourcerCreated(address(_universe), _market, _disputeCrowdsourcer, _payoutNumerators, _size, _disputeRound);
        return true;
    }

    function isKnownFeeSender(address _feeSender) public view returns (bool) {
        return _feeSender == registry["ShareToken"] || markets[_feeSender];
    }

    //
    // Transfer
    //

    function trustedCashTransfer(address _from, address _to, uint256 _amount) public returns (bool) {
        require(trustedSender[msg.sender]);
        cashTransferFrom(_from, _to, _amount);
        return true;
    }

    function isTrustedSender(address _address) public returns (bool) {
        return trustedSender[_address];
    }

    //
    // Time
    //

    /// @notice Returns Augur’s internal Unix timestamp.
    /// @return (uint256) Augur’s internal Unix timestamp
    function getTimestamp() public view returns (uint256) {
        return time.getTimestamp();
    }

    //
    // Markets
    //

    function isKnownMarket(IMarket _market) public view returns (bool) {
        return markets[address(_market)];
    }

    function getMaximumMarketEndDate() public returns (uint256) {
        uint256 _now = getTimestamp();
        while (_now > upgradeTimestamp) {
            upgradeTimestamp = upgradeTimestamp.add(Reporting.getUpgradeCadence());
        }
        uint256 _upgradeCadenceDurationEndTime = upgradeTimestamp;
        uint256 _baseDurationEndTime = _now + Reporting.getBaseMarketDurationMaximum();
        return _baseDurationEndTime.max(_upgradeCadenceDurationEndTime);
    }

    function derivePayoutDistributionHash(uint256[] memory _payoutNumerators, uint256 _numTicks, uint256 _numOutcomes) public view returns (bytes32) {
        uint256 _sum = 0;
        // This is to force an Invalid report to be entirely payed out to Invalid
        require(_payoutNumerators[0] == 0 || _payoutNumerators[0] == _numTicks);
        require(_payoutNumerators.length == _numOutcomes, "Augur.derivePayoutDistributionHash: Malformed payout length");
        for (uint256 i = 0; i < _payoutNumerators.length; i++) {
            uint256 _value = _payoutNumerators[i];
            _sum = _sum.add(_value);
        }
        require(_sum == _numTicks, "Augur.derivePayoutDistributionHash: Malformed payout sum");
        return keccak256(abi.encodePacked(_payoutNumerators));
    }

    function getMarketCreationData(IMarket _market) public view returns (MarketCreationData memory) {
        return marketCreationData[address(_market)];
    }

    function getMarketType(IMarket _market) public view returns (IMarket.MarketType _marketType) {
        return marketCreationData[address(_market)].marketType;
    }

    function getMarketOutcomes(IMarket _market) public view returns (bytes32[] memory _outcomes) {
        return marketCreationData[address(_market)].outcomes;
    }

    //
    // Logging
    //

    function onCategoricalMarketCreated(uint256 _endTime, string memory _extraInfo, IMarket _market, address _marketCreator, address _designatedReporter, uint256 _feePerCashInAttoCash, bytes32[] memory _outcomes) public returns (bool) {
        IUniverse _universe = getAndValidateUniverse(msg.sender);
        markets[address(_market)] = true;
        int256[] memory _prices = new int256[](2);
        _prices[0] = DEFAULT_MIN_PRICE;
        _prices[1] = DEFAULT_MAX_PRICE;
        marketCreationData[address(_market)].extraInfo = _extraInfo;
        marketCreationData[address(_market)].marketCreator = _marketCreator;
        marketCreationData[address(_market)].outcomes = _outcomes;
        marketCreationData[address(_market)].marketType = IMarket.MarketType.CATEGORICAL;
        emit MarketCreated(_universe, _endTime, _extraInfo, _market, _marketCreator, _designatedReporter, _feePerCashInAttoCash, _prices, IMarket.MarketType.CATEGORICAL, 100, _outcomes, _universe.getOrCacheMarketRepBond(), getTimestamp());
        return true;
    }

    function onYesNoMarketCreated(uint256 _endTime, string memory _extraInfo, IMarket _market, address _marketCreator, address _designatedReporter, uint256 _feePerCashInAttoCash) public returns (bool) {
        IUniverse _universe = getAndValidateUniverse(msg.sender);
        markets[address(_market)] = true;
        int256[] memory _prices = new int256[](2);
        _prices[0] = DEFAULT_MIN_PRICE;
        _prices[1] = DEFAULT_MAX_PRICE;
        marketCreationData[address(_market)].extraInfo = _extraInfo;
        marketCreationData[address(_market)].marketCreator = _marketCreator;
        marketCreationData[address(_market)].marketType = IMarket.MarketType.YES_NO;
        emit MarketCreated(_universe, _endTime, _extraInfo, _market, _marketCreator, _designatedReporter, _feePerCashInAttoCash, _prices, IMarket.MarketType.YES_NO, 100, new bytes32[](0), _universe.getOrCacheMarketRepBond(), getTimestamp());
        return true;
    }

    function onScalarMarketCreated(uint256 _endTime, string memory _extraInfo, IMarket _market, address _marketCreator, address _designatedReporter, uint256 _feePerCashInAttoCash, int256[] memory _prices, uint256 _numTicks)  public returns (bool) {
        IUniverse _universe = getAndValidateUniverse(msg.sender);
        require(_prices.length == 2);
        require(_prices[0] < _prices[1]);
        uint256 _priceRange = uint256(_prices[1] - _prices[0]);
        require(_priceRange > _numTicks);
        markets[address(_market)] = true;
        marketCreationData[address(_market)].extraInfo = _extraInfo;
        marketCreationData[address(_market)].marketCreator = _marketCreator;
        marketCreationData[address(_market)].displayPrices = _prices;
        marketCreationData[address(_market)].marketType = IMarket.MarketType.SCALAR;
        emit MarketCreated(_universe, _endTime, _extraInfo, _market, _marketCreator, _designatedReporter, _feePerCashInAttoCash, _prices, IMarket.MarketType.SCALAR, _numTicks, new bytes32[](0), _universe.getOrCacheMarketRepBond(), getTimestamp());
        return true;
    }

    function logInitialReportSubmitted(IUniverse _universe, address _reporter, address _market, address _initialReporter, uint256 _amountStaked, bool _isDesignatedReporter, uint256[] memory _payoutNumerators, string memory _description, uint256 _nextWindowStartTime, uint256 _nextWindowEndTime) public returns (bool) {
        require(isKnownUniverse(_universe));
        require(_universe.isContainerForMarket(IMarket(msg.sender)));
        emit InitialReportSubmitted(address(_universe), _reporter, _market, _initialReporter, _amountStaked, _isDesignatedReporter, _payoutNumerators, _description, _nextWindowStartTime, _nextWindowEndTime, getTimestamp());
        return true;
    }

    function logDisputeCrowdsourcerContribution(IUniverse _universe, address _reporter, address _market, address _disputeCrowdsourcer, uint256 _amountStaked, string memory _description, uint256[] memory _payoutNumerators, uint256 _currentStake, uint256 _stakeRemaining, uint256 _disputeRound) public returns (bool) {
        require(isKnownUniverse(_universe));
        require(_universe.isContainerForMarket(IMarket(msg.sender)));
        emit DisputeCrowdsourcerContribution(address(_universe), _reporter, _market, _disputeCrowdsourcer, _amountStaked, _description, _payoutNumerators, _currentStake, _stakeRemaining, _disputeRound, getTimestamp());
        return true;
    }

    function logDisputeCrowdsourcerCompleted(IUniverse _universe, address _market, address _disputeCrowdsourcer, uint256[] memory _payoutNumerators, uint256 _nextWindowStartTime, uint256 _nextWindowEndTime, bool _pacingOn, uint256 _totalRepStakedInPayout, uint256 _totalRepStakedInMarket, uint256 _disputeRound) public returns (bool) {
        require(isKnownUniverse(_universe));
        require(_universe.isContainerForMarket(IMarket(msg.sender)));
        emit DisputeCrowdsourcerCompleted(address(_universe), _market, _disputeCrowdsourcer, _payoutNumerators, _nextWindowStartTime, _nextWindowEndTime, _pacingOn, _totalRepStakedInPayout, _totalRepStakedInMarket, _disputeRound, getTimestamp());
        return true;
    }

    function logInitialReporterRedeemed(IUniverse _universe, address _reporter, address _market, uint256 _amountRedeemed, uint256 _repReceived, uint256[] memory _payoutNumerators) public returns (bool) {
        require(isKnownUniverse(_universe));
        require(_universe.isContainerForReportingParticipant(IReportingParticipant(msg.sender)));
        emit InitialReporterRedeemed(address(_universe), _reporter, _market, msg.sender, _amountRedeemed, _repReceived, _payoutNumerators, getTimestamp());
        return true;
    }

    function logDisputeCrowdsourcerRedeemed(IUniverse _universe, address _reporter, address _market, uint256 _amountRedeemed, uint256 _repReceived, uint256[] memory _payoutNumerators) public returns (bool) {
        IDisputeCrowdsourcer _disputeCrowdsourcer = IDisputeCrowdsourcer(msg.sender);
        require(isKnownCrowdsourcer(_disputeCrowdsourcer));
        emit DisputeCrowdsourcerRedeemed(address(_universe), _reporter, _market, address(_disputeCrowdsourcer), _amountRedeemed, _repReceived, _payoutNumerators, getTimestamp());
        return true;
    }

    function logReportingParticipantDisavowed(IUniverse _universe, IMarket _market) public returns (bool) {
        require(isKnownUniverse(_universe));
        require(_universe.isContainerForReportingParticipant(IReportingParticipant(msg.sender)));
        emit ReportingParticipantDisavowed(address(_universe), address(_market), msg.sender);
        return true;
    }

    function logMarketParticipantsDisavowed(IUniverse _universe) public returns (bool) {
        require(isKnownUniverse(_universe));
        IMarket _market = IMarket(msg.sender);
        require(_universe.isContainerForMarket(_market));
        emit MarketParticipantsDisavowed(address(_universe), address(_market));
        return true;
    }

    function logMarketFinalized(IUniverse _universe, uint256[] memory _winningPayoutNumerators) public returns (bool) {
        require(isKnownUniverse(_universe));
        IMarket _market = IMarket(msg.sender);
        require(_universe.isContainerForMarket(_market));
        emit MarketFinalized(address(_universe), address(_market), getTimestamp(), _winningPayoutNumerators);
        return true;
    }

    function logMarketMigrated(IMarket _market, IUniverse _originalUniverse) public returns (bool) {
        IUniverse _newUniverse = IUniverse(msg.sender);
        require(isKnownUniverse(_newUniverse));
        emit MarketMigrated(address(_market), address(_originalUniverse), address(_newUniverse));
        return true;
    }

    function logCompleteSetsPurchased(IUniverse _universe, IMarket _market, address _account, uint256 _numCompleteSets) public returns (bool) {
        require(msg.sender == registry["ShareToken"] || (isKnownUniverse(_universe) && _universe.isOpenInterestCash(msg.sender)));
        emit CompleteSetsPurchased(address(_universe), address(_market), _account, _numCompleteSets, getTimestamp());
        return true;
    }

    function logCompleteSetsSold(IUniverse _universe, IMarket _market, address _account, uint256 _numCompleteSets, uint256 _fees) public returns (bool) {
        require(msg.sender == registry["ShareToken"]);
        emit CompleteSetsSold(address(_universe), address(_market), _account, _numCompleteSets, _fees, getTimestamp());
        return true;
    }

    function logMarketOIChanged(IUniverse _universe, IMarket _market) public returns (bool) {
        require(msg.sender == registry["ShareToken"]);
        emit MarketOIChanged(address(_universe), address(_market), _market.getOpenInterest());
        return true;
    }

    function logTradingProceedsClaimed(IUniverse _universe, address _sender, address _market, uint256 _outcome, uint256 _numShares, uint256 _numPayoutTokens, uint256 _fees) public returns (bool) {
        require(msg.sender == registry["ShareToken"]);
        emit TradingProceedsClaimed(address(_universe), _sender, _market, _outcome, _numShares, _numPayoutTokens, _fees, getTimestamp());
        return true;
    }

    function logUniverseForked(IMarket _forkingMarket) public returns (bool) {
        require(isKnownUniverse(IUniverse(msg.sender)));
        forkCounter += 1;
        universeForkIndex[msg.sender] = forkCounter;
        emit UniverseForked(msg.sender, _forkingMarket);
        return true;
    }

    function logReputationTokensTransferred(IUniverse _universe, address _from, address _to, uint256 _value, uint256 _fromBalance, uint256 _toBalance) public returns (bool) {
        require(isKnownUniverse(_universe));
        require(_universe.getReputationToken() == IReputationToken(msg.sender));
        logTokensTransferred(address(_universe), msg.sender, _from, _to, _value, TokenType.ReputationToken, address(0), _fromBalance, _toBalance, 0);
        return true;
    }

    function logDisputeCrowdsourcerTokensTransferred(IUniverse _universe, address _from, address _to, uint256 _value, uint256 _fromBalance, uint256 _toBalance) public returns (bool) {
        IDisputeCrowdsourcer _disputeCrowdsourcer = IDisputeCrowdsourcer(msg.sender);
        require(isKnownCrowdsourcer(_disputeCrowdsourcer));
        logTokensTransferred(address(_universe), msg.sender, _from, _to, _value, TokenType.DisputeCrowdsourcer, address(_disputeCrowdsourcer.getMarket()), _fromBalance, _toBalance, 0);
        return true;
    }

    function logReputationTokensBurned(IUniverse _universe, address _target, uint256 _amount, uint256 _totalSupply, uint256 _balance) public returns (bool) {
        require(isKnownUniverse(_universe));
        require(_universe.getReputationToken() == IReputationToken(msg.sender));
        logTokensBurned(address(_universe), msg.sender, _target, _amount, TokenType.ReputationToken, address(0), _totalSupply, _balance, 0);
        return true;
    }

    function logReputationTokensMinted(IUniverse _universe, address _target, uint256 _amount, uint256 _totalSupply, uint256 _balance) public returns (bool) {
        require(isKnownUniverse(_universe));
        require(_universe.getReputationToken() == IReputationToken(msg.sender));
        logTokensMinted(address(_universe), msg.sender, _target, _amount, TokenType.ReputationToken, address(0), _totalSupply, _balance, 0);
        return true;
    }

    function logShareTokensBalanceChanged(address _account, IMarket _market, uint256 _outcome, uint256 _balance) public returns (bool) {
        require(msg.sender == registry["ShareToken"]);
        emit ShareTokenBalanceChanged(address(_market.getUniverse()), _account, address(_market), _outcome, _balance);
        return true;
    }

    function logDisputeCrowdsourcerTokensBurned(IUniverse _universe, address _target, uint256 _amount, uint256 _totalSupply, uint256 _balance) public returns (bool) {
        IDisputeCrowdsourcer _disputeCrowdsourcer = IDisputeCrowdsourcer(msg.sender);
        require(isKnownCrowdsourcer(_disputeCrowdsourcer));
        logTokensBurned(address(_universe), msg.sender, _target, _amount, TokenType.DisputeCrowdsourcer, address(_disputeCrowdsourcer.getMarket()), _totalSupply, _balance, 0);
        return true;
    }

    function logDisputeCrowdsourcerTokensMinted(IUniverse _universe, address _target, uint256 _amount, uint256 _totalSupply, uint256 _balance) public returns (bool) {
        IDisputeCrowdsourcer _disputeCrowdsourcer = IDisputeCrowdsourcer(msg.sender);
        require(isKnownCrowdsourcer(_disputeCrowdsourcer));
        logTokensMinted(address(_universe), msg.sender, _target, _amount, TokenType.DisputeCrowdsourcer, address(_disputeCrowdsourcer.getMarket()), _totalSupply, _balance, 0);
        return true;
    }

    function logDisputeWindowCreated(IDisputeWindow _disputeWindow, uint256 _id, bool _initial) public returns (bool) {
        require(isKnownUniverse(IUniverse(msg.sender)));
        emit DisputeWindowCreated(msg.sender, address(_disputeWindow), _disputeWindow.getStartTime(), _disputeWindow.getEndTime(), _id, _initial);
        return true;
    }

    function logParticipationTokensRedeemed(IUniverse _universe, address _account, uint256 _attoParticipationTokens, uint256 _feePayoutShare) public returns (bool) {
        require(isKnownUniverse(_universe));
        require(_universe.isContainerForDisputeWindow(IDisputeWindow(msg.sender)));
        emit ParticipationTokensRedeemed(address(_universe), msg.sender, _account, _attoParticipationTokens, _feePayoutShare, getTimestamp());
        return true;
    }

    function logTimestampSet(uint256 _newTimestamp) public returns (bool) {
        require(msg.sender == registry["Time"]);
        emit TimestampSet(_newTimestamp);
        return true;
    }

    function logInitialReporterTransferred(IUniverse _universe, IMarket _market, address _from, address _to) public returns (bool) {
        require(isKnownUniverse(_universe));
        require(_universe.isContainerForMarket(_market));
        require(msg.sender == address(_market.getInitialReporter()));
        emit InitialReporterTransferred(address(_universe), address(_market), _from, _to);
        return true;
    }

    function logMarketTransferred(IUniverse _universe, address _from, address _to) public returns (bool) {
        require(isKnownUniverse(_universe));
        IMarket _market = IMarket(msg.sender);
        require(_universe.isContainerForMarket(_market));
        emit MarketTransferred(address(_universe), address(_market), _from, _to);
        return true;
    }

    function logParticipationTokensTransferred(IUniverse _universe, address _from, address _to, uint256 _value, uint256 _fromBalance, uint256 _toBalance) public returns (bool) {
        require(isKnownUniverse(_universe));
        require(_universe.isContainerForDisputeWindow(IDisputeWindow(msg.sender)));
        logTokensTransferred(address(_universe), msg.sender, _from, _to, _value, TokenType.ParticipationToken, address(0), _fromBalance, _toBalance, 0);
        return true;
    }

    function logParticipationTokensBurned(IUniverse _universe, address _target, uint256 _amount, uint256 _totalSupply, uint256 _balance) public returns (bool) {
        require(isKnownUniverse(_universe));
        require(_universe.isContainerForDisputeWindow(IDisputeWindow(msg.sender)));
        logTokensBurned(address(_universe), msg.sender, _target, _amount, TokenType.ParticipationToken, address(0), _totalSupply, _balance, 0);
        return true;
    }

    function logParticipationTokensMinted(IUniverse _universe, address _target, uint256 _amount, uint256 _totalSupply, uint256 _balance) public returns (bool) {
        require(isKnownUniverse(_universe));
        require(_universe.isContainerForDisputeWindow(IDisputeWindow(msg.sender)));
        logTokensMinted(address(_universe), msg.sender, _target, _amount, TokenType.ParticipationToken, address(0), _totalSupply, _balance, 0);
        return true;
    }

    function logTokensTransferred(address _universe, address _token, address _from, address _to, uint256 _amount, TokenType _tokenType, address _market, uint256 _fromBalance, uint256 _toBalance, uint256 _outcome) private returns (bool) {
        emit TokensTransferred(_universe, _token, _from, _to, _amount, _tokenType, _market);
        emit TokenBalanceChanged(_universe, _from, _token, _tokenType, _market, _fromBalance, _outcome);
        emit TokenBalanceChanged(_universe, _to, _token, _tokenType, _market, _toBalance, _outcome);
        return true;
    }

    function logTokensBurned(address _universe, address _token, address _target, uint256 _amount, TokenType _tokenType, address _market, uint256 _totalSupply, uint256 _balance, uint256 _outcome) private returns (bool) {
        emit TokensBurned(_universe, _token, _target, _amount, _tokenType, _market, _totalSupply);
        emit TokenBalanceChanged(_universe, _target, _token, _tokenType, _market, _balance, _outcome);
        return true;
    }

    function logTokensMinted(address _universe, address _token, address _target, uint256 _amount, TokenType _tokenType, address _market, uint256 _totalSupply, uint256 _balance, uint256 _outcome) private returns (bool) {
        emit TokensMinted(_universe, _token, _target, _amount, _tokenType, _market, _totalSupply);
        emit TokenBalanceChanged(_universe, _target, _token, _tokenType, _market, _balance, _outcome);
        return true;
    }

    function logValidityBondChanged(uint256 _validityBond) public returns (bool) {
        IUniverse _universe = getAndValidateUniverse(msg.sender);
        emit ValidityBondChanged(address(_universe), _validityBond);
        return true;
    }

    function logDesignatedReportStakeChanged(uint256 _designatedReportStake) public returns (bool) {
        IUniverse _universe = getAndValidateUniverse(msg.sender);
        emit DesignatedReportStakeChanged(address(_universe), _designatedReportStake);
        return true;
    }

    function logNoShowBondChanged(uint256 _noShowBond) public returns (bool) {
        IUniverse _universe = getAndValidateUniverse(msg.sender);
        emit NoShowBondChanged(address(_universe), _noShowBond);
        return true;
    }

    function logReportingFeeChanged(uint256 _reportingFee) public returns (bool) {
        IUniverse _universe = getAndValidateUniverse(msg.sender);
        emit ReportingFeeChanged(address(_universe), _reportingFee);
        return true;
    }

    function logMarketRepBondTransferred(address _universe, address _from, address _to) public returns (bool) {
        require(isKnownMarket(IMarket(msg.sender)));
        emit MarketRepBondTransferred(_universe, msg.sender, _from, _to);
    }

    function logWarpSyncDataUpdated(address _universe, uint256 _warpSyncHash, uint256 _marketEndTime) public returns (bool) {
        require(msg.sender == registry["WarpSync"]);
        emit WarpSyncDataUpdated(_universe, _warpSyncHash, _marketEndTime);
    }

    function getAndValidateUniverse(address _untrustedUniverse) internal view returns (IUniverse) {
        IUniverse _universe = IUniverse(_untrustedUniverse);
        require(isKnownUniverse(_universe));
        return _universe;
    }
}
