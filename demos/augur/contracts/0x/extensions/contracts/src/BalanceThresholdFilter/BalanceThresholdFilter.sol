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
pragma experimental ABIEncoderV2;

import "../../../../exchange/contracts/src/interfaces/IExchange.sol";
import "./interfaces/IThresholdAsset.sol";
import "./MixinBalanceThresholdFilterCore.sol";


contract BalanceThresholdFilter is
    MixinBalanceThresholdFilterCore
{

    /// @dev Constructs BalanceThresholdFilter.
    /// @param exchange Address of 0x exchange.
    /// @param thresholdAsset The asset that must be held by makers/takers.
    /// @param balanceThreshold The minimum balance of `thresholdAsset` that must be held by makers/takers.
    constructor (
        address exchange,
        address thresholdAsset,
        uint256 balanceThreshold
    )
        public
    {
        EXCHANGE = IExchange(exchange);
        THRESHOLD_ASSET = IThresholdAsset(thresholdAsset);
        BALANCE_THRESHOLD = balanceThreshold;
    }
}
