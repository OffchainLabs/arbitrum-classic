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
import "../../../utils/contracts/src/LibEIP1271.sol";

import "../../../exchange-libs/contracts/src/LibOrder.sol";
import "../../../exchange-libs/contracts/src/LibZeroExTransaction.sol";
import "../../../exchange-libs/contracts/src/LibEIP712ExchangeDomain.sol";
import "../../../exchange-libs/contracts/src/LibExchangeRichErrors.sol";
import "./interfaces/IWallet.sol";
import "./interfaces/IEIP1271Wallet.sol";
import "./interfaces/ISignatureValidator.sol";
import "./interfaces/IEIP1271Data.sol";
import "./MixinTransactions.sol";


contract MixinSignatureValidator is
    LibEIP712ExchangeDomain,
    LibEIP1271,
    ISignatureValidator,
    MixinTransactions
{
    using LibBytes for bytes;
    using LibOrder for LibOrder.Order;
    using LibZeroExTransaction for LibZeroExTransaction.ZeroExTransaction;

    // Magic bytes to be returned by `Wallet` signature type validators.
    // bytes4(keccak256("isValidWalletSignature(bytes32,address,bytes)"))
    bytes4 private constant LEGACY_WALLET_MAGIC_VALUE = 0xb0671381;

    /// @dev Mapping of hash => signer => signed
    /// @param 0 Order hash.
    /// @param 1 Signer address.
    /// @return 0 Whether the hash is presigned.
    mapping (bytes32 => mapping (address => bool)) public preSigned;

    /// @dev Mapping of signer => validator => approved
    /// @param 0 Signer address.
    /// @param 1 Signature validator address.
    /// @return 0 Whether the validator is allowed to validate on behalf of the signer.
    mapping (address => mapping (address => bool)) public allowedValidators;

    /// @dev Approves a hash on-chain.
    ///      After presigning a hash, the preSign signature type will become valid for that hash and signer.
    /// @param hash Any 32-byte hash.
    function preSign(bytes32 hash)
        external
        payable
        refundFinalBalanceNoReentry
    {
        address signerAddress = _getCurrentContextAddress();
        preSigned[hash][signerAddress] = true;
    }

    /// @dev Approves/unnapproves a Validator contract to verify signatures on signer's behalf
    ///      using the `Validator` signature type.
    /// @param validatorAddress Address of Validator contract.
    /// @param approval Approval or disapproval of  Validator contract.
    function setSignatureValidatorApproval(
        address validatorAddress,
        bool approval
    )
        external
        payable
        refundFinalBalanceNoReentry
    {
        address signerAddress = _getCurrentContextAddress();
        allowedValidators[signerAddress][validatorAddress] = approval;
        emit SignatureValidatorApproval(
            signerAddress,
            validatorAddress,
            approval
        );
    }

    /// @dev Verifies that a hash has been signed by the given signer.
    /// @param hash Any 32-byte hash.
    /// @param signerAddress Address that should have signed the given hash.
    /// @param signature Proof that the hash has been signed by signer.
    /// @return isValid `true` if the signature is valid for the given hash and signer.
    function isValidHashSignature(
        bytes32 hash,
        address signerAddress,
        bytes memory signature
    )
        public
        view
        returns (bool isValid)
    {
        SignatureType signatureType = _readValidSignatureType(
            hash,
            signerAddress,
            signature
        );
        // Only hash-compatible signature types can be handled by this
        // function.
        if (
            signatureType == SignatureType.Validator ||
            signatureType == SignatureType.EIP1271Wallet
        ) {
            revert();
        }
        isValid = _validateHashSignatureTypes(
            signatureType,
            hash,
            signerAddress,
            signature
        );
        return isValid;
    }

    /// @dev Verifies that a signature for an order is valid.
    /// @param order The order.
    /// @param signature Proof that the order has been signed by signer.
    /// @return isValid `true` if the signature is valid for the given order and signer.
    function isValidOrderSignature(
        LibOrder.Order memory order,
        bytes memory signature
    )
        public
        view
        returns (bool isValid)
    {
        bytes32 orderHash = order.getTypedDataHash(EIP712_EXCHANGE_DOMAIN_HASH);
        isValid = _isValidOrderWithHashSignature(
            order,
            orderHash,
            signature
        );
        return isValid;
    }

    /// @dev Verifies that a signature for a transaction is valid.
    /// @param transaction The transaction.
    /// @param signature Proof that the order has been signed by signer.
    /// @return isValid `true` if the signature is valid for the given transaction and signer.
    function isValidTransactionSignature(
        LibZeroExTransaction.ZeroExTransaction memory transaction,
        bytes memory signature
    )
        public
        view
        returns (bool isValid)
    {
        bytes32 transactionHash = transaction.getTypedDataHash(EIP712_EXCHANGE_DOMAIN_HASH);
        isValid = _isValidTransactionWithHashSignature(
            transaction,
            transactionHash,
            signature
        );
        return isValid;
    }

    /// @dev Verifies that an order, with provided order hash, has been signed
    ///      by the given signer.
    /// @param order The order.
    /// @param orderHash The hash of the order.
    /// @param signature Proof that the hash has been signed by signer.
    /// @return isValid True if the signature is valid for the given order and signer.
    function _isValidOrderWithHashSignature(
        LibOrder.Order memory order,
        bytes32 orderHash,
        bytes memory signature
    )
        internal
        view
        returns (bool isValid)
    {
        address signerAddress = order.makerAddress;
        SignatureType signatureType = _readValidSignatureType(
            orderHash,
            signerAddress,
            signature
        );
        if (signatureType == SignatureType.Validator) {
            // The entire order is verified by a validator contract.
            isValid = _validateBytesWithValidator(
                _encodeEIP1271OrderWithHash(order, orderHash),
                orderHash,
                signerAddress,
                signature
            );
        } else if (signatureType == SignatureType.EIP1271Wallet) {
            // The entire order is verified by a wallet contract.
            isValid = _validateBytesWithWallet(
                _encodeEIP1271OrderWithHash(order, orderHash),
                signerAddress,
                signature
            );
        } else {
            // Otherwise, it's one of the hash-only signature types.
            isValid = _validateHashSignatureTypes(
                signatureType,
                orderHash,
                signerAddress,
                signature
            );
        }
        return isValid;
    }

    /// @dev Verifies that a transaction, with provided order hash, has been signed
    ///      by the given signer.
    /// @param transaction The transaction.
    /// @param transactionHash The hash of the transaction.
    /// @param signature Proof that the hash has been signed by signer.
    /// @return isValid True if the signature is valid for the given transaction and signer.
    function _isValidTransactionWithHashSignature(
        LibZeroExTransaction.ZeroExTransaction memory transaction,
        bytes32 transactionHash,
        bytes memory signature
    )
        internal
        view
        returns (bool isValid)
    {
        address signerAddress = transaction.signerAddress;
        SignatureType signatureType = _readValidSignatureType(
            transactionHash,
            signerAddress,
            signature
        );
        if (signatureType == SignatureType.Validator) {
            // The entire transaction is verified by a validator contract.
            isValid = _validateBytesWithValidator(
                _encodeEIP1271TransactionWithHash(transaction, transactionHash),
                transactionHash,
                signerAddress,
                signature
            );
        } else if (signatureType == SignatureType.EIP1271Wallet) {
            // The entire transaction is verified by a wallet contract.
            isValid = _validateBytesWithWallet(
                _encodeEIP1271TransactionWithHash(transaction, transactionHash),
                signerAddress,
                signature
            );
        } else {
            // Otherwise, it's one of the hash-only signature types.
            isValid = _validateHashSignatureTypes(
                signatureType,
                transactionHash,
                signerAddress,
                signature
            );
        }
        return isValid;
    }

    /// Validates a hash-only signature type
    /// (anything but `Validator` and `EIP1271Wallet`).
    function _validateHashSignatureTypes(
        SignatureType signatureType,
        bytes32 hash,
        address signerAddress,
        bytes memory signature
    )
        private
        view
        returns (bool isValid)
    {
        // Always invalid signature.
        // Like Illegal, this is always implicitly available and therefore
        // offered explicitly. It can be implicitly created by providing
        // a correctly formatted but incorrect signature.
        if (signatureType == SignatureType.Invalid) {
            if (signature.length != 1) {
                revert();
            }
            isValid = false;

        // Signature using EIP712
        } else if (signatureType == SignatureType.EIP712) {
            if (signature.length != 66) {
                revert();
            }
            uint8 v = uint8(signature[0]);
            bytes32 r = signature.readBytes32(1);
            bytes32 s = signature.readBytes32(33);
            address recovered = ecrecover(
                hash,
                v,
                r,
                s
            );
            isValid = signerAddress == recovered;

        // Signed using web3.eth_sign
        } else if (signatureType == SignatureType.EthSign) {
            if (signature.length != 66) {
                revert();
            }
            uint8 v = uint8(signature[0]);
            bytes32 r = signature.readBytes32(1);
            bytes32 s = signature.readBytes32(33);
            address recovered = ecrecover(
                keccak256(abi.encodePacked(
                    "\x19Ethereum Signed Message:\n32",
                    hash
                )),
                v,
                r,
                s
            );
            isValid = signerAddress == recovered;

        // Signature verified by wallet contract.
        } else if (signatureType == SignatureType.Wallet) {
            isValid = _validateHashWithWallet(
                hash,
                signerAddress,
                signature
            );

        // Otherwise, signatureType == SignatureType.PreSigned
        } else {
            assert(signatureType == SignatureType.PreSigned);
            // Signer signed hash previously using the preSign function.
            isValid = preSigned[hash][signerAddress];
        }
        return isValid;
    }

    /// @dev Reads the `SignatureType` from a signature with minimal validation.
    function _readSignatureType(
        bytes32 hash,
        address signerAddress,
        bytes memory signature
    )
        private
        pure
        returns (SignatureType)
    {
        if (signature.length == 0) {
            revert();
        }
        return SignatureType(uint8(signature[signature.length - 1]));
    }

    /// @dev Reads the `SignatureType` from the end of a signature and validates it.
    function _readValidSignatureType(
        bytes32 hash,
        address signerAddress,
        bytes memory signature
    )
        private
        pure
        returns (SignatureType signatureType)
    {
        // Read the signatureType from the signature
        signatureType = _readSignatureType(
            hash,
            signerAddress,
            signature
        );

        // Disallow address zero because ecrecover() returns zero on failure.
        if (signerAddress == address(0)) {
            revert();
        }

        // Ensure signature is supported
        if (uint8(signatureType) >= uint8(SignatureType.NSignatureTypes)) {
            revert();
        }

        // Always illegal signature.
        // This is always an implicit option since a signer can create a
        // signature array with invalid type or length. We may as well make
        // it an explicit option. This aids testing and analysis. It is
        // also the initialization value for the enum type.
        if (signatureType == SignatureType.Illegal) {
            revert();
        }

        return signatureType;
    }

    /// @dev ABI encodes an order and hash with a selector to be passed into
    ///      an EIP1271 compliant `isValidSignature` function.
    function _encodeEIP1271OrderWithHash(
        LibOrder.Order memory order,
        bytes32 orderHash
    )
        private
        pure
        returns (bytes memory encoded)
    {
        return abi.encodeWithSelector(
            IEIP1271Data(address(0)).OrderWithHash.selector,
            order,
            orderHash
        );
    }

    /// @dev ABI encodes a transaction and hash with a selector to be passed into
    ///      an EIP1271 compliant `isValidSignature` function.
    function _encodeEIP1271TransactionWithHash(
        LibZeroExTransaction.ZeroExTransaction memory transaction,
        bytes32 transactionHash
    )
        private
        pure
        returns (bytes memory encoded)
    {
        return abi.encodeWithSelector(
            IEIP1271Data(address(0)).ZeroExTransactionWithHash.selector,
            transaction,
            transactionHash
        );
    }

    /// @dev Verifies a hash and signature using logic defined by Wallet contract.
    /// @param hash Any 32 byte hash.
    /// @param walletAddress Address that should have signed the given hash
    ///                      and defines its own signature verification method.
    /// @param signature Proof that the hash has been signed by signer.
    /// @return True if the signature is validated by the Wallet.
    function _validateHashWithWallet(
        bytes32 hash,
        address walletAddress,
        bytes memory signature
    )
        private
        view
        returns (bool)
    {
        // Backup length of signature
        uint256 signatureLength = signature.length;
        // Temporarily remove signatureType byte from end of signature
        signature.writeLength(signatureLength - 1);
        // Encode the call data.
        bytes memory callData = abi.encodeWithSelector(
            IWallet(address(0)).isValidSignature.selector,
            hash,
            signature
        );
        // Restore the original signature length
        signature.writeLength(signatureLength);
        // Static call the verification function.
        (bool didSucceed, bytes memory returnData) = walletAddress.staticcall(callData);
        // Return the validity of the signature if the call was successful
        if (didSucceed && returnData.length == 32) {
            return returnData.readBytes4(0) == LEGACY_WALLET_MAGIC_VALUE;
        }
        // Revert if the call was unsuccessful
        revert();
    }

    /// @dev Verifies arbitrary data and a signature via an EIP1271 Wallet
    ///      contract, where the wallet address is also the signer address.
    /// @param data Arbitrary signed data.
    /// @param walletAddress Contract that will verify the data and signature.
    /// @param signature Proof that the data has been signed by signer.
    /// @return isValid True if the signature is validated by the Wallet.
    function _validateBytesWithWallet(
        bytes memory data,
        address walletAddress,
        bytes memory signature
    )
        private
        view
        returns (bool isValid)
    {
        isValid = _staticCallEIP1271WalletWithReducedSignatureLength(
            walletAddress,
            data,
            signature,
            1  // The last byte of the signature (signatureType) is removed before making the staticcall
        );
        return isValid;
    }

    /// @dev Verifies arbitrary data and a signature via an EIP1271 contract
    ///      whose address is encoded in the signature.
    /// @param data Arbitrary signed data.
    /// @param hash The hash associated with the data.
    /// @param signerAddress Address that should have signed the given hash.
    /// @param signature Proof that the data has been signed by signer.
    /// @return isValid True if the signature is validated by the validator contract.
    function _validateBytesWithValidator(
        bytes memory data,
        bytes32 hash,
        address signerAddress,
        bytes memory signature
    )
        private
        view
        returns (bool isValid)
    {
        uint256 signatureLength = signature.length;
        if (signatureLength < 21) {
            revert();
        }
        // The validator address is appended to the signature before the signatureType.
        // Read the validator address from the signature.
        address validatorAddress = signature.readAddress(signatureLength - 21);
        // Ensure signer has approved validator.
        if (!allowedValidators[signerAddress][validatorAddress]) {
            revert();
        }
        isValid = _staticCallEIP1271WalletWithReducedSignatureLength(
            validatorAddress,
            data,
            signature,
            21  // The last 21 bytes of the signature (validatorAddress + signatureType) are removed before making the staticcall
        );
        return isValid;
    }

    /// @dev Performs a staticcall to an EIP1271 compiant `isValidSignature` function and validates the output.
    /// @param verifyingContractAddress Address of EIP1271Wallet or Validator contract.
    /// @param data Arbitrary signed data.
    /// @param signature Proof that the hash has been signed by signer. Bytes will be temporarily be popped
    ///                  off of the signature before calling `isValidSignature`.
    /// @param ignoredSignatureBytesLen The amount of bytes that will be temporarily popped off the the signature.
    /// @return The validity of the signature.
    function _staticCallEIP1271WalletWithReducedSignatureLength(
        address verifyingContractAddress,
        bytes memory data,
        bytes memory signature,
        uint256 ignoredSignatureBytesLen
    )
        private
        view
        returns (bool)
    {
        // Backup length of the signature
        uint256 signatureLength = signature.length;
        // Temporarily remove bytes from signature end
        signature.writeLength(signatureLength - ignoredSignatureBytesLen);
        bytes memory callData = abi.encodeWithSelector(
            IEIP1271Wallet(address(0)).isValidSignature.selector,
            data,
            signature
        );
        // Restore original signature length
        signature.writeLength(signatureLength);
        // Static call the verification function
        (bool didSucceed, bytes memory returnData) = verifyingContractAddress.staticcall(callData);
        // Return the validity of the signature if the call was successful
        if (didSucceed && returnData.length == 32) {
            return returnData.readBytes4(0) == EIP1271_MAGIC_VALUE;
        }
        // Revert if the call was unsuccessful
        revert();
    }
}
