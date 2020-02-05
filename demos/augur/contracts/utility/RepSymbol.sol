pragma solidity 0.5.15;

import '../IAugur.sol';
import '../IAugurMarketDataGetter.sol';
import '../reporting/IUniverse.sol';
import '../libraries/math/UintToString.sol';
import '../libraries/BytesToString.sol';
import './IRepSymbol.sol';


/**
 * @title Rep Symbol
 * @notice A Utility contract used to retrieve the symbol for a REP token. We extract this out to reduce REP token contract size.
 */
contract RepSymbol is IRepSymbol {
    using UintToString for uint;
    using BytesToString for bytes32;

    function getRepSymbol(address _augurAddress, address _universeAddress) external view returns (string memory) {
        IAugur _augur = IAugur(_augurAddress);
        IUniverse _universe = IUniverse(_universeAddress);
        IUniverse _parentUniverse = _universe.getParentUniverse();

        if (_parentUniverse != IUniverse(0)) {
            uint256 _forkIndex = _augur.getUniverseForkIndex(_parentUniverse);
            IMarket _forkingMarket = _parentUniverse.getForkingMarket();
            uint256 _numTicks = _forkingMarket.getNumTicks();
            uint256[] memory _payoutNumerators = _universe.getPayoutNumerators();
            if (_payoutNumerators[0] != 0) {
                return string(abi.encodePacked("REPv2", "_", _payoutNumerators[0] == _numTicks ? "INVALID" : "MALFORMED", "_", _forkIndex.uint2str()));
            }
            IMarket.MarketType _marketType = IAugurMarketDataGetter(_augurAddress).getMarketType(_forkingMarket);
            string memory _outcome = "YES";
            if (_marketType == IMarket.MarketType.YES_NO) {
                if (_payoutNumerators[1] == _numTicks) {
                    _outcome = "NO";
                } else if (_payoutNumerators[1] != _numTicks) {
                    _outcome = "MALFORMED";
                }
            } else if (_marketType == IMarket.MarketType.CATEGORICAL) {
                uint256 _numOutcomes = _forkingMarket.getNumberOfOutcomes();
                bytes32[] memory _outcomes = IAugurMarketDataGetter(_augurAddress).getMarketOutcomes(_forkingMarket);
                for (uint256 _i = 1; _i < _numOutcomes; _i++) {
                    if (_payoutNumerators[_i] == 0) {
                        continue;
                    }
                    _outcome = _payoutNumerators[_i] != _numTicks ? "MALFORMED" : _outcomes[_i - 1].bytes32ToString();
                    break;
                }
            } else {
                _outcome = _payoutNumerators[2].uint2str();
            }

            return string(abi.encodePacked("REPv2", "_", _outcome, "_", _forkIndex.uint2str()));
        }
        return "REPv2";
    }
}
