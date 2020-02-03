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

import "../../../../exchange-libs/contracts/src/LibOrder.sol";
import "../../../../exchange-libs/contracts/src/LibZeroExTransaction.sol";


// solhint-disable
contract IEIP1271Data {

    /// @dev This function's selector is used when ABI encoding the order
    ///      and hash into a byte array before calling `isValidSignature`.
    ///      This function serves no other purpose.
    function OrderWithHash(
        LibOrder.Order calldata order,
        bytes32 orderHash
    )
        external
        pure;
    
    /// @dev This function's selector is used when ABI encoding the transaction
    ///      and hash into a byte array before calling `isValidSignature`.
    ///      This function serves no other purpose.
    function ZeroExTransactionWithHash(
        LibZeroExTransaction.ZeroExTransaction calldata transaction,
        bytes32 transactionHash
    )
        external
        pure;
}
