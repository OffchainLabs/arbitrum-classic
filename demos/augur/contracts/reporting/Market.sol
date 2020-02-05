pragma solidity 0.5.15;

import './IMarket.sol';
import '../libraries/Initializable.sol';
import '../libraries/Ownable.sol';
import './IUniverse.sol';
import './IReportingParticipant.sol';
import './IDisputeCrowdsourcer.sol';
import './IV2ReputationToken.sol';
import './IAffiliateValidator.sol';
import '../external/IDaiVat.sol';
import './IAffiliates.sol';
import '../factories/IDisputeCrowdsourcerFactory.sol';
import '../ICash.sol';
import '../factories/InitialReporterFactory.sol';
import '../libraries/math/SafeMathUint256.sol';
import './Reporting.sol';
import './IInitialReporter.sol';
import '../IWarpSync.sol';
import '../libraries/token/IERC1155.sol';
import '../CashSender.sol';


/**
 * @title Market
 * @notice The contract which encapsulates event data and payout resolution for the event
 */
contract Market is Initializable, Ownable, IMarket, CashSender {
    using SafeMathUint256 for uint256;

    // Constants
    uint256 private constant MAX_APPROVAL_AMOUNT = 2 ** 256 - 1;
    address private constant NULL_ADDRESS = address(0);

    // Contract Refs
    IUniverse private universe;
    IDisputeWindow private disputeWindow;
    IAugur public augur;
    IWarpSync public warpSync;
    IShareToken public shareToken;
    IAffiliateValidator affiliateValidator;
    IAffiliates affiliates;

    // Attributes
    uint256 private numTicks;
    uint256 private feeDivisor;
    uint256 public affiliateFeeDivisor;
    uint256 private endTime;
    uint256 private numOutcomes;
    bytes32 private winningPayoutDistributionHash;
    uint256 public validityBondAttoCash;
    uint256 private finalizationTime;
    uint256 public repBond;
    bool private disputePacingOn;
    address public repBondOwner;
    uint256 public marketCreatorFeesAttoCash;
    uint256 public totalPreFinalizationAffiliateFeesAttoCash;
    IDisputeCrowdsourcer public preemptiveDisputeCrowdsourcer;

    // Collections
    IReportingParticipant[] public participants;

    mapping(bytes32 => address) public crowdsourcers;
    uint256 public crowdsourcerGeneration;

    mapping (address => uint256) public affiliateFeesAttoCash;

    function initialize(IAugur _augur, IUniverse _universe, uint256 _endTime, uint256 _feePerCashInAttoCash, IAffiliateValidator _affiliateValidator, uint256 _affiliateFeeDivisor, address _designatedReporterAddress, address _creator, uint256 _numOutcomes, uint256 _numTicks) public beforeInitialized {
        endInitialization();
        augur = _augur;
        require(msg.sender == _augur.lookup("MarketFactory"));
        _numOutcomes += 1; // The INVALID outcome is always first
        universe = _universe;
        warpSync = IWarpSync(_augur.lookup("WarpSync"));
        require(warpSync != IWarpSync(0));
        affiliateValidator = _affiliateValidator;
        affiliates = IAffiliates(_augur.lookup("Affiliates"));
        require(affiliates != IAffiliates(0));
        require(affiliateValidator == IAffiliateValidator(0) || affiliates.affiliateValidators(address(_affiliateValidator)));
        owner = _creator;
        repBondOwner = owner;
        initializeCashSender(_augur.lookup("DaiVat"), _augur.lookup("Cash"));
        cashApprove(address(_augur), MAX_APPROVAL_AMOUNT);
        assessFees();
        endTime = _endTime;
        numOutcomes = _numOutcomes;
        numTicks = _numTicks;
        feeDivisor = _feePerCashInAttoCash == 0 ? 0 : 1 ether / _feePerCashInAttoCash;
        affiliateFeeDivisor = _affiliateFeeDivisor;
        InitialReporterFactory _initialReporterFactory = InitialReporterFactory(_augur.lookup("InitialReporterFactory"));
        participants.push(_initialReporterFactory.createInitialReporter(_augur, _designatedReporterAddress));
        shareToken = IShareToken(_augur.lookup("ShareToken"));
        require(shareToken != IShareToken(0));
    }

    function assessFees() private {
        repBond = universe.getOrCacheMarketRepBond();
        require(getReputationToken().balanceOf(address(this)) >= repBond);
        if (owner != address(warpSync)) {
            validityBondAttoCash = cashBalance(address(this));
            require(validityBondAttoCash >= universe.getOrCacheValidityBond());
            universe.deposit(address(this), validityBondAttoCash, address(this));
        }
    }

    /**
     * @notice Do the initial report for the market.
     * @param _payoutNumerators An array indicating the payout for each market outcome
     * @param _description Any additional information or justification for this report
     * @param _additionalStake Additional optional REP to stake in anticipation of a dispute. This REP will be held in a bond that only activates if the report is disputed
     * @return Bool True
     */
    function doInitialReport(uint256[] memory _payoutNumerators, string memory _description, uint256 _additionalStake) public returns (bool) {
        doInitialReportInternal(msg.sender, _payoutNumerators, _description);
        if (_additionalStake > 0) {
            contributeToTentativeInternal(msg.sender, _payoutNumerators, _additionalStake, _description);
        }
        return true;
    }

    function doInitialReportInternal(address _reporter, uint256[] memory _payoutNumerators, string memory _description) private {
        require(!universe.isForking());
        IInitialReporter _initialReporter = getInitialReporter();
        uint256 _timestamp = augur.getTimestamp();
        require(_timestamp > endTime);
        uint256 _initialReportStake = distributeInitialReportingRep(_reporter, _initialReporter);
        // The derive call will validate that an Invalid report is entirely paid out on the Invalid outcome
        bytes32 _payoutDistributionHash = derivePayoutDistributionHash(_payoutNumerators);
        disputeWindow = universe.getOrCreateNextDisputeWindow(true);
        _initialReporter.report(_reporter, _payoutDistributionHash, _payoutNumerators, _initialReportStake);
        augur.logInitialReportSubmitted(universe, _reporter, address(this), address(_initialReporter), _initialReportStake, _initialReporter.designatedReporterShowed(), _payoutNumerators, _description, disputeWindow.getStartTime(), disputeWindow.getEndTime());
    }

    function distributeInitialReportingRep(address _reporter, IInitialReporter _initialReporter) private returns (uint256) {
        IV2ReputationToken _reputationToken = getReputationToken();
        uint256 _initialReportStake = repBond;
        repBond = 0;
        // If the designated reporter showed up and is not also the rep bond owner return the rep bond to the bond owner. Otherwise it will be used as stake in the first report.
        if (_reporter == _initialReporter.getDesignatedReporter() && _reporter != repBondOwner) {
            require(_reputationToken.transfer(repBondOwner, _initialReportStake));
            _reputationToken.trustedMarketTransfer(_reporter, address(_initialReporter), _initialReportStake);
        } else {
            require(_reputationToken.transfer(address(_initialReporter), _initialReportStake));
        }
        return _initialReportStake;
    }

    /**
     * @notice Contribute REP to the tentative winning outcome in anticipation of a dispute
     * @dev This will escrow REP in a bond which will be active immediately if the tentative outcome is successfully disputed.
     * @param _payoutNumerators An array indicating the payout for each market outcome
     * @param _amount The amount of REP to contribute
     * @param _description Any additional information or justification for this dispute
     * @return Bool True
     */
    function contributeToTentative(uint256[] memory _payoutNumerators, uint256 _amount, string memory _description) public returns (bool) {
        contributeToTentativeInternal(msg.sender, _payoutNumerators, _amount, _description);
        return true;
    }

    function contributeToTentativeInternal(address _sender, uint256[] memory _payoutNumerators, uint256 _amount, string memory _description) private {
        require(!disputePacingOn);
        // The derive call will validate that an Invalid report is entirely paid out on the Invalid outcome
        bytes32 _payoutDistributionHash = derivePayoutDistributionHash(_payoutNumerators);
        require(_payoutDistributionHash == getWinningReportingParticipant().getPayoutDistributionHash());
        internalContribute(_sender, _payoutDistributionHash, _payoutNumerators, _amount, true, _description);
    }

    /**
     * @notice Contribute REP to a payout other than the tenative winning outcome in order to dispute it
     * @param _payoutNumerators An array indicating the payout for each market outcome
     * @param _amount The amount of REP to contribute
     * @param _description Any additional information or justification for this dispute
     * @return Bool True
     */
    function contribute(uint256[] memory _payoutNumerators, uint256 _amount, string memory _description) public returns (bool) {
        // The derive call will validate that an Invalid report is entirely paid out on the Invalid outcome
        bytes32 _payoutDistributionHash = derivePayoutDistributionHash(_payoutNumerators);
        require(_payoutDistributionHash != getWinningReportingParticipant().getPayoutDistributionHash());
        internalContribute(msg.sender, _payoutDistributionHash, _payoutNumerators, _amount, false, _description);
        return true;
    }

    function internalContribute(address _contributor, bytes32 _payoutDistributionHash, uint256[] memory _payoutNumerators, uint256 _amount, bool _overload, string memory _description) internal {
        if (disputePacingOn) {
            require(disputeWindow.isActive());
        } else {
            require(!disputeWindow.isOver());
        }
        // This will require that the universe is not forking
        universe.updateForkValues();
        IDisputeCrowdsourcer _crowdsourcer = getOrCreateDisputeCrowdsourcer(_payoutDistributionHash, _payoutNumerators, _overload);
        uint256 _actualAmount = _crowdsourcer.contribute(_contributor, _amount, _overload);
        uint256 _amountRemainingToFill = _overload ? 0 : _crowdsourcer.getRemainingToFill();
        augur.logDisputeCrowdsourcerContribution(universe, _contributor, address(this), address(_crowdsourcer), _actualAmount, _description, _payoutNumerators, _crowdsourcer.getStake(), _amountRemainingToFill, getNumParticipants());
        if (!_overload) {
            if (_amountRemainingToFill == 0) {
                finishedCrowdsourcingDisputeBond(_crowdsourcer);
            } else {
                require(_amountRemainingToFill >= getInitialReporter().getSize());
            }
        }
    }

    function finishedCrowdsourcingDisputeBond(IDisputeCrowdsourcer _crowdsourcer) private {
        correctLastParticipantSize();
        participants.push(_crowdsourcer);
        clearCrowdsourcers(); // disavow other crowdsourcers
        uint256 _crowdsourcerSize = IDisputeCrowdsourcer(_crowdsourcer).getSize();
        if (_crowdsourcerSize >= universe.getDisputeThresholdForFork()) {
            universe.fork();
        } else {
            if (_crowdsourcerSize >= universe.getDisputeThresholdForDisputePacing()) {
                disputePacingOn = true;
            }
            disputeWindow = universe.getOrCreateNextDisputeWindow(false);
        }
        augur.logDisputeCrowdsourcerCompleted(
            universe,
            address(this),
            address(_crowdsourcer),
            _crowdsourcer.getPayoutNumerators(),
            disputeWindow.getStartTime(),
            disputeWindow.getEndTime(),
            disputePacingOn,
            getStakeInOutcome(_crowdsourcer.getPayoutDistributionHash()),
            getParticipantStake(),
            participants.length);
        if (preemptiveDisputeCrowdsourcer != IDisputeCrowdsourcer(0)) {
            IDisputeCrowdsourcer _newCrowdsourcer = preemptiveDisputeCrowdsourcer;
            preemptiveDisputeCrowdsourcer = IDisputeCrowdsourcer(0);
            bytes32 _payoutDistributionHash = _newCrowdsourcer.getPayoutDistributionHash();
            // The size of any dispute bond should be (2 * ALL STAKE) - (3 * STAKE IN OUTCOME)
            uint256 _correctSize = getParticipantStake().mul(2).sub(getStakeInOutcome(_payoutDistributionHash).mul(3));
            _newCrowdsourcer.setSize(_correctSize);
            if (_newCrowdsourcer.getStake() >= _correctSize) {
                finishedCrowdsourcingDisputeBond(_newCrowdsourcer);
            } else {
                crowdsourcers[_payoutDistributionHash] = address(_newCrowdsourcer);
            }
        }
    }

    function correctLastParticipantSize() private {
        // A dispute has occured if there is more than one completed reporting participant
        if (participants.length > 1) {
            IDisputeCrowdsourcer(address(getWinningReportingParticipant())).correctSize();
        }
    }

    /**
     * @notice Finalize a market
     * @return Bool True
     */
    function finalize() public returns (bool) {
        require(!isFinalized());
        uint256[] memory _winningPayoutNumerators;
        if (isForkingMarket()) {
            IUniverse _winningUniverse = universe.getWinningChildUniverse();
            winningPayoutDistributionHash = _winningUniverse.getParentPayoutDistributionHash();
            _winningPayoutNumerators = _winningUniverse.getPayoutNumerators();
        } else {
            require(disputeWindow.isOver());
            require(!universe.isForking());
            IReportingParticipant _reportingParticipant = getWinningReportingParticipant();
            winningPayoutDistributionHash = _reportingParticipant.getPayoutDistributionHash();
            _winningPayoutNumerators = _reportingParticipant.getPayoutNumerators();
            warpSync.notifyMarketFinalized();
            // Make sure the dispute window for which we record finalization is the standard cadence window and not an initial dispute window
            disputeWindow = universe.getOrCreatePreviousDisputeWindow(false);
            disputeWindow.onMarketFinalized();
            universe.decrementOpenInterestFromMarket(this);
            redistributeLosingReputation();
        }
        finalizationTime = augur.getTimestamp();
        distributeValidityBondAndMarketCreatorFees();
        augur.logMarketFinalized(universe, _winningPayoutNumerators);
        return true;
    }

    function redistributeLosingReputation() private {
        // If no disputes occurred early exit
        if (participants.length == 1) {
            return;
        }

        IReportingParticipant _reportingParticipant;

        // Initial pass is to liquidate losers so we have sufficient REP to pay the winners. Participants is implicitly bounded by the floor of the initial report REP cost to be no more than 21
        for (uint256 i = 0; i < participants.length; i++) {
            _reportingParticipant = participants[i];
            if (_reportingParticipant.getPayoutDistributionHash() != winningPayoutDistributionHash) {
                _reportingParticipant.liquidateLosing();
            }
        }

        IV2ReputationToken _reputationToken = getReputationToken();
        // We burn 20% of the REP to prevent griefing attacks which rely on getting back lost REP
        _reputationToken.burnForMarket(_reputationToken.balanceOf(address(this)) / 5);

        // Now redistribute REP. Participants is implicitly bounded by the floor of the initial report REP cost to be no more than 21.
        for (uint256 j = 0; j < participants.length; j++) {
            _reportingParticipant = participants[j];
            if (_reportingParticipant.getPayoutDistributionHash() == winningPayoutDistributionHash) {
                // The last participant's owed REP will not actually be 40% ROI in the event it was created through pre-emptive contributions. We just give them all the remaining non burn REP
                uint256 amountToTransfer = j == participants.length - 1 ? _reputationToken.balanceOf(address(this)) : _reportingParticipant.getSize().mul(2) / 5;
                require(_reputationToken.transfer(address(_reportingParticipant), amountToTransfer));
            }
        }
    }

    /**
     * @return The amount any settlement proceeds are divided by in order to calculate the market creator fee portion
     */
    function getMarketCreatorSettlementFeeDivisor() public view returns (uint256) {
        return feeDivisor;
    }

    /**
     * @param _amount The total settlement proceeds of a trade or claim
     * @return The amount of fees the market creator will receive
     */
    function deriveMarketCreatorFeeAmount(uint256 _amount) public view returns (uint256) {
        return feeDivisor == 0 ? 0 : _amount / feeDivisor;
    }

    function recordMarketCreatorFees(uint256 _marketCreatorFees, address _sourceAccount, bytes32 _fingerprint) public returns (bool) {
        require(augur.isKnownFeeSender(msg.sender));

        address _affiliateAddress = affiliates.getAndValidateReferrer(_sourceAccount, affiliateValidator);
        bytes32 _affiliateFingerprint = affiliates.getAccountFingerprint(_affiliateAddress);
        if (_fingerprint == _affiliateFingerprint) {
            // don't let affiliates refer themselves
            _affiliateAddress = address(0);
        }

        if (_affiliateAddress != NULL_ADDRESS && affiliateFeeDivisor != 0) {
            uint256 _totalAffiliateFees = _marketCreatorFees / affiliateFeeDivisor;
            uint256 _sourceCut = _totalAffiliateFees / Reporting.getAffiliateSourceCutDivisor();
            uint256 _affiliateFees = _totalAffiliateFees.sub(_sourceCut);
            universe.withdraw(_sourceAccount, _sourceCut, address(this));
            affiliateFeesAttoCash[_affiliateAddress] += _affiliateFees;
            _marketCreatorFees = _marketCreatorFees.sub(_totalAffiliateFees);
            totalPreFinalizationAffiliateFeesAttoCash = totalPreFinalizationAffiliateFeesAttoCash.add(_affiliateFees);
        }

        marketCreatorFeesAttoCash = marketCreatorFeesAttoCash.add(_marketCreatorFees);

        if (isFinalized()) {
            distributeMarketCreatorAndAffiliateFees(_affiliateAddress);
        }
    }

    function distributeValidityBondAndMarketCreatorFees() private {
        // If the market resolved to invalid the bond gets sent to the dispute window. Otherwise it gets returned to the market creator.
        marketCreatorFeesAttoCash = validityBondAttoCash.add(marketCreatorFeesAttoCash);
        distributeMarketCreatorAndAffiliateFees(NULL_ADDRESS);
    }

    function distributeMarketCreatorAndAffiliateFees(address _affiliateAddress) private {
        uint256 _marketCreatorFeesAttoCash = marketCreatorFeesAttoCash;
        marketCreatorFeesAttoCash = 0;
        if (!isFinalizedAsInvalid()) {
            universe.withdraw(owner, _marketCreatorFeesAttoCash, address(this));
            if (_affiliateAddress != NULL_ADDRESS) {
                withdrawAffiliateFees(_affiliateAddress);
            }
        } else {
            universe.withdraw(address(universe.getOrCreateNextDisputeWindow(false)), _marketCreatorFeesAttoCash.add(totalPreFinalizationAffiliateFeesAttoCash), address(this));
            totalPreFinalizationAffiliateFeesAttoCash = 0;
        }
    }

    /**
     * @notice Redeems any owed affiliate fees for a particular address
     * @dev Will fail if the market is Invalid
     * @param _affiliate The address that is owed affiliate fees
     * @return Bool True
     */
    function withdrawAffiliateFees(address _affiliate) public returns (bool) {
        require(!isFinalizedAsInvalid());
        uint256 _affiliateBalance = affiliateFeesAttoCash[_affiliate];
        if (_affiliateBalance == 0) {
            return true;
        }
        affiliateFeesAttoCash[_affiliate] = 0;
        universe.withdraw(_affiliate, _affiliateBalance, address(this));
        return true;
    }

    function getOrCreateDisputeCrowdsourcer(bytes32 _payoutDistributionHash, uint256[] memory _payoutNumerators, bool _overload) private returns (IDisputeCrowdsourcer) {
        IDisputeCrowdsourcer _crowdsourcer = _overload ? preemptiveDisputeCrowdsourcer : IDisputeCrowdsourcer(getCrowdsourcer(_payoutDistributionHash));
        if (_crowdsourcer == IDisputeCrowdsourcer(0)) {
            IDisputeCrowdsourcerFactory _disputeCrowdsourcerFactory = IDisputeCrowdsourcerFactory(augur.lookup("DisputeCrowdsourcerFactory"));
            uint256 _participantStake = getParticipantStake();
            if (_overload) {
                // The stake of a dispute bond is (2 * ALL STAKE) - (3 * STAKE IN OUTCOME)
                _participantStake = _participantStake.add(_participantStake.mul(2).sub(getHighestNonTentativeParticipantStake().mul(3)));
            }
            uint256 _size = _participantStake.mul(2).sub(getStakeInOutcome(_payoutDistributionHash).mul(3));
            uint256 _crowdsourcerGeneration = crowdsourcerGeneration;
            if (_overload) {
                // If the preemptive crowdsourcer is used, it will always enter at the next generation
                _crowdsourcerGeneration += 1;
            }
            _crowdsourcer = _disputeCrowdsourcerFactory.createDisputeCrowdsourcer(augur, _size, _payoutDistributionHash, _payoutNumerators, _crowdsourcerGeneration);
            if (!_overload) {
                crowdsourcers[_payoutDistributionHash] = address(_crowdsourcer);
            } else {
                preemptiveDisputeCrowdsourcer = _crowdsourcer;
            }
            augur.disputeCrowdsourcerCreated(universe, address(this), address(_crowdsourcer), _payoutNumerators, _size, getNumParticipants());
        }
        return _crowdsourcer;
    }

    /**
     * @notice Migrates the market through a fork into the winning Universe
     * @dev This will extract a new REP no show bond from whoever calls this and if the market is in the reporting phase will require a report be made as well
     * @param _payoutNumerators An array indicating the payout for each market outcome
     * @param _description Any additional information or justification for this report
     * @return Bool True
     */
    function migrateThroughOneFork(uint256[] memory _payoutNumerators, string memory _description) public returns (bool) {
        // only proceed if the forking market is finalized
        IMarket _forkingMarket = universe.getForkingMarket();
        require(_forkingMarket.isFinalized());
        require(!isFinalized());
        require(this != warpSync.markets(address(universe)));

        disavowCrowdsourcers();

        bytes32 _winningForkPayoutDistributionHash = _forkingMarket.getWinningPayoutDistributionHash();
        IUniverse _destinationUniverse = universe.getChildUniverse(_winningForkPayoutDistributionHash);

        // follow the forking market to its universe
        if (disputeWindow != IDisputeWindow(0)) {
            // Markets go into the standard resolution period during fork migration even if they were in the initial dispute window. We want to give some time for REP to migrate.
            disputeWindow = _destinationUniverse.getOrCreateNextDisputeWindow(false);
        }
        universe.migrateMarketOut(_destinationUniverse);
        universe = _destinationUniverse;
        uint256 _numOutcomes = numOutcomes;

        // Pay the REP bond.
        repBond = universe.getOrCacheMarketRepBond();
        repBondOwner = msg.sender;
        getReputationToken().trustedMarketTransfer(repBondOwner, address(this), repBond);

        // Update the Initial Reporter
        IInitialReporter _initialReporter = getInitialReporter();
        _initialReporter.migrateToNewUniverse(msg.sender);

        // If the market is past expiration use the reporting data to make an initial report
        uint256 _timestamp = augur.getTimestamp();
        if (_timestamp > endTime) {
            doInitialReportInternal(msg.sender, _payoutNumerators, _description);
        }

        return true;
    }

    function disavowCrowdsourcers() public returns (bool) {
        IMarket _forkingMarket = getForkingMarket();
        require(_forkingMarket != this);
        require(_forkingMarket != IMarket(NULL_ADDRESS));
        require(!isFinalized());
        IInitialReporter _initialParticipant = getInitialReporter();
        delete participants;
        participants.push(_initialParticipant);
        clearCrowdsourcers();
        preemptiveDisputeCrowdsourcer = IDisputeCrowdsourcer(0);
        // Send REP from the rep bond back to the address that placed it. If a report has been made tell the InitialReporter to return that REP and reset
        if (repBond > 0) {
            IV2ReputationToken _reputationToken = getReputationToken();
            uint256 _repBond = repBond;
            require(_reputationToken.transfer(repBondOwner, _repBond));
            repBond = 0;
        } else {
            _initialParticipant.returnRepFromDisavow();
        }
        augur.logMarketParticipantsDisavowed(universe);
        return true;
    }

    function clearCrowdsourcers() private {
        crowdsourcerGeneration += 1;
    }

    function getHighestNonTentativeParticipantStake() public view returns (uint256) {
        if (participants.length < 2) {
            return 0;
        }
        bytes32 _payoutDistributionHash = participants[participants.length - 2].getPayoutDistributionHash();
        return getStakeInOutcome(_payoutDistributionHash);
    }

    /**
     * @notice Gets all REP stake in completed bonds for this market
     * @return uint256 indicating sum of all stake
     */
    function getParticipantStake() public view returns (uint256) {
        uint256 _sum;
        // Participants is implicitly bounded by the floor of the initial report REP cost to be no more than 21
        for (uint256 i = 0; i < participants.length; ++i) {
            _sum += participants[i].getStake();
        }
        return _sum;
    }

    /**
     * @param _payoutDistributionHash the payout distribution hash being checked
     * @return uint256 indicating the REP stake in a single outcome for a particular payout hash
     */
    function getStakeInOutcome(bytes32 _payoutDistributionHash) public view returns (uint256) {
        uint256 _sum;
        // Participants is implicitly bounded by the floor of the initial report REP cost to be no more than 21
        for (uint256 i = 0; i < participants.length; ++i) {
            IReportingParticipant _reportingParticipant = participants[i];
            if (_reportingParticipant.getPayoutDistributionHash() != _payoutDistributionHash) {
                continue;
            }
            _sum = _sum.add(_reportingParticipant.getStake());
        }
        return _sum;
    }

    /**
     * @return The forking market for the associated universe if one exists
     */
    function getForkingMarket() public view returns (IMarket) {
        return universe.getForkingMarket();
    }

    /**
     * @return The current bytes32 winning distribution hash if one exists
     */
    function getWinningPayoutDistributionHash() public view returns (bytes32) {
        return winningPayoutDistributionHash;
    }

    /**
     * @return Bool indicating if the market is finalized
     */
    function isFinalized() public view returns (bool) {
        return winningPayoutDistributionHash != bytes32(0);
    }

    /**
     * @return Time at which the event is considered ready to report on
     */
    function getEndTime() public view returns (uint256) {
        return endTime;
    }

    /**
     * @return Bool indicating if the market resolved as anything other than Invalid
     */
    function isFinalizedAsInvalid() public view returns (bool) {
        require(isFinalized());
        if (isForkingMarket()) {
            return getWinningChildPayout(0) > 0;
        }
        return getWinningReportingParticipant().getPayoutNumerator(0) > 0;
    }

    /**
     * @return The Initial Reporter contract
     */
    function getInitialReporter() public view returns (IInitialReporter) {
        return IInitialReporter(address(participants[0]));
    }

    /**
     * @param _payoutDistributionHash The payout distribution hash for a Dispute Crowdsourcer contract for this round of disputing
     * @return The associated Dispute Crowdsourcer contract for this round of disputing
     */
    function getCrowdsourcer(bytes32 _payoutDistributionHash) public view returns (IDisputeCrowdsourcer) {
        IDisputeCrowdsourcer _crowdsourcer = IDisputeCrowdsourcer(crowdsourcers[_payoutDistributionHash]);
        if (_crowdsourcer != IDisputeCrowdsourcer(0) && _crowdsourcer.getCrowdsourcerGeneration() == crowdsourcerGeneration) {
            return _crowdsourcer;
        }
        return IDisputeCrowdsourcer(0);
    }

    /**
     * @return The associated Initial Reporter or a Dispute Crowdsourcer contract for the current tentative winning payout
     */
    function getWinningReportingParticipant() public view returns (IReportingParticipant) {
        return participants[participants.length-1];
    }

    /**
     * @param _outcome The outcome to get a payout for
     * @return The payout for a particular outcome for the tentative winning payout
     */
    function getWinningPayoutNumerator(uint256 _outcome) public view returns (uint256) {
        if (isForkingMarket()) {
            return getWinningChildPayout(_outcome);
        }
        return getWinningReportingParticipant().getPayoutNumerator(_outcome);
    }

    /**
     * @return The Universe associated with this Market
     */
    function getUniverse() public view returns (IUniverse) {
        return universe;
    }

    /**
     * @return The Dispute Window currently associated with this Market
     */
    function getDisputeWindow() public view returns (IDisputeWindow) {
        return disputeWindow;
    }

    /**
     * @return The time the Market was finalzied as a uint256 timestmap if the market was finalized
     */
    function getFinalizationTime() public view returns (uint256) {
        return finalizationTime;
    }

    /**
     * @return The REP token associated with this Market
     */
    function getReputationToken() public view returns (IV2ReputationToken) {
        return universe.getReputationToken();
    }

    /**
     * @return The number of outcomes (including invalid) this market has
     */
    function getNumberOfOutcomes() public view returns (uint256) {
        return numOutcomes;
    }

    /**
     * @return The number of ticks for this market. The number of ticks determines the possible on chain prices for Shares of the market. (e.g. A Market with 10 ticks can have prices 1-9 and a complete set will cost 10)
     */
    function getNumTicks() public view returns (uint256) {
        return numTicks;
    }

    /**
     * @return The uint256 timestamp for when the designated reporting period is over and anyone may report
     */
    function getDesignatedReportingEndTime() public view returns (uint256) {
        return endTime.add(Reporting.getDesignatedReportingDurationSeconds());
    }

    /**
     * @return The number of rounds of reporting + disputing that have occured
     */
    function getNumParticipants() public view returns (uint256) {
        return participants.length;
    }

    /**
     * @return The size of the validity bond
     */
    function getValidityBondAttoCash() public view returns (uint256) {
        return validityBondAttoCash;
    }

    /**
     * @return Bool indicating if slow dispute rounds have turned on
     */
    function getDisputePacingOn() public view returns (bool) {
        return disputePacingOn;
    }

    /**
     * @param _payoutNumerators array of payouts per outcome
     * @return Bytes32 has of the payout for use in other functions
     */
    function derivePayoutDistributionHash(uint256[] memory _payoutNumerators) public view returns (bytes32) {
        return augur.derivePayoutDistributionHash(_payoutNumerators, numTicks, numOutcomes);
    }

    function isContainerForReportingParticipant(IReportingParticipant _shadyReportingParticipant) public view returns (bool) {
        require(_shadyReportingParticipant != IReportingParticipant(0));
        if (address(preemptiveDisputeCrowdsourcer) == address(_shadyReportingParticipant)) {
            return true;
        }
        if (getCrowdsourcer(_shadyReportingParticipant.getPayoutDistributionHash()) == _shadyReportingParticipant) {
            return true;
        }
        // Participants is implicitly bounded by the floor of the initial report REP cost to be no more than 21
        for (uint256 i = 0; i < participants.length; i++) {
            if (_shadyReportingParticipant == participants[i]) {
                return true;
            }
        }
        return false;
    }

    function onTransferOwnership(address _owner, address _newOwner) internal {
        augur.logMarketTransferred(getUniverse(), _owner, _newOwner);
    }

    /**
     * @notice Transfers ownership of the REP no-show bond
     * @param _newOwner The new REP no show bond owner
     * @return Bool True
     */
    function transferRepBondOwnership(address _newOwner) public returns (bool) {
        require(_newOwner != address(0));
        require(msg.sender == repBondOwner);
        address _oldOwner = repBondOwner;
        repBondOwner = _newOwner;
        augur.logMarketRepBondTransferred(address(universe), _oldOwner, _newOwner);
        return true;
    }

    function isForkingMarket() public view returns (bool) {
        return universe.isForkingMarket();
    }

    function getWinningChildPayout(uint256 _outcome) public view returns (uint256) {
        return universe.getWinningChildPayoutNumerator(_outcome);
    }

    function getOpenInterest() public view returns (uint256) {
        if (isFinalized()) {
            return 0;
        }
        return shareToken.totalSupplyForMarketOutcome(this, 0).mul(numTicks);
    }
}
