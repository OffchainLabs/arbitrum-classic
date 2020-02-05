pragma solidity 0.5.15;

import './ICash.sol';
import './external/IDaiVat.sol';
import './libraries/math/SafeMathUint256.sol';


contract CashSender {
    using SafeMathUint256 for uint256;

    ICash public cash;
    IDaiVat public vat;

    uint constant RAY = 10 ** 27;

    function initializeCashSender(address _vat, address _cash) internal {
        vat = IDaiVat(_vat);
        cash = ICash(_cash);
    }

    function cashBalance(address _account) public view returns (uint256 _balance) {
        _balance = cash.balanceOf(_account);
        if (vat.live() == 0) {
            _balance += vatDaiToDai(vat.dai(_account));
        }
        return _balance;
    }

    function cashAvailableForTransferFrom(address _owner, address _sender) public view returns (uint256 _available) {
        uint256 _balance = cash.balanceOf(_owner);
        uint256 _allowance = cash.allowance(_owner, _sender);
        _available = _balance.min(_allowance);
        if (vat.live() == 0 && (vat.can(_owner, _sender) == 1)) {
            _available += vatDaiToDai(vat.dai(_owner));
        }
        return _available;
    }

    function cashApprove(address _spender, uint256 _amount) internal {
        cash.approve(_spender, _amount);
        if (vat.live() == 0) {
            vat.hope(_spender);
        }
    }

    function cashTransfer(address _to, uint256 _amount) internal {
        address _from = address(this);
        if (vat.live() == 0) {
            _amount = shutdownTransfer(_from, _to, _amount);
            if (_amount == 0) {
                return;
            }
        }
        require(cash.transfer(_to, _amount));
    }

    function cashTransferFrom(address _from, address _to, uint256 _amount) internal {
        if (vat.live() == 0) {
            _amount = shutdownTransfer(_from, _to, _amount);
            if (_amount == 0) {
                return;
            }
        }
        require(cash.transferFrom(_from, _to, _amount));
    }

    function shutdownTransfer(address _from, address _to, uint256 _amount) internal returns (uint256) {
        if (cash.balanceOf(_from) < _amount) {
            uint256 _vDaiToTransfer = vat.dai(_from).min(daiToVatDai(_amount));
            vat.move(_from, _to, _vDaiToTransfer);
            _amount -= vatDaiToDai(_vDaiToTransfer);
        }
        return _amount;
    }

    function vatDaiToDai(uint256 _vDaiAmount) public pure returns (uint256) {
        return _vDaiAmount.div(RAY);
    }

    function daiToVatDai(uint256 _daiAmount) public pure returns (uint256) {
        return _daiAmount.mul(RAY);
    }
}