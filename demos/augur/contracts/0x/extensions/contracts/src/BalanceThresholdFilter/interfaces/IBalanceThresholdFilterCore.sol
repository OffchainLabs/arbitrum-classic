
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


contract IBalanceThresholdFilterCore {
    
    /// @dev Executes an Exchange transaction iff the maker and taker meet 
    ///      the hold at least `BALANCE_THRESHOLD` of the asset `THRESHOLD_ASSET` OR 
    ///      the exchange function is a cancellation.
    ///      Supported Exchange functions:
    ///         - batchFillOrders
    ///         - batchFillOrdersNoThrow
    ///         - batchFillOrKillOrders
    ///         - fillOrder
    ///         - fillOrderNoThrow
    ///         - fillOrKillOrder
    ///         - marketBuyOrders
    ///         - marketBuyOrdersNoThrow
    ///         - marketSellOrders
    ///         - marketSellOrdersNoThrow
    ///         - matchOrders
    ///         - cancelOrder
    ///         - batchCancelOrders
    ///         - cancelOrdersUpTo
    ///     Trying to call any other exchange function will throw.
    /// @param salt Arbitrary number to ensure uniqueness of transaction hash.
    /// @param signerAddress Address of transaction signer.
    /// @param signedExchangeTransaction AbiV2 encoded calldata.
    /// @param signature Proof of signer transaction by signer.
    function executeTransaction(
        uint256 salt,
        address signerAddress,
        bytes calldata signedExchangeTransaction,
        bytes calldata signature
    ) 
        external;
}
