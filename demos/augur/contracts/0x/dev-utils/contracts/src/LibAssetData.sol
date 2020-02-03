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
import "../../../exchange/contracts/src/interfaces/IExchange.sol";
import "../../../asset-proxy/contracts/src/interfaces/IAssetData.sol";
import "../../../asset-proxy/contracts/src/interfaces/IAssetProxy.sol";
import "../../../erc20/contracts/src/interfaces/IERC20Token.sol";
import "../../../erc721/contracts/src/interfaces/IERC721Token.sol";
import "../../../erc1155/contracts/src/interfaces/IERC1155.sol";
import "../../../asset-proxy/contracts/src/interfaces/IChai.sol";
import "../../../utils/contracts/src/DeploymentConstants.sol";
import "../../../exchange-libs/contracts/src/LibMath.sol";


contract LibAssetData is
    DeploymentConstants
{
    // 2^256 - 1
    uint256 constant internal _MAX_UINT256 = uint256(-1);

    using LibBytes for bytes;

    // solhint-disable var-name-mixedcase
    IExchange internal _EXCHANGE;
    address internal _ERC20_PROXY_ADDRESS;
    address internal _ERC721_PROXY_ADDRESS;
    address internal _ERC1155_PROXY_ADDRESS;
    address internal _STATIC_CALL_PROXY_ADDRESS;
    address internal _CHAI_BRIDGE_ADDRESS;
    // solhint-enable var-name-mixedcase

    constructor (
        address _exchange,
        address _chaiBridge
    )
        public
    {
        _EXCHANGE = IExchange(_exchange);
        _ERC20_PROXY_ADDRESS = _EXCHANGE.getAssetProxy(IAssetData(address(0)).ERC20Token.selector);
        _ERC721_PROXY_ADDRESS = _EXCHANGE.getAssetProxy(IAssetData(address(0)).ERC721Token.selector);
        _ERC1155_PROXY_ADDRESS = _EXCHANGE.getAssetProxy(IAssetData(address(0)).ERC1155Assets.selector);
        _STATIC_CALL_PROXY_ADDRESS = _EXCHANGE.getAssetProxy(IAssetData(address(0)).StaticCall.selector);
        _CHAI_BRIDGE_ADDRESS = _chaiBridge;
    }

    /// @dev Returns the owner's balance of the assets(s) specified in
    /// assetData.  When the asset data contains multiple assets (eg in
    /// ERC1155 or Multi-Asset), the return value indicates how many
    /// complete "baskets" of those assets are owned by owner.
    /// @param ownerAddress Owner of the assets specified by assetData.
    /// @param assetData Details of asset, encoded per the AssetProxy contract specification.
    /// @return Number of assets (or asset baskets) held by owner.
    function getBalance(address ownerAddress, bytes memory assetData)
        public
        returns (uint256 balance)
    {
        // Get id of AssetProxy contract
        bytes4 assetProxyId = assetData.readBytes4(0);

        if (assetProxyId == IAssetData(address(0)).ERC20Token.selector) {
            // Get ERC20 token address
            address tokenAddress = assetData.readAddress(16);
            balance = _erc20BalanceOf(tokenAddress, ownerAddress);

        } else if (assetProxyId == IAssetData(address(0)).ERC1155Assets.selector) {
            // Get ERC1155 token address, array of ids, and array of values
            (, address tokenAddress, uint256[] memory tokenIds, uint256[] memory tokenValues,) = decodeERC1155AssetData(assetData);

            uint256 length = tokenIds.length;
            for (uint256 i = 0; i != length; i++) {
                // Skip over the token if the corresponding value is 0.
                if (tokenValues[i] == 0) {
                    continue;
                }

                // Encode data for `balanceOf(ownerAddress, tokenIds[i])
                bytes memory balanceOfData = abi.encodeWithSelector(
                    IERC1155(address(0)).balanceOf.selector,
                    ownerAddress,
                    tokenIds[i]
                );

                // Query balance
                (bool success, bytes memory returnData) = tokenAddress.staticcall(balanceOfData);
                uint256 totalBalance = success && returnData.length == 32 ? returnData.readUint256(0) : 0;

                // Scale total balance down by corresponding value in assetData
                uint256 scaledBalance = totalBalance / tokenValues[i];
                if (scaledBalance == 0) {
                    return 0;
                }
                if (scaledBalance < balance || balance == 0) {
                    balance = scaledBalance;
                }
            }

        } else if (assetProxyId == IAssetData(address(0)).StaticCall.selector) {
            // Encode data for `staticCallProxy.transferFrom(assetData,...)`
            bytes memory transferFromData = abi.encodeWithSelector(
                IAssetProxy(address(0)).transferFrom.selector,
                assetData,
                address(0),  // `from` address is not used
                address(0),  // `to` address is not used
                0            // `amount` is not used
            );

            // Check if staticcall would be successful
            (bool success,) = _STATIC_CALL_PROXY_ADDRESS.staticcall(transferFromData);

            // Success means that the staticcall can be made an unlimited amount of times
            balance = success ? _MAX_UINT256 : 0;

        } else if (assetProxyId == IAssetData(address(0)).ERC20Bridge.selector) {
            // Get address of ERC20 token and bridge contract
            (, address tokenAddress, address bridgeAddress,) = decodeERC20BridgeAssetData(assetData);
            if (tokenAddress == _getDaiAddress() && bridgeAddress == _CHAI_BRIDGE_ADDRESS) {
                uint256 chaiBalance = _erc20BalanceOf(_getChaiAddress(), ownerAddress);
                // Calculate Dai balance
                balance = _convertChaiToDaiAmount(chaiBalance);
            }
            // Balance will be 0 if bridge is not supported

        } else if (assetProxyId == IAssetData(address(0)).MultiAsset.selector) {
            // Get array of values and array of assetDatas
            (, uint256[] memory assetAmounts, bytes[] memory nestedAssetData) = decodeMultiAssetData(assetData);

            uint256 length = nestedAssetData.length;
            for (uint256 i = 0; i != length; i++) {
                // Skip over the asset if the corresponding amount is 0.
                if (assetAmounts[i] == 0) {
                    continue;
                }

                // Query balance of individual assetData
                uint256 totalBalance = getBalance(ownerAddress, nestedAssetData[i]);

                // Scale total balance down by corresponding value in assetData
                uint256 scaledBalance = totalBalance / assetAmounts[i];
                if (scaledBalance == 0) {
                    return 0;
                }
                if (scaledBalance < balance || balance == 0) {
                    balance = scaledBalance;
                }
            }
        }

        // Balance will be 0 if assetProxyId is unknown
        return balance;
    }

    /// @dev Calls getBalance() for each element of assetData.
    /// @param ownerAddress Owner of the assets specified by assetData.
    /// @param assetData Array of asset details, each encoded per the AssetProxy contract specification.
    /// @return Array of asset balances from getBalance(), with each element
    /// corresponding to the same-indexed element in the assetData input.
    function getBatchBalances(address ownerAddress, bytes[] memory assetData)
        public
        returns (uint256[] memory balances)
    {
        uint256 length = assetData.length;
        balances = new uint256[](length);
        for (uint256 i = 0; i != length; i++) {
            balances[i] = getBalance(ownerAddress, assetData[i]);
        }
        return balances;
    }

    /// @dev Returns the number of asset(s) (described by assetData) that
    /// the corresponding AssetProxy contract is authorized to spend.  When the asset data contains
    /// multiple assets (eg for Multi-Asset), the return value indicates
    /// how many complete "baskets" of those assets may be spent by all of the corresponding
    /// AssetProxy contracts.
    /// @param ownerAddress Owner of the assets specified by assetData.
    /// @param assetData Details of asset, encoded per the AssetProxy contract specification.
    /// @return Number of assets (or asset baskets) that the corresponding AssetProxy is authorized to spend.
    function getAssetProxyAllowance(address ownerAddress, bytes memory assetData)
        public
        returns (uint256 allowance)
    {
        // Get id of AssetProxy contract
        bytes4 assetProxyId = assetData.readBytes4(0);

        if (assetProxyId == IAssetData(address(0)).MultiAsset.selector) {
            // Get array of values and array of assetDatas
            (, uint256[] memory amounts, bytes[] memory nestedAssetData) = decodeMultiAssetData(assetData);

            uint256 length = nestedAssetData.length;
            for (uint256 i = 0; i != length; i++) {
                // Skip over the asset if the corresponding amount is 0.
                if (amounts[i] == 0) {
                    continue;
                }

                // Query allowance of individual assetData
                uint256 totalAllowance = getAssetProxyAllowance(ownerAddress, nestedAssetData[i]);

                // Scale total allowance down by corresponding value in assetData
                uint256 scaledAllowance = totalAllowance / amounts[i];
                if (scaledAllowance == 0) {
                    return 0;
                }
                if (scaledAllowance < allowance || allowance == 0) {
                    allowance = scaledAllowance;
                }
            }
            return allowance;
        }

        if (assetProxyId == IAssetData(address(0)).ERC20Token.selector) {
            // Get ERC20 token address
            address tokenAddress = assetData.readAddress(16);

            // Encode data for `allowance(ownerAddress, _ERC20_PROXY_ADDRESS)`
            bytes memory allowanceData = abi.encodeWithSelector(
                IERC20Token(address(0)).allowance.selector,
                ownerAddress,
                _ERC20_PROXY_ADDRESS
            );

            // Query allowance
            (bool success, bytes memory returnData) = tokenAddress.staticcall(allowanceData);
            allowance = success && returnData.length == 32 ? returnData.readUint256(0) : 0;

        } else if (assetProxyId == IAssetData(address(0)).ERC1155Assets.selector) {
            // Get ERC1155 token address
            (, address tokenAddress, , , ) = decodeERC1155AssetData(assetData);

            // Encode data for `isApprovedForAll(ownerAddress, _ERC1155_PROXY_ADDRESS)`
            bytes memory isApprovedForAllData = abi.encodeWithSelector(
                IERC1155(address(0)).isApprovedForAll.selector,
                ownerAddress,
                _ERC1155_PROXY_ADDRESS
            );

            // Query allowance
            (bool success, bytes memory returnData) = tokenAddress.staticcall(isApprovedForAllData);
            allowance = success && returnData.length == 32 && returnData.readUint256(0) == 1 ? _MAX_UINT256 : 0;

        } else if (assetProxyId == IAssetData(address(0)).StaticCall.selector) {
            // The StaticCallProxy does not require any approvals
            allowance = _MAX_UINT256;

        } else if (assetProxyId == IAssetData(address(0)).ERC20Bridge.selector) {
            // Get address of ERC20 token and bridge contract
            (, address tokenAddress, address bridgeAddress,) = decodeERC20BridgeAssetData(assetData);
            if (tokenAddress == _getDaiAddress() && bridgeAddress == _CHAI_BRIDGE_ADDRESS) {
                bytes memory allowanceData = abi.encodeWithSelector(
                    IERC20Token(address(0)).allowance.selector,
                    ownerAddress,
                    _CHAI_BRIDGE_ADDRESS
                );
                (bool success, bytes memory returnData) = _getChaiAddress().staticcall(allowanceData);
                uint256 chaiAllowance = success && returnData.length == 32 ? returnData.readUint256(0) : 0;
                // Dai allowance is unlimited if Chai allowance is unlimited
                allowance = chaiAllowance == _MAX_UINT256 ? _MAX_UINT256 : _convertChaiToDaiAmount(chaiAllowance);
            }
            // Allowance will be 0 if bridge is not supported
        }

        // Allowance will be 0 if the assetProxyId is unknown
        return allowance;
    }

    /// @dev Calls getAssetProxyAllowance() for each element of assetData.
    /// @param ownerAddress Owner of the assets specified by assetData.
    /// @param assetData Array of asset details, each encoded per the AssetProxy contract specification.
    /// @return An array of asset allowances from getAllowance(), with each
    /// element corresponding to the same-indexed element in the assetData input.
    function getBatchAssetProxyAllowances(address ownerAddress, bytes[] memory assetData)
        public
        returns (uint256[] memory allowances)
    {
        uint256 length = assetData.length;
        allowances = new uint256[](length);
        for (uint256 i = 0; i != length; i++) {
            allowances[i] = getAssetProxyAllowance(ownerAddress, assetData[i]);
        }
        return allowances;
    }

    /// @dev Calls getBalance() and getAllowance() for assetData.
    /// @param ownerAddress Owner of the assets specified by assetData.
    /// @param assetData Details of asset, encoded per the AssetProxy contract specification.
    /// @return Number of assets (or asset baskets) held by owner, and number
    /// of assets (or asset baskets) that the corresponding AssetProxy is authorized to spend.
    function getBalanceAndAssetProxyAllowance(address ownerAddress, bytes memory assetData)
        public
        returns (uint256 balance, uint256 allowance)
    {
        balance = getBalance(ownerAddress, assetData);
        allowance = getAssetProxyAllowance(ownerAddress, assetData);
        return (balance, allowance);
    }

    /// @dev Calls getBatchBalances() and getBatchAllowances() for each element of assetData.
    /// @param ownerAddress Owner of the assets specified by assetData.
    /// @param assetData Array of asset details, each encoded per the AssetProxy contract specification.
    /// @return An array of asset balances from getBalance(), and an array of
    /// asset allowances from getAllowance(), with each element
    /// corresponding to the same-indexed element in the assetData input.
    function getBatchBalancesAndAssetProxyAllowances(address ownerAddress, bytes[] memory assetData)
        public
        returns (uint256[] memory balances, uint256[] memory allowances)
    {
        balances = getBatchBalances(ownerAddress, assetData);
        allowances = getBatchAssetProxyAllowances(ownerAddress, assetData);
        return (balances, allowances);
    }

    /// @dev Decode AssetProxy identifier
    /// @param assetData AssetProxy-compliant asset data describing an ERC-20, ERC-721, ERC1155, or MultiAsset asset.
    /// @return The AssetProxy identifier
    function decodeAssetProxyId(bytes memory assetData)
        public
        pure
        returns (
            bytes4 assetProxyId
        )
    {
        assetProxyId = assetData.readBytes4(0);

        require(
            assetProxyId == IAssetData(address(0)).ERC20Token.selector ||
            assetProxyId == IAssetData(address(0)).ERC721Token.selector ||
            assetProxyId == IAssetData(address(0)).ERC1155Assets.selector ||
            assetProxyId == IAssetData(address(0)).MultiAsset.selector ||
            assetProxyId == IAssetData(address(0)).StaticCall.selector,
            "WRONG_PROXY_ID"
        );
        return assetProxyId;
    }

    /// @dev Encode ERC-20 asset data into the format described in the AssetProxy contract specification.
    /// @param tokenAddress The address of the ERC-20 contract hosting the asset to be traded.
    /// @return AssetProxy-compliant data describing the asset.
    function encodeERC20AssetData(address tokenAddress)
        public
        pure
        returns (bytes memory assetData)
    {
        assetData = abi.encodeWithSelector(IAssetData(address(0)).ERC20Token.selector, tokenAddress);
        return assetData;
    }

    /// @dev Decode ERC-20 asset data from the format described in the AssetProxy contract specification.
    /// @param assetData AssetProxy-compliant asset data describing an ERC-20 asset.
    /// @return The AssetProxy identifier, and the address of the ERC-20
    /// contract hosting this asset.
    function decodeERC20AssetData(bytes memory assetData)
        public
        pure
        returns (
            bytes4 assetProxyId,
            address tokenAddress
        )
    {
        assetProxyId = assetData.readBytes4(0);

        require(
            assetProxyId == IAssetData(address(0)).ERC20Token.selector,
            "WRONG_PROXY_ID"
        );

        tokenAddress = assetData.readAddress(16);
        return (assetProxyId, tokenAddress);
    }

    /// @dev Encode ERC-721 asset data into the format described in the AssetProxy specification.
    /// @param tokenAddress The address of the ERC-721 contract hosting the asset to be traded.
    /// @param tokenId The identifier of the specific asset to be traded.
    /// @return AssetProxy-compliant asset data describing the asset.
    function encodeERC721AssetData(address tokenAddress, uint256 tokenId)
        public
        pure
        returns (bytes memory assetData)
    {
        assetData = abi.encodeWithSelector(
            IAssetData(address(0)).ERC721Token.selector,
            tokenAddress,
            tokenId
        );
        return assetData;
    }

    /// @dev Decode ERC-721 asset data from the format described in the AssetProxy contract specification.
    /// @param assetData AssetProxy-compliant asset data describing an ERC-721 asset.
    /// @return The ERC-721 AssetProxy identifier, the address of the ERC-721
    /// contract hosting this asset, and the identifier of the specific
    /// asset to be traded.
    function decodeERC721AssetData(bytes memory assetData)
        public
        pure
        returns (
            bytes4 assetProxyId,
            address tokenAddress,
            uint256 tokenId
        )
    {
        assetProxyId = assetData.readBytes4(0);

        require(
            assetProxyId == IAssetData(address(0)).ERC721Token.selector,
            "WRONG_PROXY_ID"
        );

        tokenAddress = assetData.readAddress(16);
        tokenId = assetData.readUint256(36);
        return (assetProxyId, tokenAddress, tokenId);
    }

    /// @dev Encode ERC-1155 asset data into the format described in the AssetProxy contract specification.
    /// @param tokenAddress The address of the ERC-1155 contract hosting the asset(s) to be traded.
    /// @param tokenIds The identifiers of the specific assets to be traded.
    /// @param tokenValues The amounts of each asset to be traded.
    /// @param callbackData Data to be passed to receiving contracts when a transfer is performed.
    /// @return AssetProxy-compliant asset data describing the set of assets.
    function encodeERC1155AssetData(
        address tokenAddress,
        uint256[] memory tokenIds,
        uint256[] memory tokenValues,
        bytes memory callbackData
    )
        public
        pure
        returns (bytes memory assetData)
    {
        assetData = abi.encodeWithSelector(
            IAssetData(address(0)).ERC1155Assets.selector,
            tokenAddress,
            tokenIds,
            tokenValues,
            callbackData
        );
        return assetData;
    }

    /// @dev Decode ERC-1155 asset data from the format described in the AssetProxy contract specification.
    /// @param assetData AssetProxy-compliant asset data describing an ERC-1155 set of assets.
    /// @return The ERC-1155 AssetProxy identifier, the address of the ERC-1155
    /// contract hosting the assets, an array of the identifiers of the
    /// assets to be traded, an array of asset amounts to be traded, and
    /// callback data.  Each element of the arrays corresponds to the
    /// same-indexed element of the other array.  Return values specified as
    /// `memory` are returned as pointers to locations within the memory of
    /// the input parameter `assetData`.
    function decodeERC1155AssetData(bytes memory assetData)
        public
        pure
        returns (
            bytes4 assetProxyId,
            address tokenAddress,
            uint256[] memory tokenIds,
            uint256[] memory tokenValues,
            bytes memory callbackData
        )
    {
        assetProxyId = assetData.readBytes4(0);

        require(
            assetProxyId == IAssetData(address(0)).ERC1155Assets.selector,
            "WRONG_PROXY_ID"
        );

        assembly {
            // Skip selector and length to get to the first parameter:
            assetData := add(assetData, 36)
            // Read the value of the first parameter:
            tokenAddress := mload(assetData)
            // Point to the next parameter's data:
            tokenIds := add(assetData, mload(add(assetData, 32)))
            // Point to the next parameter's data:
            tokenValues := add(assetData, mload(add(assetData, 64)))
            // Point to the next parameter's data:
            callbackData := add(assetData, mload(add(assetData, 96)))
        }

        return (
            assetProxyId,
            tokenAddress,
            tokenIds,
            tokenValues,
            callbackData
        );
    }

    /// @dev Encode data for multiple assets, per the AssetProxy contract specification.
    /// @param amounts The amounts of each asset to be traded.
    /// @param nestedAssetData AssetProxy-compliant data describing each asset to be traded.
    /// @return AssetProxy-compliant data describing the set of assets.
    function encodeMultiAssetData(uint256[] memory amounts, bytes[] memory nestedAssetData)
        public
        pure
        returns (bytes memory assetData)
    {
        assetData = abi.encodeWithSelector(
            IAssetData(address(0)).MultiAsset.selector,
            amounts,
            nestedAssetData
        );
        return assetData;
    }

    /// @dev Decode multi-asset data from the format described in the AssetProxy contract specification.
    /// @param assetData AssetProxy-compliant data describing a multi-asset basket.
    /// @return The Multi-Asset AssetProxy identifier, an array of the amounts
    /// of the assets to be traded, and an array of the
    /// AssetProxy-compliant data describing each asset to be traded.  Each
    /// element of the arrays corresponds to the same-indexed element of the other array.
    function decodeMultiAssetData(bytes memory assetData)
        public
        pure
        returns (
            bytes4 assetProxyId,
            uint256[] memory amounts,
            bytes[] memory nestedAssetData
        )
    {
        assetProxyId = assetData.readBytes4(0);

        require(
            assetProxyId == IAssetData(address(0)).MultiAsset.selector,
            "WRONG_PROXY_ID"
        );

        // solhint-disable indent
        (amounts, nestedAssetData) = abi.decode(
            assetData.slice(4, assetData.length),
            (uint256[], bytes[])
        );
        // solhint-enable indent
    }

    /// @dev Encode StaticCall asset data into the format described in the AssetProxy contract specification.
    /// @param staticCallTargetAddress Target address of StaticCall.
    /// @param staticCallData Data that will be passed to staticCallTargetAddress in the StaticCall.
    /// @param expectedReturnDataHash Expected Keccak-256 hash of the StaticCall return data.
    /// @return AssetProxy-compliant asset data describing the set of assets.
    function encodeStaticCallAssetData(
        address staticCallTargetAddress,
        bytes memory staticCallData,
        bytes32 expectedReturnDataHash
    )
        public
        pure
        returns (bytes memory assetData)
    {
        assetData = abi.encodeWithSelector(
            IAssetData(address(0)).StaticCall.selector,
            staticCallTargetAddress,
            staticCallData,
            expectedReturnDataHash
        );
        return assetData;
    }

    /// @dev Decode StaticCall asset data from the format described in the AssetProxy contract specification.
    /// @param assetData AssetProxy-compliant asset data describing a StaticCall asset
    /// @return The StaticCall AssetProxy identifier, the target address of the StaticCAll, the data to be
    /// passed to the target address, and the expected Keccak-256 hash of the static call return data.
    function decodeStaticCallAssetData(bytes memory assetData)
        public
        pure
        returns (
            bytes4 assetProxyId,
            address staticCallTargetAddress,
            bytes memory staticCallData,
            bytes32 expectedReturnDataHash
        )
    {
        assetProxyId = assetData.readBytes4(0);

        require(
            assetProxyId == IAssetData(address(0)).StaticCall.selector,
            "WRONG_PROXY_ID"
        );

        (staticCallTargetAddress, staticCallData, expectedReturnDataHash) = abi.decode(
            assetData.slice(4, assetData.length),
            (address, bytes, bytes32)
        );
    }

    /// @dev Decode ERC20Bridge asset data from the format described in the AssetProxy contract specification.
    /// @param assetData AssetProxy-compliant asset data describing an ERC20Bridge asset
    /// @return The ERC20BridgeProxy identifier, the address of the ERC20 token to transfer, the address
    /// of the bridge contract, and extra data to be passed to the bridge contract.
    function decodeERC20BridgeAssetData(bytes memory assetData)
        public
        pure
        returns (
            bytes4 assetProxyId,
            address tokenAddress,
            address bridgeAddress,
            bytes memory bridgeData
        )
    {
        assetProxyId = assetData.readBytes4(0);

        require(
            assetProxyId == IAssetData(address(0)).ERC20Bridge.selector,
            "WRONG_PROXY_ID"
        );

        (tokenAddress, bridgeAddress, bridgeData) = abi.decode(
            assetData.slice(4, assetData.length),
            (address, address, bytes)
        );
    }

    /// @dev Reverts if assetData is not of a valid format for its given proxy id.
    /// @param assetData AssetProxy compliant asset data.
    function revertIfInvalidAssetData(bytes memory assetData)
        public
        pure
    {
        bytes4 assetProxyId = assetData.readBytes4(0);

        if (assetProxyId == IAssetData(address(0)).ERC20Token.selector) {
            decodeERC20AssetData(assetData);
        } else if (assetProxyId == IAssetData(address(0)).ERC721Token.selector) {
            decodeERC721AssetData(assetData);
        } else if (assetProxyId == IAssetData(address(0)).ERC1155Assets.selector) {
            decodeERC1155AssetData(assetData);
        } else if (assetProxyId == IAssetData(address(0)).MultiAsset.selector) {
            decodeMultiAssetData(assetData);
        } else if (assetProxyId == IAssetData(address(0)).StaticCall.selector) {
            decodeStaticCallAssetData(assetData);
        } else if (assetProxyId == IAssetData(address(0)).ERC20Bridge.selector) {
            decodeERC20BridgeAssetData(assetData);
        } else {
            revert("WRONG_PROXY_ID");
        }
    }

    /// @dev Queries balance of an ERC20 token. Returns 0 if call was unsuccessful.
    /// @param tokenAddress Address of ERC20 token.
    /// @param ownerAddress Address of owner of ERC20 token.
    /// @return balance ERC20 token balance of owner.
    function _erc20BalanceOf(
        address tokenAddress,
        address ownerAddress
    )
        internal
        view
        returns (uint256 balance)
    {
        // Encode data for `balanceOf(ownerAddress)`
        bytes memory balanceOfData = abi.encodeWithSelector(
            IERC20Token(address(0)).balanceOf.selector,
            ownerAddress
        );

        // Query balance
        (bool success, bytes memory returnData) = tokenAddress.staticcall(balanceOfData);
        balance = success && returnData.length == 32 ? returnData.readUint256(0) : 0;
        return balance;
    }

    /// @dev Converts an amount of Chai into its equivalent Dai amount.
    ///      Also accumulates Dai from DSR if called after the last time it was collected.
    /// @param chaiAmount Amount of Chai to converts.
    function _convertChaiToDaiAmount(uint256 chaiAmount)
        internal
        returns (uint256 daiAmount)
    {
        PotLike pot = IChai(_getChaiAddress()).pot();
        // Accumulate savings if called after last time savings were collected
        uint256 chiMultiplier = (now > pot.rho())
            ? pot.drip()
            : pot.chi();
        daiAmount = LibMath.getPartialAmountFloor(chiMultiplier, 10**27, chaiAmount);
        return daiAmount;
    }
}
