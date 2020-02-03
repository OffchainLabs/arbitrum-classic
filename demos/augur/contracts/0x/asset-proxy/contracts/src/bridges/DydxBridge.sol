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

import "../../../../utils/contracts/src/DeploymentConstants.sol";
import "../../../../utils/contracts/src/LibSafeMath.sol";
import "../../../../exchange-libs/contracts/src/LibMath.sol";
import "../interfaces/IERC20Bridge.sol";
import "../interfaces/IDydxBridge.sol";
import "../interfaces/IDydx.sol";


contract DydxBridge is
    IERC20Bridge,
    IDydxBridge,
    DeploymentConstants
{

    using LibSafeMath for uint256;

    /// @dev Callback for `IERC20Bridge`. Deposits or withdraws tokens from a dydx account.
    ///      Notes:
    ///         1. This bridge must be set as an operator of the input dydx account.
    ///         2. This function may only be called in the context of the 0x Exchange.
    ///         3. The maker or taker of the 0x order must be the dydx account owner.
    ///         4. Deposits into dydx are made from the `from` address.
    ///         5. Withdrawals from dydx are made to the `to` address.
    ///         6. Calling this function must always withdraw at least `amount`,
    ///            otherwise the `ERC20Bridge` will revert.
    /// @param from The sender of the tokens and owner of the dydx account.
    /// @param to The recipient of the tokens.
    /// @param amount Minimum amount of `toTokenAddress` tokens to deposit or withdraw.
    /// @param encodedBridgeData An abi-encoded `BridgeData` struct.
    /// @return success The magic bytes if successful.
    function bridgeTransferFrom(
        address,
        address from,
        address to,
        uint256 amount,
        bytes calldata encodedBridgeData
    )
        external
        returns (bytes4 success)
    {
        // Ensure that only the `ERC20BridgeProxy` can call this function.
        require(
            msg.sender == _getERC20BridgeProxyAddress(),
            "DydxBridge/ONLY_CALLABLE_BY_ERC20_BRIDGE_PROXY"
        );

        // Decode bridge data.
        (BridgeData memory bridgeData) = abi.decode(encodedBridgeData, (BridgeData));

        // The dydx accounts are owned by the `from` address.
        IDydx.AccountInfo[] memory accounts = _createAccounts(from, bridgeData);

        // Create dydx actions to run on the dydx accounts.
        IDydx.ActionArgs[] memory actions = _createActions(
            from,
            to,
            amount,
            bridgeData
        );

        // Run operation. This will revert on failure.
        IDydx(_getDydxAddress()).operate(accounts, actions);
        return BRIDGE_SUCCESS;
    }

    /// @dev Creates an array of accounts for dydx to operate on.
    ///      All accounts must belong to the same owner.
    /// @param accountOwner Owner of the dydx account.
    /// @param bridgeData A `BridgeData` struct.
    function _createAccounts(
        address accountOwner,
        BridgeData memory bridgeData
    )
        internal
        returns (IDydx.AccountInfo[] memory accounts)
    {
        uint256[] memory accountNumbers = bridgeData.accountNumbers;
        uint256 nAccounts = accountNumbers.length;
        accounts = new IDydx.AccountInfo[](nAccounts);
        for (uint256 i = 0; i < nAccounts; ++i) {
            accounts[i] = IDydx.AccountInfo({
                owner: accountOwner,
                number: accountNumbers[i]
            });
        }
    }

    /// @dev Creates an array of actions to carry out on dydx.
    /// @param depositFrom Deposit value from this address (owner of the dydx account).
    /// @param withdrawTo Withdraw value to this address.
    /// @param amount The amount of value available to operate on.
    /// @param bridgeData A `BridgeData` struct.
    function _createActions(
        address depositFrom,
        address withdrawTo,
        uint256 amount,
        BridgeData memory bridgeData
    )
        internal
        returns (IDydx.ActionArgs[] memory actions)
    {
        BridgeAction[] memory bridgeActions = bridgeData.actions;
        uint256 nBridgeActions = bridgeActions.length;
        actions = new IDydx.ActionArgs[](nBridgeActions);
        for (uint256 i = 0; i < nBridgeActions; ++i) {
            // Cache current bridge action.
            BridgeAction memory bridgeAction = bridgeActions[i];

            // Scale amount, if conversion rate is set.
            uint256 scaledAmount;
            if (bridgeAction.conversionRateDenominator > 0) {
                scaledAmount = LibMath.safeGetPartialAmountFloor(
                    bridgeAction.conversionRateNumerator,
                    bridgeAction.conversionRateDenominator,
                    amount
                );
            } else {
                scaledAmount = amount;
            }

            // Construct dydx action.
            if (bridgeAction.actionType == BridgeActionType.Deposit) {
                // Deposit tokens from the account owner into their dydx account.
                actions[i] = _createDepositAction(
                    depositFrom,
                    scaledAmount,
                    bridgeAction
                );
            } else if (bridgeAction.actionType == BridgeActionType.Withdraw) {
                // Withdraw tokens from dydx to the `otherAccount`.
                actions[i] = _createWithdrawAction(
                    withdrawTo,
                    scaledAmount,
                    bridgeAction
                );
            } else {
                // If all values in the `Action` enum are handled then this
                // revert is unreachable: Solidity will revert when casting
                // from `uint8` to `Action`.
                revert("DydxBridge/UNRECOGNIZED_BRIDGE_ACTION");
            }
        }
    }

    /// @dev Returns a dydx `DepositAction`.
    /// @param depositFrom Deposit tokens from this address who is also the account owner.
    /// @param amount of tokens to deposit.
    /// @param bridgeAction A `BridgeAction` struct.
    /// @return depositAction The encoded dydx action.
    function _createDepositAction(
        address depositFrom,
        uint256 amount,
        BridgeAction memory bridgeAction
    )
        internal
        pure
        returns (
            IDydx.ActionArgs memory depositAction
        )
    {
        // Create dydx amount.
        IDydx.AssetAmount memory dydxAmount = IDydx.AssetAmount({
            sign: true,                                 // true if positive.
            denomination: IDydx.AssetDenomination.Wei,  // Wei => actual token amount held in account.
            ref: IDydx.AssetReference.Delta,                // Delta => a relative amount.
            value: amount                               // amount to deposit.
        });

        // Create dydx deposit action.
        depositAction = IDydx.ActionArgs({
            actionType: IDydx.ActionType.Deposit,           // deposit tokens.
            amount: dydxAmount,                             // amount to deposit.
            accountId: bridgeAction.accountId,              // index in the `accounts` when calling `operate`.
            primaryMarketId: bridgeAction.marketId,         // indicates which token to deposit.
            otherAddress: depositFrom,                      // deposit from the account owner.
            // unused parameters
            secondaryMarketId: 0,
            otherAccountId: 0,
            data: hex''
        });
    }

    /// @dev Returns a dydx `WithdrawAction`.
    /// @param withdrawTo Withdraw tokens to this address.
    /// @param amount of tokens to withdraw.
    /// @param bridgeAction A `BridgeAction` struct.
    /// @return withdrawAction The encoded dydx action.
    function _createWithdrawAction(
        address withdrawTo,
        uint256 amount,
        BridgeAction memory bridgeAction
    )
        internal
        pure
        returns (
            IDydx.ActionArgs memory withdrawAction
        )
    {
        // Create dydx amount.
        IDydx.AssetAmount memory amountToWithdraw = IDydx.AssetAmount({
            sign: false,                                    // false if negative.
            denomination: IDydx.AssetDenomination.Wei,      // Wei => actual token amount held in account.
            ref: IDydx.AssetReference.Delta,                // Delta => a relative amount.
            value: amount                                   // amount to withdraw.
        });

        // Create withdraw action.
        withdrawAction = IDydx.ActionArgs({
            actionType: IDydx.ActionType.Withdraw,          // withdraw tokens.
            amount: amountToWithdraw,                       // amount to withdraw.
            accountId: bridgeAction.accountId,              // index in the `accounts` when calling `operate`.
            primaryMarketId: bridgeAction.marketId,         // indicates which token to withdraw.
            otherAddress: withdrawTo,                       // withdraw tokens to this address.
            // unused parameters
            secondaryMarketId: 0,
            otherAccountId: 0,
            data: hex''
        });
    }
}
