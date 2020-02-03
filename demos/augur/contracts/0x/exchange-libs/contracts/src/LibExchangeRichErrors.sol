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


import "./LibOrder.sol";


library LibExchangeRichErrors {

    enum AssetProxyDispatchErrorCodes {
        INVALID_ASSET_DATA_LENGTH,
        UNKNOWN_ASSET_PROXY
    }

    enum BatchMatchOrdersErrorCodes {
        ZERO_LEFT_ORDERS,
        ZERO_RIGHT_ORDERS,
        INVALID_LENGTH_LEFT_SIGNATURES,
        INVALID_LENGTH_RIGHT_SIGNATURES
    }

    enum ExchangeContextErrorCodes {
        INVALID_MAKER,
        INVALID_TAKER,
        INVALID_SENDER
    }

    enum FillErrorCodes {
        INVALID_TAKER_AMOUNT,
        TAKER_OVERPAY,
        OVERFILL,
        INVALID_FILL_PRICE
    }

    enum SignatureErrorCodes {
        BAD_ORDER_SIGNATURE,
        BAD_TRANSACTION_SIGNATURE,
        INVALID_LENGTH,
        UNSUPPORTED,
        ILLEGAL,
        INAPPROPRIATE_SIGNATURE_TYPE,
        INVALID_SIGNER
    }

    enum TransactionErrorCodes {
        ALREADY_EXECUTED,
        EXPIRED
    }

    enum IncompleteFillErrorCode {
        INCOMPLETE_MARKET_BUY_ORDERS,
        INCOMPLETE_MARKET_SELL_ORDERS,
        INCOMPLETE_FILL_ORDER
    }

    // bytes4(keccak256("SignatureError(uint8,bytes32,address,bytes)"))
    bytes4 internal constant SIGNATURE_ERROR_SELECTOR =
        0x7e5a2318;

    // bytes4(keccak256("SignatureValidatorNotApprovedError(address,address)"))
    bytes4 internal constant SIGNATURE_VALIDATOR_NOT_APPROVED_ERROR_SELECTOR =
        0xa15c0d06;

    // bytes4(keccak256("EIP1271SignatureError(address,bytes,bytes,bytes)"))
    bytes4 internal constant EIP1271_SIGNATURE_ERROR_SELECTOR =
        0x5bd0428d;

    // bytes4(keccak256("SignatureWalletError(bytes32,address,bytes,bytes)"))
    bytes4 internal constant SIGNATURE_WALLET_ERROR_SELECTOR =
        0x1b8388f7;

    // bytes4(keccak256("OrderStatusError(bytes32,uint8)"))
    bytes4 internal constant ORDER_STATUS_ERROR_SELECTOR =
        0xfdb6ca8d;

    // bytes4(keccak256("ExchangeInvalidContextError(uint8,bytes32,address)"))
    bytes4 internal constant EXCHANGE_INVALID_CONTEXT_ERROR_SELECTOR =
        0xe53c76c8;

    // bytes4(keccak256("FillError(uint8,bytes32)"))
    bytes4 internal constant FILL_ERROR_SELECTOR =
        0xe94a7ed0;

    // bytes4(keccak256("OrderEpochError(address,address,uint256)"))
    bytes4 internal constant ORDER_EPOCH_ERROR_SELECTOR =
        0x4ad31275;

    // bytes4(keccak256("AssetProxyExistsError(bytes4,address)"))
    bytes4 internal constant ASSET_PROXY_EXISTS_ERROR_SELECTOR =
        0x11c7b720;

    // bytes4(keccak256("AssetProxyDispatchError(uint8,bytes32,bytes)"))
    bytes4 internal constant ASSET_PROXY_DISPATCH_ERROR_SELECTOR =
        0x488219a6;

    // bytes4(keccak256("AssetProxyTransferError(bytes32,bytes,bytes)"))
    bytes4 internal constant ASSET_PROXY_TRANSFER_ERROR_SELECTOR =
        0x4678472b;

    // bytes4(keccak256("NegativeSpreadError(bytes32,bytes32)"))
    bytes4 internal constant NEGATIVE_SPREAD_ERROR_SELECTOR =
        0xb6555d6f;

    // bytes4(keccak256("TransactionError(uint8,bytes32)"))
    bytes4 internal constant TRANSACTION_ERROR_SELECTOR =
        0xf5985184;

    // bytes4(keccak256("TransactionExecutionError(bytes32,bytes)"))
    bytes4 internal constant TRANSACTION_EXECUTION_ERROR_SELECTOR =
        0x20d11f61;
    
    // bytes4(keccak256("TransactionGasPriceError(bytes32,uint256,uint256)"))
    bytes4 internal constant TRANSACTION_GAS_PRICE_ERROR_SELECTOR =
        0xa26dac09;

    // bytes4(keccak256("TransactionInvalidContextError(bytes32,address)"))
    bytes4 internal constant TRANSACTION_INVALID_CONTEXT_ERROR_SELECTOR =
        0xdec4aedf;

    // bytes4(keccak256("IncompleteFillError(uint8,uint256,uint256)"))
    bytes4 internal constant INCOMPLETE_FILL_ERROR_SELECTOR =
        0x18e4b141;

    // bytes4(keccak256("BatchMatchOrdersError(uint8)"))
    bytes4 internal constant BATCH_MATCH_ORDERS_ERROR_SELECTOR =
        0xd4092f4f;

    // bytes4(keccak256("PayProtocolFeeError(bytes32,uint256,address,address,bytes)"))
    bytes4 internal constant PAY_PROTOCOL_FEE_ERROR_SELECTOR =
        0x87cb1e75;

    // solhint-disable func-name-mixedcase
    function SignatureErrorSelector()
        internal
        pure
        returns (bytes4)
    {
        return SIGNATURE_ERROR_SELECTOR;
    }

    function SignatureValidatorNotApprovedErrorSelector()
        internal
        pure
        returns (bytes4)
    {
        return SIGNATURE_VALIDATOR_NOT_APPROVED_ERROR_SELECTOR;
    }

    function EIP1271SignatureErrorSelector()
        internal
        pure
        returns (bytes4)
    {
        return EIP1271_SIGNATURE_ERROR_SELECTOR;
    }

    function SignatureWalletErrorSelector()
        internal
        pure
        returns (bytes4)
    {
        return SIGNATURE_WALLET_ERROR_SELECTOR;
    }

    function OrderStatusErrorSelector()
        internal
        pure
        returns (bytes4)
    {
        return ORDER_STATUS_ERROR_SELECTOR;
    }

    function ExchangeInvalidContextErrorSelector()
        internal
        pure
        returns (bytes4)
    {
        return EXCHANGE_INVALID_CONTEXT_ERROR_SELECTOR;
    }

    function FillErrorSelector()
        internal
        pure
        returns (bytes4)
    {
        return FILL_ERROR_SELECTOR;
    }

    function OrderEpochErrorSelector()
        internal
        pure
        returns (bytes4)
    {
        return ORDER_EPOCH_ERROR_SELECTOR;
    }

    function AssetProxyExistsErrorSelector()
        internal
        pure
        returns (bytes4)
    {
        return ASSET_PROXY_EXISTS_ERROR_SELECTOR;
    }

    function AssetProxyDispatchErrorSelector()
        internal
        pure
        returns (bytes4)
    {
        return ASSET_PROXY_DISPATCH_ERROR_SELECTOR;
    }

    function AssetProxyTransferErrorSelector()
        internal
        pure
        returns (bytes4)
    {
        return ASSET_PROXY_TRANSFER_ERROR_SELECTOR;
    }

    function NegativeSpreadErrorSelector()
        internal
        pure
        returns (bytes4)
    {
        return NEGATIVE_SPREAD_ERROR_SELECTOR;
    }

    function TransactionErrorSelector()
        internal
        pure
        returns (bytes4)
    {
        return TRANSACTION_ERROR_SELECTOR;
    }

    function TransactionExecutionErrorSelector()
        internal
        pure
        returns (bytes4)
    {
        return TRANSACTION_EXECUTION_ERROR_SELECTOR;
    }

    function IncompleteFillErrorSelector()
        internal
        pure
        returns (bytes4)
    {
        return INCOMPLETE_FILL_ERROR_SELECTOR;
    }

    function BatchMatchOrdersErrorSelector()
        internal
        pure
        returns (bytes4)
    {
        return BATCH_MATCH_ORDERS_ERROR_SELECTOR;
    }

    function TransactionGasPriceErrorSelector()
        internal
        pure
        returns (bytes4)
    {
        return TRANSACTION_GAS_PRICE_ERROR_SELECTOR;
    }

    function TransactionInvalidContextErrorSelector()
        internal
        pure
        returns (bytes4)
    {
        return TRANSACTION_INVALID_CONTEXT_ERROR_SELECTOR;
    }

    function PayProtocolFeeErrorSelector()
        internal
        pure
        returns (bytes4)
    {
        return PAY_PROTOCOL_FEE_ERROR_SELECTOR;
    }
    
    function BatchMatchOrdersError(
        BatchMatchOrdersErrorCodes errorCode
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            BATCH_MATCH_ORDERS_ERROR_SELECTOR,
            errorCode
        );
    }

    function SignatureError(
        SignatureErrorCodes errorCode,
        bytes32 hash,
        address signerAddress,
        bytes memory signature
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            SIGNATURE_ERROR_SELECTOR,
            errorCode,
            hash,
            signerAddress,
            signature
        );
    }

    function SignatureValidatorNotApprovedError(
        address signerAddress,
        address validatorAddress
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            SIGNATURE_VALIDATOR_NOT_APPROVED_ERROR_SELECTOR,
            signerAddress,
            validatorAddress
        );
    }

    function EIP1271SignatureError(
        address verifyingContractAddress,
        bytes memory data,
        bytes memory signature,
        bytes memory errorData
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            EIP1271_SIGNATURE_ERROR_SELECTOR,
            verifyingContractAddress,
            data,
            signature,
            errorData
        );
    }

    function SignatureWalletError(
        bytes32 hash,
        address walletAddress,
        bytes memory signature,
        bytes memory errorData
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            SIGNATURE_WALLET_ERROR_SELECTOR,
            hash,
            walletAddress,
            signature,
            errorData
        );
    }

    function OrderStatusError(
        bytes32 orderHash,
        LibOrder.OrderStatus orderStatus
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            ORDER_STATUS_ERROR_SELECTOR,
            orderHash,
            orderStatus
        );
    }

    function ExchangeInvalidContextError(
        ExchangeContextErrorCodes errorCode,
        bytes32 orderHash,
        address contextAddress
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            EXCHANGE_INVALID_CONTEXT_ERROR_SELECTOR,
            errorCode,
            orderHash,
            contextAddress
        );
    }

    function FillError(
        FillErrorCodes errorCode,
        bytes32 orderHash
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            FILL_ERROR_SELECTOR,
            errorCode,
            orderHash
        );
    }

    function OrderEpochError(
        address makerAddress,
        address orderSenderAddress,
        uint256 currentEpoch
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            ORDER_EPOCH_ERROR_SELECTOR,
            makerAddress,
            orderSenderAddress,
            currentEpoch
        );
    }

    function AssetProxyExistsError(
        bytes4 assetProxyId,
        address assetProxyAddress
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            ASSET_PROXY_EXISTS_ERROR_SELECTOR,
            assetProxyId,
            assetProxyAddress
        );
    }

    function AssetProxyDispatchError(
        AssetProxyDispatchErrorCodes errorCode,
        bytes32 orderHash,
        bytes memory assetData
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            ASSET_PROXY_DISPATCH_ERROR_SELECTOR,
            errorCode,
            orderHash,
            assetData
        );
    }

    function AssetProxyTransferError(
        bytes32 orderHash,
        bytes memory assetData,
        bytes memory errorData
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            ASSET_PROXY_TRANSFER_ERROR_SELECTOR,
            orderHash,
            assetData,
            errorData
        );
    }

    function NegativeSpreadError(
        bytes32 leftOrderHash,
        bytes32 rightOrderHash
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            NEGATIVE_SPREAD_ERROR_SELECTOR,
            leftOrderHash,
            rightOrderHash
        );
    }

    function TransactionError(
        TransactionErrorCodes errorCode,
        bytes32 transactionHash
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            TRANSACTION_ERROR_SELECTOR,
            errorCode,
            transactionHash
        );
    }

    function TransactionExecutionError(
        bytes32 transactionHash,
        bytes memory errorData
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            TRANSACTION_EXECUTION_ERROR_SELECTOR,
            transactionHash,
            errorData
        );
    }

    function TransactionGasPriceError(
        bytes32 transactionHash,
        uint256 actualGasPrice,
        uint256 requiredGasPrice
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            TRANSACTION_GAS_PRICE_ERROR_SELECTOR,
            transactionHash,
            actualGasPrice,
            requiredGasPrice
        );
    }

    function TransactionInvalidContextError(
        bytes32 transactionHash,
        address currentContextAddress
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            TRANSACTION_INVALID_CONTEXT_ERROR_SELECTOR,
            transactionHash,
            currentContextAddress
        );
    }

    function IncompleteFillError(
        IncompleteFillErrorCode errorCode,
        uint256 expectedAssetFillAmount,
        uint256 actualAssetFillAmount
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            INCOMPLETE_FILL_ERROR_SELECTOR,
            errorCode,
            expectedAssetFillAmount,
            actualAssetFillAmount
        );
    }

    function PayProtocolFeeError(
        bytes32 orderHash,
        uint256 protocolFee,
        address makerAddress,
        address takerAddress,
        bytes memory errorData
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            PAY_PROTOCOL_FEE_ERROR_SELECTOR,
            orderHash,
            protocolFee,
            makerAddress,
            takerAddress,
            errorData
        );
    }
}
