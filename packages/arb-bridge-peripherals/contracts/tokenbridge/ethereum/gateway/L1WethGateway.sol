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

import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";
import "../../libraries/IWETH9.sol";
import "../../test/TestWETH9.sol";
import "./L1ArbitrumExtendedGateway.sol";

contract L1WethGateway is L1ArbitrumExtendedGateway {
    using SafeERC20 for IWETH9;

    address public l1Weth;
    address public l2Weth;

    function initialize(
        address _l1Counterpart,
        address _l1Router,
        address _inbox,
        address _l1Weth,
        address _l2Weth
    ) public virtual {
        L1ArbitrumExtendedGateway._initialize(_l1Counterpart, _l1Router, _inbox);
        require(_l1Weth != address(0), "INVALID_L1WETH");
        require(_l2Weth != address(0), "INVALID_L2WETH");
        l1Weth = _l1Weth;
        l2Weth = _l2Weth;
    }

    function createOutboundTx(
        address _l1Token,
        address _from,
        address _to,
        uint256 _amount,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        uint256 _maxSubmissionCost,
        bytes memory _extraData
    ) internal virtual override returns (uint256) {
        return
            sendTxToL2(
                _from,
                _amount, // send token amount to L2 as call value
                _maxSubmissionCost,
                _maxGas,
                _gasPriceBid,
                getOutboundCalldata(_l1Token, _from, _to, _amount, _extraData)
            );
    }

    function outboundEscrowTransfer(
        address _l1Token,
        address _from,
        uint256 _amount
    ) internal virtual override {
        IWETH9(_l1Token).safeTransferFrom(_from, address(this), _amount);
        IWETH9(_l1Token).withdraw(_amount);
    }

    function inboundEscrowTransfer(
        address _l1Token,
        address _dest,
        uint256 _amount
    ) internal virtual override {
        IWETH9(_l1Token).deposit{ value: _amount }();
        IWETH9(_l1Token).safeTransfer(_dest, _amount);
    }

    /**
     * @notice Calculate the address used when bridging an ERC20 token
     * @dev this always returns the same as the L1 oracle, but may be out of date.
     * For example, a custom token may have been registered but not deploy or the contract self destructed.
     * @param l1ERC20 address of L1 token
     * @return L2 address of a bridged ERC20 token
     */
    function _calculateL2TokenAddress(address l1ERC20)
        internal
        view
        virtual
        override
        returns (address)
    {
        require(l1ERC20 == l1Weth, "WRONG_L1WETH");
        return l2Weth;
    }

    receive() external payable {}
}
