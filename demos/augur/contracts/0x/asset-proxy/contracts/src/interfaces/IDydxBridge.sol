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


interface IDydxBridge {

    /// @dev This is the subset of `IDydx.ActionType` that are supported by the bridge.
    enum BridgeActionType {
        Deposit,                    // Deposit tokens into dydx account.
        Withdraw                    // Withdraw tokens from dydx account.
    }

    struct BridgeAction {
        BridgeActionType actionType;            // Action to run on dydx account.
        uint256 accountId;                      // Index in `BridgeData.accountNumbers` for this action.
        uint256 marketId;                       // Market to operate on.
        uint256 conversionRateNumerator;        // Optional. If set, transfer amount is scaled by (conversionRateNumerator/conversionRateDenominator).
        uint256 conversionRateDenominator;      // Optional. If set, transfer amount is scaled by (conversionRateNumerator/conversionRateDenominator).
    }

    struct BridgeData {
        uint256[] accountNumbers;               // Account number used to identify the owner's specific account.
        BridgeAction[] actions;                 // Actions to carry out on the owner's accounts.
    }
}