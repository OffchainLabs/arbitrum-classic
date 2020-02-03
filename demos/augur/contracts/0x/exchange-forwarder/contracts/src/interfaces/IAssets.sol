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


contract IAssets {

    /// @dev Withdraws assets from this contract. The contract requires a ZRX balance in order to
    ///      function optimally, and this function allows the ZRX to be withdrawn by owner. It may also be
    ///      used to withdraw assets that were accidentally sent to this contract.
    /// @param assetData Byte array encoded for the respective asset proxy.
    /// @param amount Amount of ERC20 token to withdraw.
    function withdrawAsset(
        bytes calldata assetData,
        uint256 amount
    )
        external;

        /// @dev Approves the respective proxy for a given asset to transfer tokens on the Forwarder contract's behalf.
        ///      This is necessary because an order fee denominated in the maker asset (i.e. a percentage fee) is sent by the
        ///      Forwarder contract to the fee recipient.
        ///      This method needs to be called before forwarding orders of a maker asset that hasn't
        ///      previously been approved.
        /// @param assetData Byte array encoded for the respective asset proxy.
    function approveMakerAssetProxy(
        bytes calldata assetData
    )
        external;
}
