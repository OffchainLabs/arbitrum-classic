/*
 * Copyright 2019, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

pragma solidity ^0.5.3;


library SigUtils {

    function parseSignature(
        bytes memory _data,
        uint256 _start
    )
        internal
        pure
        returns (uint8 v, bytes32 r, bytes32 s)
    {
        uint256 offset = _start + 0x20;
        // The signature format is a compact form of:
        //   {bytes32 r}{bytes32 s}{uint8 v}
        // Compact means, uint8 is not padded to 32 bytes.
        assembly { // solium-disable-line security/no-inline-assembly
            r := mload(add(_data, offset))
            s := mload(add(_data, add(0x20, offset)))
            // Here we are loading the last 32 bytes, including 31 bytes
            // of 's'. There is no 'mload8' to do this.
            //
            // 'byte' is not working due to the Solidity parser, so lets
            // use the second best option, 'and'
            v := byte(0, mload(add(_data, add(offset, 0x40))))
        }

        if (v < 27) {
            v += 27;
        }

        require(v == 27 || v == 28, "Incorrect v value");
    }

    /// @notice Counts the number of signatures in a signatures bytes array. Returns 0 if the length is invalid.
    /// @param _signatures The signatures bytes array
    /// @dev Signatures are 65 bytes long and are densely packed.
    function countSignatures(bytes memory _signatures) internal pure returns (uint) {
        return _signatures.length % 65 == 0 ? _signatures.length / 65 : 0;
    }

    /// @notice Recovers an array of addresses using a message hash and a signatures bytes array.
    /// @param _messageHash The signed message hash
    /// @param _signatures The signatures bytes array
    function recoverAddresses(
        bytes32 _messageHash,
        bytes memory _signatures
    )
        internal
        pure
        returns (address[] memory)
    {
        uint8 v;
        bytes32 r;
        bytes32 s;
        uint256 count = countSignatures(_signatures);
        address[] memory addresses = new address[](count);
        bytes memory prefix = "\x19Ethereum Signed Message:\n32";
        bytes32 prefixedHash = keccak256(abi.encodePacked(prefix, _messageHash));
        for (uint256 i = 0; i < count; i++) {
            (v, r, s) = parseSignature(_signatures, i*65);
            addresses[i] = ecrecover(
                prefixedHash,
                v,
                r,
                s
            );
        }
        return addresses;
    }

    /// @notice Recovers an array of addresses using a message hash and a signatures bytes array.
    /// @param _messageHash The signed message hash
    /// @param _signature The signature bytes array
    function recoverAddress(
        bytes32 _messageHash,
        bytes memory _signature
    )
        internal
        pure
        returns (address)
    {
        return recoverAddressFromData(
            _messageHash,
            _signature,
            0x00
        );
    }

    function recoverAddressFromData(
        bytes32 _messageHash,
        bytes memory _data,
        uint256 _offset
    )
        internal
        pure
        returns (address)
    {
        uint8 v;
        bytes32 r;
        bytes32 s;
        bytes memory prefix = "\x19Ethereum Signed Message:\n32";
        bytes32 prefixedHash = keccak256(abi.encodePacked(prefix, _messageHash));
        (v, r, s) = parseSignature(_data, _offset);
        return ecrecover(
            prefixedHash,
            v,
            r,
            s
        );
    }
}
