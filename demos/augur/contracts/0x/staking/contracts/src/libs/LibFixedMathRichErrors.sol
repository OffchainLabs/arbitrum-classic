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




library LibFixedMathRichErrors {

    enum ValueErrorCodes {
        TOO_SMALL,
        TOO_LARGE
    }

    enum BinOpErrorCodes {
        ADDITION_OVERFLOW,
        MULTIPLICATION_OVERFLOW,
        DIVISION_BY_ZERO,
        DIVISION_OVERFLOW
    }

    // bytes4(keccak256("SignedValueError(uint8,int256)"))
    bytes4 internal constant SIGNED_VALUE_ERROR_SELECTOR =
        0xed2f26a1;

    // bytes4(keccak256("UnsignedValueError(uint8,uint256)"))
    bytes4 internal constant UNSIGNED_VALUE_ERROR_SELECTOR =
        0xbd79545f;

    // bytes4(keccak256("BinOpError(uint8,int256,int256)"))
    bytes4 internal constant BIN_OP_ERROR_SELECTOR =
        0x8c12dfe7;

    // solhint-disable func-name-mixedcase
    function SignedValueError(
        ValueErrorCodes error,
        int256 n
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            SIGNED_VALUE_ERROR_SELECTOR,
            uint8(error),
            n
        );
    }

    function UnsignedValueError(
        ValueErrorCodes error,
        uint256 n
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            UNSIGNED_VALUE_ERROR_SELECTOR,
            uint8(error),
            n
        );
    }

    function BinOpError(
        BinOpErrorCodes error,
        int256 a,
        int256 b
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            BIN_OP_ERROR_SELECTOR,
            uint8(error),
            a,
            b
        );
    }
}
