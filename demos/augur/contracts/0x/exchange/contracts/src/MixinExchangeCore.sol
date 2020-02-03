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
import "../../../utils/contracts/src/LibSafeMath.sol";
import "../../../utils/contracts/src/Refundable.sol";
import "../../../exchange-libs/contracts/src/LibFillResults.sol";
import "../../../exchange-libs/contracts/src/LibMath.sol";
import "../../../exchange-libs/contracts/src/LibOrder.sol";
import "../../../exchange-libs/contracts/src/LibEIP712ExchangeDomain.sol";
import "../../../exchange-libs/contracts/src/LibExchangeRichErrors.sol";
import "./interfaces/IExchangeCore.sol";
import "./MixinAssetProxyDispatcher.sol";
import "./MixinProtocolFees.sol";
import "./MixinSignatureValidator.sol";


contract MixinExchangeCore is
    IExchangeCore,
    Refundable,
    LibEIP712ExchangeDomain,
    MixinAssetProxyDispatcher,
    MixinProtocolFees,
    MixinSignatureValidator
{
    using LibOrder for LibOrder.Order;
    using LibSafeMath for uint256;
    using LibBytes for bytes;

    /// @dev Mapping of orderHash => amount of takerAsset already bought by maker
    /// @param 0 Order hash.
    /// @return 0 The amount of taker asset filled.
    mapping (bytes32 => uint256) public filled;

    /// @dev Mapping of orderHash => cancelled
    /// @param 0 Order hash.
    /// @return 0 Whether the order was cancelled.
    mapping (bytes32 => bool) public cancelled;

    /// @dev Mapping of makerAddress => senderAddress => lowest salt an order can have in order to be fillable
    ///      Orders with specified senderAddress and with a salt less than their epoch are considered cancelled
    /// @param 0 Address of the order's maker.
    /// @param 1 Address of the order's sender.
    /// @return 0 Minimum valid order epoch.
    mapping (address => mapping (address => uint256)) public orderEpoch;

    /// @dev Cancels all orders created by makerAddress with a salt less than or equal to the targetOrderEpoch
    ///      and senderAddress equal to msg.sender (or null address if msg.sender == makerAddress).
    /// @param targetOrderEpoch Orders created with a salt less or equal to this value will be cancelled.
    function cancelOrdersUpTo(uint256 targetOrderEpoch)
        external
        payable
        refundFinalBalanceNoReentry
    {
        address makerAddress = _getCurrentContextAddress();
        // If this function is called via `executeTransaction`, we only update the orderEpoch for the makerAddress/msg.sender combination.
        // This allows external filter contracts to add rules to how orders are cancelled via this function.
        address orderSenderAddress = makerAddress == msg.sender ? address(0) : msg.sender;

        // orderEpoch is initialized to 0, so to cancelUpTo we need salt + 1
        uint256 newOrderEpoch = targetOrderEpoch + 1;
        uint256 oldOrderEpoch = orderEpoch[makerAddress][orderSenderAddress];

        // Ensure orderEpoch is monotonically increasing
        if (newOrderEpoch <= oldOrderEpoch) {
            revert();
        }

        // Update orderEpoch
        orderEpoch[makerAddress][orderSenderAddress] = newOrderEpoch;
        emit CancelUpTo(
            makerAddress,
            orderSenderAddress,
            newOrderEpoch
        );
    }

    /// @dev Fills the input order.
    /// @param order Order struct containing order specifications.
    /// @param takerAssetFillAmount Desired amount of takerAsset to sell.
    /// @param signature Proof that order has been created by maker.
    /// @return fillResults Amounts filled and fees paid by maker and taker.
    function fillOrder(
        LibOrder.Order memory order,
        uint256 takerAssetFillAmount,
        bytes memory signature
    )
        public
        payable
        refundFinalBalanceNoReentry
        returns (LibFillResults.FillResults memory fillResults)
    {
        fillResults = _fillOrder(
            order,
            takerAssetFillAmount,
            signature
        );
        return fillResults;
    }

    /// @dev After calling, the order can not be filled anymore.
    /// @param order Order struct containing order specifications.
    function cancelOrder(LibOrder.Order memory order)
        public
        payable
        refundFinalBalanceNoReentry
    {
        _cancelOrder(order);
    }

    /// @dev Gets information about an order: status, hash, and amount filled.
    /// @param order Order to gather information on.
    /// @return orderInfo Information about the order and its state.
    ///         See LibOrder.OrderInfo for a complete description.
    function getOrderInfo(LibOrder.Order memory order)
        public
        view
        returns (LibOrder.OrderInfo memory orderInfo)
    {
        // Compute the order hash and fetch the amount of takerAsset that has already been filled
        (orderInfo.orderHash, orderInfo.orderTakerAssetFilledAmount) = _getOrderHashAndFilledAmount(order);

        // If order.makerAssetAmount is zero, we also reject the order.
        // While the Exchange contract handles them correctly, they create
        // edge cases in the supporting infrastructure because they have
        // an 'infinite' price when computed by a simple division.
        if (order.makerAssetAmount == 0) {
            orderInfo.orderStatus = LibOrder.OrderStatus.INVALID_MAKER_ASSET_AMOUNT;
            return orderInfo;
        }

        // If order.takerAssetAmount is zero, then the order will always
        // be considered filled because 0 == takerAssetAmount == orderTakerAssetFilledAmount
        // Instead of distinguishing between unfilled and filled zero taker
        // amount orders, we choose not to support them.
        if (order.takerAssetAmount == 0) {
            orderInfo.orderStatus = LibOrder.OrderStatus.INVALID_TAKER_ASSET_AMOUNT;
            return orderInfo;
        }

        // Validate order availability
        if (orderInfo.orderTakerAssetFilledAmount >= order.takerAssetAmount) {
            orderInfo.orderStatus = LibOrder.OrderStatus.FULLY_FILLED;
            return orderInfo;
        }

        // Validate order expiration
        // solhint-disable-next-line not-rely-on-time
        if (block.timestamp >= order.expirationTimeSeconds) {
            orderInfo.orderStatus = LibOrder.OrderStatus.EXPIRED;
            return orderInfo;
        }

        // Check if order has been cancelled
        if (cancelled[orderInfo.orderHash]) {
            orderInfo.orderStatus = LibOrder.OrderStatus.CANCELLED;
            return orderInfo;
        }
        if (orderEpoch[order.makerAddress][order.senderAddress] > order.salt) {
            orderInfo.orderStatus = LibOrder.OrderStatus.CANCELLED;
            return orderInfo;
        }

        // All other statuses are ruled out: order is Fillable
        orderInfo.orderStatus = LibOrder.OrderStatus.FILLABLE;
        return orderInfo;
    }

    /// @dev Fills the input order.
    /// @param order Order struct containing order specifications.
    /// @param takerAssetFillAmount Desired amount of takerAsset to sell.
    /// @param signature Proof that order has been created by maker.
    /// @return fillResults Amounts filled and fees paid by maker and taker.
    function _fillOrder(
        LibOrder.Order memory order,
        uint256 takerAssetFillAmount,
        bytes memory signature
    )
        internal
        returns (LibFillResults.FillResults memory fillResults)
    {
        // Fetch order info
        LibOrder.OrderInfo memory orderInfo = getOrderInfo(order);

        // Fetch taker address
        address takerAddress = _getCurrentContextAddress();

        // Assert that the order is fillable by taker
        _assertFillableOrder(
            order,
            orderInfo,
            takerAddress,
            signature
        );

        // Get amount of takerAsset to fill
        uint256 remainingTakerAssetAmount = order.takerAssetAmount.safeSub(orderInfo.orderTakerAssetFilledAmount);
        uint256 takerAssetFilledAmount = LibSafeMath.min256(takerAssetFillAmount, remainingTakerAssetAmount);

        // Compute proportional fill amounts
        fillResults = LibFillResults.calculateFillResults(
            order,
            takerAssetFilledAmount,
            protocolFeeMultiplier,
            tx.gasprice
        );

        bytes32 orderHash = orderInfo.orderHash;

        // Update exchange internal state
        _updateFilledState(
            order,
            takerAddress,
            orderHash,
            orderInfo.orderTakerAssetFilledAmount,
            fillResults
        );

        // Settle order
        _settleOrder(
            orderHash,
            order,
            takerAddress,
            fillResults
        );

        return fillResults;
    }

    /// @dev After calling, the order can not be filled anymore.
    ///      Throws if order is invalid or sender does not have permission to cancel.
    /// @param order Order to cancel. Order must be OrderStatus.FILLABLE.
    function _cancelOrder(LibOrder.Order memory order)
        internal
    {
        // Fetch current order status
        LibOrder.OrderInfo memory orderInfo = getOrderInfo(order);

        // Validate context
        _assertValidCancel(order, orderInfo);

        // Noop if order is already unfillable
        if (orderInfo.orderStatus != LibOrder.OrderStatus.FILLABLE) {
            return;
        }

        // Perform cancel
        _updateCancelledState(order, orderInfo.orderHash);
    }

    /// @dev Updates state with results of a fill order.
    /// @param order that was filled.
    /// @param takerAddress Address of taker who filled the order.
    /// @param orderTakerAssetFilledAmount Amount of order already filled.
    function _updateFilledState(
        LibOrder.Order memory order,
        address takerAddress,
        bytes32 orderHash,
        uint256 orderTakerAssetFilledAmount,
        LibFillResults.FillResults memory fillResults
    )
        internal
    {
        // Update state
        filled[orderHash] = orderTakerAssetFilledAmount.safeAdd(fillResults.takerAssetFilledAmount);

        emit Fill(
            order.makerAddress,
            order.feeRecipientAddress,
            order.makerAssetData,
            order.takerAssetData,
            order.makerFeeAssetData,
            order.takerFeeAssetData,
            orderHash,
            takerAddress,
            msg.sender,
            fillResults.makerAssetFilledAmount,
            fillResults.takerAssetFilledAmount,
            fillResults.makerFeePaid,
            fillResults.takerFeePaid,
            fillResults.protocolFeePaid
        );
    }

    /// @dev Updates state with results of cancelling an order.
    ///      State is only updated if the order is currently fillable.
    ///      Otherwise, updating state would have no effect.
    /// @param order that was cancelled.
    /// @param orderHash Hash of order that was cancelled.
    function _updateCancelledState(
        LibOrder.Order memory order,
        bytes32 orderHash
    )
        internal
    {
        // Perform cancel
        cancelled[orderHash] = true;

        // Log cancel
        emit Cancel(
            order.makerAddress,
            order.feeRecipientAddress,
            order.makerAssetData,
            order.takerAssetData,
            msg.sender,
            orderHash
        );
    }

    /// @dev Validates context for fillOrder. Succeeds or throws.
    /// @param order to be filled.
    /// @param orderInfo OrderStatus, orderHash, and amount already filled of order.
    /// @param takerAddress Address of order taker.
    /// @param signature Proof that the orders was created by its maker.
    function _assertFillableOrder(
        LibOrder.Order memory order,
        LibOrder.OrderInfo memory orderInfo,
        address takerAddress,
        bytes memory signature
    )
        internal
        view
    {
        // An order can only be filled if its status is FILLABLE.
        if (orderInfo.orderStatus != LibOrder.OrderStatus.FILLABLE) {
            revert();
        }

        // Validate sender is allowed to fill this order
        if (order.senderAddress != address(0)) {
            if (order.senderAddress != msg.sender) {
                revert();
            }
        }

        // Validate taker is allowed to fill this order
        if (order.takerAddress != address(0)) {
            if (order.takerAddress != takerAddress) {
                revert();
            }
        }

        // Validate signature
        if (!_isValidOrderWithHashSignature(
                order,
                orderInfo.orderHash,
                signature
            )
        ) {
            revert();
        }
    }

    /// @dev Validates context for cancelOrder. Succeeds or throws.
    /// @param order to be cancelled.
    /// @param orderInfo OrderStatus, orderHash, and amount already filled of order.
    function _assertValidCancel(
        LibOrder.Order memory order,
        LibOrder.OrderInfo memory orderInfo
    )
        internal
        view
    {
        // Validate sender is allowed to cancel this order
        if (order.senderAddress != address(0)) {
            if (order.senderAddress != msg.sender) {
                revert();
            }
        }

        // Validate transaction signed by maker
        address makerAddress = _getCurrentContextAddress();
        if (order.makerAddress != makerAddress) {
            revert();
        }
    }

    /// @dev Settles an order by transferring assets between counterparties.
    /// @param orderHash The order hash.
    /// @param order Order struct containing order specifications.
    /// @param takerAddress Address selling takerAsset and buying makerAsset.
    /// @param fillResults Amounts to be filled and fees paid by maker and taker.
    function _settleOrder(
        bytes32 orderHash,
        LibOrder.Order memory order,
        address takerAddress,
        LibFillResults.FillResults memory fillResults
    )
        internal
    {
        // Transfer taker -> maker
        _dispatchTransferFrom(
            orderHash,
            order.takerAssetData,
            takerAddress,
            order.makerAddress,
            fillResults.takerAssetFilledAmount
        );

        // Transfer maker -> taker
        _dispatchTransferFrom(
            orderHash,
            order.makerAssetData,
            order.makerAddress,
            takerAddress,
            fillResults.makerAssetFilledAmount
        );

        // Transfer taker fee -> feeRecipient
        _dispatchTransferFrom(
            orderHash,
            order.takerFeeAssetData,
            takerAddress,
            order.feeRecipientAddress,
            fillResults.takerFeePaid
        );

        // Transfer maker fee -> feeRecipient
        _dispatchTransferFrom(
            orderHash,
            order.makerFeeAssetData,
            order.makerAddress,
            order.feeRecipientAddress,
            fillResults.makerFeePaid
        );

        // Pay protocol fee
        bool didPayProtocolFee = _paySingleProtocolFee(
            orderHash,
            fillResults.protocolFeePaid,
            order.makerAddress,
            takerAddress
        );

        // Protocol fees are not paid if the protocolFeeCollector contract is not set
        if (!didPayProtocolFee) {
            fillResults.protocolFeePaid = 0;
        }
    }

    /// @dev Gets the order's hash and amount of takerAsset that has already been filled.
    /// @param order Order struct containing order specifications.
    /// @return The typed data hash and amount filled of the order.
    function _getOrderHashAndFilledAmount(LibOrder.Order memory order)
        internal
        view
        returns (bytes32 orderHash, uint256 orderTakerAssetFilledAmount)
    {
        orderHash = order.getTypedDataHash(EIP712_EXCHANGE_DOMAIN_HASH);
        orderTakerAssetFilledAmount = filled[orderHash];
        return (orderHash, orderTakerAssetFilledAmount);
    }
}
