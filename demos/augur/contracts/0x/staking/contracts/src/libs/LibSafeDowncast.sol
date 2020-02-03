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


import "../../../../utils/contracts/src/LibSafeMathRichErrors.sol";


library LibSafeDowncast {

    /// @dev Safely downcasts to a uint96
    /// Note that this reverts if the input value is too large.
    function downcastToUint96(uint256 a)
        internal
        pure
        returns (uint96 b)
    {
        b = uint96(a);
        if (uint256(b) != a) {
            revert();
        }
        return b;
    }

    /// @dev Safely downcasts to a uint64
    /// Note that this reverts if the input value is too large.
    function downcastToUint64(uint256 a)
        internal
        pure
        returns (uint64 b)
    {
        b = uint64(a);
        if (uint256(b) != a) {
            revert();
        }
        return b;
    }

    /// @dev Safely downcasts to a uint32
    /// Note that this reverts if the input value is too large.
    function downcastToUint32(uint256 a)
        internal
        pure
        returns (uint32 b)
    {
        b = uint32(a);
        if (uint256(b) != a) {
            revert();
        }
        return b;
    }
}
