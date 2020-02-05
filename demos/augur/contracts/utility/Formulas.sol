pragma solidity 0.5.15;

import '../reporting/IDisputeWindow.sol';
import '../libraries/math/SafeMathUint256.sol';
import '../reporting/Reporting.sol';


/**
 * @title Formulas
 * @notice A Utility contract with no state meant to house the logic for large formulas which don't fit into already large contracts
 */
contract Formulas {
    using SafeMathUint256 for uint256;

    function calculateValidityBond(IDisputeWindow  _previousDisputeWindow, uint256 _previousValidityBondInAttoCash) public view returns (uint256) {
        uint256 _totalValidityBondsInPreviousWindow = _previousDisputeWindow.validityBondTotal();
        uint256 _invalidBondsInPreviousWindow = _previousDisputeWindow.invalidMarketsTotal();
        return calculateFloatingValue(_invalidBondsInPreviousWindow, _totalValidityBondsInPreviousWindow, Reporting.getTargetInvalidMarketsDivisor(), _previousValidityBondInAttoCash, Reporting.getValidityBondFloor());
    }

    function calculateDesignatedReportStake(IDisputeWindow  _previousDisputeWindow, uint256 _previousDesignatedReportStakeInAttoRep, uint256 _initialReportMinValue) public view returns (uint256) {
        uint256 _totalInitialReportBondsInPreviousWindow = _previousDisputeWindow.initialReportBondTotal();
        uint256 _incorrectDesignatedReportBondsInPreviousWindow = _previousDisputeWindow.incorrectDesignatedReportTotal();
        return calculateFloatingValue(_incorrectDesignatedReportBondsInPreviousWindow, _totalInitialReportBondsInPreviousWindow, Reporting.getTargetIncorrectDesignatedReportMarketsDivisor(), _previousDesignatedReportStakeInAttoRep, _initialReportMinValue);
    }

    function calculateDesignatedReportNoShowBond(IDisputeWindow  _previousDisputeWindow, uint256 _previousDesignatedReportNoShowBondInAttoRep, uint256 _initialReportMinValue) public view returns (uint256) {
        uint256 _totalNoShowBondsInPreviousWindow = _previousDisputeWindow.designatedReporterNoShowBondTotal();
        uint256 _designatedReportNoShowBondsInPreviousWindow = _previousDisputeWindow.designatedReportNoShowsTotal();
        return calculateFloatingValue(_designatedReportNoShowBondsInPreviousWindow, _totalNoShowBondsInPreviousWindow, Reporting.getTargetDesignatedReportNoShowsDivisor(), _previousDesignatedReportNoShowBondInAttoRep, _initialReportMinValue);
    }

    function calculateFloatingValue(uint256 _totalBad, uint256 _total, uint256 _targetDivisor, uint256 _previousValue, uint256 _floor) public pure returns (uint256 _newValue) {
        if (_total == 0) {
            return _previousValue;
        }

        // Modify the amount based on the previous amount and the number of markets fitting the failure criteria. We want the amount to be somewhere in the range of 0.9 to 2 times its previous value where ALL markets with the condition results in 2x and 0 results in 0.9x.
        // Safe math div is redundant so we avoid here as we're at the stack limit.
        if (_totalBad <= _total / _targetDivisor) {
            // FXP formula: previous_amount * (actual_percent / (10 * target_percent) + 0.9);
            _newValue = _totalBad
                .mul(_previousValue)
                .mul(_targetDivisor);
            _newValue = _newValue / _total;
            _newValue = _newValue / 10;
            _newValue = _newValue.add(_previousValue * 9 / 10);
        } else {
            // FXP formula: previous_amount * ((1/(1 - target_percent)) * (actual_percent - target_percent) + 1);
            _newValue = _targetDivisor
                .mul(_previousValue
                    .mul(_totalBad)
                    .div(_total)
                .sub(_previousValue / _targetDivisor));
            _newValue = _newValue / (_targetDivisor - 1);
            _newValue = _newValue.add(_previousValue);
        }
        _newValue = _newValue.max(_floor);

        return _newValue;
    }
}
