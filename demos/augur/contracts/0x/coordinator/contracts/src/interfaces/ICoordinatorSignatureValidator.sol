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


contract ICoordinatorSignatureValidator {

   // Allowed signature types.
    enum SignatureType {
        Illegal,                // 0x00, default value
        Invalid,                // 0x01
        EIP712,                 // 0x02
        EthSign,                // 0x03
        Wallet,                 // 0x04
        Validator,              // 0x05
        PreSigned,              // 0x06
        EIP1271Wallet,          // 0x07
        NSignatureTypes         // 0x08, number of signature types. Always leave at end.
    }

    /// @dev Recovers the address of a signer given a hash and signature.
    /// @param hash Any 32 byte hash.
    /// @param signature Proof that the hash has been signed by signer.
    /// @return signerAddress Address of the signer. 
    function getSignerAddress(bytes32 hash, bytes memory signature)
        public
        pure
        returns (address signerAddress);
}
