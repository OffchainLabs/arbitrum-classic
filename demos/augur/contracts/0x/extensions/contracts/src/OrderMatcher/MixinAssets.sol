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

import "../../../../utils/contracts/src/LibBytes.sol";
import "../../../../utils/contracts/src/Ownable.sol";
import "../../../../erc20/contracts/src/interfaces/IERC20Token.sol";
import "../../../../erc721/contracts/src/interfaces/IERC721Token.sol";
import "./libs/LibConstants.sol";
import "./interfaces/IAssets.sol";


contract MixinAssets is
    IAssets,
    Ownable,
    LibConstants
{
    using LibBytes for bytes;

    /// @dev Withdraws assets from this contract. The contract requires a ZRX balance in order to 
    ///      function optimally, and this function allows the ZRX to be withdrawn by owner. It may also be
    ///      used to withdraw assets that were accidentally sent to this contract.
    /// @param assetData Byte array encoded for the respective asset proxy.
    /// @param amount Amount of asset to withdraw.
    function withdrawAsset(
        bytes calldata assetData,
        uint256 amount
    )
        external
        onlyOwner
    {
        _transferAssetToSender(assetData, amount);
    }

    /// @dev Approves or disapproves an AssetProxy to spend asset.
    /// @param assetData Byte array encoded for the respective asset proxy.
    /// @param amount Amount of asset to approve for respective proxy.
    function approveAssetProxy(
        bytes calldata assetData,
        uint256 amount
    )
        external
        onlyOwner
    {
        bytes4 proxyId = assetData.readBytes4(0);

        if (proxyId == ERC20_DATA_ID) {
            _approveERC20Token(assetData, amount);
        } else if (proxyId == ERC721_DATA_ID) {
            _approveERC721Token(assetData, amount);
        } else {
            revert("UNSUPPORTED_ASSET_PROXY");
        }
    }

    /// @dev Transfers given amount of asset to sender.
    /// @param assetData Byte array encoded for the respective asset proxy.
    /// @param amount Amount of asset to transfer to sender.
    function _transferAssetToSender(
        bytes memory assetData,
        uint256 amount
    )
        internal
    {
        bytes4 proxyId = assetData.readBytes4(0);

        if (proxyId == ERC20_DATA_ID) {
            _transferERC20Token(assetData, amount);
        } else if (proxyId == ERC721_DATA_ID) {
            _transferERC721Token(assetData, amount);
        } else {
            revert("UNSUPPORTED_ASSET_PROXY");
        }
    }

    /// @dev Decodes ERC20 assetData and transfers given amount to sender.
    /// @param assetData Byte array encoded for the respective asset proxy.
    /// @param amount Amount of asset to transfer to sender.
    function _transferERC20Token(
        bytes memory assetData,
        uint256 amount
    )
        internal
    {
        // 4 byte id + 12 0 bytes before ABI encoded token address.
        address token = assetData.readAddress(16);

        // Transfer tokens.
        // We do a raw call so we can check the success separate
        // from the return data.
        (bool success,) = token.call(abi.encodeWithSelector(
            ERC20_TRANSFER_SELECTOR,
            msg.sender,
            amount
        ));
        require(
            success,
            "TRANSFER_FAILED"
        );
        
        // Check return data.
        // If there is no return data, we assume the token incorrectly
        // does not return a bool. In this case we expect it to revert
        // on failure, which was handled above.
        // If the token does return data, we require that it is a single
        // value that evaluates to true.
        assembly {
            if returndatasize {
                success := 0
                if eq(returndatasize, 32) {
                    // First 64 bytes of memory are reserved scratch space
                    returndatacopy(0, 0, 32)
                    success := mload(0)
                }
            }
        }
        require(
            success,
            "TRANSFER_FAILED"
        );
    }

    /// @dev Decodes ERC721 assetData and transfers given amount to sender.
    /// @param assetData Byte array encoded for the respective asset proxy.
    /// @param amount Amount of asset to transfer to sender.
    function _transferERC721Token(
        bytes memory assetData,
        uint256 amount
    )
        internal
    {
        require(
            amount == 1,
            "INVALID_AMOUNT"
        );
        // Decode asset data.
        // 4 byte id + 12 0 bytes before ABI encoded token address.
        address token = assetData.readAddress(16);
        // 4 byte id + 32 byte ABI encoded token address before token id.
        uint256 tokenId = assetData.readUint256(36);

        // Perform transfer.
        IERC721Token(token).transferFrom(
            address(this),
            msg.sender,
            tokenId
        );
    }

    /// @dev Sets approval for ERC20 AssetProxy.
    /// @param assetData Byte array encoded for the respective asset proxy.
    /// @param amount Amount of asset to approve for respective proxy.
    function _approveERC20Token(
        bytes memory assetData,
        uint256 amount
    )
        internal
    {
        address token = assetData.readAddress(16);
        require(
            IERC20Token(token).approve(ERC20_PROXY_ADDRESS, amount),
            "APPROVAL_FAILED"
        );
    }

    /// @dev Sets approval for ERC721 AssetProxy.
    /// @param assetData Byte array encoded for the respective asset proxy.
    /// @param amount Amount of asset to approve for respective proxy.
    function _approveERC721Token(
        bytes memory assetData,
        uint256 amount
    )
        internal
    {
        address token = assetData.readAddress(16);
        bool approval = amount >= 1;
        IERC721Token(token).setApprovalForAll(ERC721_PROXY_ADDRESS, approval);
    }
}
