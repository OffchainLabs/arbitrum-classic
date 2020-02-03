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

import "../interfaces/IERC20Bridge.sol";
import "../interfaces/IChai.sol";
import "../../../../utils/contracts/src/DeploymentConstants.sol";
import "../../../../erc20/contracts/src/interfaces/IERC20Token.sol";


// solhint-disable space-after-comma
contract ChaiBridge is
    IERC20Bridge,
    DeploymentConstants
{
    /// @dev Withdraws `amount` of `from` address's Dai from the Chai contract.
    ///      Transfers `amount` of Dai to `to` address.
    /// @param from Address to transfer asset from.
    /// @param to Address to transfer asset to.
    /// @param amount Amount of asset to transfer.
    /// @return success The magic bytes `0x37708e9b` if successful.
    function bridgeTransferFrom(
        address /* tokenAddress */,
        address from,
        address to,
        uint256 amount,
        bytes calldata /* bridgeData */
    )
        external
        returns (bytes4 success)
    {
        // Ensure that only the `ERC20BridgeProxy` can call this function.
        require(
            msg.sender == _getERC20BridgeProxyAddress(),
            "ChaiBridge/ONLY_CALLABLE_BY_ERC20_BRIDGE_PROXY"
        );

        // Withdraw `from` address's Dai.
        // NOTE: This contract must be approved to spend Chai on behalf of `from`.
        bytes memory drawCalldata = abi.encodeWithSelector(
            IChai(address(0)).draw.selector,
            from,
            amount
        );

        (bool success,) = _getChaiAddress().call(drawCalldata);
        require(
            success,
            "ChaiBridge/DRAW_DAI_FAILED"
        );

        // Transfer Dai to `to`
        // This will never fail if the `draw` call was successful
        IERC20Token(_getDaiAddress()).transfer(to, amount);

        return BRIDGE_SUCCESS;
    }
}
