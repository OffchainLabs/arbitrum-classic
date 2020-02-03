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
import "../../../../exchange-libs/contracts/src/LibOrder.sol";
import "../../../../erc20/contracts/src/interfaces/IERC20Token.sol";
import "../../../../utils/contracts/src/LibBytes.sol";
import "../../../../utils/contracts/src/LibSafeMath.sol";


contract DutchAuction {

    using LibBytes for bytes;
    using LibSafeMath for uint256;

    // solhint-disable var-name-mixedcase
    IExchange internal EXCHANGE;

    struct AuctionDetails {
        uint256 beginTimeSeconds;    // Auction begin unix timestamp: sellOrder.makerAssetData
        uint256 endTimeSeconds;      // Auction end unix timestamp: sellOrder.expiryTimeSeconds
        uint256 beginAmount;         // Auction begin amount: sellOrder.makerAssetData
        uint256 endAmount;           // Auction end amount: sellOrder.takerAssetAmount
        uint256 currentAmount;       // Calculated amount given block.timestamp
        uint256 currentTimeSeconds;  // block.timestamp
    }

    constructor (address _exchange)
        public
    {
        EXCHANGE = IExchange(_exchange);
    }

    /// @dev Matches the buy and sell orders at an amount given the following: the current block time, the auction
    ///      start time and the auction begin amount. The sell order is a an order at the lowest amount
    ///      at the end of the auction. Excess from the match is transferred to the seller.
    ///      Over time the price moves from beginAmount to endAmount given the current block.timestamp.
    ///      sellOrder.expiryTimeSeconds is the end time of the auction.
    ///      sellOrder.takerAssetAmount is the end amount of the auction (lowest possible amount).
    ///      sellOrder.makerAssetData is the ABI encoded Asset Proxy data with the following data appended
    ///      buyOrder.makerAssetData is the buyers bid on the auction, must meet the amount for the current block timestamp
    ///      (uint256 beginTimeSeconds, uint256 beginAmount).
    ///      This function reverts in the following scenarios:
    ///         * Auction has not started (auctionDetails.currentTimeSeconds < auctionDetails.beginTimeSeconds)
    ///         * Auction has expired (auctionDetails.endTimeSeconds < auctionDetails.currentTimeSeconds)
    ///         * Amount is invalid: Buy order amount is too low (buyOrder.makerAssetAmount < auctionDetails.currentAmount)
    ///         * Amount is invalid: Invalid begin amount (auctionDetails.beginAmount > auctionDetails.endAmount)
    ///         * Any failure in the 0x Match Orders
    /// @param buyOrder The Buyer's order. This order is for the current expected price of the auction.
    /// @param sellOrder The Seller's order. This order is for the lowest amount (at the end of the auction).
    /// @param buySignature Proof that order was created by the buyer.
    /// @param sellSignature Proof that order was created by the seller.
    /// @return matchedFillResults amounts filled and fees paid by maker and taker of matched orders.
    function matchOrders(
        LibOrder.Order memory buyOrder,
        LibOrder.Order memory sellOrder,
        bytes memory buySignature,
        bytes memory sellSignature
    )
        public
        returns (LibFillResults.MatchedFillResults memory matchedFillResults)
    {
        AuctionDetails memory auctionDetails = getAuctionDetails(sellOrder);
        // Ensure the auction has not yet started
        require(
            auctionDetails.currentTimeSeconds >= auctionDetails.beginTimeSeconds,
            "AUCTION_NOT_STARTED"
        );
        // Ensure the auction has not expired. This will fail later in 0x but we can save gas by failing early
        require(
            sellOrder.expirationTimeSeconds > auctionDetails.currentTimeSeconds,
            "AUCTION_EXPIRED"
        );
        // Validate the buyer amount is greater than the current auction amount
        require(
            buyOrder.makerAssetAmount >= auctionDetails.currentAmount,
            "INVALID_AMOUNT"
        );
        // Match orders, maximally filling `buyOrder`
        matchedFillResults = EXCHANGE.matchOrders(
            buyOrder,
            sellOrder,
            buySignature,
            sellSignature
        );
        // The difference in sellOrder.takerAssetAmount and current amount is given as spread to the matcher
        // This may include additional spread from the buyOrder.makerAssetAmount and the currentAmount.
        // e.g currentAmount is 30, sellOrder.takerAssetAmount is 10 and buyOrder.makerAssetamount is 40.
        // 10 (40-30) is returned to the buyer, 20 (30-10) sent to the seller and 10 has previously
        // been transferred to the seller during matchOrders
        uint256 leftMakerAssetSpreadAmount = matchedFillResults.leftMakerAssetSpreadAmount;
        if (leftMakerAssetSpreadAmount > 0) {
            // ERC20 Asset data itself is encoded as follows:
            //
            // | Area     | Offset | Length  | Contents                            |
            // |----------|--------|---------|-------------------------------------|
            // | Header   | 0      | 4       | function selector                   |
            // | Params   |        | 1 * 32  | function parameters:                |
            // |          | 4      | 12      |   1. token address padding          |
            // |          | 16     | 20      |   2. token address                  |
            bytes memory assetData = sellOrder.takerAssetData;
            address token = assetData.readAddress(16);
            // Calculate the excess from the buy order. This can occur if the buyer sends in a higher
            // amount than the calculated current amount
            uint256 buyerExcessAmount = buyOrder.makerAssetAmount.safeSub(auctionDetails.currentAmount);
            uint256 sellerExcessAmount = leftMakerAssetSpreadAmount.safeSub(buyerExcessAmount);
            // Return the difference between auctionDetails.currentAmount and sellOrder.takerAssetAmount
            // to the seller
            if (sellerExcessAmount > 0) {
                IERC20Token(token).transfer(sellOrder.makerAddress, sellerExcessAmount);
            }
            // Return the difference between buyOrder.makerAssetAmount and auctionDetails.currentAmount
            // to the buyer
            if (buyerExcessAmount > 0) {
                IERC20Token(token).transfer(buyOrder.makerAddress, buyerExcessAmount);
            }
        }
        return matchedFillResults;
    }

    /// @dev Calculates the Auction Details for the given order
    /// @param order The sell order
    /// @return AuctionDetails
    function getAuctionDetails(LibOrder.Order memory order)
        public
        returns (AuctionDetails memory auctionDetails)
    {
        uint256 makerAssetDataLength = order.makerAssetData.length;
        // It is unknown the encoded data of makerAssetData, we assume the last 64 bytes
        // are the Auction Details encoding.
        // Auction Details is encoded as follows:
        //
        // | Area     | Offset | Length  | Contents                            |
        // |----------|--------|---------|-------------------------------------|
        // | Params   |        | 2 * 32  | parameters:                         |
        // |          | -64    | 32      |   1. auction begin unix timestamp   |
        // |          | -32    | 32      |   2. auction begin begin amount     |
        // ERC20 asset data length is 4+32, 64 for auction details results in min length 100
        require(
            makerAssetDataLength >= 100,
            "INVALID_ASSET_DATA"
        );
        uint256 auctionBeginTimeSeconds = order.makerAssetData.readUint256(makerAssetDataLength - 64);
        uint256 auctionBeginAmount = order.makerAssetData.readUint256(makerAssetDataLength - 32);
        // Ensure the auction has a valid begin time
        require(
            order.expirationTimeSeconds > auctionBeginTimeSeconds,
            "INVALID_BEGIN_TIME"
        );
        uint256 auctionDurationSeconds = order.expirationTimeSeconds-auctionBeginTimeSeconds;
        // Ensure the auction goes from high to low
        uint256 minAmount = order.takerAssetAmount;
        require(
            auctionBeginAmount > minAmount,
            "INVALID_AMOUNT"
        );
        uint256 amountDelta = auctionBeginAmount-minAmount;
        // solhint-disable-next-line not-rely-on-time
        uint256 timestamp = block.timestamp;
        auctionDetails.beginTimeSeconds = auctionBeginTimeSeconds;
        auctionDetails.endTimeSeconds = order.expirationTimeSeconds;
        auctionDetails.beginAmount = auctionBeginAmount;
        auctionDetails.endAmount = minAmount;
        auctionDetails.currentTimeSeconds = timestamp;

        uint256 remainingDurationSeconds = order.expirationTimeSeconds-timestamp;
        if (timestamp < auctionBeginTimeSeconds) {
            // If the auction has not yet begun the current amount is the auctionBeginAmount
            auctionDetails.currentAmount = auctionBeginAmount;
        } else if (timestamp >= order.expirationTimeSeconds) {
            // If the auction has ended the current amount is the minAmount.
            // Auction end time is guaranteed by 0x Exchange due to the order expiration
            auctionDetails.currentAmount = minAmount;
        } else {
            auctionDetails.currentAmount = minAmount.safeAdd(
                remainingDurationSeconds.safeMul(amountDelta).safeDiv(auctionDurationSeconds)
            );
        }
        return auctionDetails;
    }
}
