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

import "../../../exchange/contracts/src/interfaces/IExchange.sol";
import "../../../exchange-libs/contracts/src/LibOrder.sol";
import "../../../utils/contracts/src/LibBytes.sol";


contract LibTransactionDecoder {

    using LibBytes for bytes;

    /// @dev Decodes the call data for an Exchange contract method call.
    /// @param transactionData ABI-encoded calldata for an Exchange
    ///     contract method call.
    /// @return The name of the function called, and the parameters it was
    ///     given.  For single-order fills and cancels, the arrays will have
    ///     just one element.
    function decodeZeroExTransactionData(bytes memory transactionData)
        public
        pure
        returns(
            string memory functionName,
            LibOrder.Order[] memory orders,
            uint256[] memory takerAssetFillAmounts,
            bytes[] memory signatures
        )
    {
        bytes4 functionSelector = transactionData.readBytes4(0);

        if (functionSelector == IExchange(address(0)).batchCancelOrders.selector) {
            functionName = "batchCancelOrders";
        } else if (functionSelector == IExchange(address(0)).batchFillOrders.selector) {
            functionName = "batchFillOrders";
        } else if (functionSelector == IExchange(address(0)).batchFillOrdersNoThrow.selector) {
            functionName = "batchFillOrdersNoThrow";
        } else if (functionSelector == IExchange(address(0)).batchFillOrKillOrders.selector) {
            functionName = "batchFillOrKillOrders";
        } else if (functionSelector == IExchange(address(0)).cancelOrder.selector) {
            functionName = "cancelOrder";
        } else if (functionSelector == IExchange(address(0)).fillOrder.selector) {
            functionName = "fillOrder";
        } else if (functionSelector == IExchange(address(0)).fillOrKillOrder.selector) {
            functionName = "fillOrKillOrder";
        } else if (functionSelector == IExchange(address(0)).marketBuyOrdersNoThrow.selector) {
            functionName = "marketBuyOrdersNoThrow";
        } else if (functionSelector == IExchange(address(0)).marketSellOrdersNoThrow.selector) {
            functionName = "marketSellOrdersNoThrow";
        } else if (functionSelector == IExchange(address(0)).marketBuyOrdersFillOrKill.selector) {
            functionName = "marketBuyOrdersFillOrKill";
        } else if (functionSelector == IExchange(address(0)).marketSellOrdersFillOrKill.selector) {
            functionName = "marketSellOrdersFillOrKill";
        } else if (functionSelector == IExchange(address(0)).matchOrders.selector) {
            functionName = "matchOrders";
        } else if (
            functionSelector == IExchange(address(0)).cancelOrdersUpTo.selector ||
            functionSelector == IExchange(address(0)).executeTransaction.selector
        ) {
            revert();
        } else {
            revert();
        }

        if (functionSelector == IExchange(address(0)).batchCancelOrders.selector) {
            // solhint-disable-next-line indent
            orders = abi.decode(transactionData.slice(4, transactionData.length), (LibOrder.Order[]));
            takerAssetFillAmounts = new uint256[](0);
            signatures = new bytes[](0);
        } else if (
            functionSelector == IExchange(address(0)).batchFillOrKillOrders.selector ||
            functionSelector == IExchange(address(0)).batchFillOrders.selector ||
            functionSelector == IExchange(address(0)).batchFillOrdersNoThrow.selector
        ) {
            (orders, takerAssetFillAmounts, signatures) = _makeReturnValuesForBatchFill(transactionData);
        } else if (functionSelector == IExchange(address(0)).cancelOrder.selector) {
            orders = new LibOrder.Order[](1);
            orders[0] = abi.decode(transactionData.slice(4, transactionData.length), (LibOrder.Order));
            takerAssetFillAmounts = new uint256[](0);
            signatures = new bytes[](0);
        } else if (
            functionSelector == IExchange(address(0)).fillOrKillOrder.selector ||
            functionSelector == IExchange(address(0)).fillOrder.selector
        ) {
            (orders, takerAssetFillAmounts, signatures) = _makeReturnValuesForSingleOrderFill(transactionData);
        } else if (
            functionSelector == IExchange(address(0)).marketBuyOrdersNoThrow.selector ||
            functionSelector == IExchange(address(0)).marketSellOrdersNoThrow.selector ||
            functionSelector == IExchange(address(0)).marketBuyOrdersFillOrKill.selector ||
            functionSelector == IExchange(address(0)).marketSellOrdersFillOrKill.selector
        ) {
            (orders, takerAssetFillAmounts, signatures) = _makeReturnValuesForMarketFill(transactionData);
        } else if (functionSelector == IExchange(address(0)).matchOrders.selector) {
            (
                LibOrder.Order memory leftOrder,
                LibOrder.Order memory rightOrder,
                bytes memory leftSignature,
                bytes memory rightSignature
            ) = abi.decode(
                transactionData.slice(4, transactionData.length),
                (LibOrder.Order, LibOrder.Order, bytes, bytes)
            );

            orders = new LibOrder.Order[](2);
            orders[0] = leftOrder;
            orders[1] = rightOrder;

            takerAssetFillAmounts = new uint256[](2);
            takerAssetFillAmounts[0] = leftOrder.takerAssetAmount;
            takerAssetFillAmounts[1] = rightOrder.takerAssetAmount;

            signatures = new bytes[](2);
            signatures[0] = leftSignature;
            signatures[1] = rightSignature;
        }
    }

    function _makeReturnValuesForSingleOrderFill(bytes memory transactionData)
        private
        pure
        returns(
            LibOrder.Order[] memory orders,
            uint256[] memory takerAssetFillAmounts,
            bytes[] memory signatures
        )
    {
        orders = new LibOrder.Order[](1);
        takerAssetFillAmounts = new uint256[](1);
        signatures = new bytes[](1);
        // solhint-disable-next-line indent
        (orders[0], takerAssetFillAmounts[0], signatures[0]) = abi.decode(
            transactionData.slice(4, transactionData.length),
            (LibOrder.Order, uint256, bytes)
        );
    }

    function _makeReturnValuesForBatchFill(bytes memory transactionData)
        private
        pure
        returns(
            LibOrder.Order[] memory orders,
            uint256[] memory takerAssetFillAmounts,
            bytes[] memory signatures
        )
    {
        // solhint-disable-next-line indent
        (orders, takerAssetFillAmounts, signatures) = abi.decode(
            transactionData.slice(4, transactionData.length),
            // solhint-disable-next-line indent
            (LibOrder.Order[], uint256[], bytes[])
        );
    }

    function _makeReturnValuesForMarketFill(bytes memory transactionData)
        private
        pure
        returns(
            LibOrder.Order[] memory orders,
            uint256[] memory takerAssetFillAmounts,
            bytes[] memory signatures
        )
    {
        takerAssetFillAmounts = new uint256[](1);
        // solhint-disable-next-line indent
        (orders, takerAssetFillAmounts[0], signatures) = abi.decode(
            transactionData.slice(4, transactionData.length),
            // solhint-disable-next-line indent
            (LibOrder.Order[], uint256, bytes[])
        );
    }
}
