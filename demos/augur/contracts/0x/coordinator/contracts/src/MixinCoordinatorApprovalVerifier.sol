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

import "../../../exchange-libs/contracts/src/LibEIP712ExchangeDomain.sol";
import "../../../exchange-libs/contracts/src/LibOrder.sol";
import "../../../exchange-libs/contracts/src/LibZeroExTransaction.sol";
import "../../../utils/contracts/src/LibAddressArray.sol";
import "../../../utils/contracts/src/LibBytes.sol";
import "../../../exchange/contracts/src/interfaces/IExchange.sol";
import "./libs/LibCoordinatorApproval.sol";
import "./libs/LibCoordinatorRichErrors.sol";
import "./interfaces/ICoordinatorSignatureValidator.sol";
import "./interfaces/ICoordinatorApprovalVerifier.sol";


// solhint-disable avoid-tx-origin
contract MixinCoordinatorApprovalVerifier is
    LibCoordinatorApproval,
    LibEIP712ExchangeDomain,
    ICoordinatorSignatureValidator,
    ICoordinatorApprovalVerifier
{
    using LibBytes for bytes;
    using LibAddressArray for address[];

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
        view
    {
        // Get the orders from the the Exchange calldata in the 0x transaction
        LibOrder.Order[] memory orders = decodeOrdersFromFillData(transaction.data);

        // No approval is required for non-fill methods
        if (orders.length > 0) {
            // Revert if approval is invalid for transaction orders
            _assertValidTransactionOrdersApproval(
                transaction,
                orders,
                txOrigin,
                transactionSignature,
                approvalSignatures
            );
        }
    }

    /// @dev Decodes the orders from Exchange calldata representing any fill method.
    /// @param data Exchange calldata representing a fill method.
    /// @return orders The orders from the Exchange calldata.
    function decodeOrdersFromFillData(bytes memory data)
        public
        pure
        returns (LibOrder.Order[] memory orders)
    {
        bytes4 selector = data.readBytes4(0);
        if (
            selector == IExchange(address(0)).fillOrder.selector ||
            selector == IExchange(address(0)).fillOrKillOrder.selector
        ) {
            // Decode single order
            (LibOrder.Order memory order) = abi.decode(
                data.slice(4, data.length),
                (LibOrder.Order)
            );
            orders = new LibOrder.Order[](1);
            orders[0] = order;
        } else if (
            selector == IExchange(address(0)).batchFillOrders.selector ||
            selector == IExchange(address(0)).batchFillOrdersNoThrow.selector ||
            selector == IExchange(address(0)).batchFillOrKillOrders.selector ||
            selector == IExchange(address(0)).marketBuyOrdersNoThrow.selector ||
            selector == IExchange(address(0)).marketBuyOrdersFillOrKill.selector ||
            selector == IExchange(address(0)).marketSellOrdersNoThrow.selector ||
            selector == IExchange(address(0)).marketSellOrdersFillOrKill.selector
        ) {
            // Decode all orders
            // solhint-disable indent
            (orders) = abi.decode(
                data.slice(4, data.length),
                (LibOrder.Order[])
            );
        } else if (
            selector == IExchange(address(0)).matchOrders.selector ||
            selector == IExchange(address(0)).matchOrdersWithMaximalFill.selector
        ) {
            // Decode left and right orders
            (LibOrder.Order memory leftOrder, LibOrder.Order memory rightOrder) = abi.decode(
                data.slice(4, data.length),
                (LibOrder.Order, LibOrder.Order)
            );

            // Create array of orders
            orders = new LibOrder.Order[](2);
            orders[0] = leftOrder;
            orders[1] = rightOrder;
        }
        return orders;
    }

    /// @dev Validates that the feeRecipients of a batch of order have approved a 0x transaction.
    /// @param transaction 0x transaction containing salt, signerAddress, and data.
    /// @param orders Array of order structs containing order specifications.
    /// @param txOrigin Required signer of Ethereum transaction calling this function.
    /// @param transactionSignature Proof that the transaction has been signed by the signer.
    /// @param approvalSignatures Array of signatures that correspond to the feeRecipients of each order.
    function _assertValidTransactionOrdersApproval(
        LibZeroExTransaction.ZeroExTransaction memory transaction,
        LibOrder.Order[] memory orders,
        address txOrigin,
        bytes memory transactionSignature,
        bytes[] memory approvalSignatures
    )
        internal
        view
    {
        // Verify that Ethereum tx signer is the same as the approved txOrigin
        if (tx.origin != txOrigin) {
            revert();
        }

        // Hash 0x transaction
        bytes32 transactionHash = LibZeroExTransaction.getTypedDataHash(transaction, EIP712_EXCHANGE_DOMAIN_HASH);

        // Create empty list of approval signers
        address[] memory approvalSignerAddresses = new address[](0);

        uint256 signaturesLength = approvalSignatures.length;
        for (uint256 i = 0; i != signaturesLength; i++) {
            // Create approval message
            CoordinatorApproval memory approval = CoordinatorApproval({
                txOrigin: txOrigin,
                transactionHash: transactionHash,
                transactionSignature: transactionSignature
            });

            // Hash approval message and recover signer address
            bytes32 approvalHash = getCoordinatorApprovalHash(approval);
            address approvalSignerAddress = getSignerAddress(approvalHash, approvalSignatures[i]);

            // Add approval signer to list of signers
            approvalSignerAddresses = approvalSignerAddresses.append(approvalSignerAddress);
        }

        // Ethereum transaction signer gives implicit signature of approval
        approvalSignerAddresses = approvalSignerAddresses.append(tx.origin);

        uint256 ordersLength = orders.length;
        for (uint256 i = 0; i != ordersLength; i++) {
            // Do not check approval if the order's senderAddress is null
            if (orders[i].senderAddress == address(0)) {
                continue;
            }

            // Ensure feeRecipient of order has approved this 0x transaction
            address approverAddress = orders[i].feeRecipientAddress;
            bool isOrderApproved = approvalSignerAddresses.contains(approverAddress);
            if (!isOrderApproved) {
                revert();
            }
        }
    }
}
