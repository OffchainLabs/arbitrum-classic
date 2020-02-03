pragma solidity 0.5.15;

import 'ROOT/libraries/CloneFactory.sol';
import 'ROOT/reporting/IInitialReporter.sol';
import 'ROOT/reporting/IMarket.sol';
import 'ROOT/IAugur.sol';


/**
 * @title Initial Reporter Factory
 * @notice A Factory contract to create Initial Reporter delegate contracts
 * @dev Should not be used directly. Only intended to be used by Market contracts
 */
contract InitialReporterFactory is CloneFactory {
    function createInitialReporter(IAugur _augur, address _designatedReporter) public returns (IInitialReporter) {
        IMarket _market = IMarket(msg.sender);
        IInitialReporter _initialReporter = IInitialReporter(createClone(_augur.lookup("InitialReporter")));
        _initialReporter.initialize(_augur, _market, _designatedReporter);
        return _initialReporter;
    }
}
