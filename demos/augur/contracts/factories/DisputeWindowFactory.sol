pragma solidity 0.5.15;

import 'ROOT/libraries/CloneFactory.sol';
import 'ROOT/IAugur.sol';
import 'ROOT/reporting/IUniverse.sol';
import 'ROOT/reporting/IDisputeWindow.sol';
import 'ROOT/factories/IDisputeWindowFactory.sol';
import 'ROOT/reporting/Reporting.sol';


/**
 * @title Dispute Window Factory
 * @notice A Factory contract to create Dispute Window delegate contracts
 * @dev Cannot be used directly. Only called by Universe contracts
 */
contract DisputeWindowFactory is CloneFactory, IDisputeWindowFactory {
    function createDisputeWindow(IAugur _augur, uint256 _disputeWindowId, uint256 _windowDuration, uint256 _startTime, bool _participationTokensEnabled) public returns (IDisputeWindow) {
        IUniverse _universe = IUniverse(msg.sender);
        require(_augur.isKnownUniverse(_universe), "DisputeWindowFactory: Universe specified is unrecognized by Augur");
        IDisputeWindow _disputeWindow = IDisputeWindow(createClone(_augur.lookup("DisputeWindow")));
        _disputeWindow.initialize(_augur, _universe, _disputeWindowId, _participationTokensEnabled, _windowDuration, _startTime);
        return _disputeWindow;
    }
}
