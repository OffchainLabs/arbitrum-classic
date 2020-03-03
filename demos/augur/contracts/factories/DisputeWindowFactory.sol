pragma solidity 0.5.15;

import '../libraries/CloneFactory.sol';
import '../IAugur.sol';
import '../reporting/IUniverse.sol';
import '../reporting/IDisputeWindow.sol';
import './IDisputeWindowFactory.sol';
import '../reporting/Reporting.sol';


/**
 * @title Dispute Window Factory
 * @notice A Factory contract to create Dispute Window delegate contracts
 * @dev Cannot be used directly. Only called by Universe contracts
 */
contract DisputeWindowFactory is CloneFactory, IDisputeWindowFactory {
    function createDisputeWindow(IAugur _augur, uint256 _disputeWindowId, uint256 _windowDuration, uint256 _startTime, bool _participationTokensEnabled) public returns (IDisputeWindow) {
        IUniverse _universe = IUniverse(msg.sender);
        // require(_augur.isKnownUniverse(_universe), "DisputeWindowFactory: Universe specified is unrecognized by Augur");

        address newContractAddress = createNewContract();
        IDisputeWindow _disputeWindow = IDisputeWindow(newContractAddress);
        _disputeWindow.initialize(_augur, _universe, _disputeWindowId, _participationTokensEnabled, _windowDuration, _startTime);
        
        return _disputeWindow;
    }
}
