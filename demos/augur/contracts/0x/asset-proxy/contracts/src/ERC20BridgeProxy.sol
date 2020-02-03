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

import "../../../utils/contracts/src/LibBytes.sol";
import "../../../utils/contracts/src/LibSafeMath.sol";
import "../../../utils/contracts/src/Authorizable.sol";
import "../../../erc20/contracts/src/interfaces/IERC20Token.sol";
import "./interfaces/IAssetProxy.sol";
import "./interfaces/IERC20Bridge.sol";


contract ERC20BridgeProxy is
    IAssetProxy,
    Authorizable
{
    using LibBytes for bytes;
    using LibSafeMath for uint256;

    // @dev Id of this proxy. Also the result of a successful bridge call.
    //      bytes4(keccak256("ERC20Bridge(address,address,bytes)"))
    bytes4 constant private PROXY_ID = 0xdc1600f3;

    /// @dev Calls a bridge contract to transfer `amount` of ERC20 from `from`
    ///      to `to`. Asserts that the balance of `to` has increased by `amount`.
    /// @param assetData Abi-encoded data for this asset proxy encoded as:
    ///          abi.encodeWithSelector(
    ///             bytes4 PROXY_ID,
    ///             address tokenAddress,
    ///             address bridgeAddress,
    ///             bytes bridgeData
    ///          )
    /// @param from Address to transfer asset from.
    /// @param to Address to transfer asset to.
    /// @param amount Amount of asset to transfer.
    function transferFrom(
        bytes calldata assetData,
        address from,
        address to,
        uint256 amount
    )
        external
        onlyAuthorized
    {
        // Extract asset data fields.
        (
            address tokenAddress,
            address bridgeAddress,
            bytes memory bridgeData
        ) = abi.decode(
            assetData.sliceDestructive(4, assetData.length),
            (address, address, bytes)
        );

        // Remember the balance of `to` before calling the bridge.
        uint256 balanceBefore = balanceOf(tokenAddress, to);
        // Call the bridge, who should transfer `amount` of `tokenAddress` to
        // `to`.
        bytes4 success = IERC20Bridge(bridgeAddress).bridgeTransferFrom(
            tokenAddress,
            from,
            to,
            amount,
            bridgeData
        );
        // Bridge must return the proxy ID to indicate success.
        require(success == PROXY_ID, "BRIDGE_FAILED");
        // Ensure that the balance of `to` has increased by at least `amount`.
        require(
            balanceBefore.safeAdd(amount) <= balanceOf(tokenAddress, to),
            "BRIDGE_UNDERPAY"
        );
    }

    /// @dev Gets the proxy id associated with this asset proxy.
    /// @return proxyId The proxy id.
    function getProxyId()
        external
        pure
        returns (bytes4 proxyId)
    {
        return PROXY_ID;
    }

    /// @dev Retrieves the balance of `owner` for this asset.
    /// @return balance The balance of the ERC20 token being transferred by this
    ///         asset proxy.
    function balanceOf(bytes calldata assetData, address owner)
        external
        view
        returns (uint256 balance)
    {
        (address tokenAddress) = abi.decode(
            assetData.sliceDestructive(4, assetData.length),
            (address)
        );
        return balanceOf(tokenAddress, owner);
    }

    /// @dev Retrieves the balance of `owner` given an ERC20 address.
    /// @return balance The balance of the ERC20 token for `owner`.
    function balanceOf(address tokenAddress, address owner)
        private
        view
        returns (uint256 balance)
    {
        return IERC20Token(tokenAddress).balanceOf(owner);
    }
}
