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


import "../interfaces/IStructs.sol";


library LibStakingRichErrors {

    enum OperatorShareErrorCodes {
        OperatorShareTooLarge,
        CanOnlyDecreaseOperatorShare
    }

    enum InitializationErrorCodes {
        MixinSchedulerAlreadyInitialized,
        MixinParamsAlreadyInitialized
    }

    enum InvalidParamValueErrorCodes {
        InvalidCobbDouglasAlpha,
        InvalidRewardDelegatedStakeWeight,
        InvalidMaximumMakersInPool,
        InvalidMinimumPoolStake,
        InvalidEpochDuration
    }

    enum ExchangeManagerErrorCodes {
        ExchangeAlreadyRegistered,
        ExchangeNotRegistered
    }

    // bytes4(keccak256("OnlyCallableByExchangeError(address)"))
    bytes4 internal constant ONLY_CALLABLE_BY_EXCHANGE_ERROR_SELECTOR =
        0xb56d2df0;

    // bytes4(keccak256("ExchangeManagerError(uint8,address)"))
    bytes4 internal constant EXCHANGE_MANAGER_ERROR_SELECTOR =
        0xb9588e43;

    // bytes4(keccak256("InsufficientBalanceError(uint256,uint256)"))
    bytes4 internal constant INSUFFICIENT_BALANCE_ERROR_SELECTOR =
        0x84c8b7c9;

    // bytes4(keccak256("OnlyCallableByPoolOperatorError(address,bytes32)"))
    bytes4 internal constant ONLY_CALLABLE_BY_POOL_OPERATOR_ERROR_SELECTOR =
        0x82ded785;

    // bytes4(keccak256("BlockTimestampTooLowError(uint256,uint256)"))
    bytes4 internal constant BLOCK_TIMESTAMP_TOO_LOW_ERROR_SELECTOR =
        0xa6bcde47;

    // bytes4(keccak256("OnlyCallableByStakingContractError(address)"))
    bytes4 internal constant ONLY_CALLABLE_BY_STAKING_CONTRACT_ERROR_SELECTOR =
        0xca1d07a2;

    // bytes4(keccak256("OnlyCallableIfInCatastrophicFailureError()"))
    bytes internal constant ONLY_CALLABLE_IF_IN_CATASTROPHIC_FAILURE_ERROR = "OnlyCallableIfInCatastrophicFailureError";

    // bytes4(keccak256("OnlyCallableIfNotInCatastrophicFailureError()"))
    bytes internal constant ONLY_CALLABLE_IF_NOT_IN_CATASTROPHIC_FAILURE_ERROR = "OnlyCallableIfNotInCatastrophicFailureError";

    // bytes4(keccak256("OperatorShareError(uint8,bytes32,uint32)"))
    bytes4 internal constant OPERATOR_SHARE_ERROR_SELECTOR =
        0x22df9597;

    // bytes4(keccak256("PoolExistenceError(bytes32,bool)"))
    bytes4 internal constant POOL_EXISTENCE_ERROR_SELECTOR =
        0x9ae94f01;

    // bytes4(keccak256("ProxyDestinationCannotBeNilError()"))
    bytes internal constant PROXY_DESTINATION_CANNOT_BE_NIL_ERROR = "PROXY_DESTINATION_CANNOT_BE_NIL_ERROR";

    // bytes4(keccak256("InitializationError(uint8)"))
    bytes4 internal constant INITIALIZATION_ERROR_SELECTOR =
        0x0b02d773;

    // bytes4(keccak256("InvalidParamValueError(uint8)"))
    bytes4 internal constant INVALID_PARAM_VALUE_ERROR_SELECTOR =
        0xfc45bd11;

    // bytes4(keccak256("InvalidProtocolFeePaymentError(uint256,uint256)"))
    bytes4 internal constant INVALID_PROTOCOL_FEE_PAYMENT_ERROR_SELECTOR =
        0x31d7a505;

    // bytes4(keccak256("PreviousEpochNotFinalizedError(uint256,uint256)"))
    bytes4 internal constant PREVIOUS_EPOCH_NOT_FINALIZED_ERROR_SELECTOR =
        0x614b800a;

    // bytes4(keccak256("PoolNotFinalizedError(bytes32,uint256)"))
    bytes4 internal constant POOL_NOT_FINALIZED_ERROR_SELECTOR =
        0x5caa0b05;

    // solhint-disable func-name-mixedcase
    function OnlyCallableByExchangeError(
        address senderAddress
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            ONLY_CALLABLE_BY_EXCHANGE_ERROR_SELECTOR,
            senderAddress
        );
    }

    function ExchangeManagerError(
        ExchangeManagerErrorCodes errorCodes,
        address exchangeAddress
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            EXCHANGE_MANAGER_ERROR_SELECTOR,
            errorCodes,
            exchangeAddress
        );
    }

    function InsufficientBalanceError(
        uint256 amount,
        uint256 balance
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            INSUFFICIENT_BALANCE_ERROR_SELECTOR,
            amount,
            balance
        );
    }

    function OnlyCallableByPoolOperatorError(
        address senderAddress,
        bytes32 poolId
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            ONLY_CALLABLE_BY_POOL_OPERATOR_ERROR_SELECTOR,
            senderAddress,
            poolId
        );
    }

    function BlockTimestampTooLowError(
        uint256 epochEndTime,
        uint256 currentBlockTimestamp
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            BLOCK_TIMESTAMP_TOO_LOW_ERROR_SELECTOR,
            epochEndTime,
            currentBlockTimestamp
        );
    }

    function OnlyCallableByStakingContractError(
        address senderAddress
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            ONLY_CALLABLE_BY_STAKING_CONTRACT_ERROR_SELECTOR,
            senderAddress
        );
    }

    function OnlyCallableIfInCatastrophicFailureError()
        internal
        pure
        returns (bytes memory)
    {
        return ONLY_CALLABLE_IF_IN_CATASTROPHIC_FAILURE_ERROR;
    }

    function OnlyCallableIfNotInCatastrophicFailureError()
        internal
        pure
        returns (bytes memory)
    {
        return ONLY_CALLABLE_IF_NOT_IN_CATASTROPHIC_FAILURE_ERROR;
    }

    function OperatorShareError(
        OperatorShareErrorCodes errorCodes,
        bytes32 poolId,
        uint32 operatorShare
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            OPERATOR_SHARE_ERROR_SELECTOR,
            errorCodes,
            poolId,
            operatorShare
        );
    }

    function PoolExistenceError(
        bytes32 poolId,
        bool alreadyExists
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            POOL_EXISTENCE_ERROR_SELECTOR,
            poolId,
            alreadyExists
        );
    }

    function InvalidProtocolFeePaymentError(
        uint256 expectedProtocolFeePaid,
        uint256 actualProtocolFeePaid
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            INVALID_PROTOCOL_FEE_PAYMENT_ERROR_SELECTOR,
            expectedProtocolFeePaid,
            actualProtocolFeePaid
        );
    }

    function InitializationError(InitializationErrorCodes code)
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            INITIALIZATION_ERROR_SELECTOR,
            uint8(code)
        );
    }

    function InvalidParamValueError(InvalidParamValueErrorCodes code)
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            INVALID_PARAM_VALUE_ERROR_SELECTOR,
            uint8(code)
        );
    }

    function ProxyDestinationCannotBeNilError()
        internal
        pure
        returns (bytes memory)
    {
        return PROXY_DESTINATION_CANNOT_BE_NIL_ERROR;
    }

    function PreviousEpochNotFinalizedError(
        uint256 unfinalizedEpoch,
        uint256 unfinalizedPoolsRemaining
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            PREVIOUS_EPOCH_NOT_FINALIZED_ERROR_SELECTOR,
            unfinalizedEpoch,
            unfinalizedPoolsRemaining
        );
    }

    function PoolNotFinalizedError(
        bytes32 poolId,
        uint256 epoch
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            POOL_NOT_FINALIZED_ERROR_SELECTOR,
            poolId,
            epoch
        );
    }
}
