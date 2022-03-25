// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2020, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

pragma solidity ^0.6.11;

import "arb-bridge-eth/contracts/bridge/interfaces/IInbox.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "../../libraries/IWETH9.sol";
import "../../test/TestWETH9.sol";
import "./L1ArbitrumExtendedGateway.sol";

contract L1WethGateway is L1ArbitrumExtendedGateway {
    using SafeERC20 for IERC20;

    address public l1Weth;
    address public l2Weth;

    function initialize(
        address _l1Counterpart,
        address _l1Router,
        address _inbox,
        address _l1Weth,
        address _l2Weth
    ) public {
        L1ArbitrumExtendedGateway._initialize(_l1Counterpart, _l1Router, _inbox);
        require(_l1Weth != address(0), "INVALID_L1WETH");
        require(_l2Weth != address(0), "INVALID_L2WETH");
        l1Weth = _l1Weth;
        l2Weth = _l2Weth;
    }

    function createOutboundTxCustomRefund(
        address _refundTo,
        address _from,
        uint256 _tokenAmount,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        uint256 _maxSubmissionCost,
        bytes memory _outboundCalldata
    ) internal override returns (uint256) {
        return
            sendTxToL2(
                inbox,
                counterpartGateway,
                _from,
                // msg.value does not include weth withdrawn from user, we need to add in that amount
                msg.value + _tokenAmount,
                // send token amount to L2 as call value
                _tokenAmount,
                L2GasParams({
                    _maxSubmissionCost: _maxSubmissionCost,
                    _maxGas: _maxGas,
                    _gasPriceBid: _gasPriceBid
                }),
                _outboundCalldata
            );
    }

    function outboundEscrowTransfer(
        address _l1Token,
        address _from,
        uint256 _amount
    ) internal override returns (uint256) {
        IERC20(_l1Token).safeTransferFrom(_from, address(this), _amount);
        IWETH9(_l1Token).withdraw(_amount);
        // the weth token doesn't contain any special behaviour that changes the amount
        // when doing transfers / withdrawals. so we don't check the balanceOf
        return _amount;
    }

    function inboundEscrowTransfer(
        address _l1Token,
        address _dest,
        uint256 _amount
    ) internal override {
        IWETH9(_l1Token).deposit{ value: _amount }();
        IERC20(_l1Token).safeTransfer(_dest, _amount);
    }

    /**
     * @notice Calculate the address used when bridging an ERC20 token
     * @dev the L1 and L2 address oracles may not always be in sync.
     * For example, a custom token may have been registered but not deploy or the contract self destructed.
     * @param l1ERC20 address of L1 token
     * @return L2 address of a bridged ERC20 token
     */
    function calculateL2TokenAddress(address l1ERC20) public view override returns (address) {
        if (l1ERC20 != l1Weth) {
            // invalid L1 weth address
            return address(0);
        }
        return l2Weth;
    }

    receive() external payable {}
}
