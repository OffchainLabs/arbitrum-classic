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

import "../../../utils/contracts/src/LibAddressArrayRichErrors.sol";
import "../../../utils/contracts/src/LibBytes.sol";



library LibAddressArray {

    /// @dev Append a new address to an array of addresses.
    ///      The `addressArray` may need to be reallocated to make space
    ///      for the new address. Because of this we return the resulting
    ///      memory location of `addressArray`.
    /// @param addressArray Array of addresses.
    /// @param addressToAppend  Address to append.
    /// @return Array of addresses: [... addressArray, addressToAppend]
    function append(address[] memory addressArray, address addressToAppend)
        internal
        pure
        returns (address[] memory)
    {
        // Get stats on address array and free memory
        uint256 freeMemPtr = 0;
        uint256 addressArrayBeginPtr = 0;
        uint256 addressArrayEndPtr = 0;
        uint256 addressArrayLength = addressArray.length;
        uint256 addressArrayMemSizeInBytes = 32 + (32 * addressArrayLength);
        assembly {
            freeMemPtr := mload(0x40)
            addressArrayBeginPtr := addressArray
            addressArrayEndPtr := add(addressArray, addressArrayMemSizeInBytes)
        }

        // Cases for `freeMemPtr`:
        //  `freeMemPtr` == `addressArrayEndPtr`: Nothing occupies memory after `addressArray`
        //  `freeMemPtr` > `addressArrayEndPtr`: Some value occupies memory after `addressArray`
        //  `freeMemPtr` < `addressArrayEndPtr`: Memory has not been managed properly.
        if (freeMemPtr < addressArrayEndPtr) {
            revert();
        }

        // If free memory begins at the end of `addressArray`
        // then we can append `addressToAppend` directly.
        // Otherwise, we must copy the array to free memory
        // before appending new values to it.
        if (freeMemPtr > addressArrayEndPtr) {
            LibBytes.memCopy(freeMemPtr, addressArrayBeginPtr, addressArrayMemSizeInBytes);
            assembly {
                addressArray := freeMemPtr
                addressArrayBeginPtr := addressArray
            }
        }

        // Append `addressToAppend`
        addressArrayLength += 1;
        addressArrayMemSizeInBytes += 32;
        addressArrayEndPtr = addressArrayBeginPtr + addressArrayMemSizeInBytes;
        freeMemPtr = addressArrayEndPtr;
        assembly {
            // Store new array length
            mstore(addressArray, addressArrayLength)

            // Update `freeMemPtr`
            mstore(0x40, freeMemPtr)
        }
        addressArray[addressArrayLength - 1] = addressToAppend;
        return addressArray;
    }

    /// @dev Checks if an address array contains the target address.
    /// @param addressArray Array of addresses.
    /// @param target Address to search for in array.
    /// @return True if the addressArray contains the target.
    function contains(address[] memory addressArray, address target)
        internal
        pure
        returns (bool success)
    {
        assembly {

            // Calculate byte length of array
            let arrayByteLen := mul(mload(addressArray), 32)
            // Calculate beginning of array contents
            let arrayContentsStart := add(addressArray, 32)
            // Calclulate end of array contents
            let arrayContentsEnd := add(arrayContentsStart, arrayByteLen)

            // Loop through array
            for {let i:= arrayContentsStart} lt(i, arrayContentsEnd) {i := add(i, 32)} {

                // Load array element
                let arrayElement := mload(i)

                // Return true if array element equals target
                if eq(target, arrayElement) {
                    // Set success to true
                    success := 1
                    // Break loop
                    i := arrayContentsEnd
                }
            }
        }
        return success;
    }

    /// @dev Finds the index of an address within an array.
    /// @param addressArray Array of addresses.
    /// @param target Address to search for in array.
    /// @return Existence and index of the target in the array.
    function indexOf(address[] memory addressArray, address target)
        internal
        pure
        returns (bool success, uint256 index)
    {
        assembly {

            // Calculate byte length of array
            let arrayByteLen := mul(mload(addressArray), 32)
            // Calculate beginning of array contents
            let arrayContentsStart := add(addressArray, 32)
            // Calclulate end of array contents
            let arrayContentsEnd := add(arrayContentsStart, arrayByteLen)

            // Loop through array
            for {let i:= arrayContentsStart} lt(i, arrayContentsEnd) {i := add(i, 32)} {

                // Load array element
                let arrayElement := mload(i)

                // Return true if array element equals target
                if eq(target, arrayElement) {
                    // Set success and index
                    success := 1
                    index := div(sub(i, arrayContentsStart), 32)
                    // Break loop
                    i := arrayContentsEnd
                }
            }
        }
        return (success, index);
    }
}
