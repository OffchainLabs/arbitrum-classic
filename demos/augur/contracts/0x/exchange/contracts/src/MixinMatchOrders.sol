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

import "../../../utils/contracts/src/LibBytes.sol";
import "../../../exchange-libs/contracts/src/LibOrder.sol";
import "../../../exchange-libs/contracts/src/LibFillResults.sol";
import "../../../exchange-libs/contracts/src/LibExchangeRichErrors.sol";
import "./interfaces/IMatchOrders.sol";
import "./MixinExchangeCore.sol";


contract MixinMatchOrders is
    MixinExchangeCore,
    IMatchOrders
{
    using LibBytes for bytes;
    using LibSafeMath for uint256;
    using LibOrder for LibOrder.Order;

    /// @dev Match complementary orders that have a profitable spread.
    ///      Each order is filled at their respective price point, and
    ///      the matcher receives a profit denominated in the left maker asset.
    /// @param leftOrders Set of orders with the same maker / taker asset.
    /// @param rightOrders Set of orders to match against `leftOrders`
    /// @param leftSignatures Proof that left orders were created by the left makers.
    /// @param rightSignatures Proof that right orders were created by the right makers.
    /// @return batchMatchedFillResults Amounts filled and profit generated.
    function batchMatchOrders(
        LibOrder.Order[] memory leftOrders,
        LibOrder.Order[] memory rightOrders,
        bytes[] memory leftSignatures,
        bytes[] memory rightSignatures
    )
        public
        payable
        refundFinalBalanceNoReentry
        returns (LibFillResults.BatchMatchedFillResults memory batchMatchedFillResults)
    {
        return _batchMatchOrders(
            leftOrders,
            rightOrders,
            leftSignatures,
            rightSignatures,
            false
        );
    }

    /// @dev Match complementary orders that have a profitable spread.
    ///      Each order is maximally filled at their respective price point, and
    ///      the matcher receives a profit denominated in either the left maker asset,
    ///      right maker asset, or a combination of both.
    /// @param leftOrders Set of orders with the same maker / taker asset.
    /// @param rightOrders Set of orders to match against `leftOrders`
    /// @param leftSignatures Proof that left orders were created by the left makers.
    /// @param rightSignatures Proof that right orders were created by the right makers.
    /// @return batchMatchedFillResults Amounts filled and profit generated.
    function batchMatchOrdersWithMaximalFill(
        LibOrder.Order[] memory leftOrders,
        LibOrder.Order[] memory rightOrders,
        bytes[] memory leftSignatures,
        bytes[] memory rightSignatures
    )
        public
        payable
        refundFinalBalanceNoReentry
        returns (LibFillResults.BatchMatchedFillResults memory batchMatchedFillResults)
    {
        return _batchMatchOrders(
            leftOrders,
            rightOrders,
            leftSignatures,
            rightSignatures,
            true
        );
    }

    /// @dev Match two complementary orders that have a profitable spread.
    ///      Each order is filled at their respective price point. However, the calculations are
    ///      carried out as though the orders are both being filled at the right order's price point.
    ///      The profit made by the left order goes to the taker (who matched the two orders).
    /// @param leftOrder First order to match.
    /// @param rightOrder Second order to match.
    /// @param leftSignature Proof that order was created by the left maker.
    /// @param rightSignature Proof that order was created by the right maker.
    /// @return matchedFillResults Amounts filled and fees paid by maker and taker of matched orders.
    function matchOrders(
        LibOrder.Order memory leftOrder,
        LibOrder.Order memory rightOrder,
        bytes memory leftSignature,
        bytes memory rightSignature
    )
        public
        payable
        refundFinalBalanceNoReentry
        returns (LibFillResults.MatchedFillResults memory matchedFillResults)
    {
        return _matchOrders(
            leftOrder,
            rightOrder,
            leftSignature,
            rightSignature,
            false
        );
    }

    /// @dev Match two complementary orders that have a profitable spread.
    ///      Each order is maximally filled at their respective price point, and
    ///      the matcher receives a profit denominated in either the left maker asset,
    ///      right maker asset, or a combination of both.
    /// @param leftOrder First order to match.
    /// @param rightOrder Second order to match.
    /// @param leftSignature Proof that order was created by the left maker.
    /// @param rightSignature Proof that order was created by the right maker.
    /// @return matchedFillResults Amounts filled by maker and taker of matched orders.
    function matchOrdersWithMaximalFill(
        LibOrder.Order memory leftOrder,
        LibOrder.Order memory rightOrder,
        bytes memory leftSignature,
        bytes memory rightSignature
    )
        public
        payable
        refundFinalBalanceNoReentry
        returns (LibFillResults.MatchedFillResults memory matchedFillResults)
    {
        return _matchOrders(
            leftOrder,
            rightOrder,
            leftSignature,
            rightSignature,
            true
        );
    }

    /// @dev Validates context for matchOrders. Succeeds or throws.
    /// @param leftOrder First order to match.
    /// @param rightOrder Second order to match.
    /// @param leftOrderHash First matched order hash.
    /// @param rightOrderHash Second matched order hash.
    function _assertValidMatch(
        LibOrder.Order memory leftOrder,
        LibOrder.Order memory rightOrder,
        bytes32 leftOrderHash,
        bytes32 rightOrderHash
    )
        internal
        pure
    {
        // Make sure there is a profitable spread.
        // There is a profitable spread iff the cost per unit bought (OrderA.MakerAmount/OrderA.TakerAmount) for each order is greater
        // than the profit per unit sold of the matched order (OrderB.TakerAmount/OrderB.MakerAmount).
        // This is satisfied by the equations below:
        // <leftOrder.makerAssetAmount> / <leftOrder.takerAssetAmount> >= <rightOrder.takerAssetAmount> / <rightOrder.makerAssetAmount>
        // AND
        // <rightOrder.makerAssetAmount> / <rightOrder.takerAssetAmount> >= <leftOrder.takerAssetAmount> / <leftOrder.makerAssetAmount>
        // These equations can be combined to get the following:
        if (leftOrder.makerAssetAmount.safeMul(rightOrder.makerAssetAmount) <
            leftOrder.takerAssetAmount.safeMul(rightOrder.takerAssetAmount)) {
            revert();
        }
    }

    /// @dev Match complementary orders that have a profitable spread.
    ///      Each order is filled at their respective price point, and
    ///      the matcher receives a profit denominated in the left maker asset.
    ///      This is the reentrant version of `batchMatchOrders` and `batchMatchOrdersWithMaximalFill`.
    /// @param leftOrders Set of orders with the same maker / taker asset.
    /// @param rightOrders Set of orders to match against `leftOrders`
    /// @param leftSignatures Proof that left orders were created by the left makers.
    /// @param rightSignatures Proof that right orders were created by the right makers.
    /// @param shouldMaximallyFillOrders A value that indicates whether or not the order matching
    ///                        should be done with maximal fill.
    /// @return batchMatchedFillResults Amounts filled and profit generated.
    function _batchMatchOrders(
        LibOrder.Order[] memory leftOrders,
        LibOrder.Order[] memory rightOrders,
        bytes[] memory leftSignatures,
        bytes[] memory rightSignatures,
        bool shouldMaximallyFillOrders
    )
        internal
        returns (LibFillResults.BatchMatchedFillResults memory batchMatchedFillResults)
    {
        // Ensure that the left and right orders have nonzero lengths.
        if (leftOrders.length == 0) {
            revert();
        }
        if (rightOrders.length == 0) {
            revert();
        }

        // Ensure that the left and right arrays are compatible.
        if (leftOrders.length != leftSignatures.length) {
            revert();
        }
        if (rightOrders.length != rightSignatures.length) {
            revert();
        }

        batchMatchedFillResults.left = new LibFillResults.FillResults[](leftOrders.length);
        batchMatchedFillResults.right = new LibFillResults.FillResults[](rightOrders.length);

        // Set up initial indices.
        uint256 leftIdx = 0;
        uint256 rightIdx = 0;

        // Keep local variables for orders, order filled amounts, and signatures for efficiency.
        LibOrder.Order memory leftOrder = leftOrders[0];
        LibOrder.Order memory rightOrder = rightOrders[0];
        (, uint256 leftOrderTakerAssetFilledAmount) = _getOrderHashAndFilledAmount(leftOrder);
        (, uint256 rightOrderTakerAssetFilledAmount) = _getOrderHashAndFilledAmount(rightOrder);
        LibFillResults.FillResults memory leftFillResults;
        LibFillResults.FillResults memory rightFillResults;

        // Loop infinitely (until broken inside of the loop), but keep a counter of how
        // many orders have been matched.
        for (;;) {
            // Match the two orders that are pointed to by the left and right indices
            LibFillResults.MatchedFillResults memory matchResults = _matchOrders(
                leftOrder,
                rightOrder,
                leftSignatures[leftIdx],
                rightSignatures[rightIdx],
                shouldMaximallyFillOrders
            );

            // Update the order filled amounts with the updated takerAssetFilledAmount
            leftOrderTakerAssetFilledAmount = leftOrderTakerAssetFilledAmount.safeAdd(matchResults.left.takerAssetFilledAmount);
            rightOrderTakerAssetFilledAmount = rightOrderTakerAssetFilledAmount.safeAdd(matchResults.right.takerAssetFilledAmount);

            // Aggregate the new fill results with the previous fill results for the current orders.
            leftFillResults = LibFillResults.addFillResults(
                leftFillResults,
                matchResults.left
            );
            rightFillResults = LibFillResults.addFillResults(
                rightFillResults,
                matchResults.right
            );

            // Update the profit in the left and right maker assets using the profits from
            // the match.
            batchMatchedFillResults.profitInLeftMakerAsset = batchMatchedFillResults.profitInLeftMakerAsset.safeAdd(
                matchResults.profitInLeftMakerAsset
            );
            batchMatchedFillResults.profitInRightMakerAsset = batchMatchedFillResults.profitInRightMakerAsset.safeAdd(
                matchResults.profitInRightMakerAsset
            );

            // If the leftOrder is filled, update the leftIdx, leftOrder, and leftSignature,
            // or break out of the loop if there are no more leftOrders to match.
            if (leftOrderTakerAssetFilledAmount >= leftOrder.takerAssetAmount) {
                // Update the batched fill results once the leftIdx is updated.
                batchMatchedFillResults.left[leftIdx++] = leftFillResults;
                // Clear the intermediate fill results value.
                leftFillResults = LibFillResults.FillResults(0, 0, 0, 0, 0);

                // If all of the left orders have been filled, break out of the loop.
                // Otherwise, update the current right order.
                if (leftIdx == leftOrders.length) {
                    // Update the right batched fill results
                    batchMatchedFillResults.right[rightIdx] = rightFillResults;
                    break;
                } else {
                    leftOrder = leftOrders[leftIdx];
                    (, leftOrderTakerAssetFilledAmount) = _getOrderHashAndFilledAmount(leftOrder);
                }
            }

            // If the rightOrder is filled, update the rightIdx, rightOrder, and rightSignature,
            // or break out of the loop if there are no more rightOrders to match.
            if (rightOrderTakerAssetFilledAmount >= rightOrder.takerAssetAmount) {
                // Update the batched fill results once the rightIdx is updated.
                batchMatchedFillResults.right[rightIdx++] = rightFillResults;
                // Clear the intermediate fill results value.
                rightFillResults = LibFillResults.FillResults(0, 0, 0, 0, 0);

                // If all of the right orders have been filled, break out of the loop.
                // Otherwise, update the current right order.
                if (rightIdx == rightOrders.length) {
                    // Update the left batched fill results
                    batchMatchedFillResults.left[leftIdx] = leftFillResults;
                    break;
                } else {
                    rightOrder = rightOrders[rightIdx];
                    (, rightOrderTakerAssetFilledAmount) = _getOrderHashAndFilledAmount(rightOrder);
                }
            }
        }

        // Return the fill results from the batch match
        return batchMatchedFillResults;
    }

    /// @dev Match two complementary orders that have a profitable spread.
    ///      Each order is filled at their respective price point. However, the calculations are
    ///      carried out as though the orders are both being filled at the right order's price point.
    ///      The profit made by the left order goes to the taker (who matched the two orders). This
    ///      function is needed to allow for reentrant order matching (used by `batchMatchOrders` and
    ///      `batchMatchOrdersWithMaximalFill`).
    /// @param leftOrder First order to match.
    /// @param rightOrder Second order to match.
    /// @param leftSignature Proof that order was created by the left maker.
    /// @param rightSignature Proof that order was created by the right maker.
    /// @param shouldMaximallyFillOrders Indicates whether or not the maximal fill matching strategy should be used
    /// @return matchedFillResults Amounts filled and fees paid by maker and taker of matched orders.
    function _matchOrders(
        LibOrder.Order memory leftOrder,
        LibOrder.Order memory rightOrder,
        bytes memory leftSignature,
        bytes memory rightSignature,
        bool shouldMaximallyFillOrders
    )
        internal
        returns (LibFillResults.MatchedFillResults memory matchedFillResults)
    {
        // We assume that rightOrder.takerAssetData == leftOrder.makerAssetData and rightOrder.makerAssetData == leftOrder.takerAssetData
        // by pointing these values to the same location in memory. This is cheaper than checking equality.
        // If this assumption isn't true, the match will fail at signature validation.
        rightOrder.makerAssetData = leftOrder.takerAssetData;
        rightOrder.takerAssetData = leftOrder.makerAssetData;

        // Get left & right order info
        LibOrder.OrderInfo memory leftOrderInfo = getOrderInfo(leftOrder);
        LibOrder.OrderInfo memory rightOrderInfo = getOrderInfo(rightOrder);

        // Fetch taker address
        address takerAddress = _getCurrentContextAddress();

        // Either our context is valid or we revert
        _assertFillableOrder(
            leftOrder,
            leftOrderInfo,
            takerAddress,
            leftSignature
        );
        _assertFillableOrder(
            rightOrder,
            rightOrderInfo,
            takerAddress,
            rightSignature
        );
        _assertValidMatch(
            leftOrder,
            rightOrder,
            leftOrderInfo.orderHash,
            rightOrderInfo.orderHash
        );

        // Compute proportional fill amounts
        matchedFillResults = LibFillResults.calculateMatchedFillResults(
            leftOrder,
            rightOrder,
            leftOrderInfo.orderTakerAssetFilledAmount,
            rightOrderInfo.orderTakerAssetFilledAmount,
            protocolFeeMultiplier,
            tx.gasprice,
            shouldMaximallyFillOrders
        );

        // Update exchange state
        _updateFilledState(
            leftOrder,
            takerAddress,
            leftOrderInfo.orderHash,
            leftOrderInfo.orderTakerAssetFilledAmount,
            matchedFillResults.left
        );
        _updateFilledState(
            rightOrder,
            takerAddress,
            rightOrderInfo.orderHash,
            rightOrderInfo.orderTakerAssetFilledAmount,
            matchedFillResults.right
        );

        // Settle matched orders. Succeeds or throws.
        _settleMatchedOrders(
            leftOrderInfo.orderHash,
            rightOrderInfo.orderHash,
            leftOrder,
            rightOrder,
            takerAddress,
            matchedFillResults
        );

        return matchedFillResults;
    }

    /// @dev Settles matched order by transferring appropriate funds between order makers, taker, and fee recipient.
    /// @param leftOrderHash First matched order hash.
    /// @param rightOrderHash Second matched order hash.
    /// @param leftOrder First matched order.
    /// @param rightOrder Second matched order.
    /// @param takerAddress Address that matched the orders. The taker receives the spread between orders as profit.
    /// @param matchedFillResults Struct holding amounts to transfer between makers, taker, and fee recipients.
    function _settleMatchedOrders(
        bytes32 leftOrderHash,
        bytes32 rightOrderHash,
        LibOrder.Order memory leftOrder,
        LibOrder.Order memory rightOrder,
        address takerAddress,
        LibFillResults.MatchedFillResults memory matchedFillResults
    )
        internal
    {
        address leftMakerAddress = leftOrder.makerAddress;
        address rightMakerAddress = rightOrder.makerAddress;
        address leftFeeRecipientAddress = leftOrder.feeRecipientAddress;
        address rightFeeRecipientAddress = rightOrder.feeRecipientAddress;

        // Right maker asset -> left maker
        _dispatchTransferFrom(
            rightOrderHash,
            rightOrder.makerAssetData,
            rightMakerAddress,
            leftMakerAddress,
            matchedFillResults.left.takerAssetFilledAmount
        );

        // Left maker asset -> right maker
        _dispatchTransferFrom(
            leftOrderHash,
            leftOrder.makerAssetData,
            leftMakerAddress,
            rightMakerAddress,
            matchedFillResults.right.takerAssetFilledAmount
        );

        // Right maker fee -> right fee recipient
        _dispatchTransferFrom(
            rightOrderHash,
            rightOrder.makerFeeAssetData,
            rightMakerAddress,
            rightFeeRecipientAddress,
            matchedFillResults.right.makerFeePaid
        );

        // Left maker fee -> left fee recipient
        _dispatchTransferFrom(
            leftOrderHash,
            leftOrder.makerFeeAssetData,
            leftMakerAddress,
            leftFeeRecipientAddress,
            matchedFillResults.left.makerFeePaid
        );

        // Settle taker profits.
        _dispatchTransferFrom(
            leftOrderHash,
            leftOrder.makerAssetData,
            leftMakerAddress,
            takerAddress,
            matchedFillResults.profitInLeftMakerAsset
        );
        _dispatchTransferFrom(
            rightOrderHash,
            rightOrder.makerAssetData,
            rightMakerAddress,
            takerAddress,
            matchedFillResults.profitInRightMakerAsset
        );

        // Pay protocol fees for each maker
        bool didPayProtocolFees = _payTwoProtocolFees(
            leftOrderHash,
            rightOrderHash,
            matchedFillResults.left.protocolFeePaid,
            leftMakerAddress,
            rightMakerAddress,
            takerAddress
        );

        // Protocol fees are not paid if the protocolFeeCollector contract is not set
        if (!didPayProtocolFees) {
            matchedFillResults.left.protocolFeePaid = 0;
            matchedFillResults.right.protocolFeePaid = 0;
        }

        // Settle taker fees.
        if (
            leftFeeRecipientAddress == rightFeeRecipientAddress &&
            leftOrder.takerFeeAssetData.equals(rightOrder.takerFeeAssetData)
        ) {
            // Fee recipients and taker fee assets are identical, so we can
            // transfer them in one go.
            _dispatchTransferFrom(
                leftOrderHash,
                leftOrder.takerFeeAssetData,
                takerAddress,
                leftFeeRecipientAddress,
                matchedFillResults.left.takerFeePaid.safeAdd(matchedFillResults.right.takerFeePaid)
            );
        } else {
            // Right taker fee -> right fee recipient
            _dispatchTransferFrom(
                rightOrderHash,
                rightOrder.takerFeeAssetData,
                takerAddress,
                rightFeeRecipientAddress,
                matchedFillResults.right.takerFeePaid
            );

            // Left taker fee -> left fee recipient
            _dispatchTransferFrom(
                leftOrderHash,
                leftOrder.takerFeeAssetData,
                takerAddress,
                leftFeeRecipientAddress,
                matchedFillResults.left.takerFeePaid
            );
        }
    }
}
