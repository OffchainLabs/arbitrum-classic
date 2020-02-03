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


contract ICoordinatorApprovalVerifier {

    /// @dev Validates that the 0x transaction has been approved by all of the feeRecipients
    ///      that correspond to each order in the transaction's Exchange calldata.
    /// @param transaction 0x transaction containing salt, signerAddress, and data.
    /// @param txOrigin Required signer of Ethereum transaction calling this function.
    /// @param transactionSignature Proof that the transaction has been signed by the signer.
    /// @param approvalSignatures Array of signatures that correspond to the feeRecipients of each
    ///        order in the transaction's Exchange calldata.
    function assertValidCoordinatorApprovals(
        LibZeroExTransaction.ZeroExTransaction memory transaction,
        address txOrigin,
        bytes memory transactionSignature,
        bytes[] memory approvalSignatures
    )
        public
        view;

    /// @dev Decodes the orders from Exchange calldata representing any fill method.
    /// @param data Exchange calldata representing a fill method.
    /// @return orders The orders from the Exchange calldata.
    function decodeOrdersFromFillData(bytes memory data)
        public
        pure
        returns (LibOrder.Order[] memory orders);
}
