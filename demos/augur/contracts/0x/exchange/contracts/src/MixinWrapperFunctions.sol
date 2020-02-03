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

import "../../../utils/contracts/src/LibSafeMath.sol";

import "../../../exchange-libs/contracts/src/LibOrder.sol";
import "../../../exchange-libs/contracts/src/LibMath.sol";
import "../../../exchange-libs/contracts/src/LibFillResults.sol";
import "../../../exchange-libs/contracts/src/LibExchangeRichErrors.sol";
import "./interfaces/IExchangeCore.sol";
import "./interfaces/IWrapperFunctions.sol";
import "./MixinExchangeCore.sol";


contract MixinWrapperFunctions is
    IWrapperFunctions,
    MixinExchangeCore
{
    using LibSafeMath for uint256;

    /// @dev Fills the input order. Reverts if exact `takerAssetFillAmount` not filled.
    /// @param order Order struct containing order specifications.
    /// @param takerAssetFillAmount Desired amount of takerAsset to sell.
    /// @param signature Proof that order has been created by maker.
    /// @return fillResults Amounts filled and fees paid.
    function fillOrKillOrder(
        LibOrder.Order memory order,
        uint256 takerAssetFillAmount,
        bytes memory signature
    )
        public
        payable
        refundFinalBalanceNoReentry
        returns (LibFillResults.FillResults memory fillResults)
    {
        fillResults = _fillOrKillOrder(
            order,
            takerAssetFillAmount,
            signature
        );
        return fillResults;
    }

    /// @dev Executes multiple calls of fillOrder.
    /// @param orders Array of order specifications.
    /// @param takerAssetFillAmounts Array of desired amounts of takerAsset to sell in orders.
    /// @param signatures Proofs that orders have been created by makers.
    /// @return fillResults Array of amounts filled and fees paid by makers and taker.
    function batchFillOrders(
        LibOrder.Order[] memory orders,
        uint256[] memory takerAssetFillAmounts,
        bytes[] memory signatures
    )
        public
        payable
        refundFinalBalanceNoReentry
        returns (LibFillResults.FillResults[] memory fillResults)
    {
        uint256 ordersLength = orders.length;
        fillResults = new LibFillResults.FillResults[](ordersLength);
        for (uint256 i = 0; i != ordersLength; i++) {
            fillResults[i] = _fillOrder(
                orders[i],
                takerAssetFillAmounts[i],
                signatures[i]
            );
        }
        return fillResults;
    }

    /// @dev Executes multiple calls of fillOrKillOrder.
    /// @param orders Array of order specifications.
    /// @param takerAssetFillAmounts Array of desired amounts of takerAsset to sell in orders.
    /// @param signatures Proofs that orders have been created by makers.
    /// @return fillResults Array of amounts filled and fees paid by makers and taker.
    function batchFillOrKillOrders(
        LibOrder.Order[] memory orders,
        uint256[] memory takerAssetFillAmounts,
        bytes[] memory signatures
    )
        public
        payable
        refundFinalBalanceNoReentry
        returns (LibFillResults.FillResults[] memory fillResults)
    {
        uint256 ordersLength = orders.length;
        fillResults = new LibFillResults.FillResults[](ordersLength);
        for (uint256 i = 0; i != ordersLength; i++) {
            fillResults[i] = _fillOrKillOrder(
                orders[i],
                takerAssetFillAmounts[i],
                signatures[i]
            );
        }
        return fillResults;
    }

    /// @dev Executes multiple calls of fillOrder. If any fill reverts, the error is caught and ignored.
    /// @param orders Array of order specifications.
    /// @param takerAssetFillAmounts Array of desired amounts of takerAsset to sell in orders.
    /// @param signatures Proofs that orders have been created by makers.
    /// @return fillResults Array of amounts filled and fees paid by makers and taker.
    function batchFillOrdersNoThrow(
        LibOrder.Order[] memory orders,
        uint256[] memory takerAssetFillAmounts,
        bytes[] memory signatures
    )
        public
        payable
        disableRefundUntilEnd
        returns (LibFillResults.FillResults[] memory fillResults)
    {
        uint256 ordersLength = orders.length;
        fillResults = new LibFillResults.FillResults[](ordersLength);
        for (uint256 i = 0; i != ordersLength; i++) {
            fillResults[i] = _fillOrderNoThrow(
                orders[i],
                takerAssetFillAmounts[i],
                signatures[i]
            );
        }
        return fillResults;
    }

    /// @dev Executes multiple calls of fillOrder until total amount of takerAsset is sold by taker.
    ///      If any fill reverts, the error is caught and ignored.
    ///      NOTE: This function does not enforce that the takerAsset is the same for each order.
    /// @param orders Array of order specifications.
    /// @param takerAssetFillAmount Desired amount of takerAsset to sell.
    /// @param signatures Proofs that orders have been signed by makers.
    /// @return fillResults Amounts filled and fees paid by makers and taker.
    function marketSellOrdersNoThrow(
        LibOrder.Order[] memory orders,
        uint256 takerAssetFillAmount,
        bytes[] memory signatures
    )
        public
        payable
        disableRefundUntilEnd
        returns (LibFillResults.FillResults memory fillResults)
    {
        uint256 ordersLength = orders.length;
        for (uint256 i = 0; i != ordersLength; i++) {

            // Calculate the remaining amount of takerAsset to sell
            uint256 remainingTakerAssetFillAmount = takerAssetFillAmount.safeSub(fillResults.takerAssetFilledAmount);

            // Attempt to sell the remaining amount of takerAsset
            LibFillResults.FillResults memory singleFillResults = _fillOrderNoThrow(
                orders[i],
                remainingTakerAssetFillAmount,
                signatures[i]
            );

            // Update amounts filled and fees paid by maker and taker
            fillResults = LibFillResults.addFillResults(fillResults, singleFillResults);

            // Stop execution if the entire amount of takerAsset has been sold
            if (fillResults.takerAssetFilledAmount >= takerAssetFillAmount) {
                break;
            }
        }
        return fillResults;
    }

    /// @dev Executes multiple calls of fillOrder until total amount of makerAsset is bought by taker.
    ///      If any fill reverts, the error is caught and ignored.
    ///      NOTE: This function does not enforce that the makerAsset is the same for each order.
    /// @param orders Array of order specifications.
    /// @param makerAssetFillAmount Desired amount of makerAsset to buy.
    /// @param signatures Proofs that orders have been signed by makers.
    /// @return fillResults Amounts filled and fees paid by makers and taker.
    function marketBuyOrdersNoThrow(
        LibOrder.Order[] memory orders,
        uint256 makerAssetFillAmount,
        bytes[] memory signatures
    )
        public
        payable
        disableRefundUntilEnd
        returns (LibFillResults.FillResults memory fillResults)
    {
        uint256 ordersLength = orders.length;
        for (uint256 i = 0; i != ordersLength; i++) {

            // Calculate the remaining amount of makerAsset to buy
            uint256 remainingMakerAssetFillAmount = makerAssetFillAmount.safeSub(fillResults.makerAssetFilledAmount);

            // Convert the remaining amount of makerAsset to buy into remaining amount
            // of takerAsset to sell, assuming entire amount can be sold in the current order
            uint256 remainingTakerAssetFillAmount = LibMath.getPartialAmountCeil(
                orders[i].takerAssetAmount,
                orders[i].makerAssetAmount,
                remainingMakerAssetFillAmount
            );

            // Attempt to sell the remaining amount of takerAsset
            LibFillResults.FillResults memory singleFillResults = _fillOrderNoThrow(
                orders[i],
                remainingTakerAssetFillAmount,
                signatures[i]
            );

            // Update amounts filled and fees paid by maker and taker
            fillResults = LibFillResults.addFillResults(fillResults, singleFillResults);

            // Stop execution if the entire amount of makerAsset has been bought
            if (fillResults.makerAssetFilledAmount >= makerAssetFillAmount) {
                break;
            }
        }
        return fillResults;
    }

    /// @dev Calls marketSellOrdersNoThrow then reverts if < takerAssetFillAmount has been sold.
    ///      NOTE: This function does not enforce that the takerAsset is the same for each order.
    /// @param orders Array of order specifications.
    /// @param takerAssetFillAmount Minimum amount of takerAsset to sell.
    /// @param signatures Proofs that orders have been signed by makers.
    /// @return fillResults Amounts filled and fees paid by makers and taker.
    function marketSellOrdersFillOrKill(
        LibOrder.Order[] memory orders,
        uint256 takerAssetFillAmount,
        bytes[] memory signatures
    )
        public
        payable
        returns (LibFillResults.FillResults memory fillResults)
    {
        fillResults = marketSellOrdersNoThrow(orders, takerAssetFillAmount, signatures);
        if (fillResults.takerAssetFilledAmount < takerAssetFillAmount) {
            revert();
        }
    }

    /// @dev Calls marketBuyOrdersNoThrow then reverts if < makerAssetFillAmount has been bought.
    ///      NOTE: This function does not enforce that the makerAsset is the same for each order.
    /// @param orders Array of order specifications.
    /// @param makerAssetFillAmount Minimum amount of makerAsset to buy.
    /// @param signatures Proofs that orders have been signed by makers.
    /// @return fillResults Amounts filled and fees paid by makers and taker.
    function marketBuyOrdersFillOrKill(
        LibOrder.Order[] memory orders,
        uint256 makerAssetFillAmount,
        bytes[] memory signatures
    )
        public
        payable
        returns (LibFillResults.FillResults memory fillResults)
    {
        fillResults = marketBuyOrdersNoThrow(orders, makerAssetFillAmount, signatures);
        if (fillResults.makerAssetFilledAmount < makerAssetFillAmount) {
            revert();
        }
    }

    /// @dev Executes multiple calls of cancelOrder.
    /// @param orders Array of order specifications.
    function batchCancelOrders(LibOrder.Order[] memory orders)
        public
        payable
        refundFinalBalanceNoReentry
    {
        uint256 ordersLength = orders.length;
        for (uint256 i = 0; i != ordersLength; i++) {
            _cancelOrder(orders[i]);
        }
    }

    /// @dev Fills the input order. Reverts if exact takerAssetFillAmount not filled.
    /// @param order Order struct containing order specifications.
    /// @param takerAssetFillAmount Desired amount of takerAsset to sell.
    /// @param fillResults ignature Proof that order has been created by maker.
    function _fillOrKillOrder(
        LibOrder.Order memory order,
        uint256 takerAssetFillAmount,
        bytes memory signature
    )
        internal
        returns (LibFillResults.FillResults memory fillResults)
    {
        fillResults = _fillOrder(
            order,
            takerAssetFillAmount,
            signature
        );
        if (fillResults.takerAssetFilledAmount != takerAssetFillAmount) {
            revert();
        }
        return fillResults;
    }

    /// @dev Fills the input order.
    ///      Returns a null FillResults instance if the transaction would otherwise revert.
    /// @param order Order struct containing order specifications.
    /// @param takerAssetFillAmount Desired amount of takerAsset to sell.
    /// @param signature Proof that order has been created by maker.
    /// @return fillResults Amounts filled and fees paid by maker and taker.
    function _fillOrderNoThrow(
        LibOrder.Order memory order,
        uint256 takerAssetFillAmount,
        bytes memory signature
    )
        internal
        returns (LibFillResults.FillResults memory fillResults)
    {
        // ABI encode calldata for `fillOrder`
        bytes memory fillOrderCalldata = abi.encodeWithSelector(
            IExchangeCore(address(0)).fillOrder.selector,
            order,
            takerAssetFillAmount,
            signature
        );

        (bool didSucceed, bytes memory returnData) = address(this).delegatecall(fillOrderCalldata);
        if (didSucceed) {
            assert(returnData.length == 160);
            fillResults = abi.decode(returnData, (LibFillResults.FillResults));
        }
        // fillResults values will be 0 by default if call was unsuccessful
        return fillResults;
    }
}
