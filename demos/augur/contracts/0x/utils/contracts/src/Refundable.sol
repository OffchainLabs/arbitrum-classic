/*

  Copyright 2019 ZeroEx Intl.

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

*/

pragma solidity 0.5.15;

import "../../../utils/contracts/src/ReentrancyGuard.sol";


contract Refundable is
    ReentrancyGuard
{

    // This bool is used by the refund modifier to allow for lazily evaluated refunds.
    bool internal _shouldNotRefund;

    modifier refundFinalBalance {
        _;
        _refundNonZeroBalanceIfEnabled();
    }

    modifier refundFinalBalanceNoReentry {
        _lockMutexOrThrowIfAlreadyLocked();
        _;
        _refundNonZeroBalanceIfEnabled();
        _unlockMutex();
    }

    modifier disableRefundUntilEnd {
        if (_areRefundsDisabled()) {
            _;
        } else {
            _disableRefund();
            _;
            _enableAndRefundNonZeroBalance();
        }
    }

    function _refundNonZeroBalanceIfEnabled()
        internal
    {
        if (!_areRefundsDisabled()) {
            _refundNonZeroBalance();
        }
    }

    function _refundNonZeroBalance()
        internal
    {
        uint256 balance = address(this).balance;
        if (balance > 0) {
            msg.sender.transfer(balance);
        }
    }

    function _disableRefund()
        internal
    {
        _shouldNotRefund = true;
    }

    function _enableAndRefundNonZeroBalance()
        internal
    {
        _shouldNotRefund = false;
        _refundNonZeroBalance();
    }

    function _areRefundsDisabled()
        internal
        view
        returns (bool)
    {
        return _shouldNotRefund;
    }
}
