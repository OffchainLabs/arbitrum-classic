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

import "../../../utils/contracts/src/LibBytes.sol";


// solhint-disable no-unused-vars
contract StaticCallProxy {

    using LibBytes for bytes;

    // Id of this proxy.
    bytes4 constant internal PROXY_ID = bytes4(keccak256("StaticCall(address,bytes,bytes32)"));

    /// @dev Makes a staticcall to a target address and verifies that the data returned matches the expected return data.
    /// @param assetData Byte array encoded with staticCallTarget, staticCallData, and expectedCallResultHash
    /// @param from This value is ignored.
    /// @param to This value is ignored.
    /// @param amount This value is ignored.
    function transferFrom(
        bytes calldata assetData,
        address from,
        address to,
        uint256 amount
    )
        external
        view
    {
        // Decode params from `assetData`
        (
            address staticCallTarget,
            bytes memory staticCallData,
            bytes32 expectedReturnDataHash
        ) = abi.decode(
            assetData.sliceDestructive(4, assetData.length),
            (address, bytes, bytes32)
        );

        // Execute staticcall
        (bool success, bytes memory returnData) = staticCallTarget.staticcall(staticCallData);

        // Revert with returned data if staticcall is unsuccessful
        if (!success) {
            assembly {
                revert(add(returnData, 32), mload(returnData))
            }
        }

        // Revert if hash of return data is not as expected
        bytes32 returnDataHash = keccak256(returnData);
        require(
            expectedReturnDataHash == returnDataHash,
            "UNEXPECTED_STATIC_CALL_RESULT"
        );
    }

    /// @dev Gets the proxy id associated with the proxy address.
    /// @return Proxy id.
    function getProxyId()
        external
        pure
        returns (bytes4)
    {
        return PROXY_ID;
    }
}
