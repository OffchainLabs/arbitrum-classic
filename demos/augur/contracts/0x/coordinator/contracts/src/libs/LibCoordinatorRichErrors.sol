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


library LibCoordinatorRichErrors {
    enum SignatureErrorCodes {
        INVALID_LENGTH,
        UNSUPPORTED,
        ILLEGAL,
        INVALID
    }

    // bytes4(keccak256("SignatureError(uint8,bytes32,bytes)"))
    bytes4 internal constant SIGNATURE_ERROR_SELECTOR =
        0x779c5223;

    // bytes4(keccak256("InvalidOriginError(address)"))
    bytes4 internal constant INVALID_ORIGIN_ERROR_SELECTOR =
        0xa458d7ff;

    // bytes4(keccak256("InvalidApprovalSignatureError(bytes32,address)"))
    bytes4 internal constant INVALID_APPROVAL_SIGNATURE_ERROR_SELECTOR =
        0xd789b640;

    // solhint-disable func-name-mixedcase
    function SignatureError(
        SignatureErrorCodes errorCode,
        bytes32 hash,
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
            signature
        );
    }

    function InvalidOriginError(
        address expectedOrigin
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            INVALID_ORIGIN_ERROR_SELECTOR,
            expectedOrigin
        );
    }

    function InvalidApprovalSignatureError(
        bytes32 transactionHash,
        address approverAddress
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            INVALID_APPROVAL_SIGNATURE_ERROR_SELECTOR,
            transactionHash,
            approverAddress
        );
    }
}
