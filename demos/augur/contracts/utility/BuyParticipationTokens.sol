pragma solidity 0.5.15;

import '../reporting/IDisputeWindow.sol';
import '../reporting/IUniverse.sol';


/**
 * @title BuyParticipationTokens
 * @notice A Utility contract with no state for purchasing participation tokens
 */
contract BuyParticipationTokens {

    /**
     * @notice Redeems stake for multiple dispute bonds or Participation Tokens
     * @param _universe The Universe to buy participation tokens in
     * @return Bool True
     */
    function buyParticipationTokens(IUniverse _universe, uint256 _attotokens) public returns (bool) {
        IDisputeWindow _disputeWindow = _universe.getOrCreateCurrentDisputeWindow(false);
        _disputeWindow.trustedBuy(msg.sender, _attotokens);
        return true;
    }
}