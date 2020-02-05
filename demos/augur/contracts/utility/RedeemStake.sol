pragma solidity 0.5.15;

import '../reporting/IDisputeWindow.sol';
import '../reporting/IReportingParticipant.sol';


/**
 * @title Redeem Stake
 * @notice A Utility contract with no state for redeeming many redeemable bonds or tokens into the system in on TX
 */
contract RedeemStake {

    /**
     * @notice Redeems stake for multiple dispute bonds or Participation Tokens
     * @param _reportingParticipants Winning Initial Reporter or Dispute Crowdsourcer bonds the msg sender has stake in
     * @param _disputeWindows Dispute Windows (Participation Tokens) the msg sender has tokens for
     * @return Bool True
     */
    function redeemStake(IReportingParticipant[] memory _reportingParticipants, IDisputeWindow[] memory _disputeWindows) public returns (bool) {
        for (uint256 i=0; i < _reportingParticipants.length; i++) {
            _reportingParticipants[i].redeem(msg.sender);
        }
        for (uint256 i=0; i < _disputeWindows.length; i++) {
            _disputeWindows[i].redeem(msg.sender);
        }
        return true;
    }
}