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

import "../../../../utils/contracts/src/LibBytes.sol";

import "../../../../utils/contracts/src/LibSafeMath.sol";
import "../../../../erc20/contracts/src/LibERC20Token.sol";
import "../../../../erc721/contracts/src/interfaces/IERC721Token.sol";
import "../../../../erc1155/contracts/src/interfaces/IERC1155.sol";
import "../../../../asset-proxy/contracts/src/interfaces/IAssetData.sol";
import "./LibForwarderRichErrors.sol";


library LibAssetDataTransfer {

    using LibBytes for bytes;
    using LibSafeMath for uint256;
    using LibAssetDataTransfer for bytes;

    /// @dev Transfers given amount of asset to sender.
    /// @param assetData Byte array encoded for the respective asset proxy.
    /// @param from Address to transfer asset from.
    /// @param to Address to transfer asset to.
    /// @param amount Amount of asset to transfer to sender.
    function transferFrom(
        bytes memory assetData,
        address from,
        address to,
        uint256 amount
    )
        internal
    {
        if (amount == 0) {
            return;
        }

        bytes4 proxyId = assetData.readBytes4(0);

        if (
            proxyId == IAssetData(address(0)).ERC20Token.selector ||
            proxyId == IAssetData(address(0)).ERC20Bridge.selector
        ) {
            assetData.transferERC20Token(
                from,
                to,
                amount
            );
        } else if (proxyId == IAssetData(address(0)).ERC721Token.selector) {
            assetData.transferERC721Token(
                from,
                to,
                amount
            );
        } else if (proxyId == IAssetData(address(0)).ERC1155Assets.selector) {
            assetData.transferERC1155Assets(
                from,
                to,
                amount
            );
        } else if (proxyId == IAssetData(address(0)).MultiAsset.selector) {
            assetData.transferMultiAsset(
                from,
                to,
                amount
            );
        } else if (proxyId != IAssetData(address(0)).StaticCall.selector) {
            revert();
        }
    }

    ///@dev Transfer asset from sender to this contract.
    /// @param assetData Byte array encoded for the respective asset proxy.
    /// @param amount Amount of asset to transfer to sender.
    function transferIn(
        bytes memory assetData,
        uint256 amount
    )
        internal
    {
        assetData.transferFrom(
            msg.sender,
            address(this),
            amount
        );
    }

    ///@dev Transfer asset from this contract to sender.
    /// @param assetData Byte array encoded for the respective asset proxy.
    /// @param amount Amount of asset to transfer to sender.
    function transferOut(
        bytes memory assetData,
        uint256 amount
    )
        internal
    {
        assetData.transferFrom(
            address(this),
            msg.sender,
            amount
        );
    }

    /// @dev Decodes ERC20 or ERC20Bridge assetData and transfers given amount to sender.
    /// @param assetData Byte array encoded for the respective asset proxy.
    /// @param from Address to transfer asset from.
    /// @param to Address to transfer asset to.
    /// @param amount Amount of asset to transfer to sender.
    function transferERC20Token(
        bytes memory assetData,
        address from,
        address to,
        uint256 amount
    )
        internal
    {
        address token = assetData.readAddress(16);
        // Transfer tokens.
        if (from == address(this)) {
            LibERC20Token.transfer(
                token,
                to,
                amount
            );
        } else {
            LibERC20Token.transferFrom(
                token,
                from,
                to,
                amount
            );
        }
    }

    /// @dev Decodes ERC721 assetData and transfers given amount to sender.
    /// @param assetData Byte array encoded for the respective asset proxy.
    /// @param from Address to transfer asset from.
    /// @param to Address to transfer asset to.
    /// @param amount Amount of asset to transfer to sender.
    function transferERC721Token(
        bytes memory assetData,
        address from,
        address to,
        uint256 amount
    )
        internal
    {
        if (amount != 1) {
            revert();
        }
        // Decode asset data.
        address token = assetData.readAddress(16);
        uint256 tokenId = assetData.readUint256(36);

        // Perform transfer.
        IERC721Token(token).transferFrom(
            from,
            to,
            tokenId
        );
    }

    /// @dev Decodes ERC1155 assetData and transfers given amounts to sender.
    /// @param assetData Byte array encoded for the respective asset proxy.
    /// @param from Address to transfer asset from.
    /// @param to Address to transfer asset to.
    /// @param amount Amount of asset to transfer to sender.
    function transferERC1155Assets(
        bytes memory assetData,
        address from,
        address to,
        uint256 amount
    )
        internal
    {
        // Decode assetData
        // solhint-disable
        (
            address token,
            uint256[] memory ids,
            uint256[] memory values,
            bytes memory data
        ) = abi.decode(
            assetData.slice(4, assetData.length),
            (address, uint256[], uint256[], bytes)
        );
        // solhint-enable

        // Scale up values by `amount`
        uint256 length = values.length;
        uint256[] memory scaledValues = new uint256[](length);
        for (uint256 i = 0; i != length; i++) {
            scaledValues[i] = values[i].safeMul(amount);
        }

        // Execute `safeBatchTransferFrom` call
        // Either succeeds or throws
        IERC1155(token).safeBatchTransferFrom(
            from,
            to,
            ids,
            scaledValues,
            data
        );
    }

    /// @dev Decodes MultiAsset assetData and recursively transfers assets to sender.
    /// @param assetData Byte array encoded for the respective asset proxy.
    /// @param from Address to transfer asset from.
    /// @param to Address to transfer asset to.
    /// @param amount Amount of asset to transfer to sender.
    function transferMultiAsset(
        bytes memory assetData,
        address from,
        address to,
        uint256 amount
    )
        internal
    {
        // solhint-disable indent
        (uint256[] memory nestedAmounts, bytes[] memory nestedAssetData) = abi.decode(
            assetData.slice(4, assetData.length),
            (uint256[], bytes[])
        );
        // solhint-enable indent

        uint256 numNestedAssets = nestedAssetData.length;
        for (uint256 i = 0; i != numNestedAssets; i++) {
            transferFrom(
                nestedAssetData[i],
                from,
                to,
                amount.safeMul(nestedAmounts[i])
            );
        }
    }
}
