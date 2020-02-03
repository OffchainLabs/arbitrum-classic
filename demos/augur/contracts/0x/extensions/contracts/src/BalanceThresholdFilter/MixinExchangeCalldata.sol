    
    /*

  Copyright 2019 ZeroEx Intl.

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

*/

pragma solidity 0.5.15;

import "../../../../utils/contracts/src/LibAddressArray.sol";


contract MixinExchangeCalldata {

    using LibAddressArray for address[];

    /// @dev Emulates the `calldataload` opcode on the embedded Exchange calldata,
    ///      which is accessed through `signedExchangeTransaction`.
    /// @param offset  Offset into the Exchange calldata.
    /// @return value  Corresponding 32 byte value stored at `offset`.
    function _exchangeCalldataload(uint256 offset)
        internal
        pure
        returns (bytes32 value)
    {
        assembly {
            // Pointer to exchange transaction
            // 0x04 for calldata selector
            // 0x40 to access `signedExchangeTransaction`, which is the third parameter
            let exchangeTxPtr := calldataload(0x44)

            // Offset into Exchange calldata
            // We compute this by adding 0x24 to the `exchangeTxPtr` computed above.
            // 0x04 for calldata selector
            // 0x20 for length field of `signedExchangeTransaction`
            let exchangeCalldataOffset := add(exchangeTxPtr, add(0x24, offset))
            value := calldataload(exchangeCalldataOffset)
        }
        return value;
    }

    /// @dev Convenience function that skips the 4 byte selector when loading
    ///      from the embedded Exchange calldata.
    /// @param offset  Offset into the Exchange calldata (minus the 4 byte selector)
    /// @return value  Corresponding 32 byte value stored at `offset` + 4.
    function _loadExchangeData(uint256 offset)
        internal
        pure
        returns (bytes32 value)
    {
        value = _exchangeCalldataload(offset + 4);
        return value;
    }

    /// @dev Extracts the maker address from an order stored in the Exchange calldata
    ///      (which is embedded in `signedExchangeTransaction`).
    /// @param orderParamIndex  Index of the order in the Exchange function's signature.
    /// @return makerAddress The extracted maker address.
    function _loadMakerAddressFromOrder(uint256 orderParamIndex)
        internal
        pure
        returns (address makerAddress)
    {
        uint256 orderOffsetInBytes = orderParamIndex * 32;
        uint256 orderPtr = uint256(_loadExchangeData(orderOffsetInBytes));
        makerAddress = address(uint256(_loadExchangeData(orderPtr)));
        return makerAddress;
    }

    /// @dev Extracts the maker addresses from an array of orders stored in the Exchange calldata
    ///      (which is embedded in `signedExchangeTransaction`).
    /// @param orderArrayParamIndex  Index of the order array in the Exchange function's signature
    /// @return makerAddresses The extracted maker addresses.
    function _loadMakerAddressesFromOrderArray(uint256 orderArrayParamIndex)
        internal
        pure
        returns (address[] memory makerAddresses)
    {
        uint256 orderArrayOffsetInBytes = orderArrayParamIndex * 32;
        uint256 orderArrayPtr = uint256(_loadExchangeData(orderArrayOffsetInBytes));
        uint256 orderArrayLength = uint256(_loadExchangeData(orderArrayPtr));
        uint256 orderArrayLengthInBytes = orderArrayLength * 32;
        uint256 orderArrayElementPtr = orderArrayPtr + 32;
        uint256 orderArrayElementEndPtr = orderArrayElementPtr + orderArrayLengthInBytes;
        for (uint orderPtrOffset = orderArrayElementPtr; orderPtrOffset < orderArrayElementEndPtr; orderPtrOffset += 32) {
            uint256 orderPtr = uint256(_loadExchangeData(orderPtrOffset));
            address makerAddress = address(uint256(_loadExchangeData(orderPtr + orderArrayElementPtr)));
            makerAddresses = makerAddresses.append(makerAddress);
        }
        return makerAddresses;
    }
}
