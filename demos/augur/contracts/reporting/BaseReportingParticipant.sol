pragma solidity 0.5.15;

import './IReportingParticipant.sol';
import './IMarket.sol';
import './IDisputeWindow.sol';
import './IReputationToken.sol';


contract BaseReportingParticipant is IReportingParticipant {
    IMarket internal market;
    uint256 internal size;
    bytes32 internal payoutDistributionHash;
    uint256[] internal payoutNumerators;
    IReputationToken internal reputationToken;
    IAugur public augur;

    function liquidateLosing() public {
        require(IMarket(msg.sender) == market);
        require(market.getWinningPayoutDistributionHash() != getPayoutDistributionHash() && market.getWinningPayoutDistributionHash() != bytes32(0));
        IReputationToken _reputationToken = market.getReputationToken();
        require(_reputationToken.transfer(address(market), _reputationToken.balanceOf(address(this))));
    }

    function fork() public {
        require(market == market.getUniverse().getForkingMarket());
        IUniverse _newUniverse = market.getUniverse().createChildUniverse(payoutNumerators);
        IReputationToken _newReputationToken = _newUniverse.getReputationToken();
        uint256 _balance = reputationToken.balanceOf(address(this));
        reputationToken.migrateOutByPayout(_newUniverse.getPayoutNumerators(), _balance);
        _newReputationToken.mintForReportingParticipant(size);
        reputationToken = _newReputationToken;
        augur.logReportingParticipantDisavowed(market.getUniverse(), market);
        market = IMarket(0);
    }

    /**
     * @notice Get the size of the bond. This is the amount of REP needed to fill this bond.
     * @return The size of the bond.
     */
    function getSize() public view returns (uint256) {
        return size;
    }

    /**
     * @notice Get the payout distribution hash
     * @return The payout distribution hash
     */
    function getPayoutDistributionHash() public view returns (bytes32) {
        return payoutDistributionHash;
    }

    /**
     * @notice Get the market associated with this bond
     * @return The market associated with this bond
     */
    function getMarket() public view returns (IMarket) {
        return market;
    }

    /**
     * @notice Get bool indicating if the bond is disavowed. Disavowal occurs if a bond is not filled or if a fork occurs.
     * @return Bool indicating if the bond is disavowed
     */
    function isDisavowed() public view returns (bool) {
        return market == IMarket(0) || !market.isContainerForReportingParticipant(this);
    }

    /**
     * @notice Get the payout for a particular outcome in this bonds stated payout.
     * @return The uint256 payout for a particular outcome for this bond.
     */
    function getPayoutNumerator(uint256 _outcome) public view returns (uint256) {
        return payoutNumerators[_outcome];
    }

    /**
     * @notice Get the payout numerators for this bond.
     * @return The payout numerators for this bond.
     */
    function getPayoutNumerators() public view returns (uint256[] memory) {
        return payoutNumerators;
    }
}
