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

import "./libs/LibConstants.sol";
import "../../../../exchange-libs/contracts/src/LibOrder.sol";
import "../../../../exchange-libs/contracts/src/LibFillResults.sol";
import "../../../../utils/contracts/src/Ownable.sol";


contract MixinMatchOrders is
    Ownable,
    LibConstants
{
    /// @dev Match two complementary orders that have a profitable spread.
    ///      Each order is filled at their respective price point. However, the calculations are
    ///      carried out as though the orders are both being filled at the right order's price point.
    ///      The profit made by the left order is then used to fill the right order as much as possible.
    ///      This results in a spread being taken in terms of both assets. The spread is held within this contract.
    /// @param leftOrder First order to match.
    /// @param rightOrder Second order to match.
    /// @param leftSignature Proof that order was created by the left maker.
    /// @param rightSignature Proof that order was created by the right maker.
    function matchOrders(
        LibOrder.Order memory leftOrder,
        LibOrder.Order memory rightOrder,
        bytes memory leftSignature,
        bytes memory rightSignature
    )
        public
        onlyOwner
    {
        // Match orders, maximally filling `leftOrder`
        LibFillResults.MatchedFillResults memory matchedFillResults = EXCHANGE.matchOrders(
            leftOrder,
            rightOrder,
            leftSignature,
            rightSignature
        );

        uint256 leftMakerAssetSpreadAmount = matchedFillResults.left.makerAssetSpreadAmount;
        uint256 rightOrderTakerAssetAmount = rightOrder.takerAssetAmount;

        // Do not attempt to call `fillOrder` if no spread was taken or `rightOrder` has been completely filled
        if (leftMakerAssetSpreadAmount == 0 || matchedFillResults.right.takerAssetFilledAmount == rightOrderTakerAssetAmount) {
            return;
        }

        // The `assetData` fields of the `rightOrder` could have been null for the `matchOrders` call. We reassign them before calling `fillOrder`.
        rightOrder.makerAssetData = leftOrder.takerAssetData;
        rightOrder.takerAssetData = leftOrder.makerAssetData;

        // Query `rightOrder` info to check if it has been completely filled
        // We need to make this check in case the `rightOrder` was partially filled before the `matchOrders` call
        LibOrder.OrderInfo memory orderInfo = EXCHANGE.getOrderInfo(rightOrder);

        // Do not attempt to call `fillOrder` if order has been completely filled
        if (orderInfo.orderTakerAssetFilledAmount == rightOrderTakerAssetAmount) {
            return;
        }

        // We do not need to pass in a signature since it was already validated in the `matchOrders` call
        EXCHANGE.fillOrder(
            rightOrder,
            leftMakerAssetSpreadAmount,
            ""
        );
    }
}
