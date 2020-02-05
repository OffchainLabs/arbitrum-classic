pragma solidity 0.5.15;

import './IWarpSync.sol';
import './reporting/IUniverse.sol';
import './reporting/IMarket.sol';
import './reporting/IV2ReputationToken.sol';
import './reporting/IAffiliateValidator.sol';
import './libraries/Initializable.sol';
import './libraries/math/SafeMathUint256.sol';

/**
 * @title Warp Sync
 * @notice A contract that control the recurring creation of a market which exists to report on warp sync data for Augur
 */
contract WarpSync is IWarpSync, Initializable {
    using SafeMathUint256 for uint256;

    struct Data {
        uint256 warpSyncHash;
        uint256 timestamp;
    }

    IAugur public augur;
    mapping(address => address) public markets;
    mapping(address => Data) public data;

    uint256 public lastSweepTime;

    uint256 private constant MIN_TIME_BETWEEN_INTEREST_SWEEPS = 7 days;

    uint256 private constant MARKET_LENGTH = 1 days;
    uint256 private constant MAX_NUM_TICKS = 2 ** 256 - 2;
    int256 private INT256_MIN = int256(2**255);
    int256 private INT256_MAX = int256(2**255 - 1);
    int256[] private PRICES = [INT256_MIN, INT256_MAX];
    string private constant EXTRA_INFO = '{"description":"What will the next Augur Warp Sync hash be?","longDescription":"What will the Augur SDK warp sync hash be for the last block with a timestamp less than the reporting start timestamp for this market?"}';

    function initialize(IAugur _augur) public beforeInitialized returns (bool) {
        endInitialization();
        augur = _augur;
        lastSweepTime = block.timestamp;
        return true;
    }

    /**
     * @notice Do the initial report for the warp sync market.
     * @param _universe The universe whose warp sync market to report on
     * @param _payoutNumerators An array indicating the payout for each market outcome
     * @param _description Any additional information or justification for this report
     * @param _additionalStake Additional optional REP to stake in anticipation of a dispute. This REP will be held in a bond that only activates if the report is disputed
     * @return Bool True
     */
    function doInitialReport(IUniverse _universe, uint256[] memory _payoutNumerators, string memory _description, uint256 _additionalStake) public returns (bool) {
        IMarket _market = IMarket(markets[address(_universe)]);
        _market.doInitialReport(_payoutNumerators, _description, _additionalStake);
        _market.getInitialReporter().transferOwnership(msg.sender);
        return true;
    }

    /**
     * @notice Create the initial warp sync market. The creator will be rewarded with REP
     * @param _universe The universe to create a warp sync market in
     */
    function initializeUniverse(IUniverse _universe) public {
        require(augur.isKnownUniverse(_universe));
        require(markets[address(_universe)] == address(0));
        awardRep(_universe, getCreationReward(_universe));
        createMarket(_universe);
    }

    function notifyMarketFinalized() public {
        IMarket _market = IMarket(msg.sender);
        IUniverse _universe = _market.getUniverse();

        // NOTE: This validates that the market is legitimate. A malicious market has no way of modifying this mapping to pass here.
        if (markets[address(_universe)] != address(_market)) {
            return;
        }

        // In order to periodically sweep interest we do so here which will result in a sweep at least on recurring basis where the sweeper is compensated.
        uint256 _timestamp = block.timestamp;
        if (lastSweepTime - _timestamp >= MIN_TIME_BETWEEN_INTEREST_SWEEPS) {
            _universe.sweepInterest();
            lastSweepTime = _timestamp;
        }

        recordMarketFinalized(_market, _universe);

        if (!_universe.isForking()) {
            createMarket(_universe);
        }
    }

    function recordMarketFinalized(IMarket _market, IUniverse _universe) private {
        awardRep(_universe, getFinalizationReward(_market));
        uint256 _warpSyncHash = _market.getWinningPayoutNumerator(2);
        uint256 _endTime = _market.getEndTime();
        data[address(_universe)].warpSyncHash = _warpSyncHash;
        data[address(_universe)].timestamp = _endTime;
        augur.logWarpSyncDataUpdated(address(_universe), _warpSyncHash, _endTime);
    }

    /**
     * @notice Get the REP reward for finalizing the warp sync market
     * @param _market The market to finalize. (Should be a warp sync market)
     */
    function getFinalizationReward(IMarket _market) public view returns (uint256) {
        return getRepReward(_market.getDisputeWindow().getEndTime());
    }

    /**
     * @notice Get the REP reward for initializing a universe and creating the Warp Sync Market
     * @param _universe the universe to initialize
     */
    function getCreationReward(IUniverse _universe) public view returns (uint256) {
        return getRepReward(_universe.creationTime());
    }

    function getRepReward(uint256 _theoreticalTime) private view returns (uint256) {
        uint256 _currentTime = augur.getTimestamp();
        uint256 _timeSinceTheoreticalCreationInSeconds = _currentTime.sub(_theoreticalTime);
        // Cannot overflow in any reasonable timeline of the universe
        return (_timeSinceTheoreticalCreationInSeconds ** 3).mul(1000);
    }

    function awardRep(IUniverse _universe, uint256 _amount) private returns (bool) {
        IV2ReputationToken _reputationToken = _universe.getReputationToken();
        // Whoever was responsible for this tx occuring gets REP.
        // solium-disable-next-line security/no-tx-origin
        _reputationToken.mintForWarpSync(_amount, tx.origin);
        return true;
    }

    function createMarket(IUniverse _universe) private {
        IV2ReputationToken _reputationToken = _universe.getReputationToken();
        uint256 _repBond = _universe.getOrCacheMarketRepBond();
        _reputationToken.mintForWarpSync(_repBond, address(this));
        uint256 _endTime = augur.getTimestamp().add(MARKET_LENGTH);
        IMarket _market = _universe.createScalarMarket(_endTime, 0, IAffiliateValidator(0), 0, address(this), PRICES, MAX_NUM_TICKS, EXTRA_INFO);
        markets[address(_universe)] = address(_market);
    }
}
