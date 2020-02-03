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


contract ITransferSimulator {

    /// @dev This function may be used to simulate any amount of transfers
    /// As they would occur through the Exchange contract. Note that this function
    /// will always revert, even if all transfers are successful. However, it may
    /// be used with eth_call or with a try/catch pattern in order to simulate
    /// the results of the transfers.
    /// @param assetData Array of asset details, each encoded per the AssetProxy contract specification.
    /// @param fromAddresses Array containing the `from` addresses that correspond with each transfer.
    /// @param toAddresses Array containing the `to` addresses that correspond with each transfer.
    /// @param amounts Array containing the amounts that correspond to each transfer.
    /// @return This function does not return a value. However, it will always revert with
    /// `Error("TRANSFERS_SUCCESSFUL")` if all of the transfers were successful.
    function simulateDispatchTransferFromCalls(
        bytes[] memory assetData,
        address[] memory fromAddresses,
        address[] memory toAddresses,
        uint256[] memory amounts
    )
        public;
}
