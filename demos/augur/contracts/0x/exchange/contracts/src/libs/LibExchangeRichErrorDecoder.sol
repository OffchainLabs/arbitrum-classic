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

import "../../../../exchange-libs/contracts/src/LibOrder.sol";
import "../../../../exchange-libs/contracts/src/LibExchangeRichErrors.sol";
import "../../../../utils/contracts/src/LibBytes.sol";


contract LibExchangeRichErrorDecoder {

    using LibBytes for bytes;

    /// @dev Decompose an ABI-encoded SignatureError.
    /// @param encoded ABI-encoded revert error.
    /// @return errorCode The error code.
    /// @return signerAddress The expected signer of the hash.
    /// @return signature The full signature.
    function decodeSignatureError(bytes memory encoded)
        public
        pure
        returns (
            LibExchangeRichErrors.SignatureErrorCodes errorCode,
            bytes32 hash,
            address signerAddress,
            bytes memory signature
        )
    {
        _assertSelectorBytes(encoded, LibExchangeRichErrors.SignatureErrorSelector());
        uint8 _errorCode;
        (_errorCode, hash, signerAddress, signature) = abi.decode(
            encoded.sliceDestructive(4, encoded.length),
            (uint8, bytes32, address, bytes)
        );
        errorCode = LibExchangeRichErrors.SignatureErrorCodes(_errorCode);
    }

    /// @dev Decompose an ABI-encoded SignatureValidatorError.
    /// @param encoded ABI-encoded revert error.
    /// @return signerAddress The expected signer of the hash.
    /// @return signature The full signature bytes.
    /// @return errorData The revert data thrown by the validator contract.
    function decodeEIP1271SignatureError(bytes memory encoded)
        public
        pure
        returns (
            address verifyingContractAddress,
            bytes memory data,
            bytes memory signature,
            bytes memory errorData
        )
    {
        _assertSelectorBytes(encoded, LibExchangeRichErrors.EIP1271SignatureErrorSelector());
        (verifyingContractAddress, data, signature, errorData) = abi.decode(
            encoded.sliceDestructive(4, encoded.length),
            (address, bytes, bytes, bytes)
        );
    }

    /// @dev Decompose an ABI-encoded SignatureValidatorNotApprovedError.
    /// @param encoded ABI-encoded revert error.
    /// @return signerAddress The expected signer of the hash.
    /// @return validatorAddress The expected validator.
    function decodeSignatureValidatorNotApprovedError(bytes memory encoded)
        public
        pure
        returns (
            address signerAddress,
            address validatorAddress
        )
    {
        _assertSelectorBytes(encoded, LibExchangeRichErrors.SignatureValidatorNotApprovedErrorSelector());
        (signerAddress, validatorAddress) = abi.decode(
            encoded.sliceDestructive(4, encoded.length),
            (address, address)
        );
    }

    /// @dev Decompose an ABI-encoded SignatureWalletError.
    /// @param encoded ABI-encoded revert error.
    /// @return errorCode The error code.
    /// @return signerAddress The expected signer of the hash.
    /// @return signature The full signature bytes.
    /// @return errorData The revert data thrown by the validator contract.
    function decodeSignatureWalletError(bytes memory encoded)
        public
        pure
        returns (
            bytes32 hash,
            address signerAddress,
            bytes memory signature,
            bytes memory errorData
        )
    {
        _assertSelectorBytes(encoded, LibExchangeRichErrors.SignatureWalletErrorSelector());
        (hash, signerAddress, signature, errorData) = abi.decode(
            encoded.sliceDestructive(4, encoded.length),
            (bytes32, address, bytes, bytes)
        );
    }

    /// @dev Decompose an ABI-encoded OrderStatusError.
    /// @param encoded ABI-encoded revert error.
    /// @return orderHash The order hash.
    /// @return orderStatus The order status.
    function decodeOrderStatusError(bytes memory encoded)
        public
        pure
        returns (
            bytes32 orderHash,
            LibOrder.OrderStatus orderStatus
        )
    {
        _assertSelectorBytes(encoded, LibExchangeRichErrors.OrderStatusErrorSelector());
        uint8 _orderStatus;
        (orderHash, _orderStatus) = abi.decode(
            encoded.sliceDestructive(4, encoded.length),
            (bytes32, uint8)
        );
        orderStatus = LibOrder.OrderStatus(_orderStatus);
    }

    /// @dev Decompose an ABI-encoded OrderStatusError.
    /// @param encoded ABI-encoded revert error.
    /// @return errorCode Error code that corresponds to invalid maker, taker, or sender.
    /// @return orderHash The order hash.
    /// @return contextAddress The maker, taker, or sender address
    function decodeExchangeInvalidContextError(bytes memory encoded)
        public
        pure
        returns (
            LibExchangeRichErrors.ExchangeContextErrorCodes errorCode,
            bytes32 orderHash,
            address contextAddress
        )
    {
        _assertSelectorBytes(encoded, LibExchangeRichErrors.ExchangeInvalidContextErrorSelector());
        uint8 _errorCode;
        (_errorCode, orderHash, contextAddress) = abi.decode(
            encoded.sliceDestructive(4, encoded.length),
            (uint8, bytes32, address)
        );
        errorCode = LibExchangeRichErrors.ExchangeContextErrorCodes(_errorCode);
    }

    /// @dev Decompose an ABI-encoded FillError.
    /// @param encoded ABI-encoded revert error.
    /// @return errorCode The error code.
    /// @return orderHash The order hash.
    function decodeFillError(bytes memory encoded)
        public
        pure
        returns (
            LibExchangeRichErrors.FillErrorCodes errorCode,
            bytes32 orderHash
        )
    {
        _assertSelectorBytes(encoded, LibExchangeRichErrors.FillErrorSelector());
        uint8 _errorCode;
        (_errorCode, orderHash) = abi.decode(
            encoded.sliceDestructive(4, encoded.length),
            (uint8, bytes32)
        );
        errorCode = LibExchangeRichErrors.FillErrorCodes(_errorCode);
    }

    /// @dev Decompose an ABI-encoded OrderEpochError.
    /// @param encoded ABI-encoded revert error.
    /// @return makerAddress The order maker.
    /// @return orderSenderAddress The order sender.
    /// @return currentEpoch The current epoch for the maker.
    function decodeOrderEpochError(bytes memory encoded)
        public
        pure
        returns (
            address makerAddress,
            address orderSenderAddress,
            uint256 currentEpoch
        )
    {
        _assertSelectorBytes(encoded, LibExchangeRichErrors.OrderEpochErrorSelector());
        (makerAddress, orderSenderAddress, currentEpoch) = abi.decode(
            encoded.sliceDestructive(4, encoded.length),
            (address, address, uint256)
        );
    }

    /// @dev Decompose an ABI-encoded AssetProxyExistsError.
    /// @param encoded ABI-encoded revert error.
    /// @return assetProxyId Id of asset proxy.
    /// @return assetProxyAddress The address of the asset proxy.
    function decodeAssetProxyExistsError(bytes memory encoded)
        public
        pure
        returns (
            bytes4 assetProxyId, address assetProxyAddress)
    {
        _assertSelectorBytes(encoded, LibExchangeRichErrors.AssetProxyExistsErrorSelector());
        (assetProxyId, assetProxyAddress) = abi.decode(
            encoded.sliceDestructive(4, encoded.length),
            (bytes4, address)
        );
    }

    /// @dev Decompose an ABI-encoded AssetProxyDispatchError.
    /// @param encoded ABI-encoded revert error.
    /// @return errorCode The error code.
    /// @return orderHash Hash of the order being dispatched.
    /// @return assetData Asset data of the order being dispatched.
    function decodeAssetProxyDispatchError(bytes memory encoded)
        public
        pure
        returns (
            LibExchangeRichErrors.AssetProxyDispatchErrorCodes errorCode,
            bytes32 orderHash,
            bytes memory assetData
        )
    {
        _assertSelectorBytes(encoded, LibExchangeRichErrors.AssetProxyDispatchErrorSelector());
        uint8 _errorCode;
        (_errorCode, orderHash, assetData) = abi.decode(
            encoded.sliceDestructive(4, encoded.length),
            (uint8, bytes32, bytes)
        );
        errorCode = LibExchangeRichErrors.AssetProxyDispatchErrorCodes(_errorCode);
    }

    /// @dev Decompose an ABI-encoded AssetProxyTransferError.
    /// @param encoded ABI-encoded revert error.
    /// @return orderHash Hash of the order being dispatched.
    /// @return assetData Asset data of the order being dispatched.
    /// @return errorData ABI-encoded revert data from the asset proxy.
    function decodeAssetProxyTransferError(bytes memory encoded)
        public
        pure
        returns (
            bytes32 orderHash,
            bytes memory assetData,
            bytes memory errorData
        )
    {
        _assertSelectorBytes(encoded, LibExchangeRichErrors.AssetProxyTransferErrorSelector());
        (orderHash, assetData, errorData) = abi.decode(
            encoded.sliceDestructive(4, encoded.length),
            (bytes32, bytes, bytes)
        );
    }

    /// @dev Decompose an ABI-encoded NegativeSpreadError.
    /// @param encoded ABI-encoded revert error.
    /// @return leftOrderHash Hash of the left order being matched.
    /// @return rightOrderHash Hash of the right order being matched.
    function decodeNegativeSpreadError(bytes memory encoded)
        public
        pure
        returns (
            bytes32 leftOrderHash,
            bytes32 rightOrderHash
        )
    {
        _assertSelectorBytes(encoded, LibExchangeRichErrors.NegativeSpreadErrorSelector());
        (leftOrderHash, rightOrderHash) = abi.decode(
            encoded.sliceDestructive(4, encoded.length),
            (bytes32, bytes32)
        );
    }

    /// @dev Decompose an ABI-encoded TransactionError.
    /// @param encoded ABI-encoded revert error.
    /// @return errorCode The error code.
    /// @return transactionHash Hash of the transaction.
    function decodeTransactionError(bytes memory encoded)
        public
        pure
        returns (
            LibExchangeRichErrors.TransactionErrorCodes errorCode,
            bytes32 transactionHash
        )
    {
        _assertSelectorBytes(encoded, LibExchangeRichErrors.TransactionErrorSelector());
        uint8 _errorCode;
        (_errorCode, transactionHash) = abi.decode(
            encoded.sliceDestructive(4, encoded.length),
            (uint8, bytes32)
        );
        errorCode = LibExchangeRichErrors.TransactionErrorCodes(_errorCode);
    }

    /// @dev Decompose an ABI-encoded TransactionExecutionError.
    /// @param encoded ABI-encoded revert error.
    /// @return transactionHash Hash of the transaction.
    /// @return errorData Error thrown by exeucteTransaction().
    function decodeTransactionExecutionError(bytes memory encoded)
        public
        pure
        returns (
            bytes32 transactionHash,
            bytes memory errorData
        )
    {
        _assertSelectorBytes(encoded, LibExchangeRichErrors.TransactionExecutionErrorSelector());
        (transactionHash, errorData) = abi.decode(
            encoded.sliceDestructive(4, encoded.length),
            (bytes32, bytes)
        );
    }

    /// @dev Decompose an ABI-encoded IncompleteFillError.
    /// @param encoded ABI-encoded revert error.
    /// @return orderHash Hash of the order being filled.
    function decodeIncompleteFillError(bytes memory encoded)
        public
        pure
        returns (
            LibExchangeRichErrors.IncompleteFillErrorCode errorCode,
            uint256 expectedAssetFillAmount,
            uint256 actualAssetFillAmount
        )
    {
        _assertSelectorBytes(encoded, LibExchangeRichErrors.IncompleteFillErrorSelector());
        uint8 _errorCode;
        (_errorCode, expectedAssetFillAmount, actualAssetFillAmount) = abi.decode(
            encoded.sliceDestructive(4, encoded.length),
            (uint8, uint256, uint256)
        );
        errorCode = LibExchangeRichErrors.IncompleteFillErrorCode(_errorCode);
    }

    /// @dev Revert if the leading 4 bytes of `encoded` is not `selector`.
    function _assertSelectorBytes(bytes memory encoded, bytes4 selector)
        private
        pure
    {
        bytes4 actualSelector = LibBytes.readBytes4(encoded, 0);
        require(
            actualSelector == selector,
            "BAD_SELECTOR"
        );
    }
}
