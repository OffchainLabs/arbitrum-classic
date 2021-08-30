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

import "./L2ArbitrumGateway.sol";
import "../../libraries/IWETH9.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";

contract L2WethGateway is L2ArbitrumGateway {
    using SafeERC20 for IERC20;

    address public l1Weth;
    address public l2Weth;

    function initialize(
        address _l1Counterpart,
        address _router,
        address _l1Weth,
        address _l2Weth
    ) public {
        L2ArbitrumGateway._initialize(_l1Counterpart, _router);
        require(_l1Weth != address(0), "INVALID_L1WETH");
        require(_l2Weth != address(0), "INVALID_L2WETH");
        l1Weth = _l1Weth;
        l2Weth = _l2Weth;
    }

    /**
     * @notice internal utility function used to handle when no contract is deployed at expected address
     * @param l1ERC20 L1 address of ERC20
     * @param expectedL2Address L2 address of ERC20
     * @param deployData encoded symbol/name/decimal data for initial deploy
     */
    function handleNoContract(
        address l1ERC20,
        address expectedL2Address,
        address _from,
        address _to,
        uint256 _amount,
        bytes memory deployData
    ) internal override returns (bool shouldHalt) {
        // it is assumed that the custom token is deployed in the L2 before deposits are made
        // trigger withdrawal
        // this codepath should only be hit if the system is setup incorrectly
        // this withdrawal is for error recovery, not composing with L2 dapps, so we ignore the return value
        triggerWithdrawal(l1ERC20, address(this), _from, _amount, "");
        return true;
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

    function inboundEscrowTransfer(
        address _l2TokenAddress,
        address _dest,
        uint256 _amount
    ) internal override {
        IWETH9(_l2TokenAddress).deposit{ value: _amount }();
        IERC20(_l2TokenAddress).safeTransfer(_dest, _amount);
    }

    function createOutboundTx(
        address _from,
        uint256 _tokenAmount,
        bytes memory _outboundCalldata
    ) internal override returns (uint256) {
        return
            sendTxToL1(
                // we send the amount of weth withdrawn as callvalue to the L1 gateway
                _tokenAmount,
                _from,
                counterpartGateway,
                _outboundCalldata
            );
    }

    receive() external payable {}
}
