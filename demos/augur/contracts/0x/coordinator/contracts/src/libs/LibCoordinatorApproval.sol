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

import "./LibEIP712CoordinatorDomain.sol";


contract LibCoordinatorApproval is
    LibEIP712CoordinatorDomain
{
    // Hash for the EIP712 Coordinator approval message
    // keccak256(abi.encodePacked(
    //     "CoordinatorApproval(",
    //     "address txOrigin,",
    //     "bytes32 transactionHash,",
    //     "bytes transactionSignature",
    //     ")"
    // ));
    bytes32 constant public EIP712_COORDINATOR_APPROVAL_SCHEMA_HASH =
        0xa6511c04ca44625d50986f8c36bedc09366207a17b96e347094053a9f8507168;

    struct CoordinatorApproval {
        address txOrigin;                       // Required signer of Ethereum transaction that is submitting approval.
        bytes32 transactionHash;                // EIP712 hash of the transaction.
        bytes transactionSignature;             // Signature of the 0x transaction.
    }

    /// @dev Calculates the EIP712 hash of the Coordinator approval mesasage using the domain
    ///      separator of this contract.
    /// @param approval Coordinator approval message containing the transaction hash, and transaction
    ///        signature.
    /// @return approvalHash EIP712 hash of the Coordinator approval message with the domain
    ///         separator of this contract.
    function getCoordinatorApprovalHash(CoordinatorApproval memory approval)
        public
        view
        returns (bytes32 approvalHash)
    {
        approvalHash = _hashEIP712CoordinatorMessage(_hashCoordinatorApproval(approval));
        return approvalHash;
    }

    /// @dev Calculates the EIP712 hash of the Coordinator approval mesasage with no domain separator.
    /// @param approval Coordinator approval message containing the transaction hash, and transaction
    //         signature.
    /// @return result EIP712 hash of the Coordinator approval message with no domain separator.
    function _hashCoordinatorApproval(CoordinatorApproval memory approval)
        internal
        pure
        returns (bytes32 result)
    {
        bytes32 schemaHash = EIP712_COORDINATOR_APPROVAL_SCHEMA_HASH;
        bytes memory transactionSignature = approval.transactionSignature;
        address txOrigin = approval.txOrigin;
        bytes32 transactionHash = approval.transactionHash;

        // Assembly for more efficiently computing:
        // keccak256(abi.encodePacked(
        //     EIP712_COORDINATOR_APPROVAL_SCHEMA_HASH,
        //     approval.txOrigin,
        //     approval.transactionHash,
        //     keccak256(approval.transactionSignature)
        // ));

        assembly {
            // Compute hash of transaction signature
            let transactionSignatureHash := keccak256(add(transactionSignature, 32), mload(transactionSignature))

            // Load free memory pointer
            let memPtr := mload(64)

            mstore(memPtr, schemaHash)                               // hash of schema
            mstore(add(memPtr, 32), txOrigin)                        // txOrigin
            mstore(add(memPtr, 64), transactionHash)                 // transactionHash
            mstore(add(memPtr, 96), transactionSignatureHash)        // transactionSignatureHash
            // Compute hash
            result := keccak256(memPtr, 128)
        }
        return result;
    }
}
