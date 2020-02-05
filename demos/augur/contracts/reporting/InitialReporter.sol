pragma solidity 0.5.15;

import '../libraries/Initializable.sol';
import './IInitialReporter.sol';
import './IMarket.sol';
import './BaseReportingParticipant.sol';
import '../libraries/Ownable.sol';
import '../IAugur.sol';


/**
 * @title Initial Reporter
 * @notice The bond used to encapsulate the initial report for a Market
 */
contract InitialReporter is Ownable, BaseReportingParticipant, Initializable, IInitialReporter {
    address private designatedReporter;
    address private actualReporter;
    uint256 private reportTimestamp;

    function initialize(IAugur _augur, IMarket _market, address _designatedReporter) public beforeInitialized {
        endInitialization();
        augur = _augur;
        market = _market;
        reputationToken = market.getUniverse().getReputationToken();
        designatedReporter = _designatedReporter;
    }

    /**
     * @notice Redeems ownership of this bond for the provided redeemer in exchange for owed REP
     * @dev The address argument is ignored. There is only ever one owner of this bond, but the signature needs to match Dispute Crowdsourcer's redeem for code simplicity
     * @return bool True
     */
    function redeem(address) public returns (bool) {
        bool _isDisavowed = isDisavowed();
        if (!_isDisavowed && !market.isFinalized()) {
            market.finalize();
        }
        uint256 _repBalance = reputationToken.balanceOf(address(this));
        require(reputationToken.transfer(owner, _repBalance));
        if (!_isDisavowed) {
            augur.logInitialReporterRedeemed(market.getUniverse(), owner, address(market), size, _repBalance, payoutNumerators);
        }
        return true;
    }

    function report(address _reporter, bytes32 _payoutDistributionHash, uint256[] memory _payoutNumerators, uint256 _initialReportStake) public {
        require(IMarket(msg.sender) == market);
        require(reportTimestamp == 0, "InitialReporter.report: Report has already been placed");
        uint256 _timestamp = augur.getTimestamp();
        bool _isDesignatedReporter = _reporter == getDesignatedReporter();
        bool _designatedReportingExpired = _timestamp > market.getDesignatedReportingEndTime();
        require(_designatedReportingExpired || _isDesignatedReporter, "InitialReporter.report: Reporting time not started");
        actualReporter = _reporter;
        owner = _reporter;
        payoutDistributionHash = _payoutDistributionHash;
        reportTimestamp = _timestamp;
        payoutNumerators = _payoutNumerators;
        size = _initialReportStake;
    }

    function returnRepFromDisavow() public {
        require(IMarket(msg.sender) == market);
        require(reputationToken.transfer(owner, reputationToken.balanceOf(address(this))));
        reportTimestamp = 0;
    }

    function migrateToNewUniverse(address _designatedReporter) public {
        require(IMarket(msg.sender) == market);
        designatedReporter = _designatedReporter;
        reputationToken = market.getUniverse().getReputationToken();
    }

    /**
     * @notice Used in the event of the market forking. First forks this bond into the appropriate child universe and then redeems there.
     * @return bool True
     */
    function forkAndRedeem() public returns (bool) {
        if (!isDisavowed()) {
            augur.logInitialReporterRedeemed(market.getUniverse(), owner, address(market), size, reputationToken.balanceOf(address(this)), payoutNumerators);
        }
        fork();
        redeem(msg.sender);
        return true;
    }

    /**
     * @return The amount of REP currently staked in this bond
     */
    function getStake() public view returns (uint256) {
        return size;
    }

    /**
     * @return The designated reporter for this market / bond
     */
    function getDesignatedReporter() public view returns (address) {
        return designatedReporter;
    }

    /**
     * @return When the actual report was made if one was made
     */
    function getReportTimestamp() public view returns (uint256) {
        return reportTimestamp;
    }

    /**
     * @return Bool indicating if the report was made by the Designated Reporter
     */
    function designatedReporterShowed() public view returns (bool) {
        return actualReporter == designatedReporter;
    }

    /**
     * @return The REP token associated with this bond
     */
    function getReputationToken() public view returns (IReputationToken) {
        return reputationToken;
    }

    /**
     * @return Bool indicating if the report was ultimately the finalzied payout
     */
    function initialReporterWasCorrect() public view returns (bool) {
        return payoutDistributionHash != bytes32(0) && payoutDistributionHash == market.getWinningPayoutDistributionHash();
    }

    function onTransferOwnership(address _owner, address _newOwner) internal {
        augur.logInitialReporterTransferred(market.getUniverse(), market, _owner, _newOwner);
    }
}
