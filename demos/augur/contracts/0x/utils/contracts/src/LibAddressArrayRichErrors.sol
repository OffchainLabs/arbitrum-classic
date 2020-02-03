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


library LibAddressArrayRichErrors {

    // bytes4(keccak256("MismanagedMemoryError(uint256,uint256)"))
    bytes4 internal constant MISMANAGED_MEMORY_ERROR_SELECTOR =
        0x5fc83722;

    // solhint-disable func-name-mixedcase
    function MismanagedMemoryError(
        uint256 freeMemPtr,
        uint256 addressArrayEndPtr
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodeWithSelector(
            MISMANAGED_MEMORY_ERROR_SELECTOR,
            freeMemPtr,
            addressArrayEndPtr
        );
    }
}
