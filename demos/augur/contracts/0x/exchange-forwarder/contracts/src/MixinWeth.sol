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


import "../../../utils/contracts/src/LibSafeMath.sol";
import "./libs/LibConstants.sol";
import "./libs/LibForwarderRichErrors.sol";


contract MixinWeth is
    LibConstants
{
    using LibSafeMath for uint256;

    /// @dev Default payable function, this allows us to withdraw WETH
    function ()
        external
        payable
    {
        if (msg.sender != address(ETHER_TOKEN)) {
            revert();
        }
    }

    /// @dev Transfers ETH denominated fees to all feeRecipient addresses
    /// @param ethFeeAmounts Amounts of ETH, denominated in Wei, that are paid to corresponding feeRecipients.
    /// @param feeRecipients Addresses that will receive ETH when orders are filled.
    /// @return ethRemaining msg.value minus the amount of ETH spent on affiliate fees.
    function _transferEthFeesAndWrapRemaining(
        uint256[] memory ethFeeAmounts,
        address payable[] memory feeRecipients
    )
        internal
        returns (uint256 ethRemaining)
    {
        uint256 feesLen = ethFeeAmounts.length;
        // ethFeeAmounts len must equal feeRecipients len
        if (feesLen != feeRecipients.length) {
            revert();
        }

        // This function is always called before any other function, so we assume that
        // the ETH remaining is the entire msg.value.
        ethRemaining = msg.value;

        for (uint256 i = 0; i != feesLen; i++) {
            uint256 ethFeeAmount = ethFeeAmounts[i];
            // Ensure there is enough ETH to pay the fee
            if (ethRemaining < ethFeeAmount) {
                revert();
            }
            // Decrease ethRemaining and transfer fee to corresponding feeRecipient
            ethRemaining = ethRemaining.safeSub(ethFeeAmount);
            feeRecipients[i].transfer(ethFeeAmount);
        }

        // Convert remaining ETH to WETH.
        ETHER_TOKEN.deposit.value(ethRemaining)();

        return ethRemaining;
    }

    /// @dev Refunds any excess ETH to msg.sender.
    /// @param initialWethAmount Amount of WETH available after transferring affiliate fees.
    /// @param wethSpent Amount of WETH spent when filling orders.
    function _transferEthRefund(
        uint256 initialWethAmount,
        uint256 wethSpent
    )
        internal
    {
        // Ensure that no extra WETH owned by this contract has been spent.
        if (wethSpent > initialWethAmount) {
            revert();
        }

        // Calculate amount of WETH that hasn't been spent.
        uint256 wethRemaining = initialWethAmount.safeSub(wethSpent);

        // Do nothing if no WETH remaining
        if (wethRemaining > 0) {
            // Convert remaining WETH to ETH
            ETHER_TOKEN.withdraw(wethRemaining);
            // Transfer remaining ETH to sender
            msg.sender.transfer(wethRemaining);
        }
    }
}
